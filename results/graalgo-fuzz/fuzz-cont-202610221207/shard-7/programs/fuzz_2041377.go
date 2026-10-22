package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 bool }

func main() {
    v1 := []int{4294967295, -872, -104}
    v2 := []string{("foo bar" + "\n\t"), "", "\n\t", (("foo bar" + "\n\t") + (("" + "a") + "\n\t"))}
    v3 := make(map[string]bool)
    v4 := make(map[uint32]int)
    v5 := []string{}
    v6 := []uint32{425, 0, 1}
    var v7 uint8 = 32
    v7 += v7
    switch -9223372036854775808 {
    case 127:
        var v8 float32 = -2.25
        _ = v8
    default:
        v7 = 0
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi9
    }
    fmt.Println(v7)
    fmt.Println(v7)
    v10 := GenS1{F0: -950}
    switch 0 {
    case -150:
        v11 := GenS1{F0: -1}
        _ = v11
    case 525:
        switch "foo bar" {
        case "":
            if false {
                v7 -= 1
                v12 := v4[675]
                _ = v12
            } else {
                fmt.Println(v7)
                v3["foo bar"] = (-516 <= 180)
                v7 = v7
            }
        default:
            var v13 uint32 = 231
            _ = v13
        }
    }
    var v14 int16 = 1
    var v15 int16 = 1000
    var v16 uint = 1
    var v17 int8 = 111
    for fi18 := 0; fi18 < 7; fi18++ {
        fmt.Println(v17, v17, v17)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi18
    }
    fmt.Println(v17, v7)
    v19 := GenS2{F0: 1, F1: (1 <= v15)}
    for fi20 := 0; fi20 < 13; fi20++ {
        if true {
            v16 = 42
        }
        _ = fi20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
