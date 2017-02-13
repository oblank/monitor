package server

import (
    "log"
    "errors"
    "io/ioutil"
    "encoding/json"
    "monitor/monitor/header"
    "monitor/monitor/helper"
    "monitor/monitor/collector"
    "github.com/noaway/heartbeat"
    "monitor/monitor/collector/model"
)

type Node header.Node

func (n *Node) Verify() error {
    R, err := helper.Request(header.METHOD, header.SCHEMA + n.Addr + "/verify", n.Token)
    if err != nil {
        return err
    }
    defer R.Body.Close()
    Body, _ := ioutil.ReadAll(R.Body)
    var Answer header.Answer
    json.Unmarshal(Body, &Answer)
    if Answer.Code == header.SUCCESS {
        return nil
    } else {
        return errors.New("verify token failure")
    }
}

func (n *Node) gather(Spec int) error {
    Ht, err := heartbeat.NewTast("gather", Spec)
    if err != nil {
        return err
    }
    Ht.Start(func() error {
        Gather := model.Gather{}
        
        R, err := helper.Request(header.METHOD, header.SCHEMA + n.Addr + "/gather", string(Gather.Exec()))
        if err != nil {
            log.Printf("%v", err)
        }
        defer R.Body.Close()
        Body, _ := ioutil.ReadAll(R.Body)
        var Answer header.Answer
        json.Unmarshal(Body, &Answer)
        if Answer.Code == header.FAILURE {
            log.Println(Answer)
        }
        return nil
    })
    return nil
}

func (n *Node) RunForever() error {
    err := n.Verify()
    if err != nil {
        return err
    }
    
    // 收集信息
    go n.gather(5)
    
    return nil
}