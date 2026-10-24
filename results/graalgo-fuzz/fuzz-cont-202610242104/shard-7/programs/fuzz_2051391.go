package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 bool; F1 int64 }
func genFn0(p0 uint) int8 {
    v4 := GenS1{F0: "hello"}
    v5 := GenS1{F0: "\n\t"}
    p0 = p0
    if true {
        p0 = p0
    }
    _ = p0
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 uint32 = 923
    v2 := make(map[uint32]bool)
    v3 := []int{-2147483648}
    v6 := GenS2{F0: 933, F1: 128}
    switch "hello" {
    case "a":
        v7 := GenS2{F0: 195, F1: 128}
        _ = v7
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 int16 = 743
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        if false {
            switch "" {
            case "foo bar":
                var v11 uint64 = 0
                _ = v11
            case "":
                v12 := GenS1{F0: "foo bar"}
                _ = v12
            }
            fmt.Println(v1, fi10)
        }
        delete(v2, (v1 | (v1 & 49)))
        _ = fi10
    }
    switch 0 {
    case -191:
        v13 := GenS2{F0: 194, F1: 127}
        _ = v13
    case -2147483648:
        fmt.Println(v1)
    case 279:
        switch "" {
        case "a":
            var v14 float32 = -2.25
            _ = v14
        case "foo bar":
            v1 -= v1
        case "\n\t":
            fmt.Println(v1)
        default:
            v1 += 4294967295
        }
    }
    v15 := v2[v1]
    fmt.Println(v1, v1)
    for fi16 := 0; fi16 < 2; fi16++ {
        v15 = v15
        _ = fi16
    }
    v17 := GenS2{F0: 2, F1: 214}
    _ = len(v2)
    v1 = (v1 / (((v1 - (v1 & (^(v1))))) | 1))
    var v18 int16 = 1
    delete(v2, 2147483648)
    v19 := GenS3{F0: true, F1: 403}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v6
}
