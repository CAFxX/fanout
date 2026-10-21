package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS2 struct { F0 bool; F1 int; F2 int32 }
type GenS3 struct { F0 uint16; F1 int }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []string{"a", ""}
    var v3 int32 = 1
    var v4 uint16 = 42
    switch 0 {
    case -2147483648:
        v1 -= v1
    case 128:
        var v5 float32 = 1e10
        _ = v5
    }
    for fi6 := 0; fi6 < 15; fi6++ {
        fmt.Println(v3, fi6, v1)
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS1{F0: v4, F1: 65535, F2: 0}
            v9 := GenS1{F0: 329, F1: 1, F2: 255}
            v1 -= v1
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    v4 = (v4 * v4)
    switch 1 {
    case 42:
        v10 := GenS1{F0: 255, F1: 931, F2: 87}
        _ = v10
    case 128:
        v2 = append(v2, "")
    default:
        if false {
            var v11 int = 32767
            for fi12 := 0; fi12 < 11; fi12++ {
                var v13 uint = 2
                _ = fi12
                _ = v13
            }
            _ = v11
        }
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 uint16 = (v4 + (v4 & (v4 * 0)))
        v1 += (v1 - v1)
        _ = fi14
        _ = v15
    }
    if false {
        if (0 != -32768) {
            for fi16 := 0; fi16 < 7; fi16++ {
                v1 -= v1
                v17 := GenS1{F0: (^(v4)), F1: 42, F2: 42}
                _ = fi16
                _ = v17
            }
            for fi18 := 0; fi18 < 15; fi18++ {
                v2 = append(v2, "")
                fmt.Println(v3, v1, v3)
                v19 := GenS2{F0: (!true), F1: (-(fi18)), F2: (v3 >> ((46) & 63))}
                _ = fi18
                _ = v19
            }
        }
        switch "" {
        case "foo bar":
            fmt.Println(v1, v1)
        default:
            var v20 string = "hello"
            _ = v20
        }
    }
    v1 = (v1 | v1)
    fmt.Println(v1)
    fmt.Println(v4, v3, v1)
    var v21 uint16 = 44
    v22 := GenS3{F0: v21, F1: -2147483648}
    v2 = append(v2, ("" + "hello"))
    if true {
        var v23 int8 = -128
        _ = v23
    } else {
        if (false || true) {
            for fi24 := 0; fi24 < 3; fi24++ {
                for fi25 := 0; fi25 < 12; fi25++ {
                    var v26 int = fi25
                    if false {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if len(v2) > 0 {
                            v2[0] = ("" + "\n\t")
                        }
                    }
                    _ = fi25
                    _ = v26
                }
                _ = fi24
            }
            v1 = (v1 & (v1 ^ (v1 >> ((670) & 63))))
        }
    }
    v3 = -316
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
