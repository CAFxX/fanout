package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int8; F1 int32; F2 uint }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 uint32; F1 uint64; F2 bool }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]int)
    v3 := make(map[string]int)
    for fi4 := 0; fi4 < 11; fi4++ {
        v5 := []int{fi4, 42, -520, 2147483647}
        switch "foo bar" {
        case "foo bar":
            var v6 int = fi4
            _ = v6
        default:
            v7 := GenS1{F0: 9223372036854775807, F1: 128}
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    if (2 < 513) {
        var v8 bool = false
        v9 := GenS1{F0: 1, F1: 127}
        _ = v8
        _ = v9
    } else {
        v10 := []int{}
        _ = v10
    }
    var v11 string = "a"
    v12 := GenS2{F0: 945, F1: 42}
    v13 := GenS3{F0: 42, F1: 1, F2: 18446744073709551615}
    v14 := GenS1{F0: 0, F1: 127}
    v15 := GenS4{F0: ((float64((1e10 * -2.25)) - float64(1e10)) >= float64(0.0)), F1: 1}
    v11 = ""
    switch 42 {
    case -530:
        if (2147483648 != 545) {
            v11 = v11
            if true {
                v11 = ((("foo bar" + "\n\t") + "") + v11)
                v11 = v11
                v16 := GenS5{F0: 552, F1: 751, F2: true}
                _ = v16
            } else {
                fmt.Println(v11)
            }
        }
    case -894:
        v17 := []int{}
        _ = v17
    default:
        v11 = v11
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
