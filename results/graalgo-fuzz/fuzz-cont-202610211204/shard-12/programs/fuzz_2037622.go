package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 bool }

func main() {
    var v1 float64 = 0.0
    var v2 int8 = -128
    v3 := make(map[uint32]int)
    var v4 uint16 = 2
    if (0 == 4294967295) {
        var v5 uint = 322
        _ = v5
    } else {
        var v6 int32 = 2147483647
        _ = v6
    }
    var v7 bool = (42 < 42)
    v8 := GenS1{F0: v2, F1: -807, F2: 2}
    v4 += 255
    fmt.Println(v1, v7)
    v2 = (v2 % (((v2 ^ v2)) | 1))
    if false {
        var v9 int8 = (v2 % ((v2) | 1))
        _ = v9
    } else {
        v10 := GenS2{F0: 0, F1: 1}
        _ = v10
    }
    var v11 float64 = (float64(3.14159) + (v1 / (v1 / ((v1 * -2.25) - v1))))
    var v12 int16 = -1
    v7 = (v7 && (!v7))
    v13 := GenS3{F0: (-2147483648 >= -1)}
    if (!((true && (v7 || v7)) && v7)) {
        if (357 > 501) {
            v1 -= ((1.5 - (v11 - -2.25)) + (v11 * v11))
            v11 = v1
        }
    } else {
        v14 := GenS1{F0: (v2 / ((v2) | 1)), F1: (^(v12)), F2: 1}
        if v7 {
            v12 += (v12 << ((18446744073709551615) & 63))
            if false {
                if ((-2147483648 > 2147483647) || v7) {
                    if v7 {
                        var v15 uint8 = 42
                        _ = v15
                    } else {
                        v16 := []int{127, 32768, 1}
                        _ = len(v3)
                        _ = v16
                    }
                } else {
                    var v17 uint8 = 2
                    _ = v17
                }
                if v7 {
                    v18 := []int{756, 128, 32767, 9223372036854775807}
                    _ = v18
                }
                var v19 int64 = 1
                _ = v19
            }
        } else {
            v11 = v1
            v1 += (v1 - (float64((float32((v1 + 1e10)) + (-2.25 / (1e10 + 1.5)))) / float64((1e10 - float32((v11 + v11))))))
        }
        v1 += v11
        _ = v14
    }
    var v20 uint32 = 4294967295
    var v21 uint32 = 1
    fmt.Println(v12)
    for fi22 := 0; fi22 < 16; fi22++ {
        var v23 string = ("\n\t" + "\n\t")
        v23 = ("a" + "\n\t")
        v4 = v4
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
