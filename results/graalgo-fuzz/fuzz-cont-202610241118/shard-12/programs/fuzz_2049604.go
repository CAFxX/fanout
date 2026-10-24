package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint32; F1 string; F2 int64 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 uint }
func genFn0() int16 {
    var v7 uint8 = 11
    v8 := []int{42, -129, 1, 4294967295}
    _ = v7
    _ = v8
    return 32767
}
func genFn1(p0 int64, p1 int8, p2 bool) uint8 {
    p1 -= p1
    for fi9 := 0; fi9 < 13; fi9++ {
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 int16 = 42
            p2 = p2
            v12 := GenS1{F0: 979, F1: 65535, F2: 2147483648}
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 uint16 = 528
        _ = fi9
        _ = v13
    }
    p0 -= (p0 | p0)
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 uint = 239
    v2 := make(map[int]string)
    v3 := make(map[string]int)
    v4 := []string{"\n\t", "\n\t", "", ("hello" + ("\n\t" + ""))}
    var v5 int16 = -32768
    v6 := []string{}
    switch "a" {
    case "\n\t":
        v1 -= (v1 % ((65535) | 1))
    case "foo bar":
        for fi14 := 0; fi14 < 20; fi14++ {
            for fi15 := 0; fi15 < 2; fi15++ {
                var v16 int64 = int64(v1)
                _ = fi15
                _ = v16
            }
            v1 = (v1 ^ (v1 & (^(v1))))
            _ = len(v4)
            _ = fi14
        }
    default:
        if true {
            var v17 float64 = -2.25
            v17 -= v17
            v18 := GenS2{F0: 100, F1: 104}
            _ = v17
            _ = v18
        }
    }
    var v19 int32 = int32(v5)
    v20 := GenS1{F0: 0, F1: 686, F2: 65}
    var v21 uint8 = 89
    switch 2147483647 {
    case 0:
        var v22 float64 = (((0.0 / 1.5) / 1.0) / 1e10)
        _ = v22
    case 127:
        for fi23 := 0; fi23 < 3; fi23++ {
            fmt.Println(v21, v1)
            var v24 uint = (^(v1))
            fmt.Println(v19)
            _ = fi23
            _ = v24
        }
    }
    v19 = v19
    switch -713 {
    case 32767:
        v25 := GenS2{F0: 1, F1: -128}
        _ = v25
    case 32768:
        v1 = (v1 ^ (v1 / (((v1 - (v1 | (^(v1))))) | 1)))
    case 2147483647:
        _ = len(v3)
    default:
        var v26 uint8 = v21
        _ = v26
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        if false {
            v6 = append(v6, "hello")
        } else {
            var v28 uint64 = 2
            if true {
                for fi29 := 0; fi29 < 9; fi29++ {
                    v1 += (v1 | v1)
                    _ = fi29
                }
                v19 += -2147483648
            } else {
                var v30 int8 = 0
                if ((!(v28 < v28)) || (false && ((1.5 / float32(1e10)) <= 1e10))) {
                    v5 = v5
                }
                for fi31 := 0; fi31 < 18; fi31++ {
                    if (v19 < 713) {
                        var v32 uint8 = genFn1(0, -1, (fi27 > fi27))
                        switch (("a" + "") + ("hello" + "hello")) {
                        case "\n\t":
                            var v33 uint64 = 18446744073709551615
                            _ = v33
                        }
                        v34 := v2[(fi31 & 9223372036854775807)]
                        _ = v32
                        _ = v34
                    } else {
                        v19 += 0
                        switch "\n\t" {
                        case "a":
                            fmt.Println(v21, v28, fi31)
                        case "hello":
                            var v35 uint64 = v28
                            _ = v35
                        case "foo bar":
                            if true {
                                v36 := GenS1{F0: 0, F1: uint(v21), F2: 2147483647}
                                var v37 int32 = -836
                                _ = v36
                                _ = v37
                            } else {
                                switch "foo bar" {
                                case "\n\t":
                                    delete(v3, "\n\t")
                                case "a":
                                    var v38 uint = v1
                                    _ = v38
                                default:
                                    var v39 int = 146
                                    _ = v39
                                }
                                if true {
                                    var v40 uint16 = 32767
                                    _ = v40
                                } else {
                                    fmt.Println(v21, v28)
                                    v41 := GenS3{F0: 2, F1: "foo bar", F2: -9223372036854775808}
                                    _ = v41
                                }
                                v30 += -1
                            }
                        default:
                            if (!false) {
                                v5 += -897
                                if true {
                                    v42 := GenS4{F0: 1000, F1: 2}
                                    v43 := GenS3{F0: 42, F1: (("\n\t" + ("a" + ("a" + ""))) + ("a" + ("foo bar" + ""))), F2: 0}
                                    v44 := GenS4{F0: genFn0(), F1: v28}
                                    _ = v42
                                    _ = v43
                                    _ = v44
                                }
                            }
                        }
                    }
                    if (v19 == v19) {
                        switch ("\n\t" + ("foo bar" + "")) {
                        case "a":
                            var v45 uint = (v1 % ((v1) | 1))
                            _ = v45
                        case "":
                            if (true || (true || true)) {
                                switch int(v19) {
                                case 2147483647:
                                    var v46 int = int(v5)
                                    _ = v46
                                }
                                var v47 float32 = 1e10
                                v48 := GenS5{F0: (v1 - (v1 ^ (v1 | (v1 + (^(v1))))))}
                                var v49 int16 = v5
                                v30 -= v30
                                _ = v47
                                _ = v48
                                _ = v49
                            } else {
                                var v50 uint32 = 2
                                if len(v6) > 0 {
                                    v6[0] = "hello"
                                }
                                _ = v50
                            }
                        case "\n\t":
                            v51 := GenS1{F0: 449, F1: (v1 / ((v1) | 1)), F2: 0}
                            _ = v51
                        default:
                            fmt.Println(v5, v19, v5)
                        }
                        if true {
                            delete(v2, (-(fi31)))
                        }
                    }
                    _ = len(v2)
                    v19 = 0
                    _ = fi31
                }
                _ = v30
            }
            _ = v28
        }
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
