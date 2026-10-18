package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 int64; F1 uint32; F2 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = 141
    var v3 float64 = float64((-2.25 - 3.14159))
    v4 := []uint32{2147483648, 2, 400, 800}
    var v5 uint16 = 2
    for fi6 := 0; fi6 < 19; fi6++ {
        fmt.Println(v5)
        _ = fi6
    }
    if true {
        for fi7 := 0; fi7 < 15; fi7++ {
            v2 -= (^(v2))
            if true {
                if ((!false) || true) {
                    v2 += 9223372036854775807
                }
                fmt.Println(v5, v2, v3)
            } else {
                v8 := GenS1{F0: -1}
                _ = v8
            }
            v9 := GenS1{F0: -9223372036854775808}
            _ = fi7
            _ = v9
        }
        for fi10 := 0; fi10 < 6; fi10++ {
            v11 := GenS2{F0: uint32(v5)}
            v4 = append(v4, 701)
            _ = fi10
            _ = v11
        }
        v12 := GenS3{F0: 2, F1: 914}
        _ = v12
    }
    v3 += 1.5
    v13 := GenS3{F0: 235, F1: 2}
    fmt.Println(v2, v2)
    fmt.Println(v2)
    var v14 int64 = -1
    var v15 int = -9223372036854775808
    switch "hello" {
    case "":
        v16 := GenS4{F0: (v2 - (v2 / ((42) | 1))), F1: 29, F2: 65535}
        _ = v16
    case "\n\t":
        for fi17 := 0; fi17 < 2; fi17++ {
            v2 = (v2 - -608)
            _ = fi17
        }
    default:
        v14 -= (v2 % ((-9223372036854775808) | 1))
    }
    v5 += v5
    var v18 int = (v15 * v15)
    v1[474] = false
    v3 -= v3
    var v19 int16 = -753
    for fi20 := 0; fi20 < 10; fi20++ {
        delete(v1, 2)
        _ = fi20
    }
    var v21 int16 = (-(v19))
    for fi22 := 0; fi22 < 4; fi22++ {
        v2 += (v2 | -9223372036854775808)
        _ = fi22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
