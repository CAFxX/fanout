package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 uint64 }
type GenS5 struct { F0 uint32; F1 uint8; F2 uint64 }
func genFn0(p0 uint32) uint32 {
    if (!true) {
        p0 += 4294967295
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 int8 = 127
            v7 = (v7 ^ (v7 << ((0) & 63)))
            if (!(!((36 == (v7 >> ((2) & 63))) || true))) {
                v7 -= (v7 << ((4294967295) & 63))
                p0 += 1
            } else {
                for fi8 := 0; fi8 < 16; fi8++ {
                    v9 := GenS1{F0: 0, F1: 0, F2: 1}
                    v10 := []int{(^(fi6)), (fi6 / ((fi8) | 1)), (fi6 >> ((4294967295) & 63)), 127}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                p0 += (p0 ^ p0)
                for fi11 := 0; fi11 < 1; fi11++ {
                    var v12 bool = false
                    fmt.Println(p0)
                    _ = fi11
                    _ = v12
                }
            }
            _ = fi6
            _ = v7
        }
        var v13 int64 = -367
        _ = v13
    } else {
        v14 := GenS2{F0: 818}
        fmt.Println(p0)
        _ = v14
    }
    if false {
        switch 122 {
        case 32768:
            v15 := []int{}
            _ = v15
        }
        v16 := []int{9223372036854775807, 32767, 1, -2147483648}
        _ = v16
    }
    _ = p0
    return p0
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 2
    v3 := make(map[uint32]bool)
    v4 := []uint32{(v2 - (v2 - (v2 + v2))), 0}
    var v5 float32 = 1.5
    var v17 int = 4294967295
    switch ("foo bar" + ((((("hello" + "foo bar") + "") + "foo bar") + ("a" + "hello")) + ("\n\t" + ("a" + "hello")))) {
    case "foo bar":
        v5 -= v5
    case "\n\t":
        var v18 int = v17
        _ = v18
    default:
        fmt.Println(v2, v2)
    }
    if false {
        v19 := GenS3{F0: 14}
        _ = v19
    }
    fmt.Println(v2, v17)
    if ((2147483647 < 1) || ((-(v17)) >= (v17 << ((1) & 63)))) {
        if len(v4) > 0 {
            v4[0] = v2
        }
    } else {
        v2 = genFn0(genFn0((v2 + v2)))
        for fi20 := 0; fi20 < 1; fi20++ {
            var v21 int8 = -1
            var v22 float64 = ((0.0 / 1.0) + 1e10)
            for fi23 := 0; fi23 < 16; fi23++ {
                var v24 float64 = v22
                var v25 bool = ((!true) && (!((v21 * 127) > v21)))
                var v26 uint8 = uint8(v2)
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
            _ = fi20
            _ = v21
            _ = v22
        }
        switch (v17 / ((v17) | 1)) {
        case -928:
            var v27 int32 = 75
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 1; fi28++ {
        if false {
            v29 := GenS4{F0: 100, F1: 0}
            _ = v29
        } else {
            var v30 uint16 = 32767
            _ = v30
        }
        _ = fi28
    }
    v5 -= 1.5
    v31 := GenS5{F0: genFn0(v2), F1: 49, F2: 1}
    for fi32 := 0; fi32 < 18; fi32++ {
        switch v17 {
        case 0:
            v33 := v1["hello"]
            _ = v33
        case -9223372036854775808:
            var v34 int16 = 25
            _ = v34
        case -21:
            fmt.Println(v2)
        }
        _ = fi32
    }
    v17 = (v17 | 207)
    _ = len(v4)
    v17 -= v17
    var v35 int16 = -708
    v36 := GenS1{F0: 0, F1: 321, F2: 42}
    switch v17 {
    case -1:
        var v37 int64 = 42
        _ = v37
    }
    v38 := GenS3{F0: -128}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v38
    _ = v4
    _ = v5
}
