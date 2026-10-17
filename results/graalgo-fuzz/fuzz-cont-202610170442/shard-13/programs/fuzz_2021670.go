package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 string; F1 bool; F2 int32 }
type GenS4 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS5 struct { F0 int32; F1 uint16 }

func main() {
    var v1 string = (("a" + "hello") + "hello")
    var v2 float64 = (0.1 / (0.1 + 1e10))
    var v3 int64 = 831
    if (false && (0 < 1)) {
        v4 := GenS1{F0: 42, F1: 1, F2: -832}
        v5 := GenS2{F0: 0, F1: 65535}
        _ = v4
        _ = v5
    } else {
        v2 = v2
        v6 := []int{32768, 355, 0}
        _ = v6
    }
    if true {
        var v7 int8 = 0
        v8 := []int{0, 128}
        var v9 int8 = 100
        _ = v7
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := GenS3{F0: v1, F1: (!false), F2: -1}
            v3 -= (v3 >> ((42) & 63))
            v12 := GenS4{F0: 42, F1: 0, F2: 488}
            _ = fi10
            _ = v11
            _ = v12
        }
        v2 = v2
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        fmt.Println(v2, v1, v2)
        var v14 int8 = 1
        switch "foo bar" {
        case "":
            switch fi13 {
            case 42:
                v15 := []int{(fi13 - fi13), fi13, -2147483648, -354}
                _ = v15
            case -1:
                var v16 uint64 = 363
                _ = v16
            case 2147483647:
                v17 := []int{2147483647, fi13, (fi13 * fi13)}
                _ = v17
            }
        case "hello":
            v2 = 3.14159
        default:
            for fi18 := 0; fi18 < 1; fi18++ {
                var v19 int64 = (v3 + (v3 - 9223372036854775807))
                v2 += 0.1
                _ = fi18
                _ = v19
            }
        }
        _ = fi13
        _ = v14
    }
    fmt.Println(v2)
    v20 := GenS5{F0: 2147483647, F1: 32767}
    fmt.Println(v1, v1)
    v3 -= v3
    fmt.Println(v1)
    v21 := []int{-230, 1, 32768, 32767}
    v1 = "foo bar"
    v22 := GenS5{F0: 42, F1: 177}
    v23 := GenS3{F0: "a", F1: false, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
