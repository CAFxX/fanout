package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int32 }
type GenS2 struct { F0 int64; F1 int64; F2 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 int16; F1 uint; F2 uint16 }
type GenS6 struct { F0 bool; F1 uint8 }
type GenS7 struct { F0 int16 }

func main() {
    v1 := []int{-2147483648, -2147483648}
    v2 := []int{}
    var v3 float32 = (0.1 * -2.25)
    var v4 uint64 = 345
    v5 := GenS1{F0: 0, F1: 56, F2: 2147483647}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v6 := GenS2{F0: 9223372036854775807, F1: 9223372036854775807, F2: 442}
    var v7 uint64 = v4
    for fi8 := 0; fi8 < 2; fi8++ {
        fmt.Println(v4, v7, v4)
        v3 -= v3
        _ = fi8
    }
    switch 1 {
    case 127:
        v9 := GenS1{F0: 2, F1: 148, F2: 1}
        _ = v9
    case 32767:
        v3 = float32(((1.5 - (0.0 / 0.1)) / 1.0))
    case 32768:
        if true {
            switch ("" + "") {
            case "":
                if (false || false) {
                    for fi10 := 0; fi10 < 4; fi10++ {
                        v11 := GenS1{F0: 720, F1: 42, F2: 82}
                        _ = fi10
                        _ = v11
                    }
                    v7 = v7
                    v12 := GenS2{F0: 931, F1: 42, F2: 907}
                    _ = v12
                }
            case "\n\t":
                if false {
                    for fi13 := 0; fi13 < 9; fi13++ {
                        v7 = v4
                        v4 -= (^(v7))
                        _ = fi13
                    }
                } else {
                    var v14 float64 = 3.14159
                    fmt.Println(v4, v3, v14)
                    _ = v14
                }
            case "foo bar":
                v15 := GenS3{F0: 65535}
                _ = v15
            default:
                v1 = append(v1, 32767)
            }
        } else {
            v4 = (v4 + (v4 + 18446744073709551615))
        }
    }
    var v16 int = 32768
    for fi17 := 0; fi17 < 15; fi17++ {
        for fi18 := 0; fi18 < 2; fi18++ {
            if (127 <= 126) {
                v19 := GenS3{F0: 65535}
                if len(v1) > 0 {
                    v1[0] = -909
                }
                v2 = append(v2, (v16 / ((fi18) | 1)))
                var v20 uint32 = 1
                _ = v19
                _ = v20
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                fmt.Println(v3)
            }
            _ = fi18
        }
        v21 := GenS1{F0: 2, F1: 255, F2: 1}
        _ = fi17
        _ = v21
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        v23 := GenS1{F0: 208, F1: 128, F2: int32(v16)}
        _ = fi22
        _ = v23
    }
    var v24 uint32 = 2
    v25 := GenS2{F0: 265, F1: -777, F2: 2}
    for fi26 := 0; fi26 < 2; fi26++ {
        v7 += (v4 - v7)
        _ = fi26
    }
    v27 := GenS2{F0: 9223372036854775807, F1: 1, F2: 18446744073709551615}
    if (0.0 < 1e10) {
        for fi28 := 0; fi28 < 6; fi28++ {
            v24 += v24
            v24 = v24
            _ = fi28
        }
    } else {
        if ((int16(v16) >= 32767) || false) {
            fmt.Println(v4)
        } else {
            v29 := GenS4{F0: false, F1: 387}
            v7 -= v7
            _ = v29
        }
        for fi30 := 0; fi30 < 10; fi30++ {
            var v31 float32 = ((v3 * (1.5 + ((v3 + 0.1) - (v3 - (v3 * v3))))) / float32((((1e10 * 3.14159) - (0.0 * ((1e10 + float64((v3 + (v3 * 1.5)))) + 1.5))) - (0.1 * 1.5))))
            v32 := GenS5{F0: -545, F1: 0, F2: 42}
            var v33 uint8 = 1
            _ = fi30
            _ = v31
            _ = v32
            _ = v33
        }
    }
    switch 42 {
    case 32768:
        switch v16 {
        case -506:
            v2 = append(v2, (v16 << ((4294967295) & 63)))
        }
    case 9223372036854775807:
        if false {
            for fi34 := 0; fi34 < 12; fi34++ {
                v24 -= (v24 / ((42) | 1))
                v35 := GenS6{F0: ((-2.25 - -2.25) != ((v3 + ((v3 + ((0.0 + (v3 + v3)) / (v3 + float32((0.1 / ((3.14159 - 1e10) * ((float64(v3) + 0.1) - 1e10))))))) - ((3.14159 + (0.1 - (v3 + v3))) * v3))) / 1.0)), F1: 0}
                var v36 uint8 = 36
                _ = fi34
                _ = v35
                _ = v36
            }
        } else {
            v37 := GenS7{F0: -211}
            fmt.Println(v4)
            v24 = v24
            _ = v37
        }
    case -1:
        v16 += v16
    }
    fmt.Println(v24, v24, v24)
    v38 := GenS2{F0: 0, F1: -892, F2: 2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
