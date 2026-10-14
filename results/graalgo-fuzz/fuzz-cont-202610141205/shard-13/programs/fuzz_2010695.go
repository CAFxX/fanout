package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int64; F1 bool; F2 uint }
func genFn0() int64 {
    if (0.1 <= -2.25) {
        var v5 float64 = 1e10
        var v6 uint = 1
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int16 = 570
        if true {
            v8 = (v8 ^ (v8 + (^(v8))))
            var v9 int16 = v8
            _ = v9
        } else {
            if (1 <= 4294967295) {
                v10 := GenS1{F0: 656}
                switch (-(fi7)) {
                case 42:
                    v8 = v8
                case -1:
                    var v11 int8 = -101
                    _ = v11
                case -9223372036854775808:
                    v12 := []int{(fi7 * -244)}
                    _ = v12
                }
                v13 := []int{(fi7 >> ((904) & 63))}
                _ = v10
                _ = v13
            } else {
                var v14 int = fi7
                var v15 uint32 = 399
                _ = v14
                _ = v15
            }
        }
        switch 127 {
        case 42:
            if true {
                var v16 uint32 = 65535
                v8 = v8
                _ = v16
            } else {
                fmt.Println(fi7, fi7)
                v8 += (v8 * (v8 & (v8 / ((348) | 1))))
            }
        case 32767:
            fmt.Println(fi7)
        case 128:
            v17 := GenS1{F0: -289}
            _ = v17
        }
        _ = fi7
        _ = v8
    }
    v18 := GenS2{F0: true, F1: -345}
    var v19 int8 = 127
    v19 = v19
    _ = v18
    _ = v19
    return 958
}
func genFn1(p0 int16) int8 {
    switch -1 {
    case -37:
        p0 = (p0 * (p0 | 42))
    case -204:
        v20 := GenS2{F0: ((42 >= p0) || true), F1: -469}
        _ = v20
    case 65:
        v21 := []int{32768, 42, 127, -2147483648}
        _ = v21
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        if (1000 == 1) {
            fmt.Println(p0, p0)
            p0 = p0
            fmt.Println(fi22)
        }
        _ = fi22
    }
    if true {
        var v23 int8 = -115
        fmt.Println(p0, v23)
        _ = v23
    }
    if false {
        var v24 float64 = 0.1
        _ = v24
    }
    _ = p0
    return 127
}

func main() {
    var v1 float64 = (0.0 / float64(1.5))
    v2 := []int{0, -2147483648, -694}
    v3 := make(map[int]string)
    var v4 bool = true
    v25 := GenS3{F0: -1, F1: (-1 < 13), F2: 296}
    if (246 >= 267) {
        fmt.Println(v1, v4)
        fmt.Println(v4, v4)
        v1 += v1
    }
    v26 := GenS2{F0: v4, F1: 9223372036854775807}
    v3[127] = "foo bar"
    _ = len(v3)
    fmt.Println(v4, v4)
    v27 := GenS3{F0: 9223372036854775807, F1: v4, F2: 49}
    v4 = (!(795 >= 4294967295))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
