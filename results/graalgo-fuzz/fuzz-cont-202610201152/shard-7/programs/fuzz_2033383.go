package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 uint64; F2 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 int64 }
func genFn0() bool {
    v4 := []int{0, 2147483647}
    var v5 float64 = float64(0.1)
    if (1 != 879) {
        v5 -= float64((1e10 * -2.25))
        if true {
            for fi6 := 0; fi6 < 17; fi6++ {
                v7 := GenS1{F0: -1, F1: 2, F2: ("foo bar" + "\n\t")}
                if false {
                    if (!false) {
                        v8 := GenS1{F0: -1, F1: 377, F2: "a"}
                        _ = v8
                    }
                } else {
                    var v9 int8 = 42
                    switch (("foo bar" + "") + ("\n\t" + "foo bar")) {
                    case "foo bar":
                        v10 := GenS1{F0: -1, F1: 1, F2: "hello"}
                        _ = v10
                    case "":
                        var v11 int8 = -77
                        _ = v11
                    case "hello":
                        var v12 int32 = -2147483648
                        _ = v12
                    default:
                        if false {
                            for fi13 := 0; fi13 < 10; fi13++ {
                                var v14 int64 = 9223372036854775807
                                switch fi13 {
                                case 32767:
                                    if false {
                                        var v15 int8 = 100
                                        _ = v15
                                    }
                                case 42:
                                    v4 = append(v4, 128)
                                default:
                                    fmt.Println(v5, fi13, v9)
                                }
                                _ = fi13
                                _ = v14
                            }
                            v5 = float64(((0.0 + 0.1) + -2.25))
                        }
                    }
                    _ = len(v4)
                    _ = v9
                }
                for fi16 := 0; fi16 < 17; fi16++ {
                    v4 = append(v4, 9223372036854775807)
                    fmt.Println(v5)
                    _ = fi16
                }
                _ = fi6
                _ = v7
            }
            v4 = append(v4, -1)
        } else {
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v5)
            case "":
                if len(v4) > 0 {
                    v4[0] = -667
                }
            }
            for fi17 := 0; fi17 < 13; fi17++ {
                v5 = (v5 * float64(1.5))
                _ = fi17
            }
        }
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        fmt.Println(fi18)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi18
    }
    _ = v4
    _ = v5
    return (true && true)
}

func main() {
    var v1 uint = 42
    v2 := make(map[string]uint32)
    var v3 int = 42
    delete(v2, "foo bar")
    var v19 float32 = float32((-2.25 / (-2.25 / ((((0.1 / -2.25) + ((0.1 + 1.5) / 1e10)) / 0.1) * 3.14159))))
    for fi20 := 0; fi20 < 11; fi20++ {
        switch v3 {
        case 458:
            for fi21 := 0; fi21 < 16; fi21++ {
                v22 := GenS2{F0: 381}
                if (false && ((42 <= 470) || (!(!true)))) {
                    _ = len(v2)
                    switch v3 {
                    case -660:
                        v1 -= (v1 & v1)
                    default:
                        var v23 int32 = 1
                        _ = v23
                    }
                    v3 = (fi20 - (fi21 / ((127) | 1)))
                } else {
                    var v24 uint = (^(v1))
                    v25 := GenS1{F0: -857, F1: 831, F2: ""}
                    _ = v24
                    _ = v25
                }
                v26 := v2[""]
                v26 += (v26 % (((v26 << (((^(v1))) & 63))) | 1))
                var v27 int64 = 1
                _ = fi21
                _ = v22
                _ = v26
                _ = v27
            }
        case -563:
            if false {
                v19 -= (float32(0.0) - 0.1)
                fmt.Println(v1, fi20, v1)
            } else {
                v3 = 128
            }
        default:
            switch fi20 {
            case 380:
                var v28 uint64 = 18446744073709551615
                _ = v28
            case 2147483647:
                v19 -= (v19 - 0.1)
            case 42:
                for fi29 := 0; fi29 < 5; fi29++ {
                    var v30 float32 = -2.25
                    var v31 uint32 = 964
                    if ((-1 > -9223372036854775808) || true) {
                        if false {
                            v32 := []int{}
                            _ = v32
                        }
                        fmt.Println(v1, v3, v3)
                        var v33 int64 = int64(v1)
                        fmt.Println(fi20)
                        v33 -= (^(v33))
                        v3 -= 128
                        delete(v2, "")
                        v19 += 0.0
                        _ = v33
                    }
                    v1 = v1
                    switch "hello" {
                    case "foo bar":
                        if true {
                            v34 := GenS1{F0: 32767, F1: 734, F2: "a"}
                            _ = v34
                        }
                    default:
                        var v35 uint16 = 321
                        _ = v35
                    }
                    fmt.Println(fi29, fi20)
                    fmt.Println(fi29)
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            }
        }
        v36 := GenS2{F0: 4294967295}
        v2["foo bar"] = 2147483648
        _ = fi20
        _ = v36
    }
    v37 := GenS2{F0: 223}
    var v38 uint = 18446744073709551615
    fmt.Println(v38, v19, v19)
    for fi39 := 0; fi39 < 2; fi39++ {
        for fi40 := 0; fi40 < 7; fi40++ {
            var v41 uint8 = 152
            _ = fi40
            _ = v41
        }
        var v42 float64 = (float64((v19 * v19)) - -2.25)
        _ = fi39
        _ = v42
    }
    v43 := []int{9223372036854775807}
    v38 = 311
    if false {
        v44 := GenS1{F0: -1, F1: 32767, F2: "foo bar"}
        var v45 bool = (!false)
        _ = v44
        _ = v45
    }
    v38 = v1
    fmt.Println(v19, v19, v1)
    if (!false) {
        v3 = v3
        v3 -= -2147483648
    } else {
        for fi46 := 0; fi46 < 9; fi46++ {
            switch v3 {
            case -476:
                var v47 int64 = -9223372036854775808
                _ = v47
            }
            switch (("\n\t" + "hello") + "\n\t") {
            case "":
                if (!(v19 < v19)) {
                    v48 := GenS2{F0: 2147483647}
                    switch (v3 + v3) {
                    case 32768:
                        for fi49 := 0; fi49 < 8; fi49++ {
                            var v50 uint64 = 932
                            _ = fi49
                            _ = v50
                        }
                    case 32767:
                        var v51 bool = ((0.0 * v19) == 3.14159)
                        _ = v51
                    default:
                        var v52 int64 = 9223372036854775807
                        _ = v52
                    }
                    for fi53 := 0; fi53 < 5; fi53++ {
                        if ((!(255 >= 255)) || true) {
                            v2["\n\t"] = 2147483647
                            var v54 int64 = -971
                            fmt.Println(v19)
                            _ = v54
                        } else {
                            switch fi53 {
                            case -9223372036854775808:
                                if false {
                                    v55 := v2[""]
                                    switch 103 {
                                    case 9223372036854775807:
                                        if true {
                                            fmt.Println(fi53, fi53)
                                        }
                                    }
                                    _ = v55
                                } else {
                                    fmt.Println(v38)
                                }
                            default:
                                if len(v43) > 0 {
                                    v43[0] = (fi53 * (v3 ^ v3))
                                }
                            }
                            v56 := GenS2{F0: 4294967295}
                            if true {
                                if true {
                                    if len(v43) >= 2 {
                                        v43 = v43[1:len(v43)]
                                    }
                                    switch -293 {
                                    case 9223372036854775807:
                                        var v57 int16 = 32767
                                        _ = v57
                                    case -42:
                                        v3 = (fi46 >> (((v38 / ((v38) | 1))) & 63))
                                    default:
                                        v58 := GenS1{F0: -71, F1: 2, F2: ("" + "foo bar")}
                                        _ = v58
                                    }
                                    var v59 bool = false
                                    if v59 {
                                        var v60 uint16 = 2
                                        v61 := GenS1{F0: 0, F1: (v60 ^ v60), F2: "\n\t"}
                                        _ = v60
                                        _ = v61
                                    } else {
                                        v19 = 1e10
                                        if len(v43) >= 2 {
                                            v43 = v43[1:len(v43)]
                                        }
                                        if v59 {
                                            var v62 uint32 = 1
                                            var v63 uint32 = (v62 % (((v62 / ((574) | 1))) | 1))
                                            v62 = (v62 << ((42) & 63))
                                            _ = v62
                                            _ = v63
                                        } else {
                                            v64 := GenS2{F0: 0}
                                            var v65 uint64 = 4294967295
                                            _ = v64
                                            _ = v65
                                        }
                                    }
                                    v66 := GenS3{F0: -9223372036854775808, F1: 42, F2: ""}
                                    _ = v59
                                    _ = v66
                                } else {
                                    var v67 int = fi46
                                    var v68 string = "hello"
                                    _ = v67
                                    _ = v68
                                }
                                if false {
                                    switch "" {
                                    case "\n\t":
                                        if (fi46 == fi53) {
                                            if len(v43) >= 2 {
                                                v43 = v43[1:len(v43)]
                                            }
                                            if false {
                                                switch "a" {
                                                case "a":
                                                    v19 -= 1e10
                                                case "hello":
                                                    var v69 uint64 = 1
                                                    _ = v69
                                                }
                                                fmt.Println(v1, v1)
                                            } else {
                                                v38 = v38
                                            }
                                        }
                                    case "foo bar":
                                        var v70 int8 = -24
                                        _ = v70
                                    default:
                                        fmt.Println(v19)
                                    }
                                    var v71 float32 = ((1.5 * float32((1.5 / float64(3.14159)))) - v19)
                                    var v72 int = 42
                                    _ = v71
                                    _ = v72
                                } else {
                                    v1 += 1
                                    v19 += 3.14159
                                }
                                v73 := GenS3{F0: -1, F1: 1, F2: "a"}
                                _ = v73
                            }
                            _ = v56
                        }
                        if (4294967295 > 1) {
                            var v74 int8 = 71
                            v75 := GenS4{F0: 2}
                            v76 := GenS1{F0: -32768, F1: 755, F2: ("foo bar" + "")}
                            _ = v74
                            _ = v75
                            _ = v76
                        }
                        var v77 int32 = 840
                        _ = fi53
                        _ = v77
                    }
                    _ = v48
                } else {
                    fmt.Println(v1, fi46, v1)
                    var v78 uint8 = 128
                    v3 -= (v3 % (((^(fi46))) | 1))
                    _ = v78
                }
            case "a":
                fmt.Println(v1)
            default:
                switch 32768 {
                case 762:
                    v79 := GenS5{F0: 131, F1: 9223372036854775807}
                    _ = v79
                case 32767:
                    v3 -= (-(v3))
                case -2147483648:
                    fmt.Println(v3, v19)
                default:
                    v3 = (^(fi46))
                }
            }
            var v80 uint16 = 255
            _ = fi46
            _ = v80
        }
        if false {
            var v81 string = "hello"
            _ = len(v2)
            _ = v81
        } else {
            for fi82 := 0; fi82 < 1; fi82++ {
                var v83 int16 = 911
                var v84 uint16 = 32767
                _ = fi82
                _ = v83
                _ = v84
            }
            if (270 <= 121) {
                var v85 int = (-(v3))
                switch ("hello" + ("hello" + "foo bar")) {
                case "\n\t":
                    v3 += v85
                default:
                    v1 -= v38
                }
                _ = v85
            } else {
                for fi86 := 0; fi86 < 20; fi86++ {
                    v3 = (fi86 ^ (v3 << ((v38) & 63)))
                    _ = fi86
                }
            }
        }
    }
    for fi87 := 0; fi87 < 5; fi87++ {
        var v88 uint32 = 2147483647
        for fi89 := 0; fi89 < 11; fi89++ {
            v88 = v88
            if (false || (v88 >= v88)) {
                if len(v43) > 0 {
                    v43[0] = (v3 + 112)
                }
            }
            _ = fi89
        }
        _ = fi87
        _ = v88
    }
    var v90 int64 = -1
    _ = len(v43)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v43
    _ = v90
}
