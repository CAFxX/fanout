package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64; F1 uint32 }
func genFn0(p0 string, p1 int16, p2 bool) bool {
    v7 := GenS1{F0: p0, F1: p1}
    for fi8 := 0; fi8 < 18; fi8++ {
        fmt.Println(p0)
        v9 := []int{(fi8 & fi8), 1, -9223372036854775808}
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: 42}
    v11 := GenS3{F0: 255, F1: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v7
    return true
}
func genFn1(p0 int32, p1 uint8) uint {
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 uint64 = 0
        _ = fi12
        _ = v13
    }
    p1 = p1
    p1 -= p1
    v14 := []int{127, 0}
    _ = p0
    _ = p1
    _ = v14
    return 372
}

func main() {
    var v1 uint64 = 42
    var v2 int8 = 39
    var v3 bool = ((!(false || false)) || true)
    var v4 uint16 = 663
    var v5 float32 = -2.25
    v6 := make(map[string]uint32)
    var v15 uint = 989
    switch (("hello" + "a") + "foo bar") {
    case "\n\t":
        if v3 {
            var v16 int16 = 1
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        v18 := []int{(fi17 - fi17), 32767, (fi17 >> ((genFn1(-988, 183)) & 63)), (fi17 / ((603) | 1))}
        var v19 uint16 = (v4 | v4)
        _ = fi17
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        v15 -= v15
        _ = fi20
    }
    v21 := GenS4{F0: v2}
    v22 := []int{9223372036854775807}
    v1 = uint64(v15)
    if v3 {
        v23 := GenS5{F0: (v1 - 18446744073709551615), F1: 0}
        var v24 int8 = (v2 + 93)
        _ = v23
        _ = v24
    } else {
        var v25 uint32 = 4294967295
        if (177 <= uint8(v15)) {
            if len(v22) > 0 {
                v22[0] = -2147483648
            }
            if v3 {
                switch "" {
                case "\n\t":
                    _ = len(v22)
                case "a":
                    fmt.Println(v5, v2, v2)
                }
                var v26 uint32 = (^(v25))
                _ = v26
            } else {
                fmt.Println(v4)
            }
        } else {
            fmt.Println(v4, v2, v15)
        }
        _ = v25
    }
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    var v27 int64 = -9223372036854775808
    for fi28 := 0; fi28 < 11; fi28++ {
        var v29 int32 = 861
        _ = fi28
        _ = v29
    }
    var v30 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
