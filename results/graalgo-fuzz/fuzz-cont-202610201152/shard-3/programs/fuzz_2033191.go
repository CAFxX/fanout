package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int8; F2 uint }
func genFn0(p0 int, p1 int, p2 int) uint16 {
    fmt.Println(p2, p2)
    v4 := []int{-2147483648, 372, (p2 + p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 2
}
func genFn1(p0 string) uint16 {
    v5 := GenS1{F0: 0, F1: 306, F2: 372}
    if (false || ((128 <= 128) || true)) {
        var v6 int8 = 56
        _ = v6
    } else {
        v7 := GenS1{F0: 603, F1: 2147483648, F2: 0}
        _ = v7
    }
    switch "foo bar" {
    case "foo bar":
        if ((true && true) && (!true)) {
            switch -9223372036854775808 {
            case 32767:
                p0 = p0
            default:
                v8 := GenS1{F0: 0, F1: 42, F2: 18446744073709551615}
                _ = v8
            }
            v9 := GenS1{F0: 255, F1: 4294967295, F2: 18446744073709551615}
            _ = v9
        } else {
            var v10 uint = 255
            for fi11 := 0; fi11 < 15; fi11++ {
                p0 = p0
                var v12 int32 = 0
                v13 := []int{9223372036854775807, 1}
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = v10
        }
    }
    _ = p0
    _ = v5
    return 1
}

func main() {
    var v1 uint16 = 1
    var v2 int64 = 9223372036854775807
    var v3 int32 = 2147483647
    fmt.Println(v1)
    for fi14 := 0; fi14 < 3; fi14++ {
        v2 = (v2 % (((^(v2))) | 1))
        v3 = -2147483648
        _ = fi14
    }
    if (!false) {
        v15 := []int{480, -614, 0, 128}
        if len(v15) > 0 {
            v15[0] = 9223372036854775807
        }
        for fi16 := 0; fi16 < 16; fi16++ {
            var v17 int8 = -1
            _ = fi16
            _ = v17
        }
        _ = v15
    } else {
        v18 := []int{32768, int(v3), -2147483648, -775}
        fmt.Println(v3)
        v3 -= (v3 | v3)
        _ = v18
    }
    v1 = v1
    for fi19 := 0; fi19 < 3; fi19++ {
        v1 -= v1
        _ = fi19
    }
    var v20 string = "\n\t"
    if false {
        v21 := GenS2{F0: 0}
        fmt.Println(v2)
        _ = v21
    } else {
        switch 128 {
        case -363:
            v3 += -1
        default:
            if (v1 >= v1) {
                v3 = (v3 ^ v3)
            } else {
                v22 := GenS3{F0: 0, F1: -128, F2: 4294967295}
                var v23 uint64 = 0
                v24 := []int{}
                _ = v22
                _ = v23
                _ = v24
            }
        }
        v3 = (v3 - (v3 - (^(v3))))
    }
    switch v20 {
    case "a":
        if (-1 == -538) {
            if (!(191 >= 127)) {
                v25 := []int{}
                _ = v25
            }
            fmt.Println(v2)
        }
    }
    v26 := GenS1{F0: 523, F1: 1, F2: 18446744073709551615}
    if false {
        var v27 bool = false
        _ = v27
    } else {
        if ((0.1 >= -2.25) || false) {
            v28 := []int{-35, 0, 9223372036854775807, 127}
            for fi29 := 0; fi29 < 7; fi29++ {
                v30 := GenS1{F0: 1, F1: 276, F2: 320}
                _ = fi29
                _ = v30
            }
            _ = v28
        } else {
            v3 = (v3 % (((v3 / ((2147483647) | 1))) | 1))
            if ((v1 >> ((uint(v1)) & 63)) != 65535) {
                v3 -= (-(v3))
                v31 := GenS2{F0: 1}
                _ = v31
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v3
}
