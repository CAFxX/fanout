package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint32; F2 int }
type GenS3 struct { F0 int16; F1 int64; F2 uint }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 int32; F1 string; F2 int32 }
func genFn0() uint32 {
    switch ("hello" + "foo bar") {
    case "\n\t":
        v5 := GenS1{F0: true, F1: -119, F2: 42}
        _ = v5
    case "foo bar":
        if ((395 != 2147483647) || (false && (0 >= 1))) {
            var v6 int8 = -1
            var v7 uint8 = 167
            fmt.Println(v7, v6)
            _ = v6
            _ = v7
        }
    case "":
        v8 := []int{-211, 32768, 761}
        _ = v8
    default:
        v9 := GenS1{F0: false, F1: 0, F2: 42}
        _ = v9
    }
    var v10 int = -2147483648
    _ = v10
    return 0
}

func main() {
    var v1 float64 = (3.14159 * (1.5 + 0.0))
    var v2 string = "hello"
    v3 := make(map[string]uint32)
    var v4 uint = 352
    v4 -= 239
    for fi11 := 0; fi11 < 13; fi11++ {
        v2 = v2
        var v12 string = ""
        var v13 bool = (!false)
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 uint32 = 4294967295
    switch 42 {
    case -9223372036854775808:
        v4 += uint(v14)
    default:
        if false {
            v14 += (v14 * v14)
        }
    }
    v3["hello"] = v14
    v2 = (v2 + (((v2 + ("hello" + v2)) + ((v2 + ((v2 + (("\n\t" + (v2 + (v2 + ("hello" + v2)))) + "hello")) + "\n\t")) + v2)) + (v2 + "hello")))
    var v15 bool = (false || (42 >= 0))
    v16 := []int{102, 127, -641, -600}
    var v17 uint64 = 517
    for fi18 := 0; fi18 < 1; fi18++ {
        v15 = (!v15)
        for fi19 := 0; fi19 < 5; fi19++ {
            if v15 {
                var v20 int = fi19
                _ = v20
            }
            v21 := GenS1{F0: (65535 < 916), F1: -97, F2: v17}
            var v22 uint = v4
            _ = fi19
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 11; fi23++ {
            v24 := GenS1{F0: (false || true), F1: 100, F2: 809}
            fmt.Println(v17, v17)
            _ = fi23
            _ = v24
        }
        _ = fi18
    }
    for fi25 := 0; fi25 < 4; fi25++ {
        if (((-2.25 == 1.5) && (v15 || (false && v15))) || true) {
            v26 := GenS1{F0: (2147483648 != 723), F1: 40, F2: 18446744073709551615}
            _ = v26
        } else {
            var v27 bool = (!v15)
            _ = v27
        }
        if (((("a" + (v2 + ("hello" + v2))) < "a") || v15) || v15) {
            for fi28 := 0; fi28 < 12; fi28++ {
                if (!v15) {
                    v29 := GenS1{F0: ((v14 ^ 0) < (v14 - (^(v14)))), F1: -128, F2: (v17 % (((v17 ^ 0)) | 1))}
                    var v30 int32 = int32(fi25)
                    var v31 float32 = (((((1e10 + -2.25) * 0.0) * 0.0) - ((3.14159 - (((float32(v1) * (-2.25 * 3.14159)) * (1.5 - float32((v1 / (-2.25 * v1))))) + (float32(float64((3.14159 / 1e10))) / 1.0))) / ((float32(0.1) - 3.14159) + float32(v1)))) + 1e10)
                    fmt.Println(v30)
                    if (fi25 < fi25) {
                        var v32 int = int(v30)
                        v1 = float64(float32((v1 - 0.0)))
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                        v3[v2] = (v14 << ((1) & 63))
                        if false {
                            v31 -= v31
                            switch fi25 {
                            case 128:
                                switch (v2 + (v2 + "a")) {
                                case "foo bar":
                                    var v33 float64 = v1
                                    _ = v33
                                case "\n\t":
                                    switch 32768 {
                                    case -770:
                                        delete(v3, v2)
                                    case -819:
                                        var v34 int64 = 0
                                        _ = v34
                                    case -1:
                                        var v35 uint8 = 2
                                        _ = v35
                                    default:
                                        var v36 bool = ((!(0 > 0)) || v15)
                                        _ = v36
                                    }
                                case "hello":
                                    if (v15 || false) {
                                        v14 += 2
                                        var v37 uint16 = 0
                                        _ = v37
                                    }
                                default:
                                    var v38 int8 = 0
                                    _ = v38
                                }
                            case 1:
                                v32 = fi25
                            }
                            var v39 string = "\n\t"
                            fmt.Println(v32, v32, v17)
                            if v15 {
                                var v40 float64 = v1
                                if (9223372036854775807 > 42) {
                                    v41 := GenS2{F0: v15, F1: (v14 - genFn0()), F2: fi25}
                                    v42 := GenS2{F0: (!v15), F1: genFn0(), F2: 154}
                                    _ = v41
                                    _ = v42
                                }
                                if false {
                                    v43 := GenS1{F0: false, F1: -36, F2: v17}
                                    v44 := GenS1{F0: (v30 >= (v30 % (((v30 % (((^(v30))) | 1))) | 1))), F1: int8(v4), F2: 1}
                                    _ = v43
                                    _ = v44
                                } else {
                                    var v45 int8 = -1
                                    switch (((("" + v39) + (v39 + v39)) + "foo bar") + "a") {
                                    case "":
                                        switch 0 {
                                        case 0:
                                            v3["a"] = (v14 << ((672) & 63))
                                        case -584:
                                            v46 := v3[""]
                                            _ = v46
                                        default:
                                            if (-9223372036854775808 < -402) {
                                                fmt.Println(v39, v15)
                                                var v47 uint16 = 2
                                                _ = v47
                                            }
                                        }
                                    }
                                    v17 -= (v17 % (((v17 / (((v17 >> ((v4) & 63))) | 1))) | 1))
                                    _ = v45
                                }
                                _ = v40
                            } else {
                                var v48 float32 = v31
                                var v49 float32 = v48
                                v50 := GenS3{F0: 221, F1: 9223372036854775807, F2: (v4 ^ (v4 / (((v4 ^ 615)) | 1)))}
                                v51 := GenS1{F0: (-9223372036854775808 > 328), F1: 127, F2: (^(v17))}
                                _ = v48
                                _ = v49
                                _ = v50
                                _ = v51
                            }
                            _ = v39
                        }
                        _ = v32
                    } else {
                        v52 := GenS1{F0: (!v15), F1: 1, F2: (v17 | (v17 >> (((^(v4))) & 63)))}
                        _ = v52
                    }
                    _ = v29
                    _ = v30
                    _ = v31
                } else {
                    if ((v17 * 2) != v17) {
                        var v53 int = fi25
                        var v54 uint32 = v14
                        v55 := GenS4{F0: v2, F1: (-(fi28))}
                        _ = v53
                        _ = v54
                        _ = v55
                    }
                }
                v15 = v15
                _ = fi28
            }
        } else {
            if (!(v2 <= "")) {
                fmt.Println(v14, v4)
                v1 = 3.14159
                v17 += 477
            } else {
                var v56 int = fi25
                for fi57 := 0; fi57 < 12; fi57++ {
                    if ((-1 == 167) || true) {
                        switch (((v2 + (((((v2 + "") + "") + ("" + (("" + "foo bar") + v2))) + "hello") + v2)) + "hello") + v2) {
                        case "":
                            v14 = v14
                        case "hello":
                            fmt.Println(fi57, v56, fi25)
                        case "foo bar":
                            fmt.Println(v2)
                        }
                        switch ((v2 + "a") + v2) {
                        case "":
                            if v15 {
                                _ = len(v3)
                            } else {
                                _ = len(v3)
                            }
                        }
                        v56 += fi57
                    }
                    _ = fi57
                }
                _ = v56
            }
        }
        if (!false) {
            if ((^(v4)) > v4) {
                v16 = append(v16, fi25)
            } else {
                var v58 uint32 = 1
                v4 = v4
                v4 -= (v4 | v4)
                _ = v58
            }
            v59 := GenS2{F0: true, F1: (v14 >> (((v4 | (v4 | v4))) & 63)), F2: fi25}
            for fi60 := 0; fi60 < 2; fi60++ {
                fmt.Println(v14, v15)
                switch (((v2 + v2) + ("hello" + v2)) + v2) {
                case "foo bar":
                    if v15 {
                        v4 += uint(v17)
                    } else {
                        v2 = v2
                        var v61 uint = 2
                        _ = v61
                    }
                default:
                    switch fi60 {
                    case -583:
                        v62 := GenS5{F0: -1, F1: v2, F2: 0}
                        _ = v62
                    }
                }
                if ((!(!(!false))) && (940 < -2147483648)) {
                    v63 := GenS1{F0: v15, F1: -59, F2: (v17 + (v17 % ((v17) | 1)))}
                    var v64 int64 = 42
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    var v65 float32 = -2.25
                    _ = v63
                    _ = v64
                    _ = v65
                } else {
                    var v66 uint32 = v14
                    var v67 uint64 = (v17 + 132)
                    _ = v66
                    _ = v67
                }
                _ = fi60
            }
            _ = v59
        }
        _ = fi25
    }
    if false {
        fmt.Println(v14)
        fmt.Println(v2)
    }
    v15 = v15
    for fi68 := 0; fi68 < 18; fi68++ {
        if len(v16) > 0 {
            v16[0] = fi68
        }
        _ = fi68
    }
    if v15 {
        v69 := v3["\n\t"]
        var v70 float32 = (0.1 / 1.0)
        _ = v69
        _ = v70
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
