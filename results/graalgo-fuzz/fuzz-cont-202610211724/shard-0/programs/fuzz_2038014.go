package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 bool }
type GenS2 struct { F0 int8; F1 string; F2 uint64 }
func genFn0() uint8 {
    switch 4294967295 {
    case 42:
        v7 := []int{32768, 128, 410}
        _ = v7
    case 584:
        v8 := GenS1{F0: 127, F1: 554, F2: (65535 <= 701)}
        _ = v8
    case -1:
        var v9 uint16 = 42
        _ = v9
    default:
        v10 := GenS2{F0: 42, F1: "\n\t", F2: 1}
        _ = v10
    }
    v11 := GenS1{F0: 128, F1: -2147483648, F2: true}
    var v12 int32 = -580
    _ = v11
    _ = v12
    return 158
}
func genFn1(p0 uint8, p1 uint64) int64 {
    for fi13 := 0; fi13 < 17; fi13++ {
        if true {
            fmt.Println(fi13, p1, p1)
            for fi14 := 0; fi14 < 18; fi14++ {
                var v15 int8 = 0
                var v16 int = 32767
                v15 = v15
                _ = fi14
                _ = v15
                _ = v16
            }
        } else {
            v17 := []int{902}
            p0 -= (^(p0))
            _ = v17
        }
        _ = fi13
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        for fi19 := 0; fi19 < 11; fi19++ {
            if (-1 <= 100) {
                var v20 int16 = 1000
                v20 -= (-(v20))
                _ = v20
            }
            var v21 uint = 2
            _ = fi19
            _ = v21
        }
        v22 := GenS2{F0: 100, F1: "\n\t", F2: p1}
        _ = fi18
        _ = v22
    }
    p0 = genFn0()
    _ = p0
    _ = p1
    return int64(p0)
}

func main() {
    var v1 float32 = (-2.25 + 1e10)
    v2 := make(map[string]string)
    var v3 uint = 2
    v4 := make(map[int]int)
    var v5 uint8 = 1
    var v6 int32 = -1
    v23 := GenS2{F0: 1, F1: ("a" + "foo bar"), F2: 0}
    v1 -= 0.1
    for fi24 := 0; fi24 < 8; fi24++ {
        var v25 string = "\n\t"
        var v26 uint = v3
        _ = fi24
        _ = v25
        _ = v26
    }
    v5 = (v5 + genFn0())
    v3 = v3
    var v27 uint = v3
    v28 := v4[-283]
    var v29 float64 = 1.5
    if (v29 != 1e10) {
        v30 := GenS2{F0: 0, F1: (("foo bar" + "a") + "a"), F2: 4294967295}
        var v31 uint32 = 0
        var v32 int16 = -32768
        _ = v30
        _ = v31
        _ = v32
    }
    v33 := []int{v28, (^(v28))}
    v28 -= v28
    var v34 int64 = int64(v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
