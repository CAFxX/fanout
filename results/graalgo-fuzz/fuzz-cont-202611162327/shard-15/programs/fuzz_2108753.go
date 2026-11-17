package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int64 }
type GenS2 struct { F0 uint; F1 int8 }

func main() {
    var v1 float32 = 1.5
    v2 := make(map[string]string)
    var v3 uint = 2
    var v4 uint8 = 127
    v4 = v4
    var v5 int = -9223372036854775808
    if ("" < (("hello" + ("hello" + ((("a" + (("a" + "a") + "")) + ("hello" + ("a" + ""))) + "hello"))) + "foo bar")) {
        if true {
            fmt.Println(v1)
        } else {
            v3 += v3
            v6 := GenS1{F0: 32767, F1: -1, F2: 636}
            _ = v6
        }
        fmt.Println(v1, v1, v1)
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        if (4294967295 != 4294967295) {
            fmt.Println(v5, v1, v4)
            v5 = 6
            v8 := GenS1{F0: -1, F1: -1, F2: -9223372036854775808}
            _ = v8
        } else {
            var v9 int16 = -692
            var v10 uint64 = 598
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    if (v5 >= -2147483648) {
        v11 := GenS1{F0: 1000, F1: 32767, F2: 9223372036854775807}
        var v12 int16 = 179
        _ = v11
        _ = v12
    }
    fmt.Println(v1, v5)
    var v13 uint16 = 0
    if (0.0 > (v1 / v1)) {
        for fi14 := 0; fi14 < 2; fi14++ {
            for fi15 := 0; fi15 < 9; fi15++ {
                v16 := GenS1{F0: 0, F1: 32767, F2: -1}
                _ = fi15
                _ = v16
            }
            v17 := GenS1{F0: 42, F1: -1, F2: 1}
            _ = fi14
            _ = v17
        }
        for fi18 := 0; fi18 < 18; fi18++ {
            switch fi18 {
            case 0:
                switch fi18 {
                case -1:
                    switch "" {
                    case "a":
                        var v19 int = fi18
                        _ = v19
                    case "\n\t":
                        v20 := GenS2{F0: v3, F1: 0}
                        _ = v20
                    default:
                        if (!((992 < 0) || false)) {
                            switch "\n\t" {
                            case "":
                                v4 -= v4
                            case "a":
                                var v21 int8 = 42
                                _ = v21
                            case "foo bar":
                                v4 -= (v4 >> ((uint(v4)) & 63))
                            default:
                                v4 = v4
                            }
                            switch "\n\t" {
                            case "":
                                switch "foo bar" {
                                case "foo bar":
                                    v22 := GenS2{F0: 111, F1: 0}
                                    _ = v22
                                case "":
                                    _ = len(v2)
                                default:
                                    switch "\n\t" {
                                    case "":
                                        if false {
                                            var v23 bool = (!true)
                                            switch ("" + ("a" + "a")) {
                                            case "\n\t":
                                                switch (fi18 & (fi18 - (fi18 << (((v3 | (v3 + (^(v3))))) & 63)))) {
                                                case 299:
                                                    v1 -= 0.1
                                                default:
                                                    var v24 uint64 = 42
                                                    _ = v24
                                                }
                                            default:
                                                v25 := []int{}
                                                _ = v25
                                            }
                                            _ = v23
                                        } else {
                                            v3 = v3
                                            v5 += 248
                                            var v26 int8 = -127
                                            _ = v26
                                        }
                                    case "a":
                                        v13 = 65535
                                    case "hello":
                                        var v27 int = (-(fi18))
                                        _ = v27
                                    default:
                                        var v28 int = v5
                                        _ = v28
                                    }
                                }
                            default:
                                for fi29 := 0; fi29 < 12; fi29++ {
                                    fmt.Println(v3, fi18)
                                    if ((v13 & v13) <= 0) {
                                        var v30 int8 = 1
                                        _ = v30
                                    } else {
                                        var v31 int = (fi29 + fi29)
                                        v2[""] = "a"
                                        _ = v31
                                    }
                                    _ = fi29
                                }
                            }
                        } else {
                            v2["a"] = "hello"
                            if false {
                                delete(v2, "")
                            } else {
                                fmt.Println(v3)
                            }
                            var v32 int64 = -9223372036854775808
                            _ = v32
                        }
                    }
                case -2147483648:
                    var v33 float32 = v1
                    _ = v33
                default:
                    v34 := GenS2{F0: (v3 | (v3 + 1)), F1: 4}
                    _ = v34
                }
            case -1:
                var v35 int8 = -98
                _ = v35
            case 2147483647:
                if (false || false) {
                    v1 += 0.1
                } else {
                    var v36 float64 = -2.25
                    _ = v36
                }
            default:
                delete(v2, "a")
            }
            switch -350 {
            case -849:
                v4 = (^(v4))
            case 1:
                for fi37 := 0; fi37 < 18; fi37++ {
                    var v38 int = (fi37 / (((fi18 % (((fi37 ^ (fi37 % (((fi18 & (v5 * -236))) | 1)))) | 1))) | 1))
                    var v39 int32 = -1
                    v4 = (v4 << ((v3) & 63))
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            }
            fmt.Println(v5, v5)
            _ = fi18
        }
    } else {
        var v40 int8 = 127
        _ = v40
    }
    fmt.Println(v4, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
