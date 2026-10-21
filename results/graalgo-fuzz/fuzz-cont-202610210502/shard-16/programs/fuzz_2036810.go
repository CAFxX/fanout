package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 int32; F2 int8 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint32, p1 uint8) bool {
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 int32 = -1
        var v9 int = (fi7 & (^(fi7)))
        switch "" {
        case "hello":
            var v10 uint32 = p0
            _ = v10
        case "foo bar":
            for fi11 := 0; fi11 < 13; fi11++ {
                v8 -= v8
                _ = fi11
            }
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v12 int32 = 1
    _ = p0
    _ = p1
    _ = v12
    return (2147483647 == (p0 << ((0) & 63)))
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]string)
    v3 := []uint32{}
    var v4 float32 = 1e10
    var v5 int32 = 0
    var v6 float32 = 1e10
    v13 := GenS1{F0: 1, F1: 0, F2: (1000 == 42)}
    fmt.Println(v4, v4)
    v6 = v6
    v14 := GenS2{F0: 2147483647}
    delete(v2, "hello")
    for fi15 := 0; fi15 < 19; fi15++ {
        if (426 < 4294967295) {
            var v16 float64 = float64(v4)
            _ = v16
        } else {
            v3 = append(v3, 2147483648)
        }
        _ = fi15
    }
    v4 = (v6 - v6)
    v4 -= float32(((1e10 * (((3.14159 + ((((0.0 / 0.1) - 1.5) * float64(float32(1.5))) - (((0.0 * ((1e10 + (float64((v4 + -2.25)) / 3.14159)) / 1.0)) / float64(1.5)) + 0.1))) * 0.0) - 1e10)) * 1.5))
    v17 := GenS3{F0: -2147483648, F1: v5, F2: 127}
    delete(v2, ("foo bar" + "hello"))
    if false {
        v6 = v6
        for fi18 := 0; fi18 < 6; fi18++ {
            if (!(-1 > 0)) {
                v19 := GenS2{F0: (fi18 * (fi18 / ((1) | 1)))}
                _ = v19
            } else {
                v20 := GenS4{F0: true}
                if (!false) {
                    v21 := GenS1{F0: 9223372036854775807, F1: 4294967295, F2: (128 != 128)}
                    _ = v21
                }
                var v22 uint32 = 2147483648
                _ = v20
                _ = v22
            }
            var v23 uint8 = 255
            _ = fi18
            _ = v23
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
