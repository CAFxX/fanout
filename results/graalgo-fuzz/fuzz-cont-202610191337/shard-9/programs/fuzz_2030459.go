package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int8 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS3 struct { F0 int; F1 int32; F2 int32 }

func main() {
    var v1 uint = 0
    v2 := make(map[uint32]bool)
    var v3 uint16 = 2
    var v4 uint32 = 629
    fmt.Println(v1, v1, v4)
    for fi5 := 0; fi5 < 18; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            var v7 uint = 193
            fmt.Println(v1, v7, fi6)
            v8 := GenS1{F0: (v3 % ((v3) | 1)), F1: 0, F2: 127}
            switch "foo bar" {
            case "hello":
                switch "hello" {
                case "hello":
                    var v9 int32 = -407
                    _ = v9
                case "a":
                    switch "hello" {
                    case "hello":
                        var v10 string = ("a" + "a")
                        _ = v10
                    case "a":
                        v11 := GenS2{F0: (v3 ^ 255), F1: int8(v7), F2: 0}
                        _ = v11
                    case "\n\t":
                        if false {
                            var v12 uint8 = 42
                            v12 += (v12 & v12)
                            _ = v12
                        }
                    }
                case "":
                    v13 := GenS1{F0: v3, F1: 42, F2: 1}
                    _ = v13
                default:
                    v14 := GenS2{F0: (v3 ^ (^(v3))), F1: -41, F2: 127}
                    _ = v14
                }
            case "a":
                var v15 uint32 = v4
                _ = v15
            case "":
                v4 -= v4
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        v16 := GenS1{F0: 1, F1: -2147483648, F2: -128}
        _ = fi5
        _ = v16
    }
    var v17 float64 = (float64(3.14159) - 0.0)
    var v18 float32 = 1e10
    var v19 uint32 = (v4 - (v4 % (((v4 >> ((v1) & 63))) | 1)))
    var v20 bool = true
    v17 += (0.1 - -2.25)
    fmt.Println(v4, v4)
    fmt.Println(v4, v3)
    var v21 uint64 = 4294967295
    v22 := GenS3{F0: 32768, F1: 77, F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
