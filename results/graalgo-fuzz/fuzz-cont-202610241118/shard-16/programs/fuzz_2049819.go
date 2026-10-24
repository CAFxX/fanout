package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int; F1 uint; F2 uint64 }

func main() {
    var v1 string = ""
    var v2 int64 = -1
    var v3 int8 = 4
    var v4 uint32 = 0
    v1 = v1
    v5 := GenS1{F0: 0, F1: v2}
    v6 := []int{}
    var v7 uint16 = 0
    var v8 int = 210
    v1 = ((v1 + v1) + "foo bar")
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v3 += v3
    switch v1 {
    case "\n\t":
        v9 := GenS2{F0: 0, F1: 442, F2: 977}
        _ = v9
    default:
        if (!true) {
            var v10 uint64 = 42
            for fi11 := 0; fi11 < 6; fi11++ {
                if false {
                    v12 := GenS2{F0: 2147483647, F1: 1, F2: (v10 << ((2) & 63))}
                    _ = v12
                }
                if (78 > -128) {
                    var v13 uint8 = 255
                    _ = v13
                }
                switch (fi11 >> ((255) & 63)) {
                case -1:
                    var v14 uint8 = 128
                    _ = v14
                }
                _ = fi11
            }
            _ = v10
        }
    }
    v15 := GenS1{F0: 0, F1: int64(v3)}
    var v16 float64 = float64(1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
