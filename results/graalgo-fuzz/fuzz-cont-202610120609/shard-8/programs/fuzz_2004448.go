package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int16 }
func genFn0(p0 int) bool {
    v4 := GenS1{F0: 2}
    if true {
        var v5 int8 = 127
        _ = v5
    }
    var v6 int8 = 100
    _ = p0
    _ = v4
    _ = v6
    return ((591 == 1) || true)
}
func genFn1(p0 bool) int {
    p0 = p0
    p0 = p0
    if true {
        p0 = (!(65535 == 65535))
    }
    switch "a" {
    case "hello":
        v7 := GenS2{F0: -1, F1: -32768}
        _ = v7
    case "foo bar":
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := []int{}
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = fi8
            _ = v9
        }
    case "a":
        var v10 int8 = 127
        _ = v10
    }
    _ = p0
    return 128
}

func main() {
    var v1 int8 = 127
    var v2 float64 = (((1.5 / 1.0) / 1.0) - 1.5)
    var v3 int64 = -689
    if ("\n\t" >= ("" + "foo bar")) {
        if (65535 < 277) {
            v1 -= v1
            var v11 bool = false
            _ = v11
        } else {
            var v12 float32 = (float32(v2) / 0.1)
            var v13 float64 = float64((float32((((v2 - v2) / ((-2.25 - -2.25) + v2)) - v2)) - v12))
            _ = v12
            _ = v13
        }
        v14 := []int{42, 0}
        var v15 uint32 = 4294967295
        _ = v14
        _ = v15
    }
    v16 := []int{457, 42, genFn1(false)}
    if ((-784 > (v3 | v3)) || (!(!(!false)))) {
        v17 := GenS2{F0: (v3 + v3), F1: 1}
        _ = v17
    }
    var v18 uint = 31
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 float64 = v2
        v21 := GenS2{F0: 102, F1: 42}
        _ = fi19
        _ = v20
        _ = v21
    }
    v18 -= v18
    for fi22 := 0; fi22 < 13; fi22++ {
        for fi23 := 0; fi23 < 9; fi23++ {
            v3 = 9223372036854775807
            var v24 uint64 = 4294967295
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            _ = fi23
            _ = v24
        }
        fmt.Println(v2, v3)
        _ = len(v16)
        _ = fi22
    }
    var v25 uint8 = 245
    var v26 int8 = v1
    v25 += v25
    var v27 bool = true
    if len(v16) > 0 {
        v16[0] = 127
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
