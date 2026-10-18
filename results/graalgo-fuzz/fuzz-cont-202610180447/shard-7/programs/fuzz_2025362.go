package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int32; F1 int32; F2 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 int16; F2 bool }
type GenS6 struct { F0 uint16; F1 uint }
type GenS7 struct { F0 int; F1 uint64; F2 uint64 }
type GenS8 struct { F0 uint }
func genFn0(p0 int, p1 uint32) int16 {
    if false {
        fmt.Println(p1, p1)
    } else {
        switch "foo bar" {
        case "":
            v4 := []int{(p0 - p0)}
            _ = v4
        case "\n\t":
            v5 := []int{}
            _ = v5
        }
        fmt.Println(p1)
    }
    switch p0 {
    case -47:
        switch (p0 % (((p0 - (p0 >> ((0) & 63)))) | 1)) {
        case -641:
            v6 := []int{(p0 % ((p0) | 1)), -9223372036854775808, (p0 + (p0 << ((2) & 63))), (p0 * p0)}
            _ = v6
        }
    }
    v7 := []int{32767}
    var v8 float32 = (1e10 / 0.1)
    if (true || (!(0 == 147))) {
        v9 := GenS1{F0: p0, F1: 968}
        v7 = append(v7, 4294967295)
        p0 = p0
        _ = v9
    } else {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        v10 := GenS2{F0: 128, F1: -2147483648}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint8) bool {
    switch (("hello" + "foo bar") + ("" + ("foo bar" + "foo bar"))) {
    case "hello":
        var v11 float32 = (float32(1e10) * 3.14159)
        _ = v11
    case "\n\t":
        v12 := GenS2{F0: 2, F1: -2147483648}
        _ = v12
    }
    p0 = p0
    var v13 float32 = (0.0 + 3.14159)
    v13 = v13
    _ = p0
    _ = v13
    return true
}

func main() {
    v1 := []int{9223372036854775807}
    v2 := make(map[int]bool)
    var v3 bool = (2147483647 > 0)
    v3 = v3
    v14 := GenS3{F0: 0, F1: -2147483648, F2: 32767}
    fmt.Println(v3)
    v3 = false
    if (255 >= 200) {
        if v3 {
            var v15 float64 = 1e10
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v16 int16 = 42
            _ = v15
            _ = v16
        } else {
            var v17 float64 = (float64((float32(float64(-2.25)) + (3.14159 - ((((float32(-2.25) - -2.25) / (((1.5 - 3.14159) + 1e10) + float32(-2.25))) * -2.25) - (3.14159 - -2.25))))) / float64(1e10))
            if (995 < 1) {
                v18 := GenS4{F0: -1}
                var v19 uint8 = 130
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v18
                _ = v19
            }
            _ = v17
        }
        _ = len(v1)
    }
    switch "" {
    case "":
        v3 = ((6 != 0) || (91 <= 128))
    default:
        v20 := GenS4{F0: 120}
        _ = v20
    }
    if (682 >= 2) {
        v21 := GenS5{F0: 32767, F1: 42, F2: false}
        v3 = (42 >= genFn0(-1, 2))
        _ = v21
    } else {
        if (true && (-1 >= 0)) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    for fi22 := 0; fi22 < 12; fi22++ {
        if v3 {
            if ((42 >= 383) && (653 > 2)) {
                fmt.Println(v3, v3)
            }
            v23 := GenS6{F0: 0, F1: 4294967295}
            _ = v23
        } else {
            fmt.Println(v3)
            v3 = (1 >= 485)
            v3 = (false && (45 != -128))
        }
        var v24 string = ("foo bar" + ("hello" + ((("foo bar" + "foo bar") + ((("foo bar" + "a") + "a") + ("" + ""))) + ("" + "hello"))))
        var v25 string = (v24 + v24)
        _ = fi22
        _ = v24
        _ = v25
    }
    v3 = v3
    fmt.Println(v3)
    v3 = (!true)
    fmt.Println(v3)
    if v3 {
        if true {
            var v26 uint8 = 99
            v3 = (-94 > -128)
            switch ("foo bar" + (((("" + "\n\t") + ("\n\t" + "\n\t")) + "") + "hello")) {
            case "foo bar":
                v27 := GenS7{F0: -2147483648, F1: 42, F2: 42}
                _ = v27
            case "a":
                var v28 bool = false
                _ = v28
            default:
                if len(v1) > 0 {
                    v1[0] = 4294967295
                }
            }
            _ = v26
        }
    } else {
        v2[4294967295] = (((float32((-2.25 / -2.25)) * 1.5) / 0.1) >= float32(float64(0.0)))
    }
    v29 := GenS8{F0: 2}
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v29
    _ = v3
}
