package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint16; F1 uint8; F2 int64 }
func genFn0() int {
    v4 := []int{32768, 32767}
    var v5 int16 = -24
    v5 += (v5 ^ (v5 ^ v5))
    var v6 bool = (true && false)
    var v7 uint16 = 65535
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 881
}
func genFn1(p0 int8, p1 int16, p2 int8) int {
    var v8 uint32 = 65535
    var v9 int8 = -115
    var v10 int = 694
    if (p1 != p1) {
        v11 := []int{(v10 & v10), v10}
        v8 = (v8 + (v8 + 741))
        v8 += 733
        _ = v11
    } else {
        var v12 uint8 = 23
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 836
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 0
    v3 := []int{9223372036854775807}
    var v13 uint8 = 128
    fmt.Println(v13, v2)
    v14 := GenS1{F0: -1, F1: false}
    if true {
        v13 = (v13 >> ((255) & 63))
        for fi15 := 0; fi15 < 13; fi15++ {
            fmt.Println(v13, v2)
            _ = fi15
        }
        switch ("foo bar" + "\n\t") {
        case "hello":
            if (((true && (0 != -253)) && true) || (true && (32768 != 32767))) {
                v13 += v13
                if len(v3) > 0 {
                    v3[0] = -656
                }
                v16 := GenS1{F0: -417, F1: true}
                _ = v16
            } else {
                _ = len(v3)
                for fi17 := 0; fi17 < 8; fi17++ {
                    for fi18 := 0; fi18 < 4; fi18++ {
                        var v19 bool = (!(32767 == 136))
                        _ = fi18
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 16; fi20++ {
                        v2 += -62
                        v2 = (v2 ^ v2)
                        _ = fi20
                    }
                    fmt.Println(v2, v2)
                    _ = fi17
                }
            }
        case "\n\t":
            v13 = 158
        case "":
            fmt.Println(v2, v2)
        default:
            var v21 uint = 42
            _ = v21
        }
    } else {
        var v22 bool = (-28 == -1)
        if (42 == 42) {
            var v23 int8 = 127
            if false {
                v24 := GenS1{F0: 42, F1: v22}
                v2 += 661
                _ = v24
            }
            _ = v23
        }
        _ = v22
    }
    v25 := GenS2{F0: 1, F1: 255, F2: v2}
    v2 -= v2
    fmt.Println(v13, v13)
    v2 += (v2 % ((v2) | 1))
    v1[337] = genFn0()
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v3
}
