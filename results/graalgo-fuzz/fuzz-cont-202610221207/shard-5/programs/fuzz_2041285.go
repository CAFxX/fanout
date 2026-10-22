package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint64 }
func genFn0(p0 int32, p1 uint8, p2 uint32) int16 {
    switch -958 {
    case 32767:
        p1 -= p1
    case 1:
        fmt.Println(p0)
    default:
        switch "a" {
        case "a":
            for fi7 := 0; fi7 < 2; fi7++ {
                fmt.Println(p2)
                _ = fi7
            }
        case "hello":
            var v8 int64 = -791
            _ = v8
        }
    }
    p0 = int32(p2)
    var v9 float32 = -2.25
    v10 := []int{902}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 1000
}
func genFn1(p0 uint64) int16 {
    for fi11 := 0; fi11 < 17; fi11++ {
        fmt.Println(fi11, p0)
        switch (fi11 ^ fi11) {
        case 1:
            var v12 uint = 1
            _ = v12
        case 2147483647:
            p0 -= p0
        case -1:
            p0 += p0
        }
        var v13 uint64 = (p0 & 266)
        _ = fi11
        _ = v13
    }
    p0 += p0
    _ = p0
    return genFn0(0, 128, 46)
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{65535, 2, 2, 4294967295}
    v3 := make(map[int]int)
    var v4 int32 = 368
    v5 := make(map[string]int)
    var v6 int8 = -4
    switch ("a" + "\n\t") {
    case "a":
        v14 := GenS1{F0: (("foo bar" + (("foo bar" + "foo bar") + "a")) < "hello"), F1: 59, F2: 2}
        _ = v14
    default:
        var v15 bool = false
        _ = v15
    }
    delete(v3, -270)
    var v16 float32 = float32(0.0)
    if true {
        var v17 int16 = -32768
        switch "foo bar" {
        case "\n\t":
            for fi18 := 0; fi18 < 1; fi18++ {
                v17 -= 32767
                _ = fi18
            }
        case "foo bar":
            v19 := GenS1{F0: (127 == -2147483648), F1: 24, F2: 696}
            _ = v19
        default:
            var v20 int64 = 1
            _ = v20
        }
        _ = v17
    } else {
        var v21 int16 = -405
        var v22 int8 = (v6 ^ (v6 << ((4294967295) & 63)))
        v4 += -1
        _ = v21
        _ = v22
    }
    v23 := GenS1{F0: ((((((float64((((0.1 - (v16 / ((float32(((float64(3.14159) / 1e10) * 1.5)) + (v16 / float32(-2.25))) / 1.5))) / (0.1 + v16)) - 1.5)) * (1e10 * float64(1e10))) / 1.0) / 0.1) - 3.14159) * 3.14159) == (-2.25 - (((((0.1 * float64(0.0)) + ((1e10 - 1.5) + 0.1)) - 1.5) - 0.0) / -2.25))), F1: 116, F2: 1}
    v24 := GenS1{F0: (-821 != 0), F1: 128, F2: 690}
    fmt.Println(v6, v16, v4)
    fmt.Println(v6)
    var v25 int = -2147483648
    for fi26 := 0; fi26 < 1; fi26++ {
        switch (v25 % ((fi26) | 1)) {
        case 9223372036854775807:
            var v27 float32 = float32(3.14159)
            _ = v27
        case -1:
            var v28 uint = 18446744073709551615
            _ = v28
        }
        v4 -= 175
        for fi29 := 0; fi29 < 5; fi29++ {
            if (42 != 0) {
                v4 -= 182
                v4 += v4
            }
            _ = fi29
        }
        _ = fi26
    }
    if false {
        v6 = 62
    }
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
