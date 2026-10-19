package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 uint; F1 int; F2 int }
type GenS6 struct { F0 int; F1 int; F2 int8 }

func main() {
    var v1 uint = 65535
    v2 := make(map[int]int)
    var v3 int32 = -1
    v4 := []uint32{65535, 675}
    v5 := []uint32{2147483647, 2, 4294967295, 2}
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 bool = (true || ((false && true) && (false || ((fi6 | (fi6 ^ -1)) != 4294967295))))
        _ = fi6
        _ = v7
    }
    var v8 bool = true
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 string = "foo bar"
        v10 = (v10 + v10)
        _ = fi9
        _ = v10
    }
    if v8 {
        if false {
            if (469 <= 74) {
                var v11 int8 = 76
                switch 32767 {
                case 127:
                    switch "" {
                    case "a":
                        if (!(!true)) {
                            v8 = false
                            _ = len(v2)
                            if v8 {
                                v12 := GenS1{F0: 1000, F1: (^(v1)), F2: 1}
                                v13 := v2[405]
                                var v14 string = "a"
                                _ = v12
                                _ = v13
                                _ = v14
                            } else {
                                var v15 uint64 = 18446744073709551615
                                if (42 <= 1) {
                                    v8 = true
                                    delete(v2, 0)
                                    fmt.Println(v1)
                                } else {
                                    var v16 int16 = -882
                                    _ = v16
                                }
                                var v17 float32 = 3.14159
                                _ = v15
                                _ = v17
                            }
                        }
                    case "\n\t":
                        v1 = (v1 / (((v1 | (v1 % (((v1 % (((v1 - 843)) | 1))) | 1)))) | 1))
                    default:
                        v11 = v11
                    }
                case 32767:
                    var v18 float64 = float64(float32(((1.5 - 1.5) + 0.0)))
                    _ = v18
                case -2147483648:
                    var v19 float64 = 1.5
                    _ = v19
                default:
                    v20 := GenS1{F0: 32767, F1: 4294967295, F2: 0}
                    _ = v20
                }
                _ = v11
            }
            var v21 int32 = 2147483647
            _ = v21
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v22 := GenS1{F0: int16(v1), F1: 255, F2: 138}
    v2[127] = 918
    var v23 int16 = 0
    fmt.Println(v3, v23)
    v24 := GenS2{F0: 2147483648}
    var v25 float32 = 0.1
    for fi26 := 0; fi26 < 6; fi26++ {
        v27 := GenS3{F0: 42, F1: 127}
        if false {
            if len(v4) > 0 {
                v4[0] = 42
            }
            var v28 int16 = v23
            _ = v28
        }
        var v29 int64 = 1
        _ = fi26
        _ = v27
        _ = v29
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        v3 += v3
        _ = fi30
    }
    if v8 {
        if v8 {
            for fi31 := 0; fi31 < 11; fi31++ {
                for fi32 := 0; fi32 < 7; fi32++ {
                    switch fi32 {
                    case 32768:
                        v1 -= (v1 ^ (v1 - (v1 + v1)))
                    default:
                        _ = len(v2)
                        switch fi31 {
                        case 830:
                            var v33 uint16 = 32767
                            _ = v33
                        }
                        var v34 int16 = v23
                        v35 := GenS4{F0: (701 <= 9223372036854775807), F1: v1}
                        v36 := GenS5{F0: (v1 % (((v1 | (^(v1)))) | 1)), F1: (^(fi31)), F2: (^(fi31))}
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                    v25 -= float32((0.0 + (0.1 * 1e10)))
                    _ = fi32
                }
                if false {
                    var v37 uint16 = 255
                    var v38 float32 = v25
                    _ = v37
                    _ = v38
                }
                _ = fi31
            }
        } else {
            var v39 uint64 = 4294967295
            var v40 string = "foo bar"
            _ = v39
            _ = v40
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi41 := 0; fi41 < 13; fi41++ {
            v3 += (v3 + (-(v3)))
            if v8 {
                v1 = v1
            } else {
                for fi42 := 0; fi42 < 6; fi42++ {
                    v3 = (v3 - v3)
                    _ = len(v2)
                    v43 := GenS6{F0: fi42, F1: fi41, F2: 100}
                    _ = fi42
                    _ = v43
                }
                fmt.Println(v1, v1, v1)
            }
            _ = fi41
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
