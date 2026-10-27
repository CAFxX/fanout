package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int32 }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[string]string)
    v2 := []int{42, 9223372036854775807, 9223372036854775807, -906}
    var v3 string = "a"
    var v4 float64 = 1.5
    v5 := make(map[int]int)
    v6 := make(map[string]uint32)
    var v7 uint64 = 42
    switch 32767 {
    case -444:
        switch "a" {
        case "":
            _ = len(v1)
        case "foo bar":
            for fi8 := 0; fi8 < 2; fi8++ {
                v3 = v3
                v7 = (v7 | v7)
                switch v3 {
                case "hello":
                    v2 = append(v2, fi8)
                case "foo bar":
                    for fi9 := 0; fi9 < 11; fi9++ {
                        fmt.Println(v7, v4, fi9)
                        fmt.Println(v7, v3)
                        _ = fi9
                    }
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = fi8
            }
        }
    case 32767:
        v3 = v3
    case -1:
        _ = len(v1)
    default:
        fmt.Println(v4, v3, v7)
    }
    if len(v2) > 0 {
        v2[0] = 127
    }
    v7 = (v7 % ((89) | 1))
    var v10 float64 = v4
    switch 32768 {
    case -9223372036854775808:
        v11 := GenS1{F0: 0, F1: v7, F2: 42}
        _ = v11
    case -189:
        v7 -= 0
    }
    var v12 uint64 = v7
    if ((978 == 61) || false) {
        v13 := GenS1{F0: 803, F1: v12, F2: -1}
        fmt.Println(v10, v7)
        _ = v13
    } else {
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 int = -962
            v10 = -2.25
            _ = fi14
            _ = v15
        }
        fmt.Println(v12, v10)
    }
    v16 := GenS2{F0: -925}
    v17 := GenS2{F0: 9223372036854775807}
    var v18 string = "hello"
    for fi19 := 0; fi19 < 19; fi19++ {
        v7 += v12
        _ = fi19
    }
    var v20 uint32 = 0
    var v21 string = "hello"
    v22 := GenS2{F0: -9223372036854775808}
    v10 = (v4 * v10)
    v4 = v10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
