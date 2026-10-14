package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 int32; F2 int16 }
type GenS4 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int64; F1 bool; F2 int64 }
type GenS7 struct { F0 uint64; F1 bool; F2 int16 }

func main() {
    v1 := []uint32{}
    v2 := []uint32{42, 0, 563, 2147483648}
    v3 := make(map[int]int)
    var v4 int8 = 42
    if true {
        if true {
            if (65535 < 2147483648) {
                for fi5 := 0; fi5 < 18; fi5++ {
                    v6 := GenS1{F0: fi5, F1: 1, F2: -59}
                    var v7 uint32 = 42
                    fmt.Println(v7)
                    _ = fi5
                    _ = v6
                    _ = v7
                }
                for fi8 := 0; fi8 < 9; fi8++ {
                    for fi9 := 0; fi9 < 4; fi9++ {
                        if (v4 > 127) {
                            fmt.Println(fi9, fi8)
                            if (!true) {
                                if (false || (1 < 128)) {
                                    v4 += v4
                                    var v10 uint = 18446744073709551615
                                    v11 := GenS1{F0: int(v4), F1: 65535, F2: 42}
                                    _ = v10
                                    _ = v11
                                } else {
                                    var v12 int64 = 42
                                    v4 += v4
                                    _ = v12
                                }
                                if true {
                                    v13 := GenS1{F0: (fi8 ^ fi9), F1: 2, F2: -32768}
                                    var v14 uint16 = 166
                                    _ = v13
                                    _ = v14
                                }
                            } else {
                                fmt.Println(fi8, v4, v4)
                                if false {
                                    if false {
                                        v4 += (^(v4))
                                        _ = len(v3)
                                    }
                                    var v15 uint16 = 65535
                                    v16 := v3[-168]
                                    _ = v15
                                    _ = v16
                                } else {
                                    if (-813 == 42) {
                                        v2 = append(v2, 1)
                                        var v17 uint32 = 22
                                        switch fi9 {
                                        case 42:
                                            v18 := GenS2{F0: -2147483648}
                                            _ = v18
                                        case 22:
                                            if (false || (0 > 0)) {
                                                v4 += -54
                                                if (-2147483648 != 760) {
                                                    if true {
                                                        var v19 uint32 = v17
                                                        _ = v19
                                                    } else {
                                                        v20 := GenS3{F0: -1, F1: 0, F2: 0}
                                                        _ = v20
                                                    }
                                                    switch ("a" + "hello") {
                                                    case "foo bar":
                                                        if (!false) {
                                                            v21 := GenS4{F0: 527, F1: (v17 >> ((255) & 63)), F2: -418}
                                                            _ = v21
                                                        } else {
                                                            v17 -= v17
                                                            v22 := GenS3{F0: -1, F1: 362, F2: -32768}
                                                            fmt.Println(v17, fi8, fi9)
                                                            _ = v22
                                                        }
                                                    }
                                                    v23 := GenS1{F0: (fi9 ^ (fi9 & 549)), F1: (v17 & v17), F2: 0}
                                                    _ = v23
                                                } else {
                                                    v24 := v3[fi8]
                                                    _ = v24
                                                }
                                                v4 += 0
                                            } else {
                                                v4 = (-(v4))
                                                v4 += -114
                                                fmt.Println(fi8)
                                            }
                                        case 9223372036854775807:
                                            var v25 int = fi9
                                            _ = v25
                                        default:
                                            if ((0 > (v17 << ((65535) & 63))) && true) {
                                                var v26 uint64 = 717
                                                v17 -= 42
                                                _ = v26
                                            } else {
                                                var v27 uint16 = 0
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                _ = v27
                                            }
                                        }
                                        switch fi9 {
                                        case 128:
                                            v17 += (^(v17))
                                        }
                                        _ = v17
                                    }
                                    v4 += v4
                                }
                                v28 := GenS5{F0: -1}
                                v29 := GenS2{F0: -639}
                                _ = len(v2)
                                v4 += (v4 - v4)
                                var v30 int64 = -1
                                var v31 uint32 = 849
                                _ = v28
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                        }
                        v4 += v4
                        v4 = 79
                        var v32 int = -669
                        _ = fi9
                        _ = v32
                    }
                    var v33 uint = 0
                    _ = fi8
                    _ = v33
                }
                for fi34 := 0; fi34 < 1; fi34++ {
                    var v35 string = "foo bar"
                    if true {
                        var v36 uint8 = 42
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        delete(v3, fi34)
                        _ = v36
                    }
                    _ = fi34
                    _ = v35
                }
            }
        }
        for fi37 := 0; fi37 < 10; fi37++ {
            var v38 uint16 = 2
            if (true && false) {
                switch 127 {
                case 128:
                    v38 -= 0
                }
            }
            for fi39 := 0; fi39 < 17; fi39++ {
                var v40 float32 = ((-2.25 - 3.14159) / 1.5)
                var v41 float64 = 1e10
                v42 := GenS6{F0: -1, F1: false, F2: -9223372036854775808}
                _ = fi39
                _ = v40
                _ = v41
                _ = v42
            }
            _ = fi37
            _ = v38
        }
    }
    if false {
        if (!("\n\t" <= ((((("foo bar" + ((("\n\t" + ("foo bar" + ("a" + ((("a" + "") + ("foo bar" + "foo bar")) + ("\n\t" + "\n\t"))))) + "foo bar") + "hello")) + "foo bar") + "a") + "hello") + ""))) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v43 float64 = -2.25
            _ = v43
        } else {
            v4 += (v4 % (((v4 % ((v4) | 1))) | 1))
            var v44 int8 = v4
            var v45 bool = ((0.1 / 1.0) <= float64(1.5))
            _ = v44
            _ = v45
        }
        v3[128] = -255
    } else {
        fmt.Println(v4)
        var v46 uint = 1
        var v47 uint32 = 2147483648
        _ = v46
        _ = v47
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v2)
    v4 = (v4 ^ 100)
    switch 1 {
    case 128:
        v48 := GenS1{F0: -938, F1: 919, F2: 32767}
        _ = v48
    case 2147483647:
        fmt.Println(v4)
    case 264:
        v49 := v3[-2147483648]
        _ = v49
    }
    var v50 int64 = 0
    fmt.Println(v50, v50)
    for fi51 := 0; fi51 < 6; fi51++ {
        if true {
            var v52 uint16 = 65535
            _ = v52
        }
        _ = fi51
    }
    v53 := GenS7{F0: 1, F1: (("foo bar" + ("" + "\n\t")) < ""), F2: -603}
    v3[0] = 0
    var v54 int64 = -600
    var v55 uint8 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v53
    _ = v54
    _ = v55
}
