package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int64; F2 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 bool }
type GenS6 struct { F0 uint64; F1 int8 }
type GenS7 struct { F0 int64; F1 uint8; F2 string }
type GenS8 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int64; F1 bool }
func genFn0() int {
    if (!false) {
        var v7 int = 32768
        switch v7 {
        case -1:
            v7 += 201
        default:
            for fi8 := 0; fi8 < 3; fi8++ {
                v9 := GenS1{F0: "\n\t"}
                for fi10 := 0; fi10 < 20; fi10++ {
                    switch fi10 {
                    case 2147483647:
                        fmt.Println(fi8, v7)
                    case -517:
                        fmt.Println(fi10)
                    case 422:
                        fmt.Println(v7)
                    }
                    _ = fi10
                }
                _ = fi8
                _ = v9
            }
        }
        fmt.Println(v7)
        _ = v7
    } else {
        var v11 uint8 = 47
        _ = v11
    }
    v12 := GenS1{F0: ("hello" + ("\n\t" + ("foo bar" + "a")))}
    var v13 uint = 943
    v13 += v13
    fmt.Println(v13)
    _ = v12
    _ = v13
    return 509
}
func genFn1(p0 int, p1 string) uint8 {
    p0 -= p0
    v14 := []int{genFn0()}
    switch "foo bar" {
    case "\n\t":
        fmt.Println(p1, p0)
    default:
        var v15 int16 = 32767
        _ = v15
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        p1 = "foo bar"
        _ = fi16
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = v14
    return 215
}

func main() {
    var v1 float64 = 1.5
    var v2 int = 127
    v3 := []string{(("a" + (("" + "foo bar") + (("" + ("" + "foo bar")) + "\n\t"))) + ("\n\t" + ("" + "\n\t")))}
    var v4 int16 = 0
    var v5 int32 = 42
    var v6 uint = 284
    var v17 int64 = 9223372036854775807
    var v18 uint64 = 349
    if ((float32(1.5) - 1e10) <= 0.0) {
        if ((true && true) || (!true)) {
            v19 := GenS1{F0: "foo bar"}
            fmt.Println(v4)
            _ = v19
        }
        for fi20 := 0; fi20 < 19; fi20++ {
            v17 = 780
            _ = fi20
        }
        var v21 int8 = 127
        _ = v21
    } else {
        v22 := GenS1{F0: ("hello" + "")}
        _ = v22
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        v17 = v17
        if ("\n\t" <= "\n\t") {
            fmt.Println(v5, v1, v18)
            var v24 float64 = (((0.1 * v1) / v1) + (v1 + -2.25))
            switch "foo bar" {
            case "":
                switch "\n\t" {
                case "":
                    for fi25 := 0; fi25 < 11; fi25++ {
                        fmt.Println(fi25, v2)
                        _ = fi25
                    }
                case "a":
                    v17 -= (v17 - v17)
                default:
                    var v26 int8 = 100
                    _ = v26
                }
            case "\n\t":
                var v27 int = v2
                _ = v27
            }
            _ = v24
        }
        _ = fi23
    }
    var v28 uint32 = 0
    switch "hello" {
    case "hello":
        switch ((("foo bar" + "") + "\n\t") + "\n\t") {
        case "a":
            fmt.Println(v6)
        case "\n\t":
            var v29 uint = (v6 ^ v6)
            _ = v29
        }
    case "":
        var v30 int16 = 32767
        _ = v30
    default:
        if true {
            v31 := GenS2{F0: v17, F1: v17, F2: (^(v4))}
            fmt.Println(v1, v1)
            for fi32 := 0; fi32 < 6; fi32++ {
                _ = len(v3)
                _ = len(v3)
                _ = fi32
            }
            _ = v31
        }
    }
    v33 := GenS3{F0: 0}
    if false {
        v34 := GenS4{F0: v5}
        if true {
            switch v2 {
            case 42:
                fmt.Println(v5, v5)
            case 0:
                var v35 uint32 = (v28 - 1)
                _ = v35
            case 32768:
                switch (v2 ^ -415) {
                case 32767:
                    for fi36 := 0; fi36 < 15; fi36++ {
                        if false {
                            if (!false) {
                                fmt.Println(v2)
                                var v37 uint = (v6 * (v6 | 4294967295))
                                _ = v37
                            } else {
                                v38 := GenS5{F0: 2, F1: false}
                                fmt.Println(v28)
                                _ = v38
                            }
                            v28 += (v28 + 0)
                        }
                        if len(v3) > 0 {
                            v3[0] = ("hello" + "\n\t")
                        }
                        var v39 float64 = float64(-2.25)
                        _ = fi36
                        _ = v39
                    }
                case 852:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case -419:
                    var v40 float32 = ((-2.25 * 1e10) / 1e10)
                    _ = v40
                default:
                    var v41 bool = (v1 > v1)
                    _ = v41
                }
            }
            if ((v5 % ((-685) | 1)) == (v5 & v5)) {
                v42 := GenS5{F0: v6, F1: (379 <= 39)}
                v5 -= v5
                var v43 uint64 = (v18 & (v18 % (((v18 - (v18 + 0))) | 1)))
                _ = v42
                _ = v43
            }
            if false {
                for fi44 := 0; fi44 < 7; fi44++ {
                    for fi45 := 0; fi45 < 2; fi45++ {
                        var v46 bool = false
                        _ = fi45
                        _ = v46
                    }
                    if false {
                        switch "\n\t" {
                        case "":
                            if true {
                                for fi47 := 0; fi47 < 7; fi47++ {
                                    v4 = v4
                                    if true {
                                        v17 -= (v17 & v17)
                                    } else {
                                        switch ("foo bar" + ("hello" + ("hello" + "\n\t"))) {
                                        case "\n\t":
                                            v48 := GenS2{F0: (^(v17)), F1: -9223372036854775808, F2: 0}
                                            _ = v48
                                        case "a":
                                            var v49 string = ""
                                            _ = v49
                                        case "hello":
                                            v50 := GenS4{F0: -943}
                                            _ = v50
                                        default:
                                            var v51 uint32 = v28
                                            _ = v51
                                        }
                                        v6 = (v6 & (v6 << ((v6) & 63)))
                                    }
                                    var v52 uint8 = 255
                                    _ = fi47
                                    _ = v52
                                }
                                switch "hello" {
                                case "foo bar":
                                    if len(v3) > 0 {
                                        v3[0] = ("\n\t" + "hello")
                                    }
                                default:
                                    v53 := GenS6{F0: (v18 ^ 2), F1: 100}
                                    _ = v53
                                }
                                switch "a" {
                                case "":
                                    fmt.Println(v1, v6, v4)
                                }
                            }
                        case "a":
                            switch ("a" + "foo bar") {
                            case "foo bar":
                                fmt.Println(v1)
                            default:
                                fmt.Println(v1)
                            }
                        case "hello":
                            v2 = 42
                        }
                        for fi54 := 0; fi54 < 6; fi54++ {
                            if false {
                                if false {
                                    var v55 int8 = -85
                                    _ = v55
                                }
                            } else {
                                v6 = v6
                                v56 := GenS1{F0: "a"}
                                v57 := GenS1{F0: "hello"}
                                v58 := GenS7{F0: v17, F1: 114, F2: "\n\t"}
                                _ = v56
                                _ = v57
                                _ = v58
                            }
                            _ = fi54
                        }
                    } else {
                        switch ("hello" + "\n\t") {
                        case "a":
                            var v59 bool = false
                            _ = v59
                        case "foo bar":
                            _ = len(v3)
                        default:
                            switch "\n\t" {
                            case "foo bar":
                                v28 = (v28 << ((65535) & 63))
                            }
                        }
                        for fi60 := 0; fi60 < 5; fi60++ {
                            var v61 int = (v2 / ((v2) | 1))
                            if true {
                                fmt.Println(v5, v4, v17)
                                v1 += v1
                                v62 := GenS4{F0: (^(v5))}
                                _ = v62
                            }
                            _ = fi60
                            _ = v61
                        }
                        v63 := GenS8{F0: 840, F1: v18, F2: 1}
                        _ = v63
                    }
                    _ = fi44
                }
            }
        }
        _ = v34
    }
    v64 := GenS9{F0: v4}
    if true {
        var v65 uint16 = 555
        _ = len(v3)
        v66 := GenS10{F0: 526, F1: false}
        _ = v65
        _ = v66
    } else {
        for fi67 := 0; fi67 < 3; fi67++ {
            var v68 uint = (v6 / (((v6 & v6)) | 1))
            var v69 bool = (871 != 2)
            var v70 uint16 = 342
            _ = fi67
            _ = v68
            _ = v69
            _ = v70
        }
        v2 -= v2
    }
    var v71 string = "a"
    var v72 int8 = 0
    v72 = v72
    var v73 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v71
    _ = v72
    _ = v73
}
