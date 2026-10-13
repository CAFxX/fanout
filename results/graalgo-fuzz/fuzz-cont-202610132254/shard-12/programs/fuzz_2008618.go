package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 int; F2 uint16 }
func genFn0() bool {
    var v7 float64 = (1.5 / 1.0)
    if false {
        v8 := []int{32767, 32768, -9223372036854775808, 0}
        _ = v8
    } else {
        if false {
            for fi9 := 0; fi9 < 8; fi9++ {
                v10 := []int{(fi9 + (-(fi9))), -9223372036854775808, fi9}
                _ = fi9
                _ = v10
            }
            var v11 int = 72
            _ = v11
        }
    }
    var v12 uint8 = 32
    fmt.Println(v12, v7)
    _ = v12
    _ = v7
    return ((!(170 >= 255)) && false)
}
func genFn1(p0 uint64, p1 int64, p2 uint32) string {
    p0 = 548
    var v13 int32 = 42
    p2 -= (p2 * p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return ""
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = ((!(-944 != -1)) && true)
    v3 := make(map[int]int)
    v4 := make(map[uint32]bool)
    var v5 bool = true
    var v6 uint32 = 271
    v14 := GenS1{F0: (v6 | v6), F1: 1, F2: 0}
    var v15 uint8 = 224
    v16 := GenS2{F0: (^(v15))}
    v5 = (127 < (^(v15)))
    switch "foo bar" {
    case "\n\t":
        for fi17 := 0; fi17 < 1; fi17++ {
            var v18 bool = ((447 <= 193) && false)
            _ = fi17
            _ = v18
        }
    case "a":
        if v5 {
            v19 := GenS3{F0: 872, F1: -407, F2: 42}
            v15 = 68
            var v20 uint = 974
            _ = v19
            _ = v20
        } else {
            v21 := []int{}
            _ = v21
        }
    case "":
        v22 := []int{-747, 4294967295, 9223372036854775807, 1}
        _ = v22
    }
    v5 = v5
    switch 504 {
    case 1:
        v5 = (!v5)
    case 2147483647:
        v23 := v1[("foo bar" + ("a" + ""))]
        _ = v23
    case 32767:
        for fi24 := 0; fi24 < 14; fi24++ {
            fmt.Println(v2, v5)
            _ = fi24
        }
    }
    var v25 uint32 = (^(v6))
    if v5 {
        var v26 float64 = 1.5
        v27 := GenS3{F0: 768, F1: -881, F2: 42}
        var v28 uint16 = 0
        _ = v26
        _ = v27
        _ = v28
    } else {
        var v29 uint = 591
        _ = v29
    }
    v30 := GenS1{F0: (v25 - (^(v25))), F1: 1, F2: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
