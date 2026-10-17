package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 bool; F1 bool; F2 string }
func genFn0(p0 uint8) bool {
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 float64 = 0.1
        var v7 int16 = 1000
        var v8 uint32 = 851
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    v9 := []int{}
    _ = p0
    _ = v9
    return (847 > 260)
}
func genFn1(p0 int64) bool {
    p0 = -9223372036854775808
    var v10 float32 = -2.25
    fmt.Println(v10, p0)
    v10 += float32(((float64(v10) / 1.0) + -2.25))
    for fi11 := 0; fi11 < 17; fi11++ {
        p0 = p0
        _ = fi11
    }
    _ = p0
    _ = v10
    return (!((0 != 1) || (0 <= 1)))
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 65535
    v3 := make(map[int]int)
    v4 := make(map[uint32]string)
    for fi12 := 0; fi12 < 6; fi12++ {
        fmt.Println(fi12, fi12)
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := GenS1{F0: (fi13 % ((-167) | 1)), F1: -32768}
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    switch (("a" + "") + "a") {
    case "foo bar":
        v2 += v2
    case "hello":
        v2 += v2
    default:
        v15 := []int{267, -39, 4294967295}
        _ = v15
    }
    v2 -= (v2 * (v2 << (((v2 ^ v2)) & 63)))
    var v16 float32 = 1e10
    for fi17 := 0; fi17 < 7; fi17++ {
        for fi18 := 0; fi18 < 16; fi18++ {
            v2 = (v2 % (((v2 % (((v2 * (v2 | v2))) | 1))) | 1))
            v19 := GenS1{F0: fi17, F1: 0}
            _ = fi18
            _ = v19
        }
        v1[fi17] = fi17
        v16 -= 1e10
        _ = fi17
    }
    v16 -= v16
    var v20 uint32 = 161
    var v21 float32 = (v16 * (v16 - 3.14159))
    v22 := GenS1{F0: -2147483648, F1: -629}
    fmt.Println(v16)
    var v23 float64 = 1e10
    v24 := GenS2{F0: true, F1: ((false || (!false)) && false), F2: "a"}
    fmt.Println(v20)
    fmt.Println(v2)
    if (2 == 1) {
        var v25 int32 = -375
        var v26 bool = (!false)
        _ = v25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
