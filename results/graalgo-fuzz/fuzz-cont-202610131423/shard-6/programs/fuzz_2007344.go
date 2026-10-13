package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[int]string)
    var v2 float64 = 1.5
    v3 := []int{128}
    var v4 float32 = (-2.25 - 0.0)
    v5 := []int{2147483647}
    var v6 uint16 = 42
    var v7 uint64 = 18446744073709551615
    v4 += v4
    var v8 float32 = (float32(v2) - v4)
    fmt.Println(v6, v6)
    v9 := GenS1{F0: -32768, F1: (("" + ("a" + "a")) + "\n\t")}
    if false {
        v7 = v7
        fmt.Println(v8, v8, v6)
    }
    switch -21 {
    case -9223372036854775808:
        for fi10 := 0; fi10 < 13; fi10++ {
            v7 = 4294967295
            if (true && false) {
                for fi11 := 0; fi11 < 12; fi11++ {
                    var v12 bool = false
                    fmt.Println(v12, v7)
                    v13 := GenS2{F0: 836}
                    var v14 uint64 = v7
                    var v15 int64 = 9223372036854775807
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                    _ = v15
                }
                var v16 float64 = v2
                v17 := GenS1{F0: 1000, F1: "foo bar"}
                _ = v16
                _ = v17
            } else {
                var v18 float64 = float64(v4)
                _ = v18
            }
            _ = fi10
        }
    case 127:
        v19 := GenS1{F0: 129, F1: "hello"}
        _ = v19
    case -1:
        var v20 float64 = v2
        _ = v20
    }
    v8 = -2.25
    fmt.Println(v6, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
