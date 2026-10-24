package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 bool = ((float32(-2.25) / float32(0.1)) >= (float32(0.1) * -2.25))
    var v2 string = (("\n\t" + ("foo bar" + "a")) + "")
    var v3 uint = 2
    v4 := make(map[int]bool)
    if v1 {
        v1 = true
    }
    if (2 >= 127) {
        v1 = false
        v5 := GenS1{F0: -9223372036854775808, F1: ((v3 <= (v3 & (^(v3)))) || (0.1 <= 1e10))}
        if (v1 || ((-2.25 / float64((0.1 - 1e10))) <= (3.14159 * 0.1))) {
            v1 = true
            delete(v4, 1)
        } else {
            for fi6 := 0; fi6 < 1; fi6++ {
                var v7 string = ((v2 + "hello") + v2)
                fmt.Println(v7, v1, v2)
                v1 = v1
                _ = fi6
                _ = v7
            }
            v2 = v2
        }
        _ = v5
    }
    v2 = v2
    v2 = v2
    v8 := []int{-9223372036854775808, 2147483647, 32767, 128}
    fmt.Println(v2)
    v1 = false
    var v9 int32 = -2147483648
    if true {
        if v1 {
            if (!(v1 || v1)) {
                var v10 int8 = -1
                _ = v10
            }
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
        if (true || v1) {
            var v11 float64 = (0.0 - (0.0 / (1.5 / -2.25)))
            v9 += v9
            fmt.Println(v1, v1, v9)
            _ = v11
        } else {
            switch 127 {
            case 331:
                switch v2 {
                case "foo bar":
                    var v12 bool = (0.0 > -2.25)
                    _ = v12
                case "hello":
                    switch ("a" + v2) {
                    case "a":
                        v13 := GenS1{F0: -9223372036854775808, F1: v1}
                        _ = v13
                    case "":
                        v14 := GenS1{F0: 541, F1: v1}
                        _ = v14
                    case "\n\t":
                        var v15 int32 = 325
                        _ = v15
                    default:
                        var v16 float32 = (1e10 / 1.5)
                        _ = v16
                    }
                case "\n\t":
                    v17 := GenS2{F0: 42}
                    _ = v17
                }
            }
        }
    }
    if v1 {
        fmt.Println(v9, v3)
    }
    switch 280 {
    case -471:
        if ((!false) || v1) {
            v3 = 42
        }
    case 0:
        v9 += (v9 + (v9 * 1))
    case 748:
        fmt.Println(v3, v9)
    }
    var v18 bool = (852 > 0)
    v8 = append(v8, 4294967295)
    var v19 float64 = float64(0.0)
    delete(v4, -2147483648)
    v9 = v9
    var v20 int = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
