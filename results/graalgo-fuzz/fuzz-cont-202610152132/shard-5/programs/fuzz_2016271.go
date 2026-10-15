package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 int64; F1 uint64; F2 int }
type GenS5 struct { F0 int }
func genFn0(p0 int, p1 int8, p2 int64) int16 {
    v6 := []int{}
    p0 = p0
    switch "" {
    case "foo bar":
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 string = "foo bar"
            v9 := GenS1{F0: 0}
            _ = fi7
            _ = v8
            _ = v9
        }
    case "hello":
        var v10 uint64 = 568
        _ = v10
    case "":
        switch ("" + ("foo bar" + (("hello" + "\n\t") + ("" + "\n\t")))) {
        case "hello":
            var v11 bool = (true || false)
            _ = v11
        }
    default:
        if (55 != 127) {
            if (42 >= 60) {
                if (2 != 65535) {
                    for fi12 := 0; fi12 < 15; fi12++ {
                        var v13 int32 = 42
                        _ = fi12
                        _ = v13
                    }
                } else {
                    switch (((("foo bar" + (("a" + "hello") + ((("hello" + "") + ("" + "\n\t")) + ("" + "")))) + "a") + "") + ("hello" + "foo bar")) {
                    case "foo bar":
                        var v14 float64 = -2.25
                        _ = v14
                    case "\n\t":
                        if true {
                            p1 += p1
                            v15 := GenS2{F0: 2}
                            if false {
                                var v16 int64 = p2
                                _ = v16
                            }
                            _ = v15
                        }
                    default:
                        v6 = append(v6, (p0 - 2147483647))
                    }
                    if (702 > 0) {
                        p0 -= p0
                        if false {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            v6 = append(v6, p0)
                        } else {
                            p0 = -2147483648
                            fmt.Println(p2)
                            var v17 int8 = (p1 % ((p1) | 1))
                            _ = v17
                        }
                        v18 := GenS3{F0: p2, F1: 165}
                        _ = v18
                    } else {
                        for fi19 := 0; fi19 < 8; fi19++ {
                            fmt.Println(p0)
                            v20 := GenS3{F0: p2, F1: 163}
                            if (!(0 >= 1)) {
                                var v21 float32 = 0.0
                                p2 = (p2 ^ p2)
                                _ = v21
                            }
                            _ = fi19
                            _ = v20
                        }
                    }
                }
            } else {
                _ = len(v6)
            }
            for fi22 := 0; fi22 < 9; fi22++ {
                var v23 int8 = (p1 << ((42) & 63))
                p1 = (-(p1))
                _ = fi22
                _ = v23
            }
        }
    }
    v6 = append(v6, 4294967295)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1000
}
func genFn1(p0 string) int {
    var v24 uint32 = 476
    v25 := GenS2{F0: 75}
    var v26 string = p0
    _ = p0
    _ = v24
    _ = v25
    _ = v26
    return -1
}

func main() {
    var v1 int32 = -1
    var v2 uint16 = 576
    v3 := make(map[uint32]int)
    var v4 int = -92
    v5 := []string{"\n\t", ("\n\t" + ("a" + ("hello" + "hello")))}
    v27 := GenS2{F0: 0}
    var v28 string = "foo bar"
    switch (((v28 + ("\n\t" + v28)) + "\n\t") + ("foo bar" + v28)) {
    case "a":
        v4 = v4
    case "hello":
        var v29 uint8 = 42
        _ = v29
    case "":
        if false {
            v2 += (v2 & v2)
            var v30 int16 = -646
            var v31 int8 = 97
            _ = v30
            _ = v31
        } else {
            _ = len(v3)
            var v32 uint8 = 39
            _ = v32
        }
    }
    if ((-(v1)) <= 1) {
        switch 32767 {
        case 0:
            switch v4 {
            case 0:
                var v33 int32 = v1
                _ = v33
            case 9223372036854775807:
                _ = len(v3)
            default:
                switch v4 {
                case 128:
                    v34 := GenS3{F0: -9223372036854775808, F1: 728}
                    _ = v34
                case 579:
                    v4 -= (v4 % (((^(v4))) | 1))
                }
            }
        case 42:
            _ = len(v3)
        case 2147483647:
            v35 := GenS1{F0: uint32(v2)}
            _ = v35
        }
    } else {
        for fi36 := 0; fi36 < 2; fi36++ {
            for fi37 := 0; fi37 < 6; fi37++ {
                if (42 > 1) {
                    var v38 float32 = float32(3.14159)
                    _ = v38
                } else {
                    switch v28 {
                    case "a":
                        v4 = genFn1(v28)
                        v2 = v2
                        if ((false || ((157 != 127) || ((3.14159 >= 1.5) || (42 == -1)))) || true) {
                            var v39 int = (fi37 >> ((2) & 63))
                            if (255 != 18446744073709551615) {
                                switch (v28 + v28) {
                                case "hello":
                                    if true {
                                        var v40 int16 = -213
                                        _ = v40
                                    }
                                case "":
                                    v1 = 75
                                }
                                v41 := GenS1{F0: 65535}
                                _ = v41
                            } else {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            fmt.Println(v4, fi36, v2)
                            _ = v39
                        } else {
                            var v42 uint16 = (^(v2))
                            if false {
                                var v43 uint64 = 0
                                _ = v43
                            } else {
                                v44 := GenS2{F0: 1}
                                switch fi36 {
                                case 2147483647:
                                    if true {
                                        if true {
                                            var v45 int16 = 880
                                            v45 += 373
                                            fmt.Println(v45, v45)
                                            _ = v45
                                        } else {
                                            v4 += int(v1)
                                            v46 := GenS3{F0: 42, F1: 950}
                                            _ = len(v3)
                                            var v47 uint64 = 342
                                            v5 = append(v5, ("a" + v28))
                                            _ = len(v3)
                                            var v48 int8 = -77
                                            _ = v46
                                            _ = v47
                                            _ = v48
                                        }
                                    } else {
                                        v2 += v2
                                    }
                                case 643:
                                    v4 -= fi36
                                case 1:
                                    if (202 <= 42) {
                                        var v49 uint64 = 1
                                        v50 := GenS3{F0: 954, F1: (v49 & (^(v49)))}
                                        var v51 uint16 = 42
                                        _ = v49
                                        _ = v50
                                        _ = v51
                                    } else {
                                        var v52 int16 = 641
                                        var v53 int16 = 32767
                                        v52 = (v53 * (v53 ^ (v52 & (v52 | (v52 - v53)))))
                                        _ = v52
                                        _ = v53
                                    }
                                }
                                v54 := v3[42]
                                switch (v54 / ((fi37) | 1)) {
                                case 154:
                                    if ((false || (!(true || false))) && (18446744073709551615 >= 37)) {
                                        switch v28 {
                                        case "":
                                            if (!(!false)) {
                                                if true {
                                                    switch (v28 + v28) {
                                                    case "hello":
                                                        var v55 bool = false
                                                        _ = v55
                                                    default:
                                                        var v56 bool = false
                                                        _ = v56
                                                    }
                                                    fmt.Println(v4)
                                                    if true {
                                                        var v57 int32 = v1
                                                        _ = v57
                                                    }
                                                    _ = len(v5)
                                                } else {
                                                    var v58 string = v28
                                                    _ = v58
                                                }
                                                v3[42] = -2147483648
                                                fmt.Println(v1, v1)
                                                v42 -= 864
                                                fmt.Println(v54, fi36, v2)
                                            }
                                        case "foo bar":
                                            var v59 int8 = 0
                                            _ = v59
                                        case "hello":
                                            v4 -= fi36
                                        }
                                        switch (fi36 & fi37) {
                                        case 350:
                                            v60 := GenS4{F0: 42, F1: 1, F2: (v4 >> ((uint(v42)) & 63))}
                                            _ = v60
                                        case -9223372036854775808:
                                            if false {
                                                fmt.Println(fi37)
                                                switch "\n\t" {
                                                case "foo bar":
                                                    _ = len(v3)
                                                default:
                                                    var v61 int16 = -404
                                                    _ = v61
                                                }
                                                v2 = v2
                                            } else {
                                                fmt.Println(fi37)
                                                v62 := GenS5{F0: (v54 + fi37)}
                                                var v63 int = (fi37 ^ (v54 + fi36))
                                                _ = v62
                                                _ = v63
                                            }
                                        default:
                                            var v64 float64 = 1.5
                                            _ = v64
                                        }
                                        fmt.Println(v54, v2)
                                    }
                                default:
                                    switch v4 {
                                    case -1:
                                        v42 = (v2 / (((^(v2))) | 1))
                                    }
                                }
                                _ = v44
                                _ = v54
                            }
                            switch (^(v4)) {
                            case 0:
                                fmt.Println(v2)
                                v42 -= (^(v42))
                                if true {
                                    var v65 int32 = 0
                                    if len(v5) > 0 {
                                        v5[0] = (v28 + ((v28 + v28) + v28))
                                    }
                                    fmt.Println(v42, v1, fi36)
                                    _ = v65
                                }
                                v28 = "foo bar"
                            case 42:
                                v5 = append(v5, "a")
                            default:
                                switch v28 {
                                case "\n\t":
                                    fmt.Println(v4)
                                default:
                                    var v66 uint64 = 4294967295
                                    _ = v66
                                }
                            }
                            v28 = "foo bar"
                            _ = v42
                        }
                    case "\n\t":
                        v4 = (v4 * v4)
                        var v67 uint32 = 599
                        _ = v67
                    }
                    var v68 float64 = (0.0 / ((0.1 + 0.0) + 0.1))
                    v5 = append(v5, v28)
                    _ = v68
                }
                _ = fi37
            }
            _ = fi36
        }
        fmt.Println(v2)
        v4 += genFn1(v28)
    }
    v5 = append(v5, "\n\t")
    var v69 uint32 = 705
    fmt.Println(v69)
    for fi70 := 0; fi70 < 19; fi70++ {
        var v71 bool = true
        var v72 bool = (v28 >= ("\n\t" + ""))
        if v72 {
            var v73 float32 = 3.14159
            fmt.Println(v4, v69, v1)
            _ = v73
        }
        _ = fi70
        _ = v71
        _ = v72
    }
    v3[v69] = int(v2)
    for fi74 := 0; fi74 < 11; fi74++ {
        for fi75 := 0; fi75 < 20; fi75++ {
            var v76 uint16 = (v2 ^ 458)
            var v77 int = genFn1("")
            v2 -= (v76 & v2)
            _ = fi75
            _ = v76
            _ = v77
        }
        var v78 int = fi74
        _ = fi74
        _ = v78
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v69
}
