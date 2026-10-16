package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint64; F2 string }
func genFn0(p0 int16, p1 int, p2 uint16) uint32 {
    v6 := []int{}
    if true {
        for fi7 := 0; fi7 < 20; fi7++ {
            p2 = (p2 << ((1) & 63))
            for fi8 := 0; fi8 < 11; fi8++ {
                switch int(p0) {
                case 9223372036854775807:
                    if (true || (1 >= 231)) {
                        _ = len(v6)
                        v9 := GenS1{F0: -9223372036854775808, F1: 922}
                        fmt.Println(fi7)
                        _ = v9
                    } else {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                case 2147483647:
                    var v10 string = "hello"
                    _ = v10
                case -9223372036854775808:
                    v11 := GenS2{F0: 1, F1: 4294967295, F2: "\n\t"}
                    _ = v11
                default:
                    p0 += 1
                }
                _ = fi8
            }
            if ((0.1 != float32(((0.0 / -2.25) - (1.5 + 1.5)))) || true) {
                v12 := GenS2{F0: p2, F1: 42, F2: "hello"}
                switch "" {
                case "hello":
                    p0 = (p0 & -94)
                default:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                _ = v12
            } else {
                var v13 int64 = 1
                v14 := GenS1{F0: (v13 & -676), F1: 644}
                v13 -= (v13 - (v13 / (((v13 | (-(v13)))) | 1)))
                _ = v13
                _ = v14
            }
            _ = fi7
        }
        var v15 bool = (42 != 2)
        fmt.Println(p1, p0)
        _ = v15
    }
    if len(v6) > 0 {
        v6[0] = (p1 & (p1 + p1))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 506
}
func genFn1(p0 int32, p1 int64) uint {
    var v16 int64 = p1
    v17 := []int{}
    p0 = p0
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return 353
}

func main() {
    v1 := []int{2147483647, -1, -16}
    var v2 string = "foo bar"
    v3 := []int{128}
    v4 := []int{-131, 32768, 32767, 9223372036854775807}
    v5 := []int{562, 42}
    v2 = v2
    v2 = "\n\t"
    if len(v3) > 0 {
        v3[0] = 155
    }
    var v18 string = "hello"
    v18 = "a"
    var v19 uint8 = 255
    v20 := GenS1{F0: 42, F1: 2147483647}
    v2 = v18
    _ = len(v3)
    v2 = "hello"
    v18 = v18
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
