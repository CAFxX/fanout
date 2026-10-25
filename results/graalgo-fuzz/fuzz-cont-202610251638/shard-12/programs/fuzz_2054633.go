package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64 }

func main() {
    var v1 uint8 = 128
    var v2 uint64 = 940
    var v3 uint16 = 32767
    v4 := make(map[int]string)
    v5 := make(map[uint32]bool)
    var v6 uint64 = (^(v2))
    var v7 int16 = 0
    if true {
        fmt.Println(v7, v6)
        v7 = (v7 + (v7 | 1000))
        v8 := GenS1{F0: v1, F1: 1, F2: true}
        _ = v8
    } else {
        v2 = (v2 ^ v2)
        v6 -= 18446744073709551615
    }
    if ((^(v6)) <= (v2 << ((18446744073709551615) & 63))) {
        v2 -= 18446744073709551615
        switch "foo bar" {
        case "":
            for fi9 := 0; fi9 < 8; fi9++ {
                v2 -= v6
                v10 := []int{(^(fi9)), fi9}
                _ = fi9
                _ = v10
            }
        default:
            for fi11 := 0; fi11 < 20; fi11++ {
                v6 = v6
                for fi12 := 0; fi12 < 20; fi12++ {
                    v13 := GenS2{F0: 128}
                    _ = fi12
                    _ = v13
                }
                v2 = (^(v6))
                _ = fi11
            }
        }
    } else {
        v14 := GenS1{F0: v1, F1: v7, F2: false}
        var v15 float64 = 1.5
        _ = v14
        _ = v15
    }
    v6 = 18446744073709551615
    v16 := GenS1{F0: v1, F1: 1, F2: ((v1 - (v1 >> ((334) & 63))) > v1)}
    for fi17 := 0; fi17 < 1; fi17++ {
        if (0.1 < 1e10) {
            _ = len(v5)
            fmt.Println(v2, v6)
        }
        v6 -= v2
        v18 := GenS3{F0: -1}
        _ = fi17
        _ = v18
    }
    fmt.Println(v6, v1, v6)
    var v19 string = ("hello" + ("a" + ""))
    for fi20 := 0; fi20 < 15; fi20++ {
        _ = len(v5)
        _ = fi20
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
