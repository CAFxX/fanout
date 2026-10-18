package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 string }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16 }

func main() {
    v1 := []uint32{65535, 0, 42}
    var v2 uint64 = 401
    var v3 string = "hello"
    var v4 int = -59
    var v5 int64 = -905
    if true {
        for fi6 := 0; fi6 < 3; fi6++ {
            v7 := GenS1{F0: v4, F1: ((v3 + ((v3 + "") + v3)) > (v3 + v3)), F2: v3}
            _ = fi6
            _ = v7
        }
        switch v4 {
        case 9223372036854775807:
            v8 := GenS1{F0: int(v5), F1: (v2 < v2), F2: v3}
            _ = v8
        case 0:
            fmt.Println(v5)
        default:
            v1 = append(v1, 1)
        }
        var v9 uint16 = 693
        _ = v9
    }
    if (!(!false)) {
        switch "hello" {
        case "foo bar":
            if (!false) {
                switch 0 {
                case -1:
                    v10 := GenS2{F0: (v5 ^ 147), F1: v5}
                    _ = v10
                case 4294967295:
                    for fi11 := 0; fi11 < 4; fi11++ {
                        var v12 uint = 65535
                        _ = fi11
                        _ = v12
                    }
                }
                for fi13 := 0; fi13 < 8; fi13++ {
                    for fi14 := 0; fi14 < 11; fi14++ {
                        if len(v1) > 0 {
                            v1[0] = 1
                        }
                        var v15 uint8 = 128
                        _ = fi14
                        _ = v15
                    }
                    v3 = (v3 + "hello")
                    switch v3 {
                    case "hello":
                        if (255 <= uint8(v2)) {
                            _ = len(v1)
                            v16 := GenS2{F0: v5, F1: (-(v5))}
                            if (1 < 1) {
                                var v17 uint16 = 255
                                for fi18 := 0; fi18 < 8; fi18++ {
                                    v3 = ("hello" + "a")
                                    _ = fi18
                                }
                                switch "" {
                                case "":
                                    if false {
                                        var v19 bool = (!true)
                                        if v19 {
                                            var v20 int64 = v5
                                            _ = v20
                                        } else {
                                            _ = len(v1)
                                            if (!v19) {
                                                v2 = v2
                                                switch v3 {
                                                case "":
                                                    v4 += fi13
                                                case "a":
                                                    v3 = "foo bar"
                                                }
                                            }
                                            var v21 float64 = -2.25
                                            _ = v21
                                        }
                                        _ = v19
                                    }
                                case "hello":
                                    v3 = (("a" + "") + "foo bar")
                                }
                                _ = v17
                            } else {
                                if (42 != 693) {
                                    if false {
                                        var v22 bool = true
                                        _ = v22
                                    }
                                    fmt.Println(fi13, v2)
                                    var v23 uint16 = 879
                                    _ = v23
                                }
                            }
                            _ = v16
                        }
                    case "foo bar":
                        _ = len(v1)
                    }
                    _ = fi13
                }
                if true {
                    switch v3 {
                    case "a":
                        _ = len(v1)
                    default:
                        for fi24 := 0; fi24 < 14; fi24++ {
                            for fi25 := 0; fi25 < 1; fi25++ {
                                switch v3 {
                                case "":
                                    v26 := GenS1{F0: (fi25 % (((^(v4))) | 1)), F1: (!true), F2: "foo bar"}
                                    _ = v26
                                case "a":
                                    if len(v1) > 0 {
                                        v1[0] = 454
                                    }
                                case "hello":
                                    var v27 uint8 = 245
                                    if (fi25 <= 42) {
                                        if false {
                                            var v28 int = (fi24 & fi25)
                                            v27 = (v27 * (v27 ^ v27))
                                            _ = v28
                                        } else {
                                            v29 := GenS3{F0: v2}
                                            _ = v29
                                        }
                                    }
                                    _ = v27
                                default:
                                    v2 -= (v2 & v2)
                                }
                                _ = fi25
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v30 float64 = -2.25
                            _ = fi24
                            _ = v30
                        }
                    }
                } else {
                    v5 = (v5 % (((v5 - v5)) | 1))
                }
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v4 = -574
                if (float64(3.14159) >= 1e10) {
                    if (true || (true && (false && false))) {
                        v4 = (v4 | v4)
                    } else {
                        fmt.Println(v5)
                    }
                }
            }
        default:
            if false {
                for fi31 := 0; fi31 < 20; fi31++ {
                    var v32 uint64 = v2
                    v2 = (v2 / ((v32) | 1))
                    _ = fi31
                    _ = v32
                }
                v33 := GenS4{F0: 32767}
                _ = v33
            } else {
                v4 = v4
                var v34 string = v3
                for fi35 := 0; fi35 < 5; fi35++ {
                    v2 += v2
                    v1 = append(v1, 4294967295)
                    v2 -= v2
                    _ = fi35
                }
                _ = v34
            }
        }
    }
    _ = len(v1)
    v36 := GenS2{F0: (v5 + 878), F1: (v5 << ((2) & 63))}
    var v37 uint = 0
    var v38 int64 = v5
    for fi39 := 0; fi39 < 11; fi39++ {
        if (42 != -577) {
            var v40 uint8 = 128
            v41 := GenS4{F0: 658}
            _ = v40
            _ = v41
        } else {
            if true {
                v2 = v2
                v38 -= (v5 & (v38 << (((v37 >> (((v37 + (v37 / (((^(v37))) | 1)))) & 63))) & 63)))
                _ = len(v1)
            }
            if len(v1) > 0 {
                v1[0] = 780
            }
        }
        v2 = 0
        var v42 int64 = 0
        _ = fi39
        _ = v42
    }
    for fi43 := 0; fi43 < 3; fi43++ {
        var v44 uint32 = 473
        var v45 bool = (!(!true))
        v46 := GenS4{F0: 65535}
        _ = fi43
        _ = v44
        _ = v45
        _ = v46
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
