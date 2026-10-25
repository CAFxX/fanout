package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 int64 }

func main() {
    var v1 bool = (1 <= 18446744073709551615)
    var v2 uint8 = 42
    v3 := make(map[int]bool)
    var v4 uint = 65535
    var v5 uint64 = 18446744073709551615
    if false {
        v4 -= (v4 & v4)
        v4 -= 42
        var v6 uint16 = 255
        _ = v6
    }
    v5 -= (v5 + (^(v5)))
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := v3[fi7]
        _ = fi7
        _ = v8
    }
    if v1 {
        v9 := []int{-459, 32768}
        v5 += v5
        for fi10 := 0; fi10 < 7; fi10++ {
            v5 -= 822
            _ = fi10
        }
        _ = v9
    }
    v2 = v2
    if (-4 == 42) {
        v11 := GenS1{F0: 42, F1: 74}
        var v12 int8 = 1
        _ = v11
        _ = v12
    }
    if (!((^(v4)) == 255)) {
        var v13 uint8 = uint8(v4)
        fmt.Println(v4, v2)
        for fi14 := 0; fi14 < 16; fi14++ {
            v15 := []int{(fi14 / (((fi14 + -168)) | 1))}
            v3[fi14] = false
            var v16 int64 = 413
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v13
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        v4 -= (v4 * (^(v4)))
        v5 += v5
        _ = fi17
    }
    if v1 {
        if (false || (0 > -1)) {
            var v18 uint16 = 0
            v19 := []int{}
            var v20 int = -44
            _ = v18
            _ = v19
            _ = v20
        }
        for fi21 := 0; fi21 < 10; fi21++ {
            v5 = v5
            var v22 int = (fi21 ^ 2147483647)
            for fi23 := 0; fi23 < 19; fi23++ {
                v24 := []int{}
                _ = fi23
                _ = v24
            }
            _ = fi21
            _ = v22
        }
        v25 := GenS2{F0: v2, F1: 2}
        _ = v25
    }
    var v26 int64 = 948
    if (!(v1 && false)) {
        fmt.Println(v4, v2)
        switch "\n\t" {
        case "hello":
            for fi27 := 0; fi27 < 5; fi27++ {
                fmt.Println(v2)
                _ = fi27
            }
        case "foo bar":
            v2 -= v2
        }
        for fi28 := 0; fi28 < 12; fi28++ {
            var v29 string = "a"
            var v30 int8 = 42
            var v31 int32 = 0
            _ = fi28
            _ = v29
            _ = v30
            _ = v31
        }
    } else {
        v4 = v4
        for fi32 := 0; fi32 < 11; fi32++ {
            v26 += v26
            v26 -= 55
            _ = fi32
        }
    }
    var v33 uint = 65535
    if true {
        v26 += v26
        for fi34 := 0; fi34 < 19; fi34++ {
            var v35 uint64 = (^(v5))
            fmt.Println(v35)
            _ = fi34
            _ = v35
        }
    }
    for fi36 := 0; fi36 < 5; fi36++ {
        v37 := GenS3{F0: 1, F1: fi36}
        _ = fi36
        _ = v37
    }
    v38 := []int{-796}
    v39 := GenS1{F0: 534, F1: -1}
    for fi40 := 0; fi40 < 11; fi40++ {
        v41 := GenS1{F0: 42, F1: 1}
        _ = fi40
        _ = v41
    }
    switch 2147483647 {
    case 0:
        if v1 {
            v42 := GenS2{F0: v2, F1: uint16(v5)}
            v5 = uint64(v33)
            var v43 bool = v1
            _ = v42
            _ = v43
        }
    case -751:
        for fi44 := 0; fi44 < 12; fi44++ {
            var v45 int32 = 185
            v33 -= (^(v4))
            var v46 float64 = ((0.1 + float64(((((3.14159 / 1.0) + 0.0) + 0.1) - ((0.0 / (1e10 + float32((float64(-2.25) - 3.14159)))) / 0.1)))) * 0.1)
            _ = fi44
            _ = v45
            _ = v46
        }
    case -723:
        v2 = v2
    default:
        v47 := GenS4{F0: v26}
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v33
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
