package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 int64; F2 uint64 }
type GenS5 struct { F0 int16; F1 bool }
func genFn0(p0 uint64, p1 int8) uint64 {
    if (p1 != (p1 + -1)) {
        if false {
            for fi6 := 0; fi6 < 11; fi6++ {
                p1 = -128
                _ = fi6
            }
            var v7 float64 = 1.5
            p1 += (p1 ^ 1)
            _ = v7
        }
        var v8 uint64 = p0
        v9 := []int{-1, 32767}
        _ = v8
        _ = v9
    } else {
        fmt.Println(p0)
        v10 := []int{-896, 127, -649}
        _ = v10
    }
    v11 := GenS1{F0: (p1 | p1), F1: 0, F2: 42}
    if (("" <= "foo bar") && ("\n\t" <= "hello")) {
        fmt.Println(p1, p0)
        v12 := []int{-58, 32768, 128, 61}
        _ = v12
    }
    v13 := []int{677, 9223372036854775807, 4294967295}
    _ = p0
    _ = p1
    _ = v11
    _ = v13
    return 396
}
func genFn1(p0 uint32, p1 int32) string {
    var v14 int64 = 251
    v15 := GenS2{F0: 2, F1: 100}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return "foo bar"
}

func main() {
    v1 := []int{}
    var v2 uint32 = 651
    v3 := make(map[int]string)
    v4 := []int{570}
    var v5 string = "a"
    v16 := GenS3{F0: 255}
    var v17 uint16 = 0
    v18 := v3[32767]
    if (3.14159 >= float32(1e10)) {
        v19 := GenS4{F0: 42, F1: -487, F2: genFn0(2, 0)}
        v17 -= v17
        if true {
            for fi20 := 0; fi20 < 9; fi20++ {
                v18 = (("\n\t" + (("" + "hello") + "a")) + v5)
                if len(v1) > 0 {
                    v1[0] = (fi20 % ((fi20) | 1))
                }
                _ = fi20
            }
            v17 = (v17 ^ 32767)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            var v21 uint64 = 188
            v21 -= uint64(v2)
            _ = v21
        }
        _ = v19
    } else {
        var v22 uint32 = v2
        _ = v22
    }
    v2 += v2
    _ = len(v1)
    v23 := GenS5{F0: 42, F1: ((!(false || true)) && false)}
    if true {
        v24 := GenS3{F0: (v17 & v17)}
        v5 = "foo bar"
        _ = v24
    }
    if false {
        switch "hello" {
        case "hello":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "\n\t":
            for fi25 := 0; fi25 < 11; fi25++ {
                var v26 string = v18
                _ = fi25
                _ = v26
            }
        case "a":
            fmt.Println(v5)
        default:
            fmt.Println(v5, v18)
        }
    }
    if true {
        var v27 uint32 = 2
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v27
    }
    v5 = v5
    v17 -= (v17 & v17)
    fmt.Println(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
