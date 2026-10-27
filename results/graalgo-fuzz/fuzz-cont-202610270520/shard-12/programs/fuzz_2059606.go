package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 int; F2 uint32 }
type GenS5 struct { F0 uint16; F1 uint8; F2 bool }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 uint16; F1 uint16 }
type GenS8 struct { F0 int32; F1 uint16 }
type GenS9 struct { F0 uint64; F1 uint64 }
func genFn0(p0 uint) int16 {
    var v7 uint64 = 303
    var v8 int16 = 42
    _ = p0
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint32) int16 {
    v9 := []int{}
    if (true || (243 == 2)) {
        switch "foo bar" {
        case "\n\t":
            v10 := GenS1{F0: 0, F1: -74, F2: 65535}
            _ = v10
        }
        if (42 >= 42) {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            var v11 float64 = (((0.0 + 3.14159) + 1e10) * -2.25)
            v11 -= v11
            _ = v11
        } else {
            if (!false) {
                p0 = p0
                fmt.Println(p0)
            }
            if (65535 > 153) {
                switch "foo bar" {
                case "foo bar":
                    var v12 string = "hello"
                    _ = v12
                case "\n\t":
                    if (false || true) {
                        switch 32768 {
                        case 106:
                            var v13 uint16 = 42
                            _ = v13
                        case 2147483647:
                            if (-9223372036854775808 < -78) {
                                v14 := GenS2{F0: 559, F1: 1, F2: -128}
                                _ = v14
                            } else {
                                if true {
                                    fmt.Println(p0)
                                }
                                v15 := GenS1{F0: 65535, F1: 0, F2: 846}
                                _ = v15
                            }
                        }
                    }
                default:
                    v16 := GenS1{F0: 255, F1: 412, F2: (^(p0))}
                    _ = v16
                }
            } else {
                for fi17 := 0; fi17 < 8; fi17++ {
                    p0 -= 800
                    v18 := GenS3{F0: 127}
                    switch -1 {
                    case 42:
                        var v19 int8 = -57
                        _ = v19
                    default:
                        switch 0 {
                        case -492:
                            fmt.Println(p0, p0)
                        }
                    }
                    _ = fi17
                    _ = v18
                }
                p0 += (p0 % ((p0) | 1))
                v20 := GenS4{F0: -1, F1: int(p0), F2: 0}
                _ = v20
            }
        }
    } else {
        if false {
            var v21 string = "foo bar"
            _ = v21
        } else {
            var v22 uint = 4294967295
            v23 := GenS4{F0: -1, F1: 42, F2: 829}
            var v24 bool = ((p0 ^ p0) == (p0 & (p0 << ((18446744073709551615) & 63))))
            _ = v22
            _ = v23
            _ = v24
        }
        v25 := GenS5{F0: 0, F1: 208, F2: false}
        _ = v25
    }
    _ = p0
    _ = v9
    return 1000
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[int]int)
    var v3 int = 2147483647
    v4 := make(map[string]int)
    v5 := []string{"", ("a" + ((("foo bar" + "hello") + "foo bar") + ("hello" + "a"))), "foo bar"}
    var v6 uint16 = 790
    var v26 bool = true
    if (!false) {
        if ((float64(float32((3.14159 + 1.5))) + (0.1 / 1.0)) < 1e10) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v27 float64 = (0.1 + (float64(0.0) * -2.25))
            _ = v27
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            switch "a" {
            case "foo bar":
                var v28 int = v3
                _ = v28
            }
            for fi29 := 0; fi29 < 14; fi29++ {
                v30 := GenS6{F0: "\n\t", F1: 2}
                if ((false && v26) || (v26 || true)) {
                    for fi31 := 0; fi31 < 2; fi31++ {
                        v26 = (42 >= 110)
                        if true {
                            v32 := GenS2{F0: 42, F1: 255, F2: -128}
                            _ = v32
                        } else {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                        v1 -= (-(v1))
                        _ = fi31
                    }
                }
                v33 := GenS3{F0: -1}
                _ = fi29
                _ = v30
                _ = v33
            }
        }
        v34 := GenS6{F0: "a", F1: 161}
        if true {
            v6 += 0
            if (!(!true)) {
                fmt.Println(v26, v26, v1)
                for fi35 := 0; fi35 < 7; fi35++ {
                    v1 += (v1 ^ v1)
                    _ = fi35
                }
                fmt.Println(v6)
            }
            v1 = (v1 | (v1 ^ (v1 - (v1 & v1))))
        } else {
            fmt.Println(v1, v3, v26)
        }
        _ = v34
    }
    var v36 int64 = -1
    v37 := GenS7{F0: (v6 << ((255) & 63)), F1: 421}
    var v38 uint16 = v6
    v39 := GenS7{F0: 32767, F1: (v38 - (^(v6)))}
    v5 = append(v5, "")
    for fi40 := 0; fi40 < 3; fi40++ {
        v41 := GenS7{F0: 32767, F1: v38}
        _ = fi40
        _ = v41
    }
    v4[""] = (v3 << ((290) & 63))
    var v42 int16 = v1
    v43 := GenS8{F0: 42, F1: (v38 - v6)}
    if true {
        switch v3 {
        case 2147483647:
            switch "\n\t" {
            case "hello":
                var v44 uint = 0
                _ = v44
            case "\n\t":
                for fi45 := 0; fi45 < 11; fi45++ {
                    v46 := GenS2{F0: 449, F1: 128, F2: int8(fi45)}
                    _ = fi45
                    _ = v46
                }
            case "foo bar":
                if (float64(0.1) != float64((float32(0.1) + float32(0.1)))) {
                    var v47 bool = v26
                    _ = v47
                }
            }
        default:
            fmt.Println(v3, v26, v26)
        }
        v48 := GenS9{F0: 586, F1: 808}
        v3 += (v3 >> ((42) & 63))
        _ = v48
    } else {
        if v26 {
            v49 := GenS8{F0: 2147483647, F1: 2}
            v38 -= 569
            _ = v49
        }
        var v50 float64 = float64(1e10)
        _ = v50
    }
    v51 := GenS1{F0: v6, F1: -9223372036854775808, F2: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v51
    _ = v6
}
