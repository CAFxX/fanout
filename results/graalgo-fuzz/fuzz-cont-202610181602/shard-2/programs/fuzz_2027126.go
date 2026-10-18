package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 uint16; F1 uint }
type GenS6 struct { F0 int32 }

func main() {
    var v1 string = ""
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]bool)
    var v4 uint8 = 43
    var v5 float64 = 0.0
    v4 += (v4 | v4)
    switch v1 {
    case "foo bar":
        v6 := GenS1{F0: (^(v4)), F1: 1}
        _ = v6
    default:
        v4 += 190
    }
    var v7 uint16 = 955
    var v8 int32 = -651
    v8 = v8
    switch -1 {
    case -951:
        var v9 string = v1
        _ = v9
    case 32767:
        if true {
            v10 := []int{653, -934, -718}
            _ = v10
        } else {
            var v11 int64 = 1
            _ = v11
        }
    default:
        var v12 int32 = v8
        _ = v12
    }
    v13 := []int{9223372036854775807, 1, 127, 2147483647}
    if (1 != 301) {
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 uint8 = (v4 % (((v4 & 42)) | 1))
            for fi16 := 0; fi16 < 11; fi16++ {
                var v17 uint = 10
                v18 := GenS1{F0: (v4 - v4), F1: 9223372036854775807}
                var v19 uint = 18446744073709551615
                v15 -= 16
                v17 += uint(v15)
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            _ = len(v2)
            _ = fi14
            _ = v15
        }
        v20 := GenS1{F0: v4, F1: -1}
        for fi21 := 0; fi21 < 16; fi21++ {
            var v22 uint32 = 435
            if true {
                if true {
                    v7 = v7
                    _ = len(v3)
                    for fi23 := 0; fi23 < 11; fi23++ {
                        v8 += (v8 % ((42) | 1))
                        _ = fi23
                    }
                } else {
                    if (!((v8 % ((v8) | 1)) < (^(v8)))) {
                        _ = len(v13)
                        v1 = v1
                        switch v1 {
                        case "\n\t":
                            var v24 uint64 = 0
                            _ = v24
                        case "foo bar":
                            fmt.Println(v8, fi21, v4)
                        default:
                            if true {
                                if (v1 <= "foo bar") {
                                    fmt.Println(v22, v5)
                                } else {
                                    v22 = 2147483647
                                }
                            } else {
                                for fi25 := 0; fi25 < 18; fi25++ {
                                    switch v1 {
                                    case "":
                                        if true {
                                            if false {
                                                var v26 string = "foo bar"
                                                switch v1 {
                                                case "a":
                                                    v27 := GenS2{F0: (0 > 834), F1: 18446744073709551615}
                                                    _ = v27
                                                default:
                                                    var v28 int = fi21
                                                    v29 := GenS1{F0: (v4 - v4), F1: 760}
                                                    _ = v28
                                                    _ = v29
                                                }
                                                v22 = (^(v22))
                                                _ = v26
                                            }
                                        }
                                    case "\n\t":
                                        if (0 <= -1) {
                                            if (false && false) {
                                                var v30 int64 = -791
                                                _ = v30
                                            } else {
                                                switch "" {
                                                case "foo bar":
                                                    v31 := GenS3{F0: (^(v8))}
                                                    _ = v31
                                                default:
                                                    _ = len(v3)
                                                }
                                            }
                                            switch fi21 {
                                            case 32768:
                                                var v32 bool = (-9223372036854775808 != 1)
                                                _ = v32
                                            }
                                            var v33 int16 = 42
                                            _ = v33
                                        } else {
                                            v34 := GenS4{F0: v22, F1: true}
                                            switch fi21 {
                                            case 0:
                                                v4 = v4
                                            case 193:
                                                fmt.Println(fi21)
                                            case 32768:
                                                var v35 string = v1
                                                _ = v35
                                            default:
                                                var v36 int = (fi21 >> ((65535) & 63))
                                                _ = v36
                                            }
                                            _ = v34
                                        }
                                    }
                                    fmt.Println(fi25, v8, v8)
                                    _ = fi25
                                }
                                v22 -= v22
                                fmt.Println(v7, v7)
                            }
                        }
                    } else {
                        if true {
                            for fi37 := 0; fi37 < 17; fi37++ {
                                v1 = v1
                                fmt.Println(v1, v22)
                                v38 := GenS4{F0: 433, F1: false}
                                _ = fi37
                                _ = v38
                            }
                        }
                        switch (v1 + v1) {
                        case "hello":
                            if (0 > 9223372036854775807) {
                                v39 := GenS5{F0: (v7 & (v7 / (((v7 ^ (v7 % ((v7) | 1)))) | 1))), F1: uint(v22)}
                                v1 = "foo bar"
                                _ = v39
                            } else {
                                v22 -= (v22 + 2147483647)
                            }
                        case "a":
                            if true {
                                v22 += v22
                                switch 32768 {
                                case 0:
                                    v40 := GenS2{F0: (1 != 9223372036854775807), F1: 727}
                                    _ = v40
                                case -9223372036854775808:
                                    v41 := GenS4{F0: (v22 / ((v22) | 1)), F1: (18446744073709551615 == 163)}
                                    _ = v41
                                }
                            } else {
                                switch (^(fi21)) {
                                case 2147483647:
                                    var v42 int = 42
                                    _ = v42
                                case 127:
                                    _ = len(v3)
                                }
                                var v43 int8 = 42
                                fmt.Println(v22, v5, v43)
                                _ = v43
                            }
                        case "":
                            for fi44 := 0; fi44 < 10; fi44++ {
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                                v45 := GenS6{F0: int32(v4)}
                                _ = fi44
                                _ = v45
                            }
                        default:
                            var v46 float32 = ((0.1 * (-2.25 + -2.25)) / 1.0)
                            _ = v46
                        }
                        var v47 uint32 = (v22 & (v22 - (^(v22))))
                        _ = v47
                    }
                    var v48 int64 = -9223372036854775808
                    _ = v48
                }
                if len(v13) > 0 {
                    v13[0] = fi21
                }
                v49 := GenS6{F0: (-(v8))}
                _ = v49
            } else {
                var v50 uint = 42
                for fi51 := 0; fi51 < 14; fi51++ {
                    if false {
                        var v52 int16 = -32768
                        fmt.Println(v7, v50)
                        _ = v52
                    } else {
                        v4 = v4
                    }
                    var v53 int32 = (v8 / ((42) | 1))
                    _ = fi51
                    _ = v53
                }
                _ = v50
            }
            v13 = append(v13, (fi21 ^ (fi21 - 1)))
            _ = fi21
            _ = v22
        }
        _ = v20
    }
    v7 = v7
    if false {
        switch (("foo bar" + v1) + (v1 + v1)) {
        case "foo bar":
            var v54 string = ((v1 + "") + "")
            _ = v54
        case "":
            fmt.Println(v8)
        default:
            if true {
                if (true || ((v7 ^ 32767) != 42)) {
                    for fi55 := 0; fi55 < 4; fi55++ {
                        fmt.Println(v4, v7)
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        for fi56 := 0; fi56 < 19; fi56++ {
                            v4 = 42
                            switch fi55 {
                            case 42:
                                if len(v13) > 0 {
                                    v13[0] = 128
                                }
                            default:
                                v1 = v1
                                if true {
                                    v7 += 65535
                                    v5 += -2.25
                                    v57 := GenS4{F0: 65535, F1: (v4 <= 255)}
                                    _ = v57
                                } else {
                                    fmt.Println(v8, fi55, v8)
                                    v5 -= v5
                                    var v58 bool = false
                                    _ = v58
                                }
                                switch (("hello" + v1) + v1) {
                                case "foo bar":
                                    var v59 float32 = ((1e10 * (1.5 + ((float32(v5) * 3.14159) * 3.14159))) * 1e10)
                                    _ = v59
                                case "a":
                                    delete(v3, 2147483647)
                                case "hello":
                                    v1 = v1
                                default:
                                    v60 := GenS2{F0: true, F1: 4294967295}
                                    _ = v60
                                }
                            }
                            v8 = -429
                            _ = fi56
                        }
                        _ = fi55
                    }
                } else {
                    fmt.Println(v8)
                }
                v61 := GenS4{F0: 65535, F1: (!true)}
                v5 = v5
                _ = v61
            } else {
                fmt.Println(v7, v1, v1)
                v7 -= v7
                v4 = (v4 & 42)
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
