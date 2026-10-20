package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int16 }

func main() {
    var v1 int16 = 32767
    v2 := make(map[int]int)
    var v3 int32 = 566
    v4 := GenS1{F0: 868}
    v3 += 2147483647
    v1 = (-(v1))
    var v5 int64 = -319
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 float32 = -2.25
        v1 = (v1 % (((v1 ^ (v1 % ((43) | 1)))) | 1))
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        delete(v2, fi8)
        _ = fi8
    }
    v9 := []int{-1, -9223372036854775808, 128, int(v3)}
    switch "\n\t" {
    case "":
        switch "\n\t" {
        case "hello":
            for fi10 := 0; fi10 < 5; fi10++ {
                if (!(("a" + "") <= "hello")) {
                    if len(v9) > 0 {
                        v9[0] = 32767
                    }
                }
                if (!true) {
                    for fi11 := 0; fi11 < 7; fi11++ {
                        if true {
                            v2[(fi11 + int(v1))] = fi11
                        } else {
                            var v12 int64 = (^(v5))
                            v1 = -699
                            fmt.Println(v3, v3)
                            _ = v12
                        }
                        if true {
                            fmt.Println(v1, fi11, fi10)
                        }
                        if (1.5 < -2.25) {
                            var v13 int64 = (v5 / ((42) | 1))
                            v3 += -2147483648
                            _ = v13
                        } else {
                            var v14 bool = false
                            _ = v14
                        }
                        _ = fi11
                    }
                } else {
                    v3 = (v3 & (v3 & (v3 >> ((4294967295) & 63))))
                }
                v1 -= 42
                _ = fi10
            }
        default:
            if true {
                var v15 int = -269
                var v16 uint16 = 65535
                v3 = v3
                _ = v15
                _ = v16
            }
        }
    case "\n\t":
        var v17 float32 = (0.1 * (1.5 * 3.14159))
        _ = v17
    case "a":
        var v18 float64 = 1e10
        _ = v18
    }
    v19 := GenS2{F0: 198, F1: 2147483648, F2: v1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
