package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 int32; F2 uint32 }
type GenS3 struct { F0 int8; F1 uint32 }
func genFn0(p0 int32) bool {
    v4 := GenS1{F0: 22}
    p0 = 2147483647
    switch -631 {
    case 346:
        p0 += p0
    }
    v5 := []int{}
    switch "a" {
    case "":
        switch "a" {
        case "":
            p0 += (p0 | p0)
        default:
            switch "\n\t" {
            case "":
                v5 = append(v5, -857)
            case "\n\t":
                p0 += (p0 >> ((2) & 63))
            case "hello":
                p0 -= p0
            }
        }
    case "hello":
        p0 = p0
    }
    _ = p0
    _ = v4
    _ = v5
    return (1 != int16(p0))
}

func main() {
    var v1 bool = (false || (-1 <= 23))
    v2 := make(map[uint32]bool)
    v3 := []int{}
    for fi6 := 0; fi6 < 14; fi6++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = len(v2)
        _ = fi6
    }
    v3 = append(v3, 802)
    var v7 uint16 = 255
    v8 := GenS1{F0: (v7 % (((^(v7))) | 1))}
    if v1 {
        fmt.Println(v1, v7)
    }
    v9 := GenS2{F0: 32767, F1: 42, F2: 301}
    switch ("\n\t" + "\n\t") {
    case "hello":
        if (1 != 9223372036854775807) {
            fmt.Println(v7, v1)
        }
    }
    v1 = true
    if v1 {
        fmt.Println(v7, v1)
        var v10 int64 = 0
        v1 = (v1 || v1)
        _ = v10
    }
    v7 -= (v7 >> ((4294967295) & 63))
    v11 := GenS3{F0: 124, F1: 2147483647}
    for fi12 := 0; fi12 < 5; fi12++ {
        switch fi12 {
        case 32767:
            var v13 uint16 = 65535
            _ = v13
        case -999:
            v1 = (2 > 65535)
        default:
            var v14 uint32 = 2147483648
            _ = v14
        }
        for fi15 := 0; fi15 < 10; fi15++ {
            v7 = v7
            _ = fi15
        }
        _ = fi12
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
