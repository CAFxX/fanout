package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 string; F2 string }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64; F1 string }
type GenS7 struct { F0 int; F1 int32 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 string; F1 uint32 }
type GenS10 struct { F0 string }
func genFn0(p0 uint32, p1 uint64) bool {
    if true {
        v4 := GenS1{F0: int(p1), F1: (p1 % ((p1) | 1)), F2: p0}
        _ = v4
    }
    var v5 uint8 = 42
    v6 := GenS1{F0: -996, F1: 42, F2: p0}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (!false)
}
func genFn1(p0 int, p1 int8, p2 int16) uint16 {
    fmt.Println(p1, p2, p1)
    if (!(875 == 758)) {
        v7 := GenS2{F0: true}
        if false {
            fmt.Println(p0)
        }
        _ = v7
    } else {
        var v8 string = "foo bar"
        v9 := GenS3{F0: 825, F1: "foo bar", F2: v8}
        _ = v8
        _ = v9
    }
    var v10 int8 = p1
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 int64 = -9223372036854775808
        for fi13 := 0; fi13 < 2; fi13++ {
            v12 += v12
            _ = fi13
        }
        v10 = 15
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 255
}

func main() {
    var v1 uint = 42
    var v2 uint64 = 42
    var v3 uint = 255
    var v14 int32 = -330
    var v15 int8 = -123
    v16 := []int{1, 32767, -338, 42}
    if true {
        if (!(!true)) {
            _ = len(v16)
            v3 = uint(v2)
        } else {
            for fi17 := 0; fi17 < 12; fi17++ {
                for fi18 := 0; fi18 < 12; fi18++ {
                    if false {
                        v19 := GenS2{F0: (0.1 >= 3.14159)}
                        _ = v19
                    }
                    fmt.Println(v1)
                    var v20 uint16 = 0
                    _ = fi18
                    _ = v20
                }
                switch (fi17 % ((32767) | 1)) {
                case 1:
                    var v21 uint16 = 2
                    _ = v21
                default:
                    switch (fi17 ^ fi17) {
                    case 719:
                        var v22 uint = v3
                        _ = v22
                    case 660:
                        if ((-2.25 * 0.1) < (1e10 - (3.14159 - 0.0))) {
                            for fi23 := 0; fi23 < 18; fi23++ {
                                var v24 int16 = 1000
                                _ = fi23
                                _ = v24
                            }
                            if (-32768 < 1) {
                                fmt.Println(v15)
                                v2 = v2
                                switch "\n\t" {
                                case "\n\t":
                                    fmt.Println(v15, v2, v15)
                                }
                            }
                            if len(v16) > 0 {
                                v16[0] = fi17
                            }
                        }
                    default:
                        v25 := GenS3{F0: (v1 / ((1) | 1)), F1: "a", F2: "\n\t"}
                        _ = v25
                    }
                }
                switch "hello" {
                case "a":
                    var v26 string = "hello"
                    _ = v26
                case "\n\t":
                    v15 = (v15 | v15)
                }
                _ = fi17
            }
        }
        var v27 uint8 = 2
        _ = v27
    }
    v28 := GenS4{F0: 42, F1: 1}
    for fi29 := 0; fi29 < 10; fi29++ {
        v16 = append(v16, (^(fi29)))
        switch 4294967295 {
        case 128:
            if false {
                fmt.Println(v2)
                if len(v16) > 0 {
                    v16[0] = (fi29 & (fi29 & fi29))
                }
            }
        case -9223372036854775808:
            v30 := GenS5{F0: (v2 | 18446744073709551615)}
            _ = v30
        default:
            v31 := GenS6{F0: 0, F1: ""}
            _ = v31
        }
        var v32 uint16 = 0
        _ = fi29
        _ = v32
    }
    v1 += 18446744073709551615
    if false {
        for fi33 := 0; fi33 < 3; fi33++ {
            v15 += (v15 / (((v15 - (v15 | 0))) | 1))
            v34 := GenS6{F0: -9223372036854775808, F1: ""}
            _ = fi33
            _ = v34
        }
        v35 := GenS6{F0: 469, F1: "hello"}
        switch (("hello" + "\n\t") + ("hello" + "")) {
        case "\n\t":
            v2 -= (v2 * 616)
        case "foo bar":
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        case "":
            switch 32768 {
            case 9223372036854775807:
                var v36 uint32 = 2
                _ = v36
            case 2147483647:
                switch 4294967295 {
                case 42:
                    switch 475 {
                    case 2147483647:
                        v2 -= 1
                    case 128:
                        v37 := GenS1{F0: 9223372036854775807, F1: v2, F2: 4294967295}
                        _ = v37
                    case -2147483648:
                        _ = len(v16)
                    }
                }
            case -423:
                for fi38 := 0; fi38 < 14; fi38++ {
                    v2 = v2
                    if (!(!(true || true))) {
                        switch (fi38 - 32768) {
                        case 127:
                            v39 := GenS7{F0: 128, F1: v14}
                            _ = v39
                        case 1:
                            switch (^(fi38)) {
                            case -9223372036854775808:
                                for fi40 := 0; fi40 < 5; fi40++ {
                                    if len(v16) >= 2 {
                                        v16 = v16[1:len(v16)]
                                    }
                                    _ = len(v16)
                                    _ = fi40
                                }
                            case 434:
                                _ = len(v16)
                            }
                        }
                        switch (fi38 - -2147483648) {
                        case -703:
                            for fi41 := 0; fi41 < 12; fi41++ {
                                v1 = (v3 ^ (v3 - 42))
                                v2 -= 2
                                v3 = v3
                                _ = fi41
                            }
                        case 2147483647:
                            v42 := GenS8{F0: (v15 << ((v3) & 63))}
                            _ = v42
                        case 32768:
                            _ = len(v16)
                        default:
                            var v43 int32 = 569
                            _ = v43
                        }
                    } else {
                        fmt.Println(v2)
                        for fi44 := 0; fi44 < 17; fi44++ {
                            v3 += v1
                            v45 := GenS9{F0: "", F1: 670}
                            switch ("foo bar" + ("foo bar" + (("a" + "\n\t") + "hello"))) {
                            case "a":
                                var v46 string = "hello"
                                _ = v46
                            default:
                                var v47 int32 = (-(v14))
                                _ = v47
                            }
                            _ = fi44
                            _ = v45
                        }
                    }
                    if len(v16) > 0 {
                        v16[0] = (fi38 - 127)
                    }
                    _ = fi38
                }
            default:
                v3 += (v1 % ((v3) | 1))
            }
        default:
            for fi48 := 0; fi48 < 3; fi48++ {
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                for fi49 := 0; fi49 < 8; fi49++ {
                    v50 := GenS1{F0: fi49, F1: (v2 - (v2 | v2)), F2: 42}
                    fmt.Println(v3)
                    _ = fi49
                    _ = v50
                }
                var v51 int8 = -128
                _ = fi48
                _ = v51
            }
        }
        _ = v35
    }
    v52 := GenS9{F0: ("foo bar" + ("" + "a")), F1: uint32(v2)}
    if true {
        v2 = v2
    } else {
        v16 = append(v16, int(v14))
        v53 := GenS10{F0: "foo bar"}
        _ = v53
    }
    for fi54 := 0; fi54 < 10; fi54++ {
        var v55 int32 = v14
        if len(v16) > 0 {
            v16[0] = fi54
        }
        var v56 int32 = (v55 ^ v14)
        _ = fi54
        _ = v55
        _ = v56
    }
    for fi57 := 0; fi57 < 3; fi57++ {
        fmt.Println(fi57)
        if len(v16) > 0 {
            v16[0] = fi57
        }
        _ = fi57
    }
    fmt.Println(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v52
}
