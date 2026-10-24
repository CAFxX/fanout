package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint64; F2 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 int8; F1 int8; F2 int16 }
type GenS6 struct { F0 uint64; F1 uint }
type GenS7 struct { F0 uint64; F1 int32; F2 bool }

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 42
    var v3 uint8 = 236
    v4 := []uint32{2147483647, 169, 42, 62}
    fmt.Println(v3)
    switch ("a" + (("\n\t" + "\n\t") + "foo bar")) {
    case "hello":
        v2 += 1
    case "foo bar":
        v5 := GenS1{F0: -1, F1: 4294967295}
        _ = v5
    case "":
        fmt.Println(v3)
    default:
        v2 += 9223372036854775807
    }
    for fi6 := 0; fi6 < 6; fi6++ {
        v2 = 9223372036854775807
        _ = fi6
    }
    if (0 > 875) {
        if true {
            for fi7 := 0; fi7 < 5; fi7++ {
                v2 -= v2
                _ = fi7
            }
            for fi8 := 0; fi8 < 20; fi8++ {
                var v9 uint16 = 255
                _ = fi8
                _ = v9
            }
            var v10 int = 786
            _ = v10
        }
    } else {
        var v11 int64 = v2
        var v12 uint32 = 1
        _ = v11
        _ = v12
    }
    if true {
        if true {
            var v13 int16 = 32767
            _ = v13
        } else {
            v14 := v1["hello"]
            for fi15 := 0; fi15 < 17; fi15++ {
                var v16 int = (fi15 / (((fi15 ^ -983)) | 1))
                for fi17 := 0; fi17 < 11; fi17++ {
                    v18 := GenS1{F0: 0, F1: 2}
                    var v19 string = ((((("hello" + v14) + "hello") + v14) + "\n\t") + (v14 + v14))
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = fi15
                _ = v16
            }
            _ = v14
        }
        v20 := GenS2{F0: false, F1: 42, F2: 22}
        var v21 uint64 = 185
        _ = v20
        _ = v21
    } else {
        if ((v2 << ((9) & 63)) >= -1) {
            var v22 uint64 = 42
            _ = v22
        }
        for fi23 := 0; fi23 < 5; fi23++ {
            v24 := v1["\n\t"]
            var v25 float32 = 1e10
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    fmt.Println(v3)
    if (100 < 42) {
        switch ("hello" + "hello") {
        case "":
            var v26 float32 = (0.1 / 0.1)
            _ = v26
        default:
            var v27 float32 = 0.1
            _ = v27
        }
        var v28 int = -1
        _ = v28
    }
    if false {
        var v29 int32 = -2147483648
        _ = v29
    } else {
        var v30 int16 = -769
        v30 = v30
        _ = v30
    }
    if true {
        v3 = v3
        v31 := GenS3{F0: 191}
        switch ("foo bar" + ((("hello" + "foo bar") + ("a" + ("hello" + "hello"))) + "")) {
        case "hello":
            if (true || true) {
                delete(v1, "\n\t")
                for fi32 := 0; fi32 < 9; fi32++ {
                    v33 := GenS3{F0: 163}
                    if true {
                        delete(v1, "hello")
                        v3 -= (v3 ^ (v3 & (v3 >> ((1) & 63))))
                        var v34 float32 = (3.14159 * (0.0 - 0.1))
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 10; fi35++ {
                        v4 = append(v4, 0)
                        var v36 uint8 = 143
                        v4 = append(v4, 2147483648)
                        _ = fi35
                        _ = v36
                    }
                    _ = fi32
                    _ = v33
                }
                switch "foo bar" {
                case "a":
                    if (!((!false) || false)) {
                        switch 32768 {
                        case -634:
                            v37 := GenS1{F0: 42, F1: 42}
                            _ = v37
                        default:
                            for fi38 := 0; fi38 < 17; fi38++ {
                                if (!true) {
                                    delete(v1, "\n\t")
                                }
                                delete(v1, "hello")
                                _ = fi38
                            }
                        }
                    } else {
                        _ = len(v4)
                        if false {
                            v2 -= v2
                            var v39 int32 = int32(v3)
                            _ = v39
                        }
                        v40 := GenS4{F0: 255, F1: -448}
                        _ = v40
                    }
                case "hello":
                    for fi41 := 0; fi41 < 17; fi41++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi41
                    }
                case "\n\t":
                    v3 = (v3 * v3)
                default:
                    switch 32767 {
                    case -2147483648:
                        v42 := GenS2{F0: true, F1: 18446744073709551615, F2: -1}
                        _ = v42
                    default:
                        v2 -= (-(v2))
                    }
                }
            }
        }
        _ = v31
    } else {
        var v43 int32 = 0
        _ = v43
    }
    if true {
        v2 -= -348
    }
    v3 -= (v3 * (v3 ^ v3))
    switch ("foo bar" + ("hello" + ("a" + "hello"))) {
    case "":
        var v44 string = "foo bar"
        _ = v44
    case "foo bar":
        var v45 bool = true
        _ = v45
    case "\n\t":
        if ((-139 < 127) || true) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            delete(v1, ("a" + "hello"))
            if ((false || false) && false) {
                _ = len(v4)
                var v46 int32 = 827
                for fi47 := 0; fi47 < 12; fi47++ {
                    if (65535 <= 42) {
                        var v48 float64 = (float64(1e10) - -2.25)
                        _ = v48
                    }
                    var v49 int8 = -128
                    for fi50 := 0; fi50 < 10; fi50++ {
                        v2 += 388
                        var v51 int8 = v49
                        var v52 int8 = v51
                        var v53 string = ("\n\t" + (("" + ("\n\t" + "hello")) + (("" + "\n\t") + "a")))
                        _ = fi50
                        _ = v51
                        _ = v52
                        _ = v53
                    }
                    _ = fi47
                    _ = v49
                }
                _ = v46
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                fmt.Println(v2, v2)
                switch 2147483647 {
                case 32767:
                    switch 128 {
                    case 127:
                        var v54 uint32 = 2147483648
                        _ = v54
                    case 66:
                        if true {
                            v55 := GenS2{F0: false, F1: 1, F2: 1}
                            _ = v55
                        } else {
                            if (false && ("" >= "")) {
                                v3 += v3
                            }
                        }
                    }
                case 2147483647:
                    if (true && true) {
                        v56 := GenS5{F0: -1, F1: 0, F2: -1}
                        for fi57 := 0; fi57 < 18; fi57++ {
                            for fi58 := 0; fi58 < 13; fi58++ {
                                v2 = v2
                                fmt.Println(v2, fi57, v3)
                                if false {
                                    var v59 int32 = 2147483647
                                    fmt.Println(v59)
                                    _ = v59
                                }
                                _ = fi58
                            }
                            for fi60 := 0; fi60 < 13; fi60++ {
                                var v61 float64 = 0.0
                                var v62 float64 = v61
                                _ = fi60
                                _ = v61
                                _ = v62
                            }
                            v2 -= 100
                            _ = fi57
                        }
                        v2 -= v2
                        _ = v56
                    } else {
                        if false {
                            var v63 int64 = -9223372036854775808
                            for fi64 := 0; fi64 < 16; fi64++ {
                                var v65 float32 = 0.1
                                _ = fi64
                                _ = v65
                            }
                            var v66 float64 = -2.25
                            _ = v63
                            _ = v66
                        }
                    }
                case 743:
                    v67 := GenS6{F0: 18446744073709551615, F1: 4294967295}
                    _ = v67
                }
            }
        } else {
            v3 = 245
        }
    default:
        for fi68 := 0; fi68 < 14; fi68++ {
            v69 := GenS7{F0: 2, F1: 1, F2: false}
            v70 := GenS7{F0: 2, F1: 42, F2: true}
            _ = fi68
            _ = v69
            _ = v70
        }
    }
    v1["hello"] = "foo bar"
    v71 := GenS5{F0: 107, F1: 100, F2: -1}
    var v72 int16 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v71
    _ = v72
}
