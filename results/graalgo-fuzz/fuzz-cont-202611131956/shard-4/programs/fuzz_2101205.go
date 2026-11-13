package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 string }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint, p1 int8, p2 int8) uint {
    v7 := []int{-137, 4294967295, 127, -2147483648}
    if true {
        p0 = (p0 * (p0 * p0))
    }
    v8 := GenS1{F0: 1, F1: false, F2: "\n\t"}
    if (18446744073709551615 > 2) {
        var v9 bool = (0 == 42)
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p0 & (p0 + p0))
}

func main() {
    var v1 string = ""
    var v2 uint64 = 4294967295
    v3 := make(map[string]int)
    v4 := make(map[uint32]string)
    var v5 uint = 384
    var v6 int32 = -374
    var v10 uint = v5
    fmt.Println(v2)
    for fi11 := 0; fi11 < 4; fi11++ {
        v10 = v10
        _ = fi11
    }
    switch v1 {
    case "\n\t":
        v10 -= (v5 / ((genFn0((v5 / ((v10) | 1)), 127, 100)) | 1))
    case "":
        v12 := []int{9223372036854775807}
        _ = v12
    default:
        switch (v1 + "a") {
        case "foo bar":
            for fi13 := 0; fi13 < 6; fi13++ {
                fmt.Println(fi13, v1)
                _ = fi13
            }
        }
    }
    v2 = (^(v2))
    v1 = ("\n\t" + (v1 + v1))
    var v14 string = v1
    switch -313 {
    case 732:
        fmt.Println(v10, v6)
    case 42:
        v15 := GenS1{F0: 42, F1: false, F2: "\n\t"}
        _ = v15
    case 32767:
        v16 := GenS2{F0: 255, F1: 0, F2: (v2 ^ 4294967295)}
        _ = v16
    default:
        var v17 int32 = v6
        _ = v17
    }
    if (-908 < 32768) {
        for fi18 := 0; fi18 < 14; fi18++ {
            v10 += v5
            var v19 uint64 = (v2 - v2)
            var v20 string = "hello"
            _ = fi18
            _ = v19
            _ = v20
        }
        v10 += genFn0(v5, 127, 100)
    } else {
        for fi21 := 0; fi21 < 1; fi21++ {
            v22 := []int{fi21, fi21}
            fmt.Println(v14, fi21)
            var v23 float32 = 1e10
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    v24 := GenS3{F0: (^(v2)), F1: false}
    v25 := GenS4{F0: v6}
    v3[v14] = 128
    v14 = (v1 + ((("" + "foo bar") + "hello") + (v1 + v1)))
    v1 = (v14 + v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
