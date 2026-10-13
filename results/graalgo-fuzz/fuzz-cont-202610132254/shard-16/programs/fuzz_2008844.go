package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint16 }
type GenS2 struct { F0 int8; F1 int8; F2 int8 }
type GenS3 struct { F0 uint32; F1 int64; F2 string }
type GenS4 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS5 struct { F0 uint32 }

func main() {
    var v1 uint16 = 795
    var v2 float32 = 1e10
    v3 := []uint32{}
    var v4 int8 = 127
    var v5 uint16 = v1
    var v6 uint = 9
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 uint16 = (v1 / ((v5) | 1))
        _ = fi7
        _ = v8
    }
    fmt.Println(v4, v6)
    var v9 int64 = 1
    if true {
        if (false || false) {
            var v10 float64 = ((1.5 * (3.14159 - 0.0)) + 1e10)
            var v11 int32 = 446
            for fi12 := 0; fi12 < 9; fi12++ {
                fmt.Println(v2, v11, v9)
                _ = fi12
            }
            _ = v10
            _ = v11
        }
        for fi13 := 0; fi13 < 11; fi13++ {
            if false {
                v14 := GenS1{F0: (v9 >> ((v6) & 63)), F1: fi13, F2: (^(v5))}
                v4 = (-(v4))
                _ = v14
            }
            v9 = (v9 | v9)
            _ = fi13
        }
    } else {
        if ((0 >= v9) && false) {
            if false {
                if len(v3) > 0 {
                    v3[0] = 725
                }
            } else {
                for fi15 := 0; fi15 < 18; fi15++ {
                    if (!true) {
                        for fi16 := 0; fi16 < 5; fi16++ {
                            var v17 int32 = 1
                            _ = fi16
                            _ = v17
                        }
                        for fi18 := 0; fi18 < 19; fi18++ {
                            v19 := GenS1{F0: int64(fi18), F1: fi15, F2: 2}
                            _ = fi18
                            _ = v19
                        }
                    } else {
                        for fi20 := 0; fi20 < 11; fi20++ {
                            var v21 uint8 = 53
                            if true {
                                fmt.Println(v4, fi15, fi15)
                            }
                            v22 := GenS2{F0: -117, F1: v4, F2: (v4 - (^(v4)))}
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                        for fi23 := 0; fi23 < 15; fi23++ {
                            if (2 > 4294967295) {
                                var v24 uint8 = 127
                                var v25 int64 = v9
                                _ = v24
                                _ = v25
                            }
                            if false {
                                var v26 float64 = (0.1 * 0.1)
                                fmt.Println(fi23, v5, v2)
                                _ = v26
                            } else {
                                fmt.Println(v2)
                            }
                            fmt.Println(v6, fi23, v9)
                            _ = fi23
                        }
                        fmt.Println(v9, v6)
                    }
                    if len(v3) > 0 {
                        v3[0] = 42
                    }
                    _ = fi15
                }
            }
            for fi27 := 0; fi27 < 7; fi27++ {
                switch ("" + (("" + "hello") + "\n\t")) {
                case "":
                    var v28 uint32 = 2147483648
                    _ = v28
                case "\n\t":
                    v29 := GenS1{F0: v9, F1: 127, F2: v5}
                    _ = v29
                }
                _ = fi27
            }
        } else {
            for fi30 := 0; fi30 < 11; fi30++ {
                if true {
                    v31 := GenS1{F0: v9, F1: (fi30 - (^(fi30))), F2: v1}
                    if len(v3) > 0 {
                        v3[0] = 4294967295
                    }
                    v32 := GenS1{F0: 0, F1: fi30, F2: v5}
                    _ = v31
                    _ = v32
                } else {
                    for fi33 := 0; fi33 < 8; fi33++ {
                        v3 = append(v3, 42)
                        var v34 int32 = 902
                        v34 = (v34 >> (((v6 * v6)) & 63))
                        _ = fi33
                        _ = v34
                    }
                    v2 -= 1.5
                }
                _ = fi30
            }
            var v35 string = "a"
            _ = v35
        }
        _ = len(v3)
        var v36 bool = ((float32(1e10) + v2) == v2)
        _ = v36
    }
    var v37 float32 = (1e10 * v2)
    for fi38 := 0; fi38 < 6; fi38++ {
        for fi39 := 0; fi39 < 7; fi39++ {
            switch ("" + ("a" + "foo bar")) {
            case "foo bar":
                var v40 int64 = -1
                _ = v40
            case "\n\t":
                v41 := GenS3{F0: 1, F1: (v9 * v9), F2: "\n\t"}
                _ = v41
            }
            _ = fi39
        }
        for fi42 := 0; fi42 < 14; fi42++ {
            v43 := GenS1{F0: (v9 ^ (v9 ^ (v9 * (v9 & (v9 << (((v6 & v6)) & 63)))))), F1: -9223372036854775808, F2: 1}
            var v44 uint64 = 291
            v45 := GenS4{F0: v1, F1: -1, F2: (v44 + (^(v44)))}
            if ("hello" < ("a" + "")) {
                var v46 float32 = v2
                _ = v46
            } else {
                if len(v3) > 0 {
                    v3[0] = 0
                }
                var v47 uint32 = 65535
                _ = v47
            }
            _ = fi42
            _ = v43
            _ = v44
            _ = v45
        }
        _ = fi38
    }
    v5 -= v5
    fmt.Println(v1, v6, v37)
    switch 42 {
    case 32767:
        v3 = append(v3, 4294967295)
    case 9223372036854775807:
        for fi48 := 0; fi48 < 1; fi48++ {
            v49 := GenS5{F0: 1}
            v2 = float32(1e10)
            _ = fi48
            _ = v49
        }
    }
    v50 := GenS1{F0: -901, F1: 801, F2: (v5 | v5)}
    v37 = -2.25
    v3 = append(v3, 516)
    var v51 uint64 = 42
    fmt.Println(v1, v9, v51)
    var v52 uint16 = v5
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
    _ = v9
}
