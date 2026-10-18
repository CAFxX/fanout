package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint8; F1 int }

func main() {
    var v1 int8 = 100
    var v2 string = ("a" + "a")
    var v3 int64 = 1
    var v4 string = "hello"
    v1 = 42
    v1 = (v1 / (((v1 - (v1 | v1))) | 1))
    v2 = (v2 + v4)
    v4 = v4
    if true {
        fmt.Println(v4, v1, v3)
    } else {
        for fi5 := 0; fi5 < 7; fi5++ {
            var v6 uint64 = 534
            fmt.Println(fi5, v6)
            _ = fi5
            _ = v6
        }
        v4 = "foo bar"
        var v7 string = "\n\t"
        _ = v7
    }
    var v8 uint = 65535
    v9 := []int{}
    v8 = 4294967295
    switch ((v2 + "a") + (v2 + ((("a" + (v4 + ("\n\t" + ((v2 + v2) + "foo bar")))) + "\n\t") + ("" + (v4 + v4))))) {
    case "\n\t":
        v9 = append(v9, 127)
    case "":
        switch 127 {
        case 739:
            for fi10 := 0; fi10 < 20; fi10++ {
                v11 := GenS1{F0: 42, F1: (v3 & int64(v8))}
                _ = fi10
                _ = v11
            }
        case 0:
            switch v2 {
            case "foo bar":
                for fi12 := 0; fi12 < 8; fi12++ {
                    for fi13 := 0; fi13 < 16; fi13++ {
                        _ = len(v9)
                        v4 = v2
                        _ = fi13
                    }
                    _ = fi12
                }
            case "a":
                fmt.Println(v4, v3)
            case "\n\t":
                v4 = "foo bar"
            }
        case -473:
            v14 := GenS2{F0: 234, F1: 796}
            _ = v14
        }
    default:
        v8 = v8
    }
    var v15 uint16 = 919
    v4 = v4
    _ = len(v9)
    fmt.Println(v2, v4)
    var v16 uint = v8
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    if true {
        var v17 uint64 = 42
        for fi18 := 0; fi18 < 8; fi18++ {
            v17 = v17
            _ = fi18
        }
        _ = v17
    } else {
        v16 -= v8
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    }
    v19 := GenS2{F0: 177, F1: 42}
    for fi20 := 0; fi20 < 12; fi20++ {
        var v21 int16 = 1
        if false {
            var v22 bool = (!(-1 < v1))
            _ = v22
        }
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
