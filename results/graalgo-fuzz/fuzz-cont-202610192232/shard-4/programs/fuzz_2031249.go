package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS2 struct { F0 int16; F1 int32; F2 int64 }
type GenS3 struct { F0 int; F1 uint32; F2 uint8 }
func genFn0(p0 uint64, p1 int8) uint16 {
    if true {
        v4 := GenS1{F0: p0, F1: 2147483647, F2: p0}
        p0 += (p0 >> ((0) & 63))
        switch "\n\t" {
        case "foo bar":
            fmt.Println(p0, p1)
        }
        _ = v4
    }
    fmt.Println(p0, p1)
    var v5 int8 = (p1 % ((105) | 1))
    var v6 float64 = float64(3.14159)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 2
}

func main() {
    var v1 int32 = -1
    var v2 string = (("" + "hello") + ("" + ("hello" + (("a" + "") + ("foo bar" + ("hello" + (("foo bar" + "\n\t") + ("foo bar" + "a"))))))))
    var v3 uint64 = 18446744073709551615
    switch v2 {
    case "a":
        v1 = v1
    case "hello":
        fmt.Println(v3, v1)
    case "":
        v2 = v2
    }
    v7 := []int{-558, -862, 9223372036854775807, 283}
    switch (("\n\t" + ("a" + v2)) + v2) {
    case "":
        v8 := GenS2{F0: 1000, F1: v1, F2: -145}
        _ = v8
    default:
        var v9 int8 = -80
        _ = v9
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        fmt.Println(v3, v3, v1)
        _ = fi10
    }
    v11 := GenS2{F0: 1, F1: (v1 - -429), F2: -9223372036854775808}
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 int32 = (v1 & 42)
        _ = fi12
        _ = v13
    }
    var v14 uint32 = 65535
    var v15 int8 = 0
    for fi16 := 0; fi16 < 7; fi16++ {
        var v17 int32 = (v1 | v1)
        v18 := GenS3{F0: (fi16 * 42), F1: v14, F2: 39}
        v19 := GenS3{F0: (fi16 / ((4294967295) | 1)), F1: (^(v14)), F2: 74}
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    var v20 uint64 = v3
    if (!true) {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        for fi21 := 0; fi21 < 16; fi21++ {
            if ((0.0 + 1e10) < 1.5) {
                var v22 uint = uint(v14)
                _ = len(v7)
                v20 = (v20 & v20)
                _ = v22
            } else {
                v3 = (^(v20))
            }
            v23 := GenS3{F0: fi21, F1: v14, F2: 42}
            _ = fi21
            _ = v23
        }
    } else {
        var v24 uint = uint(v14)
        _ = v24
    }
    v25 := GenS2{F0: 0, F1: (v1 + v1), F2: 0}
    v15 += v15
    fmt.Println(v20, v20)
    for fi26 := 0; fi26 < 18; fi26++ {
        v2 = "a"
        _ = fi26
    }
    var v27 uint16 = 2
    var v28 uint8 = 128
    v29 := GenS3{F0: -456, F1: (^(v14)), F2: v28}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v7
}
