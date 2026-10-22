package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 string; F2 int64 }
func genFn0(p0 int8) uint8 {
    v7 := []int{321, 4294967295, -9223372036854775808, -546}
    switch 1 {
    case 4294967295:
        fmt.Println(p0)
    case 2147483647:
        var v8 float32 = 1e10
        _ = v8
    default:
        v7 = append(v7, 42)
    }
    _ = p0
    _ = v7
    return 102
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]bool)
    var v3 int16 = 32767
    var v4 int64 = 341
    v5 := make(map[int]uint32)
    var v6 uint16 = 65535
    v9 := GenS1{F0: (1.5 != (1.5 * (1.5 * 3.14159))), F1: "hello"}
    v6 = (^(v6))
    var v10 uint32 = 562
    switch "" {
    case "foo bar":
        if ((int8(v3) == 42) && true) {
            var v11 int64 = v4
            var v12 int = 551
            var v13 int32 = -582
            _ = v11
            _ = v12
            _ = v13
        } else {
            for fi14 := 0; fi14 < 15; fi14++ {
                var v15 int8 = -128
                switch "foo bar" {
                case "":
                    v3 -= v3
                default:
                    var v16 uint8 = 42
                    _ = v16
                }
                _ = fi14
                _ = v15
            }
        }
    case "hello":
        var v17 uint16 = (^(v6))
        _ = v17
    }
    v18 := GenS2{F0: -128}
    for fi19 := 0; fi19 < 15; fi19++ {
        fmt.Println(v4, v10, v6)
        if ((0.1 + 3.14159) >= 0.0) {
            if false {
                v6 -= (v6 + v6)
            }
        }
        _ = fi19
    }
    fmt.Println(v10, v4, v4)
    v20 := v1["\n\t"]
    v4 = v4
    var v21 string = ""
    v3 = (v3 >> ((494) & 63))
    switch 0 {
    case -9223372036854775808:
        v21 = v21
    case -509:
        v22 := GenS1{F0: false, F1: "hello"}
        _ = v22
    }
    switch 128 {
    case -22:
        for fi23 := 0; fi23 < 1; fi23++ {
            for fi24 := 0; fi24 < 16; fi24++ {
                v25 := GenS3{F0: 0, F1: "foo bar", F2: -1}
                v20 = v20
                _ = fi24
                _ = v25
            }
            _ = fi23
        }
    case 1:
        if ("foo bar" <= v20) {
            var v26 int16 = (v3 ^ v3)
            var v27 uint32 = v10
            _ = v26
            _ = v27
        }
    default:
        switch 2147483647 {
        case -499:
            var v28 string = v21
            _ = v28
        case -9223372036854775808:
            var v29 int = 457
            _ = v29
        default:
            var v30 string = (v21 + (v20 + v20))
            _ = v30
        }
    }
    var v31 uint = 42
    var v32 uint8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
