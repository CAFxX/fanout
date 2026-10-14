package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 int32; F1 uint; F2 int64 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 int8; F1 bool; F2 bool }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint16; F1 string }
type GenS9 struct { F0 bool; F1 uint }
type GenS10 struct { F0 int32; F1 int32 }
func genFn0(p0 int32) uint8 {
    if false {
        switch "\n\t" {
        case "a":
            p0 -= p0
        default:
            p0 = (p0 >> ((65535) & 63))
        }
        var v5 int32 = (p0 ^ (^(p0)))
        _ = v5
    }
    p0 = (p0 / (((^(p0))) | 1))
    v6 := []int{-1, -9223372036854775808, 42}
    v7 := GenS1{F0: 42}
    _ = p0
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int8, p1 int16) uint16 {
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 uint16 = 1
        var v10 int32 = -424
        for fi11 := 0; fi11 < 5; fi11++ {
            v9 += (v9 * (^(v9)))
            v12 := []int{(fi11 >> ((1) & 63)), (-(fi11))}
            if false {
                switch (fi8 ^ fi11) {
                case -133:
                    var v13 float64 = 3.14159
                    _ = v13
                case 340:
                    p1 = p1
                    var v14 float32 = (3.14159 + float32(-2.25))
                    _ = v14
                case 2147483647:
                    switch "" {
                    case "foo bar":
                        var v15 int32 = 42
                        _ = v15
                    }
                default:
                    if (-1 > 1) {
                        fmt.Println(fi11)
                    }
                }
                var v16 uint32 = 924
                _ = v16
            }
            _ = fi11
            _ = v12
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    if false {
        v17 := GenS2{F0: 42, F1: 642}
        fmt.Println(p0, p1)
        _ = v17
    }
    p0 = (p0 - (-(p0)))
    v18 := []int{int(p0), 32767, 127, -357}
    p0 = 1
    _ = p0
    _ = p1
    _ = v18
    return 0
}

func main() {
    var v1 int8 = 42
    var v2 bool = true
    var v3 int64 = 654
    var v4 string = "foo bar"
    for fi19 := 0; fi19 < 10; fi19++ {
        if (497 == v3) {
            v20 := []int{(fi19 ^ fi19)}
            v1 -= v1
            _ = v20
        } else {
            v21 := []int{fi19}
            _ = v21
        }
        _ = fi19
    }
    if true {
        v22 := GenS3{F0: 95, F1: (v1 - v1)}
        var v23 uint32 = 196
        fmt.Println(v23, v3)
        _ = v22
        _ = v23
    }
    var v24 int8 = v1
    switch ((((v4 + v4) + (("" + (v4 + v4)) + "a")) + v4) + (v4 + "a")) {
    case "foo bar":
        v25 := GenS2{F0: v3, F1: genFn1(v24, -32768)}
        _ = v25
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        switch (fi26 / (((fi26 | 195)) | 1)) {
        case -318:
            v27 := GenS4{F0: -1, F1: 255, F2: (v3 + (^(v3)))}
            _ = v27
        case 9223372036854775807:
            for fi28 := 0; fi28 < 6; fi28++ {
                v29 := []int{145, fi28, fi28, fi26}
                fmt.Println(v1)
                _ = fi28
                _ = v29
            }
        case 0:
            switch (v4 + v4) {
            case "\n\t":
                for fi30 := 0; fi30 < 8; fi30++ {
                    var v31 string = v4
                    v32 := GenS1{F0: 1}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            case "":
                for fi33 := 0; fi33 < 13; fi33++ {
                    fmt.Println(v3, fi33)
                    v24 = (v24 << ((2) & 63))
                    if true {
                        v2 = v2
                        switch ("\n\t" + v4) {
                        case "a":
                            v34 := GenS5{F0: 802, F1: v3}
                            _ = v34
                        }
                    } else {
                        var v35 int32 = 22
                        var v36 bool = false
                        _ = v35
                        _ = v36
                    }
                    fmt.Println(fi33)
                    v37 := []int{(^(fi33))}
                    _ = fi33
                    _ = v37
                }
            case "hello":
                var v38 int16 = 1000
                _ = v38
            }
        }
        _ = fi26
    }
    var v39 int16 = -427
    v4 = v4
    v39 += v39
    switch "hello" {
    case "a":
        switch 128 {
        case 128:
            v40 := []int{127, 612, int(v3)}
            _ = v40
        case 127:
            v41 := []int{372, -700, -491}
            _ = v41
        }
    case "":
        v42 := GenS6{F0: v1, F1: (float64(-2.25) <= ((-2.25 + 0.0) / (((1e10 * (1.5 / 1e10)) / -2.25) * (3.14159 / 3.14159)))), F2: (!true)}
        _ = v42
    default:
        if true {
            var v43 uint16 = 332
            if v2 {
                switch -409 {
                case -677:
                    fmt.Println(v43)
                default:
                    v44 := GenS1{F0: v3}
                    _ = v44
                }
                fmt.Println(v2)
            } else {
                fmt.Println(v39, v24, v1)
                var v45 int16 = v39
                var v46 uint64 = 2
                _ = v45
                _ = v46
            }
            v47 := []int{}
            _ = v43
            _ = v47
        } else {
            if v2 {
                v4 = "\n\t"
                switch 0 {
                case -943:
                    v48 := []int{}
                    _ = v48
                case 32768:
                    for fi49 := 0; fi49 < 2; fi49++ {
                        var v50 int = (fi49 & fi49)
                        var v51 bool = (255 < 0)
                        var v52 float64 = float64(3.14159)
                        _ = fi49
                        _ = v50
                        _ = v51
                        _ = v52
                    }
                case 42:
                    var v53 string = ("foo bar" + "hello")
                    _ = v53
                default:
                    v4 = ("a" + v4)
                }
            } else {
                if true {
                    v3 -= -1
                    v54 := []int{127}
                    _ = v54
                }
                if false {
                    v4 = v4
                    for fi55 := 0; fi55 < 5; fi55++ {
                        v24 = (v1 << ((305) & 63))
                        _ = fi55
                    }
                    v56 := GenS5{F0: genFn1(v24, -1), F1: (v3 << ((18446744073709551615) & 63))}
                    _ = v56
                } else {
                    if false {
                        if true {
                            switch 4294967295 {
                            case -1:
                                v4 = ("foo bar" + "\n\t")
                            case 2147483647:
                                v57 := GenS6{F0: (v24 ^ v24), F1: ((v39 == -961) || v2), F2: (-1 < 4294967295)}
                                _ = v57
                            case -2147483648:
                                v58 := GenS7{F0: 42}
                                _ = v58
                            default:
                                switch v4 {
                                case "\n\t":
                                    if (true && true) {
                                        v59 := GenS5{F0: 2, F1: -171}
                                        if false {
                                            for fi60 := 0; fi60 < 18; fi60++ {
                                                v39 = (v39 >> ((4294967295) & 63))
                                                var v61 uint8 = 0
                                                v62 := GenS1{F0: (v3 / (((-(v3))) | 1))}
                                                _ = fi60
                                                _ = v61
                                                _ = v62
                                            }
                                            v2 = v2
                                        }
                                        _ = v59
                                    }
                                case "":
                                    var v63 float64 = 1.5
                                    _ = v63
                                default:
                                    var v64 bool = (4294967295 == 0)
                                    _ = v64
                                }
                            }
                            if (v3 > (v3 + -768)) {
                                v65 := GenS2{F0: (^(v3)), F1: 32767}
                                _ = v65
                            }
                            v66 := GenS8{F0: 1, F1: "a"}
                            _ = v66
                        } else {
                            v67 := []int{174}
                            _ = v67
                        }
                        switch 1 {
                        case -203:
                            var v68 int = -82
                            _ = v68
                        default:
                            var v69 uint = 65535
                            _ = v69
                        }
                        v70 := []int{9223372036854775807, -524, 32767, -806}
                        _ = v70
                    }
                    v1 += (v24 ^ (-(v24)))
                }
                v71 := GenS9{F0: ((v39 - (v39 ^ v39)) <= -317), F1: 42}
                _ = v71
            }
            fmt.Println(v24, v39, v24)
            var v72 string = (("hello" + (v4 + (v4 + "foo bar"))) + v4)
            _ = v72
        }
    }
    for fi73 := 0; fi73 < 7; fi73++ {
        switch "hello" {
        case "foo bar":
            v74 := []int{int(v39), 42, (fi73 >> ((42) & 63)), -214}
            _ = v74
        }
        var v75 float64 = (3.14159 - 1.5)
        v76 := GenS1{F0: (v3 >> ((2) & 63))}
        _ = fi73
        _ = v75
        _ = v76
    }
    switch 1 {
    case -635:
        for fi77 := 0; fi77 < 5; fi77++ {
            v78 := []int{fi77, fi77, (fi77 >> ((4294967295) & 63))}
            _ = fi77
            _ = v78
        }
    case -1:
        fmt.Println(v4, v1, v2)
    default:
        fmt.Println(v2, v1, v4)
    }
    if (!true) {
        v79 := GenS10{F0: 42, F1: 0}
        var v80 uint32 = 4294967295
        for fi81 := 0; fi81 < 10; fi81++ {
            if (v39 > v39) {
                v82 := []int{}
                for fi83 := 0; fi83 < 8; fi83++ {
                    v3 = 0
                    var v84 uint = 378
                    var v85 uint32 = v80
                    _ = fi83
                    _ = v84
                    _ = v85
                }
                _ = v82
            }
            _ = fi81
        }
        _ = v79
        _ = v80
    }
    v1 += v24
    for fi86 := 0; fi86 < 11; fi86++ {
        if v2 {
            var v87 bool = (!(v2 || (!(660 > -564))))
            _ = v87
        }
        _ = fi86
    }
    if true {
        v88 := GenS6{F0: v1, F1: v2, F2: v2}
        v89 := []int{4294967295, -2147483648, 128, -2147483648}
        _ = v88
        _ = v89
    } else {
        var v90 int32 = 131
        _ = v90
    }
    for fi91 := 0; fi91 < 3; fi91++ {
        v92 := GenS1{F0: (v3 ^ (v3 / (((-(v3))) | 1)))}
        _ = fi91
        _ = v92
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v39
    _ = v4
}
