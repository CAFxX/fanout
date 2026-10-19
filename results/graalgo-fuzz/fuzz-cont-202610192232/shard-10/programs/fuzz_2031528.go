package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int16 }
type GenS2 struct { F0 string; F1 uint8; F2 bool }
type GenS3 struct { F0 int16; F1 int64; F2 bool }
func genFn0(p0 uint8, p1 int8) bool {
    fmt.Println(p0, p0)
    v7 := GenS1{F0: 255, F1: true, F2: 42}
    _ = p0
    _ = p1
    _ = v7
    return true
}
func genFn1(p0 uint) bool {
    if (947 > 1) {
        p0 = (p0 * (p0 % ((p0) | 1)))
        var v8 float64 = (float64(-2.25) * (1e10 * 3.14159))
        fmt.Println(p0, p0)
        _ = v8
    }
    if ((-1 != 42) || (!true)) {
        p0 -= p0
        var v9 int32 = -2147483648
        _ = v9
    } else {
        v10 := []int{1, -635}
        fmt.Println(p0)
        var v11 int64 = -275
        _ = v10
        _ = v11
    }
    _ = p0
    return false
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[int]string)
    var v3 uint8 = 170
    v4 := []string{("" + ("foo bar" + "")), "hello", "\n\t", "foo bar"}
    var v5 int16 = 126
    var v6 uint = 4294967295
    switch 4294967295 {
    case -499:
        fmt.Println(v3, v3, v3)
    case 9223372036854775807:
        v12 := GenS2{F0: "foo bar", F1: v3, F2: false}
        _ = v12
    case -840:
        _ = len(v4)
    default:
        fmt.Println(v3)
    }
    var v13 float32 = 1e10
    v14 := v2[int(v5)]
    var v15 bool = false
    v13 -= (1e10 / v13)
    v16 := GenS2{F0: v14, F1: (v3 | 38), F2: false}
    for fi17 := 0; fi17 < 18; fi17++ {
        v14 = v14
        _ = fi17
    }
    var v18 int64 = 0
    _ = len(v4)
    if v15 {
        fmt.Println(v1, v14)
        var v19 int = 42
        switch (v19 / (((v19 + v19)) | 1)) {
        case 32767:
            v20 := GenS2{F0: "foo bar", F1: v3, F2: v15}
            _ = v20
        }
        _ = v19
    } else {
        delete(v2, 431)
        for fi21 := 0; fi21 < 13; fi21++ {
            for fi22 := 0; fi22 < 14; fi22++ {
                var v23 int = fi21
                fmt.Println(v3)
                v23 += v23
                _ = fi22
                _ = v23
            }
            if v15 {
                var v24 string = ((v14 + "\n\t") + (v14 + v14))
                _ = v24
            }
            v25 := GenS3{F0: (v5 * int16(v6)), F1: v18, F2: (v14 >= "hello")}
            _ = fi21
            _ = v25
        }
        v15 = v15
    }
    v1 -= (v1 / (((v1 | (v1 ^ v1))) | 1))
    fmt.Println(v14, v15, v3)
    fmt.Println(v6)
    if (false && v15) {
        if v15 {
            switch "" {
            case "\n\t":
                var v26 uint64 = 42
                _ = v26
            }
            v3 -= (^(v3))
        } else {
            if (977 != 613) {
                _ = len(v2)
            }
            for fi27 := 0; fi27 < 7; fi27++ {
                delete(v2, fi27)
                _ = fi27
            }
            var v28 uint32 = 2
            _ = v28
        }
        var v29 uint16 = v1
        var v30 int16 = 89
        _ = v29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
