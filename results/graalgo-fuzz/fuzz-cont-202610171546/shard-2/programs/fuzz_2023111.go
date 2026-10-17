package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 int8; F1 bool; F2 uint32 }

func main() {
    var v1 uint = 730
    var v2 uint = v1
    v3 := []int{384, 32767, 0, 259}
    v4 := make(map[string]uint32)
    var v5 float64 = ((0.1 + (3.14159 + 1e10)) * (0.1 / 1.0))
    v1 = v2
    if (973 <= -1) {
        fmt.Println(v2, v1, v1)
        switch 128 {
        case -892:
            var v6 uint8 = 143
            _ = v6
        case 32768:
            switch ("\n\t" + "") {
            case "":
                for fi7 := 0; fi7 < 13; fi7++ {
                    var v8 uint64 = 42
                    var v9 uint32 = 915
                    if true {
                        v10 := GenS1{F0: (v8 | (v8 * 84))}
                        _ = v10
                    }
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            case "hello":
                v1 -= (v2 ^ (^(v2)))
            default:
                if (true || (1 != 756)) {
                    v11 := GenS1{F0: 2}
                    _ = v11
                }
            }
        case 42:
            for fi12 := 0; fi12 < 4; fi12++ {
                if len(v3) > 0 {
                    v3[0] = fi12
                }
                _ = fi12
            }
        default:
            v13 := GenS2{F0: ""}
            _ = v13
        }
    }
    v4[""] = 4294967295
    for fi14 := 0; fi14 < 6; fi14++ {
        v4[("" + ("foo bar" + "hello"))] = 42
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 int16 = 1
            var v17 int64 = 326
            var v18 int32 = 1
            fmt.Println(fi15, v16)
            _ = len(v4)
            switch "foo bar" {
            case "hello":
                if ((v18 - (v18 ^ (^(v18)))) >= (v18 << ((2) & 63))) {
                    v18 += v18
                } else {
                    var v19 string = ("\n\t" + "a")
                    v4[v19] = 65535
                    switch fi15 {
                    case 32768:
                        v1 += v2
                    case -9223372036854775808:
                        v18 += v18
                    }
                    _ = v19
                }
            case "foo bar":
                v20 := GenS1{F0: 2}
                _ = v20
            default:
                v21 := GenS1{F0: 0}
                _ = v21
            }
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        for fi22 := 0; fi22 < 11; fi22++ {
            v2 = (v2 + (v1 / (((v2 % (((v2 >> ((18446744073709551615) & 63))) | 1))) | 1)))
            _ = fi22
        }
        _ = fi14
    }
    switch ("" + "\n\t") {
    case "\n\t":
        for fi23 := 0; fi23 < 10; fi23++ {
            v24 := GenS3{F0: 692, F1: "a"}
            if false {
                for fi25 := 0; fi25 < 4; fi25++ {
                    v26 := GenS2{F0: "hello"}
                    if len(v3) > 0 {
                        v3[0] = fi23
                    }
                    _ = fi25
                    _ = v26
                }
                v4["a"] = 1
            } else {
                var v27 bool = false
                v1 -= v2
                fmt.Println(fi23, v2)
                _ = v27
            }
            _ = fi23
            _ = v24
        }
    case "a":
        if true {
            for fi28 := 0; fi28 < 20; fi28++ {
                var v29 float32 = 1.5
                v1 = v1
                v2 += (v1 | v1)
                _ = fi28
                _ = v29
            }
            for fi30 := 0; fi30 < 18; fi30++ {
                for fi31 := 0; fi31 < 11; fi31++ {
                    if false {
                        if false {
                            switch fi31 {
                            case 32768:
                                v32 := GenS4{F0: 0}
                                _ = v32
                            case -317:
                                fmt.Println(fi31, v1, fi30)
                            }
                            var v33 int8 = 1
                            _ = v33
                        }
                        if (1 <= 32767) {
                            var v34 string = "a"
                            var v35 string = v34
                            _ = v34
                            _ = v35
                        } else {
                            if (!(65535 != 0)) {
                                switch "hello" {
                                case "\n\t":
                                    v36 := v4["hello"]
                                    _ = v36
                                case "foo bar":
                                    _ = len(v4)
                                case "":
                                    v1 = (v1 + v1)
                                default:
                                    v1 = v2
                                }
                                v37 := GenS4{F0: 65535}
                                _ = v37
                            }
                            v2 = 2
                        }
                    } else {
                        switch fi31 {
                        case 2147483647:
                            var v38 float32 = -2.25
                            _ = v38
                        case 128:
                            if (0.0 == v5) {
                                if true {
                                    fmt.Println(v1, v1)
                                    v39 := GenS4{F0: 0}
                                    switch "" {
                                    case "":
                                        _ = len(v4)
                                    }
                                    v40 := GenS1{F0: 99}
                                    _ = v39
                                    _ = v40
                                } else {
                                    v1 += (v1 - v2)
                                }
                            }
                        }
                        if false {
                            var v41 float32 = 1.5
                            v42 := GenS4{F0: 255}
                            var v43 uint16 = 47
                            _ = v41
                            _ = v42
                            _ = v43
                        } else {
                            var v44 string = ("" + "")
                            if true {
                                v2 -= v1
                                _ = len(v4)
                                fmt.Println(fi30, v44, fi31)
                            } else {
                                var v45 int8 = 127
                                _ = v45
                            }
                            v44 = ""
                            _ = v44
                        }
                        if true {
                            var v46 uint32 = 966
                            switch (^(fi31)) {
                            case 9223372036854775807:
                                v46 = v46
                                var v47 int32 = 42
                                _ = len(v3)
                                _ = v47
                            case 127:
                                v46 -= 2147483648
                            case -180:
                                var v48 uint = (v2 | v2)
                                _ = v48
                            default:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            v5 += (0.1 / ((float64(float32((float64((((1.5 / (1e10 - 1.5)) * 1e10) - 0.1)) * ((v5 - (v5 + v5)) + ((3.14159 / 1e10) - v5))))) * float64(((0.0 + 1e10) + 0.1))) - (v5 + (v5 / v5))))
                            v49 := GenS1{F0: 42}
                            var v50 uint8 = 1
                            var v51 string = "a"
                            v52 := GenS4{F0: 42}
                            _ = v46
                            _ = v49
                            _ = v50
                            _ = v51
                            _ = v52
                        } else {
                            var v53 int16 = -287
                            _ = v53
                        }
                    }
                    if (!false) {
                        fmt.Println(v1, fi30, v5)
                        v54 := GenS5{F0: 824, F1: -1}
                        fmt.Println(v5)
                        _ = v54
                    } else {
                        v5 = v5
                        var v55 uint = (v2 + v1)
                        _ = v55
                    }
                    v3 = append(v3, (fi31 - fi31))
                    _ = fi31
                }
                _ = fi30
            }
        } else {
            v1 = v1
            switch "foo bar" {
            case "":
                v5 = (v5 * 3.14159)
            case "a":
                var v56 float32 = 0.1
                _ = v56
            case "hello":
                v5 = v5
            default:
                fmt.Println(v1, v1)
            }
        }
    case "hello":
        delete(v4, "hello")
    default:
        var v57 int64 = 1
        _ = v57
    }
    if ("a" == ("foo bar" + "foo bar")) {
        var v58 uint32 = 4294967295
        _ = v58
    }
    v1 -= v2
    v4[("" + "\n\t")] = 1
    if true {
        v2 += 715
        switch "a" {
        case "a":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "hello":
            v4[("\n\t" + "")] = 1
        case "foo bar":
            if (!(-880 < -753)) {
                switch 32767 {
                case 32767:
                    v2 += 258
                case 907:
                    if false {
                        v5 = v5
                        v59 := GenS6{F0: -36, F1: false, F2: 975}
                        _ = v59
                    } else {
                        if (427 <= 2) {
                            switch -889 {
                            case 32768:
                                v1 = (^(v2))
                            case -106:
                                var v60 uint64 = 0
                                _ = v60
                            case 2147483647:
                                var v61 uint = v2
                                _ = v61
                            }
                            v3 = append(v3, 2147483647)
                            if (42 > 128) {
                                var v62 int64 = 9223372036854775807
                                v3 = append(v3, 9223372036854775807)
                                _ = v62
                            }
                        } else {
                            v63 := GenS4{F0: 223}
                            _ = v63
                        }
                        v5 -= v5
                    }
                case 42:
                    fmt.Println(v5, v1)
                }
                v64 := GenS1{F0: 530}
                _ = v64
            } else {
                fmt.Println(v1, v5, v5)
            }
        }
    } else {
        fmt.Println(v1, v2)
    }
    for fi65 := 0; fi65 < 12; fi65++ {
        if ((false && false) && ((v5 + v5) < v5)) {
            fmt.Println(v2, fi65)
            switch fi65 {
            case 9223372036854775807:
                _ = len(v3)
            case 0:
                var v66 uint32 = 730
                _ = v66
            case 2147483647:
                if ((("foo bar" + "hello") + "hello") >= "a") {
                    for fi67 := 0; fi67 < 17; fi67++ {
                        v4["a"] = 636
                        _ = fi67
                    }
                    switch "a" {
                    case "a":
                        for fi68 := 0; fi68 < 16; fi68++ {
                            switch ("hello" + "\n\t") {
                            case "hello":
                                v2 = 1
                            default:
                                switch (("" + "") + "hello") {
                                case "hello":
                                    fmt.Println(v2)
                                default:
                                    v5 += v5
                                }
                            }
                            switch "a" {
                            case "hello":
                                _ = len(v4)
                            case "":
                                var v69 uint16 = 32767
                                _ = v69
                            default:
                                var v70 int = fi68
                                _ = v70
                            }
                            v5 = 1e10
                            _ = fi68
                        }
                    case "\n\t":
                        v5 = v5
                    case "":
                        v71 := GenS1{F0: 0}
                        _ = v71
                    }
                } else {
                    v2 = (v2 * 415)
                    v5 += float64(((-2.25 / (float32(((v5 - 0.1) - 3.14159)) + (0.1 / 1.0))) + 0.1))
                }
            default:
                _ = len(v3)
            }
        }
        _ = fi65
    }
    var v72 uint16 = 327
    switch 2147483647 {
    case 9223372036854775807:
        for fi73 := 0; fi73 < 13; fi73++ {
            v74 := v4["\n\t"]
            _ = fi73
            _ = v74
        }
    case 32768:
        if true {
            var v75 uint32 = 42
            _ = v75
        }
    case 127:
        var v76 uint16 = v72
        _ = v76
    }
    var v77 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v72:", v72)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v72
    _ = v77
}
