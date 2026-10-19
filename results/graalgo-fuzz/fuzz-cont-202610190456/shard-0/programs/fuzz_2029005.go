package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS8 struct { F0 int16; F1 int16 }
func genFn0(p0 int64) int8 {
    v5 := GenS1{F0: -80, F1: 0, F2: 337}
    p0 -= -9223372036854775808
    _ = p0
    _ = v5
    return 127
}
func genFn1(p0 uint, p1 uint64, p2 int8) string {
    var v6 int = 985
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := []int{(fi7 & fi7)}
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        if false {
            v9 := GenS2{F0: 630}
            _ = len(v8)
            _ = v9
        } else {
            var v10 uint64 = (p1 * p1)
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    switch "hello" {
    case "a":
        fmt.Println(v6, p2)
    case "foo bar":
        var v11 string = ("hello" + "a")
        _ = v11
    case "":
        v6 = (v6 >> (((p0 * p0)) & 63))
    }
    if (!(255 < 0)) {
        var v12 float64 = (((((0.0 + (-2.25 * (0.1 / 1e10))) - -2.25) * (-2.25 - ((0.0 * ((0.1 / (1e10 / float64(-2.25))) + 0.1)) * float64(float32(0.0))))) / -2.25) - 0.0)
        p0 = p0
        _ = v12
    }
    var v13 int64 = -9223372036854775808
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v6
    return ("foo bar" + "foo bar")
}

func main() {
    var v1 int32 = 1
    var v2 uint16 = 65535
    v3 := make(map[uint32]bool)
    v4 := make(map[int]bool)
    v1 = (v1 + (v1 - -810))
    if (0 >= 135) {
        v1 = 276
        if true {
            var v14 int16 = 1000
            var v15 uint8 = 166
            v16 := GenS3{F0: (v2 ^ (^(v2)))}
            _ = v14
            _ = v15
            _ = v16
        }
        switch "\n\t" {
        case "foo bar":
            if (true && (42 == 42)) {
                switch "foo bar" {
                case "hello":
                    switch "a" {
                    case "foo bar":
                        switch "a" {
                        case "":
                            v17 := GenS1{F0: -1, F1: -83, F2: 32768}
                            _ = v17
                        case "\n\t":
                            fmt.Println(v1, v2)
                        case "a":
                            v1 = (^(v1))
                        }
                    case "hello":
                        var v18 uint = 42
                        _ = v18
                    case "a":
                        var v19 int64 = 1
                        _ = v19
                    }
                case "":
                    if false {
                        if ((^(v1)) >= v1) {
                            switch 803 {
                            case 127:
                                v20 := GenS4{F0: 2}
                                _ = v20
                            }
                        } else {
                            for fi21 := 0; fi21 < 1; fi21++ {
                                var v22 int8 = -128
                                var v23 uint8 = 95
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                        }
                        _ = len(v4)
                        v24 := GenS5{F0: 100, F1: 32767}
                        _ = v24
                    } else {
                        for fi25 := 0; fi25 < 4; fi25++ {
                            var v26 int = fi25
                            _ = fi25
                            _ = v26
                        }
                    }
                }
                switch 32767 {
                case 9223372036854775807:
                    if false {
                        var v27 uint64 = 18446744073709551615
                        if true {
                            v28 := GenS5{F0: 0, F1: (v2 >> ((757) & 63))}
                            var v29 string = ""
                            _ = v28
                            _ = v29
                        }
                        switch 0 {
                        case 1:
                            var v30 uint64 = 597
                            _ = v30
                        }
                        _ = v27
                    } else {
                        v1 += (v1 - 2147483647)
                        var v31 uint16 = (v2 + v2)
                        _ = v31
                    }
                case -341:
                    if true {
                        v32 := GenS1{F0: 1, F1: 100, F2: 0}
                        _ = v32
                    } else {
                        var v33 int16 = 294
                        fmt.Println(v2, v33)
                        _ = v33
                    }
                }
            } else {
                fmt.Println(v2, v1)
                if false {
                    if true {
                        v34 := []int{-2147483648, 2147483647, 1}
                        v35 := GenS6{F0: 42}
                        for fi36 := 0; fi36 < 18; fi36++ {
                            v37 := GenS7{F0: (^(v2)), F1: 131, F2: -1}
                            _ = fi36
                            _ = v37
                        }
                        _ = v34
                        _ = v35
                    } else {
                        v1 -= (v1 - (v1 | 42))
                        var v38 uint16 = 1
                        _ = v38
                    }
                    v1 += v1
                }
            }
        default:
            var v39 float32 = 0.1
            _ = v39
        }
    } else {
        fmt.Println(v1)
    }
    v40 := GenS8{F0: 32767, F1: 1000}
    var v41 int = 9223372036854775807
    var v42 uint32 = 944
    var v43 bool = true
    v44 := GenS7{F0: 251, F1: v2, F2: -823}
    v45 := GenS4{F0: 1}
    v2 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
