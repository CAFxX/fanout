package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 int; F1 string; F2 int64 }

func main() {
    var v1 int8 = 100
    var v2 int16 = 363
    var v3 int16 = (v2 * (v2 ^ v2))
    v4 := make(map[string]int)
    v5 := make(map[int]uint32)
    switch ("foo bar" + ("foo bar" + "a")) {
    case "a":
        v6 := GenS1{F0: 392}
        _ = v6
    case "foo bar":
        v7 := []int{-355, -9223372036854775808, 32767}
        _ = v7
    }
    fmt.Println(v1, v3, v2)
    var v8 bool = (1 >= 2)
    var v9 float32 = 3.14159
    var v10 int = 128
    var v11 uint8 = 127
    v12 := GenS2{F0: (v10 + (v10 % (((v10 << ((1) & 63))) | 1))), F1: (v11 + (v11 ^ 0))}
    switch 128 {
    case 42:
        delete(v5, 2147483647)
    }
    var v13 int = (^(v10))
    v4[(("a" + ("hello" + "")) + "")] = (v10 * v13)
    var v14 string = ("a" + "")
    switch ("hello" + (v14 + v14)) {
    case "foo bar":
        if (!(!false)) {
            fmt.Println(v11, v8, v11)
            _ = len(v5)
        }
    default:
        fmt.Println(v10)
    }
    switch (("hello" + v14) + v14) {
    case "a":
        if v8 {
            for fi15 := 0; fi15 < 9; fi15++ {
                var v16 bool = v8
                var v17 int = fi15
                _ = fi15
                _ = v16
                _ = v17
            }
        } else {
            fmt.Println(v1)
        }
    default:
        if (!(!(((0.1 - 0.1) / float64((v9 * ((1.5 / v9) - v9)))) != float64((0.0 / 1.0))))) {
            for fi18 := 0; fi18 < 10; fi18++ {
                v19 := []int{}
                v9 = (3.14159 / ((v9 + float32((((1e10 * 0.0) - 0.0) - float64(1.5)))) * (3.14159 - (3.14159 + float32(1.5)))))
                _ = fi18
                _ = v19
            }
            v20 := GenS3{F0: v13, F1: v14, F2: -877}
            _ = v20
        } else {
            var v21 uint16 = 1
            v22 := []int{507, (v10 % ((v13) | 1))}
            _ = v21
            _ = v22
        }
    }
    if (151 < 2) {
        switch v14 {
        case "hello":
            v3 = (v2 * v3)
        case "\n\t":
            if v8 {
                var v23 int64 = 9223372036854775807
                var v24 float64 = (-2.25 - 3.14159)
                var v25 uint8 = 127
                _ = v23
                _ = v24
                _ = v25
            }
        case "a":
            var v26 uint = 247
            _ = v26
        }
    }
    v14 = v14
    switch v14 {
    case "hello":
        var v27 bool = (false && v8)
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
