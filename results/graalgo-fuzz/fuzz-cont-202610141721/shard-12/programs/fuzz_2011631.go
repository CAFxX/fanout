package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 int; F1 uint; F2 int64 }
func genFn0(p0 bool) uint16 {
    if p0 {
        var v4 bool = (true && ((1 == 1) || (float32(0.0) != ((float32(-2.25) / ((float32(-2.25) + 0.0) / (1e10 / -2.25))) - 1.5))))
        _ = v4
    }
    var v5 string = "foo bar"
    p0 = (!(1 != 42))
    _ = p0
    _ = v5
    return 0
}
func genFn1(p0 uint, p1 int32) int8 {
    p0 = (^(p0))
    var v6 uint8 = 82
    _ = p0
    _ = p1
    _ = v6
    return -1
}

func main() {
    var v1 int16 = -935
    var v2 uint64 = 2
    var v3 uint8 = 91
    var v7 int = -230
    v8 := GenS1{F0: 1, F1: v7}
    var v9 int8 = genFn1(18446744073709551615, -1)
    var v10 uint = 4294967295
    var v11 float32 = 1.5
    for fi12 := 0; fi12 < 16; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            if true {
                if (v11 < (v11 + 1e10)) {
                    var v14 int = (fi13 / (((fi13 & 4294967295)) | 1))
                    _ = v14
                } else {
                    var v15 uint32 = 0
                    _ = v15
                }
                var v16 uint64 = (v2 % (((v2 % ((v2) | 1))) | 1))
                v11 += v11
                _ = v16
            } else {
                v11 = v11
            }
            _ = fi13
        }
        _ = fi12
    }
    v17 := []int{v7, v7, int(v9)}
    var v18 int16 = v1
    v7 -= (-(v7))
    _ = len(v17)
    if ((v3 - (^(v3))) <= v3) {
        if (1e10 == -2.25) {
            v19 := GenS2{F0: 9223372036854775807}
            v3 = v3
            v9 += genFn1((v10 & v10), 2147483647)
            _ = v19
        } else {
            var v20 int64 = 831
            if false {
                v21 := GenS3{F0: false, F1: false}
                _ = v21
            }
            v22 := GenS4{F0: (v7 - (v7 & int(v18))), F1: (v10 >> (((v10 / (((v10 * uint(v2))) | 1))) & 63)), F2: (v20 & 0)}
            _ = v20
            _ = v22
        }
        v2 += (^(v2))
        for fi23 := 0; fi23 < 5; fi23++ {
            for fi24 := 0; fi24 < 18; fi24++ {
                v2 = uint64(v10)
                _ = len(v17)
                var v25 float32 = 0.0
                var v26 int8 = (v9 << (((v10 - (v10 & 1))) & 63))
                if false {
                    v9 = (-(v26))
                }
                _ = fi24
                _ = v25
                _ = v26
            }
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
            fmt.Println(v18, v18)
            _ = fi23
        }
    }
    var v27 int = v7
    var v28 uint32 = 2
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    v7 -= v27
    var v29 bool = (!(-576 > 42))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
