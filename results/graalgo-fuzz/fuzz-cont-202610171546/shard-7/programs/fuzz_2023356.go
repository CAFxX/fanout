package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int8; F1 uint64 }

func main() {
    var v1 uint16 = 0
    var v2 uint8 = 127
    v3 := []string{("foo bar" + "")}
    var v4 float64 = -2.25
    v5 := GenS1{F0: 501, F1: int16(v1)}
    v6 := GenS1{F0: 1, F1: -536}
    if (false || true) {
        if false {
            var v7 int32 = int32(v1)
            fmt.Println(v4, v2, v1)
            if true {
                v3 = append(v3, "\n\t")
            }
            _ = v7
        } else {
            for fi8 := 0; fi8 < 4; fi8++ {
                for fi9 := 0; fi9 < 17; fi9++ {
                    v3 = append(v3, "a")
                    _ = fi9
                }
                for fi10 := 0; fi10 < 20; fi10++ {
                    var v11 int8 = 42
                    if false {
                        v2 = (v2 >> ((65535) & 63))
                        if len(v3) > 0 {
                            v3[0] = ("" + ((("" + "a") + "foo bar") + ("hello" + "")))
                        }
                    }
                    _ = fi10
                    _ = v11
                }
                if len(v3) > 0 {
                    v3[0] = "a"
                }
                _ = fi8
            }
            v12 := GenS1{F0: -1, F1: -1}
            _ = v12
        }
        v2 = v2
        v1 -= (v1 | v1)
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v13 uint64 = 1
    v3 = append(v3, "a")
    for fi14 := 0; fi14 < 9; fi14++ {
        for fi15 := 0; fi15 < 20; fi15++ {
            v3 = append(v3, "")
            v16 := GenS1{F0: 939, F1: 32767}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi15
            _ = v16
        }
        v3 = append(v3, "foo bar")
        switch "hello" {
        case "\n\t":
            for fi17 := 0; fi17 < 13; fi17++ {
                v2 = v2
                fmt.Println(v1, v2, v13)
                if (!(-874 < 0)) {
                    var v18 uint16 = v1
                    _ = v18
                } else {
                    v19 := GenS1{F0: 1, F1: 847}
                    _ = v19
                }
                v1 = 0
                _ = fi17
            }
        default:
            switch ("hello" + "a") {
            case "foo bar":
                fmt.Println(fi14)
            case "":
                v2 = v2
            case "a":
                v1 = (v1 / ((2) | 1))
            default:
                if (-2147483648 >= int32(fi14)) {
                    for fi20 := 0; fi20 < 6; fi20++ {
                        var v21 int16 = -1
                        v1 = (v1 << ((497) & 63))
                        _ = fi20
                        _ = v21
                    }
                    if ((969 > 42) || ((0.0 <= (3.14159 * 0.1)) || false)) {
                        var v22 uint = 0
                        v13 = 4294967295
                        _ = v22
                    } else {
                        fmt.Println(v13, v2)
                        for fi23 := 0; fi23 < 20; fi23++ {
                            v3 = append(v3, "hello")
                            v4 += float64(-2.25)
                            if true {
                                if true {
                                    if true {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        var v24 int32 = 1
                                        var v25 uint32 = 632
                                        _ = v24
                                        _ = v25
                                    } else {
                                        var v26 int64 = -63
                                        if true {
                                            v1 -= (^(v1))
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v1 = (v1 % ((1) | 1))
                                        } else {
                                            var v27 uint32 = 1
                                            _ = v27
                                        }
                                        var v28 uint16 = v1
                                        _ = v26
                                        _ = v28
                                    }
                                    var v29 int16 = -1
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v29
                                }
                                v4 -= (v4 * (v4 - -2.25))
                                v2 -= v2
                            } else {
                                var v30 float32 = (-2.25 - ((0.0 * 3.14159) - float32(((v4 + float64(-2.25)) / 1e10))))
                                if len(v3) > 0 {
                                    v3[0] = ""
                                }
                                v30 -= v30
                                _ = v30
                            }
                            _ = fi23
                        }
                    }
                } else {
                    if ((fi14 & fi14) >= fi14) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    } else {
                        v31 := GenS1{F0: 1, F1: 1}
                        v32 := GenS1{F0: -1, F1: 612}
                        var v33 string = ("\n\t" + "hello")
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    v3 = append(v3, ("a" + ""))
                    _ = len(v3)
                }
            }
        }
        _ = fi14
    }
    v4 += 0.0
    v13 -= (^(v13))
    v13 = uint64(v1)
    v34 := GenS2{F0: 100, F1: (v13 % (((v13 % ((533) | 1))) | 1))}
    for fi35 := 0; fi35 < 19; fi35++ {
        v36 := GenS1{F0: int32(v13), F1: 1}
        var v37 string = "foo bar"
        _ = fi35
        _ = v36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
