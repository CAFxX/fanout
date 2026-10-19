package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 string }
type GenS2 struct { F0 uint8; F1 bool; F2 uint64 }
type GenS3 struct { F0 int; F1 uint64; F2 int8 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 bool = (0 >= 42)
    var v3 float32 = 0.1
    var v4 uint32 = 2147483648
    if (1.5 > 3.14159) {
        v3 -= ((float32(3.14159) * float32((1.5 + float64(v3)))) - -2.25)
        v5 := GenS1{F0: 0, F1: 0, F2: ""}
        if v2 {
            var v6 bool = false
            switch "" {
            case "a":
                v7 := GenS2{F0: 1, F1: v6, F2: 521}
                _ = v7
            case "hello":
                v8 := []int{127}
                _ = v8
            }
            switch "a" {
            case "foo bar":
                v6 = (!false)
            case "hello":
                switch "" {
                case "foo bar":
                    v6 = (!(((v1 >> ((255) & 63)) <= v1) && (true && v2)))
                case "hello":
                    v9 := GenS2{F0: 42, F1: v6, F2: 2}
                    _ = v9
                }
            case "\n\t":
                v1 -= (-(v1))
            }
            _ = v6
        }
        _ = v5
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        v11 := []int{-9223372036854775808, (fi10 / (((fi10 & fi10)) | 1)), fi10}
        _ = fi10
        _ = v11
    }
    v1 += (v1 * (v1 | (^(v1))))
    v12 := GenS3{F0: 2147483647, F1: 2, F2: 1}
    switch 9223372036854775807 {
    case -595:
        fmt.Println(v3, v1)
    case 32767:
        var v13 uint64 = 42
        _ = v13
    }
    var v14 float64 = 0.0
    var v15 int8 = 42
    for fi16 := 0; fi16 < 2; fi16++ {
        v17 := []int{}
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
