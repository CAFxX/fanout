package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int64; F1 int32; F2 int16 }
type GenS3 struct { F0 uint16; F1 string; F2 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int32; F1 int8; F2 bool }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint64; F1 uint8; F2 int16 }

func main() {
    v1 := []uint32{2}
    var v2 uint16 = 2
    var v3 int16 = 448
    var v4 uint = 1
    var v5 bool = (-1 == -1)
    var v6 float32 = ((1.5 - float32(0.1)) - -2.25)
    for fi7 := 0; fi7 < 13; fi7++ {
        fmt.Println(v6)
        v4 -= 914
        _ = len(v1)
        _ = fi7
    }
    switch ((("a" + ("" + "")) + "\n\t") + "a") {
    case "":
        switch 4294967295 {
        case 32767:
            var v8 string = "hello"
            _ = v8
        case 2147483647:
            if false {
                if false {
                    var v9 int32 = 2147483647
                    _ = v9
                } else {
                    fmt.Println(v2, v6, v5)
                    v1 = append(v1, 2147483648)
                    var v10 uint32 = 65535
                    _ = v10
                }
                v6 += float32(1.5)
            }
        case 661:
            for fi11 := 0; fi11 < 5; fi11++ {
                v5 = (12 == 127)
                if len(v1) > 0 {
                    v1[0] = 0
                }
                for fi12 := 0; fi12 < 11; fi12++ {
                    if (9223372036854775807 <= 0) {
                        var v13 int32 = -1
                        _ = len(v1)
                        _ = v13
                    } else {
                        v1 = append(v1, 4294967295)
                        v14 := GenS1{F0: 1, F1: uint8(v2), F2: (v2 * v2)}
                        v4 -= (v4 << ((4294967295) & 63))
                        _ = v14
                    }
                    _ = fi12
                }
                _ = fi11
            }
        }
    case "hello":
        v2 = (^(v2))
    case "foo bar":
        v2 = (v2 - 2)
    default:
        for fi15 := 0; fi15 < 19; fi15++ {
            var v16 int16 = int16(fi15)
            _ = fi15
            _ = v16
        }
    }
    if false {
        v17 := GenS1{F0: 100, F1: 127, F2: (v2 ^ v2)}
        for fi18 := 0; fi18 < 19; fi18++ {
            for fi19 := 0; fi19 < 13; fi19++ {
                v2 = (v2 * v2)
                _ = len(v1)
                v1 = append(v1, 4294967295)
                _ = fi19
            }
            if v5 {
                var v20 uint32 = 42
                v21 := GenS2{F0: -697, F1: -2147483648, F2: (v3 ^ (v3 % (((v3 ^ v3)) | 1)))}
                fmt.Println(v20)
                _ = v20
                _ = v21
            }
            switch (-(fi18)) {
            case 32767:
                v22 := GenS3{F0: (v2 + v2), F1: ((("" + ("" + ("\n\t" + ""))) + (("" + ("foo bar" + "a")) + "a")) + "hello"), F2: (v3 - (^(v3)))}
                _ = v22
            }
            _ = fi18
        }
        v3 -= v3
        _ = v17
    }
    switch 127 {
    case 0:
        for fi23 := 0; fi23 < 11; fi23++ {
            var v24 float64 = -2.25
            if (("foo bar" + "") > "\n\t") {
                v4 += v4
                var v25 bool = (!v5)
                _ = v25
            } else {
                for fi26 := 0; fi26 < 11; fi26++ {
                    v2 = v2
                    _ = fi26
                }
            }
            _ = fi23
            _ = v24
        }
    case -773:
        v27 := GenS3{F0: 199, F1: ("a" + "\n\t"), F2: v3}
        _ = v27
    case 1:
        v28 := GenS4{F0: int32(v3)}
        _ = v28
    default:
        for fi29 := 0; fi29 < 1; fi29++ {
            v30 := GenS2{F0: -1, F1: 2147483647, F2: v3}
            for fi31 := 0; fi31 < 1; fi31++ {
                var v32 int64 = 1
                _ = len(v1)
                _ = fi31
                _ = v32
            }
            v1 = append(v1, 42)
            _ = fi29
            _ = v30
        }
    }
    for fi33 := 0; fi33 < 15; fi33++ {
        if v5 {
            for fi34 := 0; fi34 < 14; fi34++ {
                var v35 int8 = -1
                switch fi33 {
                case 584:
                    v36 := GenS5{F0: -1}
                    _ = v36
                case -631:
                    var v37 uint16 = (v2 / (((v2 / ((65535) | 1))) | 1))
                    _ = v37
                case -9223372036854775808:
                    v6 += 1.5
                default:
                    var v38 int32 = -2147483648
                    fmt.Println(v2, fi33, v35)
                    _ = v38
                }
                if (0.1 != v6) {
                    var v39 float32 = v6
                    var v40 int64 = 1
                    _ = v39
                    _ = v40
                } else {
                    v5 = true
                    _ = len(v1)
                }
                var v41 int = (fi33 & fi33)
                if len(v1) > 0 {
                    v1[0] = 2147483648
                }
                var v42 int64 = 167
                _ = len(v1)
                v35 = v35
                if v5 {
                    v43 := GenS6{F0: int32(v42), F1: int8(v3), F2: v5}
                    var v44 uint32 = 553
                    var v45 uint = 18446744073709551615
                    _ = v43
                    _ = v44
                    _ = v45
                } else {
                    _ = len(v1)
                }
                _ = fi34
                _ = v35
                _ = v41
                _ = v42
            }
            v6 = 3.14159
        }
        _ = fi33
    }
    if (100 <= 100) {
        switch "foo bar" {
        case "a":
            v3 += v3
        case "":
            var v46 int8 = 42
            _ = v46
        case "foo bar":
            if (42 <= 18446744073709551615) {
                fmt.Println(v5, v6)
                var v47 uint8 = 127
                _ = v47
            } else {
                var v48 int32 = 1
                var v49 int8 = 1
                _ = len(v1)
                _ = v48
                _ = v49
            }
        default:
            var v50 uint8 = 34
            _ = v50
        }
    }
    v5 = ((42 <= 121) || ((^(v4)) != v4))
    var v51 int64 = 9223372036854775807
    v52 := GenS6{F0: 2147483647, F1: 73, F2: ((917 != 0) || (0.1 < 3.14159))}
    fmt.Println(v5)
    for fi53 := 0; fi53 < 11; fi53++ {
        if v5 {
            if (false && (78 == 208)) {
                for fi54 := 0; fi54 < 19; fi54++ {
                    v2 = v2
                    if v5 {
                        switch "foo bar" {
                        case "\n\t":
                            var v55 int32 = 2147483647
                            _ = v55
                        case "a":
                            var v56 int64 = (-(v51))
                            _ = v56
                        case "":
                            if (!(v5 && true)) {
                                if len(v1) > 0 {
                                    v1[0] = 2
                                }
                                fmt.Println(v5)
                            } else {
                                if ((v2 >> (((^(v4))) & 63)) == (^(v2))) {
                                    var v57 uint64 = 158
                                    v6 = ((v6 - v6) * float32((0.1 - (0.1 + (0.0 - -2.25)))))
                                    v5 = (!((4294967295 >= v57) && v5))
                                    _ = v57
                                } else {
                                    v58 := GenS7{F0: -56}
                                    _ = v58
                                }
                            }
                        }
                        if len(v1) > 0 {
                            v1[0] = 621
                        }
                    } else {
                        v6 = v6
                        fmt.Println(fi53, v4)
                        var v59 float64 = 1.5
                        _ = v59
                    }
                    fmt.Println(v4)
                    _ = fi54
                }
                if false {
                    v60 := GenS8{F0: 2, F1: 144, F2: -415}
                    if ((v3 >> ((v4) & 63)) == v3) {
                        if false {
                            var v61 int16 = (v3 * int16(v4))
                            _ = v61
                        }
                        switch 2147483647 {
                        case 42:
                            var v62 uint64 = 661
                            _ = v62
                        }
                        var v63 uint32 = 2147483648
                        _ = v63
                    }
                    _ = v60
                }
                var v64 uint32 = 0
                _ = v64
            } else {
                var v65 int16 = int16(v2)
                _ = v65
            }
            v51 = (-(v51))
            for fi66 := 0; fi66 < 9; fi66++ {
                fmt.Println(v51, fi53, fi53)
                _ = fi66
            }
        }
        var v67 uint = v4
        _ = fi53
        _ = v67
    }
    v1 = append(v1, 4294967295)
    v68 := GenS5{F0: 42}
    var v69 uint32 = 74
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v6
    _ = v68
    _ = v69
}
