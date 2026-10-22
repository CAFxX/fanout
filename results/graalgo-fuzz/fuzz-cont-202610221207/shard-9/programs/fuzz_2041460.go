package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int32; F1 int32 }

func main() {
    var v1 uint64 = 0
    var v2 uint32 = 4294967295
    v3 := make(map[string]int)
    v1 = v1
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := []int{fi4}
        for fi6 := 0; fi6 < 2; fi6++ {
            if true {
                v7 := GenS1{F0: (fi4 - fi6), F1: "a"}
                fmt.Println(fi6)
                v8 := GenS1{F0: fi4, F1: "foo bar"}
                _ = v7
                _ = v8
            }
            if ((v1 >> ((65535) & 63)) >= v1) {
                var v9 uint16 = 0
                v10 := GenS2{F0: 1}
                fmt.Println(fi6, v9, v1)
                _ = v10
                _ = v9
            } else {
                v11 := GenS2{F0: 0}
                v1 += (v1 >> ((0) & 63))
                _ = v11
            }
            var v12 string = "\n\t"
            _ = fi6
            _ = v12
        }
        fmt.Println(v1, v2, v1)
        _ = fi4
        _ = v5
    }
    v1 = 18446744073709551615
    v13 := []int{2147483647, -761}
    var v14 int64 = 0
    fmt.Println(v2, v1)
    if true {
        var v15 float64 = 0.0
        for fi16 := 0; fi16 < 20; fi16++ {
            if (2 <= (v1 | v1)) {
                v17 := GenS3{F0: 1}
                var v18 uint = 65535
                _ = v17
                _ = v18
            } else {
                v1 = (^(v1))
            }
            v14 -= int64(fi16)
            _ = fi16
        }
        _ = v15
    } else {
        v14 = v14
        for fi19 := 0; fi19 < 14; fi19++ {
            switch ((("a" + "\n\t") + "") + "hello") {
            case "hello":
                if true {
                    for fi20 := 0; fi20 < 7; fi20++ {
                        fmt.Println(fi20, v14)
                        var v21 uint8 = uint8(v1)
                        _ = fi20
                        _ = v21
                    }
                    if false {
                        for fi22 := 0; fi22 < 4; fi22++ {
                            var v23 uint8 = 42
                            v2 = (^(v2))
                            fmt.Println(v14)
                            switch 32768 {
                            case 32767:
                                var v24 uint = 65535
                                v25 := GenS1{F0: (fi19 % ((fi22) | 1)), F1: ""}
                                fmt.Println(v14, v2, v1)
                                if false {
                                    switch "a" {
                                    case "hello":
                                        switch fi19 {
                                        case 4294967295:
                                            fmt.Println(v14)
                                        default:
                                            v1 += 653
                                            v26 := GenS2{F0: -525}
                                            _ = v26
                                        }
                                    default:
                                        if (!(1000 > 0)) {
                                            var v27 uint16 = 32767
                                            _ = v27
                                        }
                                        v28 := GenS4{F0: uint16(v23), F1: ((("\n\t" + "a") + ("foo bar" + "\n\t")) + "hello")}
                                        v2 += v2
                                        _ = v28
                                    }
                                }
                                v1 -= 4294967295
                                _ = v24
                                _ = v25
                            case 802:
                                _ = len(v3)
                            case 0:
                                var v29 bool = ((3.14159 - (3.14159 / 0.1)) == 0.0)
                                _ = v29
                            default:
                                delete(v3, ("foo bar" + "a"))
                                var v30 float64 = 1.5
                                v31 := GenS3{F0: 255}
                                _ = v30
                                _ = v31
                            }
                            switch "foo bar" {
                            case "\n\t":
                                if len(v13) > 0 {
                                    v13[0] = (fi19 >> ((152) & 63))
                                }
                            case "a":
                                fmt.Println(fi22)
                            case "foo bar":
                                var v32 uint16 = 0
                                _ = v32
                            }
                            _ = fi22
                            _ = v23
                        }
                        for fi33 := 0; fi33 < 15; fi33++ {
                            switch "" {
                            case "":
                                switch "\n\t" {
                                case "\n\t":
                                    if ((!((v1 | (v1 * (v1 & v1))) != v1)) && ((v1 <= uint64(v2)) && (0 <= -1))) {
                                        v34 := GenS1{F0: fi19, F1: "a"}
                                        var v35 int32 = -2147483648
                                        _ = v34
                                        _ = v35
                                    }
                                default:
                                    switch 9223372036854775807 {
                                    case -9223372036854775808:
                                        switch "hello" {
                                        case "a":
                                            fmt.Println(v1, v14)
                                        case "foo bar":
                                            v36 := GenS2{F0: 603}
                                            _ = v36
                                        case "":
                                            v1 -= v1
                                        }
                                    default:
                                        var v37 int64 = v14
                                        _ = v37
                                    }
                                }
                            default:
                                if true {
                                    fmt.Println(fi19, v2)
                                }
                            }
                            _ = fi33
                        }
                        _ = len(v3)
                    }
                }
            case "":
                switch fi19 {
                case 0:
                    v38 := v3[("\n\t" + "")]
                    _ = v38
                case 42:
                    v39 := GenS1{F0: (fi19 % ((-9223372036854775808) | 1)), F1: "a"}
                    _ = v39
                default:
                    switch fi19 {
                    case 0:
                        for fi40 := 0; fi40 < 13; fi40++ {
                            delete(v3, ("\n\t" + (("a" + (("" + "hello") + ("hello" + ("foo bar" + ("\n\t" + ""))))) + ((("\n\t" + "\n\t") + "a") + ("a" + "\n\t")))))
                            v41 := v3["a"]
                            fmt.Println(fi40, fi19)
                            _ = fi40
                            _ = v41
                        }
                    }
                }
            case "a":
                v42 := GenS2{F0: 0}
                _ = v42
            default:
                var v43 uint16 = 4
                _ = v43
            }
            _ = len(v3)
            for fi44 := 0; fi44 < 15; fi44++ {
                v45 := GenS5{F0: v14}
                _ = fi44
                _ = v45
            }
            _ = fi19
        }
    }
    for fi46 := 0; fi46 < 2; fi46++ {
        switch "\n\t" {
        case "foo bar":
            if (fi46 != (-(fi46))) {
                v47 := GenS6{F0: (^(v1))}
                if ((v2 / ((v2) | 1)) < v2) {
                    if (526 == fi46) {
                        var v48 int16 = -32768
                        var v49 uint32 = 526
                        if (-2147483648 == 1) {
                            for fi50 := 0; fi50 < 11; fi50++ {
                                var v51 int32 = 0
                                v52 := GenS7{F0: (^(v51)), F1: v51}
                                _ = fi50
                                _ = v51
                                _ = v52
                            }
                            for fi53 := 0; fi53 < 18; fi53++ {
                                v2 = 2147483647
                                _ = fi53
                            }
                        } else {
                            var v54 uint16 = 32767
                            _ = v54
                        }
                        _ = v48
                        _ = v49
                    }
                } else {
                    var v55 uint16 = 42
                    fmt.Println(v1, v2)
                    _ = v55
                }
                _ = v47
            }
        case "a":
            for fi56 := 0; fi56 < 7; fi56++ {
                _ = len(v3)
                if (false || false) {
                    var v57 int16 = -559
                    if (!(((v57 != v57) || ((((42 != (v57 << ((255) & 63))) || false) || false) && true)) && (v2 > (^(v2))))) {
                        fmt.Println(v1)
                        v13 = append(v13, (fi46 & int(v2)))
                    }
                    _ = v57
                } else {
                    var v58 float64 = 1e10
                    v3["\n\t"] = fi46
                    _ = v58
                }
                _ = fi56
            }
        case "\n\t":
            switch "" {
            case "\n\t":
                v1 += v1
            default:
                var v59 uint64 = v1
                _ = v59
            }
        }
        var v60 uint16 = 985
        if (1 < 42) {
            v1 = 18446744073709551615
            fmt.Println(fi46, v60)
        }
        _ = fi46
        _ = v60
    }
    for fi61 := 0; fi61 < 6; fi61++ {
        var v62 int = (-(fi61))
        var v63 uint = 1
        _ = fi61
        _ = v62
        _ = v63
    }
    if ((-2.25 / 1.0) == (3.14159 - (3.14159 + (-2.25 - (0.1 / 1.0))))) {
        var v64 bool = true
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
