package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 uint; F2 uint64 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 bool; F1 string; F2 int64 }
type GenS6 struct { F0 int64; F1 int16 }

func main() {
    var v1 int32 = -2147483648
    var v2 int16 = -32768
    var v3 uint8 = 6
    var v4 float64 = 0.1
    if true {
        switch -1 {
        case 9223372036854775807:
            v2 -= v2
        case -2147483648:
            if true {
                if false {
                    if (!false) {
                        v5 := GenS1{F0: v1, F1: 1, F2: 18446744073709551615}
                        var v6 uint8 = v3
                        var v7 bool = (255 != 18446744073709551615)
                        _ = v5
                        _ = v6
                        _ = v7
                    }
                }
                v4 = (-2.25 - 3.14159)
            } else {
                for fi8 := 0; fi8 < 16; fi8++ {
                    if false {
                        v9 := []int{fi8, fi8}
                        var v10 bool = true
                        if len(v9) > 0 {
                            v9[0] = (^(fi8))
                        }
                        _ = v10
                        _ = v9
                    }
                    v3 -= 2
                    _ = fi8
                }
            }
        case 766:
            v4 -= (float64(1.5) * 0.1)
        default:
            v1 = v1
        }
        fmt.Println(v1, v4)
        v11 := []int{2147483647, -9223372036854775808, 42, 32768}
        _ = v11
    }
    var v12 uint32 = 2147483647
    var v13 bool = (false && (2 != 1))
    var v14 uint64 = 4294967295
    for fi15 := 0; fi15 < 3; fi15++ {
        fmt.Println(v12, v2, v1)
        _ = fi15
    }
    var v16 uint8 = (v3 - v3)
    var v17 float32 = 1.5
    if false {
        var v18 int32 = (-(v1))
        if v13 {
            var v19 uint32 = (v12 * v12)
            for fi20 := 0; fi20 < 14; fi20++ {
                v2 = (v2 << ((65535) & 63))
                fmt.Println(v2, fi20, v2)
                _ = fi20
            }
            _ = v19
        } else {
            v17 += 1.5
        }
        var v21 uint16 = 0
        _ = v18
        _ = v21
    } else {
        v22 := []int{2147483647, -98}
        var v23 uint32 = (v12 - v12)
        _ = v22
        _ = v23
    }
    v24 := GenS2{F0: 0}
    for fi25 := 0; fi25 < 6; fi25++ {
        v26 := GenS2{F0: 4294967295}
        if (-208 == 814) {
            var v27 string = "a"
            var v28 int16 = v2
            _ = v27
            _ = v28
        }
        if ((v17 + -2.25) != (((v17 - v17) * v17) - -2.25)) {
            if (!(((-2.25 * -2.25) < ((3.14159 * float32(v4)) + (v17 * 1e10))) || (v17 == (v17 - 3.14159)))) {
                if v13 {
                    for fi29 := 0; fi29 < 15; fi29++ {
                        if (v2 > -32768) {
                            v16 = (v3 | 1)
                        }
                        _ = fi29
                    }
                } else {
                    v17 -= v17
                    fmt.Println(v14, v16, v4)
                    for fi30 := 0; fi30 < 18; fi30++ {
                        v31 := []int{fi30}
                        var v32 uint64 = v14
                        v14 = 883
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                }
                var v33 uint = 467
                _ = v33
            }
            var v34 int8 = 1
            _ = v34
        } else {
            if v13 {
                if true {
                    fmt.Println(v16, v12)
                } else {
                    var v35 float32 = (1.5 / (1e10 / (v17 + 1.5)))
                    v36 := []int{2147483647, (-(fi25)), (fi25 / (((fi25 & int(v1))) | 1)), fi25}
                    if len(v36) >= 2 {
                        v36 = v36[1:len(v36)]
                    }
                    _ = v35
                    _ = v36
                }
                for fi37 := 0; fi37 < 4; fi37++ {
                    v38 := []int{-2147483648, fi25, (fi37 & 128), 128}
                    if true {
                        if v13 {
                            var v39 int = 9223372036854775807
                            var v40 uint64 = (v14 | 4294967295)
                            v41 := GenS3{F0: 42, F1: uint(v16), F2: (v14 & 913)}
                            _ = v39
                            _ = v40
                            _ = v41
                        } else {
                            var v42 uint16 = 23
                            _ = v42
                        }
                        v43 := GenS4{F0: -128, F1: 32767}
                        fmt.Println(v16, v3, v13)
                        _ = v43
                    }
                    _ = fi37
                    _ = v38
                }
                if v13 {
                    var v44 uint64 = 18446744073709551615
                    var v45 uint32 = (^(v12))
                    var v46 int64 = 42
                    _ = v44
                    _ = v45
                    _ = v46
                } else {
                    for fi47 := 0; fi47 < 9; fi47++ {
                        v4 = v4
                        if (v17 >= v17) {
                            var v48 int8 = 75
                            var v49 int16 = v2
                            _ = v48
                            _ = v49
                        } else {
                            v2 -= v2
                        }
                        _ = fi47
                    }
                }
            }
            v16 = (v3 * (v3 / (((v3 & v3)) | 1)))
        }
        _ = fi25
        _ = v26
    }
    v4 += (0.0 + (v4 + 1.5))
    for fi50 := 0; fi50 < 11; fi50++ {
        var v51 int32 = (v1 | (v1 ^ (v1 >> ((255) & 63))))
        _ = fi50
        _ = v51
    }
    if v13 {
        v52 := []int{135, int(v2)}
        v53 := GenS5{F0: (v17 > ((3.14159 + 1.5) / 1e10)), F1: "hello", F2: 63}
        for fi54 := 0; fi54 < 8; fi54++ {
            v1 = (v1 ^ (v1 + (-(v1))))
            v14 = 42
            for fi55 := 0; fi55 < 10; fi55++ {
                fmt.Println(v3, v3)
                var v56 float32 = v17
                var v57 bool = ((fi55 & fi55) == fi55)
                _ = len(v52)
                _ = fi55
                _ = v56
                _ = v57
            }
            _ = fi54
        }
        _ = v52
        _ = v53
    }
    switch "" {
    case "\n\t":
        v58 := []int{-9223372036854775808}
        _ = v58
    case "":
        v59 := []int{42, 131, 9223372036854775807, 9223372036854775807}
        _ = v59
    case "hello":
        var v60 int16 = (v2 - v2)
        _ = v60
    }
    v13 = (!(true || (2 < (v3 - (v16 >> ((uint(v3)) & 63))))))
    for fi61 := 0; fi61 < 10; fi61++ {
        v62 := []int{9223372036854775807, (^(fi61)), 42, fi61}
        fmt.Println(v2)
        _ = fi61
        _ = v62
    }
    if v13 {
        v63 := GenS6{F0: int64(v12), F1: v2}
        _ = v63
    } else {
        var v64 int64 = 1
        v16 = (v3 / (((^(v3))) | 1))
        var v65 uint = 1
        _ = v64
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
