package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint16 }

func main() {
    v1 := []uint32{}
    v2 := make(map[int]uint32)
    var v3 bool = (65535 >= 42)
    for fi4 := 0; fi4 < 12; fi4++ {
        switch "foo bar" {
        case "\n\t":
            var v5 int = fi4
            _ = v5
        default:
            v3 = (!false)
        }
        fmt.Println(fi4, v3)
        var v6 int32 = 0
        _ = fi4
        _ = v6
    }
    v7 := GenS1{F0: (42 <= 42), F1: 260, F2: 255}
    v1 = append(v1, 895)
    var v8 uint32 = 65535
    v3 = true
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v8 += 748
    if ((!((0.1 + (3.14159 + 3.14159)) > 0.0)) && (!((v8 << ((255) & 63)) > v8))) {
        var v9 float64 = 1e10
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := GenS1{F0: true, F1: 32767, F2: 514}
        var v12 float32 = ((0.0 / (3.14159 - (float32((-2.25 - 0.1)) * 1.5))) / 1e10)
        fmt.Println(fi10)
        _ = fi10
        _ = v11
        _ = v12
    }
    if ((0 <= 457) && (65535 == 240)) {
        for fi13 := 0; fi13 < 20; fi13++ {
            if v3 {
                v8 = 523
                fmt.Println(fi13, fi13)
                _ = len(v2)
            }
            _ = fi13
        }
    }
    var v14 float32 = -2.25
    switch "hello" {
    case "foo bar":
        var v15 int16 = 32767
        _ = v15
    case "a":
        for fi16 := 0; fi16 < 12; fi16++ {
            fmt.Println(v8, v3)
            v8 = (^(v8))
            _ = fi16
        }
    case "":
        v3 = v3
    default:
        v17 := GenS1{F0: false, F1: -32768, F2: 260}
        _ = v17
    }
    fmt.Println(v14, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
