package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint8; F2 bool }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 uint32; F1 int; F2 uint64 }
type GenS5 struct { F0 bool; F1 int32; F2 uint16 }
func genFn0(p0 int32, p1 uint) uint8 {
    for fi6 := 0; fi6 < 14; fi6++ {
        fmt.Println(p1)
        var v7 bool = (false || (2147483648 > 2147483648))
        v8 := []int{}
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := []int{127}
    for fi10 := 0; fi10 < 2; fi10++ {
        fmt.Println(p0, p1, p1)
        v11 := GenS1{F0: -1}
        fmt.Println(fi10, fi10, p0)
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        if (true || (false || false)) {
            var v13 bool = (658 == 4294967295)
            switch 752 {
            case 9223372036854775807:
                fmt.Println(p1)
            }
            for fi14 := 0; fi14 < 4; fi14++ {
                fmt.Println(fi12)
                _ = len(v9)
                _ = fi14
            }
            _ = v13
        } else {
            v15 := GenS1{F0: 1}
            p1 = 645
            _ = v15
        }
        var v16 string = "hello"
        _ = fi12
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v9
    return 255
}
func genFn1(p0 string, p1 uint16) int64 {
    v17 := []int{-9223372036854775808, 9223372036854775807}
    if len(v17) > 0 {
        v17[0] = 117
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v17
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]uint32)
    v3 := []int{292, 864}
    v4 := make(map[uint32]string)
    v5 := []uint32{2, 42}
    var v18 uint16 = 42
    fmt.Println(v18)
    var v19 uint32 = 1
    v18 = (v18 / ((v18) | 1))
    fmt.Println(v19)
    fmt.Println(v18)
    for fi20 := 0; fi20 < 19; fi20++ {
        for fi21 := 0; fi21 < 16; fi21++ {
            var v22 int16 = -32768
            var v23 float64 = (0.0 + 0.0)
            v24 := GenS1{F0: (v22 | 1000)}
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        if true {
            if false {
                v18 += v18
                if (v19 < v19) {
                    v25 := GenS2{F0: (((-2.25 * (float32(-2.25) / 1.0)) <= 1.5) && false), F1: 127, F2: (!true)}
                    _ = v25
                } else {
                    var v26 bool = false
                    if (1e10 <= ((float64(float32((0.1 / -2.25))) + (3.14159 * 1e10)) / (-2.25 + (0.1 + (-2.25 / (((3.14159 - 0.0) + ((((3.14159 - (3.14159 * 0.1)) - ((((float64(3.14159) + float64(1e10)) - (0.1 / 1e10)) + float64((1.5 * (1.5 + 0.1)))) / 1.0)) + (float64(0.0) / 1.0)) / 3.14159)) - 0.1)))))) {
                        var v27 bool = (!((-9223372036854775808 <= 1) || ((848 < 75) || (877 <= 42))))
                        _ = v27
                    } else {
                        v3 = append(v3, (-(fi20)))
                        var v28 int = 32768
                        _ = v28
                    }
                    _ = v26
                }
            }
        }
        fmt.Println(v19)
        _ = fi20
    }
    for fi29 := 0; fi29 < 19; fi29++ {
        if (-9223372036854775808 < 0) {
            v30 := GenS2{F0: (!false), F1: genFn0(2147483647, 501), F2: (false && (-128 < -1))}
            v31 := GenS1{F0: -1}
            switch "\n\t" {
            case "foo bar":
                v32 := GenS3{F0: 18446744073709551615, F1: (v19 / ((2147483647) | 1))}
                _ = v32
            default:
                v33 := GenS4{F0: v19, F1: fi29, F2: uint64(v18)}
                _ = v33
            }
            _ = v30
            _ = v31
        }
        v19 = (v19 | (v19 << ((195) & 63)))
        _ = fi29
    }
    v19 = 1
    v34 := GenS2{F0: (!(-943 > 322)), F1: 128, F2: ((v19 - 778) > v19)}
    switch -1 {
    case 9223372036854775807:
        var v35 string = ("" + ("foo bar" + "\n\t"))
        _ = v35
    case -165:
        for fi36 := 0; fi36 < 19; fi36++ {
            v37 := GenS5{F0: true, F1: -2147483648, F2: v18}
            v19 += (v19 >> ((601) & 63))
            _ = fi36
            _ = v37
        }
    }
    var v38 bool = (132 != -2147483648)
    _ = len(v3)
    fmt.Println(v38)
    var v39 string = "a"
    v38 = (!v38)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
