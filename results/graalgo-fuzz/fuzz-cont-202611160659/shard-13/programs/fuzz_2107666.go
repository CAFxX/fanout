package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint64; F1 int16; F2 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 int; F1 uint64; F2 bool }
type GenS6 struct { F0 int32; F1 string; F2 uint32 }
type GenS7 struct { F0 string }
func genFn0(p0 uint64, p1 uint8, p2 uint64) uint16 {
    var v5 uint32 = 2
    fmt.Println(p0)
    v5 = v5
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := []int{(fi6 >> ((716) & 63)), fi6, fi6, fi6}
        v8 := GenS1{F0: v5, F1: -173}
        if false {
            v9 := GenS2{F0: p2, F1: 986, F2: 2}
            if ((!false) && (1.5 > 1.5)) {
                switch "\n\t" {
                case "hello":
                    var v10 uint = 0
                    _ = v10
                case "a":
                    v5 = 65535
                case "foo bar":
                    p0 -= (p2 + (p0 % (((p0 - 111)) | 1)))
                }
            }
            _ = v9
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    if false {
        var v11 uint8 = 186
        var v12 int32 = 2147483647
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 255
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int = 364
    var v3 float32 = -2.25
    var v4 uint8 = 0
    v13 := []int{v2, (v2 >> ((18446744073709551615) & 63))}
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 float64 = float64((((-2.25 - ((v3 - v3) * -2.25)) + v3) / 1e10))
        _ = fi14
        _ = v15
    }
    var v16 uint8 = v4
    v17 := GenS3{F0: (-(v2))}
    switch v2 {
    case 1:
        fmt.Println(v16, v1)
    case -9223372036854775808:
        var v18 int16 = 794
        _ = v18
    default:
        fmt.Println(v16, v3, v1)
    }
    var v19 int = 180
    if len(v13) > 0 {
        v13[0] = (^(v19))
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        switch "foo bar" {
        case "a":
            var v21 int8 = -79
            _ = v21
        case "":
            for fi22 := 0; fi22 < 12; fi22++ {
                if false {
                    if true {
                        var v23 string = "foo bar"
                        if false {
                            fmt.Println(v3, v3, v1)
                        }
                        _ = v23
                    }
                }
                if len(v13) > 0 {
                    v13[0] = fi20
                }
                _ = fi22
            }
        case "foo bar":
            v24 := GenS4{F0: (-2147483648 != -292), F1: 4294967295}
            _ = v24
        default:
            v25 := GenS5{F0: v19, F1: (v1 & 90), F2: false}
            _ = v25
        }
        for fi26 := 0; fi26 < 10; fi26++ {
            if false {
                if (!true) {
                    var v27 uint32 = 4294967295
                    var v28 int = v2
                    v16 = (v16 + (v16 | (v16 | 135)))
                    v3 = v3
                    _ = v27
                    _ = v28
                } else {
                    var v29 int64 = 42
                    v30 := GenS6{F0: -1, F1: ("a" + ""), F2: 2147483647}
                    _ = v29
                    _ = v30
                }
                fmt.Println(v2, fi20, v19)
                var v31 int32 = 929
                _ = v31
            }
            _ = fi26
        }
        switch (("a" + "hello") + ("a" + "\n\t")) {
        case "a":
            if true {
                v19 += (v2 / ((v2) | 1))
                var v32 int32 = 42
                fmt.Println(v19, v4, v32)
                _ = v32
            }
        case "\n\t":
            v16 = 255
        case "foo bar":
            v1 = v1
        default:
            v4 -= (v16 % (((v16 / ((0) | 1))) | 1))
        }
        _ = fi20
    }
    for fi33 := 0; fi33 < 13; fi33++ {
        v34 := GenS1{F0: 456, F1: 1}
        v35 := GenS7{F0: "\n\t"}
        _ = fi33
        _ = v34
        _ = v35
    }
    for fi36 := 0; fi36 < 16; fi36++ {
        if len(v13) > 0 {
            v13[0] = (v19 + v19)
        }
        v3 = v3
        fmt.Println(v2)
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
