package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 uint; F2 uint64 }

func main() {
    v1 := []uint32{2, 1}
    var v2 int32 = 1
    v3 := make(map[uint32]string)
    var v4 uint8 = 255
    v5 := make(map[int]uint32)
    switch -717 {
    case 32767:
        v6 := GenS1{F0: 65535}
        _ = v6
    case -2147483648:
        v2 = -2147483648
    case -216:
        v5[127] = 373
    default:
        v7 := v3[4294967295]
        _ = v7
    }
    v4 -= v4
    fmt.Println(v2, v2)
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint8 = v4
        _ = fi8
        _ = v9
    }
    if (((1.5 + float64((((float32(0.1) * 1e10) / 0.1) / ((1e10 - 3.14159) - 1e10)))) / 3.14159) != (float64(float32(float64((1.5 * 1.5)))) - 1.5)) {
        for fi10 := 0; fi10 < 13; fi10++ {
            v4 += (v4 * (v4 - 255))
            for fi11 := 0; fi11 < 10; fi11++ {
                v12 := GenS2{F0: (fi10 % ((398) | 1))}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi11
                _ = v12
            }
            v13 := GenS3{F0: 127, F1: 18446744073709551615, F2: 1}
            _ = fi10
            _ = v13
        }
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 uint16 = 2
            _ = fi14
            _ = v15
        }
        var v16 uint32 = 1
        _ = v16
    }
    var v17 uint32 = uint32(v4)
    v18 := GenS1{F0: 1}
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 string = "a"
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        v22 := GenS1{F0: 2}
        _ = fi21
        _ = v22
    }
    var v23 uint8 = (v4 >> ((4294967295) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
