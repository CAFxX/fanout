package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS6 struct { F0 uint8; F1 bool }
type GenS7 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS8 struct { F0 int; F1 int16 }
type GenS9 struct { F0 int8; F1 int64 }
func genFn0(p0 uint8, p1 string, p2 string) int16 {
    var v4 int32 = 42
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := GenS1{F0: -32768, F1: (121 != 100)}
        _ = fi5
        _ = v6
    }
    v7 := []int{-213, 0, 4294967295}
    fmt.Println(p0, p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return 0
}
func genFn1() bool {
    if false {
        var v8 uint16 = 32767
        v9 := GenS1{F0: 1, F1: false}
        if true {
            v10 := []int{-2147483648, 32767, 32767}
            _ = v10
        } else {
            v11 := []int{588}
            _ = v11
        }
        _ = v8
        _ = v9
    } else {
        var v12 int16 = 0
        for fi13 := 0; fi13 < 4; fi13++ {
            fmt.Println(v12)
            var v14 int32 = -505
            if ((!true) || true) {
                var v15 int = (fi13 / ((fi13) | 1))
                v16 := []int{fi13}
                _ = v15
                _ = v16
            } else {
                v17 := GenS1{F0: -1, F1: false}
                v18 := GenS2{F0: v14}
                _ = v17
                _ = v18
            }
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    v19 := []int{405}
    switch (("a" + "\n\t") + "\n\t") {
    case "hello":
        var v20 string = ("\n\t" + "hello")
        _ = v20
    }
    var v21 float64 = 3.14159
    _ = v19
    _ = v21
    return true
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 1
    v3 := []string{"foo bar"}
    var v22 uint32 = 42
    v23 := v1[(v22 % ((v22) | 1))]
    if (-1 > 1) {
        if (42 <= 1) {
            switch "\n\t" {
            case "foo bar":
                v24 := GenS3{F0: 65535}
                _ = v24
            case "":
                v25 := GenS2{F0: -1}
                _ = v25
            case "a":
                for fi26 := 0; fi26 < 11; fi26++ {
                    for fi27 := 0; fi27 < 18; fi27++ {
                        v22 = v2
                        if (!true) {
                            v22 += 0
                        } else {
                            v2 = 2147483647
                            switch "foo bar" {
                            case "":
                                if true {
                                    v28 := GenS2{F0: 42}
                                    v23 = (fi27 & v23)
                                    _ = v28
                                }
                            case "foo bar":
                                if true {
                                    delete(v1, (^(v2)))
                                } else {
                                    if true {
                                        v2 = 538
                                        if ((true || true) || (!(!(-758 < 0)))) {
                                            v29 := GenS4{F0: 75, F1: (!(!true))}
                                            _ = v29
                                        } else {
                                            var v30 uint8 = 42
                                            if true {
                                                var v31 float64 = (-2.25 / 1.5)
                                                fmt.Println(v31, v31)
                                                v32 := GenS2{F0: -437}
                                                _ = v31
                                                _ = v32
                                            }
                                            if (((v2 / ((v22) | 1)) != v2) || (42 < -2147483648)) {
                                                v33 := GenS5{F0: 255, F1: uint16(v22), F2: 625}
                                                if ((127 > 127) || true) {
                                                    if ((639 != 905) && true) {
                                                        var v34 int16 = 42
                                                        delete(v1, (^(v2)))
                                                        v35 := GenS2{F0: 42}
                                                        _ = v34
                                                        _ = v35
                                                    } else {
                                                        var v36 uint = 1
                                                        _ = v36
                                                    }
                                                    if (genFn0(v30, ("hello" + "a"), "a") <= genFn0((v30 % (((v30 | v30)) | 1)), "", "a")) {
                                                        var v37 bool = false
                                                        v38 := GenS6{F0: 1, F1: v37}
                                                        var v39 uint16 = 32767
                                                        v40 := GenS7{F0: (0 <= 127), F1: (v39 ^ 955), F2: 0}
                                                        _ = v37
                                                        _ = v38
                                                        _ = v39
                                                        _ = v40
                                                    } else {
                                                        var v41 string = (("hello" + "") + "a")
                                                        v22 = v22
                                                        var v42 int16 = 1000
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                } else {
                                                    var v43 uint8 = 82
                                                    _ = v43
                                                }
                                                var v44 int8 = 127
                                                var v45 int32 = int32(v30)
                                                _ = v33
                                                _ = v44
                                                _ = v45
                                            }
                                            _ = v30
                                        }
                                    }
                                }
                            case "hello":
                                switch 2147483647 {
                                case 42:
                                    if ((0.1 * 1.5) < (-2.25 * (-2.25 * 0.1))) {
                                        var v46 int16 = 732
                                        v2 = 266
                                        switch (^(fi27)) {
                                        case 127:
                                            _ = len(v1)
                                        case -2147483648:
                                            var v47 int = (fi26 ^ 772)
                                            _ = v47
                                        case -1:
                                            v22 = 405
                                        default:
                                            var v48 int32 = -2147483648
                                            _ = v48
                                        }
                                        _ = v46
                                    } else {
                                        fmt.Println(fi26, v22)
                                    }
                                case 879:
                                    v23 += fi26
                                case -1:
                                    var v49 bool = false
                                    _ = v49
                                default:
                                    v2 += (v2 % ((v2) | 1))
                                    v2 = (v2 / ((135) | 1))
                                    if false {
                                        if (4294967295 <= (v22 & v2)) {
                                            var v50 float64 = float64(float32((1.5 - 0.1)))
                                            v50 += v50
                                            _ = v50
                                        }
                                        var v51 uint8 = 2
                                        v3 = append(v3, "hello")
                                        var v52 uint = 792
                                        var v53 int = fi26
                                        _ = v51
                                        _ = v52
                                        _ = v53
                                    }
                                }
                            }
                            _ = len(v1)
                        }
                        _ = fi27
                    }
                    v54 := GenS2{F0: 586}
                    switch (v23 << ((0) & 63)) {
                    case 801:
                        var v55 int = fi26
                        _ = v55
                    default:
                        var v56 uint16 = 808
                        _ = v56
                    }
                    _ = fi26
                    _ = v54
                }
            default:
                v1[v2] = v23
            }
            var v57 int64 = 0
            v58 := GenS8{F0: -2147483648, F1: int16(v57)}
            _ = v57
            _ = v58
        } else {
            var v59 int8 = -128
            _ = v59
        }
    }
    if true {
        for fi60 := 0; fi60 < 12; fi60++ {
            var v61 int64 = -854
            v61 += 348
            _ = fi60
            _ = v61
        }
    }
    for fi62 := 0; fi62 < 6; fi62++ {
        var v63 int8 = 42
        for fi64 := 0; fi64 < 16; fi64++ {
            v65 := GenS7{F0: false, F1: 413, F2: 2}
            switch (v23 + -9223372036854775808) {
            case 2147483647:
                switch "foo bar" {
                case "foo bar":
                    v3 = append(v3, "a")
                case "":
                    v63 -= -1
                case "hello":
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                }
            case 129:
                v2 -= v2
            }
            var v66 int32 = -451
            _ = fi64
            _ = v65
            _ = v66
        }
        v67 := v1[v2]
        _ = fi62
        _ = v63
        _ = v67
    }
    if false {
        v22 = 0
        v68 := GenS4{F0: 0, F1: false}
        _ = v68
    } else {
        var v69 uint16 = 255
        if false {
            for fi70 := 0; fi70 < 14; fi70++ {
                _ = len(v3)
                if (false || (true && false)) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    var v71 uint16 = 42
                    _ = v71
                }
                var v72 uint16 = v69
                _ = fi70
                _ = v72
            }
        } else {
            if true {
                if (128 >= 0) {
                    for fi73 := 0; fi73 < 13; fi73++ {
                        fmt.Println(v23, fi73)
                        if len(v3) > 0 {
                            v3[0] = ("\n\t" + ("\n\t" + ""))
                        }
                        if false {
                            var v74 uint8 = 255
                            var v75 uint32 = 568
                            _ = v74
                            _ = v75
                        }
                        _ = fi73
                    }
                    _ = len(v1)
                } else {
                    fmt.Println(v2, v23, v23)
                    for fi76 := 0; fi76 < 17; fi76++ {
                        v77 := v1[508]
                        fmt.Println(v23, v77, v77)
                        _ = fi76
                        _ = v77
                    }
                }
                switch 42 {
                case 9223372036854775807:
                    v78 := GenS9{F0: 127, F1: 42}
                    _ = v78
                case -2147483648:
                    var v79 int8 = -86
                    _ = v79
                case 32768:
                    for fi80 := 0; fi80 < 13; fi80++ {
                        if true {
                            fmt.Println(v22, v2, v2)
                        } else {
                            for fi81 := 0; fi81 < 1; fi81++ {
                                v23 = (-(fi80))
                                _ = fi81
                            }
                            v2 = (v22 | 42)
                        }
                        var v82 int = (fi80 * 0)
                        delete(v1, (^(v2)))
                        _ = fi80
                        _ = v82
                    }
                }
                if (42 < 65535) {
                    var v83 int8 = 0
                    fmt.Println(v2, v22, v83)
                    _ = v83
                }
            } else {
                var v84 int = (v23 % (((v23 * (v23 / ((v23) | 1)))) | 1))
                var v85 uint = 0
                _ = v84
                _ = v85
            }
        }
        _ = v69
    }
    fmt.Println(v22, v22, v22)
    v3 = append(v3, ("\n\t" + "foo bar"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
