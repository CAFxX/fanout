package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 int32; F1 int8 }
type GenS6 struct { F0 string; F1 string }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int32; F1 int16; F2 int32 }
func genFn0(p0 int16, p1 uint64, p2 int64) int32 {
    var v6 int = -144
    v7 := []int{(^(v6)), (v6 | (v6 - -9223372036854775808)), (v6 - (v6 ^ v6)), (-(v6))}
    if ((!(4294967295 > v6)) || (!(true && (2147483648 > 42)))) {
        if (-1 > 42) {
            v7 = append(v7, (v6 >> ((255) & 63)))
            if len(v7) > 0 {
                v7[0] = (^(v6))
            }
        } else {
            fmt.Println(p2)
            v8 := GenS1{F0: (p2 + (p2 % ((-9223372036854775808) | 1))), F1: p2, F2: (p2 * -9223372036854775808)}
            switch ("a" + "a") {
            case "\n\t":
                for fi9 := 0; fi9 < 18; fi9++ {
                    if len(v7) > 0 {
                        v7[0] = 2147483647
                    }
                    _ = fi9
                }
            case "foo bar":
                v10 := GenS1{F0: p2, F1: p2, F2: (p2 & (p2 * (p2 * int64(p0))))}
                _ = v10
            }
            _ = v8
        }
    } else {
        v7 = append(v7, 718)
        v6 = int(p1)
    }
    if (true || false) {
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := GenS2{F0: (p2 | (p2 | p2))}
            _ = fi11
            _ = v12
        }
    } else {
        var v13 uint32 = 102
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        var v15 float32 = ((float32((1.5 + ((((0.1 + (0.1 - float64(-2.25))) * 0.0) - -2.25) - 0.0))) * -2.25) + (1.5 + 1.5))
        if (!(p2 == (p2 / (((p2 * 852)) | 1)))) {
            for fi16 := 0; fi16 < 7; fi16++ {
                p2 = p2
                _ = fi16
            }
            v15 = v15
        } else {
            v6 += v6
        }
        if false {
            for fi17 := 0; fi17 < 13; fi17++ {
                v18 := GenS3{F0: "foo bar", F1: false}
                v19 := GenS4{F0: 307, F1: 65535}
                var v20 uint = 735
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
            v21 := GenS5{F0: 841, F1: 0}
            v22 := GenS1{F0: p2, F1: (p2 << ((269) & 63)), F2: (p2 - (^(p2)))}
            _ = v21
            _ = v22
        } else {
            v23 := GenS2{F0: p2}
            fmt.Println(v15, v6)
            var v24 uint16 = 65535
            _ = v23
            _ = v24
        }
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 float64 = 1.5
    v2 := make(map[uint32]uint32)
    var v3 int16 = 84
    v4 := make(map[int]int)
    var v5 uint64 = 371
    switch 0 {
    case -565:
        if (2147483648 >= 0) {
            if true {
                if false {
                    for fi25 := 0; fi25 < 20; fi25++ {
                        var v26 int64 = 1
                        _ = fi25
                        _ = v26
                    }
                    for fi27 := 0; fi27 < 19; fi27++ {
                        var v28 int64 = -1
                        if (!((false && false) && false)) {
                            var v29 uint64 = 2
                            fmt.Println(fi27)
                            _ = v29
                        }
                        v28 = (v28 - -688)
                        _ = fi27
                        _ = v28
                    }
                    v1 = 0.0
                }
            }
        } else {
            fmt.Println(v3, v5)
            switch 611 {
            case 2147483647:
                var v30 uint32 = 42
                _ = v30
            case 42:
                v31 := GenS6{F0: ("" + "foo bar"), F1: "a"}
                _ = v31
            case 621:
                switch "hello" {
                case "foo bar":
                    v1 += 3.14159
                case "\n\t":
                    v32 := []int{-186}
                    _ = v32
                case "hello":
                    switch ("\n\t" + "") {
                    case "":
                        v1 = float64((1e10 * 1e10))
                    case "foo bar":
                        if false {
                            for fi33 := 0; fi33 < 18; fi33++ {
                                if false {
                                    var v34 string = "hello"
                                    v1 = v1
                                    if true {
                                        v35 := GenS7{F0: 212}
                                        var v36 int8 = 42
                                        for fi37 := 0; fi37 < 8; fi37++ {
                                            v38 := GenS4{F0: 255, F1: 65535}
                                            var v39 float32 = 0.1
                                            _ = fi37
                                            _ = v38
                                            _ = v39
                                        }
                                        _ = v35
                                        _ = v36
                                    }
                                    _ = v34
                                } else {
                                    for fi40 := 0; fi40 < 16; fi40++ {
                                        v3 += -181
                                        fmt.Println(fi40, fi40, fi33)
                                        _ = fi40
                                    }
                                }
                                v41 := GenS8{F0: 2147483647, F1: (v3 * v3), F2: genFn0(v3, v5, -9223372036854775808)}
                                _ = fi33
                                _ = v41
                            }
                        }
                    case "\n\t":
                        fmt.Println(v5)
                    }
                default:
                    delete(v2, 2147483647)
                }
            }
        }
    default:
        if true {
            var v42 float64 = (((v1 + v1) * (-2.25 + (v1 * 1.5))) + float64(1e10))
            var v43 int32 = -570
            _ = v42
            _ = v43
        } else {
            v44 := []int{}
            _ = v44
        }
    }
    for fi45 := 0; fi45 < 14; fi45++ {
        v3 += v3
        _ = fi45
    }
    var v46 int8 = 100
    var v47 uint64 = v5
    if true {
        v48 := GenS4{F0: 1, F1: 65535}
        v49 := v4[42]
        _ = v48
        _ = v49
    }
    var v50 float32 = 0.0
    var v51 uint64 = 1
    for fi52 := 0; fi52 < 1; fi52++ {
        v53 := GenS2{F0: 1}
        v54 := v2[2147483648]
        var v55 int = int(v3)
        _ = fi52
        _ = v53
        _ = v54
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v50
    _ = v51
}
