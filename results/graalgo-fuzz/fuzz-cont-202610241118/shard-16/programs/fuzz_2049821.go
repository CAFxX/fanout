package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 string }
type GenS2 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 int16; F1 int8; F2 int64 }
type GenS5 struct { F0 int64; F1 int64; F2 int8 }

func main() {
    var v1 uint = 2
    var v2 string = ("hello" + ("a" + ("hello" + "")))
    var v3 string = (v2 + (v2 + v2))
    var v4 int16 = 32767
    v5 := make(map[string]string)
    v2 = v3
    switch 127 {
    case 32767:
        var v6 float64 = (-2.25 / float64((0.0 - ((((3.14159 + 0.1) / -2.25) * float32(1.5)) - (3.14159 - (float32(0.0) - -2.25))))))
        _ = v6
    default:
        for fi7 := 0; fi7 < 3; fi7++ {
            for fi8 := 0; fi8 < 2; fi8++ {
                var v9 uint = (v1 % (((v1 >> (((v1 & 809)) & 63))) | 1))
                _ = len(v5)
                _ = fi8
                _ = v9
            }
            if true {
                var v10 int64 = 1
                if (("" + (v2 + (v2 + v3))) > ("" + ("foo bar" + (((("\n\t" + v3) + "\n\t") + (v2 + v3)) + v2)))) {
                    var v11 int16 = int16(v10)
                    v12 := []int{(fi7 / ((fi7) | 1)), (fi7 % ((fi7) | 1))}
                    _ = len(v5)
                    _ = v11
                    _ = v12
                } else {
                    v13 := GenS1{F0: 65535, F1: ("foo bar" + "a"), F2: ("a" + v3)}
                    _ = v13
                }
                _ = v10
            } else {
                v14 := GenS1{F0: 0, F1: "hello", F2: (("a" + v2) + v2)}
                v15 := GenS1{F0: 65535, F1: (((v2 + v3) + ((v2 + (v2 + v2)) + v3)) + "a"), F2: (v2 + "\n\t")}
                _ = v14
                _ = v15
            }
            _ = fi7
        }
    }
    fmt.Println(v4, v2, v3)
    v16 := []int{32768, -435}
    if false {
        for fi17 := 0; fi17 < 2; fi17++ {
            var v18 int64 = 0
            for fi19 := 0; fi19 < 20; fi19++ {
                var v20 uint8 = 42
                _ = len(v16)
                v21 := v5[v3]
                _ = fi19
                _ = v20
                _ = v21
            }
            v22 := GenS1{F0: 0, F1: "foo bar", F2: v2}
            _ = fi17
            _ = v18
            _ = v22
        }
        fmt.Println(v1)
        switch v2 {
        case "hello":
            for fi23 := 0; fi23 < 17; fi23++ {
                for fi24 := 0; fi24 < 12; fi24++ {
                    v4 = v4
                    switch (v3 + v2) {
                    case "hello":
                        v25 := GenS1{F0: 4294967295, F1: v3, F2: "hello"}
                        _ = v25
                    case "a":
                        switch (fi23 << ((65535) & 63)) {
                        case 1:
                            if (42 == 4294967295) {
                                if ((v1 * 18446744073709551615) == v1) {
                                    v26 := GenS1{F0: uint32(v1), F1: (v2 + (v2 + v2)), F2: ("\n\t" + (v3 + (v3 + ((v3 + v2) + "a"))))}
                                    v5[(v2 + "a")] = "foo bar"
                                    v27 := GenS1{F0: 808, F1: ("\n\t" + ("a" + ("" + "a"))), F2: v3}
                                    v28 := GenS1{F0: 1, F1: v3, F2: v3}
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                } else {
                                    v4 = v4
                                }
                                if true {
                                    if true {
                                        v3 = v2
                                        var v29 int64 = 0
                                        if ((v29 != (v29 & (-(v29)))) || ((-1 < (^(v4))) || ((!(((0.0 + 3.14159) / -2.25) == -2.25)) || false))) {
                                            switch v3 {
                                            case "":
                                                var v30 string = "a"
                                                _ = v30
                                            }
                                            v29 -= v29
                                        }
                                        if len(v16) >= 2 {
                                            v16 = v16[1:len(v16)]
                                        }
                                        _ = v29
                                    } else {
                                        v31 := GenS2{F0: 2147483648, F1: (^(v1)), F2: 114}
                                        var v32 string = v3
                                        _ = v31
                                        _ = v32
                                    }
                                }
                                fmt.Println(fi24)
                                v2 = "\n\t"
                                _ = len(v5)
                                v33 := v5["a"]
                                v33 = ""
                                v4 += (v4 | (v4 ^ int16(fi23)))
                                var v34 int32 = -627
                                _ = v33
                                _ = v34
                            } else {
                                v1 = v1
                            }
                        }
                        v35 := GenS2{F0: 945, F1: v1, F2: 2}
                        switch (fi23 + (fi23 ^ 2147483647)) {
                        case 499:
                            var v36 int16 = 0
                            _ = v36
                        case 1:
                            if (false && true) {
                                var v37 string = v3
                                switch 9223372036854775807 {
                                case -1:
                                    if true {
                                        v38 := GenS1{F0: 2, F1: "", F2: v37}
                                        var v39 uint = (v1 % ((v1) | 1))
                                        v3 = "a"
                                        _ = v38
                                        _ = v39
                                    }
                                case -190:
                                    v3 = v37
                                case 9223372036854775807:
                                    if (0.0 <= 1.5) {
                                        switch -25 {
                                        case 127:
                                            v37 = v2
                                        case 42:
                                            switch ((v3 + v3) + v2) {
                                            case "a":
                                                _ = len(v16)
                                            case "foo bar":
                                                fmt.Println(v37)
                                            case "\n\t":
                                                fmt.Println(v4)
                                            }
                                        default:
                                            var v40 uint32 = 65535
                                            var v41 uint64 = 2
                                            _ = v40
                                            _ = v41
                                        }
                                        v5["\n\t"] = v2
                                        fmt.Println(v4, v4, v37)
                                        fmt.Println(fi24, v1, v1)
                                    }
                                }
                                v37 = v2
                                _ = v37
                            } else {
                                var v42 int16 = v4
                                var v43 uint8 = 42
                                var v44 uint64 = 2
                                var v45 uint8 = (v43 + v43)
                                fmt.Println(v4)
                                _ = v42
                                _ = v43
                                _ = v44
                                _ = v45
                            }
                        case -2147483648:
                            fmt.Println(v1, v4)
                        default:
                            v46 := GenS3{F0: 42, F1: (^(v4))}
                            _ = v46
                        }
                        fmt.Println(fi23, v4)
                        _ = v35
                    case "\n\t":
                        v4 += (v4 >> ((18446744073709551615) & 63))
                    }
                    var v47 int64 = 9223372036854775807
                    _ = fi24
                    _ = v47
                }
                var v48 int64 = 42
                _ = fi23
                _ = v48
            }
        case "foo bar":
            v4 = v4
        default:
            fmt.Println(v1, v4)
        }
    } else {
        for fi49 := 0; fi49 < 10; fi49++ {
            v50 := GenS1{F0: 0, F1: v2, F2: (v2 + v3)}
            if (v3 == (v2 + ("hello" + (v3 + "foo bar")))) {
                var v51 int8 = -80
                v52 := GenS4{F0: (v4 << ((839) & 63)), F1: v51, F2: -1}
                _ = v51
                _ = v52
            }
            _ = fi49
            _ = v50
        }
        switch v3 {
        case "a":
            var v53 uint64 = 1
            _ = v53
        default:
            var v54 uint16 = 39
            _ = v54
        }
        var v55 int8 = -8
        _ = v55
    }
    v2 = v2
    if false {
        v56 := GenS3{F0: 1, F1: (v4 % ((v4) | 1))}
        _ = v56
    } else {
        for fi57 := 0; fi57 < 19; fi57++ {
            var v58 string = v3
            for fi59 := 0; fi59 < 12; fi59++ {
                v2 = (v58 + (v3 + ((v2 + ((v3 + (v2 + (v3 + v58))) + "a")) + ("foo bar" + v58))))
                var v60 uint64 = 1
                switch (v2 + v2) {
                case "hello":
                    v58 = v3
                case "":
                    if true {
                        var v61 int8 = -65
                        _ = v61
                    }
                case "a":
                    v1 = v1
                default:
                    var v62 int8 = 100
                    _ = v62
                }
                if (42 >= v4) {
                    v63 := GenS2{F0: 4294967295, F1: v1, F2: 2147483647}
                    v64 := v5[v3]
                    _ = v63
                    _ = v64
                }
                if ((v4 | (v4 >> ((v1) & 63))) < (v4 >> (((^(v1))) & 63))) {
                    v60 += 42
                } else {
                    v16 = append(v16, 1)
                    var v65 int = fi57
                    _ = v65
                }
                v66 := v5[((v58 + ("\n\t" + v58)) + (v58 + (v3 + v3)))]
                _ = fi59
                _ = v60
                _ = v66
            }
            _ = fi57
            _ = v58
        }
        delete(v5, v2)
        if (127 == 205) {
            v67 := GenS1{F0: 2, F1: "\n\t", F2: v2}
            for fi68 := 0; fi68 < 2; fi68++ {
                v5[v2] = v2
                if (-1 == (^(v4))) {
                    fmt.Println(v1, v1)
                }
                _ = fi68
            }
            v5[(v3 + v2)] = ("foo bar" + "\n\t")
            _ = v67
        } else {
            v3 = ((v2 + ((v3 + "hello") + "hello")) + v3)
        }
    }
    v69 := GenS2{F0: 0, F1: (^(v1)), F2: 0}
    if false {
        v2 = ("" + "")
        _ = len(v16)
    }
    var v70 uint8 = 0
    var v71 bool = (false || true)
    v16 = append(v16, -584)
    if (-671 < 777) {
        v72 := GenS5{F0: 0, F1: 1, F2: 1}
        v73 := GenS2{F0: 918, F1: 527, F2: 232}
        _ = v72
        _ = v73
    }
    if len(v16) > 0 {
        v16[0] = 32767
    }
    switch -2147483648 {
    case 0:
        v1 = v1
    case -755:
        v4 -= (v4 + 244)
    case 2147483647:
        var v74 string = "hello"
        _ = v74
    default:
        var v75 float32 = 0.0
        _ = v75
    }
    if true {
        var v76 string = v2
        v76 = v3
        var v77 uint32 = 42
        _ = v76
        _ = v77
    }
    v78 := GenS3{F0: 9223372036854775807, F1: v4}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v69
    _ = v70
    _ = v71
    _ = v78
}
