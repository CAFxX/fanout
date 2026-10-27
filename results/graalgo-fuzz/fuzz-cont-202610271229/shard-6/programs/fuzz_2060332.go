package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int }
type GenS2 struct { F0 bool }
func genFn0(p0 int8) string {
    v4 := GenS1{F0: -9223372036854775808, F1: p0, F2: 1}
    v5 := GenS1{F0: 699, F1: (p0 / ((1) | 1)), F2: -2147483648}
    p0 += (p0 << ((42) & 63))
    _ = p0
    _ = v4
    _ = v5
    return "a"
}
func genFn1(p0 uint32, p1 int8) int {
    var v6 int64 = -91
    var v7 int16 = -32768
    p0 = 4294967295
    var v8 int8 = (^(p1))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 0
    v3 := make(map[string]string)
    v9 := []int{128, genFn1(65535, 42), 127}
    fmt.Println(v2)
    switch "" {
    case "\n\t":
        switch "a" {
        case "foo bar":
            if (true || (!true)) {
                v2 += v2
                v2 = v2
                v10 := GenS1{F0: 32768, F1: -87, F2: -1}
                _ = v10
            } else {
                var v11 uint64 = (^(v2))
                switch "a" {
                case "foo bar":
                    v12 := GenS1{F0: 0, F1: -54, F2: -469}
                    _ = v12
                case "hello":
                    switch "a" {
                    case "foo bar":
                        var v13 string = "foo bar"
                        _ = v13
                    }
                }
                _ = v11
            }
        }
    }
    var v14 string = ("a" + "\n\t")
    switch ("" + v14) {
    case "hello":
        var v15 uint = 255
        _ = v15
    case "":
        delete(v3, "foo bar")
    case "\n\t":
        v14 = (v14 + v14)
    default:
        v2 = 2
    }
    v2 = 2
    if (-128 >= -1) {
        _ = len(v3)
        v16 := GenS2{F0: (true || (float64((3.14159 + float32((float64(1.5) * 1.5)))) <= 1e10))}
        _ = v16
    } else {
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 uint8 = 222
            v14 = (v14 + v14)
            fmt.Println(v18)
            _ = fi17
            _ = v18
        }
        var v19 int16 = -32768
        switch (((((("hello" + "foo bar") + v14) + (v14 + "hello")) + "\n\t") + (v14 + ("foo bar" + (v14 + v14)))) + v14) {
        case "hello":
            if true {
                fmt.Println(v19, v19)
            }
        }
        _ = v19
    }
    v14 = v14
    fmt.Println(v2, v2)
    v20 := GenS1{F0: 4294967295, F1: 1, F2: 0}
    var v21 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v9
}
