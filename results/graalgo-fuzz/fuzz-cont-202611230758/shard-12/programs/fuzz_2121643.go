package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 int; F1 bool; F2 uint32 }
type GenS7 struct { F0 string; F1 uint32 }
type GenS8 struct { F0 uint16; F1 uint16; F2 uint }
type GenS9 struct { F0 bool; F1 int32; F2 int64 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int; F1 int16 }

func main() {
    var v1 int8 = 1
    v2 := []int{128}
    var v3 uint = 55
    v4 := []uint32{52, 761}
    v5 := []string{("\n\t" + ("hello" + "")), ("hello" + ("\n\t" + "a")), "a", ""}
    var v6 int = 929
    v3 -= v3
    v7 := GenS1{F0: 512}
    v8 := GenS1{F0: 42}
    for fi9 := 0; fi9 < 9; fi9++ {
        for fi10 := 0; fi10 < 1; fi10++ {
            v11 := GenS1{F0: -320}
            switch ("foo bar" + "a") {
            case "a":
                var v12 bool = true
                _ = v12
            case "foo bar":
                var v13 int = (fi9 - 128)
                _ = v13
            case "":
                var v14 float32 = (1e10 - 1.5)
                _ = v14
            default:
                _ = len(v2)
            }
            _ = fi10
            _ = v11
        }
        if true {
            switch int(v3) {
            case -2147483648:
                if (-1 >= 1) {
                    v15 := GenS2{F0: v1, F1: 42}
                    _ = v15
                } else {
                    switch "\n\t" {
                    case "foo bar":
                        if (191 == 36) {
                            var v16 uint32 = 4294967295
                            fmt.Println(v16, v1, v1)
                            var v17 uint64 = 42
                            _ = v16
                            _ = v17
                        }
                    case "\n\t":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case "a":
                        v18 := GenS2{F0: v1, F1: 128}
                        _ = v18
                    default:
                        switch v6 {
                        case 331:
                            if (!(844 != -9223372036854775808)) {
                                if ((!true) && (false && ((-(v1)) >= v1))) {
                                    if false {
                                        var v19 uint32 = 589
                                        for fi20 := 0; fi20 < 2; fi20++ {
                                            if (111 == 0) {
                                                if (!(!(uint64(v3) > 0))) {
                                                    v3 = v3
                                                    v21 := GenS2{F0: v1, F1: 42}
                                                    v6 = 1
                                                    _ = v21
                                                }
                                                v1 = v1
                                                v1 -= (-(v1))
                                            }
                                            _ = fi20
                                        }
                                        fmt.Println(v6)
                                        _ = v19
                                    } else {
                                        for fi22 := 0; fi22 < 18; fi22++ {
                                            v1 -= (v1 & (v1 & (v1 - (v1 ^ -1))))
                                            fmt.Println(v1, fi9)
                                            v6 = 32767
                                            _ = fi22
                                        }
                                        switch (("hello" + "foo bar") + "a") {
                                        case "\n\t":
                                            for fi23 := 0; fi23 < 15; fi23++ {
                                                var v24 uint16 = 42
                                                _ = fi23
                                                _ = v24
                                            }
                                        case "a":
                                            v25 := GenS3{F0: 42, F1: 2147483647}
                                            _ = v25
                                        }
                                    }
                                }
                                for fi26 := 0; fi26 < 10; fi26++ {
                                    v27 := GenS4{F0: 0, F1: v1}
                                    _ = fi26
                                    _ = v27
                                }
                            }
                        case -291:
                            v1 -= v1
                        case 1:
                            for fi28 := 0; fi28 < 9; fi28++ {
                                v29 := GenS5{F0: 176, F1: "foo bar"}
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v30 := GenS5{F0: (^(v3)), F1: "foo bar"}
                                var v31 int8 = v1
                                _ = fi28
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                        default:
                            for fi32 := 0; fi32 < 10; fi32++ {
                                var v33 uint16 = 1
                                if (true || (!(true || (2 >= 18446744073709551615)))) {
                                    if true {
                                        v34 := GenS6{F0: (v6 * (v6 / ((fi9) | 1))), F1: false, F2: 616}
                                        v3 -= 4294967295
                                        _ = v34
                                    }
                                    v35 := GenS5{F0: (v3 * 1), F1: ("\n\t" + "foo bar")}
                                    v36 := GenS7{F0: ((("a" + "") + "") + (("hello" + "hello") + "\n\t")), F1: 65535}
                                    _ = v35
                                    _ = v36
                                }
                                v3 = 2
                                _ = fi32
                                _ = v33
                            }
                        }
                    }
                    for fi37 := 0; fi37 < 5; fi37++ {
                        v38 := GenS8{F0: 32767, F1: 848, F2: (v3 - v3)}
                        v1 -= v1
                        var v39 string = ("" + "hello")
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                }
            case 127:
                var v40 uint = (v3 & 0)
                _ = v40
            }
            switch "a" {
            case "\n\t":
                v6 = (-(v6))
            case "hello":
                for fi41 := 0; fi41 < 10; fi41++ {
                    v42 := GenS9{F0: true, F1: -1, F2: 343}
                    v43 := GenS10{F0: v3}
                    _ = fi41
                    _ = v42
                    _ = v43
                }
            default:
                switch (v6 * fi9) {
                case 1:
                    var v44 float32 = float32(1e10)
                    _ = v44
                default:
                    for fi45 := 0; fi45 < 15; fi45++ {
                        var v46 uint64 = 4294967295
                        _ = fi45
                        _ = v46
                    }
                }
            }
        } else {
            switch fi9 {
            case 4294967295:
                v47 := GenS9{F0: (-1 == -1), F1: 42, F2: 371}
                _ = v47
            case 458:
                fmt.Println(fi9, fi9, v6)
            default:
                var v48 int8 = v1
                _ = v48
            }
            if true {
                var v49 int64 = 9223372036854775807
                v50 := GenS11{F0: -1, F1: -1}
                _ = v49
                _ = v50
            }
        }
        _ = fi9
    }
    fmt.Println(v1)
    if len(v2) > 0 {
        v2[0] = (v6 % (((v6 | -947)) | 1))
    }
    var v51 uint16 = 255
    if true {
        if (1 >= uint8(v51)) {
            v3 -= 42
        } else {
            switch (v6 * v6) {
            case -74:
                var v52 uint16 = 42
                _ = v52
            case 4294967295:
                v53 := GenS4{F0: 903, F1: v1}
                _ = v53
            case -1:
                var v54 uint64 = 627
                _ = v54
            default:
                v51 += (v51 * v51)
            }
            if false {
                var v55 uint8 = 128
                _ = v55
            } else {
                v56 := GenS11{F0: (v6 ^ 1), F1: int16(v6)}
                _ = v56
            }
        }
    } else {
        fmt.Println(v6, v3, v3)
        var v57 int64 = -669
        _ = v57
    }
    v3 = (v3 - v3)
    var v58 uint64 = 2
    fmt.Println(v6, v58, v6)
    var v59 float64 = 1e10
    switch (v6 << ((314) & 63)) {
    case 822:
        for fi60 := 0; fi60 < 2; fi60++ {
            if ((false || false) || false) {
                var v61 uint64 = (^(v58))
                for fi62 := 0; fi62 < 14; fi62++ {
                    if len(v5) > 0 {
                        v5[0] = ""
                    }
                    var v63 uint8 = uint8(v58)
                    v2 = append(v2, fi60)
                    _ = fi62
                    _ = v63
                }
                v64 := GenS9{F0: false, F1: -485, F2: 0}
                _ = v61
                _ = v64
            } else {
                v51 = v51
                switch (-(v6)) {
                case 128:
                    var v65 string = "a"
                    _ = v65
                case -714:
                    v58 = 0
                }
                for fi66 := 0; fi66 < 7; fi66++ {
                    var v67 uint32 = 65535
                    if len(v4) > 0 {
                        v4[0] = (v67 | 42)
                    }
                    _ = fi66
                    _ = v67
                }
            }
            fmt.Println(v6)
            _ = fi60
        }
    default:
        fmt.Println(v3, v51, v1)
    }
    var v68 uint8 = uint8(v58)
    v58 -= 18446744073709551615
    v59 -= ((v59 / v59) + ((v59 * v59) + 0.1))
    if true {
        var v69 uint = uint(v68)
        var v70 string = "a"
        v71 := GenS2{F0: v1, F1: 127}
        _ = v69
        _ = v70
        _ = v71
    } else {
        var v72 uint8 = 255
        _ = v72
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v58
    _ = v59
    _ = v6
    _ = v68
    _ = v7
    _ = v8
}
