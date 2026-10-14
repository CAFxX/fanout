package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 bool; F1 int }
func genFn0(p0 uint16, p1 uint32, p2 uint8) uint16 {
    p1 -= (^(p1))
    if (-86 >= -128) {
        if true {
            v4 := []int{32767, -9223372036854775808, 42}
            _ = v4
        }
    }
    var v5 uint = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return uint16(p1)
}
func genFn1(p0 int16, p1 int64) uint {
    if false {
        p1 = p1
        v6 := GenS1{F0: 83, F1: 128, F2: 4294967295}
        v7 := []int{32767, 9223372036854775807, 1}
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := GenS1{F0: 4294967295, F1: (fi8 >> ((2) & 63)), F2: 504}
        v10 := []int{0, (fi8 + (^(fi8)))}
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: "foo bar"}
    p1 += (p1 + (p1 * 42))
    var v12 uint = 4294967295
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 986
}

func main() {
    var v1 uint16 = 2
    var v2 float64 = float64(((1e10 * (float32(1e10) * float32(((float64(-2.25) / 1e10) + 3.14159)))) * float32(0.1)))
    v3 := []string{}
    var v13 uint8 = 2
    var v14 int8 = 1
    v14 -= (-(v14))
    if (false || (!false)) {
        v15 := GenS3{F0: uint32(v1), F1: v1}
        _ = v15
    }
    var v16 uint16 = 255
    var v17 int8 = v14
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 += (v16 ^ 65535)
    v16 = (v1 ^ v16)
    var v18 int8 = (v17 | -16)
    switch "\n\t" {
    case "a":
        var v19 string = ("" + "foo bar")
        _ = v19
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        var v21 int32 = -2147483648
        if (true && true) {
            var v22 int16 = 42
            v23 := GenS2{F0: (("foo bar" + ("hello" + "")) + "\n\t")}
            switch fi20 {
            case 1:
                var v24 float64 = v2
                _ = v24
            default:
                var v25 uint16 = v1
                _ = v25
            }
            _ = v22
            _ = v23
        }
        v14 -= (v17 << ((4294967295) & 63))
        _ = fi20
        _ = v21
    }
    fmt.Println(v16, v1)
    for fi26 := 0; fi26 < 17; fi26++ {
        v3 = append(v3, "a")
        switch "" {
        case "foo bar":
            if (!false) {
                v27 := GenS4{F0: (v17 <= (v17 | 1)), F1: fi26}
                v14 -= 39
                _ = v27
            }
        case "":
            switch (("" + "hello") + "\n\t") {
            case "a":
                var v28 int16 = 1000
                _ = v28
            }
        case "a":
            v18 = v18
        }
        _ = fi26
    }
    switch 42 {
    case 1:
        if (("a" + "\n\t") == "hello") {
            var v29 int32 = 1
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
