package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
func genFn0(p0 uint) int16 {
    var v7 int16 = 32767
    v8 := []int{-731, -646, -9223372036854775808}
    _ = p0
    _ = v7
    _ = v8
    return (-(v7))
}

func main() {
    var v1 int16 = 42
    var v2 uint32 = 4294967295
    var v3 int8 = 42
    var v4 string = ("hello" + "foo bar")
    var v5 uint16 = 652
    v6 := make(map[int]string)
    v5 -= uint16(v2)
    v9 := GenS1{F0: 118, F1: v5}
    if (18446744073709551615 < uint(v5)) {
        fmt.Println(v4, v5)
    } else {
        if true {
            var v10 uint16 = v5
            _ = v10
        }
    }
    v11 := []int{-1, 9223372036854775807}
    if (v5 > (v5 >> ((126) & 63))) {
        for fi12 := 0; fi12 < 17; fi12++ {
            v2 -= 2147483647
            for fi13 := 0; fi13 < 16; fi13++ {
                var v14 uint16 = v5
                v15 := GenS1{F0: 646, F1: 65535}
                _ = fi13
                _ = v14
                _ = v15
            }
            v6[fi12] = v4
            _ = fi12
        }
        _ = len(v11)
    }
    fmt.Println(v3, v4, v4)
    fmt.Println(v4, v5)
    var v16 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
