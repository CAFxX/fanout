package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int16 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 uint8 = 42
    var v2 int16 = 42
    v3 := make(map[string]uint32)
    var v4 int16 = 390
    v5 := []string{"", "a"}
    var v6 uint8 = (v1 << ((2) & 63))
    v1 += v6
    v6 -= (^(v6))
    var v7 float64 = (float64(1e10) + 1.5)
    switch -9223372036854775808 {
    case -2147483648:
        var v8 int16 = 1
        _ = v8
    case 822:
        var v9 uint16 = 1
        _ = v9
    case 42:
        if (false && false) {
            for fi10 := 0; fi10 < 8; fi10++ {
                _ = len(v5)
                if (v4 >= v2) {
                    var v11 uint16 = 255
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    switch ("hello" + ("hello" + ("a" + ""))) {
                    case "":
                        var v12 int16 = v2
                        _ = v12
                    case "\n\t":
                        if false {
                            fmt.Println(v2)
                            v13 := GenS1{F0: uint32(v1), F1: 0, F2: v2}
                            _ = v13
                        }
                    case "hello":
                        v1 = v6
                    }
                    _ = v11
                } else {
                    switch "a" {
                    case "foo bar":
                        v7 = v7
                    case "hello":
                        for fi14 := 0; fi14 < 16; fi14++ {
                            v6 = 128
                            var v15 float32 = 0.0
                            if (1 > 65535) {
                                var v16 int32 = 428
                                var v17 uint64 = 727
                                _ = v16
                                _ = v17
                            }
                            fmt.Println(v4)
                            v1 = (^(v6))
                            _ = fi14
                            _ = v15
                        }
                    case "\n\t":
                        var v18 float64 = -2.25
                        _ = v18
                    default:
                        delete(v3, "foo bar")
                    }
                    v7 += v7
                    var v19 int32 = 42
                    _ = v19
                }
                _ = fi10
            }
            var v20 uint16 = 0
            v21 := GenS1{F0: 65535, F1: 42, F2: v4}
            _ = v20
            _ = v21
        } else {
            var v22 int16 = v4
            _ = v22
        }
    default:
        var v23 string = "foo bar"
        _ = v23
    }
    var v24 uint64 = 317
    if true {
        v4 = v4
        if true {
            if (255 == (v1 << ((uint(v6)) & 63))) {
                v6 -= v1
                var v25 string = "foo bar"
                _ = v25
            } else {
                for fi26 := 0; fi26 < 14; fi26++ {
                    v27 := GenS2{F0: (v1 % (((v1 - (v6 >> ((255) & 63)))) | 1)), F1: 127}
                    v1 -= (v1 | (v1 & uint8(v24)))
                    _ = fi26
                    _ = v27
                }
                v28 := GenS1{F0: 65535, F1: 0, F2: (v2 % ((42) | 1))}
                var v29 uint8 = (^(v1))
                _ = v28
                _ = v29
            }
            var v30 uint16 = 255
            for fi31 := 0; fi31 < 16; fi31++ {
                v32 := GenS1{F0: 2, F1: int32(v30), F2: (v4 % (((v2 + (v2 + v4))) | 1))}
                for fi33 := 0; fi33 < 6; fi33++ {
                    var v34 string = ("" + ("a" + ""))
                    _ = fi33
                    _ = v34
                }
                for fi35 := 0; fi35 < 12; fi35++ {
                    if false {
                        v2 = 1
                        var v36 int32 = -496
                        if (uint16(v24) >= v30) {
                            var v37 uint8 = (v6 >> ((18446744073709551615) & 63))
                            fmt.Println(v37, v36)
                            v24 += uint64(v30)
                            _ = v37
                        }
                        _ = v36
                    }
                    v7 += 1e10
                    if true {
                        v7 = 0.1
                    } else {
                        if ((v24 - v24) <= (v24 - uint64(v30))) {
                            switch "hello" {
                            case "hello":
                                v24 += v24
                            default:
                                delete(v3, "hello")
                            }
                            v4 += (v4 / ((-32768) | 1))
                            var v38 int32 = 0
                            var v39 string = "hello"
                            _ = v38
                            _ = v39
                        } else {
                            if (18446744073709551615 == 18446744073709551615) {
                                if false {
                                    v30 += v30
                                } else {
                                    var v40 float64 = float64(((0.1 * (((0.1 * (float32((v7 - v7)) * (1.5 - 0.0))) * -2.25) + 0.1)) - (-2.25 + 0.1)))
                                    if (100 >= 0) {
                                        v41 := GenS2{F0: (^(v1)), F1: 1}
                                        v42 := GenS3{F0: (v4 % ((947) | 1))}
                                        v43 := v3[((("\n\t" + "") + "hello") + ((("hello" + ("hello" + "a")) + (("" + ("a" + ("hello" + "a"))) + "foo bar")) + ("\n\t" + "foo bar")))]
                                        _ = v41
                                        _ = v42
                                        _ = v43
                                    } else {
                                        v4 = 32767
                                    }
                                    _ = v40
                                }
                                v1 -= v6
                                var v44 uint8 = (v1 & v6)
                                _ = v44
                            } else {
                                if true {
                                    v45 := v3[("" + ("hello" + "foo bar"))]
                                    _ = v45
                                }
                                if (float64(-2.25) > v7) {
                                    v46 := v3["hello"]
                                    v2 = int16(v46)
                                    _ = len(v5)
                                    _ = v46
                                }
                                switch fi31 {
                                case 32768:
                                    var v47 uint8 = (v1 | v1)
                                    _ = v47
                                default:
                                    fmt.Println(v24, v30)
                                }
                                var v48 bool = true
                                fmt.Println(v2)
                                fmt.Println(v6, v24, v6)
                                var v49 uint64 = v24
                                _ = v48
                                _ = v49
                            }
                            v50 := GenS2{F0: v1, F1: -86}
                            var v51 string = "foo bar"
                            if (782 < 2) {
                                v30 += (v30 + v30)
                                switch (fi31 % ((fi35) | 1)) {
                                case 2147483647:
                                    switch "a" {
                                    case "foo bar":
                                        v52 := GenS4{F0: 2}
                                        v24 += v24
                                        _ = v52
                                    case "\n\t":
                                        var v53 int8 = -128
                                        _ = v53
                                    default:
                                        v7 = float64(float32(v7))
                                    }
                                case 127:
                                    var v54 int8 = -123
                                    _ = v54
                                case -1:
                                    v55 := v3[v51]
                                    _ = v55
                                }
                                switch fi35 {
                                case -799:
                                    v56 := GenS1{F0: 498, F1: -1, F2: v4}
                                    _ = v56
                                case 127:
                                    v30 -= (v30 / (((v30 & (v30 | (^(v30))))) | 1))
                                }
                            } else {
                                v7 += v7
                                var v57 uint64 = v24
                                var v58 float64 = 0.1
                                _ = v57
                                _ = v58
                            }
                            _ = v50
                            _ = v51
                        }
                        v59 := GenS2{F0: (v1 << ((uint(v1)) & 63)), F1: 127}
                        v5 = append(v5, "")
                        _ = v59
                    }
                    _ = fi35
                }
                _ = fi31
                _ = v32
            }
            _ = v30
        } else {
            var v60 int = 321
            _ = v60
        }
    } else {
        _ = len(v3)
        for fi61 := 0; fi61 < 16; fi61++ {
            var v62 uint = 0
            _ = fi61
            _ = v62
        }
    }
    var v63 uint32 = 57
    var v64 uint8 = 127
    switch "foo bar" {
    case "hello":
        v1 += 1
    case "\n\t":
        v4 += v4
    case "foo bar":
        fmt.Println(v6, v63, v24)
    default:
        v63 -= 2147483648
    }
    var v65 uint32 = (v63 * v63)
    v66 := GenS2{F0: v6, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v64
    _ = v65
    _ = v66
    _ = v7
}
