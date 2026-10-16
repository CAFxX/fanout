package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint8, p1 int16) int {
    v6 := []int{-2147483648, 0, 176}
    v7 := GenS1{F0: 0, F1: 241, F2: 4294967295}
    var v8 uint16 = 726
    fmt.Println(v8, v8, v8)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return -944
}
func genFn1(p0 int, p1 string) int64 {
    var v9 int8 = 9
    v10 := []int{genFn0(30, -1)}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 0
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[string]bool)
    v3 := []int{-9223372036854775808}
    var v4 float32 = 1.5
    var v5 uint8 = 128
    switch ("hello" + "\n\t") {
    case "foo bar":
        if (-9223372036854775808 >= -332) {
            fmt.Println(v4, v1, v4)
            var v11 uint16 = 0
            var v12 string = "hello"
            _ = v11
            _ = v12
        }
    case "\n\t":
        if ((!(!(591 >= 255))) || true) {
            v13 := GenS1{F0: 1, F1: 0, F2: uint(v1)}
            var v14 float32 = ((v4 + float32(0.0)) * float32((0.1 - 1.5)))
            if len(v3) > 0 {
                v3[0] = -88
            }
            _ = v13
            _ = v14
        } else {
            var v15 uint8 = (v5 * 2)
            v15 = (v5 - 127)
            if (!((296 == -855) && true)) {
                var v16 float32 = 1e10
                v5 -= (v5 + (v5 & (v1 >> ((456) & 63))))
                _ = v16
            } else {
                if false {
                    switch 97 {
                    case 2147483647:
                        switch 853 {
                        case 176:
                            v17 := GenS2{F0: 59, F1: -2147483648}
                            _ = v17
                        }
                    case 128:
                        var v18 uint64 = 1
                        _ = v18
                    default:
                        v1 += v15
                    }
                    var v19 uint = 1
                    _ = v19
                }
                var v20 uint = 255
                fmt.Println(v20, v4, v4)
                _ = v20
            }
            _ = v15
        }
    case "":
        fmt.Println(v4, v5)
    default:
        if false {
            v21 := GenS2{F0: 9223372036854775807, F1: 1}
            _ = v21
        }
    }
    v22 := GenS1{F0: genFn1(32768, ("" + "a")), F1: 42, F2: 1}
    for fi23 := 0; fi23 < 12; fi23++ {
        if (510 > uint16(v1)) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        var v24 uint = 255
        v24 -= 42
        _ = fi23
        _ = v24
    }
    v1 += (v1 | v5)
    var v25 float64 = 3.14159
    v26 := GenS1{F0: -165, F1: 268, F2: 42}
    v27 := GenS3{F0: 0}
    if true {
        _ = len(v2)
        if ((2147483647 <= 65535) || (("hello" + "hello") >= "\n\t")) {
            fmt.Println(v25)
            v28 := GenS3{F0: (^(v1))}
            var v29 int64 = -749
            _ = v28
            _ = v29
        } else {
            fmt.Println(v1, v5, v25)
        }
    }
    var v30 int64 = -1
    v25 -= ((1e10 * ((1.5 + (v25 / ((((0.1 - (v25 - (v25 / (1.5 * v25)))) / 3.14159) + v25) - 3.14159))) / (3.14159 * v25))) * v25)
    fmt.Println(v1, v25)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
