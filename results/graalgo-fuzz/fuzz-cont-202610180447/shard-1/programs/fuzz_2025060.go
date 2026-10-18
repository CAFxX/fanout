package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 uint16; F1 int16; F2 int64 }
type GenS5 struct { F0 uint32; F1 int8; F2 int16 }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 int8; F1 int }
type GenS8 struct { F0 uint16; F1 uint }
func genFn0(p0 int32, p1 int16, p2 string) uint64 {
    p2 = ((p2 + "hello") + p2)
    v7 := []int{701, 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 908
}
func genFn1(p0 int, p1 int8) int64 {
    p0 -= (p0 | p0)
    v8 := GenS1{F0: (p0 | p0)}
    v9 := []int{p0, p0, (p0 * (-(p0))), (p0 / (((p0 * (p0 ^ p0))) | 1))}
    var v10 uint8 = 1
    switch ("" + "hello") {
    case "a":
        if false {
            var v11 uint64 = 1
            _ = len(v9)
            p1 -= (p1 * int8(v10))
            _ = v11
        } else {
            if (!false) {
                fmt.Println(p0, p0)
                var v12 int = (p0 ^ p0)
                _ = v12
            }
        }
    case "hello":
        switch "a" {
        case "\n\t":
            p0 = (p0 % ((p0) | 1))
        }
    default:
        _ = len(v9)
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 int8 = -128
    v2 := []uint32{65535, 1, 1, 396}
    var v3 float64 = 0.0
    var v4 string = "\n\t"
    var v5 int32 = -550
    v6 := []int{1, -2147483648}
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 uint16 = 65535
        var v15 bool = (640 < 65535)
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 float64 = 3.14159
            _ = fi16
            _ = v17
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    v18 := GenS2{F0: 2, F1: 4294967295, F2: 383}
    v19 := GenS2{F0: 19, F1: 2, F2: 259}
    _ = len(v6)
    v20 := GenS3{F0: -9223372036854775808, F1: (v4 + v4)}
    var v21 uint8 = 2
    var v22 uint16 = uint16(v21)
    for fi23 := 0; fi23 < 12; fi23++ {
        var v24 uint32 = 2
        for fi25 := 0; fi25 < 16; fi25++ {
            switch v4 {
            case "hello":
                switch ("foo bar" + v4) {
                case "\n\t":
                    fmt.Println(v22, v4)
                    var v26 bool = true
                    var v27 int8 = -128
                    _ = v26
                    _ = v27
                case "hello":
                    var v28 int8 = (v1 << ((18446744073709551615) & 63))
                    _ = v28
                default:
                    v29 := GenS4{F0: (v22 + v22), F1: 529, F2: 585}
                    _ = v29
                }
            case "a":
                v24 = v24
            case "":
                var v30 bool = false
                _ = v30
            default:
                if true {
                    var v31 string = v4
                    var v32 float64 = float64((1.5 + 0.1))
                    var v33 float32 = 0.0
                    _ = v31
                    _ = v32
                    _ = v33
                } else {
                    if ((0.0 + (((1e10 / ((-2.25 / float32(v3)) * 3.14159)) * float32(v3)) / 1.0)) == 1.5) {
                        var v34 uint8 = (v21 | 127)
                        var v35 uint32 = v24
                        if false {
                            var v36 uint16 = (v22 >> ((500) & 63))
                            _ = v36
                        }
                        _ = v34
                        _ = v35
                    } else {
                        switch fi25 {
                        case 32767:
                            v37 := GenS3{F0: -207, F1: v4}
                            _ = v37
                        }
                        v38 := GenS5{F0: 458, F1: 63, F2: -29}
                        var v39 uint8 = 255
                        _ = v38
                        _ = v39
                    }
                    v40 := GenS5{F0: v24, F1: 1, F2: 42}
                    var v41 int64 = int64(v1)
                    _ = v40
                    _ = v41
                }
            }
            if true {
                v2 = append(v2, (^(v24)))
                _ = len(v6)
                v4 = ("foo bar" + "")
            } else {
                v42 := GenS6{F0: v5, F1: true}
                _ = v42
            }
            if len(v6) > 0 {
                v6[0] = int(v1)
            }
            var v43 string = "foo bar"
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi25
            _ = v43
        }
        _ = fi23
        _ = v24
    }
    var v44 uint8 = 1
    v45 := GenS7{F0: v1, F1: 1}
    v2 = append(v2, 1)
    v22 -= (v22 - 32767)
    v46 := GenS7{F0: (v1 ^ int8(v5)), F1: 206}
    v47 := GenS8{F0: (v22 / ((2) | 1)), F1: 255}
    v44 = (v21 / (((v21 * (^(v21)))) | 1))
    var v48 int32 = (v5 * (v5 | (v5 * (v5 & (v5 & 2147483647)))))
    var v49 uint64 = 948
    v5 = (-(v5))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
