package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int }
type GenS2 struct { F0 string; F1 int64; F2 uint }
type GenS3 struct { F0 uint8; F1 uint32 }
func genFn0() uint16 {
    var v6 int = 32767
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := []int{32767, 32768, -1}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(fi9, v6)
        fmt.Println(v6, v6)
        _ = fi9
    }
    var v10 uint8 = 0
    _ = v10
    _ = v6
    return 673
}
func genFn1(p0 bool, p1 string, p2 uint16) int16 {
    fmt.Println(p0, p2, p1)
    var v11 string = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 32767
}

func main() {
    var v1 int16 = 42
    var v2 uint8 = 128
    v3 := make(map[string]uint32)
    var v4 float32 = 1.5
    var v5 int16 = 0
    v12 := GenS1{F0: 18446744073709551615, F1: 128, F2: 9223372036854775807}
    v2 = v2
    v1 = 0
    v5 -= v1
    v13 := GenS2{F0: "", F1: -9223372036854775808, F2: 42}
    fmt.Println(v2)
    fmt.Println(v1, v5, v5)
    if (0 > 2) {
        for fi14 := 0; fi14 < 9; fi14++ {
            v2 = 128
            _ = fi14
        }
        v15 := GenS3{F0: v2, F1: 2147483648}
        _ = v15
    }
    v5 = 0
    v2 = (v2 - (v2 / ((128) | 1)))
    v5 = v5
    for fi16 := 0; fi16 < 2; fi16++ {
        switch (fi16 & -544) {
        case 55:
            switch (fi16 - fi16) {
            case 734:
                var v17 int8 = -1
                _ = v17
            case 9223372036854775807:
                v5 = v5
            }
        case 1:
            v5 -= -32768
        }
        for fi18 := 0; fi18 < 15; fi18++ {
            v5 = (v5 | v1)
            fmt.Println(fi18, v5)
            var v19 uint64 = 4294967295
            _ = fi18
            _ = v19
        }
        _ = fi16
    }
    switch (("\n\t" + "\n\t") + "a") {
    case "":
        for fi20 := 0; fi20 < 18; fi20++ {
            for fi21 := 0; fi21 < 7; fi21++ {
                var v22 float32 = v4
                _ = fi21
                _ = v22
            }
            v1 -= (^(v1))
            fmt.Println(v1, v1, v5)
            _ = fi20
        }
    case "a":
        for fi23 := 0; fi23 < 10; fi23++ {
            v3[("a" + "hello")] = 895
            _ = fi23
        }
    default:
        var v24 uint64 = 1
        _ = v24
    }
    v2 = v2
    var v25 int = 945
    for fi26 := 0; fi26 < 10; fi26++ {
        fmt.Println(v1, v2, v5)
        var v27 uint = 4294967295
        delete(v3, (("foo bar" + "foo bar") + "hello"))
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
