package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 bool; F1 int64 }

func main() {
    var v1 uint16 = 32767
    var v2 float64 = 1e10
    var v3 bool = false
    var v4 uint = 786
    v5 := []uint32{}
    var v6 int16 = 42
    _ = len(v5)
    fmt.Println(v4)
    v6 += v6
    if v3 {
        if (!((3.14159 >= -2.25) && v3)) {
            for fi7 := 0; fi7 < 13; fi7++ {
                v3 = (v4 > v4)
                if (v4 != (v4 + v4)) {
                    v5 = append(v5, 2)
                    v1 = v1
                } else {
                    v3 = (false && (true && true))
                    if (!(v3 || ((1e10 > -2.25) && true))) {
                        _ = len(v5)
                    } else {
                        _ = len(v5)
                        for fi8 := 0; fi8 < 5; fi8++ {
                            var v9 uint16 = 370
                            var v10 bool = (v6 != (v6 - 0))
                            _ = fi8
                            _ = v10
                            _ = v9
                        }
                    }
                    var v11 float32 = (3.14159 / (0.1 + 3.14159))
                    _ = v11
                }
                _ = fi7
            }
        } else {
            v2 += -2.25
        }
    } else {
        switch "hello" {
        case "a":
            var v12 uint8 = 2
            _ = v12
        case "hello":
            for fi13 := 0; fi13 < 17; fi13++ {
                if v3 {
                    var v14 float64 = v2
                    v15 := GenS1{F0: (v1 / (((v1 - 2)) | 1)), F1: "hello"}
                    v1 -= (^(v1))
                    _ = v14
                    _ = v15
                } else {
                    var v16 int8 = 42
                    v17 := GenS1{F0: 255, F1: "hello"}
                    _ = v16
                    _ = v17
                }
                v18 := GenS2{F0: (!v3), F1: -1}
                fmt.Println(v6, fi13)
                _ = fi13
                _ = v18
            }
        default:
            fmt.Println(v1, v6)
        }
        fmt.Println(v6, v4, v1)
    }
    if (("" + "foo bar") >= "a") {
        _ = len(v5)
        if (v3 && true) {
            if false {
                v5 = append(v5, 67)
                if v3 {
                    _ = len(v5)
                    fmt.Println(v4, v3)
                } else {
                    var v19 int16 = 32767
                    if (0 < -992) {
                        v19 = (v6 * 1000)
                        var v20 int32 = 2147483647
                        _ = v20
                    }
                    var v21 uint16 = v1
                    _ = v19
                    _ = v21
                }
            } else {
                v3 = (-1 < 219)
            }
        }
    }
    if v3 {
        switch ((("" + "foo bar") + "hello") + ("foo bar" + ("\n\t" + ("\n\t" + "foo bar")))) {
        case "hello":
            if len(v5) > 0 {
                v5[0] = 0
            }
        default:
            fmt.Println(v3)
        }
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        for fi23 := 0; fi23 < 7; fi23++ {
            fmt.Println(v1, fi23)
            var v24 uint64 = uint64(v1)
            switch "foo bar" {
            case "hello":
                if v3 {
                    v24 = v24
                    v6 = v6
                    v5 = append(v5, 2147483648)
                } else {
                    fmt.Println(v24, v3)
                    var v25 float64 = 3.14159
                    _ = v25
                }
            case "\n\t":
                var v26 uint8 = 2
                _ = v26
            }
            var v27 uint32 = 42
            _ = fi23
            _ = v24
            _ = v27
        }
        _ = fi22
    }
    switch int(v4) {
    case 374:
        var v28 float32 = (float32(1.5) + ((1.5 / 0.1) * 0.0))
        _ = v28
    default:
        var v29 int32 = int32(v6)
        _ = v29
    }
    v30 := GenS2{F0: (2147483647 < 0), F1: 173}
    _ = len(v5)
    var v31 string = "hello"
    v32 := GenS2{F0: false, F1: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
