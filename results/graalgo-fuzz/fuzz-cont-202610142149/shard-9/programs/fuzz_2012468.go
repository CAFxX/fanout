package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 int16; F1 int32; F2 bool }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint8; F1 uint32 }
type GenS9 struct { F0 int16; F1 int64 }

func main() {
    v1 := []uint32{4294967295, 0, 454}
    v2 := make(map[int]bool)
    var v3 float32 = (1.5 - 1e10)
    v3 += 3.14159
    switch "\n\t" {
    case "\n\t":
        _ = len(v1)
    case "foo bar":
        fmt.Println(v3)
    case "a":
        var v4 uint32 = 42
        _ = v4
    }
    var v5 int = 127
    for fi6 := 0; fi6 < 4; fi6++ {
        v5 -= fi6
        _ = fi6
    }
    v1 = append(v1, 65535)
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            fmt.Println(fi8)
            switch "a" {
            case "":
                var v9 int32 = -1
                _ = v9
            case "hello":
                v10 := v2[-501]
                _ = v10
            case "\n\t":
                var v11 string = (("hello" + ("a" + ("" + "a"))) + "foo bar")
                _ = v11
            }
            if (true || false) {
                v3 = 0.1
                var v12 uint32 = 42
                v13 := GenS1{F0: (!(-528 < 32767))}
                _ = v12
                _ = v13
            }
            _ = fi8
        }
        _ = fi7
    }
    fmt.Println(v3, v3)
    v14 := v2[v5]
    v15 := GenS2{F0: 42, F1: 128}
    v3 = v3
    v1 = append(v1, 42)
    var v16 uint8 = 1
    if (!v14) {
        v16 = (v16 % (((v16 ^ 255)) | 1))
        var v17 int16 = 1000
        if (128 <= 178) {
            if v14 {
                v5 -= (v5 << ((0) & 63))
            }
        }
        _ = v17
    }
    if (v14 || ((v14 && v14) || true)) {
        fmt.Println(v3, v3, v3)
        switch v5 {
        case 128:
            v14 = (v14 && (!false))
        case -1:
            for fi18 := 0; fi18 < 18; fi18++ {
                var v19 int64 = 42
                var v20 string = "foo bar"
                if v14 {
                    if true {
                        v20 = (((((v20 + "\n\t") + "hello") + v20) + (v20 + (((((((v20 + "hello") + v20) + v20) + (v20 + v20)) + v20) + "foo bar") + v20))) + "hello")
                    } else {
                        fmt.Println(v5, fi18)
                    }
                    if (1e10 > 1.5) {
                        for fi21 := 0; fi21 < 18; fi21++ {
                            v22 := GenS2{F0: 671, F1: v16}
                            if v14 {
                                var v23 uint16 = 772
                                v24 := GenS3{F0: 615, F1: 42}
                                delete(v2, (^(v5)))
                                _ = v23
                                _ = v24
                            }
                            var v25 uint8 = v16
                            _ = fi21
                            _ = v22
                            _ = v25
                        }
                        switch (v20 + ((v20 + (("hello" + "\n\t") + "")) + "\n\t")) {
                        case "":
                            if true {
                                var v26 uint16 = 255
                                if len(v1) > 0 {
                                    v1[0] = 4294967295
                                }
                                _ = v26
                            } else {
                                var v27 uint32 = 4294967295
                                var v28 float32 = ((0.0 / (v3 + v3)) + ((v3 * (v3 + (-2.25 * v3))) * v3))
                                _ = v27
                                _ = v28
                            }
                        case "\n\t":
                            for fi29 := 0; fi29 < 14; fi29++ {
                                var v30 uint = 4294967295
                                _ = fi29
                                _ = v30
                            }
                        }
                        v31 := GenS4{F0: 32767, F1: 0, F2: v14}
                        _ = v31
                    } else {
                        v32 := GenS5{F0: 1, F1: v19}
                        switch "" {
                        case "":
                            v5 -= (-(v5))
                        case "\n\t":
                            if true {
                                if (100 < 114) {
                                    var v33 float64 = 0.1
                                    _ = v33
                                }
                                switch (v5 << ((1) & 63)) {
                                case 42:
                                    if ((v20 + v20) >= v20) {
                                        var v34 uint8 = v16
                                        for fi35 := 0; fi35 < 16; fi35++ {
                                            v19 -= int64(v16)
                                            if (528 > 42) {
                                                var v36 float64 = 1.5
                                                v37 := v2[(fi35 * (v5 - (v5 % (((fi18 * fi35)) | 1))))]
                                                if (42 < 18446744073709551615) {
                                                    if v14 {
                                                        v38 := v2[fi35]
                                                        if (-128 != 127) {
                                                            v39 := GenS3{F0: -1, F1: fi35}
                                                            v40 := GenS1{F0: v38}
                                                            _ = v39
                                                            _ = v40
                                                        } else {
                                                            if false {
                                                                fmt.Println(v5, v5)
                                                                if (2 >= 42) {
                                                                    fmt.Println(v5, v19)
                                                                    if false {
                                                                        fmt.Println(v19)
                                                                        var v41 int16 = 0
                                                                        _ = v41
                                                                    } else {
                                                                        var v42 uint = 2
                                                                        v38 = v14
                                                                        if (333 >= -32768) {
                                                                            var v43 string = v20
                                                                            switch fi35 {
                                                                            case 127:
                                                                                switch v5 {
                                                                                case 32767:
                                                                                    v44 := GenS6{F0: 32767, F1: v16}
                                                                                    if v14 {
                                                                                        var v45 float32 = 1e10
                                                                                        _ = v45
                                                                                    }
                                                                                    if v37 {
                                                                                        var v46 bool = (-1 > -507)
                                                                                        var v47 uint8 = 94
                                                                                        v19 = (v19 >> (((^(v42))) & 63))
                                                                                        _ = v46
                                                                                        _ = v47
                                                                                    } else {
                                                                                        var v48 uint8 = 255
                                                                                        fmt.Println(v42, v16)
                                                                                        var v49 float64 = (((float64((v3 + v3)) + 1.5) * 1e10) / (-2.25 * ((v36 - (v36 - ((v36 / v36) + v36))) * v36)))
                                                                                        _ = v48
                                                                                        _ = v49
                                                                                    }
                                                                                    _ = v44
                                                                                case -2147483648:
                                                                                    v1 = append(v1, 2)
                                                                                default:
                                                                                    v36 = (v36 / (0.1 / ((v36 * (v36 + v36)) - float64(3.14159))))
                                                                                }
                                                                                var v50 string = ("hello" + (((v20 + v20) + (v43 + v43)) + (v43 + v43)))
                                                                                switch v20 {
                                                                                case "":
                                                                                    var v51 int16 = -32768
                                                                                    switch (v43 + v43) {
                                                                                    case "\n\t":
                                                                                        v20 = ("foo bar" + v20)
                                                                                    case "hello":
                                                                                        var v52 float64 = (v36 + ((-2.25 * v36) - 1.5))
                                                                                        _ = v52
                                                                                    case "foo bar":
                                                                                        switch v50 {
                                                                                        case "hello":
                                                                                            switch -721 {
                                                                                            case 42:
                                                                                                v53 := GenS7{F0: 127}
                                                                                                _ = v53
                                                                                            }
                                                                                        case "":
                                                                                            if true {
                                                                                                if (v51 <= 32767) {
                                                                                                    _ = len(v1)
                                                                                                }
                                                                                                v16 -= v16
                                                                                            }
                                                                                        default:
                                                                                            v3 = float32(1e10)
                                                                                        }
                                                                                    }
                                                                                    if (v16 < (v34 + v34)) {
                                                                                        if (!(!v14)) {
                                                                                            var v54 int16 = 734
                                                                                            v38 = v38
                                                                                            var v55 uint16 = 572
                                                                                            _ = v54
                                                                                            _ = v55
                                                                                        }
                                                                                    } else {
                                                                                        v34 = (v34 - 1)
                                                                                    }
                                                                                    _ = v51
                                                                                case "foo bar":
                                                                                    _ = len(v1)
                                                                                    if v38 {
                                                                                        var v56 int16 = 1000
                                                                                        var v57 float32 = (v3 / 1.5)
                                                                                        v58 := GenS3{F0: (^(v56)), F1: (v5 / ((int(v56)) | 1))}
                                                                                        _ = v56
                                                                                        _ = v57
                                                                                        _ = v58
                                                                                    } else {
                                                                                        fmt.Println(v34, v50, v43)
                                                                                        var v59 int64 = 1
                                                                                        var v60 int16 = 278
                                                                                        _ = v59
                                                                                        _ = v60
                                                                                    }
                                                                                case "a":
                                                                                    var v61 float32 = ((0.0 + (1e10 * ((1.5 - (((v3 * (((v3 / (v3 + v3)) - ((v3 - v3) - -2.25)) - (-2.25 + v3))) / (v3 + v3)) + 0.1)) + 1.5))) + (v3 / -2.25))
                                                                                    _ = v61
                                                                                default:
                                                                                    var v62 int32 = 930
                                                                                    _ = v62
                                                                                }
                                                                                _ = v50
                                                                            }
                                                                            if false {
                                                                                v63 := GenS4{F0: 42, F1: 1, F2: (!(!v14))}
                                                                                if (!v38) {
                                                                                    if (!(v14 && (42 != 475))) {
                                                                                        var v64 string = (("" + v43) + ("" + v20))
                                                                                        _ = v64
                                                                                    } else {
                                                                                        if (145 < 570) {
                                                                                            if true {
                                                                                                var v65 int8 = -9
                                                                                                if v38 {
                                                                                                    fmt.Println(v5, v5)
                                                                                                }
                                                                                                _ = v65
                                                                                            } else {
                                                                                                var v66 int8 = int8(v16)
                                                                                                if len(v1) >= 2 {
                                                                                                    v1 = v1[1:len(v1)]
                                                                                                }
                                                                                                var v67 uint64 = 42
                                                                                                _ = v66
                                                                                                _ = v67
                                                                                            }
                                                                                        } else {
                                                                                            var v68 uint = (v42 >> (((^(v42))) & 63))
                                                                                            _ = v68
                                                                                        }
                                                                                        v43 = v43
                                                                                    }
                                                                                    var v69 int16 = int16(v42)
                                                                                    v70 := GenS3{F0: (v69 - (v69 ^ 1)), F1: (-(fi35))}
                                                                                    var v71 uint32 = 4294967295
                                                                                    _ = v69
                                                                                    _ = v70
                                                                                    _ = v71
                                                                                }
                                                                                var v72 float64 = float64(v3)
                                                                                _ = v63
                                                                                _ = v72
                                                                            } else {
                                                                                v1 = append(v1, 65535)
                                                                                fmt.Println(v42, v34)
                                                                                var v73 float64 = v36
                                                                                _ = v73
                                                                            }
                                                                            _ = v43
                                                                        }
                                                                        _ = v42
                                                                    }
                                                                } else {
                                                                    switch fi18 {
                                                                    case -9223372036854775808:
                                                                        v1 = append(v1, 0)
                                                                    case -1:
                                                                        switch ("a" + v20) {
                                                                        case "":
                                                                            v34 = v16
                                                                        case "foo bar":
                                                                            v38 = false
                                                                        case "\n\t":
                                                                            v74 := GenS6{F0: 953, F1: 0}
                                                                            _ = v74
                                                                        }
                                                                    default:
                                                                        if (!(42 > 1)) {
                                                                            var v75 float32 = float32(v36)
                                                                            _ = v75
                                                                        }
                                                                    }
                                                                    v76 := GenS1{F0: (!true)}
                                                                    _ = v76
                                                                }
                                                            } else {
                                                                v77 := GenS8{F0: 127, F1: 65535}
                                                                if (((!v38) && (1 >= 0)) || false) {
                                                                    v78 := GenS6{F0: 255, F1: (v34 << ((0) & 63))}
                                                                    _ = v78
                                                                }
                                                                if v38 {
                                                                    if (v37 && (v16 == 128)) {
                                                                        fmt.Println(v19)
                                                                    } else {
                                                                        var v79 int64 = (v19 << ((869) & 63))
                                                                        fmt.Println(fi18)
                                                                        _ = v79
                                                                    }
                                                                    if v38 {
                                                                        switch v20 {
                                                                        case "foo bar":
                                                                            v3 = 0.0
                                                                        case "\n\t":
                                                                            var v80 int = -1
                                                                            _ = v80
                                                                        case "hello":
                                                                            var v81 int64 = v19
                                                                            _ = v81
                                                                        default:
                                                                            fmt.Println(v14)
                                                                            v82 := GenS3{F0: 926, F1: (fi35 & fi18)}
                                                                            var v83 int = v5
                                                                            _ = v82
                                                                            _ = v83
                                                                        }
                                                                    }
                                                                    _ = len(v2)
                                                                    v3 -= 0.0
                                                                } else {
                                                                    var v84 string = ("\n\t" + "foo bar")
                                                                    fmt.Println(v14)
                                                                    var v85 uint = 18446744073709551615
                                                                    _ = v84
                                                                    _ = v85
                                                                }
                                                                v19 = v19
                                                                if v14 {
                                                                    var v86 int16 = 32767
                                                                    v87 := GenS8{F0: v34, F1: 42}
                                                                    _ = v86
                                                                    _ = v87
                                                                } else {
                                                                    fmt.Println(v36, v14)
                                                                    switch (fi18 >> ((4294967295) & 63)) {
                                                                    case 4294967295:
                                                                        v38 = (221 <= v19)
                                                                    }
                                                                }
                                                                v88 := v2[(-(fi18))]
                                                                v36 -= 1e10
                                                                _ = v77
                                                                _ = v88
                                                            }
                                                            v2[42] = (((v34 & v34) < 2) || (114 == 100))
                                                            var v89 uint = 65535
                                                            _ = v89
                                                        }
                                                        _ = v38
                                                    } else {
                                                        switch "\n\t" {
                                                        case "":
                                                            var v90 int32 = -2147483648
                                                            _ = v90
                                                        case "a":
                                                            var v91 uint32 = 1
                                                            _ = v91
                                                        case "\n\t":
                                                            switch -935 {
                                                            case -1:
                                                                v92 := GenS7{F0: (fi18 << ((2) & 63))}
                                                                _ = v92
                                                            case 0:
                                                                v20 = ("a" + "a")
                                                            default:
                                                                var v93 uint64 = 450
                                                                _ = v93
                                                            }
                                                        default:
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    v94 := GenS4{F0: 1000, F1: -883, F2: (!(!(v37 || ((!(((v19 * (v19 ^ v19)) != v19) || (122 != 74))) && (224 != -402)))))}
                                                    v5 = fi35
                                                    _ = v94
                                                }
                                                _ = v36
                                                _ = v37
                                            }
                                            v95 := GenS7{F0: (fi35 ^ v5)}
                                            _ = fi35
                                            _ = v95
                                        }
                                        if v14 {
                                            fmt.Println(v5)
                                            if v14 {
                                                v2[42] = (1 <= 1)
                                            }
                                            var v96 int32 = 1
                                            _ = v96
                                        }
                                        _ = v34
                                    }
                                case 4294967295:
                                    switch "hello" {
                                    case "hello":
                                        v16 = (v16 | 121)
                                    case "foo bar":
                                        fmt.Println(v5, v14, v20)
                                    }
                                case 32767:
                                    v97 := v2[(fi18 >> ((0) & 63))]
                                    _ = v97
                                default:
                                    v3 = v3
                                }
                            } else {
                                fmt.Println(v20, fi18, v20)
                            }
                        case "a":
                            fmt.Println(v16, v3)
                        }
                        _ = v32
                    }
                }
                _ = fi18
                _ = v19
                _ = v20
            }
        default:
            for fi98 := 0; fi98 < 5; fi98++ {
                var v99 string = "foo bar"
                _ = fi98
                _ = v99
            }
        }
        if v14 {
            var v100 uint = 18446744073709551615
            v16 = v16
            _ = v100
        } else {
            v16 = (v16 ^ (v16 & v16))
            for fi101 := 0; fi101 < 7; fi101++ {
                var v102 float64 = float64(float32((((0.0 / (3.14159 + 3.14159)) / (1e10 / 1.0)) - (1e10 / 1.5))))
                for fi103 := 0; fi103 < 10; fi103++ {
                    v104 := GenS9{F0: 42, F1: int64(v16)}
                    _ = fi103
                    _ = v104
                }
                _ = fi101
                _ = v102
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    _ = len(v2)
    if v14 {
        v1 = append(v1, 65535)
    } else {
        v16 -= 255
        v3 = float32(1.5)
    }
    v14 = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v5
}
