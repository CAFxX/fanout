package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 bool; F2 bool }
func genFn0() uint {
    var v4 int32 = 1
    v5 := GenS1{F0: 100, F1: ("\n\t" + "\n\t")}
    var v6 int16 = -32768
    v7 := []int{2147483647, -9223372036854775808, 128, -2147483648}
    fmt.Println(v4, v4)
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 255
}

func main() {
    var v1 int16 = 40
    v2 := make(map[int]int)
    var v3 int16 = 1000
    v8 := GenS1{F0: 127, F1: ("hello" + "\n\t")}
    v2[69] = -874
    v1 -= (v1 * (v3 | -32768))
    for fi9 := 0; fi9 < 11; fi9++ {
        v3 = (v1 ^ v1)
        fmt.Println(v1, v1, fi9)
        _ = fi9
    }
    switch ("hello" + ("\n\t" + "a")) {
    case "":
        _ = len(v2)
    case "\n\t":
        var v10 int64 = 1
        _ = v10
    case "a":
        v11 := v2[-9223372036854775808]
        _ = v11
    default:
        switch "" {
        case "":
            v3 -= v1
        default:
            var v12 uint16 = 834
            _ = v12
        }
    }
    switch "hello" {
    case "hello":
        if false {
            var v13 uint32 = 533
            var v14 string = "a"
            for fi15 := 0; fi15 < 16; fi15++ {
                if false {
                    for fi16 := 0; fi16 < 9; fi16++ {
                        fmt.Println(v1, v14)
                        v3 = v3
                        v17 := []int{-9223372036854775808, (fi15 + fi16), (fi16 >> ((636) & 63)), fi15}
                        _ = fi16
                        _ = v17
                    }
                } else {
                    if false {
                        switch v14 {
                        case "":
                            if (fi15 > (-(fi15))) {
                                if ((fi15 << ((222) & 63)) < fi15) {
                                    v1 = 0
                                }
                                var v18 int32 = 938
                                var v19 float32 = float32(1.5)
                                _ = v18
                                _ = v19
                            } else {
                                fmt.Println(v13, v1)
                            }
                        case "foo bar":
                            var v20 float64 = 0.0
                            _ = v20
                        case "hello":
                            for fi21 := 0; fi21 < 3; fi21++ {
                                fmt.Println(v13)
                                v13 -= v13
                                fmt.Println(v1, v3)
                                _ = fi21
                            }
                        default:
                            switch v14 {
                            case "":
                                var v22 uint64 = 18446744073709551615
                                _ = v22
                            case "hello":
                                if true {
                                    var v23 uint64 = 706
                                    for fi24 := 0; fi24 < 6; fi24++ {
                                        if true {
                                            var v25 bool = true
                                            _ = v25
                                        } else {
                                            v1 = 1000
                                            var v26 uint = 0
                                            v23 -= (v23 << (((v26 / (((v26 - genFn0())) | 1))) & 63))
                                            _ = v26
                                        }
                                        v23 = (v23 * 4294967295)
                                        v23 -= (v23 % ((v23) | 1))
                                        _ = fi24
                                    }
                                    _ = v23
                                } else {
                                    for fi27 := 0; fi27 < 18; fi27++ {
                                        switch (fi27 - (fi27 & fi27)) {
                                        case 127:
                                            switch "\n\t" {
                                            case "hello":
                                                var v28 uint = 925
                                                _ = v28
                                            }
                                        case 2147483647:
                                            v1 = (v3 | 743)
                                        }
                                        var v29 string = v14
                                        fmt.Println(v29, v14, v14)
                                        _ = fi27
                                        _ = v29
                                    }
                                    var v30 uint32 = 42
                                    _ = v30
                                }
                            }
                        }
                        if false {
                            var v31 uint64 = 42
                            fmt.Println(fi15, v13, v31)
                            _ = v31
                        } else {
                            fmt.Println(v13, v13)
                            v32 := GenS2{F0: 1}
                            v33 := GenS2{F0: (v1 & v1)}
                            _ = v32
                            _ = v33
                        }
                        var v34 uint32 = (v13 + 0)
                        _ = v34
                    } else {
                        var v35 uint32 = (v13 ^ v13)
                        _ = v35
                    }
                    var v36 int8 = 1
                    _ = v36
                }
                v14 = v14
                for fi37 := 0; fi37 < 20; fi37++ {
                    switch (v14 + (v14 + (("hello" + v14) + "foo bar"))) {
                    case "a":
                        v13 = (v13 / ((1) | 1))
                        if false {
                            var v38 int8 = 42
                            var v39 uint = 255
                            switch (fi15 / (((fi15 + (fi15 << ((18446744073709551615) & 63)))) | 1)) {
                            case 32768:
                                v40 := []int{-268}
                                _ = v40
                            default:
                                var v41 int8 = v38
                                _ = v41
                            }
                            _ = v38
                            _ = v39
                        }
                        if (1 == 42) {
                            var v42 uint32 = 2147483647
                            v3 -= (v3 >> ((238) & 63))
                            _ = v42
                        } else {
                            v13 += 4294967295
                            v43 := GenS1{F0: 20, F1: v14}
                            _ = v43
                        }
                    case "":
                        v3 = (-(v3))
                    }
                    v1 = 1000
                    v44 := GenS1{F0: 42, F1: "a"}
                    _ = fi37
                    _ = v44
                }
                _ = fi15
            }
            _ = v13
            _ = v14
        }
    case "a":
        if false {
            v3 = 694
            var v45 float32 = 1e10
            _ = v45
        }
    }
    var v46 uint64 = 820
    var v47 int8 = -70
    switch 32768 {
    case 127:
        var v48 uint16 = 0
        _ = v48
    case -783:
        fmt.Println(v3, v47)
    case -1:
        for fi49 := 0; fi49 < 9; fi49++ {
            fmt.Println(fi49, v3)
            if false {
                var v50 float64 = 0.0
                v51 := GenS1{F0: 0, F1: ("\n\t" + ("hello" + "foo bar"))}
                var v52 int8 = int8(fi49)
                _ = v50
                _ = v51
                _ = v52
            } else {
                for fi53 := 0; fi53 < 7; fi53++ {
                    v46 += (v46 - (v46 ^ (v46 + v46)))
                    _ = fi53
                }
            }
            fmt.Println(v47, v3)
            _ = fi49
        }
    }
    v47 = v47
    for fi54 := 0; fi54 < 19; fi54++ {
        if false {
            v55 := GenS2{F0: (-(v3))}
            v46 = (v46 << ((65535) & 63))
            for fi56 := 0; fi56 < 19; fi56++ {
                v57 := v2[0]
                _ = fi56
                _ = v57
            }
            _ = v55
        }
        for fi58 := 0; fi58 < 8; fi58++ {
            v59 := []int{fi58, int(v46)}
            switch "" {
            case "a":
                if true {
                    fmt.Println(fi58, v46)
                }
            case "foo bar":
                fmt.Println(fi58, fi54)
            case "hello":
                fmt.Println(v46, fi58)
            }
            var v60 uint64 = v46
            switch "hello" {
            case "":
                if (true && (!false)) {
                    if true {
                        v2[fi54] = 245
                        v60 -= v60
                    }
                }
            case "\n\t":
                var v61 uint16 = 255
                _ = v61
            default:
                v60 += v46
            }
            v46 = v46
            v62 := GenS3{F0: 2, F1: (true && false), F2: ((-(v1)) < (-(v3)))}
            _ = fi58
            _ = v59
            _ = v60
            _ = v62
        }
        if true {
            switch ("hello" + ("a" + ("foo bar" + "foo bar"))) {
            case "":
                for fi63 := 0; fi63 < 11; fi63++ {
                    v64 := []int{(-(fi63)), fi63}
                    fmt.Println(v3, v3)
                    _ = fi63
                    _ = v64
                }
            case "\n\t":
                var v65 uint64 = v46
                _ = v65
            case "foo bar":
                v47 = (^(v47))
            default:
                for fi66 := 0; fi66 < 5; fi66++ {
                    var v67 uint16 = 0
                    var v68 uint16 = 32767
                    if ((v68 / (((v68 << ((313) & 63))) | 1)) >= uint16(v46)) {
                        if true {
                            fmt.Println(v68, v68, v47)
                            v68 = (v68 ^ v67)
                            switch "" {
                            case "foo bar":
                                v68 = (v67 | v67)
                            case "":
                                if ("" <= "") {
                                    if false {
                                        switch (^(fi54)) {
                                        case 128:
                                            _ = len(v2)
                                        default:
                                            fmt.Println(fi66, v68, v68)
                                        }
                                        v67 = (v68 | 351)
                                        v69 := GenS1{F0: v47, F1: "foo bar"}
                                        _ = v69
                                    } else {
                                        fmt.Println(v68)
                                        var v70 int64 = 875
                                        _ = v70
                                    }
                                    fmt.Println(v46, v3, fi66)
                                }
                                var v71 int16 = v1
                                _ = v71
                            case "\n\t":
                                v72 := GenS1{F0: (v47 >> ((255) & 63)), F1: ("foo bar" + "hello")}
                                _ = v72
                            }
                        }
                        var v73 float32 = 1e10
                        switch 128 {
                        case 2147483647:
                            v74 := []int{}
                            _ = v74
                        case 4294967295:
                            var v75 int64 = -475
                            _ = v75
                        }
                        _ = v73
                    } else {
                        v67 = 0
                        var v76 float64 = (((0.1 * 1.5) / 1.0) / 1e10)
                        _ = v76
                    }
                    _ = fi66
                    _ = v67
                    _ = v68
                }
            }
        }
        _ = fi54
    }
    if true {
        for fi77 := 0; fi77 < 20; fi77++ {
            v78 := GenS3{F0: 2, F1: false, F2: (("hello" + "\n\t") < ("a" + (("hello" + "") + (("" + ((("a" + ("foo bar" + ("" + ((("foo bar" + (("a" + "foo bar") + "foo bar")) + "foo bar") + "")))) + ("" + ("a" + "foo bar"))) + "hello")) + ("foo bar" + "\n\t")))))}
            v47 -= (v47 | v47)
            var v79 uint32 = 2
            _ = fi77
            _ = v78
            _ = v79
        }
        if false {
            switch 4294967295 {
            case 4294967295:
                v47 += (v47 << ((706) & 63))
            default:
                v1 = (v1 >> ((genFn0()) & 63))
            }
            var v80 uint16 = 65535
            _ = v80
        }
    }
    v46 -= v46
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v46
    _ = v47
    _ = v8
}
