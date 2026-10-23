package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 int; F1 uint; F2 int32 }
func genFn0(p0 int, p1 string) uint8 {
    p0 = p0
    p0 -= p0
    fmt.Println(p0)
    for fi5 := 0; fi5 < 12; fi5++ {
        p1 = "a"
        if true {
            v6 := GenS1{F0: 32767, F1: 4294967295, F2: 128}
            if ((0.0 < float32(-2.25)) || false) {
                if (!false) {
                    v7 := GenS1{F0: 1, F1: 2147483647, F2: 2}
                    p0 = -9223372036854775808
                    var v8 uint8 = 1
                    _ = v7
                    _ = v8
                } else {
                    var v9 uint8 = 2
                    var v10 float32 = ((0.0 * 3.14159) + 0.1)
                    _ = v10
                    _ = v9
                }
            }
            if (2 >= 848) {
                v11 := GenS1{F0: 220, F1: 0, F2: 2}
                _ = v11
            }
            _ = v6
        } else {
            var v12 float32 = 3.14159
            var v13 int32 = 497
            _ = v12
            _ = v13
        }
        v14 := GenS1{F0: 25, F1: 855, F2: 1}
        _ = fi5
        _ = v14
    }
    _ = p0
    _ = p1
    return 128
}

func main() {
    var v1 int16 = -11
    v2 := []uint32{}
    v3 := make(map[uint32]int)
    v4 := make(map[string]uint32)
    v1 += (^(v1))
    v4[(("\n\t" + "a") + "\n\t")] = 675
    var v15 int32 = int32(v1)
    v15 = v15
    var v16 float64 = 0.1
    if (((v1 / (((v1 / ((v1) | 1))) | 1)) > v1) && ((128 >= 1) || (!(!(-2147483648 < 42))))) {
        v17 := GenS2{F0: (-505 >= 32768), F1: 2}
        for fi18 := 0; fi18 < 9; fi18++ {
            for fi19 := 0; fi19 < 13; fi19++ {
                v20 := GenS3{F0: 70, F1: (v15 << ((65535) & 63))}
                _ = fi19
                _ = v20
            }
            var v21 bool = (true && false)
            _ = fi18
            _ = v21
        }
        v22 := GenS4{F0: 9223372036854775807, F1: 4294967295, F2: (v15 & (v15 ^ v15))}
        _ = v17
        _ = v22
    }
    v1 += (v1 / ((0) | 1))
    var v23 bool = (false || true)
    var v24 float64 = 3.14159
    v16 -= (-2.25 - ((float64(3.14159) * v24) + v24))
    var v25 uint8 = 92
    for fi26 := 0; fi26 < 12; fi26++ {
        var v27 uint = 4294967295
        _ = fi26
        _ = v27
    }
    var v28 uint8 = (v25 + (v25 / ((128) | 1)))
    v23 = v23
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
}
