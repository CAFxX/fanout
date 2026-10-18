package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint8 }

func main() {
    var v1 int64 = 0
    var v2 bool = (!(!false))
    var v3 int32 = 1
    var v4 int16 = -32768
    var v5 float64 = 1.5
    v6 := []int{}
    if false {
        v2 = (4294967295 == 255)
        switch 9223372036854775807 {
        case 343:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        case 1:
            v7 := GenS1{F0: 65535}
            _ = v7
        case 42:
            v5 -= (v5 - ((1.5 * 1.5) + 3.14159))
        }
    }
    v5 -= float64((-2.25 - 0.0))
    fmt.Println(v5)
    if len(v6) > 0 {
        v6[0] = 0
    }
    switch 4294967295 {
    case 127:
        if true {
            for fi8 := 0; fi8 < 17; fi8++ {
                if ((!((v1 * (v1 >> ((65535) & 63))) != v1)) && (!(!v2))) {
                    if (58 < 42) {
                        v9 := GenS2{F0: (v3 + v3), F1: 2}
                        switch "a" {
                        case "a":
                            if false {
                                var v10 int32 = (v3 & v3)
                                var v11 float32 = float32((-2.25 / v5))
                                var v12 bool = (((v5 / ((3.14159 / v5) - 1.5)) / -2.25) < (0.0 - v5))
                                _ = v10
                                _ = v11
                                _ = v12
                            } else {
                                v2 = false
                                switch "a" {
                                case "a":
                                    v3 += v3
                                }
                            }
                        }
                        _ = len(v6)
                        _ = v9
                    } else {
                        for fi13 := 0; fi13 < 15; fi13++ {
                            switch "hello" {
                            case "foo bar":
                                var v14 uint = 255
                                _ = v14
                            case "a":
                                if ((!((255 != 4294967295) && v2)) && true) {
                                    if ((42 >= v4) || true) {
                                        v15 := GenS1{F0: 42}
                                        _ = v15
                                    }
                                    v16 := GenS3{F0: 89}
                                    _ = v16
                                }
                            }
                            _ = fi13
                        }
                        if ((((v5 / 1.0) + float64(0.1)) / v5) <= float64((3.14159 / ((3.14159 + (1e10 + 1.5)) * 1.5)))) {
                            if (!(!false)) {
                                v1 -= (v1 ^ 42)
                                switch "foo bar" {
                                case "":
                                    switch -687 {
                                    case 2147483647:
                                        var v17 uint64 = 4294967295
                                        _ = v17
                                    }
                                case "foo bar":
                                    v18 := GenS4{F0: v3, F1: 65}
                                    _ = v18
                                }
                                var v19 float64 = float64((0.1 + (3.14159 / -2.25)))
                                _ = v19
                            }
                            var v20 uint = 25
                            switch ("a" + "foo bar") {
                            case "hello":
                                fmt.Println(v4, v2, v20)
                            default:
                                if len(v6) > 0 {
                                    v6[0] = (fi8 << (((v20 ^ (v20 << ((v20) & 63)))) & 63))
                                }
                            }
                            _ = v20
                        }
                    }
                    var v21 uint64 = 196
                    if false {
                        _ = len(v6)
                        if v2 {
                            v2 = v2
                            var v22 uint = 4294967295
                            fmt.Println(v4)
                            _ = v22
                        } else {
                            v21 -= 4294967295
                            for fi23 := 0; fi23 < 3; fi23++ {
                                v24 := GenS5{F0: "foo bar"}
                                fmt.Println(v1)
                                _ = fi23
                                _ = v24
                            }
                        }
                        switch "foo bar" {
                        case "\n\t":
                            if (!true) {
                                for fi25 := 0; fi25 < 12; fi25++ {
                                    v5 += (1e10 * 3.14159)
                                    var v26 uint64 = (v21 - v21)
                                    v27 := GenS1{F0: 1}
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                }
                                if (false && v2) {
                                    v2 = (v2 || v2)
                                } else {
                                    v28 := GenS5{F0: "\n\t"}
                                    _ = v28
                                }
                            }
                        default:
                            switch "\n\t" {
                            case "\n\t":
                                if v2 {
                                    var v29 string = "hello"
                                    if (v5 == (1e10 * float64(-2.25))) {
                                        v3 += -1
                                        for fi30 := 0; fi30 < 6; fi30++ {
                                            var v31 int32 = v3
                                            v3 += -1
                                            var v32 float32 = float32(((-2.25 * v5) / (0.1 * (0.0 + v5))))
                                            _ = fi30
                                            _ = v31
                                            _ = v32
                                        }
                                    }
                                    _ = v29
                                } else {
                                    fmt.Println(v2, v2)
                                    fmt.Println(v5)
                                    _ = len(v6)
                                }
                            case "foo bar":
                                fmt.Println(v21)
                            case "hello":
                                v4 = 0
                            default:
                                var v33 uint32 = 2
                                _ = v33
                            }
                        }
                    } else {
                        if v2 {
                            v21 = 0
                            v34 := GenS3{F0: 42}
                            var v35 int16 = v4
                            _ = v34
                            _ = v35
                        } else {
                            var v36 uint64 = v21
                            for fi37 := 0; fi37 < 3; fi37++ {
                                v1 = v1
                                v6 = append(v6, (fi8 + 887))
                                _ = fi37
                            }
                            _ = v36
                        }
                    }
                    _ = v21
                } else {
                    v3 -= 135
                    if len(v6) > 0 {
                        v6[0] = (^(fi8))
                    }
                }
                if (v4 > (v4 * 32767)) {
                    _ = len(v6)
                    var v38 uint64 = 18446744073709551615
                    _ = v38
                }
                var v39 uint = 2
                _ = fi8
                _ = v39
            }
            v6 = append(v6, -833)
        } else {
            var v40 int32 = v3
            v4 -= (v4 & (v4 / ((1) | 1)))
            v6 = append(v6, 128)
            _ = v40
        }
    case -2147483648:
        v5 = (0.0 - 0.1)
    default:
        var v41 float32 = 1e10
        _ = v41
    }
    var v42 uint8 = 42
    v4 += (-(v4))
    var v43 int64 = (v1 - v1)
    _ = len(v6)
    for fi44 := 0; fi44 < 7; fi44++ {
        v3 -= (-(v3))
        for fi45 := 0; fi45 < 18; fi45++ {
            var v46 int16 = (v4 << ((345) & 63))
            _ = fi45
            _ = v46
        }
        _ = fi44
    }
    for fi47 := 0; fi47 < 12; fi47++ {
        v42 += (^(v42))
        _ = fi47
    }
    var v48 float64 = (v5 + 1.5)
    var v49 float32 = 0.0
    v50 := GenS6{F0: 0}
    var v51 uint64 = 18446744073709551615
    for fi52 := 0; fi52 < 12; fi52++ {
        var v53 float64 = float64(float32(float64(1.5)))
        for fi54 := 0; fi54 < 6; fi54++ {
            v1 -= -43
            var v55 int64 = v43
            _ = fi54
            _ = v55
        }
        _ = fi52
        _ = v53
    }
    v56 := GenS3{F0: v42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v56
    _ = v6
}
