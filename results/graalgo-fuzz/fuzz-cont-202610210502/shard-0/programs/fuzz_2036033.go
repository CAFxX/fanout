package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 int64; F2 int32 }
func genFn0(p0 uint16, p1 uint64) uint16 {
    var v4 float64 = 1.5
    switch 9223372036854775807 {
    case -1:
        var v5 bool = (float32(0.0) <= 1.5)
        _ = v5
    default:
        for fi6 := 0; fi6 < 19; fi6++ {
            v7 := GenS1{F0: (fi6 % (((^(fi6))) | 1))}
            v8 := GenS1{F0: (fi6 / (((fi6 & fi6)) | 1))}
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    if false {
        var v9 uint16 = 42
        var v10 float32 = -2.25
        _ = v10
        _ = v9
    } else {
        v4 += 0.0
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 uint = 959
            _ = fi11
            _ = v12
        }
        switch -281 {
        case 684:
            fmt.Println(p0, v4)
        case 32768:
            p0 -= 276
        }
    }
    v4 += (1e10 + ((float64(((0.0 - ((-2.25 * 3.14159) - (((0.0 * (3.14159 - (0.0 - (1.5 + 3.14159)))) + 3.14159) + ((0.0 + 0.0) + 0.1)))) * -2.25)) + (v4 * v4)) * (float64(1.5) / v4)))
    _ = p0
    _ = p1
    _ = v4
    return (p0 - (p0 ^ (^(p0))))
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{32767, 396}
    v3 := make(map[int]string)
    var v13 float32 = 1.5
    v14 := GenS1{F0: 0}
    switch 0 {
    case -2147483648:
        for fi15 := 0; fi15 < 16; fi15++ {
            v13 = 0.0
            _ = fi15
        }
    }
    fmt.Println(v13)
    v3[-2147483648] = "foo bar"
    var v16 int32 = -1
    v13 = (1e10 * (float32(float64(v13)) - 0.1))
    for fi17 := 0; fi17 < 7; fi17++ {
        for fi18 := 0; fi18 < 17; fi18++ {
            v19 := GenS1{F0: (fi17 >> ((0) & 63))}
            var v20 float64 = (((-2.25 - ((float64((v13 + (1.5 - (v13 - (-2.25 - 0.0))))) + float64(0.1)) * (((1e10 + 1e10) / 0.1) * (0.1 - 0.0)))) / 1.0) / float64((((float32(1e10) + (0.0 - 0.0)) + v13) - (v13 - (float32(1e10) * (((0.0 * float32(1e10)) - 1e10) / ((0.1 + float32(0.1)) + ((3.14159 / 3.14159) - v13))))))))
            v21 := v1["a"]
            v20 = -2.25
            v22 := GenS1{F0: fi18}
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
            _ = v22
        }
        v23 := v1["hello"]
        _ = fi17
        _ = v23
    }
    v24 := GenS2{F0: 1, F1: 1, F2: v16}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v24
    _ = v3
}
