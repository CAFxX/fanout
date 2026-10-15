package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int; F1 int64; F2 uint8 }
type GenS7 struct { F0 uint32; F1 uint16; F2 bool }
type GenS8 struct { F0 uint8; F1 int32; F2 int32 }

func main() {
    var v1 uint16 = 2
    var v2 uint16 = v1
    var v3 int = 32767
    var v4 string = "\n\t"
    v5 := []string{}
    for fi6 := 0; fi6 < 12; fi6++ {
        v2 -= (v2 ^ v1)
        for fi7 := 0; fi7 < 12; fi7++ {
            v1 = (v1 | v1)
            v1 = v1
            _ = fi7
        }
        _ = fi6
    }
    if true {
        if (413 < 0) {
            var v8 uint16 = (^(v2))
            _ = v8
        } else {
            v9 := GenS1{F0: 0, F1: (((v4 + (v4 + "")) + (v4 + "")) + (v4 + v4)), F2: 4294967295}
            fmt.Println(v3)
            v2 += (v2 | (v1 & (v1 / (((v1 + v2)) | 1))))
            _ = v9
        }
        fmt.Println(v3, v2, v1)
        fmt.Println(v4, v2, v3)
    } else {
        fmt.Println(v3, v4)
        for fi10 := 0; fi10 < 11; fi10++ {
            for fi11 := 0; fi11 < 15; fi11++ {
                var v12 uint64 = 42
                v13 := GenS1{F0: 128, F1: "\n\t", F2: 2147483647}
                if false {
                    if (v1 <= v1) {
                        if false {
                            switch v4 {
                            case "":
                                var v14 uint8 = 1
                                _ = v14
                            default:
                                v15 := GenS1{F0: 42, F1: (v4 + (("\n\t" + v4) + (v4 + "foo bar"))), F2: 0}
                                _ = v15
                            }
                            fmt.Println(v1)
                        } else {
                            v4 = v4
                        }
                        fmt.Println(v2, v3, v4)
                    } else {
                        v16 := GenS2{F0: v3}
                        var v17 int16 = 32767
                        _ = v16
                        _ = v17
                    }
                    if ((2 > (v12 | (v12 - (v12 << ((0) & 63))))) || false) {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v4 = "a"
                        v18 := GenS1{F0: 42, F1: v4, F2: 2}
                        _ = v18
                    }
                    if false {
                        fmt.Println(v4, v12, v3)
                        fmt.Println(v4, v12)
                        v2 = 32767
                        v2 -= v1
                        v12 = v12
                    }
                }
                _ = fi11
                _ = v12
                _ = v13
            }
            v19 := GenS3{F0: 0}
            if true {
                for fi20 := 0; fi20 < 9; fi20++ {
                    if len(v5) > 0 {
                        v5[0] = v4
                    }
                    if ((true && (!false)) || false) {
                        v4 = v4
                    } else {
                        v21 := GenS4{F0: (559 >= v2), F1: 65535}
                        _ = v21
                    }
                    _ = fi20
                }
                v22 := GenS3{F0: 1}
                _ = v22
            }
            _ = fi10
            _ = v19
        }
        v23 := GenS4{F0: false, F1: (v1 | (v1 & v1))}
        _ = v23
    }
    var v24 float32 = -2.25
    v4 = v4
    fmt.Println(v2, v24, v2)
    v25 := GenS3{F0: 42}
    for fi26 := 0; fi26 < 17; fi26++ {
        if (v4 > v4) {
            for fi27 := 0; fi27 < 4; fi27++ {
                fmt.Println(fi26, v2, v4)
                _ = fi27
            }
            var v28 float64 = 1.5
            switch (fi26 >> ((42) & 63)) {
            case -1:
                for fi29 := 0; fi29 < 19; fi29++ {
                    v1 = v1
                    if false {
                        v30 := GenS4{F0: false, F1: (^(v1))}
                        switch "" {
                        case "foo bar":
                            v5 = append(v5, ("foo bar" + "a"))
                        }
                        if len(v5) > 0 {
                            v5[0] = "hello"
                        }
                        _ = v30
                    } else {
                        var v31 uint32 = 65535
                        v31 -= (v31 % ((v31) | 1))
                        _ = v31
                    }
                    _ = fi29
                }
            }
            _ = v28
        } else {
            for fi32 := 0; fi32 < 12; fi32++ {
                if len(v5) > 0 {
                    v5[0] = v4
                }
                _ = fi32
            }
            fmt.Println(fi26, v4)
        }
        var v33 int32 = 42
        fmt.Println(v4, v33)
        _ = fi26
        _ = v33
    }
    var v34 float32 = v24
    for fi35 := 0; fi35 < 9; fi35++ {
        fmt.Println(v1, v3, v3)
        _ = fi35
    }
    if (!false) {
        var v36 string = v4
        _ = v36
    } else {
        if true {
            fmt.Println(v3, v3, v4)
            _ = len(v5)
        }
        var v37 uint16 = v1
        _ = v37
    }
    v38 := GenS3{F0: 883}
    for fi39 := 0; fi39 < 9; fi39++ {
        var v40 float64 = 1e10
        v41 := GenS2{F0: (fi39 & (fi39 % (((fi39 / ((v3) | 1))) | 1)))}
        _ = fi39
        _ = v40
        _ = v41
    }
    var v42 int16 = 42
    v4 = (v4 + ((v4 + v4) + (v4 + (v4 + ("hello" + "hello")))))
    for fi43 := 0; fi43 < 1; fi43++ {
        if (767 < 4294967295) {
            v34 = 1e10
        }
        v34 += v24
        _ = fi43
    }
    fmt.Println(v3, v1)
    switch "a" {
    case "":
        var v44 bool = ((v42 | v42) >= v42)
        _ = v44
    case "hello":
        for fi45 := 0; fi45 < 11; fi45++ {
            v1 = v1
            if false {
                for fi46 := 0; fi46 < 20; fi46++ {
                    v47 := GenS5{F0: 0}
                    if false {
                        v1 -= v2
                        v48 := GenS6{F0: -853, F1: 0, F2: 66}
                        _ = v48
                    } else {
                        switch -2147483648 {
                        case 4294967295:
                            v49 := GenS1{F0: 182, F1: "", F2: 42}
                            var v50 bool = (-2.25 <= 1.5)
                            _ = v49
                            _ = v50
                        case -644:
                            fmt.Println(v1)
                        default:
                            v1 += v1
                            var v51 string = "a"
                            v4 = ("foo bar" + v51)
                            _ = v51
                        }
                        if ((!(!(false && true))) || true) {
                            switch (fi45 - fi46) {
                            case 42:
                                fmt.Println(v34, v4, fi46)
                            case -692:
                                switch (fi46 + (^(fi46))) {
                                case -993:
                                    var v52 float32 = v24
                                    _ = v52
                                case 32767:
                                    var v53 float64 = 1.5
                                    _ = v53
                                }
                            }
                        } else {
                            switch ("" + v4) {
                            case "hello":
                                switch (v3 * (fi46 >> ((42) & 63))) {
                                case 42:
                                    var v54 string = v4
                                    _ = v54
                                default:
                                    v24 += 3.14159
                                }
                            case "a":
                                if false {
                                    v5 = append(v5, "")
                                } else {
                                    var v55 uint32 = 2147483647
                                    _ = v55
                                }
                            case "":
                                var v56 int16 = (v42 - (v42 << ((0) & 63)))
                                v24 = ((v34 + -2.25) - float32(((3.14159 + (0.0 + -2.25)) * ((-2.25 - (0.0 - -2.25)) + 3.14159))))
                                switch v4 {
                                case "foo bar":
                                    switch v3 {
                                    case -9223372036854775808:
                                        v57 := GenS1{F0: 193, F1: "a", F2: 2}
                                        _ = v57
                                    case -1:
                                        if (fi45 != (^(fi46))) {
                                            switch ((v4 + v4) + (((v4 + ((v4 + v4) + "hello")) + v4) + v4)) {
                                            case "":
                                                v5 = append(v5, "\n\t")
                                            default:
                                                if true {
                                                    v24 -= ((v34 - (v24 - 1e10)) + float32(float64(v24)))
                                                } else {
                                                    fmt.Println(v4)
                                                    v58 := GenS1{F0: 154, F1: v4, F2: 2147483648}
                                                    _ = v58
                                                }
                                            }
                                            _ = len(v5)
                                        }
                                    default:
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        if false {
                                            var v59 float32 = -2.25
                                            _ = v59
                                        }
                                    }
                                case "\n\t":
                                    fmt.Println(fi45, v42, fi45)
                                default:
                                    v5 = append(v5, "foo bar")
                                }
                                _ = v56
                            default:
                                if true {
                                    if true {
                                        _ = len(v5)
                                    } else {
                                        v34 += v34
                                        v60 := GenS7{F0: 781, F1: (v2 << ((18446744073709551615) & 63)), F2: (false || (false && (v4 <= v4)))}
                                        fmt.Println(v1)
                                        _ = v60
                                    }
                                }
                            }
                            _ = len(v5)
                        }
                        v5 = append(v5, (v4 + ("a" + v4)))
                    }
                    if false {
                        switch (v3 * 128) {
                        case -302:
                            if false {
                                var v61 int64 = 282
                                v2 = 65535
                                var v62 uint16 = 255
                                _ = v61
                                _ = v62
                            }
                        default:
                            v4 = (v4 + "")
                            v42 = 1000
                            var v63 int8 = -1
                            _ = v63
                        }
                    }
                    fmt.Println(v3)
                    _ = fi46
                    _ = v47
                }
                var v64 uint16 = 131
                _ = v64
            } else {
                v65 := GenS5{F0: 71}
                _ = len(v5)
                _ = v65
            }
            _ = fi45
        }
    }
    v66 := GenS8{F0: 1, F1: -1, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v34
    _ = v38
    _ = v4
    _ = v42
    _ = v5
    _ = v66
}
