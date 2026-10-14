package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 int32; F1 uint8 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]bool)
    var v4 int = 128
    if (!(22 > -128)) {
        v5 := []int{}
        _ = v5
    }
    var v6 uint8 = 2
    switch (v4 << ((2) & 63)) {
    case -200:
        v7 := GenS1{F0: 467, F1: 825, F2: 1}
        _ = v7
    case -187:
        if (true && false) {
            v4 = (v4 | v4)
            switch "\n\t" {
            case "foo bar":
                var v8 bool = false
                _ = v8
            case "hello":
                switch (v4 * v4) {
                case 1:
                    var v9 uint32 = 4294967295
                    _ = v9
                case 42:
                    if (((3.14159 - -2.25) == 1.5) && ((39 <= uint(v6)) && (!false))) {
                        var v10 uint8 = 84
                        _ = v10
                    }
                default:
                    v11 := GenS2{F0: 65535, F1: -32768}
                    _ = v11
                }
            case "\n\t":
                v12 := GenS3{F0: 973, F1: (v6 + (v6 / ((v6) | 1)))}
                _ = v12
            }
            var v13 int = v4
            _ = v13
        } else {
            v14 := GenS1{F0: 990, F1: 42, F2: int64(v6)}
            _ = v14
        }
    case 32768:
        v6 += v6
    }
    var v15 int16 = 572
    v16 := []int{(v4 | (v4 | v4)), v4, 49, 1}
    v4 -= (v4 ^ v4)
    v6 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
