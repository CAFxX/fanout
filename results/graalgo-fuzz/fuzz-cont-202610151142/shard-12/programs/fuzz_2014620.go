package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint }
type GenS2 struct { F0 int32; F1 int8; F2 uint16 }
func genFn0(p0 bool) int32 {
    fmt.Println(p0)
    var v6 uint = 2
    _ = p0
    _ = v6
    return -1
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    var v3 uint = 18446744073709551615
    var v4 uint = (v3 + 18446744073709551615)
    v5 := []uint32{165, 549, 2, 2147483648}
    var v7 uint16 = 2
    v2[42] = 0
    var v8 float32 = ((float32((float64(3.14159) / 1.0)) + -2.25) + 0.0)
    if false {
        v9 := v2[65535]
        if (!((!false) && true)) {
            v10 := GenS1{F0: uint32(v7), F1: v7, F2: v3}
            if ((true || (uint(v7) != v4)) || true) {
                v11 := GenS1{F0: v9, F1: 255, F2: 405}
                _ = v11
            } else {
                var v12 uint32 = (v9 << ((0) & 63))
                if false {
                    if false {
                        var v13 uint = uint(v7)
                        _ = v13
                    }
                    var v14 int8 = 127
                    _ = len(v5)
                    _ = v14
                }
                if false {
                    var v15 uint64 = 1
                    var v16 string = "a"
                    _ = v15
                    _ = v16
                } else {
                    v4 = 255
                    if false {
                        if false {
                            for fi17 := 0; fi17 < 18; fi17++ {
                                var v18 int64 = 9223372036854775807
                                if true {
                                    if false {
                                        switch (^(fi17)) {
                                        case 128:
                                            v12 -= (v12 / ((2147483647) | 1))
                                        case 32768:
                                            v19 := GenS1{F0: (v9 ^ (v9 / (((v12 >> ((1) & 63))) | 1))), F1: v7, F2: v4}
                                            _ = v19
                                        case 42:
                                            if (!false) {
                                                fmt.Println(v3, v8, v9)
                                                if (71 >= 128) {
                                                    v20 := v2[218]
                                                    if ((-1 >= -807) && false) {
                                                        delete(v2, 0)
                                                        if true {
                                                            if (true || (((!false) || (!((v7 + v7) > 65535))) || (!true))) {
                                                                delete(v2, 2147483648)
                                                                for fi21 := 0; fi21 < 8; fi21++ {
                                                                    var v22 int = fi21
                                                                    _ = fi21
                                                                    _ = v22
                                                                }
                                                            }
                                                        }
                                                        fmt.Println(fi17, v8, v7)
                                                    }
                                                    _ = v20
                                                }
                                            } else {
                                                v23 := GenS1{F0: (v9 + (v9 % (((v12 + (v12 % (((v9 % (((v9 | 65535)) | 1))) | 1)))) | 1))), F1: v7, F2: 2}
                                                var v24 int64 = int64(v7)
                                                _ = v23
                                                _ = v24
                                            }
                                        }
                                        v18 -= v18
                                    }
                                    v25 := GenS1{F0: v12, F1: 2, F2: v3}
                                    v5 = append(v5, v9)
                                    _ = v25
                                }
                                var v26 float32 = 0.1
                                _ = fi17
                                _ = v18
                                _ = v26
                            }
                        } else {
                            v12 -= v12
                            var v27 float64 = 0.1
                            v28 := GenS2{F0: 2147483647, F1: 19, F2: v7}
                            _ = v27
                            _ = v28
                        }
                    }
                    if len(v5) > 0 {
                        v5[0] = (^(v12))
                    }
                }
                _ = v12
            }
            delete(v1, 0)
            _ = v10
        }
        v9 = v9
        _ = v9
    }
    v29 := GenS2{F0: -2147483648, F1: -128, F2: v7}
    var v30 float64 = float64(v8)
    var v31 int64 = 686
    var v32 uint16 = v7
    var v33 uint = 4294967295
    var v34 int = 1
    var v35 int8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
