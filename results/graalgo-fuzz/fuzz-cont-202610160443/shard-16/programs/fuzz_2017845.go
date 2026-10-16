package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 int8; F2 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint }
func genFn0() string {
    if (1 != 42) {
        var v6 int32 = 0
        v7 := []int{805}
        v6 = (-(v6))
        _ = v6
        _ = v7
    } else {
        v8 := GenS1{F0: 2}
        v9 := GenS1{F0: 795}
        var v10 int16 = -662
        _ = v10
        _ = v8
        _ = v9
    }
    v11 := []int{4294967295, 4294967295, 17, 0}
    _ = v11
    return "foo bar"
}
func genFn1() int8 {
    var v12 int = 0
    for fi13 := 0; fi13 < 11; fi13++ {
        v14 := GenS1{F0: 18446744073709551615}
        v12 = fi13
        _ = fi13
        _ = v14
    }
    if true {
        var v15 bool = false
        _ = v15
    } else {
        if (float32(1e10) != (3.14159 + 3.14159)) {
            v12 = (v12 << ((540) & 63))
        }
        for fi16 := 0; fi16 < 9; fi16++ {
            v12 = v12
            for fi17 := 0; fi17 < 4; fi17++ {
                v12 += fi16
                v12 = -899
                _ = fi17
            }
            v18 := []int{(fi16 % ((fi16) | 1))}
            _ = fi16
            _ = v18
        }
        var v19 uint16 = 65535
        _ = v19
    }
    v12 = (v12 + 128)
    _ = v12
    return 100
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = ""
    v3 := make(map[uint32]bool)
    var v4 int16 = 42
    v5 := make(map[string]string)
    if (0 > (v4 & v4)) {
        v4 = 42
    } else {
        if (v2 >= "") {
            v2 = v2
        }
        v20 := GenS1{F0: 42}
        v3[42] = (!true)
        _ = v20
    }
    var v21 float64 = (-2.25 - (float64((((1e10 * (1e10 - 3.14159)) / float32(1e10)) + -2.25)) * (0.0 - 3.14159)))
    var v22 uint32 = 478
    var v23 int = 32768
    var v24 string = "a"
    v25 := []int{}
    if (true || (!false)) {
        if (849 != 0) {
            var v26 int8 = 42
            _ = v26
        }
    } else {
        v22 = (v22 + v22)
    }
    v23 -= 128
    if (0 < 42) {
        if false {
            if (552 > 255) {
                switch v2 {
                case "hello":
                    v24 = (v24 + (v24 + (v24 + "a")))
                }
            }
            var v27 float64 = v21
            switch (v2 + v2) {
            case "foo bar":
                v5[(v24 + v2)] = v2
            }
            _ = v27
        }
        _ = len(v1)
        if ((1e10 > 1.5) || false) {
            switch 1 {
            case 32768:
                v28 := GenS2{F0: v23, F1: -1, F2: 9223372036854775807}
                _ = v28
            case -9223372036854775808:
                for fi29 := 0; fi29 < 8; fi29++ {
                    switch 4294967295 {
                    case -2147483648:
                        v21 = (v21 + (0.1 - 0.0))
                    case 884:
                        v21 = -2.25
                    }
                    var v30 uint64 = 42
                    switch v2 {
                    case "foo bar":
                        if true {
                            for fi31 := 0; fi31 < 11; fi31++ {
                                switch "a" {
                                case "":
                                    if ("foo bar" != v2) {
                                        if (9223372036854775807 >= 42) {
                                            v32 := GenS2{F0: -849, F1: 113, F2: fi31}
                                            v22 = v22
                                            _ = v32
                                        } else {
                                            if true {
                                                v33 := GenS2{F0: (fi29 << ((400) & 63)), F1: genFn1(), F2: 32767}
                                                v4 = v4
                                                _ = v33
                                            }
                                        }
                                        v22 += v22
                                        v22 -= (v22 << ((uint(v22)) & 63))
                                        v24 = v2
                                        v25 = append(v25, (v23 / ((v23) | 1)))
                                        var v34 uint32 = v22
                                        _ = v34
                                    } else {
                                        v2 = "\n\t"
                                        v22 = (v22 ^ 1)
                                    }
                                case "\n\t":
                                    v2 = "\n\t"
                                case "hello":
                                    v35 := GenS3{F0: v2}
                                    _ = v35
                                default:
                                    var v36 int32 = 2147483647
                                    _ = v36
                                }
                                v37 := GenS3{F0: (("\n\t" + (v24 + (("hello" + v24) + ""))) + (v2 + (((v24 + v2) + ((("" + "hello") + "hello") + v24)) + "hello")))}
                                _ = fi31
                                _ = v37
                            }
                            if true {
                                var v38 int = (v23 | (v23 << ((255) & 63)))
                                _ = v38
                            }
                            v23 += 0
                        } else {
                            v22 = 65535
                        }
                    default:
                        if (!true) {
                            if true {
                                for fi39 := 0; fi39 < 18; fi39++ {
                                    v40 := GenS1{F0: (v30 ^ v30)}
                                    var v41 uint32 = 331
                                    _ = len(v25)
                                    fmt.Println(v24)
                                    if len(v25) >= 2 {
                                        v25 = v25[1:len(v25)]
                                    }
                                    _ = fi39
                                    _ = v40
                                    _ = v41
                                }
                            } else {
                                if true {
                                    fmt.Println(v21, v2)
                                } else {
                                    var v42 float64 = 0.1
                                    switch v2 {
                                    case "a":
                                        var v43 float64 = 1.5
                                        _ = v43
                                    case "foo bar":
                                        switch "\n\t" {
                                        case "hello":
                                            for fi44 := 0; fi44 < 19; fi44++ {
                                                v45 := GenS3{F0: (v24 + v2)}
                                                v46 := GenS1{F0: v30}
                                                _ = fi44
                                                _ = v45
                                                _ = v46
                                            }
                                        case "\n\t":
                                            v4 += v4
                                        }
                                    }
                                    fmt.Println(fi29)
                                    _ = v42
                                }
                            }
                            for fi47 := 0; fi47 < 10; fi47++ {
                                v48 := GenS2{F0: fi47, F1: -73, F2: fi29}
                                _ = fi47
                                _ = v48
                            }
                        }
                    }
                    _ = fi29
                    _ = v30
                }
            case 804:
                switch 9223372036854775807 {
                case -120:
                    var v49 float64 = v21
                    _ = v49
                case 128:
                    v50 := v5["a"]
                    _ = v50
                case 32768:
                    if (1.5 >= 3.14159) {
                        var v51 bool = ((127 > 1) || true)
                        v52 := GenS1{F0: 2}
                        _ = v51
                        _ = v52
                    } else {
                        if (!(-408 == 42)) {
                            for fi53 := 0; fi53 < 15; fi53++ {
                                v54 := GenS2{F0: 1, F1: 127, F2: (-(fi53))}
                                _ = fi53
                                _ = v54
                            }
                        } else {
                            var v55 int64 = 264
                            v56 := GenS1{F0: 376}
                            var v57 int32 = 42
                            _ = v55
                            _ = v56
                            _ = v57
                        }
                    }
                }
            default:
                if true {
                    if (2147483647 >= v22) {
                        if (v21 > v21) {
                            v21 += 0.0
                            switch 245 {
                            case -9223372036854775808:
                                v58 := GenS4{F0: (-2.25 >= -2.25)}
                                _ = v58
                            case 875:
                                v59 := GenS4{F0: false}
                                _ = v59
                            default:
                                var v60 bool = (true || ((1 <= 42) && (v23 >= int(v4))))
                                _ = v60
                            }
                            _ = len(v25)
                        }
                        if true {
                            if len(v25) >= 2 {
                                v25 = v25[1:len(v25)]
                            }
                            if true {
                                switch v24 {
                                case "\n\t":
                                    _ = len(v25)
                                }
                                v22 += (^(v22))
                                var v61 int64 = 507
                                _ = v61
                            } else {
                                var v62 float64 = v21
                                _ = v62
                            }
                        }
                        if len(v25) >= 2 {
                            v25 = v25[1:len(v25)]
                        }
                    } else {
                        fmt.Println(v2, v22, v4)
                        v63 := v3[240]
                        _ = v63
                    }
                    if (!((0 == 879) || true)) {
                        var v64 int8 = -9
                        v65 := GenS5{F0: (true && false)}
                        v66 := GenS2{F0: 1, F1: 95, F2: v23}
                        _ = v64
                        _ = v65
                        _ = v66
                    } else {
                        v24 = v2
                    }
                }
            }
            _ = len(v5)
        } else {
            switch (("a" + "foo bar") + v24) {
            case "":
                switch int(v4) {
                case 4294967295:
                    if true {
                        v21 -= v21
                        v67 := GenS2{F0: (v23 + -2147483648), F1: 42, F2: 2147483647}
                        v68 := GenS3{F0: "hello"}
                        _ = v67
                        _ = v68
                    }
                case -1:
                    if true {
                        if ((v2 + v24) <= ("hello" + v2)) {
                            switch "foo bar" {
                            case "foo bar":
                                var v69 int32 = 2147483647
                                _ = v69
                            case "":
                                fmt.Println(v4, v23, v4)
                            case "a":
                                if (false && (((v23 + 528) > int(v22)) || true)) {
                                    v22 = (v22 ^ (v22 ^ v22))
                                    v70 := GenS3{F0: ""}
                                    _ = v70
                                } else {
                                    switch (-(v23)) {
                                    case 128:
                                        switch v23 {
                                        case 32768:
                                            if len(v25) >= 2 {
                                                v25 = v25[1:len(v25)]
                                            }
                                        case 4294967295:
                                            for fi71 := 0; fi71 < 8; fi71++ {
                                                v3[2147483647] = true
                                                for fi72 := 0; fi72 < 3; fi72++ {
                                                    delete(v1, v22)
                                                    _ = fi72
                                                }
                                                _ = fi71
                                            }
                                        }
                                    case 32768:
                                        var v73 uint16 = 42
                                        _ = v73
                                    }
                                    switch 464 {
                                    case 127:
                                        switch (v23 / ((v23) | 1)) {
                                        case -1:
                                            var v74 int8 = 42
                                            _ = v74
                                        }
                                    case 964:
                                        fmt.Println(v21, v23, v22)
                                    case 42:
                                        v2 = "hello"
                                    }
                                    var v75 int = 43
                                    _ = v75
                                }
                            }
                            delete(v1, (v22 - v22))
                        }
                    } else {
                        var v76 uint8 = 128
                        if false {
                            if ((255 < 255) && true) {
                                fmt.Println(v23)
                                for fi77 := 0; fi77 < 10; fi77++ {
                                    _ = len(v25)
                                    _ = fi77
                                }
                                for fi78 := 0; fi78 < 14; fi78++ {
                                    if len(v25) >= 2 {
                                        v25 = v25[1:len(v25)]
                                    }
                                    _ = fi78
                                }
                            }
                            for fi79 := 0; fi79 < 4; fi79++ {
                                fmt.Println(v24)
                                if false {
                                    if ((false && false) || (true || true)) {
                                        v80 := GenS6{F0: 42}
                                        v81 := GenS1{F0: 62}
                                        _ = v80
                                        _ = v81
                                    }
                                }
                                _ = fi79
                            }
                        } else {
                            v22 = v22
                            for fi82 := 0; fi82 < 7; fi82++ {
                                _ = len(v25)
                                v83 := v5["a"]
                                var v84 int = (-(fi82))
                                _ = fi82
                                _ = v83
                                _ = v84
                            }
                        }
                        switch v24 {
                        case "\n\t":
                            v21 -= v21
                        case "":
                            v4 = -32768
                        case "a":
                            var v85 int16 = (v4 - (v4 << ((679) & 63)))
                            _ = v85
                        default:
                            var v86 float64 = float64(((1e10 * (1.5 - float32(v21))) * (((3.14159 - 3.14159) - (1.5 + -2.25)) / (float32(v21) + float32((0.0 + (v21 / (-2.25 / -2.25))))))))
                            _ = v86
                        }
                        _ = v76
                    }
                }
            case "hello":
                switch "\n\t" {
                case "hello":
                    v22 = v22
                case "\n\t":
                    v87 := v3[(^(v22))]
                    _ = v87
                }
            default:
                fmt.Println(v4, v23)
            }
        }
    }
    var v88 int32 = 99
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v88
}
