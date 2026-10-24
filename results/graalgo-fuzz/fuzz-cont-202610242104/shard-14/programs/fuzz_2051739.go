package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint32; F2 uint }
func genFn0(p0 uint, p1 int32) int16 {
    for fi7 := 0; fi7 < 20; fi7++ {
        fmt.Println(p1, p1, fi7)
        var v8 int = fi7
        v9 := []int{}
        _ = fi7
        _ = v8
        _ = v9
    }
    if false {
        if false {
            var v10 int8 = -125
            var v11 int32 = (p1 & (p1 >> ((1) & 63)))
            _ = v10
            _ = v11
        } else {
            p1 = (p1 % (((p1 - 391)) | 1))
        }
    } else {
        fmt.Println(p1)
        var v12 int = 32768
        _ = v12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        switch ("foo bar" + "hello") {
        case "\n\t":
            var v14 uint64 = uint64(p0)
            _ = v14
        case "hello":
            p1 -= 0
        }
        _ = fi13
    }
    _ = p0
    _ = p1
    return -32768
}
func genFn1(p0 int16, p1 uint, p2 uint16) uint {
    v15 := GenS1{F0: 32768}
    v16 := []int{-1, 127, 9223372036854775807}
    var v17 int16 = 42
    var v18 uint16 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    return p1
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 2147483647
    var v3 uint64 = 4294967295
    v4 := make(map[uint32]int)
    var v5 float64 = 0.0
    var v6 int64 = 0
    switch ("" + ("\n\t" + ("hello" + ("foo bar" + "a")))) {
    case "\n\t":
        var v19 uint16 = 274
        _ = v19
    default:
        fmt.Println(v6, v3, v5)
    }
    fmt.Println(v6)
    switch 42 {
    case 174:
        var v20 int64 = (v6 + 457)
        _ = v20
    default:
        v21 := v4[0]
        _ = v21
    }
    var v22 uint64 = v3
    if true {
        v2 += -384
    }
    v3 += (^(v22))
    v23 := []int{919, 4294967295}
    _ = len(v23)
    v24 := GenS1{F0: 79}
    v5 -= v5
    v5 = 1.5
    v4[2147483647] = -281
    var v25 string = "a"
    var v26 bool = false
    for fi27 := 0; fi27 < 4; fi27++ {
        var v28 float32 = (1e10 / 0.1)
        v26 = (1000 >= -32768)
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 8; fi29++ {
        var v30 float64 = ((float64(0.0) - (0.1 / (1e10 / 1.0))) / 1.0)
        for fi31 := 0; fi31 < 18; fi31++ {
            if (!true) {
                var v32 uint32 = 2147483647
                v30 -= (1e10 / ((((0.1 - 0.0) + 0.1) * 0.1) * v30))
                v30 = v30
                _ = v32
            } else {
                v33 := GenS2{F0: int(v2), F1: 653, F2: 4294967295}
                _ = v33
            }
            v2 -= v2
            _ = fi31
        }
        _ = fi29
        _ = v30
    }
    fmt.Println(v3, v3, v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
