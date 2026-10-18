package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint64; F1 bool; F2 int64 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 bool = false
    v2 := []int{-1, 9223372036854775807, 42, 0}
    var v3 uint16 = 32767
    v3 = v3
    fmt.Println(v3, v3)
    var v4 int16 = -32768
    switch 32768 {
    case 1:
        v5 := GenS1{F0: 942, F1: (-(v4))}
        _ = v5
    case -2147483648:
        var v6 int8 = -1
        _ = v6
    case 299:
        for fi7 := 0; fi7 < 16; fi7++ {
            v2 = append(v2, 1)
            v2 = append(v2, 360)
            v3 = (v3 ^ 255)
            _ = fi7
        }
    default:
        for fi8 := 0; fi8 < 17; fi8++ {
            if len(v2) > 0 {
                v2[0] = (fi8 & 1)
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v9 := GenS2{F0: 946, F1: (false || (v1 || v1)), F2: 9223372036854775807}
            _ = fi8
            _ = v9
        }
    }
    var v10 int8 = 1
    v11 := GenS3{F0: 127}
    var v12 int16 = (v4 << ((549) & 63))
    v1 = v1
    if (!("" > "foo bar")) {
        v10 = (v10 % ((-106) | 1))
    } else {
        v4 = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
