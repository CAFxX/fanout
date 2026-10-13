package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 string; F1 uint8 }
func genFn0(p0 int16, p1 uint64, p2 bool) uint32 {
    fmt.Println(p0)
    var v6 uint8 = 249
    p1 = p1
    var v7 float32 = 0.1
    if p2 {
        p1 += p1
        var v8 int16 = p0
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 5
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"\n\t", "a", ""}
    var v3 uint = 195
    v4 := make(map[string]uint32)
    var v5 string = "a"
    var v9 uint8 = 1
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 int8 = 0
        fmt.Println(v3, v11, fi10)
        v12 := GenS1{F0: 32767, F1: 1, F2: 1}
        _ = fi10
        _ = v11
        _ = v12
    }
    v9 -= 1
    v3 = v3
    _ = len(v2)
    _ = len(v2)
    if (9223372036854775807 < 42) {
        var v13 int16 = 1000
        _ = len(v2)
        for fi14 := 0; fi14 < 10; fi14++ {
            v4[v5] = 2
            if len(v2) > 0 {
                v2[0] = "hello"
            }
            fmt.Println(fi14, v13, fi14)
            _ = fi14
        }
        _ = v13
    }
    var v15 int8 = 42
    for fi16 := 0; fi16 < 3; fi16++ {
        v3 = (^(v3))
        if (v5 == ("hello" + "foo bar")) {
            for fi17 := 0; fi17 < 17; fi17++ {
                v2 = append(v2, ((v5 + ((("hello" + v5) + (("foo bar" + v5) + "\n\t")) + ("hello" + ""))) + v5))
                v18 := v1[v5]
                var v19 uint16 = 2
                var v20 int16 = int16(v18)
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
        }
        _ = fi16
    }
    if true {
        for fi21 := 0; fi21 < 16; fi21++ {
            for fi22 := 0; fi22 < 18; fi22++ {
                if len(v2) > 0 {
                    v2[0] = v5
                }
                var v23 uint16 = 1
                var v24 int16 = -573
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = fi21
        }
        delete(v1, v5)
        if false {
            var v25 int = 128
            _ = v25
        } else {
            for fi26 := 0; fi26 < 16; fi26++ {
                var v27 int32 = -1
                v28 := GenS1{F0: 499, F1: 1, F2: 42}
                _ = fi26
                _ = v27
                _ = v28
            }
        }
    }
    if false {
        var v29 uint16 = 24
        for fi30 := 0; fi30 < 5; fi30++ {
            if true {
                v31 := GenS2{F0: (v3 / (((^(v3))) | 1))}
                v15 += v15
                if (false || (false || true)) {
                    v5 = v5
                } else {
                    switch v5 {
                    case "hello":
                        var v32 int32 = 1
                        _ = v32
                    case "a":
                        v15 += v15
                    case "foo bar":
                        var v33 int32 = -1
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 3; fi34++ {
                        if false {
                            var v35 int8 = 0
                            _ = v35
                        } else {
                            var v36 uint8 = v9
                            v37 := GenS1{F0: v29, F1: -9223372036854775808, F2: 977}
                            _ = v36
                            _ = v37
                        }
                        _ = fi34
                    }
                }
                _ = v31
            } else {
                v3 += 593
                switch int(v29) {
                case -403:
                    for fi38 := 0; fi38 < 8; fi38++ {
                        var v39 uint64 = 18446744073709551615
                        _ = fi38
                        _ = v39
                    }
                default:
                    if false {
                        var v40 int8 = (^(v15))
                        v29 = (v29 % ((99) | 1))
                        if len(v2) > 0 {
                            v2[0] = "\n\t"
                        }
                        _ = v40
                    } else {
                        var v41 int8 = int8(v3)
                        for fi42 := 0; fi42 < 7; fi42++ {
                            v29 -= 0
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi42
                        }
                        v43 := GenS3{F0: 2, F1: 174}
                        _ = v41
                        _ = v43
                    }
                }
            }
            if true {
                for fi44 := 0; fi44 < 8; fi44++ {
                    var v45 float32 = -2.25
                    if (65535 <= v3) {
                        var v46 float64 = 1.5
                        _ = v46
                    } else {
                        v47 := GenS4{F0: (v5 + ("\n\t" + ("a" + "foo bar"))), F1: v9}
                        switch v5 {
                        case "\n\t":
                            switch fi30 {
                            case -1:
                                var v48 float32 = -2.25
                                _ = v48
                            case 32768:
                                if ((v3 % (((^(v3))) | 1)) == v3) {
                                    _ = len(v1)
                                }
                            }
                        case "":
                            fmt.Println(v5, v15)
                        }
                        v9 -= (^(v9))
                        _ = v47
                    }
                    v29 += 65535
                    _ = fi44
                    _ = v45
                }
                v29 += 32767
                switch fi30 {
                case -338:
                    switch "" {
                    case "foo bar":
                        fmt.Println(v3)
                    default:
                        if true {
                            for fi49 := 0; fi49 < 3; fi49++ {
                                var v50 uint16 = v29
                                _ = fi49
                                _ = v50
                            }
                            var v51 int8 = v15
                            _ = v51
                        }
                    }
                case 128:
                    v52 := GenS4{F0: v5, F1: (v9 & uint8(v29))}
                    _ = v52
                case 2147483647:
                    _ = len(v1)
                default:
                    fmt.Println(v15, fi30)
                }
            }
            fmt.Println(fi30, v5, v5)
            _ = fi30
        }
        var v53 int32 = int32(v15)
        _ = v29
        _ = v53
    } else {
        if (!false) {
            v15 = (^(v15))
        } else {
            switch 32768 {
            case 224:
                if (127 >= 0) {
                    for fi54 := 0; fi54 < 4; fi54++ {
                        var v55 bool = ((0.1 < (3.14159 + (1.5 * -2.25))) || (1 != 92))
                        v56 := v4[("\n\t" + v5)]
                        _ = fi54
                        _ = v55
                        _ = v56
                    }
                } else {
                    fmt.Println(v3, v3)
                    for fi57 := 0; fi57 < 3; fi57++ {
                        for fi58 := 0; fi58 < 16; fi58++ {
                            var v59 string = ((v5 + ("\n\t" + v5)) + v5)
                            var v60 int = fi57
                            _ = fi58
                            _ = v59
                            _ = v60
                        }
                        switch v5 {
                        case "\n\t":
                            fmt.Println(v15, v15, v15)
                        case "hello":
                            v61 := GenS4{F0: ((v5 + v5) + (v5 + ((v5 + v5) + "\n\t"))), F1: v9}
                            _ = v61
                        case "":
                            if ((false && (255 < 821)) || true) {
                                var v62 bool = false
                                _ = v62
                            } else {
                                fmt.Println(v15, fi57)
                            }
                        default:
                            switch fi57 {
                            case -1:
                                v63 := GenS1{F0: 65535, F1: -9223372036854775808, F2: 408}
                                _ = v63
                            case -962:
                                var v64 int32 = 1
                                _ = v64
                            }
                        }
                        _ = fi57
                    }
                    for fi65 := 0; fi65 < 10; fi65++ {
                        v66 := GenS3{F0: 42, F1: v9}
                        v67 := v1["hello"]
                        _ = fi65
                        _ = v66
                        _ = v67
                    }
                }
            }
            for fi68 := 0; fi68 < 3; fi68++ {
                v5 = v5
                _ = len(v2)
                for fi69 := 0; fi69 < 12; fi69++ {
                    v70 := GenS1{F0: 0, F1: 9223372036854775807, F2: 398}
                    fmt.Println(v5, fi68)
                    if false {
                        var v71 int64 = -761
                        v72 := v4[v5]
                        v72 -= (v72 / (((^(v72))) | 1))
                        _ = v71
                        _ = v72
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v73 int64 = 42
                        v73 += (^(v73))
                        if len(v2) > 0 {
                            v2[0] = "foo bar"
                        }
                        var v74 float32 = 3.14159
                        _ = v73
                        _ = v74
                    }
                    _ = fi69
                    _ = v70
                }
                _ = fi68
            }
            switch 2147483647 {
            case 9223372036854775807:
                v9 = (v9 >> ((v3) & 63))
            case -2147483648:
                v75 := GenS2{F0: 675}
                _ = v75
            default:
                var v76 int8 = (v15 ^ 1)
                _ = v76
            }
        }
        var v77 int8 = v15
        _ = v77
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
