package main

import (
    //"github.com/shirou/gopsutil/load"
    //"github.com/gizak/termui"
    "fmt"
    //"runtime"
    //"time"
    //"github.com/shirou/gopsutil/cpu"
    //"github.com/shirou/gopsutil/mem"
    //"github.com/shirou/gopsutil/disk"
    //"github.com/shirou/gopsutil/docker"
    "monitor/slave"
)

func main() {

    fmt.Println(slave.GetSysInfo().ToString())
    fmt.Println("--------")


    //avgStat, _ := load.Avg()
    //miscStat, _ := load.Misc()
    //c, _ := cpu.Times(true)
    ////c ,_ := cpu.Percent(time.Second, false)
    ////fmt.Println(avgStat.String(), miscStat, cpuInfo)
    //fmt.Println(c)
    //
    //m, _ := mem.SwapMemory()
    //v, _ := mem.VirtualMemory()
    //fmt.Println(m, v)
    //
    //
    //d, _ := disk.Usage("/")
    //fmt.Println(d)
    //
    //do, _ := docker.GetDockerIDList()
    //fmt.Println(do)
    //err := termui.Init()
    //if err != nil {
    //    panic(err)
    //}
    //defer termui.Close()
    //
    ////termui.UseTheme("helloworld")
    //
    //data := []int{4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6}
    //spl0 := termui.NewSparkline()
    //spl0.Data = data[3:]
    //spl0.Title = "Sparkline 0"
    //spl0.LineColor = termui.ColorGreen
    //
    //// single
    //spls0 := termui.NewSparklines(spl0)
    //spls0.Height = 2
    //spls0.Width = 20
    //spls0.Border = false
    //
    //spl1 := termui.NewSparkline()
    //spl1.Data = data
    //spl1.Title = "Sparkline 1"
    //spl1.LineColor = termui.ColorRed
    //
    //spl2 := termui.NewSparkline()
    //spl2.Data = data[5:]
    //spl2.Title = "Sparkline 2"
    //spl2.LineColor = termui.ColorMagenta
    //
    //// group
    //spls1 := termui.NewSparklines(spl0, spl1, spl2)
    //spls1.Height = 8
    //spls1.Width = 20
    //spls1.Y = 3
    //spls1.BorderLabel = "Group Sparklines"
    //
    //spl3 := termui.NewSparkline()
    //spl3.Data = data
    //spl3.Title = "Enlarged Sparkline"
    //spl3.Height = 8
    //spl3.LineColor = termui.ColorYellow
    //
    //spls2 := termui.NewSparklines(spl3)
    //spls2.Height = 11
    //spls2.Width = 30
    //spls2.BorderFg = termui.ColorCyan
    //spls2.X = 21
    //spls2.BorderLabel = "Tweeked Sparkline"
    //
    //termui.Render(spls0, spls1, spls2)
    //
    //termui.Handle("/sys/kbd/q", func(termui.Event) {
    //    termui.StopLoop()
    //})
    //termui.Loop()
    //
    //
    //fmt.
}