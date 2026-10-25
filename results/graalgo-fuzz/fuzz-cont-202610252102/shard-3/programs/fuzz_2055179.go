package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := []uint32{}
    var v2 string = ("a" + "a")
    var v3 int16 = -1
    v4 := make(map[string]string)
    var v5 string = "a"
    var v6 float32 = 0.0
    var v7 uint8 = 14
    var v8 float64 = float64((v6 + 1.5))
    var v9 uint64 = 4294967295
    if (42 <= 128) {
        var v10 uint16 = 1
        switch v5 {
        case "a":
            var v11 uint = 65535
            _ = v11
        case "\n\t":
            for fi12 := 0; fi12 < 5; fi12++ {
                v7 -= v7
                var v13 int16 = v3
                v14 := GenS1{F0: 1}
                _ = fi12
                _ = v13
                _ = v14
            }
        default:
            var v15 float64 = (float64(v6) / (3.14159 + v8))
            _ = v15
        }
        _ = v10
    }
    fmt.Println(v5, v8)
    delete(v4, (v2 + "a"))
    v16 := GenS1{F0: -9223372036854775808}
    var v17 string = "\n\t"
    fmt.Println(v6, v9)
    if len(v1) > 0 {
        v1[0] = 0
    }
    var v18 int8 = -1
    v6 -= (float32(-2.25) - (0.0 - v6))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
