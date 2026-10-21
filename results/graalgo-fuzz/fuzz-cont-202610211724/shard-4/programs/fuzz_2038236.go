package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint }

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 0
    v3 := make(map[string]bool)
    var v4 float32 = 3.14159
    v5 := []int{}
    v6 := []int{9223372036854775807, 207, 2147483647, 128}
    var v7 int16 = 651
    switch ("" + "\n\t") {
    case "hello":
        var v8 int64 = 9223372036854775807
        _ = v8
    case "\n\t":
        switch 2147483647 {
        case 42:
            v7 -= v7
        case 83:
            if len(v5) > 0 {
                v5[0] = 2147483647
            }
        default:
            v7 -= v7
        }
    }
    if true {
        v7 -= (v7 / (((^(v7))) | 1))
    } else {
        if false {
            for fi9 := 0; fi9 < 15; fi9++ {
                var v10 int = (fi9 >> ((0) & 63))
                _ = fi9
                _ = v10
            }
            fmt.Println(v7, v4, v2)
            v2 += v2
        } else {
            for fi11 := 0; fi11 < 16; fi11++ {
                for fi12 := 0; fi12 < 6; fi12++ {
                    fmt.Println(v7, v4, v2)
                    _ = fi12
                }
                if false {
                    _ = len(v3)
                }
                var v13 float32 = float32(((0.1 - (-2.25 / 3.14159)) - 1.5))
                _ = fi11
                _ = v13
            }
            if (!(!true)) {
                v14 := GenS1{F0: 1, F1: true, F2: 127}
                v7 += 898
                v4 = v4
                _ = v14
            }
        }
    }
    var v15 bool = (1 > (^(v2)))
    switch "hello" {
    case "":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case "hello":
        for fi16 := 0; fi16 < 15; fi16++ {
            fmt.Println(v7, v15)
            switch ("" + "\n\t") {
            case "a":
                fmt.Println(v7, v15, fi16)
            case "hello":
                v2 += (v2 >> ((42) & 63))
            case "\n\t":
                if (v7 == v7) {
                    v5 = append(v5, (fi16 + fi16))
                } else {
                    var v17 float32 = v4
                    v15 = v15
                    _ = v17
                }
            default:
                v2 += v2
            }
            _ = fi16
        }
    case "\n\t":
        var v18 int64 = 585
        _ = v18
    default:
        var v19 uint8 = 2
        _ = v19
    }
    var v20 uint16 = v2
    var v21 uint64 = 42
    var v22 int64 = 9223372036854775807
    v1[32767] = 0
    if false {
        fmt.Println(v22)
    } else {
        var v23 int64 = (-(v22))
        v20 += 32767
        _ = v23
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        v25 := GenS2{F0: v2, F1: 605}
        var v26 uint32 = 1
        _ = fi24
        _ = v25
        _ = v26
    }
    v15 = v15
    v20 -= 65535
    v22 = v22
    v2 = (v20 & 440)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
