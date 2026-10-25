package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 uint; F2 uint8 }

func main() {
    var v1 uint64 = 4294967295
    var v2 uint64 = v1
    v3 := []int{1, -9223372036854775808, 4294967295, -1}
    fmt.Println(v1, v2)
    v3 = append(v3, 0)
    v3 = append(v3, 32767)
    v2 = 466
    v1 += v2
    v1 += (^(v1))
    fmt.Println(v1, v1)
    v1 = v1
    switch 32767 {
    case 4294967295:
        for fi4 := 0; fi4 < 14; fi4++ {
            v1 = (v2 >> ((223) & 63))
            for fi5 := 0; fi5 < 2; fi5++ {
                switch "foo bar" {
                case "\n\t":
                    var v6 float64 = 1.5
                    _ = v6
                case "":
                    var v7 uint32 = 504
                    _ = v7
                case "foo bar":
                    v8 := GenS1{F0: 42}
                    _ = v8
                }
                if (32767 < uint16(v2)) {
                    v3 = append(v3, fi5)
                    var v9 uint8 = 82
                    var v10 float32 = 0.0
                    _ = v10
                    _ = v9
                } else {
                    v3 = append(v3, 42)
                    var v11 float32 = (0.1 / 3.14159)
                    _ = v11
                }
                _ = fi5
            }
            v12 := GenS2{F0: 961}
            _ = fi4
            _ = v12
        }
    case 0:
        fmt.Println(v1)
    }
    v13 := GenS3{F0: 32767}
    switch "hello" {
    case "a":
        for fi14 := 0; fi14 < 8; fi14++ {
            v3 = append(v3, -1)
            _ = fi14
        }
    case "\n\t":
        v2 -= v1
    default:
        v15 := GenS4{F0: -1, F1: 10, F2: 0}
        _ = v15
    }
    v2 = (v1 % (((^(v2))) | 1))
    v1 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
}
