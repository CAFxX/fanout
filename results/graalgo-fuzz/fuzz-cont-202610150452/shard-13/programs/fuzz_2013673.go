package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint16; F2 bool }
func genFn0(p0 int64, p1 uint32, p2 bool) uint8 {
    p2 = true
    p0 -= (p0 ^ (p0 >> ((65535) & 63)))
    for fi6 := 0; fi6 < 13; fi6++ {
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := []int{fi6, (^(fi6)), fi6}
            fmt.Println(fi7, p2, p1)
            if len(v8) > 0 {
                v8[0] = 127
            }
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: 42}
        _ = fi6
        _ = v9
    }
    v10 := []int{0, -9223372036854775808, 9223372036854775807}
    v11 := GenS1{F0: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 255
}

func main() {
    v1 := []int{}
    var v2 int = 127
    v3 := []int{32767}
    var v4 uint64 = 18446744073709551615
    v5 := make(map[int]string)
    if ((0 != -1) && true) {
        if (v4 < v4) {
            if false {
                for fi12 := 0; fi12 < 3; fi12++ {
                    v4 += (v4 >> ((65535) & 63))
                    if (-719 >= -1) {
                        _ = len(v3)
                    }
                    v13 := GenS1{F0: -171}
                    _ = fi12
                    _ = v13
                }
            }
        } else {
            var v14 float32 = (1e10 / -2.25)
            _ = v14
        }
    }
    v2 = v2
    v5[(^(v2))] = (("\n\t" + "foo bar") + "hello")
    for fi15 := 0; fi15 < 17; fi15++ {
        v2 += fi15
        v4 -= 4294967295
        fmt.Println(v4, v4)
        _ = fi15
    }
    fmt.Println(v2, v4)
    v2 = v2
    switch "hello" {
    case "hello":
        var v16 bool = true
        _ = v16
    case "\n\t":
        v2 = 32768
    case "a":
        switch (v2 ^ (v2 | v2)) {
        case 9223372036854775807:
            delete(v5, v2)
        default:
            _ = len(v3)
        }
    default:
        v17 := GenS2{F0: ((-2.25 - 0.0) >= 1e10), F1: 242, F2: ((float64(1e10) - (1.5 / 0.1)) >= 1.5)}
        _ = v17
    }
    v2 -= (v2 / (((v2 * -143)) | 1))
    fmt.Println(v2, v4)
    v18 := GenS2{F0: (!true), F1: 2, F2: true}
    v4 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
