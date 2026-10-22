package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 string = "foo bar"
    var v2 uint32 = 2
    v3 := []int{9223372036854775807, 127, 128}
    fmt.Println(v2, v2)
    for fi4 := 0; fi4 < 8; fi4++ {
        v1 = "hello"
        v1 = (v1 + ("" + v1))
        for fi5 := 0; fi5 < 11; fi5++ {
            switch v1 {
            case "foo bar":
                _ = len(v3)
            case "":
                fmt.Println(fi4, v2)
            case "\n\t":
                v2 = v2
            default:
                v3 = append(v3, (fi4 | 880))
            }
            _ = fi5
        }
        _ = fi4
    }
    v2 = v2
    switch -1 {
    case -2147483648:
        v1 = "hello"
    case 32767:
        v6 := GenS1{F0: 1, F1: 2, F2: -1}
        _ = v6
    default:
        v7 := GenS2{F0: (v2 & (v2 + v2))}
        _ = v7
    }
    var v8 int8 = -128
    for fi9 := 0; fi9 < 9; fi9++ {
        v10 := GenS2{F0: (v2 / ((511) | 1))}
        _ = len(v3)
        _ = fi9
        _ = v10
    }
    _ = len(v3)
    if (88 == 0) {
        var v11 int64 = 868
        var v12 uint64 = 0
        _ = v11
        _ = v12
    } else {
        v2 = 67
    }
    fmt.Println(v1, v8)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
}
