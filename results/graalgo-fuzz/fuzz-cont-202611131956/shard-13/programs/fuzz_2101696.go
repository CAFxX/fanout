package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 int16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 641
    v3 := []int{}
    v4 := []uint32{65535, 602, 0, 2147483647}
    v5 := make(map[uint32]string)
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 uint = 42
        v3 = append(v3, (-(fi6)))
        switch fi6 {
        case -495:
            switch "" {
            case "foo bar":
                fmt.Println(v2)
            }
        case 9223372036854775807:
            v7 = (v7 ^ 2)
        case 407:
            if len(v3) > 0 {
                v3[0] = fi6
            }
        default:
            if ((-2.25 * 0.0) <= 0.1) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                for fi8 := 0; fi8 < 12; fi8++ {
                    v9 := GenS1{F0: "", F1: v2, F2: v7}
                    switch ("\n\t" + (("hello" + "a") + "a")) {
                    case "foo bar":
                        fmt.Println(v7, fi8)
                    default:
                        v2 = -202
                    }
                    if (false || ((!((-2.25 / ((3.14159 - 1.5) * 1e10)) <= 3.14159)) && false)) {
                        v10 := GenS1{F0: "a", F1: -2147483648, F2: v7}
                        _ = v10
                    }
                    _ = fi8
                    _ = v9
                }
            } else {
                v11 := GenS1{F0: (("foo bar" + "foo bar") + ((("a" + ("\n\t" + ("" + "\n\t"))) + "hello") + "")), F1: v2, F2: v7}
                _ = v11
            }
        }
        _ = fi6
        _ = v7
    }
    v2 += (v2 / (((v2 ^ 788)) | 1))
    for fi12 := 0; fi12 < 20; fi12++ {
        switch "\n\t" {
        case "foo bar":
            v2 = v2
        case "hello":
            switch (fi12 ^ fi12) {
            case -2147483648:
                var v13 float64 = (3.14159 - ((0.1 / 1.0) * 1e10))
                _ = v13
            default:
                for fi14 := 0; fi14 < 8; fi14++ {
                    v2 += (^(v2))
                    v15 := GenS1{F0: "\n\t", F1: v2, F2: 18446744073709551615}
                    _ = fi14
                    _ = v15
                }
            }
        case "":
            v2 = (^(v2))
        }
        v2 = 42
        _ = fi12
    }
    var v16 string = "\n\t"
    v17 := GenS1{F0: ((v16 + (v16 + ("" + ((v16 + ((("hello" + v16) + "foo bar") + (v16 + (v16 + "hello")))) + (v16 + (v16 + v16)))))) + ("" + v16)), F1: v2, F2: 617}
    switch 32767 {
    case -171:
        var v18 uint16 = 32767
        _ = v18
    }
    switch 445 {
    case 936:
        switch v16 {
        case "":
            fmt.Println(v16, v16)
        case "foo bar":
            for fi19 := 0; fi19 < 4; fi19++ {
                if false {
                    v20 := GenS2{F0: 255, F1: 648, F2: 2}
                    _ = v20
                } else {
                    var v21 uint32 = 2147483647
                    var v22 uint16 = 1
                    _ = v21
                    _ = v22
                }
                v16 = v16
                v2 = v2
                _ = fi19
            }
        case "\n\t":
            _ = len(v1)
        default:
            v23 := GenS3{F0: true, F1: -82}
            _ = v23
        }
    case 4294967295:
        if (2147483647 == (^(v2))) {
            v24 := GenS4{F0: 427}
            var v25 uint64 = 71
            _ = len(v3)
            _ = v24
            _ = v25
        } else {
            if (100 < -56) {
                v3 = append(v3, 32767)
                if (42 <= 1) {
                    var v26 uint32 = 0
                    _ = v26
                }
                v27 := v5[0]
                _ = v27
            }
            v4 = append(v4, 1)
        }
    }
    for fi28 := 0; fi28 < 8; fi28++ {
        var v29 float64 = -2.25
        var v30 uint64 = 636
        _ = fi28
        _ = v29
        _ = v30
    }
    if (128 != 0) {
        var v31 float64 = 3.14159
        _ = v31
    }
    v2 = v2
    if (!((-1 > 42) || false)) {
        v32 := GenS3{F0: (128 != 0), F1: -128}
        _ = v32
    }
    if (!false) {
        var v33 uint32 = 42
        v33 = v33
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
