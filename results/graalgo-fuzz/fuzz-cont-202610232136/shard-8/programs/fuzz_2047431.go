package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint64; F1 bool }
type GenS9 struct { F0 bool; F1 bool }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint32; F1 int8; F2 uint }
type GenS12 struct { F0 uint32 }
func genFn0() uint64 {
    v7 := GenS1{F0: 539}
    var v8 bool = (-491 <= -9223372036854775808)
    if v8 {
        if (!(42 == 324)) {
            v9 := GenS2{F0: 42, F1: -989}
            _ = v9
        }
        v8 = (-2147483648 >= 2147483647)
    }
    fmt.Println(v8)
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := []int{32767, 42, 127, 922}
    v2 := make(map[string]string)
    var v3 float32 = ((((float32(((1e10 * 0.0) / (((1e10 + 0.0) * float64(0.0)) + 0.1))) + 1e10) + 3.14159) / 3.14159) + (3.14159 / 3.14159))
    v4 := make(map[uint32]uint32)
    var v5 string = "hello"
    var v6 string = (v5 + v5)
    v10 := GenS3{F0: 1, F1: genFn0(), F2: 159}
    var v11 string = (v5 + "a")
    var v12 int64 = 1
    if (!(-2147483648 == 0)) {
        if false {
            if true {
                if ((587 >= 0) || (127 >= 255)) {
                    v13 := GenS4{F0: 127}
                    var v14 int16 = 1000
                    v6 = ("foo bar" + (v11 + ((v11 + v11) + v6)))
                    _ = v13
                    _ = v14
                }
            } else {
                v6 = ("" + ("hello" + (v5 + ("\n\t" + "foo bar"))))
                if false {
                    v15 := v4[462]
                    for fi16 := 0; fi16 < 16; fi16++ {
                        v17 := GenS5{F0: 835}
                        _ = fi16
                        _ = v17
                    }
                    v2["\n\t"] = ((v5 + (v5 + v11)) + v11)
                    _ = v15
                }
                fmt.Println(v12, v5)
            }
        } else {
            v11 = v6
        }
    } else {
        if (genFn0() != 0) {
            v6 = ("\n\t" + v5)
        } else {
            for fi18 := 0; fi18 < 9; fi18++ {
                var v19 uint = 4294967295
                switch 1 {
                case 1:
                    switch (("\n\t" + "a") + ("hello" + v11)) {
                    case "a":
                        for fi20 := 0; fi20 < 1; fi20++ {
                            v11 = v11
                            v21 := GenS4{F0: 42}
                            _ = fi20
                            _ = v21
                        }
                    case "":
                        var v22 uint16 = 65535
                        _ = v22
                    case "\n\t":
                        v19 -= v19
                    default:
                        v1 = append(v1, (fi18 & fi18))
                    }
                case -462:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case 127:
                    switch v6 {
                    case "":
                        for fi23 := 0; fi23 < 14; fi23++ {
                            delete(v4, 538)
                            v24 := GenS3{F0: v19, F1: genFn0(), F2: 213}
                            v25 := GenS1{F0: 65535}
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                    }
                }
                switch ((((v6 + v11) + "\n\t") + v5) + "") {
                case "foo bar":
                    fmt.Println(v12)
                default:
                    fmt.Println(v5, v11)
                }
                _ = fi18
                _ = v19
            }
            switch 864 {
            case -1:
                if (0.1 == 0.1) {
                    v26 := GenS6{F0: -2147483648}
                    _ = v26
                } else {
                    v27 := GenS7{F0: 141}
                    _ = v27
                }
            case 32768:
                if (1 != 255) {
                    if true {
                        var v28 int32 = 1
                        switch -9223372036854775808 {
                        case 32768:
                            for fi29 := 0; fi29 < 13; fi29++ {
                                v3 = 1e10
                                for fi30 := 0; fi30 < 17; fi30++ {
                                    fmt.Println(v28)
                                    v5 = ""
                                    _ = fi30
                                }
                                var v31 int64 = (v12 & v12)
                                _ = fi29
                                _ = v31
                            }
                        default:
                            for fi32 := 0; fi32 < 19; fi32++ {
                                fmt.Println(v5)
                                v11 = "foo bar"
                                var v33 int32 = (v28 - v28)
                                _ = fi32
                                _ = v33
                            }
                        }
                        var v34 uint16 = 740
                        _ = v28
                        _ = v34
                    }
                    var v35 int8 = -128
                    _ = v35
                } else {
                    if (!false) {
                        v5 = v5
                        v36 := GenS6{F0: 0}
                        switch v11 {
                        case "":
                            for fi37 := 0; fi37 < 16; fi37++ {
                                v11 = v6
                                _ = fi37
                            }
                        case "foo bar":
                            delete(v2, (((v11 + v5) + v11) + v5))
                        case "\n\t":
                            var v38 int16 = 42
                            _ = v38
                        }
                        _ = v36
                    }
                    for fi39 := 0; fi39 < 18; fi39++ {
                        v6 = "hello"
                        if true {
                            v3 = 1e10
                            v2[(v6 + (("\n\t" + "foo bar") + ("foo bar" + "\n\t")))] = v6
                        }
                        v6 = "foo bar"
                        _ = fi39
                    }
                }
            case 0:
                v5 = "\n\t"
            default:
                var v40 float64 = 1e10
                _ = v40
            }
            var v41 uint8 = 2
            _ = v41
        }
    }
    v42 := GenS1{F0: 360}
    v43 := GenS8{F0: 18446744073709551615, F1: false}
    var v44 string = ""
    if false {
        if false {
            v45 := GenS7{F0: 4294967295}
            var v46 float64 = 0.1
            _ = v45
            _ = v46
        } else {
            v6 = "\n\t"
            var v47 uint8 = 110
            if (!(!((1 >= 0) || (849 <= 597)))) {
                var v48 uint8 = 1
                v49 := GenS9{F0: (1000 > 0), F1: true}
                for fi50 := 0; fi50 < 11; fi50++ {
                    if (42 == 4294967295) {
                        var v51 int = 4294967295
                        fmt.Println(v48, v3)
                        for fi52 := 0; fi52 < 18; fi52++ {
                            var v53 int64 = v12
                            if ((((959 <= 881) || false) || (false && (false || true))) && false) {
                                v54 := GenS6{F0: 0}
                                var v55 int32 = -2147483648
                                _ = v54
                                _ = v55
                            }
                            switch v5 {
                            case "":
                                fmt.Println(v48, v6)
                            case "a":
                                var v56 uint64 = 0
                                _ = v56
                            }
                            _ = fi52
                            _ = v53
                        }
                        _ = v51
                    }
                    for fi57 := 0; fi57 < 5; fi57++ {
                        v6 = "foo bar"
                        var v58 bool = true
                        v59 := GenS9{F0: v58, F1: (88 < 0)}
                        _ = fi57
                        _ = v58
                        _ = v59
                    }
                    v47 += (^(v48))
                    _ = fi50
                }
                _ = v48
                _ = v49
            } else {
                var v60 uint16 = 246
                v61 := GenS5{F0: 0}
                _ = v60
                _ = v61
            }
            _ = v47
        }
        fmt.Println(v6, v11)
        if ((982 != 42) || ((2147483647 < 207) || (18446744073709551615 < 4294967295))) {
            var v62 uint8 = 0
            var v63 int32 = 0
            _ = v62
            _ = v63
        }
    } else {
        if (v5 >= v44) {
            v64 := GenS10{F0: v11}
            v65 := GenS11{F0: 21, F1: -1, F2: 0}
            var v66 float64 = -2.25
            _ = v64
            _ = v65
            _ = v66
        } else {
            v3 += ((float32(0.1) + 0.0) - 1.5)
            v6 = (("hello" + v6) + v44)
            if true {
                for fi67 := 0; fi67 < 8; fi67++ {
                    for fi68 := 0; fi68 < 19; fi68++ {
                        v69 := GenS11{F0: 255, F1: 100, F2: 255}
                        fmt.Println(v11, fi68, v11)
                        _ = fi68
                        _ = v69
                    }
                    _ = fi67
                }
                v3 = v3
                var v70 int64 = (v12 >> ((255) & 63))
                _ = v70
            } else {
                v71 := GenS9{F0: true, F1: (42 >= 128)}
                v72 := GenS12{F0: 10}
                v44 = v44
                _ = v71
                _ = v72
            }
        }
    }
    var v73 uint16 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v73
}
