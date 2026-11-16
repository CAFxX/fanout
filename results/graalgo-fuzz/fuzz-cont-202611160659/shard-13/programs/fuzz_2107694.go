package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 string }
func genFn0(p0 uint8, p1 uint16) bool {
    v5 := []int{118, 127, -901, 32767}
    v6 := GenS1{F0: 0}
    fmt.Println(p1, p0)
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            switch fi7 {
            case 32768:
                p1 = p1
            case 0:
                _ = len(v5)
            case -1:
                if false {
                    p0 = (p0 + p0)
                } else {
                    var v9 uint64 = 0
                    p0 = p0
                    _ = v9
                }
            default:
                var v10 uint8 = p0
                _ = v10
            }
            _ = fi8
        }
        var v11 uint16 = p1
        _ = fi7
        _ = v11
    }
    if ((((1e10 + (0.0 + 0.0)) / 1e10) + float32(float64((-2.25 - float32(0.1))))) != 0.0) {
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 int16 = 32767
            var v14 uint = 253
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 int8 = -35
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return false
}

func main() {
    var v1 float64 = 1.5
    v2 := make(map[uint32]bool)
    var v3 float32 = float32(0.1)
    v4 := make(map[string]int)
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 uint8 = 0
        if (false && true) {
            v18 := GenS2{F0: 35, F1: "foo bar"}
            _ = v18
        } else {
            v3 = v3
            v3 -= ((1e10 - ((v3 - float32(v1)) * 1.5)) / 1e10)
        }
        _ = fi16
        _ = v17
    }
    v1 += v1
    fmt.Println(v1)
    v19 := []int{4294967295, -2147483648}
    var v20 uint64 = 659
    fmt.Println(v20)
    var v21 uint16 = 0
    v22 := GenS2{F0: 128, F1: ""}
    if true {
        var v23 int = 0
        _ = v23
    }
    switch "hello" {
    case "foo bar":
        var v24 uint8 = 0
        _ = v24
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        v20 += v20
        _ = fi25
    }
    var v26 uint8 = 228
    var v27 int8 = 127
    v20 = v20
    v28 := v2[640]
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
