package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int64 }

func main() {
    var v1 int64 = 0
    var v2 int = 2147483647
    v3 := make(map[string]uint32)
    var v4 int64 = (^(v1))
    v5 := GenS1{F0: v1, F1: 2147483647, F2: -9223372036854775808}
    var v6 string = (("foo bar" + "foo bar") + "hello")
    var v7 int16 = -191
    v4 = (v4 - (-(v1)))
    if false {
        switch "" {
        case "hello":
            var v8 int = (v2 ^ (-(v2)))
            _ = v8
        case "foo bar":
            v3[((v6 + ("\n\t" + "foo bar")) + (v6 + (((v6 + (v6 + v6)) + v6) + v6)))] = 0
        default:
            if true {
                fmt.Println(v4, v6, v4)
                var v9 float32 = 3.14159
                _ = v9
            } else {
                v4 = v4
                v10 := []int{(-(v2))}
                if true {
                    var v11 uint8 = 127
                    var v12 int64 = -9223372036854775808
                    _ = v11
                    _ = v12
                } else {
                    var v13 int64 = v4
                    _ = v13
                }
                _ = v10
            }
        }
        fmt.Println(v4)
    }
    var v14 uint16 = 952
    var v15 string = ("hello" + v6)
    fmt.Println(v2)
    v1 = (v1 * 789)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
