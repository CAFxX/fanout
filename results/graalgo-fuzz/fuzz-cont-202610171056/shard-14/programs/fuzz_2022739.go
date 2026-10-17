package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 bool; F1 bool; F2 uint16 }
func genFn0(p0 uint8) bool {
    if (2 < 706) {
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 int16 = 1
            v9 := GenS1{F0: 2147483648, F1: 555}
            if (!true) {
                v8 += v8
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        p0 -= 14
        for fi10 := 0; fi10 < 20; fi10++ {
            fmt.Println(p0, p0)
            _ = fi10
        }
    } else {
        var v11 int16 = -353
        _ = v11
    }
    p0 = 2
    var v12 uint = 65535
    v13 := []int{-9223372036854775808, 151, 32768, 9223372036854775807}
    var v14 int = 127
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return true
}
func genFn1(p0 string) uint32 {
    for fi15 := 0; fi15 < 9; fi15++ {
        v16 := []int{(fi15 & (fi15 * 42)), (fi15 | fi15), fi15, fi15}
        _ = fi15
        _ = v16
    }
    p0 = p0
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 int32 = -2147483648
        _ = fi17
        _ = v18
    }
    v19 := GenS1{F0: 672, F1: 18446744073709551615}
    if false {
        var v20 uint16 = 573
        var v21 uint16 = v20
        _ = v20
        _ = v21
    } else {
        switch ((("hello" + (p0 + ((p0 + (("a" + p0) + (p0 + (p0 + "\n\t")))) + p0))) + (p0 + p0)) + p0) {
        case "hello":
            p0 = "\n\t"
        case "\n\t":
            for fi22 := 0; fi22 < 7; fi22++ {
                for fi23 := 0; fi23 < 14; fi23++ {
                    var v24 int64 = 338
                    _ = fi23
                    _ = v24
                }
                if true {
                    fmt.Println(p0, fi22)
                }
                _ = fi22
            }
        case "a":
            if (false || false) {
                v25 := []int{2147483647, 1}
                _ = v25
            } else {
                p0 = "a"
                var v26 int16 = 32767
                _ = v26
            }
        }
        v27 := GenS1{F0: 2147483647, F1: 42}
        _ = v27
    }
    _ = p0
    _ = v19
    return 0
}

func main() {
    v1 := []int{32768, -689, 61}
    var v2 int = 111
    var v3 int32 = 625
    v4 := make(map[uint32]uint32)
    var v5 uint = 1
    v6 := []int{(-(v2)), (v2 & (v2 & 127))}
    var v28 int32 = -1
    for fi29 := 0; fi29 < 13; fi29++ {
        if true {
            var v30 int32 = (v28 << ((v5) & 63))
            _ = v30
        } else {
            v31 := GenS2{F0: 492, F1: 4294967295}
            if false {
                switch "foo bar" {
                case "foo bar":
                    _ = len(v1)
                }
                v32 := GenS1{F0: 2147483648, F1: v5}
                v3 = (v28 % ((909) | 1))
                _ = v32
            }
            _ = v31
        }
        _ = len(v1)
        _ = fi29
    }
    switch v2 {
    case -9223372036854775808:
        v5 -= v5
    }
    var v33 string = ("\n\t" + "foo bar")
    if true {
        if false {
            fmt.Println(v2)
            v28 = (v3 * v3)
        } else {
            fmt.Println(v33, v5, v28)
            var v34 int = v2
            _ = v34
        }
        var v35 int = 4294967295
        var v36 int = (v2 & 127)
        _ = v35
        _ = v36
    }
    v2 = v2
    fmt.Println(v3, v28, v28)
    fmt.Println(v33)
    for fi37 := 0; fi37 < 9; fi37++ {
        for fi38 := 0; fi38 < 7; fi38++ {
            if true {
                switch 585 {
                case 2147483647:
                    v1 = append(v1, 495)
                default:
                    if len(v6) > 0 {
                        v6[0] = (v2 >> (((^(v5))) & 63))
                    }
                }
                v39 := GenS3{F0: 2147483647, F1: v5}
                var v40 int8 = 0
                if false {
                    fmt.Println(v5, v2)
                    v2 = (v2 / (((v2 * v2)) | 1))
                } else {
                    v41 := GenS4{F0: 7}
                    _ = v41
                }
                var v42 uint32 = genFn1((((v33 + v33) + v33) + "foo bar"))
                v43 := GenS3{F0: (v28 + (-(v3))), F1: v5}
                var v44 int64 = 1
                _ = v39
                _ = v40
                _ = v42
                _ = v43
                _ = v44
            } else {
                if false {
                    v33 = (v33 + ("a" + v33))
                    var v45 int8 = 42
                    var v46 int64 = 0
                    _ = v45
                    _ = v46
                } else {
                    v28 = v28
                    v47 := GenS5{F0: (v2 << (((v5 - (v5 ^ v5))) & 63)), F1: (127 < -102)}
                    switch int(v28) {
                    case 358:
                        var v48 int32 = 462
                        _ = v48
                    default:
                        v2 = v2
                    }
                    _ = v47
                }
            }
            v49 := GenS3{F0: (v28 >> (((^(v5))) & 63)), F1: (v5 / ((255) | 1))}
            if (127 > 1) {
                switch ("foo bar" + (v33 + (((((("a" + "hello") + "hello") + "hello") + v33) + ("" + v33)) + "hello"))) {
                case "foo bar":
                    delete(v4, 1)
                case "\n\t":
                    v50 := GenS3{F0: (v28 ^ (v3 / ((-2147483648) | 1))), F1: v5}
                    _ = v50
                case "":
                    switch v33 {
                    case "":
                        switch (fi38 & (fi37 + (fi38 / ((4294967295) | 1)))) {
                        case -318:
                            fmt.Println(v28)
                        default:
                            if false {
                                switch fi38 {
                                case 42:
                                    if (65535 > 2147483647) {
                                        var v51 int = fi38
                                        _ = v51
                                    }
                                }
                                if (!(!true)) {
                                    v4[42] = 2
                                } else {
                                    var v52 float32 = float32(1.5)
                                    var v53 uint8 = 42
                                    v54 := GenS4{F0: 1}
                                    _ = v52
                                    _ = v53
                                    _ = v54
                                }
                                v3 = 0
                            }
                            if false {
                                delete(v4, 430)
                                _ = len(v6)
                                var v55 int32 = -2147483648
                                var v56 uint16 = 255
                                var v57 uint8 = 128
                                _ = v55
                                _ = v56
                                _ = v57
                            }
                            v6 = append(v6, fi38)
                            if false {
                                fmt.Println(fi37, fi37, v33)
                                v28 -= (v28 << (((v5 - 42)) & 63))
                            } else {
                                _ = len(v4)
                                if false {
                                    switch v33 {
                                    case "a":
                                        if (!true) {
                                            v33 = "foo bar"
                                        }
                                    }
                                    var v58 int64 = -67
                                    _ = v58
                                }
                            }
                            var v59 uint8 = 127
                            _ = v59
                        }
                    case "foo bar":
                        v60 := v4[40]
                        _ = v60
                    case "hello":
                        if true {
                            var v61 int16 = 569
                            _ = v61
                        } else {
                            var v62 int8 = 0
                            switch "foo bar" {
                            case "a":
                                if false {
                                    fmt.Println(v28)
                                }
                            case "\n\t":
                                if false {
                                    v6 = append(v6, v2)
                                    v5 += v5
                                } else {
                                    delete(v4, 618)
                                    v2 -= v2
                                    if len(v6) > 0 {
                                        v6[0] = fi37
                                    }
                                }
                            case "foo bar":
                                v33 = v33
                            default:
                                v33 = "foo bar"
                            }
                            switch fi37 {
                            case 32767:
                                v33 = v33
                            }
                            var v63 uint32 = 190
                            if (uint(v63) > (v5 % (((v5 + v5)) | 1))) {
                                var v64 int64 = -715
                                fmt.Println(fi38, fi38, v28)
                                var v65 int16 = 1000
                                _ = v64
                                _ = v65
                            }
                            switch v33 {
                            case "a":
                                var v66 uint16 = 32767
                                _ = v66
                            default:
                                v1 = append(v1, fi37)
                            }
                            fmt.Println(v28, v2, v5)
                            _ = v62
                            _ = v63
                        }
                    }
                default:
                    if false {
                        var v67 float64 = (-2.25 + (3.14159 - -2.25))
                        v68 := GenS2{F0: 18446744073709551615, F1: 111}
                        var v69 int32 = (v28 - (v28 * (v28 % ((-1) | 1))))
                        _ = v67
                        _ = v68
                        _ = v69
                    } else {
                        v28 = 741
                    }
                }
                var v70 uint16 = 32767
                v71 := GenS4{F0: 2}
                _ = v70
                _ = v71
            } else {
                v72 := GenS6{F0: (0 > 42), F1: (true && false), F2: 32767}
                if (v33 != ("foo bar" + "hello")) {
                    var v73 uint16 = 0
                    var v74 uint64 = 542
                    _ = v73
                    _ = v74
                }
                _ = v72
            }
            _ = fi38
            _ = v49
        }
        _ = fi37
    }
    delete(v4, 2)
    fmt.Println(v33)
    v75 := GenS3{F0: (v3 << ((1) & 63)), F1: v5}
    var v76 uint64 = 2
    switch 2147483647 {
    case 1:
        var v77 uint8 = uint8(v5)
        _ = v77
    case -254:
        var v78 int = v2
        _ = v78
    }
    fmt.Println(v3, v76)
    fmt.Println(v2)
    if false {
        v33 = v33
        _ = len(v4)
    }
    var v79 bool = (v33 >= v33)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v76:", v76)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v75
    _ = v76
    _ = v79
}
