package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
func genFn0(p0 int8, p1 uint32) uint32 {
    if (false || false) {
        fmt.Println(p1, p0)
        fmt.Println(p0)
        p1 = p1
    } else {
        var v4 uint = 88
        _ = v4
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return p1
}
func genFn1(p0 bool, p1 uint32) bool {
    switch 32767 {
    case 289:
        v5 := GenS1{F0: 2, F1: -370}
        _ = v5
    case 650:
        var v6 string = ""
        _ = v6
    default:
        v7 := []int{-2147483648, -2147483648}
        _ = v7
    }
    var v8 string = ("foo bar" + "\n\t")
    var v9 int8 = 0
    var v10 uint64 = 2
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return true
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[int]string)
    var v3 int32 = -2147483648
    fmt.Println(v3, v1)
    var v11 int8 = 33
    v12 := GenS1{F0: 772, F1: v3}
    v11 += 42
    if (0 >= 533) {
        var v13 uint = 766
        _ = v13
    } else {
        switch "hello" {
        case "a":
            var v14 int16 = 1000
            _ = v14
        case "":
            fmt.Println(v11, v1)
        case "\n\t":
            switch "hello" {
            case "foo bar":
                var v15 int16 = -32768
                _ = v15
            case "hello":
                var v16 float64 = (float64(-2.25) * 1e10)
                _ = v16
            default:
                fmt.Println(v11)
            }
        default:
            if (!true) {
                var v17 string = ""
                _ = v17
            }
        }
    }
    v3 = v3
    v18 := []int{807}
    var v19 int64 = 0
    v3 = v3
    fmt.Println(v11, v19)
    _ = len(v18)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
