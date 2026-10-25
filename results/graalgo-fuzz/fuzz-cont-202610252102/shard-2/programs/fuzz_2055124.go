package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 bool }
func genFn0(p0 bool, p1 int64, p2 int8) uint64 {
    switch "hello" {
    case "foo bar":
        for fi6 := 0; fi6 < 14; fi6++ {
            if p0 {
                p0 = ((3.14159 == (3.14159 * 0.0)) && p0)
                fmt.Println(p1, p1)
            }
            v7 := GenS1{F0: 0, F1: 32767}
            for fi8 := 0; fi8 < 15; fi8++ {
                switch fi6 {
                case 9223372036854775807:
                    v9 := GenS1{F0: 1, F1: 753}
                    _ = v9
                case -682:
                    if (128 >= 255) {
                        var v10 uint64 = 672
                        _ = v10
                    } else {
                        v11 := GenS1{F0: 428, F1: 1}
                        p1 = p1
                        _ = v11
                    }
                case 32768:
                    v12 := GenS2{F0: false}
                    _ = v12
                default:
                    var v13 int64 = p1
                    _ = v13
                }
                fmt.Println(p2)
                _ = fi8
            }
            _ = fi6
            _ = v7
        }
    case "hello":
        for fi14 := 0; fi14 < 8; fi14++ {
            for fi15 := 0; fi15 < 9; fi15++ {
                p1 = (p1 ^ 0)
                _ = fi15
            }
            var v16 float32 = (-2.25 + 3.14159)
            _ = fi14
            _ = v16
        }
    default:
        var v17 int = -9223372036854775808
        _ = v17
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        if (("\n\t" + "") < "a") {
            if p0 {
                var v19 bool = (p2 > 127)
                _ = v19
            }
            v20 := []int{fi18, (fi18 % ((-2147483648) | 1)), -667}
            _ = v20
        }
        p2 = p2
        var v21 int32 = 741
        _ = fi18
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    return 18446744073709551615
}
func genFn1(p0 int32) uint8 {
    fmt.Println(p0)
    v22 := GenS1{F0: genFn0(true, 1, -128), F1: 42}
    switch "hello" {
    case "\n\t":
        switch 4294967295 {
        case 2147483647:
            p0 = (-(p0))
        case -680:
            var v23 int = 128
            _ = v23
        case -1:
            p0 = (^(p0))
        }
    default:
        v24 := []int{435}
        _ = v24
    }
    _ = p0
    _ = v22
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]bool)
    var v3 int32 = -2147483648
    v4 := []string{"foo bar"}
    var v5 int = 9223372036854775807
    v5 = v5
    v4 = append(v4, "foo bar")
    v3 = -2147483648
    var v25 float32 = -2.25
    v5 = 461
    v26 := GenS1{F0: 97, F1: 1000}
    fmt.Println(v5)
    _ = len(v2)
    var v27 int32 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
