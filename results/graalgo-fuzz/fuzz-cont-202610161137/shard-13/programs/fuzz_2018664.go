package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 uint; F1 uint8 }
func genFn0(p0 uint64, p1 uint16, p2 int) int16 {
    fmt.Println(p1)
    p1 += p1
    var v7 int16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -695
}
func genFn1() int16 {
    v8 := GenS1{F0: 65535}
    var v9 float32 = 1e10
    var v10 uint = 42
    fmt.Println(v9, v10)
    v11 := []int{9223372036854775807, 0}
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return genFn0(1, 32767, -9223372036854775808)
}

func main() {
    var v1 int16 = 665
    v2 := make(map[uint32]int)
    v3 := make(map[int]uint32)
    v4 := make(map[int]string)
    v5 := make(map[uint32]string)
    v6 := make(map[string]string)
    v12 := []int{-1, -859, -1, 9223372036854775807}
    v1 = v1
    v13 := GenS2{F0: v1, F1: 128}
    v1 = (v1 & -1)
    for fi14 := 0; fi14 < 17; fi14++ {
        switch fi14 {
        case -752:
            v15 := GenS2{F0: (^(v1)), F1: 0}
            _ = v15
        case -1:
            var v16 uint16 = 2
            _ = v16
        case 32768:
            if ("" < (("\n\t" + "hello") + ("foo bar" + "foo bar"))) {
                v17 := v3[209]
                if true {
                    switch "" {
                    case "":
                        var v18 float64 = (-2.25 - 3.14159)
                        _ = v18
                    case "\n\t":
                        v1 -= (v1 ^ (v1 % (((^(v1))) | 1)))
                    }
                }
                _ = v17
            }
        default:
            v19 := GenS1{F0: 2}
            _ = v19
        }
        v1 = v1
        v20 := GenS3{F0: "hello", F1: 2}
        _ = fi14
        _ = v20
    }
    if (260 == 743) {
        var v21 uint16 = 255
        var v22 float32 = ((1e10 + 1.5) * 3.14159)
        if len(v12) > 0 {
            v12[0] = -2147483648
        }
        _ = v21
        _ = v22
    }
    fmt.Println(v1)
    v23 := GenS3{F0: "\n\t", F1: 127}
    delete(v6, (("hello" + "foo bar") + "hello"))
    if false {
        v1 = genFn0(697, 0, -2147483648)
    }
    switch 42 {
    case 4294967295:
        fmt.Println(v1)
    default:
        var v24 uint8 = 42
        _ = v24
    }
    v25 := GenS4{F0: 883, F1: 1}
    fmt.Println(v1)
    v26 := GenS5{F0: 676, F1: 159}
    v1 = (v1 + -431)
    var v27 int = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
