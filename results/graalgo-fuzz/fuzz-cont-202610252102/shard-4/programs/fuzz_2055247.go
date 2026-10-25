package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 int64 }

func main() {
    var v1 int8 = -1
    var v2 uint8 = 15
    var v3 uint64 = 233
    var v4 uint32 = 880
    var v5 int8 = v1
    if true {
        v4 = v4
    } else {
        v6 := []int{814, -150, int(v4), 725}
        var v7 float32 = 0.1
        v8 := GenS1{F0: v3}
        _ = v6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        if (1 >= -185) {
            fmt.Println(v5)
        } else {
            var v10 int8 = v1
            switch "\n\t" {
            case "\n\t":
                v4 = (^(v4))
            case "foo bar":
                v4 = v4
            }
            _ = v10
        }
        v5 = (v5 | v1)
        _ = fi9
    }
    var v11 int8 = (v5 + 100)
    var v12 uint64 = v3
    v13 := GenS2{F0: 1, F1: -1}
    var v14 int64 = 9223372036854775807
    v3 -= 239
    var v15 int8 = v11
    switch "" {
    case "foo bar":
        switch "hello" {
        case "foo bar":
            var v16 uint8 = (v2 / (((v2 / ((v2) | 1))) | 1))
            _ = v16
        case "\n\t":
            v5 -= (v11 & v1)
        }
    case "":
        if (!false) {
            switch 127 {
            case 9223372036854775807:
                v5 += -69
            }
            v17 := GenS3{F0: 22, F1: 32767}
            switch "hello" {
            case "a":
                v2 -= (v2 | (v2 << ((508) & 63)))
            case "\n\t":
                var v18 uint = 633
                _ = v18
            }
            _ = v17
        }
    }
    if (!((v4 - 943) <= (^(v4)))) {
        switch 9223372036854775807 {
        case 1:
            var v19 int16 = -1
            _ = v19
        case -2147483648:
            for fi20 := 0; fi20 < 6; fi20++ {
                v21 := GenS1{F0: v3}
                _ = fi20
                _ = v21
            }
        case -345:
            v5 += v5
        default:
            var v22 int = 2147483647
            _ = v22
        }
    }
    var v23 uint8 = v2
    var v24 int16 = -93
    if true {
        if false {
            var v25 bool = (true || ((!false) || (42 != 241)))
            _ = v25
        } else {
            v26 := []int{0, -802, 32767, 23}
            if true {
                v2 -= uint8(v4)
                v27 := GenS1{F0: (v12 / (((v12 + (^(v12)))) | 1))}
                _ = v27
            } else {
                switch 4294967295 {
                case -2147483648:
                    v26 = append(v26, 4294967295)
                case -276:
                    if true {
                        var v28 bool = false
                        if (!(!(2 <= 255))) {
                            if len(v26) > 0 {
                                v26[0] = 931
                            }
                        } else {
                            v28 = v28
                            if ("a" >= "a") {
                                if (v28 || (v3 <= (v3 / (((v12 & (v3 | 1))) | 1)))) {
                                    if v28 {
                                        switch ("a" + (("" + "") + "")) {
                                        case "\n\t":
                                            fmt.Println(v4, v23, v14)
                                        }
                                    }
                                    var v29 uint64 = v3
                                    v1 = (v15 / ((v1) | 1))
                                    _ = v29
                                } else {
                                    var v30 string = "hello"
                                    v31 := GenS4{F0: (v2 + 157), F1: (v24 ^ -616)}
                                    v24 += -32768
                                    _ = v30
                                    _ = v31
                                }
                            } else {
                                var v32 uint8 = v23
                                _ = v32
                            }
                        }
                        fmt.Println(v23, v3, v3)
                        _ = v28
                    }
                }
            }
            v33 := GenS2{F0: 320, F1: 127}
            _ = v26
            _ = v33
        }
    } else {
        if true {
            v34 := []int{-9223372036854775808, 32767, 32768, int(v14)}
            _ = v34
        }
        switch "hello" {
        case "\n\t":
            v15 = 1
        default:
            if (int32(v24) == -1) {
                v12 -= v3
                switch 127 {
                case 42:
                    var v35 uint16 = 255
                    _ = v35
                case 127:
                    v36 := GenS3{F0: 5, F1: v24}
                    _ = v36
                case -764:
                    v12 = 471
                default:
                    v37 := GenS5{F0: 9223372036854775807}
                    _ = v37
                }
                v38 := GenS2{F0: 790, F1: 10}
                _ = v38
            }
        }
        var v39 int16 = (v24 | v24)
        _ = v39
    }
    fmt.Println(v11, v1, v2)
    var v40 int64 = -912
    v41 := []int{-1}
    var v42 uint32 = 91
    v24 += (v24 % ((v24) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
