package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint16 }

func main() {
    var v1 int = 1
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    var v4 uint8 = 128
    var v5 int = v1
    switch ("a" + "\n\t") {
    case "foo bar":
        for fi6 := 0; fi6 < 8; fi6++ {
            fmt.Println(fi6)
            _ = fi6
        }
    }
    var v7 uint16 = 473
    v5 -= 32767
    var v8 uint = 0
    switch (v5 * (v1 - (v1 + v5))) {
    case 42:
        switch (v5 - v5) {
        case 127:
            v9 := GenS1{F0: -1, F1: (v8 * (v8 - v8))}
            _ = v9
        case 807:
            switch (v1 + v1) {
            case -1:
                switch "foo bar" {
                case "\n\t":
                    v10 := GenS1{F0: -1, F1: v8}
                    _ = v10
                case "":
                    v11 := GenS1{F0: -1, F1: (v8 - v8)}
                    _ = v11
                }
            default:
                fmt.Println(v8, v8)
            }
        }
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        for fi13 := 0; fi13 < 2; fi13++ {
            v4 -= v4
            _ = fi13
        }
        fmt.Println(fi12)
        _ = fi12
    }
    var v14 bool = true
    if ((0.1 / -2.25) > (0.1 - float64(1.5))) {
        fmt.Println(v14)
        for fi15 := 0; fi15 < 1; fi15++ {
            v16 := []int{fi15, -28, (fi15 * (v5 ^ fi15)), v1}
            v17 := GenS1{F0: -9223372036854775808, F1: 2}
            switch (^(v1)) {
            case 830:
                var v18 int16 = 32767
                _ = v18
            default:
                v19 := GenS1{F0: 0, F1: 42}
                _ = v19
            }
            _ = fi15
            _ = v16
            _ = v17
        }
        v5 -= -9223372036854775808
    }
    fmt.Println(v8, v5)
    var v20 uint64 = 18446744073709551615
    var v21 int = (v5 << ((v8) & 63))
    v22 := GenS2{F0: v4, F1: 435, F2: 2}
    v23 := []int{}
    v1 -= v5
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
