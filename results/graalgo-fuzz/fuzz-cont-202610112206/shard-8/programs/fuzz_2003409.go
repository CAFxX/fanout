package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }

func main() {
    var v1 int8 = -1
    v2 := []int{-2147483648, 4294967295}
    v3 := make(map[string]int)
    v4 := make(map[string]uint32)
    v5 := []uint32{693, 625}
    v6 := GenS1{F0: (v1 + -3), F1: 176}
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int = (-(fi7))
        for fi9 := 0; fi9 < 19; fi9++ {
            fmt.Println(v1, fi7, v8)
            v8 += (-(fi7))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi9
        }
        v8 = -787
        _ = fi7
        _ = v8
    }
    var v10 uint32 = 65535
    for fi11 := 0; fi11 < 19; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            if (0.0 != (3.14159 / (-2.25 * (((-2.25 - 1.5) + float32(0.0)) / 0.1)))) {
                var v13 string = (("\n\t" + "hello") + "a")
                switch fi11 {
                case -1:
                    var v14 int32 = -480
                    _ = v14
                case 127:
                    v10 -= v10
                case 0:
                    v13 = (v13 + v13)
                }
                _ = v13
            }
            var v15 uint8 = 230
            _ = fi12
            _ = v15
        }
        v1 += (^(v1))
        _ = fi11
    }
    var v16 uint32 = v10
    var v17 uint16 = 2
    v10 = v10
    v17 = 255
    fmt.Println(v1, v17)
    v17 = 42
    var v18 uint64 = 677
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
