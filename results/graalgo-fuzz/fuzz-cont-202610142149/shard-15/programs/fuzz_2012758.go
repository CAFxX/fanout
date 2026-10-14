package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 uint32 }

func main() {
    var v1 bool = (!(1 != -128))
    v2 := make(map[int]uint32)
    var v3 uint = 4294967295
    v4 := []string{"foo bar", "", "foo bar"}
    v5 := make(map[uint32]int)
    var v6 string = ("hello" + "")
    if (-543 > 1) {
        v7 := GenS1{F0: 42}
        v8 := GenS1{F0: 0}
        switch -501 {
        case 664:
            switch -9223372036854775808 {
            case 1:
                switch 0 {
                case -9223372036854775808:
                    switch v6 {
                    case "foo bar":
                        var v9 int8 = -8
                        _ = v9
                    case "hello":
                        v3 = (v3 << ((v3) & 63))
                    }
                default:
                    v10 := v2[4294967295]
                    _ = v10
                }
            default:
                fmt.Println(v6, v6, v1)
            }
        case 2147483647:
            var v11 string = v6
            _ = v11
        case -842:
            if true {
                if (v1 || v1) {
                    v3 -= 489
                    var v12 int16 = -733
                    var v13 bool = v1
                    _ = v12
                    _ = v13
                }
            } else {
                for fi14 := 0; fi14 < 2; fi14++ {
                    fmt.Println(v1, fi14)
                    var v15 uint64 = 2
                    if (!(!(((!true) && false) && v1))) {
                        var v16 float64 = (1e10 + 3.14159)
                        var v17 int8 = 100
                        if (!v1) {
                            v3 = (v3 << (((v3 - v3)) & 63))
                            fmt.Println(v3)
                        } else {
                            var v18 string = "hello"
                            v4 = append(v4, v6)
                            switch v18 {
                            case "foo bar":
                                fmt.Println(v1)
                            case "":
                                v1 = ((((v6 + "a") + v18) + (v6 + ("\n\t" + "\n\t"))) == (("a" + ("a" + v6)) + (v18 + "a")))
                            case "a":
                                var v19 uint64 = v15
                                _ = v19
                            }
                            _ = v18
                        }
                        _ = v16
                        _ = v17
                    } else {
                        var v20 int = fi14
                        v21 := v2[v20]
                        _ = v20
                        _ = v21
                    }
                    _ = fi14
                    _ = v15
                }
            }
        default:
            var v22 uint32 = 1
            _ = v22
        }
        _ = v7
        _ = v8
    } else {
        v23 := GenS1{F0: 4294967295}
        if (int32(v3) == -2147483648) {
            for fi24 := 0; fi24 < 1; fi24++ {
                var v25 float64 = 1e10
                fmt.Println(v1, v25, v3)
                _ = len(v4)
                _ = fi24
                _ = v25
            }
            var v26 uint16 = 0
            v27 := GenS2{F0: 908, F1: 98}
            _ = v26
            _ = v27
        }
        _ = v23
    }
    if true {
        v3 = 1
        if ((v3 - v3) >= v3) {
            if (!true) {
                fmt.Println(v1, v1, v6)
            } else {
                if v1 {
                    if (128 == 32767) {
                        if (!(!(!(true || ((1 >= -216) || (42 > 65535)))))) {
                            if (v1 && true) {
                                if v1 {
                                    for fi28 := 0; fi28 < 5; fi28++ {
                                        var v29 int64 = 0
                                        _ = len(v5)
                                        v1 = (!false)
                                        _ = fi28
                                        _ = v29
                                    }
                                    var v30 int32 = 1
                                    _ = v30
                                }
                                fmt.Println(v6, v3)
                                for fi31 := 0; fi31 < 17; fi31++ {
                                    switch (fi31 % ((fi31) | 1)) {
                                    case -648:
                                        v32 := GenS3{F0: 1, F1: 0}
                                        _ = v32
                                    default:
                                        var v33 string = (v6 + v6)
                                        _ = v33
                                    }
                                    var v34 bool = (42 < 57)
                                    _ = fi31
                                    _ = v34
                                }
                            } else {
                                var v35 int8 = 1
                                v35 = v35
                                switch 9223372036854775807 {
                                case -9223372036854775808:
                                    if len(v4) > 0 {
                                        v4[0] = (("" + "") + "foo bar")
                                    }
                                case 127:
                                    v6 = ""
                                default:
                                    if v1 {
                                        v3 = (v3 + v3)
                                        fmt.Println(v1, v35)
                                    }
                                }
                                _ = v35
                            }
                        }
                        var v36 uint32 = 2
                        for fi37 := 0; fi37 < 3; fi37++ {
                            if (42 != 1) {
                                var v38 uint32 = v36
                                var v39 int64 = -668
                                _ = v38
                                _ = v39
                            }
                            _ = fi37
                        }
                        _ = v36
                    }
                } else {
                    var v40 string = (("\n\t" + v6) + "\n\t")
                    var v41 string = "a"
                    _ = v40
                    _ = v41
                }
            }
        } else {
            v1 = false
            v3 = v3
            fmt.Println(v3)
        }
        v4 = append(v4, (v6 + (v6 + "hello")))
    } else {
        var v42 bool = (42 >= 0)
        if (!v42) {
            v4 = append(v4, "hello")
        } else {
            v5[2147483647] = 127
            var v43 uint32 = 42
            var v44 int16 = 1
            _ = v43
            _ = v44
        }
        _ = v42
    }
    var v45 int64 = -9223372036854775808
    switch "\n\t" {
    case "a":
        v6 = "foo bar"
    default:
        if false {
            delete(v2, 127)
        }
    }
    var v46 uint64 = 344
    for fi47 := 0; fi47 < 17; fi47++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if (255 >= (v3 / ((uint(v46)) | 1))) {
            v48 := GenS4{F0: (v46 ^ 0), F1: "hello"}
            v45 = v45
            _ = v48
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v49 int = 573
            var v50 string = ("" + "hello")
            _ = v49
            _ = v50
        }
        v51 := GenS5{F0: 0}
        _ = fi47
        _ = v51
    }
    var v52 int16 = 0
    delete(v5, 943)
    for fi53 := 0; fi53 < 10; fi53++ {
        if (255 <= 2) {
            _ = len(v4)
        } else {
            var v54 int64 = v45
            _ = v54
        }
        v45 = -9223372036854775808
        _ = fi53
    }
    if v1 {
        var v55 uint16 = 330
        if v1 {
            v3 = v3
            fmt.Println(v46, v55, v6)
        } else {
            fmt.Println(v45, v55)
            var v56 string = v6
            _ = v56
        }
        switch "\n\t" {
        case "hello":
            var v57 string = v6
            _ = v57
        case "\n\t":
            switch ("" + (v6 + v6)) {
            case "hello":
                var v58 uint = (v3 | uint(v46))
                _ = v58
            case "foo bar":
                var v59 uint16 = v55
                _ = v59
            case "a":
                for fi60 := 0; fi60 < 7; fi60++ {
                    var v61 int16 = 0
                    _ = fi60
                    _ = v61
                }
            }
        }
        _ = v55
    } else {
        var v62 uint = (v3 ^ (v3 | v3))
        _ = v62
    }
    for fi63 := 0; fi63 < 19; fi63++ {
        var v64 uint8 = uint8(v3)
        _ = fi63
        _ = v64
    }
    var v65 string = ("foo bar" + v6)
    for fi66 := 0; fi66 < 12; fi66++ {
        if (128 == 118) {
            v4 = append(v4, v65)
            v67 := GenS4{F0: 18446744073709551615, F1: v65}
            _ = v67
        } else {
            fmt.Println(v46)
            var v68 int32 = -2147483648
            v69 := GenS4{F0: (^(v46)), F1: ((v6 + v65) + v6)}
            _ = v68
            _ = v69
        }
        fmt.Println(v1)
        for fi70 := 0; fi70 < 19; fi70++ {
            if (v1 || v1) {
                var v71 int16 = (v52 - 1000)
                _ = len(v4)
                var v72 float32 = 0.1
                fmt.Println(v3)
                _ = v71
                _ = v72
            }
            _ = fi70
        }
        _ = fi66
    }
    v3 += 255
    if (v1 || ((0.0 + 1e10) > 1.5)) {
        fmt.Println(v3)
    } else {
        if v1 {
            var v73 uint8 = 42
            if false {
                if (!false) {
                    var v74 bool = (!true)
                    _ = v74
                } else {
                    var v75 int8 = -128
                    if (v1 && true) {
                        if v1 {
                            v46 -= v46
                            v65 = v65
                        }
                    }
                    _ = v75
                }
            }
            for fi76 := 0; fi76 < 7; fi76++ {
                _ = len(v4)
                _ = fi76
            }
            _ = v73
        }
    }
    if v1 {
        var v77 int8 = -1
        var v78 float32 = 1e10
        _ = v77
        _ = v78
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v52
    _ = v6
    _ = v65
}
