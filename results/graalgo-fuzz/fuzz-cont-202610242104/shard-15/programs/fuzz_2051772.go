package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }

func main() {
    var v1 uint32 = 65535
    var v2 uint16 = 487
    var v3 uint64 = 4294967295
    v4 := make(map[uint32]int)
    var v5 int8 = 42
    v5 = v5
    v2 = 42
    var v6 uint64 = v3
    switch 0 {
    case 2147483647:
        if (-128 == -116) {
            v6 += v3
            v2 = 65535
        } else {
            for fi7 := 0; fi7 < 17; fi7++ {
                v8 := GenS1{F0: 128, F1: 65535}
                _ = fi7
                _ = v8
            }
        }
    case -290:
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 float32 = 1.5
            if (!(65535 != 1)) {
                for fi11 := 0; fi11 < 12; fi11++ {
                    v12 := GenS1{F0: 102, F1: 65535}
                    switch "a" {
                    case "a":
                        v4[uint32(v2)] = (fi9 >> ((522) & 63))
                    case "\n\t":
                        v13 := []int{(fi11 | fi11), (fi11 % ((-174) | 1)), (-(fi11)), 127}
                        _ = v13
                    default:
                        fmt.Println(v6, v6)
                    }
                    _ = fi11
                    _ = v12
                }
            } else {
                v10 = v10
                v2 += (v2 - uint16(v1))
            }
            _ = fi9
            _ = v10
        }
    case 674:
        for fi14 := 0; fi14 < 9; fi14++ {
            for fi15 := 0; fi15 < 4; fi15++ {
                var v16 int64 = -9223372036854775808
                fmt.Println(v5)
                _ = fi15
                _ = v16
            }
            v17 := []int{}
            _ = fi14
            _ = v17
        }
    default:
        v18 := v4[(^(v1))]
        _ = v18
    }
    var v19 uint = 18446744073709551615
    v20 := GenS1{F0: 255, F1: uint(v1)}
    fmt.Println(v19, v3)
    v21 := []int{1, -1}
    v6 = 42
    var v22 int = -9223372036854775808
    v6 -= uint64(v1)
    for fi23 := 0; fi23 < 19; fi23++ {
        if len(v21) >= 2 {
            v21 = v21[1:len(v21)]
        }
        fmt.Println(v22, v19, v1)
        _ = fi23
    }
    var v24 int16 = int16(v5)
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
