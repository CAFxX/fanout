package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 uint64 }
func genFn0() int32 {
    var v7 float64 = 0.0
    v8 := []int{-946}
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint8 = 255
    v2 := []uint32{}
    var v3 int = -9223372036854775808
    var v4 float32 = -2.25
    v5 := make(map[int]int)
    v6 := make(map[int]int)
    var v9 int = (v3 ^ (v3 % (((v3 << ((42) & 63))) | 1)))
    switch v9 {
    case 0:
        if true {
            v1 = 123
        } else {
            if true {
                var v10 uint8 = v1
                for fi11 := 0; fi11 < 13; fi11++ {
                    for fi12 := 0; fi12 < 19; fi12++ {
                        delete(v5, 0)
                        _ = fi12
                    }
                    switch -818 {
                    case 9223372036854775807:
                        for fi13 := 0; fi13 < 15; fi13++ {
                            v10 -= (v1 % (((v1 * v1)) | 1))
                            switch 2147483647 {
                            case 486:
                                v3 = (fi11 - v9)
                            case -9223372036854775808:
                                fmt.Println(fi11, v3)
                            default:
                                if (false || false) {
                                    switch (^(v3)) {
                                    case 9223372036854775807:
                                        v14 := GenS1{F0: 1000, F1: (v10 << ((4294967295) & 63))}
                                        _ = v14
                                    }
                                    v4 = 1e10
                                    v4 = 1e10
                                } else {
                                    if true {
                                        v15 := GenS1{F0: -329, F1: v10}
                                        var v16 int = v3
                                        _ = v15
                                        _ = v16
                                    } else {
                                        if (-429 > 172) {
                                            if ((3.14159 * 1e10) <= 0.0) {
                                                v4 += 1e10
                                                delete(v5, 32768)
                                            } else {
                                                v17 := GenS1{F0: 1, F1: (v10 ^ (v10 * (^(v1))))}
                                                _ = v17
                                            }
                                            v3 = v9
                                        } else {
                                            v4 = 0.0
                                        }
                                        var v18 uint = 0
                                        switch fi13 {
                                        case -33:
                                            if false {
                                                v1 -= (v10 ^ 128)
                                            }
                                        default:
                                            v10 -= v1
                                        }
                                        v19 := GenS2{F0: "hello", F1: 18446744073709551615}
                                        _ = v18
                                        _ = v19
                                    }
                                    v9 = (v9 >> ((255) & 63))
                                }
                            }
                            if (1e10 <= 1.5) {
                                v9 += 9223372036854775807
                            } else {
                                v1 = v1
                            }
                            _ = fi13
                        }
                    case 350:
                        if true {
                            v1 = (v10 | v10)
                            v20 := GenS1{F0: 0, F1: 242}
                            _ = len(v6)
                            _ = v20
                        } else {
                            v4 = -2.25
                        }
                    }
                    _ = fi11
                }
                fmt.Println(v1, v10, v9)
                _ = v10
            }
            v6[(-(v3))] = -1
            v2 = append(v2, 4294967295)
        }
    case 128:
        if true {
            if len(v2) > 0 {
                v2[0] = 2147483648
            }
            fmt.Println(v3, v1, v9)
            var v21 float32 = (v4 + float32((-2.25 / 3.14159)))
            _ = v21
        }
    default:
        _ = len(v5)
    }
    if ((-2.25 / (v4 - 0.1)) == (v4 / v4)) {
        v4 -= 0.0
    } else {
        v22 := GenS3{F0: ""}
        if false {
            var v23 int16 = -344
            fmt.Println(v1, v1)
            v24 := GenS4{F0: 1, F1: 0}
            _ = v23
            _ = v24
        }
        _ = v22
    }
    v25 := GenS3{F0: (("\n\t" + "a") + "hello")}
    fmt.Println(v9, v1)
    _ = len(v5)
    v4 = 0.1
    v1 = v1
    var v26 uint32 = 0
    switch 9223372036854775807 {
    case -2147483648:
        fmt.Println(v26, v26, v4)
    case 127:
        switch v3 {
        case -156:
            v27 := GenS3{F0: "\n\t"}
            _ = v27
        }
    }
    var v28 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
