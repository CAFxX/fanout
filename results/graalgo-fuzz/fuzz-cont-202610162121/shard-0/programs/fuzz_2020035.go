package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 bool; F1 uint; F2 int }
type GenS5 struct { F0 string; F1 int8 }
type GenS6 struct { F0 bool; F1 int8; F2 string }

func main() {
    var v1 int = 32767
    var v2 float64 = (-2.25 * 1e10)
    var v3 float32 = 0.0
    for fi4 := 0; fi4 < 14; fi4++ {
        var v5 int16 = -825
        _ = fi4
        _ = v5
    }
    v2 = v2
    v6 := GenS1{F0: (v1 + v1)}
    var v7 int16 = 42
    v8 := GenS1{F0: (v1 & 4294967295)}
    if (!(false || false)) {
        fmt.Println(v3)
        v2 -= -2.25
        var v9 int = (^(v1))
        _ = v9
    }
    switch ("hello" + ("hello" + "a")) {
    case "\n\t":
        var v10 uint = 279
        _ = v10
    default:
        var v11 bool = (!false)
        _ = v11
    }
    var v12 int64 = 9223372036854775807
    if false {
        var v13 uint64 = 4294967295
        _ = v13
    }
    v12 = (v12 & v12)
    var v14 float32 = 3.14159
    v15 := []int{int(v12)}
    if (true && false) {
        fmt.Println(v7, v1)
        for fi16 := 0; fi16 < 6; fi16++ {
            fmt.Println(v3, v7, v1)
            if true {
                if true {
                    v17 := GenS2{F0: 2, F1: (!(42 > -1))}
                    for fi18 := 0; fi18 < 15; fi18++ {
                        var v19 uint8 = 25
                        _ = len(v15)
                        _ = fi18
                        _ = v19
                    }
                    if ((!true) && true) {
                        v20 := GenS1{F0: (fi16 ^ v1)}
                        v1 = -2147483648
                        _ = v20
                    }
                    _ = v17
                } else {
                    for fi21 := 0; fi21 < 9; fi21++ {
                        var v22 float32 = -2.25
                        if (false || (2 <= 57)) {
                            fmt.Println(fi16)
                            _ = len(v15)
                            var v23 int64 = v12
                            _ = v23
                        }
                        _ = len(v15)
                        v15 = append(v15, fi16)
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        _ = fi21
                        _ = v22
                    }
                    v15 = append(v15, (v1 % (((-(fi16))) | 1)))
                }
            } else {
                for fi24 := 0; fi24 < 7; fi24++ {
                    v12 = (-(v12))
                    fmt.Println(v14, v14)
                    if true {
                        v15 = append(v15, (^(fi24)))
                        switch fi24 {
                        case 32767:
                            v2 += v2
                        case 32768:
                            fmt.Println(v3, v7)
                        }
                        v25 := GenS1{F0: -1}
                        _ = v25
                    } else {
                        fmt.Println(v3, v14)
                        switch 127 {
                        case 1:
                            fmt.Println(fi24)
                        }
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                    }
                    fmt.Println(v12)
                    v14 -= v14
                    _ = fi24
                }
                for fi26 := 0; fi26 < 5; fi26++ {
                    v27 := GenS3{F0: 42, F1: true}
                    v14 -= ((v3 - 1.5) * v3)
                    _ = fi26
                    _ = v27
                }
                var v28 uint = 218
                _ = v28
            }
            for fi29 := 0; fi29 < 8; fi29++ {
                v7 = v7
                switch fi29 {
                case 4294967295:
                    v2 -= v2
                default:
                    var v30 int32 = 10
                    _ = v30
                }
                var v31 float64 = 1.5
                if (false && (fi16 > fi29)) {
                    switch "a" {
                    case "\n\t":
                        var v32 string = "foo bar"
                        _ = v32
                    }
                    if len(v15) > 0 {
                        v15[0] = fi29
                    }
                } else {
                    v1 = fi29
                    if true {
                        if len(v15) > 0 {
                            v15[0] = (fi29 & (v1 >> ((42) & 63)))
                        }
                        switch v1 {
                        case -2147483648:
                            var v33 uint16 = 502
                            var v34 uint = 255
                            switch 32767 {
                            case 519:
                                if false {
                                    v35 := GenS4{F0: false, F1: (v34 & v34), F2: v1}
                                    v14 = v14
                                    _ = v35
                                } else {
                                    v36 := GenS5{F0: "foo bar", F1: 1}
                                    _ = v36
                                }
                            default:
                                if (true || (true || (((((v14 / float32(v2)) - (float32((3.14159 / 0.1)) + v3)) + ((v14 * float32(v31)) * v14)) > v3) && false))) {
                                    var v37 int = (v1 & (v1 - fi16))
                                    _ = v37
                                } else {
                                    var v38 int32 = -2147483648
                                    _ = v38
                                }
                            }
                            _ = v33
                            _ = v34
                        case 9223372036854775807:
                            switch "foo bar" {
                            case "\n\t":
                                v12 = -9223372036854775808
                            default:
                                v39 := GenS5{F0: "\n\t", F1: -128}
                                _ = v39
                            }
                        case 161:
                            var v40 int64 = v12
                            _ = v40
                        }
                    }
                }
                _ = fi29
                _ = v31
            }
            _ = fi16
        }
        v41 := GenS1{F0: 32768}
        _ = v41
    } else {
        if true {
            v12 -= (v12 + (v12 * -617))
            v2 -= -2.25
            if (!false) {
                if true {
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                }
            }
        }
        v1 = int(v7)
    }
    var v42 string = ("" + "\n\t")
    switch 128 {
    case -703:
        var v43 int16 = 42
        _ = v43
    case 0:
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
    }
    fmt.Println(v3, v3)
    for fi44 := 0; fi44 < 10; fi44++ {
        _ = len(v15)
        v45 := GenS6{F0: false, F1: -1, F2: "\n\t"}
        if (324 == 255) {
            fmt.Println(v42, v3)
            var v46 int8 = 42
            _ = v46
        } else {
            fmt.Println(v2, v2)
            v3 += (0.1 / 1.5)
        }
        _ = fi44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v42
    _ = v6
    _ = v7
    _ = v8
}
