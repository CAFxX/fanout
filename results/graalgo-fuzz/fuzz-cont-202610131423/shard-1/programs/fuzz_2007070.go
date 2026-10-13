package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 uint32; F2 string }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64; F1 uint64; F2 int64 }
func genFn0(p0 uint16, p1 int16) string {
    for fi4 := 0; fi4 < 14; fi4++ {
        var v5 uint16 = p0
        p1 = (p1 ^ p1)
        fmt.Println(p1, fi4)
        _ = fi4
        _ = v5
    }
    p0 = p0
    v6 := GenS1{F0: -1}
    fmt.Println(p0)
    var v7 int = -301
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return ("a" + "foo bar")
}

func main() {
    var v1 uint = 42
    var v2 int16 = 192
    v3 := []string{}
    fmt.Println(v1)
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 int8 = -115
        var v10 uint = (v1 | (^(v1)))
        _ = fi8
        _ = v10
        _ = v9
    }
    if true {
        if ((65535 <= 0) && (!((float32(3.14159) * (((0.1 + ((-2.25 + 1.5) - -2.25)) / 1.5) + 0.0)) <= 3.14159))) {
            if true {
                for fi11 := 0; fi11 < 17; fi11++ {
                    v2 = (v2 / ((v2) | 1))
                    v12 := GenS1{F0: -2147483648}
                    v1 = 824
                    _ = fi11
                    _ = v12
                }
            } else {
                for fi13 := 0; fi13 < 7; fi13++ {
                    switch (fi13 >> ((v1) & 63)) {
                    case -793:
                        if false {
                            _ = len(v3)
                        } else {
                            var v14 uint32 = 518
                            _ = v14
                        }
                    case 42:
                        switch "foo bar" {
                        case "hello":
                            switch "foo bar" {
                            case "a":
                                var v15 uint = v1
                                _ = v15
                            default:
                                v16 := GenS2{F0: (v1 | 255), F1: 433, F2: "foo bar"}
                                _ = v16
                            }
                        case "foo bar":
                            var v17 int8 = 42
                            _ = v17
                        case "":
                            v1 = v1
                        }
                    default:
                        if (!true) {
                            if (true && ((v1 % ((v1) | 1)) >= (v1 % ((v1) | 1)))) {
                                if ((!(880 < -1)) && false) {
                                    for fi18 := 0; fi18 < 1; fi18++ {
                                        v2 = (v2 ^ v2)
                                        v19 := GenS3{F0: (v2 + -507), F1: 42}
                                        var v20 int64 = 0
                                        _ = fi18
                                        _ = v19
                                        _ = v20
                                    }
                                } else {
                                    var v21 int16 = (v2 + v2)
                                    _ = len(v3)
                                    _ = v21
                                }
                                v22 := GenS3{F0: (v2 - int16(v1)), F1: 440}
                                _ = v22
                            }
                            v2 -= (v2 & (v2 - v2))
                        } else {
                            v2 += (v2 << ((v1) & 63))
                            v2 += v2
                        }
                    }
                    _ = len(v3)
                    _ = fi13
                }
            }
            var v23 float64 = 1e10
            v24 := GenS4{F0: 0}
            _ = v23
            _ = v24
        } else {
            _ = len(v3)
            v25 := GenS2{F0: 170, F1: 0, F2: ((("foo bar" + "a") + "") + "")}
            _ = v25
        }
        var v26 float32 = 3.14159
        _ = v26
    }
    if len(v3) > 0 {
        v3[0] = "a"
    }
    v27 := GenS5{F0: 42, F1: 4294967295, F2: 42}
    v2 += v2
    for fi28 := 0; fi28 < 4; fi28++ {
        fmt.Println(v1)
        var v29 string = "a"
        _ = fi28
        _ = v29
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
}
