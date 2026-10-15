package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint8 }
func genFn0() uint64 {
    var v6 uint8 = 42
    v6 -= (v6 % ((0) | 1))
    if ((71 < 100) && false) {
        var v7 uint8 = v6
        _ = v7
    } else {
        var v8 uint = uint(v6)
        _ = v8
    }
    v9 := []int{1, 0, 32767, -300}
    _ = v6
    _ = v9
    return 134
}
func genFn1() int32 {
    switch ("" + "a") {
    case "":
        var v10 uint16 = 0
        _ = v10
    default:
        var v11 uint64 = 64
        _ = v11
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        if false {
            var v13 bool = (2 < genFn0())
            var v14 float64 = 0.0
            _ = v13
            _ = v14
        } else {
            var v15 uint8 = 90
            var v16 int64 = 0
            _ = v15
            _ = v16
        }
        var v17 uint16 = 32767
        _ = fi12
        _ = v17
    }
    return -1
}

func main() {
    var v1 int = 32767
    var v2 int16 = 1000
    var v3 float64 = 0.0
    v4 := []string{"foo bar", (((("" + "hello") + "") + "\n\t") + "a"), "", "hello"}
    v5 := make(map[uint32]uint32)
    if true {
        fmt.Println(v2, v2, v3)
        for fi18 := 0; fi18 < 14; fi18++ {
            v19 := GenS1{F0: (v1 * v1), F1: -9223372036854775808}
            var v20 float64 = v3
            for fi21 := 0; fi21 < 2; fi21++ {
                v22 := GenS1{F0: v1, F1: 1}
                _ = fi21
                _ = v22
            }
            _ = fi18
            _ = v19
            _ = v20
        }
        var v23 int8 = 1
        _ = v23
    } else {
        v24 := GenS1{F0: 9223372036854775807, F1: -1}
        var v25 int = v1
        _ = v24
        _ = v25
    }
    _ = len(v5)
    var v26 uint16 = 50
    if len(v4) > 0 {
        v4[0] = "a"
    }
    var v27 float32 = (0.1 - ((1.5 + -2.25) - -2.25))
    for fi28 := 0; fi28 < 19; fi28++ {
        for fi29 := 0; fi29 < 12; fi29++ {
            v30 := GenS1{F0: (v1 >> ((20) & 63)), F1: 1}
            var v31 bool = (false && true)
            var v32 float32 = float32(((v3 + v3) / ((v3 / (float64((v27 * v27)) * v3)) / v3)))
            _ = fi29
            _ = v30
            _ = v31
            _ = v32
        }
        var v33 int16 = (v2 * (v2 % ((1000) | 1)))
        v34 := GenS2{F0: 65535, F1: 87}
        _ = fi28
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 17; fi35++ {
        v2 = int16(fi35)
        _ = fi35
    }
    v1 = v1
    var v36 int64 = -1
    v37 := v5[2]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
