package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool; F1 int; F2 uint }
func genFn0(p0 uint16, p1 uint) string {
    var v7 uint8 = 105
    switch 322 {
    case -2147483648:
        v7 = v7
    }
    if false {
        var v8 uint8 = v7
        v9 := GenS1{F0: 0, F1: 0, F2: 2}
        _ = v8
        _ = v9
    } else {
        fmt.Println(p0, p0)
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = v7
    return "hello"
}

func main() {
    var v1 int16 = 1000
    v2 := []uint32{2147483647, 2147483648}
    var v3 uint = 65535
    v4 := []string{}
    v5 := []string{"\n\t"}
    var v6 int8 = 100
    var v10 int16 = (v1 + -1)
    switch 32767 {
    case 9223372036854775807:
        var v11 uint64 = 2
        _ = v11
    }
    v12 := GenS2{F0: v10}
    v3 -= (v3 << (((v3 << (((^(v3))) & 63))) & 63))
    v13 := GenS1{F0: 2147483648, F1: (-(v6)), F2: 65535}
    fmt.Println(v3, v3, v1)
    switch "foo bar" {
    case "\n\t":
        if true {
            fmt.Println(v10, v6, v6)
            if (int(v3) == 32767) {
                var v14 uint64 = 0
                v15 := GenS3{F0: 2, F1: 4294967295, F2: 42}
                _ = v14
                _ = v15
            } else {
                fmt.Println(v6, v6, v10)
            }
            switch -234 {
            case 1:
                var v16 uint = v3
                _ = v16
            case 4294967295:
                var v17 bool = true
                _ = v17
            }
        }
    case "hello":
        v18 := GenS4{F0: 32767}
        _ = v18
    case "a":
        v19 := GenS5{F0: false}
        _ = v19
    }
    _ = len(v4)
    v6 = (v6 - v6)
    if (0 == 42) {
        var v20 uint8 = 149
        _ = v20
    } else {
        v1 -= (-(v10))
    }
    v21 := GenS4{F0: 128}
    v10 = v1
    if len(v2) > 0 {
        v2[0] = 107
    }
    var v22 uint = (v3 << ((v3) & 63))
    var v23 uint8 = 88
    v24 := GenS6{F0: (v23 >= 25), F1: 9223372036854775807, F2: v3}
    v22 += (v22 / ((v22) | 1))
    v25 := GenS3{F0: 0, F1: 337, F2: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
