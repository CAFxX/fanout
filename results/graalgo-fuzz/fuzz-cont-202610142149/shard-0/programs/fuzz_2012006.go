package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint }
func genFn0(p0 int8, p1 uint16, p2 uint16) int {
    fmt.Println(p1, p1)
    for fi6 := 0; fi6 < 19; fi6++ {
        p0 = (p0 - (p0 & (p0 & p0)))
        p0 = p0
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 uint64 = 2
            fmt.Println(p2, v8)
            fmt.Println(p2, v8, p1)
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    return int(p0)
}

func main() {
    var v1 uint64 = 62
    v2 := make(map[string]bool)
    var v3 int = -2147483648
    v4 := []int{}
    var v5 bool = false
    for fi9 := 0; fi9 < 5; fi9++ {
        v3 = v3
        var v10 int16 = 1000
        if len(v4) > 0 {
            v4[0] = genFn0(-128, 32767, 1)
        }
        _ = fi9
        _ = v10
    }
    if (v5 || v5) {
        for fi11 := 0; fi11 < 20; fi11++ {
            for fi12 := 0; fi12 < 1; fi12++ {
                var v13 float64 = -2.25
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 uint16 = 1
                switch "\n\t" {
                case "hello":
                    v16 := GenS1{F0: v5}
                    _ = v16
                }
                v15 = (v15 / (((v15 + uint16(v1))) | 1))
                _ = fi14
                _ = v15
            }
            var v17 float64 = (3.14159 + 1e10)
            _ = fi11
            _ = v17
        }
        for fi18 := 0; fi18 < 17; fi18++ {
            if false {
                var v19 float64 = (1e10 + float64(float32((-2.25 + 0.0))))
                v20 := GenS2{F0: 0, F1: 4294967295}
                _ = v19
                _ = v20
            }
            _ = fi18
        }
    } else {
        if ((1e10 * -2.25) != -2.25) {
            var v21 uint16 = 32767
            _ = v21
        }
        switch "a" {
        case "foo bar":
            var v22 int32 = 42
            _ = v22
        case "":
            for fi23 := 0; fi23 < 8; fi23++ {
                delete(v2, "hello")
                v4 = append(v4, -9223372036854775808)
                fmt.Println(v3, v1)
                _ = fi23
            }
        }
    }
    v1 += v1
    v3 = (v3 ^ v3)
    v24 := GenS2{F0: 877, F1: 1}
    for fi25 := 0; fi25 < 6; fi25++ {
        for fi26 := 0; fi26 < 19; fi26++ {
            v27 := GenS1{F0: false}
            v3 = (v3 * 127)
            _ = fi26
            _ = v27
        }
        v5 = v5
        _ = fi25
    }
    var v28 uint16 = 32767
    v29 := GenS3{F0: 65535, F1: 53, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
