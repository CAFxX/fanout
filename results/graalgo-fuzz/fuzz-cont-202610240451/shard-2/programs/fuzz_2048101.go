package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 bool; F2 int16 }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 uint64; F1 string; F2 uint32 }
type GenS6 struct { F0 string; F1 int; F2 string }
func genFn0(p0 uint64) int64 {
    for fi7 := 0; fi7 < 20; fi7++ {
        switch ("\n\t" + "foo bar") {
        case "":
            for fi8 := 0; fi8 < 19; fi8++ {
                v9 := GenS1{F0: 124, F1: 127, F2: (false && (42 == 230))}
                _ = fi8
                _ = v9
            }
        default:
            for fi10 := 0; fi10 < 15; fi10++ {
                var v11 uint16 = 2
                _ = fi10
                _ = v11
            }
        }
        _ = fi7
    }
    v12 := GenS1{F0: -6, F1: 33, F2: (true && true)}
    var v13 uint16 = 32767
    fmt.Println(v13, p0)
    p0 = 135
    _ = p0
    _ = v12
    _ = v13
    return -9223372036854775808
}
func genFn1(p0 int16, p1 bool, p2 uint16) int32 {
    if p1 {
        fmt.Println(p1, p1, p2)
        switch 2147483647 {
        case 128:
            var v14 bool = (!false)
            _ = v14
        case -9223372036854775808:
            var v15 uint = 255
            _ = v15
        case 2147483647:
            var v16 float64 = float64(1e10)
            _ = v16
        }
    }
    v17 := []int{42}
    if (!(-2147483648 < 0)) {
        v18 := GenS1{F0: -128, F1: 1, F2: p1}
        var v19 uint32 = 213
        _ = v18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return -1
}

func main() {
    var v1 uint16 = 181
    var v2 int = 128
    var v3 uint32 = 257
    var v4 int8 = -80
    var v5 uint64 = 937
    var v6 float64 = 1e10
    fmt.Println(v3, v6, v3)
    if false {
        var v20 uint64 = v5
        v21 := GenS2{F0: 128}
        _ = v20
        _ = v21
    } else {
        for fi22 := 0; fi22 < 14; fi22++ {
            var v23 bool = false
            _ = fi22
            _ = v23
        }
        v1 = (v1 / ((v1) | 1))
    }
    var v24 string = "a"
    if (!(v24 >= "hello")) {
        fmt.Println(v4)
        v25 := GenS3{F0: -128, F1: (true && (1 >= -1)), F2: 137}
        v26 := GenS4{F0: 127, F1: "hello"}
        _ = v25
        _ = v26
    }
    switch v2 {
    case 32767:
        if false {
            var v27 uint32 = (v3 ^ uint32(v5))
            fmt.Println(v1, v6)
            _ = v27
        } else {
            v4 -= -80
            switch (v2 - v2) {
            case 32768:
                fmt.Println(v4, v5)
            }
        }
    }
    v28 := GenS5{F0: (^(v5)), F1: v24, F2: v3}
    for fi29 := 0; fi29 < 14; fi29++ {
        v2 -= v2
        v3 = uint32(v1)
        _ = fi29
    }
    v2 += v2
    v30 := GenS6{F0: v24, F1: 128, F2: ("\n\t" + v24)}
    var v31 uint32 = (v3 | (v3 + v3))
    if false {
        var v32 bool = false
        _ = v32
    }
    v4 += v4
    v2 = 801
    v33 := []int{-511, (^(v2)), 42, 423}
    v33 = append(v33, v2)
    fmt.Println(v24, v4, v6)
    var v34 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
