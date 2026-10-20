package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 int16 = -32768
    var v2 float32 = (float32(-2.25) - float32((3.14159 - 1e10)))
    var v3 uint32 = 2
    v4 := make(map[int]string)
    var v5 int32 = 2147483647
    if ((v5 | -2147483648) == v5) {
        if false {
            if (42 < 42) {
                var v6 int32 = (v5 % ((42) | 1))
                _ = v6
            }
        } else {
            v7 := []int{-1, -304}
            var v8 uint16 = 2
            _ = v7
            _ = v8
        }
        _ = len(v4)
        fmt.Println(v3)
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 uint64 = 18446744073709551615
            v3 = 4294967295
            v12 := []int{(fi10 / ((-458) | 1))}
            var v13 string = ("hello" + "foo bar")
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        var v14 int32 = v5
        _ = fi9
        _ = v14
    }
    var v15 int = 128
    var v16 float64 = float64((0.1 - 0.0))
    for fi17 := 0; fi17 < 10; fi17++ {
        _ = len(v4)
        var v18 float64 = v16
        _ = fi17
        _ = v18
    }
    if true {
        var v19 int16 = (v1 | -1)
        _ = v19
    }
    fmt.Println(v16, v5)
    fmt.Println(v3, v3)
    fmt.Println(v5, v1)
    var v20 uint64 = 18446744073709551615
    var v21 uint32 = (v3 ^ v3)
    v22 := GenS1{F0: (^(v20)), F1: v3}
    fmt.Println(v2, v1, v16)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
