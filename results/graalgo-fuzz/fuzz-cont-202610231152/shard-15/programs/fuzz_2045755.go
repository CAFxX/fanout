package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 uint; F1 uint64 }
type GenS6 struct { F0 int; F1 uint64; F2 uint16 }
type GenS7 struct { F0 int8 }
func genFn0(p0 int, p1 int) int16 {
    v4 := []int{}
    switch (p1 * p1) {
    case 131:
        fmt.Println(p1)
    case -1:
        if (!(-32768 < 1000)) {
            v5 := GenS1{F0: 215, F1: 255}
            p1 = (p0 >> ((1) & 63))
            p0 -= 128
            _ = v5
        }
    case 1:
        fmt.Println(p0, p1)
    }
    var v6 bool = (18446744073709551615 <= 42)
    var v7 uint64 = 2
    _ = p0
    _ = p1
    _ = v4
    _ = v6
    _ = v7
    return 459
}

func main() {
    v1 := []string{"hello"}
    v2 := make(map[int]string)
    var v3 bool = false
    if v3 {
        v3 = v3
        v8 := GenS2{F0: 2147483647}
        _ = v8
    } else {
        switch "" {
        case "foo bar":
            v3 = v3
        }
        switch -2147483648 {
        case 128:
            v9 := GenS3{F0: 2, F1: -70}
            _ = v9
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    var v10 float32 = 0.1
    _ = len(v2)
    v11 := GenS4{F0: 42, F1: 0}
    v3 = (571 <= 1)
    v12 := GenS3{F0: 237, F1: -103}
    if false {
        if true {
            _ = len(v1)
            switch 127 {
            case 1:
                v3 = false
            case -9223372036854775808:
                v10 -= v10
            }
            var v13 uint32 = 2147483648
            _ = v13
        } else {
            v3 = (-1 != 9223372036854775807)
        }
        if false {
            if (42 < 540) {
                switch 9223372036854775807 {
                case 32767:
                    var v14 float64 = 0.0
                    _ = v14
                case 128:
                    fmt.Println(v3)
                default:
                    switch 127 {
                    case 4294967295:
                        if (386 != 2) {
                            var v15 float64 = -2.25
                            _ = v15
                        }
                    case -2147483648:
                        for fi16 := 0; fi16 < 13; fi16++ {
                            v3 = ((!(!(0 > 1))) && (0 > 241))
                            v3 = v3
                            _ = fi16
                        }
                    case 9223372036854775807:
                        switch ("" + "a") {
                        case "\n\t":
                            for fi17 := 0; fi17 < 18; fi17++ {
                                if len(v1) > 0 {
                                    v1[0] = ("foo bar" + "\n\t")
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = fi17
                            }
                        case "foo bar":
                            v18 := GenS5{F0: 18446744073709551615, F1: 116}
                            _ = v18
                        default:
                            var v19 uint8 = 158
                            _ = v19
                        }
                    default:
                        if true {
                            for fi20 := 0; fi20 < 16; fi20++ {
                                var v21 string = "foo bar"
                                var v22 int = (fi20 ^ fi20)
                                switch (v22 - fi20) {
                                case -2147483648:
                                    var v23 int16 = 0
                                    _ = v23
                                default:
                                    v3 = false
                                }
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                        } else {
                            var v24 uint32 = 1
                            _ = v24
                        }
                    }
                }
            }
            var v25 uint32 = 524
            if v3 {
                v3 = ((1 == -2147483648) && true)
                switch "hello" {
                case "hello":
                    v26 := GenS6{F0: 4294967295, F1: 1, F2: 614}
                    _ = v26
                default:
                    switch "foo bar" {
                    case "foo bar":
                        if len(v1) > 0 {
                            v1[0] = "hello"
                        }
                    case "a":
                        v3 = (((("foo bar" + ("hello" + ((("" + "foo bar") + "foo bar") + ""))) + ("hello" + (("" + "\n\t") + "\n\t"))) < "hello") && (("" + "hello") != "foo bar"))
                    case "hello":
                        _ = len(v1)
                    }
                }
                for fi27 := 0; fi27 < 14; fi27++ {
                    for fi28 := 0; fi28 < 4; fi28++ {
                        fmt.Println(v25, v3)
                        switch ("hello" + "a") {
                        case "a":
                            var v29 int8 = 0
                            _ = v29
                        case "":
                            v25 += v25
                        case "hello":
                            if (!false) {
                                v25 += 0
                            } else {
                                v30 := GenS5{F0: uint(v25), F1: 18446744073709551615}
                                _ = v30
                            }
                        default:
                            var v31 float32 = v10
                            _ = v31
                        }
                        v25 = 645
                        _ = fi28
                    }
                    var v32 bool = ((0.1 * 1.5) < 1e10)
                    _ = fi27
                    _ = v32
                }
            }
            _ = v25
        }
        if (((1e10 / (3.14159 + -2.25)) >= 0.0) && (0 != -623)) {
            v10 = v10
            v33 := GenS6{F0: 835, F1: 4294967295, F2: 42}
            _ = v33
        } else {
            delete(v2, 122)
            v34 := GenS2{F0: -513}
            _ = v34
        }
    }
    fmt.Println(v10)
    v3 = false
    v35 := GenS7{F0: 127}
    var v36 int64 = 42
    delete(v2, 386)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v35
    _ = v36
}
