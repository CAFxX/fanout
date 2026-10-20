package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 int8; F1 bool; F2 uint64 }
type GenS6 struct { F0 int32; F1 string }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 0
    v3 := make(map[uint32]bool)
    v4 := make(map[int]bool)
    var v5 int32 = 42
    if (255 < 128) {
        if ((!(2 < 527)) && (-333 == -523)) {
            if false {
                if true {
                    switch "foo bar" {
                    case "":
                        v6 := v3[2147483648]
                        _ = v6
                    case "foo bar":
                        var v7 int32 = (v5 * 1)
                        _ = v7
                    case "a":
                        v2 = (v2 / ((v2) | 1))
                    default:
                        var v8 uint16 = (v2 << ((18446744073709551615) & 63))
                        _ = v8
                    }
                }
                v9 := GenS1{F0: 9223372036854775807, F1: 127, F2: v5}
                _ = v9
            }
            switch 1 {
            case 1:
                v10 := GenS2{F0: (v2 * v2), F1: 42}
                _ = v10
            default:
                switch -770 {
                case 2147483647:
                    switch 127 {
                    case 550:
                        v11 := GenS2{F0: v2, F1: -1}
                        _ = v11
                    default:
                        if ((!true) && true) {
                            if (-128 >= -128) {
                                v2 += v2
                                v2 = v2
                                v12 := []int{-1, 2147483647, 58}
                                _ = v12
                            }
                            if false {
                                v2 = v2
                                v13 := GenS2{F0: v2, F1: -43}
                                v5 -= v5
                                _ = v13
                            } else {
                                switch "" {
                                case "hello":
                                    if true {
                                        v2 = (v2 % ((335) | 1))
                                        if ((0.1 * 0.1) != 0.1) {
                                            var v14 float64 = (0.0 / float64((float32(float64(0.1)) / 1.0)))
                                            _ = v14
                                        }
                                        var v15 uint16 = 32767
                                        _ = v15
                                    }
                                case "foo bar":
                                    if (18446744073709551615 != 42) {
                                        v2 += v2
                                        v2 -= (v2 * (v2 | (^(v2))))
                                    }
                                case "a":
                                    v5 -= v5
                                default:
                                    var v16 uint8 = 42
                                    _ = v16
                                }
                                var v17 bool = true
                                _ = v17
                            }
                        } else {
                            v2 = v2
                            var v18 int64 = 1
                            _ = v18
                        }
                    }
                default:
                    v2 -= (v2 | (v2 / (((v2 & v2)) | 1)))
                }
            }
            var v19 float32 = 1e10
            _ = v19
        }
        var v20 uint = 255
        fmt.Println(v5, v20)
        _ = v20
    }
    if (!true) {
        var v21 int = 4294967295
        if true {
            v22 := GenS1{F0: 1, F1: 0, F2: v5}
            _ = v22
        }
        v23 := []int{-347, int(v2), (v21 >> ((255) & 63)), (v21 * 32767)}
        _ = v21
        _ = v23
    } else {
        v24 := []int{-9223372036854775808, -84}
        if (3.14159 < 0.0) {
            v25 := GenS3{F0: 276}
            _ = v25
        } else {
            v26 := GenS1{F0: 87, F1: 0, F2: (v5 % ((v5) | 1))}
            _ = v26
        }
        fmt.Println(v2, v5)
        _ = v24
    }
    v2 = v2
    switch "" {
    case "hello":
        if true {
            fmt.Println(v2, v5)
            if ((true && ((-2.25 / 1.0) != 0.0)) || false) {
                v27 := GenS4{F0: ("hello" + "foo bar"), F1: -32768}
                v28 := v1[-384]
                _ = v27
                _ = v28
            }
        }
    default:
        fmt.Println(v2)
    }
    if false {
        v29 := GenS3{F0: (v2 + v2)}
        _ = v29
    }
    for fi30 := 0; fi30 < 15; fi30++ {
        switch fi30 {
        case 966:
            for fi31 := 0; fi31 < 15; fi31++ {
                v2 = (v2 << ((230) & 63))
                v2 += 2
                _ = fi31
            }
        case 0:
            v32 := GenS5{F0: -128, F1: ((fi30 ^ -184) <= fi30), F2: 1}
            _ = v32
        case -784:
            v2 = v2
        }
        switch ("" + (("hello" + "") + "")) {
        case "a":
            for fi33 := 0; fi33 < 5; fi33++ {
                var v34 int8 = 100
                if false {
                    var v35 int8 = (v34 + v34)
                    _ = v35
                } else {
                    v3[920] = (0 >= -1)
                    v2 = (v2 | v2)
                }
                v36 := []int{}
                _ = fi33
                _ = v34
                _ = v36
            }
        }
        var v37 string = ""
        _ = fi30
        _ = v37
    }
    v38 := GenS2{F0: 458, F1: 0}
    if true {
        v39 := []int{-2147483648, -673, 664}
        v40 := GenS6{F0: v5, F1: "foo bar"}
        var v41 float32 = 1.5
        _ = v39
        _ = v40
        _ = v41
    } else {
        var v42 uint64 = 18446744073709551615
        _ = v42
    }
    fmt.Println(v5, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
