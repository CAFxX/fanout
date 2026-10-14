package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 uint64 }

func main() {
    var v1 uint16 = 0
    var v2 int16 = -32768
    var v3 uint = 42
    v4 := GenS1{F0: "hello", F1: -1}
    var v5 uint8 = 156
    v6 := []int{127, 9223372036854775807, 733}
    switch 0 {
    case 2147483647:
        if false {
            v7 := GenS1{F0: "", F1: 42}
            _ = v7
        } else {
            v8 := GenS2{F0: 2, F1: v2}
            if true {
                switch ("hello" + "") {
                case "hello":
                    _ = len(v6)
                }
            } else {
                if true {
                    var v9 float32 = 0.1
                    v2 -= (v2 / (((^(v2))) | 1))
                    if ("foo bar" == "foo bar") {
                        if true {
                            v6 = append(v6, 32767)
                            var v10 float32 = v9
                            _ = v10
                        } else {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            if false {
                                if true {
                                    fmt.Println(v3, v3)
                                }
                                for fi11 := 0; fi11 < 15; fi11++ {
                                    if false {
                                        var v12 bool = true
                                        v6 = append(v6, -591)
                                        _ = v12
                                    }
                                    _ = fi11
                                }
                                var v13 float32 = ((v9 * -2.25) * -2.25)
                                _ = v13
                            }
                            v2 = v2
                        }
                        fmt.Println(v5, v9)
                    }
                    _ = v9
                } else {
                    for fi14 := 0; fi14 < 17; fi14++ {
                        var v15 uint16 = 2
                        switch "" {
                        case "":
                            v5 = 255
                        }
                        _ = fi14
                        _ = v15
                    }
                    var v16 int16 = (v2 << ((v3) & 63))
                    v17 := GenS2{F0: 0, F1: (v16 / ((v16) | 1))}
                    _ = v16
                    _ = v17
                }
                _ = len(v6)
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            for fi18 := 0; fi18 < 12; fi18++ {
                for fi19 := 0; fi19 < 12; fi19++ {
                    v20 := GenS1{F0: ("hello" + ""), F1: 1}
                    _ = fi19
                    _ = v20
                }
                _ = fi18
            }
            _ = v8
        }
    case 127:
        v1 -= (v1 >> (((v3 * v3)) & 63))
    }
    v21 := GenS2{F0: v1, F1: v2}
    if (!true) {
        v22 := GenS1{F0: "\n\t", F1: -1}
        if len(v6) > 0 {
            v6[0] = -1
        }
        if len(v6) > 0 {
            v6[0] = 1
        }
        _ = v22
    }
    if true {
        switch "foo bar" {
        case "hello":
            fmt.Println(v2, v3)
        case "foo bar":
            if ((!false) || false) {
                switch 733 {
                case -624:
                    for fi23 := 0; fi23 < 6; fi23++ {
                        v24 := GenS3{F0: v3}
                        _ = fi23
                        _ = v24
                    }
                case 32768:
                    v3 = (v3 >> ((uint(v5)) & 63))
                case 80:
                    v2 = v2
                default:
                    v2 = 42
                }
                _ = len(v6)
            } else {
                v5 = v5
                var v25 int8 = -48
                _ = v25
            }
        case "\n\t":
            v26 := GenS3{F0: (v3 ^ (v3 * (v3 * 4294967295)))}
            _ = v26
        }
        _ = len(v6)
        var v27 bool = false
        _ = v27
    } else {
        v28 := GenS1{F0: "\n\t", F1: 0}
        fmt.Println(v2)
        _ = v28
    }
    for fi29 := 0; fi29 < 2; fi29++ {
        for fi30 := 0; fi30 < 13; fi30++ {
            var v31 float32 = float32(((((((3.14159 * 1.5) / (((3.14159 - -2.25) / 1.5) - 0.0)) * (0.0 - (0.1 * 1e10))) / (((-2.25 * ((-2.25 - 0.0) / (1e10 / (0.0 - -2.25)))) * 0.1) - 0.0)) + 0.0) * (-2.25 / (-2.25 * 3.14159))))
            var v32 bool = false
            _ = fi30
            _ = v31
            _ = v32
        }
        var v33 uint64 = 2
        for fi34 := 0; fi34 < 13; fi34++ {
            if true {
                var v35 int16 = (v2 % (((v2 * int16(v1))) | 1))
                if ((0.0 + (0.1 / 1.0)) == (0.1 / (((1.5 / (3.14159 * 0.1)) * 0.1) + (3.14159 + ((3.14159 + 0.1) - 0.0))))) {
                    v5 += v5
                    if true {
                        if ((("foo bar" == "\n\t") || false) && (!(0 <= 1))) {
                            v3 -= v3
                            fmt.Println(fi34, v33, fi29)
                            v6 = append(v6, (fi29 & fi29))
                        } else {
                            v36 := GenS4{F0: (v5 | (^(v5))), F1: (^(v33))}
                            _ = v36
                        }
                        if false {
                            var v37 int8 = 100
                            if len(v6) > 0 {
                                v6[0] = fi29
                            }
                            _ = v37
                        }
                        v2 -= 0
                    } else {
                        v5 = 1
                        var v38 string = "foo bar"
                        _ = v38
                    }
                    var v39 string = ""
                    _ = v39
                }
                v2 = v2
                _ = v35
            } else {
                switch int(v1) {
                case 32768:
                    var v40 int64 = 0
                    _ = v40
                case 715:
                    if (-1 >= 122) {
                        var v41 uint64 = (v33 * 2)
                        _ = v41
                    } else {
                        v3 += v3
                        var v42 uint32 = 2
                        _ = v42
                    }
                }
                var v43 uint64 = (v33 + v33)
                switch (fi29 >> (((^(v3))) & 63)) {
                case 32767:
                    _ = len(v6)
                default:
                    v44 := GenS3{F0: v3}
                    _ = v44
                }
                _ = v43
            }
            _ = fi34
        }
        _ = fi29
        _ = v33
    }
    if true {
        v45 := GenS2{F0: (v1 % ((v1) | 1)), F1: v2}
        v6 = append(v6, 513)
        _ = v45
    } else {
        var v46 uint8 = (^(v5))
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
