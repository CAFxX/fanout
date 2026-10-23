package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint32 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 uint8; F1 uint64; F2 uint }
func genFn0(p0 int16, p1 int, p2 int) int {
    var v6 bool = (((0.1 * 1.5) <= float32(1e10)) && true)
    var v7 int8 = 96
    var v8 int8 = 51
    p0 -= (p0 % (((p0 >> ((2) & 63))) | 1))
    switch (p2 / (((p2 / (((p1 | (p2 ^ -2147483648))) | 1))) | 1)) {
    case 127:
        var v9 float64 = 0.1
        _ = v9
    case 32767:
        if true {
            if false {
                switch p1 {
                case 42:
                    var v10 int64 = 9223372036854775807
                    _ = v10
                case 214:
                    var v11 int8 = -128
                    _ = v11
                }
            } else {
                p1 = (p1 % ((p2) | 1))
                v12 := GenS1{F0: ((("a" + "\n\t") + "hello") + "a"), F1: 1, F2: 42}
                _ = v12
            }
            var v13 string = "hello"
            var v14 float64 = (float64(3.14159) / ((1.5 * (0.1 + 1.5)) / 3.14159))
            _ = v13
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return p2
}

func main() {
    var v1 uint64 = 652
    var v2 uint = 1
    v3 := make(map[int]int)
    v4 := []string{((("hello" + "hello") + ("hello" + "a")) + "foo bar"), (("a" + "hello") + "\n\t"), "a", ("" + "\n\t")}
    var v5 uint = v2
    var v15 int = 4294967295
    var v16 uint16 = 255
    v17 := GenS2{F0: 884, F1: 663}
    v15 = (v15 | 270)
    if true {
        fmt.Println(v1, v2)
        fmt.Println(v5, v15)
        v18 := GenS3{F0: -32768, F1: v16}
        _ = v18
    }
    v16 = 65535
    fmt.Println(v1, v2)
    v15 -= v15
    switch "a" {
    case "":
        if (-1 <= int8(v15)) {
            var v19 float32 = -2.25
            var v20 uint = v5
            _ = v19
            _ = v20
        }
    case "hello":
        v21 := GenS4{F0: 2, F1: 1}
        _ = v21
    case "a":
        if ((2147483648 < 42) && true) {
            fmt.Println(v2, v2)
            v22 := GenS5{F0: 1, F1: v1, F2: (v2 | v2)}
            _ = v22
        }
    }
    var v23 uint16 = (v16 * (v16 + 32767))
    var v24 uint8 = 42
    if true {
        v15 += (v15 | (v15 & v15))
        var v25 int64 = -9223372036854775808
        v25 = (v25 * (v25 * v25))
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
