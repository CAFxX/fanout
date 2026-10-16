package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 bool; F2 uint16 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 int64; F2 int16 }
type GenS6 struct { F0 string; F1 int32; F2 int8 }
type GenS7 struct { F0 bool; F1 int }
type GenS8 struct { F0 int8; F1 bool }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int16 }
func genFn0() uint16 {
    var v7 float32 = ((1.5 + (1e10 * 0.0)) + (-2.25 + ((-2.25 - (float32(0.0) - (3.14159 - -2.25))) / 0.1)))
    var v8 bool = true
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 uint16, p1 uint16, p2 uint64) int32 {
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS1{F0: 1000}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := GenS2{F0: false, F1: ((1 > -349) && true), F2: genFn0()}
        p1 += (p0 << ((199) & 63))
        v13 := []int{(fi11 + fi11), fi11, fi11, 9223372036854775807}
        _ = fi11
        _ = v12
        _ = v13
    }
    if ((p0 | 32767) >= p1) {
        switch -9223372036854775808 {
        case -2147483648:
            p1 = p1
        case 222:
            var v14 int = 4294967295
            _ = v14
        }
        p2 += p2
    } else {
        for fi15 := 0; fi15 < 2; fi15++ {
            switch fi15 {
            case 9223372036854775807:
                var v16 int16 = -238
                _ = v16
            case 1:
                var v17 int8 = 55
                _ = v17
            case 42:
                v18 := GenS3{F0: "\n\t", F1: -883}
                _ = v18
            }
            v19 := GenS3{F0: "foo bar", F1: 389}
            _ = fi15
            _ = v19
        }
        p2 -= (p2 - p2)
        if true {
            if (("\n\t" + ("\n\t" + "foo bar")) < "a") {
                var v20 uint32 = 1
                _ = v20
            } else {
                v21 := []int{-1, 127}
                for fi22 := 0; fi22 < 18; fi22++ {
                    if true {
                        if (0 <= -783) {
                            fmt.Println(fi22)
                        }
                        for fi23 := 0; fi23 < 20; fi23++ {
                            var v24 int32 = 1
                            _ = fi23
                            _ = v24
                        }
                    } else {
                        v25 := GenS1{F0: 0}
                        switch ("" + ("\n\t" + ("foo bar" + ("hello" + (("foo bar" + "\n\t") + "hello"))))) {
                        case "foo bar":
                            var v26 uint = 527
                            _ = v26
                        case "\n\t":
                            _ = len(v21)
                        case "hello":
                            for fi27 := 0; fi27 < 14; fi27++ {
                                fmt.Println(p1, p0)
                                switch fi22 {
                                case -917:
                                    v28 := GenS1{F0: 1}
                                    _ = v28
                                case 535:
                                    var v29 int = fi27
                                    _ = v29
                                case 2147483647:
                                    var v30 string = "\n\t"
                                    _ = v30
                                }
                                _ = fi27
                            }
                        default:
                            if (true && false) {
                                if len(v21) >= 2 {
                                    v21 = v21[1:len(v21)]
                                }
                            }
                        }
                        _ = v25
                    }
                    fmt.Println(p2)
                    _ = fi22
                }
                _ = v21
            }
            p2 = p2
            switch -212 {
            case 327:
                v31 := []int{69, 42}
                _ = v31
            default:
                for fi32 := 0; fi32 < 11; fi32++ {
                    fmt.Println(p1, fi32, p0)
                    p0 = (p1 | p0)
                    v33 := []int{}
                    _ = fi32
                    _ = v33
                }
            }
        }
    }
    for fi34 := 0; fi34 < 11; fi34++ {
        var v35 int64 = 0
        if true {
            v36 := GenS4{F0: p0}
            switch ("hello" + "") {
            case "a":
                v37 := GenS2{F0: (true || (p1 >= p1)), F1: false, F2: genFn0()}
                _ = v37
            case "":
                if true {
                    var v38 int = -331
                    _ = v38
                } else {
                    for fi39 := 0; fi39 < 16; fi39++ {
                        var v40 int8 = -128
                        switch "a" {
                        case "foo bar":
                            v41 := GenS2{F0: true, F1: ((p1 < p1) && (2147483647 <= 2147483648)), F2: p1}
                            if true {
                                var v42 uint8 = 2
                                fmt.Println(fi34, p1)
                                v43 := GenS5{F0: v40, F1: v35, F2: 32767}
                                _ = v42
                                _ = v43
                            }
                            switch "a" {
                            case "a":
                                v44 := GenS6{F0: ("foo bar" + "foo bar"), F1: 308, F2: (v40 * (v40 - (v40 / ((v40) | 1))))}
                                v35 -= (v35 | (v35 | v35))
                                if true {
                                    v45 := []int{fi34, int(v35), fi34, 42}
                                    _ = v45
                                } else {
                                    if true {
                                        v46 := []int{4294967295, fi34}
                                        if len(v46) > 0 {
                                            v46[0] = fi34
                                        }
                                        _ = v46
                                    }
                                }
                                _ = v44
                            default:
                                v47 := GenS7{F0: (!false), F1: (^(fi34))}
                                _ = v47
                            }
                            _ = v41
                        case "\n\t":
                            var v48 int16 = int16(p2)
                            _ = v48
                        case "hello":
                            fmt.Println(fi39, fi34, v35)
                        default:
                            var v49 string = "\n\t"
                            v35 = (-(v35))
                            var v50 int8 = 1
                            _ = v49
                            _ = v50
                        }
                        var v51 int16 = -32768
                        _ = fi39
                        _ = v40
                        _ = v51
                    }
                    switch (fi34 | 42) {
                    case 602:
                        var v52 uint64 = 605
                        _ = v52
                    case 2147483647:
                        switch (fi34 ^ (fi34 * fi34)) {
                        case -1:
                            if true {
                                fmt.Println(fi34, p0)
                                p2 = (p2 | 1)
                                if (!false) {
                                    v53 := GenS3{F0: ("foo bar" + "hello"), F1: (v35 / ((v35) | 1))}
                                    v54 := []int{(-(fi34)), (fi34 - -823), 4294967295, (fi34 * (fi34 - (^(fi34))))}
                                    fmt.Println(v35, p1)
                                    _ = v53
                                    _ = v54
                                }
                            } else {
                                fmt.Println(v35, p0)
                                if true {
                                    v35 -= v35
                                    v55 := GenS5{F0: 0, F1: 0, F2: 42}
                                    if true {
                                        v56 := []int{}
                                        _ = v56
                                    }
                                    _ = v55
                                }
                            }
                        case 2147483647:
                            v57 := GenS8{F0: 27, F1: true}
                            _ = v57
                        }
                    case -765:
                        p1 = p1
                    default:
                        switch fi34 {
                        case -2147483648:
                            v35 += 9223372036854775807
                        default:
                            v58 := GenS9{F0: "hello"}
                            _ = v58
                        }
                    }
                }
            case "hello":
                switch "" {
                case "hello":
                    for fi59 := 0; fi59 < 19; fi59++ {
                        v60 := GenS8{F0: -48, F1: true}
                        v61 := GenS7{F0: (!((v35 >> ((4294967295) & 63)) < (v35 + v35))), F1: fi34}
                        var v62 uint = 4294967295
                        var v63 int64 = -645
                        _ = fi59
                        _ = v60
                        _ = v61
                        _ = v62
                        _ = v63
                    }
                default:
                    switch (fi34 - fi34) {
                    case 42:
                        for fi64 := 0; fi64 < 1; fi64++ {
                            var v65 int = (fi34 | -866)
                            _ = fi64
                            _ = v65
                        }
                    }
                }
            default:
                if (42 != 1) {
                    var v66 int16 = -931
                    _ = v66
                }
            }
            _ = v36
        }
        var v67 uint32 = 2147483648
        _ = fi34
        _ = v35
        _ = v67
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 uint32 = 4294967295
    var v2 string = "hello"
    var v3 uint = 4294967295
    var v4 uint = (v3 & v3)
    var v5 uint64 = 42
    var v6 int32 = -886
    var v68 bool = (0 <= -787)
    v69 := []int{-729}
    v6 += 1
    v6 = v6
    var v70 uint64 = (v5 % (((v5 | 0)) | 1))
    var v71 uint = 42
    var v72 uint = 839
    v70 += v70
    if false {
        var v73 uint64 = (v5 * v70)
        _ = v73
    } else {
        var v74 string = "hello"
        _ = v74
    }
    if len(v69) >= 2 {
        v69 = v69[1:len(v69)]
    }
    for fi75 := 0; fi75 < 3; fi75++ {
        if len(v69) >= 2 {
            v69 = v69[1:len(v69)]
        }
        _ = fi75
    }
    fmt.Println(v72)
    fmt.Println(v2)
    v72 -= v4
    v76 := GenS10{F0: -803}
    var v77 uint = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    fmt.Println("slicelen v69:", len(v69), v69)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v69
    _ = v70
    _ = v71
    _ = v72
    _ = v76
    _ = v77
}
