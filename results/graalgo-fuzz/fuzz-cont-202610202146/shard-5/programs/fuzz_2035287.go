package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 uint8; F1 string }

func main() {
    var v1 float64 = 1e10
    var v2 float32 = float32((v1 + v1))
    v3 := make(map[string]uint32)
    var v4 uint16 = 649
    v5 := make(map[int]bool)
    v6 := make(map[string]int)
    var v7 int32 = 1
    fmt.Println(v4)
    if true {
        v8 := GenS1{F0: 241, F1: 2}
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 int32 = 42
            var v11 uint32 = 2147483648
            v4 = uint16(v11)
            _ = fi9
            _ = v10
            _ = v11
        }
        fmt.Println(v1)
        _ = v8
    } else {
        v12 := GenS2{F0: 591, F1: 0, F2: -1}
        var v13 int32 = (^(v7))
        _ = v12
        _ = v13
    }
    v14 := GenS3{F0: ("\n\t" + ("foo bar" + "foo bar")), F1: 42}
    fmt.Println(v2, v7)
    v1 = (v1 * (1e10 / v1))
    switch "hello" {
    case "\n\t":
        switch -135 {
        case 127:
            v4 = (v4 >> ((75) & 63))
        case 1:
            var v15 float32 = ((((v2 / v2) / float32((((0.0 - v1) + v1) * v1))) - 1.5) / v2)
            _ = v15
        case -2147483648:
            v16 := GenS4{F0: (v7 | v7), F1: 42}
            _ = v16
        }
    case "a":
        var v17 int8 = -9
        _ = v17
    default:
        v7 += v7
    }
    if false {
        if false {
            v18 := GenS5{F0: 2, F1: ("foo bar" + "foo bar")}
            _ = v18
        } else {
            var v19 uint16 = v4
            v7 -= (v7 ^ (v7 | (v7 ^ v7)))
            _ = v19
        }
        fmt.Println(v1, v2, v7)
    } else {
        for fi20 := 0; fi20 < 12; fi20++ {
            if ((true && (!(("" + ("a" + "")) > ("a" + "foo bar")))) || true) {
                var v21 uint = 2
                _ = v21
            }
            _ = fi20
        }
        switch (("" + "foo bar") + "") {
        case "foo bar":
            var v22 int8 = -128
            _ = v22
        default:
            for fi23 := 0; fi23 < 8; fi23++ {
                delete(v3, "a")
                _ = fi23
            }
        }
        var v24 uint64 = 42
        _ = v24
    }
    fmt.Println(v7)
    switch 657 {
    case 370:
        v7 = -176
    case 810:
        v25 := []int{2147483647, -2147483648, 1, 390}
        _ = v25
    default:
        var v26 uint16 = (^(v4))
        _ = v26
    }
    if (-2147483648 > -2147483648) {
        fmt.Println(v7, v4, v7)
        var v27 float32 = 3.14159
        _ = v27
    }
    switch 0 {
    case -371:
        v6[("" + "foo bar")] = -602
    case 128:
        switch int(v7) {
        case -9223372036854775808:
            fmt.Println(v2, v7)
        }
    case -2147483648:
        switch ("" + "foo bar") {
        case "":
            v1 -= 1.5
        case "a":
            for fi28 := 0; fi28 < 11; fi28++ {
                fmt.Println(v7, fi28, v7)
                var v29 bool = true
                v30 := v6[""]
                _ = fi28
                _ = v29
                _ = v30
            }
        default:
            var v31 uint = 787
            _ = v31
        }
    default:
        fmt.Println(v1, v7, v7)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
