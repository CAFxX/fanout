package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int64 = -1
    v2 := []int{-9223372036854775808, 128}
    v3 := []int{-9223372036854775808, 127, 533, 32767}
    v4 := make(map[int]bool)
    v5 := GenS1{F0: 42, F1: 18446744073709551615}
    v3 = append(v3, -9223372036854775808)
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 float64 = -2.25
        _ = fi6
        _ = v7
    }
    var v8 int8 = -116
    if (!false) {
        var v9 uint32 = 55
        v10 := GenS1{F0: -2147483648, F1: 1}
        var v11 int16 = -1
        _ = v10
        _ = v11
        _ = v9
    } else {
        _ = len(v4)
        if len(v3) > 0 {
            v3[0] = -2147483648
        }
        var v12 uint64 = 4294967295
        _ = v12
    }
    switch 42 {
    case 32767:
        v8 += 117
    case 42:
        switch "a" {
        case "\n\t":
            var v13 float32 = ((0.0 * 1.5) / 3.14159)
            _ = v13
        default:
            v4[1] = false
        }
    case 2147483647:
        v14 := GenS1{F0: -2147483648, F1: 65535}
        _ = v14
    default:
        v1 = 0
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 uint32 = 929
        v8 = (-(v8))
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        if false {
            var v18 int16 = -32768
            switch "\n\t" {
            case "foo bar":
                v19 := v4[fi17]
                _ = v19
            case "hello":
                var v20 int8 = v8
                _ = v20
            case "\n\t":
                fmt.Println(v8, v8)
            }
            _ = v18
        }
        v21 := GenS2{F0: 2147483647}
        _ = fi17
        _ = v21
    }
    fmt.Println(v8)
    var v22 int32 = -2147483648
    v8 = (v8 << ((65535) & 63))
    fmt.Println(v22, v8, v1)
    var v23 string = ("hello" + "hello")
    v8 = 127
    if false {
        fmt.Println(v8, v22)
        v4[2147483647] = (!(934 > v22))
        fmt.Println(v1)
    }
    var v24 uint64 = 1
    switch v23 {
    case "foo bar":
        v23 = v23
    case "\n\t":
        var v25 int8 = (v8 - v8)
        _ = v25
    case "":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v8 = (v8 ^ (v8 ^ (v8 & (v8 & v8))))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
