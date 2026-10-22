package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 int32; F1 int64; F2 string }
func genFn0(p0 uint, p1 int32) int32 {
    v6 := []int{1, 699}
    v7 := GenS1{F0: p0}
    var v8 string = "a"
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return (p1 - -2147483648)
}
func genFn1(p0 int8, p1 uint16) uint {
    var v9 uint = 433
    if true {
        v10 := []int{32768}
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 int32 = 2147483647
            if true {
                for fi13 := 0; fi13 < 8; fi13++ {
                    if len(v10) > 0 {
                        v10[0] = fi13
                    }
                    p1 += p1
                    _ = fi13
                }
                for fi14 := 0; fi14 < 18; fi14++ {
                    switch 127 {
                    case 32767:
                        v9 -= 127
                    case 9223372036854775807:
                        if false {
                            if (false && (true && (false || ((388 < 975) && false)))) {
                                if true {
                                    var v15 uint64 = 1
                                    v16 := GenS2{F0: 1000}
                                    switch fi14 {
                                    case -9223372036854775808:
                                        v9 -= 0
                                    case 9223372036854775807:
                                        v10 = append(v10, (fi14 & 4294967295))
                                        v17 := GenS2{F0: 1000}
                                        p1 += uint16(v15)
                                        _ = v17
                                    case -2147483648:
                                        v12 += (v12 / ((genFn0((v9 * (v9 ^ (v9 | uint(v15)))), (^(v12)))) | 1))
                                    }
                                    _ = v15
                                    _ = v16
                                }
                            }
                            fmt.Println(v12, fi11, p1)
                        }
                    }
                    _ = fi14
                }
                var v18 int = (-(fi11))
                _ = v18
            }
            for fi19 := 0; fi19 < 11; fi19++ {
                var v20 int64 = 9223372036854775807
                _ = fi19
                _ = v20
            }
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    var v21 int64 = -9223372036854775808
    v22 := []int{4294967295, 4294967295, 42}
    _ = p0
    _ = p1
    _ = v21
    _ = v22
    _ = v9
    return v9
}

func main() {
    var v1 uint64 = 42
    var v2 uint16 = 1
    var v3 uint8 = 1
    var v4 uint32 = 0
    v5 := make(map[string]bool)
    v1 -= uint64(v4)
    if (!(true && false)) {
        v5["hello"] = true
    } else {
        v4 -= 1
        v4 = (v4 * (v4 - (v4 / (((v4 << ((42) & 63))) | 1))))
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        _ = len(v5)
        _ = fi23
    }
    var v24 int = -2147483648
    switch "foo bar" {
    case "\n\t":
        v1 = v1
    case "foo bar":
        v25 := []int{}
        _ = v25
    default:
        v26 := []int{v24, int(v2)}
        _ = v26
    }
    for fi27 := 0; fi27 < 18; fi27++ {
        v4 = v4
        _ = fi27
    }
    delete(v5, ("" + "foo bar"))
    if false {
        for fi28 := 0; fi28 < 13; fi28++ {
            for fi29 := 0; fi29 < 3; fi29++ {
                if ((fi28 % (((-(v24))) | 1)) == (^(fi29))) {
                    var v30 int8 = 20
                    _ = v30
                } else {
                    fmt.Println(v1, v1, v4)
                    if true {
                        var v31 uint64 = (^(v1))
                        switch ((("" + "") + ("a" + "\n\t")) + "") {
                        case "a":
                            var v32 bool = (!true)
                            _ = v32
                        default:
                            v33 := GenS1{F0: 42}
                            _ = v33
                        }
                        _ = v31
                    } else {
                        switch fi29 {
                        case 32768:
                            var v34 bool = true
                            switch fi29 {
                            case 128:
                                if ((v3 % ((0) | 1)) <= (v3 & 2)) {
                                    var v35 uint8 = 128
                                    _ = v35
                                }
                            case -42:
                                var v36 int = v24
                                _ = v36
                            }
                            if true {
                                fmt.Println(v3, fi29)
                            } else {
                                fmt.Println(v24)
                            }
                            _ = v34
                        default:
                            if false {
                                switch ((("foo bar" + (("foo bar" + ("\n\t" + "")) + "")) + "") + "a") {
                                case "a":
                                    var v37 int8 = 42
                                    _ = v37
                                case "":
                                    v2 -= (v2 / ((v2) | 1))
                                case "\n\t":
                                    v5[("hello" + "a")] = (2 != 0)
                                default:
                                    switch (^(fi29)) {
                                    case 0:
                                        _ = len(v5)
                                    case -996:
                                        var v38 int64 = 0
                                        _ = v38
                                    }
                                }
                            } else {
                                v3 = v3
                            }
                        }
                        v39 := GenS3{F0: 0, F1: (v4 - 1)}
                        switch "hello" {
                        case "foo bar":
                            switch -886 {
                            case 150:
                                v40 := GenS4{F0: -795, F1: 9223372036854775807, F2: "a"}
                                _ = v40
                            case -710:
                                if false {
                                    var v41 int32 = -537
                                    v42 := GenS1{F0: 255}
                                    v43 := []int{}
                                    _ = v41
                                    _ = v42
                                    _ = v43
                                } else {
                                    v3 = v3
                                    v44 := []int{}
                                    if (127 > 17) {
                                        fmt.Println(v24, v24, v1)
                                        if false {
                                            v5[("a" + "foo bar")] = true
                                        } else {
                                            var v45 uint = genFn1(42, v2)
                                            _ = v45
                                        }
                                        v1 = v1
                                    } else {
                                        if true {
                                            fmt.Println(v3)
                                            v46 := GenS4{F0: 0, F1: -87, F2: "hello"}
                                            _ = v46
                                        } else {
                                            var v47 uint64 = (v1 | v1)
                                            _ = v47
                                        }
                                    }
                                    _ = v44
                                }
                            }
                        case "hello":
                            switch (fi28 ^ v24) {
                            case 42:
                                if false {
                                    v2 = 1
                                    v4 -= v4
                                } else {
                                    v5[""] = (-105 <= 100)
                                    var v48 int32 = 2147483647
                                    if true {
                                        v2 = 0
                                        v2 = v2
                                        var v49 int64 = int64(v2)
                                        _ = v49
                                    } else {
                                        if false {
                                            var v50 float32 = 3.14159
                                            if ((v4 % ((v4) | 1)) >= v4) {
                                                var v51 int32 = (v48 | genFn0(1, 1))
                                                var v52 string = ("foo bar" + (("foo bar" + "\n\t") + "\n\t"))
                                                if false {
                                                    if (18446744073709551615 > 65535) {
                                                        var v53 uint64 = v1
                                                        var v54 uint16 = v2
                                                        _ = v53
                                                        _ = v54
                                                    }
                                                } else {
                                                    if (!true) {
                                                        var v55 uint64 = v1
                                                        if (!(v55 > v55)) {
                                                            var v56 int = 903
                                                            v57 := GenS2{F0: int16(v1)}
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                        _ = v55
                                                    }
                                                    switch "\n\t" {
                                                    case "":
                                                        v58 := GenS1{F0: 255}
                                                        _ = v58
                                                    }
                                                }
                                                _ = v51
                                                _ = v52
                                            }
                                            fmt.Println(v2, fi29, v1)
                                            delete(v5, ("\n\t" + "foo bar"))
                                            var v59 uint32 = (^(v4))
                                            var v60 uint = 784
                                            _ = v50
                                            _ = v59
                                            _ = v60
                                        } else {
                                            if true {
                                                switch ("a" + "") {
                                                case "":
                                                    v61 := []int{v24, 128}
                                                    _ = v61
                                                case "a":
                                                    v2 -= v2
                                                case "hello":
                                                    fmt.Println(v3, v1, v24)
                                                default:
                                                    var v62 bool = true
                                                    _ = v62
                                                }
                                            } else {
                                                switch 127 {
                                                case -2147483648:
                                                    v4 = v4
                                                case 32768:
                                                    var v63 uint8 = v3
                                                    _ = v63
                                                }
                                            }
                                            v64 := GenS3{F0: 981, F1: 4294967295}
                                            _ = v64
                                        }
                                    }
                                    _ = v48
                                }
                            }
                        }
                        _ = v39
                    }
                    var v65 uint = 891
                    _ = v65
                }
                if ((0 <= 42) || ((-(fi28)) >= (fi29 - fi29))) {
                    fmt.Println(fi29, v24)
                } else {
                    var v66 float64 = 1.5
                    v5["foo bar"] = (!(-1 >= 1000))
                    var v67 string = "foo bar"
                    var v68 uint32 = v4
                    _ = v66
                    _ = v67
                    _ = v68
                }
                _ = fi29
            }
            for fi69 := 0; fi69 < 16; fi69++ {
                fmt.Println(fi69, v24)
                var v70 bool = false
                var v71 uint16 = 255
                _ = fi69
                _ = v70
                _ = v71
            }
            _ = fi28
        }
    }
    switch ("foo bar" + "foo bar") {
    case "":
        var v72 int = 32767
        _ = v72
    case "a":
        if true {
            var v73 int = v24
            _ = v73
        }
    default:
        var v74 float64 = -2.25
        _ = v74
    }
    for fi75 := 0; fi75 < 1; fi75++ {
        var v76 int8 = 127
        _ = fi75
        _ = v76
    }
    switch v24 {
    case 127:
        var v77 uint32 = 2147483648
        _ = v77
    case 0:
        v4 = 1
    case -2147483648:
        var v78 uint32 = v4
        _ = v78
    default:
        v5[""] = (2 <= v4)
    }
    if (!false) {
        var v79 int32 = 42
        var v80 float32 = -2.25
        _ = v79
        _ = v80
    } else {
        var v81 bool = false
        delete(v5, "foo bar")
        _ = v81
    }
    fmt.Println(v24, v1, v24)
    for fi82 := 0; fi82 < 7; fi82++ {
        if true {
            v5[""] = ((^(fi82)) != (v24 ^ v24))
            if (1 <= -886) {
                v1 = (v1 + v1)
            }
        }
        v83 := []int{}
        var v84 string = "a"
        _ = fi82
        _ = v83
        _ = v84
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
