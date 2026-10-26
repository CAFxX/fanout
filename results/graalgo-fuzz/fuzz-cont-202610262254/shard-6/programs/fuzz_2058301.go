package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 int = -2147483648
    var v2 float32 = 0.1
    var v3 float64 = (3.14159 / (0.0 + -2.25))
    var v4 float64 = (3.14159 - v3)
    fmt.Println(v3)
    var v5 uint = 176
    v6 := GenS1{F0: 65535, F1: 544, F2: 1}
    v4 += v4
    v5 = (v5 | v5)
    var v7 int8 = 0
    for fi8 := 0; fi8 < 10; fi8++ {
        if false {
            v9 := []int{(fi8 << ((4294967295) & 63))}
            _ = v9
        } else {
            switch (fi8 / (((v1 + (fi8 ^ (fi8 & v1)))) | 1)) {
            case 128:
                if true {
                    if (v1 <= (-(v1))) {
                        switch -657 {
                        case 9223372036854775807:
                            v10 := []int{fi8, fi8}
                            _ = v10
                        case -381:
                            var v11 uint8 = 42
                            _ = v11
                        case 42:
                            fmt.Println(fi8, v3, v3)
                        default:
                            switch "a" {
                            case "hello":
                                v1 = (v1 >> ((2) & 63))
                            case "foo bar":
                                v12 := []int{}
                                _ = v12
                            default:
                                fmt.Println(v1, v1, fi8)
                            }
                        }
                    }
                }
            case 217:
                v13 := []int{(v1 / (((v1 * fi8)) | 1)), 127}
                _ = v13
            }
            fmt.Println(v1)
        }
        v14 := GenS1{F0: 2, F1: 981, F2: 255}
        var v15 uint8 = 0
        _ = fi8
        _ = v14
        _ = v15
    }
    v7 -= 1
    for fi16 := 0; fi16 < 9; fi16++ {
        v17 := GenS2{F0: 42}
        var v18 int = (fi16 ^ 2147483647)
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: 255, F1: 692, F2: 42}
    v20 := GenS1{F0: 255, F1: 376, F2: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
