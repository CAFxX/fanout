package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint32; F1 uint; F2 int64 }

func main() {
    var v1 uint = 255
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    v1 = v1
    for fi4 := 0; fi4 < 10; fi4++ {
        v1 = v1
        _ = fi4
    }
    v5 := GenS1{F0: (v1 & v1), F1: 4294967295}
    switch ("hello" + "\n\t") {
    case "hello":
        switch ("" + "hello") {
        case "\n\t":
            var v6 int8 = -55
            _ = v6
        case "a":
            switch "foo bar" {
            case "hello":
                v1 += 42
            default:
                v7 := GenS1{F0: v1, F1: 1}
                _ = v7
            }
        default:
            v1 = v1
        }
    case "\n\t":
        v1 = 4294967295
    case "a":
        v1 = v1
    }
    var v8 uint8 = 2
    switch "a" {
    case "":
        v8 = (v8 >> (((v1 % ((385) | 1))) & 63))
    case "a":
        for fi9 := 0; fi9 < 1; fi9++ {
            delete(v2, fi9)
            _ = fi9
        }
    case "hello":
        fmt.Println(v1, v1)
    default:
        var v10 bool = true
        _ = v10
    }
    if (false || (!false)) {
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 uint16 = 2
            var v13 int32 = 2147483647
            _ = fi11
            _ = v12
            _ = v13
        }
        if (!true) {
            var v14 int = 687
            v8 += 1
            _ = v14
        }
        v1 += (v1 << (((^(v1))) & 63))
    }
    v8 += v8
    for fi15 := 0; fi15 < 16; fi15++ {
        v16 := []int{fi15, (fi15 >> (((^(v1))) & 63)), (^(fi15)), (-(fi15))}
        switch fi15 {
        case -2147483648:
            v17 := GenS1{F0: 65535, F1: fi15}
            _ = v17
        case -831:
            v18 := GenS2{F0: 42, F1: (v1 + v1), F2: 9223372036854775807}
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    fmt.Println(v8)
    v8 = 42
    var v19 uint = 1
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v5
    _ = v8
}
