package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint; F2 string }

func main() {
    var v1 int8 = -28
    v2 := make(map[uint32]uint32)
    v3 := []string{}
    v4 := make(map[int]uint32)
    var v5 uint8 = 128
    v6 := make(map[uint32]string)
    v6[250] = ("\n\t" + ("a" + "hello"))
    var v7 uint16 = uint16(v5)
    _ = len(v3)
    v8 := GenS1{F0: -9223372036854775808, F1: 4294967295, F2: (v5 << ((42) & 63))}
    v1 += v1
    v1 -= (v1 ^ 1)
    for fi9 := 0; fi9 < 4; fi9++ {
        v7 = (v7 & v7)
        _ = fi9
    }
    v7 += (^(v7))
    if (true || true) {
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 int32 = -832
            _ = fi10
            _ = v11
        }
        _ = len(v2)
    } else {
        switch 0 {
        case 2147483647:
            if true {
                var v12 float64 = 0.1
                _ = v12
            } else {
                fmt.Println(v5, v5, v7)
            }
        }
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        _ = len(v4)
        if (true || false) {
            if (!true) {
                if (!false) {
                    v14 := v2[324]
                    v15 := GenS2{F0: -344, F1: 18446744073709551615, F2: ("foo bar" + "hello")}
                    v7 -= (^(v7))
                    _ = v14
                    _ = v15
                } else {
                    v5 += (v5 / (((^(v5))) | 1))
                    for fi16 := 0; fi16 < 1; fi16++ {
                        if false {
                            v1 += -1
                        }
                        var v17 uint16 = (^(v7))
                        v7 += 50
                        _ = fi16
                        _ = v17
                    }
                    v7 -= 42
                }
            }
            var v18 float64 = 1e10
            _ = v18
        } else {
            if (1000 <= 1) {
                v3 = append(v3, "foo bar")
                var v19 uint16 = v7
                _ = v19
            } else {
                for fi20 := 0; fi20 < 14; fi20++ {
                    var v21 float64 = ((1.5 * 3.14159) / (1.5 * (1.5 + (((((0.1 - 0.0) - 3.14159) - -2.25) - ((0.1 + 1.5) / ((3.14159 * float64(((0.1 + 0.1) * (-2.25 / 1e10)))) + 3.14159))) - float64(-2.25)))))
                    _ = fi20
                    _ = v21
                }
                v1 = (v1 - (v1 & v1))
            }
            v5 = (v5 >> ((0) & 63))
        }
        _ = fi13
    }
    var v22 int64 = -244
    var v23 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
