package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int; F1 string; F2 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 uint32; F2 uint }
func genFn0(p0 int16, p1 string) string {
    if (561 >= 1) {
        var v6 int64 = 0
        v7 := GenS1{F0: 255, F1: (^(p0))}
        if (false && false) {
            v8 := GenS2{F0: 128, F1: "\n\t", F2: -2147483648}
            v9 := GenS2{F0: 348, F1: p1, F2: -949}
            if false {
                if false {
                    var v10 uint64 = 604
                    var v11 float64 = float64(-2.25)
                    _ = v10
                    _ = v11
                } else {
                    if (!(true || (!(true || (1 != 1))))) {
                        var v12 uint32 = 2
                        p1 = p1
                        fmt.Println(p0, p1, v12)
                        _ = v12
                    }
                }
                fmt.Println(v6)
                v13 := GenS2{F0: 745, F1: "a", F2: -881}
                _ = v13
            }
            _ = v8
            _ = v9
        } else {
            v14 := []int{-920, 4294967295}
            _ = v14
        }
        _ = v6
        _ = v7
    } else {
        v15 := GenS3{F0: 32767}
        _ = v15
    }
    if (p1 == p1) {
        p0 = p0
        p1 = (p1 + (p1 + p1))
    }
    fmt.Println(p0, p1)
    var v16 int = 504
    _ = p0
    _ = p1
    _ = v16
    return p1
}
func genFn1(p0 uint64) uint16 {
    fmt.Println(p0)
    v17 := GenS3{F0: 1000}
    var v18 float64 = 0.0
    p0 += 4294967295
    fmt.Println(p0, p0)
    _ = p0
    _ = v17
    _ = v18
    return 284
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 1
    var v3 uint = 220
    v4 := make(map[int]bool)
    var v5 int = -167
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 uint8 = (v2 << (((v3 * (v3 % ((v3) | 1)))) & 63))
        _ = fi19
        _ = v20
    }
    v21 := GenS2{F0: 32768, F1: ("\n\t" + (("hello" + "") + ("" + "hello"))), F2: (v5 - (v5 >> (((v3 % ((v3) | 1))) & 63)))}
    var v22 bool = false
    v3 = v3
    v23 := GenS3{F0: 444}
    v2 += 2
    v24 := GenS2{F0: 127, F1: "foo bar", F2: (v5 * (v5 - v5))}
    fmt.Println(v22, v22, v5)
    for fi25 := 0; fi25 < 10; fi25++ {
        if (!v22) {
            v26 := GenS4{F0: 42, F1: 523, F2: 4294967295}
            v22 = (!v22)
            _ = v26
        }
        fmt.Println(v3)
        v27 := []int{}
        _ = fi25
        _ = v27
    }
    v22 = v22
    v28 := []int{}
    var v29 int = -856
    if false {
        switch (-(v29)) {
        case 334:
            v29 -= v29
        case 4294967295:
            fmt.Println(v3, v3, v3)
        case 0:
            v30 := v4[(v29 ^ -9223372036854775808)]
            _ = v30
        }
    }
    var v31 float32 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
