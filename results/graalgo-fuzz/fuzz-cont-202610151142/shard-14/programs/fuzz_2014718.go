package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 uint; F1 int8; F2 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint64; F1 int64; F2 int32 }

func main() {
    v1 := []string{"a"}
    v2 := []uint32{42, 4294967295}
    var v3 uint8 = 56
    v3 = 42
    fmt.Println(v3)
    v1 = append(v1, "a")
    fmt.Println(v3)
    if false {
        var v4 uint64 = uint64(v3)
        v4 = (^(v4))
        _ = v4
    } else {
        v3 += v3
        v5 := GenS1{F0: 42}
        for fi6 := 0; fi6 < 14; fi6++ {
            v3 -= (^(v3))
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi6
        }
        _ = v5
    }
    v3 -= (v3 + 255)
    fmt.Println(v3)
    var v7 uint8 = v3
    var v8 uint64 = 42
    if (101 < 113) {
        v8 = 1
        v9 := GenS2{F0: "hello", F1: -32768}
        fmt.Println(v3)
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        if false {
            fmt.Println(v3, fi10, v7)
            fmt.Println(fi10, v7)
            v11 := GenS3{F0: 2147483647, F1: "a"}
            _ = v11
        } else {
            v12 := GenS1{F0: -128}
            v13 := GenS1{F0: 1}
            _ = v12
            _ = v13
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi10
    }
    v14 := GenS4{F0: 65535, F1: 100, F2: 4294967295}
    v15 := GenS1{F0: 1}
    var v16 uint8 = v7
    v17 := GenS5{F0: 18446744073709551615}
    if (!false) {
        if true {
            fmt.Println(v8, v3, v8)
            var v18 float64 = 3.14159
            v8 -= (v8 + 42)
            _ = v18
        } else {
            var v19 float64 = (-2.25 + 0.0)
            v19 -= v19
            _ = v19
        }
        v20 := GenS2{F0: ("hello" + "hello"), F1: 1000}
        v3 = 42
        _ = v20
    }
    if false {
        v21 := GenS6{F0: 4294967295, F1: int64(v8), F2: 1}
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
