package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint64; F1 int16 }
func genFn0(p0 uint64, p1 int32) int32 {
    v6 := []int{128, -844, 0, 32767}
    v7 := GenS1{F0: 65535, F1: 1}
    for fi8 := 0; fi8 < 5; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            if true {
                p0 += (p0 & p0)
            }
            p0 = p0
            _ = fi9
        }
        var v10 string = "hello"
        _ = fi8
        _ = v10
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return p1
}
func genFn1(p0 int32, p1 uint) uint {
    v11 := []int{0, 239}
    p1 = (p1 - p1)
    v12 := GenS1{F0: 65535, F1: 314}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 2
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[int]int)
    var v3 uint16 = 42
    var v4 int16 = 32767
    var v5 bool = (!true)
    v13 := GenS1{F0: uint16(v1), F1: 2147483648}
    var v14 bool = (!v5)
    if v5 {
        v3 = 665
        if ((false && false) && true) {
            var v15 uint64 = (v1 & (v1 % ((185) | 1)))
            v16 := GenS1{F0: v3, F1: 980}
            _ = v15
            _ = v16
        } else {
            v17 := GenS2{F0: -2147483648, F1: 255}
            _ = v17
        }
    }
    if (uint(v1) > genFn1(-852, 489)) {
        v18 := []int{}
        v14 = v14
        _ = v18
    }
    v19 := GenS3{F0: (^(v1)), F1: (^(v4))}
    v5 = false
    v20 := GenS1{F0: (v3 & v3), F1: 4294967295}
    var v21 float32 = 1e10
    v22 := []int{32767}
    v23 := GenS1{F0: (v3 << ((1) & 63)), F1: 2147483647}
    v14 = (!(!false))
    v3 = (v3 | v3)
    var v24 bool = ((1.5 / 1e10) > 1.5)
    var v25 int64 = 0
    var v26 bool = (v1 != 1)
    var v27 float32 = (v21 + (float32(1.5) / 1e10))
    switch "a" {
    case "":
        v28 := v2[32767]
        _ = v28
    case "\n\t":
        v1 -= 2
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        v26 = v26
        if ((true && v26) || (!v14)) {
            switch fi29 {
            case 1:
                var v30 string = ("\n\t" + ((("foo bar" + ("a" + "")) + (("foo bar" + "a") + "hello")) + "hello"))
                _ = v30
            case -187:
                if (!((0 <= -1) && (622 >= (fi29 * (fi29 << ((210) & 63)))))) {
                    var v31 uint = 2
                    _ = v31
                }
            default:
                var v32 int64 = (v25 ^ v25)
                _ = v32
            }
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
