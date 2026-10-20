package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int16; F1 uint32 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint8 = 128
    v3 := make(map[uint32]int)
    var v4 float64 = ((0.0 * (1e10 * 0.0)) - 1.5)
    v5 := make(map[string]int)
    var v6 int16 = -1
    v7 := GenS1{F0: 42, F1: (1e10 != 0.0)}
    var v8 uint16 = 65535
    v4 = v4
    v9 := []int{4294967295, 9223372036854775807, 746}
    for fi10 := 0; fi10 < 13; fi10++ {
        switch (^(fi10)) {
        case 32768:
            v9 = append(v9, 32768)
        case -672:
            var v11 int16 = 1
            _ = v11
        default:
            switch "foo bar" {
            case "":
                fmt.Println(v1, fi10, v4)
            case "hello":
                _ = len(v5)
            }
        }
        var v12 string = "\n\t"
        _ = fi10
        _ = v12
    }
    if (true || ((v8 ^ 573) <= 65535)) {
        v2 -= 255
        v2 = (v2 + uint8(v8))
        if true {
            switch -319 {
            case -420:
                var v13 bool = (int16(v8) != -1)
                _ = v13
            case 32767:
                v9 = append(v9, 2147483647)
            case -1:
                v4 = ((v4 - float64(((0.1 + 0.1) - 1.5))) / v4)
            default:
                var v14 int32 = 1
                _ = v14
            }
            if false {
                _ = len(v3)
            } else {
                _ = len(v3)
                for fi15 := 0; fi15 < 6; fi15++ {
                    for fi16 := 0; fi16 < 20; fi16++ {
                        v17 := GenS1{F0: 829, F1: false}
                        v18 := GenS1{F0: 1, F1: ("" == ("\n\t" + "\n\t"))}
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    _ = fi15
                }
            }
        }
    }
    v19 := GenS1{F0: int32(v8), F1: ((v2 >> (((v1 | 4294967295)) & 63)) <= 50)}
    if true {
        v1 -= (v1 + (^(v1)))
        fmt.Println(v4, v8)
        v2 += (^(v2))
    } else {
        _ = len(v5)
        v20 := GenS1{F0: int32(v2), F1: ((v1 / ((v1) | 1)) <= (v1 << ((uint(v2)) & 63)))}
        _ = v20
    }
    var v21 int64 = -1
    if false {
        var v22 float32 = (-2.25 * float32((float64(1e10) / -2.25)))
        _ = v22
    }
    var v23 int64 = (v21 * (v21 + (v21 + 537)))
    if (float64(float32(0.0)) > 1e10) {
        if ("a" == (("a" + "foo bar") + "")) {
            var v24 float32 = (3.14159 + ((3.14159 + float32(1.5)) / 3.14159))
            _ = v24
        }
        for fi25 := 0; fi25 < 8; fi25++ {
            fmt.Println(v6, v8, v4)
            _ = fi25
        }
    } else {
        switch -2147483648 {
        case -9223372036854775808:
            var v26 bool = (v2 == v2)
            _ = v26
        case 42:
            fmt.Println(v23)
        case 639:
            v4 = 1.5
        default:
            fmt.Println(v8, v4)
        }
        for fi27 := 0; fi27 < 3; fi27++ {
            v28 := GenS2{F0: v6, F1: 42}
            _ = fi27
            _ = v28
        }
        for fi29 := 0; fi29 < 9; fi29++ {
            v30 := GenS2{F0: (^(v6)), F1: 2}
            v8 -= v8
            _ = fi29
            _ = v30
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
