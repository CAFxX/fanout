package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint8 }
func genFn0(p0 uint32, p1 int16, p2 string) uint8 {
    var v7 int16 = (^(p1))
    v8 := []int{127, -9223372036854775808, 32767}
    switch -371 {
    case 4294967295:
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 uint8 = 1
            for fi11 := 0; fi11 < 6; fi11++ {
                v12 := GenS1{F0: -592, F1: 681, F2: (^(v10))}
                var v13 int = fi11
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = fi9
            _ = v10
        }
    default:
        switch -626 {
        case 4294967295:
            switch 32767 {
            case 0:
                _ = len(v8)
            default:
                var v14 float64 = 1.5
                _ = v14
            }
        case -9223372036854775808:
            if len(v8) > 0 {
                v8[0] = 1
            }
        }
    }
    v15 := GenS2{F0: 65535, F1: 111, F2: 2}
    var v16 int16 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 int32, p1 uint8) uint16 {
    fmt.Println(p0)
    var v17 int64 = 532
    v18 := GenS2{F0: 65535, F1: 0, F2: p1}
    if (1e10 != (0.0 / float32(0.1))) {
        v19 := []int{}
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 0
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 2
    var v3 uint32 = 370
    var v4 uint32 = v3
    v5 := make(map[uint32]bool)
    v6 := make(map[string]int)
    delete(v6, "\n\t")
    for fi20 := 0; fi20 < 7; fi20++ {
        var v21 uint64 = (v2 + (v2 + (v2 * 183)))
        _ = fi20
        _ = v21
    }
    var v22 uint16 = genFn1(42, 42)
    for fi23 := 0; fi23 < 1; fi23++ {
        v3 += v3
        var v24 string = "foo bar"
        v24 = (v24 + "foo bar")
        _ = fi23
        _ = v24
    }
    var v25 uint8 = 1
    v26 := []int{}
    switch -31 {
    case -338:
        fmt.Println(v3, v2, v22)
    case 9223372036854775807:
        var v27 uint8 = (v25 >> ((2) & 63))
        _ = v27
    default:
        var v28 float32 = 3.14159
        _ = v28
    }
    if len(v26) > 0 {
        v26[0] = 32767
    }
    var v29 bool = false
    for fi30 := 0; fi30 < 11; fi30++ {
        var v31 float64 = 0.1
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 5; fi32++ {
        delete(v5, 2147483648)
        for fi33 := 0; fi33 < 6; fi33++ {
            v22 = v22
            fmt.Println(v3)
            _ = fi33
        }
        _ = fi32
    }
    var v34 uint8 = v25
    switch 127 {
    case 2147483647:
        for fi35 := 0; fi35 < 12; fi35++ {
            var v36 int16 = 1000
            fmt.Println(v29, v34)
            var v37 uint16 = 32767
            _ = fi35
            _ = v36
            _ = v37
        }
    case 32768:
        var v38 int16 = 1000
        _ = v38
    case 4294967295:
        _ = len(v26)
    }
    v2 -= (^(v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
