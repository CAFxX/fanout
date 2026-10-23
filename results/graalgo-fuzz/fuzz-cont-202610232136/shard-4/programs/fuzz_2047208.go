package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 int; F2 uint32 }
type GenS3 struct { F0 string; F1 int16; F2 uint8 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 uint8; F1 int; F2 int }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 bool; F1 string }
type GenS8 struct { F0 string; F1 int64; F2 uint64 }
func genFn0(p0 uint16, p1 uint64, p2 int) bool {
    p0 = (p0 / ((p0) | 1))
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return true
}
func genFn1(p0 uint8) int16 {
    if (("\n\t" + ("" + ("\n\t" + "hello"))) > "") {
        fmt.Println(p0)
    } else {
        p0 = 128
    }
    v4 := GenS1{F0: 11}
    var v5 uint = 65535
    switch "foo bar" {
    case "foo bar":
        if (!false) {
            v5 = (v5 & v5)
            var v6 uint16 = 624
            _ = v6
        } else {
            var v7 bool = ((!false) || false)
            v8 := GenS1{F0: 127}
            if v7 {
                var v9 int8 = -128
                v9 = 0
                var v10 bool = v7
                _ = v10
                _ = v9
            }
            _ = v7
            _ = v8
        }
    case "\n\t":
        var v11 bool = false
        _ = v11
    case "a":
        var v12 int16 = 1
        _ = v12
    default:
        for fi13 := 0; fi13 < 14; fi13++ {
            p0 = uint8(v5)
            _ = fi13
        }
    }
    _ = p0
    _ = v4
    _ = v5
    return -1
}

func main() {
    v1 := make(map[int]string)
    var v2 string = "hello"
    var v3 float32 = (0.0 + (1e10 / (0.1 - 1.5)))
    for fi14 := 0; fi14 < 1; fi14++ {
        for fi15 := 0; fi15 < 13; fi15++ {
            v16 := GenS2{F0: 980, F1: 127, F2: 4294967295}
            var v17 int = -2147483648
            _ = fi15
            _ = v16
            _ = v17
        }
        fmt.Println(fi14, v3, v2)
        for fi18 := 0; fi18 < 14; fi18++ {
            if true {
                fmt.Println(v3, fi18, fi18)
                v3 = ((float32(((-2.25 * (1.5 * (0.0 / -2.25))) * (3.14159 * 1e10))) - v3) / (float32(0.1) / 1.0))
                if (!false) {
                    v19 := []int{}
                    _ = v19
                }
            }
            v20 := GenS1{F0: -128}
            _ = fi18
            _ = v20
        }
        _ = fi14
    }
    if (((false && true) || ((true || ((true && (0 >= 0)) || false)) && false)) && false) {
        fmt.Println(v2, v3)
        v21 := GenS3{F0: v2, F1: -32768, F2: 127}
        v22 := v1[42]
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        fmt.Println(fi23, v3, v3)
        for fi24 := 0; fi24 < 17; fi24++ {
            var v25 int16 = 42
            v2 = v2
            v26 := GenS4{F0: -260, F1: 100}
            v25 = (v25 % ((v25) | 1))
            _ = fi24
            _ = v25
            _ = v26
        }
        v27 := GenS5{F0: 0, F1: fi23, F2: fi23}
        _ = fi23
        _ = v27
    }
    v2 = (v2 + v2)
    if true {
        v2 = "hello"
    } else {
        v3 -= float32(0.1)
        var v28 float32 = v3
        var v29 uint16 = 65535
        _ = v28
        _ = v29
    }
    v30 := GenS6{F0: -427, F1: ("\n\t" + v2)}
    v3 += (v3 * 1e10)
    switch "a" {
    case "\n\t":
        var v31 uint8 = 42
        _ = v31
    case "hello":
        if (true && (-795 == -159)) {
            fmt.Println(v3, v3)
            switch v2 {
            case "":
                fmt.Println(v3, v3)
            case "\n\t":
                fmt.Println(v3, v3)
            }
            v32 := GenS7{F0: false, F1: ("" + (v2 + v2))}
            _ = v32
        }
    default:
        v2 = v2
    }
    switch -131 {
    case 0:
        var v33 int64 = -296
        _ = v33
    case 127:
        for fi34 := 0; fi34 < 14; fi34++ {
            var v35 int = fi34
            switch fi34 {
            case -9223372036854775808:
                var v36 int32 = 317
                _ = v36
            }
            _ = fi34
            _ = v35
        }
    case -1:
        var v37 uint64 = 42
        _ = v37
    default:
        fmt.Println(v2)
    }
    v38 := GenS8{F0: "", F1: -9223372036854775808, F2: 563}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v38
}
