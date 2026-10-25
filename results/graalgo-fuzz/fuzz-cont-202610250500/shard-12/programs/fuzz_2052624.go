package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS4 struct { F0 int64; F1 uint16 }
func genFn0(p0 int64, p1 int8, p2 uint) uint32 {
    var v5 uint32 = 956
    v6 := []int{9223372036854775807, 0, -839}
    switch ("" + "foo bar") {
    case "foo bar":
        var v7 uint64 = 0
        _ = v7
    default:
        fmt.Println(p1, p2, p2)
    }
    switch ("a" + "a") {
    case "":
        var v8 int = 32768
        _ = v8
    case "foo bar":
        var v9 int32 = 42
        _ = v9
    case "hello":
        switch 32768 {
        case -708:
            p2 = (p2 & (p2 >> ((p2) & 63)))
        case -9223372036854775808:
            var v10 int = 127
            _ = v10
        case 2147483647:
            fmt.Println(p0, p1)
        default:
            if len(v6) > 0 {
                v6[0] = 32767
            }
        }
    default:
        fmt.Println(p2, p1, p0)
    }
    v11 := GenS1{F0: "a"}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v5
    _ = v6
    return (v5 + 2147483648)
}
func genFn1(p0 uint32) int {
    for fi12 := 0; fi12 < 14; fi12++ {
        if ((((((1e10 / -2.25) - 0.0) + (float64((((1.5 / float32(3.14159)) * ((0.0 + 0.1) * (float32(-2.25) * ((0.1 / 0.1) + 0.0)))) / 1.5)) * ((0.1 + float64(1e10)) - ((-2.25 / (float64(0.1) + float64(1e10))) + (0.1 * float64(float32(-2.25))))))) * 1e10) + ((3.14159 / -2.25) + (1.5 - 1e10))) == (0.1 + 0.0)) {
            v13 := []int{328, fi12}
            p0 = 843
            for fi14 := 0; fi14 < 18; fi14++ {
                var v15 bool = false
                _ = fi14
                _ = v15
            }
            _ = v13
        } else {
            v16 := GenS2{F0: 32767, F1: 187}
            v17 := []int{9223372036854775807, fi12}
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 14; fi18++ {
            var v19 int8 = 100
            _ = fi18
            _ = v19
        }
        v20 := []int{}
        _ = fi12
        _ = v20
    }
    v21 := []int{-461, 944}
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    _ = v21
    return -914
}

func main() {
    var v1 int32 = 106
    v2 := []int{-1}
    v3 := make(map[uint32]string)
    v4 := []uint32{42}
    var v22 uint8 = 128
    _ = len(v2)
    v1 = (-(v1))
    v23 := GenS1{F0: "foo bar"}
    v22 += (v22 ^ v22)
    switch -2147483648 {
    case 127:
        for fi24 := 0; fi24 < 11; fi24++ {
            fmt.Println(v1, v1, v1)
            _ = fi24
        }
    case -9223372036854775808:
        switch "" {
        case "":
            var v25 uint = 255
            _ = v25
        default:
            fmt.Println(v22, v22)
        }
    default:
        var v26 int32 = 799
        _ = v26
    }
    var v27 uint16 = 707
    for fi28 := 0; fi28 < 14; fi28++ {
        var v29 int32 = v1
        v30 := GenS3{F0: genFn0(1, 21, 42), F1: int32(v22), F2: v22}
        v31 := GenS2{F0: -306, F1: (v27 | 38)}
        _ = fi28
        _ = v29
        _ = v30
        _ = v31
    }
    v32 := GenS4{F0: 9223372036854775807, F1: 1}
    var v33 float64 = ((-2.25 / 1.0) * 0.0)
    for fi34 := 0; fi34 < 6; fi34++ {
        v1 -= 42
        v35 := GenS2{F0: 1000, F1: (v27 ^ (v27 & v27))}
        for fi36 := 0; fi36 < 9; fi36++ {
            v22 = 195
            var v37 uint64 = 156
            var v38 int16 = -32768
            v22 -= 128
            v37 = uint64(v27)
            var v39 uint16 = 255
            _ = fi36
            _ = v37
            _ = v38
            _ = v39
        }
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
