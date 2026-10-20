package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int8; F1 uint8; F2 int }

func main() {
    v1 := []int{127, -134}
    v2 := make(map[int]string)
    v3 := make(map[uint32]uint32)
    var v4 int8 = 1
    v5 := GenS1{F0: 32, F1: 508}
    for fi6 := 0; fi6 < 15; fi6++ {
        v4 -= 1
        v7 := GenS2{F0: (-(v4)), F1: 0, F2: (fi6 ^ fi6)}
        _ = fi6
        _ = v7
    }
    v8 := v3[42]
    switch 128 {
    case -1:
        if (!true) {
            _ = len(v2)
        } else {
            var v9 float32 = (float32(float64(1e10)) + float32(1e10))
            for fi10 := 0; fi10 < 2; fi10++ {
                if (945 == 0) {
                    v8 = v8
                }
                for fi11 := 0; fi11 < 6; fi11++ {
                    _ = len(v3)
                    fmt.Println(fi10, v9, fi11)
                    _ = fi11
                }
                _ = fi10
            }
            _ = v9
        }
    }
    var v12 int8 = (v4 >> ((723) & 63))
    v13 := v3[v8]
    for fi14 := 0; fi14 < 9; fi14++ {
        if false {
            v13 = 1
            v15 := GenS2{F0: (v4 % (((v12 | -32)) | 1)), F1: 241, F2: (fi14 / ((fi14) | 1))}
            v16 := GenS2{F0: 106, F1: 2, F2: (fi14 | fi14)}
            _ = v15
            _ = v16
        } else {
            fmt.Println(v4, v13)
            v17 := GenS1{F0: 731, F1: 2}
            switch ("foo bar" + "hello") {
            case "\n\t":
                var v18 int32 = -2147483648
                _ = v18
            case "foo bar":
                if len(v1) > 0 {
                    v1[0] = (fi14 >> ((782) & 63))
                }
            default:
                v19 := GenS1{F0: -1, F1: 1}
                _ = v19
            }
            _ = v17
        }
        _ = fi14
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := v3[1]
        _ = fi20
        _ = v21
    }
    fmt.Println(v12)
    var v22 uint8 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
