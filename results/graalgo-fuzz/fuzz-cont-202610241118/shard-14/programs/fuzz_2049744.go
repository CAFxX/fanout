package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS5 struct { F0 int64; F1 int8; F2 int64 }
type GenS6 struct { F0 int; F1 int32; F2 int32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int8; F1 bool; F2 uint8 }
func genFn0(p0 int8, p1 string) uint8 {
    if (!(128 > 0)) {
        p0 -= p0
    } else {
        v5 := GenS1{F0: p0, F1: 1}
        fmt.Println(p0)
        _ = v5
    }
    v6 := []int{-1}
    _ = p0
    _ = p1
    _ = v6
    return 0
}
func genFn1() uint16 {
    if (!((4294967295 == 604) || (!(true || (-1 <= 42))))) {
        v7 := GenS2{F0: 32767}
        if (!(!(0 < 600))) {
            v8 := []int{398}
            var v9 float32 = 0.1
            var v10 bool = false
            _ = v10
            _ = v8
            _ = v9
        } else {
            for fi11 := 0; fi11 < 9; fi11++ {
                v12 := []int{fi11}
                for fi13 := 0; fi13 < 20; fi13++ {
                    v14 := GenS1{F0: 42, F1: 0}
                    _ = fi13
                    _ = v14
                }
                _ = fi11
                _ = v12
            }
            var v15 int8 = 40
            fmt.Println(v15)
            _ = v15
        }
        _ = v7
    } else {
        var v16 float32 = 0.1
        v17 := GenS3{F0: 860, F1: 42}
        _ = v16
        _ = v17
    }
    var v18 float32 = float32((1.5 / 1.5))
    v18 = ((v18 * ((((v18 * v18) - -2.25) - 1e10) * v18)) / (1e10 * (v18 + v18)))
    _ = v18
    return 413
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[int]uint32)
    v3 := []string{("foo bar" + ("" + "hello")), "a"}
    var v4 uint16 = 925
    switch -2147483648 {
    case -2147483648:
        v19 := GenS2{F0: 255}
        _ = v19
    case 9223372036854775807:
        var v20 int32 = 595
        _ = v20
    case 42:
        v1 += v1
    default:
        v1 += (1.5 / v1)
    }
    switch "" {
    case "a":
        for fi21 := 0; fi21 < 1; fi21++ {
            var v22 int = (fi21 / ((-2147483648) | 1))
            _ = fi21
            _ = v22
        }
    case "\n\t":
        if (696 > 1) {
            v4 += genFn1()
            switch "a" {
            case "a":
                var v23 uint16 = genFn1()
                _ = v23
            case "foo bar":
                if ((18446744073709551615 < 4294967295) && false) {
                    var v24 uint32 = 4294967295
                    _ = v24
                }
            case "":
                v25 := GenS1{F0: -7, F1: 97}
                _ = v25
            }
        } else {
            v26 := GenS4{F0: (v4 / (((v4 >> ((1) & 63))) | 1)), F1: 0, F2: 18446744073709551615}
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 8; fi27++ {
        var v28 float32 = float32(float64((3.14159 + 0.0)))
        v1 = float64(v28)
        for fi29 := 0; fi29 < 9; fi29++ {
            var v30 uint16 = 2
            if len(v3) > 0 {
                v3[0] = ""
            }
            v1 = v1
            v31 := GenS5{F0: 329, F1: 100, F2: -9223372036854775808}
            v28 -= 0.1
            v1 -= (v1 * 0.1)
            v32 := GenS6{F0: (fi27 ^ fi29), F1: 42, F2: 42}
            _ = fi29
            _ = v30
            _ = v31
            _ = v32
        }
        _ = fi27
        _ = v28
    }
    v33 := GenS7{F0: 1}
    v4 = 816
    fmt.Println(v4, v1)
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v1)
    default:
        if (255 <= 255) {
            if (!(!(4294967295 <= 2147483647))) {
                var v34 int16 = -436
                v4 -= genFn1()
                _ = v34
            }
            for fi35 := 0; fi35 < 13; fi35++ {
                _ = len(v3)
                v4 = (v4 << ((65535) & 63))
                _ = fi35
            }
            var v36 string = ("hello" + "a")
            _ = v36
        }
    }
    if (-912 >= -9223372036854775808) {
        if true {
            v1 = (v1 * 0.0)
            if (32767 > -1) {
                v4 -= 796
            } else {
                v37 := GenS8{F0: 104, F1: (42 != 2147483647), F2: 128}
                v1 += v1
                _ = v37
            }
        }
        var v38 uint64 = 48
        _ = v38
    }
    switch ("" + "\n\t") {
    case "a":
        v4 = 399
    default:
        fmt.Println(v1, v1)
    }
    v1 -= v1
    v4 += (^(v4))
    v4 -= genFn1()
    var v39 uint = 4294967295
    var v40 int = 0
    v40 = v40
    v4 -= (v4 + v4)
    if (1 < 1000) {
        for fi41 := 0; fi41 < 19; fi41++ {
            v3 = append(v3, ("foo bar" + "foo bar"))
            if ("hello" != "") {
                v1 -= (((((float64(3.14159) - v1) / v1) / v1) * (0.0 / (float64(1.5) + (v1 * v1)))) * -2.25)
                for fi42 := 0; fi42 < 12; fi42++ {
                    v1 += (0.1 - v1)
                    var v43 int64 = 9223372036854775807
                    v43 = 258
                    _ = fi42
                    _ = v43
                }
            }
            var v44 int64 = 435
            _ = fi41
            _ = v44
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v39
    _ = v4
    _ = v40
}
