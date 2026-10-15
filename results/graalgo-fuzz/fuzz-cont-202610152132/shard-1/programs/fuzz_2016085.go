package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int8; F1 uint; F2 uint }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 int64; F1 uint64 }
type GenS7 struct { F0 int16; F1 int8 }
func genFn0(p0 bool, p1 uint) uint {
    for fi4 := 0; fi4 < 6; fi4++ {
        for fi5 := 0; fi5 < 16; fi5++ {
            v6 := GenS1{F0: -1, F1: -749}
            var v7 bool = true
            p0 = v7
            _ = fi5
            _ = v6
            _ = v7
        }
        var v8 float64 = -2.25
        _ = fi4
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 uint16 = 255
        if (100 <= -1) {
            if p0 {
                var v11 uint8 = 42
                for fi12 := 0; fi12 < 13; fi12++ {
                    v13 := []int{42}
                    _ = fi12
                    _ = v13
                }
                if false {
                    v14 := []int{}
                    _ = v14
                }
                _ = v11
            } else {
                for fi15 := 0; fi15 < 13; fi15++ {
                    fmt.Println(p1, fi15)
                    v16 := GenS2{F0: -128, F1: 65535, F2: (^(p1))}
                    p1 += (p1 & p1)
                    _ = fi15
                    _ = v16
                }
                switch "\n\t" {
                case "foo bar":
                    var v17 uint = (p1 * p1)
                    _ = v17
                case "a":
                    v18 := GenS3{F0: (fi9 * fi9), F1: ((float32(0.0) + ((3.14159 - 0.0) / ((0.0 / 1.0) + 0.1))) <= 3.14159)}
                    _ = v18
                }
                if p0 {
                    fmt.Println(p0, p1)
                    switch "\n\t" {
                    case "a":
                        v19 := GenS4{F0: 32767, F1: ((("" + "") + "\n\t") + ("foo bar" + ""))}
                        _ = v19
                    case "":
                        var v20 uint8 = 195
                        _ = v20
                    default:
                        v21 := GenS2{F0: 100, F1: p1, F2: (p1 >> ((p1) & 63))}
                        _ = v21
                    }
                    p0 = p0
                }
            }
            p1 += (p1 % ((p1) | 1))
        } else {
            if p0 {
                var v22 float64 = 1e10
                v23 := GenS4{F0: fi9, F1: ("" + "foo bar")}
                switch "hello" {
                case "foo bar":
                    switch (fi9 * (fi9 | fi9)) {
                    case 128:
                        v24 := GenS2{F0: 94, F1: p1, F2: (p1 ^ (^(p1)))}
                        _ = v24
                    case 258:
                        if (42 <= 65535) {
                            for fi25 := 0; fi25 < 15; fi25++ {
                                v26 := GenS2{F0: -1, F1: (p1 | (p1 % ((p1) | 1))), F2: (p1 & (p1 * (^(p1))))}
                                fmt.Println(v10, v10)
                                _ = fi25
                                _ = v26
                            }
                            for fi27 := 0; fi27 < 2; fi27++ {
                                var v28 string = ""
                                var v29 uint = (p1 & uint(v10))
                                if (!(!false)) {
                                    v10 += v10
                                    p1 -= p1
                                    fmt.Println(v22, fi9)
                                }
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                            var v30 string = ("hello" + "hello")
                            _ = v30
                        } else {
                            p1 = 2
                            p1 += (p1 % ((p1) | 1))
                            if (true || ((p1 & p1) <= p1)) {
                                v31 := GenS5{F0: "hello", F1: (-1 < 42)}
                                if (-32768 < 1000) {
                                    var v32 string = "\n\t"
                                    _ = v32
                                }
                                _ = v31
                            }
                        }
                    default:
                        v33 := GenS2{F0: -122, F1: p1, F2: (p1 << ((p1) & 63))}
                        _ = v33
                    }
                case "":
                    switch ("hello" + ("a" + ("hello" + ("foo bar" + "")))) {
                    case "a":
                        if p0 {
                            v34 := []int{}
                            _ = v34
                        } else {
                            v10 = 65535
                            v22 += -2.25
                            var v35 float32 = float32(((((v22 - v22) - 0.0) - v22) * (((1e10 - float64(float32(v22))) - v22) - v22)))
                            _ = v35
                        }
                    case "":
                        p0 = p0
                    case "\n\t":
                        fmt.Println(v10, p0)
                    default:
                        var v36 int16 = -32768
                        _ = v36
                    }
                case "a":
                    if (0.0 > 0.1) {
                        fmt.Println(fi9, p0, p0)
                        for fi37 := 0; fi37 < 17; fi37++ {
                            var v38 int16 = -640
                            if false {
                                switch fi9 {
                                case 2147483647:
                                    var v39 uint64 = 699
                                    _ = v39
                                default:
                                    v40 := []int{783, (fi9 & fi37), 1}
                                    _ = v40
                                }
                                v10 = (v10 | 32767)
                                var v41 int32 = 2147483647
                                _ = v41
                            } else {
                                var v42 uint32 = 643
                                if false {
                                    v43 := GenS6{F0: -9223372036854775808, F1: 18446744073709551615}
                                    _ = v43
                                }
                                v44 := []int{fi9, (fi37 & fi9)}
                                _ = v42
                                _ = v44
                            }
                            _ = fi37
                            _ = v38
                        }
                    }
                default:
                    fmt.Println(v10)
                }
                _ = v22
                _ = v23
            } else {
                if (3.14159 >= 1.5) {
                    fmt.Println(p1, p0)
                }
                for fi45 := 0; fi45 < 16; fi45++ {
                    p0 = (2 >= 1)
                    _ = fi45
                }
                fmt.Println(v10, v10)
            }
            v46 := []int{(fi9 + fi9), (fi9 + fi9), (fi9 % ((0) | 1))}
            p1 -= p1
            _ = v46
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 float32 = 3.14159
    var v2 uint = 2
    v3 := []string{"hello", ""}
    if (v1 >= v1) {
        if (!false) {
            var v47 bool = (-103 >= 40)
            v48 := GenS7{F0: 42, F1: -128}
            v3 = append(v3, ("" + ("\n\t" + "")))
            _ = v47
            _ = v48
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v2, v1)
    } else {
        v1 += 1.5
        if len(v3) > 0 {
            v3[0] = "hello"
        }
    }
    for fi49 := 0; fi49 < 6; fi49++ {
        for fi50 := 0; fi50 < 19; fi50++ {
            var v51 uint64 = 0
            _ = fi50
            _ = v51
        }
        _ = fi49
    }
    var v52 uint32 = 671
    var v53 string = "\n\t"
    for fi54 := 0; fi54 < 12; fi54++ {
        v2 = (v2 / (((v2 & (v2 % ((v2) | 1)))) | 1))
        _ = fi54
    }
    _ = len(v3)
    switch -1 {
    case 1:
        if false {
            var v55 uint16 = 2
            _ = v55
        }
    case 42:
        v3 = append(v3, "hello")
    default:
        if ((42 == -2147483648) && (0.0 <= 0.1)) {
            switch v53 {
            case "a":
                for fi56 := 0; fi56 < 1; fi56++ {
                    var v57 int16 = 504
                    _ = fi56
                    _ = v57
                }
            default:
                v2 -= v2
            }
        } else {
            var v58 uint64 = 2
            _ = v58
        }
    }
    fmt.Println(v53, v53, v2)
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v52
    _ = v53
}
