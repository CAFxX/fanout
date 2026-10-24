package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
func genFn0(p0 uint, p1 int64, p2 int32) uint8 {
    switch "hello" {
    case "hello":
        var v6 int = 9223372036854775807
        _ = v6
    case "\n\t":
        v7 := []int{127}
        _ = v7
    default:
        switch "\n\t" {
        case "":
            fmt.Println(p2, p1)
        case "foo bar":
            v8 := GenS1{F0: 2147483647, F1: -128}
            _ = v8
        default:
            fmt.Println(p2, p1, p0)
        }
    }
    v9 := []int{2147483647, 128, -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 217
}
func genFn1(p0 int32, p1 int) uint16 {
    v10 := []int{}
    p0 += p0
    for fi11 := 0; fi11 < 16; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 int32 = (^(p0))
            if len(v10) > 0 {
                v10[0] = p1
            }
            var v14 int32 = 2147483647
            _ = fi12
            _ = v13
            _ = v14
        }
        p1 = fi11
        p1 -= 13
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    return 2
}

func main() {
    v1 := []string{"foo bar", "hello", (("a" + "") + "hello"), "a"}
    v2 := make(map[string]string)
    var v3 int16 = 472
    v4 := []string{"", "\n\t", "a", (((("foo bar" + "a") + "foo bar") + "foo bar") + "a")}
    var v5 uint = 18446744073709551615
    var v15 bool = true
    v3 -= v3
    var v16 float32 = ((0.1 - 1e10) * (1e10 * (1.5 / 1.0)))
    v17 := GenS1{F0: 4294967295, F1: -26}
    if (false && true) {
        var v18 string = ""
        switch 4294967295 {
        case 32767:
            v5 = 4294967295
        }
        _ = v18
    }
    switch 127 {
    case 1:
        if false {
            var v19 uint32 = 1
            fmt.Println(v3)
            _ = v19
        }
    case -1:
        fmt.Println(v3, v16)
    }
    v3 += v3
    fmt.Println(v5)
    var v20 bool = v15
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
