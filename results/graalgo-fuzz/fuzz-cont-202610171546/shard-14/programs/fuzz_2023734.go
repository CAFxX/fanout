package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int }
type GenS2 struct { F0 int16; F1 int; F2 bool }
type GenS3 struct { F0 int16; F1 int64; F2 uint64 }
type GenS4 struct { F0 int8; F1 uint32; F2 uint64 }

func main() {
    var v1 int8 = -128
    v2 := []int{9223372036854775807}
    var v3 uint8 = 0
    v4 := GenS1{F0: 42, F1: -1, F2: -148}
    var v5 bool = (!false)
    v2 = append(v2, -657)
    var v6 int32 = -518
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    var v7 uint16 = 255
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v2 = append(v2, 32768)
    var v8 uint64 = 1
    fmt.Println(v6, v8)
    for fi9 := 0; fi9 < 8; fi9++ {
        if (true || v5) {
            if ((v8 << ((42) & 63)) < 4294967295) {
                for fi10 := 0; fi10 < 1; fi10++ {
                    _ = len(v2)
                    v11 := GenS2{F0: 0, F1: (-(fi9)), F2: ((v6 / (((v6 / ((v6) | 1))) | 1)) >= v6)}
                    _ = fi10
                    _ = v11
                }
                if ((0.0 / -2.25) < (3.14159 / 1.0)) {
                    fmt.Println(v5)
                } else {
                    var v12 int64 = 1
                    _ = v12
                }
            } else {
                v3 -= (v3 << ((4294967295) & 63))
                var v13 uint16 = (v7 | 255)
                _ = v13
            }
            var v14 int64 = -1
            _ = v14
        }
        _ = fi9
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        v7 = v7
        _ = fi15
    }
    for fi16 := 0; fi16 < 20; fi16++ {
        for fi17 := 0; fi17 < 10; fi17++ {
            switch "" {
            case "\n\t":
                if true {
                    v8 -= (^(v8))
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v18 string = "hello"
                    v7 = v7
                    if ((false || (65535 == (v7 / ((32767) | 1)))) || v5) {
                        v6 = -1
                        switch (fi16 << ((255) & 63)) {
                        case -296:
                            if len(v2) > 0 {
                                v2[0] = (-(fi17))
                            }
                        default:
                            v5 = (32767 != v7)
                        }
                        if v5 {
                            var v19 int8 = -119
                            v2 = append(v2, fi17)
                            if len(v2) > 0 {
                                v2[0] = (fi17 ^ (fi16 - 1))
                            }
                            _ = v19
                        }
                    }
                    switch v18 {
                    case "hello":
                        var v20 uint8 = v3
                        _ = v20
                    case "\n\t":
                        fmt.Println(fi16, v5, v5)
                    case "foo bar":
                        var v21 int = fi17
                        _ = v21
                    default:
                        if (!false) {
                            switch (-(fi16)) {
                            case 128:
                                v7 -= v7
                                switch v18 {
                                case "":
                                    switch (^(fi16)) {
                                    case -1:
                                        var v22 int32 = 2147483647
                                        _ = v22
                                    case 32767:
                                        v1 = v1
                                        if false {
                                            fmt.Println(fi16)
                                            var v23 int32 = (v6 / (((v6 << ((65535) & 63))) | 1))
                                            var v24 float32 = 3.14159
                                            _ = v23
                                            _ = v24
                                        } else {
                                            v6 = v6
                                            if true {
                                                if v5 {
                                                    if (((v8 + (^(v8))) >= v8) && true) {
                                                        v25 := GenS3{F0: int16(v3), F1: -207, F2: v8}
                                                        fmt.Println(v5, fi16, fi16)
                                                        v26 := GenS4{F0: 42, F1: 0, F2: (v8 & v8)}
                                                        if (0.1 == (1.5 / (((-2.25 - (-2.25 - -2.25)) + (((float64(float32(-2.25)) / -2.25) * (0.0 * (float64(-2.25) / -2.25))) - (3.14159 - 0.0))) + (0.1 - (1e10 - 0.0))))) {
                                                            v3 += v3
                                                            var v27 int16 = 0
                                                            v28 := GenS4{F0: (v1 + (v1 >> ((0) & 63))), F1: 0, F2: (v8 >> ((65535) & 63))}
                                                            _ = v27
                                                            _ = v28
                                                        }
                                                        _ = v25
                                                        _ = v26
                                                    }
                                                } else {
                                                    var v29 string = v18
                                                    switch v29 {
                                                    case "foo bar":
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    case "hello":
                                                        _ = len(v2)
                                                    default:
                                                        fmt.Println(v3)
                                                        switch v29 {
                                                        case "foo bar":
                                                            v30 := GenS1{F0: 40, F1: fi16, F2: (fi17 ^ fi17)}
                                                            v29 = "foo bar"
                                                            _ = v30
                                                        case "":
                                                            var v31 int64 = -9223372036854775808
                                                            _ = v31
                                                        }
                                                        v32 := GenS2{F0: -32768, F1: (fi17 >> ((65535) & 63)), F2: (v6 != (-(v6)))}
                                                        _ = v32
                                                    }
                                                    _ = v29
                                                }
                                                var v33 uint32 = 65535
                                                if len(v2) > 0 {
                                                    v2[0] = (fi17 >> ((18446744073709551615) & 63))
                                                }
                                                v34 := GenS4{F0: v1, F1: (v33 << ((255) & 63)), F2: 18446744073709551615}
                                                v3 += v3
                                                _ = v33
                                                _ = v34
                                            }
                                            var v35 int = fi16
                                            _ = v35
                                        }
                                        if ((0 > 1) && true) {
                                            v1 += int8(v8)
                                        }
                                    default:
                                        v7 = (v7 / ((v7) | 1))
                                    }
                                case "foo bar":
                                    v5 = v5
                                }
                                v18 = (v18 + "foo bar")
                            case -2147483648:
                                if ((fi17 >> ((18446744073709551615) & 63)) <= (fi16 + fi16)) {
                                    v7 = (^(v7))
                                    _ = len(v2)
                                }
                            case 0:
                                v36 := GenS3{F0: 360, F1: -1, F2: (v8 ^ (v8 - uint64(v7)))}
                                _ = v36
                            default:
                                fmt.Println(fi17)
                            }
                        }
                    }
                    var v37 int16 = 32767
                    _ = v18
                    _ = v37
                } else {
                    fmt.Println(v3)
                    fmt.Println(fi17, fi17, v5)
                    var v38 uint16 = 2
                    var v39 float64 = 1e10
                    v5 = (-1 < 0)
                    _ = v38
                    _ = v39
                }
            }
            switch "" {
            case "hello":
                var v40 int32 = (v6 ^ int32(v8))
                var v41 int64 = 0
                _ = v40
                _ = v41
            case "":
                fmt.Println(v5)
                v7 = v7
                switch (("hello" + "hello") + ((("hello" + "\n\t") + "") + "a")) {
                case "foo bar":
                    var v42 int8 = v1
                    _ = v42
                }
            case "foo bar":
                var v43 uint32 = 42
                _ = v43
            default:
                var v44 string = "a"
                _ = v44
            }
            var v45 uint8 = 2
            _ = fi17
            _ = v45
        }
        for fi46 := 0; fi46 < 18; fi46++ {
            if (v5 && v5) {
                fmt.Println(v8)
                if len(v2) > 0 {
                    v2[0] = fi46
                }
                if v5 {
                    var v47 int16 = 960
                    _ = len(v2)
                    _ = v47
                } else {
                    v3 = v3
                    var v48 uint8 = v3
                    var v49 uint16 = 827
                    _ = v48
                    _ = v49
                }
            }
            _ = len(v2)
            _ = fi46
        }
        _ = fi16
    }
    var v50 uint = 4294967295
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
    _ = v7
    _ = v8
}
