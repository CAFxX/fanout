package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 int8 }
func genFn0(p0 int, p1 int, p2 int32) int8 {
    v4 := []int{(p1 + (-(p0))), 42, (p1 << ((65535) & 63)), p0}
    v5 := GenS1{F0: p0, F1: 1}
    if ((-(p0)) <= 0) {
        var v6 bool = (p2 < 42)
        var v7 int8 = -24
        for fi8 := 0; fi8 < 14; fi8++ {
            fmt.Println(v7)
            var v9 bool = (!false)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi8
            _ = v9
        }
        _ = v6
        _ = v7
    }
    switch p1 {
    case 0:
        fmt.Println(p2)
    case 42:
        if (!true) {
            var v10 string = ("\n\t" + ("a" + "foo bar"))
            _ = v10
        }
    case -2147483648:
        fmt.Println(p0)
    default:
        if true {
            fmt.Println(p0, p2)
            var v11 float64 = (0.1 / 1.5)
            _ = v11
        } else {
            var v12 int16 = 806
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 100
}

func main() {
    var v1 int64 = 0
    v2 := make(map[string]uint32)
    v3 := []uint32{}
    var v13 int8 = -1
    v14 := GenS2{F0: 814, F1: 178}
    if (-355 >= 42) {
        fmt.Println(v13)
        var v15 uint16 = 42
        _ = v15
    }
    var v16 string = ""
    v17 := GenS3{F0: 317}
    v18 := GenS3{F0: -1}
    fmt.Println(v1)
    v1 = v1
    for fi19 := 0; fi19 < 4; fi19++ {
        if ((!false) && false) {
            for fi20 := 0; fi20 < 18; fi20++ {
                v3 = append(v3, 0)
                if (0 >= 65535) {
                    v13 += (v13 | -38)
                    v2["foo bar"] = 2147483647
                    switch "hello" {
                    case "foo bar":
                        var v21 uint = 554
                        _ = v21
                    case "":
                        delete(v2, "")
                    case "hello":
                        v1 -= (v1 / (((v1 % ((-223) | 1))) | 1))
                    default:
                        v16 = ""
                    }
                    v16 = v16
                    v1 += v1
                    fmt.Println(v1)
                }
                v16 = ((v16 + "foo bar") + (v16 + v16))
                _ = fi20
            }
        } else {
            fmt.Println(v13, v1)
        }
        if (0.0 < 0.1) {
            v22 := GenS4{F0: 2147483647}
            v23 := GenS5{F0: 127, F1: v13}
            _ = v22
            _ = v23
        } else {
            for fi24 := 0; fi24 < 9; fi24++ {
                var v25 bool = (!true)
                v25 = (!v25)
                _ = fi24
                _ = v25
            }
            var v26 int32 = 1
            v26 += 1
            _ = v26
        }
        _ = len(v2)
        _ = fi19
    }
    var v27 int64 = -1
    _ = len(v3)
    switch 127 {
    case -817:
        if true {
            for fi28 := 0; fi28 < 7; fi28++ {
                v13 = (v13 ^ 42)
                if (!(false && (("foo bar" < (v16 + "a")) && false))) {
                    v1 += -593
                } else {
                    switch ((v16 + "a") + v16) {
                    case "hello":
                        if (!false) {
                            v29 := GenS4{F0: 2}
                            var v30 int = int(v13)
                            _ = v29
                            _ = v30
                        }
                    case "a":
                        v27 += 42
                    case "\n\t":
                        if true {
                            if (!false) {
                                v31 := GenS5{F0: 32768, F1: int8(v1)}
                                switch v16 {
                                case "hello":
                                    switch 32767 {
                                    case 127:
                                        for fi32 := 0; fi32 < 2; fi32++ {
                                            fmt.Println(v27, v13)
                                            _ = fi32
                                        }
                                    default:
                                        var v33 uint64 = 63
                                        _ = v33
                                    }
                                case "":
                                    switch fi28 {
                                    case 1:
                                        for fi34 := 0; fi34 < 14; fi34++ {
                                            delete(v2, ((v16 + v16) + "hello"))
                                            _ = fi34
                                        }
                                    case 9223372036854775807:
                                        switch fi28 {
                                        case 2147483647:
                                            v13 = (v13 ^ -1)
                                        default:
                                            v13 += (v13 | 83)
                                        }
                                    }
                                }
                                _ = v31
                            }
                        } else {
                            var v35 uint8 = 1
                            v27 -= (v27 * int64(fi28))
                            v35 = 127
                            _ = v35
                        }
                    }
                    v36 := GenS4{F0: 42}
                    _ = v36
                }
                switch v16 {
                case "a":
                    if (42 >= 2) {
                        v27 -= v27
                    }
                case "":
                    var v37 uint64 = 4294967295
                    _ = v37
                default:
                    for fi38 := 0; fi38 < 14; fi38++ {
                        v2[v16] = 2147483648
                        var v39 int32 = 1
                        if false {
                            switch (-(fi38)) {
                            case -932:
                                v2["hello"] = 661
                            default:
                                var v40 int32 = 374
                                _ = v40
                            }
                        } else {
                            v41 := GenS4{F0: 2}
                            var v42 int16 = -178
                            var v43 int = 4294967295
                            if false {
                                v44 := GenS5{F0: (fi28 | fi38), F1: v13}
                                _ = v44
                            }
                            v13 -= genFn0((v43 << ((255) & 63)), 127, (-(v39)))
                            _ = v41
                            _ = v42
                            _ = v43
                        }
                        _ = fi38
                        _ = v39
                    }
                }
                _ = fi28
            }
        }
    }
    v27 = 1
    for fi45 := 0; fi45 < 20; fi45++ {
        _ = len(v2)
        _ = fi45
    }
    var v46 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v46
}
