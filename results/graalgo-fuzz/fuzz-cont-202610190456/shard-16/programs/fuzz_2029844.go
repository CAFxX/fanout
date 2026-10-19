package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 int32; F2 int64 }
func genFn0(p0 uint64, p1 uint32, p2 uint) string {
    fmt.Println(p0, p0, p1)
    v5 := GenS1{F0: (p1 / (((p1 * p1)) | 1))}
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return ("a" + "")
}

func main() {
    v1 := []int{-992, 32768, 128, -187}
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    var v4 bool = (!(100 >= -128))
    if (false && true) {
        var v6 uint8 = 235
        switch 613 {
        case -738:
            v6 -= (v6 | v6)
        case -1:
            fmt.Println(v6, v6)
        default:
            var v7 uint = 299
            _ = v7
        }
        _ = v6
    }
    v8 := GenS2{F0: 4294967295}
    fmt.Println(v4)
    var v9 uint16 = 1
    for fi10 := 0; fi10 < 5; fi10++ {
        v3[(fi10 << ((18446744073709551615) & 63))] = 1
        _ = fi10
    }
    var v11 uint32 = 832
    fmt.Println(v9, v11, v9)
    v9 -= v9
    if (0 <= int8(v11)) {
        fmt.Println(v4)
        if (v4 || (!(42 <= 0))) {
            _ = len(v1)
            if (4 != v9) {
                v4 = (v9 > v9)
                v9 += (v9 % ((154) | 1))
            }
            if true {
                fmt.Println(v4, v4, v4)
            }
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi12 := 0; fi12 < 12; fi12++ {
                v9 = v9
                if false {
                    if len(v1) > 0 {
                        v1[0] = (fi12 & (fi12 / (((fi12 + fi12)) | 1)))
                    }
                    for fi13 := 0; fi13 < 10; fi13++ {
                        if ((^(v11)) < v11) {
                            var v14 int16 = 1000
                            v15 := GenS2{F0: 4294967295}
                            _ = v14
                            _ = v15
                        } else {
                            fmt.Println(v4, v11)
                        }
                        var v16 int = (^(fi12))
                        if (454 >= -942) {
                            var v17 int = v16
                            fmt.Println(v11, fi13)
                            _ = len(v2)
                            _ = v17
                        }
                        _ = fi13
                        _ = v16
                    }
                }
                _ = fi12
            }
            if false {
                v18 := GenS2{F0: v11}
                fmt.Println(v4)
                _ = v18
            }
        }
    } else {
        var v19 int32 = 720
        _ = v19
    }
    fmt.Println(v11, v4, v9)
    v9 -= 0
    v20 := GenS3{F0: "\n\t", F1: 42, F2: 9223372036854775807}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
