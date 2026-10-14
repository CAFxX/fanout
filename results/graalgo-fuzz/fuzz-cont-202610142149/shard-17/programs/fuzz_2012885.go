package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint }
type GenS2 struct { F0 uint16; F1 int; F2 uint }
type GenS3 struct { F0 int32; F1 int64; F2 uint32 }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint16; F1 int; F2 uint16 }
type GenS8 struct { F0 uint64; F1 int8; F2 string }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int }
func genFn0() int64 {
    if (-938 == 0) {
        var v4 uint = 4294967295
        fmt.Println(v4)
        _ = v4
    } else {
        var v5 uint64 = 1
        v6 := []int{32767, 9223372036854775807, -802, -9223372036854775808}
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 2147483647, F1: 42, F2: 42}
    _ = v7
    return 1
}

func main() {
    v1 := []string{}
    var v2 float64 = (1e10 * 0.0)
    var v3 int32 = -518
    v8 := GenS1{F0: 853, F1: 675, F2: 4294967295}
    switch ("foo bar" + "a") {
    case "foo bar":
        if (-2.25 == 0.1) {
            if true {
                v2 = float64(1e10)
                if true {
                    var v9 int32 = v3
                    _ = v9
                }
                fmt.Println(v3)
            } else {
                fmt.Println(v3, v2)
            }
            for fi10 := 0; fi10 < 7; fi10++ {
                v3 = (-(v3))
                if false {
                    fmt.Println(fi10, v3, v2)
                    v3 = (v3 << ((966) & 63))
                }
                _ = fi10
            }
        } else {
            fmt.Println(v3, v3)
        }
    default:
        v2 += -2.25
    }
    var v11 float64 = (v2 / 1e10)
    var v12 bool = (371 <= 2147483648)
    switch "a" {
    case "\n\t":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    switch "a" {
    case "foo bar":
        v13 := GenS1{F0: 103, F1: 1, F2: 4294967295}
        _ = v13
    case "":
        switch -257 {
        case -2147483648:
            v2 += v2
        case 4294967295:
            v3 += -1
        }
    default:
        v14 := GenS2{F0: 1, F1: 32768, F2: 658}
        _ = v14
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        for fi16 := 0; fi16 < 5; fi16++ {
            v17 := GenS2{F0: 2, F1: (fi16 - fi15), F2: 4294967295}
            _ = fi16
            _ = v17
        }
        fmt.Println(v2, fi15, v11)
        _ = fi15
    }
    fmt.Println(v12, v11, v2)
    if (!((-1 >= -1) && (1 >= 773))) {
        _ = len(v1)
        if false {
            v18 := GenS1{F0: 65535, F1: 42, F2: 255}
            var v19 int16 = 514
            _ = v18
            _ = v19
        } else {
            v2 = (1e10 - (v11 / 1.5))
            var v20 uint16 = 1
            _ = v20
        }
        fmt.Println(v3, v2)
    }
    if len(v1) > 0 {
        v1[0] = ("a" + ("hello" + "foo bar"))
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        switch ("foo bar" + ("foo bar" + "foo bar")) {
        case "":
            var v22 uint = 42
            _ = v22
        case "foo bar":
            var v23 int = (fi21 << ((64) & 63))
            _ = v23
        default:
            for fi24 := 0; fi24 < 18; fi24++ {
                v25 := GenS3{F0: (-(v3)), F1: -1, F2: 2147483647}
                _ = fi24
                _ = v25
            }
        }
        for fi26 := 0; fi26 < 11; fi26++ {
            if v12 {
                if ((3.14159 / 3.14159) <= 1e10) {
                    var v27 float64 = 3.14159
                    switch "\n\t" {
                    case "foo bar":
                        switch (fi21 - (fi26 + fi26)) {
                        case -680:
                            var v28 float32 = 0.0
                            _ = v28
                        case 983:
                            var v29 int = fi21
                            _ = v29
                        case 32768:
                            if (0 >= 4294967295) {
                                var v30 bool = (400 != 42)
                                v31 := GenS1{F0: 2, F1: 32767, F2: 0}
                                _ = v30
                                _ = v31
                            } else {
                                var v32 bool = (72 >= 77)
                                if v12 {
                                    if v32 {
                                        switch "\n\t" {
                                        case "a":
                                            switch "\n\t" {
                                            case "\n\t":
                                                v33 := GenS4{F0: (v12 || v12), F1: 0}
                                                _ = v33
                                            case "a":
                                                if (4294967295 != 42) {
                                                    var v34 int8 = -128
                                                    switch "hello" {
                                                    case "foo bar":
                                                        var v35 uint8 = 151
                                                        _ = v35
                                                    default:
                                                        if (!((42 < 18446744073709551615) && (false && v12))) {
                                                            if len(v1) > 0 {
                                                                v1[0] = "hello"
                                                            }
                                                            v36 := GenS4{F0: true, F1: 0}
                                                            _ = v36
                                                        } else {
                                                            fmt.Println(v34)
                                                            v37 := GenS4{F0: false, F1: 1}
                                                            var v38 uint = 731
                                                            v39 := GenS1{F0: 273, F1: 141, F2: 93}
                                                            _ = v37
                                                            _ = v38
                                                            _ = v39
                                                        }
                                                    }
                                                    if (!(363 > 2147483648)) {
                                                        var v40 int32 = v3
                                                        if false {
                                                            fmt.Println(v2, v2)
                                                            v2 = (v27 - (float64(-2.25) - (((0.1 * v27) - (-2.25 - v27)) - 0.0)))
                                                            if v32 {
                                                                var v41 int = fi21
                                                                if len(v1) > 0 {
                                                                    v1[0] = ""
                                                                }
                                                                _ = v41
                                                            }
                                                        }
                                                        v3 -= v40
                                                        _ = v40
                                                    }
                                                    _ = v34
                                                } else {
                                                    if ((fi21 * fi21) == 578) {
                                                        var v42 int = fi21
                                                        if (!false) {
                                                            if v12 {
                                                                var v43 int = 428
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                var v44 uint = 255
                                                                _ = v43
                                                                _ = v44
                                                            }
                                                        } else {
                                                            if len(v1) > 0 {
                                                                v1[0] = ((("foo bar" + (("a" + ("\n\t" + "hello")) + "foo bar")) + ("foo bar" + ((("\n\t" + ("foo bar" + ("hello" + "a"))) + "") + ""))) + ("" + "hello"))
                                                            }
                                                        }
                                                        var v45 int32 = v3
                                                        _ = v42
                                                        _ = v45
                                                    }
                                                    var v46 int32 = v3
                                                    _ = v46
                                                }
                                            }
                                        case "foo bar":
                                            switch "a" {
                                            case "hello":
                                                v47 := GenS1{F0: 2147483647, F1: 65535, F2: 0}
                                                _ = v47
                                            case "foo bar":
                                                v1 = append(v1, ("a" + "\n\t"))
                                            case "a":
                                                v48 := GenS1{F0: 659, F1: 42, F2: 1}
                                                _ = len(v1)
                                                v12 = (!v12)
                                                _ = v48
                                            }
                                        default:
                                            v12 = false
                                        }
                                    }
                                    fmt.Println(v27, fi21)
                                    var v49 float32 = (float32((v11 + v2)) + (1e10 - (3.14159 / (1e10 / float32(v2)))))
                                    _ = v49
                                }
                                _ = len(v1)
                                _ = v32
                            }
                        }
                        var v50 int = (fi26 ^ (fi21 | (fi26 >> ((0) & 63))))
                        _ = v50
                    case "\n\t":
                        v51 := GenS2{F0: 727, F1: (fi21 * (fi21 >> ((255) & 63))), F2: 255}
                        _ = v51
                    }
                    fmt.Println(v11, v27, v12)
                    _ = v27
                } else {
                    var v52 uint64 = 0
                    v2 = v11
                    v53 := GenS1{F0: 2147483648, F1: 0, F2: 4294967295}
                    if v12 {
                        v54 := GenS5{F0: 65535, F1: fi26}
                        v2 += 1e10
                        _ = v54
                    } else {
                        var v55 uint16 = 2
                        v12 = (v12 || (v52 != v52))
                        v55 = (v55 ^ (v55 + (v55 >> ((160) & 63))))
                        _ = v55
                    }
                    _ = v52
                    _ = v53
                }
                if (-1 > 127) {
                    v11 += v2
                    v56 := GenS6{F0: -1}
                    var v57 string = ""
                    _ = v56
                    _ = v57
                } else {
                    switch 86 {
                    case 1:
                        if false {
                            if v12 {
                                v2 -= 1.5
                                v11 -= (v2 - (3.14159 + float64(-2.25)))
                            } else {
                                if true {
                                    v58 := GenS2{F0: 255, F1: (fi21 ^ (fi26 & (fi26 >> ((4294967295) & 63)))), F2: 139}
                                    _ = v58
                                } else {
                                    var v59 int = -707
                                    fmt.Println(fi26)
                                    var v60 int32 = (v3 & v3)
                                    v59 = fi26
                                    _ = v59
                                    _ = v60
                                }
                                if (v3 <= v3) {
                                    v61 := GenS7{F0: 0, F1: 4294967295, F2: 228}
                                    fmt.Println(v3, fi26)
                                    var v62 int16 = 32767
                                    v63 := GenS8{F0: 4294967295, F1: -1, F2: "hello"}
                                    _ = v61
                                    _ = v62
                                    _ = v63
                                } else {
                                    fmt.Println(v12)
                                    v11 = v2
                                    v2 -= 1e10
                                }
                                if v12 {
                                    v11 += (v2 / float64((float32(float64(3.14159)) - float32(0.0))))
                                } else {
                                    var v64 int8 = -1
                                    if true {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v65 := GenS7{F0: 2, F1: fi21, F2: 42}
                                        var v66 uint8 = 255
                                        _ = v65
                                        _ = v66
                                    }
                                    if (42 > 127) {
                                        if len(v1) > 0 {
                                            v1[0] = ("\n\t" + "\n\t")
                                        }
                                        switch ("a" + "hello") {
                                        case "\n\t":
                                            fmt.Println(fi21, v12, v11)
                                        default:
                                            v12 = (!(1 != -32768))
                                        }
                                    } else {
                                        var v67 int32 = 0
                                        _ = v67
                                    }
                                    _ = v64
                                }
                                var v68 bool = v12
                                _ = v68
                            }
                        }
                    default:
                        var v69 int = fi26
                        _ = v69
                    }
                    v70 := GenS9{F0: 1}
                    var v71 uint32 = 2
                    var v72 float32 = -2.25
                    var v73 int64 = 1
                    _ = v70
                    _ = v71
                    _ = v72
                    _ = v73
                }
            } else {
                v11 += v2
                fmt.Println(v11, fi26, v3)
                var v74 uint = 376
                if false {
                    v75 := GenS8{F0: 4294967295, F1: 127, F2: ("" + "\n\t")}
                    var v76 float32 = 0.1
                    _ = v75
                    _ = v76
                } else {
                    v77 := GenS10{F0: fi26}
                    if (v3 < (v3 * (v3 * -2147483648))) {
                        v74 += 255
                        v11 = v11
                        v12 = true
                    } else {
                        var v78 int = 0
                        v12 = ((fi21 + (fi26 / ((v78) | 1))) <= 1)
                        if true {
                            switch (fi26 - -1) {
                            case 582:
                                v3 = (v3 ^ v3)
                            case -490:
                                if (v12 || true) {
                                    var v79 bool = ((0 == 32767) || ((-(v3)) >= (v3 & v3)))
                                    _ = v79
                                } else {
                                    var v80 uint16 = 35
                                    switch "" {
                                    case "\n\t":
                                        switch "foo bar" {
                                        case "foo bar":
                                            v11 = 1e10
                                        case "hello":
                                            var v81 int64 = 42
                                            _ = v81
                                        case "a":
                                            fmt.Println(v74, v2)
                                        }
                                    default:
                                        v80 = 65535
                                    }
                                    _ = v80
                                }
                            }
                            if (9223372036854775807 != fi26) {
                                var v82 int = (v78 & (fi21 / ((fi26) | 1)))
                                var v83 int16 = -890
                                _ = v82
                                _ = v83
                            } else {
                                if v12 {
                                    var v84 int64 = 853
                                    if v12 {
                                        fmt.Println(v12)
                                        v12 = (5 != 1)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    _ = v84
                                } else {
                                    var v85 float64 = v11
                                    _ = v85
                                }
                                switch (fi26 ^ (fi21 | v78)) {
                                case 127:
                                    var v86 int16 = 184
                                    _ = v86
                                }
                                var v87 uint64 = 993
                                _ = v87
                            }
                        }
                        _ = v78
                    }
                    var v88 uint16 = 1
                    _ = v77
                    _ = v88
                }
                _ = v74
            }
            _ = fi26
        }
        var v89 float32 = float32(v11)
        _ = fi21
        _ = v89
    }
    var v90 uint = 481
    for fi91 := 0; fi91 < 10; fi91++ {
        var v92 uint = (v90 >> ((v90) & 63))
        var v93 bool = v12
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi91
        _ = v92
        _ = v93
    }
    v1 = append(v1, "foo bar")
    v2 += (-2.25 * float64((((3.14159 * 1e10) + ((1e10 - 0.1) / -2.25)) - 0.0)))
    switch 9223372036854775807 {
    case 32767:
        fmt.Println(v11, v11, v12)
    case 9223372036854775807:
        var v94 int8 = -128
        _ = v94
    case -215:
        var v95 bool = ((!(!(v12 || true))) && false)
        _ = v95
    }
    var v96 uint64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v90:", v90)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v8
    _ = v90
    _ = v96
}
