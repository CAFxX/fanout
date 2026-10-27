package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint16 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint32 }
func genFn0() uint64 {
    for fi4 := 0; fi4 < 17; fi4++ {
        var v5 int64 = -1
        v6 := GenS1{F0: (-(v5)), F1: false, F2: 255}
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 string = ("" + "")
    var v8 int16 = -32768
    v9 := GenS1{F0: -9223372036854775808, F1: (!true), F2: 32767}
    v10 := GenS2{F0: -449, F1: true}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 float32 = float32(0.0)
    var v2 uint32 = 0
    v3 := make(map[int]uint32)
    v1 += (3.14159 * ((v1 - 0.1) + float32(float64(v1))))
    v2 += (v2 << ((1) & 63))
    for fi11 := 0; fi11 < 2; fi11++ {
        for fi12 := 0; fi12 < 6; fi12++ {
            if (v2 >= (v2 + v2)) {
                v1 += (0.0 - 0.1)
                switch "hello" {
                case "hello":
                    var v13 uint32 = (^(v2))
                    _ = v13
                case "\n\t":
                    fmt.Println(v1, v2)
                case "a":
                    v2 = v2
                default:
                    v1 = 1e10
                }
                var v14 bool = false
                _ = v14
            }
            v15 := GenS1{F0: 9223372036854775807, F1: (0 > -128), F2: 255}
            _ = fi12
            _ = v15
        }
        _ = fi11
    }
    if (!(("\n\t" != "hello") && ((1.5 == v1) || false))) {
        switch 0 {
        case 1:
            v1 = (v1 * float32(1.5))
        }
    }
    v16 := GenS1{F0: 0, F1: (778 <= -203), F2: 1}
    v17 := GenS3{F0: v2}
    v18 := []int{1, 32767, -509, int(v2)}
    _ = len(v3)
    fmt.Println(v1, v1)
    fmt.Println(v2)
    var v19 int64 = 0
    _ = len(v3)
    fmt.Println(v1, v1, v1)
    v2 = (v2 | v2)
    for fi20 := 0; fi20 < 17; fi20++ {
        v2 = (v2 * 2)
        switch fi20 {
        case 79:
            var v21 uint = 0
            _ = v21
        case 127:
            v3[fi20] = v2
        case 0:
            v22 := GenS1{F0: (v19 - v19), F1: false, F2: 758}
            _ = v22
        }
        var v23 int = fi20
        _ = fi20
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
