package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32 }
func genFn0(p0 int8, p1 string) uint8 {
    v4 := GenS1{F0: 18, F1: -399, F2: 42}
    p0 += -1
    _ = p0
    _ = p1
    _ = v4
    return 215
}
func genFn1(p0 int16, p1 int8, p2 int64) uint {
    v5 := []int{317}
    for fi6 := 0; fi6 < 11; fi6++ {
        p0 = p0
        _ = fi6
    }
    p0 = p0
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: 255, F1: p0, F2: 138}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 212
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]bool)
    var v3 uint32 = 65535
    for fi9 := 0; fi9 < 11; fi9++ {
        v1 += v1
        _ = fi9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        v1 += v1
        v11 := GenS1{F0: 1, F1: 41, F2: 255}
        v3 = v3
        _ = fi10
        _ = v11
    }
    if (!(false || true)) {
        for fi12 := 0; fi12 < 12; fi12++ {
            if (-2.25 >= -2.25) {
                v3 -= (v3 + v3)
            }
            _ = fi12
        }
        switch -9223372036854775808 {
        case -910:
            var v13 uint32 = 2147483647
            _ = v13
        }
        for fi14 := 0; fi14 < 13; fi14++ {
            v1 = v1
            var v15 float64 = 1e10
            _ = fi14
            _ = v15
        }
    }
    v3 += 4294967295
    v16 := GenS1{F0: uint8(v1), F1: -32768, F2: uint8(v3)}
    v17 := GenS2{F0: ("" + (("a" + "") + "foo bar"))}
    var v18 uint64 = 620
    var v19 int64 = -527
    fmt.Println(v19, v18)
    for fi20 := 0; fi20 < 5; fi20++ {
        fmt.Println(v3, fi20, v19)
        _ = fi20
    }
    fmt.Println(v19, v3)
    v21 := GenS2{F0: ""}
    var v22 float32 = ((float32((0.0 / 1.0)) * 0.0) * ((1.5 / 1e10) - ((0.1 + 1e10) - 0.0)))
    v23 := []int{2147483647, 127}
    v24 := GenS3{F0: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
