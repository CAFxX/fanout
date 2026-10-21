package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 uint16; F2 int }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS7 struct { F0 int16 }
func genFn0(p0 int64, p1 int8, p2 int32) uint64 {
    v6 := GenS1{F0: 0, F1: 73}
    p1 += (p1 | 83)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 144
}
func genFn1(p0 int64, p1 int64) uint {
    var v7 uint = 4294967295
    v8 := GenS2{F0: -128, F1: -2147483648}
    v9 := GenS2{F0: -19, F1: 42}
    v10 := GenS1{F0: 32767, F1: 1}
    fmt.Println(p1, v7, p0)
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return v7
}

func main() {
    v1 := []int{-228}
    var v2 uint64 = 4294967295
    var v3 float64 = (-2.25 - 0.0)
    v4 := make(map[string]string)
    v5 := make(map[int]string)
    var v11 int16 = -263
    v12 := GenS3{F0: ""}
    if (!(true && true)) {
        v3 = 0.0
        v4[(("a" + "") + ("" + "foo bar"))] = "a"
    } else {
        if false {
            v13 := GenS2{F0: 127, F1: 42}
            for fi14 := 0; fi14 < 9; fi14++ {
                fmt.Println(fi14)
                _ = fi14
            }
            v15 := GenS3{F0: ""}
            _ = v13
            _ = v15
        }
        var v16 uint = 0
        var v17 int32 = 0
        _ = v16
        _ = v17
    }
    switch "foo bar" {
    case "hello":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v3 = v3
    for fi18 := 0; fi18 < 9; fi18++ {
        switch (fi18 >> ((18446744073709551615) & 63)) {
        case 1:
            v19 := GenS4{F0: "a", F1: 0, F2: fi18}
            _ = v19
        default:
            fmt.Println(v2)
        }
        var v20 uint8 = 127
        fmt.Println(v2, v20, v11)
        _ = fi18
        _ = v20
    }
    var v21 float64 = v3
    v22 := GenS3{F0: ""}
    if (!false) {
        fmt.Println(v21)
    } else {
        var v23 float64 = v3
        for fi24 := 0; fi24 < 11; fi24++ {
            for fi25 := 0; fi25 < 2; fi25++ {
                var v26 int16 = v11
                _ = fi25
                _ = v26
            }
            _ = fi24
        }
        _ = v23
    }
    v27 := GenS1{F0: 2, F1: 46}
    v28 := GenS1{F0: 2, F1: 2}
    var v29 bool = true
    var v30 int64 = 42
    for fi31 := 0; fi31 < 5; fi31++ {
        v32 := GenS2{F0: 42, F1: -2147483648}
        v33 := GenS5{F0: 54, F1: 2}
        var v34 uint32 = 963
        _ = fi31
        _ = v32
        _ = v33
        _ = v34
    }
    if true {
        if (-782 <= 9223372036854775807) {
            var v35 uint16 = 526
            var v36 float64 = 1.5
            v3 -= (float64((-2.25 / 1.0)) - v36)
            _ = v35
            _ = v36
        } else {
            var v37 uint32 = 2147483648
            v37 = 2147483647
            for fi38 := 0; fi38 < 4; fi38++ {
                v39 := GenS6{F0: (-(v30)), F1: (^(v37)), F2: (v2 ^ (v2 * genFn0(v30, 1, -159)))}
                v40 := GenS7{F0: v11}
                v41 := GenS5{F0: v2, F1: 593}
                _ = fi38
                _ = v39
                _ = v40
                _ = v41
            }
            _ = v37
        }
        v42 := GenS5{F0: v2, F1: 458}
        v1 = append(v1, 128)
        _ = v42
    }
    var v43 int16 = v11
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v43
    _ = v5
}
