package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int64 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint32; F1 int16 }
func genFn0(p0 uint32) uint8 {
    p0 = (p0 ^ p0)
    p0 += (^(p0))
    v7 := []int{32768, 32768, 32768}
    _ = p0
    _ = v7
    return uint8(p0)
}
func genFn1(p0 int, p1 string) uint16 {
    v8 := GenS1{F0: 1, F1: 61, F2: 550}
    if false {
        v9 := GenS1{F0: 204, F1: -128, F2: 317}
        _ = v9
    } else {
        switch ("\n\t" + p1) {
        case "\n\t":
            var v10 int32 = -2147483648
            _ = v10
        case "hello":
            if false {
                if true {
                    for fi11 := 0; fi11 < 11; fi11++ {
                        p0 += p0
                        _ = fi11
                    }
                }
                fmt.Println(p0)
            } else {
                var v12 int64 = -9223372036854775808
                for fi13 := 0; fi13 < 20; fi13++ {
                    v14 := []int{}
                    for fi15 := 0; fi15 < 16; fi15++ {
                        v16 := GenS1{F0: 104, F1: 11, F2: 9223372036854775807}
                        var v17 uint = 255
                        var v18 int32 = -609
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    _ = fi13
                    _ = v14
                }
                p1 = p1
                _ = v12
            }
        case "a":
            v19 := []int{(p0 >> ((0) & 63)), p0}
            _ = v19
        default:
            var v20 float32 = 0.0
            _ = v20
        }
        v21 := GenS2{F0: p0, F1: 694}
        _ = v21
    }
    if (-128 != 76) {
        p0 = -866
        var v22 float64 = 0.1
        _ = v22
    }
    _ = p0
    _ = p1
    _ = v8
    return 255
}

func main() {
    var v1 int32 = -2147483648
    var v2 uint8 = 42
    var v3 uint64 = 810
    var v4 int8 = 7
    var v5 uint16 = 1
    var v6 float64 = float64(-2.25)
    var v23 bool = false
    switch (("foo bar" + ("a" + "foo bar")) + "foo bar") {
    case "hello":
        for fi24 := 0; fi24 < 9; fi24++ {
            v25 := []int{0, fi24, 477}
            var v26 string = "\n\t"
            _ = fi24
            _ = v25
            _ = v26
        }
    }
    var v27 uint32 = 353
    v28 := []int{-41}
    v29 := GenS3{F0: (v27 & 29), F1: 42}
    v2 += 1
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    switch "foo bar" {
    case "foo bar":
        var v30 int8 = (v4 - 42)
        _ = v30
    default:
        var v31 int32 = (v1 + (v1 % (((v1 + v1)) | 1)))
        _ = v31
    }
    var v32 string = (("foo bar" + "a") + ("foo bar" + "foo bar"))
    var v33 float32 = float32((-2.25 - 3.14159))
    _ = len(v28)
    v23 = false
    switch 9223372036854775807 {
    case 44:
        _ = len(v28)
    case -9223372036854775808:
        var v34 uint64 = v3
        _ = v34
    case -865:
        fmt.Println(v4, v23, v2)
    default:
        var v35 uint = 2
        _ = v35
    }
    _ = len(v28)
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
