package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int16 }
type GenS2 struct { F0 int16; F1 string; F2 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 uint32; F1 bool }

func main() {
    v1 := []int{4294967295, 885}
    var v2 uint32 = 2
    var v3 int8 = -1
    var v4 int = 32767
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 string = "hello"
        var v7 bool = true
        fmt.Println(v2)
        _ = fi5
        _ = v6
        _ = v7
    }
    switch "hello" {
    case "\n\t":
        v8 := GenS1{F0: 42, F1: ((^(v2)) <= (v2 / (((^(v2))) | 1))), F2: 1}
        _ = v8
    case "hello":
        if false {
            var v9 float64 = 1e10
            _ = v9
        }
    case "":
        switch "hello" {
        case "hello":
            fmt.Println(v3)
        case "foo bar":
            v3 = (v3 - v3)
        case "":
            var v10 bool = (432 >= 255)
            _ = v10
        }
    default:
        if true {
            if len(v1) > 0 {
                v1[0] = v4
            }
            v11 := GenS2{F0: 32767, F1: "\n\t", F2: 255}
            fmt.Println(v4, v4)
            _ = v11
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v12 := GenS3{F0: -1}
    if len(v1) > 0 {
        v1[0] = (-(v4))
    }
    v2 = (v2 % ((1) | 1))
    fmt.Println(v2, v2, v2)
    if (!true) {
        for fi13 := 0; fi13 < 4; fi13++ {
            v4 += (fi13 + fi13)
            for fi14 := 0; fi14 < 5; fi14++ {
                v4 = fi13
                var v15 uint = 255
                var v16 int64 = 42
                _ = fi14
                _ = v15
                _ = v16
            }
            if false {
                v3 = v3
                switch ((("foo bar" + "foo bar") + ("" + "\n\t")) + "hello") {
                case "foo bar":
                    v1 = append(v1, (v4 + fi13))
                default:
                    var v17 bool = (42 != (v2 % ((v2) | 1)))
                    _ = v17
                }
                var v18 float32 = float32(1.5)
                _ = v18
            }
            _ = fi13
        }
    } else {
        switch "\n\t" {
        case "a":
            for fi19 := 0; fi19 < 7; fi19++ {
                v3 -= (v3 % ((v3) | 1))
                for fi20 := 0; fi20 < 4; fi20++ {
                    fmt.Println(v4, v4)
                    var v21 uint32 = v2
                    _ = fi20
                    _ = v21
                }
                _ = fi19
            }
        case "":
            switch v4 {
            case -700:
                _ = len(v1)
            case 800:
                v22 := GenS2{F0: int16(v2), F1: "hello", F2: 255}
                _ = v22
            case -1:
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    for fi23 := 0; fi23 < 16; fi23++ {
                        if len(v1) > 0 {
                            v1[0] = (fi23 / ((fi23) | 1))
                        }
                        if false {
                            if false {
                                var v24 uint = 0
                                v25 := GenS3{F0: 323}
                                v2 = 2147483647
                                _ = v24
                                _ = v25
                            } else {
                                v2 += v2
                                var v26 float32 = (1.5 + (0.1 + 0.1))
                                _ = v26
                            }
                            for fi27 := 0; fi27 < 6; fi27++ {
                                if len(v1) > 0 {
                                    v1[0] = fi23
                                }
                                fmt.Println(fi23)
                                _ = fi27
                            }
                        } else {
                            var v28 int8 = v3
                            v29 := GenS4{F0: 891, F1: fi23}
                            _ = v28
                            _ = v29
                        }
                        _ = fi23
                    }
                }
            default:
                v2 += 42
            }
        default:
            var v30 uint8 = 217
            _ = v30
        }
        var v31 string = "a"
        _ = v31
    }
    v32 := GenS3{F0: -76}
    for fi33 := 0; fi33 < 9; fi33++ {
        v3 += 1
        _ = fi33
    }
    fmt.Println(v2, v3, v2)
    for fi34 := 0; fi34 < 15; fi34++ {
        var v35 uint = 2
        switch v4 {
        case -9223372036854775808:
            if true {
                for fi36 := 0; fi36 < 5; fi36++ {
                    v37 := GenS3{F0: 1}
                    fmt.Println(v35, v2)
                    _ = fi36
                    _ = v37
                }
                _ = len(v1)
                for fi38 := 0; fi38 < 14; fi38++ {
                    _ = len(v1)
                    v39 := GenS5{F0: -1, F1: 42}
                    fmt.Println(v3, v2, fi34)
                    v4 = (fi38 ^ 540)
                    _ = fi38
                    _ = v39
                }
            }
        default:
            _ = len(v1)
        }
        v40 := GenS6{F0: (v2 << ((v35) & 63)), F1: (v2 == v2)}
        _ = fi34
        _ = v35
        _ = v40
    }
    var v41 bool = (-2147483648 < 2147483647)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v41
}
