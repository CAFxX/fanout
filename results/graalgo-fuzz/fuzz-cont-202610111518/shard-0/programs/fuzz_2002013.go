package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint }
type GenS2 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64; F1 int8; F2 int8 }
type GenS6 struct { F0 int64 }

func main() {
    var v1 float64 = (3.14159 / 1.0)
    v2 := make(map[string]string)
    var v3 int64 = 9223372036854775807
    var v4 bool = false
    var v5 int64 = -202
    v1 = v1
    v4 = v4
    var v6 uint = 65535
    var v7 int = 32767
    v3 = (v3 << (((^(v6))) & 63))
    if ((^(v7)) != v7) {
        _ = len(v2)
        for fi8 := 0; fi8 < 17; fi8++ {
            _ = len(v2)
            v1 += (-2.25 - v1)
            v9 := GenS1{F0: 0, F1: 2, F2: v6}
            _ = fi8
            _ = v9
        }
    } else {
        v10 := GenS2{F0: 42, F1: -104, F2: 192}
        v6 += (^(v6))
        _ = v10
    }
    fmt.Println(v4, v4, v1)
    var v11 uint8 = 128
    switch v7 {
    case 628:
        v12 := GenS1{F0: 0, F1: 42, F2: (v6 | v6)}
        _ = v12
    }
    v6 = (v6 % ((v6) | 1))
    switch ("foo bar" + "a") {
    case "foo bar":
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 string = ((((("foo bar" + "") + "a") + "\n\t") + "hello") + "a")
            _ = fi13
            _ = v14
        }
    case "\n\t":
        switch "foo bar" {
        case "\n\t":
            switch "a" {
            case "a":
                var v15 int = (^(v7))
                _ = v15
            case "\n\t":
                v7 = v7
            case "foo bar":
                switch "\n\t" {
                case "foo bar":
                    v5 = v3
                case "\n\t":
                    v16 := v2[""]
                    _ = v16
                case "hello":
                    var v17 string = "\n\t"
                    _ = v17
                default:
                    v18 := GenS2{F0: 1, F1: -128, F2: v11}
                    _ = v18
                }
            default:
                fmt.Println(v3, v3)
            }
        default:
            if (3.14159 <= 0.0) {
                v19 := []int{32767, v7, -1}
                _ = v19
            } else {
                v20 := GenS3{F0: 1, F1: v6}
                var v21 int64 = v5
                if ("\n\t" >= "hello") {
                    v22 := v2["\n\t"]
                    if (967 > 2147483647) {
                        v23 := GenS2{F0: 2147483648, F1: 0, F2: v11}
                        var v24 int8 = 42
                        _ = v23
                        _ = v24
                    } else {
                        fmt.Println(v11, v6)
                        for fi25 := 0; fi25 < 11; fi25++ {
                            switch "" {
                            case "\n\t":
                                if (2147483647 == uint32(v11)) {
                                    var v26 uint = (v6 / ((v6) | 1))
                                    v27 := GenS2{F0: 10, F1: 42, F2: v11}
                                    v1 += -2.25
                                    _ = v26
                                    _ = v27
                                } else {
                                    var v28 string = v22
                                    var v29 uint16 = 65535
                                    _ = v28
                                    _ = v29
                                }
                            case "hello":
                                v3 -= (v21 >> (((v6 & v6)) & 63))
                            case "a":
                                switch v22 {
                                case "":
                                    v4 = (v4 || v4)
                                case "\n\t":
                                    fmt.Println(v4, v21)
                                default:
                                    v22 = (v22 + "hello")
                                }
                            default:
                                fmt.Println(v3, v4, v4)
                            }
                            _ = fi25
                        }
                    }
                    for fi30 := 0; fi30 < 12; fi30++ {
                        for fi31 := 0; fi31 < 17; fi31++ {
                            v11 = 37
                            var v32 int64 = v3
                            _ = fi31
                            _ = v32
                        }
                        v33 := GenS4{F0: true}
                        for fi34 := 0; fi34 < 10; fi34++ {
                            v3 += v5
                            v1 = v1
                            if (-2.25 < (3.14159 - 1.5)) {
                                if ((v5 - (v5 - v21)) == -531) {
                                    var v35 int8 = 0
                                    switch v22 {
                                    case "a":
                                        v21 -= v5
                                    case "foo bar":
                                        var v36 uint8 = (^(v11))
                                        _ = v36
                                    default:
                                        v37 := GenS5{F0: 0, F1: (v35 >> ((v6) & 63)), F2: v35}
                                        _ = v37
                                    }
                                    fmt.Println(fi34)
                                    _ = v35
                                }
                                fmt.Println(v6, v22)
                            }
                            if true {
                                if (42 >= 1) {
                                    switch "a" {
                                    case "":
                                        v38 := v2[v22]
                                        _ = v38
                                    case "foo bar":
                                        fmt.Println(v22, v3, v6)
                                    case "a":
                                        v11 -= 92
                                    default:
                                        v7 -= fi34
                                    }
                                    v22 = v22
                                } else {
                                    if (!v4) {
                                        v22 = v22
                                    }
                                    fmt.Println(v11, v21, v5)
                                }
                                fmt.Println(fi34)
                            }
                            v39 := GenS6{F0: v5}
                            _ = fi34
                            _ = v39
                        }
                        _ = fi30
                        _ = v33
                    }
                    _ = v22
                } else {
                    var v40 int32 = 1
                    _ = v40
                }
                _ = v20
                _ = v21
            }
        }
    case "a":
        var v41 bool = (0.1 <= (0.1 - 3.14159))
        _ = v41
    default:
        var v42 int8 = -1
        _ = v42
    }
    if v4 {
        v43 := []int{v7}
        _ = len(v2)
        v11 -= (v11 / ((127) | 1))
        _ = v43
    } else {
        if v4 {
            switch "foo bar" {
            case "":
                var v44 uint32 = 2
                _ = v44
            case "hello":
                if false {
                    fmt.Println(v4, v5, v5)
                } else {
                    var v45 uint16 = 65535
                    var v46 int64 = (v5 | (v3 >> ((v6) & 63)))
                    _ = v45
                    _ = v46
                }
            case "\n\t":
                v2["hello"] = ""
            default:
                v47 := GenS5{F0: 1, F1: int8(v5), F2: 100}
                _ = v47
            }
        } else {
            v48 := GenS6{F0: v3}
            _ = v48
        }
    }
    var v49 uint16 = 42
    v50 := GenS3{F0: 4294967295, F1: 2}
    var v51 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v6
    _ = v7
}
