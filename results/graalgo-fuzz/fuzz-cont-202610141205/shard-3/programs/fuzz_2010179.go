package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int64; F2 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 string; F2 uint64 }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 uint8; F1 int16; F2 int8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 bool; F1 bool }
type GenS9 struct { F0 string; F1 int8; F2 int }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint; F1 int; F2 int }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 18446744073709551615
    v3 := make(map[int]uint32)
    var v4 int16 = 32767
    v5 := []int{984}
    fmt.Println(v4)
    var v6 int32 = -660
    if (v4 != 32767) {
        var v7 uint = 18446744073709551615
        _ = v7
    }
    if ((2 == 72) || false) {
        fmt.Println(v2)
        v5 = append(v5, 0)
        if (32767 > 42) {
            v2 += (v2 % (((^(v2))) | 1))
            v6 += v6
            if false {
                v8 := GenS1{F0: 0}
                _ = v8
            } else {
                if (((v2 >= v2) || (-1 != -1)) && false) {
                    if (((v6 ^ v6) < 529) || ("\n\t" == ((("a" + ("hello" + "")) + "\n\t") + ("a" + (("" + "") + "foo bar"))))) {
                        if true {
                            v9 := v1["hello"]
                            for fi10 := 0; fi10 < 10; fi10++ {
                                switch ("" + "a") {
                                case "":
                                    if false {
                                        if false {
                                            var v11 uint8 = 1
                                            _ = v11
                                        }
                                        v12 := v3[1]
                                        v13 := GenS2{F0: (v4 << ((1) & 63)), F1: 9223372036854775807, F2: (^(fi10))}
                                        _ = v12
                                        _ = v13
                                    } else {
                                        for fi14 := 0; fi14 < 5; fi14++ {
                                            if ((v9 >> ((uint(v2)) & 63)) > v9) {
                                                fmt.Println(v9, v6, v9)
                                                v15 := GenS3{F0: 2147483648}
                                                if true {
                                                    var v16 int = (fi10 >> ((0) & 63))
                                                    if (v4 >= v4) {
                                                        v17 := GenS1{F0: int(v6)}
                                                        if true {
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                        } else {
                                                            v9 = 101
                                                            switch v16 {
                                                            case -838:
                                                                var v18 uint64 = 18446744073709551615
                                                                _ = v18
                                                            }
                                                            var v19 bool = ("a" == "\n\t")
                                                            fmt.Println(v9, v19, v4)
                                                            _ = v19
                                                        }
                                                        var v20 int64 = 42
                                                        _ = v17
                                                        _ = v20
                                                    }
                                                    v21 := GenS2{F0: (v4 - -605), F1: 42, F2: fi10}
                                                    _ = v16
                                                    _ = v21
                                                } else {
                                                    v4 += (v4 << ((18446744073709551615) & 63))
                                                    v2 -= (^(v2))
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                }
                                                _ = v15
                                            } else {
                                                v4 = 43
                                                switch fi10 {
                                                case -714:
                                                    v9 -= v9
                                                case 32768:
                                                    v5 = append(v5, (fi10 >> ((93) & 63)))
                                                }
                                            }
                                            v4 = (v4 | (v4 & (v4 * -485)))
                                            _ = fi14
                                        }
                                    }
                                case "\n\t":
                                    v6 = 1
                                case "foo bar":
                                    if (false && true) {
                                        var v22 string = ""
                                        for fi23 := 0; fi23 < 19; fi23++ {
                                            var v24 int = (-(fi10))
                                            fmt.Println(v9)
                                            _ = fi23
                                            _ = v24
                                        }
                                        for fi25 := 0; fi25 < 13; fi25++ {
                                            v9 = (^(v9))
                                            _ = fi25
                                        }
                                        _ = v22
                                    }
                                }
                                _ = fi10
                            }
                            _ = v9
                        }
                        var v26 uint16 = 255
                        _ = v26
                    }
                }
                switch 499 {
                case 1:
                    if ((717 == -129) || false) {
                        v4 -= (v4 - -1)
                        if (0 == -128) {
                            fmt.Println(v6, v2)
                            v27 := GenS2{F0: (v4 | (v4 % ((v4) | 1))), F1: -444, F2: -481}
                            v6 -= v6
                            _ = v27
                        } else {
                            switch "hello" {
                            case "":
                                v6 = v6
                            case "a":
                                v2 = v2
                            case "\n\t":
                                for fi28 := 0; fi28 < 16; fi28++ {
                                    v5 = append(v5, fi28)
                                    _ = fi28
                                }
                            default:
                                v1["foo bar"] = 42
                            }
                            if (v2 <= 42) {
                                v29 := GenS2{F0: v4, F1: -383, F2: -631}
                                _ = v29
                            } else {
                                v30 := v3[32767]
                                v30 -= (v30 | v30)
                                _ = v30
                            }
                            var v31 string = "foo bar"
                            _ = v31
                        }
                        v4 -= (v4 * v4)
                    }
                case 32767:
                    var v32 uint = 4294967295
                    _ = v32
                default:
                    if (42 >= 784) {
                        for fi33 := 0; fi33 < 7; fi33++ {
                            fmt.Println(fi33)
                            v6 = (v6 + (v6 << ((413) & 63)))
                            _ = fi33
                        }
                    }
                }
                switch -238 {
                case -2147483648:
                    _ = len(v5)
                default:
                    for fi34 := 0; fi34 < 6; fi34++ {
                        var v35 float64 = -2.25
                        _ = fi34
                        _ = v35
                    }
                }
            }
        } else {
            if (!(42 < (v6 + (v6 + v6)))) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            } else {
                for fi36 := 0; fi36 < 7; fi36++ {
                    for fi37 := 0; fi37 < 20; fi37++ {
                        switch "foo bar" {
                        case "a":
                            if ((v4 | v4) != (^(v4))) {
                                if ((0.0 * 0.0) != 0.1) {
                                    var v38 string = "hello"
                                    v39 := GenS1{F0: fi36}
                                    _ = v38
                                    _ = v39
                                } else {
                                    v6 = (v6 ^ v6)
                                }
                                switch fi36 {
                                case -9223372036854775808:
                                    switch -2147483648 {
                                    case 42:
                                        if false {
                                            switch (fi37 ^ 528) {
                                            case 281:
                                                var v40 uint16 = 408
                                                _ = v40
                                            case 0:
                                                if false {
                                                    delete(v1, "a")
                                                    var v41 float64 = 1.5
                                                    var v42 int8 = -96
                                                    _ = v41
                                                    _ = v42
                                                } else {
                                                    v5 = append(v5, (fi37 - int(v2)))
                                                    fmt.Println(fi36, v6, fi37)
                                                    v6 = v6
                                                }
                                            case 4294967295:
                                                switch 1 {
                                                case 4294967295:
                                                    v43 := GenS4{F0: 0, F1: "\n\t", F2: (v2 ^ (^(v2)))}
                                                    _ = v43
                                                default:
                                                    if true {
                                                        v6 = (^(v6))
                                                        var v44 float32 = -2.25
                                                        switch fi37 {
                                                        case -590:
                                                            v45 := GenS5{F0: (fi36 & (-(fi36))), F1: (^(v2))}
                                                            _ = v45
                                                        case 632:
                                                            var v46 int64 = -143
                                                            _ = v46
                                                        case 128:
                                                            fmt.Println(v2, v6)
                                                            v2 += v2
                                                        default:
                                                            _ = len(v1)
                                                        }
                                                        _ = v44
                                                    } else {
                                                        switch fi36 {
                                                        case 127:
                                                            v47 := GenS3{F0: 2}
                                                            _ = v47
                                                        default:
                                                            fmt.Println(fi37, v2, v2)
                                                        }
                                                        fmt.Println(v6, v2)
                                                    }
                                                }
                                            }
                                            v6 = -2147483648
                                            if ((1.5 > 3.14159) && ((false || (1.5 >= float64((3.14159 + 1.5)))) || false)) {
                                                v48 := GenS6{F0: 42, F1: v4, F2: 54}
                                                v4 += 32767
                                                _ = v48
                                            }
                                        }
                                    case -1:
                                        _ = len(v5)
                                    }
                                case 0:
                                    var v49 int8 = -128
                                    _ = v49
                                case 32768:
                                    v50 := GenS7{F0: 2147483648}
                                    _ = v50
                                }
                                v6 += v6
                            } else {
                                var v51 int = fi36
                                if (!(((^(v2)) > v2) || (("\n\t" + "\n\t") != "foo bar"))) {
                                    if true {
                                        fmt.Println(v6, v4, v4)
                                        var v52 uint16 = 255
                                        v5 = append(v5, int(v6))
                                        _ = v52
                                    } else {
                                        v51 -= (v51 | int(v2))
                                    }
                                    v53 := GenS6{F0: 2, F1: int16(v2), F2: 42}
                                    if (2 < 0) {
                                        switch "hello" {
                                        case "hello":
                                            var v54 uint64 = v2
                                            _ = v54
                                        case "foo bar":
                                            v3[(v51 % ((v51) | 1))] = 2147483648
                                        case "\n\t":
                                            var v55 uint8 = 0
                                            _ = v55
                                        default:
                                            switch v51 {
                                            case 128:
                                                var v56 uint = 1
                                                _ = v56
                                            case 32768:
                                                v57 := GenS8{F0: (2 != 65535), F1: (1 < 128)}
                                                _ = v57
                                            }
                                        }
                                        var v58 int16 = (^(v4))
                                        _ = v58
                                    } else {
                                        if len(v5) > 0 {
                                            v5[0] = 0
                                        }
                                        v2 = 1
                                        v59 := GenS9{F0: (("a" + "hello") + (("foo bar" + "a") + "a")), F1: 0, F2: 158}
                                        _ = v59
                                    }
                                    var v60 int = fi36
                                    var v61 int32 = 591
                                    _ = v53
                                    _ = v60
                                    _ = v61
                                } else {
                                    if ((fi37 <= (fi37 >> ((928) & 63))) && false) {
                                        v4 = 42
                                        var v62 int = -558
                                        _ = v62
                                    } else {
                                        fmt.Println(fi36)
                                        switch (fi37 | (fi36 | (fi37 / ((fi37) | 1)))) {
                                        case -9223372036854775808:
                                            v6 += v6
                                        case 229:
                                            v6 = v6
                                        case 32767:
                                            var v63 float32 = -2.25
                                            _ = v63
                                        }
                                    }
                                    v64 := GenS10{F0: (fi37 ^ fi36)}
                                    _ = v64
                                }
                                _ = v51
                            }
                        case "hello":
                            if false {
                                _ = len(v5)
                                var v65 uint = 2
                                _ = v65
                            } else {
                                v4 = v4
                                v3[(fi37 % (((fi36 << ((uint(v2)) & 63))) | 1))] = 0
                                var v66 float64 = 3.14159
                                _ = v66
                            }
                        }
                        var v67 int16 = 382
                        v4 += (v4 & int16(v2))
                        var v68 uint = 953
                        v67 += v4
                        _ = fi37
                        _ = v67
                        _ = v68
                    }
                    switch fi36 {
                    case -142:
                        var v69 bool = (2147483648 < 732)
                        _ = v69
                    case -1:
                        var v70 uint8 = 255
                        _ = v70
                    }
                    _ = fi36
                }
                for fi71 := 0; fi71 < 1; fi71++ {
                    var v72 int64 = 9223372036854775807
                    _ = fi71
                    _ = v72
                }
                if true {
                    fmt.Println(v6, v2, v2)
                    var v73 uint8 = 17
                    _ = v73
                } else {
                    v74 := GenS11{F0: 2, F1: 42, F2: -748}
                    _ = v74
                }
            }
        }
    }
    var v75 string = ((("foo bar" + ("\n\t" + (("a" + "a") + "foo bar"))) + (("" + "a") + ("a" + "\n\t"))) + (((("\n\t" + "") + ("" + "foo bar")) + (("hello" + "") + ("foo bar" + (("foo bar" + "hello") + "\n\t")))) + "hello"))
    var v76 int16 = 0
    v77 := GenS10{F0: 819}
    for fi78 := 0; fi78 < 2; fi78++ {
        for fi79 := 0; fi79 < 7; fi79++ {
            fmt.Println(v6, v76)
            _ = fi79
        }
        _ = fi78
    }
    var v80 string = v75
    switch v75 {
    case "foo bar":
        for fi81 := 0; fi81 < 16; fi81++ {
            delete(v1, v80)
            _ = fi81
        }
    }
    var v82 string = v80
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v6 = 0
    switch -1 {
    case 576:
        fmt.Println(v76, v6, v82)
    case 128:
        switch 32767 {
        case 342:
            fmt.Println(v6)
        case 42:
            var v83 float32 = float32(-2.25)
            _ = v83
        case -600:
            var v84 uint32 = 519
            _ = v84
        default:
            _ = len(v5)
        }
    case 2147483647:
        if (false || (127 == -2)) {
            v85 := GenS8{F0: (!false), F1: true}
            _ = len(v5)
            _ = v85
        } else {
            var v86 uint64 = v2
            _ = v86
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v80:", v80)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v75
    _ = v76
    _ = v77
    _ = v80
    _ = v82
}
