package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint8; F1 bool; F2 int8 }
type GenS4 struct { F0 int64; F1 bool; F2 int64 }

func main() {
    v1 := []int{36, 9223372036854775807, 4294967295, 42}
    v2 := []uint32{4294967295, 1, 4294967295, 65535}
    var v3 uint64 = 0
    v4 := []int{0, 325}
    var v5 uint32 = 2147483647
    fmt.Println(v5, v3)
    switch "\n\t" {
    case "hello":
        v6 := GenS1{F0: (v5 / ((2147483648) | 1))}
        _ = v6
    case "":
        v7 := GenS2{F0: 255, F1: ("a" + "\n\t")}
        _ = v7
    default:
        switch "hello" {
        case "":
            var v8 uint16 = 255
            _ = v8
        case "hello":
            v3 -= 0
        case "foo bar":
            for fi9 := 0; fi9 < 16; fi9++ {
                if len(v4) > 0 {
                    v4[0] = 0
                }
                if (v5 < (v5 | 2147483648)) {
                    v10 := GenS1{F0: 42}
                    var v11 uint16 = 2
                    _ = v10
                    _ = v11
                }
                _ = fi9
            }
        }
    }
    v5 += v5
    if true {
        var v12 float32 = 3.14159
        _ = v12
    } else {
        for fi13 := 0; fi13 < 6; fi13++ {
            var v14 uint16 = 301
            var v15 float64 = (-2.25 * ((1.5 - 0.0) / (0.0 - 0.1)))
            switch "hello" {
            case "a":
                var v16 uint16 = 42
                _ = v16
            case "hello":
                v17 := GenS3{F0: 225, F1: false, F2: -128}
                _ = v17
            case "":
                v3 += uint64(v5)
            }
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    switch "foo bar" {
    case "":
        var v18 uint16 = 255
        _ = v18
    case "hello":
        for fi19 := 0; fi19 < 10; fi19++ {
            v5 = (v5 - (v5 * (^(v5))))
            v5 = 65535
            var v20 uint16 = 32767
            _ = fi19
            _ = v20
        }
    case "foo bar":
        var v21 uint = 65535
        _ = v21
    }
    v5 = 65535
    v22 := GenS4{F0: -1, F1: ("hello" == "\n\t"), F2: 0}
    v23 := GenS4{F0: 1, F1: false, F2: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
