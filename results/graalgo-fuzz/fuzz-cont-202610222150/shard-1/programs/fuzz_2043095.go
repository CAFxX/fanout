package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint8 }
func genFn0(p0 int32, p1 bool, p2 int16) bool {
    fmt.Println(p2, p2, p1)
    var v4 uint = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (-266 <= 32768)
}

func main() {
    var v1 uint8 = 146
    v2 := []string{("foo bar" + ("foo bar" + ("a" + "a"))), "a", "a", "\n\t"}
    v3 := make(map[int]bool)
    fmt.Println(v1)
    v1 -= (^(v1))
    var v5 uint = 1
    switch 910 {
    case -95:
        v1 = 1
    case 4294967295:
        if true {
            for fi6 := 0; fi6 < 8; fi6++ {
                fmt.Println(fi6, fi6)
                _ = fi6
            }
            v1 = v1
            for fi7 := 0; fi7 < 11; fi7++ {
                var v8 int = fi7
                _ = fi7
                _ = v8
            }
        }
    case -1:
        var v9 int = -2147483648
        _ = v9
    }
    var v10 uint = v5
    if (255 < v1) {
        var v11 int16 = -32768
        _ = v11
    } else {
        fmt.Println(v5)
    }
    if false {
        fmt.Println(v1, v10, v5)
        var v12 float64 = -2.25
        _ = v12
    } else {
        var v13 uint16 = 370
        switch 9223372036854775807 {
        case 32767:
            for fi14 := 0; fi14 < 18; fi14++ {
                for fi15 := 0; fi15 < 10; fi15++ {
                    var v16 uint32 = 4294967295
                    _ = fi15
                    _ = v16
                }
                if (true && true) {
                    var v17 int16 = 1
                    v18 := v3[fi14]
                    v13 += 326
                    _ = v17
                    _ = v18
                }
                _ = fi14
            }
        case 852:
            var v19 uint32 = 698
            _ = v19
        case 0:
            for fi20 := 0; fi20 < 8; fi20++ {
                v5 += 78
                _ = len(v2)
                v21 := GenS1{F0: (true && true), F1: (v13 ^ (v13 * 28)), F2: (v1 / (((v1 - 2)) | 1))}
                _ = fi20
                _ = v21
            }
        default:
            var v22 float64 = (float64(3.14159) / 3.14159)
            _ = v22
        }
        _ = v13
    }
    if false {
        if (-546 == 9223372036854775807) {
            var v23 uint8 = (v1 & (v1 * 32))
            for fi24 := 0; fi24 < 15; fi24++ {
                fmt.Println(v10, v23)
                switch (fi24 - (fi24 * (fi24 << ((v10) & 63)))) {
                case -1:
                    v1 -= v23
                case -2147483648:
                    for fi25 := 0; fi25 < 14; fi25++ {
                        v10 = v5
                        _ = len(v2)
                        switch "foo bar" {
                        case "hello":
                            var v26 uint16 = 1
                            _ = v26
                        default:
                            v2 = append(v2, ("" + ((("" + "hello") + "a") + "foo bar")))
                        }
                        v5 = v5
                        var v27 uint64 = 42
                        _ = fi25
                        _ = v27
                    }
                }
                var v28 uint16 = 42
                _ = fi24
                _ = v28
            }
            var v29 uint64 = 276
            _ = v23
            _ = v29
        }
        if (905 == 0) {
            var v30 int64 = -1
            v30 += v30
            _ = v30
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v5
}
