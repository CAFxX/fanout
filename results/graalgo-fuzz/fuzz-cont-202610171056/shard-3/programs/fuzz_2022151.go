package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint8; F2 uint32 }
type GenS3 struct { F0 uint; F1 uint32; F2 int32 }
type GenS4 struct { F0 string; F1 int32; F2 uint64 }

func main() {
    v1 := []int{-491, -420, -1, 32767}
    var v2 int = -881
    var v3 uint32 = 95
    var v4 float64 = -2.25
    v4 -= (v4 - (1.5 * 0.1))
    var v5 int16 = -498
    if (!true) {
        v6 := GenS1{F0: (18446744073709551615 == 921)}
        _ = v6
    }
    switch v2 {
    case 1:
        v5 = (v5 / ((42) | 1))
    default:
        var v7 uint = 0
        _ = v7
    }
    v4 -= v4
    if false {
        v4 = float64(0.0)
    } else {
        v8 := GenS2{F0: "hello", F1: 255, F2: (v3 >> ((255) & 63))}
        for fi9 := 0; fi9 < 5; fi9++ {
            if true {
                fmt.Println(fi9, fi9)
            }
            if false {
                v2 = v2
                if false {
                    for fi10 := 0; fi10 < 8; fi10++ {
                        v2 += fi10
                        var v11 uint = 1
                        _ = fi10
                        _ = v11
                    }
                    v3 = 2147483648
                }
            } else {
                if true {
                    switch 0 {
                    case 9223372036854775807:
                        for fi12 := 0; fi12 < 20; fi12++ {
                            if (183 != 1) {
                                if false {
                                    switch 4294967295 {
                                    case -500:
                                        var v13 uint8 = 42
                                        v1 = append(v1, v2)
                                        var v14 int64 = 843
                                        _ = v13
                                        _ = v14
                                    case 9223372036854775807:
                                        var v15 int64 = -9223372036854775808
                                        _ = v15
                                    case 127:
                                        var v16 int = fi12
                                        _ = v16
                                    default:
                                        switch ("hello" + "") {
                                        case "":
                                            switch -30 {
                                            case 128:
                                                var v17 string = "foo bar"
                                                _ = v17
                                            case 9223372036854775807:
                                                v2 += fi12
                                            case 32767:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            default:
                                                if (!(42 != 1)) {
                                                    v18 := GenS3{F0: 42, F1: v3, F2: 0}
                                                    _ = v18
                                                } else {
                                                    v2 += 2147483647
                                                    v2 -= fi9
                                                    var v19 int16 = v5
                                                    _ = v19
                                                }
                                            }
                                        default:
                                            var v20 float32 = (0.0 * (0.1 * float32(v4)))
                                            _ = v20
                                        }
                                    }
                                    switch v2 {
                                    case -2147483648:
                                        var v21 float64 = v4
                                        _ = v21
                                    case 4294967295:
                                        v2 = 1
                                    }
                                    v3 = (^(v3))
                                }
                                var v22 uint32 = v3
                                v2 += (fi9 % ((v2) | 1))
                                _ = v22
                            }
                            var v23 int8 = 127
                            _ = fi12
                            _ = v23
                        }
                    case 116:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case 127:
                        if (!(-32768 < v5)) {
                            v5 = 1000
                        } else {
                            v2 = (fi9 | (fi9 | (fi9 % (((v2 % ((int(v5)) | 1))) | 1))))
                            v24 := GenS4{F0: "\n\t", F1: 0, F2: 2}
                            var v25 bool = false
                            _ = v24
                            _ = v25
                        }
                    }
                    var v26 uint16 = 1
                    var v27 float32 = -2.25
                    _ = v26
                    _ = v27
                }
                v3 = 0
                var v28 uint32 = v3
                _ = v28
            }
            _ = fi9
        }
        _ = v8
    }
    switch "" {
    case "hello":
        var v29 uint8 = 42
        _ = v29
    default:
        v30 := GenS1{F0: true}
        _ = v30
    }
    fmt.Println(v3, v5)
    _ = len(v1)
    v3 -= v3
    if len(v1) > 0 {
        v1[0] = (v2 + 0)
    }
    v4 += v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
