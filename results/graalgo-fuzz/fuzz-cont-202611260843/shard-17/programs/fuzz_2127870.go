package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 255
    v4 := []uint32{}
    var v5 float64 = 3.14159
    var v6 uint64 = 813
    if ((2 > 65535) && false) {
        for fi7 := 0; fi7 < 9; fi7++ {
            switch "\n\t" {
            case "":
                v5 = 1e10
            case "a":
                var v8 int64 = 9223372036854775807
                _ = v8
            case "foo bar":
                v4 = append(v4, 41)
            }
            v5 = v5
            _ = fi7
        }
        switch 4294967295 {
        case -474:
            var v9 string = ("" + "")
            _ = v9
        }
    } else {
        if true {
            fmt.Println(v3, v6)
            v10 := GenS1{F0: 1, F1: 248, F2: (!("a" != "foo bar"))}
            if (702 > -172) {
                v5 -= (1e10 * -2.25)
                if len(v4) > 0 {
                    v4[0] = uint32(v6)
                }
                v3 += (^(v3))
            }
            _ = v10
        }
        v2[2147483647] = 0
        v5 = (v5 + ((0.1 + 0.1) - v5))
    }
    var v11 uint16 = 255
    fmt.Println(v5)
    var v12 int64 = 858
    fmt.Println(v6, v6)
    var v13 uint8 = 128
    v13 -= (v3 & v3)
    if true {
        v6 += v6
        switch 875 {
        case 0:
            v14 := GenS1{F0: v13, F1: 0, F2: (v5 == 3.14159)}
            _ = v14
        case -179:
            for fi15 := 0; fi15 < 13; fi15++ {
                switch ("foo bar" + "\n\t") {
                case "a":
                    if (v5 < (0.0 / 1e10)) {
                        switch "a" {
                        case "foo bar":
                            for fi16 := 0; fi16 < 3; fi16++ {
                                switch ("\n\t" + "hello") {
                                case "\n\t":
                                    v17 := GenS1{F0: (v3 | 127), F1: -32768, F2: false}
                                    _ = v17
                                case "a":
                                    v3 = v13
                                }
                                _ = len(v4)
                                switch (fi15 & (-(fi15))) {
                                case 2147483647:
                                    fmt.Println(v6, v5, fi15)
                                case -245:
                                    v4 = append(v4, 65535)
                                }
                                _ = fi16
                            }
                        case "hello":
                            if (63 <= 4294967295) {
                                var v18 int32 = 42
                                v19 := GenS2{F0: v12}
                                _ = v18
                                _ = v19
                            }
                        case "a":
                            var v20 int8 = 100
                            _ = v20
                        }
                    }
                case "hello":
                    if false {
                        v5 -= 1e10
                        fmt.Println(v5, fi15, v5)
                    }
                case "\n\t":
                    switch (fi15 / ((fi15) | 1)) {
                    case -357:
                        v3 -= (v3 & (v13 - (v3 * v3)))
                    default:
                        v11 = v11
                    }
                }
                _ = fi15
            }
        case -958:
            for fi21 := 0; fi21 < 1; fi21++ {
                v11 = 1
                v3 = 1
                if false {
                    v4 = append(v4, 1)
                }
                _ = fi21
            }
        default:
            var v22 uint64 = (v6 ^ v6)
            _ = v22
        }
        for fi23 := 0; fi23 < 4; fi23++ {
            var v24 uint = 0
            _ = fi23
            _ = v24
        }
    } else {
        if true {
            if true {
                switch "" {
                case "":
                    var v25 int16 = 715
                    _ = v25
                case "hello":
                    var v26 string = ""
                    _ = v26
                }
                v27 := GenS2{F0: v12}
                var v28 string = "foo bar"
                _ = v27
                _ = v28
            }
        }
        v29 := GenS3{F0: -996}
        if (!false) {
            var v30 int = -869
            if len(v4) > 0 {
                v4[0] = 34
            }
            var v31 uint64 = 646
            _ = v30
            _ = v31
        }
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
