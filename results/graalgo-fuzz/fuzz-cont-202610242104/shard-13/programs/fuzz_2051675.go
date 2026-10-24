package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int; F1 uint8; F2 string }
type GenS3 struct { F0 int }

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 100
    v3 := []uint32{0, 2147483647, 736, 2147483648}
    switch 127 {
    case -707:
        if (false || false) {
            _ = len(v1)
            fmt.Println(v2)
            v4 := GenS1{F0: -1, F1: -1}
            _ = v4
        }
    case 42:
        v2 = (v2 * (v2 << ((652) & 63)))
    default:
        delete(v1, "hello")
    }
    var v5 uint32 = 818
    for fi6 := 0; fi6 < 4; fi6++ {
        v5 = (v5 ^ 65535)
        v7 := GenS1{F0: -2147483648, F1: fi6}
        switch 127 {
        case 128:
            for fi8 := 0; fi8 < 7; fi8++ {
                if false {
                    v2 = v2
                    v3 = append(v3, 471)
                    var v9 uint8 = 1
                    _ = v9
                }
                switch fi6 {
                case 1:
                    var v10 int32 = -2147483648
                    _ = v10
                case -2147483648:
                    var v11 int = (^(fi6))
                    _ = v11
                }
                v12 := GenS1{F0: 548, F1: (fi8 & fi8)}
                _ = fi8
                _ = v12
            }
        }
        _ = fi6
        _ = v7
    }
    delete(v1, (("" + "a") + (("foo bar" + "") + "\n\t")))
    var v13 float64 = -2.25
    var v14 uint8 = 2
    v14 += v14
    if false {
        v15 := GenS2{F0: 2147483647, F1: (v14 << ((2) & 63)), F2: "\n\t"}
        var v16 int = -9223372036854775808
        var v17 float32 = (float32(v13) + 0.0)
        _ = v15
        _ = v16
        _ = v17
    }
    v18 := GenS3{F0: 0}
    if true {
        var v19 float64 = float64((1e10 + 1e10))
        for fi20 := 0; fi20 < 19; fi20++ {
            v13 += 0.1
            fmt.Println(v2, v13)
            _ = fi20
        }
        var v21 string = "a"
        _ = v19
        _ = v21
    } else {
        var v22 int16 = 504
        delete(v1, "foo bar")
        _ = v22
    }
    if true {
        v23 := GenS3{F0: int(v2)}
        fmt.Println(v13, v13, v5)
        switch "a" {
        case "a":
            v5 = v5
        case "":
            for fi24 := 0; fi24 < 2; fi24++ {
                var v25 float32 = -2.25
                _ = fi24
                _ = v25
            }
        case "\n\t":
            for fi26 := 0; fi26 < 20; fi26++ {
                v27 := GenS1{F0: -1, F1: 32767}
                v28 := v1["foo bar"]
                switch "foo bar" {
                case "a":
                    var v29 uint64 = 287
                    _ = v29
                case "\n\t":
                    var v30 int = (v28 * (fi26 << ((2) & 63)))
                    _ = v30
                default:
                    if (false || true) {
                        for fi31 := 0; fi31 < 20; fi31++ {
                            v3 = append(v3, (^(v5)))
                            var v32 bool = (false && true)
                            _ = fi31
                            _ = v32
                        }
                    } else {
                        if ((239 != 0) && ((fi26 / (((v28 & fi26)) | 1)) > 9223372036854775807)) {
                            if len(v3) > 0 {
                                v3[0] = (v5 - v5)
                            }
                        } else {
                            v28 = v28
                            v28 += (fi26 + (fi26 ^ -593))
                            v13 += -2.25
                        }
                        var v33 float64 = v13
                        _ = v33
                    }
                }
                _ = fi26
                _ = v27
                _ = v28
            }
        default:
            var v34 float64 = (0.1 * ((((float64((((0.1 + 0.1) * -2.25) - 1.5)) * -2.25) / 3.14159) + 1.5) / 1.5))
            _ = v34
        }
        _ = v23
    }
    v13 += 1e10
    var v35 uint16 = 65535
    var v36 float64 = float64(0.0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v5
}
