package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 uint64; F1 int16 }

func main() {
    var v1 int16 = -51
    var v2 uint64 = 1
    var v3 float64 = 0.1
    fmt.Println(v3, v2, v3)
    switch 32768 {
    case 13:
        if (-1 == 853) {
            v4 := GenS1{F0: 42, F1: 884, F2: 0}
            _ = v4
        }
    case -2147483648:
        if (!(44 > 32767)) {
            v3 += 3.14159
            var v5 int16 = v1
            _ = v5
        } else {
            for fi6 := 0; fi6 < 13; fi6++ {
                v7 := []int{}
                _ = fi6
                _ = v7
            }
            if ((v2 + (^(v2))) <= 814) {
                if (-515 > 217) {
                    var v8 uint64 = 271
                    _ = v8
                } else {
                    v1 = v1
                    v9 := GenS1{F0: 255, F1: 2, F2: 0}
                    _ = v9
                }
            } else {
                v1 = v1
                switch 127 {
                case 248:
                    if true {
                        var v10 float32 = -2.25
                        for fi11 := 0; fi11 < 17; fi11++ {
                            fmt.Println(fi11)
                            v12 := []int{(fi11 + (fi11 & 2147483647)), (fi11 % ((fi11) | 1)), (^(fi11)), (fi11 ^ (fi11 | int(v1)))}
                            _ = fi11
                            _ = v12
                        }
                        _ = v10
                    } else {
                        v3 = 1.5
                        v13 := []int{9223372036854775807}
                        _ = v13
                    }
                case -1:
                    for fi14 := 0; fi14 < 6; fi14++ {
                        var v15 float64 = float64(float32(v3))
                        fmt.Println(v1, v15)
                        _ = fi14
                        _ = v15
                    }
                case 128:
                    var v16 int16 = v1
                    _ = v16
                default:
                    if false {
                        var v17 int64 = -435
                        switch -396 {
                        case 32767:
                            if false {
                                v17 = (v17 % (((v17 - 0)) | 1))
                            } else {
                                switch 9223372036854775807 {
                                case 1:
                                    v18 := GenS2{F0: 32767, F1: -1}
                                    _ = v18
                                case -9223372036854775808:
                                    v19 := []int{999, 128, 0, 2147483647}
                                    _ = v19
                                default:
                                    var v20 int16 = (v1 % ((-198) | 1))
                                    _ = v20
                                }
                                if ((32767 != 91) && false) {
                                    v1 += 1
                                    var v21 uint = 0
                                    if (!(!(4 > -573))) {
                                        var v22 uint32 = 0
                                        v17 = 42
                                        v1 = (v1 - (v1 ^ v1))
                                        _ = v22
                                    }
                                    _ = v21
                                } else {
                                    for fi23 := 0; fi23 < 7; fi23++ {
                                        for fi24 := 0; fi24 < 10; fi24++ {
                                            v25 := []int{fi24, fi24}
                                            v26 := GenS3{F0: 4294967295, F1: 1}
                                            _ = fi24
                                            _ = v25
                                            _ = v26
                                        }
                                        for fi27 := 0; fi27 < 1; fi27++ {
                                            v2 -= v2
                                            _ = fi27
                                        }
                                        var v28 uint16 = 506
                                        _ = fi23
                                        _ = v28
                                    }
                                    for fi29 := 0; fi29 < 2; fi29++ {
                                        for fi30 := 0; fi30 < 5; fi30++ {
                                            v31 := GenS3{F0: (^(v2)), F1: 1}
                                            _ = fi30
                                            _ = v31
                                        }
                                        v3 = v3
                                        _ = fi29
                                    }
                                    v3 = v3
                                }
                                v1 -= v1
                            }
                        case 353:
                            var v32 int16 = 0
                            _ = v32
                        }
                        v1 += (v1 % (((v1 | (v1 ^ (v1 & v1)))) | 1))
                        _ = v17
                    }
                }
            }
            fmt.Println(v3, v3, v2)
        }
    default:
        var v33 int16 = v1
        _ = v33
    }
    switch 0 {
    case -283:
        fmt.Println(v2, v3, v1)
    default:
        v34 := []int{-740, 128, 2147483647}
        _ = v34
    }
    v3 += (0.1 * v3)
    v35 := []int{42}
    switch 861 {
    case 4294967295:
        var v36 int64 = 9223372036854775807
        _ = v36
    }
    v2 -= 1
    fmt.Println(v3, v3)
    fmt.Println(v2)
    fmt.Println(v2, v1, v1)
    for fi37 := 0; fi37 < 12; fi37++ {
        var v38 int8 = 41
        var v39 int = fi37
        _ = fi37
        _ = v38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
}
