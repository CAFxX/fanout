package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 uint32; F2 string }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 uint; F1 uint8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int; F1 int8; F2 int64 }

func main() {
    var v1 int16 = 42
    var v2 int16 = -402
    v3 := make(map[uint32]string)
    var v4 float32 = 0.0
    v5 := GenS1{F0: -1, F1: -1}
    if (0.1 < 0.1) {
        fmt.Println(v4)
        var v6 string = "hello"
        if (false || (18446744073709551615 != 379)) {
            var v7 uint8 = 224
            _ = v7
        } else {
            v1 = (v1 - v2)
            for fi8 := 0; fi8 < 13; fi8++ {
                v3[42] = "a"
                if (false && true) {
                    if (!(-1 != 42)) {
                        v9 := v3[2]
                        switch 32767 {
                        case 32767:
                            v10 := []int{fi8}
                            _ = v10
                        case -423:
                            v11 := GenS1{F0: (fi8 / ((fi8) | 1)), F1: 1}
                            _ = v11
                        case 9223372036854775807:
                            if true {
                                v6 = (("foo bar" + ("a" + "\n\t")) + (v9 + (v6 + (v9 + ((v9 + v6) + v9)))))
                                for fi12 := 0; fi12 < 9; fi12++ {
                                    var v13 uint16 = 32767
                                    var v14 int8 = int8(v1)
                                    _ = fi12
                                    _ = v13
                                    _ = v14
                                }
                                v1 = v1
                            } else {
                                for fi15 := 0; fi15 < 10; fi15++ {
                                    var v16 bool = false
                                    _ = fi15
                                    _ = v16
                                }
                            }
                        default:
                            for fi17 := 0; fi17 < 11; fi17++ {
                                if (false && false) {
                                    v18 := GenS2{F0: 42}
                                    if ((65535 >= 42) || (2 < 2147483648)) {
                                        fmt.Println(v4)
                                        _ = len(v3)
                                    }
                                    v19 := GenS1{F0: fi17, F1: 342}
                                    v20 := GenS2{F0: 724}
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                }
                                fmt.Println(fi17, v9)
                                var v21 int16 = v2
                                _ = fi17
                                _ = v21
                            }
                        }
                        v9 = v9
                        _ = v9
                    } else {
                        v2 = (v2 * -1)
                        v22 := GenS2{F0: -2147483648}
                        _ = v22
                    }
                }
                _ = fi8
            }
            var v23 int8 = -128
            _ = v23
        }
        _ = v6
    } else {
        v24 := GenS2{F0: 42}
        _ = v24
    }
    v25 := GenS3{F0: 42, F1: 905, F2: ((("" + "foo bar") + "") + ("\n\t" + "hello"))}
    v26 := []int{42, -971, -211, -2147483648}
    fmt.Println(v1, v4, v4)
    switch -1 {
    case 32767:
        var v27 uint = 788
        _ = v27
    case 1:
        v2 = v2
    }
    var v28 string = (("a" + ("" + "a")) + "foo bar")
    v28 = v28
    var v29 bool = false
    for fi30 := 0; fi30 < 8; fi30++ {
        if false {
            v31 := GenS2{F0: -2147483648}
            var v32 uint64 = 2
            fmt.Println(v1)
            _ = v31
            _ = v32
        }
        v3[2147483647] = (v28 + v28)
        switch (fi30 % ((fi30) | 1)) {
        case -843:
            var v33 int32 = -210
            _ = v33
        case -9223372036854775808:
            v34 := GenS4{F0: 1, F1: int8(v2)}
            _ = v34
        }
        _ = fi30
    }
    v35 := GenS4{F0: 2, F1: 1}
    switch "" {
    case "foo bar":
        switch 127 {
        case -52:
            v2 = v1
        default:
            var v36 int32 = 0
            _ = v36
        }
    case "a":
        if v29 {
            var v37 int16 = v1
            v3[2] = ((v28 + "\n\t") + v28)
            _ = v37
        } else {
            v1 = v1
            v4 -= v4
            _ = len(v26)
        }
    default:
        v2 -= v1
    }
    if (false || (!v29)) {
        var v38 uint8 = 0
        for fi39 := 0; fi39 < 12; fi39++ {
            if v29 {
                if false {
                    if (v29 || true) {
                        delete(v3, 157)
                        var v40 uint8 = (v38 + v38)
                        _ = v40
                    } else {
                        v41 := GenS5{F0: 0, F1: v38}
                        v42 := GenS2{F0: 1}
                        _ = v41
                        _ = v42
                    }
                    fmt.Println(v1, v1, v4)
                    v28 = v28
                } else {
                    var v43 int8 = 42
                    _ = v43
                }
                switch 0 {
                case 972:
                    for fi44 := 0; fi44 < 17; fi44++ {
                        var v45 uint16 = 32767
                        _ = fi44
                        _ = v45
                    }
                }
                var v46 int64 = 42
                _ = v46
            }
            v47 := GenS3{F0: 0, F1: 42, F2: "foo bar"}
            switch (fi39 ^ -2147483648) {
            case 87:
                v48 := GenS4{F0: 42, F1: -1}
                _ = v48
            }
            _ = fi39
            _ = v47
        }
        _ = v38
    } else {
        var v49 uint64 = 18446744073709551615
        _ = v49
    }
    var v50 uint64 = 928
    var v51 string = (("a" + "\n\t") + "\n\t")
    switch 0 {
    case -2147483648:
        var v52 uint64 = v50
        _ = v52
    case 4294967295:
        if v29 {
            if (42 < 32767) {
                for fi53 := 0; fi53 < 5; fi53++ {
                    for fi54 := 0; fi54 < 18; fi54++ {
                        v26 = append(v26, (fi54 ^ fi53))
                        switch v28 {
                        case "hello":
                            var v55 bool = v29
                            _ = v55
                        case "":
                            if len(v26) >= 2 {
                                v26 = v26[1:len(v26)]
                            }
                            if ((-9223372036854775808 < 0) && (v1 != (v1 ^ (-(v2))))) {
                                fmt.Println(v2, v2, fi53)
                            }
                        }
                        var v56 uint32 = 2147483648
                        _ = fi54
                        _ = v56
                    }
                    _ = fi53
                }
                v57 := GenS1{F0: 0, F1: -9223372036854775808}
                var v58 float64 = 0.1
                _ = v57
                _ = v58
            } else {
                for fi59 := 0; fi59 < 4; fi59++ {
                    fmt.Println(v4)
                    for fi60 := 0; fi60 < 8; fi60++ {
                        v29 = ((true && (v51 >= v51)) || (0 == 235))
                        v28 = ((v28 + ((v28 + "") + v28)) + "a")
                        v29 = v29
                        _ = fi60
                    }
                    v26 = append(v26, (fi59 / ((fi59) | 1)))
                    _ = fi59
                }
                v61 := GenS4{F0: 0, F1: 42}
                if (!(uint16(v50) < 834)) {
                    if ((!((0.1 * (0.1 * (-2.25 + 0.0))) == -2.25)) || (4294967295 != 182)) {
                        if (2 < 65535) {
                            for fi62 := 0; fi62 < 4; fi62++ {
                                var v63 string = ((v28 + (v28 + "")) + v28)
                                var v64 int64 = 0
                                v65 := GenS3{F0: -1, F1: 4294967295, F2: v51}
                                _ = fi62
                                _ = v63
                                _ = v64
                                _ = v65
                            }
                        } else {
                            v66 := GenS6{F0: 0}
                            _ = v66
                        }
                        fmt.Println(v28, v1, v51)
                    } else {
                        v29 = (!false)
                        v1 = (^(v2))
                        v1 -= v1
                    }
                    v2 += (v2 % ((v1) | 1))
                } else {
                    var v67 int = 4294967295
                    if true {
                        v50 = v50
                        v1 -= v2
                    } else {
                        if (199 >= 1) {
                            if (42 < 127) {
                                v4 = v4
                                v51 = v51
                            } else {
                                if (-1 >= -9223372036854775808) {
                                    var v68 uint64 = (v50 & 1)
                                    _ = v68
                                }
                                v3[583] = v28
                                v29 = v29
                            }
                        }
                        var v69 bool = true
                        var v70 uint8 = 127
                        _ = v69
                        _ = v70
                    }
                    _ = v67
                }
                _ = v61
            }
        } else {
            switch (("hello" + (v51 + (v51 + ((("foo bar" + v28) + (("hello" + ("hello" + (v28 + "\n\t"))) + v28)) + "foo bar")))) + v28) {
            case "hello":
                _ = len(v26)
            case "foo bar":
                for fi71 := 0; fi71 < 4; fi71++ {
                    var v72 int = fi71
                    _ = fi71
                    _ = v72
                }
            default:
                v73 := GenS7{F0: -2147483648, F1: 42, F2: 0}
                _ = v73
            }
        }
    }
    var v74 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v74
}
