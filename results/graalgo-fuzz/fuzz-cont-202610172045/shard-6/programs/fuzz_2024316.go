package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int32 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 int; F1 uint }
func genFn0(p0 int16, p1 int16) int16 {
    for fi5 := 0; fi5 < 18; fi5++ {
        if true {
            switch -206 {
            case 2147483647:
                if false {
                    var v6 int32 = 42
                    var v7 int16 = p1
                    _ = v6
                    _ = v7
                } else {
                    if false {
                        v8 := GenS1{F0: -841, F1: -655, F2: -1}
                        for fi9 := 0; fi9 < 8; fi9++ {
                            var v10 int32 = 42
                            _ = fi9
                            _ = v10
                        }
                        v11 := GenS2{F0: p0, F1: 464}
                        _ = v11
                        _ = v8
                    }
                    fmt.Println(p0, fi5, fi5)
                }
            default:
                fmt.Println(p0, fi5, fi5)
            }
            if ((255 < 32767) || (true || false)) {
                var v12 int64 = 748
                _ = v12
            } else {
                p1 = (p1 | (p0 % (((p1 / ((p1) | 1))) | 1)))
                p1 = (p0 | (p0 & p1))
                var v13 int64 = 1
                _ = v13
            }
            var v14 int32 = -127
            _ = v14
        }
        for fi15 := 0; fi15 < 16; fi15++ {
            switch fi15 {
            case 1:
                p1 = (p0 ^ p0)
            case 597:
                p0 = 32767
            case 2147483647:
                if false {
                    fmt.Println(fi5, fi15, fi5)
                }
            }
            p0 = (p0 + p0)
            if true {
                v16 := GenS1{F0: 1, F1: 0, F2: -1}
                v17 := GenS1{F0: -1, F1: 862, F2: 1}
                p1 = (p0 % ((p1) | 1))
                _ = v16
                _ = v17
            }
            _ = fi15
        }
        _ = fi5
    }
    p0 -= (p0 ^ -1)
    fmt.Println(p0)
    _ = p0
    _ = p1
    return p1
}

func main() {
    var v1 int16 = 1
    var v2 uint16 = 74
    var v3 uint32 = 65535
    v4 := []int{-815, 32767, 32768, -120}
    v18 := GenS3{F0: 0, F1: 18446744073709551615}
    v2 = v2
    fmt.Println(v1)
    var v19 int32 = 42
    v2 = 2
    fmt.Println(v3, v1, v3)
    var v20 uint32 = 42
    fmt.Println(v2, v2, v1)
    v21 := GenS2{F0: v1, F1: v20}
    var v22 uint = uint(v2)
    v23 := GenS1{F0: 0, F1: -9223372036854775808, F2: (v19 * (v19 >> ((18446744073709551615) & 63)))}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
