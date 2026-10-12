package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 uint; F1 uint; F2 string }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 int32; F1 uint64 }
type GenS7 struct { F0 uint8; F1 bool; F2 uint64 }
type GenS8 struct { F0 uint64 }

func main() {
    v1 := []int{2147483647, -9223372036854775808, 9223372036854775807, 0}
    v2 := make(map[string]string)
    var v3 uint8 = 188
    v4 := make(map[string]uint32)
    var v5 bool = ((!true) || true)
    v5 = (v3 <= (v3 >> ((649) & 63)))
    for fi6 := 0; fi6 < 19; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 uint8 = (v3 & v3)
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    var v9 int64 = 42
    switch -1 {
    case 648:
        for fi10 := 0; fi10 < 11; fi10++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi11 := 0; fi11 < 12; fi11++ {
                v4[("a" + (((("foo bar" + "foo bar") + "a") + "") + "foo bar"))] = 2147483648
                _ = fi11
            }
            fmt.Println(v5)
            _ = fi10
        }
    }
    var v12 uint8 = v3
    var v13 int8 = 40
    var v14 uint16 = 0
    if true {
        var v15 bool = (388 == 2)
        _ = v15
    }
    v9 = 42
    for fi16 := 0; fi16 < 4; fi16++ {
        if (v5 || (v5 && (v3 < (^(v3))))) {
            for fi17 := 0; fi17 < 1; fi17++ {
                v18 := GenS1{F0: (v14 - v14), F1: (v12 >> ((4294967295) & 63))}
                v19 := GenS1{F0: v14, F1: v12}
                _ = fi17
                _ = v18
                _ = v19
            }
        } else {
            v12 += 95
            var v20 bool = v5
            _ = v20
        }
        var v21 uint8 = 1
        v22 := GenS2{F0: ""}
        _ = fi16
        _ = v21
        _ = v22
    }
    if v5 {
        for fi23 := 0; fi23 < 6; fi23++ {
            switch fi23 {
            case 127:
                var v24 uint64 = 0
                _ = v24
            }
            if v5 {
                for fi25 := 0; fi25 < 16; fi25++ {
                    fmt.Println(v5, fi23)
                    var v26 uint64 = 0
                    if true {
                        v2[("hello" + "hello")] = "foo bar"
                        if v5 {
                            if (v5 && ((42 <= 42) && (v5 || v5))) {
                                v27 := GenS3{F0: (0.1 <= -2.25), F1: (v13 & int8(v26))}
                                v3 = (v3 & 127)
                                _ = v27
                            }
                        } else {
                            v5 = (((!(!v5)) && (v13 < (v13 << ((266) & 63)))) || false)
                            v26 += 4294967295
                            if v5 {
                                fmt.Println(v3)
                                var v28 uint16 = v14
                                if (v5 || true) {
                                    var v29 int16 = -210
                                    if (65535 == 570) {
                                        switch (-(fi23)) {
                                        case 32768:
                                            if (571 != 4294967295) {
                                                var v30 bool = v5
                                                _ = v30
                                            }
                                        case -9223372036854775808:
                                            var v31 int = (fi25 << ((18446744073709551615) & 63))
                                            _ = v31
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    fmt.Println(fi25, v29)
                                    _ = v29
                                }
                                _ = v28
                            }
                            var v32 bool = (uint8(v26) > 197)
                            if ((float64(1.5) / 3.14159) < 3.14159) {
                                var v33 int64 = v9
                                _ = v33
                            }
                            _ = v32
                        }
                        fmt.Println(v9)
                        var v34 uint8 = 0
                        var v35 uint = 2
                        _ = v34
                        _ = v35
                    } else {
                        if (31 < 0) {
                            v13 += (v13 << ((65535) & 63))
                        }
                        var v36 float32 = 1.5
                        v37 := GenS3{F0: ((v12 | (v12 ^ 255)) <= 128), F1: 127}
                        _ = v36
                        _ = v37
                    }
                    _ = fi25
                    _ = v26
                }
                v38 := v4["a"]
                v39 := v2[(((("\n\t" + ("hello" + "")) + "\n\t") + "\n\t") + ("a" + ""))]
                _ = v38
                _ = v39
            }
            _ = fi23
        }
    } else {
        var v40 uint64 = 18446744073709551615
        _ = v40
    }
    switch "a" {
    case "":
        v41 := GenS4{F0: 983, F1: 145, F2: "foo bar"}
        _ = v41
    }
    var v42 uint = 180
    for fi43 := 0; fi43 < 2; fi43++ {
        for fi44 := 0; fi44 < 3; fi44++ {
            var v45 int64 = (v9 ^ (v9 / (((v9 - (v9 | 516))) | 1)))
            fmt.Println(v13, v9, v14)
            v5 = v5
            _ = fi44
            _ = v45
        }
        if v5 {
            if false {
                v46 := GenS5{F0: 65535, F1: (1 >= 633)}
                for fi47 := 0; fi47 < 10; fi47++ {
                    v13 = (v13 - (v13 >> ((v42) & 63)))
                    _ = fi47
                }
                _ = v46
            } else {
                for fi48 := 0; fi48 < 1; fi48++ {
                    if (false || (1.5 >= -2.25)) {
                        v49 := GenS5{F0: 4294967295, F1: true}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v50 bool = (1 > v14)
                        _ = v49
                        _ = v50
                    }
                    switch 383 {
                    case -2147483648:
                        var v51 uint = 0
                        _ = v51
                    case 32768:
                        switch "" {
                        case "":
                            v52 := GenS1{F0: v14, F1: v12}
                            _ = v52
                        }
                    }
                    var v53 int64 = 1
                    _ = fi48
                    _ = v53
                }
                switch fi43 {
                case 776:
                    v9 = (v9 << ((v42) & 63))
                case -9223372036854775808:
                    switch fi43 {
                    case -2147483648:
                        if (("a" < ((("\n\t" + "hello") + "hello") + "\n\t")) || ("foo bar" >= "hello")) {
                            v54 := GenS2{F0: ""}
                            _ = v54
                        } else {
                            v13 += v13
                        }
                    default:
                        if ((v14 << (((^(v42))) & 63)) != v14) {
                            delete(v4, (("" + "a") + "a"))
                            _ = len(v1)
                            for fi55 := 0; fi55 < 19; fi55++ {
                                if v5 {
                                    if v5 {
                                        if false {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            if (!((v12 ^ uint8(v42)) > (v3 & (v12 << (((v42 + v42)) & 63))))) {
                                                v56 := GenS6{F0: -950, F1: 2}
                                                v12 -= (v3 % ((255) | 1))
                                                var v57 int = (^(fi43))
                                                if ((v3 % (((v12 - v12)) | 1)) < (^(v3))) {
                                                    v3 = 254
                                                    v5 = ((v14 / ((255) | 1)) != 65535)
                                                    var v58 string = "\n\t"
                                                    _ = v58
                                                }
                                                _ = v56
                                                _ = v57
                                            }
                                        } else {
                                            v42 += (^(v42))
                                            if (v5 || (v12 == 2)) {
                                                var v59 uint8 = 42
                                                _ = v59
                                            } else {
                                                var v60 float32 = float32((-2.25 - (3.14159 - float64(0.1))))
                                                _ = v60
                                            }
                                            if true {
                                                if ((52 < 42) && true) {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    v61 := GenS2{F0: ""}
                                                    _ = v61
                                                } else {
                                                    v62 := GenS6{F0: -439, F1: 42}
                                                    var v63 int32 = 2147483647
                                                    v5 = (646 >= -226)
                                                    switch (-(fi55)) {
                                                    case -148:
                                                        switch "hello" {
                                                        case "\n\t":
                                                            if v5 {
                                                                var v64 int32 = (-(v63))
                                                                fmt.Println(v5, fi55, v14)
                                                                _ = len(v1)
                                                                _ = v64
                                                            } else {
                                                                fmt.Println(v9, v14, fi55)
                                                                var v65 int16 = -32768
                                                                _ = v65
                                                            }
                                                        default:
                                                            v66 := GenS6{F0: (v63 >> (((v42 >> (((v42 * v42)) & 63))) & 63)), F1: 704}
                                                            _ = v66
                                                        }
                                                    case -432:
                                                        var v67 uint16 = v14
                                                        _ = v67
                                                    }
                                                    _ = v62
                                                    _ = v63
                                                }
                                                v5 = (((v9 | v9) > v9) || true)
                                                var v68 float32 = 1.5
                                                _ = v68
                                            } else {
                                                var v69 uint16 = uint16(v12)
                                                switch (fi55 | (fi55 | fi55)) {
                                                case 115:
                                                    v70 := GenS3{F0: (4294967295 < 700), F1: v13}
                                                    _ = v70
                                                case 1:
                                                    if (true || v5) {
                                                        v71 := v4[(("hello" + (("foo bar" + "hello") + ("a" + (("foo bar" + "\n\t") + "hello")))) + "hello")]
                                                        if v5 {
                                                            v72 := GenS4{F0: 18446744073709551615, F1: (v42 & (v42 & v42)), F2: ("a" + "hello")}
                                                            v12 = uint8(v71)
                                                            v13 = (v13 ^ v13)
                                                            _ = v72
                                                        } else {
                                                            v73 := GenS7{F0: v3, F1: false, F2: 18446744073709551615}
                                                            _ = v73
                                                        }
                                                        if true {
                                                            if false {
                                                                if ((uint16(v3) < v14) || (0 > uint64(v12))) {
                                                                    fmt.Println(v9, v5)
                                                                    if (4294967295 >= 457) {
                                                                        var v74 uint32 = uint32(v14)
                                                                        if false {
                                                                            _ = len(v4)
                                                                            v75 := GenS1{F0: 255, F1: (^(v12))}
                                                                            v76 := GenS1{F0: v14, F1: (v3 % ((v3) | 1))}
                                                                            _ = v75
                                                                            _ = v76
                                                                        } else {
                                                                            v13 = (^(v13))
                                                                            v77 := GenS8{F0: 70}
                                                                            v78 := GenS7{F0: (v12 / ((1) | 1)), F1: true, F2: uint64(v74)}
                                                                            _ = v77
                                                                            _ = v78
                                                                        }
                                                                        switch fi43 {
                                                                        case -2147483648:
                                                                            var v79 int = int(v71)
                                                                            _ = v79
                                                                        case -448:
                                                                            v3 = v3
                                                                        }
                                                                        _ = v74
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = v71
                                                    } else {
                                                        var v80 int = fi55
                                                        _ = v80
                                                    }
                                                }
                                                v12 += v3
                                                _ = v69
                                            }
                                        }
                                    } else {
                                        var v81 int8 = (-(v13))
                                        _ = v81
                                    }
                                    var v82 uint8 = v3
                                    v3 += 132
                                    _ = v82
                                }
                                var v83 int32 = -2147483648
                                _ = fi55
                                _ = v83
                            }
                        }
                    }
                }
            }
            var v84 uint64 = 0
            _ = v84
        }
        v85 := GenS7{F0: 0, F1: (-1 >= 813), F2: 4294967295}
        _ = fi43
        _ = v85
    }
    if false {
        for fi86 := 0; fi86 < 14; fi86++ {
            for fi87 := 0; fi87 < 16; fi87++ {
                fmt.Println(v12, v5, fi87)
                _ = len(v2)
                _ = fi87
            }
            var v88 int = 4294967295
            _ = fi86
            _ = v88
        }
        v42 = (v42 >> (((v42 >> (((v42 << ((v42) & 63))) & 63))) & 63))
    }
    fmt.Println(v14, v3, v42)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v9
}
