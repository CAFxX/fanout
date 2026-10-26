package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint; F2 int16 }

func main() {
    var v1 bool = (1 <= 128)
    v2 := []uint32{4294967295, 2, 2}
    var v3 float64 = (1e10 * 0.1)
    v4 := make(map[uint32]uint32)
    var v5 int64 = 42
    var v6 int8 = 0
    var v7 int16 = 32767
    fmt.Println(v7, v7, v6)
    v1 = v1
    if true {
        v8 := GenS1{F0: 940}
        var v9 int = 64
        _ = v8
        _ = v9
    }
    var v10 uint = 1
    v11 := GenS1{F0: 2}
    v5 -= v5
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := GenS1{F0: 0}
        _ = fi12
        _ = v13
    }
    var v14 uint8 = 2
    v4[65535] = 4294967295
    if (1 != 65535) {
        if (false && ((v14 | (v14 - (v14 / ((127) | 1)))) >= (^(v14)))) {
            for fi15 := 0; fi15 < 16; fi15++ {
                var v16 uint32 = 2
                _ = fi15
                _ = v16
            }
            fmt.Println(v3, v1, v14)
        } else {
            v3 += v3
        }
        v2 = append(v2, 65535)
    } else {
        if v1 {
            for fi17 := 0; fi17 < 6; fi17++ {
                v1 = v1
                var v18 float32 = 0.0
                _ = fi17
                _ = v18
            }
            var v19 int64 = (v5 % (((v5 + v5)) | 1))
            switch ("\n\t" + "") {
            case "\n\t":
                var v20 int8 = v6
                _ = v20
            case "a":
                v14 = 0
            }
            _ = v19
        } else {
            if (v7 >= 0) {
                for fi21 := 0; fi21 < 11; fi21++ {
                    if (128 == -868) {
                        var v22 uint = 687
                        v7 = (v7 >> (((v10 / ((v10) | 1))) & 63))
                        delete(v4, 65535)
                        _ = v22
                    } else {
                        var v23 int64 = (v5 & (v5 % ((v5) | 1)))
                        switch fi21 {
                        case 2147483647:
                            v5 -= -865
                        case 9223372036854775807:
                            v23 += v5
                        default:
                            var v24 int = int(v7)
                            _ = v24
                        }
                        _ = v23
                    }
                    _ = fi21
                }
                switch -9223372036854775808 {
                case 955:
                    v1 = ((0.0 / ((v3 + v3) * (0.1 / v3))) != -2.25)
                case -198:
                    fmt.Println(v7, v5, v7)
                case 0:
                    fmt.Println(v7, v5, v7)
                }
                switch 127 {
                case -691:
                    var v25 int32 = 1
                    _ = v25
                case 32767:
                    switch "foo bar" {
                    case "hello":
                        var v26 int8 = v6
                        _ = v26
                    default:
                        v27 := GenS1{F0: 0}
                        _ = v27
                    }
                default:
                    if (((1e10 * 0.0) * (-2.25 * 1e10)) < (0.0 * 1.5)) {
                        switch "" {
                        case "\n\t":
                            if false {
                                v1 = v1
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            } else {
                                var v28 uint = 0
                                _ = v28
                            }
                        }
                        for fi29 := 0; fi29 < 1; fi29++ {
                            switch -1 {
                            case 127:
                                for fi30 := 0; fi30 < 2; fi30++ {
                                    if (("" < "\n\t") || v1) {
                                        var v31 int = 42
                                        _ = v31
                                    }
                                    _ = fi30
                                }
                            case -866:
                                if (301 > (v10 & (v10 | v10))) {
                                    fmt.Println(v1, v3)
                                } else {
                                    for fi32 := 0; fi32 < 10; fi32++ {
                                        var v33 uint32 = 1
                                        fmt.Println(v1, fi32, v33)
                                        switch (fi32 >> ((v10) & 63)) {
                                        case 32768:
                                            var v34 int64 = (v5 >> ((65) & 63))
                                            _ = v34
                                        default:
                                            fmt.Println(fi29, v3)
                                        }
                                        _ = fi32
                                        _ = v33
                                    }
                                    v10 = (^(v10))
                                    fmt.Println(v10)
                                }
                            case 0:
                                fmt.Println(v5)
                            }
                            for fi35 := 0; fi35 < 11; fi35++ {
                                var v36 float64 = v3
                                _ = fi35
                                _ = v36
                            }
                            v6 = (v6 ^ v6)
                            _ = fi29
                        }
                    }
                }
            } else {
                for fi37 := 0; fi37 < 5; fi37++ {
                    for fi38 := 0; fi38 < 12; fi38++ {
                        _ = len(v2)
                        _ = fi38
                    }
                    _ = fi37
                }
                for fi39 := 0; fi39 < 2; fi39++ {
                    for fi40 := 0; fi40 < 12; fi40++ {
                        v14 = 0
                        fmt.Println(v10, v6)
                        _ = fi40
                    }
                    _ = fi39
                }
                v41 := GenS1{F0: 65535}
                _ = v41
            }
            var v42 float32 = 0.1
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v42
        }
        v1 = true
        v43 := GenS1{F0: 2}
        _ = v43
    }
    for fi44 := 0; fi44 < 6; fi44++ {
        var v45 float64 = (3.14159 + 0.0)
        v46 := GenS2{F0: (v6 & -71)}
        if false {
            var v47 float64 = v45
            _ = v47
        } else {
            var v48 int8 = int8(v14)
            v49 := GenS3{F0: v10, F1: (^(v10)), F2: 32767}
            _ = v48
            _ = v49
        }
        _ = fi44
        _ = v45
        _ = v46
    }
    if len(v2) > 0 {
        v2[0] = 0
    }
    _ = len(v4)
    if (1e10 > -2.25) {
        if v1 {
            v1 = v1
        } else {
            var v50 float64 = v3
            _ = v50
        }
    } else {
        fmt.Println(v1, v7)
        v6 -= int8(v14)
        _ = len(v2)
    }
    v51 := v4[2147483648]
    switch ("a" + "") {
    case "hello":
        var v52 int32 = 1
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
    _ = v7
}
