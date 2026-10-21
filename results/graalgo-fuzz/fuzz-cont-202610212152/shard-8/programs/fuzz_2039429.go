package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 int64 = 1
    v2 := make(map[string]int)
    var v3 uint32 = 830
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]uint32)
    v6 := []uint32{(v3 - 65535), (^(v3)), v3, (v3 / ((42) | 1))}
    v1 += -9223372036854775808
    v3 -= v3
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 bool = (fi7 != (fi7 / ((fi7) | 1)))
        fmt.Println(v8, v3, v8)
        for fi9 := 0; fi9 < 19; fi9++ {
            v3 = 2
            v1 = -9223372036854775808
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    var v10 bool = true
    for fi11 := 0; fi11 < 14; fi11++ {
        if (float64(0.1) > float64(1.5)) {
            for fi12 := 0; fi12 < 4; fi12++ {
                if len(v6) > 0 {
                    v6[0] = v3
                }
                v3 = 950
                v3 -= (v3 << ((65535) & 63))
                _ = fi12
            }
            v13 := GenS1{F0: 553}
            _ = v13
        } else {
            v14 := GenS2{F0: 2, F1: (v3 % (((v3 + v3)) | 1))}
            if ((v3 - (v3 / ((v3) | 1))) < v3) {
                var v15 uint32 = (v3 % (((v3 << ((204) & 63))) | 1))
                _ = v15
            }
            var v16 float32 = -2.25
            _ = v14
            _ = v16
        }
        _ = fi11
    }
    v10 = v10
    for fi17 := 0; fi17 < 17; fi17++ {
        v18 := GenS3{F0: 38}
        _ = fi17
        _ = v18
    }
    fmt.Println(v1, v10, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
