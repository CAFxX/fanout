package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 bool, p1 int32) int {
    var v6 int = 9223372036854775807
    switch (v6 % ((v6) | 1)) {
    case 32767:
        p1 = (p1 | (p1 & p1))
    case 128:
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 int8 = 114
            _ = fi7
            _ = v8
        }
    default:
        v9 := []int{}
        _ = v9
    }
    v10 := []int{v6}
    v11 := GenS1{F0: "hello"}
    v10 = append(v10, (v6 / (((v6 & v6)) | 1)))
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v6
    return (^(v6))
}
func genFn1() string {
    var v12 uint32 = 4294967295
    v12 = v12
    var v13 int16 = -32768
    _ = v12
    _ = v13
    return "hello"
}

func main() {
    var v1 uint64 = 289
    v2 := make(map[uint32]string)
    var v3 float64 = ((((((1.5 * float64(1.5)) * ((((0.1 - (0.1 * 1e10)) / (0.1 / 1.5)) / ((1e10 * (0.1 + 1.5)) / 1e10)) / 1.0)) * float64(3.14159)) + -2.25) + -2.25) - 1.5)
    v4 := make(map[string]bool)
    v5 := []string{}
    v14 := v2[2147483648]
    v1 -= (v1 & 18446744073709551615)
    if true {
        var v15 int8 = 1
        var v16 bool = true
        var v17 float64 = v3
        _ = v15
        _ = v16
        _ = v17
    }
    v3 += v3
    if true {
        fmt.Println(v3)
        fmt.Println(v3, v14, v1)
    } else {
        if false {
            fmt.Println(v1, v1)
            var v18 uint8 = 144
            _ = v18
        } else {
            if (false || true) {
                v3 = v3
                var v19 int = 257
                _ = v19
            }
            var v20 uint = 1
            if false {
                delete(v4, (v14 + v14))
                var v21 uint32 = 2147483647
                v4[v14] = true
                _ = v21
            }
            _ = v20
        }
        v14 = "hello"
    }
    switch v14 {
    case "a":
        v22 := GenS1{F0: "foo bar"}
        _ = v22
    case "":
        if (2 <= 55) {
            if true {
                var v23 uint64 = (v1 % (((v1 & (v1 / ((18446744073709551615) | 1)))) | 1))
                fmt.Println(v3, v23, v3)
                delete(v2, 146)
                _ = v23
            } else {
                _ = len(v4)
                _ = len(v4)
                fmt.Println(v14, v1, v1)
            }
        } else {
            v1 -= (^(v1))
            v1 = 1
            fmt.Println(v14, v1)
        }
    case "\n\t":
        var v24 uint8 = 128
        _ = v24
    }
    var v25 uint = 18446744073709551615
    if true {
        var v26 uint16 = 2
        _ = v26
    } else {
        _ = len(v5)
        fmt.Println(v3, v14, v14)
        var v27 float64 = -2.25
        _ = v27
    }
    v25 += 1
    v3 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
