package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 int32; F2 uint16 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 int8; F1 string; F2 bool }
func genFn0() int8 {
    var v5 uint64 = 42
    var v6 int32 = -2147483648
    var v7 uint32 = 2147483648
    for fi8 := 0; fi8 < 20; fi8++ {
        v7 -= v7
        v9 := []int{(fi8 / (((fi8 | 214)) | 1)), (^(fi8))}
        v7 = v7
        _ = fi8
        _ = v9
    }
    _ = v5
    _ = v6
    _ = v7
    return -113
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[int]uint32)
    var v3 string = "\n\t"
    v4 := make(map[uint32]int)
    var v10 uint16 = 42
    v11 := GenS1{F0: 806}
    v12 := []int{0}
    var v13 uint8 = 175
    if (true || true) {
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 uint64 = 831
            _ = fi14
            _ = v15
        }
        v16 := GenS1{F0: (^(v1))}
        fmt.Println(v1, v13)
        _ = v16
    }
    var v17 float32 = 1e10
    v12 = append(v12, 22)
    var v18 int = 42
    for fi19 := 0; fi19 < 8; fi19++ {
        var v20 uint64 = uint64(v10)
        for fi21 := 0; fi21 < 11; fi21++ {
            v3 = v3
            if (!false) {
                var v22 int16 = 32767
                if false {
                    if false {
                        if false {
                            var v23 bool = ((uint64(v1) >= v20) || false)
                            v13 += v13
                            _ = v23
                        } else {
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                            var v24 float64 = 0.1
                            if (!false) {
                                if ((^(v20)) <= v20) {
                                    var v25 int8 = 1
                                    if (!false) {
                                        v13 = v13
                                        v25 -= -103
                                    } else {
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                        if ((1 <= v1) && true) {
                                            fmt.Println(v1, v22)
                                        }
                                        v1 = v1
                                        switch (fi21 + (^(fi21))) {
                                        case 42:
                                            v26 := GenS2{F0: 32768}
                                            _ = v26
                                        case 2147483647:
                                            v13 -= v13
                                        default:
                                            var v27 int16 = 32767
                                            _ = v27
                                        }
                                    }
                                    v18 += (v18 % ((32768) | 1))
                                    var v28 int32 = -2147483648
                                    switch v3 {
                                    case "hello":
                                        fmt.Println(v20, v3, v17)
                                    case "\n\t":
                                        v29 := GenS3{F0: "hello", F1: 1, F2: uint16(v13)}
                                        _ = v29
                                    case "a":
                                        v30 := GenS1{F0: (^(v1))}
                                        _ = v30
                                    default:
                                        if (42 > 960) {
                                            if true {
                                                v31 := GenS4{F0: 9223372036854775807, F1: 18446744073709551615}
                                                v3 = ((((v3 + v3) + (v3 + v3)) + "\n\t") + (v3 + ""))
                                                _ = v31
                                            }
                                            if ((("hello" + v3) >= v3) || false) {
                                                v32 := GenS4{F0: -9223372036854775808, F1: 18446744073709551615}
                                                var v33 float64 = (3.14159 + v24)
                                                v33 += v24
                                                _ = v32
                                                _ = v33
                                            }
                                        } else {
                                            var v34 int64 = 9223372036854775807
                                            _ = v34
                                        }
                                    }
                                    _ = v25
                                    _ = v28
                                } else {
                                    v17 = v17
                                    var v35 uint8 = 128
                                    _ = v35
                                }
                                v17 += float32((float64(float32(v24)) - v24))
                            } else {
                                switch 2147483647 {
                                case -1:
                                    if len(v12) > 0 {
                                        v12[0] = fi21
                                    }
                                default:
                                    v12 = append(v12, fi21)
                                }
                                delete(v2, fi21)
                                switch 926 {
                                case -540:
                                    var v36 int32 = 1
                                    _ = v36
                                case -2147483648:
                                    v17 += 0.1
                                case 233:
                                    var v37 bool = true
                                    _ = v37
                                }
                            }
                            _ = v24
                        }
                    } else {
                        fmt.Println(v22, fi21, v22)
                        v20 = v20
                    }
                } else {
                    v3 = "a"
                    var v38 uint32 = v1
                    v20 = (v20 & (^(v20)))
                    v18 -= fi21
                    _ = v38
                }
                _ = v22
            }
            switch 42 {
            case 963:
                fmt.Println(v1, v20)
            case 830:
                v39 := GenS5{F0: 0, F1: "a", F2: true}
                var v40 uint16 = uint16(v13)
                _ = v39
                _ = v40
            case -2147483648:
                v10 = uint16(v20)
            }
            _ = fi21
        }
        _ = fi19
        _ = v20
    }
    var v41 int8 = -84
    var v42 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
}
