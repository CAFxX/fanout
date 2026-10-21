package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 string; F1 string; F2 uint64 }
type GenS3 struct { F0 int; F1 uint; F2 int }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint32, p1 int8, p2 int8) uint64 {
    fmt.Println(p1, p1)
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := GenS1{F0: 0, F1: 2}
        for fi9 := 0; fi9 < 18; fi9++ {
            p2 = p2
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    var v10 uint = 209
    if ((0.0 > 1e10) || (((!((p0 >= p0) && (true && true))) && (126 >= 0)) || (-839 > 1))) {
        switch -1 {
        case 43:
            v10 = (v10 ^ (v10 + uint(p0)))
        case 42:
            if (!false) {
                v11 := GenS2{F0: ("" + "foo bar"), F1: "", F2: 0}
                _ = v11
            } else {
                v10 += 2
            }
        default:
            v12 := []int{4294967295}
            _ = v12
        }
        var v13 float32 = 0.1
        p0 = (^(p0))
        _ = v13
    }
    switch ("\n\t" + "a") {
    case "\n\t":
        v14 := []int{122, -9223372036854775808, 32768}
        _ = v14
    case "hello":
        if false {
            v15 := GenS1{F0: 890, F1: 314}
            for fi16 := 0; fi16 < 6; fi16++ {
                for fi17 := 0; fi17 < 15; fi17++ {
                    fmt.Println(fi16)
                    var v18 float32 = 1e10
                    v19 := []int{(fi17 << (((v10 - v10)) & 63))}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                v20 := []int{(fi16 | 516), 127}
                _ = fi16
                _ = v20
            }
            _ = v15
        } else {
            if true {
                v10 = (^(v10))
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 4294967295
}
func genFn1(p0 uint32, p1 int32) uint16 {
    p0 -= (p0 << ((65535) & 63))
    if false {
        switch ("a" + "") {
        case "\n\t":
            if true {
                if false {
                    var v21 string = (("\n\t" + ("foo bar" + ("foo bar" + "a"))) + "\n\t")
                    v22 := []int{9223372036854775807, 1, -149, 32768}
                    for fi23 := 0; fi23 < 15; fi23++ {
                        v22 = append(v22, (-(fi23)))
                        _ = len(v22)
                        v21 = "\n\t"
                        _ = fi23
                    }
                    _ = v21
                    _ = v22
                }
                for fi24 := 0; fi24 < 2; fi24++ {
                    p1 = (-(p1))
                    switch (fi24 | fi24) {
                    case -1:
                        v25 := []int{}
                        _ = v25
                    case -41:
                        v26 := GenS3{F0: (fi24 | fi24), F1: 0, F2: (fi24 ^ (fi24 ^ int(p1)))}
                        _ = v26
                    default:
                        switch "" {
                        case "":
                            fmt.Println(fi24, p0)
                        case "foo bar":
                            fmt.Println(fi24)
                        case "a":
                            for fi27 := 0; fi27 < 1; fi27++ {
                                if false {
                                    var v28 float64 = ((0.1 / 1.0) - (((0.1 * ((1e10 - -2.25) + -2.25)) + 1.5) - 3.14159))
                                    _ = v28
                                }
                                v29 := GenS4{F0: 988}
                                fmt.Println(fi24, p1, p0)
                                var v30 int8 = 0
                                _ = fi27
                                _ = v29
                                _ = v30
                            }
                        default:
                            var v31 int64 = 42
                            _ = v31
                        }
                    }
                    p1 = p1
                    _ = fi24
                }
            }
        case "a":
            switch 4294967295 {
            case 0:
                v32 := GenS2{F0: ("a" + "foo bar"), F1: "\n\t", F2: 0}
                _ = v32
            case -75:
                fmt.Println(p1, p1)
            case -1:
                for fi33 := 0; fi33 < 1; fi33++ {
                    for fi34 := 0; fi34 < 13; fi34++ {
                        var v35 float64 = (1e10 / 3.14159)
                        if (float64((1e10 / 1e10)) != v35) {
                            v36 := GenS2{F0: ("hello" + "\n\t"), F1: "", F2: 4294967295}
                            p0 = p0
                            p0 += 4294967295
                            _ = v36
                        }
                        _ = fi34
                        _ = v35
                    }
                    p0 -= (p0 & p0)
                    _ = fi33
                }
            default:
                for fi37 := 0; fi37 < 2; fi37++ {
                    if false {
                        switch "foo bar" {
                        case "hello":
                            p1 = p1
                        case "foo bar":
                            if (769 <= 0) {
                                p1 = p1
                                var v38 uint = 0
                                fmt.Println(p0, fi37)
                                _ = v38
                            }
                        }
                        for fi39 := 0; fi39 < 9; fi39++ {
                            v40 := []int{}
                            p0 = 427
                            _ = fi39
                            _ = v40
                        }
                        var v41 int8 = 42
                        _ = v41
                    }
                    _ = fi37
                }
            }
        case "hello":
            p1 += 1
        }
        v42 := GenS4{F0: p1}
        _ = v42
    }
    if (!(!(((p0 << ((255) & 63)) == p0) && true))) {
        p0 += p0
    }
    _ = p0
    _ = p1
    return 32767
}

func main() {
    var v1 int = -948
    var v2 uint64 = 2
    var v3 uint64 = v2
    v4 := make(map[string]uint32)
    var v5 string = ""
    v6 := []int{v1, v1, (v1 / (((^(v1))) | 1)), 42}
    if false {
        if ((!false) && (!(658 >= 1))) {
            v5 = (v5 + (v5 + v5))
            for fi43 := 0; fi43 < 16; fi43++ {
                v3 -= 627
                _ = fi43
            }
        } else {
            v44 := GenS1{F0: v3, F1: 255}
            _ = v44
        }
    } else {
        v3 += (^(v2))
    }
    switch 32768 {
    case -238:
        for fi45 := 0; fi45 < 5; fi45++ {
            if ((true || false) || (true && ((!false) || (true && (520 > (^(v3))))))) {
                v5 = v5
                var v46 float32 = (0.1 + 0.1)
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v46
            } else {
                var v47 float64 = 1e10
                _ = v47
            }
            _ = fi45
        }
    default:
        var v48 float64 = 0.1
        _ = v48
    }
    v49 := GenS1{F0: (^(v3)), F1: genFn1(1, -953)}
    v50 := GenS1{F0: v3, F1: 42}
    var v51 int16 = 0
    var v52 uint8 = 128
    _ = len(v4)
    v3 += (v3 % ((genFn0(0, int8(v52), 0)) | 1))
    var v53 uint = 65535
    var v54 int = (v1 / (((-(v1))) | 1))
    fmt.Println(v53, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v6
}
