package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32; F1 int64; F2 string }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int16; F1 uint64 }
type GenS8 struct { F0 int8; F1 int8 }
func genFn0() int16 {
    var v5 float32 = 0.0
    if (42 < 42) {
        var v6 int16 = 1000
        v7 := []int{127, 599, 2147483647}
        var v8 int64 = -231
        _ = v6
        _ = v7
        _ = v8
    } else {
        var v9 uint32 = 542
        _ = v9
    }
    _ = v5
    return 133
}
func genFn1(p0 uint32) int {
    for fi10 := 0; fi10 < 19; fi10++ {
        switch 127 {
        case 2147483647:
            var v11 int32 = 558
            _ = v11
        }
        _ = fi10
    }
    if false {
        v12 := GenS1{F0: 1, F1: 42, F2: (!true)}
        var v13 int8 = 69
        var v14 int16 = 1000
        _ = v12
        _ = v13
        _ = v14
    }
    p0 += 2147483647
    _ = p0
    return 4294967295
}

func main() {
    var v1 uint8 = 0
    v2 := []uint32{}
    var v3 string = ""
    var v4 string = (((("hello" + "foo bar") + v3) + "hello") + "a")
    v3 = ((v3 + (v3 + ("" + (v4 + v3)))) + ("" + v4))
    for fi15 := 0; fi15 < 3; fi15++ {
        for fi16 := 0; fi16 < 4; fi16++ {
            if (65535 <= 65535) {
                var v17 int64 = 669
                _ = v17
            } else {
                var v18 uint = 872
                switch fi15 {
                case 42:
                    v19 := GenS1{F0: (^(v1)), F1: 366, F2: (18446744073709551615 >= 18446744073709551615)}
                    _ = v19
                case -9223372036854775808:
                    v1 = v1
                case -564:
                    if ((-2.25 / (1.5 + 0.1)) <= 1e10) {
                        switch (-(fi16)) {
                        case 1:
                            v20 := GenS1{F0: v1, F1: 875, F2: (((v1 * v1) < v1) && true)}
                            _ = v20
                        case 0:
                            v3 = v3
                        }
                        v1 = (v1 >> ((v18) & 63))
                        var v21 string = v4
                        if (!(65535 > v18)) {
                            if false {
                                if true {
                                    v4 = v3
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                v22 := GenS1{F0: 127, F1: 2, F2: (float32(3.14159) >= 1.5)}
                                _ = v22
                            }
                            switch ((v21 + "") + v4) {
                            case "hello":
                                var v23 float32 = 0.1
                                _ = v23
                            case "a":
                                v24 := GenS1{F0: (v1 | (^(v1))), F1: 441, F2: ((-278 > 777) || true)}
                                _ = v24
                            default:
                                if (!false) {
                                    switch fi16 {
                                    case 127:
                                        v25 := GenS1{F0: 118, F1: 255, F2: (false || false)}
                                        _ = v25
                                    default:
                                        if (true || ((9223372036854775807 < -520) && (false && false))) {
                                            var v26 int8 = 100
                                            var v27 uint32 = 239
                                            _ = v26
                                            _ = v27
                                        } else {
                                            v4 = ""
                                            var v28 bool = (("foo bar" < "\n\t") || true)
                                            _ = v28
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        var v29 uint = (v18 >> (((v18 & v18)) & 63))
                                        _ = v29
                                    }
                                    if (!false) {
                                        switch v3 {
                                        case "a":
                                            switch ("hello" + ((("a" + "foo bar") + v21) + v21)) {
                                            case "":
                                                fmt.Println(v1, fi15)
                                            case "hello":
                                                if true {
                                                    var v30 uint64 = 849
                                                    if false {
                                                        if (true || (((884 != -367) && (fi15 <= int(v18))) || (42 > -2147483648))) {
                                                            var v31 uint16 = uint16(v30)
                                                            _ = v31
                                                        } else {
                                                            fmt.Println(v4, v21, v21)
                                                            switch v4 {
                                                            case "hello":
                                                                var v32 float64 = (0.1 / -2.25)
                                                                _ = v32
                                                            case "foo bar":
                                                                var v33 int = fi15
                                                                _ = v33
                                                            case "\n\t":
                                                                if true {
                                                                    v34 := GenS1{F0: 127, F1: 598, F2: ((^(v30)) != (^(v30)))}
                                                                    v35 := GenS1{F0: 2, F1: 2, F2: true}
                                                                    _ = v34
                                                                    _ = v35
                                                                } else {
                                                                    fmt.Println(v4)
                                                                }
                                                            default:
                                                                v3 = v4
                                                            }
                                                        }
                                                        v4 = v3
                                                        switch v4 {
                                                        case "a":
                                                            v2 = append(v2, 42)
                                                        case "":
                                                            switch "\n\t" {
                                                            case "hello":
                                                                var v36 uint16 = 2
                                                                v37 := GenS2{F0: 42}
                                                                v36 -= v36
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                v38 := GenS2{F0: -1}
                                                                v39 := GenS1{F0: 128, F1: v36, F2: (!(v21 < (v3 + (v4 + v4))))}
                                                                _ = v36
                                                                _ = v37
                                                                _ = v38
                                                                _ = v39
                                                            case "foo bar":
                                                                if false {
                                                                    var v40 uint = (^(v18))
                                                                    v41 := GenS3{F0: 65535, F1: (v1 ^ v1)}
                                                                    _ = v40
                                                                    _ = v41
                                                                } else {
                                                                    var v42 string = ("hello" + v21)
                                                                    _ = v42
                                                                }
                                                            }
                                                        }
                                                        if false {
                                                            var v43 int = (fi16 % ((genFn1(0)) | 1))
                                                            _ = v43
                                                        }
                                                        v18 = (v18 << ((18446744073709551615) & 63))
                                                        var v44 uint = (v18 + (v18 / ((2) | 1)))
                                                        _ = v44
                                                    } else {
                                                        var v45 float32 = 0.1
                                                        var v46 float32 = v45
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                    _ = v30
                                                } else {
                                                    v47 := GenS3{F0: 2, F1: 219}
                                                    _ = v47
                                                }
                                            default:
                                                fmt.Println(fi15)
                                            }
                                        case "\n\t":
                                            if true {
                                                var v48 uint64 = 991
                                                v2 = append(v2, uint32(v1))
                                                _ = v48
                                            } else {
                                                v1 = v1
                                                if false {
                                                    v49 := GenS4{F0: -9223372036854775808}
                                                    _ = v49
                                                }
                                            }
                                        }
                                        fmt.Println(v18, v4)
                                        if false {
                                            if true {
                                                var v50 float32 = ((1.5 - ((((3.14159 + 0.1) + (0.0 * 0.0)) + (1e10 / 3.14159)) / 0.1)) - (3.14159 + (0.1 / 1.0)))
                                                _ = v50
                                            } else {
                                                switch v4 {
                                                case "a":
                                                    var v51 uint8 = v1
                                                    _ = v51
                                                }
                                                var v52 int = genFn1(682)
                                                _ = v52
                                            }
                                            if true {
                                                v1 += 42
                                            } else {
                                                v18 += 808
                                                v53 := GenS1{F0: (v1 ^ 197), F1: 99, F2: ((!false) && false)}
                                                switch 1 {
                                                case 2147483647:
                                                    var v54 int64 = 1
                                                    _ = v54
                                                case 477:
                                                    var v55 float32 = -2.25
                                                    _ = v55
                                                }
                                                _ = v53
                                            }
                                            v56 := GenS5{F0: 2147483647, F1: 618, F2: "hello"}
                                            _ = v56
                                        }
                                    }
                                }
                            }
                            switch (("foo bar" + "") + v21) {
                            case "hello":
                                switch ("\n\t" + v4) {
                                case "\n\t":
                                    var v57 int8 = 100
                                    _ = v57
                                default:
                                    v21 = v4
                                }
                            case "a":
                                if (v18 < v18) {
                                    var v58 uint32 = 1
                                    _ = v58
                                }
                            case "":
                                var v59 int = (fi15 * (^(fi15)))
                                _ = v59
                            }
                        }
                        _ = v21
                    }
                    var v60 string = ((v4 + v4) + ((v3 + ((v3 + v3) + (((v4 + ("foo bar" + "\n\t")) + "a") + "hello"))) + v3))
                    v60 = v4
                    _ = v60
                }
                if false {
                    if len(v2) > 0 {
                        v2[0] = uint32(v1)
                    }
                } else {
                    if false {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if false {
                            fmt.Println(v4, v1)
                            if ((!false) && true) {
                                var v61 string = v4
                                switch (-(fi16)) {
                                case 32768:
                                    v4 = v3
                                case -211:
                                    if (false || true) {
                                        v62 := GenS1{F0: 128, F1: uint16(v1), F2: true}
                                        if len(v2) > 0 {
                                            v2[0] = 42
                                        }
                                        v61 = v4
                                        _ = v62
                                    } else {
                                        var v63 int = fi16
                                        _ = v63
                                    }
                                case 0:
                                    var v64 uint64 = 42
                                    _ = v64
                                }
                                _ = v61
                            }
                            var v65 bool = true
                            _ = v65
                        }
                        v66 := GenS6{F0: 1}
                        _ = v66
                    } else {
                        v67 := GenS7{F0: 1000, F1: 676}
                        var v68 uint = v18
                        fmt.Println(v3, fi15)
                        if false {
                            fmt.Println(v1, v3)
                        } else {
                            fmt.Println(v3, v1)
                            var v69 bool = false
                            v4 = v3
                            _ = v69
                        }
                        _ = v67
                        _ = v68
                    }
                }
                _ = v18
            }
            var v70 string = "foo bar"
            _ = fi16
            _ = v70
        }
        _ = fi15
    }
    fmt.Println(v4, v1, v4)
    switch ((v3 + v3) + v3) {
    case "":
        v71 := GenS1{F0: v1, F1: 32767, F2: (!true)}
        _ = v71
    case "a":
        switch 9223372036854775807 {
        case -201:
            _ = len(v2)
        case 96:
            var v72 uint64 = 0
            _ = v72
        case 210:
            var v73 int64 = 1
            _ = v73
        default:
            var v74 float32 = 0.0
            _ = v74
        }
    case "foo bar":
        v1 = (v1 % ((v1) | 1))
    }
    var v75 uint64 = 18
    v76 := GenS2{F0: 0}
    var v77 int = -40
    v78 := GenS8{F0: 100, F1: -1}
    v75 -= 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v75:", v75)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v75
    _ = v76
    _ = v77
    _ = v78
}
