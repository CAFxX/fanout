package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int32; F1 int64; F2 int16 }

func main() {
    var v1 uint64 = 2
    v2 := make(map[uint32]int)
    var v3 int32 = -1
    var v4 int64 = 9223372036854775807
    if (194 != 42) {
        for fi5 := 0; fi5 < 12; fi5++ {
            for fi6 := 0; fi6 < 6; fi6++ {
                v7 := GenS1{F0: 0, F1: (v3 * v3)}
                switch ("hello" + "\n\t") {
                case "hello":
                    v8 := []int{fi6, 2147483647, fi6, fi6}
                    _ = v8
                }
                v2[65535] = 0
                _ = fi6
                _ = v7
            }
            v4 -= -260
            v1 -= 280
            _ = fi5
        }
    }
    v3 = (v3 + 42)
    v9 := []int{2, -22, -133, 449}
    _ = len(v2)
    if false {
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 int64 = (v4 << ((587) & 63))
            v3 += (v3 << ((0) & 63))
            _ = fi10
            _ = v11
        }
    } else {
        for fi12 := 0; fi12 < 12; fi12++ {
            v4 = v4
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 float64 = 1e10
                v15 := GenS2{F0: (v3 >> ((255) & 63)), F1: v4, F2: 42}
                var v16 uint = 135
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            _ = fi12
        }
    }
    switch "a" {
    case "\n\t":
        v17 := GenS2{F0: (v3 << ((727) & 63)), F1: 0, F2: 32767}
        _ = v17
    case "foo bar":
        var v18 uint16 = 65535
        _ = v18
    case "":
        v19 := v2[474]
        _ = v19
    }
    _ = len(v2)
    if false {
        var v20 int16 = 1
        var v21 bool = false
        var v22 int64 = v4
        _ = v20
        _ = v21
        _ = v22
    } else {
        fmt.Println(v1, v3)
        switch 0 {
        case -1:
            v23 := GenS1{F0: 255, F1: (v3 ^ 2147483647)}
            _ = v23
        case 42:
            v3 = v3
        case 0:
            v9 = append(v9, 32767)
        default:
            for fi24 := 0; fi24 < 7; fi24++ {
                switch fi24 {
                case 35:
                    fmt.Println(fi24, fi24)
                case 9223372036854775807:
                    var v25 int = 127
                    _ = v25
                default:
                    var v26 bool = false
                    _ = v26
                }
                _ = fi24
            }
        }
    }
    var v27 uint32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v9
}
