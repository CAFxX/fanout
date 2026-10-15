package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 int32 }

func main() {
    var v1 int16 = 1
    var v2 int16 = (v1 | 1000)
    var v3 uint16 = 255
    var v4 float32 = (-2.25 / 1.0)
    if (1 != 1) {
        switch "hello" {
        case "hello":
            v5 := []int{2147483647, 368, 2147483647}
            _ = v5
        case "foo bar":
            for fi6 := 0; fi6 < 11; fi6++ {
                var v7 string = "a"
                for fi8 := 0; fi8 < 17; fi8++ {
                    var v9 int64 = -9223372036854775808
                    v1 = (v1 | 1000)
                    var v10 uint = 18446744073709551615
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                fmt.Println(v7, v4, v1)
                _ = fi6
                _ = v7
            }
        }
        if false {
            for fi11 := 0; fi11 < 17; fi11++ {
                v12 := GenS1{F0: 580}
                switch (fi11 - 32768) {
                case 836:
                    if ((float64(v4) <= 1e10) && true) {
                        v13 := GenS1{F0: 2147483647}
                        _ = v13
                    } else {
                        switch -9223372036854775808 {
                        case 128:
                            var v14 int8 = int8(v2)
                            _ = v14
                        case 1:
                            for fi15 := 0; fi15 < 6; fi15++ {
                                if false {
                                    v1 = (-(v2))
                                } else {
                                    v16 := GenS2{F0: (false || true), F1: 42}
                                    _ = v16
                                }
                                if (false || true) {
                                    v3 = 65535
                                } else {
                                    if false {
                                        var v17 int64 = -9223372036854775808
                                        _ = v17
                                    } else {
                                        v18 := GenS1{F0: fi11}
                                        _ = v18
                                    }
                                }
                                var v19 uint = 65535
                                var v20 int64 = 9223372036854775807
                                v3 += (^(v3))
                                var v21 uint = (v19 - v19)
                                _ = fi15
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                        default:
                            for fi22 := 0; fi22 < 3; fi22++ {
                                var v23 uint64 = 16
                                v24 := []int{(-(fi22)), (-(fi11)), (fi22 % ((128) | 1)), (fi22 | (-(fi11)))}
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                        }
                        v3 = 656
                        if (((1 < 2) && (1 > 117)) && (!((v2 << ((384) & 63)) > 42))) {
                            for fi25 := 0; fi25 < 5; fi25++ {
                                var v26 int16 = (v1 + 1000)
                                switch -42 {
                                case -2147483648:
                                    var v27 uint32 = 2147483648
                                    _ = v27
                                case -9223372036854775808:
                                    fmt.Println(v1, v3, fi25)
                                case -543:
                                    var v28 int8 = -1
                                    _ = v28
                                }
                                switch fi11 {
                                case -137:
                                    v4 -= (((v4 + float32(1.5)) + 1e10) * float32((float64(v4) * float64((v4 - -2.25)))))
                                case -2147483648:
                                    fmt.Println(v3, v4)
                                default:
                                    v26 = v1
                                }
                                v3 = 944
                                var v29 int8 = -128
                                _ = fi25
                                _ = v26
                                _ = v29
                            }
                        }
                    }
                case 0:
                    v30 := GenS3{F0: 1}
                    _ = v30
                case -2147483648:
                    v31 := []int{-169, 128, fi11, (fi11 << ((972) & 63))}
                    _ = v31
                default:
                    fmt.Println(v2, v2, v3)
                }
                var v32 float64 = ((0.0 - -2.25) - 1e10)
                _ = fi11
                _ = v12
                _ = v32
            }
            var v33 int32 = -489
            v33 = (v33 >> ((255) & 63))
            _ = v33
        } else {
            if true {
                var v34 bool = false
                _ = v34
            }
            v35 := []int{128, 923}
            var v36 uint8 = 16
            _ = v35
            _ = v36
        }
    }
    if (true || false) {
        var v37 int64 = 1
        _ = v37
    }
    var v38 float64 = 0.1
    fmt.Println(v1)
    var v39 int8 = 127
    if false {
        fmt.Println(v1, v2)
        fmt.Println(v38, v39, v39)
    }
    for fi40 := 0; fi40 < 17; fi40++ {
        v41 := GenS3{F0: 0}
        v42 := []int{fi40, fi40}
        _ = fi40
        _ = v41
        _ = v42
    }
    v43 := GenS3{F0: 42}
    fmt.Println(v1, v2, v4)
    if (255 == 1) {
        if (0 > 0) {
            v44 := []int{32768, -1, 1, -2147483648}
            v3 += v3
            _ = v44
        }
        var v45 float64 = float64((((-2.25 - 1e10) * (v4 - (v4 * ((v4 / -2.25) * ((3.14159 + v4) - (v4 - v4)))))) + (0.0 * v4)))
        _ = v45
    }
    v3 -= v3
    switch "a" {
    case "foo bar":
        if false {
            v46 := GenS4{F0: (^(v2)), F1: -2147483648}
            for fi47 := 0; fi47 < 17; fi47++ {
                for fi48 := 0; fi48 < 12; fi48++ {
                    var v49 int16 = v2
                    if false {
                        v4 = 3.14159
                        switch ("a" + "hello") {
                        case "":
                            var v50 int16 = (-(v1))
                            _ = v50
                        }
                        v3 = (v3 - (v3 & 182))
                    } else {
                        v51 := []int{}
                        if len(v51) > 0 {
                            v51[0] = (fi47 ^ int(v39))
                        }
                        v3 = v3
                        _ = v51
                    }
                    _ = fi48
                    _ = v49
                }
                var v52 int16 = (v2 << ((395) & 63))
                v53 := GenS2{F0: (!((false || false) && false)), F1: 0}
                _ = fi47
                _ = v52
                _ = v53
            }
            _ = v46
        }
    }
    v38 += v38
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v43
}
