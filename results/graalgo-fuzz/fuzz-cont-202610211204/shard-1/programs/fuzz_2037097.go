package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 bool }

func main() {
    var v1 bool = false
    var v2 uint64 = 591
    var v3 int64 = 9223372036854775807
    var v4 int32 = 0
    var v5 int64 = v3
    v6 := make(map[int]string)
    fmt.Println(v1, v1, v1)
    var v7 int8 = -119
    _ = len(v6)
    switch 32767 {
    case -593:
        var v8 uint32 = 440
        _ = v8
    case 54:
        var v9 string = ("foo bar" + "a")
        _ = v9
    default:
        switch "hello" {
        case "":
            if v1 {
                v3 = v5
            } else {
                v7 -= v7
                var v10 float32 = 0.1
                _ = v10
            }
        }
    }
    switch "\n\t" {
    case "a":
        var v11 uint32 = 4294967295
        _ = v11
    default:
        if ("a" == ("" + "foo bar")) {
            if v1 {
                var v12 int64 = v3
                v13 := GenS1{F0: (v3 | (v5 + (v5 & (v5 % ((9223372036854775807) | 1))))), F1: (((v4 ^ (v4 | v4)) == (v4 ^ -2147483648)) || (((1.5 - 3.14159) - float32(3.14159)) <= (1.5 - 0.0))), F2: (!false)}
                _ = v12
                _ = v13
            } else {
                fmt.Println(v4, v3, v5)
                v4 = (v4 - v4)
            }
            v2 += v2
            v3 += (v3 * v5)
        } else {
            v5 = (v3 & -1)
            v14 := v6[1]
            _ = v14
        }
    }
    switch "\n\t" {
    case "\n\t":
        if ((!((v2 | (v2 << ((767) & 63))) != (v2 << ((486) & 63)))) && false) {
            v1 = false
            v15 := GenS1{F0: v3, F1: v1, F2: false}
            _ = v15
        } else {
            var v16 int64 = v3
            _ = v16
        }
    case "a":
        for fi17 := 0; fi17 < 8; fi17++ {
            fmt.Println(v1, v2)
            v18 := []int{fi17, fi17, fi17, 9223372036854775807}
            _ = fi17
            _ = v18
        }
    case "foo bar":
        var v19 uint32 = 2147483648
        _ = v19
    }
    fmt.Println(v7, v3, v5)
    var v20 int8 = v7
    v4 = (v4 ^ (v4 / ((v4) | 1)))
    v21 := []int{-2147483648, -1, 9223372036854775807}
    _ = len(v21)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
