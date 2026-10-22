package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int32 }
type GenS2 struct { F0 int8; F1 uint; F2 int64 }
type GenS3 struct { F0 bool; F1 string; F2 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 int32; F2 int8 }

func main() {
    var v1 float32 = 1e10
    var v2 uint = 524
    v3 := make(map[string]uint32)
    var v4 uint64 = 301
    v5 := GenS1{F0: 982, F1: 100, F2: 42}
    v6 := GenS2{F0: -50, F1: v2, F2: 9223372036854775807}
    if true {
        v7 := GenS2{F0: 127, F1: (^(v2)), F2: 0}
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := []int{fi8, fi8, fi8}
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 7; fi10++ {
            switch -9223372036854775808 {
            case 0:
                v11 := []int{-687, fi10, fi10, (^(fi10))}
                _ = v11
            }
            v4 -= v4
            if false {
                var v12 float32 = 0.0
                v2 -= (v2 - (v2 >> ((42) & 63)))
                _ = v12
            }
            _ = fi10
        }
        _ = v7
    } else {
        switch "foo bar" {
        case "foo bar":
            v13 := GenS3{F0: false, F1: "hello", F2: 1}
            _ = v13
        default:
            if (9223372036854775807 > -9223372036854775808) {
                switch 83 {
                case -636:
                    for fi14 := 0; fi14 < 4; fi14++ {
                        for fi15 := 0; fi15 < 1; fi15++ {
                            switch "\n\t" {
                            case "a":
                                var v16 uint64 = 1
                                _ = v16
                            }
                            _ = fi15
                        }
                        _ = fi14
                    }
                case 9223372036854775807:
                    switch "hello" {
                    case "":
                        switch "" {
                        case "hello":
                            if (uint32(v4) != 790) {
                                for fi17 := 0; fi17 < 19; fi17++ {
                                    v2 = v2
                                    for fi18 := 0; fi18 < 2; fi18++ {
                                        var v19 int64 = 1
                                        if ((0.0 < -2.25) && false) {
                                            if (true || (!((uint(v4) == v2) && (((1.5 + 3.14159) / -2.25) >= (((3.14159 / 1e10) - (-2.25 * 1e10)) + 3.14159))))) {
                                                if false {
                                                    var v20 uint16 = 904
                                                    if false {
                                                        var v21 int64 = v19
                                                        v1 = float32((1e10 / (-2.25 * 3.14159)))
                                                        v1 -= v1
                                                        _ = v21
                                                    }
                                                    var v22 int16 = -144
                                                    _ = v20
                                                    _ = v22
                                                }
                                                if true {
                                                    var v23 int = -9223372036854775808
                                                    if (272 < v2) {
                                                        if (!(-2.25 != 1e10)) {
                                                            fmt.Println(fi18, fi18, v23)
                                                        }
                                                    } else {
                                                        fmt.Println(v19, v4)
                                                    }
                                                    _ = v23
                                                } else {
                                                    var v24 bool = true
                                                    _ = v24
                                                }
                                            }
                                        }
                                        if false {
                                            v1 = (v1 - v1)
                                            v19 = v19
                                            v25 := []int{}
                                            _ = v25
                                        }
                                        _ = fi18
                                        _ = v19
                                    }
                                    v26 := GenS1{F0: 1, F1: 127, F2: 2147483647}
                                    _ = fi17
                                    _ = v26
                                }
                            } else {
                                switch "foo bar" {
                                case "\n\t":
                                    v4 += v4
                                case "foo bar":
                                    var v27 int8 = -7
                                    _ = v27
                                case "":
                                    fmt.Println(v1, v2, v4)
                                }
                            }
                        case "\n\t":
                            v2 = v2
                        case "a":
                            v28 := GenS4{F0: uint8(v2)}
                            _ = v28
                        }
                    case "hello":
                        for fi29 := 0; fi29 < 11; fi29++ {
                            fmt.Println(v4, v1)
                            if (false && (0.0 < 3.14159)) {
                                for fi30 := 0; fi30 < 19; fi30++ {
                                    v31 := []int{(fi29 & fi29), (fi30 * int(v2))}
                                    v32 := GenS5{F0: (v2 - (v2 * 95)), F1: 724, F2: -44}
                                    _ = fi30
                                    _ = v31
                                    _ = v32
                                }
                            } else {
                                var v33 uint = (^(v2))
                                fmt.Println(v2)
                                var v34 bool = false
                                _ = v33
                                _ = v34
                            }
                            _ = fi29
                        }
                    case "a":
                        var v35 uint64 = 2
                        _ = v35
                    }
                default:
                    v1 = 3.14159
                }
            } else {
                if (-2147483648 < -351) {
                    for fi36 := 0; fi36 < 1; fi36++ {
                        var v37 string = "foo bar"
                        _ = fi36
                        _ = v37
                    }
                }
            }
        }
        fmt.Println(v4, v4)
        var v38 uint32 = 713
        _ = v38
    }
    for fi39 := 0; fi39 < 13; fi39++ {
        if (false || (255 != 780)) {
            for fi40 := 0; fi40 < 15; fi40++ {
                v41 := GenS3{F0: false, F1: ("hello" + ""), F2: fi39}
                if false {
                    v42 := GenS4{F0: 1}
                    _ = v42
                }
                v43 := GenS2{F0: 42, F1: (v2 << ((v2) & 63)), F2: -834}
                _ = fi40
                _ = v41
                _ = v43
            }
            var v44 string = (((("a" + (("a" + ("" + (("foo bar" + "\n\t") + "a"))) + "")) + ("foo bar" + "hello")) + "hello") + "")
            if true {
                v45 := GenS5{F0: v2, F1: 2147483647, F2: 1}
                fmt.Println(fi39, v4)
                if true {
                    v46 := GenS3{F0: (255 == 127), F1: ("" + "a"), F2: -208}
                    fmt.Println(v44, v44, v44)
                    if (((!false) && true) && false) {
                        switch 128 {
                        case 202:
                            for fi47 := 0; fi47 < 17; fi47++ {
                                v48 := GenS4{F0: 42}
                                var v49 uint = (v2 ^ (v2 * 0))
                                _ = fi47
                                _ = v48
                                _ = v49
                            }
                        case 0:
                            v44 = v44
                        }
                    }
                    _ = v46
                } else {
                    v50 := GenS1{F0: -792, F1: 42, F2: 42}
                    var v51 int = (fi39 + (fi39 - 1))
                    if true {
                        for fi52 := 0; fi52 < 3; fi52++ {
                            v53 := []int{fi52, -631, (^(v51))}
                            v54 := v3[(v44 + ("hello" + v44))]
                            _ = fi52
                            _ = v53
                            _ = v54
                        }
                    }
                    _ = v50
                    _ = v51
                }
                _ = v45
            } else {
                var v55 uint64 = (^(v4))
                var v56 int = -9223372036854775808
                fmt.Println(v56)
                _ = v55
                _ = v56
            }
            _ = v44
        } else {
            v4 += (v4 >> ((v2) & 63))
            fmt.Println(v1, v1)
        }
        _ = fi39
    }
    var v57 float32 = v1
    var v58 int8 = -78
    fmt.Println(v2, v4)
    fmt.Println(v1, v4)
    var v59 bool = true
    v4 += v4
    var v60 float64 = float64(v57)
    v61 := v3[(("hello" + (("hello" + ("foo bar" + ("\n\t" + ((("foo bar" + "\n\t") + ("" + "hello")) + "\n\t")))) + "hello")) + "\n\t")]
    var v62 int64 = -740
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v62
}
