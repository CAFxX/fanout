package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint8 }
func genFn0(p0 uint16) uint32 {
    var v5 int8 = 1
    v6 := []int{}
    _ = p0
    _ = v5
    _ = v6
    return 4294967295
}
func genFn1(p0 int64) uint64 {
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 int8 = -3
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 float64 = 0.1
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 13; fi11++ {
            fmt.Println(fi7, fi7)
            _ = fi11
        }
        _ = fi7
        _ = v8
    }
    if (100 == -74) {
        p0 += (p0 * (p0 ^ p0))
    }
    var v12 uint32 = 2147483648
    _ = p0
    _ = v12
    return 2
}

func main() {
    var v1 float64 = (((0.1 + 0.0) / 3.14159) + 3.14159)
    v2 := make(map[int]int)
    var v3 int64 = 9223372036854775807
    v4 := make(map[uint32]string)
    if (4294967295 >= 4294967295) {
        v1 = (v1 + v1)
        switch 128 {
        case 4294967295:
            var v13 int = 9223372036854775807
            _ = v13
        }
    } else {
        var v14 uint = 65535
        v15 := GenS1{F0: 0, F1: (v3 + v3), F2: uint8(v14)}
        var v16 uint = (v14 >> (((v14 | 790)) & 63))
        _ = v14
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: genFn0(36), F1: 960, F2: 133}
    v1 = v1
    for fi18 := 0; fi18 < 11; fi18++ {
        v3 = v3
        _ = fi18
    }
    v19 := GenS1{F0: genFn0(589), F1: (v3 & v3), F2: 252}
    switch 489 {
    case 4294967295:
        fmt.Println(v3, v1)
    case 9223372036854775807:
        v20 := []int{786, 4294967295, -43, 9223372036854775807}
        _ = v20
    }
    v21 := []int{1, -9223372036854775808}
    delete(v2, -2147483648)
    v1 += v1
    v1 = (v1 - v1)
    var v22 int32 = 1
    var v23 float32 = 0.0
    var v24 uint32 = 2
    switch "foo bar" {
    case "a":
        var v25 int16 = 1
        _ = v25
    default:
        for fi26 := 0; fi26 < 12; fi26++ {
            var v27 int = (fi26 * (fi26 >> ((2) & 63)))
            _ = fi26
            _ = v27
        }
    }
    v2[10] = 1
    var v28 float32 = v23
    for fi29 := 0; fi29 < 16; fi29++ {
        var v30 uint8 = 128
        var v31 float32 = -2.25
        var v32 string = (("" + ("\n\t" + "hello")) + "foo bar")
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    v24 = genFn0(42)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v4
}
