package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 int16; F2 int64 }
func genFn0(p0 uint32, p1 uint32) int32 {
    p0 += (^(p0))
    p1 -= (p1 - p0)
    v5 := []int{2147483647, 4294967295, 32767}
    _ = p0
    _ = p1
    _ = v5
    return 42
}
func genFn1(p0 uint16, p1 int32, p2 uint32) int {
    fmt.Println(p0, p1, p1)
    var v6 float64 = ((1.5 + -2.25) + float64(0.0))
    p0 += p0
    fmt.Println(p1, p0, v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 4294967295
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = true
    var v3 int16 = 42
    v4 := make(map[uint32]bool)
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 int = (fi7 - genFn1(255, 643, 4294967295))
        v3 += -97
        for fi9 := 0; fi9 < 18; fi9++ {
            if v2 {
                v2 = v2
            }
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        if false {
            v2 = v2
        } else {
            v3 += (v3 + (v3 ^ (v3 / ((v3) | 1))))
            v11 := GenS1{F0: 834, F1: true}
            fmt.Println(v2, v2, v2)
            _ = v11
        }
        if (220 == 742) {
            v2 = true
        }
        var v12 uint64 = 4294967295
        _ = fi10
        _ = v12
    }
    var v13 int8 = 42
    v14 := GenS2{F0: ((("foo bar" + ("\n\t" + (("\n\t" + "hello") + "hello"))) + ((((("" + "foo bar") + ("" + "foo bar")) + "") + "a") + "hello")) + "a")}
    v15 := v1[2147483647]
    v3 = -1
    v16 := GenS3{F0: 473, F1: (v3 | (v3 ^ (v3 / (((v3 / ((v3) | 1))) | 1)))), F2: -545}
    switch "" {
    case "\n\t":
        var v17 uint32 = 1
        _ = v17
    case "foo bar":
        v18 := []int{-541, 32767, genFn1(0, -1, 2147483648)}
        _ = v18
    case "":
        var v19 uint = 44
        _ = v19
    default:
        var v20 uint = 4294967295
        _ = v20
    }
    v13 = v13
    var v21 float64 = (1e10 + -2.25)
    var v22 uint = 42
    v15 = v15
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
