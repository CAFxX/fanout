package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint16; F1 uint8; F2 bool }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 uint; F1 int16; F2 int32 }
type GenS5 struct { F0 bool; F1 bool; F2 int32 }
type GenS6 struct { F0 int32; F1 int32; F2 uint8 }
func genFn0(p0 uint, p1 string) string {
    if true {
        if ((0.1 / (0.1 + (1.5 * 0.1))) <= 1.5) {
            if (-203 <= -32768) {
                v5 := GenS1{F0: 9223372036854775807, F1: 18446744073709551615}
                _ = v5
            }
            fmt.Println(p1)
        } else {
            p0 += p0
        }
        var v6 uint32 = 2147483648
        switch 128 {
        case -582:
            fmt.Println(p1, p1, p0)
        case 4294967295:
            fmt.Println(v6, p1, p1)
        }
        _ = v6
    }
    p1 = p1
    v7 := []int{-1, 9223372036854775807, -9223372036854775808}
    if false {
        switch -2147483648 {
        case -42:
            v8 := GenS2{F0: 42, F1: 127, F2: false}
            _ = v8
        }
    } else {
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = v7
    return p1
}
func genFn1(p0 uint32, p1 uint, p2 uint8) int {
    var v9 bool = false
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 2147483647
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{}
    v3 := make(map[int]int)
    v4 := []string{"foo bar", "foo bar", "a"}
    var v10 int = 1
    _ = len(v2)
    v11 := GenS2{F0: 852, F1: 23, F2: ((3.14159 + (float32(float64(float32((1e10 + 1.5)))) / (0.0 + 3.14159))) < -2.25)}
    for fi12 := 0; fi12 < 13; fi12++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v13 := GenS3{F0: true, F1: 1}
        _ = fi12
        _ = v13
    }
    fmt.Println(v10)
    v14 := GenS4{F0: 0, F1: 42, F2: -1}
    var v15 uint64 = 2
    var v16 uint16 = 255
    v15 = (v15 * v15)
    if (true || true) {
        var v17 uint8 = 255
        v15 += (v15 << ((42) & 63))
        _ = v17
    }
    switch (v10 - (v10 * (v10 % ((v10) | 1)))) {
    case 32767:
        _ = len(v4)
    case 871:
        v15 = (^(v15))
    default:
        if false {
            var v18 bool = (!true)
            v15 -= (^(v15))
            if len(v4) > 0 {
                v4[0] = (("a" + "hello") + "")
            }
            _ = v18
        } else {
            var v19 int8 = -1
            v16 -= uint16(v15)
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        for fi21 := 0; fi21 < 10; fi21++ {
            v15 -= v15
            if false {
                v10 = (^(fi20))
                if (true && false) {
                    v15 = (v15 - (v15 << ((274) & 63)))
                } else {
                    v22 := GenS5{F0: (false || true), F1: true, F2: int32(v16)}
                    var v23 string = ("hello" + "a")
                    _ = v22
                    _ = v23
                }
                v4 = append(v4, ("" + "foo bar"))
                fmt.Println(v15)
                v24 := GenS6{F0: 1, F1: 2147483647, F2: 1}
                _ = v24
            }
            v25 := v1[4294967295]
            _ = fi21
            _ = v25
        }
        var v26 float32 = (1.5 * 3.14159)
        _ = fi20
        _ = v26
    }
    var v27 uint64 = v15
    fmt.Println(v16, v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
