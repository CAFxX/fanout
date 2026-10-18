package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint; F1 uint; F2 int64 }
type GenS4 struct { F0 int64; F1 int8; F2 int64 }
type GenS5 struct { F0 int64; F1 uint16; F2 uint16 }
type GenS6 struct { F0 int32; F1 int8; F2 int32 }
type GenS7 struct { F0 int64; F1 uint32 }
type GenS8 struct { F0 uint64 }

func main() {
    var v1 int = 42
    v2 := make(map[string]string)
    var v3 string = "hello"
    v4 := []string{v3, (v3 + "a")}
    _ = len(v4)
    var v5 int16 = 1000
    if false {
        switch (v3 + "hello") {
        case "":
            fmt.Println(v5, v1)
        case "hello":
            for fi6 := 0; fi6 < 9; fi6++ {
                for fi7 := 0; fi7 < 2; fi7++ {
                    switch "foo bar" {
                    case "":
                        delete(v2, (((v3 + (v3 + ((v3 + v3) + "foo bar"))) + v3) + (v3 + "hello")))
                    case "hello":
                        v3 = ("foo bar" + v3)
                    default:
                        v5 -= v5
                    }
                    _ = fi7
                }
                v5 = v5
                _ = fi6
            }
        default:
            if (!true) {
                v3 = v3
                var v8 uint8 = 0
                var v9 int8 = -1
                _ = v8
                _ = v9
            } else {
                if false {
                    var v10 uint8 = 2
                    _ = v10
                } else {
                    if (true || false) {
                        if true {
                            fmt.Println(v3)
                            switch v1 {
                            case 0:
                                v11 := GenS1{F0: -982}
                                _ = v11
                            default:
                                fmt.Println(v1, v1)
                            }
                        }
                    } else {
                        v3 = (v3 + "")
                        v3 = "a"
                    }
                    _ = len(v2)
                    v5 = v5
                }
                switch v3 {
                case "hello":
                    var v12 int64 = 219
                    _ = v12
                case "\n\t":
                    if true {
                        for fi13 := 0; fi13 < 15; fi13++ {
                            if false {
                                var v14 int16 = v5
                                _ = v14
                            } else {
                                if (3.14159 > float32(float64(1e10))) {
                                    if true {
                                        v3 = ""
                                        if (!(0.0 >= 3.14159)) {
                                            for fi15 := 0; fi15 < 13; fi15++ {
                                                v5 = v5
                                                v16 := GenS1{F0: 9223372036854775807}
                                                _ = fi15
                                                _ = v16
                                            }
                                            for fi17 := 0; fi17 < 14; fi17++ {
                                                if (0.0 > 1.5) {
                                                    var v18 int32 = -973
                                                    _ = v18
                                                }
                                                fmt.Println(v5, fi13)
                                                _ = fi17
                                            }
                                        } else {
                                            if (71 > 1) {
                                                v19 := GenS1{F0: 42}
                                                _ = v19
                                            }
                                            fmt.Println(v5)
                                        }
                                        switch -417 {
                                        case -1:
                                            if len(v4) > 0 {
                                                v4[0] = ("" + "\n\t")
                                            }
                                        case 127:
                                            v20 := GenS2{F0: 774, F1: ""}
                                            _ = v20
                                        }
                                    }
                                    v5 = -71
                                    var v21 int8 = -128
                                    _ = v21
                                }
                            }
                            _ = fi13
                        }
                    }
                case "":
                    for fi22 := 0; fi22 < 12; fi22++ {
                        for fi23 := 0; fi23 < 11; fi23++ {
                            v24 := GenS1{F0: 989}
                            if false {
                                var v25 uint16 = 65535
                                _ = v25
                            } else {
                                delete(v2, v3)
                                if (-32768 <= v5) {
                                    if false {
                                        v26 := GenS3{F0: 0, F1: 65535, F2: 346}
                                        if (fi23 > (fi23 ^ (fi23 & (v1 | (-(fi23)))))) {
                                            v3 = "a"
                                            v27 := GenS4{F0: -9223372036854775808, F1: 27, F2: 1}
                                            _ = v27
                                        }
                                        _ = v26
                                    }
                                    if false {
                                        var v28 bool = false
                                        var v29 uint64 = 0
                                        _ = v28
                                        _ = v29
                                    } else {
                                        _ = len(v2)
                                        var v30 uint32 = 187
                                        _ = v30
                                    }
                                    var v31 uint16 = 1
                                    _ = v31
                                } else {
                                    if (0 > 1) {
                                        v1 = fi22
                                    }
                                    if (true && false) {
                                        switch v3 {
                                        case "":
                                            if ((0 == 0) || (!(127 == v5))) {
                                                v2[v3] = v3
                                            } else {
                                                switch (v1 ^ (fi22 - 1)) {
                                                case 1:
                                                    if len(v4) > 0 {
                                                        v4[0] = "a"
                                                    }
                                                case 127:
                                                    var v32 float32 = 0.0
                                                    _ = v32
                                                case 42:
                                                    v1 = v1
                                                default:
                                                    v33 := v2[(((v3 + (v3 + ("\n\t" + v3))) + v3) + (v3 + ((v3 + v3) + (v3 + v3))))]
                                                    _ = v33
                                                }
                                            }
                                        case "foo bar":
                                            var v34 int16 = 42
                                            _ = v34
                                        }
                                    }
                                    v3 = "\n\t"
                                }
                            }
                            _ = fi23
                            _ = v24
                        }
                        if (4294967295 >= 2147483647) {
                            var v35 uint8 = 1
                            var v36 float64 = 0.0
                            v37 := GenS1{F0: -1}
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                        _ = fi22
                    }
                }
            }
        }
        if false {
            v3 = v3
            var v38 int16 = (-(v5))
            _ = v38
        } else {
            for fi39 := 0; fi39 < 8; fi39++ {
                if (!(!false)) {
                    var v40 int32 = 120
                    _ = v40
                } else {
                    for fi41 := 0; fi41 < 18; fi41++ {
                        if false {
                            v42 := GenS2{F0: 726, F1: v3}
                            switch fi41 {
                            case 1:
                                var v43 bool = ((-1 >= -59) || true)
                                _ = v43
                            }
                            switch "hello" {
                            case "a":
                                if (2147483647 <= 960) {
                                    var v44 uint = 425
                                    v4 = append(v4, "\n\t")
                                    _ = v44
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                            case "\n\t":
                                fmt.Println(v3)
                            default:
                                if len(v4) > 0 {
                                    v4[0] = v3
                                }
                                v3 = v3
                                v5 -= (v5 % ((-32768) | 1))
                            }
                            _ = v42
                        }
                        v3 = "a"
                        _ = len(v2)
                        _ = fi41
                    }
                }
                switch (((v3 + (v3 + ((v3 + "") + "hello"))) + (v3 + v3)) + "foo bar") {
                case "":
                    if (472 == 0) {
                        if (0 < 100) {
                            v1 -= v1
                            v5 -= (v5 % (((v5 % (((v5 << ((241) & 63))) | 1))) | 1))
                            var v45 int64 = 42
                            _ = v45
                        } else {
                            if true {
                                if ((1 > 2) && (-2147483648 <= 1)) {
                                    for fi46 := 0; fi46 < 9; fi46++ {
                                        if true {
                                            v5 += v5
                                        } else {
                                            var v47 int16 = (v5 | (v5 << ((2) & 63)))
                                            v48 := GenS5{F0: -1, F1: 82, F2: 0}
                                            var v49 uint64 = 524
                                            v4 = append(v4, v3)
                                            _ = v47
                                            _ = v48
                                            _ = v49
                                        }
                                        _ = fi46
                                    }
                                    _ = len(v4)
                                    var v50 uint64 = 1
                                    _ = v50
                                }
                                if false {
                                    v51 := GenS1{F0: -80}
                                    v2[v3] = (((v3 + (v3 + v3)) + (("a" + "hello") + "\n\t")) + "foo bar")
                                    _ = v51
                                }
                            }
                            if true {
                                fmt.Println(v3, v5, v3)
                            } else {
                                v52 := GenS6{F0: 508, F1: 24, F2: 977}
                                _ = v52
                            }
                        }
                        v1 = (v1 - 142)
                    }
                case "foo bar":
                    for fi53 := 0; fi53 < 2; fi53++ {
                        _ = len(v2)
                        v3 = v3
                        if (!(-128 < -128)) {
                            v54 := GenS7{F0: -1, F1: 4294967295}
                            v1 = -9223372036854775808
                            v3 = ("" + (v3 + v3))
                            var v55 uint = 1
                            if true {
                                delete(v2, v3)
                                switch fi53 {
                                case 2147483647:
                                    v5 -= -32768
                                case 1:
                                    var v56 int64 = 0
                                    _ = v56
                                }
                            }
                            var v57 uint16 = 2
                            _ = v54
                            _ = v55
                            _ = v57
                        } else {
                            v58 := GenS8{F0: 42}
                            switch v1 {
                            case 128:
                                if true {
                                    var v59 uint8 = 120
                                    _ = v59
                                }
                            default:
                                v1 += 0
                            }
                            _ = v58
                        }
                        _ = fi53
                    }
                }
                v60 := GenS7{F0: 42, F1: 2147483648}
                _ = fi39
                _ = v60
            }
        }
    }
    v5 = (v5 | (-(v5)))
    v5 = 1
    if false {
        var v61 float64 = (float64(0.1) - 0.0)
        v62 := v2[v3]
        _ = v61
        _ = v62
    }
    for fi63 := 0; fi63 < 10; fi63++ {
        var v64 int = (^(v1))
        _ = fi63
        _ = v64
    }
    v5 += v5
    v65 := GenS4{F0: 9223372036854775807, F1: 100, F2: 9223372036854775807}
    var v66 uint64 = 18446744073709551615
    fmt.Println(v1, v3, v66)
    v67 := GenS5{F0: 1, F1: 983, F2: 255}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v66
    _ = v67
}
