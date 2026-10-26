package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 int16; F2 int64 }

func main() {
    v1 := []uint32{65535}
    v2 := []uint32{2147483648, 542}
    var v3 bool = (1.5 != ((1.5 + float32(0.1)) - -2.25))
    var v4 int64 = -1
    var v5 float32 = 1e10
    v4 -= v4
    switch 933 {
    case -2147483648:
        var v6 string = "foo bar"
        _ = v6
    case 547:
        var v7 uint64 = 42
        _ = v7
    case -177:
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 string = "foo bar"
            var v10 uint64 = 42
            _ = fi8
            _ = v10
            _ = v9
        }
    default:
        var v11 float64 = 0.1
        _ = v11
    }
    v2 = append(v2, 887)
    var v12 float64 = 1.5
    if v3 {
        var v13 uint32 = 291
        v14 := GenS1{F0: 18446744073709551615}
        _ = v13
        _ = v14
    } else {
        if len(v1) > 0 {
            v1[0] = 4294967295
        }
        v12 = float64(v5)
        for fi15 := 0; fi15 < 12; fi15++ {
            fmt.Println(v4)
            _ = fi15
        }
    }
    if true {
        v3 = true
        var v16 uint8 = 215
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v16
    } else {
        v3 = ((v3 && v3) && v3)
        if true {
            fmt.Println(v12)
            var v17 int64 = 1
            _ = v17
        } else {
            fmt.Println(v4, v4, v4)
            var v18 bool = ((false && true) || v3)
            for fi19 := 0; fi19 < 20; fi19++ {
                var v20 int16 = 1
                _ = len(v2)
                _ = fi19
                _ = v20
            }
            _ = v18
        }
    }
    var v21 float32 = v5
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 431 {
    case 42:
        var v22 int32 = -871
        _ = v22
    case 128:
        v3 = v3
    case -9223372036854775808:
        fmt.Println(v4)
    }
    var v23 uint8 = 1
    v4 = v4
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    fmt.Println(v4, v5, v21)
    if v3 {
        v24 := GenS1{F0: 948}
        _ = v24
    } else {
        v23 += (^(v23))
        v21 = ((0.0 + v21) - v5)
        var v25 uint64 = 65
        _ = v25
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        if false {
            v27 := GenS1{F0: 18446744073709551615}
            v28 := GenS1{F0: 622}
            _ = v27
            _ = v28
        }
        for fi29 := 0; fi29 < 20; fi29++ {
            v5 = (3.14159 - -2.25)
            v4 -= v4
            v30 := GenS1{F0: 42}
            if v3 {
                if v3 {
                    fmt.Println(fi29, fi29, v23)
                    var v31 uint64 = 18446744073709551615
                    var v32 int32 = -890
                    _ = v31
                    _ = v32
                } else {
                    if v3 {
                        if v3 {
                            switch "a" {
                            case "a":
                                if v3 {
                                    v33 := GenS2{F0: ""}
                                    _ = v33
                                }
                            default:
                                var v34 string = "hello"
                                _ = v34
                            }
                        }
                        fmt.Println(fi29)
                        var v35 float64 = v12
                        _ = v35
                    } else {
                        switch (fi26 / ((fi29) | 1)) {
                        case 9223372036854775807:
                            if true {
                                v36 := GenS3{F0: 287}
                                v37 := GenS3{F0: 18446744073709551615}
                                fmt.Println(v5, v5)
                                _ = v36
                                _ = v37
                            } else {
                                if (0 < 18446744073709551615) {
                                    v2 = append(v2, 0)
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v2 = append(v2, 42)
                                    if (356 != 931) {
                                        v38 := GenS4{F0: (fi29 % (((-(fi26))) | 1)), F1: 42, F2: -256}
                                        fmt.Println(v12, v21)
                                        _ = v38
                                    }
                                }
                            }
                        case -488:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        v39 := GenS4{F0: (fi29 << ((65535) & 63)), F1: 42, F2: 73}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v39
                    }
                    fmt.Println(v5, fi29, v23)
                }
                _ = len(v2)
            } else {
                switch ("\n\t" + "\n\t") {
                case "\n\t":
                    v40 := GenS2{F0: "\n\t"}
                    _ = v40
                }
                var v41 float32 = float32((1.5 + (v12 - v12)))
                _ = v41
            }
            _ = fi29
            _ = v30
        }
        _ = fi26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
