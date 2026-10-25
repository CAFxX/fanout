package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 bool; F1 int; F2 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16 }

func main() {
    var v1 int16 = 0
    v2 := make(map[int]uint32)
    v3 := []string{("a" + "a")}
    v4 := []uint32{}
    var v5 string = "\n\t"
    v6 := make(map[string]uint32)
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 float64 = 0.1
            v9 = 3.14159
            var v10 int = fi7
            _ = fi8
            _ = v10
            _ = v9
        }
        if true {
            v11 := GenS1{F0: 100, F1: 42}
            if (!false) {
                v12 := GenS1{F0: 42, F1: (^(v1))}
                for fi13 := 0; fi13 < 16; fi13++ {
                    v14 := v6[((v5 + v5) + v5)]
                    if len(v4) > 0 {
                        v4[0] = v14
                    }
                    _ = fi13
                    _ = v14
                }
                v15 := GenS1{F0: 127, F1: (-(v1))}
                _ = v12
                _ = v15
            }
            var v16 int = (fi7 + int(v1))
            _ = v11
            _ = v16
        }
        var v17 int64 = 9223372036854775807
        _ = fi7
        _ = v17
    }
    v1 -= (v1 % (((v1 + v1)) | 1))
    fmt.Println(v5, v1)
    var v18 int = -1
    var v19 float32 = (-2.25 * 3.14159)
    if (false && false) {
        for fi20 := 0; fi20 < 7; fi20++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            fmt.Println(fi20, v5, v1)
            for fi21 := 0; fi21 < 13; fi21++ {
                fmt.Println(v18)
                if (!false) {
                    fmt.Println(fi21, v18, fi21)
                    v22 := GenS2{F0: (-740 == fi20), F1: fi20, F2: ("hello" + "foo bar")}
                    if len(v4) > 0 {
                        v4[0] = 676
                    }
                    _ = v22
                } else {
                    v23 := GenS2{F0: true, F1: fi21, F2: ("a" + "")}
                    _ = v23
                }
                _ = fi21
            }
            _ = fi20
        }
        _ = len(v4)
        v24 := GenS2{F0: true, F1: (v18 - (v18 ^ (v18 * v18))), F2: "hello"}
        _ = v24
    } else {
        var v25 float32 = -2.25
        v19 -= v25
        v26 := v2[(v18 + (v18 + v18))]
        _ = v25
        _ = v26
    }
    switch (v18 % ((v18) | 1)) {
    case 809:
        v18 = (v18 | (^(v18)))
    case -2147483648:
        var v27 string = (v5 + v5)
        _ = v27
    default:
        switch v5 {
        case "a":
            v28 := GenS1{F0: -111, F1: v1}
            _ = v28
        case "\n\t":
            var v29 int32 = 42
            _ = v29
        }
    }
    var v30 uint16 = 0
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (779 != 1) {
        switch (v18 >> ((255) & 63)) {
        case 32767:
            v1 = v1
        case -333:
            switch ("\n\t" + v5) {
            case "a":
                if false {
                    if (v19 != float32(0.1)) {
                        v18 -= (v18 | v18)
                    }
                    _ = len(v3)
                } else {
                    if true {
                        delete(v6, v5)
                        var v31 int16 = (v1 % ((v1) | 1))
                        _ = v31
                    } else {
                        if true {
                            v19 += v19
                            var v32 int64 = -422
                            delete(v6, "")
                            _ = v32
                        } else {
                            var v33 bool = false
                            _ = v33
                        }
                        v34 := GenS3{F0: -101}
                        for fi35 := 0; fi35 < 12; fi35++ {
                            fmt.Println(v18, v18)
                            switch v5 {
                            case "foo bar":
                                switch "" {
                                case "a":
                                    fmt.Println(v5, v19)
                                case "foo bar":
                                    v36 := GenS4{F0: 4294967295, F1: 54}
                                    _ = v36
                                case "\n\t":
                                    if (!(true || false)) {
                                        fmt.Println(v18)
                                        var v37 int64 = -84
                                        _ = v37
                                    } else {
                                        var v38 int16 = -463
                                        v18 -= (v18 >> ((880) & 63))
                                        _ = v38
                                    }
                                default:
                                    if ((!(1 != 543)) || (2 < v30)) {
                                        if false {
                                            v18 = v18
                                            for fi39 := 0; fi39 < 18; fi39++ {
                                                var v40 int8 = 100
                                                var v41 int64 = 664
                                                var v42 int16 = (v1 / (((v1 * (v1 * 511))) | 1))
                                                _ = fi39
                                                _ = v40
                                                _ = v41
                                                _ = v42
                                            }
                                        } else {
                                            var v43 int16 = 0
                                            v1 -= (-(v1))
                                            for fi44 := 0; fi44 < 7; fi44++ {
                                                var v45 uint32 = 546
                                                var v46 int = (fi44 * (v18 >> ((65535) & 63)))
                                                _ = fi44
                                                _ = v45
                                                _ = v46
                                            }
                                            _ = v43
                                        }
                                    } else {
                                        v4 = append(v4, 1)
                                        v47 := GenS5{F0: v1}
                                        v6["foo bar"] = 2147483648
                                        _ = v47
                                    }
                                }
                            }
                            _ = fi35
                        }
                        _ = v34
                    }
                }
            case "":
                var v48 uint16 = (v30 + v30)
                _ = v48
            case "\n\t":
                if (!false) {
                    var v49 uint = 1
                    v50 := GenS1{F0: 62, F1: v1}
                    fmt.Println(v5, v19)
                    _ = v49
                    _ = v50
                }
            default:
                for fi51 := 0; fi51 < 11; fi51++ {
                    v52 := GenS5{F0: v1}
                    _ = fi51
                    _ = v52
                }
            }
        case 786:
            v53 := GenS6{F0: 42}
            _ = v53
        }
        v30 = (v30 + v30)
    } else {
        v1 = (v1 | (-(v1)))
        fmt.Println(v30, v1)
        switch v18 {
        case 42:
            v18 += 648
        case 641:
            v5 = v5
        default:
            if ((!false) && false) {
                v54 := GenS7{F0: v1}
                v55 := GenS1{F0: 1, F1: (-(v1))}
                _ = len(v4)
                _ = v54
                _ = v55
            } else {
                delete(v2, int(v30))
                var v56 float32 = (3.14159 - v19)
                v56 = v19
                _ = v56
            }
        }
    }
    v5 = (v5 + v5)
    switch (v18 % (((-(v18))) | 1)) {
    case 42:
        v57 := GenS4{F0: 2, F1: 0}
        _ = v57
    case -653:
        v5 = (v5 + ("\n\t" + (v5 + (v5 + ((v5 + (((v5 + v5) + "hello") + (v5 + ("a" + "foo bar")))) + "")))))
    case -1:
        if false {
            for fi58 := 0; fi58 < 2; fi58++ {
                v59 := GenS2{F0: (float32((0.0 / (0.0 - -2.25))) > float32(1e10)), F1: (fi58 ^ fi58), F2: ((v5 + v5) + "")}
                _ = len(v6)
                _ = fi58
                _ = v59
            }
            v60 := GenS2{F0: (true && false), F1: (^(v18)), F2: "\n\t"}
            var v61 uint32 = 2147483648
            _ = v60
            _ = v61
        }
    default:
        var v62 int16 = 46
        _ = v62
    }
    v1 = (v1 ^ 827)
    var v63 bool = (v19 <= ((v19 / (v19 + v19)) + ((v19 - ((v19 - (0.1 * v19)) * v19)) * ((0.1 + v19) + ((v19 + v19) - float32(1.5))))))
    fmt.Println(v18)
    v5 = "\n\t"
    for fi64 := 0; fi64 < 15; fi64++ {
        v30 += (v30 % ((v30) | 1))
        if false {
            if (212 != 127) {
                for fi65 := 0; fi65 < 17; fi65++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v66 uint = 255
                    var v67 string = v5
                    _ = fi65
                    _ = v66
                    _ = v67
                }
            }
            v6["\n\t"] = 4294967295
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            v68 := v6[""]
            if (!(65535 < 174)) {
                var v69 int8 = 127
                fmt.Println(v68)
                if (((-(v1)) < v1) || v63) {
                    for fi70 := 0; fi70 < 17; fi70++ {
                        var v71 string = v5
                        var v72 int8 = int8(v1)
                        v63 = (((^(v30)) == 0) && v63)
                        _ = fi70
                        _ = v71
                        _ = v72
                    }
                }
                _ = v69
            }
            _ = v68
        }
        fmt.Println(v30)
        _ = fi64
    }
    var v73 uint = 750
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v63:", v63)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v73
}
