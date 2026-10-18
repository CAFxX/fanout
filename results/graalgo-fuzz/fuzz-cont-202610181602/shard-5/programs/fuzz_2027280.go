package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 bool; F1 uint; F2 bool }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 int; F1 int64; F2 int16 }
func genFn0(p0 bool, p1 uint8, p2 int32) uint32 {
    v7 := GenS1{F0: 42, F1: 42, F2: 845}
    v8 := []int{-118, -1}
    switch (("hello" + (((("hello" + "a") + "hello") + "hello") + (("hello" + (("foo bar" + ("" + "a")) + "a")) + ""))) + ("foo bar" + "a")) {
    case "a":
        var v9 float64 = 3.14159
        _ = v9
    case "":
        p1 -= p1
    }
    v10 := GenS2{F0: 0, F1: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    return 2147483647
}
func genFn1(p0 int) int8 {
    v11 := []int{-264, p0, p0}
    var v12 int32 = -767
    v13 := GenS2{F0: 2, F1: 180}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 0
}

func main() {
    v1 := []int{}
    var v2 float32 = 1.5
    var v3 float32 = (1.5 + v2)
    var v4 uint8 = 127
    var v5 uint16 = 0
    v6 := make(map[int]bool)
    v14 := GenS2{F0: 2, F1: 2147483647}
    var v15 int64 = 9223372036854775807
    if false {
        for fi16 := 0; fi16 < 11; fi16++ {
            switch ("hello" + ("\n\t" + "a")) {
            case "hello":
                var v17 int64 = (v15 >> ((2) & 63))
                _ = v17
            case "a":
                switch (("foo bar" + ("hello" + "foo bar")) + (("hello" + (("a" + "") + ("foo bar" + "foo bar"))) + "a")) {
                case "a":
                    for fi18 := 0; fi18 < 16; fi18++ {
                        v3 = (-2.25 * v3)
                        v19 := GenS3{F0: 0, F1: -2147483648}
                        v6[(fi16 - -1)] = (false || (1 <= 290))
                        v20 := GenS1{F0: 1000, F1: 65535, F2: 2}
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                case "":
                    v21 := GenS3{F0: 65535, F1: (fi16 << ((65535) & 63))}
                    _ = v21
                case "\n\t":
                    var v22 int64 = v15
                    _ = v22
                default:
                    if ((v5 + (^(v5))) <= (v5 & 146)) {
                        for fi23 := 0; fi23 < 15; fi23++ {
                            v4 = v4
                            var v24 int16 = 42
                            _ = fi23
                            _ = v24
                        }
                    }
                }
            default:
                if (v2 == -2.25) {
                    switch fi16 {
                    case 32768:
                        delete(v6, (-(fi16)))
                    case 42:
                        var v25 uint16 = v5
                        _ = v25
                    case 0:
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v26 int32 = 1
                            v27 := GenS4{F0: 2, F1: "hello"}
                            _ = v26
                            _ = v27
                        }
                    }
                }
            }
            switch "\n\t" {
            case "\n\t":
                v3 = v2
            }
            _ = fi16
        }
        var v28 float32 = (1e10 * v2)
        _ = v28
    } else {
        var v29 int8 = -1
        _ = v29
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        fmt.Println(v3, v15)
        v2 = v2
        _ = fi30
    }
    _ = len(v1)
    var v31 uint16 = 2
    var v32 uint64 = 673
    for fi33 := 0; fi33 < 2; fi33++ {
        switch (fi33 + -992) {
        case -2147483648:
            if true {
                v3 -= (1e10 / 1.5)
                var v34 uint8 = v4
                v35 := GenS5{F0: true, F1: 461, F2: ((((v32 >> ((0) & 63)) > (v32 << ((42) & 63))) || true) && true)}
                _ = v34
                _ = v35
            } else {
                v2 -= v2
            }
        case 1:
            fmt.Println(v5)
        case -9223372036854775808:
            var v36 int16 = 0
            _ = v36
        }
        var v37 int64 = v15
        for fi38 := 0; fi38 < 17; fi38++ {
            var v39 uint = 42
            _ = fi38
            _ = v39
        }
        _ = fi33
        _ = v37
    }
    if ((v2 > (0.0 * ((v2 * -2.25) - v2))) && (-1 >= -1)) {
        _ = len(v6)
    }
    for fi40 := 0; fi40 < 2; fi40++ {
        var v41 bool = true
        _ = fi40
        _ = v41
    }
    if false {
        delete(v6, 873)
        v42 := GenS6{F0: -1, F1: 4294967295}
        _ = v42
    } else {
        var v43 uint16 = v31
        _ = v43
    }
    if ((v4 / ((v4) | 1)) == (^(v4))) {
        var v44 bool = false
        _ = v44
    }
    v32 -= (v32 % ((v32) | 1))
    v45 := GenS7{F0: 656, F1: (^(v15)), F2: 1000}
    if false {
        v31 = (v5 << ((327) & 63))
    } else {
        v1 = append(v1, -380)
        if (float64(1.5) > ((-2.25 * (1e10 * (-2.25 + 0.0))) + 1.5)) {
            for fi46 := 0; fi46 < 5; fi46++ {
                var v47 int64 = -138
                _ = fi46
                _ = v47
            }
            if ((v32 >> ((2) & 63)) >= (v32 << ((625) & 63))) {
                v5 = 2
                fmt.Println(v4, v2)
                var v48 int16 = -1
                _ = v48
            } else {
                if true {
                    v2 -= ((v2 - v3) - (((v2 + 3.14159) * 0.1) / v2))
                    for fi49 := 0; fi49 < 8; fi49++ {
                        for fi50 := 0; fi50 < 6; fi50++ {
                            v5 = v5
                            v5 = uint16(v32)
                            fmt.Println(v2, v32)
                            _ = fi50
                        }
                        _ = fi49
                    }
                    for fi51 := 0; fi51 < 2; fi51++ {
                        var v52 float64 = 1.5
                        for fi53 := 0; fi53 < 6; fi53++ {
                            v15 -= v15
                            if len(v1) > 0 {
                                v1[0] = fi51
                            }
                            _ = fi53
                        }
                        _ = fi51
                        _ = v52
                    }
                } else {
                    switch "hello" {
                    case "":
                        v54 := GenS2{F0: (v32 >> ((1) & 63)), F1: 2147483647}
                        _ = v54
                    case "\n\t":
                        v55 := GenS7{F0: 42, F1: v15, F2: -1}
                        _ = v55
                    default:
                        v56 := GenS1{F0: int16(v32), F1: genFn0(false, v4, -1), F2: (^(v32))}
                        _ = v56
                    }
                    fmt.Println(v3, v32, v2)
                }
            }
            for fi57 := 0; fi57 < 16; fi57++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = len(v6)
                var v58 int8 = 100
                _ = fi57
                _ = v58
            }
        } else {
            var v59 uint16 = 573
            _ = v59
        }
        v60 := v6[0]
        _ = v60
    }
    switch 2147483647 {
    case 1:
        for fi61 := 0; fi61 < 19; fi61++ {
            fmt.Println(v32)
            _ = fi61
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
