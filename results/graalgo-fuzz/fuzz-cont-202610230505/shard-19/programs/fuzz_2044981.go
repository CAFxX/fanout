package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 bool; F2 string }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 uint32; F1 uint16; F2 int }

func main() {
    var v1 int16 = 32767
    v2 := []uint32{2147483647, 555}
    v3 := make(map[uint32]string)
    var v4 uint64 = 42
    fmt.Println(v4)
    var v5 int = -2147483648
    var v6 uint16 = 2
    switch "hello" {
    case "":
        var v7 int = 128
        _ = v7
    case "a":
        var v8 int8 = 127
        _ = v8
    default:
        fmt.Println(v4, v1, v6)
    }
    v2 = append(v2, 1)
    if false {
        switch ("hello" + "foo bar") {
        case "hello":
            var v9 int = -1
            _ = v9
        case "\n\t":
            switch v5 {
            case 695:
                v1 += (v1 % (((v1 | (v1 % ((v1) | 1)))) | 1))
            case -1:
                for fi10 := 0; fi10 < 7; fi10++ {
                    switch 102 {
                    case 32768:
                        v1 = v1
                    case -929:
                        v6 = v6
                    default:
                        switch fi10 {
                        case 522:
                            switch 9223372036854775807 {
                            case 0:
                                _ = len(v3)
                            default:
                                var v11 int8 = -128
                                _ = v11
                            }
                        case 9223372036854775807:
                            var v12 float64 = ((float64((1e10 / 1e10)) + (1.5 / ((3.14159 + (((0.1 + 1.5) / 3.14159) + -2.25)) - 0.0))) - 0.1)
                            _ = v12
                        case 841:
                            v1 -= -32768
                        default:
                            v1 = 42
                        }
                    }
                    v1 = v1
                    _ = fi10
                }
            }
        case "":
            switch int(v6) {
            case 9223372036854775807:
                var v13 int = v5
                _ = v13
            case 127:
                fmt.Println(v4)
            default:
                fmt.Println(v1, v1, v5)
            }
        }
        fmt.Println(v4, v5, v1)
    } else {
        var v14 bool = ("\n\t" <= "a")
        if false {
            v14 = (v14 || false)
            fmt.Println(v1, v6)
            var v15 uint16 = v6
            _ = v15
        } else {
            _ = len(v3)
            delete(v3, 4294967295)
        }
        switch "hello" {
        case "hello":
            for fi16 := 0; fi16 < 11; fi16++ {
                if (!v14) {
                    v17 := GenS1{F0: v4, F1: 4294967295, F2: 42}
                    v18 := GenS1{F0: (v4 - v4), F1: 688, F2: (v4 | 1)}
                    _ = v17
                    _ = v18
                }
                _ = fi16
            }
        case "a":
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
        case "":
            delete(v3, 2147483647)
        default:
            var v19 int16 = v1
            _ = v19
        }
        _ = v14
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        v21 := GenS2{F0: 42}
        var v22 uint16 = 1
        switch (v5 * -2147483648) {
        case 1:
            for fi23 := 0; fi23 < 5; fi23++ {
                if len(v2) > 0 {
                    v2[0] = 2
                }
                v24 := GenS3{F0: -551, F1: (4294967295 == 4294967295), F2: ("a" + "hello")}
                v25 := GenS2{F0: 4294967295}
                _ = fi23
                _ = v24
                _ = v25
            }
        case 421:
            switch fi20 {
            case -2147483648:
                v26 := GenS1{F0: (^(v4)), F1: 284, F2: (^(v4))}
                _ = v26
            case -1:
                var v27 bool = ((2 > 223) && true)
                _ = v27
            default:
                switch ("foo bar" + "foo bar") {
                case "foo bar":
                    for fi28 := 0; fi28 < 12; fi28++ {
                        var v29 int8 = -84
                        var v30 uint8 = 42
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                case "a":
                    _ = len(v3)
                }
            }
        case 2147483647:
            v6 += (v6 % (((^(v6))) | 1))
        }
        _ = fi20
        _ = v21
        _ = v22
    }
    if ("" > "a") {
        var v31 uint64 = 18446744073709551615
        _ = v31
    } else {
        var v32 int32 = -2147483648
        v33 := GenS4{F0: 4294967295, F1: v6}
        if ((v4 | v4) >= v4) {
            v6 = (v6 + v6)
        } else {
            if (false || false) {
                v4 -= (v4 / ((v4) | 1))
                if len(v2) > 0 {
                    v2[0] = 0
                }
            } else {
                v34 := v3[0]
                _ = v34
            }
            if (v1 == v1) {
                v35 := GenS5{F0: 612, F1: (^(v6)), F2: 2147483647}
                _ = v35
            }
            var v36 string = "foo bar"
            _ = v36
        }
        _ = v32
        _ = v33
    }
    v5 = -227
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
