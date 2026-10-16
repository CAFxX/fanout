package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int8; F1 int32; F2 string }
type GenS4 struct { F0 int64; F1 uint; F2 int32 }
func genFn0(p0 int, p1 int16) int {
    fmt.Println(p0, p0)
    v5 := GenS1{F0: 42, F1: 1000}
    p1 = 663
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := []int{(p0 / (((fi6 - fi6)) | 1)), (p0 ^ -739), (p0 | (fi6 << ((530) & 63)))}
        _ = fi6
        _ = v7
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = v5
    return p0
}

func main() {
    var v1 uint64 = 2
    v2 := []uint32{2147483647}
    var v3 int16 = 32767
    v4 := make(map[int]int)
    v8 := GenS2{F0: 0, F1: 895}
    var v9 float32 = 0.0
    var v10 uint32 = uint32(v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v10 -= 2147483647
    v11 := GenS3{F0: 0, F1: 905, F2: "foo bar"}
    if true {
        v1 -= v1
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        v1 = 0
        switch genFn0(fi12, -533) {
        case 0:
            for fi13 := 0; fi13 < 4; fi13++ {
                if (118 >= 100) {
                    if (42 >= 2) {
                        v9 -= v9
                        v14 := v4[fi13]
                        var v15 float32 = 1.5
                        var v16 int8 = -50
                        _ = v14
                        _ = v15
                        _ = v16
                    } else {
                        var v17 int64 = 42
                        if (uint16(v1) >= 1) {
                            fmt.Println(fi13, v9, v9)
                            v18 := v4[(fi13 * genFn0((fi13 | (fi12 - fi12)), v3))]
                            _ = v18
                        } else {
                            var v19 uint8 = 42
                            var v20 bool = false
                            var v21 uint8 = v19
                            v1 += v1
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                        v22 := GenS4{F0: v17, F1: 18446744073709551615, F2: 506}
                        _ = v17
                        _ = v22
                    }
                    v23 := GenS1{F0: (v3 % ((0) | 1)), F1: v3}
                    _ = v23
                }
                var v24 float32 = float32(3.14159)
                v1 -= 276
                _ = fi13
                _ = v24
            }
        case -187:
            for fi25 := 0; fi25 < 20; fi25++ {
                v26 := GenS4{F0: 0, F1: 82, F2: 42}
                switch genFn0(fi12, v3) {
                case -123:
                    if true {
                        fmt.Println(v9)
                        switch "a" {
                        case "a":
                            var v27 int = fi12
                            _ = v27
                        case "\n\t":
                            v3 -= v3
                        case "foo bar":
                            v9 = v9
                        }
                        if true {
                            if true {
                                v10 += (^(v10))
                            } else {
                                v28 := GenS3{F0: 28, F1: 2147483647, F2: "hello"}
                                _ = v28
                            }
                            v9 = v9
                        } else {
                            fmt.Println(fi25, v9, fi12)
                            var v29 uint16 = 2
                            v1 = (v1 - uint64(v10))
                            if true {
                                if (0 != v29) {
                                    switch (^(fi12)) {
                                    case 4294967295:
                                        var v30 int16 = v3
                                        _ = v30
                                    default:
                                        switch (fi12 ^ fi12) {
                                        case 2147483647:
                                            if false {
                                                v3 = v3
                                            }
                                        default:
                                            var v31 uint8 = uint8(v1)
                                            _ = v31
                                        }
                                    }
                                    v32 := GenS2{F0: 128, F1: uint(v10)}
                                    _ = v32
                                }
                                var v33 int8 = -1
                                v34 := GenS2{F0: 255, F1: 1}
                                _ = v33
                                _ = v34
                            } else {
                                v1 = v1
                                delete(v4, (fi12 - 0))
                            }
                            _ = v29
                        }
                    } else {
                        var v35 int64 = -412
                        _ = v35
                    }
                case -2147483648:
                    v2 = append(v2, (v10 & (v10 & v10)))
                case 9223372036854775807:
                    switch -160 {
                    case 0:
                        var v36 int8 = -128
                        _ = v36
                    }
                }
                _ = fi25
                _ = v26
            }
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
