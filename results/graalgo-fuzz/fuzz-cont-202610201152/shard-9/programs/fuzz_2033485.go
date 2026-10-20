package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int8; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint32 }
func genFn0() int8 {
    switch "" {
    case "":
        for fi7 := 0; fi7 < 6; fi7++ {
            fmt.Println(fi7)
            _ = fi7
        }
    case "\n\t":
        switch 9223372036854775807 {
        case 32767:
            switch ("" + (("\n\t" + (("hello" + (("" + "a") + "foo bar")) + "")) + (("" + ("a" + "hello")) + (((((("" + ("hello" + "foo bar")) + "") + "foo bar") + ("a" + ("\n\t" + "a"))) + "") + "hello")))) {
            case "hello":
                var v8 bool = true
                _ = v8
            case "a":
                var v9 float64 = (1.5 / 1e10)
                _ = v9
            case "\n\t":
                var v10 float32 = 3.14159
                _ = v10
            }
        default:
            for fi11 := 0; fi11 < 1; fi11++ {
                var v12 uint16 = 1
                v13 := []int{fi11, 32767, fi11, (fi11 | (fi11 >> ((261) & 63)))}
                switch "\n\t" {
                case "hello":
                    v12 = v12
                default:
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                }
                _ = fi11
                _ = v12
                _ = v13
            }
        }
    case "hello":
        var v14 uint32 = 1
        _ = v14
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        switch "hello" {
        case "a":
            switch -9223372036854775808 {
            case 32768:
                var v16 int32 = -1
                _ = v16
            case 127:
                for fi17 := 0; fi17 < 19; fi17++ {
                    var v18 float32 = -2.25
                    _ = fi17
                    _ = v18
                }
            default:
                var v19 int16 = -1
                _ = v19
            }
        case "":
            switch "a" {
            case "foo bar":
                var v20 int32 = 432
                _ = v20
            }
        }
        switch "hello" {
        case "a":
            var v21 int32 = 42
            _ = v21
        }
        _ = fi15
    }
    var v22 int8 = 0
    _ = v22
    return 127
}
func genFn1() uint8 {
    var v23 uint8 = 13
    for fi24 := 0; fi24 < 1; fi24++ {
        var v25 bool = (!(false && (201 < fi24)))
        fmt.Println(v25, fi24, fi24)
        v23 += 198
        _ = fi24
        _ = v25
    }
    _ = v23
    return 0
}

func main() {
    var v1 int8 = -128
    var v2 uint = 4294967295
    var v3 uint8 = 77
    v4 := make(map[int]string)
    var v5 uint32 = 42
    var v6 int8 = (v1 << (((v2 & (v2 % ((316) | 1)))) & 63))
    if false {
        switch -2147483648 {
        case -665:
            var v26 uint64 = 520
            _ = v26
        default:
            var v27 uint = v2
            _ = v27
        }
        v4[42] = "\n\t"
    } else {
        if (int16(v1) > -102) {
            v2 = (v2 & 1)
            var v28 uint32 = (v5 & 0)
            var v29 uint = v2
            _ = v28
            _ = v29
        } else {
            fmt.Println(v6, v6)
            v30 := []int{0, -693, -9223372036854775808, 0}
            if true {
                if true {
                    for fi31 := 0; fi31 < 1; fi31++ {
                        if ((^(v3)) > (v3 | 128)) {
                            for fi32 := 0; fi32 < 19; fi32++ {
                                var v33 bool = true
                                var v34 float32 = -2.25
                                v35 := GenS1{F0: 0, F1: 4294967295}
                                _ = fi32
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                            if (true || false) {
                                switch 9223372036854775807 {
                                case 2147483647:
                                    var v36 float64 = float64((1e10 + -2.25))
                                    _ = v36
                                default:
                                    if (((!((!true) || true)) && (!true)) || true) {
                                        v37 := v4[fi31]
                                        _ = v37
                                    }
                                }
                            } else {
                                v38 := v4[296]
                                switch v38 {
                                case "\n\t":
                                    v38 = v38
                                case "hello":
                                    var v39 bool = (true && true)
                                    _ = v39
                                case "foo bar":
                                    v5 += v5
                                }
                                switch (("hello" + v38) + v38) {
                                case "foo bar":
                                    v38 = v38
                                case "\n\t":
                                    var v40 uint32 = (v5 / ((v5) | 1))
                                    _ = v40
                                }
                                _ = v38
                            }
                        }
                        v6 -= v6
                        var v41 float64 = (0.1 / 1.5)
                        _ = fi31
                        _ = v41
                    }
                }
                v42 := GenS1{F0: -1, F1: 2}
                v1 -= v1
                _ = v42
            } else {
                switch -367 {
                case 142:
                    delete(v4, 32768)
                case 128:
                    var v43 int32 = -874
                    _ = v43
                case -790:
                    var v44 float32 = -2.25
                    _ = v44
                default:
                    fmt.Println(v6)
                }
                fmt.Println(v3, v5, v6)
            }
            _ = v30
        }
    }
    _ = len(v4)
    v45 := []int{4294967295, -800, -105}
    if len(v45) > 0 {
        v45[0] = 1
    }
    v46 := GenS2{F0: (v1 ^ v6), F1: 4294967295, F2: (v1 - (v1 / ((int8(v3)) | 1)))}
    v47 := GenS3{F0: (v5 + v5)}
    fmt.Println(v6, v2)
    var v48 uint32 = (^(v5))
    v2 += 4294967295
    if ((v2 - v2) > uint(v48)) {
        v49 := GenS1{F0: 0, F1: 2}
        _ = v49
    } else {
        v45 = append(v45, 127)
        v5 = (v48 / (((^(v48))) | 1))
    }
    v4[-9223372036854775808] = "a"
    fmt.Println(v48, v6, v6)
    var v50 float32 = 1.5
    if true {
        if ("a" > "a") {
            for fi51 := 0; fi51 < 3; fi51++ {
                v1 -= v1
                switch "a" {
                case "":
                    v3 = 85
                case "\n\t":
                    v2 -= 65535
                }
                _ = fi51
            }
            var v52 uint = v2
            switch int(v1) {
            case -9223372036854775808:
                if false {
                    v53 := GenS3{F0: v48}
                    _ = v53
                } else {
                    if false {
                        _ = len(v45)
                        var v54 int64 = -1
                        var v55 string = "hello"
                        _ = v54
                        _ = v55
                    }
                }
            default:
                v1 -= (v6 + v1)
            }
            _ = v52
        } else {
            if true {
                v5 += v48
            } else {
                v45 = append(v45, -2147483648)
            }
        }
        for fi56 := 0; fi56 < 16; fi56++ {
            if true {
                for fi57 := 0; fi57 < 14; fi57++ {
                    var v58 uint32 = v48
                    if len(v45) >= 2 {
                        v45 = v45[1:len(v45)]
                    }
                    _ = fi57
                    _ = v58
                }
                v59 := GenS1{F0: -1, F1: 1}
                for fi60 := 0; fi60 < 14; fi60++ {
                    v4[(fi60 & fi60)] = ("hello" + "")
                    switch 4294967295 {
                    case 2147483647:
                        if (255 == genFn1()) {
                            v45 = append(v45, (^(fi56)))
                            var v61 bool = true
                            var v62 int64 = -307
                            if len(v45) > 0 {
                                v45[0] = fi60
                            }
                            v62 += 9223372036854775807
                            fmt.Println(fi56, fi60)
                            _ = v61
                            _ = v62
                        }
                    case 1:
                        v4[(-(fi56))] = "hello"
                    case 42:
                        var v63 int32 = 467
                        _ = v63
                    default:
                        v64 := v4[2147483647]
                        _ = v64
                    }
                    _ = fi60
                }
                _ = v59
            }
            _ = fi56
        }
    }
    v3 = (v3 / ((v3) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v50
    _ = v6
}
