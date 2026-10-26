package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 bool }

func main() {
    var v1 uint16 = 681
    var v2 float64 = ((0.0 * 3.14159) + float64(3.14159))
    var v3 float32 = 0.0
    var v4 uint = 4294967295
    var v5 int64 = 1
    var v6 int32 = int32(v4)
    fmt.Println(v3, v6)
    v3 -= v3
    if (0 == 1) {
        v4 += 4294967295
        fmt.Println(v2, v3)
        switch ("hello" + (("\n\t" + "\n\t") + "foo bar")) {
        case "\n\t":
            for fi7 := 0; fi7 < 9; fi7++ {
                if (!(true && true)) {
                    v8 := []int{fi7, fi7, -2147483648}
                    _ = v8
                }
                _ = fi7
            }
        case "a":
            if (v1 != v1) {
                for fi9 := 0; fi9 < 7; fi9++ {
                    if (1000 > -1) {
                        var v10 int8 = -1
                        _ = v10
                    } else {
                        var v11 float64 = (0.1 * (((v2 - ((v2 + v2) + v2)) / float64(v3)) - float64(3.14159)))
                        _ = v11
                    }
                    v6 = v6
                    switch "" {
                    case "a":
                        var v12 int = (fi9 - -1)
                        _ = v12
                    case "foo bar":
                        v13 := []int{}
                        _ = v13
                    }
                    _ = fi9
                }
                if ((0.0 - (0.0 + 1e10)) < v2) {
                    fmt.Println(v1, v5, v2)
                    if true {
                        var v14 bool = (!false)
                        v5 -= v5
                        _ = v14
                    } else {
                        v3 -= (v3 - (v3 - v3))
                        fmt.Println(v2)
                        v15 := []int{}
                        _ = v15
                    }
                    v16 := GenS1{F0: v6, F1: uint(v1)}
                    _ = v16
                } else {
                    v17 := GenS2{F0: (v6 << (((v4 ^ v4)) & 63)), F1: 133}
                    v18 := GenS3{F0: false}
                    _ = v17
                    _ = v18
                }
            }
        default:
            if (!("foo bar" >= "")) {
                fmt.Println(v3, v3)
                if (false || false) {
                    switch ("" + ("" + ("foo bar" + "\n\t"))) {
                    case "foo bar":
                        fmt.Println(v3)
                    case "hello":
                        fmt.Println(v4, v6, v1)
                    }
                } else {
                    v19 := []int{127, 4294967295}
                    _ = v19
                }
                v6 -= 1
            } else {
                var v20 float64 = (1.5 - 0.1)
                _ = v20
            }
        }
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        v2 -= (v2 * (((v2 * v2) * v2) + v2))
        v22 := GenS1{F0: 1, F1: 65535}
        _ = fi21
        _ = v22
    }
    var v23 float32 = v3
    fmt.Println(v1)
    v5 = (v5 + (v5 / ((9223372036854775807) | 1)))
    switch "hello" {
    case "":
        v6 = (v6 >> (((v4 << ((v4) & 63))) & 63))
    case "a":
        v24 := []int{}
        _ = v24
    }
    fmt.Println(v3, v1)
    v23 = (((float32(v2) / ((v3 * 1e10) * v3)) + ((v3 + v23) * v3)) - v3)
    v25 := []int{-9223372036854775808, 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
