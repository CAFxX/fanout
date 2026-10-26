package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int64; F1 int8 }
func genFn0(p0 uint32, p1 int8, p2 uint32) uint16 {
    var v5 uint32 = (p0 - p0)
    var v6 int16 = int16(p1)
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: p2, F1: p1}
        var v9 uint64 = 274
        if true {
            p0 -= p2
        } else {
            p0 = (^(v5))
            fmt.Println(fi7, v5, v9)
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    p1 -= p1
    var v10 float64 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    _ = v6
    return 948
}
func genFn1(p0 int) int64 {
    if (859 == 42) {
        var v11 uint64 = 4294967295
        v12 := []int{}
        fmt.Println(v11, v11)
        _ = v11
        _ = v12
    } else {
        p0 += (p0 * 1)
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 int32 = 0
        _ = fi13
        _ = v14
    }
    _ = p0
    return 42
}

func main() {
    var v1 string = "a"
    v2 := []uint32{1}
    var v3 int8 = 42
    v4 := make(map[uint32]string)
    var v15 float64 = (3.14159 + (1.5 * -2.25))
    if false {
        for fi16 := 0; fi16 < 5; fi16++ {
            var v17 uint16 = genFn0(125, (v3 - v3), 2)
            _ = fi16
            _ = v17
        }
        v18 := GenS1{F0: 2147483647, F1: 100}
        _ = v18
    } else {
        v3 = v3
        v19 := GenS1{F0: 65535, F1: 0}
        _ = v19
    }
    v2 = append(v2, 2147483647)
    v20 := GenS1{F0: 65535, F1: (v3 + v3)}
    var v21 uint8 = 192
    var v22 uint32 = 2147483648
    if false {
        var v23 int8 = v3
        _ = v23
    }
    if (1 >= -32768) {
        v15 = 1.5
        var v24 uint8 = (v21 - 42)
        _ = v24
    } else {
        var v25 uint8 = uint8(v22)
        _ = v25
    }
    var v26 int8 = (v3 % (((v3 >> ((255) & 63))) | 1))
    switch int(v26) {
    case -422:
        fmt.Println(v21, v21, v21)
    case -1:
        switch -235 {
        case 4294967295:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 0:
            v26 += v26
        case -1:
            var v27 int8 = 0
            _ = v27
        }
    }
    switch v1 {
    case "hello":
        var v28 bool = (false && true)
        _ = v28
    case "a":
        fmt.Println(v22, v3)
    default:
        v29 := GenS1{F0: v22, F1: v26}
        _ = v29
    }
    var v30 float64 = (-2.25 - v15)
    switch v1 {
    case "a":
        var v31 int32 = 0
        _ = v31
    case "hello":
        var v32 bool = true
        _ = v32
    case "foo bar":
        v22 -= (v22 ^ v22)
    }
    if (42 >= 634) {
        var v33 float64 = v15
        if (-196 == -1) {
            v26 = v3
            var v34 float32 = ((1e10 * (1e10 / 0.1)) / 1.5)
            v30 -= (float64(0.1) / -2.25)
            _ = v34
        }
        v35 := GenS2{F0: -9223372036854775808, F1: v3}
        _ = v33
        _ = v35
    } else {
        v2 = append(v2, 581)
        if (false || (false || true)) {
            var v36 int64 = -721
            var v37 uint = 565
            _ = v36
            _ = v37
        }
    }
    if false {
        fmt.Println(v3)
    }
    var v38 uint16 = 160
    var v39 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v30
    _ = v38
    _ = v39
    _ = v4
}
