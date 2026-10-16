package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint8 }

func main() {
    v1 := []string{((("" + "") + "hello") + "hello"), "", ""}
    var v2 uint32 = 65535
    v3 := make(map[uint32]bool)
    v4 := make(map[int]string)
    var v5 bool = (1 != -1)
    var v6 int64 = 0
    if (65535 <= uint16(v2)) {
        var v7 int32 = -1
        var v8 bool = false
        _ = v7
        _ = v8
    } else {
        v5 = (!(!v5))
        var v9 bool = (v5 && (0 == 42))
        _ = v9
    }
    v10 := GenS1{F0: ("a" + ""), F1: 583}
    var v11 uint = 18446744073709551615
    delete(v4, 2147483647)
    v12 := GenS1{F0: "a", F1: 18446744073709551615}
    fmt.Println(v11, v5, v2)
    fmt.Println(v2)
    v11 = 42
    if v5 {
        v11 = v11
        if (1.5 > 0.1) {
            switch 679 {
            case -427:
                v13 := GenS1{F0: ((("a" + "foo bar") + "") + "foo bar"), F1: 2}
                _ = v13
            }
        } else {
            for fi14 := 0; fi14 < 2; fi14++ {
                if v5 {
                    if (882 != 4294967295) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v15 bool = (((v5 && false) && v5) && (v5 || (32767 >= 0)))
                        _ = v15
                    } else {
                        _ = len(v3)
                    }
                } else {
                    v16 := GenS2{F0: 0}
                    _ = v16
                }
                v17 := GenS2{F0: 42}
                _ = fi14
                _ = v17
            }
            fmt.Println(v6, v6)
            _ = len(v1)
        }
        switch "foo bar" {
        case "\n\t":
            for fi18 := 0; fi18 < 10; fi18++ {
                switch "\n\t" {
                case "a":
                    var v19 uint = (v11 - (v11 - 65535))
                    _ = v19
                case "\n\t":
                    switch "foo bar" {
                    case "a":
                        if v5 {
                            for fi20 := 0; fi20 < 19; fi20++ {
                                v21 := GenS1{F0: ("a" + "foo bar"), F1: 42}
                                _ = fi20
                                _ = v21
                            }
                        } else {
                            for fi22 := 0; fi22 < 19; fi22++ {
                                var v23 uint8 = 209
                                _ = fi22
                                _ = v23
                            }
                        }
                    default:
                        v24 := v4[fi18]
                        _ = v24
                    }
                default:
                    v25 := v4[391]
                    _ = v25
                }
                _ = fi18
            }
        case "a":
            if (0.0 == (float32((1.5 / 3.14159)) - 0.0)) {
                if ("" != "") {
                    var v26 float64 = (-2.25 * 0.0)
                    _ = v26
                }
                for fi27 := 0; fi27 < 2; fi27++ {
                    var v28 int = fi27
                    _ = fi27
                    _ = v28
                }
            } else {
                var v29 uint64 = 1
                for fi30 := 0; fi30 < 5; fi30++ {
                    if (float32(0.0) < 3.14159) {
                        if false {
                            v31 := v3[4294967295]
                            _ = v31
                        } else {
                            var v32 int32 = 2147483647
                            var v33 string = "foo bar"
                            _ = v32
                            _ = v33
                        }
                        var v34 int = 32767
                        var v35 int8 = 1
                        _ = v34
                        _ = v35
                    } else {
                        v29 += (^(v29))
                        v11 = (v11 % ((18446744073709551615) | 1))
                        v36 := GenS3{F0: 42, F1: int8(v6)}
                        _ = v36
                    }
                    _ = fi30
                }
                _ = v29
            }
        default:
            switch "\n\t" {
            case "\n\t":
                v5 = ("hello" != "a")
            default:
                for fi37 := 0; fi37 < 16; fi37++ {
                    var v38 int32 = 439
                    _ = fi37
                    _ = v38
                }
            }
        }
    } else {
        var v39 bool = true
        var v40 uint8 = 2
        for fi41 := 0; fi41 < 7; fi41++ {
            v5 = ("hello" != ("foo bar" + "\n\t"))
            v42 := GenS4{F0: 0, F1: 2}
            v43 := GenS3{F0: 42, F1: -1}
            _ = fi41
            _ = v42
            _ = v43
        }
        _ = v39
        _ = v40
    }
    v44 := GenS5{F0: (v11 & v11)}
    if ((!false) || (32767 > 32768)) {
        v45 := GenS6{F0: 1}
        v5 = (false && true)
        fmt.Println(v2, v11, v5)
        _ = v45
    }
    var v46 bool = v5
    for fi47 := 0; fi47 < 11; fi47++ {
        var v48 float32 = -2.25
        v48 = 1e10
        _ = fi47
        _ = v48
    }
    v5 = (v5 && v46)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v46
    _ = v5
    _ = v6
}
