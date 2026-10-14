package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int8 }
type GenS2 struct { F0 int; F1 string; F2 uint }
type GenS3 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 int64; F1 int32; F2 int16 }

func main() {
    var v1 uint32 = 0
    v2 := make(map[uint32]uint32)
    var v3 uint = 255
    v4 := []int{636, 32768, 0}
    v5 := v2[(^(v1))]
    v1 += (v1 & 406)
    v6 := GenS1{F0: v3, F1: -465, F2: 127}
    v1 = v5
    var v7 float32 = ((float32(0.0) - (1.5 * 0.1)) / (1.5 + 1.5))
    v8 := GenS2{F0: 19, F1: "hello", F2: v3}
    v3 = 4294967295
    switch 1 {
    case 0:
        switch ("a" + ("" + "foo bar")) {
        case "":
            _ = len(v4)
        case "\n\t":
            var v9 bool = ((v3 / (((^(v3))) | 1)) >= (v3 << ((v3) & 63)))
            _ = v9
        case "hello":
            v1 = (v1 | (v1 - 2147483648))
        default:
            v5 += 65535
        }
    case 459:
        _ = len(v2)
    default:
        v10 := v2[2147483648]
        _ = v10
    }
    if true {
        v11 := GenS1{F0: (v3 >> ((42) & 63)), F1: 9223372036854775807, F2: 42}
        switch ("\n\t" + "\n\t") {
        case "foo bar":
            if true {
                v4 = append(v4, 4294967295)
                v12 := GenS3{F0: 1, F1: -831, F2: 42}
                if true {
                    switch "a" {
                    case "":
                        v13 := GenS4{F0: true, F1: 2147483647}
                        _ = v13
                    case "\n\t":
                        var v14 int16 = 616
                        _ = v14
                    case "foo bar":
                        v5 = (v1 + 65535)
                    default:
                        if false {
                            switch 42 {
                            case -1:
                                v7 -= ((v7 * (v7 + ((v7 * float32(0.1)) / 1e10))) + v7)
                            case -2147483648:
                                var v15 uint = (v3 + (v3 / ((v3) | 1)))
                                _ = v15
                            case -273:
                                v4 = append(v4, 193)
                            default:
                                if true {
                                    for fi16 := 0; fi16 < 9; fi16++ {
                                        if true {
                                            v7 -= (float32(((3.14159 / ((float64(v7) / 3.14159) - 1e10)) - (1e10 * float64((v7 + float32(0.1)))))) + v7)
                                        }
                                        _ = fi16
                                    }
                                }
                            }
                            _ = len(v2)
                            for fi17 := 0; fi17 < 16; fi17++ {
                                var v18 int = fi17
                                _ = fi17
                                _ = v18
                            }
                        }
                    }
                    if (!true) {
                        switch "\n\t" {
                        case "":
                            var v19 uint64 = 4294967295
                            _ = v19
                        case "foo bar":
                            var v20 uint64 = 1
                            _ = v20
                        case "\n\t":
                            v1 = uint32(v3)
                        default:
                            fmt.Println(v1, v7)
                        }
                        v21 := GenS5{F0: 1000, F1: "a"}
                        v7 -= ((1e10 + v7) * v7)
                        _ = v21
                    } else {
                        v22 := GenS5{F0: 734, F1: ""}
                        if false {
                            if false {
                                var v23 int = 32767
                                if ((true || (255 != 458)) && true) {
                                    for fi24 := 0; fi24 < 11; fi24++ {
                                        for fi25 := 0; fi25 < 2; fi25++ {
                                            v7 = 0.0
                                            fmt.Println(v5)
                                            _ = fi25
                                        }
                                        if (1 <= 496) {
                                            switch v23 {
                                            case 127:
                                                v3 = 645
                                            case 2147483647:
                                                switch 42 {
                                                case -567:
                                                    for fi26 := 0; fi26 < 7; fi26++ {
                                                        if false {
                                                            var v27 uint64 = 42
                                                            _ = v27
                                                        }
                                                        _ = fi26
                                                    }
                                                case 9223372036854775807:
                                                    var v28 string = "foo bar"
                                                    _ = v28
                                                }
                                            case 467:
                                                _ = len(v2)
                                            default:
                                                switch (fi24 | -1) {
                                                case -34:
                                                    fmt.Println(fi24, v7)
                                                default:
                                                    v1 -= v1
                                                }
                                            }
                                            var v29 bool = true
                                            for fi30 := 0; fi30 < 7; fi30++ {
                                                var v31 int8 = -128
                                                if (v29 || v29) {
                                                    v32 := GenS4{F0: v29, F1: v23}
                                                    if ("a" < (("foo bar" + "foo bar") + "foo bar")) {
                                                        v5 = (^(v1))
                                                    }
                                                    _ = v32
                                                } else {
                                                    if v29 {
                                                        switch fi24 {
                                                        case -508:
                                                            if v29 {
                                                                v33 := GenS5{F0: 42, F1: "hello"}
                                                                v2[0] = v1
                                                                v5 = (v5 << ((v3) & 63))
                                                                v3 = v3
                                                                _ = v33
                                                            } else {
                                                                v34 := GenS6{F0: 9223372036854775807, F1: 0, F2: 42}
                                                                fmt.Println(v1)
                                                                v7 = 1e10
                                                                _ = v34
                                                            }
                                                        }
                                                    } else {
                                                        delete(v2, (v5 ^ 1))
                                                    }
                                                }
                                                if (("hello" + "a") > "\n\t") {
                                                    var v35 int64 = int64(fi30)
                                                    _ = v35
                                                }
                                                _ = fi30
                                                _ = v31
                                            }
                                            _ = v29
                                        } else {
                                            if false {
                                                v23 = (v23 | (fi24 / ((646) | 1)))
                                                var v36 int8 = 27
                                                _ = v36
                                            }
                                            if (!false) {
                                                fmt.Println(v7)
                                            } else {
                                                var v37 int = (-(fi24))
                                                for fi38 := 0; fi38 < 9; fi38++ {
                                                    var v39 uint32 = v5
                                                    var v40 uint8 = 91
                                                    v39 -= (v1 << ((v3) & 63))
                                                    _ = fi38
                                                    _ = v39
                                                    _ = v40
                                                }
                                                delete(v2, v1)
                                                _ = v37
                                            }
                                            switch "a" {
                                            case "foo bar":
                                                v7 = v7
                                            }
                                        }
                                        var v41 int16 = 253
                                        _ = fi24
                                        _ = v41
                                    }
                                    if false {
                                        v23 -= v23
                                    }
                                    for fi42 := 0; fi42 < 17; fi42++ {
                                        var v43 uint32 = v1
                                        v23 += (v23 + (fi42 >> ((18446744073709551615) & 63)))
                                        fmt.Println(v3, v3, v43)
                                        _ = fi42
                                        _ = v43
                                    }
                                }
                                _ = v23
                            }
                        }
                        v2[v1] = v5
                        _ = v22
                    }
                    for fi44 := 0; fi44 < 2; fi44++ {
                        switch ("foo bar" + "\n\t") {
                        case "foo bar":
                            if true {
                                for fi45 := 0; fi45 < 14; fi45++ {
                                    var v46 uint = (v3 >> ((365) & 63))
                                    _ = fi45
                                    _ = v46
                                }
                            }
                        }
                        _ = fi44
                    }
                }
                _ = v12
            } else {
                var v47 int16 = -134
                var v48 bool = false
                v5 -= 778
                _ = v47
                _ = v48
            }
        case "a":
            v49 := GenS1{F0: v3, F1: 826, F2: -128}
            _ = v49
        case "":
            var v50 int = 2147483647
            _ = v50
        default:
            v7 = ((v7 - 0.0) - (v7 * float32(1.5)))
        }
        _ = v11
    } else {
        v1 = v5
    }
    v51 := GenS5{F0: 0, F1: ("foo bar" + "")}
    v3 = 68
    v3 -= 65535
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v4)
    v2[2147483647] = 2
    fmt.Println(v3, v3, v7)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
    _ = v7
    _ = v8
}
