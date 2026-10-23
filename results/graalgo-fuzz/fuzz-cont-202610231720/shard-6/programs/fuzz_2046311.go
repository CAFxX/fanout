package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint16; F1 int16; F2 string }
type GenS8 struct { F0 uint32; F1 int8 }
type GenS9 struct { F0 uint32; F1 uint; F2 int16 }
type GenS10 struct { F0 uint16; F1 string }
type GenS11 struct { F0 uint16; F1 bool }
func genFn0(p0 uint8, p1 int, p2 uint) string {
    p2 = (^(p2))
    if (p1 == (p1 >> ((uint(p0)) & 63))) {
        fmt.Println(p2, p0, p2)
        v4 := []int{2147483647, 4294967295, (-(p1)), (^(p1))}
        var v5 bool = false
        _ = v4
        _ = v5
    }
    v6 := GenS1{F0: 42, F1: (!true)}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (("\n\t" + "\n\t") + "")
}

func main() {
    v1 := []uint32{977}
    v2 := make(map[string]string)
    v3 := make(map[uint32]uint32)
    for fi7 := 0; fi7 < 17; fi7++ {
        if true {
            _ = len(v2)
            _ = len(v2)
        }
        var v8 int8 = 60
        v8 = (v8 / ((v8) | 1))
        _ = fi7
        _ = v8
    }
    v9 := GenS2{F0: 255, F1: 0}
    var v10 int8 = 0
    var v11 uint16 = 42
    fmt.Println(v11)
    v10 -= v10
    var v12 int32 = 2147483647
    fmt.Println(v12, v12, v11)
    for fi13 := 0; fi13 < 10; fi13++ {
        if true {
            var v14 uint8 = 2
            fmt.Println(fi13, v12)
            var v15 uint16 = (v11 >> ((18446744073709551615) & 63))
            _ = v14
            _ = v15
        }
        _ = fi13
    }
    v11 = v11
    for fi16 := 0; fi16 < 2; fi16++ {
        if (true || true) {
            for fi17 := 0; fi17 < 3; fi17++ {
                switch fi16 {
                case 42:
                    v18 := GenS2{F0: (v11 >> ((uint(v11)) & 63)), F1: 0}
                    _ = v18
                case 163:
                    v10 = v10
                case 4294967295:
                    v11 = (v11 + (v11 - (v11 + 65535)))
                }
                _ = fi17
            }
            if true {
                v19 := GenS1{F0: v11, F1: (128 != 1)}
                var v20 string = "a"
                fmt.Println(fi16, v12, v12)
                _ = v19
                _ = v20
            } else {
                var v21 uint16 = (v11 >> ((2) & 63))
                _ = v21
            }
            switch "" {
            case "foo bar":
                var v22 uint8 = 42
                _ = v22
            default:
                for fi23 := 0; fi23 < 13; fi23++ {
                    _ = len(v3)
                    var v24 uint32 = 0
                    v25 := GenS1{F0: (v11 / (((v11 << ((995) & 63))) | 1)), F1: false}
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
        } else {
            v26 := GenS2{F0: v11, F1: 42}
            for fi27 := 0; fi27 < 20; fi27++ {
                v28 := GenS1{F0: (v11 % ((v11) | 1)), F1: true}
                _ = fi27
                _ = v28
            }
            _ = v26
        }
        for fi29 := 0; fi29 < 1; fi29++ {
            if true {
                if (-305 > 354) {
                    v2["\n\t"] = "foo bar"
                    v10 = (v10 << ((2) & 63))
                    var v30 float64 = -2.25
                    _ = v30
                } else {
                    if (false && ((v12 | (v12 % (((v12 * (v12 ^ (v12 >> ((615) & 63))))) | 1))) < v12)) {
                        if (!true) {
                            fmt.Println(fi16)
                            fmt.Println(v11, v12, v10)
                            fmt.Println(v11, v11)
                            v31 := GenS1{F0: v11, F1: (42 == 0)}
                            switch fi29 {
                            case -948:
                                var v32 string = "\n\t"
                                _ = v32
                            case 9223372036854775807:
                                v33 := GenS2{F0: 575, F1: 0}
                                _ = v33
                            }
                            v11 = v11
                            switch 128 {
                            case 32767:
                                var v34 int16 = 1
                                _ = v34
                            }
                            _ = v31
                        }
                        fmt.Println(v10, v11, fi16)
                        v35 := v3[2147483648]
                        v36 := GenS2{F0: v11, F1: (^(v35))}
                        v2["a"] = "hello"
                        _ = v35
                        _ = v36
                    } else {
                        var v37 int = fi29
                        v37 -= v37
                        if true {
                            _ = len(v1)
                            v3[65535] = 2
                            fmt.Println(v10)
                        }
                        var v38 uint32 = 2147483648
                        switch "" {
                        case "foo bar":
                            switch 9223372036854775807 {
                            case 834:
                                if (true && true) {
                                    if len(v1) > 0 {
                                        v1[0] = (v38 ^ 1)
                                    }
                                    var v39 uint32 = 42
                                    _ = v39
                                }
                            }
                            var v40 uint64 = 0
                            v37 = (^(fi16))
                            _ = v40
                        case "\n\t":
                            var v41 int = (v37 * (^(fi29)))
                            _ = v41
                        default:
                            v38 = (v38 * v38)
                            v12 += 1
                            var v42 uint16 = (v11 & (v11 / (((v11 * (v11 | 32))) | 1)))
                            v43 := GenS2{F0: 642, F1: v38}
                            _ = v42
                            _ = v43
                        }
                        v10 = (^(v10))
                        var v44 float32 = ((1.5 / 3.14159) - (float32(1e10) - 1.5))
                        _ = v37
                        _ = v38
                        _ = v44
                    }
                    if ((v12 * (^(v12))) >= (v12 & v12)) {
                        fmt.Println(v10, fi16, v10)
                        v45 := GenS3{F0: "\n\t", F1: (v10 | -128)}
                        v12 += (v12 | 0)
                        if false {
                            v46 := GenS3{F0: ("" + "hello"), F1: (v10 << ((255) & 63))}
                            v47 := GenS4{F0: v12}
                            if (false || (242 == (v11 ^ v11))) {
                                fmt.Println(v10)
                                if true {
                                    v10 += (^(v10))
                                    var v48 uint32 = 790
                                    fmt.Println(v48, v10, v12)
                                    _ = v48
                                }
                            }
                            if false {
                                if (2 <= 166) {
                                    v49 := GenS5{F0: -9223372036854775808, F1: (!false)}
                                    _ = v49
                                } else {
                                    var v50 int16 = 42
                                    switch "a" {
                                    case "foo bar":
                                        v51 := GenS4{F0: (v12 / ((v12) | 1))}
                                        _ = v51
                                    case "hello":
                                        v11 += 82
                                    default:
                                        var v52 int32 = 987
                                        _ = v52
                                    }
                                    _ = v50
                                }
                                var v53 uint16 = (v11 | 65535)
                                v2[("foo bar" + "")] = "foo bar"
                                _ = v53
                            } else {
                                if ((("foo bar" + "\n\t") + "\n\t") < "hello") {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v54 := GenS6{F0: 2}
                                    if (((true && (4294967295 <= 299)) && (true && false)) || (!true)) {
                                        v55 := GenS7{F0: (v11 >> ((255) & 63)), F1: -32768, F2: "\n\t"}
                                        _ = v55
                                    }
                                    v12 += (v12 | (v12 + 42))
                                    var v56 uint8 = 15
                                    _ = len(v1)
                                    _ = v54
                                    _ = v56
                                } else {
                                    var v57 int = int(v11)
                                    var v58 int8 = v10
                                    v59 := GenS8{F0: 42, F1: v10}
                                    fmt.Println(fi29, fi29, v57)
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                }
                                var v60 int8 = (v10 - v10)
                                v61 := GenS9{F0: 0, F1: 18446744073709551615, F2: -1}
                                _ = v60
                                _ = v61
                            }
                            v10 = (v10 ^ -113)
                            _ = v46
                            _ = v47
                        } else {
                            fmt.Println(fi29, v12, fi29)
                            if true {
                                v62 := GenS5{F0: 288, F1: true}
                                v12 += v12
                                v10 += (v10 - (-(v10)))
                                _ = v62
                            }
                        }
                        _ = v45
                    } else {
                        if false {
                            if (42 == 4294967295) {
                                v63 := GenS10{F0: 635, F1: "foo bar"}
                                var v64 uint8 = 42
                                _ = v63
                                _ = v64
                            } else {
                                var v65 uint64 = 2
                                v65 = 2
                                _ = v65
                            }
                        }
                        var v66 float64 = 1e10
                        if false {
                            var v67 uint8 = 255
                            var v68 float32 = 0.1
                            fmt.Println(v10, v68, fi29)
                            fmt.Println(fi16, v12, fi29)
                            _ = v67
                            _ = v68
                        }
                        _ = v66
                    }
                }
                if (v12 <= 1) {
                    var v69 string = "foo bar"
                    v70 := GenS7{F0: (v11 + (^(v11))), F1: 32767, F2: "a"}
                    v71 := GenS11{F0: v11, F1: true}
                    _ = v69
                    _ = v70
                    _ = v71
                }
                _ = len(v1)
                var v72 int32 = 2147483647
                _ = v72
            }
            var v73 uint64 = 1
            var v74 uint = 1
            _ = fi29
            _ = v73
            _ = v74
        }
        var v75 uint8 = 239
        _ = fi16
        _ = v75
    }
    for fi76 := 0; fi76 < 16; fi76++ {
        _ = len(v1)
        for fi77 := 0; fi77 < 20; fi77++ {
            var v78 uint = 167
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v79 uint8 = 134
            _ = fi77
            _ = v78
            _ = v79
        }
        _ = fi76
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v9
}
