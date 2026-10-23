package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint; F2 int32 }
type GenS3 struct { F0 int64; F1 int; F2 string }
type GenS4 struct { F0 uint; F1 uint64; F2 uint8 }
func genFn0(p0 uint8, p1 int16, p2 uint32) uint32 {
    switch -1 {
    case -1:
        v6 := GenS1{F0: "\n\t", F1: 32767}
        _ = v6
    case 32768:
        if false {
            var v7 uint16 = 172
            _ = v7
        }
    case 42:
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 int16 = p1
            _ = fi8
            _ = v9
        }
    }
    var v10 uint32 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return p2
}
func genFn1() int16 {
    v11 := []int{-595, 1, 2147483647, -9223372036854775808}
    var v12 int16 = 1000
    switch 128 {
    case -1:
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    default:
        _ = len(v11)
    }
    if (!true) {
        if true {
            v13 := GenS1{F0: "", F1: 2}
            _ = v13
        } else {
            _ = len(v11)
            if true {
                v14 := GenS2{F0: 42, F1: 1, F2: -2147483648}
                var v15 float32 = ((float32(-2.25) / 3.14159) + 3.14159)
                _ = v14
                _ = v15
            }
            v12 += (v12 ^ (-(v12)))
        }
        fmt.Println(v12)
    }
    switch "a" {
    case "\n\t":
        var v16 bool = false
        _ = v16
    case "foo bar":
        v17 := GenS1{F0: (("a" + "a") + ""), F1: 42}
        _ = v17
    default:
        v12 = -32768
    }
    _ = v11
    _ = v12
    return v12
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{2147483648, 1}
    var v3 uint32 = 72
    var v4 uint16 = 903
    v5 := make(map[string]string)
    var v18 int64 = -495
    var v19 bool = (true || true)
    for fi20 := 0; fi20 < 10; fi20++ {
        delete(v5, ("\n\t" + ("\n\t" + "\n\t")))
        fmt.Println(v18, v4)
        var v21 int8 = -1
        _ = fi20
        _ = v21
    }
    v18 = 1
    v18 = (-(v18))
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 bool = ((^(v3)) >= (v3 * v3))
        fmt.Println(fi22, v23)
        _ = fi22
        _ = v23
    }
    var v24 float64 = 1e10
    switch "\n\t" {
    case "":
        v2 = append(v2, (v3 ^ v3))
    case "a":
        v2 = append(v2, v3)
    case "foo bar":
        var v25 int8 = 100
        _ = v25
    }
    switch -9223372036854775808 {
    case 32768:
        v26 := GenS1{F0: "hello", F1: (v4 ^ (v4 << ((4294967295) & 63)))}
        _ = v26
    case 128:
        v27 := GenS3{F0: v18, F1: 32768, F2: "foo bar"}
        _ = v27
    default:
        switch 2147483647 {
        case 127:
            v1[(v3 >> ((0) & 63))] = genFn0(2, -32768, genFn0(2, 1, v3))
        case 32768:
            v24 -= ((v24 + ((((float64(1.5) * (v24 - v24)) - -2.25) + float64((((3.14159 * ((float32(((3.14159 / v24) * (0.1 / v24))) / ((float32((v24 / 0.1)) / 1e10) / 1.0)) + 1e10)) - (-2.25 / (float32(((0.1 - ((v24 / (1e10 + (v24 * (v24 / float64(((3.14159 + ((0.0 * 1e10) + ((1.5 + float32(v24)) + 0.1))) * 1.5)))))) / ((v24 - (1e10 + v24)) / v24))) * float64(0.1))) + 1e10))) + 3.14159))) + 0.0)) - v24)
        }
    }
    switch ("" + "\n\t") {
    case "a":
        var v28 uint32 = (^(v3))
        _ = v28
    case "":
        var v29 int8 = 42
        _ = v29
    default:
        v24 += (((0.1 - v24) - v24) + v24)
    }
    fmt.Println(v24, v18)
    var v30 int8 = 42
    for fi31 := 0; fi31 < 10; fi31++ {
        var v32 int64 = (v18 << ((255) & 63))
        _ = fi31
        _ = v32
    }
    if false {
        v33 := GenS4{F0: 278, F1: 0, F2: 51}
        _ = v33
    } else {
        var v34 bool = v19
        var v35 uint32 = genFn0(1, 413, v3)
        _ = v34
        _ = v35
    }
    if (!("\n\t" >= ("foo bar" + "hello"))) {
        if true {
            _ = len(v1)
            var v36 int32 = 905
            _ = v36
        }
        v37 := GenS1{F0: "hello", F1: 1}
        var v38 uint8 = 2
        _ = v37
        _ = v38
    } else {
        fmt.Println(v30, v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
