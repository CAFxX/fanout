package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int16, p1 string, p2 int16) bool {
    var v5 uint16 = 0
    var v6 int32 = 2147483647
    for fi7 := 0; fi7 < 19; fi7++ {
        switch 0 {
        case 32767:
            for fi8 := 0; fi8 < 10; fi8++ {
                var v9 float64 = (0.0 * (3.14159 * 0.0))
                v10 := []int{(fi7 | 629), (-(fi8)), -9223372036854775808}
                v11 := GenS1{F0: -101}
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
        case -1:
            var v12 uint8 = 0
            _ = v12
        case -2147483648:
            if (42 > 575) {
                v13 := []int{fi7, (fi7 & 0)}
                p0 += p0
                _ = v13
            } else {
                fmt.Println(v6)
            }
        default:
            v14 := GenS1{F0: 0}
            _ = v14
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return false
}
func genFn1(p0 uint64, p1 uint64) int16 {
    p1 += (p0 - (p0 * (^(p0))))
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 int64 = 0
        var v17 uint64 = p1
        v18 := []int{32767, (fi15 / (((fi15 - (fi15 & (fi15 & fi15)))) | 1)), fi15, fi15}
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    v19 := []int{127, 128, 890, -1}
    _ = p0
    _ = p1
    _ = v19
    return 0
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[int]uint32)
    var v3 float64 = float64(float32(3.14159))
    var v4 int32 = 744
    var v20 int = 9223372036854775807
    v21 := v2[(v20 << ((4294967295) & 63))]
    v22 := GenS1{F0: -1}
    v4 += (v4 + v4)
    v4 -= v4
    var v23 uint16 = (v1 << ((220) & 63))
    var v24 int32 = (v4 << ((42) & 63))
    fmt.Println(v21, v1)
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
