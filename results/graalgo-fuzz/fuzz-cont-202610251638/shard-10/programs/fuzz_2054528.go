package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 string }

func main() {
    var v1 int16 = 0
    var v2 uint16 = 0
    var v3 uint8 = 255
    var v4 float64 = (0.0 * 1e10)
    v5 := make(map[uint32]string)
    v6 := make(map[string]bool)
    v6["\n\t"] = (0 <= -9223372036854775808)
    if (530 != -9223372036854775808) {
        fmt.Println(v2, v1)
        v1 -= (v1 >> ((18446744073709551615) & 63))
    } else {
        if true {
            fmt.Println(v1)
            switch "a" {
            case "\n\t":
                fmt.Println(v1, v4)
            case "foo bar":
                var v7 uint64 = 18446744073709551615
                _ = v7
            }
        } else {
            _ = len(v5)
            v8 := GenS1{F0: (!true), F1: -107, F2: ""}
            _ = v8
        }
    }
    switch 32767 {
    case -556:
        for fi9 := 0; fi9 < 6; fi9++ {
            switch (fi9 >> ((4294967295) & 63)) {
            case -2147483648:
                var v10 int = (fi9 << ((0) & 63))
                _ = v10
            case -713:
                var v11 float64 = v4
                _ = v11
            }
            for fi12 := 0; fi12 < 8; fi12++ {
                fmt.Println(v1, v2)
                var v13 uint = 162
                var v14 bool = (false && false)
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = fi9
        }
    case -9223372036854775808:
        var v15 int = 4294967295
        _ = v15
    default:
        if true {
            v2 += v2
        } else {
            v3 += v3
            v16 := []int{32767, 127, -9223372036854775808, 127}
            _ = v16
        }
    }
    if ((!true) || (!(!true))) {
        _ = len(v5)
    } else {
        var v17 int16 = -32768
        v1 = (v17 & v1)
        v4 -= v4
        _ = v17
    }
    v1 += (v1 ^ 32767)
    v4 += v4
    v1 = 736
    var v18 int8 = 42
    fmt.Println(v3, v18)
    var v19 string = "hello"
    if false {
        v20 := v5[2]
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
