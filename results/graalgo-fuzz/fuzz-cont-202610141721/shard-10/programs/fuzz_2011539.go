package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 uint64; F1 uint32; F2 string }

func main() {
    var v1 uint8 = 255
    v2 := []uint32{4294967295, 65535, 2147483648}
    v3 := []uint32{778}
    var v4 uint32 = 232
    var v5 uint32 = v4
    v3 = append(v3, (v5 - 382))
    if len(v2) > 0 {
        v2[0] = (v5 * (^(v4)))
    }
    for fi6 := 0; fi6 < 17; fi6++ {
        if (false && true) {
            var v7 int32 = 719
            for fi8 := 0; fi8 < 12; fi8++ {
                fmt.Println(v7, v7, v7)
                fmt.Println(fi8, v1)
                var v9 uint16 = 1
                _ = fi8
                _ = v9
            }
            if (false && (true && true)) {
                fmt.Println(v5)
                switch fi6 {
                case -2147483648:
                    for fi10 := 0; fi10 < 1; fi10++ {
                        v11 := GenS1{F0: 0, F1: 0}
                        var v12 string = ("\n\t" + "\n\t")
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                case 937:
                    if (fi6 == fi6) {
                        v3 = append(v3, 971)
                    }
                }
                fmt.Println(v1, v4)
            } else {
                v1 += v1
            }
            _ = v7
        } else {
            var v13 int32 = 651
            _ = v13
        }
        _ = fi6
    }
    switch "a" {
    case "foo bar":
        var v14 int64 = 928
        _ = v14
    default:
        v2 = append(v2, 863)
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 string = ("" + "\n\t")
        if true {
            for fi17 := 0; fi17 < 4; fi17++ {
                switch fi15 {
                case 536:
                    v18 := GenS2{F0: -879, F1: (fi15 | (fi15 << ((21) & 63)))}
                    if true {
                        var v19 uint8 = 0
                        _ = v19
                    }
                    var v20 uint64 = 2
                    _ = v18
                    _ = v20
                case -447:
                    _ = len(v3)
                case 127:
                    v2 = append(v2, (v4 / ((v5) | 1)))
                }
                var v21 int = fi17
                _ = fi17
                _ = v21
            }
            fmt.Println(v16, v4)
        }
        switch 957 {
        case 128:
            var v22 int = (fi15 - fi15)
            _ = v22
        case -326:
            v4 = 544
        case 9223372036854775807:
            switch 127 {
            case -421:
                v4 += v4
            }
        }
        _ = fi15
        _ = v16
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v4)
    v5 += v4
    var v23 uint64 = 4294967295
    for fi24 := 0; fi24 < 13; fi24++ {
        switch (("foo bar" + "hello") + "a") {
        case "\n\t":
            fmt.Println(v5, fi24)
        }
        v4 = 65535
        if (true && (!(!(18446744073709551615 <= 4294967295)))) {
            v25 := GenS3{F0: ("a" + "hello"), F1: uint16(v23)}
            switch "" {
            case "hello":
                if true {
                    if (2147483647 != 1) {
                        if ((((v5 >> ((0) & 63)) != (v4 ^ (v4 / ((42) | 1)))) || true) && true) {
                            if false {
                                var v26 int64 = 42
                                var v27 int32 = 2
                                _ = v26
                                _ = v27
                            }
                            for fi28 := 0; fi28 < 14; fi28++ {
                                v4 += v5
                                _ = fi28
                            }
                            v5 = v5
                        }
                        fmt.Println(v5, v23, v4)
                    }
                } else {
                    switch "hello" {
                    case "hello":
                        var v29 uint32 = v5
                        _ = v29
                    case "\n\t":
                        var v30 float64 = (1e10 / 3.14159)
                        _ = v30
                    }
                    for fi31 := 0; fi31 < 19; fi31++ {
                        if (uint64(v4) != v23) {
                            v32 := GenS3{F0: "foo bar", F1: 184}
                            v23 = (^(v23))
                            _ = v32
                        }
                        v4 = 65535
                        _ = fi31
                    }
                    v33 := GenS4{F0: v23, F1: 65535, F2: "a"}
                    _ = v33
                }
            case "\n\t":
                if true {
                    var v34 int32 = 0
                    _ = v34
                }
            }
            _ = v25
        }
        _ = fi24
    }
    for fi35 := 0; fi35 < 8; fi35++ {
        var v36 int64 = 42
        v37 := GenS3{F0: "", F1: 255}
        for fi38 := 0; fi38 < 9; fi38++ {
            var v39 float64 = (((float64(0.0) - float64(((-2.25 / 0.1) / 0.1))) / 1e10) / -2.25)
            _ = fi38
            _ = v39
        }
        _ = fi35
        _ = v36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
