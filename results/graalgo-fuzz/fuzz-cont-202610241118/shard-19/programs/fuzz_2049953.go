package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool; F1 uint8; F2 int8 }

func main() {
    var v1 bool = (("a" + "hello") < ("hello" + ""))
    var v2 bool = (v1 && true)
    var v3 int = 32767
    v4 := []string{}
    v5 := []int{(-(v3)), (v3 * v3)}
    var v6 int = (v3 ^ 32767)
    v3 += (v6 >> ((42) & 63))
    v2 = (((((3.14159 * -2.25) < (1.5 / ((((-2.25 - (((0.1 / 1.5) - float64(3.14159)) - 0.0)) - (1e10 - 0.1)) + -2.25) * 0.1))) || (v2 || v1)) || (849 > 80)) && v1)
    if v2 {
        if true {
            for fi7 := 0; fi7 < 3; fi7++ {
                for fi8 := 0; fi8 < 17; fi8++ {
                    if false {
                        var v9 uint8 = 128
                        _ = v9
                    }
                    var v10 uint32 = 42
                    if v1 {
                        switch (fi7 ^ v3) {
                        case -801:
                            v11 := GenS1{F0: 1000}
                            _ = v11
                        case 32768:
                            if true {
                                v5 = append(v5, 4294967295)
                                var v12 int16 = 42
                                fmt.Println(fi7, v2)
                                _ = v12
                            } else {
                                switch v6 {
                                case -9223372036854775808:
                                    v13 := GenS2{F0: (v10 + (v10 ^ v10)), F1: 255, F2: -320}
                                    _ = v13
                                case 2147483647:
                                    if ((v6 >> ((0) & 63)) >= 9223372036854775807) {
                                        fmt.Println(v2)
                                        v6 += fi8
                                        switch "\n\t" {
                                        case "a":
                                            var v14 float64 = 0.0
                                            _ = v14
                                        case "hello":
                                            v15 := GenS2{F0: (v10 >> ((610) & 63)), F1: 2, F2: 792}
                                            _ = v15
                                        default:
                                            var v16 bool = ("hello" < "a")
                                            v1 = false
                                            if ((116 > 128) || (fi8 < (fi7 >> ((uint(v10)) & 63)))) {
                                                var v17 float64 = (0.0 / 1.5)
                                                if true {
                                                    if false {
                                                        v2 = true
                                                    }
                                                    v18 := GenS3{F0: (779 > 1)}
                                                    v3 = -2147483648
                                                    _ = v18
                                                } else {
                                                    switch "a" {
                                                    case "hello":
                                                        fmt.Println(v6, v17)
                                                    case "foo bar":
                                                        v19 := GenS1{F0: 0}
                                                        _ = v19
                                                    }
                                                }
                                                _ = v17
                                            }
                                            _ = v16
                                        }
                                    }
                                default:
                                    v20 := GenS4{F0: 118}
                                    _ = v20
                                }
                            }
                        default:
                            if len(v4) > 0 {
                                v4[0] = ("\n\t" + "\n\t")
                            }
                        }
                        if (2 == 255) {
                            v2 = true
                            fmt.Println(v6, v6, v3)
                        }
                        fmt.Println(v6)
                    }
                    _ = fi8
                    _ = v10
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v21 := GenS5{F0: v1, F1: 42, F2: -111}
                _ = fi7
                _ = v21
            }
            switch "\n\t" {
            case "hello":
                var v22 string = (("a" + ("\n\t" + "foo bar")) + "")
                _ = v22
            case "":
                for fi23 := 0; fi23 < 19; fi23++ {
                    if (!v1) {
                        _ = len(v5)
                        var v24 bool = (519 < 279)
                        v25 := GenS2{F0: 2147483648, F1: 127, F2: 1}
                        _ = v24
                        _ = v25
                    } else {
                        v26 := GenS5{F0: (!v1), F1: 0, F2: -1}
                        if v2 {
                            var v27 uint8 = 0
                            if len(v4) > 0 {
                                v4[0] = ("hello" + ("" + ""))
                            }
                            fmt.Println(v1, v6, v27)
                            _ = v27
                        }
                        v28 := GenS3{F0: (65535 < 65535)}
                        _ = v26
                        _ = v28
                    }
                    v2 = v1
                    _ = fi23
                }
            }
        }
        switch ("a" + "foo bar") {
        case "\n\t":
            v29 := GenS2{F0: 546, F1: 228, F2: 1000}
            _ = v29
        }
    } else {
        if true {
            fmt.Println(v2)
            if v1 {
                v2 = ((((!v2) || false) || ((false || ((v2 || v1) || ((1 != 0) || v1))) && ((((130 != 584) && (341 != 662)) && true) && true))) && true)
                v3 = (v3 & 9223372036854775807)
            }
        } else {
            var v30 string = ("\n\t" + ("\n\t" + "foo bar"))
            _ = v30
        }
        var v31 uint64 = 919
        for fi32 := 0; fi32 < 6; fi32++ {
            for fi33 := 0; fi33 < 13; fi33++ {
                v3 -= (v3 >> ((18446744073709551615) & 63))
                _ = fi33
            }
            v6 = fi32
            fmt.Println(v3, v3)
            _ = fi32
        }
        _ = v31
    }
    v34 := GenS4{F0: 42}
    if len(v4) > 0 {
        v4[0] = ((("hello" + ("" + "hello")) + "") + "a")
    }
    if (!v2) {
        v35 := GenS5{F0: (1.5 == -2.25), F1: 139, F2: 0}
        v6 -= (v6 - (v3 ^ v6))
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v35
    }
    var v36 bool = (!v2)
    v37 := GenS3{F0: v36}
    if v1 {
        var v38 int64 = -1
        _ = v38
    }
    fmt.Println(v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
