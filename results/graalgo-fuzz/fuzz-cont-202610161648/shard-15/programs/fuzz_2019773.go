package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 int8; F2 int }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string; F1 int8 }
func genFn0(p0 uint64, p1 int32, p2 int16) int32 {
    fmt.Println(p2, p0, p2)
    var v7 int16 = (p2 ^ p2)
    var v8 uint16 = 577
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p1
}

func main() {
    v1 := []string{"a", "hello", "a", "foo bar"}
    v2 := make(map[int]bool)
    var v3 bool = (!false)
    var v4 uint32 = 2
    var v5 string = ("" + ("foo bar" + ""))
    var v6 bool = (-746 == 0)
    var v9 int16 = -32768
    v9 = (^(v9))
    for fi10 := 0; fi10 < 5; fi10++ {
        v5 = v5
        _ = fi10
    }
    var v11 bool = (!(!v3))
    switch "" {
    case "foo bar":
        var v12 int32 = -346
        _ = v12
    }
    var v13 uint64 = 18446744073709551615
    var v14 int32 = -237
    var v15 int64 = 1
    var v16 int32 = 0
    if v3 {
        var v17 uint32 = v4
        v18 := v2[42]
        if len(v1) > 0 {
            v1[0] = "hello"
        }
        _ = v17
        _ = v18
    } else {
        var v19 float32 = (-2.25 - (((-2.25 - 0.0) + (float32(float64(1e10)) / (-2.25 + 0.0))) + 0.0))
        var v20 int32 = (v16 / ((v14) | 1))
        _ = v19
        _ = v20
    }
    v21 := GenS1{F0: (v14 | genFn0((^(v13)), (v16 * v14), v9)), F1: (-(v15))}
    if (("hello" + (("a" + v5) + v5)) == v5) {
        if v3 {
            var v22 string = "a"
            v1 = append(v1, v22)
            _ = v22
        } else {
            v23 := GenS2{F0: v13}
            if v11 {
                var v24 uint32 = (^(v4))
                switch -841 {
                case -530:
                    v9 = 42
                case 127:
                    switch 4294967295 {
                    case 32767:
                        v25 := GenS3{F0: "a", F1: -93, F2: 4294967295}
                        _ = v25
                    case -692:
                        v15 = -9223372036854775808
                    default:
                        v24 = (v24 & v24)
                    }
                case -685:
                    fmt.Println(v15, v13)
                default:
                    if (-9223372036854775808 != 230) {
                        if (v11 || v3) {
                            v15 = (v15 & v15)
                            v3 = false
                            var v26 bool = v3
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 13; fi27++ {
                            var v28 float32 = float32(0.0)
                            _ = fi27
                            _ = v28
                        }
                    }
                }
                if v3 {
                    fmt.Println(v9, v11)
                    if true {
                        if v3 {
                            v29 := GenS4{F0: (v16 % ((1) | 1))}
                            var v30 float64 = 0.0
                            _ = v29
                            _ = v30
                        }
                        if (!(float64(3.14159) < 1e10)) {
                            v31 := GenS4{F0: (v16 ^ genFn0((v13 | v13), v16, 1))}
                            _ = v31
                        }
                    } else {
                        var v32 int = 9223372036854775807
                        for fi33 := 0; fi33 < 17; fi33++ {
                            v24 = (v4 | v4)
                            v2[(-(v32))] = false
                            v11 = ((v15 - (^(v15))) <= -9223372036854775808)
                            _ = fi33
                        }
                        _ = v32
                    }
                    if ((^(v4)) != v4) {
                        v34 := GenS4{F0: 2147483647}
                        _ = v34
                    }
                } else {
                    switch v5 {
                    case "foo bar":
                        switch 32767 {
                        case 32767:
                            v2[-1] = false
                        case 1:
                            if v11 {
                                v35 := v2[-964]
                                _ = v35
                            }
                        case -1:
                            v36 := GenS5{F0: v5, F1: -16}
                            _ = v36
                        default:
                            for fi37 := 0; fi37 < 19; fi37++ {
                                v2[(fi37 << ((65535) & 63))] = false
                                _ = fi37
                            }
                        }
                    }
                    for fi38 := 0; fi38 < 17; fi38++ {
                        for fi39 := 0; fi39 < 11; fi39++ {
                            v6 = (-19 < 119)
                            v4 -= v24
                            _ = fi39
                        }
                        _ = fi38
                    }
                    if v11 {
                        var v40 int16 = v9
                        _ = v40
                    }
                }
                _ = v24
            }
            _ = v23
        }
        for fi41 := 0; fi41 < 7; fi41++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi41
        }
        var v42 int = -1
        _ = v42
    } else {
        fmt.Println(v14)
        if len(v1) > 0 {
            v1[0] = (("\n\t" + (v5 + v5)) + "")
        }
    }
    var v43 uint8 = 128
    var v44 int16 = (v9 | v9)
    if ((v13 + v13) <= v13) {
        v14 = -750
        if false {
            v4 = 815
        } else {
            v15 -= (-(v15))
        }
        delete(v2, 32768)
    } else {
        v15 += (v15 & (v15 ^ v15))
        for fi45 := 0; fi45 < 5; fi45++ {
            delete(v2, fi45)
            _ = fi45
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v9
}
