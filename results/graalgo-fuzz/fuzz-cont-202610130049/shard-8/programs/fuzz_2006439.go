package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 uint8; F2 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 int64; F2 uint16 }
type GenS5 struct { F0 uint }

func main() {
    var v1 uint64 = 575
    var v2 uint8 = 2
    var v3 int16 = 32767
    v4 := []string{"\n\t", "hello"}
    var v5 int64 = -9223372036854775808
    var v6 int64 = (v5 % ((42) | 1))
    for fi7 := 0; fi7 < 7; fi7++ {
        fmt.Println(v3, fi7, v5)
        _ = fi7
    }
    v1 += 771
    v2 -= (v2 * v2)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        for fi9 := 0; fi9 < 11; fi9++ {
            if true {
                var v10 uint32 = 2
                switch "" {
                case "\n\t":
                    var v11 string = ""
                    _ = v11
                case "":
                    v2 = (v2 / ((v2) | 1))
                case "hello":
                    v4 = append(v4, "hello")
                    var v12 float64 = -2.25
                    var v13 int = fi8
                    _ = v12
                    _ = v13
                default:
                    if ((true || (!(0 != (-(v6))))) || true) {
                        v14 := GenS1{F0: fi9}
                        _ = v14
                    }
                }
                _ = v10
            }
            var v15 uint = 298
            var v16 int32 = 800
            _ = fi9
            _ = v15
            _ = v16
        }
        if false {
            for fi17 := 0; fi17 < 14; fi17++ {
                var v18 uint16 = 0
                v19 := GenS1{F0: (fi8 | (fi17 / (((fi17 / ((fi17) | 1))) | 1)))}
                _ = fi17
                _ = v18
                _ = v19
            }
        } else {
            var v20 uint = 65535
            switch "foo bar" {
            case "a":
                switch (fi8 + (^(fi8))) {
                case 1:
                    var v21 int16 = (v3 | (v3 / (((v3 + v3)) | 1)))
                    _ = v21
                default:
                    v3 += (v3 ^ -1)
                }
            default:
                for fi22 := 0; fi22 < 1; fi22++ {
                    if (int64(fi22) < (v6 << ((0) & 63))) {
                        if false {
                            var v23 uint8 = (v2 + v2)
                            if true {
                                var v24 int8 = 76
                                if ((v20 / (((v20 & v20)) | 1)) != v20) {
                                    v5 -= (^(v6))
                                    var v25 int16 = 65
                                    switch "" {
                                    case "":
                                        v26 := GenS2{F0: (v20 << (((v20 / ((4294967295) | 1))) & 63)), F1: (^(v2)), F2: (fi22 / (((fi22 | fi22)) | 1))}
                                        _ = v26
                                    case "\n\t":
                                        _ = len(v4)
                                    case "hello":
                                        if true {
                                            var v27 float32 = 1e10
                                            if false {
                                                v28 := GenS2{F0: (v20 % (((v20 % ((v20) | 1))) | 1)), F1: v2, F2: (fi22 * fi22)}
                                                v6 = v6
                                                v29 := GenS2{F0: v20, F1: (v23 & v23), F2: fi8}
                                                _ = v28
                                                _ = v29
                                            }
                                            v30 := GenS1{F0: fi22}
                                            v5 += -1
                                            var v31 uint8 = (^(v2))
                                            _ = v27
                                            _ = v30
                                            _ = v31
                                        }
                                    }
                                    _ = v25
                                }
                                switch -615 {
                                case 354:
                                    v32 := GenS3{F0: 42}
                                    _ = v32
                                case -407:
                                    var v33 float64 = 3.14159
                                    _ = v33
                                default:
                                    if len(v4) > 0 {
                                        v4[0] = ("foo bar" + ("a" + ""))
                                    }
                                }
                                _ = v24
                            } else {
                                v1 += (v1 >> (((v20 << (((^(v20))) & 63))) & 63))
                                var v34 bool = (((-2.25 - 0.0) * 1.5) != 0.0)
                                _ = v34
                            }
                            _ = v23
                        } else {
                            v20 = 656
                            var v35 float64 = 3.14159
                            _ = v35
                        }
                        v20 = (v20 & (v20 << ((v20) & 63)))
                        var v36 uint = (v20 / (((v20 ^ (v20 << (((v20 - 65535)) & 63)))) | 1))
                        _ = v36
                    }
                    _ = fi22
                }
            }
            var v37 bool = false
            _ = v20
            _ = v37
        }
        _ = fi8
    }
    switch -799 {
    case 2147483647:
        v4 = append(v4, ("" + "foo bar"))
    default:
        v5 += 485
    }
    switch 1 {
    case 42:
        var v38 uint64 = v1
        _ = v38
    case 512:
        var v39 int8 = -1
        _ = v39
    default:
        for fi40 := 0; fi40 < 5; fi40++ {
            var v41 uint32 = 807
            _ = fi40
            _ = v41
        }
    }
    v5 -= -9223372036854775808
    v2 = v2
    if len(v4) > 0 {
        v4[0] = "foo bar"
    }
    for fi42 := 0; fi42 < 15; fi42++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v43 int32 = -262
        _ = fi42
        _ = v43
    }
    if false {
        v44 := GenS4{F0: -702, F1: v6, F2: 607}
        if (1 != -1) {
            for fi45 := 0; fi45 < 8; fi45++ {
                switch fi45 {
                case 127:
                    v3 = (v3 / ((v3) | 1))
                }
                v3 -= int16(v1)
                v2 -= (v2 + v2)
                _ = fi45
            }
        }
        _ = v44
    } else {
        v46 := GenS5{F0: 4294967295}
        for fi47 := 0; fi47 < 5; fi47++ {
            fmt.Println(v5)
            _ = fi47
        }
        _ = v46
    }
    if (!(65535 <= 42)) {
        var v48 int = -1
        var v49 bool = false
        switch ("hello" + "hello") {
        case "hello":
            var v50 int8 = 44
            _ = v50
        }
        _ = v48
        _ = v49
    } else {
        for fi51 := 0; fi51 < 7; fi51++ {
            v1 += v1
            fmt.Println(v2, v3)
            _ = fi51
        }
        var v52 int64 = (v5 / (((v5 % ((v5) | 1))) | 1))
        _ = len(v4)
        _ = v52
    }
    fmt.Println(v2)
    v5 -= 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
