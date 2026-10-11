package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int16; F1 bool; F2 uint16 }
func genFn0() bool {
    var v4 int32 = -415
    v4 = 2147483647
    _ = v4
    return (4294967295 > 4294967295)
}
func genFn1(p0 uint, p1 uint32) uint {
    p0 = 0
    v5 := GenS1{F0: (-5 > -1), F1: 128, F2: 56}
    switch 142 {
    case -732:
        v6 := []int{-534, 2147483647, -1, 4294967295}
        _ = v6
    case -2147483648:
        switch 128 {
        case -1:
            if false {
                for fi7 := 0; fi7 < 11; fi7++ {
                    var v8 int16 = 0
                    switch fi7 {
                    case 32768:
                        v8 = -1
                    }
                    fmt.Println(p0, p0, v8)
                    _ = fi7
                    _ = v8
                }
            } else {
                for fi9 := 0; fi9 < 14; fi9++ {
                    var v10 int64 = -1
                    _ = fi9
                    _ = v10
                }
                for fi11 := 0; fi11 < 11; fi11++ {
                    v12 := []int{-9223372036854775808, int(p0), (fi11 << (((p0 / (((p0 * 34)) | 1))) & 63))}
                    var v13 uint16 = 65535
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                p1 = 2147483648
            }
        case 405:
            v14 := GenS2{F0: -960, F1: (!(1e10 < 0.1)), F2: 0}
            _ = v14
        case 42:
            p1 += 1
        default:
            var v15 uint64 = 2
            _ = v15
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return (p0 + (p0 | p0))
}

func main() {
    v1 := make(map[int]string)
    v2 := []int{127, 2147483647, 32767}
    var v3 float32 = float32(0.1)
    v16 := GenS2{F0: 1000, F1: (false && (812 >= -2147483648)), F2: 2}
    for fi17 := 0; fi17 < 11; fi17++ {
        v3 -= (v3 + (1e10 - -2.25))
        _ = fi17
    }
    var v18 float32 = v3
    fmt.Println(v18)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 += float32(((0.0 + float64(v3)) / (-2.25 - float64(float32(((0.1 + 1e10) * ((-2.25 - (-2.25 - -2.25)) * 1.5)))))))
    fmt.Println(v18, v3)
    v19 := GenS1{F0: ((((0.0 + 1.5) + float64(v3)) / -2.25) != (0.1 - 0.0)), F1: 128, F2: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
