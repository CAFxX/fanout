package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int16; F2 int8 }
type GenS3 struct { F0 uint }

func main() {
    var v1 int32 = 2147483647
    var v2 int8 = -128
    var v3 string = "a"
    var v4 uint64 = 402
    var v5 uint64 = (v4 * v4)
    if ("hello" >= (v3 + ("\n\t" + v3))) {
        fmt.Println(v2)
        v3 = "foo bar"
        fmt.Println(v5, v5)
    } else {
        var v6 int16 = 32767
        fmt.Println(v1, v3)
        _ = v6
    }
    v5 += 4294967295
    v7 := GenS1{F0: (v2 & v2)}
    var v8 float64 = -2.25
    switch "hello" {
    case "":
        if (222 < 127) {
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 string = (v3 + (((v3 + v3) + ("hello" + (v3 + v3))) + "hello"))
                _ = fi9
                _ = v10
            }
            switch "foo bar" {
            case "":
                var v11 int64 = 9223372036854775807
                _ = v11
            case "foo bar":
                v5 = (v5 ^ v5)
            case "\n\t":
                fmt.Println(v4, v8)
            default:
                v12 := GenS1{F0: v2}
                _ = v12
            }
        }
    case "\n\t":
        v8 += (v8 * v8)
    case "a":
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 uint = 18446744073709551615
            for fi15 := 0; fi15 < 14; fi15++ {
                var v16 float64 = (v8 - v8)
                v17 := GenS1{F0: v2}
                var v18 float64 = v16
                v1 = v1
                fmt.Println(v1)
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            _ = fi13
            _ = v14
        }
    }
    if (((v3 + (v3 + v3)) + ((v3 + ((v3 + (v3 + v3)) + ("a" + v3))) + "hello")) != "\n\t") {
        fmt.Println(v4, v5)
        v19 := GenS2{F0: 42, F1: 1, F2: (v2 % ((0) | 1))}
        _ = v19
    }
    switch v3 {
    case "hello":
        var v20 uint16 = 292
        _ = v20
    case "\n\t":
        for fi21 := 0; fi21 < 9; fi21++ {
            v22 := []int{fi21, 1}
            v23 := GenS3{F0: 18446744073709551615}
            fmt.Println(v3)
            _ = fi21
            _ = v22
            _ = v23
        }
    default:
        switch 4294967295 {
        case 801:
            switch "foo bar" {
            case "hello":
                if (false && true) {
                    if (-810 >= 26) {
                        v1 = (v1 % ((v1) | 1))
                    }
                    fmt.Println(v4, v2)
                    v24 := []int{1, 974, 2147483647}
                    _ = v24
                }
            case "\n\t":
                v25 := GenS2{F0: 937, F1: 1000, F2: -128}
                _ = v25
            case "":
                v2 = (-(v2))
            }
        case 42:
            v26 := []int{127}
            _ = v26
        case 128:
            var v27 int = 4294967295
            _ = v27
        }
    }
    v8 += 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
