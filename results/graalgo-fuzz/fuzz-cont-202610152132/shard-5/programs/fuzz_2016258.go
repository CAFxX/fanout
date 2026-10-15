package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 string }

func main() {
    v1 := make(map[string]bool)
    var v2 int = 42
    var v3 float64 = 0.0
    v4 := make(map[uint32]bool)
    v5 := []string{}
    fmt.Println(v3, v2)
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 int16 = -1
            var v9 int8 = 127
            _ = fi7
            _ = v8
            _ = v9
        }
        if true {
            var v10 uint = 945
            if true {
                fmt.Println(v2, fi6, v2)
                if len(v5) > 0 {
                    v5[0] = ((("a" + (("" + "foo bar") + "\n\t")) + (("foo bar" + "foo bar") + "a")) + "\n\t")
                }
            } else {
                v11 := GenS1{F0: -527, F1: 42, F2: "\n\t"}
                for fi12 := 0; fi12 < 16; fi12++ {
                    switch 4294967295 {
                    case 128:
                        switch "hello" {
                        case "\n\t":
                            var v13 uint16 = 0
                            _ = v13
                        case "":
                            fmt.Println(v2, v2)
                        }
                    default:
                        var v14 int8 = 100
                        _ = v14
                    }
                    v10 = v10
                    _ = fi12
                }
                _ = v11
            }
            _ = v10
        } else {
            switch 193 {
            case 1:
                v2 = (v2 << ((255) & 63))
            case 9223372036854775807:
                _ = len(v4)
            case 32767:
                v2 = fi6
            }
            var v15 int = (fi6 ^ (v2 & (v2 ^ (fi6 - (fi6 / ((32768) | 1))))))
            _ = v15
        }
        _ = fi6
    }
    if false {
        switch "\n\t" {
        case "foo bar":
            v16 := GenS1{F0: 351, F1: 1, F2: "\n\t"}
            _ = v16
        case "":
            var v17 uint64 = 2
            _ = v17
        case "\n\t":
            if len(v5) > 0 {
                v5[0] = ""
            }
        default:
            _ = len(v1)
        }
        var v18 uint32 = 766
        _ = v18
    } else {
        if (9223372036854775807 < 0) {
            var v19 uint32 = 1
            if (true && (false || false)) {
                fmt.Println(v3, v19, v3)
                _ = len(v4)
            }
            for fi20 := 0; fi20 < 1; fi20++ {
                if false {
                    var v21 uint = 2
                    if true {
                        for fi22 := 0; fi22 < 17; fi22++ {
                            v21 += (v21 ^ (^(v21)))
                            v2 -= fi22
                            _ = fi22
                        }
                    }
                    for fi23 := 0; fi23 < 19; fi23++ {
                        v3 = v3
                        _ = fi23
                    }
                    _ = v21
                } else {
                    v19 += 2147483647
                }
                var v24 uint8 = 2
                if true {
                    fmt.Println(v2, v24)
                }
                _ = fi20
                _ = v24
            }
            _ = v19
        } else {
            switch v2 {
            case 803:
                v5 = append(v5, "\n\t")
            }
            if true {
                v5 = append(v5, "\n\t")
                v25 := GenS1{F0: -1, F1: 1, F2: "a"}
                switch v2 {
                case -427:
                    v26 := GenS1{F0: 1, F1: -128, F2: "hello"}
                    _ = v26
                case -9223372036854775808:
                    v2 += (v2 | (v2 << ((394) & 63)))
                }
                _ = v25
            }
        }
    }
    v3 += v3
    _ = len(v1)
    var v27 int64 = 1
    v3 += v3
    v28 := GenS1{F0: 0, F1: 127, F2: "a"}
    fmt.Println(v2, v2)
    var v29 uint8 = 195
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
