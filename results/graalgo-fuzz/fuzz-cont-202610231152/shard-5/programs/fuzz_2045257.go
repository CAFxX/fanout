package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int32 = -101
    v2 := make(map[int]uint32)
    v3 := []uint32{776, 2, 284, 679}
    v4 := make(map[uint32]string)
    var v5 float64 = (1.5 - (1.5 - 0.1))
    v6 := GenS1{F0: 232, F1: 138}
    if len(v3) > 0 {
        v3[0] = 1
    }
    if len(v3) > 0 {
        v3[0] = 65535
    }
    switch 391 {
    case 32768:
        if true {
            _ = len(v3)
            var v7 uint8 = 9
            v8 := GenS1{F0: v7, F1: 0}
            _ = v7
            _ = v8
        } else {
            var v9 int8 = 100
            v10 := GenS1{F0: 128, F1: 540}
            switch -928 {
            case 127:
                _ = len(v4)
            case 4294967295:
                var v11 uint16 = 1
                _ = v11
            }
            _ = v10
            _ = v9
        }
    default:
        var v12 bool = false
        _ = v12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        fmt.Println(fi13, v1)
        _ = fi13
    }
    if (true || true) {
        for fi14 := 0; fi14 < 11; fi14++ {
            v5 = v5
            _ = fi14
        }
    }
    if (false || false) {
        v1 = 42
    }
    if false {
        var v15 float32 = -2.25
        switch 128 {
        case 0:
            v16 := GenS2{F0: 0}
            _ = v16
        default:
            fmt.Println(v5, v5, v5)
        }
        if (255 == 116) {
            var v17 bool = false
            for fi18 := 0; fi18 < 15; fi18++ {
                v15 = (-2.25 * (float32(1.5) / -2.25))
                v15 += v15
                _ = fi18
            }
            _ = v17
        }
        _ = v15
    }
    var v19 bool = false
    for fi20 := 0; fi20 < 13; fi20++ {
        var v21 uint8 = 17
        _ = fi20
        _ = v21
    }
    fmt.Println(v1, v1, v5)
    for fi22 := 0; fi22 < 11; fi22++ {
        var v23 bool = ((("a" + "\n\t") + "hello") >= ("" + ("foo bar" + "")))
        _ = fi22
        _ = v23
    }
    v3 = append(v3, 2)
    var v24 int64 = 42
    var v25 int8 = -114
    switch "foo bar" {
    case "hello":
        var v26 uint32 = 0
        _ = v26
    case "foo bar":
        if (((v24 * (v24 * v24)) < 0) && false) {
            var v27 bool = (42 != 2)
            v28 := GenS2{F0: 2147483648}
            v29 := GenS1{F0: 255, F1: -1}
            _ = v27
            _ = v28
            _ = v29
        }
    }
    v25 = (v25 % ((0) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
