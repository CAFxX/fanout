package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 string }
func genFn0() uint8 {
    var v4 int32 = 42
    fmt.Println(v4)
    fmt.Println(v4)
    var v5 int64 = -9223372036854775808
    v6 := []int{-2147483648}
    _ = v4
    _ = v5
    _ = v6
    return 80
}

func main() {
    var v1 uint16 = 42
    var v2 float64 = -2.25
    v3 := make(map[string]bool)
    v2 -= v2
    var v7 uint64 = 18446744073709551615
    for fi8 := 0; fi8 < 5; fi8++ {
        fmt.Println(fi8, fi8)
        _ = fi8
    }
    v7 -= v7
    v9 := []int{-516, -680, -1}
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 15; fi11++ {
            v12 := GenS1{F0: 9223372036854775807, F1: 127, F2: ((("\n\t" + ((("foo bar" + ("hello" + "foo bar")) + ("\n\t" + "")) + "a")) + ("a" + "a")) + ("\n\t" + ("hello" + ("a" + ((("a" + ((("\n\t" + "\n\t") + "\n\t") + "a")) + "a") + "\n\t")))))}
            _ = fi11
            _ = v12
        }
        v2 = (float64(float32(0.1)) * (v2 - (v2 / (((1.5 / v2) / v2) - v2))))
        var v13 uint8 = 128
        _ = fi10
        _ = v13
    }
    fmt.Println(v2, v1)
    var v14 float64 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v7
    _ = v9
}
