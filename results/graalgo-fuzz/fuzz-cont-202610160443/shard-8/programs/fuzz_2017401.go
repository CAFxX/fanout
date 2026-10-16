package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int; F1 bool; F2 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 int32; F2 bool }

func main() {
    var v1 int = 9223372036854775807
    v2 := []uint32{1, 4294967295}
    var v3 string = ""
    var v4 uint64 = 448
    if (42 != 2) {
        v3 = "hello"
        v5 := GenS1{F0: -424, F1: -28}
        _ = v5
    }
    if (2 < 1) {
        if true {
            fmt.Println(v3, v3, v4)
            fmt.Println(v3)
            if true {
                for fi6 := 0; fi6 < 14; fi6++ {
                    switch v3 {
                    case "\n\t":
                        fmt.Println(v4, fi6)
                    case "foo bar":
                        for fi7 := 0; fi7 < 13; fi7++ {
                            v8 := GenS2{F0: v1, F1: (false && true), F2: (v4 + v4)}
                            v1 += 612
                            _ = fi7
                            _ = v8
                        }
                    }
                    _ = fi6
                }
                if ((^(v4)) == 4294967295) {
                    v3 = "a"
                    var v9 uint = 18446744073709551615
                    var v10 uint8 = 2
                    _ = v10
                    _ = v9
                } else {
                    var v11 bool = ((-1 >= 32767) && (255 >= 32767))
                    _ = v11
                }
            }
        }
        if (66 >= 4294967295) {
            v4 -= v4
            v4 -= v4
            v4 += (v4 & 4294967295)
        } else {
            if false {
                fmt.Println(v1, v4, v4)
                v12 := GenS3{F0: v3}
                v2 = append(v2, 2)
                _ = v12
            } else {
                v13 := GenS3{F0: v3}
                switch (("foo bar" + (v3 + v3)) + ("a" + (v3 + v3))) {
                case "\n\t":
                    if true {
                        _ = len(v2)
                        var v14 int32 = -2147483648
                        for fi15 := 0; fi15 < 3; fi15++ {
                            v4 += (v4 >> ((150) & 63))
                            v16 := GenS4{F0: 0, F1: (v14 & v14), F2: true}
                            _ = fi15
                            _ = v16
                        }
                        _ = v14
                    }
                case "foo bar":
                    if false {
                        switch "" {
                        case "hello":
                            var v17 int32 = int32(v4)
                            _ = v17
                        case "foo bar":
                            v18 := GenS3{F0: (v3 + ("hello" + v3))}
                            _ = v18
                        case "":
                            v1 -= (v1 - 2147483647)
                        }
                        switch (("" + (v3 + v3)) + v3) {
                        case "a":
                            switch v3 {
                            case "":
                                if false {
                                    v4 = v4
                                }
                            case "\n\t":
                                if true {
                                    var v19 int64 = -9223372036854775808
                                    for fi20 := 0; fi20 < 15; fi20++ {
                                        var v21 uint = 302
                                        var v22 int = fi20
                                        _ = fi20
                                        _ = v21
                                        _ = v22
                                    }
                                    fmt.Println(v3)
                                    _ = v19
                                } else {
                                    if true {
                                        var v23 bool = ((v1 + (v1 << ((2) & 63))) <= 2147483647)
                                        var v24 bool = (506 != 0)
                                        if len(v2) > 0 {
                                            v2[0] = 558
                                        }
                                        _ = v23
                                        _ = v24
                                    }
                                    v25 := GenS3{F0: v3}
                                    var v26 uint32 = 765
                                    _ = v25
                                    _ = v26
                                }
                            case "hello":
                                v3 = "foo bar"
                            }
                        case "foo bar":
                            var v27 float64 = 0.1
                            _ = v27
                        }
                    }
                default:
                    var v28 bool = true
                    _ = v28
                }
                _ = v13
            }
        }
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        v3 = ("hello" + (v3 + v3))
        _ = fi29
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (-71 == 42) {
        if len(v2) > 0 {
            v2[0] = 1
        }
        if false {
            v30 := GenS3{F0: (v3 + v3)}
            _ = v30
        }
        v4 = (^(v4))
    }
    if false {
        var v31 int8 = 127
        var v32 float32 = 1.5
        _ = v31
        _ = v32
    } else {
        var v33 uint16 = 166
        _ = v33
    }
    if false {
        _ = len(v2)
        switch v1 {
        case 888:
            v4 += 18446744073709551615
        }
    } else {
        v3 = (("\n\t" + "foo bar") + v3)
    }
    var v34 uint8 = uint8(v4)
    var v35 int32 = -38
    for fi36 := 0; fi36 < 11; fi36++ {
        v3 = ""
        for fi37 := 0; fi37 < 4; fi37++ {
            var v38 int = 9223372036854775807
            v38 = (v1 / (((v38 >> ((836) & 63))) | 1))
            _ = fi37
            _ = v38
        }
        for fi39 := 0; fi39 < 1; fi39++ {
            fmt.Println(v35, v4, v34)
            v1 -= fi36
            _ = fi39
        }
        _ = fi36
    }
    fmt.Println(v1, v4)
    var v40 float64 = float64(-2.25)
    if len(v2) > 0 {
        v2[0] = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v40
}
