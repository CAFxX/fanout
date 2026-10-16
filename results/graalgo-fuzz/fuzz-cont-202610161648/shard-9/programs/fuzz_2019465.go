package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int32; F2 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 uint32; F2 string }
type GenS5 struct { F0 int32; F1 uint; F2 string }
type GenS6 struct { F0 uint16 }

func main() {
    var v1 int = -1
    var v2 float32 = 1.5
    v3 := []uint32{65535}
    var v4 int8 = 61
    var v5 uint32 = 1
    if true {
        v6 := GenS1{F0: "a"}
        v7 := GenS1{F0: "foo bar"}
        for fi8 := 0; fi8 < 15; fi8++ {
            v1 += (fi8 ^ 32768)
            v4 += 42
            v9 := GenS2{F0: 42, F1: -575, F2: false}
            _ = fi8
            _ = v9
        }
        _ = v6
        _ = v7
    }
    v4 = (v4 | v4)
    v2 = float32(1e10)
    v5 = 0
    if true {
        v4 -= v4
        if false {
            v5 -= (^(v5))
            for fi10 := 0; fi10 < 6; fi10++ {
                v1 += fi10
                if (v2 == float32(((1.5 + ((0.0 + 3.14159) + (float64(((float32(-2.25) + v2) + (1.5 - float32((float64((v2 - v2)) + 1e10))))) / -2.25))) + float64(((((v2 * float32(0.0)) * v2) - float32(1e10)) / (((v2 + v2) - float32(1e10)) + (((v2 - (v2 / (v2 + v2))) / v2) / (((v2 * (v2 / 1.5)) * 3.14159) + v2)))))))) {
                    for fi11 := 0; fi11 < 11; fi11++ {
                        v4 = v4
                        _ = fi11
                    }
                }
                var v12 int8 = v4
                _ = fi10
                _ = v12
            }
            for fi13 := 0; fi13 < 2; fi13++ {
                if true {
                    for fi14 := 0; fi14 < 10; fi14++ {
                        fmt.Println(fi13, v1, v1)
                        v1 += (fi13 % ((v1) | 1))
                        v3 = append(v3, v5)
                        _ = fi14
                    }
                    var v15 uint32 = v5
                    _ = v15
                }
                _ = fi13
            }
        } else {
            for fi16 := 0; fi16 < 13; fi16++ {
                var v17 string = ("foo bar" + ("hello" + (("\n\t" + "foo bar") + "a")))
                v18 := GenS2{F0: -1, F1: -134, F2: false}
                _ = fi16
                _ = v17
                _ = v18
            }
            var v19 uint16 = 0
            _ = v19
        }
        switch (v1 % (((v1 & v1)) | 1)) {
        case 1:
            switch v1 {
            case 42:
                fmt.Println(v1, v4, v1)
            case 4294967295:
                switch 32767 {
                case 0:
                    v20 := GenS2{F0: 42, F1: -873, F2: ("hello" > "\n\t")}
                    _ = v20
                default:
                    _ = len(v3)
                }
            case -9223372036854775808:
                if (42 != 1) {
                    for fi21 := 0; fi21 < 8; fi21++ {
                        v4 = 127
                        for fi22 := 0; fi22 < 7; fi22++ {
                            var v23 uint64 = 792
                            if len(v3) > 0 {
                                v3[0] = 300
                            }
                            var v24 int64 = 9223372036854775807
                            switch "\n\t" {
                            case "hello":
                                if true {
                                    var v25 uint16 = 65535
                                    fmt.Println(v2)
                                    fmt.Println(v24)
                                    _ = v25
                                } else {
                                    var v26 uint32 = 2
                                    v27 := GenS1{F0: "hello"}
                                    _ = v26
                                    _ = v27
                                }
                                var v28 int32 = -2147483648
                                switch (fi21 / ((104) | 1)) {
                                case -643:
                                    v29 := GenS3{F0: v4}
                                    _ = v29
                                case -261:
                                    if (42 >= 0) {
                                        var v30 uint = 255
                                        fmt.Println(v23)
                                        v3 = append(v3, (v5 - v5))
                                        _ = v30
                                    } else {
                                        if true {
                                            fmt.Println(fi21, fi21)
                                        } else {
                                            v31 := GenS4{F0: (v23 | (v23 * v23)), F1: 332, F2: "\n\t"}
                                            var v32 float32 = v2
                                            _ = v31
                                            _ = v32
                                        }
                                    }
                                default:
                                    v24 = (v24 & v24)
                                }
                                _ = v28
                            case "\n\t":
                                v33 := GenS5{F0: 284, F1: 1, F2: "\n\t"}
                                _ = v33
                            case "":
                                fmt.Println(fi21)
                                var v34 uint = 65535
                                _ = v34
                            default:
                                v3 = append(v3, 2147483647)
                            }
                            fmt.Println(v4, v1)
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        _ = fi21
                    }
                } else {
                    var v35 uint64 = 42
                    var v36 int8 = (v4 - (v4 ^ (v4 | (v4 + (v4 ^ 51)))))
                    if (false || true) {
                        for fi37 := 0; fi37 < 7; fi37++ {
                            fmt.Println(v1)
                            _ = fi37
                        }
                    } else {
                        v38 := GenS3{F0: (v36 * v36)}
                        _ = v38
                    }
                    _ = v35
                    _ = v36
                }
            default:
                if (4294967295 > 962) {
                    _ = len(v3)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    for fi39 := 0; fi39 < 7; fi39++ {
                        _ = len(v3)
                        var v40 float32 = v2
                        _ = fi39
                        _ = v40
                    }
                }
            }
        case 42:
            for fi41 := 0; fi41 < 1; fi41++ {
                var v42 bool = (2 >= 18446744073709551615)
                _ = fi41
                _ = v42
            }
        case 591:
            fmt.Println(v1, v2)
        }
    }
    var v43 uint16 = 32767
    for fi44 := 0; fi44 < 16; fi44++ {
        v1 = fi44
        v45 := GenS2{F0: 0, F1: -2147483648, F2: true}
        _ = fi44
        _ = v45
    }
    for fi46 := 0; fi46 < 9; fi46++ {
        if true {
            if (481 >= 2147483647) {
                v1 = 4294967295
                switch fi46 {
                case 731:
                    var v47 uint16 = v43
                    _ = v47
                default:
                    for fi48 := 0; fi48 < 6; fi48++ {
                        if (-351 >= 2147483647) {
                            fmt.Println(v43, v2)
                            if true {
                                if len(v3) > 0 {
                                    v3[0] = v5
                                }
                            } else {
                                _ = len(v3)
                                if (!true) {
                                    if true {
                                        v1 = (-(fi46))
                                    }
                                    v43 = v43
                                }
                                var v49 uint16 = 0
                                _ = v49
                            }
                            v43 += v43
                            if len(v3) > 0 {
                                v3[0] = v5
                            }
                            v43 = (^(v43))
                        }
                        _ = fi48
                    }
                }
                _ = len(v3)
            } else {
                if false {
                    if false {
                        v5 = uint32(v43)
                        switch "" {
                        case "hello":
                            v3 = append(v3, v5)
                        case "foo bar":
                            var v50 int = 462
                            _ = v50
                        default:
                            for fi51 := 0; fi51 < 16; fi51++ {
                                v52 := GenS1{F0: "\n\t"}
                                switch "hello" {
                                case "hello":
                                    v5 = (v5 - v5)
                                }
                                if false {
                                    fmt.Println(v2, fi51, fi51)
                                }
                                _ = fi51
                                _ = v52
                            }
                        }
                        fmt.Println(fi46, v5, v2)
                    } else {
                        for fi53 := 0; fi53 < 18; fi53++ {
                            var v54 int16 = 0
                            _ = fi53
                            _ = v54
                        }
                        if true {
                            if len(v3) > 0 {
                                v3[0] = v5
                            }
                            var v55 uint8 = 246
                            fmt.Println(v2, v43)
                            _ = v55
                        }
                    }
                    var v56 uint = 2
                    if ((128 > 127) && false) {
                        if true {
                            switch "\n\t" {
                            case "\n\t":
                                fmt.Println(v43, v4)
                            }
                            var v57 int8 = 127
                            if true {
                                for fi58 := 0; fi58 < 5; fi58++ {
                                    fmt.Println(v2)
                                    _ = fi58
                                }
                            }
                            _ = v57
                        }
                    } else {
                        var v59 uint16 = v43
                        _ = v59
                    }
                    _ = v56
                } else {
                    if ((0 > 2) && true) {
                        if (true && ((v2 / float32(float64(-2.25))) < 0.1)) {
                            var v60 float32 = v2
                            v61 := GenS6{F0: v43}
                            var v62 int64 = -9223372036854775808
                            _ = v60
                            _ = v61
                            _ = v62
                        } else {
                            var v63 int64 = -665
                            _ = v63
                        }
                        for fi64 := 0; fi64 < 8; fi64++ {
                            v4 += v4
                            fmt.Println(v4, fi64)
                            _ = fi64
                        }
                    }
                }
            }
        } else {
            v5 -= v5
            if len(v3) > 0 {
                v3[0] = (v5 / ((v5) | 1))
            }
            if len(v3) > 0 {
                v3[0] = v5
            }
        }
        var v65 uint64 = 4294967295
        _ = fi46
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
}
