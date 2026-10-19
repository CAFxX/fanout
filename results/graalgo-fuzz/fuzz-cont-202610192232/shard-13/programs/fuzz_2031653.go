package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 uint16 }

func main() {
    var v1 int16 = 215
    v2 := make(map[uint32]bool)
    var v3 uint16 = 0
    v4 := []int{32767, 32767}
    switch 122 {
    case -2147483648:
        if false {
            v5 := GenS1{F0: -2147483648, F1: v1}
            v6 := GenS2{F0: 0, F1: 128}
            if ((131 >= 255) || false) {
                if false {
                    for fi7 := 0; fi7 < 16; fi7++ {
                        if (!false) {
                            var v8 float32 = 3.14159
                            for fi9 := 0; fi9 < 2; fi9++ {
                                v1 -= v1
                                var v10 uint8 = 141
                                v3 += (v3 & 43)
                                _ = fi9
                                _ = v10
                            }
                            _ = v8
                        } else {
                            var v11 uint8 = 255
                            _ = v11
                        }
                        v1 -= (v1 << ((18446744073709551615) & 63))
                        if len(v4) > 0 {
                            v4[0] = (fi7 & fi7)
                        }
                        _ = fi7
                    }
                    _ = len(v4)
                    var v12 float32 = 1e10
                    _ = v12
                }
                if ((false && false) && (0 < 351)) {
                    v1 = -32768
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v13 := v2[0]
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 2; fi14++ {
                        v1 = -660
                        _ = fi14
                    }
                }
                v15 := v2[4294967295]
                _ = v15
            } else {
                if (!true) {
                    v16 := GenS2{F0: 854, F1: 4294967295}
                    v1 = (^(v1))
                    _ = v16
                }
                switch 42 {
                case 832:
                    if (70 >= 54) {
                        if (-1 == 0) {
                            for fi17 := 0; fi17 < 10; fi17++ {
                                v18 := GenS3{F0: 1}
                                v1 += 42
                                if true {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v1 -= -32768
                                    v1 += 42
                                } else {
                                    var v19 int8 = -1
                                    var v20 uint = 2
                                    var v21 uint8 = 42
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                }
                                _ = fi17
                                _ = v18
                            }
                            for fi22 := 0; fi22 < 4; fi22++ {
                                v3 = (v3 - v3)
                                var v23 uint32 = 2
                                switch "hello" {
                                case "hello":
                                    v24 := GenS4{F0: -37, F1: v3}
                                    _ = v24
                                default:
                                    fmt.Println(v1, v3, v3)
                                }
                                _ = fi22
                                _ = v23
                            }
                            v1 = v1
                        } else {
                            v3 = 255
                        }
                    }
                case 4294967295:
                    var v25 uint64 = 42
                    _ = v25
                default:
                    _ = len(v4)
                }
                if false {
                    var v26 float32 = float32((0.0 * -2.25))
                    switch 32768 {
                    case 2147483647:
                        for fi27 := 0; fi27 < 19; fi27++ {
                            v28 := GenS3{F0: 1}
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v29 int64 = 9223372036854775807
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                    case 32767:
                        v1 += (v1 + 42)
                    case 815:
                        v3 -= v3
                    default:
                        var v30 uint = 0
                        _ = v30
                    }
                    _ = v26
                } else {
                    _ = len(v2)
                    v31 := GenS1{F0: 2147483647, F1: (v1 << ((42) & 63))}
                    _ = v31
                }
            }
            _ = v5
            _ = v6
        } else {
            fmt.Println(v3)
        }
    case 127:
        fmt.Println(v3, v3)
    }
    var v32 uint8 = 42
    fmt.Println(v32, v32, v32)
    for fi33 := 0; fi33 < 16; fi33++ {
        v3 -= (v3 | v3)
        _ = fi33
    }
    var v34 int8 = -77
    v32 += v32
    if (42 != -9223372036854775808) {
        var v35 int64 = 42
        _ = v35
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v36 bool = (((false && (!true)) && false) || (1.5 != ((0.0 + (float64(3.14159) * float64(0.1))) * 0.0)))
    switch 128 {
    case 42:
        var v37 float64 = 1e10
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v36
    _ = v4
}
