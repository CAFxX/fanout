package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint16, p1 string) uint8 {
    fmt.Println(p1, p1)
    var v4 uint32 = 0
    _ = p0
    _ = p1
    _ = v4
    return 94
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 42
    v3 := make(map[uint32]string)
    var v5 uint8 = 2
    v2 = 0
    switch 128 {
    case -1:
        v3[27] = (("" + "a") + "hello")
    default:
        for fi6 := 0; fi6 < 8; fi6++ {
            if false {
                for fi7 := 0; fi7 < 4; fi7++ {
                    fmt.Println(v5)
                    v8 := []int{(fi7 | 0), 900, fi6, (fi6 + (^(fi6)))}
                    switch fi7 {
                    case 42:
                        var v9 uint64 = 541
                        _ = v9
                    }
                    _ = fi7
                    _ = v8
                }
            } else {
                if false {
                    var v10 uint8 = v5
                    v11 := GenS1{F0: int8(fi6)}
                    for fi12 := 0; fi12 < 19; fi12++ {
                        v13 := []int{fi12, fi12, (fi12 * fi6)}
                        _ = fi12
                        _ = v13
                    }
                    _ = v10
                    _ = v11
                }
            }
            _ = fi6
        }
    }
    v14 := GenS1{F0: -128}
    var v15 int8 = 100
    fmt.Println(v5, v5)
    v2 += (v2 / (((v2 - v2)) | 1))
    v16 := v3[2147483647]
    if false {
        for fi17 := 0; fi17 < 4; fi17++ {
            var v18 uint8 = 0
            _ = fi17
            _ = v18
        }
        v19 := v3[28]
        _ = v19
    }
    fmt.Println(v15, v2, v5)
    var v20 float64 = 0.1
    v5 -= (v5 & v5)
    v21 := GenS2{F0: 2147483647}
    switch 2147483647 {
    case 32768:
        v1[uint32(v2)] = 2147483647
    case -2147483648:
        v22 := []int{}
        _ = v22
    }
    if ((-2147483648 == 915) || ((0 <= 32767) && (1.5 != 3.14159))) {
        if true {
            v20 = (3.14159 + 0.0)
        } else {
            v15 = (v15 << ((v2) & 63))
            v20 += v20
            for fi23 := 0; fi23 < 9; fi23++ {
                var v24 float64 = 1.5
                v15 += (v15 % (((v15 | v15)) | 1))
                _ = fi23
                _ = v24
            }
        }
    } else {
        v2 = v2
        v15 = v15
        fmt.Println(v5, v15, v16)
    }
    for fi25 := 0; fi25 < 18; fi25++ {
        v26 := []int{(fi25 << (((v2 % (((v2 ^ (v2 | v2))) | 1))) & 63))}
        _ = fi25
        _ = v26
    }
    v27 := v3[2147483647]
    var v28 uint8 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v28
    _ = v3
    _ = v5
}
