package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint16; F1 string; F2 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 bool; F2 uint }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 int32; F1 string; F2 uint16 }
type GenS8 struct { F0 int; F1 int64 }
type GenS9 struct { F0 uint16; F1 uint8 }
type GenS10 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS11 struct { F0 uint }
func genFn0() int32 {
    var v5 int32 = -763
    v6 := GenS1{F0: v5, F1: 127}
    _ = v5
    _ = v6
    return 2147483647
}
func genFn1() uint64 {
    if true {
        v7 := GenS1{F0: 1, F1: -128}
        v8 := GenS2{F0: 1, F1: ("hello" + "a"), F2: 88}
        _ = v7
        _ = v8
    } else {
        v9 := GenS2{F0: 918, F1: "a", F2: 0}
        _ = v9
    }
    var v10 uint8 = 1
    for fi11 := 0; fi11 < 3; fi11++ {
        for fi12 := 0; fi12 < 5; fi12++ {
            v13 := GenS3{F0: genFn0()}
            _ = fi12
            _ = v13
        }
        var v14 int8 = -1
        v10 -= v10
        _ = fi11
        _ = v14
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        v10 = v10
        v10 += v10
        _ = fi15
    }
    v16 := GenS4{F0: 2147483647, F1: false, F2: 820}
    _ = v10
    _ = v16
    return 864
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{4294967295}
    v3 := []int{-471}
    var v4 uint32 = 303
    v4 -= (v4 & (v4 % ((0) | 1)))
    fmt.Println(v4)
    for fi17 := 0; fi17 < 17; fi17++ {
        fmt.Println(v4, fi17)
        if len(v2) > 0 {
            v2[0] = (^(v4))
        }
        _ = fi17
    }
    v18 := GenS4{F0: genFn0(), F1: false, F2: 4294967295}
    var v19 float32 = -2.25
    v4 = 669
    fmt.Println(v4)
    if true {
        v2 = append(v2, v4)
    } else {
        if (756 < 835) {
            v20 := GenS5{F0: genFn0()}
            for fi21 := 0; fi21 < 16; fi21++ {
                for fi22 := 0; fi22 < 15; fi22++ {
                    v4 = v4
                    fmt.Println(fi21, fi22, fi22)
                    fmt.Println(v4)
                    if ((v19 + v19) <= v19) {
                        v19 -= -2.25
                        v19 = -2.25
                        if (!false) {
                            var v23 uint16 = 0
                            _ = v23
                        } else {
                            v19 = (v19 - v19)
                            fmt.Println(v4)
                            v24 := v1["a"]
                            _ = v24
                        }
                    } else {
                        v4 += (v4 * (^(v4)))
                        v4 = v4
                    }
                    var v25 uint = 915
                    v25 += v25
                    if true {
                        switch (-(fi21)) {
                        case 32768:
                            var v26 int32 = 42
                            _ = v26
                        default:
                            var v27 int = fi22
                            _ = v27
                        }
                        switch "\n\t" {
                        case "a":
                            v19 += v19
                        case "hello":
                            if (0 >= 18446744073709551615) {
                                var v28 string = ""
                                var v29 uint16 = 960
                                fmt.Println(v25, v19)
                                _ = v28
                                _ = v29
                            }
                            if true {
                                fmt.Println(v25, v25)
                                v30 := GenS6{F0: ("foo bar" + ""), F1: 0}
                                var v31 int8 = -128
                                _ = v30
                                _ = v31
                            } else {
                                var v32 bool = true
                                var v33 int8 = 1
                                _ = v32
                                _ = v33
                            }
                            v19 = v19
                        default:
                            if true {
                                if (!((false || true) && (42 != 316))) {
                                    var v34 uint = (v25 % (((v25 - 4294967295)) | 1))
                                    _ = v34
                                }
                                if true {
                                    v25 += (v25 % (((v25 * 2)) | 1))
                                    v19 -= (v19 / 1e10)
                                    if false {
                                        var v35 int64 = -1
                                        v36 := GenS5{F0: 435}
                                        _ = v35
                                        _ = v36
                                    } else {
                                        v25 = (v25 ^ v25)
                                        switch fi22 {
                                        case 43:
                                            var v37 int = -229
                                            _ = v37
                                        default:
                                            if false {
                                                v4 -= 654
                                                if (((v25 + v25) > (v25 >> ((v25) & 63))) || (32767 == 0)) {
                                                    v25 = 787
                                                    var v38 float64 = float64(v19)
                                                    _ = v38
                                                }
                                                var v39 bool = false
                                                _ = v39
                                            }
                                        }
                                        v19 += (float32(0.1) / v19)
                                    }
                                }
                            }
                        }
                        var v40 int8 = 127
                        _ = v40
                    } else {
                        if true {
                            var v41 int16 = 0
                            _ = v41
                        }
                        var v42 uint64 = 4294967295
                        if true {
                            if ((v25 | uint(v4)) >= v25) {
                                delete(v1, "")
                                switch -2147483648 {
                                case 1:
                                    var v43 int16 = int16(v42)
                                    _ = v43
                                case 32768:
                                    if (!false) {
                                        v42 += genFn1()
                                    } else {
                                        v25 = (^(v25))
                                        if false {
                                            var v44 float64 = (3.14159 * (3.14159 - float64(v19)))
                                            _ = v44
                                        }
                                        v19 = v19
                                    }
                                case -9223372036854775808:
                                    if false {
                                        var v45 int32 = -889
                                        if true {
                                            v46 := GenS5{F0: (v45 << (((^(v25))) & 63))}
                                            var v47 uint64 = v42
                                            _ = v46
                                            _ = v47
                                        } else {
                                            v1[(("\n\t" + "a") + "foo bar")] = 32767
                                            var v48 uint32 = v4
                                            var v49 uint32 = (^(v48))
                                            _ = v48
                                            _ = v49
                                        }
                                        v45 -= (v45 & (v45 | v45))
                                        _ = v45
                                    } else {
                                        var v50 float64 = (1e10 / (0.1 * (0.0 - float64((v19 - v19)))))
                                        v51 := GenS3{F0: 187}
                                        switch 42 {
                                        case 348:
                                            v42 -= v42
                                        case -2147483648:
                                            var v52 int8 = -128
                                            _ = v52
                                        default:
                                            switch fi22 {
                                            case -1:
                                                var v53 int32 = 953
                                                _ = v53
                                            case 128:
                                                var v54 float32 = -2.25
                                                _ = v54
                                            default:
                                                fmt.Println(v4, v42, v42)
                                                fmt.Println(v50, v50, fi22)
                                            }
                                        }
                                        _ = v50
                                        _ = v51
                                    }
                                    var v55 bool = false
                                    switch fi21 {
                                    case -2147483648:
                                        v56 := GenS1{F0: 42, F1: 0}
                                        fmt.Println(fi21, v42)
                                        if false {
                                            v57 := GenS5{F0: 794}
                                            var v58 uint8 = 188
                                            var v59 int64 = int64(v25)
                                            v42 -= uint64(v4)
                                            _ = v57
                                            _ = v58
                                            _ = v59
                                        }
                                        _ = v56
                                    case 841:
                                        switch fi21 {
                                        case 127:
                                            v60 := GenS7{F0: 0, F1: "a", F2: 2}
                                            _ = v60
                                        case -1:
                                            if (0.0 > v19) {
                                                var v61 uint8 = 255
                                                switch (fi21 / ((fi21) | 1)) {
                                                case 128:
                                                    v62 := GenS1{F0: -1, F1: 42}
                                                    _ = v62
                                                case 9223372036854775807:
                                                    if v55 {
                                                        v63 := v1["\n\t"]
                                                        _ = v63
                                                    }
                                                case 251:
                                                    v64 := v1[("a" + "hello")]
                                                    _ = v64
                                                default:
                                                    v65 := GenS7{F0: genFn0(), F1: "foo bar", F2: 1}
                                                    _ = v65
                                                }
                                                switch fi21 {
                                                case 127:
                                                    fmt.Println(v55)
                                                case 4294967295:
                                                    v42 += (v42 % ((v42) | 1))
                                                default:
                                                    v25 -= v25
                                                }
                                                _ = v61
                                            } else {
                                                v66 := GenS8{F0: fi21, F1: -1}
                                                var v67 bool = v55
                                                var v68 uint32 = v4
                                                _ = v66
                                                _ = v67
                                                _ = v68
                                            }
                                        default:
                                            if len(v3) > 0 {
                                                v3[0] = 32768
                                            }
                                        }
                                        fmt.Println(v4, v25)
                                        v69 := GenS9{F0: 255, F1: 128}
                                        _ = v69
                                    case 128:
                                        v42 -= v42
                                    }
                                    _ = v55
                                }
                                delete(v1, ((("hello" + "a") + "foo bar") + "hello"))
                                v19 -= float32((float64(((-2.25 - (3.14159 - v19)) / v19)) + 1e10))
                            }
                        }
                        if len(v2) > 0 {
                            v2[0] = (v4 / ((v4) | 1))
                        }
                        _ = v42
                    }
                    _ = fi22
                    _ = v25
                }
                _ = fi21
            }
            var v70 int8 = 100
            _ = v20
            _ = v70
        }
        if (316 != genFn0()) {
            switch 1 {
            case -24:
                var v71 bool = ((-2147483648 != 42) && (!false))
                _ = v71
            case 102:
                v19 = (v19 - v19)
            }
        }
    }
    if ((genFn1() <= 758) || (1.5 == 1.5)) {
        for fi72 := 0; fi72 < 12; fi72++ {
            fmt.Println(v4)
            _ = fi72
        }
    } else {
        var v73 uint16 = 1
        v19 += -2.25
        _ = v73
    }
    v19 = (((v19 * ((((((v19 + v19) - 1.5) / float32(float64(1e10))) - v19) / (v19 - v19)) - (((v19 - v19) * (3.14159 * ((v19 * (v19 * v19)) / (1.5 - v19)))) * v19))) + (0.1 / v19)) * v19)
    v19 = v19
    v19 -= v19
    for fi74 := 0; fi74 < 12; fi74++ {
        var v75 uint8 = 60
        _ = fi74
        _ = v75
    }
    v76 := GenS10{F0: 243, F1: 1000, F2: 77}
    var v77 float32 = float32(3.14159)
    v78 := GenS11{F0: 4294967295}
    switch "a" {
    case "\n\t":
        var v79 int8 = -1
        _ = v79
    case "":
        v4 += 1
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v76
    _ = v77
    _ = v78
}
