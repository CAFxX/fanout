package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }

func main() {
    var v1 bool = (-970 >= 32767)
    v2 := []uint32{}
    v3 := make(map[int]string)
    var v4 float32 = 1e10
    switch (("\n\t" + "a") + "\n\t") {
    case "\n\t":
        v4 += v4
    default:
        var v5 int64 = -114
        _ = v5
    }
    delete(v3, -9223372036854775808)
    _ = len(v2)
    var v6 int64 = 9223372036854775807
    switch 2147483647 {
    case 945:
        var v7 bool = (!v1)
        _ = v7
    case 1:
        for fi8 := 0; fi8 < 11; fi8++ {
            if (-2147483648 != 2147483647) {
                for fi9 := 0; fi9 < 2; fi9++ {
                    if len(v2) > 0 {
                        v2[0] = 395
                    }
                    _ = fi9
                }
            }
            var v10 int8 = 1
            _ = fi8
            _ = v10
        }
    case 42:
        var v11 uint16 = 546
        _ = v11
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        fmt.Println(v6)
        _ = fi12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := GenS1{F0: (65535 != 2147483647), F1: ""}
        var v15 uint = 18446744073709551615
        _ = fi13
        _ = v14
        _ = v15
    }
    v4 -= (v4 - v4)
    var v16 bool = (32768 < -9223372036854775808)
    v6 = -9223372036854775808
    v17 := GenS1{F0: false, F1: "hello"}
    fmt.Println(v1, v1, v6)
    fmt.Println(v6, v6)
    fmt.Println(v1)
    if len(v2) > 0 {
        v2[0] = 65535
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
