package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint8 }

func main() {
    var v1 uint32 = 138
    v2 := make(map[uint32]string)
    var v3 uint16 = 32767
    var v4 uint16 = 42
    v5 := []string{"\n\t", ("a" + "\n\t"), "", "foo bar"}
    switch 0 {
    case 0:
        v3 = 255
    default:
        v3 += (v4 << ((uint(v1)) & 63))
    }
    if true {
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 uint32 = v1
            v4 = (v4 ^ v4)
            _ = fi6
            _ = v7
        }
    } else {
        var v8 string = "\n\t"
        _ = v8
    }
    v9 := GenS1{F0: ("" + ""), F1: (^(v3)), F2: 2}
    var v10 int = 0
    for fi11 := 0; fi11 < 12; fi11++ {
        if false {
            var v12 float32 = -2.25
            _ = v12
        }
        var v13 uint8 = 0
        var v14 int32 = 988
        _ = fi11
        _ = v13
        _ = v14
    }
    fmt.Println(v4, v1, v10)
    v10 = (v10 & (v10 + 4294967295))
    if ((-710 != 1) || true) {
        var v15 uint64 = 467
        fmt.Println(v15, v10)
        for fi16 := 0; fi16 < 15; fi16++ {
            var v17 uint64 = (^(v15))
            _ = fi16
            _ = v17
        }
        _ = v15
    } else {
        v18 := GenS1{F0: ((("hello" + "\n\t") + "\n\t") + "foo bar"), F1: 255, F2: 176}
        fmt.Println(v4)
        _ = v18
    }
    v4 = (v3 & v3)
    if true {
        v3 = (v4 - (v4 | 42))
        switch (v10 ^ (v10 * v10)) {
        case -9223372036854775808:
            v19 := GenS1{F0: "\n\t", F1: 42, F2: 50}
            _ = v19
        case -400:
            if true {
                if true {
                    fmt.Println(v4, v1, v3)
                } else {
                    v4 += v3
                    for fi20 := 0; fi20 < 14; fi20++ {
                        fmt.Println(v10)
                        _ = fi20
                    }
                }
                if true {
                    if false {
                        var v21 int64 = 145
                        switch "" {
                        case "\n\t":
                            fmt.Println(v10)
                        case "":
                            var v22 float64 = -2.25
                            _ = v22
                        case "foo bar":
                            for fi23 := 0; fi23 < 4; fi23++ {
                                fmt.Println(fi23)
                                v2[42] = "a"
                                _ = fi23
                            }
                        }
                        _ = v21
                    }
                    var v24 uint32 = 2
                    v3 = v3
                    _ = v24
                }
                v3 = v4
            }
        case 2147483647:
            var v25 float64 = 3.14159
            _ = v25
        default:
            if (!false) {
                switch "a" {
                case "foo bar":
                    if false {
                        if (88 > 255) {
                            var v26 int = v10
                            v4 -= 32767
                            _ = v26
                        }
                    }
                case "":
                    switch (v10 >> ((116) & 63)) {
                    case 1:
                        for fi27 := 0; fi27 < 8; fi27++ {
                            v10 -= (fi27 + fi27)
                            switch 0 {
                            case 127:
                                var v28 bool = true
                                _ = v28
                            case 42:
                                switch 32767 {
                                case 0:
                                    var v29 int8 = 127
                                    _ = v29
                                default:
                                    v10 = v10
                                }
                            }
                            _ = fi27
                        }
                    }
                default:
                    var v30 int16 = 62
                    _ = v30
                }
                v1 = uint32(v3)
            }
        }
        fmt.Println(v1, v3, v10)
    }
    v1 -= v1
    var v31 uint8 = 83
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v9
}
