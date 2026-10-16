package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS6 struct { F0 uint8; F1 uint }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint32; F1 uint; F2 int }
type GenS10 struct { F0 uint64; F1 uint16 }
type GenS11 struct { F0 string; F1 uint8; F2 int32 }

func main() {
    var v1 uint64 = 2
    v2 := []uint32{42}
    v3 := []uint32{65535, 4294967295, 2, 2}
    var v4 uint = 403
    v5 := make(map[int]string)
    for fi6 := 0; fi6 < 4; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            v4 = (^(v4))
            var v8 bool = false
            _ = fi7
            _ = v8
        }
        if true {
            _ = len(v3)
        } else {
            if len(v2) > 0 {
                v2[0] = 479
            }
            v1 += (v1 % ((v1) | 1))
        }
        v1 = v1
        _ = fi6
    }
    v4 = 0
    fmt.Println(v4, v1)
    v9 := GenS1{F0: 1, F1: (3.14159 >= (3.14159 - 0.1))}
    if (true || true) {
        if (v1 > v1) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        if true {
            v4 -= v4
            switch "" {
            case "hello":
                v10 := GenS2{F0: false}
                _ = v10
            case "a":
                if false {
                    var v11 int32 = 1
                    _ = v11
                }
            }
            v1 += (^(v1))
        }
    }
    var v12 uint = 4294967295
    v4 += 18446744073709551615
    v4 = 0
    var v13 uint16 = 10
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v14 uint8 = 128
    var v15 uint16 = (v13 | (v13 ^ 437))
    var v16 uint32 = uint32(v1)
    if (-1 >= 2147483647) {
        for fi17 := 0; fi17 < 4; fi17++ {
            v16 -= (^(v16))
            switch "hello" {
            case "a":
                if false {
                    var v18 int8 = 53
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 16; fi19++ {
                        if ((v16 / ((4294967295) | 1)) < v16) {
                            v20 := v5[fi17]
                            v3 = append(v3, v16)
                            switch v20 {
                            case "foo bar":
                                fmt.Println(v14, v16)
                                var v21 int32 = -1
                                _ = v21
                            case "hello":
                                v15 = uint16(v1)
                            default:
                                fmt.Println(v14, v15)
                            }
                            _ = v20
                        }
                        if true {
                            if false {
                                if false {
                                    var v22 uint = v12
                                    switch "hello" {
                                    case "foo bar":
                                        v23 := GenS1{F0: (v13 / ((v13) | 1)), F1: (false && true)}
                                        _ = v23
                                    case "a":
                                        if false {
                                            v1 = uint64(v13)
                                            if ("\n\t" < "foo bar") {
                                                if (367 > v12) {
                                                    v24 := GenS3{F0: (uint32(v22) >= v16)}
                                                    var v25 int8 = -2
                                                    _ = v24
                                                    _ = v25
                                                } else {
                                                    switch fi17 {
                                                    case -1:
                                                        delete(v5, fi19)
                                                    case 0:
                                                        var v26 uint32 = (v16 >> ((42) & 63))
                                                        _ = v26
                                                    default:
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    }
                                                    v27 := GenS1{F0: 898, F1: false}
                                                    v28 := GenS4{F0: 127}
                                                    _ = v27
                                                    _ = v28
                                                }
                                                switch fi17 {
                                                case 2147483647:
                                                    v5[(fi19 | -9223372036854775808)] = ("foo bar" + (("hello" + "") + "a"))
                                                case 0:
                                                    switch fi17 {
                                                    case 0:
                                                        if false {
                                                            v12 = (^(v12))
                                                            fmt.Println(v12, v16)
                                                            if (!((1e10 - (float32((float64((1e10 + float32(1.5))) * (-2.25 / (float64(-2.25) * 0.1)))) - 1.5)) < ((-2.25 + (0.1 / 1e10)) + 3.14159))) {
                                                                if (("hello" + ("a" + "a")) == "foo bar") {
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    v29 := GenS3{F0: true}
                                                                    _ = v29
                                                                }
                                                                fmt.Println(fi17, fi17)
                                                                v16 -= (v16 + (v16 * 2147483647))
                                                            }
                                                        } else {
                                                            switch 545 {
                                                            case 9223372036854775807:
                                                                if true {
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        if false {
                                                                            v30 := GenS4{F0: -128}
                                                                            v12 = v12
                                                                            if len(v2) > 0 {
                                                                                v2[0] = v16
                                                                            }
                                                                            _ = v30
                                                                        }
                                                                        fmt.Println(v22, fi19)
                                                                    }
                                                                    v22 = 478
                                                                    v22 -= uint(v14)
                                                                } else {
                                                                    v13 += (v15 % ((v15) | 1))
                                                                    var v31 uint64 = 4294967295
                                                                    var v32 uint = (^(v22))
                                                                    _ = v31
                                                                    _ = v32
                                                                }
                                                            case 128:
                                                                _ = len(v5)
                                                            case 515:
                                                                v33 := GenS5{F0: (v14 & v14), F1: 1, F2: 2}
                                                                _ = v33
                                                            default:
                                                                if len(v3) > 0 {
                                                                    v3[0] = (^(v16))
                                                                }
                                                            }
                                                        }
                                                    }
                                                case 718:
                                                    v12 += uint(v14)
                                                default:
                                                    v34 := v5[fi19]
                                                    _ = v34
                                                }
                                            } else {
                                                switch (-(fi19)) {
                                                case 788:
                                                    v14 -= uint8(v13)
                                                case -9223372036854775808:
                                                    v35 := GenS6{F0: v14, F1: (^(v22))}
                                                    _ = v35
                                                case -2147483648:
                                                    if true {
                                                        var v36 string = "foo bar"
                                                        _ = v36
                                                    }
                                                default:
                                                    switch -1 {
                                                    case -1:
                                                        if false {
                                                            var v37 int16 = 42
                                                            v38 := GenS7{F0: 42}
                                                            fmt.Println(v16)
                                                            _ = v37
                                                            _ = v38
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                    case "hello":
                                        v12 = 65535
                                    }
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v39 uint8 = (v14 % ((v14) | 1))
                                        _ = v39
                                    case "hello":
                                        v13 += 65535
                                    case "a":
                                        v4 = 1
                                        if true {
                                            fmt.Println(v13, v13)
                                        }
                                    }
                                    _ = v22
                                }
                            } else {
                                if false {
                                    var v40 int32 = 0
                                    _ = v40
                                }
                                switch "" {
                                case "":
                                    v41 := GenS7{F0: (fi17 * 1)}
                                    _ = v41
                                case "foo bar":
                                    v42 := GenS1{F0: (^(v15)), F1: ((3.14159 == float32(float64(3.14159))) || false)}
                                    _ = v42
                                case "hello":
                                    v43 := GenS8{F0: uint16(v16)}
                                    _ = v43
                                }
                                var v44 bool = false
                                _ = v44
                            }
                            v45 := GenS6{F0: uint8(v1), F1: v12}
                            _ = v45
                        } else {
                            switch 42 {
                            case 680:
                                if (v13 < (v15 >> ((v12) & 63))) {
                                    var v46 bool = true
                                    fmt.Println(v46, v4, v16)
                                    _ = v46
                                } else {
                                    if false {
                                        if (v4 > v12) {
                                            switch (fi17 >> ((42) & 63)) {
                                            case -62:
                                                v47 := GenS4{F0: 1}
                                                _ = v47
                                            case 9223372036854775807:
                                                var v48 uint = uint(v16)
                                                _ = v48
                                            case 167:
                                                _ = len(v5)
                                            default:
                                                v49 := GenS3{F0: false}
                                                _ = v49
                                            }
                                            v50 := GenS3{F0: (("foo bar" <= "") && true)}
                                            delete(v5, -9223372036854775808)
                                            _ = v50
                                        } else {
                                            v51 := GenS6{F0: v14, F1: (v12 ^ v12)}
                                            _ = v51
                                        }
                                        if true {
                                            fmt.Println(fi19)
                                            switch fi19 {
                                            case -2147483648:
                                                var v52 float64 = (0.0 / 1e10)
                                                _ = v52
                                            }
                                        } else {
                                            v14 = uint8(v1)
                                            v53 := GenS9{F0: (^(v16)), F1: uint(v14), F2: -2147483648}
                                            switch (fi17 / (((fi19 & fi17)) | 1)) {
                                            case 409:
                                                if true {
                                                    var v54 uint16 = (v15 / (((v15 % (((v13 | 65535)) | 1))) | 1))
                                                    fmt.Println(v4)
                                                    v55 := GenS10{F0: v1, F1: (v13 ^ (^(v13)))}
                                                    var v56 string = ("hello" + "")
                                                    _ = v54
                                                    _ = v55
                                                    _ = v56
                                                }
                                            case -446:
                                                fmt.Println(v12)
                                                switch int(v16) {
                                                case -296:
                                                    switch "" {
                                                    case "\n\t":
                                                        fmt.Println(v1, v14, v13)
                                                    }
                                                }
                                            case 127:
                                                fmt.Println(v14, v1)
                                            default:
                                                v13 += (v15 / ((960) | 1))
                                            }
                                            _ = v53
                                        }
                                    }
                                }
                            case -312:
                                v12 -= uint(v14)
                            }
                            switch "" {
                            case "\n\t":
                                v13 += (v15 % ((v13) | 1))
                            case "a":
                                v1 -= 4294967295
                            case "":
                                var v57 uint64 = v1
                                _ = v57
                            default:
                                var v58 int = fi19
                                _ = v58
                            }
                        }
                        var v59 int8 = -1
                        _ = fi19
                        _ = v59
                    }
                    if (true && ((!true) && true)) {
                        switch 1 {
                        case 48:
                            if (223 == v14) {
                                var v60 uint8 = 127
                                v14 = uint8(v12)
                                _ = v60
                            }
                        case 1:
                            var v61 uint32 = 4294967295
                            _ = v61
                        case 4294967295:
                            v62 := GenS1{F0: v15, F1: false}
                            _ = v62
                        default:
                            v63 := GenS11{F0: "a", F1: v14, F2: -63}
                            _ = v63
                        }
                        v14 = (^(v14))
                    }
                }
            case "hello":
                var v64 float32 = (1e10 - 0.0)
                _ = v64
            case "":
                v2 = append(v2, 65535)
            }
            _ = fi17
        }
    }
    v15 = (v13 ^ v13)
    var v65 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v65:", v65)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v9
}
