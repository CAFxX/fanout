package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int64; F1 uint16; F2 int32 }
type GenS3 struct { F0 int; F1 bool }

func main() {
    var v1 int16 = 42
    var v2 float64 = 1.5
    var v3 uint16 = 255
    v4 := make(map[string]int)
    if true {
        if true {
            var v5 int = 260
            _ = v5
        }
        v6 := []int{32768, -786, -9223372036854775808, 32767}
        var v7 string = "\n\t"
        _ = v6
        _ = v7
    } else {
        if true {
            var v8 uint32 = 4294967295
            _ = len(v4)
            var v9 int64 = 0
            _ = v8
            _ = v9
        }
        var v10 uint8 = 255
        v11 := GenS1{F0: 4294967295, F1: 87}
        _ = v10
        _ = v11
    }
    fmt.Println(v1)
    v12 := v4[("a" + ((((("\n\t" + "\n\t") + "hello") + "") + ("foo bar" + "foo bar")) + "a"))]
    v13 := GenS1{F0: 1, F1: 42}
    if false {
        var v14 int16 = (-(v1))
        _ = v14
    } else {
        switch (v12 | -2147483648) {
        case 9223372036854775807:
            v15 := GenS2{F0: 0, F1: v3, F2: 955}
            _ = v15
        case 42:
            v12 = (^(v12))
        }
        if true {
            var v16 float32 = 0.0
            var v17 int64 = 1
            _ = v16
            _ = v17
        } else {
            if (!(255 < 2)) {
                v12 = (v12 % ((v12) | 1))
                v18 := GenS1{F0: 447, F1: 1}
                fmt.Println(v12)
                _ = v18
            }
        }
    }
    v12 -= -677
    if (float32(0.0) <= 3.14159) {
        fmt.Println(v3, v12, v1)
    } else {
        for fi19 := 0; fi19 < 11; fi19++ {
            switch ("a" + "foo bar") {
            case "a":
                v1 += v1
            case "foo bar":
                var v20 int16 = v1
                _ = v20
            case "\n\t":
                v12 -= 127
            default:
                v21 := GenS3{F0: fi19, F1: (128 == 196)}
                _ = v21
            }
            _ = fi19
        }
    }
    var v22 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
