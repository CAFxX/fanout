package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 bool }
func genFn0() bool {
    var v7 int16 = -32768
    v7 += (v7 ^ (v7 << ((1) & 63)))
    _ = v7
    return false
}
func genFn1(p0 uint) int16 {
    if (3.14159 > 0.0) {
        if false {
            p0 += (p0 ^ p0)
            if (105 <= 217) {
                p0 = (p0 >> ((p0) & 63))
            }
            var v8 uint64 = 764
            _ = v8
        } else {
            v9 := []int{32767, 32767}
            _ = v9
        }
        p0 = p0
    }
    if false {
        switch 372 {
        case 4294967295:
            var v10 uint16 = 42
            _ = v10
        case 1:
            v11 := GenS1{F0: 1}
            _ = v11
        case -163:
            p0 = p0
        default:
            var v12 uint = (p0 & p0)
            _ = v12
        }
        v13 := []int{9223372036854775807, 1, -86}
        p0 -= 212
        _ = v13
    }
    v14 := GenS1{F0: 127}
    _ = p0
    _ = v14
    return -220
}

func main() {
    var v1 uint = 953
    var v2 int = 0
    var v3 uint32 = 42
    var v4 bool = true
    var v5 float64 = 3.14159
    var v6 int = (v2 << (((v1 | (^(v1)))) & 63))
    var v15 uint16 = 65535
    if v4 {
        for fi16 := 0; fi16 < 3; fi16++ {
            fmt.Println(fi16, v4, v5)
            v17 := []int{(fi16 >> ((v1) & 63)), 128, fi16, 9223372036854775807}
            _ = fi16
            _ = v17
        }
        fmt.Println(v15, v2, v3)
        var v18 string = "a"
        _ = v18
    } else {
        for fi19 := 0; fi19 < 20; fi19++ {
            if false {
                v6 = 32767
            } else {
                v20 := GenS1{F0: 128}
                var v21 int8 = 127
                _ = v20
                _ = v21
            }
            _ = fi19
        }
        for fi22 := 0; fi22 < 8; fi22++ {
            v5 = 1.5
            v23 := GenS2{F0: v4}
            _ = fi22
            _ = v23
        }
        v5 += 1.5
    }
    if v4 {
        fmt.Println(v6)
        var v24 int16 = 1000
        if ((!v4) || v4) {
            var v25 bool = (!false)
            _ = v25
        }
        _ = v24
    }
    v26 := GenS2{F0: v4}
    var v27 int16 = 1000
    var v28 uint32 = v3
    var v29 int64 = -1
    v27 -= (v27 ^ (v27 << ((v1) & 63)))
    for fi30 := 0; fi30 < 18; fi30++ {
        fmt.Println(v3)
        _ = fi30
    }
    fmt.Println(v3, v1)
    v31 := GenS2{F0: (!v4)}
    v32 := []int{32767}
    v4 = false
    var v33 float32 = float32(0.0)
    var v34 uint = v1
    switch "foo bar" {
    case "foo bar":
        v35 := GenS3{F0: (^(v27))}
        _ = v35
    case "hello":
        switch 464 {
        case -9223372036854775808:
            fmt.Println(v2, v4, v29)
        case 234:
            v6 = v2
        case 130:
            if (genFn1(v1) != (v27 << ((v34) & 63))) {
                v27 = genFn1(18446744073709551615)
            } else {
                var v36 string = "hello"
                var v37 float64 = (v5 + ((v5 / float64(((0.0 - v33) - v33))) * (v5 - 1.5)))
                _ = v36
                _ = v37
            }
        default:
            if false {
                if v4 {
                    v2 -= v2
                    fmt.Println(v27)
                    if ((!false) && (0.1 < v5)) {
                        fmt.Println(v2)
                        for fi38 := 0; fi38 < 1; fi38++ {
                            _ = len(v32)
                            v2 = (fi38 + v2)
                            _ = fi38
                        }
                        switch "\n\t" {
                        case "":
                            var v39 uint = 0
                            _ = v39
                        case "\n\t":
                            if (v4 || (true || ((-(v29)) > -9223372036854775808))) {
                                v6 = v2
                            } else {
                                for fi40 := 0; fi40 < 15; fi40++ {
                                    fmt.Println(v29, v28, v4)
                                    _ = fi40
                                }
                                var v41 float64 = ((float64((v33 - v33)) / v5) / 1.5)
                                for fi42 := 0; fi42 < 4; fi42++ {
                                    v43 := GenS1{F0: 0}
                                    _ = fi42
                                    _ = v43
                                }
                                _ = v41
                            }
                        case "foo bar":
                            fmt.Println(v15)
                        default:
                            v44 := GenS2{F0: v4}
                            _ = v44
                        }
                    }
                }
            }
        }
    case "":
        for fi45 := 0; fi45 < 3; fi45++ {
            v32 = append(v32, v6)
            for fi46 := 0; fi46 < 10; fi46++ {
                var v47 uint32 = 285
                fmt.Println(v2, v33, v5)
                if (-1 > 127) {
                    if (100 != -15) {
                        v5 = 1e10
                        v48 := GenS1{F0: 128}
                        v27 += genFn1((v1 - v34))
                        _ = v48
                    } else {
                        fmt.Println(v29, fi46, v28)
                        switch "a" {
                        case "\n\t":
                            v15 = (v15 & (v15 / ((v15) | 1)))
                        case "foo bar":
                            var v49 uint = v34
                            _ = v49
                        case "hello":
                            var v50 uint16 = (v15 % (((v15 >> (((v1 + v1)) & 63))) | 1))
                            _ = v50
                        }
                    }
                    v51 := GenS4{F0: ("\n\t" + "hello"), F1: ("a" + "a")}
                    v52 := GenS5{F0: v4}
                    v32 = append(v32, fi46)
                    _ = v51
                    _ = v52
                }
                v27 += (v27 << ((v34) & 63))
                _ = fi46
                _ = v47
            }
            _ = fi45
        }
    default:
        v53 := GenS4{F0: "hello", F1: "hello"}
        _ = v53
    }
    if (128 == 42) {
        if true {
            v3 += (^(v28))
            for fi54 := 0; fi54 < 9; fi54++ {
                switch 326 {
                case -567:
                    v55 := GenS3{F0: (v27 >> ((v34) & 63))}
                    _ = v55
                case 702:
                    v56 := GenS2{F0: (67 > -1)}
                    _ = v56
                }
                v3 += 2147483648
                _ = fi54
            }
        }
        if (v4 || (true && (127 < -1))) {
            if (false && true) {
                for fi57 := 0; fi57 < 11; fi57++ {
                    var v58 float32 = (float32(v5) * (((3.14159 * v33) / (v33 * float32(3.14159))) * float32(v5)))
                    v27 -= v27
                    _ = fi57
                    _ = v58
                }
                v1 -= v1
                var v59 uint16 = (v15 % (((v15 ^ v15)) | 1))
                _ = v59
            }
        } else {
            v60 := GenS2{F0: v4}
            v27 += -1
            v5 -= 1.5
            _ = v60
        }
    } else {
        _ = len(v32)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
