package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint8 }
func genFn0() uint8 {
    if (0 >= 1) {
        switch "a" {
        case "":
            var v7 float32 = 1.5
            _ = v7
        default:
            var v8 string = ("foo bar" + "\n\t")
            _ = v8
        }
    }
    v9 := []int{42, 9223372036854775807}
    var v10 uint64 = 42
    var v11 uint32 = 2
    _ = v10
    _ = v11
    _ = v9
    return 0
}
func genFn1(p0 uint8) uint {
    v12 := GenS1{F0: (^(p0)), F1: 4294967295}
    var v13 int32 = 2147483647
    v13 += (v13 % ((v13) | 1))
    _ = p0
    _ = v12
    _ = v13
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]int)
    var v3 float32 = 1e10
    var v4 uint16 = 2
    v5 := make(map[int]uint32)
    var v6 float32 = float32((1.5 * (float64(v3) - 1e10)))
    fmt.Println(v4, v3, v3)
    for fi14 := 0; fi14 < 16; fi14++ {
        v3 += 0.1
        fmt.Println(v4, v6, fi14)
        var v15 int64 = -45
        _ = fi14
        _ = v15
    }
    fmt.Println(v6, v4, v3)
    switch "foo bar" {
    case "a":
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 uint = 4294967295
            _ = fi16
            _ = v17
        }
    case "\n\t":
        switch "\n\t" {
        case "a":
            v6 += v6
        case "":
            if false {
                v18 := GenS1{F0: 128, F1: 0}
                var v19 int16 = -32768
                var v20 uint16 = (^(v4))
                _ = v18
                _ = v19
                _ = v20
            }
        }
    case "hello":
        v3 = v6
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        v22 := []int{}
        var v23 uint = 4294967295
        _ = fi21
        _ = v22
        _ = v23
    }
    if (((-128 <= -74) || false) || (0.1 <= v3)) {
        var v24 float64 = 0.1
        v25 := GenS1{F0: 114, F1: 691}
        var v26 string = ("foo bar" + (("" + "a") + ("" + "hello")))
        _ = v24
        _ = v25
        _ = v26
    } else {
        var v27 int8 = 0
        _ = v27
    }
    fmt.Println(v3, v6, v3)
    v4 = (v4 + (v4 & v4))
    if (73 == 131) {
        v1["hello"] = 989
    }
    v28 := GenS1{F0: 75, F1: 468}
    if (v3 < (-2.25 - -2.25)) {
        var v29 bool = true
        var v30 float64 = (1.5 * -2.25)
        var v31 uint32 = 42
        _ = v29
        _ = v30
        _ = v31
    } else {
        v3 += v6
        v32 := GenS1{F0: 128, F1: 18446744073709551615}
        _ = v32
    }
    v4 += (^(v4))
    v33 := []int{42}
    fmt.Println(v3)
    for fi34 := 0; fi34 < 12; fi34++ {
        for fi35 := 0; fi35 < 10; fi35++ {
            switch 848 {
            case 128:
                var v36 uint64 = 1
                _ = v36
            case 32767:
                v37 := GenS1{F0: 2, F1: 2}
                _ = v37
            case -1:
                v38 := v5[(-(fi34))]
                _ = v38
            default:
                v3 = v6
            }
            _ = fi35
        }
        var v39 string = "foo bar"
        if false {
            if (!(!true)) {
                var v40 float64 = 0.0
                _ = v40
            }
            v4 = (v4 - (v4 | v4))
            v3 -= v3
        } else {
            var v41 float32 = (float32(0.1) - (((((v6 * -2.25) * v6) * 0.0) + 1e10) / v3))
            v41 = float32(((0.1 * 1.5) * 1.5))
            v42 := GenS2{F0: 73}
            _ = v41
            _ = v42
        }
        _ = fi34
        _ = v39
    }
    var v43 uint64 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
