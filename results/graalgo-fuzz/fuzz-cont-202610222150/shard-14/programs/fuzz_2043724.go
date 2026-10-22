package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint32 }
func genFn0(p0 uint32) int16 {
    var v5 bool = (!(128 == 0))
    var v6 string = "hello"
    _ = p0
    _ = v5
    _ = v6
    return 331
}

func main() {
    var v1 uint = 45
    v2 := make(map[string]string)
    var v3 uint32 = 2
    v4 := make(map[int]string)
    v3 += (v3 << ((v1) & 63))
    v1 += v1
    v7 := GenS1{F0: 127, F1: uint32(v1), F2: 743}
    var v8 bool = false
    var v9 int16 = 32767
    v10 := v4[9223372036854775807]
    fmt.Println(v9, v10)
    var v11 float32 = 0.0
    fmt.Println(v10)
    v12 := GenS1{F0: -1, F1: 2147483648, F2: 2147483647}
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 uint8 = 241
        if v8 {
            var v15 int64 = 0
            fmt.Println(v8, v15)
            if ((!true) && v8) {
                v10 = (v10 + (v10 + "foo bar"))
            }
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 float64 = (3.14159 / 1e10)
        _ = fi16
        _ = v17
    }
    v3 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
