package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32 }
func genFn0() uint {
    if (1.5 <= 1e10) {
        v7 := GenS1{F0: 9223372036854775807, F1: 4294967295}
        _ = v7
    } else {
        var v8 int16 = 42
        _ = v8
    }
    var v9 uint16 = 255
    v10 := []int{-925}
    v11 := GenS2{F0: -317, F1: 94}
    _ = v10
    _ = v11
    _ = v9
    return 287
}
func genFn1(p0 uint, p1 int) uint32 {
    for fi12 := 0; fi12 < 16; fi12++ {
        v13 := GenS2{F0: p1, F1: 1}
        if (((fi12 / ((fi12) | 1)) == p1) && (737 == 1)) {
            var v14 int32 = -2147483648
            var v15 bool = true
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 10; fi16++ {
            fmt.Println(p0, p0, fi16)
            fmt.Println(fi12)
            p1 = (fi16 * 32767)
            p1 -= (fi16 >> ((255) & 63))
            _ = fi16
        }
        _ = fi12
        _ = v13
    }
    var v17 uint8 = 2
    for fi18 := 0; fi18 < 16; fi18++ {
        v19 := []int{471, (p1 | (p1 / ((4294967295) | 1)))}
        var v20 string = "hello"
        if true {
            v17 = 145
            v21 := GenS3{F0: v17}
            _ = v21
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    v17 += (^(v17))
    var v22 int = p1
    _ = p0
    _ = p1
    _ = v17
    _ = v22
    return 168
}

func main() {
    var v1 int64 = -1
    var v2 uint = 170
    var v3 int8 = 127
    var v4 uint8 = 128
    v5 := make(map[uint32]uint32)
    v6 := []string{"", "", ("foo bar" + "foo bar")}
    _ = len(v5)
    v2 -= (v2 ^ genFn0())
    v23 := GenS3{F0: 0}
    v1 += (v1 % ((v1) | 1))
    var v24 int8 = v3
    v4 = v4
    v25 := GenS4{F0: genFn1((v2 % ((v2) | 1)), 42)}
    if false {
        v26 := GenS4{F0: 0}
        v27 := GenS1{F0: v1, F1: (v2 << ((genFn0()) & 63))}
        _ = v26
        _ = v27
    } else {
        v24 = v24
        v28 := GenS3{F0: v4}
        v4 -= v4
        _ = v28
    }
    switch ("a" + ("\n\t" + "foo bar")) {
    case "a":
        var v29 float32 = 1e10
        _ = v29
    case "foo bar":
        var v30 int8 = v24
        _ = v30
    }
    v31 := GenS2{F0: -31, F1: v24}
    if ((v1 % ((v1) | 1)) > v1) {
        v1 = (v1 << ((412) & 63))
    }
    var v32 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
