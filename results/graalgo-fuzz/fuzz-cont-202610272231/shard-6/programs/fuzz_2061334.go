package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 string; F1 int64; F2 string }
type GenS3 struct { F0 uint32; F1 uint; F2 int }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS6 struct { F0 string; F1 int16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 539
    v3 := []uint32{42, 149}
    var v4 int16 = 32767
    var v5 bool = false
    var v6 string = ("" + "foo bar")
    v5 = (991 >= 2)
    if true {
        v3 = append(v3, 1)
        var v7 int = v2
        if v5 {
            v3 = append(v3, 859)
            if v5 {
                v8 := GenS1{F0: v4, F1: 65535}
                _ = v8
            }
        } else {
            if (((127 != 133) && true) && true) {
                v9 := v1[937]
                for fi10 := 0; fi10 < 7; fi10++ {
                    switch (v7 & (v2 ^ v2)) {
                    case 32767:
                        switch v7 {
                        case 442:
                            fmt.Println(v7, v2, fi10)
                        }
                    case -465:
                        var v11 string = v6
                        _ = v11
                    case 0:
                        for fi12 := 0; fi12 < 2; fi12++ {
                            v2 = 127
                            v13 := GenS1{F0: v4, F1: 0}
                            _ = fi12
                            _ = v13
                        }
                    }
                    v14 := GenS1{F0: (v4 * v4), F1: 2}
                    _ = fi10
                    _ = v14
                }
                _ = v9
            } else {
                switch (("" + "") + v6) {
                case "a":
                    for fi15 := 0; fi15 < 8; fi15++ {
                        _ = len(v1)
                        for fi16 := 0; fi16 < 12; fi16++ {
                            fmt.Println(v7, fi15)
                            _ = fi16
                        }
                        if len(v3) > 0 {
                            v3[0] = 173
                        }
                        _ = fi15
                    }
                case "":
                    if true {
                        if v5 {
                            var v17 float32 = (0.0 / float32(-2.25))
                            v3 = append(v3, 2147483648)
                            if ((!true) || (true || (float32(-2.25) <= 0.0))) {
                                switch (v6 + v6) {
                                case "\n\t":
                                    for fi18 := 0; fi18 < 17; fi18++ {
                                        fmt.Println(fi18)
                                        v17 += v17
                                        _ = fi18
                                    }
                                case "foo bar":
                                    for fi19 := 0; fi19 < 20; fi19++ {
                                        var v20 int32 = 1
                                        _ = fi19
                                        _ = v20
                                    }
                                case "":
                                    fmt.Println(v5, v7)
                                default:
                                    v21 := GenS2{F0: "hello", F1: 1, F2: v6}
                                    _ = v21
                                }
                            }
                            _ = v17
                        }
                        var v22 float32 = 0.1
                        switch v6 {
                        case "":
                            if (4294967295 > 154) {
                                for fi23 := 0; fi23 < 17; fi23++ {
                                    if false {
                                        if (65535 >= 1) {
                                            v7 = (-(fi23))
                                        } else {
                                            fmt.Println(v7, v22, v6)
                                            v24 := GenS3{F0: 4294967295, F1: 255, F2: (-(v7))}
                                            _ = v24
                                        }
                                        v4 += (v4 % (((v4 / ((v4) | 1))) | 1))
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = 4294967295
                                    }
                                    _ = fi23
                                }
                                if v5 {
                                    switch ((v6 + "foo bar") + "a") {
                                    case "":
                                        var v25 float64 = ((-2.25 + (0.1 + 1e10)) * 0.1)
                                        _ = v25
                                    case "foo bar":
                                        var v26 int16 = v4
                                        _ = v26
                                    case "\n\t":
                                        v2 -= -491
                                    }
                                }
                            }
                        }
                        _ = v22
                    }
                case "foo bar":
                    var v27 int = (v7 | -2147483648)
                    _ = v27
                }
                fmt.Println(v4)
            }
        }
        _ = v7
    } else {
        for fi28 := 0; fi28 < 18; fi28++ {
            if (255 <= 210) {
                switch 0 {
                case 127:
                    if (!true) {
                        fmt.Println(v5, v4)
                    }
                case -2147483648:
                    v3 = append(v3, 2)
                }
                v2 -= 167
            } else {
                for fi29 := 0; fi29 < 13; fi29++ {
                    var v30 uint = 930
                    _ = fi29
                    _ = v30
                }
                v31 := GenS4{F0: 42, F1: 420}
                v6 = v6
                _ = v31
            }
            fmt.Println(v4, v5, v2)
            var v32 string = v6
            _ = fi28
            _ = v32
        }
        v4 -= (v4 % (((v4 % (((v4 ^ (v4 * (-(v4))))) | 1))) | 1))
    }
    for fi33 := 0; fi33 < 4; fi33++ {
        fmt.Println(v5, v2)
        delete(v1, 2147483648)
        fmt.Println(v4, fi33, v2)
        _ = fi33
    }
    if false {
        for fi34 := 0; fi34 < 10; fi34++ {
            switch 2147483647 {
            case 457:
                var v35 float32 = 3.14159
                _ = v35
            default:
                v4 = (^(v4))
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi34
        }
        v5 = (v5 || v5)
    }
    fmt.Println(v4, v5, v4)
    if (v4 == 32767) {
        v5 = (0 >= -240)
        fmt.Println(v2, v6)
    } else {
        if v5 {
            v36 := GenS3{F0: 65535, F1: 0, F2: (v2 & v2)}
            _ = v36
        }
        v37 := GenS2{F0: (v6 + ("foo bar" + "")), F1: 1, F2: ""}
        _ = v37
    }
    for fi38 := 0; fi38 < 18; fi38++ {
        var v39 bool = true
        if len(v3) > 0 {
            v3[0] = 979
        }
        for fi40 := 0; fi40 < 12; fi40++ {
            fmt.Println(v39, v2, fi38)
            v39 = (!(32767 == 0))
            if true {
                var v41 int = (v2 - 14)
                _ = v41
            } else {
                var v42 bool = (127 > -75)
                _ = len(v3)
                _ = v42
            }
            var v43 float64 = 1e10
            _ = fi40
            _ = v43
        }
        _ = fi38
        _ = v39
    }
    var v44 bool = true
    v45 := GenS2{F0: "a", F1: -963, F2: v6}
    v46 := GenS4{F0: 2147483648, F1: 697}
    v47 := GenS5{F0: 2147483647, F1: 42, F2: 127}
    v48 := GenS2{F0: (v6 + v6), F1: -857, F2: v6}
    for fi49 := 0; fi49 < 3; fi49++ {
        fmt.Println(v44, fi49)
        _ = fi49
    }
    v6 = ("\n\t" + "foo bar")
    if true {
        for fi50 := 0; fi50 < 7; fi50++ {
            v2 = fi50
            _ = len(v1)
            _ = fi50
        }
        v44 = v44
    }
    if (331 <= -1) {
        v51 := GenS5{F0: 0, F1: 1, F2: 42}
        v44 = false
        _ = v51
    } else {
        switch "hello" {
        case "foo bar":
            v4 = v4
        case "":
            delete(v1, 475)
        case "a":
            v52 := GenS1{F0: 42, F1: 1}
            _ = v52
        }
        _ = len(v3)
    }
    for fi53 := 0; fi53 < 19; fi53++ {
        switch -585 {
        case -548:
            v54 := GenS2{F0: "\n\t", F1: -9223372036854775808, F2: v6}
            _ = v54
        case 4294967295:
            if v44 {
                var v55 int8 = 4
                fmt.Println(v2)
                if ((0.1 + (0.1 * -2.25)) < float64(0.1)) {
                    v2 = -150
                }
                _ = v55
            } else {
                var v56 int64 = 0
                _ = v56
            }
        case -725:
            var v57 uint32 = 4294967295
            _ = v57
        }
        v58 := GenS6{F0: ("hello" + v6), F1: v4}
        _ = len(v3)
        _ = fi53
        _ = v58
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
