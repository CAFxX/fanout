package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int32; F2 int16 }
type GenS3 struct { F0 uint64; F1 uint; F2 uint32 }
type GenS4 struct { F0 int8; F1 uint8; F2 int8 }

func main() {
    var v1 uint32 = 1
    var v2 float32 = (0.0 - -2.25)
    var v3 int32 = 151
    switch ("a" + "hello") {
    case "foo bar":
        v4 := []int{-1}
        _ = v4
    }
    v2 = float32((3.14159 + 3.14159))
    v1 = 2147483647
    v5 := GenS1{F0: false}
    if (!(v1 < v1)) {
        switch ("foo bar" + ("\n\t" + "hello")) {
        case "\n\t":
            for fi6 := 0; fi6 < 3; fi6++ {
                v1 = v1
                v7 := GenS2{F0: 1, F1: 42, F2: 0}
                _ = fi6
                _ = v7
            }
        default:
            for fi8 := 0; fi8 < 1; fi8++ {
                for fi9 := 0; fi9 < 15; fi9++ {
                    if false {
                        v1 -= (v1 % ((1) | 1))
                        switch "\n\t" {
                        case "":
                            if (false && ((!(!(100 <= 42))) || (false && (!false)))) {
                                v10 := GenS1{F0: false}
                                switch (fi9 * (fi9 / (((fi8 & 944)) | 1))) {
                                case 32767:
                                    fmt.Println(fi8, fi9, v3)
                                default:
                                    v11 := []int{fi9, (fi9 % ((0) | 1)), 32767, 32768}
                                    _ = v11
                                }
                                _ = v10
                            } else {
                                var v12 bool = true
                                _ = v12
                            }
                        case "a":
                            if false {
                                v13 := GenS1{F0: (!(128 != 117))}
                                v14 := []int{fi8, 32767, (fi9 + 127), 2147483647}
                                _ = v13
                                _ = v14
                            } else {
                                var v15 float64 = float64(1e10)
                                _ = v15
                            }
                            v2 -= v2
                            v16 := GenS1{F0: (false && true)}
                            _ = v16
                        case "foo bar":
                            if false {
                                v17 := []int{(fi8 / ((2147483647) | 1)), fi8, 32768, (fi9 * (fi9 / ((fi9) | 1)))}
                                switch (fi9 & 4294967295) {
                                case -2147483648:
                                    fmt.Println(v3, v1)
                                    var v18 int16 = -1
                                    var v19 uint32 = v1
                                    _ = v18
                                    _ = v19
                                case 32767:
                                    v20 := GenS3{F0: 42, F1: 93, F2: (v1 % ((320) | 1))}
                                    _ = v20
                                default:
                                    var v21 bool = (42 < -31)
                                    _ = v21
                                }
                                var v22 int = (fi8 >> ((4294967295) & 63))
                                v3 += (v3 + (v3 & 2147483647))
                                _ = v17
                                _ = v22
                            } else {
                                v3 = v3
                                switch "foo bar" {
                                case "a":
                                    v23 := []int{(fi8 + 1), 128}
                                    _ = v23
                                case "":
                                    var v24 uint8 = 2
                                    var v25 float64 = 3.14159
                                    var v26 uint32 = (^(v1))
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                case "\n\t":
                                    if ((^(v3)) >= -1) {
                                        v1 = v1
                                        if ("" == "a") {
                                            if ((!(false || (false || false))) || (!true)) {
                                                var v27 int16 = 669
                                                v27 += v27
                                                _ = v27
                                            }
                                            v28 := GenS4{F0: 127, F1: 42, F2: int8(v3)}
                                            _ = v28
                                        } else {
                                            fmt.Println(v3, v3)
                                            var v29 float32 = float32(0.1)
                                            _ = v29
                                        }
                                    } else {
                                        v1 += (v1 >> ((65535) & 63))
                                    }
                                }
                            }
                            v2 += (((((0.0 * -2.25) / ((float32(-2.25) / (float32(3.14159) / 1.5)) - float32((3.14159 * ((1.5 - -2.25) + (float64(3.14159) - (1e10 + 0.1))))))) * v2) + (v2 + v2)) * (((((v2 - float32(3.14159)) - v2) - 1e10) - 3.14159) + v2))
                        }
                    }
                    switch fi8 {
                    case 32767:
                        switch 2147483647 {
                        case 0:
                            var v30 uint8 = 255
                            _ = v30
                        case -258:
                            v1 -= (v1 * (v1 | 147))
                        case 517:
                            v31 := []int{(fi9 >> ((42) & 63)), (fi9 >> ((18446744073709551615) & 63)), fi8}
                            var v32 int8 = 42
                            var v33 uint = 65535
                            _ = v31
                            _ = v32
                            _ = v33
                        default:
                            if false {
                                fmt.Println(fi9, v1)
                            }
                        }
                    case -834:
                        var v34 float64 = 0.1
                        _ = v34
                    }
                    _ = fi9
                }
                v35 := []int{fi8, fi8, 275}
                var v36 int = (^(fi8))
                _ = fi8
                _ = v35
                _ = v36
            }
        }
        v37 := []int{4294967295, 32767}
        for fi38 := 0; fi38 < 4; fi38++ {
            v39 := GenS2{F0: 255, F1: (v3 >> ((345) & 63)), F2: 32767}
            if false {
                var v40 uint8 = 2
                var v41 uint = 255
                v41 += (v41 & v41)
                _ = v40
                _ = v41
            }
            _ = fi38
            _ = v39
        }
        _ = v37
    }
    for fi42 := 0; fi42 < 19; fi42++ {
        if false {
            if (!true) {
                fmt.Println(v2, v2, v2)
                var v43 string = ((("a" + ("hello" + "\n\t")) + "") + ((((("" + "hello") + "\n\t") + ("\n\t" + "a")) + "") + (("hello" + ((("foo bar" + ("foo bar" + "hello")) + "hello") + "hello")) + "")))
                for fi44 := 0; fi44 < 20; fi44++ {
                    v45 := []int{(fi42 << ((18446744073709551615) & 63)), (fi44 | (-(fi42))), (^(fi42)), (fi44 % ((32768) | 1))}
                    var v46 uint8 = 255
                    var v47 int32 = (v3 ^ int32(v46))
                    _ = fi44
                    _ = v45
                    _ = v46
                    _ = v47
                }
                _ = v43
            } else {
                v1 -= (v1 - v1)
            }
            v3 -= -1
        }
        v1 -= v1
        v48 := []int{0, (fi42 / ((128) | 1))}
        _ = fi42
        _ = v48
    }
    if (v1 >= v1) {
        v49 := GenS2{F0: 2, F1: -1, F2: -32768}
        _ = v49
    }
    for fi50 := 0; fi50 < 4; fi50++ {
        var v51 uint = 0
        _ = fi50
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
}
