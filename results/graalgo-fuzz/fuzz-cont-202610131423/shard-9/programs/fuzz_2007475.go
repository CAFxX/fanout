package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 bool }
type GenS2 struct { F0 int32; F1 int8; F2 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 uint8; F1 uint; F2 uint16 }
func genFn0(p0 int, p1 uint16) int {
    fmt.Println(p1, p1)
    var v6 int8 = 1
    v6 = (v6 << ((65535) & 63))
    var v7 uint = 2
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -2147483648
}

func main() {
    v1 := []int{42, -170, 32768, 536}
    var v2 int16 = 227
    v3 := []int{-13, 0}
    var v4 bool = (127 > -106)
    v5 := make(map[int]bool)
    var v8 int32 = 651
    var v9 int = -2147483648
    var v10 uint16 = 194
    var v11 float32 = (1e10 - 0.1)
    v2 = -32768
    v2 = (v2 / ((v2) | 1))
    for fi12 := 0; fi12 < 20; fi12++ {
        for fi13 := 0; fi13 < 7; fi13++ {
            if ((!false) || v4) {
                if (42 <= -202) {
                    var v14 uint32 = 2
                    var v15 float64 = (0.1 - -2.25)
                    _ = v14
                    _ = v15
                }
                if v4 {
                    if ((v10 >> ((1) & 63)) == 317) {
                        var v16 uint16 = 0
                        _ = v16
                    } else {
                        v5[-680] = (-2.25 == 1e10)
                        var v17 int8 = -1
                        _ = v17
                    }
                    if len(v3) > 0 {
                        v3[0] = -9223372036854775808
                    }
                    if (0 < 0) {
                        v10 = 2
                        switch (fi13 * genFn0((v9 / (((fi13 & fi13)) | 1)), (v10 - (v10 | (v10 - (v10 - 42)))))) {
                        case 42:
                            v9 += 0
                        default:
                            switch "\n\t" {
                            case "\n\t":
                                fmt.Println(fi12, fi13)
                            default:
                                var v18 string = "\n\t"
                                if v4 {
                                    v18 = "foo bar"
                                    var v19 uint64 = 2
                                    if ((v10 / (((v10 ^ v10)) | 1)) >= v10) {
                                        v19 += v19
                                        var v20 uint16 = v10
                                        var v21 string = v18
                                        _ = v20
                                        _ = v21
                                    } else {
                                        v10 = v10
                                        var v22 uint = 1
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v22
                                    }
                                    _ = v19
                                }
                                v10 = (v10 + 255)
                                _ = len(v5)
                                v23 := GenS1{F0: (v10 | v10), F1: 9223372036854775807, F2: (!((fi12 + 4294967295) < (v9 & v9)))}
                                v2 = (v2 - (v2 ^ -366))
                                v24 := GenS2{F0: v8, F1: int8(fi12), F2: -32768}
                                _ = v18
                                _ = v23
                                _ = v24
                            }
                        }
                        v25 := GenS3{F0: -1}
                        _ = v25
                    }
                }
            } else {
                v26 := GenS2{F0: (v8 % ((42) | 1)), F1: -128, F2: -32768}
                _ = v26
            }
            _ = fi13
        }
        fmt.Println(v11, v9, v9)
        for fi27 := 0; fi27 < 11; fi27++ {
            if v4 {
                v28 := GenS4{F0: 791, F1: (fi12 % (((fi12 / (((fi27 ^ (fi12 << ((42) & 63)))) | 1))) | 1))}
                fmt.Println(v10)
                v29 := GenS1{F0: 32767, F1: 23, F2: false}
                _ = v28
                _ = v29
            }
            v30 := GenS5{F0: 255, F1: 306, F2: (v10 << ((123) & 63))}
            _ = fi27
            _ = v30
        }
        _ = fi12
    }
    var v31 int64 = -9223372036854775808
    for fi32 := 0; fi32 < 5; fi32++ {
        fmt.Println(v11, v2, fi32)
        v1 = append(v1, int(v2))
        _ = fi32
    }
    if v4 {
        for fi33 := 0; fi33 < 2; fi33++ {
            var v34 uint16 = 32767
            _ = fi33
            _ = v34
        }
        for fi35 := 0; fi35 < 10; fi35++ {
            for fi36 := 0; fi36 < 2; fi36++ {
                v2 += (v2 ^ v2)
                _ = fi36
            }
            var v37 int16 = -1
            var v38 int = genFn0(127, (v10 / (((v10 - 2)) | 1)))
            _ = fi35
            _ = v37
            _ = v38
        }
    } else {
        v11 += float32(0.0)
        v2 = 32767
        v3 = append(v3, 2147483647)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v39 := GenS4{F0: (v8 >> ((0) & 63)), F1: (v9 * v9)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v39
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
