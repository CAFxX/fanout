package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 255
    v3 := make(map[uint32]string)
    var v4 int64 = 9223372036854775807
    var v5 float64 = (0.1 * 0.1)
    if true {
        fmt.Println(v5)
        v4 = (^(v4))
        v5 += (-2.25 - -2.25)
    }
    var v6 uint64 = 18446744073709551615
    var v7 int8 = -122
    _ = len(v3)
    var v8 uint16 = 541
    v4 = v4
    v5 -= 0.1
    var v9 uint8 = 128
    v10 := v3[148]
    v11 := []int{4294967295, 42}
    switch 328 {
    case 128:
        var v12 int = 2147483647
        _ = v12
    case -653:
        v13 := GenS1{F0: uint32(v2), F1: (!true), F2: int64(v2)}
        _ = v13
    case -1:
        v6 += (v6 ^ v6)
    }
    var v14 int16 = 42
    if ((((v4 ^ 9223372036854775807) >= (^(v4))) && (true || false)) || ((v6 * (v6 % (((^(v6))) | 1))) < v6)) {
        v14 += -1
        var v15 int = 32768
        var v16 bool = false
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 20; fi17++ {
            v6 += (v6 + v6)
            _ = fi17
        }
    }
    v11 = append(v11, -343)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
