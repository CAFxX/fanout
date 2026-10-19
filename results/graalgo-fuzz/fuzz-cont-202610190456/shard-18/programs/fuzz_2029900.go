package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int; F1 int; F2 uint64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int32 }
func genFn0(p0 uint64, p1 int8) uint32 {
    v6 := []int{}
    if true {
        for fi7 := 0; fi7 < 18; fi7++ {
            fmt.Println(fi7, fi7, fi7)
            fmt.Println(p1, fi7, fi7)
            _ = fi7
        }
        var v8 uint = 65535
        _ = v8
    }
    var v9 int16 = -1
    p1 += p1
    v10 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    _ = v9
    return 2147483647
}
func genFn1(p0 int8, p1 int32) uint8 {
    if (!(42 < 2)) {
        var v11 int = 32767
        _ = v11
    } else {
        v12 := []int{}
        if (true || false) {
            if true {
                v13 := GenS2{F0: 0}
                _ = v13
            } else {
                fmt.Println(p0, p0)
                switch -660 {
                case 42:
                    v14 := GenS3{F0: 32767, F1: 47}
                    _ = v14
                case 987:
                    fmt.Println(p0, p0)
                case 2147483647:
                    switch ("\n\t" + (((("hello" + "\n\t") + "hello") + "foo bar") + "hello")) {
                    case "hello":
                        var v15 uint = 2
                        _ = v15
                    default:
                        if true {
                            fmt.Println(p1)
                            _ = len(v12)
                            if false {
                                for fi16 := 0; fi16 < 15; fi16++ {
                                    for fi17 := 0; fi17 < 3; fi17++ {
                                        p1 -= p1
                                        var v18 int32 = p1
                                        _ = fi17
                                        _ = v18
                                    }
                                    var v19 uint16 = 1
                                    if true {
                                        var v20 int32 = p1
                                        v20 = (p1 % ((p1) | 1))
                                        _ = v20
                                    }
                                    _ = fi16
                                    _ = v19
                                }
                            } else {
                                for fi21 := 0; fi21 < 2; fi21++ {
                                    if false {
                                        v22 := GenS4{F0: (^(p1))}
                                        v12 = append(v12, (fi21 - (fi21 << ((4294967295) & 63))))
                                        if false {
                                            for fi23 := 0; fi23 < 5; fi23++ {
                                                p1 += -2147483648
                                                if false {
                                                    v24 := GenS5{F0: (fi21 | (fi23 ^ 882)), F1: fi23, F2: 2}
                                                    p1 = 1
                                                    _ = v24
                                                } else {
                                                    p1 = (p1 % (((p1 / (((p1 % ((p1) | 1))) | 1))) | 1))
                                                    _ = len(v12)
                                                }
                                                p1 -= p1
                                                _ = fi23
                                            }
                                            for fi25 := 0; fi25 < 11; fi25++ {
                                                var v26 uint8 = 255
                                                v26 += v26
                                                v26 -= v26
                                                _ = fi25
                                                _ = v26
                                            }
                                        }
                                        _ = v22
                                    } else {
                                        switch "a" {
                                        case "foo bar":
                                            fmt.Println(p0, fi21, p1)
                                        case "":
                                            p1 = int32(fi21)
                                        case "\n\t":
                                            var v27 int16 = int16(p0)
                                            _ = v27
                                        default:
                                            p1 = (-(p1))
                                        }
                                    }
                                    var v28 int32 = (p1 * p1)
                                    var v29 int = -689
                                    _ = fi21
                                    _ = v28
                                    _ = v29
                                }
                            }
                        } else {
                            v12 = append(v12, 32768)
                            v30 := GenS2{F0: p1}
                            var v31 int16 = 493
                            _ = v30
                            _ = v31
                        }
                    }
                default:
                    for fi32 := 0; fi32 < 4; fi32++ {
                        v33 := GenS5{F0: fi32, F1: (fi32 - 32768), F2: 0}
                        var v34 string = "a"
                        for fi35 := 0; fi35 < 19; fi35++ {
                            if false {
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                                switch (v34 + "a") {
                                case "hello":
                                    var v36 uint = 255
                                    _ = v36
                                case "foo bar":
                                    var v37 uint = 793
                                    _ = v37
                                case "":
                                    v38 := GenS1{F0: -757}
                                    _ = v38
                                }
                            } else {
                                var v39 float32 = (1.5 - 0.0)
                                v40 := GenS3{F0: -1, F1: (p0 << ((4294967295) & 63))}
                                if false {
                                    v41 := GenS5{F0: (fi32 ^ -482), F1: (fi32 & (fi35 | 128)), F2: 2}
                                    v12 = append(v12, fi35)
                                    _ = v41
                                }
                                v42 := GenS6{F0: 1}
                                v43 := GenS2{F0: (p1 << ((2) & 63))}
                                var v44 uint = 2
                                var v45 int = (fi32 - fi35)
                                _ = v39
                                _ = v40
                                _ = v42
                                _ = v43
                                _ = v44
                                _ = v45
                            }
                            fmt.Println(p0, fi32, v34)
                            _ = fi35
                        }
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                }
                for fi46 := 0; fi46 < 5; fi46++ {
                    p0 -= (p0 - (p0 & (p0 - int8(p1))))
                    _ = fi46
                }
            }
        }
        _ = v12
    }
    if (((1.5 + 1.5) * 1.5) <= ((3.14159 + 1e10) - 3.14159)) {
        v47 := GenS5{F0: 1, F1: 1, F2: 18446744073709551615}
        if (4294967295 < 42) {
            fmt.Println(p1)
            v48 := []int{42}
            _ = v48
        }
        _ = v47
    }
    p1 -= p1
    if (true && (!true)) {
        p1 -= (p1 * (p1 & p1))
    } else {
        p0 = p0
        if ((4294967295 == 0) && (p0 < 42)) {
            var v49 uint = 23
            var v50 uint8 = 203
            _ = v49
            _ = v50
        }
    }
    p1 = p1
    _ = p0
    _ = p1
    return 127
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{65535, 982, 2147483648, 551}
    v3 := []string{}
    var v4 uint64 = 2
    var v5 uint = 0
    if false {
        if true {
            var v51 uint = (v5 * (^(v5)))
            _ = v51
        } else {
            v52 := GenS7{F0: 323}
            v53 := GenS8{F0: (-9223372036854775808 > 130)}
            v54 := GenS8{F0: (true || false)}
            _ = v52
            _ = v53
            _ = v54
        }
    }
    switch -819 {
    case 1:
        var v55 int8 = -128
        _ = v55
    default:
        if true {
            fmt.Println(v4)
            fmt.Println(v4, v4)
        } else {
            var v56 uint32 = 333
            v57 := GenS9{F0: 1}
            v56 += 65535
            _ = v56
            _ = v57
        }
    }
    var v58 uint8 = uint8(v4)
    for fi59 := 0; fi59 < 11; fi59++ {
        fmt.Println(v4, v58)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if true {
            var v60 string = "a"
            for fi61 := 0; fi61 < 1; fi61++ {
                v5 += (v5 / ((v5) | 1))
                switch 9223372036854775807 {
                case 4294967295:
                    v62 := GenS4{F0: 42}
                    _ = v62
                default:
                    fmt.Println(v5, v5)
                }
                if (1.5 <= 0.1) {
                    var v63 int8 = -72
                    v63 = 23
                    _ = v63
                } else {
                    v60 = (v60 + v60)
                    v64 := v1[(v60 + "foo bar")]
                    if (591 != 1) {
                        v65 := v1[("a" + v60)]
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v66 := GenS1{F0: 801}
                        _ = v65
                        _ = v66
                    } else {
                        var v67 uint64 = v4
                        v68 := GenS10{F0: -1}
                        var v69 uint16 = 88
                        _ = v67
                        _ = v68
                        _ = v69
                    }
                    _ = v64
                }
                _ = fi61
            }
            _ = v60
        } else {
            for fi70 := 0; fi70 < 10; fi70++ {
                v4 += 1
                _ = fi70
            }
        }
        _ = fi59
    }
    var v71 bool = false
    v71 = (!v71)
    var v72 int16 = 0
    v5 -= 18446744073709551615
    for fi73 := 0; fi73 < 2; fi73++ {
        switch 32768 {
        case 9223372036854775807:
            v72 -= 42
        default:
            v71 = v71
        }
        if true {
            v5 -= v5
            fmt.Println(v4)
            _ = len(v3)
        } else {
            if ((v5 <= v5) || true) {
                var v74 int16 = v72
                var v75 bool = ((266 <= 255) && ("foo bar" <= ""))
                _ = v74
                _ = v75
            } else {
                for fi76 := 0; fi76 < 6; fi76++ {
                    var v77 int16 = 1000
                    v4 -= (v4 % ((v4) | 1))
                    v1["foo bar"] = "\n\t"
                    _ = fi76
                    _ = v77
                }
                switch fi73 {
                case 128:
                    var v78 float64 = 0.0
                    _ = v78
                }
                if false {
                    fmt.Println(fi73, v4)
                    fmt.Println(fi73, v72)
                }
            }
            for fi79 := 0; fi79 < 10; fi79++ {
                switch "hello" {
                case "hello":
                    v58 += v58
                case "a":
                    v4 = (v4 + (v4 ^ v4))
                case "\n\t":
                    v58 += v58
                default:
                    fmt.Println(v5)
                }
                var v80 uint64 = (^(v4))
                var v81 string = ("hello" + "\n\t")
                _ = fi79
                _ = v80
                _ = v81
            }
        }
        _ = fi73
    }
    var v82 int32 = 562
    v71 = v71
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v71
    _ = v72
    _ = v82
}
