package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }

func main() {
    v1 := []uint32{42, 2147483648, 0}
    var v2 int32 = 0
    var v3 float64 = (1.5 / 3.14159)
    var v4 int = 42
    v5 := make(map[int]uint32)
    var v6 int16 = 714
    v7 := v5[v4]
    var v8 int16 = int16(v2)
    var v9 uint8 = 74
    fmt.Println(v3, v6)
    switch ("\n\t" + "hello") {
    case "foo bar":
        v7 = (^(v7))
    case "\n\t":
        v5[9223372036854775807] = v7
    case "a":
        var v10 int8 = 100
        _ = v10
    }
    v7 -= v7
    v11 := GenS1{F0: 2, F1: 371}
    switch (v4 ^ v4) {
    case 1:
        for fi12 := 0; fi12 < 17; fi12++ {
            switch 1 {
            case -2147483648:
                v13 := GenS1{F0: 78, F1: 85}
                _ = v13
            case 4294967295:
                delete(v5, -923)
            case 777:
                var v14 int16 = int16(v7)
                _ = v14
            }
            var v15 uint = 18446744073709551615
            _ = fi12
            _ = v15
        }
    case -2147483648:
        v16 := v5[386]
        _ = v16
    default:
        v7 -= 49
    }
    var v17 int = v4
    v9 += (v9 + (v9 * (v9 * v9)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
