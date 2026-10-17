package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 0
    v3 := []uint32{}
    var v4 int = 0
    v5 := []uint32{2147483648, 195}
    v4 += 185
    for fi6 := 0; fi6 < 4; fi6++ {
        v4 = fi6
        v4 -= fi6
        _ = fi6
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 bool = (807 <= 61)
        _ = fi7
        _ = v8
    }
    fmt.Println(v2)
    v9 := GenS1{F0: -858, F1: -292}
    if false {
        switch ("foo bar" + "") {
        case "foo bar":
            var v10 int8 = (v2 & (^(v2)))
            _ = v10
        case "hello":
            for fi11 := 0; fi11 < 17; fi11++ {
                for fi12 := 0; fi12 < 4; fi12++ {
                    _ = len(v1)
                    v2 = (v2 ^ v2)
                    _ = fi12
                }
                _ = fi11
            }
        }
    } else {
        var v13 int64 = -71
        _ = v13
    }
    var v14 float64 = (1.5 * float64(0.0))
    for fi15 := 0; fi15 < 3; fi15++ {
        v4 = fi15
        _ = fi15
    }
    v14 = -2.25
    v3 = append(v3, 137)
    var v16 int32 = 1
    v14 -= v14
    v17 := GenS1{F0: (v16 / ((int32(v4)) | 1)), F1: 1}
    v2 -= v2
    for fi18 := 0; fi18 < 8; fi18++ {
        fmt.Println(v4, v14, fi18)
        _ = fi18
    }
    v19 := v1[2147483648]
    v19 = v19
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
