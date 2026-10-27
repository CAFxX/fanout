package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 bool }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 uint8; F1 int32; F2 int32 }

func main() {
    var v1 int64 = 42
    v2 := make(map[int]bool)
    v3 := make(map[int]int)
    var v4 float32 = (1e10 + 0.0)
    v5 := []int{-41}
    var v6 int64 = 42
    for fi7 := 0; fi7 < 16; fi7++ {
        v1 = v1
        _ = fi7
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int64 = v1
        _ = fi8
        _ = v9
    }
    switch 906 {
    case 0:
        switch 546 {
        case 32767:
            var v10 int = 32767
            _ = v10
        case -1:
            fmt.Println(v6, v6, v4)
        case 4294967295:
            var v11 int32 = -481
            _ = v11
        }
    default:
        v12 := GenS1{F0: 0, F1: -128, F2: false}
        _ = v12
    }
    var v13 uint32 = 4
    var v14 int32 = -1
    var v15 uint8 = 0
    for fi16 := 0; fi16 < 2; fi16++ {
        _ = len(v5)
        _ = fi16
    }
    switch ("hello" + (("a" + ("" + "\n\t")) + ("a" + ("foo bar" + (("\n\t" + "a") + ""))))) {
    case "foo bar":
        v4 -= v4
    case "\n\t":
        v6 += (v6 - v1)
    case "a":
        for fi17 := 0; fi17 < 7; fi17++ {
            switch ("hello" + "\n\t") {
            case "":
                v18 := GenS2{F0: (v13 - (v13 / (((v13 >> ((18446744073709551615) & 63))) | 1))), F1: 316}
                _ = v18
            default:
                var v19 float32 = 1.5
                _ = v19
            }
            var v20 uint32 = v13
            _ = fi17
            _ = v20
        }
    default:
        v13 -= 2147483647
    }
    v21 := GenS2{F0: 4294967295, F1: 1}
    if false {
        for fi22 := 0; fi22 < 1; fi22++ {
            for fi23 := 0; fi23 < 7; fi23++ {
                switch "\n\t" {
                case "hello":
                    fmt.Println(v15)
                case "":
                    if (fi23 == fi23) {
                        switch ("\n\t" + "a") {
                        case "foo bar":
                            if len(v5) > 0 {
                                v5[0] = (fi23 | fi22)
                            }
                        default:
                            if (!false) {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v4 += 1.5
                            }
                        }
                        v24 := GenS3{F0: "a", F1: 526}
                        switch ("\n\t" + "a") {
                        case "":
                            delete(v3, (-(fi23)))
                        }
                        _ = v24
                    } else {
                        v13 = v13
                        _ = len(v3)
                    }
                }
                fmt.Println(v4, fi23)
                v15 = 128
                if true {
                    switch "a" {
                    case "":
                        fmt.Println(v1, v14, v1)
                    case "a":
                        fmt.Println(fi22, v4, fi22)
                    }
                    v5 = append(v5, (fi22 & -636))
                }
                var v25 uint64 = 42
                _ = fi23
                _ = v25
            }
            _ = fi22
        }
        v26 := GenS2{F0: (v13 + v13), F1: 1}
        for fi27 := 0; fi27 < 12; fi27++ {
            v28 := GenS2{F0: 0, F1: 2}
            switch "\n\t" {
            case "hello":
                v14 -= (v14 + (v14 | (v14 << ((0) & 63))))
            case "\n\t":
                v4 -= (v4 - 1e10)
            case "a":
                fmt.Println(v13, v14)
            }
            _ = fi27
            _ = v28
        }
        _ = v26
    }
    v15 -= v15
    v13 -= (v13 | v13)
    for fi29 := 0; fi29 < 1; fi29++ {
        if true {
            v30 := GenS4{F0: 1, F1: -2147483648, F2: v14}
            if false {
                for fi31 := 0; fi31 < 13; fi31++ {
                    fmt.Println(fi31, v13, v15)
                    var v32 int = -300
                    _ = fi31
                    _ = v32
                }
                var v33 int32 = v14
                _ = v33
            } else {
                switch ("hello" + "foo bar") {
                case "":
                    for fi34 := 0; fi34 < 3; fi34++ {
                        var v35 int16 = 32767
                        _ = fi34
                        _ = v35
                    }
                case "foo bar":
                    if ((v14 >> ((4294967295) & 63)) >= v14) {
                        switch ("hello" + "a") {
                        case "a":
                            for fi36 := 0; fi36 < 9; fi36++ {
                                v6 = v1
                                _ = fi36
                            }
                        }
                        v1 -= (^(v1))
                    }
                case "a":
                    delete(v2, 466)
                }
                v6 += v1
            }
            _ = v30
        } else {
            if (v15 != uint8(v13)) {
                v14 = (^(v14))
                if true {
                    var v37 uint32 = 2147483647
                    _ = v37
                }
            }
        }
        _ = fi29
    }
    if false {
        v38 := GenS3{F0: "", F1: 594}
        _ = v38
    }
    switch "a" {
    case "a":
        v39 := GenS2{F0: 4294967295, F1: 1}
        _ = v39
    case "foo bar":
        var v40 int64 = v6
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
