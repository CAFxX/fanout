package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int }
func genFn0(p0 uint8, p1 int8) uint8 {
    v5 := GenS1{F0: p0}
    for fi6 := 0; fi6 < 4; fi6++ {
        p0 += 42
        _ = fi6
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        p0 += (p0 << ((65535) & 63))
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v5
    return (p0 - 128)
}
func genFn1() uint {
    v8 := []int{}
    var v9 uint = 42
    v9 += 0
    _ = v8
    _ = v9
    return (v9 & 2)
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[int]uint32)
    var v3 float32 = 1e10
    var v4 int32 = -2147483648
    v1 = v1
    for fi10 := 0; fi10 < 9; fi10++ {
        v4 = 430
        for fi11 := 0; fi11 < 2; fi11++ {
            if true {
                var v12 int64 = -975
                fmt.Println(v3)
                _ = v12
            }
            v1 -= (((v1 + v1) - 3.14159) - v1)
            _ = fi11
        }
        v13 := GenS2{F0: 42}
        _ = fi10
        _ = v13
    }
    fmt.Println(v1)
    switch -449 {
    case -453:
        fmt.Println(v3, v1)
    default:
        for fi14 := 0; fi14 < 9; fi14++ {
            if (42 > 42) {
                for fi15 := 0; fi15 < 19; fi15++ {
                    var v16 uint = 1
                    switch 32767 {
                    case 689:
                        var v17 float64 = v1
                        _ = v17
                    case -2147483648:
                        v1 = v1
                    default:
                        if (!true) {
                            v2[(fi15 / (((fi14 & 0)) | 1))] = 1
                        }
                    }
                    _ = fi15
                    _ = v16
                }
                v18 := GenS2{F0: 42}
                _ = v18
            }
            _ = fi14
        }
    }
    var v19 float32 = v3
    v1 = (v1 + float64(v19))
    v1 = ((float64(v19) / v1) - (1.5 / 1.0))
    v20 := GenS3{F0: 297}
    var v21 float32 = (v3 * v3)
    fmt.Println(v4, v4, v4)
    v22 := GenS2{F0: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
