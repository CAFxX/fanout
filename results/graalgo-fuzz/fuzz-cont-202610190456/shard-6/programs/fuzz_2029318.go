package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 uint16 }

func main() {
    v1 := []int{-632, 127, 166}
    v2 := []uint32{}
    v3 := []uint32{332}
    v4 := make(map[uint32]uint32)
    var v5 int32 = 1
    v6 := GenS1{F0: (v5 + -1)}
    fmt.Println(v5)
    v5 -= (v5 % ((v5) | 1))
    if (!(!true)) {
        fmt.Println(v5)
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint32 = 127
        _ = fi7
        _ = v8
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch -656 {
    case -45:
        v5 = v5
    default:
        switch (("a" + "a") + ("foo bar" + "a")) {
        case "hello":
            for fi9 := 0; fi9 < 6; fi9++ {
                var v10 float64 = ((3.14159 * (-2.25 / ((1.5 * 1.5) + 1.5))) / -2.25)
                _ = fi9
                _ = v10
            }
        case "":
            v5 += 1
        case "\n\t":
            v5 -= (v5 + 1)
        default:
            for fi11 := 0; fi11 < 14; fi11++ {
                for fi12 := 0; fi12 < 4; fi12++ {
                    v3 = append(v3, 0)
                    v5 -= -1
                    _ = fi12
                }
                _ = fi11
            }
        }
    }
    v5 += v5
    if (((3.14159 / 1.5) / 3.14159) < 3.14159) {
        for fi13 := 0; fi13 < 17; fi13++ {
            v5 = v5
            var v14 int64 = 1
            v15 := GenS2{F0: (fi13 | -1), F1: 1}
            _ = fi13
            _ = v14
            _ = v15
        }
    } else {
        v5 -= -1
    }
    v4[42] = 165
    if (228 != 128) {
        switch -2147483648 {
        case 32768:
            delete(v4, 0)
        default:
            v5 = v5
        }
    } else {
        var v16 bool = false
        _ = v16
    }
    fmt.Println(v5)
    v17 := v4[989]
    var v18 int16 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
