package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int) uint {
    v4 := GenS1{F0: "\n\t", F1: ""}
    if (249 <= 4294967295) {
        v5 := []int{(p0 * (p0 & -901))}
        _ = v5
    }
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS2{F0: 175, F1: -931}
        var v8 uint32 = 560
        v9 := []int{fi6, (^(p0)), (fi6 / ((4294967295) | 1))}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v4
    return 2
}

func main() {
    var v1 uint = 1
    var v2 float64 = float64((float32(((1.5 / ((0.1 / 0.1) / -2.25)) / 1.0)) - 0.1))
    v3 := make(map[int]uint32)
    for fi10 := 0; fi10 < 11; fi10++ {
        switch 32768 {
        case 0:
            v2 -= (0.1 + 1e10)
        default:
            if (42 != -937) {
                var v11 int32 = 2147483647
                if (!(v1 >= (v1 % ((v1) | 1)))) {
                    v11 = v11
                    fmt.Println(v11)
                } else {
                    if (((v2 - (float64(1e10) - ((0.1 - (v2 - ((float64(3.14159) * (1.5 * v2)) - 0.1))) / (float64((1e10 - (1e10 + float32(v2)))) + ((v2 + 0.0) - v2))))) / v2) <= -2.25) {
                        if false {
                            v1 -= genFn0(fi10)
                            v11 -= -2147483648
                        }
                        fmt.Println(v2)
                    }
                    v1 = v1
                }
                v11 = (v11 / (((v11 << ((genFn0(fi10)) & 63))) | 1))
                _ = v11
            }
        }
        _ = fi10
    }
    v12 := GenS3{F0: 179, F1: 778}
    for fi13 := 0; fi13 < 13; fi13++ {
        _ = len(v3)
        _ = fi13
    }
    v3[2147483647] = uint32(v1)
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 int32 = 0
        v2 = v2
        _ = fi14
        _ = v15
    }
    v1 = 0
    var v16 int16 = -1
    v17 := []int{266, -895}
    delete(v3, 4294967295)
    v2 = -2.25
    delete(v3, 863)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
