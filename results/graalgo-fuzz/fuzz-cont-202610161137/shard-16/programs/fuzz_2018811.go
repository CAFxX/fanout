package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int64; F2 string }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 uint32; F1 uint }
type GenS5 struct { F0 uint64; F1 int8; F2 int16 }
type GenS6 struct { F0 bool; F1 uint64 }
type GenS7 struct { F0 bool; F1 int32 }
func genFn0(p0 uint32) int16 {
    for fi5 := 0; fi5 < 10; fi5++ {
        v6 := []int{(fi5 << ((661) & 63)), (-(fi5)), (^(fi5))}
        v7 := GenS1{F0: 1}
        _ = fi5
        _ = v6
        _ = v7
    }
    if (!false) {
        if (58 >= 128) {
            if false {
                v8 := GenS2{F0: 540, F1: 1, F2: (("" + "\n\t") + "")}
                v9 := GenS1{F0: 226}
                v10 := []int{-960, 127}
                _ = v10
                _ = v8
                _ = v9
            }
            for fi11 := 0; fi11 < 20; fi11++ {
                v12 := GenS3{F0: 42, F1: "foo bar"}
                for fi13 := 0; fi13 < 12; fi13++ {
                    p0 = 561
                    _ = fi13
                }
                v14 := GenS3{F0: 0, F1: "\n\t"}
                _ = fi11
                _ = v12
                _ = v14
            }
            for fi15 := 0; fi15 < 15; fi15++ {
                fmt.Println(p0, fi15)
                switch fi15 {
                case 32767:
                    p0 -= 2
                case 4294967295:
                    switch fi15 {
                    case -428:
                        var v16 uint = 100
                        _ = v16
                    }
                }
                _ = fi15
            }
        }
        var v17 uint32 = p0
        var v18 float64 = (0.0 + 1e10)
        _ = v17
        _ = v18
    }
    p0 = p0
    _ = p0
    return -1
}
func genFn1(p0 int64) int8 {
    fmt.Println(p0)
    switch "hello" {
    case "a":
        for fi19 := 0; fi19 < 20; fi19++ {
            p0 += p0
            fmt.Println(fi19, fi19)
            _ = fi19
        }
    case "":
        if true {
            var v20 int32 = -2147483648
            _ = v20
        }
    case "foo bar":
        if true {
            v21 := []int{533, -2147483648, -983}
            p0 -= p0
            _ = v21
        }
    }
    v22 := GenS3{F0: -291, F1: "\n\t"}
    _ = p0
    _ = v22
    return 100
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 1
    v3 := []string{""}
    var v4 uint8 = 127
    v3 = append(v3, "")
    for fi23 := 0; fi23 < 7; fi23++ {
        for fi24 := 0; fi24 < 7; fi24++ {
            var v25 uint64 = 4294967295
            if (1 != 42) {
                delete(v1, "hello")
                if false {
                    fmt.Println(fi24, v2, v4)
                } else {
                    _ = len(v1)
                }
                v2 = int(v25)
                var v26 uint16 = 49
                v27 := GenS1{F0: fi23}
                _ = v26
                _ = v27
            } else {
                var v28 string = "\n\t"
                v29 := GenS3{F0: 2147483647, F1: "\n\t"}
                switch v28 {
                case "":
                    var v30 uint32 = 2
                    _ = v30
                case "a":
                    var v31 bool = false
                    if (!(v25 != v25)) {
                        switch -9223372036854775808 {
                        case -569:
                            var v32 uint = 141
                            _ = v32
                        }
                    } else {
                        if v31 {
                            fmt.Println(v25, v4, v28)
                            v25 = uint64(v4)
                            switch fi24 {
                            case -515:
                                var v33 float64 = (float64(-2.25) + 0.0)
                                _ = v33
                            }
                            v34 := GenS1{F0: v2}
                            switch (v2 / (((-(fi24))) | 1)) {
                            case 763:
                                switch (-(v2)) {
                                case 1:
                                    var v35 float32 = 3.14159
                                    _ = v35
                                }
                            }
                            _ = v34
                        }
                        v36 := GenS2{F0: 99, F1: 42, F2: v28}
                        if v31 {
                            var v37 int16 = 1
                            v38 := GenS4{F0: 158, F1: 42}
                            switch (v28 + "a") {
                            case "":
                                var v39 uint64 = 2
                                _ = v39
                            case "hello":
                                fmt.Println(fi24, fi23, fi23)
                            }
                            v2 += (v2 + (v2 >> ((979) & 63)))
                            _ = v37
                            _ = v38
                        } else {
                            if false {
                                var v40 uint = 128
                                switch v28 {
                                case "hello":
                                    if (v31 || (v4 >= (v4 >> (((v40 ^ 42)) & 63)))) {
                                        switch v28 {
                                        case "a":
                                            switch (v28 + v28) {
                                            case "foo bar":
                                                switch (v28 + "foo bar") {
                                                case "hello":
                                                    var v41 uint64 = v25
                                                    _ = v41
                                                case "a":
                                                    var v42 int64 = 42
                                                    _ = v42
                                                default:
                                                    if v31 {
                                                        var v43 uint = 65535
                                                        var v44 int64 = -165
                                                        v4 -= v4
                                                        var v45 float32 = ((0.1 + (3.14159 - float32(0.1))) * (0.0 * ((0.1 * (1.5 + ((float32(float64(3.14159)) * 1.5) / (-2.25 / 3.14159)))) / 0.1)))
                                                        _ = v43
                                                        _ = v44
                                                        _ = v45
                                                    } else {
                                                        var v46 int16 = genFn0(uint32(v40))
                                                        switch v28 {
                                                        case "\n\t":
                                                            if (v40 != v40) {
                                                                switch 542 {
                                                                case -771:
                                                                    v47 := GenS3{F0: 698, F1: v28}
                                                                    _ = v47
                                                                case 2147483647:
                                                                    if (42 == -2147483648) {
                                                                        var v48 int8 = 100
                                                                        var v49 int = v2
                                                                        _ = v48
                                                                        _ = v49
                                                                    } else {
                                                                        v50 := GenS4{F0: 0, F1: (^(v40))}
                                                                        _ = v50
                                                                    }
                                                                case 803:
                                                                    v51 := v1[(v28 + (v28 + ((("foo bar" + v28) + "\n\t") + v28)))]
                                                                    _ = v51
                                                                default:
                                                                    if true {
                                                                        v31 = v31
                                                                        fmt.Println(fi24)
                                                                        v52 := GenS4{F0: 730, F1: (v40 & (v40 - (v40 & v40)))}
                                                                        _ = v52
                                                                    }
                                                                }
                                                            } else {
                                                                v4 = v4
                                                                v46 = (v46 & v46)
                                                            }
                                                        case "a":
                                                            if (!true) {
                                                                if v31 {
                                                                    v46 += v46
                                                                    var v53 uint32 = 65535
                                                                    v54 := GenS5{F0: 292, F1: 20, F2: (v46 * (v46 - v46))}
                                                                    fmt.Println(v25, fi23, fi24)
                                                                    var v55 int16 = -32768
                                                                    _ = v53
                                                                    _ = v54
                                                                    _ = v55
                                                                }
                                                                switch (fi24 | v2) {
                                                                case 128:
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                case 127:
                                                                    v25 -= 926
                                                                default:
                                                                    v46 = -706
                                                                }
                                                                v46 -= -32768
                                                                var v56 uint64 = (^(v25))
                                                                _ = v56
                                                            } else {
                                                                var v57 int32 = -81
                                                                if (0 < 570) {
                                                                    v1["a"] = 2147483647
                                                                    var v58 int32 = (v57 % ((1) | 1))
                                                                    _ = v58
                                                                }
                                                                _ = v57
                                                            }
                                                        case "":
                                                            var v59 int8 = 0
                                                            _ = v59
                                                        }
                                                        _ = v46
                                                    }
                                                }
                                            case "":
                                                _ = len(v1)
                                            case "a":
                                                if len(v3) > 0 {
                                                    v3[0] = v28
                                                }
                                            default:
                                                v60 := GenS5{F0: v25, F1: 42, F2: int16(v40)}
                                                _ = v60
                                            }
                                        default:
                                            fmt.Println(v2, v28, v25)
                                        }
                                        if (((v28 + "a") == "a") && (!(!(0.1 <= 0.0)))) {
                                            v61 := GenS3{F0: -1, F1: v28}
                                            v3 = append(v3, v28)
                                            v62 := GenS2{F0: (^(v40)), F1: 42, F2: "a"}
                                            _ = v61
                                            _ = v62
                                        }
                                        if true {
                                            v63 := GenS3{F0: 42, F1: v28}
                                            _ = v63
                                        } else {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            fmt.Println(v25)
                                        }
                                    } else {
                                        if (v31 || true) {
                                            if (genFn1(-1) == 42) {
                                                var v64 float64 = 0.0
                                                var v65 uint64 = v25
                                                v2 = v2
                                                if (518 > 0) {
                                                    var v66 int64 = -9223372036854775808
                                                    _ = v66
                                                } else {
                                                    if v31 {
                                                        v3 = append(v3, (v28 + (v28 + "")))
                                                    }
                                                }
                                                fmt.Println(v28)
                                                if true {
                                                    fmt.Println(v25, v4, fi24)
                                                }
                                                _ = v64
                                                _ = v65
                                            }
                                        }
                                        var v67 float64 = 0.0
                                        if false {
                                            fmt.Println(v40, v31, v28)
                                            if v31 {
                                                v68 := GenS2{F0: v40, F1: 42, F2: (v28 + v28)}
                                                var v69 uint16 = 1
                                                _ = v68
                                                _ = v69
                                            }
                                            v3 = append(v3, "hello")
                                        } else {
                                            if (!(3.14159 != 0.1)) {
                                                if len(v3) > 0 {
                                                    v3[0] = ((("foo bar" + ((v28 + ((v28 + v28) + v28)) + v28)) + v28) + "\n\t")
                                                }
                                                v67 = v67
                                                var v70 int16 = 32767
                                                _ = v70
                                            }
                                        }
                                        _ = v67
                                    }
                                default:
                                    v71 := GenS3{F0: 172, F1: (v28 + (v28 + v28))}
                                    _ = v71
                                }
                                _ = v40
                            } else {
                                var v72 float64 = 0.1
                                _ = v72
                            }
                            if v31 {
                                var v73 int64 = 42
                                if (v25 == (^(v25))) {
                                    v73 -= v73
                                }
                                _ = v73
                            }
                        }
                        var v74 int8 = genFn1(-1)
                        v31 = (!v31)
                        _ = v36
                        _ = v74
                    }
                    fmt.Println(v31, v31)
                    _ = v31
                case "\n\t":
                    if false {
                        _ = len(v3)
                    } else {
                        if false {
                            if false {
                                fmt.Println(v25, v25, v25)
                                var v75 float32 = float32(1e10)
                                _ = v75
                            }
                            var v76 float32 = 0.0
                            _ = v76
                        } else {
                            fmt.Println(v25, fi23)
                            v28 = (v28 + "")
                        }
                        if (!true) {
                            fmt.Println(fi23, fi23, v2)
                            v77 := GenS6{F0: (!(-996 != -2147483648)), F1: (^(v25))}
                            _ = v77
                        }
                        var v78 uint8 = (^(v4))
                        _ = v78
                    }
                default:
                    v79 := GenS1{F0: fi24}
                    _ = v79
                }
                _ = v28
                _ = v29
            }
            v80 := GenS7{F0: true, F1: 42}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi24
            _ = v25
            _ = v80
        }
        if (-281 != 2147483647) {
            switch "" {
            case "a":
                switch (fi23 % ((v2) | 1)) {
                case -627:
                    var v81 int = (fi23 | int(v4))
                    _ = v81
                case 0:
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                case 9223372036854775807:
                    var v82 uint = 42
                    _ = v82
                default:
                    for fi83 := 0; fi83 < 3; fi83++ {
                        var v84 uint16 = 32767
                        if (647 == 4294967295) {
                            var v85 int16 = 1000
                            var v86 uint32 = 1
                            var v87 uint = 1
                            _ = v85
                            _ = v86
                            _ = v87
                        }
                        _ = fi83
                        _ = v84
                    }
                }
            default:
                for fi88 := 0; fi88 < 14; fi88++ {
                    if (!true) {
                        var v89 uint8 = v4
                        _ = v89
                    } else {
                        v1[(("hello" + "a") + "foo bar")] = 2147483647
                        var v90 float64 = 3.14159
                        _ = v90
                    }
                    delete(v1, "foo bar")
                    fmt.Println(fi88, v4, v4)
                    _ = fi88
                }
            }
        }
        _ = fi23
    }
    v91 := GenS1{F0: 9223372036854775807}
    switch "a" {
    case "\n\t":
        v4 = 128
    }
    if len(v3) > 0 {
        v3[0] = ("foo bar" + ("hello" + ((("\n\t" + "foo bar") + ("" + "hello")) + "foo bar")))
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch (v2 / ((v2) | 1)) {
    case 2147483647:
        v1["a"] = 2
    }
    delete(v1, "\n\t")
    fmt.Println(v4, v2)
    fmt.Println(v4, v4)
    v2 -= -910
    v92 := GenS2{F0: 18446744073709551615, F1: -9223372036854775808, F2: "foo bar"}
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v91
    _ = v92
}
