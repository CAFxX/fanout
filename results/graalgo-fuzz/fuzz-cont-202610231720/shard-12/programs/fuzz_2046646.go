package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int8 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 int16; F1 int8 }

func main() {
    var v1 uint16 = 1
    v2 := []int{128, 313}
    var v3 uint64 = 972
    var v4 int32 = 42
    switch -154 {
    case 2147483647:
        v5 := GenS1{F0: 376, F1: 127, F2: 100}
        _ = v5
    case 276:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 55:
        switch "\n\t" {
        case "\n\t":
            if true {
                switch "hello" {
                case "foo bar":
                    if len(v2) > 0 {
                        v2[0] = 4294967295
                    }
                case "hello":
                    if true {
                        v4 += v4
                        var v6 string = ("\n\t" + "\n\t")
                        _ = v6
                    }
                case "\n\t":
                    v4 -= v4
                }
                if len(v2) > 0 {
                    v2[0] = -152
                }
            }
        case "foo bar":
            var v7 int8 = 0
            _ = v7
        }
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 float32 = 3.14159
        var v10 int16 = -761
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 string = "foo bar"
    if (!(true && ((35 < 227) && false))) {
        if (!(!(float64(1e10) != 0.0))) {
            for fi12 := 0; fi12 < 4; fi12++ {
                var v13 float64 = float64(0.1)
                v14 := GenS1{F0: v3, F1: 12, F2: 31}
                _ = fi12
                _ = v13
                _ = v14
            }
        } else {
            v11 = ("" + ("" + "\n\t"))
        }
        if (((-2.25 - 3.14159) == 1e10) && true) {
            v11 = "foo bar"
            var v15 uint = 47
            _ = v15
        } else {
            var v16 uint8 = 0
            _ = v16
        }
        if (640 <= 0) {
            _ = len(v2)
            for fi17 := 0; fi17 < 2; fi17++ {
                v3 -= 154
                v18 := GenS1{F0: 697, F1: 127, F2: -128}
                switch 0 {
                case 355:
                    var v19 uint8 = 1
                    _ = v19
                default:
                    v1 = v1
                }
                _ = fi17
                _ = v18
            }
            switch v11 {
            case "\n\t":
                var v20 int16 = -32768
                _ = v20
            case "foo bar":
                switch -863 {
                case -314:
                    var v21 int8 = -128
                    _ = v21
                case -1:
                    for fi22 := 0; fi22 < 20; fi22++ {
                        if len(v2) > 0 {
                            v2[0] = -9223372036854775808
                        }
                        _ = fi22
                    }
                case 533:
                    for fi23 := 0; fi23 < 1; fi23++ {
                        if len(v2) > 0 {
                            v2[0] = (fi23 >> ((0) & 63))
                        }
                        _ = fi23
                    }
                }
            default:
                v24 := GenS1{F0: 0, F1: 0, F2: 42}
                _ = v24
            }
        }
    } else {
        v4 -= v4
        switch ((v11 + "\n\t") + "foo bar") {
        case "a":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "\n\t":
            fmt.Println(v11, v3)
        }
    }
    v1 = 324
    v25 := GenS1{F0: 953, F1: 0, F2: -40}
    if true {
        v26 := GenS2{F0: 0, F1: (v4 ^ (v4 ^ (v4 / (((v4 & v4)) | 1))))}
        v3 = 18446744073709551615
        _ = v26
    } else {
        v1 += v1
        var v27 uint8 = 98
        var v28 int16 = 42
        _ = v27
        _ = v28
    }
    switch v11 {
    case "":
        v4 = v4
    case "hello":
        for fi29 := 0; fi29 < 4; fi29++ {
            v30 := GenS3{F0: 0, F1: 1}
            _ = fi29
            _ = v30
        }
    case "foo bar":
        var v31 float64 = 0.0
        _ = v31
    default:
        for fi32 := 0; fi32 < 8; fi32++ {
            var v33 uint = 42
            for fi34 := 0; fi34 < 7; fi34++ {
                v2 = append(v2, 1)
                _ = fi34
            }
            _ = fi32
            _ = v33
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
