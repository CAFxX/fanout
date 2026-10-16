package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 bool }
type GenS2 struct { F0 int32; F1 int16; F2 int16 }
func genFn0(p0 uint32) uint16 {
    for fi7 := 0; fi7 < 17; fi7++ {
        p0 = p0
        v8 := []int{-1, (fi7 & fi7)}
        p0 -= p0
        _ = fi7
        _ = v8
    }
    if (0 < 1) {
        if false {
            var v9 int32 = 2147483647
            switch ("\n\t" + "\n\t") {
            case "a":
                v10 := []int{}
                _ = v10
            case "\n\t":
                var v11 float32 = 0.1
                _ = v11
            case "hello":
                fmt.Println(v9, p0)
            }
            _ = v9
        } else {
            for fi12 := 0; fi12 < 10; fi12++ {
                p0 = 2147483647
                v13 := []int{128, (fi12 >> ((2) & 63)), fi12, fi12}
                p0 = 2147483648
                _ = fi12
                _ = v13
            }
            p0 = p0
        }
    } else {
        var v14 int16 = 0
        _ = v14
    }
    v15 := GenS1{F0: (p0 * (^(p0))), F1: false, F2: (42 == -1)}
    v16 := GenS1{F0: (p0 ^ p0), F1: false, F2: (!(!false))}
    p0 += p0
    _ = p0
    _ = v15
    _ = v16
    return 2
}
func genFn1(p0 uint64) string {
    p0 = (p0 / ((18446744073709551615) | 1))
    var v17 int8 = 42
    _ = p0
    _ = v17
    return ("hello" + "")
}

func main() {
    var v1 uint64 = 858
    var v2 float32 = (((-2.25 / 0.1) + (float32(0.1) * 1e10)) / (-2.25 * 0.1))
    v3 := make(map[uint32]string)
    var v4 uint = 496
    var v5 int = -652
    v6 := make(map[string]bool)
    v18 := []int{4294967295, (^(v5))}
    v1 = (v1 % ((42) | 1))
    v19 := GenS1{F0: 2147483647, F1: (false || true), F2: (!(true || (-1 < -976)))}
    if len(v18) > 0 {
        v18[0] = (v5 ^ 127)
    }
    v1 = (v1 * v1)
    fmt.Println(v1, v1, v2)
    v20 := GenS1{F0: 42, F1: true, F2: (!((^(v4)) == (^(v4))))}
    var v21 int32 = int32(v5)
    v22 := GenS2{F0: 42, F1: 32767, F2: 296}
    fmt.Println(v4)
    var v23 int64 = -1
    v18 = append(v18, (v5 << ((v4) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
