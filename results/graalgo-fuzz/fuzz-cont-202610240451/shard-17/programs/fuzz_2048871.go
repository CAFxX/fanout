package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint16 = 255
    v2 := make(map[uint32]int)
    var v3 float64 = 1.5
    var v4 int32 = -355
    var v5 uint8 = 24
    var v6 float32 = 0.1
    fmt.Println(v4, v6, v3)
    v6 += v6
    switch 1 {
    case 2147483647:
        v1 = v1
    default:
        v7 := []int{}
        _ = v7
    }
    fmt.Println(v6)
    var v8 int8 = 127
    if (v3 > 1.5) {
        v8 = v8
    }
    v5 += v5
    for fi9 := 0; fi9 < 16; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            v2[2147483648] = fi10
            v5 -= uint8(v1)
            fmt.Println(v8, fi9)
            _ = fi10
        }
        v6 = ((float32((0.0 / v3)) + 0.0) - float32(0.0))
        _ = fi9
    }
    var v11 uint = 116
    v4 -= 2147483647
    for fi12 := 0; fi12 < 1; fi12++ {
        v1 += v1
        if true {
            var v13 bool = (((^(v5)) >= 128) && false)
            v3 = v3
            _ = v13
        } else {
            if false {
                v14 := GenS1{F0: 4294967295}
                _ = v14
            }
            v2[2] = int(v1)
        }
        v2[587] = 9223372036854775807
        _ = fi12
    }
    v8 = (v8 * v8)
    v8 = (^(v8))
    var v15 bool = ((859 >= -9223372036854775808) || true)
    fmt.Println(v1, v4)
    if (!v15) {
        var v16 uint32 = 65535
        v17 := GenS1{F0: (v16 ^ v16)}
        _ = v16
        _ = v17
    } else {
        for fi18 := 0; fi18 < 13; fi18++ {
            v2[0] = int(v1)
            v19 := []int{-1}
            _ = fi18
            _ = v19
        }
        if ((1000 >= -1) || true) {
            for fi20 := 0; fi20 < 13; fi20++ {
                if true {
                    if (v15 || ((v11 + (v11 << (((^(v11))) & 63))) == (v11 % ((630) | 1)))) {
                        v21 := GenS1{F0: 2147483647}
                        _ = v21
                    }
                    v3 = -2.25
                    _ = len(v2)
                }
                _ = fi20
            }
        }
        v8 += (^(v8))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
