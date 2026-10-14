package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint16 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 uint64; F1 int8; F2 int32 }
type GenS5 struct { F0 int64; F1 bool; F2 int16 }
func genFn0(p0 uint16, p1 int16) int32 {
    fmt.Println(p1)
    if (!true) {
        var v6 uint32 = 0
        var v7 uint8 = 62
        v7 = 67
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    return -892
}

func main() {
    var v1 string = "\n\t"
    v2 := []string{}
    var v3 string = v1
    v4 := make(map[int]bool)
    var v5 uint8 = 2
    _ = len(v2)
    switch 42 {
    case -435:
        fmt.Println(v1, v1)
    case -9223372036854775808:
        for fi8 := 0; fi8 < 2; fi8++ {
            switch "hello" {
            case "a":
                var v9 float32 = 0.1
                _ = v9
            case "":
                if false {
                    var v10 string = ("" + "\n\t")
                    if true {
                        fmt.Println(v10, v10, fi8)
                        var v11 int32 = genFn0(uint16(v5), 32767)
                        switch v3 {
                        case "":
                            for fi12 := 0; fi12 < 10; fi12++ {
                                v13 := GenS1{F0: 4294967295, F1: false, F2: 801}
                                _ = fi12
                                _ = v13
                            }
                        case "\n\t":
                            if len(v2) > 0 {
                                v2[0] = (v1 + v1)
                            }
                        }
                        _ = v11
                    } else {
                        if (((fi8 >> ((255) & 63)) <= -580) && (v3 < (v10 + v1))) {
                            if ((true && true) && true) {
                                for fi14 := 0; fi14 < 15; fi14++ {
                                    fmt.Println(v5)
                                    _ = fi14
                                }
                                v15 := v4[-440]
                                _ = v15
                            }
                            for fi16 := 0; fi16 < 17; fi16++ {
                                v17 := GenS2{F0: 0, F1: ("" + ((v10 + v3) + "hello"))}
                                _ = fi16
                                _ = v17
                            }
                        } else {
                            v18 := GenS3{F0: -51, F1: 127}
                            _ = v18
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    var v19 uint32 = 4294967295
                    _ = v10
                    _ = v19
                } else {
                    var v20 string = ("a" + v1)
                    if (false || true) {
                        var v21 uint = 536
                        var v22 float32 = ((float32(((3.14159 / 1e10) - -2.25)) / 1.5) + ((1e10 * float32(0.1)) * 0.0))
                        var v23 float64 = 1.5
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    _ = v20
                }
            case "\n\t":
                if true {
                    for fi24 := 0; fi24 < 3; fi24++ {
                        switch ("foo bar" + v3) {
                        case "\n\t":
                            v1 = ("a" + "foo bar")
                        default:
                            var v25 uint32 = 42
                            _ = v25
                        }
                        var v26 int16 = 818
                        _ = fi24
                        _ = v26
                    }
                    var v27 uint = 42
                    _ = v27
                } else {
                    var v28 uint8 = v5
                    v29 := GenS4{F0: 2, F1: int8(fi8), F2: 985}
                    _ = v28
                    _ = v29
                }
            }
            _ = fi8
        }
    case 127:
        var v30 uint16 = 295
        _ = v30
    default:
        fmt.Println(v1)
    }
    v3 = ""
    v1 = v1
    var v31 uint64 = 4294967295
    var v32 int8 = 100
    _ = len(v4)
    switch -640 {
    case 0:
        var v33 int64 = -1
        _ = v33
    }
    _ = len(v4)
    v32 = 100
    var v34 int64 = -9223372036854775808
    _ = len(v2)
    var v35 uint32 = 0
    for fi36 := 0; fi36 < 16; fi36++ {
        if false {
            v37 := GenS2{F0: 533, F1: v1}
            if false {
                if len(v2) > 0 {
                    v2[0] = (v3 + v3)
                }
            } else {
                var v38 int32 = genFn0(65535, 0)
                v35 -= 360
                v3 = "\n\t"
                _ = v38
            }
            var v39 uint32 = (v35 / ((v35) | 1))
            _ = v37
            _ = v39
        } else {
            delete(v4, -960)
            v40 := GenS5{F0: (v34 % (((v34 - (v34 - v34))) | 1)), F1: false, F2: 0}
            _ = v40
        }
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
