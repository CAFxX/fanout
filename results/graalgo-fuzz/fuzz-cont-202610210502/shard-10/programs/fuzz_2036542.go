package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 int8; F2 string }
type GenS5 struct { F0 uint32; F1 int64; F2 int }
func genFn0(p0 int16) uint8 {
    if false {
        p0 += p0
    } else {
        v6 := GenS1{F0: -128, F1: 728, F2: 9223372036854775807}
        var v7 int64 = 9223372036854775807
        switch 1 {
        case 2147483647:
            switch 182 {
            case 128:
                v8 := GenS1{F0: -128, F1: 18446744073709551615, F2: (v7 * (v7 >> ((1) & 63)))}
                _ = v8
            case 491:
                var v9 int16 = (^(p0))
                _ = v9
            default:
                fmt.Println(v7)
            }
        }
        _ = v6
        _ = v7
    }
    p0 = (p0 & (p0 >> ((2) & 63)))
    var v10 int64 = -1
    v11 := []int{32767, 127, 4294967295, 815}
    _ = p0
    _ = v10
    _ = v11
    return 142
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = 0
    v3 := make(map[string]string)
    v4 := make(map[int]string)
    var v5 uint8 = 128
    v5 -= v5
    if (true && ("" >= "\n\t")) {
        v12 := GenS1{F0: 127, F1: 801, F2: 9223372036854775807}
        v13 := GenS1{F0: 1, F1: 197, F2: -9223372036854775808}
        if false {
            fmt.Println(v2, v5)
            fmt.Println(v2, v5)
        } else {
            if false {
                v14 := []int{(v2 >> ((295) & 63)), v2}
                if (!(true || true)) {
                    v5 = (v5 >> ((267) & 63))
                }
                var v15 uint64 = 814
                _ = v14
                _ = v15
            }
            v16 := GenS1{F0: 42, F1: 2, F2: -9223372036854775808}
            if false {
                var v17 uint = 0
                for fi18 := 0; fi18 < 8; fi18++ {
                    var v19 float32 = 1.5
                    if (((("hello" + "foo bar") + ("" + "\n\t")) <= "foo bar") || true) {
                        v17 = v17
                        switch "" {
                        case "hello":
                            var v20 string = "foo bar"
                            _ = v20
                        case "a":
                            if (!(false || (false && true))) {
                                if true {
                                    switch fi18 {
                                    case 87:
                                        v2 -= fi18
                                    case 32767:
                                        for fi21 := 0; fi21 < 12; fi21++ {
                                            v4[fi18] = "hello"
                                            _ = fi21
                                        }
                                    case 32768:
                                        v22 := GenS1{F0: 49, F1: 157, F2: 1}
                                        _ = v22
                                    default:
                                        if true {
                                            for fi23 := 0; fi23 < 17; fi23++ {
                                                v5 = (v5 % (((v5 / (((^(v5))) | 1))) | 1))
                                                _ = fi23
                                            }
                                            var v24 int16 = -1
                                            var v25 uint8 = (v5 + genFn0(v24))
                                            _ = v24
                                            _ = v25
                                        }
                                    }
                                    v2 -= fi18
                                }
                            }
                        }
                    }
                    _ = fi18
                    _ = v19
                }
                _ = v17
            } else {
                v5 += 255
            }
            _ = v16
        }
        _ = v12
        _ = v13
    }
    for fi26 := 0; fi26 < 10; fi26++ {
        var v27 string = ("a" + "\n\t")
        v27 = (("\n\t" + v27) + v27)
        if (2 > 880) {
            _ = len(v1)
            if (0 == 65535) {
                v28 := GenS1{F0: 127, F1: 2, F2: -110}
                _ = v28
            }
            v29 := GenS1{F0: -12, F1: 0, F2: -289}
            _ = v29
        } else {
            if ((fi26 % ((v2) | 1)) >= fi26) {
                var v30 bool = (9223372036854775807 == 0)
                fmt.Println(v2, v27)
                if (!(663 <= 42)) {
                    var v31 uint8 = v5
                    switch 9223372036854775807 {
                    case 1:
                        v1[2] = v30
                    default:
                        if (uint16(v31) >= 70) {
                            var v32 int = -776
                            var v33 int = (v2 % ((v2) | 1))
                            v27 = v27
                            _ = v32
                            _ = v33
                        }
                    }
                    _ = v31
                } else {
                    var v34 float64 = (1e10 / 3.14159)
                    _ = v34
                }
                _ = v30
            }
        }
        _ = fi26
        _ = v27
    }
    for fi35 := 0; fi35 < 1; fi35++ {
        if true {
            var v36 int = v2
            for fi37 := 0; fi37 < 19; fi37++ {
                v38 := GenS1{F0: 40, F1: 65535, F2: 42}
                var v39 int8 = -107
                v40 := GenS2{F0: v39}
                fmt.Println(v2, v2)
                if (!false) {
                    v41 := v3["foo bar"]
                    if true {
                        if true {
                            fmt.Println(v39)
                        }
                        v4[v36] = ("\n\t" + "hello")
                        v42 := []int{42, v36, (fi37 >> ((18446744073709551615) & 63)), -378}
                        _ = v42
                    } else {
                        v43 := GenS1{F0: (v39 / ((100) | 1)), F1: 477, F2: 1}
                        v2 += v36
                        _ = v43
                    }
                    v44 := v4[(fi35 % ((-264) | 1))]
                    _ = v41
                    _ = v44
                }
                v45 := GenS3{F0: (^(v5))}
                var v46 int16 = int16(v5)
                v47 := GenS4{F0: (v39 * (v39 >> ((2) & 63))), F1: (^(v39)), F2: "\n\t"}
                _ = fi37
                _ = v38
                _ = v39
                _ = v40
                _ = v45
                _ = v46
                _ = v47
            }
            _ = v36
        } else {
            v48 := GenS1{F0: 100, F1: 255, F2: 9223372036854775807}
            var v49 int8 = 39
            _ = v48
            _ = v49
        }
        switch ("a" + "foo bar") {
        case "hello":
            var v50 int32 = 0
            _ = v50
        case "\n\t":
            v2 = 32767
        case "a":
            var v51 uint16 = 0
            _ = v51
        default:
            if (0.0 == float32(3.14159)) {
                v52 := GenS2{F0: -1}
                for fi53 := 0; fi53 < 1; fi53++ {
                    switch (^(v2)) {
                    case 0:
                        var v54 int32 = -617
                        _ = v54
                    default:
                        fmt.Println(fi35)
                    }
                    _ = fi53
                }
                if false {
                    for fi55 := 0; fi55 < 2; fi55++ {
                        v4[fi55] = ""
                        switch (v2 / (((fi55 - 0)) | 1)) {
                        case 330:
                            if (!(true && true)) {
                                var v56 float64 = (1.5 * (float64((3.14159 / 0.1)) - (1e10 - ((float64(1e10) * 1.5) / (0.1 / 1.0)))))
                                _ = v56
                            } else {
                                if false {
                                    v2 -= 42
                                    v2 = fi35
                                    v57 := GenS4{F0: 100, F1: -14, F2: "hello"}
                                    delete(v1, 2147483647)
                                    switch ((("a" + ("hello" + "hello")) + ("\n\t" + "\n\t")) + "\n\t") {
                                    case "foo bar":
                                        if (42 != 0) {
                                            var v58 float32 = float32(3.14159)
                                            v5 = (v5 % ((0) | 1))
                                            _ = v58
                                        } else {
                                            var v59 uint16 = 0
                                            var v60 uint32 = 42
                                            var v61 string = "foo bar"
                                            v5 = (v5 + 36)
                                            v3[v61] = (v61 + (v61 + "foo bar"))
                                            _ = v59
                                            _ = v60
                                            _ = v61
                                        }
                                    default:
                                        if false {
                                            if true {
                                                v2 = -328
                                            }
                                            fmt.Println(fi35, v5, v2)
                                            var v62 bool = ((((-1 >= 42) || false) || false) || false)
                                            v63 := []int{}
                                            v64 := GenS1{F0: 40, F1: 301, F2: 1}
                                            var v65 float64 = 0.1
                                            _ = v62
                                            _ = v63
                                            _ = v64
                                            _ = v65
                                        } else {
                                            var v66 float64 = float64(0.0)
                                            v67 := []int{fi55, (fi55 / ((fi35) | 1)), fi55, (fi55 * -2147483648)}
                                            var v68 int64 = -138
                                            _ = v66
                                            _ = v67
                                            _ = v68
                                        }
                                    }
                                    switch ("" + "") {
                                    case "a":
                                        var v69 uint64 = 821
                                        v70 := []int{(fi55 >> ((398) & 63)), fi55, v2}
                                        _ = v69
                                        _ = v70
                                    case "\n\t":
                                        fmt.Println(v2, fi55, v5)
                                    default:
                                        v71 := []int{fi35, -2147483648, fi35, fi35}
                                        _ = v71
                                    }
                                    v72 := []int{32768, (fi35 * (fi35 % ((fi55) | 1)))}
                                    _ = v57
                                    _ = v72
                                } else {
                                    var v73 int64 = -1
                                    var v74 uint16 = 0
                                    v5 = 128
                                    _ = v73
                                    _ = v74
                                }
                                var v75 uint16 = 252
                                v76 := GenS1{F0: -76, F1: 0, F2: 144}
                                _ = v75
                                _ = v76
                            }
                        }
                        if ((v5 * (v5 - v5)) == genFn0(1000)) {
                            var v77 string = "\n\t"
                            var v78 int8 = -128
                            var v79 int64 = 598
                            var v80 int64 = (v79 * 0)
                            var v81 uint32 = 1
                            _ = v77
                            _ = v78
                            _ = v79
                            _ = v80
                            _ = v81
                        }
                        _ = fi55
                    }
                } else {
                    fmt.Println(v2, v2, v5)
                    var v82 int = (fi35 ^ (v2 % ((fi35) | 1)))
                    _ = v82
                }
                _ = v52
            } else {
                for fi83 := 0; fi83 < 1; fi83++ {
                    v5 -= 255
                    v5 += v5
                    var v84 int32 = 435
                    _ = fi83
                    _ = v84
                }
                var v85 uint16 = 1
                var v86 int64 = -149
                _ = v85
                _ = v86
            }
        }
        _ = fi35
    }
    if (42 >= 230) {
        switch (("hello" + ("\n\t" + "")) + ("\n\t" + "foo bar")) {
        case "a":
            for fi87 := 0; fi87 < 1; fi87++ {
                v2 -= fi87
                v88 := GenS5{F0: 678, F1: 634, F2: -342}
                _ = fi87
                _ = v88
            }
        }
    }
    var v89 uint16 = 32767
    if true {
        v89 -= v89
        v90 := []int{v2, (^(v2))}
        _ = v90
    } else {
        v2 -= -81
    }
    for fi91 := 0; fi91 < 16; fi91++ {
        var v92 float64 = -2.25
        _ = fi91
        _ = v92
    }
    v89 += v89
    v93 := []int{}
    for fi94 := 0; fi94 < 12; fi94++ {
        if (1 > 4294967295) {
            for fi95 := 0; fi95 < 19; fi95++ {
                var v96 int8 = 115
                var v97 float64 = (3.14159 - 1.5)
                _ = fi95
                _ = v96
                _ = v97
            }
        } else {
            var v98 uint = 65535
            var v99 float64 = 0.0
            var v100 int32 = 1
            _ = v100
            _ = v98
            _ = v99
        }
        v93 = append(v93, fi94)
        v93 = append(v93, (fi94 + (v2 % ((v2) | 1))))
        _ = fi94
    }
    var v101 int8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v101:", v101)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v89:", v89)
    fmt.Println("slicelen v93:", len(v93), v93)
    _ = v1
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v89
    _ = v93
}
