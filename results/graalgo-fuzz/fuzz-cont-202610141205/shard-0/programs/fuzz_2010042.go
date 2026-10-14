package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 uint16; F1 uint8 }
func genFn0(p0 string, p1 uint64, p2 uint8) bool {
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := []int{(fi8 & (fi7 * 42)), fi7, 272}
            switch (fi8 & fi8) {
            case 1:
                var v10 int = fi8
                _ = v10
            case 128:
                if len(v9) > 0 {
                    v9[0] = fi7
                }
            case -130:
                var v11 uint16 = 2
                _ = v11
            default:
                var v12 uint32 = 2147483647
                _ = v12
            }
            var v13 int8 = 1
            v14 := GenS1{F0: 4294967295}
            v13 += v13
            _ = fi8
            _ = v13
            _ = v14
            _ = v9
        }
        _ = fi7
    }
    var v15 uint16 = 255
    if (-9223372036854775808 < -745) {
        p2 = uint8(v15)
        var v16 float64 = 0.0
        var v17 int = 1
        _ = v16
        _ = v17
    } else {
        v18 := []int{-1}
        _ = v18
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        v20 := []int{127, -9223372036854775808}
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        v15 = (v15 >> ((0) & 63))
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        p1 = p1
        v22 := GenS1{F0: 65535}
        var v23 int = fi21
        _ = fi21
        _ = v22
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return false
}
func genFn1() int64 {
    var v24 int8 = -128
    var v25 bool = true
    _ = v24
    _ = v25
    return -9223372036854775808
}

func main() {
    v1 := []string{"", "\n\t"}
    v2 := make(map[uint32]int)
    var v3 bool = true
    var v4 int16 = 926
    var v5 uint8 = 128
    var v6 int32 = 42
    v26 := v2[65535]
    for fi27 := 0; fi27 < 12; fi27++ {
        v3 = v3
        switch "foo bar" {
        case "a":
            switch ("\n\t" + "hello") {
            case "\n\t":
                fmt.Println(v3)
            case "foo bar":
                var v28 uint = 502
                _ = v28
            case "":
                if v3 {
                    fmt.Println(v26, v5, v3)
                    var v29 uint64 = 2
                    _ = v29
                } else {
                    v30 := GenS1{F0: 0}
                    _ = v30
                }
            default:
                for fi31 := 0; fi31 < 11; fi31++ {
                    v32 := GenS1{F0: 1}
                    v33 := GenS2{F0: 49}
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            }
        case "hello":
            var v34 float64 = 1.5
            _ = v34
        case "\n\t":
            if v3 {
                if (v3 && v3) {
                    if len(v1) > 0 {
                        v1[0] = ("" + ("hello" + "foo bar"))
                    }
                    v6 += (v6 ^ (^(v6)))
                } else {
                    var v35 string = ("" + "hello")
                    var v36 bool = (814 == 420)
                    _ = v35
                    _ = v36
                }
                if (v3 || (!false)) {
                    v37 := GenS2{F0: (v5 % ((2) | 1))}
                    _ = v37
                }
            } else {
                var v38 uint16 = 42
                _ = len(v2)
                for fi39 := 0; fi39 < 12; fi39++ {
                    switch "hello" {
                    case "foo bar":
                        switch (("\n\t" + "\n\t") + "hello") {
                        case "foo bar":
                            v26 = int(v4)
                        case "\n\t":
                            if (42 >= (v38 % ((v38) | 1))) {
                                fmt.Println(fi27, v3)
                                v38 += v38
                                _ = len(v1)
                            }
                        case "a":
                            fmt.Println(fi27)
                        default:
                            v40 := GenS3{F0: (v4 | (v4 ^ (v4 ^ -1))), F1: v4}
                            _ = v40
                        }
                    }
                    v41 := GenS2{F0: (v5 ^ 42)}
                    _ = fi39
                    _ = v41
                }
                _ = v38
            }
        default:
            var v42 string = "foo bar"
            _ = v42
        }
        _ = fi27
    }
    for fi43 := 0; fi43 < 7; fi43++ {
        var v44 int = 2147483647
        _ = fi43
        _ = v44
    }
    for fi45 := 0; fi45 < 8; fi45++ {
        for fi46 := 0; fi46 < 9; fi46++ {
            if (!v3) {
                if (-1 >= 48) {
                    v3 = (v3 && (!(!(v4 < (^(v4))))))
                }
            }
            v47 := GenS3{F0: -64, F1: (^(v4))}
            fmt.Println(v4, v3)
            _ = fi46
            _ = v47
        }
        v48 := v2[65535]
        _ = fi45
        _ = v48
    }
    var v49 bool = ((v26 * (^(v26))) > (v26 / (((v26 & v26)) | 1)))
    var v50 uint8 = v5
    v2[1] = v26
    v49 = (42 == genFn1())
    for fi51 := 0; fi51 < 3; fi51++ {
        switch "" {
        case "":
            _ = len(v2)
        case "a":
            if v3 {
                var v52 uint16 = 1
                _ = v52
            }
        default:
            for fi53 := 0; fi53 < 19; fi53++ {
                fmt.Println(v50, v50)
                _ = fi53
            }
        }
        _ = fi51
    }
    var v54 uint64 = 461
    switch v26 {
    case 233:
        var v55 int64 = 386
        _ = v55
    case -1:
        for fi56 := 0; fi56 < 18; fi56++ {
            if false {
                switch fi56 {
                case -282:
                    fmt.Println(fi56, v6, v49)
                case -637:
                    var v57 int = v26
                    _ = v57
                }
            } else {
                var v58 int8 = 100
                for fi59 := 0; fi59 < 11; fi59++ {
                    v5 = (v5 >> ((4294967295) & 63))
                    var v60 float32 = float32(0.1)
                    _ = fi59
                    _ = v60
                }
                if ((v54 | v54) != (^(v54))) {
                    for fi61 := 0; fi61 < 15; fi61++ {
                        v6 -= v6
                        delete(v2, 2147483648)
                        v62 := v2[0]
                        _ = fi61
                        _ = v62
                    }
                    for fi63 := 0; fi63 < 4; fi63++ {
                        switch (fi63 >> ((967) & 63)) {
                        case -2147483648:
                            switch v26 {
                            case 4294967295:
                                if (true || false) {
                                    fmt.Println(v50)
                                }
                            case 1:
                                v64 := GenS3{F0: v4, F1: v4}
                                _ = v64
                            }
                        case 1:
                            if v49 {
                                if v49 {
                                    var v65 int16 = (v4 << ((756) & 63))
                                    v58 = (v58 ^ (v58 * -128))
                                    v66 := GenS4{F0: v54, F1: v50}
                                    _ = v65
                                    _ = v66
                                } else {
                                    switch "" {
                                    case "\n\t":
                                        fmt.Println(v5, v5)
                                    default:
                                        var v67 int64 = -584
                                        _ = v67
                                    }
                                }
                                switch ("" + (("hello" + ("hello" + "a")) + "a")) {
                                case "":
                                    fmt.Println(v26)
                                case "\n\t":
                                    if ((v54 | v54) != v54) {
                                        var v68 bool = true
                                        if true {
                                            var v69 int16 = (v4 & (v4 * 42))
                                            _ = v69
                                        } else {
                                            var v70 float32 = (((-2.25 - 3.14159) - 1.5) * 0.0)
                                            _ = v70
                                        }
                                        if true {
                                            v71 := v2[65535]
                                            fmt.Println(v26, v3)
                                            v71 = v71
                                            v72 := GenS5{F0: 42, F1: 232}
                                            if v3 {
                                                _ = len(v2)
                                                v3 = (false || false)
                                            } else {
                                                switch "\n\t" {
                                                case "hello":
                                                    _ = len(v2)
                                                default:
                                                    _ = len(v2)
                                                }
                                            }
                                            _ = v71
                                            _ = v72
                                        } else {
                                            v50 += v50
                                            v26 -= fi63
                                        }
                                        _ = v68
                                    } else {
                                        if (65535 != 15) {
                                            v54 += v54
                                            var v73 uint64 = (^(v54))
                                            _ = len(v2)
                                            _ = v73
                                        }
                                        switch "foo bar" {
                                        case "a":
                                            var v74 string = ""
                                            _ = v74
                                        }
                                    }
                                    fmt.Println(v50, v58, v26)
                                    var v75 int8 = v58
                                    _ = v75
                                case "foo bar":
                                    if v49 {
                                        v3 = false
                                        v58 += (v58 ^ (v58 | v58))
                                        var v76 uint = 49
                                        _ = v76
                                    }
                                }
                                v2[2] = (fi56 % (((fi56 - 32767)) | 1))
                            }
                        case -9223372036854775808:
                            v58 += (v58 | (v58 - (-(v58))))
                        }
                        switch (fi63 + 924) {
                        case 127:
                            v77 := v2[2147483648]
                            _ = v77
                        case -304:
                            v58 = 109
                        default:
                            fmt.Println(v49, v50, fi56)
                        }
                        var v78 float32 = float32((-2.25 - (1e10 + 0.0)))
                        v49 = v49
                        var v79 uint64 = 615
                        switch (("hello" + ("" + "\n\t")) + "a") {
                        case "":
                            v5 += v50
                        case "a":
                            switch (fi56 % (((v26 * v26)) | 1)) {
                            case 9223372036854775807:
                                v1 = append(v1, "a")
                            case -269:
                                var v80 float32 = (v78 * 3.14159)
                                _ = v80
                            default:
                                var v81 bool = ((-(v4)) > v4)
                                _ = v81
                            }
                        case "hello":
                            var v82 uint16 = 65535
                            _ = v82
                        }
                        var v83 int = fi63
                        v1 = append(v1, "foo bar")
                        var v84 float32 = v78
                        _ = fi63
                        _ = v78
                        _ = v79
                        _ = v83
                        _ = v84
                    }
                    v6 += (-(v6))
                }
                _ = v58
            }
            var v85 uint = 895
            switch int(v54) {
            case 32767:
                v6 -= v6
            }
            _ = fi56
            _ = v85
        }
    }
    v3 = (!v49)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v54
    _ = v6
}
