package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 string; F1 int32; F2 uint16 }
type GenS5 struct { F0 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 int = -84
    var v3 int8 = -1
    v2 = v2
    fmt.Println(v3, v3)
    if (false && true) {
        if (0 <= 792) {
            if false {
                if true {
                    v3 += (v3 & (-(v3)))
                } else {
                    switch 178 {
                    case -254:
                        v4 := GenS1{F0: "a", F1: 42, F2: 0}
                        _ = v4
                    default:
                        if false {
                            var v5 int32 = 1
                            fmt.Println(v2, v5, v5)
                            _ = v5
                        }
                    }
                    var v6 string = "\n\t"
                    var v7 int8 = v3
                    _ = v6
                    _ = v7
                }
                var v8 uint32 = 42
                fmt.Println(v3, v3, v2)
                _ = v8
            }
            v9 := []int{(v2 / ((v2) | 1))}
            _ = v9
        }
        var v10 int64 = 176
        _ = v10
    }
    fmt.Println(v3, v2)
    fmt.Println(v2, v2)
    fmt.Println(v3, v2)
    fmt.Println(v2, v3)
    v11 := GenS2{F0: 1, F1: 696}
    switch "" {
    case "hello":
        if (true || (((0.0 + (((0.1 * (1e10 + (1.5 / float32(1e10)))) / ((float32(0.1) / 1e10) - 1e10)) * 1.5)) > float32((1.5 * (3.14159 / float64(1e10))))) || false)) {
            for fi12 := 0; fi12 < 5; fi12++ {
                v3 -= (v3 * (v3 | (v3 << ((0) & 63))))
                _ = fi12
            }
        } else {
            var v13 string = "foo bar"
            _ = v13
        }
    case "\n\t":
        var v14 int = (-(v2))
        _ = v14
    default:
        if false {
            if ((0 > 729) || true) {
                if false {
                    var v15 uint32 = 740
                    switch v2 {
                    case 128:
                        v16 := GenS1{F0: "", F1: v15, F2: 4294967295}
                        _ = v16
                    default:
                        if true {
                            v1[v2] = true
                            v15 = v15
                        }
                    }
                    if (2 < 1) {
                        v17 := GenS3{F0: 18446744073709551615, F1: 966}
                        var v18 int8 = 79
                        for fi19 := 0; fi19 < 15; fi19++ {
                            switch "\n\t" {
                            case "hello":
                                var v20 string = "hello"
                                _ = v20
                            case "a":
                                v21 := []int{0, v2, (v2 ^ v2)}
                                _ = v21
                            case "":
                                fmt.Println(v15, v18, v3)
                            default:
                                var v22 bool = false
                                _ = v22
                            }
                            v2 = -950
                            _ = fi19
                        }
                        _ = v17
                        _ = v18
                    } else {
                        for fi23 := 0; fi23 < 15; fi23++ {
                            if (false && true) {
                                v24 := GenS2{F0: 1, F1: 0}
                                _ = v24
                            }
                            _ = fi23
                        }
                        var v25 int8 = (v3 / ((int8(v15)) | 1))
                        var v26 int32 = 42
                        _ = v25
                        _ = v26
                    }
                    _ = v15
                } else {
                    var v27 float32 = 0.1
                    v28 := []int{}
                    _ = v27
                    _ = v28
                }
                v29 := []int{(^(v2)), (v2 / ((-9223372036854775808) | 1))}
                if (false && true) {
                    v2 += (v2 - (v2 >> ((34) & 63)))
                    if false {
                        v30 := GenS4{F0: "", F1: 1, F2: 42}
                        _ = v30
                    } else {
                        delete(v1, (v2 | (^(v2))))
                    }
                } else {
                    v31 := GenS5{F0: (^(v3))}
                    _ = v31
                }
                _ = v29
            }
        }
    }
    v32 := []int{32767, 128, 774, 2147483647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v32
}
