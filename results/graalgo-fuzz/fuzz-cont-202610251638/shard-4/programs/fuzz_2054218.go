package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int8; F1 int64; F2 int8 }
type GenS4 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS5 struct { F0 int8; F1 int16; F2 uint }
type GenS6 struct { F0 bool; F1 bool; F2 bool }
type GenS7 struct { F0 int16; F1 int64 }
func genFn0(p0 uint, p1 uint, p2 uint32) uint32 {
    switch "\n\t" {
    case "foo bar":
        v7 := []int{629}
        _ = v7
    case "\n\t":
        v8 := GenS1{F0: uint8(p2)}
        _ = v8
    case "hello":
        p0 = p1
    }
    v9 := []int{811, -370, 32767}
    p1 = p1
    if (!(1 <= -1)) {
        if len(v9) > 0 {
            v9[0] = 1
        }
        fmt.Println(p0, p2, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return p2
}
func genFn1(p0 string, p1 uint16) bool {
    if false {
        if (255 >= 42) {
            v10 := []int{4294967295, 2147483647, -183, 1}
            _ = v10
        } else {
            if true {
                var v11 int16 = 0
                p1 -= (^(p1))
                _ = v11
            } else {
                fmt.Println(p1)
                p0 = "a"
            }
            v12 := []int{}
            _ = v12
        }
        var v13 uint16 = (p1 | p1)
        _ = v13
    } else {
        v14 := GenS2{F0: 179, F1: (^(p1))}
        var v15 string = p0
        switch 32768 {
        case -1:
            var v16 uint64 = 42
            _ = v16
        case 9223372036854775807:
            var v17 float32 = 1e10
            _ = v17
        default:
            var v18 uint8 = 1
            _ = v18
        }
        _ = v14
        _ = v15
    }
    v19 := []int{-1, 9223372036854775807}
    v20 := GenS2{F0: 908, F1: (p1 & (^(p1)))}
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    return true
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{288, 1, 42, 32767}
    v3 := make(map[string]bool)
    v4 := make(map[int]uint32)
    var v5 uint32 = 65535
    v6 := []uint32{v5, v5, (v5 ^ (v5 ^ (v5 * v5)))}
    for fi21 := 0; fi21 < 9; fi21++ {
        fmt.Println(v5, v5)
        switch "hello" {
        case "":
            var v22 uint32 = (v5 + 2147483647)
            _ = v22
        case "a":
            v5 = genFn0(18446744073709551615, 42, genFn0(255, 4294967295, v5))
        default:
            if ("" < "foo bar") {
                var v23 float64 = -2.25
                for fi24 := 0; fi24 < 7; fi24++ {
                    v25 := GenS2{F0: (v5 * v5), F1: 2}
                    _ = fi24
                    _ = v25
                }
                _ = v23
            } else {
                for fi26 := 0; fi26 < 11; fi26++ {
                    if len(v6) > 0 {
                        v6[0] = v5
                    }
                    switch ((("" + "hello") + "foo bar") + (("a" + "foo bar") + ((((("\n\t" + "hello") + "hello") + "\n\t") + "") + "foo bar"))) {
                    case "hello":
                        var v27 int16 = -495
                        _ = v27
                    case "foo bar":
                        v5 = 42
                    }
                    _ = fi26
                }
                v5 += 489
            }
        }
        fmt.Println(v5, fi21)
        _ = fi21
    }
    v28 := GenS3{F0: 100, F1: 1, F2: -43}
    if (65535 >= 2) {
        for fi29 := 0; fi29 < 11; fi29++ {
            v30 := GenS4{F0: 42, F1: (true || false), F2: 43}
            v31 := GenS3{F0: 72, F1: -357, F2: 127}
            for fi32 := 0; fi32 < 1; fi32++ {
                switch ("" + "hello") {
                case "hello":
                    v5 += (^(v5))
                case "":
                    var v33 float64 = 3.14159
                    _ = v33
                default:
                    v5 += v5
                }
                _ = fi32
            }
            _ = fi29
            _ = v30
            _ = v31
        }
    } else {
        fmt.Println(v5)
    }
    fmt.Println(v5)
    switch ("" + ("a" + "")) {
    case "foo bar":
        fmt.Println(v5)
    case "a":
        for fi34 := 0; fi34 < 9; fi34++ {
            v35 := GenS4{F0: 234, F1: (!false), F2: 207}
            _ = fi34
            _ = v35
        }
    }
    for fi36 := 0; fi36 < 12; fi36++ {
        v37 := GenS3{F0: -128, F1: 1, F2: -128}
        v38 := GenS5{F0: 22, F1: -525, F2: 18446744073709551615}
        v5 = (v5 / ((genFn0(4294967295, 2, 632)) | 1))
        _ = fi36
        _ = v37
        _ = v38
    }
    v39 := GenS6{F0: ((4294967295 <= 412) && ((((202 != 2) && ((!true) && (42 != 0))) || ((float64((0.1 * -2.25)) - (1.5 - float64(3.14159))) != -2.25)) || (65535 < 451))), F1: false, F2: (-1 > -283)}
    v5 -= (v5 | (v5 + 122))
    _ = len(v6)
    if true {
        var v40 uint32 = v5
        fmt.Println(v40, v5)
        if (true && false) {
            v5 += genFn0(65535, 65535, (v40 | v40))
            v40 = 65535
            v5 -= genFn0(2, 42, v40)
        } else {
            fmt.Println(v5, v40)
            delete(v3, "a")
        }
        _ = v40
    }
    var v41 uint = 2
    fmt.Println(v41)
    var v42 float64 = (3.14159 + 0.1)
    v42 = 1.5
    if true {
        var v43 uint32 = v5
        if false {
            v44 := GenS7{F0: 42, F1: -87}
            var v45 int32 = 2147483647
            _ = v44
            _ = v45
        } else {
            var v46 int32 = 0
            fmt.Println(v42, v41)
            _ = v46
        }
        _ = v43
    }
    if (!false) {
        var v47 uint8 = 136
        v5 = genFn0(v41, (v41 - v41), v5)
        _ = v47
    } else {
        var v48 uint64 = 4294967295
        if false {
            _ = len(v2)
            for fi49 := 0; fi49 < 5; fi49++ {
                var v50 int64 = -9223372036854775808
                if len(v6) > 0 {
                    v6[0] = genFn0(uint(v48), 4294967295, (v5 * v5))
                }
                _ = fi49
                _ = v50
            }
        }
        _ = v48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v39
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
