package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int8; F1 int8; F2 int16 }
func genFn0(p0 bool) uint32 {
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 uint16 = 24
        v9 := []int{fi7, fi7, (fi7 >> ((0) & 63)), fi7}
        switch (fi7 * fi7) {
        case 4294967295:
            var v10 bool = (v8 < 589)
            _ = v10
        case -2147483648:
            var v11 uint32 = 4294967295
            _ = v11
        default:
            for fi12 := 0; fi12 < 13; fi12++ {
                fmt.Println(v8)
                p0 = (1 <= 2)
                var v13 int16 = 1
                v9 = append(v9, 543)
                var v14 uint = 42
                var v15 uint = v14
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    v16 := GenS1{F0: 0}
    for fi17 := 0; fi17 < 17; fi17++ {
        v18 := GenS2{F0: 51}
        v19 := GenS2{F0: -96}
        _ = fi17
        _ = v18
        _ = v19
    }
    switch 32767 {
    case -9223372036854775808:
        if p0 {
            p0 = (!true)
            p0 = (559 < 972)
        }
    case 856:
        v20 := GenS3{F0: 255, F1: 1}
        _ = v20
    case 4294967295:
        for fi21 := 0; fi21 < 2; fi21++ {
            v22 := GenS3{F0: 721, F1: -306}
            for fi23 := 0; fi23 < 19; fi23++ {
                fmt.Println(p0, fi23)
                var v24 int8 = 0
                _ = fi23
                _ = v24
            }
            switch (fi21 / (((fi21 & fi21)) | 1)) {
            case -9223372036854775808:
                v25 := GenS1{F0: 42}
                _ = v25
            default:
                p0 = ((41 >= 32767) && p0)
            }
            _ = fi21
            _ = v22
        }
    default:
        for fi26 := 0; fi26 < 9; fi26++ {
            if p0 {
                p0 = p0
                switch "" {
                case "foo bar":
                    var v27 int32 = -2147483648
                    _ = v27
                case "hello":
                    if ((2 >= 2) || p0) {
                        if p0 {
                            for fi28 := 0; fi28 < 19; fi28++ {
                                var v29 uint64 = 42
                                var v30 int64 = 0
                                v31 := GenS2{F0: 100}
                                v32 := GenS3{F0: 255, F1: -2147483648}
                                var v33 uint64 = (v29 >> ((2) & 63))
                                v33 -= 1
                                _ = fi28
                                _ = v29
                                _ = v30
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                            if false {
                                fmt.Println(p0, p0)
                                p0 = (p0 || p0)
                                if p0 {
                                    var v34 uint = 42
                                    var v35 int16 = 0
                                    v35 = v35
                                    _ = v34
                                    _ = v35
                                } else {
                                    fmt.Println(p0, fi26)
                                    var v36 uint = 513
                                    fmt.Println(p0, p0, fi26)
                                    _ = v36
                                }
                            } else {
                                var v37 float64 = 1e10
                                v37 = 1.5
                                _ = v37
                            }
                        } else {
                            p0 = p0
                            v38 := GenS1{F0: 0}
                            for fi39 := 0; fi39 < 8; fi39++ {
                                var v40 uint8 = 169
                                p0 = ((p0 && p0) || p0)
                                _ = fi39
                                _ = v40
                            }
                            _ = v38
                        }
                        v41 := []int{fi26}
                        p0 = p0
                        _ = v41
                    } else {
                        v42 := GenS4{F0: p0, F1: 42}
                        v43 := GenS5{F0: 4294967295, F1: 288}
                        _ = v42
                        _ = v43
                    }
                case "":
                    for fi44 := 0; fi44 < 14; fi44++ {
                        v45 := GenS4{F0: (29 <= -9223372036854775808), F1: 910}
                        _ = fi44
                        _ = v45
                    }
                default:
                    fmt.Println(p0, p0)
                }
                if (p0 || (!true)) {
                    p0 = false
                    fmt.Println(fi26, p0)
                } else {
                    fmt.Println(p0, fi26)
                    if (-1 == -838) {
                        var v46 uint32 = 79
                        if (851 < 0) {
                            fmt.Println(v46)
                        } else {
                            v47 := []int{42}
                            _ = v47
                        }
                        v48 := GenS5{F0: (v46 >> ((42) & 63)), F1: 0}
                        _ = v46
                        _ = v48
                    } else {
                        for fi49 := 0; fi49 < 10; fi49++ {
                            var v50 uint = 1
                            v51 := []int{fi49, (^(fi49)), (^(fi26)), fi49}
                            var v52 uint = (v50 ^ (^(v50)))
                            _ = fi49
                            _ = v50
                            _ = v51
                            _ = v52
                        }
                        p0 = (128 != 0)
                    }
                }
            }
            if true {
                v53 := GenS6{F0: 215, F1: 930}
                _ = v53
            } else {
                p0 = (!(0 > 168))
                v54 := []int{fi26, (^(fi26))}
                if (p0 || p0) {
                    fmt.Println(p0, p0)
                    var v55 float64 = float64(float32(((3.14159 / 1.5) / (1.5 / (1e10 * 1.5)))))
                    fmt.Println(v55, v55, p0)
                    _ = v55
                }
                _ = v54
            }
            var v56 float64 = (1e10 + (1e10 / 1.0))
            _ = fi26
            _ = v56
        }
    }
    _ = p0
    _ = v16
    return 0
}

func main() {
    var v1 int8 = -1
    var v2 uint = 1
    var v3 float32 = 0.0
    v4 := []uint32{}
    v5 := []string{}
    var v6 int8 = v1
    v3 += (v3 * 1e10)
    v5 = append(v5, "\n\t")
    switch -1 {
    case -1:
        for fi57 := 0; fi57 < 4; fi57++ {
            v6 += 127
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi57
        }
    case 127:
        if (((0 <= -535) || (!(false && false))) && false) {
            var v58 int64 = 1
            v4 = append(v4, 2)
            _ = v58
        } else {
            if (!false) {
                for fi59 := 0; fi59 < 9; fi59++ {
                    v1 = (v1 - v6)
                    for fi60 := 0; fi60 < 6; fi60++ {
                        if true {
                            v3 = 1.5
                            v61 := GenS7{F0: (v6 + v1)}
                            v1 = 7
                            _ = v61
                        }
                        v4 = append(v4, 440)
                        v2 = 2
                        _ = fi60
                    }
                    _ = fi59
                }
                var v62 uint64 = 18446744073709551615
                v1 -= (^(v6))
                _ = v62
            }
        }
    case 9223372036854775807:
        v3 = v3
    }
    for fi63 := 0; fi63 < 9; fi63++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = len(v5)
        _ = fi63
    }
    for fi64 := 0; fi64 < 7; fi64++ {
        var v65 uint64 = 42
        _ = fi64
        _ = v65
    }
    var v66 int64 = 0
    switch "hello" {
    case "hello":
        v6 = (v1 | (-(v1)))
    case "\n\t":
        v2 += (^(v2))
    default:
        v67 := GenS8{F0: (v6 | 0), F1: (v1 - v6), F2: 1000}
        _ = v67
    }
    v1 -= (v6 - (^(v1)))
    switch ("" + (("a" + "foo bar") + "foo bar")) {
    case "a":
        v3 = (v3 - v3)
    default:
        for fi68 := 0; fi68 < 2; fi68++ {
            v1 = v1
            var v69 int64 = (v66 ^ (-(v66)))
            if true {
                v69 = 42
            }
            _ = fi68
            _ = v69
        }
    }
    var v70 int64 = (v66 + 9223372036854775807)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v70
}
