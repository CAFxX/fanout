package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
func genFn0(p0 int32, p1 bool, p2 int16) bool {
    p2 = (^(p2))
    p2 += 1
    if ((p0 | 42) >= -2147483648) {
        fmt.Println(p0, p1)
    }
    v5 := []int{0, -2147483648, 4294967295, -11}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p1
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{937, 502}
    var v3 int64 = -636
    var v4 int8 = -24
    _ = len(v2)
    switch "" {
    case "\n\t":
        if true {
            var v6 float64 = float64(1.5)
            _ = v6
        }
    case "a":
        v3 -= (v3 * (-(v3)))
    default:
        var v7 bool = (!(1e10 < -2.25))
        _ = v7
    }
    var v8 uint8 = 128
    v8 -= v8
    for fi9 := 0; fi9 < 18; fi9++ {
        switch (fi9 * fi9) {
        case 9223372036854775807:
            for fi10 := 0; fi10 < 14; fi10++ {
                if true {
                    v11 := GenS1{F0: 955}
                    fmt.Println(v4, fi9, v8)
                    switch ("a" + ("\n\t" + "hello")) {
                    case "\n\t":
                        v8 = (v8 >> ((776) & 63))
                    default:
                        delete(v1, fi9)
                    }
                    _ = v11
                } else {
                    v12 := GenS1{F0: 1}
                    v4 -= 127
                    _ = v12
                }
                fmt.Println(fi10, v8, fi9)
                _ = fi10
            }
        }
        _ = fi9
    }
    v4 = v4
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    var v13 uint32 = 579
    v14 := GenS2{F0: (false && (v4 != -94))}
    var v15 int16 = 0
    fmt.Println(v4, v3, v15)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
