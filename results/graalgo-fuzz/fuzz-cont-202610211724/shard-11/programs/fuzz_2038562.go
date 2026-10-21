package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint }
func genFn0(p0 int64) uint8 {
    v7 := []int{568, 0, 4294967295}
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 uint8 = 2
        v7 = append(v7, (fi8 * 9223372036854775807))
        switch 127 {
        case 128:
            fmt.Println(p0, v9, p0)
        case 2147483647:
            fmt.Println(fi8, p0)
        case 4294967295:
            fmt.Println(fi8, fi8, v9)
        }
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 16; fi11++ {
            if true {
                p0 += (p0 + p0)
                fmt.Println(p0)
            }
            p0 = (p0 % ((p0) | 1))
            _ = fi11
        }
        var v12 uint64 = 595
        switch "foo bar" {
        case "hello":
            switch -1 {
            case 1:
                p0 = p0
            }
        case "":
            var v13 float64 = (3.14159 / (3.14159 * 1e10))
            _ = v13
        case "foo bar":
            v14 := GenS1{F0: 473, F1: fi10}
            _ = v14
        }
        _ = fi10
        _ = v12
    }
    v15 := GenS1{F0: 2147483647, F1: -243}
    var v16 int16 = 1
    _ = p0
    _ = v15
    _ = v16
    _ = v7
    return 68
}
func genFn1(p0 int64, p1 bool, p2 uint) uint8 {
    fmt.Println(p0, p2)
    var v17 string = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return genFn0(-688)
}

func main() {
    v1 := make(map[int]string)
    var v2 float64 = (float64(float32(-2.25)) / 1.0)
    var v3 int16 = 32767
    var v4 int64 = 42
    v5 := make(map[string]string)
    var v6 uint64 = 1
    var v18 uint64 = v6
    switch ("foo bar" + "a") {
    case "":
        fmt.Println(v3, v3, v4)
    }
    switch (("" + "a") + "a") {
    case "a":
        if (!false) {
            if true {
                v19 := GenS2{F0: 42, F1: "\n\t"}
                v3 += (v3 | -32768)
                v18 -= 0
                _ = v19
            } else {
                v3 += v3
                v6 += (v18 >> ((1) & 63))
            }
        }
    case "\n\t":
        var v20 bool = (!false)
        _ = v20
    }
    var v21 int8 = 116
    var v22 int16 = v3
    v23 := GenS3{F0: 255}
    fmt.Println(v6)
    v2 = ((0.1 / 1.5) * (-2.25 * ((1e10 - v2) * v2)))
    v24 := GenS1{F0: 42, F1: 4294967295}
    fmt.Println(v18, v2)
    v3 -= 1
    fmt.Println(v21, v6, v22)
    v18 = v6
    v25 := GenS1{F0: -2147483648, F1: -85}
    var v26 uint32 = 65535
    fmt.Println(v18, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
