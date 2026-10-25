package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 uint64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32; F1 uint8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int; F1 bool; F2 bool }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int; F1 int32; F2 string }
type GenS11 struct { F0 uint64; F1 int16 }
func genFn0(p0 int16) bool {
    var v4 uint8 = 1
    fmt.Println(p0, v4)
    if (-1 != 51) {
        v5 := GenS1{F0: 42, F1: (v4 + v4)}
        v4 += 42
        _ = v5
    } else {
        fmt.Println(v4, v4)
    }
    _ = p0
    _ = v4
    return (-1 > -703)
}

func main() {
    var v1 uint64 = 650
    var v2 int64 = 42
    v3 := []int{128}
    v6 := GenS2{F0: 128}
    v7 := GenS2{F0: 127}
    if (889 != 42) {
        v2 = v2
        var v8 uint64 = (^(v1))
        switch "" {
        case "foo bar":
            switch 575 {
            case 4294967295:
                if len(v3) > 0 {
                    v3[0] = -1
                }
            case -401:
                fmt.Println(v8)
            case -9223372036854775808:
                if (0.0 <= float32(float64(0.0))) {
                    fmt.Println(v8)
                    v9 := GenS2{F0: -2147483648}
                    _ = v9
                } else {
                    fmt.Println(v8, v8)
                    switch ("hello" + "") {
                    case "a":
                        if ((!(119 <= 598)) || ((v2 ^ (v2 | (v2 >> ((2) & 63)))) <= v2)) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if false {
                                var v10 uint = 2
                                v8 = 1
                                _ = v10
                            }
                            if true {
                                v2 = v2
                                if true {
                                    v1 += v8
                                    var v11 uint8 = 2
                                    _ = v11
                                }
                            } else {
                                if ((-2.25 + 0.1) >= 3.14159) {
                                    if false {
                                        var v12 uint64 = (v8 | 42)
                                        _ = v12
                                    }
                                    v2 += (v2 << ((0) & 63))
                                    if false {
                                        v13 := GenS3{F0: 9223372036854775807}
                                        v14 := GenS3{F0: 0}
                                        _ = v13
                                        _ = v14
                                    }
                                }
                                for fi15 := 0; fi15 < 17; fi15++ {
                                    for fi16 := 0; fi16 < 9; fi16++ {
                                        if false {
                                            v17 := GenS4{F0: 0, F1: (^(v8))}
                                            v2 = 336
                                            if true {
                                                _ = len(v3)
                                            }
                                            _ = v17
                                        }
                                        var v18 string = "foo bar"
                                        v1 += (^(v8))
                                        var v19 int64 = v2
                                        fmt.Println(v19)
                                        v20 := GenS5{F0: 60}
                                        v2 += 42
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = fi16
                                        _ = v18
                                        _ = v19
                                        _ = v20
                                    }
                                    _ = fi15
                                }
                            }
                        }
                    case "\n\t":
                        v2 = v2
                    }
                }
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        case "\n\t":
            v21 := GenS6{F0: 4294967295, F1: 255}
            _ = v21
        default:
            switch "" {
            case "\n\t":
                fmt.Println(v2)
            case "hello":
                for fi22 := 0; fi22 < 12; fi22++ {
                    v2 = v2
                    _ = fi22
                }
            default:
                fmt.Println(v8, v2)
            }
        }
        _ = v8
    }
    _ = len(v3)
    if (706 <= 0) {
        var v23 uint32 = 2
        v3 = append(v3, 179)
        _ = v23
    } else {
        var v24 int = 9223372036854775807
        fmt.Println(v24, v1)
        _ = v24
    }
    fmt.Println(v2, v1)
    for fi25 := 0; fi25 < 12; fi25++ {
        switch "hello" {
        case "foo bar":
            v1 += 194
        case "\n\t":
            if (false || (!(!(1 <= 255)))) {
                v26 := GenS6{F0: uint32(v1), F1: 2}
                _ = v26
            } else {
                var v27 int32 = 2147483647
                _ = v27
            }
        }
        _ = fi25
    }
    if true {
        for fi28 := 0; fi28 < 16; fi28++ {
            if true {
                v2 += 908
            } else {
                v1 -= (v1 - v1)
            }
            v29 := GenS7{F0: v2}
            v2 = v2
            _ = fi28
            _ = v29
        }
        fmt.Println(v1)
    } else {
        var v30 string = ("hello" + "")
        fmt.Println(v30)
        var v31 int32 = -2147483648
        _ = v30
        _ = v31
    }
    switch ((("" + "a") + ("hello" + "\n\t")) + ("\n\t" + ("foo bar" + ""))) {
    case "\n\t":
        var v32 float64 = -2.25
        _ = v32
    default:
        v33 := GenS8{F0: 302, F1: (false && false), F2: true}
        _ = v33
    }
    v2 -= (v2 | (v2 & (v2 & 42)))
    v2 = (-(v2))
    switch 40 {
    case -9223372036854775808:
        v34 := GenS2{F0: 32767}
        _ = v34
    case -1000:
        var v35 int = 32768
        _ = v35
    case 9223372036854775807:
        v2 += 9223372036854775807
    default:
        fmt.Println(v2)
    }
    v36 := GenS9{F0: 2}
    v37 := GenS10{F0: -1, F1: -302, F2: "\n\t"}
    if true {
        for fi38 := 0; fi38 < 10; fi38++ {
            switch "\n\t" {
            case "foo bar":
                v39 := GenS11{F0: 0, F1: -155}
                _ = v39
            case "hello":
                switch (fi38 << ((19) & 63)) {
                case 3:
                    var v40 int8 = 14
                    _ = v40
                default:
                    fmt.Println(v2)
                }
            }
            fmt.Println(fi38)
            switch (fi38 - fi38) {
            case -353:
                _ = len(v3)
            case 128:
                var v41 uint = 65535
                _ = v41
            case 32767:
                var v42 bool = (true && true)
                _ = v42
            default:
                v1 += (^(v1))
            }
            _ = fi38
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v6
    _ = v7
}
