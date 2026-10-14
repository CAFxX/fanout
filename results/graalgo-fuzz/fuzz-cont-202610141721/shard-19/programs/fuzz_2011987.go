package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 uint; F2 string }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16 }
func genFn0(p0 int8, p1 uint) int8 {
    var v4 uint16 = 0
    switch "a" {
    case "hello":
        if false {
            if true {
                var v5 float64 = ((float64(1e10) / (0.0 + (float64(1.5) * (1e10 / 1.0)))) - 0.0)
                _ = v5
            }
            var v6 uint16 = 826
            v6 = 634
            _ = v6
        }
    case "\n\t":
        if false {
            v7 := []int{9223372036854775807, 32768}
            _ = v7
        } else {
            var v8 uint32 = 2147483647
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        p0 -= (p0 % ((p0) | 1))
        if (false && false) {
            var v10 uint8 = 255
            _ = v10
        } else {
            v4 += (v4 << ((p1) & 63))
            var v11 int32 = -429
            if false {
                var v12 float64 = 1e10
                for fi13 := 0; fi13 < 15; fi13++ {
                    switch (fi9 + (fi9 * (fi13 << ((p1) & 63)))) {
                    case 4294967295:
                        fmt.Println(v12)
                    }
                    _ = fi13
                }
                var v14 int32 = (v11 - -1)
                _ = v12
                _ = v14
            } else {
                var v15 uint8 = 1
                _ = v15
            }
            _ = v11
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v4
    return (p0 << ((0) & 63))
}
func genFn1(p0 uint8) uint8 {
    if true {
        if false {
            if (32767 <= 1) {
                switch ("" + "") {
                case "":
                    v16 := []int{2147483647, 868}
                    _ = v16
                case "hello":
                    v17 := []int{-2147483648, 0, 32768}
                    _ = v17
                }
                v18 := []int{-60, 0, 2147483647, -9223372036854775808}
                _ = v18
            }
            var v19 int32 = 1
            _ = v19
        }
        var v20 float32 = 0.1
        _ = v20
    } else {
        switch "" {
        case "a":
            fmt.Println(p0)
        }
        v21 := GenS1{F0: 0}
        v22 := GenS2{F0: false, F1: 255, F2: "foo bar"}
        _ = v21
        _ = v22
    }
    p0 = p0
    for fi23 := 0; fi23 < 10; fi23++ {
        v24 := GenS3{F0: (false || (!true)), F1: (-(fi23))}
        for fi25 := 0; fi25 < 4; fi25++ {
            switch (fi25 | (fi25 / ((fi23) | 1))) {
            case 9223372036854775807:
                p0 = 78
            }
            _ = fi25
        }
        p0 += p0
        _ = fi23
        _ = v24
    }
    switch 42 {
    case 32767:
        if true {
            fmt.Println(p0)
            p0 = (p0 << ((557) & 63))
            v26 := []int{-403}
            _ = v26
        }
    }
    _ = p0
    return p0
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 int64 = -1
    var v3 int16 = 32767
    var v27 int64 = (v2 | int64(v1))
    v28 := GenS4{F0: (v1 ^ 0)}
    for fi29 := 0; fi29 < 12; fi29++ {
        for fi30 := 0; fi30 < 3; fi30++ {
            switch "hello" {
            case "":
                var v31 int32 = 2147483647
                _ = v31
            case "\n\t":
                switch ((("\n\t" + "foo bar") + "a") + ("" + "\n\t")) {
                case "hello":
                    v27 -= (v2 + v2)
                case "foo bar":
                    fmt.Println(fi29, v27, fi29)
                case "a":
                    switch fi29 {
                    case 656:
                        v2 += 0
                    case -2147483648:
                        v3 = 32767
                    case 32767:
                        var v32 uint = (v1 * v1)
                        _ = v32
                    default:
                        if true {
                            v33 := []int{}
                            v1 -= v1
                            if len(v33) >= 2 {
                                v33 = v33[1:len(v33)]
                            }
                            _ = v33
                        } else {
                            v3 += v3
                            var v34 int32 = 1
                            _ = v34
                        }
                    }
                }
                var v35 int64 = (v2 << ((v1) & 63))
                switch fi30 {
                case -9223372036854775808:
                    if true {
                        v36 := []int{}
                        fmt.Println(fi30)
                        if (96 != 0) {
                            if false {
                                if true {
                                    var v37 float64 = 3.14159
                                    var v38 string = "hello"
                                    var v39 int64 = v35
                                    v38 = v38
                                    _ = len(v36)
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                }
                            }
                        }
                        if (v2 != (^(v35))) {
                            var v40 int64 = v35
                            _ = v40
                        } else {
                            switch (fi29 / (((fi29 >> ((v1) & 63))) | 1)) {
                            case 127:
                                if false {
                                    v41 := GenS5{F0: 842}
                                    var v42 int64 = (v27 & (v27 | 1))
                                    v35 = -759
                                    var v43 float64 = (1.5 - float64((0.1 * ((0.1 - 0.0) + 0.1))))
                                    _ = v41
                                    _ = v42
                                    _ = v43
                                } else {
                                    v36 = append(v36, fi29)
                                    v44 := GenS5{F0: 0}
                                    fmt.Println(v27, v35, v35)
                                    _ = v44
                                }
                            case -804:
                                switch int(v1) {
                                case -1:
                                    v27 -= 358
                                case 572:
                                    v1 = v1
                                default:
                                    v27 = (v35 - (-(v35)))
                                }
                            case -1:
                                fmt.Println(v35, v2, v1)
                            }
                            v1 += v1
                            var v45 int = fi29
                            _ = v45
                        }
                        v3 -= v3
                        _ = v36
                    } else {
                        v2 = v2
                        v35 -= v2
                        if true {
                            fmt.Println(v3, v3, v27)
                            var v46 int16 = v3
                            v47 := []int{-9223372036854775808, 127, (fi30 >> (((v1 - v1)) & 63)), (fi29 >> (((v1 ^ v1)) & 63))}
                            v35 += v2
                            _ = v46
                            _ = v47
                        }
                        fmt.Println(v27, v2)
                    }
                case 0:
                    var v48 float64 = 0.1
                    _ = v48
                case 2147483647:
                    var v49 uint32 = 2
                    _ = v49
                default:
                    var v50 float64 = 1e10
                    _ = v50
                }
                _ = v35
            case "a":
                v51 := GenS4{F0: v1}
                _ = v51
            }
            if (171 >= 132) {
                if false {
                    var v52 float32 = float32((((0.1 + ((-2.25 - 3.14159) / (1.5 - (3.14159 / float64(3.14159))))) / 1e10) - (1e10 / 1.0)))
                    var v53 int8 = 42
                    var v54 bool = ("\n\t" != ((("\n\t" + (((("foo bar" + (("a" + "hello") + ("a" + "a"))) + "\n\t") + "\n\t") + "")) + "") + "a"))
                    _ = v52
                    _ = v53
                    _ = v54
                }
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(fi29)
                default:
                    fmt.Println(fi29)
                }
            } else {
                var v55 int64 = v27
                v56 := GenS2{F0: (32767 == 693), F1: v1, F2: "hello"}
                _ = v55
                _ = v56
            }
            switch "foo bar" {
            case "":
                var v57 uint32 = 77
                _ = v57
            default:
                fmt.Println(fi30)
            }
            v27 = (v2 ^ 9223372036854775807)
            _ = fi30
        }
        v58 := []int{(fi29 >> ((v1) & 63)), (fi29 & fi29), (fi29 | (fi29 | (-(fi29))))}
        v3 += (v3 - (v3 | (^(v3))))
        _ = fi29
        _ = v58
    }
    v59 := []int{564}
    v1 = (^(v1))
    var v60 int32 = 909
    if (1 < 2) {
        v61 := GenS2{F0: false, F1: (v1 >> (((v1 + 886)) & 63)), F2: ("a" + ((("foo bar" + (((((("\n\t" + "foo bar") + ("\n\t" + "a")) + ("a" + "hello")) + (("" + (("" + (("\n\t" + "hello") + ("foo bar" + ""))) + "\n\t")) + "")) + "") + (("" + ("hello" + "")) + ("hello" + "")))) + "hello") + "a"))}
        if (!false) {
            for fi62 := 0; fi62 < 1; fi62++ {
                v2 = v27
                _ = fi62
            }
        } else {
            if len(v59) >= 2 {
                v59 = v59[1:len(v59)]
            }
            switch "\n\t" {
            case "hello":
                fmt.Println(v1)
            case "a":
                var v63 int32 = -1
                _ = v63
            case "\n\t":
                v64 := GenS2{F0: true, F1: v1, F2: ""}
                _ = v64
            default:
                fmt.Println(v2, v60)
            }
            if len(v59) >= 2 {
                v59 = v59[1:len(v59)]
            }
        }
        _ = v61
    } else {
        var v65 int32 = (v60 | (^(v60)))
        _ = v65
    }
    var v66 uint = v1
    for fi67 := 0; fi67 < 16; fi67++ {
        var v68 uint64 = 0
        var v69 uint = v1
        _ = fi67
        _ = v68
        _ = v69
    }
    for fi70 := 0; fi70 < 11; fi70++ {
        if true {
            var v71 bool = ((v3 >> ((4294967295) & 63)) <= v3)
            _ = v71
        }
        for fi72 := 0; fi72 < 14; fi72++ {
            var v73 int64 = 0
            v74 := GenS3{F0: (true && (!(!true))), F1: 29}
            _ = fi72
            _ = v73
            _ = v74
        }
        if false {
            if len(v59) >= 2 {
                v59 = v59[1:len(v59)]
            }
        }
        _ = fi70
    }
    _ = len(v59)
    _ = len(v59)
    if len(v59) > 0 {
        v59[0] = -905
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v59:", len(v59), v59)
    fmt.Println("v60:", v60)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v59
    _ = v60
    _ = v66
}
