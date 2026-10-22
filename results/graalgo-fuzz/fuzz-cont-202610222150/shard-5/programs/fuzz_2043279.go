package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int32 }

func main() {
    v1 := make(map[int]string)
    v2 := []int{9223372036854775807}
    var v3 bool = true
    v3 = v3
    v3 = v3
    v3 = (-2147483648 < 9223372036854775807)
    var v4 string = ((("hello" + "") + ("\n\t" + "\n\t")) + "hello")
    fmt.Println(v4, v3)
    if ((-1 == 2147483647) || v3) {
        v5 := GenS1{F0: 32767, F1: v3, F2: 2147483647}
        _ = v5
    }
    delete(v1, 2147483647)
    var v6 uint16 = 32767
    v6 = v6
    v6 = v6
    var v7 uint64 = 18446744073709551615
    for fi8 := 0; fi8 < 3; fi8++ {
        v4 = v4
        var v9 int8 = 42
        var v10 int8 = v9
        _ = fi8
        _ = v10
        _ = v9
    }
    v3 = true
    if v3 {
        if v3 {
            fmt.Println(v4, v6)
        } else {
            v11 := GenS1{F0: -32768, F1: false, F2: -1}
            for fi12 := 0; fi12 < 17; fi12++ {
                var v13 uint8 = 2
                v14 := GenS1{F0: 317, F1: v3, F2: -807}
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = v11
        }
        var v15 int8 = 100
        v1[-2147483648] = v4
        _ = v15
    } else {
        var v16 float32 = 0.0
        fmt.Println(v4, v16, v16)
        var v17 int = 128
        _ = v16
        _ = v17
    }
    v3 = (32768 == 32767)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
