package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 int8; F2 uint }
type GenS4 struct { F0 int; F1 int8 }
func genFn0(p0 uint16) bool {
    if ((!true) && false) {
        v5 := GenS1{F0: 1, F1: 2, F2: 1}
        var v6 int8 = -1
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 18446744073709551615, F1: 255, F2: 42}
    var v8 uint = 65535
    var v9 float32 = (-2.25 / 3.14159)
    v10 := []int{-188, -1, -133}
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return false
}

func main() {
    v1 := []int{}
    var v2 uint32 = 2
    var v3 float32 = (1e10 + (0.1 / -2.25))
    var v4 uint64 = 0
    v1 = append(v1, 42)
    fmt.Println(v2, v4)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 = ((1e10 - v3) * v3)
    v11 := GenS1{F0: 0, F1: 2, F2: 216}
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        v2 += 42
        var v12 int32 = 1
        _ = v12
    }
    var v13 float32 = v3
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        v4 += v4
        v1 = append(v1, -1)
        v4 = v4
        _ = fi14
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 bool = true
        var v17 uint8 = 127
        _ = fi15
        _ = v16
        _ = v17
    }
    v2 = 0
    if (((false && (!true)) && false) && ((!(!(!((2 <= 95) && (-75 > -128))))) && (!(!false)))) {
        switch "hello" {
        case "hello":
            if (0 <= 0) {
                var v18 uint8 = 178
                for fi19 := 0; fi19 < 9; fi19++ {
                    if true {
                        if false {
                            v20 := GenS2{F0: -2147483648}
                            if (!true) {
                                v21 := GenS3{F0: uint16(v4), F1: 42, F2: 694}
                                if (222 >= 2) {
                                    for fi22 := 0; fi22 < 6; fi22++ {
                                        fmt.Println(v2, fi22, v3)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        var v23 uint = 42
                                        _ = fi22
                                        _ = v23
                                    }
                                    var v24 float32 = 0.1
                                    _ = len(v1)
                                    _ = v24
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                _ = v21
                            }
                            _ = v20
                        } else {
                            var v25 string = "foo bar"
                            _ = v25
                        }
                    }
                    var v26 int32 = 42
                    _ = fi19
                    _ = v26
                }
                if len(v1) > 0 {
                    v1[0] = -9223372036854775808
                }
                _ = v18
            } else {
                if (!false) {
                    var v27 uint32 = 997
                    _ = v27
                }
                fmt.Println(v2, v3, v13)
                switch "" {
                case "hello":
                    for fi28 := 0; fi28 < 12; fi28++ {
                        if len(v1) > 0 {
                            v1[0] = (fi28 / ((-333) | 1))
                        }
                        v13 += v3
                        _ = fi28
                    }
                case "":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
        case "foo bar":
            var v29 int64 = -1
            _ = v29
        default:
            fmt.Println(v3, v2)
        }
    } else {
        _ = len(v1)
        var v30 uint64 = (v4 - v4)
        for fi31 := 0; fi31 < 5; fi31++ {
            fmt.Println(fi31, v13)
            _ = fi31
        }
        _ = v30
    }
    for fi32 := 0; fi32 < 10; fi32++ {
        v33 := GenS4{F0: (fi32 * (fi32 ^ fi32)), F1: -48}
        v2 = 2147483648
        _ = fi32
        _ = v33
    }
    v3 = v13
    switch "\n\t" {
    case "hello":
        for fi34 := 0; fi34 < 11; fi34++ {
            var v35 string = "hello"
            _ = fi34
            _ = v35
        }
    case "\n\t":
        var v36 uint = 65535
        _ = v36
    }
    v3 -= v13
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
