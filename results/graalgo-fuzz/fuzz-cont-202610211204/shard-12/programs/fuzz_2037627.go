package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint8 }
type GenS2 struct { F0 string; F1 uint8; F2 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 int32; F1 uint16; F2 uint16 }

func main() {
    var v1 int32 = 0
    var v2 int16 = -433
    v3 := []uint32{580, 2147483648, 42, 1}
    var v4 bool = (413 < 18446744073709551615)
    v5 := []uint32{2}
    var v6 int16 = (v2 | (-(v2)))
    v7 := GenS1{F0: 65535, F1: "hello", F2: 220}
    if len(v3) > 0 {
        v3[0] = 0
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        for fi9 := 0; fi9 < 16; fi9++ {
            v6 = v6
            switch fi9 {
            case -1:
                fmt.Println(v1, v1, v4)
            default:
                if len(v3) > 0 {
                    v3[0] = 0
                }
            }
            _ = fi9
        }
        _ = fi8
    }
    if true {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v1 = (^(v1))
    }
    if (957 == 65535) {
        switch 0 {
        case 32768:
            var v10 float64 = 0.0
            _ = v10
        case 127:
            var v11 float64 = ((((((-2.25 / 0.1) - 3.14159) / 1.5) / (1e10 / -2.25)) * 1.5) - 0.0)
            _ = v11
        case 774:
            if len(v3) > 0 {
                v3[0] = 746
            }
        }
        v1 += (v1 | (v1 << ((2) & 63)))
        var v12 int8 = -5
        _ = v12
    }
    v4 = ("a" < "foo bar")
    switch ("" + "foo bar") {
    case "a":
        if v4 {
            _ = len(v5)
        } else {
            switch ("\n\t" + ("a" + "foo bar")) {
            case "a":
                if len(v3) > 0 {
                    v3[0] = 2
                }
            }
        }
    case "hello":
        if (255 < 0) {
            v4 = v4
            for fi13 := 0; fi13 < 19; fi13++ {
                switch (^(fi13)) {
                case 2147483647:
                    v14 := GenS1{F0: 65535, F1: ("\n\t" + "foo bar"), F2: 127}
                    _ = v14
                case 11:
                    for fi15 := 0; fi15 < 3; fi15++ {
                        switch 32767 {
                        case 1:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v16 string = ""
                            fmt.Println(v2)
                            fmt.Println(fi15)
                            fmt.Println(fi15)
                            _ = v16
                        }
                        _ = fi15
                    }
                case 127:
                    fmt.Println(v4, fi13, fi13)
                default:
                    if ((v1 % (((^(v1))) | 1)) > -2147483648) {
                        switch fi13 {
                        case 0:
                            var v17 string = "\n\t"
                            _ = v17
                        case -313:
                            v2 = -192
                        }
                    }
                }
                v5 = append(v5, 725)
                fmt.Println(v4, v6, v1)
                _ = fi13
            }
            if (!true) {
                if len(v5) > 0 {
                    v5[0] = 1
                }
            } else {
                v18 := GenS2{F0: ("\n\t" + "a"), F1: 2, F2: (v1 * -458)}
                fmt.Println(v2)
                _ = v18
            }
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v19 := GenS3{F0: -77}
            v20 := GenS2{F0: "hello", F1: 226, F2: v1}
            _ = v19
            _ = v20
        }
    case "\n\t":
        if false {
            var v21 uint64 = 18446744073709551615
            v2 = (v2 - v2)
            _ = v21
        } else {
            v22 := GenS4{F0: 55, F1: -9223372036854775808}
            switch ("foo bar" + ("hello" + "hello")) {
            case "hello":
                var v23 float64 = (3.14159 / -2.25)
                _ = v23
            case "":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = v22
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v4 = v4
    _ = len(v5)
    var v24 int16 = (v6 & -133)
    if v4 {
        fmt.Println(v2, v2)
        fmt.Println(v24, v24)
        v25 := GenS5{F0: v1, F1: 255, F2: 65535}
        _ = v25
    }
    var v26 uint8 = 40
    v4 = (0 <= 42)
    if len(v3) > 0 {
        v3[0] = 1
    }
    var v27 uint16 = 65535
    fmt.Println(v2, v26, v26)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
