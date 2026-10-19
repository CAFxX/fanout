package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 int16; F2 int32 }
type GenS4 struct { F0 uint32; F1 int; F2 uint }
type GenS5 struct { F0 uint8; F1 int32 }
func genFn0(p0 uint32, p1 bool, p2 uint64) int {
    v5 := []int{-855, 1}
    if len(v5) > 0 {
        v5[0] = 1
    }
    switch "" {
    case "a":
        if (42 != 65535) {
            switch ("" + "a") {
            case "\n\t":
                p0 = (^(p0))
            case "hello":
                p1 = p1
            case "a":
                var v6 uint = 162
                _ = v6
            }
            var v7 uint32 = uint32(p2)
            fmt.Println(p0)
            _ = v7
        } else {
            switch "hello" {
            case "":
                var v8 bool = (!p1)
                _ = v8
            case "hello":
                _ = len(v5)
            default:
                var v9 uint64 = (p2 | p2)
                _ = v9
            }
        }
    default:
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 string = "\n\t"
            for fi12 := 0; fi12 < 18; fi12++ {
                v13 := GenS1{F0: p0}
                _ = fi12
                _ = v13
            }
            if (p1 || true) {
                var v14 int64 = 42
                _ = len(v5)
                if (!p1) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                } else {
                    if p1 {
                        v15 := GenS2{F0: p1}
                        v11 = v11
                        for fi16 := 0; fi16 < 4; fi16++ {
                            v11 = v11
                            _ = fi16
                        }
                        _ = v15
                    } else {
                        p1 = p1
                    }
                    switch "hello" {
                    case "\n\t":
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case "foo bar":
                        for fi17 := 0; fi17 < 9; fi17++ {
                            p2 = 18446744073709551615
                            v18 := GenS2{F0: (!(!(((^(fi10)) > (fi10 / (((fi10 << ((uint(p0)) & 63))) | 1))) || p1)))}
                            _ = fi17
                            _ = v18
                        }
                    case "hello":
                        var v19 uint32 = (p0 / (((p0 - 4294967295)) | 1))
                        _ = v19
                    default:
                        var v20 uint64 = (p2 / (((p2 % (((p2 << ((673) & 63))) | 1))) | 1))
                        _ = v20
                    }
                    var v21 float32 = 1e10
                    _ = v21
                }
                _ = v14
            }
            _ = fi10
            _ = v11
        }
    }
    var v22 int8 = -1
    var v23 bool = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    _ = v5
    return -2147483648
}

func main() {
    var v1 int = -9223372036854775808
    var v2 string = ("" + "hello")
    var v3 uint64 = 18446744073709551615
    var v4 uint = 42
    v24 := []int{(v1 * v1), -1, -907, 0}
    v1 = genFn0(uint32(v4), false, v3)
    v3 = v3
    var v25 int32 = 1
    v24 = append(v24, v1)
    v26 := GenS3{F0: -96, F1: -32768, F2: 32}
    var v27 float64 = 1.5
    switch v2 {
    case "hello":
        for fi28 := 0; fi28 < 9; fi28++ {
            v29 := GenS4{F0: 2147483648, F1: (v1 / (((v1 >> (((v4 / (((^(v4))) | 1))) & 63))) | 1)), F2: (v4 ^ 4294967295)}
            switch 4294967295 {
            case 32768:
                var v30 uint32 = 2147483647
                _ = v30
            case 18:
                var v31 bool = true
                _ = v31
            case 42:
                v4 = (v4 + (v4 >> (((v4 + (v4 / ((379) | 1)))) & 63)))
            default:
                v2 = v2
            }
            if len(v24) >= 2 {
                v24 = v24[1:len(v24)]
            }
            _ = fi28
            _ = v29
        }
    default:
        v32 := GenS5{F0: 0, F1: (v25 & (v25 * 42))}
        _ = v32
    }
    v27 += float64(1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
