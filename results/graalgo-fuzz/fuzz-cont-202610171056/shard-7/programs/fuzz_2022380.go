package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint }
type GenS2 struct { F0 int; F1 int16 }

func main() {
    var v1 uint8 = 0
    v2 := []int{32767, 0, -1, 9223372036854775807}
    v3 := []uint32{}
    var v4 uint64 = 302
    v5 := make(map[int]bool)
    v3 = append(v3, 42)
    if (("\n\t" + "") == "a") {
        v3 = append(v3, 2147483648)
        v6 := GenS1{F0: 2147483647, F1: 18446744073709551615, F2: 4294967295}
        if (-2 >= 32768) {
            var v7 int16 = 0
            _ = v7
        } else {
            v4 += 2
            for fi8 := 0; fi8 < 14; fi8++ {
                v3 = append(v3, 2147483648)
                for fi9 := 0; fi9 < 1; fi9++ {
                    v10 := GenS1{F0: 190, F1: 208, F2: 631}
                    switch fi8 {
                    case -1:
                        v1 -= (^(v1))
                    case 128:
                        var v11 int32 = -2147483648
                        _ = v11
                    case -386:
                        if ((!false) || ((fi8 * 396) == -2147483648)) {
                            var v12 float64 = 3.14159
                            v12 -= float64(0.1)
                            v2 = append(v2, fi9)
                            _ = v12
                        }
                    }
                    _ = fi9
                    _ = v10
                }
                _ = fi8
            }
        }
        _ = v6
    }
    v13 := GenS1{F0: 2, F1: 255, F2: 0}
    var v14 float32 = 0.1
    v4 -= 4294967295
    v4 += (v4 & v4)
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 uint32 = 601
        _ = fi15
        _ = v16
    }
    v17 := GenS2{F0: 0, F1: -32768}
    var v18 int16 = -32768
    v14 -= v14
    for fi19 := 0; fi19 < 20; fi19++ {
        v20 := GenS2{F0: (fi19 / ((fi19) | 1)), F1: v18}
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 int = (-(fi21))
        fmt.Println(fi21)
        delete(v5, fi21)
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
