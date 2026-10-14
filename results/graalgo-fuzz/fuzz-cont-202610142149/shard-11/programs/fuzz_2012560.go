package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int, p1 uint, p2 uint32) int8 {
    var v6 int = (p0 * p0)
    switch "a" {
    case "\n\t":
        v7 := []int{p0}
        _ = v7
    case "foo bar":
        fmt.Println(p2)
    case "hello":
        for fi8 := 0; fi8 < 2; fi8++ {
            p1 = p1
            _ = fi8
        }
    default:
        var v9 float32 = 0.1
        _ = v9
    }
    if (true && (((-2.25 * 1e10) / 3.14159) != -2.25)) {
        var v10 float32 = (-2.25 + 1.5)
        v11 := []int{(v6 & (v6 - 32768)), p0, 128, 0}
        _ = v10
        _ = v11
    } else {
        var v12 int16 = 32767
        v6 = (^(p0))
        v13 := GenS1{F0: (false || (true || false)), F1: (!true), F2: 0}
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: false, F1: false, F2: p1}
    switch ("foo bar" + "hello") {
    case "\n\t":
        var v15 bool = (uint8(p2) == 56)
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v6
    return 0
}

func main() {
    v1 := []uint32{531, 2147483647, 0}
    v2 := []int{-1, 0, -2147483648}
    var v3 int16 = -32768
    var v4 int64 = -1
    v5 := make(map[string]uint32)
    for fi16 := 0; fi16 < 18; fi16++ {
        if true {
            var v17 string = "a"
            _ = v17
        }
        for fi18 := 0; fi18 < 4; fi18++ {
            var v19 uint64 = 4294967295
            var v20 uint = 161
            v21 := GenS1{F0: (42 < 2), F1: (((v19 & (v19 % (((v19 - (v19 >> (((v20 / ((v20) | 1))) & 63)))) | 1))) != (v19 * (^(v19)))) || false), F2: uint(v19)}
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
        _ = fi16
    }
    if true {
        v22 := v5["foo bar"]
        _ = v22
    } else {
        fmt.Println(v4, v3)
        switch "\n\t" {
        case "foo bar":
            _ = len(v2)
        default:
            v4 -= -9223372036854775808
        }
    }
    var v23 float64 = (float64((float32(((-2.25 * -2.25) / 1.0)) - (1e10 + (-2.25 / float32((((((float64((float32(3.14159) / 1e10)) / 1e10) - float64(0.0)) * 3.14159) * 0.0) - (0.1 / float64(float32((3.14159 / -2.25)))))))))) / 1.0)
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    v4 -= (v4 & v4)
    delete(v5, "hello")
    if (!(!false)) {
        v23 += (float64(1e10) + ((v23 - 0.1) + (v23 - float64(-2.25))))
        var v24 uint64 = 4294967295
        v24 -= (v24 & v24)
        _ = v24
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        var v26 int16 = (^(v3))
        _ = fi25
        _ = v26
    }
    v4 += 901
    if true {
        var v27 int16 = v3
        v28 := GenS1{F0: (false || false), F1: true, F2: 141}
        _ = v27
        _ = v28
    } else {
        fmt.Println(v3, v23)
    }
    var v29 bool = true
    for fi30 := 0; fi30 < 20; fi30++ {
        v29 = v29
        _ = fi30
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        v32 := GenS2{F0: 257}
        v33 := GenS1{F0: v29, F1: ((42 == 65535) && v29), F2: 42}
        _ = fi31
        _ = v32
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
