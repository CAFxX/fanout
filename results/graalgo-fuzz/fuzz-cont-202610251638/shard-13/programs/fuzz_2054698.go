package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 string; F2 uint }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 uint8; F1 int64; F2 int64 }
func genFn0(p0 int8, p1 int8) int32 {
    var v4 uint8 = 72
    var v5 float32 = (-2.25 * 0.1)
    v6 := GenS1{F0: 2, F1: 2, F2: 2147483647}
    v7 := GenS2{F0: true}
    if (1 == 517) {
        v8 := GenS1{F0: 119, F1: 1, F2: 2147483647}
        switch 2147483647 {
        case 941:
            p0 = (p1 & 127)
        default:
            p0 += p1
        }
        switch 0 {
        case 128:
            var v9 uint8 = (v4 - (v4 | 175))
            _ = v9
        case 1:
            if false {
                var v10 float64 = -2.25
                for fi11 := 0; fi11 < 20; fi11++ {
                    v12 := []int{-2147483648, 4294967295}
                    var v13 bool = (!(182 <= 18446744073709551615))
                    v14 := GenS3{F0: 0, F1: "hello", F2: 428}
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
                _ = v10
            }
        case -9223372036854775808:
            if (42 >= -750) {
                p0 += (p1 * p1)
            } else {
                var v15 uint16 = 0
                v16 := []int{int(v4), 32767, 4294967295, 32767}
                _ = v15
                _ = v16
            }
        default:
            var v17 uint16 = 65535
            _ = v17
        }
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint8) uint8 {
    var v18 uint32 = 2147483647
    v19 := GenS4{F0: -1, F1: (v18 % ((0) | 1))}
    p0 -= (p0 + 42)
    _ = p0
    _ = v18
    _ = v19
    return (p0 * p0)
}

func main() {
    var v1 int16 = 1000
    var v2 string = ("foo bar" + "a")
    v3 := make(map[uint32]uint32)
    v1 -= (v1 * -1)
    if true {
        var v20 int = 128
        _ = v20
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        var v22 uint16 = 395
        _ = fi21
        _ = v22
    }
    v23 := GenS5{F0: genFn1(1), F1: -978, F2: -1}
    v1 += (v1 ^ (v1 + 1))
    var v24 uint64 = 42
    var v25 uint64 = v24
    var v26 int32 = genFn0(-106, -128)
    if (!(!false)) {
        for fi27 := 0; fi27 < 2; fi27++ {
            var v28 int16 = (v1 - v1)
            _ = fi27
            _ = v28
        }
        var v29 uint8 = uint8(v25)
        v3[51] = 65535
        _ = v29
    }
    switch (v2 + ("\n\t" + v2)) {
    case "hello":
        v2 = v2
    default:
        v25 = v24
    }
    v24 = v24
    switch (("hello" + (("" + "a") + "a")) + v2) {
    case "foo bar":
        var v30 uint16 = 406
        _ = v30
    case "":
        v1 = (^(v1))
    }
    for fi31 := 0; fi31 < 4; fi31++ {
        v3[506] = 308
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
