package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 string }

func main() {
    var v1 uint32 = 679
    var v2 string = "foo bar"
    var v3 float64 = -2.25
    var v4 int32 = -913
    var v5 uint64 = 4294967295
    var v6 uint32 = 65535
    v3 -= float64(float32((float64((1e10 + 1.5)) + float64(-2.25))))
    v2 = v2
    v3 = 3.14159
    v1 = (v1 * v6)
    v6 += (v1 - (^(v6)))
    fmt.Println(v1)
    v1 = v1
    var v7 int16 = -1
    v8 := GenS1{F0: -1, F1: 9223372036854775807, F2: v2}
    v7 -= (^(v7))
    switch 744 {
    case 956:
        v1 = 0
    case 537:
        var v9 uint16 = 479
        _ = v9
    default:
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 int8 = 0
            v12 := []int{4294967295}
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    var v13 uint32 = 753
    var v14 float32 = (0.1 * float32(0.1))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
