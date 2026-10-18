package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 uint; F1 int; F2 int16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint8; F1 int8; F2 bool }
type GenS8 struct { F0 int16; F1 int64 }
func genFn0(p0 uint16) int32 {
    fmt.Println(p0)
    if false {
        v7 := GenS1{F0: (988 > 1), F1: -113}
        p0 = 32767
        v8 := []int{1, -2147483648}
        _ = v7
        _ = v8
    } else {
        if false {
            p0 = (p0 / ((p0) | 1))
            p0 += p0
        }
        if (!true) {
            p0 = (p0 >> ((18446744073709551615) & 63))
        }
        p0 += p0
    }
    switch "foo bar" {
    case "":
        switch "hello" {
        case "\n\t":
            v9 := GenS2{F0: ("foo bar" + "a"), F1: 32768}
            _ = v9
        case "foo bar":
            var v10 int8 = 1
            _ = v10
        default:
            var v11 uint32 = 0
            _ = v11
        }
    case "hello":
        var v12 int16 = -32768
        _ = v12
    case "a":
        v13 := GenS3{F0: 2}
        _ = v13
    default:
        var v14 string = "\n\t"
        _ = v14
    }
    p0 += p0
    for fi15 := 0; fi15 < 3; fi15++ {
        fmt.Println(p0)
        _ = fi15
    }
    _ = p0
    return 1
}
func genFn1(p0 uint64) int {
    switch 9223372036854775807 {
    case 127:
        switch 4294967295 {
        case 42:
            v16 := GenS4{F0: -565, F1: 473}
            _ = v16
        case 128:
            v17 := []int{32767, 127, 4294967295}
            _ = v17
        case 127:
            v18 := []int{-9223372036854775808, -616}
            _ = v18
        }
    case 38:
        p0 = (p0 % ((1) | 1))
    default:
        for fi19 := 0; fi19 < 10; fi19++ {
            p0 = 1
            _ = fi19
        }
    }
    v20 := []int{32768}
    if (!false) {
        v21 := GenS2{F0: "hello", F1: 42}
        _ = len(v20)
        if false {
            var v22 uint64 = 4294967295
            for fi23 := 0; fi23 < 11; fi23++ {
                if true {
                    if len(v20) > 0 {
                        v20[0] = (fi23 << ((784) & 63))
                    }
                    var v24 float32 = 0.1
                    _ = v24
                }
                _ = fi23
            }
            _ = v22
        }
        _ = v21
    }
    _ = p0
    _ = v20
    return 4294967295
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 32767
    var v3 int = 9223372036854775807
    var v4 bool = true
    var v5 string = "foo bar"
    var v6 float32 = (0.0 * 0.1)
    v5 = "a"
    fmt.Println(v3)
    for fi25 := 0; fi25 < 6; fi25++ {
        v3 += fi25
        _ = fi25
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        v27 := GenS1{F0: (675 >= 0), F1: -128}
        v28 := []int{}
        _ = fi26
        _ = v27
        _ = v28
    }
    if (!v4) {
        v29 := v1[uint32(v2)]
        var v30 float64 = 0.0
        v31 := GenS5{F0: 47, F1: (^(v3)), F2: -32768}
        _ = v29
        _ = v30
        _ = v31
    }
    v6 = 0.0
    var v32 uint8 = 42
    var v33 uint8 = (^(v32))
    v4 = false
    v33 = v32
    v34 := GenS6{F0: v3}
    if true {
        v32 = v32
    }
    for fi35 := 0; fi35 < 13; fi35++ {
        v3 -= (fi35 ^ 32768)
        _ = fi35
    }
    for fi36 := 0; fi36 < 5; fi36++ {
        var v37 uint16 = (v2 ^ 1)
        _ = fi36
        _ = v37
    }
    if true {
        v5 = v5
    } else {
        v38 := GenS4{F0: v3, F1: -9223372036854775808}
        v3 = (v3 ^ 42)
        _ = v38
    }
    v39 := GenS7{F0: (v32 % (((v32 / ((v33) | 1))) | 1)), F1: 0, F2: v4}
    v40 := GenS8{F0: 1000, F1: -9223372036854775808}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
