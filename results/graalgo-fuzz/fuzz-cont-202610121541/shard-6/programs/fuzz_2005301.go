package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint }
type GenS2 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint, p1 uint16, p2 uint16) uint8 {
    p2 -= (p2 & p1)
    var v7 uint = 326
    switch -1 {
    case 0:
        if (-128 > 1) {
            if false {
                fmt.Println(p2)
            }
            v8 := GenS1{F0: 1, F1: ("\n\t" + ""), F2: (^(p0))}
            var v9 float64 = float64(0.1)
            _ = v8
            _ = v9
        }
    case -599:
        v10 := []int{162}
        _ = v10
    case 128:
        v11 := GenS1{F0: 42, F1: "a", F2: p0}
        _ = v11
    default:
        var v12 uint32 = 2147483647
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}
func genFn1(p0 int) string {
    v13 := GenS1{F0: 0, F1: "a", F2: 0}
    v14 := []int{128}
    for fi15 := 0; fi15 < 15; fi15++ {
        fmt.Println(p0, p0)
        _ = len(v14)
        _ = len(v14)
        _ = fi15
    }
    fmt.Println(p0)
    _ = p0
    _ = v13
    _ = v14
    return "hello"
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]string)
    var v3 uint32 = 696
    v4 := make(map[uint32]bool)
    var v5 int = 162
    var v6 int16 = -1
    var v16 int16 = v6
    v17 := GenS2{F0: 128, F1: 2}
    v18 := v1[""]
    v16 = (v6 << ((4294967295) & 63))
    v16 += v6
    if (v18 || v18) {
        if v18 {
            v19 := v4[(v3 % (((v3 / (((v3 - v3)) | 1))) | 1))]
            _ = v19
        }
    }
    fmt.Println(v16, v6)
    if ((1.5 / float32(float64(1.5))) < (1e10 / -2.25)) {
        var v20 bool = ((v18 && true) && v18)
        _ = v20
    } else {
        v5 -= (-(v5))
    }
    v16 += 1000
    switch "foo bar" {
    case "a":
        switch (v5 + -728) {
        case 127:
            v3 += (v3 * v3)
        case 0:
            v21 := v1["a"]
            _ = v21
        }
    case "":
        if false {
            v3 = 65535
        } else {
            v22 := []int{42, v5}
            _ = v22
        }
    case "foo bar":
        var v23 int64 = 42
        _ = v23
    default:
        v6 += (v6 | (^(v6)))
    }
    var v24 int = 0
    v25 := GenS2{F0: genFn0(255, 32767, 26), F1: 42}
    v26 := GenS1{F0: 290, F1: ("a" + ("\n\t" + "")), F2: 65535}
    v27 := GenS1{F0: int32(v3), F1: "hello", F2: 18446744073709551615}
    var v28 int64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
