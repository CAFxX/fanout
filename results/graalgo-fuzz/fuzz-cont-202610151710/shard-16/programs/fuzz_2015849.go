package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = 3.14159
    var v3 uint8 = 246
    var v4 uint16 = 995
    v5 := GenS1{F0: 1}
    if (false || (0.0 != 1e10)) {
        fmt.Println(v4)
        if ((^(v3)) > (v3 << ((42) & 63))) {
            switch "\n\t" {
            case "hello":
                v6 := GenS2{F0: -1, F1: -112}
                _ = v6
            case "foo bar":
                for fi7 := 0; fi7 < 6; fi7++ {
                    var v8 string = (("foo bar" + ("foo bar" + "hello")) + "\n\t")
                    v8 = v8
                    fmt.Println(v8)
                    _ = fi7
                    _ = v8
                }
            case "":
                if ((2147483648 >= 2) || true) {
                    for fi9 := 0; fi9 < 2; fi9++ {
                        v2 += v2
                        _ = fi9
                    }
                    var v10 float64 = 0.0
                    _ = v10
                }
            default:
                v11 := GenS3{F0: (v4 * v4)}
                _ = v11
            }
        }
        v1[65535] = 2147483648
    }
    v2 -= ((float32(0.0) + ((0.0 / 1.5) - v2)) / 0.1)
    for fi12 := 0; fi12 < 15; fi12++ {
        switch "\n\t" {
        case "foo bar":
            v13 := []int{32767, 9223372036854775807}
            _ = v13
        }
        if false {
            v14 := []int{(fi12 & (fi12 - -2147483648))}
            v2 -= v2
            v15 := v1[65535]
            _ = v14
            _ = v15
        }
        v16 := []int{fi12, 4294967295}
        _ = fi12
        _ = v16
    }
    v2 = v2
    fmt.Println(v2, v3)
    v17 := GenS1{F0: 9223372036854775807}
    if true {
        v18 := GenS1{F0: 42}
        _ = v18
    } else {
        v19 := GenS2{F0: 127, F1: 1}
        v4 = 569
        if false {
            switch "\n\t" {
            case "hello":
                if (-1 >= -2147483648) {
                    for fi20 := 0; fi20 < 12; fi20++ {
                        v2 -= (-2.25 - v2)
                        v21 := GenS4{F0: 42, F1: 100}
                        _ = fi20
                        _ = v21
                    }
                    fmt.Println(v4, v2, v2)
                    v22 := []int{}
                    _ = v22
                }
            case "a":
                v4 = (v4 / ((v4) | 1))
            }
        } else {
            v2 = 1.5
        }
        _ = v19
    }
    if ((0 != v4) || false) {
        var v23 uint = 18446744073709551615
        _ = v23
    }
    var v24 bool = true
    v2 = 0.0
    var v25 int64 = 149
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
