package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 int }

func main() {
    var v1 int16 = -32768
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]string)
    var v4 float32 = -2.25
    var v5 float32 = (v4 - v4)
    v6 := make(map[uint32]uint32)
    v5 += v5
    v1 = 32767
    fmt.Println(v5, v5, v1)
    v7 := GenS1{F0: 163, F1: 1, F2: 320}
    fmt.Println(v4)
    v1 = v1
    switch -162 {
    case 399:
        for fi8 := 0; fi8 < 7; fi8++ {
            v4 -= v5
            _ = len(v3)
            _ = fi8
        }
    }
    v9 := GenS1{F0: 65535, F1: 181, F2: 423}
    v5 -= 1.5
    if ((2 > 2147483648) && true) {
        if (1 != -128) {
            var v10 int64 = 42
            _ = v10
        } else {
            var v11 int64 = 42
            var v12 float32 = v5
            v13 := []int{-293, 9223372036854775807}
            _ = v11
            _ = v12
            _ = v13
        }
        v14 := GenS1{F0: 2147483648, F1: 1, F2: 1}
        _ = v14
    } else {
        switch "foo bar" {
        case "a":
            var v15 float64 = float64((0.0 / float32(-2.25)))
            _ = v15
        case "hello":
            switch 4294967295 {
            case 32768:
                v5 = ((3.14159 + v4) - float32(1.5))
            case -9223372036854775808:
                v5 -= (v4 * -2.25)
            case -741:
                switch 32767 {
                case 0:
                    v1 -= (v1 >> ((42) & 63))
                case -9223372036854775808:
                    for fi16 := 0; fi16 < 11; fi16++ {
                        v1 = (v1 << ((65535) & 63))
                        for fi17 := 0; fi17 < 20; fi17++ {
                            v1 = (v1 - (v1 - -1))
                            if true {
                                var v18 float64 = 3.14159
                                _ = v18
                            } else {
                                var v19 int32 = 1
                                delete(v6, 42)
                                _ = v19
                            }
                            _ = fi17
                        }
                        _ = fi16
                    }
                case 42:
                    delete(v3, 0)
                }
            default:
                if (463 <= 0) {
                    fmt.Println(v4)
                    var v20 int16 = -1
                    for fi21 := 0; fi21 < 20; fi21++ {
                        var v22 int64 = 9223372036854775807
                        if (!false) {
                            switch fi21 {
                            case 32768:
                                var v23 bool = false
                                _ = v23
                            default:
                                if false {
                                    v24 := []int{fi21, 364, (fi21 >> ((65535) & 63))}
                                    _ = v24
                                }
                            }
                            var v25 int64 = -9223372036854775808
                            var v26 int16 = (v1 + (v1 >> ((18446744073709551615) & 63)))
                            _ = v25
                            _ = v26
                        } else {
                            v27 := GenS2{F0: 42}
                            _ = v27
                        }
                        v28 := GenS3{F0: 299, F1: fi21}
                        _ = fi21
                        _ = v22
                        _ = v28
                    }
                    _ = v20
                } else {
                    for fi29 := 0; fi29 < 6; fi29++ {
                        fmt.Println(fi29)
                        _ = fi29
                    }
                }
            }
        case "":
            _ = len(v2)
        default:
            v5 -= float32(0.1)
        }
        var v30 int8 = 60
        var v31 uint32 = 558
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 9; fi32++ {
        v4 = 1.5
        _ = fi32
    }
    v5 = (((v5 / v5) + v4) / v5)
    var v33 int32 = -1
    if (!(173 < 87)) {
        fmt.Println(v5, v1, v5)
        for fi34 := 0; fi34 < 9; fi34++ {
            v33 -= (v33 - v33)
            _ = fi34
        }
    } else {
        v35 := GenS2{F0: 2}
        v4 = 1e10
        var v36 int32 = 0
        _ = v35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
