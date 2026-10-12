package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 uint; F1 uint }
func genFn0(p0 uint64, p1 string, p2 int64) bool {
    v6 := GenS1{F0: 655}
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 int32 = 1
        _ = fi7
        _ = v8
    }
    if (p0 == 0) {
        v9 := GenS2{F0: 0, F1: p0}
        v10 := []int{1, 63, 32767}
        _ = v10
        _ = v9
    }
    var v11 int = -9223372036854775808
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v6
    return (((1 == uint(p0)) && (true || (!(true && ((!(!(!false))) || (0.1 > 1.5)))))) || ((!false) && (!true)))
}
func genFn1(p0 uint8) uint8 {
    p0 += 26
    var v12 float64 = 3.14159
    var v13 int16 = 32767
    v13 = 32767
    _ = p0
    _ = v12
    _ = v13
    return (p0 * p0)
}

func main() {
    var v1 uint = 1
    var v2 uint8 = 127
    var v3 uint = 2
    v4 := make(map[string]uint32)
    v5 := []uint32{}
    for fi14 := 0; fi14 < 3; fi14++ {
        var v15 bool = false
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: 42, F1: 1}
    v17 := v4["\n\t"]
    if ("foo bar" > "foo bar") {
        switch 1 {
        case -9223372036854775808:
            v18 := GenS2{F0: 0, F1: 1}
            _ = v18
        case 0:
            var v19 uint8 = (v2 % (((v2 ^ (^(v2)))) | 1))
            _ = v19
        case 510:
            for fi20 := 0; fi20 < 13; fi20++ {
                v3 = v1
                var v21 uint8 = (v2 * 0)
                _ = fi20
                _ = v21
            }
        }
        fmt.Println(v17, v3)
    } else {
        v3 += (^(v3))
        if true {
            fmt.Println(v1)
        }
        v1 = 1
    }
    if ((v2 % (((^(v2))) | 1)) <= (v2 >> ((42) & 63))) {
        var v22 uint = v3
        var v23 bool = (0 > v17)
        _ = v22
        _ = v23
    } else {
        for fi24 := 0; fi24 < 20; fi24++ {
            v25 := GenS3{F0: 43, F1: -1}
            var v26 uint64 = 18446744073709551615
            v17 = v17
            _ = fi24
            _ = v25
            _ = v26
        }
        if true {
            var v27 float32 = -2.25
            var v28 int8 = 1
            v29 := GenS1{F0: 2147483647}
            _ = v27
            _ = v28
            _ = v29
        }
        v30 := v4[("" + "a")]
        _ = v30
    }
    if (false || (true && false)) {
        for fi31 := 0; fi31 < 9; fi31++ {
            v32 := GenS1{F0: 2147483647}
            _ = fi31
            _ = v32
        }
        fmt.Println(v2)
    } else {
        for fi33 := 0; fi33 < 1; fi33++ {
            fmt.Println(v2, v2)
            for fi34 := 0; fi34 < 10; fi34++ {
                v1 -= (v1 ^ (v1 ^ (v1 >> ((0) & 63))))
                v3 = (v1 | v1)
                v35 := GenS4{F0: 65535, F1: 0}
                fmt.Println(v1)
                var v36 uint = v1
                fmt.Println(v2)
                _ = fi34
                _ = v35
                _ = v36
            }
            _ = fi33
        }
    }
    var v37 uint64 = 598
    for fi38 := 0; fi38 < 4; fi38++ {
        var v39 int32 = -1
        _ = fi38
        _ = v39
    }
    switch -65 {
    case 128:
        fmt.Println(v1, v2, v2)
    case 936:
        var v40 float64 = (-2.25 / (-2.25 - 3.14159))
        _ = v40
    default:
        var v41 uint16 = 148
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
