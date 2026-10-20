package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint16; F1 int; F2 int8 }
func genFn0(p0 int64, p1 int32, p2 uint) string {
    var v6 uint32 = 2
    v6 -= (v6 / ((v6) | 1))
    switch "hello" {
    case "\n\t":
        if false {
            v7 := GenS1{F0: 1, F1: 65535}
            v6 += v6
            _ = v7
        } else {
            var v8 float64 = ((3.14159 / 3.14159) - float64((1e10 + (0.1 - 1e10))))
            for fi9 := 0; fi9 < 10; fi9++ {
                v10 := GenS2{F0: 42, F1: ((true || true) || (0 < 1))}
                _ = fi9
                _ = v10
            }
            _ = v8
        }
    default:
        v11 := GenS1{F0: 0, F1: 0}
        _ = v11
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return ""
}
func genFn1(p0 int32, p1 uint32, p2 uint8) int32 {
    v12 := GenS2{F0: (p2 * p2), F1: true}
    switch int(p0) {
    case -2147483648:
        v13 := GenS3{F0: 0, F1: p0}
        _ = v13
    default:
        if (4294967295 >= 4294967295) {
            p2 += p2
            if false {
                switch -460 {
                case -1:
                    switch -9223372036854775808 {
                    case 829:
                        p2 += (p2 | 0)
                    default:
                        fmt.Println(p2, p2, p0)
                    }
                case -5:
                    var v14 uint8 = (p2 / (((p2 << ((458) & 63))) | 1))
                    _ = v14
                case 32767:
                    p0 = (p0 >> ((2) & 63))
                }
                var v15 bool = true
                v16 := GenS4{F0: v15, F1: -619}
                _ = v15
                _ = v16
            }
        } else {
            v17 := GenS4{F0: true, F1: 126}
            _ = v17
        }
    }
    v18 := []int{127}
    fmt.Println(p2, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v18
    return p0
}

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[uint32]uint32)
    var v3 int8 = 52
    var v4 float32 = float32(0.0)
    var v5 uint = 836
    v19 := GenS5{F0: 1}
    fmt.Println(v5, v1, v5)
    if (!(1e10 < (3.14159 - 1e10))) {
        fmt.Println(v4, v5, v4)
        var v20 float32 = 1e10
        _ = v20
    }
    var v21 float64 = float64(3.14159)
    switch ("hello" + "foo bar") {
    case "hello":
        var v22 float32 = 0.0
        _ = v22
    case "":
        switch ("\n\t" + ("a" + ((("" + ("" + ("\n\t" + ""))) + "\n\t") + "foo bar"))) {
        case "a":
            if true {
                fmt.Println(v3, v3)
                v5 = 4294967295
                if false {
                    for fi23 := 0; fi23 < 10; fi23++ {
                        var v24 int8 = v3
                        fmt.Println(v21, v1, v3)
                        _ = fi23
                        _ = v24
                    }
                }
            }
        case "\n\t":
            var v25 int16 = 1
            _ = v25
        }
    }
    v26 := []int{}
    var v27 int64 = 693
    var v28 bool = true
    v29 := GenS6{F0: 1}
    var v30 int64 = v27
    fmt.Println(v21, v30, v4)
    var v31 uint32 = 4294967295
    for fi32 := 0; fi32 < 3; fi32++ {
        var v33 uint32 = v1
        v5 = v5
        v1 += v33
        _ = fi32
        _ = v33
    }
    v28 = (v27 >= (v30 ^ (v30 & 42)))
    if v28 {
        var v34 uint32 = 36
        v4 -= (v4 / v4)
        _ = v34
    } else {
        v35 := GenS7{F0: 42, F1: 4294967295, F2: v3}
        fmt.Println(v21)
        _ = v35
    }
    if ("hello" != "foo bar") {
        for fi36 := 0; fi36 < 8; fi36++ {
            v4 -= 1.5
            var v37 uint8 = 42
            fmt.Println(v21)
            _ = fi36
            _ = v37
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
