package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 bool }

func main() {
    var v1 int64 = 1
    v2 := make(map[int]bool)
    var v3 uint64 = 0
    var v4 uint16 = 0
    if false {
        v5 := v2[-9223372036854775808]
        var v6 int32 = 0
        switch "" {
        case "hello":
            var v7 uint64 = v3
            _ = v7
        }
        _ = v5
        _ = v6
    } else {
        v8 := GenS1{F0: 2147483648}
        var v9 int64 = (v1 | (v1 << ((361) & 63)))
        _ = v8
        _ = v9
    }
    if true {
        for fi10 := 0; fi10 < 11; fi10++ {
            v3 -= v3
            var v11 int = fi10
            v11 -= v11
            _ = fi10
            _ = v11
        }
        v1 -= v1
        v12 := GenS2{F0: false, F1: 577}
        _ = v12
    }
    switch -687 {
    case -206:
        if (false || true) {
            for fi13 := 0; fi13 < 9; fi13++ {
                v4 -= v4
                fmt.Println(fi13, v1)
                var v14 bool = ((!(0.1 < ((float32(3.14159) - (1e10 / (3.14159 - 0.0))) / -2.25))) && (!((2147483648 < 860) || true)))
                _ = fi13
                _ = v14
            }
            var v15 uint32 = 965
            _ = v15
        }
    case 1:
        for fi16 := 0; fi16 < 2; fi16++ {
            if (1 < 2) {
                fmt.Println(v4)
                var v17 int16 = -32768
                _ = v17
            } else {
                for fi18 := 0; fi18 < 11; fi18++ {
                    fmt.Println(v1)
                    if true {
                        switch "hello" {
                        case "a":
                            var v19 int = (fi18 - fi18)
                            _ = v19
                        case "":
                            v4 += v4
                        case "\n\t":
                            switch "hello" {
                            case "a":
                                _ = len(v2)
                            }
                            var v20 bool = false
                            switch (fi18 / ((fi18) | 1)) {
                            case -9223372036854775808:
                                var v21 int = fi18
                                _ = v21
                            case 590:
                                fmt.Println(v20, fi18)
                            }
                            var v22 uint16 = v4
                            _ = v20
                            _ = v22
                        default:
                            var v23 int = 9223372036854775807
                            _ = v23
                        }
                    } else {
                        v1 -= v1
                        switch "hello" {
                        case "":
                            var v24 int8 = 127
                            _ = v24
                        case "\n\t":
                            var v25 int32 = 1
                            _ = v25
                        case "a":
                            v26 := []int{-9223372036854775808, (-(fi16)), (^(fi18)), (fi16 | fi16)}
                            _ = v26
                        default:
                            v3 = 0
                        }
                    }
                    v3 = v3
                    switch "hello" {
                    case "":
                        v27 := GenS1{F0: 2147483647}
                        _ = v27
                    case "a":
                        v28 := GenS2{F0: false, F1: -269}
                        _ = v28
                    case "hello":
                        var v29 bool = ((v3 | v3) <= 832)
                        _ = v29
                    }
                    v30 := GenS3{F0: v3}
                    v31 := v2[(fi18 - fi18)]
                    if v31 {
                        var v32 uint16 = (v4 & 1)
                        _ = v32
                    } else {
                        v2[fi18] = v31
                        if ((v4 | (^(v4))) >= (v4 + v4)) {
                            var v33 int32 = -608
                            _ = v33
                        } else {
                            var v34 uint = 2
                            if v31 {
                                switch fi16 {
                                case 291:
                                    var v35 uint16 = (v4 >> ((1) & 63))
                                    _ = v35
                                case 4294967295:
                                    var v36 float64 = 1.5
                                    _ = v36
                                default:
                                    var v37 uint8 = 255
                                    _ = v37
                                }
                            } else {
                                fmt.Println(v34, v4)
                                var v38 string = "foo bar"
                                _ = v38
                            }
                            _ = v34
                        }
                        v39 := v2[fi18]
                        if v39 {
                            fmt.Println(v31, v1)
                            v40 := GenS2{F0: (65535 < 65535), F1: 0}
                            _ = v40
                        }
                        _ = v39
                    }
                    _ = len(v2)
                    _ = fi18
                    _ = v30
                    _ = v31
                }
                for fi41 := 0; fi41 < 12; fi41++ {
                    v42 := v2[(fi16 ^ (fi16 - (fi41 & fi41)))]
                    switch fi16 {
                    case 891:
                        switch 128 {
                        case 751:
                            v43 := []int{(fi16 + fi16), int(v3), (fi16 | fi41)}
                            _ = v43
                        case 32767:
                            v4 = 255
                        case 128:
                            if v42 {
                                v42 = true
                            } else {
                                var v44 int8 = 100
                                fmt.Println(v42)
                                _ = v44
                            }
                        }
                    case 627:
                        v3 = (v3 - 42)
                    case 32768:
                        v3 = 18446744073709551615
                    default:
                        v45 := []int{127, fi41, 128, (fi41 >> ((255) & 63))}
                        _ = v45
                    }
                    v46 := GenS2{F0: true, F1: -2147483648}
                    v4 += (v4 >> ((4294967295) & 63))
                    _ = fi41
                    _ = v42
                    _ = v46
                }
            }
            _ = fi16
        }
    default:
        v47 := GenS4{F0: 33}
        _ = v47
    }
    fmt.Println(v4, v1)
    var v48 bool = (("\n\t" + ("" + "\n\t")) < ("hello" + "foo bar"))
    var v49 string = "a"
    for fi50 := 0; fi50 < 8; fi50++ {
        var v51 uint = uint(v3)
        v52 := GenS1{F0: 4294967295}
        _ = fi50
        _ = v51
        _ = v52
    }
    v49 = "\n\t"
    var v53 uint32 = 2
    switch v49 {
    case "hello":
        v54 := v2[-1]
        _ = v54
    case "":
        var v55 uint16 = (v4 % ((v4) | 1))
        _ = v55
    case "a":
        if ((312 != 1) && false) {
            v48 = true
            switch "" {
            case "a":
                switch 32768 {
                case 42:
                    for fi56 := 0; fi56 < 3; fi56++ {
                        v57 := GenS5{F0: 1000, F1: false}
                        var v58 int64 = (v1 << ((983) & 63))
                        _ = fi56
                        _ = v57
                        _ = v58
                    }
                case 9223372036854775807:
                    if (!(v53 != (v53 | uint32(v3)))) {
                        v4 -= (v4 ^ 42)
                        for fi59 := 0; fi59 < 5; fi59++ {
                            var v60 uint64 = 1
                            _ = fi59
                            _ = v60
                        }
                        switch "hello" {
                        case "\n\t":
                            v53 = (v53 << ((903) & 63))
                        case "hello":
                            v61 := []int{127, 32768, 42}
                            _ = v61
                        }
                    }
                default:
                    if (true || (("hello" + "foo bar") >= v49)) {
                        v3 -= (v3 & v3)
                        for fi62 := 0; fi62 < 13; fi62++ {
                            var v63 int8 = 100
                            if ((v53 | 1) != 0) {
                                var v64 uint8 = 128
                                _ = v64
                            } else {
                                if (int(v63) > 4294967295) {
                                    var v65 int32 = 2147483647
                                    _ = v65
                                } else {
                                    v53 = (v53 | (v53 * (v53 & v53)))
                                }
                                var v66 int64 = v1
                                var v67 uint64 = 42
                                _ = v66
                                _ = v67
                            }
                            var v68 int32 = 1
                            _ = fi62
                            _ = v63
                            _ = v68
                        }
                        if v48 {
                            switch 4294967295 {
                            case 2147483647:
                                delete(v2, -858)
                            }
                            fmt.Println(v1)
                        }
                    } else {
                        for fi69 := 0; fi69 < 11; fi69++ {
                            v49 = (("hello" + "") + "\n\t")
                            _ = fi69
                        }
                    }
                }
            case "hello":
                for fi70 := 0; fi70 < 12; fi70++ {
                    var v71 uint64 = 922
                    _ = fi70
                    _ = v71
                }
            default:
                if (-482 > 1) {
                    for fi72 := 0; fi72 < 19; fi72++ {
                        v1 = 42
                        _ = fi72
                    }
                }
            }
            v1 = v1
        } else {
            v73 := []int{32767, -801, 1, -2147483648}
            var v74 uint64 = v3
            _ = v73
            _ = v74
        }
    }
    v75 := []int{32767, -553, -947, -592}
    if v48 {
        if v48 {
            v1 = (-(v1))
            fmt.Println(v4)
        } else {
            v4 -= (v4 - (v4 ^ v4))
            switch ((v49 + v49) + v49) {
            case "":
                for fi76 := 0; fi76 < 7; fi76++ {
                    v3 = (v3 | (^(v3)))
                    _ = fi76
                }
            case "a":
                var v77 int64 = (-(v1))
                _ = v77
            case "foo bar":
                if len(v75) > 0 {
                    v75[0] = -1
                }
            }
        }
        var v78 uint16 = v4
        _ = v78
    }
    if v48 {
        for fi79 := 0; fi79 < 18; fi79++ {
            for fi80 := 0; fi80 < 13; fi80++ {
                if (2 != 188) {
                    if len(v75) > 0 {
                        v75[0] = (^(fi79))
                    }
                    var v81 bool = v48
                    if (!v48) {
                        fmt.Println(v53, v48)
                        if v48 {
                            v82 := GenS5{F0: -63, F1: true}
                            var v83 int = fi79
                            v84 := GenS1{F0: v53}
                            delete(v2, (fi79 + fi80))
                            v75 = append(v75, 127)
                            if false {
                                var v85 int8 = -16
                                var v86 float64 = (0.0 - ((1e10 / float64((((1e10 / -2.25) / 1e10) / (0.1 / -2.25)))) * 1e10))
                                _ = v85
                                _ = v86
                            } else {
                                var v87 int64 = (v1 | v1)
                                var v88 int32 = 2147483647
                                _ = v87
                                _ = v88
                            }
                            _ = v82
                            _ = v83
                            _ = v84
                        }
                    } else {
                        var v89 bool = (494 > 255)
                        _ = v89
                    }
                    _ = v81
                } else {
                    fmt.Println(v3, v4, fi80)
                    fmt.Println(fi80, v1, v1)
                }
                var v90 bool = (23 > 255)
                v91 := GenS4{F0: 42}
                _ = fi80
                _ = v90
                _ = v91
            }
            v1 += -324
            fmt.Println(v4)
            _ = fi79
        }
    }
    v3 -= (v3 - v3)
    v75 = append(v75, 32767)
    v1 += v1
    fmt.Println(v53, v3)
    v1 -= 798
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v75:", len(v75), v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v53
    _ = v75
}
