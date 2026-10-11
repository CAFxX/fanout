package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 int16; F2 int8 }
func genFn0(p0 int8) int {
    v5 := GenS1{F0: true, F1: 175, F2: 0}
    v6 := GenS1{F0: (!(2147483647 <= -285)), F1: 1, F2: 2}
    _ = p0
    _ = v5
    _ = v6
    return -690
}
func genFn1() int8 {
    v7 := GenS1{F0: true, F1: 255, F2: 42}
    var v8 uint = 42
    for fi9 := 0; fi9 < 19; fi9++ {
        for fi10 := 0; fi10 < 14; fi10++ {
            switch (fi9 & fi9) {
            case -2147483648:
                var v11 float32 = (float32(((1.5 / float64(3.14159)) + ((3.14159 * (1.5 + float64(0.1))) * 0.0))) * float32(0.1))
                _ = v11
            case 32767:
                var v12 bool = true
                _ = v12
            }
            _ = fi10
        }
        v8 -= 4294967295
        _ = fi9
    }
    var v13 int64 = 0
    _ = v13
    _ = v7
    _ = v8
    return 53
}

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 1
    var v3 int8 = 1
    var v4 bool = (773 == 533)
    for fi14 := 0; fi14 < 5; fi14++ {
        if (v4 || false) {
            for fi15 := 0; fi15 < 10; fi15++ {
                if true {
                    v16 := []int{-319, fi14, fi15, (-(fi14))}
                    _ = v16
                }
                if v4 {
                    switch "a" {
                    case "":
                        v2 = v2
                    }
                }
                v2 = (v2 % (((v2 % ((42) | 1))) | 1))
                v17 := GenS1{F0: v4, F1: 87, F2: 1}
                if v4 {
                    v18 := GenS1{F0: (false && (2 < 183)), F1: 0, F2: 994}
                    v3 -= (v3 * v3)
                    _ = v18
                }
                switch -2147483648 {
                case 128:
                    var v19 float64 = 1.5
                    _ = v19
                case 454:
                    v20 := GenS1{F0: (0.0 >= (0.0 + (-2.25 * 0.1))), F1: 128, F2: 2147483648}
                    _ = v20
                default:
                    switch fi14 {
                    case 4294967295:
                        if v4 {
                            v4 = (v4 && (2147483647 == -2147483648))
                            v2 = (-(v2))
                            v21 := v1["\n\t"]
                            v22 := v1[("a" + (v21 + ("foo bar" + v21)))]
                            _ = v21
                            _ = v22
                        } else {
                            v23 := GenS1{F0: ((1e10 - 0.1) != (3.14159 - (3.14159 / 1.0))), F1: 227, F2: 2147483648}
                            fmt.Println(fi15, fi14, fi14)
                            var v24 int64 = 42
                            var v25 uint16 = 0
                            v24 += 42
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                        var v26 float64 = 1.5
                        var v27 uint8 = 19
                        _ = v26
                        _ = v27
                    case -574:
                        var v28 uint8 = 218
                        _ = v28
                    default:
                        v29 := []int{-92, fi15, genFn0(genFn1())}
                        _ = v29
                    }
                }
                _ = fi15
                _ = v17
            }
        }
        switch (fi14 + (^(fi14))) {
        case 0:
            v2 += (-(v2))
        case 2147483647:
            switch "" {
            case "hello":
                if true {
                    v30 := GenS2{F0: (fi14 + genFn0(v3))}
                    fmt.Println(v2, v4)
                    _ = v30
                } else {
                    if (-2.25 < (float64(0.0) / 1e10)) {
                        v2 -= v2
                        var v31 uint8 = 254
                        _ = v31
                    } else {
                        _ = len(v1)
                        fmt.Println(fi14, fi14, v2)
                        switch (fi14 ^ -9223372036854775808) {
                        case -2147483648:
                            fmt.Println(v4, v2, fi14)
                        case 32768:
                            for fi32 := 0; fi32 < 4; fi32++ {
                                v3 += (v3 ^ genFn1())
                                var v33 float32 = (1.5 / ((-2.25 * ((float32((-2.25 + 0.0)) * 3.14159) + 1.5)) - 0.1))
                                fmt.Println(v4)
                                _ = fi32
                                _ = v33
                            }
                        }
                    }
                    fmt.Println(v2, fi14)
                    fmt.Println(v4, v4)
                }
            case "\n\t":
                var v34 bool = (42 <= 3)
                _ = v34
            case "a":
                if false {
                    var v35 uint64 = 18446744073709551615
                    fmt.Println(v2, v4, v4)
                    _ = v35
                }
            }
        case 32767:
            fmt.Println(fi14, v3)
        default:
            v2 = 42
        }
        v2 = v2
        _ = fi14
    }
    v36 := GenS2{F0: 1}
    v37 := []int{977, 2147483647, -140}
    for fi38 := 0; fi38 < 12; fi38++ {
        for fi39 := 0; fi39 < 6; fi39++ {
            switch 42 {
            case 42:
                v40 := GenS2{F0: fi39}
                _ = v40
            case 128:
                if v4 {
                    var v41 uint16 = 581
                    var v42 uint8 = 2
                    v43 := GenS2{F0: (fi39 / ((fi38) | 1))}
                    _ = v41
                    _ = v42
                    _ = v43
                }
            default:
                fmt.Println(v2, fi38, fi39)
            }
            v44 := GenS2{F0: (fi38 / (((fi38 | (fi38 + (-(fi39))))) | 1))}
            if v4 {
                switch fi39 {
                case 827:
                    switch (fi38 % (((fi38 / ((genFn0((^(v3)))) | 1))) | 1)) {
                    case -521:
                        if len(v37) > 0 {
                            v37[0] = (fi38 >> ((4294967295) & 63))
                        }
                    default:
                        var v45 float32 = -2.25
                        if v4 {
                            v46 := v1[("a" + "hello")]
                            var v47 int64 = v2
                            _ = v46
                            _ = v47
                        }
                        _ = len(v1)
                        _ = v45
                    }
                case -1:
                    v3 -= genFn1()
                default:
                    if (275 >= 18446744073709551615) {
                        fmt.Println(fi39, fi39)
                    }
                }
            } else {
                v1[("a" + ("a" + "hello"))] = "foo bar"
                var v48 bool = true
                v49 := GenS1{F0: (((1 > 538) || false) || (false || (2 >= 42))), F1: 42, F2: 65535}
                var v50 uint64 = 686
                _ = v48
                _ = v49
                _ = v50
            }
            _ = fi39
            _ = v44
        }
        _ = fi38
    }
    switch genFn0(0) {
    case 127:
        if (4294967295 >= 2) {
            for fi51 := 0; fi51 < 19; fi51++ {
                if ((1.5 < float64(0.0)) && (1 <= 42)) {
                    delete(v1, "a")
                    for fi52 := 0; fi52 < 1; fi52++ {
                        var v53 int = (fi51 >> ((1) & 63))
                        _ = fi52
                        _ = v53
                    }
                }
                v4 = (1 >= 55)
                v54 := GenS2{F0: (^(fi51))}
                _ = fi51
                _ = v54
            }
            var v55 int64 = v2
            if (42 > 125) {
                var v56 uint32 = 42
                for fi57 := 0; fi57 < 11; fi57++ {
                    _ = len(v37)
                    _ = len(v37)
                    _ = fi57
                }
                _ = v56
            } else {
                v4 = (!((v55 | v2) <= (v55 / ((v55) | 1))))
                if v4 {
                    var v58 uint32 = 1
                    _ = v58
                }
                v55 -= (v2 ^ v55)
            }
            _ = v55
        } else {
            for fi59 := 0; fi59 < 9; fi59++ {
                v60 := GenS1{F0: true, F1: 255, F2: 125}
                v2 += v2
                var v61 bool = v4
                _ = fi59
                _ = v60
                _ = v61
            }
        }
    case 535:
        v3 += v3
    case 32767:
        var v62 int32 = 42
        _ = v62
    }
    if (false || v4) {
        v2 = v2
        fmt.Println(v3)
        delete(v1, ((("hello" + "foo bar") + "\n\t") + ("foo bar" + "foo bar")))
    } else {
        switch "a" {
        case "foo bar":
            v63 := GenS1{F0: v4, F1: 245, F2: 404}
            _ = v63
        case "a":
            var v64 uint64 = 1
            _ = v64
        case "\n\t":
            for fi65 := 0; fi65 < 20; fi65++ {
                v66 := v1[("hello" + "a")]
                if len(v37) > 0 {
                    v37[0] = (-(fi65))
                }
                _ = fi65
                _ = v66
            }
        }
        v4 = (-589 >= -32768)
        v3 += (v3 ^ v3)
    }
    var v67 int64 = (v2 - (^(v2)))
    var v68 int32 = -2147483648
    v67 += v2
    if v4 {
        for fi69 := 0; fi69 < 13; fi69++ {
            var v70 uint64 = 1
            if v4 {
                if (v67 > v67) {
                    v4 = v4
                    _ = len(v37)
                    v71 := GenS3{F0: false}
                    _ = v71
                }
            } else {
                var v72 int64 = (v2 << ((255) & 63))
                var v73 int8 = (-(v3))
                _ = v72
                _ = v73
            }
            v74 := v1[("\n\t" + "hello")]
            _ = fi69
            _ = v70
            _ = v74
        }
    }
    for fi75 := 0; fi75 < 14; fi75++ {
        switch 9223372036854775807 {
        case 128:
            if v4 {
                v76 := GenS4{F0: 750, F1: 1, F2: v3}
                var v77 int = (^(fi75))
                if (309 > 42) {
                    for fi78 := 0; fi78 < 20; fi78++ {
                        if (((0.1 + -2.25) * ((-2.25 + ((1e10 - 1e10) * float64(float32((1.5 - ((3.14159 / (0.0 - 1e10)) + 0.0)))))) * -2.25)) <= float64(float32(((3.14159 / ((float64(0.0) / 1.5) - 1e10)) + 0.1)))) {
                            v68 += int32(v3)
                        } else {
                            if ((!((1.5 <= 1.5) && (!(!v4)))) && (true || (v68 >= 670))) {
                                if (!(0.0 == (-2.25 + float64((3.14159 + 1.5))))) {
                                    delete(v1, ("\n\t" + ("hello" + ("foo bar" + "a"))))
                                }
                                if ((!v4) || false) {
                                    var v79 bool = true
                                    _ = v79
                                } else {
                                    var v80 float64 = 1e10
                                    _ = v80
                                }
                            }
                            if len(v37) > 0 {
                                v37[0] = 838
                            }
                            var v81 int32 = v68
                            _ = v81
                        }
                        if (true || (-32768 == -288)) {
                            v4 = v4
                            if (!(((v2 >> ((890) & 63)) == (v2 + 1)) || (genFn0(v3) >= (v77 << ((941) & 63))))) {
                                if len(v37) > 0 {
                                    v37[0] = (^(fi75))
                                }
                                var v82 float64 = 0.0
                                if (-531 > -1) {
                                    v3 = 127
                                    switch (fi75 ^ v77) {
                                    case 32768:
                                        v83 := GenS3{F0: v4}
                                        _ = v83
                                    case -759:
                                        if (v3 != v3) {
                                            v67 = v2
                                            v77 -= fi75
                                        }
                                    }
                                } else {
                                    v84 := GenS1{F0: (float32((1.5 + v82)) != 1.5), F1: 185, F2: 928}
                                    v37 = append(v37, (fi75 << ((1) & 63)))
                                    _ = v84
                                }
                                _ = v82
                            }
                            v85 := v1[""]
                            switch (^(fi78)) {
                            case 907:
                                _ = len(v1)
                            }
                            fmt.Println(v67)
                            v86 := GenS1{F0: v4, F1: 2, F2: 0}
                            v3 = (v3 & genFn1())
                            _ = v85
                            _ = v86
                        } else {
                            var v87 float64 = 0.1
                            var v88 int32 = v68
                            _ = v87
                            _ = v88
                        }
                        var v89 uint = 65535
                        _ = fi78
                        _ = v89
                    }
                    v77 = fi75
                    var v90 int32 = (v68 | v68)
                    _ = v90
                }
                _ = v76
                _ = v77
            } else {
                var v91 int = fi75
                var v92 int64 = v67
                _ = v91
                _ = v92
            }
        case -282:
            var v93 int16 = 67
            _ = v93
        default:
            var v94 int32 = 604
            _ = v94
        }
        v37 = append(v37, 35)
        var v95 uint32 = 42
        _ = fi75
        _ = v95
    }
    switch "" {
    case "\n\t":
        v96 := GenS3{F0: v4}
        _ = v96
    case "":
        for fi97 := 0; fi97 < 13; fi97++ {
            v68 = 42
            switch ((("" + "a") + "hello") + ("a" + "hello")) {
            case "foo bar":
                v3 = genFn1()
            default:
                var v98 bool = false
                _ = v98
            }
            _ = fi97
        }
    case "a":
        var v99 string = "hello"
        _ = v99
    }
    v2 = -615
    v37 = append(v37, 128)
    var v100 bool = (!v4)
    v37 = append(v37, 2147483647)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v67
    _ = v68
}
