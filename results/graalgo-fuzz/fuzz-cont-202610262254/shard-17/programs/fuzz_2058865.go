package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 string; F1 uint64; F2 uint32 }
type GenS5 struct { F0 bool; F1 int8 }
type GenS6 struct { F0 uint; F1 uint32; F2 uint8 }
func genFn0(p0 uint64, p1 uint32, p2 int) uint8 {
    v5 := GenS1{F0: (p1 + p1), F1: -437, F2: p2}
    if false {
        p0 = 42
        v6 := GenS1{F0: (p1 | (p1 >> ((4294967295) & 63))), F1: 777, F2: p2}
        if false {
            v7 := GenS1{F0: p1, F1: 0, F2: 248}
            if true {
                var v8 uint32 = (p1 - 203)
                _ = v8
            } else {
                p2 = 32768
                var v9 int8 = -1
                _ = v9
            }
            v10 := []int{(p2 * p2), (p2 * (p2 >> ((2) & 63))), p2, int(p0)}
            _ = v10
            _ = v7
        }
        _ = v6
    }
    if (4294967295 <= 42) {
        v11 := GenS1{F0: (p1 + p1), F1: 9223372036854775807, F2: p2}
        _ = v11
    } else {
        v12 := []int{(p2 << ((0) & 63)), (p2 + (p2 + -219)), p2}
        var v13 uint = 516
        var v14 uint = v13
        _ = v12
        _ = v13
        _ = v14
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 222
}

func main() {
    var v1 uint8 = 128
    v2 := []uint32{4294967295, 2147483647, 338, 236}
    var v3 uint16 = 1
    v4 := make(map[int]uint32)
    if false {
        var v15 uint16 = (^(v3))
        v1 -= 42
        _ = v15
    } else {
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        if ((!("hello" <= (("hello" + (("" + (("" + "a") + "\n\t")) + "foo bar")) + "foo bar"))) || true) {
            for fi16 := 0; fi16 < 13; fi16++ {
                switch "foo bar" {
                case "a":
                    v17 := v4[(fi16 + fi16)]
                    _ = v17
                }
                v18 := GenS2{F0: "", F1: false}
                _ = fi16
                _ = v18
            }
        }
        for fi19 := 0; fi19 < 9; fi19++ {
            for fi20 := 0; fi20 < 5; fi20++ {
                v21 := GenS3{F0: (!(true || true)), F1: (v1 < (v1 / (((v1 >> ((42) & 63))) | 1)))}
                _ = fi20
                _ = v21
            }
            switch "\n\t" {
            case "hello":
                var v22 int = fi19
                _ = v22
            case "foo bar":
                if ("" <= "") {
                    var v23 int8 = 127
                    _ = v23
                }
            }
            var v24 uint8 = v1
            _ = fi19
            _ = v24
        }
    }
    var v25 float64 = 0.1
    switch "hello" {
    case "foo bar":
        var v26 uint = 1
        _ = v26
    case "hello":
        var v27 int64 = 9223372036854775807
        _ = v27
    case "\n\t":
        switch "hello" {
        case "\n\t":
            var v28 uint8 = v1
            _ = v28
        case "hello":
            if len(v2) > 0 {
                v2[0] = 2147483648
            }
        default:
            var v29 uint16 = v3
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        v1 -= v1
        _ = fi30
    }
    v3 -= v3
    v31 := GenS1{F0: 2147483647, F1: -1, F2: -9223372036854775808}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi32 := 0; fi32 < 18; fi32++ {
        var v33 string = "foo bar"
        if true {
            v2 = append(v2, 0)
        }
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 18; fi34++ {
        v35 := GenS3{F0: true, F1: true}
        v1 -= (^(v1))
        _ = fi34
        _ = v35
    }
    for fi36 := 0; fi36 < 2; fi36++ {
        switch fi36 {
        case -2147483648:
            for fi37 := 0; fi37 < 11; fi37++ {
                fmt.Println(v1, v25, v3)
                switch 42 {
                case -2147483648:
                    if true {
                        var v38 string = "a"
                        _ = v38
                    }
                case -499:
                    switch (fi36 / ((fi36) | 1)) {
                    case -409:
                        v25 += 1.5
                    case 42:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                default:
                    if false {
                        if ((fi36 <= (fi37 ^ -9223372036854775808)) || (2 >= 4294967295)) {
                            v3 -= v3
                            if true {
                                switch (^(fi37)) {
                                case 1:
                                    switch fi37 {
                                    case -990:
                                        v39 := GenS4{F0: "", F1: uint64(v3), F2: 760}
                                        _ = v39
                                    case 42:
                                        v25 += v25
                                    case 1:
                                        var v40 int32 = 1
                                        _ = v40
                                    default:
                                        if len(v2) > 0 {
                                            v2[0] = 4294967295
                                        }
                                    }
                                }
                                switch "\n\t" {
                                case "a":
                                    v41 := GenS3{F0: (0.1 < v25), F1: false}
                                    _ = v41
                                }
                            }
                        }
                    } else {
                        fmt.Println(v1, fi37, fi36)
                        switch -9223372036854775808 {
                        case 4294967295:
                            v1 = (^(v1))
                        case -2147483648:
                            v25 -= 1.5
                        case -571:
                            var v42 uint16 = v3
                            _ = v42
                        default:
                            var v43 int16 = 32767
                            _ = v43
                        }
                        if false {
                            v1 = genFn0(1, 0, fi37)
                            v44 := v4[-220]
                            if ((true && ((v44 | v44) == (v44 ^ (v44 ^ (v44 >> ((uint(v44)) & 63)))))) && true) {
                                var v45 uint64 = 4294967295
                                v25 -= v25
                                v46 := GenS5{F0: false, F1: 0}
                                if ((v44 ^ v44) < uint32(v1)) {
                                    v47 := GenS6{F0: 28, F1: (v44 - (v44 >> ((587) & 63))), F2: (v1 - (^(v1)))}
                                    v1 += (v1 * v1)
                                    _ = v47
                                } else {
                                    var v48 int32 = -1
                                    delete(v4, 2147483647)
                                    v49 := GenS5{F0: true, F1: 127}
                                    _ = v48
                                    _ = v49
                                }
                                fmt.Println(v25)
                                _ = v45
                                _ = v46
                            } else {
                                if (v44 <= (v44 / ((977) | 1))) {
                                    v44 -= (^(v44))
                                    var v50 int8 = 42
                                    if (65535 > 65535) {
                                        var v51 float32 = float32(0.0)
                                        v52 := GenS3{F0: ((1.5 / 3.14159) < (v25 * v25)), F1: ("\n\t" == "a")}
                                        _ = v51
                                        _ = v52
                                    }
                                    v53 := GenS4{F0: ("" + "hello"), F1: 2, F2: v44}
                                    fmt.Println(fi36, v25, v50)
                                    if len(v2) > 0 {
                                        v2[0] = (v44 << ((255) & 63))
                                    }
                                    var v54 uint32 = 1
                                    v25 = v25
                                    v55 := v4[(^(fi36))]
                                    _ = v50
                                    _ = v53
                                    _ = v54
                                    _ = v55
                                }
                            }
                            _ = v44
                        }
                    }
                }
                _ = fi37
            }
        case -535:
            fmt.Println(v25, v3, v25)
        default:
            var v56 int8 = 127
            _ = v56
        }
        _ = fi36
    }
    v57 := v4[-1]
    switch 32767 {
    case 4294967295:
        fmt.Println(v3, v1)
    }
    v1 = v1
    var v58 int = 127
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v4
    _ = v57
    _ = v58
}
