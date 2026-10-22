package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int32; F1 uint16 }
type GenS8 struct { F0 int16; F1 int16; F2 uint8 }
type GenS9 struct { F0 string; F1 uint8 }
type GenS10 struct { F0 uint8; F1 int8; F2 string }
type GenS11 struct { F0 uint32; F1 int8 }
type GenS12 struct { F0 uint }
func genFn0() uint {
    var v4 uint64 = 42
    v5 := []int{0, 127, 0, 4294967295}
    for fi6 := 0; fi6 < 16; fi6++ {
        v4 += v4
        _ = fi6
    }
    v4 = v4
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 int8 = -128
    var v2 int8 = (v1 / ((79) | 1))
    var v3 uint64 = 42
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := GenS1{F0: 42}
        if (888 <= 428) {
            var v9 uint8 = 0
            if (124 > 657) {
                v2 = 0
                if true {
                    for fi10 := 0; fi10 < 17; fi10++ {
                        v9 -= v9
                        _ = fi10
                    }
                    fmt.Println(v2, v2, v2)
                    if true {
                        if (true || true) {
                            fmt.Println(v3, v3, v2)
                            v11 := GenS2{F0: (v9 << ((4294967295) & 63)), F1: (v3 == (v3 * v3)), F2: 4294967295}
                            var v12 int16 = int16(v2)
                            _ = v11
                            _ = v12
                        } else {
                            fmt.Println(v3)
                            v1 -= (v2 + (v1 / ((v1) | 1)))
                        }
                        var v13 uint16 = 0
                        _ = v13
                    }
                }
                v14 := []int{}
                _ = v14
            }
            if false {
                for fi15 := 0; fi15 < 1; fi15++ {
                    var v16 int64 = -9223372036854775808
                    v3 = 680
                    fmt.Println(v2, fi7)
                    v9 -= v9
                    v17 := GenS3{F0: true, F1: 2}
                    v9 = 42
                    if true {
                        var v18 float64 = 1.5
                        var v19 int8 = v1
                        _ = v18
                        _ = v19
                    } else {
                        var v20 int64 = v16
                        v21 := []int{(fi7 + 0), (-(fi7))}
                        var v22 uint64 = (^(v3))
                        v22 = v3
                        fmt.Println(v1, v2, fi7)
                        if (743 <= 0) {
                            var v23 bool = false
                            _ = v23
                        }
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                switch -989 {
                case 42:
                    v24 := GenS2{F0: 198, F1: ((((((3.14159 - -2.25) + 0.1) / float64(3.14159)) > 0.0) || true) && false), F2: 1}
                    _ = v24
                default:
                    v1 += v1
                }
                for fi25 := 0; fi25 < 13; fi25++ {
                    var v26 string = ("foo bar" + "\n\t")
                    v27 := GenS4{F0: 1, F1: v9}
                    switch fi25 {
                    case -9223372036854775808:
                        v9 += (v9 | 2)
                    case 32767:
                        v2 = v1
                    case 32768:
                        var v28 int8 = (v2 - v1)
                        _ = v28
                    default:
                        if (255 == 255) {
                            v1 = v2
                        } else {
                            if true {
                                v2 = 1
                                v29 := []int{28}
                                var v30 float32 = -2.25
                                if false {
                                    switch "" {
                                    case "foo bar":
                                        v1 = v2
                                    case "":
                                        _ = len(v29)
                                    case "a":
                                        switch (^(fi25)) {
                                        case 374:
                                            var v31 uint8 = 2
                                            _ = v31
                                        case 128:
                                            var v32 int32 = 474
                                            _ = v32
                                        case -1:
                                            v26 = (v26 + ("a" + v26))
                                        }
                                    }
                                } else {
                                    fmt.Println(v3, v26)
                                    if true {
                                        var v33 int = (fi7 >> ((4294967295) & 63))
                                        if false {
                                            v3 += v3
                                        }
                                        _ = v33
                                    }
                                    v1 = v2
                                }
                                v34 := GenS5{F0: (fi25 << ((65535) & 63)), F1: 2}
                                _ = v29
                                _ = v30
                                _ = v34
                            } else {
                                v9 = v9
                                var v35 uint8 = 128
                                v36 := GenS2{F0: (v35 + (v9 + v35)), F1: false, F2: 4294967295}
                                _ = v35
                                _ = v36
                            }
                            var v37 uint16 = 449
                            if (0 <= 0) {
                                var v38 string = ""
                                v39 := []int{}
                                _ = v38
                                _ = v39
                            }
                            _ = v37
                        }
                    }
                    v40 := GenS6{F0: 1}
                    _ = fi25
                    _ = v26
                    _ = v27
                    _ = v40
                }
            }
            _ = v9
        } else {
            fmt.Println(v2, v3, v2)
            var v41 string = ""
            var v42 uint = 0
            _ = v41
            _ = v42
        }
        _ = fi7
        _ = v8
    }
    if false {
        switch "hello" {
        case "\n\t":
            for fi43 := 0; fi43 < 15; fi43++ {
                for fi44 := 0; fi44 < 14; fi44++ {
                    v2 -= (v2 - 1)
                    v45 := GenS5{F0: fi44, F1: 0}
                    _ = fi44
                    _ = v45
                }
                var v46 uint = 1
                v46 = 255
                _ = fi43
                _ = v46
            }
        default:
            for fi47 := 0; fi47 < 3; fi47++ {
                for fi48 := 0; fi48 < 13; fi48++ {
                    var v49 int8 = (v2 | (v1 | (v1 + (v1 & (-(v1))))))
                    _ = fi48
                    _ = v49
                }
                v50 := []int{(fi47 / ((32768) | 1))}
                _ = len(v50)
                _ = fi47
                _ = v50
            }
        }
    } else {
        for fi51 := 0; fi51 < 15; fi51++ {
            v1 += v1
            v52 := []int{42}
            _ = fi51
            _ = v52
        }
        v53 := []int{0, 9223372036854775807, 128, 1}
        if true {
            v54 := GenS5{F0: 2147483647, F1: 65535}
            _ = v54
        } else {
            var v55 int64 = 42
            _ = v55
        }
        _ = v53
    }
    for fi56 := 0; fi56 < 6; fi56++ {
        for fi57 := 0; fi57 < 16; fi57++ {
            v1 = v1
            if ("foo bar" == "\n\t") {
                v58 := GenS7{F0: 1, F1: 42}
                v59 := []int{}
                v2 -= -1
                v60 := GenS3{F0: (32767 <= 0), F1: 171}
                _ = v58
                _ = v59
                _ = v60
            } else {
                var v61 uint64 = v3
                _ = v61
            }
            _ = fi57
        }
        var v62 int64 = 1
        switch (fi56 & 232) {
        case 2147483647:
            fmt.Println(v3, v1)
        default:
            var v63 int = int(v1)
            _ = v63
        }
        _ = fi56
        _ = v62
    }
    switch 0 {
    case -2147483648:
        v64 := GenS8{F0: -1, F1: 0, F2: 42}
        _ = v64
    }
    v65 := GenS9{F0: ("foo bar" + "a"), F1: 2}
    v3 = (v3 / (((v3 + v3)) | 1))
    v3 -= 0
    v66 := GenS10{F0: 223, F1: (v1 | (v1 & -128)), F2: "a"}
    for fi67 := 0; fi67 < 17; fi67++ {
        v2 += (v1 << ((853) & 63))
        var v68 int16 = 580
        switch fi67 {
        case 32767:
            fmt.Println(v3, v3)
        default:
            v69 := GenS11{F0: 504, F1: (v2 | (v2 << ((uint(v3)) & 63)))}
            _ = v69
        }
        _ = fi67
        _ = v68
    }
    var v70 string = (((("\n\t" + "") + "hello") + "hello") + "hello")
    var v71 string = v70
    switch 0 {
    case 9223372036854775807:
        v70 = (((v70 + v70) + "\n\t") + v71)
    case 2147483647:
        if (428 < 32767) {
            var v72 int16 = 32767
            if true {
                v73 := GenS12{F0: 2}
                v74 := []int{32768, 308, -485, -273}
                _ = v73
                _ = v74
            } else {
                v1 -= (v1 + v2)
                v75 := GenS10{F0: 42, F1: (v2 / ((-26) | 1)), F2: (v70 + v71)}
                v76 := GenS5{F0: 0, F1: genFn0()}
                _ = v75
                _ = v76
            }
            v70 = v70
            _ = v72
        }
    default:
        v77 := []int{4294967295, 2147483647, 4294967295, 145}
        _ = v77
    }
    v70 = ("foo bar" + v71)
    var v78 int8 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v65
    _ = v66
    _ = v70
    _ = v71
    _ = v78
}
