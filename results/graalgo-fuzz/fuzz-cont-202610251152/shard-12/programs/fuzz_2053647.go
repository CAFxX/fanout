package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int; F2 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8; F1 int32 }
type GenS7 struct { F0 int8; F1 int64; F2 int8 }
type GenS8 struct { F0 int8; F1 uint64 }
type GenS9 struct { F0 bool; F1 int8 }

func main() {
    var v1 bool = false
    v2 := make(map[string]uint32)
    var v3 int64 = 0
    v4 := []int{}
    if true {
        v5 := GenS1{F0: 1}
        _ = v5
    }
    v3 += v3
    v6 := GenS2{F0: (((("\n\t" + "hello") + "hello") + "") + "\n\t"), F1: 32768, F2: (v3 & 42)}
    if ((("a" + "a") + "hello") > "foo bar") {
        for fi7 := 0; fi7 < 14; fi7++ {
            v8 := GenS1{F0: 116}
            switch ("\n\t" + (("" + "hello") + "\n\t")) {
            case "":
                v9 := GenS2{F0: "foo bar", F1: fi7, F2: 1}
                _ = v9
            case "hello":
                var v10 int64 = (v3 ^ v3)
                _ = v10
            case "a":
                switch "" {
                case "\n\t":
                    var v11 bool = v1
                    _ = v11
                case "hello":
                    switch 32767 {
                    case 102:
                        var v12 uint = 0
                        _ = v12
                    case -262:
                        v3 -= v3
                    case 4294967295:
                        _ = len(v4)
                    }
                }
            default:
                fmt.Println(fi7, fi7, fi7)
            }
            _ = fi7
            _ = v8
        }
        var v13 uint16 = 0
        v14 := v2["foo bar"]
        _ = v13
        _ = v14
    } else {
        delete(v2, "\n\t")
    }
    v1 = (v1 || (2147483647 == 181))
    var v15 bool = v1
    if v1 {
        var v16 int64 = 1
        v17 := GenS3{F0: "\n\t"}
        if (!(779 <= (v3 + (v16 & 1)))) {
            if (1 == 42) {
                var v18 float64 = 3.14159
                v16 -= (-(v3))
                if v1 {
                    v19 := GenS4{F0: (-32768 >= -1)}
                    v1 = true
                    _ = v19
                }
                _ = v18
            } else {
                v20 := GenS3{F0: "hello"}
                var v21 uint8 = 2
                v16 += v16
                _ = v20
                _ = v21
            }
        }
        _ = v16
        _ = v17
    } else {
        if (!v15) {
            for fi22 := 0; fi22 < 14; fi22++ {
                fmt.Println(v3, v3)
                switch (fi22 << ((42) & 63)) {
                case 127:
                    for fi23 := 0; fi23 < 15; fi23++ {
                        var v24 bool = true
                        _ = fi23
                        _ = v24
                    }
                case -401:
                    if v1 {
                        v25 := GenS2{F0: "\n\t", F1: 524, F2: 0}
                        for fi26 := 0; fi26 < 17; fi26++ {
                            if ("foo bar" <= "a") {
                                if false {
                                    var v27 string = "\n\t"
                                    _ = v27
                                } else {
                                    _ = len(v4)
                                    delete(v2, "hello")
                                    v28 := GenS5{F0: 0}
                                    _ = v28
                                }
                                var v29 string = "hello"
                                switch 42 {
                                case 32767:
                                    v30 := GenS6{F0: 42, F1: 0}
                                    _ = v30
                                case -1:
                                    switch (("a" + v29) + v29) {
                                    case "foo bar":
                                        fmt.Println(v29)
                                    case "a":
                                        var v31 int = 0
                                        v32 := v2[v29]
                                        _ = v31
                                        _ = v32
                                    case "":
                                        var v33 string = v29
                                        _ = len(v4)
                                        _ = v33
                                    default:
                                        v15 = (!(42 <= 706))
                                    }
                                }
                                v3 = v3
                                _ = v29
                            } else {
                                if v1 {
                                    if false {
                                        v3 -= 71
                                    }
                                    var v34 float32 = (0.0 + 3.14159)
                                    v34 = 1.5
                                    if false {
                                        var v35 uint = 1
                                        var v36 int64 = 0
                                        v35 += (v35 % ((143) | 1))
                                        _ = v35
                                        _ = v36
                                    } else {
                                        delete(v2, "foo bar")
                                    }
                                    var v37 string = "foo bar"
                                    _ = len(v2)
                                    v38 := GenS6{F0: -13, F1: 2147483647}
                                    _ = v34
                                    _ = v37
                                    _ = v38
                                }
                            }
                            var v39 int32 = 0
                            v15 = v15
                            var v40 uint64 = 2
                            var v41 float64 = 0.0
                            fmt.Println(v41)
                            v41 += -2.25
                            fmt.Println(fi26, v3)
                            v40 -= 394
                            _ = fi26
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                        _ = v25
                    } else {
                        v3 += (v3 << ((42) & 63))
                        v3 = v3
                        for fi42 := 0; fi42 < 8; fi42++ {
                            _ = len(v2)
                            v3 = v3
                            _ = fi42
                        }
                    }
                }
                _ = fi22
            }
            v43 := GenS7{F0: -128, F1: v3, F2: 19}
            _ = v43
        } else {
            _ = len(v4)
            v15 = (!v1)
        }
        var v44 int8 = int8(v3)
        v45 := GenS5{F0: -1}
        _ = v44
        _ = v45
    }
    for fi46 := 0; fi46 < 7; fi46++ {
        for fi47 := 0; fi47 < 14; fi47++ {
            if false {
                if (-2.25 < 1e10) {
                    fmt.Println(v15, v1, v1)
                } else {
                    var v48 int8 = -40
                    _ = v48
                }
            } else {
                v4 = append(v4, 32768)
                v3 -= (v3 & v3)
            }
            v1 = (1 != 32767)
            v49 := GenS8{F0: 1, F1: 18446744073709551615}
            _ = fi47
            _ = v49
        }
        for fi50 := 0; fi50 < 19; fi50++ {
            switch (fi46 & (fi46 << ((75) & 63))) {
            case 0:
                if (0 != 0) {
                    var v51 uint8 = 47
                    _ = v51
                } else {
                    fmt.Println(v1, v1)
                    v52 := GenS5{F0: -2147483648}
                    fmt.Println(v3, v3, v3)
                    _ = v52
                }
            default:
                switch -865 {
                case -9223372036854775808:
                    v15 = (((0.1 == 1.5) || (v1 && (219 >= 1))) && (250 != 1))
                case 32768:
                    if v1 {
                        _ = len(v4)
                        v53 := GenS3{F0: "a"}
                        if true {
                            _ = len(v4)
                            var v54 int16 = 1
                            fmt.Println(v3, fi46)
                            _ = v54
                        }
                        _ = v53
                    }
                case -2147483648:
                    var v55 uint8 = 42
                    _ = v55
                }
            }
            switch fi50 {
            case 0:
                var v56 uint8 = 2
                _ = v56
            case 127:
                switch fi50 {
                case -1:
                    switch fi46 {
                    case 32767:
                        switch ("hello" + "\n\t") {
                        case "\n\t":
                            if len(v4) > 0 {
                                v4[0] = -587
                            }
                        }
                    case 127:
                        if false {
                            if v1 {
                                var v57 int8 = -108
                                if v1 {
                                    v58 := GenS5{F0: 2147483647}
                                    var v59 int = (fi46 | (fi50 << ((1) & 63)))
                                    var v60 int8 = v57
                                    switch ("\n\t" + ("hello" + "foo bar")) {
                                    case "\n\t":
                                        switch ("foo bar" + "a") {
                                        case "\n\t":
                                            switch (fi50 + -9223372036854775808) {
                                            case 292:
                                                v57 = (v60 | (v57 + (v60 * v60)))
                                                if (true || (!v15)) {
                                                    if v1 {
                                                        _ = len(v4)
                                                        v57 = (^(v60))
                                                        v15 = (!(!(2 == 42)))
                                                        v3 = (v3 * (v3 % (((v3 & v3)) | 1)))
                                                    }
                                                    v61 := GenS6{F0: v57, F1: 2147483647}
                                                    if (1e10 <= float32(0.0)) {
                                                        v62 := v2[("a" + "foo bar")]
                                                        if (float32(-2.25) == 1.5) {
                                                            var v63 bool = v15
                                                            v64 := v2["a"]
                                                            _ = v63
                                                            _ = v64
                                                        }
                                                        var v65 uint32 = v62
                                                        v62 = (^(v62))
                                                        var v66 int64 = v3
                                                        var v67 string = "\n\t"
                                                        v67 = v67
                                                        _ = v62
                                                        _ = v65
                                                        _ = v66
                                                        _ = v67
                                                    } else {
                                                        var v68 int16 = -979
                                                        v59 = (fi46 % ((fi50) | 1))
                                                        if true {
                                                            v15 = true
                                                            if (283 == 1) {
                                                                v1 = v1
                                                                _ = len(v2)
                                                                if (!(v60 != v57)) {
                                                                    v59 = fi50
                                                                    v69 := GenS7{F0: v57, F1: v3, F2: (v57 + (v57 ^ int8(v68)))}
                                                                    _ = v69
                                                                }
                                                            } else {
                                                                fmt.Println(v15, fi50)
                                                            }
                                                        } else {
                                                            if v1 {
                                                                var v70 int64 = -9223372036854775808
                                                                _ = v70
                                                            }
                                                        }
                                                        _ = v68
                                                    }
                                                    _ = v61
                                                } else {
                                                    _ = len(v2)
                                                }
                                                var v71 int16 = 1
                                                _ = v71
                                            default:
                                                var v72 int = 42
                                                _ = v72
                                            }
                                        }
                                    case "foo bar":
                                        var v73 int32 = 356
                                        _ = v73
                                    case "hello":
                                        v2["hello"] = 65535
                                    default:
                                        v74 := GenS5{F0: 2147483647}
                                        _ = v74
                                    }
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                }
                                v75 := v2[(((("hello" + "\n\t") + "hello") + "foo bar") + ("hello" + ""))]
                                _ = v57
                                _ = v75
                            } else {
                                if (0 >= 551) {
                                    v15 = (!v15)
                                }
                                v3 -= (v3 % ((0) | 1))
                                switch (fi46 * 127) {
                                case 898:
                                    var v76 uint = 18446744073709551615
                                    _ = v76
                                case 771:
                                    v77 := GenS8{F0: -15, F1: 18446744073709551615}
                                    _ = v77
                                }
                            }
                            v78 := GenS1{F0: 65535}
                            if (-69 == 127) {
                                if (!(!v15)) {
                                    v79 := GenS6{F0: -128, F1: 921}
                                    var v80 int = (fi50 >> ((2) & 63))
                                    _ = v79
                                    _ = v80
                                } else {
                                    var v81 int8 = int8(fi50)
                                    if false {
                                        switch (^(fi46)) {
                                        case 32767:
                                            if (-32768 == 79) {
                                                var v82 int64 = -9223372036854775808
                                                _ = v82
                                            }
                                        default:
                                            if len(v4) > 0 {
                                                v4[0] = -1
                                            }
                                            var v83 int8 = v81
                                            _ = v83
                                        }
                                        var v84 uint32 = 182
                                        _ = len(v2)
                                        _ = v84
                                    } else {
                                        var v85 uint16 = 65535
                                        _ = v85
                                    }
                                    _ = v81
                                }
                                var v86 int16 = -1
                                v1 = (4294967295 <= 4294967295)
                                _ = v86
                            }
                            _ = v78
                        } else {
                            if (32767 != 42) {
                                v87 := GenS4{F0: v1}
                                if v1 {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                v88 := GenS4{F0: v1}
                                _ = v87
                                _ = v88
                            } else {
                                v2["hello"] = 2
                                switch (fi46 + (fi46 * fi50)) {
                                case 187:
                                    delete(v2, "a")
                                case -426:
                                    v1 = v15
                                }
                            }
                        }
                    }
                case 32768:
                    v89 := v2[""]
                    _ = v89
                case -943:
                    if v15 {
                        if (255 <= 32767) {
                            v15 = (488 == 802)
                        }
                        v15 = (!(float32(1.5) > (0.0 + 0.1)))
                        var v90 int32 = 384
                        _ = v90
                    } else {
                        switch fi50 {
                        case 128:
                            v1 = (!v1)
                        }
                        delete(v2, ("\n\t" + "hello"))
                    }
                }
            default:
                var v91 uint64 = 42
                _ = v91
            }
            v3 += (v3 / ((v3) | 1))
            _ = fi50
        }
        fmt.Println(v3, v3)
        _ = fi46
    }
    var v92 int16 = 1
    v1 = v15
    var v93 int16 = (v92 - v92)
    v3 -= v3
    v94 := GenS1{F0: 154}
    if v15 {
        v95 := GenS9{F0: v15, F1: 127}
        _ = v95
    }
    for fi96 := 0; fi96 < 3; fi96++ {
        for fi97 := 0; fi97 < 9; fi97++ {
            switch int(v3) {
            case 0:
                switch (-(fi97)) {
                case 127:
                    var v98 float64 = 1e10
                    _ = v98
                }
            case 528:
                fmt.Println(fi97, fi96)
            }
            _ = fi97
        }
        _ = fi96
    }
    v99 := GenS8{F0: 127, F1: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v92
    _ = v93
    _ = v94
    _ = v99
}
