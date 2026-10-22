package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 uint64; F2 int8 }
type GenS3 struct { F0 int; F1 int; F2 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 int32; F2 int32 }

func main() {
    var v1 uint = 4294967295
    var v2 uint64 = 42
    v3 := []uint32{219, 2147483647, 875, 791}
    fmt.Println(v2, v1)
    v4 := GenS1{F0: true}
    v3 = append(v3, 0)
    v3 = append(v3, 2147483648)
    if (!((25 < 42) && true)) {
        var v5 uint8 = 128
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v2 -= v2
        _ = v5
    } else {
        v6 := GenS1{F0: (-661 == -1)}
        v2 -= 42
        v7 := GenS2{F0: 4294967295, F1: v2, F2: -91}
        _ = v6
        _ = v7
    }
    v3 = append(v3, 848)
    v3 = append(v3, 4294967295)
    var v8 float32 = (-2.25 + 0.1)
    for fi9 := 0; fi9 < 20; fi9++ {
        v3 = append(v3, uint32(v2))
        v1 += (^(v1))
        _ = fi9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        for fi11 := 0; fi11 < 10; fi11++ {
            fmt.Println(v2, fi11, v2)
            var v12 int = 686
            _ = fi11
            _ = v12
        }
        var v13 string = ""
        var v14 uint32 = 1
        _ = fi10
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        switch "" {
        case "":
            _ = len(v3)
        case "\n\t":
            switch "hello" {
            case "foo bar":
                for fi16 := 0; fi16 < 14; fi16++ {
                    _ = len(v3)
                    switch (fi16 << (((v1 >> ((v1) & 63))) & 63)) {
                    case -2147483648:
                        fmt.Println(v1, v2)
                    case 128:
                        fmt.Println(v8, v8)
                    case -878:
                        v1 += 65535
                    }
                    _ = fi16
                }
            case "a":
                var v17 float64 = (0.0 * 0.0)
                _ = v17
            default:
                for fi18 := 0; fi18 < 9; fi18++ {
                    var v19 int16 = -32768
                    _ = fi18
                    _ = v19
                }
            }
        }
        _ = fi15
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        v2 += (v2 + v2)
        _ = len(v3)
        v21 := GenS2{F0: (fi20 >> ((255) & 63)), F1: v2, F2: 0}
        _ = fi20
        _ = v21
    }
    v1 = v1
    var v22 uint32 = 4294967295
    v23 := GenS3{F0: 32767, F1: 2147483647, F2: 255}
    if (v1 < 42) {
        var v24 int16 = 574
        _ = v24
    } else {
        if (v2 == v2) {
            switch ("hello" + ("a" + ("hello" + "\n\t"))) {
            case "foo bar":
                if (65535 <= v1) {
                    var v25 uint32 = v22
                    _ = v25
                } else {
                    if (32767 != 128) {
                        if false {
                            for fi26 := 0; fi26 < 4; fi26++ {
                                if (1 == v1) {
                                    var v27 uint32 = (^(v22))
                                    var v28 int = 9223372036854775807
                                    v22 = v27
                                    _ = v27
                                    _ = v28
                                } else {
                                    var v29 uint = 0
                                    _ = v29
                                }
                                switch fi26 {
                                case -2147483648:
                                    for fi30 := 0; fi30 < 12; fi30++ {
                                        v2 = uint64(v1)
                                        fmt.Println(v2, fi30)
                                        if false {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v31 := GenS1{F0: (!false)}
                                            _ = v31
                                        } else {
                                            if false {
                                                var v32 float64 = 1e10
                                                var v33 uint64 = 1
                                                _ = v32
                                                _ = v33
                                            } else {
                                                fmt.Println(fi26, v2, fi26)
                                                var v34 float32 = 0.1
                                                var v35 int32 = 252
                                                _ = v34
                                                _ = v35
                                            }
                                            v1 = v1
                                        }
                                        _ = fi30
                                    }
                                }
                                var v36 uint8 = 249
                                _ = fi26
                                _ = v36
                            }
                        } else {
                            v3 = append(v3, (^(v22)))
                            _ = len(v3)
                        }
                        if false {
                            var v37 uint = v1
                            var v38 uint = v1
                            _ = v37
                            _ = v38
                        }
                    } else {
                        var v39 string = ("a" + "\n\t")
                        var v40 uint32 = (v22 * v22)
                        var v41 bool = true
                        _ = v39
                        _ = v40
                        _ = v41
                    }
                }
            case "\n\t":
                v42 := GenS4{F0: 0}
                _ = v42
            case "hello":
                v43 := GenS2{F0: 1, F1: (v2 / ((v2) | 1)), F2: 0}
                _ = v43
            }
        } else {
            for fi44 := 0; fi44 < 11; fi44++ {
                for fi45 := 0; fi45 < 6; fi45++ {
                    switch "hello" {
                    case "a":
                        switch "\n\t" {
                        case "\n\t":
                            if (0.0 < 1.5) {
                                fmt.Println(v22)
                                switch (fi45 ^ fi44) {
                                case -679:
                                    fmt.Println(v22, v1)
                                case 0:
                                    v46 := GenS2{F0: 42, F1: (v2 & (^(v2))), F2: 127}
                                    _ = v46
                                case 1:
                                    if false {
                                        var v47 float32 = v8
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v47
                                    } else {
                                        _ = len(v3)
                                    }
                                default:
                                    fmt.Println(fi45, v8, fi44)
                                }
                            }
                        default:
                            v8 += (v8 * float32(float64(0.0)))
                        }
                    case "foo bar":
                        var v48 int16 = -512
                        _ = v48
                    case "":
                        v1 -= v1
                    default:
                        var v49 uint32 = (v22 % ((v22) | 1))
                        _ = v49
                    }
                    _ = fi45
                }
                _ = fi44
            }
            v50 := GenS2{F0: 32768, F1: 1, F2: 21}
            if false {
                if ((((2147483647 > 42) && (false || (-128 >= -110))) && ((v2 >> ((v1) & 63)) != 4294967295)) && true) {
                    fmt.Println(v8)
                    for fi51 := 0; fi51 < 15; fi51++ {
                        var v52 int64 = 0
                        v52 = (v52 % ((0) | 1))
                        _ = fi51
                        _ = v52
                    }
                } else {
                    var v53 int = 679
                    var v54 uint16 = 42
                    if false {
                        for fi55 := 0; fi55 < 9; fi55++ {
                            if false {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            if (true && true) {
                                if len(v3) > 0 {
                                    v3[0] = 1
                                }
                                for fi56 := 0; fi56 < 10; fi56++ {
                                    if ((v53 <= 127) && ((!true) || (v53 <= (^(fi56))))) {
                                        v54 = (v54 % ((359) | 1))
                                        v22 -= 42
                                        v57 := GenS5{F0: false, F1: int32(v22), F2: int32(v53)}
                                        _ = v57
                                    }
                                    if (596 == 1) {
                                        v58 := GenS2{F0: fi55, F1: v2, F2: -59}
                                        _ = v58
                                    }
                                    var v59 int16 = 239
                                    v60 := GenS1{F0: true}
                                    _ = fi56
                                    _ = v59
                                    _ = v60
                                }
                            }
                            _ = fi55
                        }
                    }
                    _ = v53
                    _ = v54
                }
            } else {
                var v61 int = 1
                for fi62 := 0; fi62 < 20; fi62++ {
                    if false {
                        var v63 int32 = 2147483647
                        fmt.Println(v61, v22, v61)
                        var v64 float64 = 3.14159
                        _ = v63
                        _ = v64
                    }
                    if len(v3) > 0 {
                        v3[0] = v22
                    }
                    _ = fi62
                }
                _ = v61
            }
            _ = v50
        }
        fmt.Println(v1, v22, v8)
    }
    fmt.Println(v1, v2)
    var v65 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v65:", v65)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v65
    _ = v8
}
