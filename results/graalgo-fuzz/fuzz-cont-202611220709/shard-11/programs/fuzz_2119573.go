package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 string }
type GenS2 struct { F0 uint32; F1 int64; F2 int }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 int16; F1 uint16 }

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 0
    var v3 uint16 = 450
    var v4 float64 = 1e10
    var v5 uint32 = 0
    v4 = 0.1
    v2 = (^(v2))
    v4 -= float64(0.0)
    v1["a"] = "\n\t"
    v4 += ((v4 * 1.5) / -2.25)
    v6 := []int{-1, 4294967295, -9223372036854775808}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v7 int16 = 42
    v7 -= v7
    v7 = -32768
    if (false || ((uint(v3) <= 65535) && true)) {
        if false {
            for fi8 := 0; fi8 < 18; fi8++ {
                if ((!(!(0 <= (v2 + v2)))) || true) {
                    v9 := v1["foo bar"]
                    _ = v9
                } else {
                    var v10 uint8 = 148
                    var v11 string = "hello"
                    switch fi8 {
                    case 4294967295:
                        switch v11 {
                        case "":
                            v12 := GenS1{F0: (false || (-732 <= -147)), F1: 42, F2: (v11 + v11)}
                            _ = v12
                        case "hello":
                            switch fi8 {
                            case 773:
                                var v13 bool = (-1 <= (fi8 >> ((65535) & 63)))
                                _ = v13
                            case -9223372036854775808:
                                for fi14 := 0; fi14 < 10; fi14++ {
                                    v15 := v1[((((v11 + ("hello" + v11)) + v11) + (v11 + "\n\t")) + v11)]
                                    v16 := GenS2{F0: (v5 + v5), F1: 42, F2: (fi8 / (((fi8 & (-(fi14)))) | 1))}
                                    if true {
                                        switch "\n\t" {
                                        case "\n\t":
                                            if true {
                                                v4 = (1e10 / -2.25)
                                                fmt.Println(v3, fi8)
                                                v2 = 4294967295
                                            } else {
                                                if (uint64(v3) >= (v2 << ((uint(v5)) & 63))) {
                                                    v17 := GenS3{F0: 18446744073709551615, F1: (false && (42 == (v3 * v3)))}
                                                    var v18 float32 = (1.5 + (3.14159 - ((3.14159 + 1.5) + (1.5 - ((1e10 / float32((v4 - -2.25))) - 0.1)))))
                                                    v5 -= 2147483647
                                                    _ = v17
                                                    _ = v18
                                                }
                                                v19 := GenS4{F0: (v7 >> ((255) & 63)), F1: (v3 / ((32767) | 1))}
                                                _ = v19
                                            }
                                        case "":
                                            v7 = v7
                                        }
                                        if (-1 < 0) {
                                            var v20 int64 = 0
                                            _ = v20
                                        } else {
                                            v21 := GenS3{F0: 65535, F1: false}
                                            var v22 string = ("\n\t" + "foo bar")
                                            _ = v21
                                            _ = v22
                                        }
                                    } else {
                                        switch fi8 {
                                        case 127:
                                            fmt.Println(v2, fi8, v5)
                                            var v23 int = (fi14 | fi8)
                                            _ = v23
                                        case 0:
                                            delete(v1, "a")
                                        case 128:
                                            v10 = (^(v10))
                                        default:
                                            if (!(1 >= -128)) {
                                                var v24 int = fi14
                                                v2 -= (v2 % ((v2) | 1))
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                _ = v24
                                            }
                                        }
                                        var v25 uint16 = 255
                                        _ = v25
                                    }
                                    if (!true) {
                                        v26 := GenS4{F0: v7, F1: v3}
                                        _ = v26
                                    }
                                    _ = fi14
                                    _ = v15
                                    _ = v16
                                }
                            case -121:
                                v3 = 2
                            default:
                                var v27 bool = true
                                _ = v27
                            }
                        default:
                            var v28 int16 = v7
                            _ = v28
                        }
                    default:
                        v29 := GenS4{F0: (v7 - (v7 << ((65535) & 63))), F1: 42}
                        _ = v29
                    }
                    _ = v10
                    _ = v11
                }
                _ = fi8
            }
            for fi30 := 0; fi30 < 7; fi30++ {
                v4 = (float64(((0.1 / 0.1) + 1e10)) - -2.25)
                switch fi30 {
                case 0:
                    v7 -= v7
                case 32767:
                    var v31 int = fi30
                    _ = v31
                case -662:
                    for fi32 := 0; fi32 < 19; fi32++ {
                        switch (-(fi30)) {
                        case -389:
                            v33 := GenS3{F0: uint(v3), F1: false}
                            _ = v33
                        case -483:
                            delete(v1, "a")
                        case 1:
                            var v34 int = (^(fi32))
                            _ = v34
                        default:
                            v35 := GenS3{F0: 2, F1: true}
                            _ = v35
                        }
                        v3 = (v3 >> ((1) & 63))
                        if (!(v4 > v4)) {
                            v36 := GenS2{F0: (v5 << ((0) & 63)), F1: 9223372036854775807, F2: int(v5)}
                            if true {
                                v1["hello"] = ("hello" + "hello")
                            }
                            _ = v36
                        }
                        _ = fi32
                    }
                default:
                    v5 += v5
                }
                for fi37 := 0; fi37 < 9; fi37++ {
                    v7 = int16(v5)
                    var v38 uint64 = (v2 * v2)
                    var v39 int16 = v7
                    _ = fi37
                    _ = v38
                    _ = v39
                }
                _ = fi30
            }
            var v40 uint64 = (v2 + 1)
            _ = v40
        }
        v4 += 0.1
        if len(v6) > 0 {
            v6[0] = 127
        }
    } else {
        fmt.Println(v3, v5, v5)
        v41 := GenS3{F0: 178, F1: false}
        _ = v41
    }
    if (0 < 127) {
        var v42 float64 = 0.0
        _ = v42
    } else {
        if false {
            var v43 int32 = 1
            switch 32767 {
            case -250:
                var v44 int = 32767
                _ = v44
            default:
                if (9223372036854775807 <= 9223372036854775807) {
                    v45 := GenS3{F0: 255, F1: false}
                    v7 -= v7
                    v46 := GenS1{F0: true, F1: 390, F2: ((("" + "foo bar") + "\n\t") + "a")}
                    _ = v45
                    _ = v46
                } else {
                    v43 = v43
                    if (843 <= 127) {
                        v43 = (v43 << ((0) & 63))
                    } else {
                        v5 = (v5 | (^(v5)))
                    }
                }
            }
            v47 := GenS2{F0: (v5 % ((v5) | 1)), F1: 9223372036854775807, F2: 303}
            _ = v43
            _ = v47
        }
        v5 -= (v5 | (v5 & uint32(v2)))
    }
    _ = len(v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
