package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }

func main() {
    v1 := []uint32{4294967295, 1}
    var v2 bool = ("" == ("hello" + ""))
    var v3 string = "hello"
    v4 := []int{9223372036854775807, 1, 9223372036854775807, -9223372036854775808}
    var v5 string = v3
    var v6 int8 = 1
    var v7 int16 = 42
    for fi8 := 0; fi8 < 7; fi8++ {
        if (v7 != (v7 & -31)) {
            v3 = "\n\t"
        }
        _ = fi8
    }
    switch v3 {
    case "":
        var v9 uint64 = 18446744073709551615
        _ = v9
    case "hello":
        if (true && v2) {
            var v10 int32 = 392
            v6 = v6
            switch "\n\t" {
            case "hello":
                v7 = -71
            default:
                var v11 uint64 = 946
                _ = v11
            }
            _ = v10
        } else {
            _ = len(v1)
        }
    default:
        for fi12 := 0; fi12 < 16; fi12++ {
            v13 := GenS1{F0: 42, F1: 2}
            v2 = v2
            _ = fi12
            _ = v13
        }
    }
    v14 := GenS2{F0: (^(v6))}
    var v15 string = "a"
    v7 = 1000
    v5 = "hello"
    var v16 float64 = (1e10 - -2.25)
    switch v5 {
    case "a":
        switch v15 {
        case "foo bar":
            v3 = v15
        default:
            v17 := GenS3{F0: 0}
            _ = v17
        }
    }
    switch 211 {
    case 4294967295:
        v5 = v3
    case -1:
        if v2 {
            v7 += (v7 % (((v7 ^ 1000)) | 1))
        }
    default:
        for fi18 := 0; fi18 < 12; fi18++ {
            v19 := GenS3{F0: (fi18 << ((4294967295) & 63))}
            _ = fi18
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 float64 = (v16 * (0.0 / 1.5))
        _ = fi20
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
