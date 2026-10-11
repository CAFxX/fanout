package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 uint; F2 uint8 }
func genFn0(p0 uint8, p1 uint32) int16 {
    v7 := GenS1{F0: 119}
    if (-2.25 == 1e10) {
        switch "\n\t" {
        case "\n\t":
            switch 127 {
            case 127:
                v8 := []int{1, -9223372036854775808}
                _ = v8
            case -878:
                switch ("foo bar" + "foo bar") {
                case "":
                    v9 := GenS2{F0: 357}
                    _ = v9
                case "foo bar":
                    v10 := GenS2{F0: -363}
                    _ = v10
                case "\n\t":
                    var v11 float32 = 0.1
                    _ = v11
                }
            case 32767:
                var v12 float32 = float32(((1.5 * 1e10) - (3.14159 * -2.25)))
                _ = v12
            }
        case "hello":
            fmt.Println(p0, p1)
        }
    }
    v13 := []int{361, -912}
    _ = p0
    _ = p1
    _ = v13
    _ = v7
    return -32768
}
func genFn1(p0 int, p1 uint, p2 uint64) uint64 {
    p1 = p1
    var v14 int8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return (p2 - (^(p2)))
}

func main() {
    v1 := []int{253, 1, 839, 108}
    var v2 uint8 = 127
    var v3 uint16 = 1
    var v4 uint32 = 669
    v5 := make(map[string]bool)
    var v6 int64 = -721
    v3 -= v3
    var v15 uint = 843
    v16 := GenS1{F0: (v15 >> (((v15 - v15)) & 63))}
    _ = len(v1)
    switch "" {
    case "hello":
        var v17 uint8 = v2
        _ = v17
    }
    v18 := GenS3{F0: 817, F1: v15, F2: v2}
    for fi19 := 0; fi19 < 7; fi19++ {
        v6 = (v6 >> (((v15 >> (((^(v15))) & 63))) & 63))
        _ = fi19
    }
    v1 = append(v1, 1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
