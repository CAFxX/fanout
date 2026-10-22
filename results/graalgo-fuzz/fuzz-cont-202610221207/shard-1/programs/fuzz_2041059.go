package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8 }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[string]int)
    var v3 int16 = 509
    var v4 uint64 = 18446744073709551615
    var v5 int16 = v3
    switch v1 {
    case "":
        v6 := GenS1{F0: ((true && (-2147483648 < 32768)) && false), F1: 442}
        _ = v6
    case "a":
        var v7 string = (v1 + ("\n\t" + "\n\t"))
        _ = v7
    case "hello":
        for fi8 := 0; fi8 < 11; fi8++ {
            fmt.Println(v4)
            _ = fi8
        }
    }
    v9 := GenS1{F0: false, F1: -392}
    v10 := []int{-696, -786}
    v10 = append(v10, -9223372036854775808)
    if true {
        if len(v10) > 0 {
            v10[0] = 473
        }
    } else {
        v11 := GenS1{F0: (false || (uint16(v4) < 32767)), F1: 1}
        for fi12 := 0; fi12 < 5; fi12++ {
            fmt.Println(v1)
            _ = fi12
        }
        var v13 float64 = 3.14159
        _ = v11
        _ = v13
    }
    var v14 uint64 = 816
    if false {
        var v15 int16 = v5
        for fi16 := 0; fi16 < 16; fi16++ {
            fmt.Println(v1, v4)
            _ = fi16
        }
        _ = v15
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        for fi18 := 0; fi18 < 12; fi18++ {
            v14 = v14
            _ = fi18
        }
        _ = fi17
    }
    var v19 float64 = 3.14159
    v20 := GenS1{F0: (false && false), F1: -695}
    v21 := v2[(((((v1 + "a") + "hello") + v1) + v1) + "foo bar")]
    if (false && false) {
        v19 = (((((float64(1e10) * 3.14159) / 1e10) / v19) - v19) - v19)
        v3 += v3
        if ((^(v4)) == v14) {
            if (v21 != (^(v21))) {
                var v22 float32 = 0.1
                var v23 uint = 255
                switch (v21 << ((255) & 63)) {
                case -659:
                    fmt.Println(v14, v23)
                }
                _ = v22
                _ = v23
            } else {
                if (v4 > (v14 & (^(v4)))) {
                    fmt.Println(v4, v3, v21)
                    fmt.Println(v21, v14)
                    var v24 uint8 = 255
                    _ = v24
                }
                var v25 int32 = 1
                v19 = v19
                _ = v25
            }
        }
    } else {
        v4 += v4
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        for fi27 := 0; fi27 < 15; fi27++ {
            var v28 bool = (!false)
            var v29 int8 = -122
            switch v1 {
            case "foo bar":
                switch fi27 {
                case 39:
                    if v28 {
                        v30 := v2[(v1 + v1)]
                        _ = v30
                    }
                case 128:
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                default:
                    v31 := GenS1{F0: v28, F1: -663}
                    v21 -= (fi27 * 693)
                    var v32 string = ((("\n\t" + "\n\t") + ("a" + ((v1 + (v1 + (("\n\t" + "hello") + v1))) + (v1 + v1)))) + (v1 + v1))
                    _ = v31
                    _ = v32
                }
            case "a":
                var v33 uint8 = 127
                _ = v33
            case "\n\t":
                if (!(0 == -288)) {
                    v29 += 27
                } else {
                    _ = len(v10)
                    var v34 int64 = 42
                    _ = v34
                }
            }
            _ = fi27
            _ = v28
            _ = v29
        }
        for fi35 := 0; fi35 < 10; fi35++ {
            if true {
                v10 = append(v10, 969)
                if (v19 > v19) {
                    v14 += 2
                } else {
                    _ = len(v2)
                    v36 := GenS2{F0: 1}
                    v37 := GenS2{F0: 42}
                    v38 := v2["\n\t"]
                    var v39 int8 = -1
                    var v40 uint = 4294967295
                    _ = v36
                    _ = v37
                    _ = v38
                    _ = v39
                    _ = v40
                }
                delete(v2, v1)
            }
            switch (^(v21)) {
            case -2147483648:
                var v41 uint = 42
                _ = v41
            case -154:
                v21 = 32767
            default:
                v42 := v2[("foo bar" + (v1 + v1))]
                _ = v42
            }
            v2["hello"] = v21
            v43 := GenS2{F0: 2}
            switch (-(fi26)) {
            case 112:
                if (1e10 >= -2.25) {
                    if (1.5 != float64(0.1)) {
                        v44 := GenS3{F0: fi26, F1: -9223372036854775808}
                        v4 = v14
                        var v45 int = 128
                        _ = v44
                        _ = v45
                    }
                    delete(v2, "\n\t")
                } else {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    v19 += 1.5
                    v14 = (v4 % ((v14) | 1))
                    if (-846 > -2147483648) {
                        fmt.Println(v19, v19, v1)
                        fmt.Println(v19, v21, v21)
                    }
                    if (1 > 5) {
                        var v46 float32 = -2.25
                        v46 = (v46 + (3.14159 * (v46 / (0.1 + (v46 * (v46 + v46))))))
                        _ = v46
                    } else {
                        fmt.Println(v1, v14)
                        var v47 uint8 = 0
                        switch 2147483647 {
                        case -9223372036854775808:
                            var v48 int8 = -82
                            _ = v48
                        case -963:
                            v10 = append(v10, 2147483647)
                            var v49 int8 = -128
                            var v50 uint16 = 154
                            var v51 uint8 = 128
                            var v52 uint8 = v47
                            v53 := GenS3{F0: (v21 + 346), F1: -888}
                            fmt.Println(v1)
                            v51 -= 116
                            _ = v49
                            _ = v50
                            _ = v51
                            _ = v52
                            _ = v53
                        case -126:
                            v54 := GenS3{F0: (fi35 % (((fi26 % ((127) | 1))) | 1)), F1: 0}
                            _ = v54
                        default:
                            var v55 uint16 = 65535
                            _ = v55
                        }
                        _ = v47
                    }
                }
            case -905:
                var v56 int32 = 0
                _ = v56
            }
            var v57 uint = 18446744073709551615
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = fi35
            _ = v43
            _ = v57
        }
        _ = fi26
    }
    v19 = v19
    if (42 < 0) {
        for fi58 := 0; fi58 < 12; fi58++ {
            for fi59 := 0; fi59 < 8; fi59++ {
                v60 := GenS4{F0: 42}
                var v61 int16 = (v3 + -1)
                var v62 int32 = 2147483647
                if ((true || false) && true) {
                    var v63 bool = false
                    v63 = false
                    _ = v63
                } else {
                    var v64 int32 = 42
                    fmt.Println(v3, v1, v62)
                    _ = v64
                }
                if false {
                    if (!(!(!true))) {
                        _ = len(v10)
                        var v65 uint32 = 42
                        _ = v65
                    } else {
                        v1 = v1
                        v66 := GenS4{F0: 2}
                        _ = v66
                    }
                } else {
                    fmt.Println(v21)
                    var v67 uint32 = 1
                    _ = v67
                }
                _ = fi59
                _ = v60
                _ = v61
                _ = v62
            }
            switch "hello" {
            case "a":
                v68 := GenS3{F0: -238, F1: 626}
                _ = v68
            case "foo bar":
                v5 -= v3
            case "hello":
                for fi69 := 0; fi69 < 9; fi69++ {
                    v1 = "foo bar"
                    var v70 float64 = v19
                    v14 -= v14
                    _ = fi69
                    _ = v70
                }
            default:
                var v71 int64 = 9223372036854775807
                _ = v71
            }
            if (0.1 >= 3.14159) {
                if true {
                    v72 := v2[v1]
                    var v73 float32 = (((((float32((v19 / ((1.5 + 0.1) / (v19 / float64(-2.25))))) - float32(((v19 * v19) - v19))) / (3.14159 - (float32((0.0 - 3.14159)) * 3.14159))) * 0.0) / (0.0 / float32(v19))) / 1e10)
                    _ = v72
                    _ = v73
                }
                v19 += v19
            }
            _ = fi58
        }
    } else {
        var v74 uint8 = 42
        v19 = float64(1.5)
        _ = v74
    }
    var v75 int = v21
    for fi76 := 0; fi76 < 1; fi76++ {
        if (127 < -128) {
            var v77 int32 = 1
            _ = v77
        }
        if false {
            switch (("foo bar" + v1) + v1) {
            case "hello":
                var v78 uint64 = v4
                _ = v78
            case "a":
                fmt.Println(v5)
            case "foo bar":
                v79 := GenS5{F0: 1}
                _ = v79
            }
        } else {
            if (656 < 157) {
                var v80 string = (v1 + v1)
                v2[("foo bar" + ("a" + ""))] = (fi76 % (((fi76 + (^(v75)))) | 1))
                _ = v80
            } else {
                v81 := GenS4{F0: 25}
                _ = v81
            }
            v5 -= (v5 / ((32767) | 1))
        }
        switch ("" + v1) {
        case "a":
            v10 = append(v10, fi76)
        case "foo bar":
            var v82 int = 42
            _ = v82
        case "hello":
            if true {
                for fi83 := 0; fi83 < 2; fi83++ {
                    v84 := GenS2{F0: 2}
                    var v85 float64 = 1e10
                    _ = fi83
                    _ = v84
                    _ = v85
                }
                fmt.Println(fi76)
            }
        }
        _ = fi76
    }
    var v86 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v75:", v75)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v10
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v75
    _ = v86
    _ = v9
}
