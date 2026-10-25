package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS4 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = -483
    var v3 bool = (42 >= -2147483648)
    var v4 string = ("hello" + ("hello" + "foo bar"))
    var v5 int8 = -128
    var v6 uint64 = 437
    switch -2147483648 {
    case -882:
        if v3 {
            var v7 int = -1
            _ = v7
        } else {
            if true {
                var v8 int8 = 0
                _ = v8
            }
            v9 := GenS1{F0: 0, F1: 990}
            _ = v9
        }
    case 42:
        v10 := []int{128}
        _ = v10
    default:
        v11 := v1[4294967295]
        _ = v11
    }
    v12 := GenS2{F0: 4294967295}
    if (!v3) {
        switch 128 {
        case 4294967295:
            v5 -= (v5 >> ((18446744073709551615) & 63))
        case 42:
            for fi13 := 0; fi13 < 9; fi13++ {
                fmt.Println(v3)
                delete(v1, (fi13 % ((4294967295) | 1)))
                _ = fi13
            }
        case 32768:
            switch 0 {
            case -811:
                v1[901] = (-2.25 == 1.5)
            case 4294967295:
                v14 := []int{-898}
                _ = v14
            }
        default:
            v6 += 598
        }
        if (v3 && false) {
            var v15 int32 = int32(v6)
            _ = v15
        }
        switch 32767 {
        case 127:
            v2 = (v2 / (((v2 - (v2 / (((v2 << ((0) & 63))) | 1)))) | 1))
        case -314:
            if v3 {
                v2 -= 1
                _ = len(v1)
                v4 = v4
            } else {
                v6 += v6
                fmt.Println(v4, v6, v2)
                fmt.Println(v6, v3)
            }
        case 1:
            if false {
                for fi16 := 0; fi16 < 20; fi16++ {
                    delete(v1, fi16)
                    v5 = v5
                    v5 = -128
                    _ = fi16
                }
                var v17 int64 = -1
                v4 = (v4 + (v4 + v4))
                _ = v17
            }
        default:
            var v18 int64 = -25
            _ = v18
        }
    } else {
        var v19 int32 = -85
        switch 32768 {
        case 852:
            v4 = ("\n\t" + (v4 + "foo bar"))
        case -201:
            v20 := v1[4294967295]
            _ = v20
        case 605:
            if true {
                switch -9223372036854775808 {
                case 1:
                    for fi21 := 0; fi21 < 12; fi21++ {
                        v22 := GenS1{F0: 83, F1: 387}
                        var v23 int32 = v19
                        for fi24 := 0; fi24 < 9; fi24++ {
                            if v3 {
                                switch (((((("" + v4) + "a") + v4) + v4) + "\n\t") + "") {
                                case "hello":
                                    v19 += (v19 << ((0) & 63))
                                default:
                                    var v25 bool = (!false)
                                    _ = v25
                                }
                                switch ("\n\t" + "") {
                                case "a":
                                    switch (fi21 | 128) {
                                    case -2147483648:
                                        if v3 {
                                            switch (v4 + "a") {
                                            case "":
                                                switch (v4 + "foo bar") {
                                                case "":
                                                    var v26 int16 = 0
                                                    _ = v26
                                                case "a":
                                                    switch fi21 {
                                                    case -2147483648:
                                                        v6 = (v6 << ((18446744073709551615) & 63))
                                                    }
                                                case "foo bar":
                                                    v27 := GenS2{F0: 486}
                                                    _ = v27
                                                }
                                                var v28 uint64 = 42
                                                if v3 {
                                                    v29 := GenS3{F0: int32(v2), F1: 504, F2: 0}
                                                    var v30 int = (-(fi21))
                                                    _ = v29
                                                    _ = v30
                                                } else {
                                                    v31 := []int{fi21, (fi21 * fi21)}
                                                    fmt.Println(fi21)
                                                    if len(v31) > 0 {
                                                        v31[0] = (^(fi21))
                                                    }
                                                    _ = v31
                                                }
                                                _ = v28
                                            default:
                                                var v32 float64 = 3.14159
                                                _ = v32
                                            }
                                            if v3 {
                                                v33 := GenS4{F0: 18446744073709551615, F1: 286, F2: 0}
                                                var v34 int = 471
                                                var v35 float32 = float32(0.0)
                                                _ = v33
                                                _ = v34
                                                _ = v35
                                            }
                                            v1[4294967295] = v3
                                        }
                                    case 265:
                                        fmt.Println(fi24)
                                    case 32767:
                                        fmt.Println(v4, v4, fi21)
                                    default:
                                        v36 := GenS2{F0: 54}
                                        v37 := GenS4{F0: (v6 >> ((1) & 63)), F1: 840, F2: 824}
                                        v38 := GenS5{F0: 2147483647, F1: v5}
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                    }
                                case "\n\t":
                                    if false {
                                        var v39 int8 = (v5 ^ 42)
                                        v6 += v6
                                        _ = v39
                                    } else {
                                        switch 42 {
                                        case 9223372036854775807:
                                            v40 := GenS2{F0: 65535}
                                            _ = v40
                                        default:
                                            switch fi21 {
                                            case 32767:
                                                v41 := GenS4{F0: 248, F1: 2147483648, F2: 2}
                                                _ = v41
                                            case 127:
                                                switch v4 {
                                                case "a":
                                                    v4 = (v4 + ("\n\t" + (v4 + v4)))
                                                case "\n\t":
                                                    v4 = (v4 + "a")
                                                }
                                            default:
                                                v42 := GenS5{F0: 4294967295, F1: 100}
                                                _ = v42
                                            }
                                        }
                                        var v43 float32 = 0.0
                                        _ = v43
                                    }
                                case "foo bar":
                                    var v44 uint8 = 2
                                    _ = v44
                                default:
                                    v1[(fi21 + (fi24 + -1))] = true
                                }
                            } else {
                                v6 += v6
                                v23 += (v23 + 1)
                            }
                            v2 = (v2 + (v2 >> ((65535) & 63)))
                            var v45 float32 = float32(((float64((float32(-2.25) + (0.1 / 0.1))) + 1e10) * 0.1))
                            _ = fi24
                            _ = v45
                        }
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                case 4294967295:
                    v46 := []int{32768, 994}
                    _ = v46
                }
                for fi47 := 0; fi47 < 12; fi47++ {
                    v2 -= (v2 >> ((uint(v6)) & 63))
                    fmt.Println(v2)
                    _ = fi47
                }
            }
        default:
            var v48 uint32 = 2147483647
            _ = v48
        }
        _ = v19
    }
    var v49 float64 = 1e10
    v50 := GenS6{F0: (v5 - (v5 >> ((62) & 63)))}
    var v51 float64 = float64(1e10)
    switch 1 {
    case 477:
        for fi52 := 0; fi52 < 2; fi52++ {
            var v53 float64 = v49
            delete(v1, (fi52 / (((fi52 / (((fi52 << ((518) & 63))) | 1))) | 1)))
            _ = fi52
            _ = v53
        }
    case -2147483648:
        v54 := GenS6{F0: -128}
        _ = v54
    case -678:
        var v55 float64 = (v51 / ((1.5 + (v49 / 1e10)) + -2.25))
        _ = v55
    default:
        if ((1 == 2147483647) && (-1 <= -672)) {
            v2 = (v2 / ((0) | 1))
        }
    }
    var v56 uint8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v56
    _ = v6
}
