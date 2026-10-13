package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 int32; F2 int }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS6 struct { F0 int64 }
func genFn0(p0 int, p1 int16, p2 int8) uint8 {
    var v6 uint = 2
    fmt.Println(p0, p2)
    p2 = p2
    switch "hello" {
    case "a":
        if true {
            v7 := GenS1{F0: false, F1: 929, F2: (-(p2))}
            p2 = 100
            var v8 bool = true
            _ = v7
            _ = v8
        } else {
            v9 := []int{p0, (p0 << ((v6) & 63)), int(p1), 1}
            var v10 string = "\n\t"
            _ = v10
            _ = v9
        }
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 uint16 = 255
        if false {
            fmt.Println(v12)
            p2 = 57
        }
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}
func genFn1(p0 bool, p1 int8, p2 uint16) int32 {
    v13 := []int{-225}
    if (-1 != 1) {
        p0 = p0
        var v14 int16 = -32768
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return -376
}

func main() {
    var v1 int16 = 1000
    var v2 float64 = 3.14159
    v3 := []string{}
    v4 := []uint32{897, 166, 146, 822}
    var v5 int8 = 64
    v2 -= ((0.0 / (((1e10 / v2) + v2) + float64(((0.0 / -2.25) - 1.5)))) - (v2 * v2))
    switch (("a" + "foo bar") + "foo bar") {
    case "a":
        _ = len(v3)
    default:
        var v15 uint32 = 65535
        _ = v15
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        for fi17 := 0; fi17 < 13; fi17++ {
            if (v2 > 1.5) {
                if ("foo bar" < ("hello" + "\n\t")) {
                    v3 = append(v3, "\n\t")
                    var v18 uint64 = 370
                    fmt.Println(v2)
                    _ = v18
                } else {
                    if true {
                        if len(v3) > 0 {
                            v3[0] = ((("foo bar" + "hello") + "\n\t") + "foo bar")
                        }
                        var v19 float32 = ((1e10 / 0.1) * 1.5)
                        v20 := GenS1{F0: ((!true) && false), F1: 42, F2: 100}
                        v19 = (1e10 * (v19 / float32(float64(((v19 - v19) * v19)))))
                        _ = v19
                        _ = v20
                    } else {
                        var v21 uint64 = 2
                        if ((965 <= v21) || (!(("\n\t" + ("foo bar" + ("\n\t" + (("foo bar" + ("foo bar" + "\n\t")) + "")))) >= "foo bar"))) {
                            switch (-(fi16)) {
                            case 0:
                                if (false && (1 <= (^(v21)))) {
                                    var v22 int8 = 1
                                    if (!(91 != -1)) {
                                        if false {
                                            switch ("foo bar" + ("" + (("a" + "\n\t") + "\n\t"))) {
                                            case "a":
                                                switch "\n\t" {
                                                case "hello":
                                                    var v23 uint8 = 2
                                                    _ = v23
                                                case "\n\t":
                                                    v1 -= 1000
                                                }
                                            case "":
                                                v22 = v22
                                            case "\n\t":
                                                v24 := GenS2{F0: -128}
                                                _ = v24
                                            default:
                                                v1 = 1000
                                            }
                                            fmt.Println(v21)
                                            var v25 float64 = v2
                                            v26 := GenS3{F0: 0, F1: 0, F2: 0}
                                            _ = len(v4)
                                            _ = v25
                                            _ = v26
                                        }
                                        var v27 int16 = v1
                                        _ = v27
                                    } else {
                                        v2 -= v2
                                        switch ("foo bar" + ("foo bar" + ("hello" + "a"))) {
                                        case "hello":
                                            var v28 int64 = int64(v21)
                                            _ = v28
                                        default:
                                            var v29 int16 = (v1 & v1)
                                            _ = v29
                                        }
                                        fmt.Println(v5)
                                    }
                                    _ = v22
                                } else {
                                    fmt.Println(v5, v21)
                                    var v30 uint16 = 188
                                    _ = v30
                                }
                            case 4294967295:
                                var v31 int8 = v5
                                _ = v31
                            case -733:
                                v21 -= 578
                            }
                        } else {
                            var v32 float32 = 3.14159
                            v2 = (1.5 - (-2.25 - (((v2 * v2) * (v2 - v2)) - (float64(float32((v2 - v2))) - 3.14159))))
                            _ = v32
                        }
                        _ = v21
                    }
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            fmt.Println(v5, fi16, fi17)
            v33 := GenS1{F0: true, F1: 18446744073709551615, F2: (v5 | -128)}
            _ = fi17
            _ = v33
        }
        for fi34 := 0; fi34 < 10; fi34++ {
            if false {
                v4 = append(v4, 2)
                switch "" {
                case "":
                    var v35 float64 = 0.0
                    _ = v35
                case "\n\t":
                    v1 = v1
                    var v36 string = "foo bar"
                    v37 := GenS2{F0: v5}
                    _ = v36
                    _ = v37
                case "foo bar":
                    v5 = (^(v5))
                default:
                    switch ("a" + "hello") {
                    case "a":
                        fmt.Println(v5)
                    }
                }
                if (v5 > 127) {
                    v5 += v5
                } else {
                    _ = len(v3)
                }
            }
            fmt.Println(fi34, v1)
            _ = fi34
        }
        _ = fi16
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v2, v5)
    switch 2147483647 {
    case 443:
        v38 := GenS1{F0: (!true), F1: 907, F2: -1}
        _ = v38
    case 9223372036854775807:
        v39 := GenS4{F0: "foo bar", F1: false}
        _ = v39
    default:
        for fi40 := 0; fi40 < 9; fi40++ {
            var v41 uint32 = 4294967295
            fmt.Println(fi40)
            _ = fi40
            _ = v41
        }
    }
    if (18446744073709551615 == 4294967295) {
        v5 = (v5 / ((v5) | 1))
        var v42 bool = false
        _ = v42
    }
    var v43 uint64 = 18446744073709551615
    var v44 float32 = 1e10
    if true {
        for fi45 := 0; fi45 < 10; fi45++ {
            fmt.Println(v44, v43)
            var v46 string = ""
            _ = fi45
            _ = v46
        }
        v2 = 1.5
        for fi47 := 0; fi47 < 12; fi47++ {
            v4 = append(v4, 4294967295)
            var v48 float32 = (((v44 * v44) + ((0.0 + float32((1e10 / (float64(1.5) * v2)))) / 0.1)) * -2.25)
            v43 += (v43 << ((18446744073709551615) & 63))
            _ = fi47
            _ = v48
        }
    } else {
        fmt.Println(v2, v43, v2)
        var v49 int = -949
        _ = v49
    }
    for fi50 := 0; fi50 < 1; fi50++ {
        if (false || false) {
            for fi51 := 0; fi51 < 1; fi51++ {
                if (true || (!true)) {
                    v2 += v2
                }
                var v52 int = int(v1)
                _ = fi51
                _ = v52
            }
        } else {
            v53 := GenS5{F0: 980, F1: 930, F2: 644}
            _ = v53
        }
        var v54 uint8 = 127
        var v55 uint32 = 65535
        _ = fi50
        _ = v54
        _ = v55
    }
    v4 = append(v4, 2147483648)
    for fi56 := 0; fi56 < 18; fi56++ {
        v57 := GenS6{F0: 950}
        v43 = v43
        v43 -= (v43 / ((1) | 1))
        _ = fi56
        _ = v57
    }
    v5 += (v5 ^ v5)
    for fi58 := 0; fi58 < 12; fi58++ {
        var v59 uint32 = 2
        var v60 int16 = v1
        _ = fi58
        _ = v59
        _ = v60
    }
    v61 := GenS2{F0: 113}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v61
}
