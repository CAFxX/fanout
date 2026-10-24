package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int32 }
type GenS2 struct { F0 int64; F1 bool; F2 uint16 }
type GenS3 struct { F0 int64; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint64; F1 uint; F2 int32 }
type GenS5 struct { F0 int; F1 uint; F2 uint8 }
type GenS6 struct { F0 uint16; F1 int16 }
type GenS7 struct { F0 uint32 }

func main() {
    var v1 uint64 = 0
    var v2 float32 = float32(-2.25)
    var v3 int32 = -583
    v4 := []int{128, 4294967295, -496}
    var v5 int = -64
    v6 := make(map[string]int)
    v7 := GenS1{F0: 42, F1: 0, F2: v3}
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 int32 = (v3 + v3)
        if ((^(v5)) <= fi8) {
            v2 += v2
            for fi10 := 0; fi10 < 4; fi10++ {
                var v11 bool = ((v5 >= fi8) && true)
                switch (fi8 * fi8) {
                case 0:
                    var v12 uint32 = 65535
                    _ = v12
                case 539:
                    if v11 {
                        fmt.Println(v1, fi8)
                    } else {
                        var v13 float64 = 1e10
                        var v14 uint = 122
                        switch (^(v5)) {
                        case -9223372036854775808:
                            var v15 int8 = 71
                            _ = v15
                        }
                        v16 := GenS1{F0: 128, F1: 313, F2: (^(v9))}
                        _ = v13
                        _ = v14
                        _ = v16
                    }
                }
                _ = fi10
                _ = v11
            }
            v6[(("a" + "foo bar") + "\n\t")] = -1
        } else {
            v6["hello"] = (-(fi8))
        }
        fmt.Println(v2, v5, v1)
        _ = fi8
        _ = v9
    }
    if true {
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 uint = 2
            v3 = -2147483648
            var v19 uint32 = 2
            _ = fi17
            _ = v18
            _ = v19
        }
    } else {
        fmt.Println(v2, v3)
        var v20 float64 = 3.14159
        var v21 int16 = int16(v5)
        _ = v20
        _ = v21
    }
    var v22 uint16 = 42
    var v23 uint32 = 420
    for fi24 := 0; fi24 < 11; fi24++ {
        for fi25 := 0; fi25 < 5; fi25++ {
            if (true || true) {
                var v26 int32 = -704
                _ = v26
            } else {
                var v27 int32 = (v3 << ((18446744073709551615) & 63))
                _ = v27
            }
            _ = fi25
        }
        fmt.Println(v5, v23)
        v3 = (v3 >> ((0) & 63))
        _ = fi24
    }
    switch "foo bar" {
    case "hello":
        switch ("" + "") {
        case "hello":
            var v28 float64 = float64(0.1)
            _ = v28
        case "a":
            v2 = (1.5 / (((v2 * 0.0) / 1.0) * v2))
        case "\n\t":
            if true {
                v4 = append(v4, v5)
                v29 := GenS1{F0: v5, F1: v22, F2: (v3 * (^(v3)))}
                _ = v29
            } else {
                var v30 int = v5
                if true {
                    if true {
                        for fi31 := 0; fi31 < 18; fi31++ {
                            switch (v5 + (v30 << ((771) & 63))) {
                            case -771:
                                var v32 uint = uint(v1)
                                _ = v32
                            }
                            _ = fi31
                        }
                        var v33 string = ("\n\t" + "foo bar")
                        var v34 string = v33
                        _ = v33
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 19; fi35++ {
                        v36 := GenS2{F0: 349, F1: true, F2: 379}
                        switch (v30 * fi35) {
                        case 1:
                            switch (("a" + ("foo bar" + "\n\t")) + "\n\t") {
                            case "foo bar":
                                if false {
                                    for fi37 := 0; fi37 < 6; fi37++ {
                                        v3 -= (v3 % ((v3) | 1))
                                        v38 := GenS3{F0: -9223372036854775808, F1: 0, F2: 199}
                                        _ = fi37
                                        _ = v38
                                    }
                                    var v39 uint = 2
                                    v3 = v3
                                    _ = v39
                                }
                            case "":
                                v40 := v6[("a" + "a")]
                                _ = v40
                            default:
                                v5 = 42
                            }
                        case 978:
                            delete(v6, (("\n\t" + ("a" + ("foo bar" + "a"))) + (("a" + ("\n\t" + ("" + ("foo bar" + "")))) + "")))
                        }
                        _ = fi35
                        _ = v36
                    }
                } else {
                    var v41 int8 = 89
                    var v42 uint16 = v22
                    if true {
                        if true {
                            var v43 int8 = -9
                            for fi44 := 0; fi44 < 2; fi44++ {
                                var v45 int8 = v43
                                delete(v6, "")
                                v30 = 127
                                _ = fi44
                                _ = v45
                            }
                            _ = v43
                        } else {
                            v46 := GenS4{F0: (v1 ^ (v1 - (^(v1)))), F1: 65535, F2: v3}
                            switch "a" {
                            case "\n\t":
                                v47 := GenS4{F0: v1, F1: 4294967295, F2: (v3 / ((0) | 1))}
                                _ = v47
                            case "":
                                var v48 int32 = (v3 << ((313) & 63))
                                _ = v48
                            }
                            _ = v46
                        }
                        switch 817 {
                        case 127:
                            var v49 int64 = 9223372036854775807
                            _ = v49
                        default:
                            var v50 uint16 = (v22 - (v42 * (^(v42))))
                            _ = v50
                        }
                        if len(v4) > 0 {
                            v4[0] = v5
                        }
                    }
                    _ = v41
                    _ = v42
                }
                switch "a" {
                case "foo bar":
                    if true {
                        v51 := GenS5{F0: (v5 / (((v30 & v5)) | 1)), F1: 18446744073709551615, F2: 239}
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v52 float32 = v2
                        _ = v51
                        _ = v52
                    } else {
                        _ = len(v4)
                        var v53 int64 = 9223372036854775807
                        _ = v53
                    }
                }
                _ = v30
            }
        }
    case "foo bar":
        for fi54 := 0; fi54 < 16; fi54++ {
            fmt.Println(v2, v2, v1)
            fmt.Println(v2, v1)
            switch -829 {
            case 872:
                v2 = v2
            case -2147483648:
                v2 += v2
            }
            _ = fi54
        }
    default:
        var v55 int8 = -128
        _ = v55
    }
    if true {
        switch int(v22) {
        case 32767:
            fmt.Println(v3, v22)
        case 42:
            for fi56 := 0; fi56 < 2; fi56++ {
                var v57 int32 = (v3 >> ((2) & 63))
                _ = fi56
                _ = v57
            }
        case 662:
            fmt.Println(v1)
        }
    }
    switch (v5 * 127) {
    case 127:
        for fi58 := 0; fi58 < 10; fi58++ {
            var v59 int8 = -1
            _ = fi58
            _ = v59
        }
    case 443:
        v60 := GenS6{F0: (v22 / ((837) | 1)), F1: 0}
        _ = v60
    case 350:
        v1 = v1
    }
    _ = len(v4)
    var v61 string = (("" + "hello") + (("foo bar" + ("a" + ((("" + "") + "") + ""))) + "\n\t"))
    for fi62 := 0; fi62 < 20; fi62++ {
        fmt.Println(fi62)
        v6[v61] = v5
        var v63 bool = (((v23 << ((325) & 63)) != 1) && false)
        _ = fi62
        _ = v63
    }
    if ((v2 / v2) < v2) {
        if len(v4) > 0 {
            v4[0] = 9223372036854775807
        }
        delete(v6, v61)
        v23 -= 4294967295
    }
    var v64 string = v61
    for fi65 := 0; fi65 < 16; fi65++ {
        _ = len(v4)
        if false {
            v66 := GenS4{F0: (v1 % ((18446744073709551615) | 1)), F1: 288, F2: (v3 + v3)}
            v67 := GenS3{F0: 42, F1: 0, F2: 237}
            switch fi65 {
            case 32767:
                v23 += v23
            case -2147483648:
                fmt.Println(v23, v61)
            case -1:
                v2 = v2
            }
            _ = v66
            _ = v67
        } else {
            var v68 string = "hello"
            v69 := GenS7{F0: (^(v23))}
            var v70 int = v5
            _ = v68
            _ = v69
            _ = v70
        }
        var v71 uint64 = (v1 / ((uint64(v23)) | 1))
        _ = fi65
        _ = v71
    }
    v72 := GenS2{F0: -1, F1: true, F2: (v22 >> ((2) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v61:", v61)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v64
    _ = v7
    _ = v72
}
