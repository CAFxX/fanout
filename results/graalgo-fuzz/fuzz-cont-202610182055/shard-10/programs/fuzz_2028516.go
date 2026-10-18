package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int32; F1 uint8; F2 int32 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 uint; F2 uint }
type GenS6 struct { F0 int32; F1 int16; F2 bool }
type GenS7 struct { F0 uint; F1 int }

func main() {
    v1 := []int{32767, 1, 504}
    var v2 float64 = float64(3.14159)
    v3 := make(map[int]uint32)
    var v4 string = "hello"
    v2 = 0.1
    if (100 < -81) {
        _ = len(v3)
        v3[9223372036854775807] = 739
    } else {
        v4 = v4
        fmt.Println(v2)
        var v5 int16 = 1
        _ = v5
    }
    if false {
        if true {
            switch "\n\t" {
            case "":
                var v6 uint16 = 593
                _ = v6
            case "\n\t":
                var v7 int16 = 1000
                _ = v7
            }
            v8 := GenS1{F0: 42, F1: ((v4 + "hello") + v4)}
            _ = len(v3)
            _ = v8
        }
        v4 = ((("foo bar" + ((v4 + ("" + v4)) + v4)) + "foo bar") + v4)
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := GenS2{F0: 42, F1: 87, F2: 0}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 uint32 = 42
        var v13 bool = false
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int8 = -72
    v15 := GenS3{F0: 2, F1: true}
    if (false && (0 >= (^(v14)))) {
        for fi16 := 0; fi16 < 7; fi16++ {
            v17 := GenS4{F0: 490}
            _ = fi16
            _ = v17
        }
    }
    delete(v3, 9223372036854775807)
    _ = len(v3)
    switch 32768 {
    case -940:
        switch "foo bar" {
        case "\n\t":
            var v18 int32 = -1
            _ = v18
        case "hello":
            var v19 int32 = 42
            _ = v19
        default:
            if (!false) {
                var v20 int = -1
                switch v20 {
                case -995:
                    for fi21 := 0; fi21 < 12; fi21++ {
                        fmt.Println(v4, fi21, v4)
                        v22 := v3[fi21]
                        _ = fi21
                        _ = v22
                    }
                default:
                    v14 = 42
                }
                v23 := GenS5{F0: v14, F1: 208, F2: 1}
                _ = v20
                _ = v23
            }
        }
    case -1:
        var v24 float32 = float32(v2)
        _ = v24
    default:
        fmt.Println(v14, v14, v4)
    }
    for fi25 := 0; fi25 < 10; fi25++ {
        var v26 bool = (127 > (v14 | 42))
        _ = fi25
        _ = v26
    }
    v2 = 3.14159
    v27 := v3[918]
    if false {
        for fi28 := 0; fi28 < 2; fi28++ {
            _ = len(v1)
            var v29 uint = 450
            for fi30 := 0; fi30 < 1; fi30++ {
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v31 string = "hello"
                    v32 := GenS5{F0: (v14 | v14), F1: v29, F2: 46}
                    var v33 uint32 = 4294967295
                    var v34 bool = (v31 > "\n\t")
                    _ = v31
                    _ = v32
                    _ = v33
                    _ = v34
                } else {
                    var v35 string = "a"
                    _ = v35
                }
                _ = fi30
            }
            _ = fi28
            _ = v29
        }
        v27 += 2147483648
    }
    if false {
        if false {
            if ((((v14 & v14) == (-(v14))) && true) && true) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if true {
                    for fi36 := 0; fi36 < 6; fi36++ {
                        v37 := GenS1{F0: 42, F1: ("\n\t" + "")}
                        if (966 >= 484) {
                            delete(v3, 42)
                            var v38 uint = 271
                            v39 := GenS6{F0: -1, F1: 1000, F2: (!(fi36 > fi36))}
                            _ = v38
                            _ = v39
                        } else {
                            v14 -= (v14 % ((v14) | 1))
                        }
                        _ = fi36
                        _ = v37
                    }
                }
                v40 := GenS7{F0: 0, F1: 0}
                _ = v40
            } else {
                v14 = v14
                if (v4 != (v4 + v4)) {
                    v14 -= (v14 % ((v14) | 1))
                }
                if false {
                    for fi41 := 0; fi41 < 7; fi41++ {
                        if true {
                            v27 += 0
                            var v42 uint16 = 0
                            _ = v42
                        }
                        v27 = v27
                        _ = fi41
                    }
                }
            }
            v27 = v27
        }
        if (4294967295 < -2147483648) {
            if false {
                fmt.Println(v27)
            }
        }
    } else {
        v43 := GenS7{F0: 142, F1: 127}
        _ = v43
    }
    for fi44 := 0; fi44 < 2; fi44++ {
        for fi45 := 0; fi45 < 2; fi45++ {
            fmt.Println(v27, v2)
            _ = fi45
        }
        var v46 int64 = 9223372036854775807
        _ = fi44
        _ = v46
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
