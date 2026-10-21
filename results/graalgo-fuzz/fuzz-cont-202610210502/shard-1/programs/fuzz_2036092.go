package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 string }
type GenS2 struct { F0 int8; F1 int16; F2 uint8 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = 0.1
    var v3 uint16 = 255
    v4 := make(map[int]int)
    var v5 string = "\n\t"
    var v6 float32 = 1e10
    var v7 int8 = -89
    v8 := []int{127}
    v6 -= v2
    var v9 uint16 = 0
    v1[9223372036854775807] = v5
    if true {
        switch 688 {
        case 4294967295:
            v7 = v7
        default:
            v10 := GenS1{F0: 808, F1: 0, F2: (v5 + ((v5 + v5) + (v5 + (((v5 + "\n\t") + v5) + v5))))}
            _ = v10
        }
        if ((v5 + "\n\t") == (((((v5 + ((("a" + v5) + (v5 + "foo bar")) + "foo bar")) + v5) + "hello") + v5) + ("\n\t" + v5))) {
            v11 := GenS2{F0: v7, F1: 1, F2: 0}
            v12 := GenS3{F0: 9223372036854775807}
            _ = v11
            _ = v12
        }
        var v13 bool = true
        _ = v13
    }
    v9 += (v3 & 462)
    if ((^(v9)) > (v9 % ((255) | 1))) {
        v3 = v9
    }
    var v14 int = 42
    for fi15 := 0; fi15 < 1; fi15++ {
        fmt.Println(v9)
        _ = fi15
    }
    var v16 string = (v5 + ("\n\t" + v5))
    if true {
        v17 := GenS1{F0: 473, F1: 596, F2: v16}
        _ = v17
    }
    v9 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
