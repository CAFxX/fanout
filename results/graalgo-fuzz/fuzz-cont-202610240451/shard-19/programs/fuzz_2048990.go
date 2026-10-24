package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 int32; F2 uint }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 int32; F1 int64 }
type GenS6 struct { F0 int16; F1 bool; F2 int }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 string }
func genFn0() uint8 {
    var v6 int16 = -1
    fmt.Println(v6)
    if ((true || false) || (-945 <= 42)) {
        v6 += v6
    } else {
        fmt.Println(v6)
        var v7 uint32 = 0
        v8 := []int{127, 42, 1, 32767}
        _ = v7
        _ = v8
    }
    fmt.Println(v6)
    _ = v6
    return 42
}

func main() {
    var v1 int8 = 127
    var v2 string = ("foo bar" + ("hello" + ("foo bar" + "")))
    var v3 int16 = 42
    var v4 float32 = ((1.5 + 3.14159) - (1e10 - float32(1.5)))
    v5 := []uint32{161}
    v3 -= (^(v3))
    v4 -= v4
    for fi9 := 0; fi9 < 10; fi9++ {
        v1 -= (v1 | v1)
        if (genFn0() < 11) {
            for fi10 := 0; fi10 < 8; fi10++ {
                fmt.Println(fi10, v1)
                if false {
                    if ((20 >= 65535) && (42 > 42)) {
                        var v11 uint8 = 163
                        if (false && true) {
                            v2 = v2
                            if true {
                                switch ((v2 + v2) + "") {
                                case "hello":
                                    v5 = append(v5, 2147483648)
                                case "foo bar":
                                    var v12 int32 = -2147483648
                                    _ = v12
                                case "\n\t":
                                    v13 := GenS1{F0: 414}
                                    _ = v13
                                default:
                                    switch (fi10 - (fi10 | 792)) {
                                    case -152:
                                        v14 := GenS1{F0: 837}
                                        _ = v14
                                    case 29:
                                        fmt.Println(v1)
                                    case 128:
                                        if (42 <= v11) {
                                            if (0 == 1) {
                                                v11 -= v11
                                            }
                                        } else {
                                            var v15 int = (fi9 + (fi10 / ((fi9) | 1)))
                                            v5 = append(v5, 4294967295)
                                            v16 := GenS1{F0: 2}
                                            _ = v15
                                            _ = v16
                                        }
                                    }
                                }
                                v17 := GenS1{F0: 4294967295}
                                if len(v5) > 0 {
                                    v5[0] = 2
                                }
                                _ = v17
                            } else {
                                v3 = (v3 << ((418) & 63))
                            }
                            v11 += (v11 << ((92) & 63))
                        }
                        _ = v11
                    }
                    v2 = v2
                    if ((723 != 65535) || (357 < -2147483648)) {
                        var v18 uint8 = genFn0()
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v18
                    }
                }
                _ = fi10
            }
            if false {
                var v19 int = (^(fi9))
                for fi20 := 0; fi20 < 4; fi20++ {
                    if true {
                        v2 = ("" + v2)
                        v21 := GenS2{F0: 1}
                        _ = v21
                    } else {
                        v22 := GenS3{F0: 1, F1: 42, F2: 2}
                        if true {
                            v5 = append(v5, 0)
                            v3 -= v3
                            if (2 < 32767) {
                                if true {
                                    if true {
                                        if (554 >= -2147483648) {
                                            v3 = 0
                                            var v23 uint64 = 4294967295
                                            _ = v23
                                        }
                                    } else {
                                        _ = len(v5)
                                        v24 := GenS2{F0: (v3 + (v3 >> ((105) & 63)))}
                                        _ = v24
                                    }
                                    var v25 float32 = v4
                                    _ = v25
                                } else {
                                    var v26 string = v2
                                    _ = v26
                                }
                                v3 = v3
                                if ((1e10 <= 1.5) || false) {
                                    fmt.Println(v3, v4, fi20)
                                    v5 = append(v5, 0)
                                    v5 = append(v5, 1)
                                } else {
                                    var v27 float32 = 1.5
                                    if true {
                                        fmt.Println(v3, v27, v3)
                                        fmt.Println(v1)
                                        v28 := GenS1{F0: 42}
                                        _ = v28
                                    } else {
                                        if (!(!true)) {
                                            if false {
                                                fmt.Println(v27)
                                            }
                                            v3 -= 1
                                            switch (fi20 * v19) {
                                            case 913:
                                                v3 -= v3
                                            case -1:
                                                v4 -= v27
                                            case -9223372036854775808:
                                                var v29 int64 = -9223372036854775808
                                                _ = v29
                                            default:
                                                v30 := GenS2{F0: (^(v3))}
                                                _ = v30
                                            }
                                            v31 := GenS4{F0: 803, F1: 9223372036854775807}
                                            _ = v31
                                        } else {
                                            v32 := GenS3{F0: 1, F1: 1, F2: 1}
                                            switch fi9 {
                                            case 0:
                                                fmt.Println(v27)
                                            }
                                            var v33 uint16 = 0
                                            _ = v32
                                            _ = v33
                                        }
                                        if (944 != 27) {
                                            var v34 float64 = ((1e10 / 1.0) * 1e10)
                                            v34 = v34
                                            _ = len(v5)
                                            _ = v34
                                        } else {
                                            if len(v5) > 0 {
                                                v5[0] = 1
                                            }
                                            switch v2 {
                                            case "foo bar":
                                                if (true && true) {
                                                    v1 = -107
                                                }
                                            default:
                                                v19 = fi9
                                            }
                                            var v35 int32 = 1
                                            if len(v5) > 0 {
                                                v5[0] = 2147483647
                                            }
                                            fmt.Println(v4, v3, v1)
                                            _ = v35
                                        }
                                    }
                                    v36 := GenS1{F0: 2}
                                    if ((3.14159 / 3.14159) <= float64((v4 / v27))) {
                                        if true {
                                            fmt.Println(fi20, v1, fi9)
                                            fmt.Println(v1, v19)
                                            if len(v5) > 0 {
                                                v5[0] = 2147483647
                                            }
                                        } else {
                                            var v37 int = (fi20 + fi20)
                                            _ = v37
                                        }
                                        if len(v5) > 0 {
                                            v5[0] = 0
                                        }
                                    } else {
                                        v38 := GenS1{F0: 0}
                                        _ = v38
                                    }
                                    if (65535 > 0) {
                                        v39 := GenS3{F0: 0, F1: -833, F2: 18446744073709551615}
                                        switch (v2 + v2) {
                                        case "foo bar":
                                            v40 := GenS5{F0: -1, F1: 1}
                                            _ = v40
                                        default:
                                            var v41 string = "hello"
                                            _ = v41
                                        }
                                        _ = v39
                                    }
                                    _ = v27
                                    _ = v36
                                }
                                var v42 uint32 = 649
                                if false {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                                _ = v42
                            } else {
                                if false {
                                    switch -973 {
                                    case 4294967295:
                                        v43 := GenS2{F0: (v3 % ((v3) | 1))}
                                        v5 = append(v5, 65535)
                                        _ = v43
                                    case 127:
                                        var v44 float64 = (-2.25 / -2.25)
                                        _ = v44
                                    default:
                                        var v45 uint64 = 4294967295
                                        _ = v45
                                    }
                                    var v46 bool = (!true)
                                    var v47 int64 = 0
                                    _ = v46
                                    _ = v47
                                } else {
                                    _ = len(v5)
                                    var v48 float32 = v4
                                    v4 += v48
                                    switch v19 {
                                    case -840:
                                        if len(v5) > 0 {
                                            v5[0] = 243
                                        }
                                    default:
                                        v48 -= v48
                                    }
                                    _ = v48
                                }
                                fmt.Println(v4, v1, v1)
                            }
                        }
                        _ = v22
                    }
                    v49 := GenS3{F0: 42, F1: -1, F2: 2}
                    _ = fi20
                    _ = v49
                }
                _ = v19
            }
        } else {
            var v50 float32 = (((v4 * 3.14159) + v4) * 1.5)
            var v51 bool = (v3 < v3)
            _ = v50
            _ = v51
        }
        _ = fi9
    }
    fmt.Println(v1, v3, v3)
    fmt.Println(v1, v2)
    if (((v4 + v4) * v4) != (((0.0 / 3.14159) * 1e10) / v4)) {
        v3 -= (v3 - v3)
        v3 = (^(v3))
        v3 += v3
    }
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    if false {
        if len(v5) > 0 {
            v5[0] = 596
        }
        for fi52 := 0; fi52 < 20; fi52++ {
            for fi53 := 0; fi53 < 17; fi53++ {
                var v54 uint8 = 202
                switch fi53 {
                case -2147483648:
                    if false {
                        if len(v5) > 0 {
                            v5[0] = 65535
                        }
                        var v55 int16 = 1
                        fmt.Println(v1, v4, v1)
                        _ = v55
                    } else {
                        var v56 int16 = 686
                        var v57 string = v2
                        if true {
                            v58 := GenS6{F0: 42, F1: false, F2: fi52}
                            v59 := GenS7{F0: 2}
                            _ = v58
                            _ = v59
                        } else {
                            v60 := GenS4{F0: 0, F1: 0}
                            _ = v60
                        }
                        _ = v56
                        _ = v57
                    }
                case 0:
                    if true {
                        var v61 uint16 = 2
                        _ = v61
                    }
                default:
                    v4 += v4
                }
                var v62 uint16 = 266
                _ = fi53
                _ = v54
                _ = v62
            }
            _ = fi52
        }
        if true {
            fmt.Println(v2, v3, v3)
        } else {
            switch "foo bar" {
            case "foo bar":
                v3 -= 1000
            case "\n\t":
                var v63 uint = 373
                _ = v63
            case "a":
                if (0 <= -9223372036854775808) {
                    var v64 bool = (!false)
                    _ = v64
                }
            default:
                switch ("\n\t" + "") {
                case "":
                    var v65 int = 4294967295
                    _ = v65
                default:
                    var v66 uint32 = 469
                    _ = v66
                }
            }
        }
    } else {
        var v67 uint32 = 2
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if len(v5) > 0 {
            v5[0] = (^(v67))
        }
        _ = v67
    }
    var v68 uint64 = 42
    v69 := GenS8{F0: "a"}
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v68
    _ = v69
}
