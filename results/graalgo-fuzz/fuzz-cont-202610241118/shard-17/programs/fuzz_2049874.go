package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 bool }

func main() {
    var v1 int64 = 0
    v2 := []string{}
    v3 := []int{-9223372036854775808, -873, 1}
    var v4 uint8 = 127
    fmt.Println(v4)
    v1 -= (-(v1))
    var v5 string = "foo bar"
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 uint = 2
        _ = fi6
        _ = v7
    }
    if true {
        v4 = v4
        v4 = (^(v4))
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v4 += 2
        for fi8 := 0; fi8 < 11; fi8++ {
            for fi9 := 0; fi9 < 1; fi9++ {
                if (!true) {
                    var v10 int8 = 0
                    _ = len(v2)
                    v4 += 2
                    _ = v10
                }
                _ = fi9
            }
            var v11 uint32 = 42
            if (65535 != 42) {
                v11 = (v11 + v11)
            }
            _ = fi8
            _ = v11
        }
    }
    if (((1.5 * ((0.0 - 3.14159) + -2.25)) + 0.1) >= 0.0) {
        var v12 uint16 = 173
        if (2147483648 == 4294967295) {
            if true {
                var v13 uint32 = 0
                _ = v13
            } else {
                var v14 float64 = 0.1
                var v15 int16 = 234
                fmt.Println(v5, v1)
                _ = v14
                _ = v15
            }
            if ((uint16(v4) > v12) || (0 >= 1)) {
                if true {
                    v12 += v12
                }
            }
        } else {
            if (!true) {
                if false {
                    var v16 uint32 = 395
                    if false {
                        for fi17 := 0; fi17 < 11; fi17++ {
                            v4 = (v4 * (^(v4)))
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi17
                        }
                    } else {
                        if false {
                            if (!(((float64(0.0) * 3.14159) <= -2.25) && false)) {
                                v12 = (^(v12))
                                fmt.Println(v5, v5, v1)
                            }
                        } else {
                            for fi18 := 0; fi18 < 15; fi18++ {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = fi18
                            }
                            v1 += v1
                        }
                    }
                    _ = v16
                }
                v5 = ((v5 + "") + v5)
            }
        }
        _ = v12
    } else {
        switch v5 {
        case "":
            for fi19 := 0; fi19 < 10; fi19++ {
                if len(v2) > 0 {
                    v2[0] = ("hello" + ("" + v5))
                }
                _ = fi19
            }
        default:
            if false {
                for fi20 := 0; fi20 < 15; fi20++ {
                    if true {
                        if len(v3) > 0 {
                            v3[0] = (fi20 ^ 0)
                        }
                    }
                    _ = fi20
                }
            } else {
                switch 127 {
                case 2147483647:
                    var v21 uint16 = 446
                    _ = v21
                case 42:
                    fmt.Println(v5)
                }
                var v22 string = ""
                _ = v22
            }
        }
        var v23 bool = ((!(42 != v1)) || (!(!(-400 >= -1))))
        _ = v23
    }
    var v24 uint16 = 1
    if len(v3) > 0 {
        v3[0] = int(v24)
    }
    v25 := GenS1{F0: v5, F1: (v1 - v1)}
    var v26 int64 = v1
    if (-775 > 32767) {
        v1 += v1
        v1 = v26
        var v27 uint16 = (v24 | (v24 + v24))
        _ = v27
    }
    v26 -= v26
    v28 := GenS2{F0: true}
    for fi29 := 0; fi29 < 19; fi29++ {
        if len(v2) > 0 {
            v2[0] = (v5 + v5)
        }
        for fi30 := 0; fi30 < 4; fi30++ {
            var v31 uint32 = 42
            _ = fi30
            _ = v31
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
