package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int32 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint16; F1 string; F2 uint8 }
type GenS4 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS5 struct { F0 int64; F1 uint64; F2 uint }

func main() {
    var v1 bool = (false || (!("foo bar" >= ("hello" + ""))))
    v2 := []int{2147483647, 4294967295}
    var v3 bool = (-1 == -2147483648)
    v2 = append(v2, 250)
    switch 10 {
    case 9223372036854775807:
        var v4 float64 = float64(0.0)
        _ = v4
    default:
        if v1 {
            for fi5 := 0; fi5 < 7; fi5++ {
                if v3 {
                    v6 := GenS1{F0: -1, F1: 0, F2: 917}
                    _ = v6
                } else {
                    var v7 uint = 1
                    fmt.Println(v3, v7)
                    _ = v7
                }
                var v8 int64 = 9223372036854775807
                _ = fi5
                _ = v8
            }
        }
    }
    v1 = true
    switch ("a" + ("" + "hello")) {
    case "":
        if (!(42 >= 18446744073709551615)) {
            _ = len(v2)
            var v9 int16 = 0
            _ = v9
        } else {
            if v3 {
                if true {
                    for fi10 := 0; fi10 < 10; fi10++ {
                        var v11 uint16 = 723
                        _ = fi10
                        _ = v11
                    }
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                    var v12 int32 = 1
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 11; fi13++ {
                        for fi14 := 0; fi14 < 18; fi14++ {
                            v3 = (79 >= 3)
                            v1 = false
                            v15 := GenS2{F0: 794, F1: 2}
                            _ = fi14
                            _ = v15
                        }
                        _ = fi13
                    }
                    var v16 int = -1
                    _ = v16
                }
                switch "foo bar" {
                case "\n\t":
                    v2 = append(v2, 1)
                case "hello":
                    switch "\n\t" {
                    case "\n\t":
                        var v17 float64 = (float64(-2.25) * 3.14159)
                        _ = v17
                    case "":
                        if (!(!v3)) {
                            v2 = append(v2, -1)
                        } else {
                            if false {
                                var v18 int64 = 546
                                var v19 uint8 = 128
                                _ = v18
                                _ = v19
                            }
                            v2 = append(v2, 2147483647)
                            if true {
                                v2 = append(v2, 4294967295)
                                v1 = v3
                                for fi20 := 0; fi20 < 10; fi20++ {
                                    for fi21 := 0; fi21 < 6; fi21++ {
                                        v2 = append(v2, (fi21 % (((fi20 % (((fi20 & fi21)) | 1))) | 1)))
                                        var v22 uint = 61
                                        switch "a" {
                                        case "a":
                                            switch 2147483647 {
                                            case 42:
                                                var v23 float64 = ((1e10 - 3.14159) - 3.14159)
                                                _ = v23
                                            }
                                        default:
                                            v24 := GenS1{F0: 23, F1: int8(fi21), F2: 0}
                                            _ = v24
                                        }
                                        _ = fi21
                                        _ = v22
                                    }
                                    if (0.0 >= -2.25) {
                                        switch fi20 {
                                        case -9223372036854775808:
                                            fmt.Println(fi20)
                                        default:
                                            v1 = (2147483647 == 65535)
                                        }
                                        var v25 int16 = 32767
                                        for fi26 := 0; fi26 < 16; fi26++ {
                                            fmt.Println(fi20, fi20, v25)
                                            v25 = (v25 | 32767)
                                            _ = fi26
                                        }
                                        _ = v25
                                    }
                                    _ = fi20
                                }
                            }
                        }
                    }
                default:
                    if len(v2) > 0 {
                        v2[0] = -631
                    }
                }
                for fi27 := 0; fi27 < 7; fi27++ {
                    for fi28 := 0; fi28 < 16; fi28++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi28
                    }
                    for fi29 := 0; fi29 < 14; fi29++ {
                        v1 = v3
                        var v30 uint = 42
                        _ = fi29
                        _ = v30
                    }
                    var v31 int16 = -735
                    _ = fi27
                    _ = v31
                }
            }
            var v32 string = (("foo bar" + "a") + ("" + ""))
            for fi33 := 0; fi33 < 16; fi33++ {
                for fi34 := 0; fi34 < 3; fi34++ {
                    v35 := GenS3{F0: 160, F1: (v32 + ("" + (("\n\t" + v32) + ((v32 + v32) + "")))), F2: 42}
                    _ = len(v2)
                    v3 = ((1 == 1) || ((((v3 && v1) && false) || (610 < 2)) || v1))
                    var v36 float32 = 0.1
                    _ = fi34
                    _ = v35
                    _ = v36
                }
                for fi37 := 0; fi37 < 14; fi37++ {
                    _ = len(v2)
                    switch (((("a" + ("a" + "\n\t")) + "foo bar") + v32) + (v32 + "hello")) {
                    case "foo bar":
                        v38 := GenS4{F0: 2, F1: 4294967295, F2: 127}
                        _ = v38
                    case "hello":
                        var v39 int64 = 950
                        _ = v39
                    case "":
                        v32 = "hello"
                    default:
                        v40 := GenS5{F0: 9223372036854775807, F1: 18446744073709551615, F2: 18446744073709551615}
                        _ = v40
                    }
                    _ = fi37
                }
                _ = fi33
            }
            _ = v32
        }
    }
    v1 = false
    for fi41 := 0; fi41 < 19; fi41++ {
        _ = len(v2)
        _ = fi41
    }
    if len(v2) > 0 {
        v2[0] = -9223372036854775808
    }
    v3 = (2 >= 42)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
