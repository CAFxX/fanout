package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }

func main() {
    var v1 uint16 = 32767
    var v2 int = 2147483647
    var v3 int = (-(v2))
    v4 := []int{(v3 * v2), v3, (v3 & (v2 / (((v2 * (-(v3)))) | 1))), v2}
    v2 = v3
    var v5 uint16 = (v1 >> ((2) & 63))
    fmt.Println(v2, v1, v1)
    var v6 int32 = 0
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v7 int16 = -376
    if ("foo bar" <= "") {
        if true {
            v1 = v5
            v8 := GenS1{F0: -457, F1: (v7 * (^(v7)))}
            _ = v8
        } else {
            v4 = append(v4, v2)
        }
        fmt.Println(v1, v7, v5)
    } else {
        for fi9 := 0; fi9 < 20; fi9++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi9
        }
    }
    v5 += (v1 ^ 1)
    var v10 float64 = 1.5
    for fi11 := 0; fi11 < 11; fi11++ {
        if (2 != 69) {
            for fi12 := 0; fi12 < 13; fi12++ {
                v4 = append(v4, fi11)
                switch "foo bar" {
                case "foo bar":
                    var v13 bool = false
                    _ = v13
                case "hello":
                    v2 -= (fi12 | v3)
                default:
                    fmt.Println(v2)
                    v14 := GenS1{F0: -32768, F1: (v7 & v7)}
                    _ = v14
                }
                _ = fi12
            }
            var v15 uint8 = 159
            v16 := GenS1{F0: v7, F1: v7}
            _ = v15
            _ = v16
        } else {
            var v17 bool = false
            if len(v4) > 0 {
                v4[0] = 24
            }
            switch fi11 {
            case 2147483647:
                var v18 float32 = -2.25
                _ = v18
            case 9223372036854775807:
                v4 = append(v4, (v3 >> ((0) & 63)))
            case 127:
                for fi19 := 0; fi19 < 18; fi19++ {
                    if len(v4) > 0 {
                        v4[0] = fi11
                    }
                    _ = fi19
                }
            }
            _ = v17
        }
        _ = fi11
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        v21 := GenS1{F0: -252, F1: v7}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi20
        _ = v21
    }
    v22 := GenS1{F0: v7, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
