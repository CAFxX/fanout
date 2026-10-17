package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 int8; F2 int16 }
type GenS4 struct { F0 bool; F1 int64; F2 int16 }
type GenS5 struct { F0 uint; F1 string }
func genFn0(p0 uint16, p1 string, p2 string) uint16 {
    p2 = p2
    var v6 uint64 = 4294967295
    if true {
        if true {
            fmt.Println(v6, p0)
            v7 := GenS1{F0: uint64(p0), F1: (float64((1.5 / 0.1)) != 0.0)}
            switch 4294967295 {
            case 42:
                v8 := GenS2{F0: 180}
                _ = v8
            }
            _ = v7
        } else {
            var v9 string = "foo bar"
            _ = v9
        }
        var v10 bool = true
        _ = v10
    }
    v11 := GenS3{F0: 1000, F1: 42, F2: 605}
    if false {
        for fi12 := 0; fi12 < 1; fi12++ {
            fmt.Println(v6, fi12, fi12)
            _ = fi12
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v6
    return (p0 + (p0 & p0))
}

func main() {
    var v1 float32 = (0.1 / (-2.25 + float32(0.0)))
    var v2 float32 = v1
    var v3 uint16 = 0
    v4 := []int{-9223372036854775808, 127, 1}
    v5 := []int{32767, 268, 953, 9223372036854775807}
    var v13 float64 = ((((1.5 * ((0.1 / ((((((1e10 * 0.0) - 0.0) + -2.25) / (-2.25 + 0.1)) + (float64((v1 - v2)) / float64(v1))) * 0.0)) + 0.0)) * 3.14159) + 0.1) / -2.25)
    v14 := GenS4{F0: (2147483648 > 149), F1: 275, F2: 32767}
    if true {
        for fi15 := 0; fi15 < 1; fi15++ {
            for fi16 := 0; fi16 < 17; fi16++ {
                v1 = 0.0
                _ = fi16
            }
            if (255 <= 166) {
                v17 := GenS1{F0: 2, F1: (2147483648 != 25)}
                v4 = append(v4, (^(fi15)))
                _ = v17
            }
            var v18 float64 = v13
            _ = fi15
            _ = v18
        }
        if (691 < (v3 << ((464) & 63))) {
            v3 = v3
            _ = len(v5)
        } else {
            v19 := GenS3{F0: 32767, F1: -43, F2: 42}
            var v20 uint8 = 128
            _ = v19
            _ = v20
        }
        v3 -= v3
    } else {
        v2 = -2.25
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        v1 = v1
        if true {
            v22 := GenS4{F0: false, F1: 42, F2: -1}
            v13 -= v13
            _ = v22
        } else {
            var v23 int32 = int32(v3)
            _ = v23
        }
        _ = fi21
    }
    if true {
        var v24 uint16 = (v3 | v3)
        _ = v24
    } else {
        for fi25 := 0; fi25 < 3; fi25++ {
            if (-1 > 1000) {
                if (float32(0.1) <= v1) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            }
            var v26 string = ""
            v1 -= ((v2 / v2) + 0.1)
            _ = fi25
            _ = v26
        }
        v5 = append(v5, -2147483648)
    }
    if (false && (65535 == 2147483647)) {
        var v27 int8 = 1
        _ = v27
    } else {
        v28 := GenS5{F0: 255, F1: ((((("hello" + "\n\t") + "hello") + (("\n\t" + "a") + "a")) + "\n\t") + "\n\t")}
        _ = v28
    }
    switch 32767 {
    case 254:
        v3 += (^(v3))
    case 9223372036854775807:
        var v29 float32 = v1
        _ = v29
    case 1:
        fmt.Println(v1)
    }
    v13 += float64(v1)
    fmt.Println(v2, v1, v13)
    switch -788 {
    case 4294967295:
        for fi30 := 0; fi30 < 7; fi30++ {
            var v31 float32 = v1
            _ = len(v4)
            _ = fi30
            _ = v31
        }
    case 946:
        fmt.Println(v3, v2)
    default:
        for fi32 := 0; fi32 < 6; fi32++ {
            var v33 int64 = -184
            fmt.Println(v33, fi32)
            _ = fi32
            _ = v33
        }
    }
    var v34 uint8 = 1
    var v35 uint8 = (^(v34))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
