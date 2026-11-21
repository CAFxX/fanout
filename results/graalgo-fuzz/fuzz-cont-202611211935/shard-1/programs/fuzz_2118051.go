package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint }
type GenS2 struct { F0 uint16; F1 int32; F2 int64 }
type GenS3 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 int16; F1 uint64 }
func genFn0(p0 uint32, p1 uint) int {
    if false {
        var v4 float64 = (0.0 / -2.25)
        switch "\n\t" {
        case "":
            var v5 float32 = 3.14159
            _ = v5
        case "a":
            v6 := GenS1{F0: 82, F1: 127, F2: p1}
            _ = v6
        case "\n\t":
            p1 -= uint(p0)
        default:
            var v7 uint64 = 18446744073709551615
            _ = v7
        }
        _ = v4
    } else {
        var v8 int32 = 2147483647
        if false {
            var v9 string = "a"
            v10 := GenS2{F0: 2, F1: v8, F2: -1}
            switch "" {
            case "\n\t":
                v9 = v9
            case "a":
                var v11 int8 = 100
                _ = v11
            case "":
                switch (((v9 + ("" + "\n\t")) + v9) + v9) {
                case "":
                    switch 32767 {
                    case 1:
                        v12 := GenS2{F0: 1, F1: (v8 & v8), F2: 9223372036854775807}
                        _ = v12
                    case 9223372036854775807:
                        switch ("foo bar" + (v9 + "\n\t")) {
                        case "foo bar":
                            v13 := GenS3{F0: 65535, F1: -627, F2: 18446744073709551615}
                            _ = v13
                        }
                    case 0:
                        fmt.Println(p0, v8, v8)
                    default:
                        fmt.Println(p1, v8)
                    }
                case "\n\t":
                    var v14 int = 9223372036854775807
                    _ = v14
                }
            }
            _ = v10
            _ = v9
        }
        if ((false || true) || false) {
            p1 = p1
        }
        _ = v8
    }
    var v15 int8 = -114
    v15 = 1
    for fi16 := 0; fi16 < 5; fi16++ {
        p1 -= 1
        _ = fi16
    }
    switch "foo bar" {
    case "foo bar":
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = v15
    return 4294967295
}
func genFn1() uint16 {
    if (42 > 0) {
        var v17 int8 = 127
        _ = v17
    } else {
        v18 := []int{genFn0(462, 65535), 2147483647}
        if ((936 <= 385) && (!false)) {
            for fi19 := 0; fi19 < 16; fi19++ {
                switch (fi19 | -354) {
                case -2147483648:
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                case 28:
                    if len(v18) > 0 {
                        v18[0] = genFn0(1, 1)
                    }
                default:
                    var v20 float32 = (3.14159 / 1.0)
                    _ = v20
                }
                _ = len(v18)
                switch "" {
                case "a":
                    v18 = append(v18, (fi19 | (fi19 / (((fi19 << ((4294967295) & 63))) | 1))))
                default:
                    v21 := GenS2{F0: 344, F1: 1, F2: 44}
                    _ = v21
                }
                _ = fi19
            }
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
        }
        _ = v18
    }
    var v22 uint = 198
    var v23 float32 = float32((0.1 / 1.5))
    _ = v22
    _ = v23
    return 1
}

func main() {
    var v1 uint64 = 2
    var v2 int32 = -432
    var v3 uint32 = 65535
    for fi24 := 0; fi24 < 4; fi24++ {
        if true {
            v25 := GenS4{F0: 1, F1: 42}
            _ = v25
        } else {
            var v26 uint8 = 2
            _ = v26
        }
        v27 := GenS3{F0: (v3 << ((2) & 63)), F1: -662, F2: (v1 - v1)}
        v3 = (v3 / ((2) | 1))
        _ = fi24
        _ = v27
    }
    var v28 int16 = -32768
    var v29 string = "foo bar"
    v1 = 18446744073709551615
    v30 := []int{}
    switch -102 {
    case -107:
        v3 = v3
    case -2147483648:
        var v31 uint16 = 255
        _ = v31
    }
    var v32 int64 = int64(v2)
    switch v29 {
    case "foo bar":
        var v33 string = v29
        _ = v33
    case "":
        v32 = v32
    default:
        if (987 == 176) {
            v3 -= v3
            var v34 float32 = ((3.14159 / (float32(3.14159) * (1.5 + 0.0))) + (3.14159 + 0.1))
            _ = v34
        } else {
            var v35 float64 = (1.5 - (-2.25 + 1e10))
            v36 := GenS3{F0: (v3 >> ((357) & 63)), F1: (v32 & (v32 % ((v32) | 1))), F2: (^(v1))}
            _ = v35
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 2; fi37++ {
        var v38 string = ""
        switch v29 {
        case "foo bar":
            v3 -= 2147483647
        case "\n\t":
            for fi39 := 0; fi39 < 10; fi39++ {
                if false {
                    v1 = v1
                    fmt.Println(fi39)
                } else {
                    fmt.Println(v28, v32)
                }
                _ = fi39
            }
        }
        _ = fi37
        _ = v38
    }
    if len(v30) > 0 {
        v30[0] = -1
    }
    var v40 uint16 = 65535
    v28 -= v28
    fmt.Println(v32, v2)
    if len(v30) >= 2 {
        v30 = v30[1:len(v30)]
    }
    switch ("foo bar" + "foo bar") {
    case "":
        fmt.Println(v29, v1)
    case "\n\t":
        switch v29 {
        case "foo bar":
            v41 := GenS2{F0: (^(v40)), F1: (^(v2)), F2: 9223372036854775807}
            _ = v41
        default:
            if len(v30) > 0 {
                v30[0] = 522
            }
        }
    case "a":
        if false {
            v42 := GenS5{F0: v28, F1: (v1 ^ v1)}
            var v43 int16 = (-(v28))
            _ = v42
            _ = v43
        } else {
            if false {
                v2 = (v2 % ((int32(v40)) | 1))
                if (!true) {
                    v44 := GenS1{F0: int(v2), F1: -72, F2: 908}
                    _ = v44
                } else {
                    for fi45 := 0; fi45 < 15; fi45++ {
                        v40 += 894
                        _ = fi45
                    }
                    var v46 float64 = 1.5
                    fmt.Println(v46, v28)
                    _ = v46
                }
            }
            v47 := GenS3{F0: v3, F1: (v32 - v32), F2: (v1 % ((v1) | 1))}
            _ = v47
        }
    default:
        switch "foo bar" {
        case "":
            for fi48 := 0; fi48 < 19; fi48++ {
                v28 = 1000
                v49 := GenS3{F0: (v3 & 2147483648), F1: v32, F2: v1}
                _ = fi48
                _ = v49
            }
        }
    }
    v50 := GenS3{F0: 4294967295, F1: (v32 << ((787) & 63)), F2: (v1 >> ((278) & 63))}
    v28 -= v28
    if len(v30) > 0 {
        v30[0] = -141
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v32:", v32)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v40
    _ = v50
}
