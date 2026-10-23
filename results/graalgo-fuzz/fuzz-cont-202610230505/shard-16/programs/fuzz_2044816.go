package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint8 }
func genFn0(p0 uint8, p1 int8, p2 uint64) string {
    fmt.Println(p1, p0)
    p0 = (p0 + p0)
    p1 = (p1 + (p1 / ((p1) | 1)))
    _ = p0
    _ = p1
    _ = p2
    return ""
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    var v3 uint64 = 18446744073709551615
    v4 := GenS1{F0: 0}
    v5 := GenS1{F0: -451}
    v3 = (v3 << ((41) & 63))
    v3 += v3
    for fi6 := 0; fi6 < 1; fi6++ {
        switch -2147483648 {
        case 32767:
            if false {
                switch fi6 {
                case 845:
                    v3 -= (v3 * v3)
                case 578:
                    var v7 int16 = 253
                    _ = v7
                default:
                    var v8 float32 = (1.5 / (float32(0.0) - 3.14159))
                    _ = v8
                }
            } else {
                for fi9 := 0; fi9 < 1; fi9++ {
                    var v10 uint = 42
                    _ = fi9
                    _ = v10
                }
            }
        case 899:
            fmt.Println(v3, v3)
        }
        v11 := GenS1{F0: -1}
        if true {
            v3 -= v3
        } else {
            switch "\n\t" {
            case "":
                var v12 uint32 = 648
                _ = v12
            case "\n\t":
                switch (^(fi6)) {
                case -2147483648:
                    v13 := GenS1{F0: 0}
                    _ = v13
                case 32767:
                    for fi14 := 0; fi14 < 5; fi14++ {
                        var v15 string = ""
                        fmt.Println(v15, fi14)
                        _ = fi14
                        _ = v15
                    }
                case 32768:
                    v3 = (v3 + v3)
                }
            case "a":
                _ = len(v1)
            default:
                var v16 uint = 255
                _ = v16
            }
            v17 := GenS2{F0: 122, F1: 133}
            _ = v17
        }
        _ = fi6
        _ = v11
    }
    if true {
        if (!true) {
            switch "a" {
            case "foo bar":
                if false {
                    fmt.Println(v3)
                    v3 = v3
                }
            default:
                v3 -= 0
            }
        }
    }
    if ((!(-2147483648 >= 42)) || ((!(0 == 4294967295)) || (1 <= -1))) {
        var v18 uint = 1
        var v19 int16 = 0
        var v20 float32 = ((1.5 - (0.0 * (0.0 * 3.14159))) * 1e10)
        _ = v18
        _ = v19
        _ = v20
    }
    var v21 int = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
