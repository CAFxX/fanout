package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint8 = 255
    var v2 uint16 = 32767
    v3 := make(map[uint32]string)
    v4 := make(map[string]uint32)
    var v5 uint32 = uint32(v1)
    if false {
        var v6 float32 = (0.1 * 1e10)
        _ = v6
    } else {
        v7 := []int{-416, 2147483647}
        _ = v7
    }
    var v8 uint = 4294967295
    if ((((!true) && ((v1 << ((65535) & 63)) > v1)) && ("" == "foo bar")) || (0 >= -128)) {
        var v9 float64 = (3.14159 / -2.25)
        var v10 float32 = ((1.5 / (((3.14159 / 1.0) / 1e10) + (0.0 + ((1e10 / 3.14159) * 3.14159)))) / -2.25)
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        switch "foo bar" {
        case "hello":
            if (uint32(v1) <= 1) {
                v1 += v1
            } else {
                var v12 int8 = 42
                _ = v12
            }
        }
        v1 += v1
        switch "foo bar" {
        case "hello":
            fmt.Println(v8)
        case "":
            v8 = v8
        case "foo bar":
            var v13 int64 = -1
            _ = v13
        }
        _ = fi11
    }
    var v14 int32 = 894
    var v15 int = 286
    v16 := GenS1{F0: false}
    var v17 float32 = 1e10
    _ = len(v3)
    var v18 uint32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
