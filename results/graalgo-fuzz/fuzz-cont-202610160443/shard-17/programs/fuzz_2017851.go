package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint8; F1 int32; F2 int64 }
func genFn0(p0 uint64, p1 int8) int32 {
    if true {
        for fi4 := 0; fi4 < 7; fi4++ {
            if ((255 == 248) || (126 < (p0 ^ (p0 >> ((1) & 63))))) {
                if (45 >= 1) {
                    p1 -= (p1 % ((0) | 1))
                    p0 = (p0 ^ p0)
                } else {
                    for fi5 := 0; fi5 < 3; fi5++ {
                        var v6 uint8 = 2
                        var v7 uint16 = 2
                        _ = fi5
                        _ = v6
                        _ = v7
                    }
                    v8 := []int{}
                    _ = v8
                }
                var v9 int = (^(fi4))
                var v10 int8 = (p1 + p1)
                _ = v10
                _ = v9
            } else {
                for fi11 := 0; fi11 < 7; fi11++ {
                    var v12 int32 = 42
                    p1 -= (p1 / (((p1 & (p1 * (-(p1))))) | 1))
                    _ = fi11
                    _ = v12
                }
                fmt.Println(p0, p0, p0)
                v13 := []int{fi4}
                _ = v13
            }
            _ = fi4
        }
    } else {
        fmt.Println(p1, p0)
        var v14 int32 = -2147483648
        for fi15 := 0; fi15 < 1; fi15++ {
            v16 := []int{fi15, fi15}
            if (65535 <= 513) {
                v14 = v14
                v17 := GenS1{F0: 1}
                _ = v17
            } else {
                if len(v16) > 0 {
                    v16[0] = (-(fi15))
                }
                p0 -= 4294967295
                v18 := GenS1{F0: 255}
                _ = v18
            }
            var v19 uint16 = 2
            _ = fi15
            _ = v16
            _ = v19
        }
        _ = v14
    }
    var v20 int16 = 32767
    v21 := GenS2{F0: ((float64(3.14159) * 3.14159) >= (-2.25 - 1e10)), F1: -585}
    v22 := []int{9223372036854775807, 2147483647, 32768}
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    _ = v22
    return -2147483648
}

func main() {
    v1 := []int{}
    v2 := []uint32{4294967295, 42, 4294967295, 4294967295}
    var v3 bool = false
    var v23 uint64 = 18446744073709551615
    var v24 uint64 = v23
    v2 = append(v2, 1)
    if len(v2) > 0 {
        v2[0] = 369
    }
    if len(v2) > 0 {
        v2[0] = 617
    }
    v25 := GenS3{F0: 105, F1: 393, F2: -642}
    fmt.Println(v23)
    v2 = append(v2, 65535)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
