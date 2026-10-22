package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS4 struct { F0 uint8; F1 int8; F2 int16 }
type GenS5 struct { F0 uint32; F1 int; F2 int64 }
type GenS6 struct { F0 uint8; F1 int }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int8; F1 int64 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int; F1 uint8; F2 uint64 }

func main() {
    var v1 int16 = 876
    v2 := make(map[string]string)
    var v3 string = "\n\t"
    switch v3 {
    case "\n\t":
        v4 := v2["foo bar"]
        _ = v4
    }
    var v5 uint16 = 871
    var v6 bool = (false && (2 != 429))
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := GenS1{F0: (("a" + "foo bar") + v3)}
        if ((!true) || (-1 < 2147483647)) {
            v5 = (^(v5))
        } else {
            if true {
                if ((1.5 * -2.25) >= float32(1e10)) {
                    if v6 {
                        var v9 float64 = (1.5 - 1.5)
                        _ = v9
                    } else {
                        var v10 bool = true
                        _ = v10
                    }
                    switch ((v3 + v3) + "hello") {
                    case "foo bar":
                        v3 = "\n\t"
                    case "a":
                        v11 := GenS2{F0: 0}
                        _ = v11
                    default:
                        if false {
                            v1 += (v1 / (((v1 | (v1 * (-(v1))))) | 1))
                            var v12 float64 = 3.14159
                            _ = v12
                        }
                    }
                    fmt.Println(v3, v3, v5)
                }
                switch 2147483647 {
                case 2147483647:
                    fmt.Println(v5, fi7, v1)
                default:
                    var v13 int = -588
                    _ = v13
                }
            } else {
                var v14 int16 = 32767
                var v15 int8 = 42
                _ = v14
                _ = v15
            }
        }
        _ = fi7
        _ = v8
    }
    if v6 {
        if (0.0 < (1e10 / float32(0.1))) {
            for fi16 := 0; fi16 < 12; fi16++ {
                if (fi16 >= 32768) {
                    switch (fi16 % ((fi16) | 1)) {
                    case -158:
                        v5 = v5
                    case 128:
                        v6 = v6
                    case -2147483648:
                        v17 := []int{fi16, 128, fi16}
                        _ = v17
                    default:
                        v1 = (v1 % ((0) | 1))
                    }
                    for fi18 := 0; fi18 < 8; fi18++ {
                        v19 := []int{128, -1, 127}
                        _ = fi18
                        _ = v19
                    }
                    v20 := []int{fi16, fi16}
                    _ = v20
                } else {
                    delete(v2, v3)
                    var v21 uint = 228
                    for fi22 := 0; fi22 < 7; fi22++ {
                        var v23 uint32 = 2
                        v24 := v2[((v3 + ((v3 + "foo bar") + "foo bar")) + v3)]
                        v25 := []int{(-(fi22))}
                        _ = fi22
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                    _ = v21
                }
                _ = fi16
            }
        } else {
            var v26 int8 = -128
            v27 := GenS1{F0: v3}
            v3 = v3
            _ = v26
            _ = v27
        }
        switch ("hello" + ("a" + (((v3 + ("" + "")) + ((v3 + v3) + "a")) + ("" + v3)))) {
        case "foo bar":
            var v28 uint32 = 2
            _ = v28
        case "\n\t":
            v29 := GenS3{F0: 0, F1: 90, F2: 42}
            _ = v29
        case "":
            if (!(100 >= -34)) {
                var v30 int32 = 2147483647
                fmt.Println(v30)
                _ = v30
            } else {
                var v31 uint16 = v5
                for fi32 := 0; fi32 < 19; fi32++ {
                    for fi33 := 0; fi33 < 9; fi33++ {
                        var v34 int8 = -128
                        _ = fi33
                        _ = v34
                    }
                    _ = fi32
                }
                _ = v31
            }
        }
        if ("a" < ((v3 + v3) + "hello")) {
            v35 := GenS4{F0: 2, F1: 69, F2: 32767}
            v1 -= v1
            var v36 string = (v3 + (v3 + v3))
            _ = v35
            _ = v36
        }
    }
    _ = len(v2)
    v3 = v3
    switch v3 {
    case "foo bar":
        v1 = 1000
    case "\n\t":
        switch v3 {
        case "a":
            for fi37 := 0; fi37 < 17; fi37++ {
                for fi38 := 0; fi38 < 11; fi38++ {
                    if v6 {
                        var v39 float32 = 3.14159
                        var v40 uint8 = 128
                        _ = v39
                        _ = v40
                    }
                    _ = fi38
                }
                _ = fi37
            }
        default:
            var v41 bool = (-642 != -952)
            _ = v41
        }
    case "hello":
        v6 = (!false)
    }
    v42 := GenS5{F0: 65535, F1: 0, F2: 783}
    if (v5 == 249) {
        if (((!((v3 + "\n\t") == "foo bar")) || (1 != 1)) || v6) {
            var v43 uint = 314
            switch -2147483648 {
            case 1:
                if v6 {
                    v44 := GenS6{F0: 12, F1: -540}
                    _ = v44
                } else {
                    v3 = v3
                    v45 := GenS7{F0: v43}
                    _ = v45
                }
            case 42:
                for fi46 := 0; fi46 < 10; fi46++ {
                    v47 := GenS3{F0: 127, F1: 128, F2: 642}
                    for fi48 := 0; fi48 < 11; fi48++ {
                        v1 = 32767
                        var v49 int32 = 1
                        switch (fi48 << ((v43) & 63)) {
                        case 32768:
                            v50 := GenS8{F0: 14, F1: 42}
                            _ = v50
                        case 32767:
                            fmt.Println(v5)
                        }
                        _ = fi48
                        _ = v49
                    }
                    switch v3 {
                    case "foo bar":
                        v51 := v2["hello"]
                        _ = v51
                    case "a":
                        v5 = (v5 << (((v43 >> (((v43 ^ v43)) & 63))) & 63))
                    default:
                        if v6 {
                            var v52 float32 = 0.1
                            _ = v52
                        } else {
                            fmt.Println(v6, v1)
                            switch "hello" {
                            case "\n\t":
                                v53 := []int{42, 517, -89, 192}
                                _ = v53
                            case "":
                                for fi54 := 0; fi54 < 14; fi54++ {
                                    v2[v3] = (("hello" + ("" + ((v3 + "foo bar") + ""))) + "foo bar")
                                    v5 += uint16(v43)
                                    if (!(false && v6)) {
                                        if ((v43 * (^(v43))) >= (^(v43))) {
                                            v55 := GenS9{F0: -9223372036854775808}
                                            v56 := GenS6{F0: 1, F1: -9223372036854775808}
                                            var v57 int64 = -144
                                            _ = v55
                                            _ = v56
                                            _ = v57
                                        }
                                        var v58 bool = (!v6)
                                        var v59 int16 = (v1 & (v1 >> (((v43 % ((v43) | 1))) & 63)))
                                        _ = v58
                                        _ = v59
                                    }
                                    delete(v2, v3)
                                    switch v3 {
                                    case "a":
                                        var v60 string = ""
                                        _ = v60
                                    }
                                    var v61 uint64 = 695
                                    v1 += v1
                                    if v6 {
                                        v43 = v43
                                        var v62 string = ("\n\t" + ((v3 + "foo bar") + "hello"))
                                        _ = v62
                                    } else {
                                        switch v3 {
                                        case "foo bar":
                                            v63 := []int{-272, (fi46 % (((fi54 << ((v43) & 63))) | 1))}
                                            _ = v63
                                        case "\n\t":
                                            if ((3.14159 - 0.1) > 3.14159) {
                                                v64 := GenS5{F0: 0, F1: fi54, F2: -966}
                                                v3 = "hello"
                                                var v65 int8 = -1
                                                _ = v64
                                                _ = v65
                                            } else {
                                                v66 := GenS10{F0: int(v5), F1: 1, F2: (v61 + v61)}
                                                _ = v66
                                            }
                                        }
                                    }
                                    _ = fi54
                                    _ = v61
                                }
                            case "a":
                                var v67 int = (^(fi46))
                                _ = v67
                            }
                        }
                    }
                    _ = fi46
                    _ = v47
                }
            case 878:
                switch (v3 + "\n\t") {
                case "hello":
                    var v68 float32 = ((1.5 - (-2.25 - ((3.14159 - 1e10) + ((1e10 * float32(-2.25)) - 1e10)))) / 1e10)
                    _ = v68
                case "":
                    var v69 string = ((v3 + (v3 + ((v3 + (v3 + v3)) + v3))) + ((v3 + v3) + v3))
                    _ = v69
                }
            default:
                v2[v3] = "hello"
            }
            _ = v43
        }
    }
    fmt.Println(v1, v6)
    switch -2147483648 {
    case -1:
        var v70 int32 = -71
        _ = v70
    default:
        switch v3 {
        case "":
            for fi71 := 0; fi71 < 3; fi71++ {
                for fi72 := 0; fi72 < 11; fi72++ {
                    v1 -= -32768
                    var v73 int8 = -1
                    v74 := []int{(fi71 & -2147483648), (-(fi72)), (fi72 + 42), -441}
                    _ = fi72
                    _ = v73
                    _ = v74
                }
                _ = fi71
            }
        }
    }
    for fi75 := 0; fi75 < 3; fi75++ {
        v76 := []int{fi75, (-(fi75))}
        var v77 uint8 = 128
        _ = fi75
        _ = v76
        _ = v77
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v5
    _ = v6
}
