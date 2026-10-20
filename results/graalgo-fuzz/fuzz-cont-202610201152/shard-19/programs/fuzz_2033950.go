package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 int32 }

func main() {
    var v1 uint8 = 135
    var v2 int32 = -1
    var v3 string = ("a" + "\n\t")
    if ((!true) || false) {
        var v4 float64 = 0.0
        fmt.Println(v3)
        if false {
            if (("a" >= "") || ((-(v2)) <= v2)) {
                v5 := []int{-528, 9223372036854775807}
                _ = v5
            } else {
                v6 := GenS1{F0: 2, F1: 0, F2: ((255 == 2) || (v2 == (v2 - -2147483648)))}
                _ = v6
            }
            switch "hello" {
            case "hello":
                switch 32767 {
                case 32767:
                    v3 = ((v3 + v3) + (("a" + "\n\t") + "hello"))
                case 42:
                    switch v3 {
                    case "foo bar":
                        fmt.Println(v4, v4)
                    case "hello":
                        var v7 uint16 = 729
                        _ = v7
                    }
                }
            case "\n\t":
                switch "a" {
                case "foo bar":
                    var v8 float64 = float64(1e10)
                    _ = v8
                default:
                    v9 := []int{2147483647, 4294967295, 42}
                    _ = v9
                }
            }
            v10 := []int{}
            _ = v10
        }
        _ = v4
    }
    fmt.Println(v3)
    fmt.Println(v2)
    v11 := GenS2{F0: -128}
    v12 := GenS3{F0: 1, F1: (^(v2))}
    v3 = v3
    fmt.Println(v3, v3, v2)
    if false {
        v2 = (v2 / ((v2) | 1))
    }
    var v13 int64 = -34
    v13 = v13
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
