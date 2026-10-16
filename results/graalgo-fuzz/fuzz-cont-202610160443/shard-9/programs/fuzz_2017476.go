package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 uint16 }
func genFn0(p0 uint32) int16 {
    for fi5 := 0; fi5 < 2; fi5++ {
        for fi6 := 0; fi6 < 18; fi6++ {
            var v7 uint32 = p0
            p0 -= v7
            _ = fi6
            _ = v7
        }
        var v8 int = (fi5 ^ (fi5 % ((int(p0)) | 1)))
        var v9 int64 = 42
        _ = fi5
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        v11 := []int{(^(fi10)), fi10, -865, fi10}
        _ = fi10
        _ = v11
    }
    var v12 uint = 65535
    _ = p0
    _ = v12
    return 1000
}

func main() {
    var v1 string = ("" + "foo bar")
    v2 := make(map[uint32]string)
    var v3 uint = 42
    var v4 string = v1
    var v13 string = "hello"
    var v14 int64 = 42
    var v15 float64 = -2.25
    fmt.Println(v4, v1, v14)
    var v16 uint16 = 308
    v13 = v13
    v17 := GenS1{F0: 496, F1: v3}
    v14 -= (v14 - v14)
    switch -1 {
    case -9223372036854775808:
        v18 := []int{-605, 42, 2147483647}
        _ = v18
    case -2147483648:
        v19 := GenS1{F0: 741, F1: (v3 >> (((v3 >> ((65535) & 63))) & 63))}
        _ = v19
    }
    var v20 float64 = 1.5
    var v21 float32 = 3.14159
    v22 := GenS1{F0: 2147483647, F1: (v3 ^ (v3 ^ 4294967295))}
    if (false || true) {
        if false {
            v23 := []int{-661}
            fmt.Println(v14, v13)
            v14 = 284
            _ = v23
        } else {
            var v24 string = v13
            fmt.Println(v1, v21, v13)
            _ = v24
        }
        var v25 float64 = (float64(((((1e10 * float32((v15 / 1e10))) * (v21 * (float32(v20) * (0.1 + 0.1)))) - v21) * -2.25)) / 1.5)
        _ = v25
    }
    for fi26 := 0; fi26 < 7; fi26++ {
        v27 := GenS2{F0: 65535}
        _ = fi26
        _ = v27
    }
    fmt.Println(v3)
    var v28 string = v1
    if (!false) {
        v29 := GenS3{F0: (v14 & v14), F1: v16}
        v20 -= 1e10
        var v30 string = ("a" + "foo bar")
        _ = v29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v28
    _ = v3
    _ = v4
}
