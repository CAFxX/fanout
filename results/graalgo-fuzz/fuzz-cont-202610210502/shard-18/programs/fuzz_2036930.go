package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS5 struct { F0 uint32; F1 int8; F2 int }
type GenS6 struct { F0 int8; F1 uint32 }
type GenS7 struct { F0 int32; F1 int }
type GenS8 struct { F0 uint32; F1 int64; F2 int }
type GenS9 struct { F0 int32; F1 int64 }
type GenS10 struct { F0 int; F1 int8; F2 uint8 }
func genFn0(p0 int64, p1 uint32) string {
    var v6 int32 = 0
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 int32 = int32(fi7)
        _ = fi7
        _ = v8
    }
    if ((v6 >> ((4294967295) & 63)) >= (v6 ^ (v6 << ((369) & 63)))) {
        fmt.Println(p0, v6)
        switch "a" {
        case "\n\t":
            v9 := GenS1{F0: 4294967295, F1: "hello"}
            _ = v9
        case "hello":
            v10 := GenS1{F0: -333, F1: ""}
            _ = v10
        case "foo bar":
            fmt.Println(p1, p1, v6)
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return "\n\t"
}

func main() {
    var v1 int = -71
    var v2 int = -553
    var v3 bool = true
    var v4 float32 = -2.25
    var v5 uint32 = 179
    v11 := GenS1{F0: 4294967295, F1: "foo bar"}
    switch "\n\t" {
    case "a":
        v3 = (209 <= 42)
    case "foo bar":
        if ((!(v3 || (2147483648 != 2147483648))) && true) {
            switch (v1 ^ v1) {
            case 812:
                v12 := GenS1{F0: (v2 / ((1) | 1)), F1: "a"}
                _ = v12
            case -48:
                if (470 < 2) {
                    v13 := []int{(v2 + (v1 + v2)), (v2 % ((v2) | 1)), 128}
                    var v14 bool = (127 > 127)
                    _ = v13
                    _ = v14
                } else {
                    v2 = v2
                    v4 -= (v4 - 3.14159)
                    fmt.Println(v2)
                }
            case -182:
                v2 = v2
            }
            for fi15 := 0; fi15 < 4; fi15++ {
                if (((!(((fi15 * v2) == v1) && (v3 || v3))) && (!(v3 && (9223372036854775807 == 1)))) && true) {
                    var v16 bool = ((v4 != 1e10) && (!v3))
                    var v17 int64 = 9223372036854775807
                    _ = v16
                    _ = v17
                } else {
                    v18 := []int{}
                    v5 = 0
                    for fi19 := 0; fi19 < 2; fi19++ {
                        v18 = append(v18, fi15)
                        v20 := GenS1{F0: 1, F1: "a"}
                        if v3 {
                            v2 = v1
                            v21 := GenS1{F0: fi15, F1: "\n\t"}
                            v2 = fi19
                            _ = v21
                        } else {
                            v1 = 32767
                            if v3 {
                                if (true || ((v4 < v4) && (!(!(!v3))))) {
                                    v18 = append(v18, fi15)
                                    fmt.Println(fi19, v3, fi19)
                                    if (v4 == ((0.1 - v4) - v4)) {
                                        fmt.Println(v5, v4, fi19)
                                        if v3 {
                                            if (v4 <= 0.1) {
                                                fmt.Println(v5)
                                            } else {
                                                fmt.Println(v3, v5)
                                            }
                                            var v22 uint = 970
                                            if len(v18) > 0 {
                                                v18[0] = (-(fi19))
                                            }
                                            var v23 string = "hello"
                                            var v24 string = v23
                                            var v25 float64 = (1e10 + -2.25)
                                            _ = v22
                                            _ = v23
                                            _ = v24
                                            _ = v25
                                        }
                                        v5 = v5
                                    } else {
                                        v18 = append(v18, fi15)
                                    }
                                }
                                var v26 uint64 = uint64(v5)
                                var v27 int32 = -1
                                v28 := GenS1{F0: fi19, F1: ("foo bar" + ("\n\t" + "foo bar"))}
                                _ = v26
                                _ = v27
                                _ = v28
                            }
                        }
                        _ = fi19
                        _ = v20
                    }
                    _ = v18
                }
                v1 = v1
                v29 := GenS2{F0: 42}
                _ = fi15
                _ = v29
            }
            v30 := []int{}
            _ = v30
        } else {
            switch "hello" {
            case "hello":
                for fi31 := 0; fi31 < 8; fi31++ {
                    v32 := GenS3{F0: (v3 || (!(!((v3 && v3) || (0.1 == 1e10))))), F1: (1 < 348)}
                    v1 = 127
                    _ = fi31
                    _ = v32
                }
            case "foo bar":
                for fi33 := 0; fi33 < 15; fi33++ {
                    var v34 int16 = 1000
                    _ = fi33
                    _ = v34
                }
            }
        }
    case "\n\t":
        v2 -= 1
    default:
        switch v1 {
        case 2147483647:
            var v35 uint32 = 215
            _ = v35
        case 127:
            for fi36 := 0; fi36 < 19; fi36++ {
                var v37 int16 = 1
                _ = fi36
                _ = v37
            }
        case 32767:
            v38 := []int{127}
            _ = v38
        }
    }
    for fi39 := 0; fi39 < 18; fi39++ {
        v40 := []int{v1, 32767, 1}
        switch fi39 {
        case 1:
            var v41 uint8 = 65
            _ = v41
        default:
            v40 = append(v40, 0)
        }
        for fi42 := 0; fi42 < 4; fi42++ {
            v43 := GenS2{F0: 26}
            if v3 {
                v44 := GenS4{F0: v5, F1: 127, F2: 9223372036854775807}
                _ = v44
            }
            if len(v40) > 0 {
                v40[0] = (^(v2))
            }
            _ = fi42
            _ = v43
        }
        _ = fi39
        _ = v40
    }
    var v45 uint8 = 2
    if true {
        var v46 bool = false
        var v47 string = "a"
        var v48 uint = 434
        _ = v46
        _ = v47
        _ = v48
    }
    if v3 {
        for fi49 := 0; fi49 < 10; fi49++ {
            v50 := GenS3{F0: (!(float32(0.0) < (3.14159 / v4))), F1: true}
            _ = fi49
            _ = v50
        }
        var v51 int32 = -1
        _ = v51
    } else {
        v52 := []int{(-(v2)), v2}
        var v53 int = -437
        v1 -= v2
        _ = v52
        _ = v53
    }
    fmt.Println(v1, v45, v45)
    if v3 {
        if v3 {
            if (0 < 0) {
                if (true || v3) {
                    v45 = 37
                }
            }
        } else {
            if false {
                v3 = v3
                v54 := GenS3{F0: (213 > 221), F1: (156 <= 0)}
                _ = v54
            } else {
                if (!(!(v3 && (v3 || (false && (-516 == -1)))))) {
                    if (42 >= 42) {
                        switch "a" {
                        case "a":
                            var v55 float32 = (1.5 - v4)
                            _ = v55
                        }
                    } else {
                        v56 := []int{}
                        _ = v56
                    }
                    v45 += 128
                }
                switch v2 {
                case -268:
                    v57 := GenS1{F0: (^(v1)), F1: (("a" + "") + "\n\t")}
                    _ = v57
                case -557:
                    for fi58 := 0; fi58 < 17; fi58++ {
                        switch v1 {
                        case 0:
                            if (("" != "hello") || (!false)) {
                                for fi59 := 0; fi59 < 18; fi59++ {
                                    v60 := []int{(-(v2))}
                                    v1 -= fi59
                                    _ = fi59
                                    _ = v60
                                }
                            }
                        case 128:
                            var v61 float64 = 0.1
                            _ = v61
                        case -2147483648:
                            var v62 float64 = 0.0
                            _ = v62
                        default:
                            var v63 uint32 = (v5 & (^(v5)))
                            _ = v63
                        }
                        _ = fi58
                    }
                }
            }
            v64 := []int{}
            for fi65 := 0; fi65 < 5; fi65++ {
                v66 := GenS5{F0: (v5 + (v5 - v5)), F1: 19, F2: fi65}
                v67 := GenS6{F0: 0, F1: v5}
                v68 := GenS6{F0: -32, F1: v5}
                _ = fi65
                _ = v66
                _ = v67
                _ = v68
            }
            _ = v64
        }
        for fi69 := 0; fi69 < 14; fi69++ {
            switch ("" + "foo bar") {
            case "":
                v2 = (fi69 - -928)
            case "foo bar":
                v70 := GenS4{F0: v5, F1: v45, F2: -9223372036854775808}
                _ = v70
            case "\n\t":
                var v71 int16 = -463
                _ = v71
            default:
                v72 := []int{1}
                _ = v72
            }
            switch "" {
            case "foo bar":
                if (!v3) {
                    for fi73 := 0; fi73 < 9; fi73++ {
                        var v74 string = "foo bar"
                        fmt.Println(v1, v3)
                        var v75 uint16 = 2
                        _ = fi73
                        _ = v74
                        _ = v75
                    }
                    v76 := GenS7{F0: 1, F1: 4294967295}
                    _ = v76
                }
            case "":
                v45 -= 40
            }
            fmt.Println(fi69)
            _ = fi69
        }
        v4 = (v4 * v4)
    } else {
        switch ("foo bar" + ((("a" + "\n\t") + "hello") + ("foo bar" + "foo bar"))) {
        case "a":
            v77 := GenS8{F0: 2, F1: -698, F2: 675}
            _ = v77
        case "hello":
            v5 = 0
        case "foo bar":
            v5 = v5
        }
    }
    if (uint8(v5) > (^(v45))) {
        var v78 uint8 = v45
        _ = v78
    } else {
        var v79 int16 = 1000
        var v80 float32 = v4
        _ = v79
        _ = v80
    }
    v81 := GenS9{F0: -2147483648, F1: 0}
    v82 := GenS10{F0: (v2 & v1), F1: 127, F2: v45}
    v1 = (v2 << ((0) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v81
    _ = v82
}
