package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int8; F2 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 int16; F1 uint32; F2 uint16 }

func main() {
    var v1 int8 = -88
    v2 := []string{("foo bar" + (((("a" + "\n\t") + "hello") + ("a" + "foo bar")) + (("foo bar" + "hello") + ""))), "foo bar", "a", ("foo bar" + "\n\t")}
    var v3 float32 = (float32((3.14159 / 0.1)) - 1.5)
    var v4 bool = false
    v5 := make(map[int]uint32)
    var v6 uint64 = 18446744073709551615
    v7 := v5[127]
    var v8 int8 = -1
    if (0 < 255) {
        v4 = (128 > 2147483647)
        var v9 int = 9223372036854775807
        _ = v9
    }
    fmt.Println(v4, v8, v7)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    delete(v5, 128)
    var v10 int16 = 103
    for fi11 := 0; fi11 < 1; fi11++ {
        switch fi11 {
        case 32768:
            var v12 uint = 65535
            _ = v12
        case 0:
            var v13 uint16 = 340
            _ = v13
        case 842:
            delete(v5, fi11)
        default:
            v10 = 1000
        }
        if v4 {
            var v14 uint16 = 42
            for fi15 := 0; fi15 < 9; fi15++ {
                var v16 uint32 = (^(v7))
                switch ("" + "") {
                case "":
                    fmt.Println(v10)
                case "hello":
                    var v17 uint64 = 2
                    _ = v17
                default:
                    var v18 int = (-(fi15))
                    fmt.Println(v1, v1)
                    v1 += 127
                    _ = v18
                }
                v19 := GenS1{F0: ("a" + ("\n\t" + ("" + "\n\t")))}
                if v4 {
                    if (v4 || (v4 || (255 != 255))) {
                        v7 = 65535
                    } else {
                        v16 = 980
                    }
                    delete(v5, 127)
                }
                fmt.Println(v8, v7, v6)
                _ = fi15
                _ = v16
                _ = v19
            }
            _ = v14
        }
        v3 -= (((((3.14159 - v3) * 3.14159) / v3) * float32(1e10)) / 1.0)
        _ = fi11
    }
    _ = len(v5)
    for fi20 := 0; fi20 < 5; fi20++ {
        for fi21 := 0; fi21 < 4; fi21++ {
            v22 := v5[(fi20 - -648)]
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    v2 = append(v2, "foo bar")
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v23 := GenS2{F0: (v7 * v7), F1: int8(v10), F2: (1 == v7)}
    fmt.Println(v4)
    var v24 uint64 = v6
    v25 := GenS2{F0: v7, F1: v8, F2: (!v4)}
    if v4 {
        if true {
            if v4 {
                switch "foo bar" {
                case "foo bar":
                    v26 := GenS3{F0: -9223372036854775808}
                    _ = v26
                }
                v27 := GenS1{F0: "hello"}
                _ = v27
            } else {
                var v28 bool = (v7 < (^(v7)))
                var v29 float32 = v3
                _ = v28
                _ = v29
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v30 uint = 1
        _ = v30
    }
    if (1 != -425) {
        var v31 int = 1
        _ = v31
    } else {
        if (!v4) {
            for fi32 := 0; fi32 < 5; fi32++ {
                var v33 int32 = 2147483647
                var v34 int64 = 42
                _ = fi32
                _ = v33
                _ = v34
            }
            if v4 {
                delete(v5, -92)
            } else {
                v35 := GenS4{F0: ("" <= "hello"), F1: 922}
                for fi36 := 0; fi36 < 1; fi36++ {
                    for fi37 := 0; fi37 < 4; fi37++ {
                        v38 := GenS4{F0: (0 != 42), F1: 188}
                        if (107 == int64(v8)) {
                            fmt.Println(fi37, v1)
                            switch "foo bar" {
                            case "\n\t":
                                fmt.Println(fi36)
                            case "foo bar":
                                v4 = true
                            case "hello":
                                if false {
                                    var v39 int8 = (v8 + v1)
                                    _ = v39
                                }
                            }
                            v40 := GenS5{F0: (v10 / ((v10) | 1)), F1: 4294967295, F2: 487}
                            var v41 uint32 = v7
                            fmt.Println(v3, fi36, fi37)
                            fmt.Println(v7, fi36)
                            _ = v40
                            _ = v41
                        } else {
                            fmt.Println(v1, v1)
                            var v42 uint = 18446744073709551615
                            v43 := GenS2{F0: (v7 | v7), F1: -1, F2: (v4 || (!false))}
                            var v44 uint64 = v6
                            _ = v42
                            _ = v43
                            _ = v44
                        }
                        var v45 int8 = (v8 / (((v1 ^ (v8 & -128))) | 1))
                        _ = fi37
                        _ = v38
                        _ = v45
                    }
                    _ = fi36
                }
                _ = v35
            }
        }
        v10 = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
