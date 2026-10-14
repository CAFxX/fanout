package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 bool; F2 string }
type GenS3 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 int8; F1 int8; F2 int32 }

func main() {
    v1 := []uint32{42, 2147483648, 739}
    var v2 uint8 = 127
    var v3 int16 = -32768
    var v4 int64 = -612
    var v5 float32 = ((-2.25 - 0.1) / 1e10)
    var v6 int64 = (^(v4))
    var v7 string = ("\n\t" + "hello")
    var v8 uint = 42
    var v9 bool = ((v3 != (v3 | v3)) || (2 != 1))
    if v9 {
        for fi10 := 0; fi10 < 16; fi10++ {
            if true {
                var v11 uint64 = 441
                v12 := GenS1{F0: (v3 >> ((42) & 63))}
                _ = v11
                _ = v12
            }
            _ = fi10
        }
        var v13 int8 = 0
        _ = v13
    } else {
        var v14 int64 = v4
        if false {
            if ((((v7 + "hello") + (v7 + (v7 + v7))) + v7) >= "") {
                v14 = (^(v4))
                var v15 int8 = 42
                v16 := GenS1{F0: (v3 << ((uint(v2)) & 63))}
                _ = v15
                _ = v16
            } else {
                if (!v9) {
                    switch -2147483648 {
                    case -925:
                        if (-9223372036854775808 == -436) {
                            var v17 uint8 = v2
                            var v18 int = 2
                            if true {
                                fmt.Println(v9)
                                var v19 uint = 737
                                _ = v19
                            } else {
                                v20 := GenS2{F0: (v7 + v7), F1: v9, F2: v7}
                                fmt.Println(v18, v4, v2)
                                _ = v20
                            }
                            _ = v17
                            _ = v18
                        }
                    }
                } else {
                    if (v9 || true) {
                        _ = len(v1)
                        for fi21 := 0; fi21 < 15; fi21++ {
                            if v9 {
                                fmt.Println(v4)
                            } else {
                                fmt.Println(v6)
                                v2 = v2
                                v6 = 749
                            }
                            v22 := GenS1{F0: 42}
                            for fi23 := 0; fi23 < 12; fi23++ {
                                if ((v3 + (v3 ^ v3)) <= v3) {
                                    _ = len(v1)
                                } else {
                                    v9 = true
                                    var v24 string = v7
                                    _ = v24
                                }
                                v1 = append(v1, 4294967295)
                                _ = len(v1)
                                _ = fi23
                            }
                            _ = fi21
                            _ = v22
                        }
                        var v25 uint64 = 1
                        _ = v25
                    } else {
                        if v9 {
                            v3 = v3
                            v7 = v7
                            for fi26 := 0; fi26 < 20; fi26++ {
                                switch -274 {
                                case 494:
                                    v27 := GenS2{F0: "\n\t", F1: v9, F2: v7}
                                    _ = v27
                                default:
                                    fmt.Println(v14)
                                }
                                var v28 int8 = -42
                                if v9 {
                                    fmt.Println(v5)
                                    var v29 int8 = v28
                                    for fi30 := 0; fi30 < 1; fi30++ {
                                        v31 := GenS1{F0: (v3 ^ -138)}
                                        _ = fi30
                                        _ = v31
                                    }
                                    _ = v29
                                }
                                _ = fi26
                                _ = v28
                            }
                        } else {
                            var v32 int32 = 42
                            if v9 {
                                v2 -= (v2 + v2)
                            }
                            v14 = v6
                            _ = v32
                        }
                        switch 4294967295 {
                        case -907:
                            var v33 int64 = 9223372036854775807
                            _ = v33
                        case 1:
                            var v34 uint16 = 0
                            _ = v34
                        }
                        var v35 int16 = v3
                        _ = v35
                    }
                    v36 := GenS3{F0: 2, F1: 1, F2: 42}
                    v37 := GenS4{F0: (v3 ^ v3), F1: -1}
                    _ = v36
                    _ = v37
                }
            }
        }
        v7 = v7
        _ = v14
    }
    var v38 uint32 = 2147483647
    if (!true) {
        fmt.Println(v38)
    } else {
        switch -200 {
        case 9223372036854775807:
            if false {
                v1 = append(v1, (v38 % (((v38 >> ((65535) & 63))) | 1)))
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v3 = -266
                for fi39 := 0; fi39 < 14; fi39++ {
                    switch v7 {
                    case "foo bar":
                        var v40 int64 = (-(v6))
                        _ = v40
                    case "hello":
                        for fi41 := 0; fi41 < 7; fi41++ {
                            var v42 float64 = (1.5 / 3.14159)
                            if len(v1) > 0 {
                                v1[0] = 88
                            }
                            switch "\n\t" {
                            case "a":
                                switch (fi41 - fi41) {
                                case 1:
                                    v4 += v6
                                case -2147483648:
                                    if (65535 == 255) {
                                        v7 = ("\n\t" + "\n\t")
                                    } else {
                                        v38 -= 1
                                        var v43 int = 2147483647
                                        _ = v43
                                    }
                                case 32767:
                                    v8 = 848
                                default:
                                    if true {
                                        fmt.Println(v3, v6, v8)
                                    }
                                }
                            case "":
                                v42 += (1e10 - v42)
                            }
                            _ = fi41
                            _ = v42
                        }
                    default:
                        switch (fi39 >> ((v8) & 63)) {
                        case 128:
                            v44 := GenS4{F0: (v3 - v3), F1: -128}
                            _ = v44
                        case 317:
                            fmt.Println(v4, v6)
                        case 42:
                            v3 = (v3 * int16(v38))
                        }
                    }
                    for fi45 := 0; fi45 < 10; fi45++ {
                        var v46 uint16 = 1
                        var v47 string = ((v7 + v7) + ("hello" + ((v7 + "") + ((v7 + v7) + (v7 + ((v7 + v7) + (v7 + ("foo bar" + "foo bar"))))))))
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                    v48 := GenS2{F0: (v7 + ("" + (("foo bar" + (v7 + (v7 + (v7 + v7)))) + v7))), F1: (v9 && v9), F2: v7}
                    _ = fi39
                    _ = v48
                }
            }
        case 317:
            v38 += v38
        }
        v7 = v7
        v8 -= 4294967295
    }
    v2 = v2
    switch -1 {
    case 32768:
        var v49 float32 = v5
        _ = v49
    case -1:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 127:
        var v50 int8 = -57
        _ = v50
    default:
        v7 = v7
    }
    for fi51 := 0; fi51 < 5; fi51++ {
        v52 := GenS2{F0: ("" + (v7 + v7)), F1: (32767 <= 0), F2: ("foo bar" + (v7 + v7))}
        fmt.Println(v8, v9)
        _ = fi51
        _ = v52
    }
    v53 := GenS5{F0: 1, F1: -39, F2: 1}
    switch v7 {
    case "foo bar":
        _ = len(v1)
    default:
        var v54 int16 = int16(v6)
        _ = v54
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v53
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
