package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 string }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 bool }
func genFn0(p0 bool) int32 {
    v6 := []int{-641, -545, 1}
    var v7 uint16 = 0
    v8 := GenS1{F0: -1, F1: 127, F2: "\n\t"}
    var v9 int64 = 571
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return -2147483648
}
func genFn1(p0 int, p1 uint16, p2 int16) uint64 {
    var v10 string = "foo bar"
    var v11 int64 = 1
    var v12 float32 = 3.14159
    v13 := GenS1{F0: v11, F1: uint8(p1), F2: ((v10 + "\n\t") + "hello")}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return 0
}

func main() {
    var v1 int32 = -724
    var v2 uint32 = 2147483648
    var v3 float32 = (((-2.25 - (3.14159 - (1.5 + -2.25))) / 1e10) - 0.1)
    var v4 int32 = v1
    var v5 int8 = 0
    var v14 int16 = -976
    var v15 uint64 = 42
    if (!(true && (float64((v3 / ((v3 * 3.14159) * (v3 * (v3 * v3))))) < ((1.5 * ((-2.25 * 3.14159) + 1.5)) * -2.25)))) {
        v2 = (^(v2))
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 float64 = (-2.25 * ((1e10 * (1e10 / -2.25)) / (1e10 * (((0.1 / ((((float64(float32((0.0 * 1.5))) * (float64(v3) / -2.25)) * float64(v3)) / 1.0) - 3.14159)) - 3.14159) + 0.1))))
            _ = fi16
            _ = v17
        }
    } else {
        var v18 float64 = 1e10
        for fi19 := 0; fi19 < 18; fi19++ {
            var v20 int32 = genFn0(true)
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 10; fi21++ {
            for fi22 := 0; fi22 < 4; fi22++ {
                v23 := []int{fi22, (fi21 & -9223372036854775808), (-(fi21))}
                _ = fi22
                _ = v23
            }
            for fi24 := 0; fi24 < 20; fi24++ {
                if (v18 > (0.0 * 1e10)) {
                    v25 := []int{(-(fi21))}
                    switch fi21 {
                    case 32768:
                        if true {
                            fmt.Println(v5, v14)
                        }
                    default:
                        if true {
                            if false {
                                fmt.Println(v18)
                                if len(v25) >= 2 {
                                    v25 = v25[1:len(v25)]
                                }
                                switch ("\n\t" + "a") {
                                case "hello":
                                    fmt.Println(v18, v1, v18)
                                case "a":
                                    v1 = -1
                                case "foo bar":
                                    var v26 int32 = genFn0(true)
                                    _ = v26
                                default:
                                    if true {
                                        v27 := GenS1{F0: -1, F1: 127, F2: "hello"}
                                        v14 -= v14
                                        _ = v27
                                    } else {
                                        if ((v14 * v14) == 1000) {
                                            var v28 int64 = 937
                                            switch (fi21 * fi24) {
                                            case -147:
                                                v29 := GenS1{F0: (v28 % ((v28) | 1)), F1: 127, F2: "\n\t"}
                                                _ = v29
                                            case 1:
                                                if false {
                                                    var v30 uint16 = 1
                                                    _ = v30
                                                } else {
                                                    var v31 uint = 0
                                                    var v32 uint16 = 70
                                                    _ = v31
                                                    _ = v32
                                                }
                                            case -195:
                                                if false {
                                                    if (!false) {
                                                        if false {
                                                            switch "\n\t" {
                                                            case "a":
                                                                v33 := GenS1{F0: (-(v28)), F1: 204, F2: "a"}
                                                                _ = v33
                                                            }
                                                            v15 = (^(v15))
                                                        } else {
                                                            var v34 int32 = v1
                                                            v28 = v28
                                                            v18 = 0.1
                                                            _ = v34
                                                        }
                                                    }
                                                }
                                            }
                                            _ = v28
                                        } else {
                                            var v35 float64 = ((v18 / float64((v3 / v3))) * v18)
                                            if ((v3 != 0.1) || (0 == 2)) {
                                                var v36 int64 = -9223372036854775808
                                                v18 += v35
                                                _ = v36
                                            }
                                            v1 += 42
                                            _ = v35
                                        }
                                    }
                                    var v37 int16 = -32768
                                    var v38 int32 = (v4 >> ((4294967295) & 63))
                                    _ = v37
                                    _ = v38
                                }
                            } else {
                                switch ((("foo bar" + ("\n\t" + "\n\t")) + (("foo bar" + (("" + "foo bar") + ("a" + "a"))) + "a")) + "\n\t") {
                                case "\n\t":
                                    if false {
                                        v1 = int32(v14)
                                        if true {
                                            if (!((v5 < (v5 / (((-(v5))) | 1))) && true)) {
                                                fmt.Println(v14, v2)
                                                var v39 bool = ((v14 ^ (v14 & v14)) != (v14 >> ((125) & 63)))
                                                _ = v39
                                            }
                                            v1 = v1
                                            if true {
                                                fmt.Println(fi21, v1)
                                            } else {
                                                var v40 int8 = v5
                                                var v41 uint16 = 2
                                                _ = v40
                                                _ = v41
                                            }
                                            fmt.Println(v1)
                                            var v42 uint64 = v15
                                            if true {
                                                var v43 float32 = v3
                                                _ = v43
                                            } else {
                                                var v44 uint = 0
                                                _ = v44
                                            }
                                            if len(v25) >= 2 {
                                                v25 = v25[1:len(v25)]
                                            }
                                            _ = v42
                                        } else {
                                            var v45 uint = 42
                                            _ = v45
                                        }
                                        v25 = append(v25, int(v14))
                                        var v46 int8 = -67
                                        var v47 float64 = v18
                                        if (!true) {
                                            v14 = 1
                                        }
                                        if false {
                                            var v48 int = (fi21 / ((1) | 1))
                                            if false {
                                                v3 = (v3 / float32(v47))
                                            } else {
                                                v48 -= (fi24 >> ((2) & 63))
                                                var v49 int16 = v14
                                                var v50 string = ("\n\t" + ("" + "\n\t"))
                                                _ = v49
                                                _ = v50
                                            }
                                            _ = v48
                                        } else {
                                            if (v2 <= 1) {
                                                v15 -= v15
                                                fmt.Println(v47, v47)
                                                fmt.Println(v5, fi24, v18)
                                                switch "hello" {
                                                case "a":
                                                    v3 += (v3 + (0.1 + v3))
                                                default:
                                                    var v51 int32 = v1
                                                    _ = v51
                                                }
                                            } else {
                                                var v52 uint = 18446744073709551615
                                                var v53 int = fi21
                                                var v54 uint32 = 42
                                                _ = v52
                                                _ = v53
                                                _ = v54
                                            }
                                        }
                                        _ = v46
                                        _ = v47
                                    } else {
                                        var v55 uint8 = 23
                                        if len(v25) >= 2 {
                                            v25 = v25[1:len(v25)]
                                        }
                                        _ = v55
                                    }
                                default:
                                    v1 -= (v4 + genFn0((65535 > 65535)))
                                }
                                v2 = v2
                                if true {
                                    v14 = (v14 % ((v14) | 1))
                                    v3 = v3
                                    if (!true) {
                                        switch "foo bar" {
                                        case "\n\t":
                                            _ = len(v25)
                                        case "":
                                            if (0 >= 182) {
                                                switch fi24 {
                                                case -1:
                                                    var v56 uint32 = 304
                                                    _ = v56
                                                case 9223372036854775807:
                                                    var v57 int8 = (v5 >> ((145) & 63))
                                                    _ = v57
                                                case 128:
                                                    var v58 float64 = (((((float64((0.1 - v3)) * float64(3.14159)) * (v18 + 1.5)) * (0.1 / v18)) + (float64((0.0 + ((v3 * 1.5) + (v3 * (v3 - 1.5))))) - ((v18 - 1.5) / (float64((float32(v18) / v3)) / (1.5 * 0.1))))) + float64(v3))
                                                    _ = v58
                                                }
                                            }
                                        case "foo bar":
                                            v2 = (v2 | v2)
                                        }
                                        switch "" {
                                        case "a":
                                            v59 := GenS1{F0: -9223372036854775808, F1: 131, F2: ("hello" + "")}
                                            var v60 float32 = (v3 + (v3 - 3.14159))
                                            _ = v59
                                            _ = v60
                                        }
                                    }
                                    var v61 string = (("hello" + ("hello" + "\n\t")) + (((("hello" + "\n\t") + "a") + "foo bar") + (("\n\t" + "") + "")))
                                    if len(v25) >= 2 {
                                        v25 = v25[1:len(v25)]
                                    }
                                    if len(v25) > 0 {
                                        v25[0] = fi21
                                    }
                                    if true {
                                        switch -9223372036854775808 {
                                        case 42:
                                            fmt.Println(v1)
                                        }
                                    } else {
                                        fmt.Println(v15, v1, v14)
                                    }
                                    if len(v25) >= 2 {
                                        v25 = v25[1:len(v25)]
                                    }
                                    if len(v25) > 0 {
                                        v25[0] = (fi24 | (^(fi24)))
                                    }
                                    _ = v61
                                }
                            }
                            v62 := GenS1{F0: 0, F1: 93, F2: "foo bar"}
                            v15 -= 42
                            _ = v62
                        }
                        v63 := GenS2{F0: 4294967295, F1: v2}
                        _ = v63
                    }
                    _ = v25
                }
                v14 -= v14
                v64 := []int{}
                _ = fi24
                _ = v64
            }
            _ = fi21
        }
        _ = v18
    }
    var v65 int32 = v1
    v66 := GenS2{F0: 0, F1: (v2 - v2)}
    for fi67 := 0; fi67 < 17; fi67++ {
        v4 -= v1
        _ = fi67
    }
    switch 9223372036854775807 {
    case 9223372036854775807:
        var v68 uint8 = 255
        _ = v68
    case -9223372036854775808:
        fmt.Println(v4, v2, v2)
    default:
        fmt.Println(v2, v2)
    }
    v69 := GenS3{F0: (v15 != (v15 | (v15 & (v15 + v15))))}
    for fi70 := 0; fi70 < 13; fi70++ {
        v65 = (-(v4))
        _ = fi70
    }
    v3 -= (((v3 - v3) - v3) - v3)
    fmt.Println(v15, v1)
    for fi71 := 0; fi71 < 19; fi71++ {
        switch -9223372036854775808 {
        case 42:
            switch (("hello" + "hello") + "hello") {
            case "":
                if false {
                    var v72 float64 = -2.25
                    v73 := GenS1{F0: 1, F1: 255, F2: "hello"}
                    v74 := GenS3{F0: true}
                    _ = v72
                    _ = v73
                    _ = v74
                }
            }
        }
        var v75 uint64 = (v15 / (((v15 / ((v15) | 1))) | 1))
        v1 += v1
        _ = fi71
        _ = v75
    }
    var v76 float32 = v3
    for fi77 := 0; fi77 < 20; fi77++ {
        var v78 float32 = float32(1.5)
        var v79 int8 = -118
        _ = fi77
        _ = v78
        _ = v79
    }
    if true {
        var v80 int8 = v5
        var v81 uint32 = v2
        var v82 float64 = 0.0
        _ = v80
        _ = v81
        _ = v82
    } else {
        v4 -= v4
    }
    v83 := GenS2{F0: 65535, F1: (v2 - v2)}
    fmt.Println(v4, v15)
    switch "foo bar" {
    case "":
        v84 := GenS2{F0: 1, F1: (v2 + v2)}
        _ = v84
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v65:", v65)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v66
    _ = v69
    _ = v76
    _ = v83
}
