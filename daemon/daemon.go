package daemon

import (
    "os"
    "runtime"
    "syscall"
    "errors"
)

type Daemon struct {
    LogFile string
}

func (d *Daemon) Start(ChDir, Close int) (int, error) {

    darwin := runtime.GOOS == "darwin"

    // 已经以daemon启动
    if syscall.Getppid() == 1 {
        return 0, nil
    }

    // fork子进程
    r1, r2, err := syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)
    if err != 0 {
        return -1, errors.New("fork fail")
    }
    if r2 < 0 {
        os.Exit(-1)
        return -1, errors.New("fork fail")
    }

    // 处理darwin的异常
    if darwin && r2 == 1 {
        r1 = 0
    }

    // 子进程成功启动，然后退出父进程
    if r1 > 0 {
        os.Exit(0)
    }

    syscall.Umask(0)

    if pid, err := syscall.Setsid(); err != nil || pid < 0 {
        if err != nil {
            return -1, err
        }
        if pid < 0 {
            return -1, errors.New("setsid fail")
        }
    }

    if ChDir > 0 {
        os.Chdir("/")
    }

    // 判断是否输出日志
    if Close > 0 || len(d.LogFile) == 0 {
        File, err := os.OpenFile("/dev/null", os.O_RDWR, 0)
        if err == nil {
            fd := File.Fd()
            syscall.Dup2(int(fd), int(os.Stdin.Fd()))
            syscall.Dup2(int(fd), int(os.Stdout.Fd()))
            syscall.Dup2(int(fd), int(os.Stderr.Fd()))
        }
    } else {
        File, err := os.OpenFile(d.LogFile, os.O_RDWR, 0)
        if err == nil {
            fd := File.Fd()
            syscall.Dup2(int(fd), int(os.Stdin.Fd()))
            syscall.Dup2(int(fd), int(os.Stdout.Fd()))
            syscall.Dup2(int(fd), int(os.Stderr.Fd()))
        }
    }

    return 0, nil
}