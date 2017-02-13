package command

import (
    "fmt"
    "monitor/command/common"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "monitor/monitor/daemon"
    "monitor/monitor"
)

var MainCmd = &common.Command{
    Subject: &monitor.Monitor{},
    Viper: viper.GetViper(),
    
}

func UsageTemplate() string {
    return `Usage:{{if .Runnable}}
  {{if .HasAvailableFlags}}{{appendIfNotPresent .UseLine "[flags]"}}{{else}}{{.UseLine}}{{end}}{{end}}{{if .HasAvailableSubCommands}}
  {{ .CommandPath}} [command]{{end}}{{if gt .Aliases 0}}

Aliases:
  {{.NameAndAliases}}
{{end}}{{if .HasExample}}

Examples:
{{ .Example }}{{end}}{{ if .HasAvailableSubCommands}}

Commands:{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{ if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimRightSpace}}{{end}}{{ if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimRightSpace}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsHelpCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{ if .HasAvailableSubCommands }}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`
}

func init() {
    
    MainCmd.Command = &cobra.Command{
        Use: common.CMD_MAIN,
        Short: "Linux server status monitor",
        Long: "Powerful Linux server status monitor server",
        PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
            if MainCmd.ReadConf() != nil {
                fmt.Println("No config file found. Using built-in defaults.");
            }
            return nil
        },
        RunE: func(cmd *cobra.Command, args []string) error {
            Conf := MainCmd.Configure.Server
            
            Daemon := &daemon.Daemon{
                PidFile: Conf.Pid,
                UnixFile: Conf.Unix,
                LogFile: Conf.Log,
            }
            
            if Conf.Daemon == true {
                Daemon.Daemon(MainCmd.Scheduler)
                return nil
            }
            
            err := Daemon.CreatePidFile()
            if err == nil {
                err := Daemon.WritePidFile()
                if err == nil {
                    go Daemon.Signal()
                    Daemon.UnixListen(MainCmd.Scheduler)
                    return nil
                }
            }
            
            fmt.Println(err)
            
            return nil
        },
    }
    
    Flags := MainCmd.Command.Flags()
    Flags.BoolVarP(&MainCmd.Flags.Main.Daemon, "daemon", "d", false, "to start by daemon")
    Flags.StringVarP(&MainCmd.Flags.Main.Pid, "pid", "", "", "full path of pid file")
    Flags.StringVarP(&MainCmd.Flags.Main.Log, "log", "", "", "full path of log file")
    Flags.StringVarP(&MainCmd.Flags.Main.Config, "config", "c", "/etc/monitor.toml", "configuration file specifying additional options")
    
    V := MainCmd.Viper
    V.BindPFlag("server.pid", Flags.Lookup("pid"))
    V.BindPFlag("server.log", Flags.Lookup("log"))
    V.BindPFlag("server.daemon", Flags.Lookup("daemon"))
    
    //Flags = JoinCmd.Command.Flags()
    //Flags.StringVarP(&JoinCmd.Flags.Join.Addr, "addr", "", "", "manager addr")
    //Flags.StringVarP(&JoinCmd.Flags.Join.Token, "token", "", "", "join token")
    //
    //Flags = InitCmd.Command.Flags()
    //Flags.StringVarP(&MainCmd.Configure.MongoDB.Host, "host", "", "127.0.0.1", "mongodb host")
    //Flags.Int32VarP(&MainCmd.Configure.MongoDB.Port, "port", "", 27017, "mongodb port")
    //Flags.StringVarP(&MainCmd.Configure.MongoDB.Auth, "auth", "", "admin", "auth database")
    //Flags.StringVarP(&MainCmd.Configure.MongoDB.Username, "user", "", "", "username")
    //Flags.StringVarP(&MainCmd.Configure.MongoDB.Password, "pwd", "", "", "password")
    //
    //Flags.StringVarP(&MainCmd.Configure.Server.Addr, "addr", "a", "0.0.0.0:3647", "web server address")
    MainCmd.Command.SetUsageTemplate(UsageTemplate())
    MainCmd.NewChildren(JoinCmd, RoleCmd, ServerCmd, VersionCmd)
}