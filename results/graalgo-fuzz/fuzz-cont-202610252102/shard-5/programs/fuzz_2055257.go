package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 bool = (("a" + ("a" + "hello")) == "\n\t")
    var v2 float64 = 1e10
    v3 := make(map[int]string)
    var v4 uint32 = 815
    var v5 float64 = (v2 / v2)
    _ = len(v3)
    fmt.Println(v2, v1)
    v2 += 0.1
    switch "\n\t" {
    case "a":
        var v6 float32 = 3.14159
        _ = v6
    case "hello":
        if v1 {
            v4 += 65535
            v7 := GenS1{F0: 0, F1: 129}
            _ = v7
        }
    default:
        for fi8 := 0; fi8 < 10; fi8++ {
            for fi9 := 0; fi9 < 17; fi9++ {
                v10 := v3[fi8]
                v11 := GenS2{F0: 42, F1: v1}
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi8
        }
    }
    fmt.Println(v5, v2, v2)
    for fi12 := 0; fi12 < 5; fi12++ {
        switch "a" {
        case "\n\t":
            switch ("a" + ("\n\t" + ("foo bar" + "hello"))) {
            case "a":
                if false {
                    if ((4294967295 >= 255) || (2 >= (v4 & 2147483647))) {
                        var v13 uint16 = 773
                        v2 -= v5
                        _ = v13
                    }
                }
            default:
                switch "a" {
                case "":
                    v14 := GenS3{F0: uint8(v4)}
                    _ = v14
                }
            }
        }
        var v15 string = ("foo bar" + "hello")
        _ = fi12
        _ = v15
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 uint8 = 167
        v18 := GenS1{F0: 32767, F1: 42}
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 uint16 = 373
    fmt.Println(v5, v5, v4)
    v20 := []int{251, -9223372036854775808, -172, -527}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
