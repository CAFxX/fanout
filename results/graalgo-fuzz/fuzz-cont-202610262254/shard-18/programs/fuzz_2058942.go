package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint32; F2 int }
type GenS3 struct { F0 int16; F1 bool; F2 uint16 }
type GenS4 struct { F0 int8; F1 int64; F2 int16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string; F1 int; F2 int8 }
type GenS7 struct { F0 uint8 }
func genFn0(p0 int, p1 int32) int64 {
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 int32 = (p1 | (p1 >> ((65535) & 63)))
        p1 = 546
        if (!((p0 ^ (fi5 * (fi5 >> ((65535) & 63)))) < (p0 ^ fi5))) {
            var v7 int32 = p1
            var v8 uint = 42
            _ = v7
            _ = v8
        } else {
            var v9 uint8 = 42
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    var v10 bool = (("hello" + "") > "")
    _ = p0
    _ = p1
    _ = v10
    return 1
}
func genFn1(p0 int32, p1 int32, p2 uint) int32 {
    p1 -= p0
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 uint8 = 127
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: p1}
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return (-(p0))
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 1
    var v3 uint32 = 65535
    v4 := make(map[uint32]string)
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 uint16 = 65535
        var v16 float64 = (1.5 / 1.5)
        v17 := GenS1{F0: 42}
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        for fi19 := 0; fi19 < 10; fi19++ {
            v20 := []int{(fi18 >> ((4294967295) & 63)), (fi18 & (fi18 << ((0) & 63))), 38}
            if true {
                if false {
                    fmt.Println(fi19, fi19)
                }
            } else {
                fmt.Println(v3)
                var v21 int8 = 127
                v22 := GenS2{F0: true, F1: v2, F2: fi18}
                _ = v21
                _ = v22
            }
            var v23 uint32 = v3
            if false {
                v24 := GenS1{F0: -2147483648}
                _ = v24
            } else {
                v23 = v3
                switch fi19 {
                case -709:
                    switch ((("\n\t" + "foo bar") + "\n\t") + (((((("" + ("hello" + (("" + "\n\t") + "a"))) + "a") + "foo bar") + "a") + "a") + "a")) {
                    case "a":
                        var v25 uint64 = 4294967295
                        _ = v25
                    case "":
                        switch fi18 {
                        case 32768:
                            var v26 string = "\n\t"
                            _ = v26
                        case 163:
                            fmt.Println(fi19, v3, v23)
                        case 127:
                            delete(v1, "foo bar")
                        }
                    case "hello":
                        var v27 bool = false
                        _ = v27
                    }
                default:
                    v2 -= 2147483648
                }
                v2 += v23
            }
            _ = fi19
            _ = v20
            _ = v23
        }
        var v28 int = (-(fi18))
        switch ((("foo bar" + "\n\t") + ("a" + ("foo bar" + "\n\t"))) + ("a" + "foo bar")) {
        case "":
            switch "a" {
            case "foo bar":
                if true {
                    v29 := GenS2{F0: true, F1: v2, F2: (fi18 << ((295) & 63))}
                    if (749 > 42) {
                        if true {
                            var v30 uint16 = 718
                            v31 := GenS2{F0: true, F1: (v2 % (((v3 * (v2 << ((1000) & 63)))) | 1)), F2: 540}
                            _ = v30
                            _ = v31
                        } else {
                            var v32 uint64 = 941
                            _ = v32
                        }
                        v33 := []int{32767, (fi18 + v28), 1}
                        _ = v33
                    } else {
                        v1["a"] = (-1 == 316)
                        v3 = (v3 + (^(v2)))
                    }
                    var v34 int = fi18
                    _ = v29
                    _ = v34
                }
            case "\n\t":
                v35 := GenS1{F0: -2147483648}
                _ = v35
            }
        default:
            if false {
                switch "\n\t" {
                case "":
                    for fi36 := 0; fi36 < 4; fi36++ {
                        var v37 bool = (((529 >= 0) || (false || (1 >= 9223372036854775807))) && (1 == -980))
                        v3 += v3
                        switch "a" {
                        case "a":
                            v38 := GenS1{F0: 2147483647}
                            _ = v38
                        case "hello":
                            v28 += (v28 ^ (fi36 + (fi18 / (((fi18 & -2147483648)) | 1))))
                        case "\n\t":
                            if v37 {
                                if (true || true) {
                                    switch (-(fi18)) {
                                    case -68:
                                        v39 := GenS3{F0: 1, F1: (-210 != -9223372036854775808), F2: 46}
                                        _ = v39
                                    case -1:
                                        if false {
                                            var v40 int64 = 42
                                            if (0.0 > float32(1e10)) {
                                                var v41 uint64 = 126
                                                v2 -= (^(v2))
                                                _ = v41
                                            }
                                            v42 := []int{v28}
                                            _ = v40
                                            _ = v42
                                        }
                                    case 4294967295:
                                        fmt.Println(v28, v37)
                                    default:
                                        v28 += -1
                                    }
                                    v43 := []int{(fi18 * fi36), (-(v28))}
                                    _ = v43
                                }
                            } else {
                                fmt.Println(v2)
                                v37 = v37
                                fmt.Println(fi18)
                            }
                            v44 := GenS3{F0: 1000, F1: false, F2: 915}
                            _ = v44
                        default:
                            if v37 {
                                var v45 uint16 = 2
                                _ = v45
                            } else {
                                v46 := GenS4{F0: 0, F1: genFn0(fi18, -2147483648), F2: 693}
                                var v47 bool = v37
                                _ = v46
                                _ = v47
                            }
                        }
                        _ = fi36
                        _ = v37
                    }
                default:
                    if true {
                        v28 = fi18
                        var v48 int8 = -1
                        fmt.Println(v48, v3)
                        _ = v48
                    }
                }
                fmt.Println(v3)
            } else {
                v3 = 42
            }
        }
        _ = fi18
        _ = v28
    }
    v49 := GenS5{F0: 1}
    var v50 uint64 = 308
    for fi51 := 0; fi51 < 12; fi51++ {
        var v52 uint8 = 1
        var v53 int32 = genFn1(596, 1, 42)
        fmt.Println(v2, v53)
        _ = fi51
        _ = v52
        _ = v53
    }
    switch ("" + ("a" + ("hello" + "foo bar"))) {
    case "hello":
        var v54 uint8 = 255
        _ = v54
    case "foo bar":
        v3 -= (v2 >> ((255) & 63))
    case "a":
        v55 := v1["\n\t"]
        _ = v55
    default:
        v56 := []int{-9223372036854775808}
        _ = v56
    }
    if ((-1 <= 1) || true) {
        switch 4294967295 {
        case 32767:
            switch "\n\t" {
            case "":
                v57 := []int{127}
                _ = v57
            case "\n\t":
                _ = len(v4)
            default:
                var v58 int8 = -128
                _ = v58
            }
        case -698:
            for fi59 := 0; fi59 < 14; fi59++ {
                fmt.Println(fi59, v3, fi59)
                if (!true) {
                    for fi60 := 0; fi60 < 15; fi60++ {
                        fmt.Println(v3)
                        _ = fi60
                    }
                }
                _ = fi59
            }
        case 0:
            for fi61 := 0; fi61 < 10; fi61++ {
                switch -999 {
                case 2147483647:
                    if true {
                        v2 = (v2 >> ((282) & 63))
                        var v62 int8 = 42
                        _ = v62
                    } else {
                        v50 += (v50 * (v50 ^ v50))
                    }
                default:
                    var v63 int64 = 9223372036854775807
                    _ = v63
                }
                var v64 uint64 = v50
                _ = len(v4)
                _ = fi61
                _ = v64
            }
        default:
            v65 := GenS6{F0: "hello", F1: 32767, F2: -1}
            _ = v65
        }
    } else {
        delete(v1, "hello")
    }
    v66 := GenS4{F0: -128, F1: 1, F2: 32767}
    v67 := GenS7{F0: 128}
    v2 += (v3 ^ 2147483647)
    v50 = 1
    switch 2147483647 {
    case 715:
        var v68 bool = (!(!((65535 < 26) || false)))
        _ = v68
    default:
        var v69 int32 = genFn1(-2147483648, genFn1(genFn1(-1, -1, 2), -2147483648, 42), 471)
        _ = v69
    }
    var v70 int64 = 0
    v71 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v50:", v50)
    fmt.Println("v70:", v70)
    fmt.Println("slicelen v71:", len(v71), v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v50
    _ = v66
    _ = v67
    _ = v70
    _ = v71
}
