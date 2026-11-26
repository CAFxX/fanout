package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 int32; F1 int; F2 int32 }

func main() {
    v1 := []uint32{1}
    var v2 bool = false
    var v3 int64 = 9223372036854775807
    var v4 uint16 = 289
    v5 := []string{}
    for fi6 := 0; fi6 < 7; fi6++ {
        v3 = int64(fi6)
        switch fi6 {
        case 127:
            var v7 float64 = (float64((0.0 - 1e10)) / 1.0)
            _ = v7
        case 626:
            if len(v5) > 0 {
                v5[0] = "foo bar"
            }
        }
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 int = 363
            _ = fi8
            _ = v9
        }
        _ = fi6
    }
    if false {
        if false {
            fmt.Println(v4)
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v10 uint8 = 172
        _ = v10
    } else {
        fmt.Println(v2, v4)
        if v2 {
            v5 = append(v5, "a")
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v5 = append(v5, "")
        } else {
            fmt.Println(v2, v3)
            v11 := GenS1{F0: "a"}
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 uint16 = (v4 | (v4 & v4))
        _ = fi12
        _ = v13
    }
    var v14 float64 = 1.5
    var v15 uint8 = 42
    v1 = append(v1, 615)
    if (!v2) {
        v15 = v15
    } else {
        v16 := GenS2{F0: (v15 / ((uint8(v4)) | 1)), F1: 42}
        v2 = ((("hello" + "hello") + "a") < "foo bar")
        v17 := GenS2{F0: v15, F1: 42}
        _ = v16
        _ = v17
    }
    v4 += v4
    if (!false) {
        if v2 {
            for fi18 := 0; fi18 < 8; fi18++ {
                v19 := GenS3{F0: v2, F1: -2147483648}
                for fi20 := 0; fi20 < 1; fi20++ {
                    var v21 uint64 = 2
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    switch (-(fi18)) {
                    case 9223372036854775807:
                        var v22 uint64 = (v21 << ((18446744073709551615) & 63))
                        _ = v22
                    case -598:
                        v4 = 839
                        if len(v5) > 0 {
                            v5[0] = "foo bar"
                        }
                    default:
                        var v23 int32 = 42
                        _ = v23
                    }
                    var v24 float32 = float32(float64(-2.25))
                    _ = fi20
                    _ = v21
                    _ = v24
                }
                _ = fi18
                _ = v19
            }
            if v2 {
                var v25 float32 = 0.0
                _ = v25
            } else {
                for fi26 := 0; fi26 < 8; fi26++ {
                    v27 := GenS4{F0: 1, F1: (fi26 / (((fi26 + -9223372036854775808)) | 1)), F2: 2147483647}
                    _ = fi26
                    _ = v27
                }
            }
        }
    }
    if false {
        for fi28 := 0; fi28 < 8; fi28++ {
            if (false && v2) {
                switch fi28 {
                case 572:
                    var v29 uint8 = 221
                    _ = v29
                case 32767:
                    for fi30 := 0; fi30 < 12; fi30++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi30
                    }
                case -77:
                    if len(v1) > 0 {
                        v1[0] = uint32(v4)
                    }
                default:
                    v5 = append(v5, "")
                }
                fmt.Println(v14, v3, v14)
            }
            switch fi28 {
            case -378:
                var v31 string = "hello"
                _ = v31
            case 9223372036854775807:
                var v32 uint = 449
                _ = v32
            case 36:
                v14 = (1e10 - (3.14159 * v14))
            default:
                v14 = float64(1.5)
            }
            var v33 uint32 = 1
            _ = fi28
            _ = v33
        }
        switch ("" + "") {
        case "hello":
            var v34 string = (((("" + "") + ("" + "")) + ("\n\t" + "foo bar")) + "foo bar")
            _ = v34
        default:
            for fi35 := 0; fi35 < 5; fi35++ {
                v4 += 255
                if v2 {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    if len(v1) > 0 {
                        v1[0] = 0
                    }
                } else {
                    v15 = (v15 ^ (^(v15)))
                }
                _ = fi35
            }
        }
        v3 = (v3 | (v3 >> ((uint(v15)) & 63)))
    }
    v14 = (v14 * ((((v14 / float64((0.0 - (1e10 * (0.0 - 3.14159))))) / (v14 / (v14 - float64(3.14159)))) / 1e10) + -2.25))
    fmt.Println(v14, v14, v14)
    var v36 int32 = -183
    if (("hello" + (("hello" + (("foo bar" + "foo bar") + "")) + "hello")) <= ("\n\t" + ("a" + "hello"))) {
        var v37 uint64 = 18446744073709551615
        var v38 string = "\n\t"
        _ = v37
        _ = v38
    }
    v3 += (v3 + v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
