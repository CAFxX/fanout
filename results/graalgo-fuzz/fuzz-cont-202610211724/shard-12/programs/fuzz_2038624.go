package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 int64; F1 int8; F2 int }
type GenS4 struct { F0 int; F1 uint; F2 uint }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool; F1 uint; F2 int16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int64; F1 int32 }
type GenS9 struct { F0 int; F1 uint64; F2 int64 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]uint32)
    var v3 uint8 = 127
    v4 := []string{"\n\t"}
    var v5 int = 114
    var v6 int8 = 80
    for fi7 := 0; fi7 < 11; fi7++ {
        v5 = 9223372036854775807
        _ = fi7
    }
    var v8 uint8 = 1
    fmt.Println(v5, v8, v5)
    switch 32768 {
    case 9223372036854775807:
        fmt.Println(v3)
    case 0:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    switch 2147483647 {
    case 128:
        var v9 int16 = 1
        _ = v9
    case 9223372036854775807:
        if false {
            switch v5 {
            case 42:
                v10 := GenS1{F0: 42, F1: "a"}
                _ = v10
            default:
                for fi11 := 0; fi11 < 17; fi11++ {
                    switch (v5 | fi11) {
                    case 0:
                        for fi12 := 0; fi12 < 18; fi12++ {
                            fmt.Println(fi12)
                            v5 += 128
                            v8 = (^(v8))
                            _ = fi12
                        }
                    case -895:
                        var v13 float64 = (1.5 / -2.25)
                        _ = v13
                    case -9223372036854775808:
                        if ("hello" <= "") {
                            if ((((1e10 / 1.5) + -2.25) != -2.25) && false) {
                                fmt.Println(fi11)
                            }
                            v2["a"] = uint32(v3)
                        }
                    default:
                        v14 := GenS1{F0: 0, F1: ""}
                        _ = v14
                    }
                    _ = fi11
                }
            }
            switch v5 {
            case -777:
                v4 = append(v4, "\n\t")
            }
            v15 := GenS2{F0: false, F1: 721}
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        switch (fi16 + (fi16 - (v5 - fi16))) {
        case -917:
            var v17 uint = 0
            _ = v17
        case -145:
            switch fi16 {
            case 9223372036854775807:
                var v18 string = (("hello" + (("\n\t" + ("" + "\n\t")) + ("\n\t" + "foo bar"))) + ("" + "a"))
                _ = v18
            }
        default:
            _ = len(v2)
        }
        fmt.Println(fi16, v8, v6)
        _ = fi16
    }
    for fi19 := 0; fi19 < 18; fi19++ {
        switch fi19 {
        case 42:
            var v20 int = (^(v5))
            _ = v20
        case -9223372036854775808:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case 128:
            var v21 int64 = -1
            _ = v21
        }
        if false {
            switch (v5 * v5) {
            case -399:
                var v22 int8 = v6
                _ = v22
            case 694:
                v23 := GenS1{F0: -845, F1: "foo bar"}
                _ = v23
            case 128:
                v6 = (^(v6))
            default:
                v24 := GenS2{F0: true, F1: -1}
                _ = v24
            }
            var v25 int8 = v6
            for fi26 := 0; fi26 < 11; fi26++ {
                v27 := GenS3{F0: -9223372036854775808, F1: v25, F2: v5}
                if false {
                    var v28 string = "a"
                    v29 := GenS4{F0: v5, F1: 42, F2: 65535}
                    var v30 uint8 = (v3 - v8)
                    _ = len(v2)
                    v8 += 127
                    _ = v28
                    _ = v29
                    _ = v30
                } else {
                    fmt.Println(v3, v5)
                }
                if false {
                    var v31 uint16 = uint16(v3)
                    var v32 uint64 = 2
                    if len(v4) > 0 {
                        v4[0] = "foo bar"
                    }
                    _ = v31
                    _ = v32
                } else {
                    v33 := GenS2{F0: (189 >= 1), F1: 42}
                    var v34 float64 = 1.5
                    v5 = 32767
                    if true {
                        fmt.Println(fi26, v34, v3)
                        var v35 string = "\n\t"
                        _ = v35
                    } else {
                        fmt.Println(v6, v8, v5)
                        fmt.Println(fi26)
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v33
                    _ = v34
                }
                _ = fi26
                _ = v27
            }
            _ = v25
        } else {
            switch v5 {
            case -743:
                var v36 uint = 18446744073709551615
                _ = v36
            case -9223372036854775808:
                v3 = 128
            case 42:
                if true {
                    switch v5 {
                    case 4294967295:
                        if true {
                            for fi37 := 0; fi37 < 1; fi37++ {
                                var v38 bool = (!(v8 >= v8))
                                switch "foo bar" {
                                case "foo bar":
                                    var v39 uint8 = v8
                                    _ = v39
                                case "hello":
                                    v3 += v8
                                    var v40 string = "hello"
                                    if v38 {
                                        v6 += v6
                                    } else {
                                        v5 = v5
                                    }
                                    _ = v40
                                case "a":
                                    var v41 int32 = 2147483647
                                    _ = v41
                                default:
                                    if len(v4) > 0 {
                                        v4[0] = "hello"
                                    }
                                    v6 = 42
                                    var v42 int8 = (v6 + (v6 / (((v6 >> ((1) & 63))) | 1)))
                                    fmt.Println(v8)
                                    v43 := GenS5{F0: false}
                                    _ = v42
                                    _ = v43
                                }
                                v5 += fi19
                                switch ("\n\t" + "hello") {
                                case "a":
                                    fmt.Println(v3)
                                default:
                                    v44 := GenS3{F0: -1, F1: (^(v6)), F2: (fi37 >> ((512) & 63))}
                                    _ = v44
                                }
                                _ = fi37
                                _ = v38
                            }
                            var v45 uint16 = 196
                            _ = v45
                        }
                    case 128:
                        if ((float64((1e10 - 0.1)) / -2.25) > 3.14159) {
                            v5 = v5
                        } else {
                            switch (v5 * 0) {
                            case -9223372036854775808:
                                v5 -= (fi19 ^ 4294967295)
                            case 4294967295:
                                var v46 uint16 = 2
                                _ = v46
                            case -1:
                                v47 := GenS6{F0: (!(false || ((0.1 * float32(3.14159)) != (0.0 - 1.5)))), F1: 1, F2: 152}
                                _ = v47
                            }
                            v48 := GenS7{F0: -1}
                            delete(v1, 1)
                            _ = v48
                        }
                    case -2147483648:
                        var v49 float32 = ((-2.25 - 1e10) - -2.25)
                        _ = v49
                    default:
                        if true {
                            var v50 int16 = 748
                            var v51 float32 = (1e10 + (0.0 - 3.14159))
                            var v52 uint64 = 1
                            _ = v50
                            _ = v51
                            _ = v52
                        }
                    }
                    var v53 uint = 65535
                    if (false && (v8 == v3)) {
                        for fi54 := 0; fi54 < 19; fi54++ {
                            var v55 uint32 = 702
                            switch fi54 {
                            case 32767:
                                if (0.0 == 0.0) {
                                    var v56 int32 = -2147483648
                                    var v57 string = "hello"
                                    if (!false) {
                                        var v58 int32 = (v56 * (-(v56)))
                                        _ = v58
                                    } else {
                                        var v59 uint64 = 1
                                        _ = v59
                                    }
                                    _ = v56
                                    _ = v57
                                } else {
                                    var v60 uint = (v53 * (v53 + 42))
                                    v61 := GenS8{F0: 42, F1: 0}
                                    _ = v60
                                    _ = v61
                                }
                            case -1:
                                if len(v4) > 0 {
                                    v4[0] = "\n\t"
                                }
                            case 128:
                                v6 -= v6
                                var v62 int64 = 42
                                v8 -= (v8 & v8)
                                fmt.Println(v53, v8)
                                _ = v62
                            default:
                                v63 := GenS7{F0: 2147483647}
                                _ = v63
                            }
                            v64 := GenS9{F0: v5, F1: 42, F2: 1}
                            _ = fi54
                            _ = v55
                            _ = v64
                        }
                    }
                    _ = v53
                } else {
                    v65 := GenS4{F0: (-(v5)), F1: 4294967295, F2: 18446744073709551615}
                    fmt.Println(v8, v5, v3)
                    _ = v65
                }
            }
        }
        _ = fi19
    }
    var v66 uint8 = (v3 - 0)
    v67 := GenS3{F0: 0, F1: v6, F2: (v5 << ((1) & 63))}
    var v68 uint = 1
    if true {
        switch "foo bar" {
        case "\n\t":
            for fi69 := 0; fi69 < 8; fi69++ {
                v3 += 189
                v4 = append(v4, "\n\t")
                _ = fi69
            }
        default:
            for fi70 := 0; fi70 < 19; fi70++ {
                fmt.Println(v3, v66)
                var v71 uint32 = 145
                _ = fi70
                _ = v71
            }
        }
    } else {
        v3 -= v3
        fmt.Println(v6, v68)
    }
    fmt.Println(v66)
    fmt.Println(v68, v6)
    v4 = append(v4, "a")
    v3 = (v3 % ((v66) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("v68:", v68)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v67
    _ = v68
    _ = v8
}
