package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 string; F2 uint64 }
type GenS5 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint32; F1 uint16 }
type GenS9 struct { F0 uint32; F1 int16; F2 int }
func genFn0() int64 {
    if false {
        var v7 string = "hello"
        v7 = v7
        var v8 int16 = 527
        _ = v7
        _ = v8
    } else {
        var v9 string = "hello"
        for fi10 := 0; fi10 < 10; fi10++ {
            for fi11 := 0; fi11 < 12; fi11++ {
                if (-934 >= 2147483647) {
                    var v12 int = fi11
                    _ = v12
                } else {
                    v9 = ("foo bar" + "\n\t")
                }
                v13 := GenS1{F0: 671, F1: 217, F2: 2}
                switch 178 {
                case 32768:
                    var v14 float64 = ((1e10 + 3.14159) / float64(1e10))
                    _ = v14
                case -632:
                    v9 = v9
                }
                _ = fi11
                _ = v13
            }
            var v15 float64 = 1.5
            var v16 int32 = 2147483647
            _ = fi10
            _ = v15
            _ = v16
        }
        if (!(-2.25 <= 0.1)) {
            var v17 int = 840
            v18 := GenS2{F0: 2, F1: 255, F2: 0}
            for fi19 := 0; fi19 < 7; fi19++ {
                switch (v17 / ((v17) | 1)) {
                case -2147483648:
                    fmt.Println(v9, v17, fi19)
                case 32767:
                    var v20 int8 = -128
                    _ = v20
                case 710:
                    v21 := []int{(fi19 * -762)}
                    _ = v21
                default:
                    v22 := GenS3{F0: 1}
                    _ = v22
                }
                for fi23 := 0; fi23 < 6; fi23++ {
                    if true {
                        if ((0.1 - 1.5) < 1.5) {
                            if (42 > 377) {
                                v24 := GenS1{F0: 18446744073709551615, F1: 42, F2: 1}
                                _ = v24
                            }
                            v17 -= fi19
                            var v25 string = "hello"
                            _ = v25
                        }
                    }
                    if (!false) {
                        v17 = (v17 ^ v17)
                        v9 = (v9 + ("hello" + (("a" + v9) + ("hello" + v9))))
                        switch v17 {
                        case 32767:
                            v17 -= v17
                        case -999:
                            var v26 int8 = -3
                            _ = v26
                        default:
                            v9 = (("\n\t" + (v9 + ("" + "foo bar"))) + v9)
                        }
                    } else {
                        v17 = (fi19 + 32767)
                    }
                    v17 -= (fi19 + 657)
                    _ = fi23
                }
                _ = fi19
            }
            _ = v17
            _ = v18
        } else {
            switch v9 {
            case "":
                v9 = "a"
            case "\n\t":
                v27 := GenS4{F0: ((3.14159 < -2.25) && ((true && ((0 >= 240) && true)) && false)), F1: v9, F2: 1}
                _ = v27
            case "foo bar":
                switch -154 {
                case 2147483647:
                    for fi28 := 0; fi28 < 15; fi28++ {
                        switch (fi28 / ((fi28) | 1)) {
                        case 32767:
                            fmt.Println(fi28)
                        case -9223372036854775808:
                            v29 := []int{}
                            _ = v29
                        case -173:
                            if true {
                                fmt.Println(v9)
                            } else {
                                var v30 int = fi28
                                _ = v30
                            }
                        }
                        for fi31 := 0; fi31 < 8; fi31++ {
                            switch (("" + (v9 + v9)) + (("foo bar" + v9) + v9)) {
                            case "":
                                var v32 int8 = 100
                                _ = v32
                            case "a":
                                v9 = v9
                            case "foo bar":
                                var v33 uint64 = 0
                                _ = v33
                            default:
                                v9 = v9
                            }
                            _ = fi31
                        }
                        v34 := []int{}
                        _ = fi28
                        _ = v34
                    }
                default:
                    if false {
                        var v35 bool = false
                        if (0 > 2147483647) {
                            for fi36 := 0; fi36 < 9; fi36++ {
                                v9 = ""
                                v9 = ((v9 + v9) + (v9 + v9))
                                _ = fi36
                            }
                        } else {
                            var v37 int64 = 1
                            var v38 float64 = 1.5
                            _ = v37
                            _ = v38
                        }
                        _ = v35
                    }
                }
            }
            v39 := GenS5{F0: 4294967295, F1: 118, F2: 255}
            _ = v39
        }
        _ = v9
    }
    var v40 float64 = (3.14159 - 1.5)
    v41 := GenS6{F0: 9223372036854775807}
    v40 -= (v40 / v40)
    _ = v40
    _ = v41
    return 697
}
func genFn1(p0 int64) uint16 {
    p0 = genFn0()
    var v42 int8 = 42
    _ = p0
    _ = v42
    return 0
}

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 42
    var v3 int = 128
    var v4 uint8 = (v2 - 230)
    var v5 uint = 42
    var v6 int8 = 117
    var v43 float64 = (((0.1 * 3.14159) / 1e10) / 3.14159)
    var v44 float32 = 1e10
    if (65535 > 597) {
        v1[(("foo bar" + (("\n\t" + "hello") + "\n\t")) + "")] = "\n\t"
        for fi45 := 0; fi45 < 9; fi45++ {
            var v46 int = (v3 + (v3 & fi45))
            var v47 int8 = int8(v46)
            _ = fi45
            _ = v46
            _ = v47
        }
    }
    v48 := []int{}
    v2 = v4
    var v49 int16 = 1
    var v50 int = (v3 - v3)
    v43 = (v43 * 0.1)
    if true {
        if (float32(v43) >= (float32(v43) * (0.1 - 3.14159))) {
            var v51 int16 = 32767
            _ = v51
        } else {
            if (v50 >= (v50 % (((v50 << ((v5) & 63))) | 1))) {
                v4 -= v4
                for fi52 := 0; fi52 < 13; fi52++ {
                    if true {
                        for fi53 := 0; fi53 < 17; fi53++ {
                            var v54 int = 42
                            _ = fi53
                            _ = v54
                        }
                    }
                    if ((!false) && ("hello" != "")) {
                        var v55 uint = (v5 << (((v5 >> ((42) & 63))) & 63))
                        v56 := GenS7{F0: ""}
                        var v57 int8 = (v6 & -1)
                        _ = v55
                        _ = v56
                        _ = v57
                    }
                    var v58 string = "a"
                    _ = fi52
                    _ = v58
                }
            } else {
                if true {
                    v6 = v6
                    for fi59 := 0; fi59 < 14; fi59++ {
                        if (!(true || true)) {
                            var v60 uint32 = 42
                            _ = v60
                        }
                        switch ((("hello" + "") + "\n\t") + "\n\t") {
                        case "foo bar":
                            var v61 float32 = 3.14159
                            _ = v61
                        }
                        for fi62 := 0; fi62 < 12; fi62++ {
                            v44 = 1e10
                            var v63 float64 = v43
                            var v64 int64 = 42
                            _ = fi62
                            _ = v63
                            _ = v64
                        }
                        _ = fi59
                    }
                }
                for fi65 := 0; fi65 < 16; fi65++ {
                    var v66 uint = v5
                    for fi67 := 0; fi67 < 16; fi67++ {
                        if (true && true) {
                            var v68 int8 = v6
                            _ = v68
                        }
                        _ = fi67
                    }
                    delete(v1, "foo bar")
                    _ = fi65
                    _ = v66
                }
                if ((!(!(!false))) && false) {
                    if len(v48) >= 2 {
                        v48 = v48[1:len(v48)]
                    }
                    v69 := GenS8{F0: 2147483648, F1: 2}
                    _ = v69
                }
            }
            _ = len(v1)
            if false {
                var v70 int64 = -938
                v43 += v43
                _ = v70
            }
        }
        _ = len(v48)
        v48 = append(v48, -2147483648)
    } else {
        v4 -= v4
        v71 := GenS9{F0: 246, F1: 0, F2: v3}
        v72 := GenS6{F0: v50}
        _ = v71
        _ = v72
    }
    v1[("\n\t" + ("a" + ""))] = ("\n\t" + "hello")
    var v73 int16 = (-(v49))
    for fi74 := 0; fi74 < 8; fi74++ {
        switch (^(v3)) {
        case 0:
            if true {
                var v75 bool = false
                _ = v75
            }
        case -2147483648:
            if false {
                v2 -= v4
                for fi76 := 0; fi76 < 16; fi76++ {
                    v2 += (v2 * (v2 - v2))
                    _ = fi76
                }
                fmt.Println(v44, v50, v43)
            }
        case 977:
            for fi77 := 0; fi77 < 18; fi77++ {
                if (("" + "a") == "a") {
                    v43 = v43
                }
                _ = fi77
            }
        default:
            var v78 bool = ("foo bar" == ("a" + "foo bar"))
            _ = v78
        }
        _ = fi74
    }
    if false {
        var v79 uint32 = 42
        _ = v79
    }
    for fi80 := 0; fi80 < 14; fi80++ {
        var v81 float64 = (float64(3.14159) * (0.0 + v43))
        v82 := GenS4{F0: true, F1: "hello", F2: 1}
        var v83 uint32 = 4294967295
        _ = fi80
        _ = v81
        _ = v82
        _ = v83
    }
    switch "hello" {
    case "\n\t":
        v5 -= v5
    case "":
        var v84 float32 = (0.1 + v44)
        _ = v84
    case "a":
        v3 += v3
    default:
        v3 = v50
    }
    fmt.Println(v2, v6, v6)
    v1[("hello" + "")] = ("foo bar" + ("" + "foo bar"))
    fmt.Println(v43, v73)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v73
}
