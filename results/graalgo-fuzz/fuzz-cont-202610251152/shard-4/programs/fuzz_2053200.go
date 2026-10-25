package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 uint }
func genFn0(p0 int32) int32 {
    p0 += (-(p0))
    var v4 uint64 = 18446744073709551615
    _ = p0
    _ = v4
    return (^(p0))
}

func main() {
    var v1 uint8 = 2
    v2 := []uint32{0}
    v3 := make(map[int]bool)
    v1 -= (^(v1))
    var v5 string = ((("\n\t" + "a") + "hello") + ("foo bar" + "foo bar"))
    for fi6 := 0; fi6 < 17; fi6++ {
        if false {
            if len(v2) > 0 {
                v2[0] = uint32(v1)
            }
            v7 := GenS1{F0: uint16(v1), F1: (true && false), F2: 65535}
            _ = v7
        } else {
            v1 += (v1 - v1)
            switch fi6 {
            case 4294967295:
                switch v5 {
                case "a":
                    if ((false || (!(1 == 581))) && (42 != 2)) {
                        var v8 bool = (!true)
                        if v8 {
                            var v9 int32 = 1
                            _ = v9
                        } else {
                            if false {
                                var v10 uint = 18446744073709551615
                                v11 := GenS1{F0: 1, F1: ((0.1 - (((1.5 + (-2.25 / 1.0)) - 0.1) * 3.14159)) != -2.25), F2: v10}
                                _ = v10
                                _ = v11
                            } else {
                                fmt.Println(v5, v5, v8)
                            }
                            v5 = ""
                            _ = len(v3)
                        }
                        _ = v8
                    }
                default:
                    var v12 int = fi6
                    _ = v12
                }
            default:
                if len(v2) > 0 {
                    v2[0] = 1
                }
            }
            _ = len(v2)
        }
        fmt.Println(v1, v5)
        var v13 int16 = int16(fi6)
        _ = fi6
        _ = v13
    }
    v5 = (v5 + "hello")
    v14 := GenS1{F0: 750, F1: (166 != 32767), F2: 203}
    v15 := GenS2{F0: -1, F1: 4294967295}
    v16 := GenS1{F0: 0, F1: false, F2: 941}
    v1 += v1
    if (v1 >= v1) {
        for fi17 := 0; fi17 < 13; fi17++ {
            v5 = v5
            _ = fi17
        }
        var v18 int32 = -2147483648
        fmt.Println(v1, v5)
        _ = v18
    } else {
        fmt.Println(v1, v5)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v19 int = 42
    switch v5 {
    case "":
        v1 -= (v1 << ((13) & 63))
    case "hello":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "a":
        switch (v19 ^ (v19 + v19)) {
        case 128:
            switch (v5 + "\n\t") {
            case "hello":
                v19 = (v19 * v19)
            case "\n\t":
                if true {
                    if true {
                        if ((1.5 - 3.14159) != 3.14159) {
                            var v20 int16 = 0
                            if false {
                                v5 = v5
                            }
                            _ = v20
                        } else {
                            if (((53 > v1) && false) && false) {
                                _ = len(v2)
                            } else {
                                v21 := GenS1{F0: 74, F1: (-1 > -555), F2: 713}
                                if true {
                                    var v22 int = (v19 % ((1) | 1))
                                    for fi23 := 0; fi23 < 19; fi23++ {
                                        var v24 uint = 2
                                        _ = fi23
                                        _ = v24
                                    }
                                    _ = v22
                                }
                                var v25 int16 = 251
                                _ = v21
                                _ = v25
                            }
                        }
                    }
                } else {
                    for fi26 := 0; fi26 < 5; fi26++ {
                        for fi27 := 0; fi27 < 19; fi27++ {
                            var v28 uint32 = 2
                            var v29 bool = (!true)
                            v19 += v19
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                        for fi30 := 0; fi30 < 20; fi30++ {
                            if (-2.25 == ((0.0 - 1e10) / -2.25)) {
                                var v31 uint8 = 0
                                _ = v31
                            } else {
                                var v32 bool = (false || true)
                                fmt.Println(v1, v19)
                                var v33 uint8 = v1
                                v2 = append(v2, 4294967295)
                                var v34 float32 = ((1.5 * 1.5) / 0.1)
                                _ = v32
                                _ = v33
                                _ = v34
                            }
                            _ = fi30
                        }
                        _ = fi26
                    }
                }
            case "a":
                v35 := GenS1{F0: 0, F1: ((^(v19)) > (v19 / (((v19 % (((v19 ^ v19)) | 1))) | 1))), F2: 1}
                _ = v35
            }
        case 32767:
            var v36 int = (v19 % ((v19) | 1))
            _ = v36
        default:
            for fi37 := 0; fi37 < 6; fi37++ {
                if len(v2) > 0 {
                    v2[0] = 1
                }
                _ = fi37
            }
        }
    }
    v1 += v1
    for fi38 := 0; fi38 < 7; fi38++ {
        fmt.Println(fi38, v19, v1)
        v1 -= (v1 % ((v1) | 1))
        _ = fi38
    }
    fmt.Println(v19, v5)
    switch v5 {
    case "":
        v39 := GenS2{F0: 1, F1: 2}
        _ = v39
    case "hello":
        for fi40 := 0; fi40 < 19; fi40++ {
            if (-529 > 31) {
                for fi41 := 0; fi41 < 10; fi41++ {
                    var v42 float64 = (0.1 - (3.14159 * 1.5))
                    var v43 int32 = -1
                    v19 = (v19 ^ (fi40 / (((-(fi41))) | 1)))
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                var v44 uint32 = 766
                _ = v44
            } else {
                v1 = (v1 / ((128) | 1))
                for fi45 := 0; fi45 < 20; fi45++ {
                    if false {
                        fmt.Println(fi40, fi40, fi45)
                        if true {
                            v46 := GenS1{F0: 339, F1: true, F2: 65535}
                            _ = v46
                        }
                        v5 = v5
                    }
                    if true {
                        _ = len(v3)
                        v1 -= 255
                        var v47 int64 = 849
                        fmt.Println(v47, fi45)
                        fmt.Println(fi40, v47, v1)
                        _ = v47
                    }
                    switch fi40 {
                    case -417:
                        if true {
                            if (-2147483648 < -1) {
                                v1 -= v1
                            }
                            switch (fi45 >> ((4294967295) & 63)) {
                            case -9223372036854775808:
                                v48 := GenS3{F0: true, F1: ("a" + "hello")}
                                _ = v48
                            case 32767:
                                v5 = ("hello" + "a")
                            case 32768:
                                fmt.Println(v1, fi45, v1)
                            }
                            v1 = v1
                        } else {
                            v49 := GenS2{F0: 41, F1: 903}
                            _ = len(v3)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v49
                        }
                    default:
                        switch 2147483647 {
                        case -9223372036854775808:
                            var v50 uint8 = (v1 / ((v1) | 1))
                            _ = v50
                        case 127:
                            v51 := GenS4{F0: 42}
                            _ = v51
                        case 4294967295:
                            fmt.Println(v19, v5)
                        default:
                            v52 := GenS4{F0: 2}
                            _ = v52
                        }
                    }
                    var v53 float32 = 1e10
                    var v54 bool = (65535 != 718)
                    v53 = 1.5
                    if (-89 < 127) {
                        var v55 int = (fi40 ^ 128)
                        _ = v55
                    } else {
                        var v56 float32 = 3.14159
                        _ = v56
                    }
                    _ = fi45
                    _ = v53
                    _ = v54
                }
                fmt.Println(fi40)
            }
            for fi57 := 0; fi57 < 13; fi57++ {
                v58 := v3[(fi57 << ((42) & 63))]
                var v59 bool = (666 < 42)
                _ = fi57
                _ = v58
                _ = v59
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi40
        }
    default:
        v60 := GenS2{F0: -128, F1: 878}
        _ = v60
    }
    if true {
        v61 := GenS3{F0: true, F1: (v5 + v5)}
        fmt.Println(v1, v19, v19)
        _ = v61
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v5
}
