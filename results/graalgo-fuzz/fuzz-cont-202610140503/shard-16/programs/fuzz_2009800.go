package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 uint8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32; F1 uint16 }
type GenS8 struct { F0 uint64; F1 int64 }

func main() {
    var v1 uint = 71
    v2 := make(map[int]int)
    var v3 uint8 = 0
    v4 := []string{"a"}
    var v5 uint8 = (v3 | (v3 - (^(v3))))
    if (!(((^(v5)) == 127) || ("\n\t" <= ""))) {
        var v6 int = 128
        if true {
            v7 := GenS1{F0: 4294967295, F1: (v1 + (v1 - (v1 + 255))), F2: 42}
            for fi8 := 0; fi8 < 7; fi8++ {
                var v9 int32 = 1
                _ = fi8
                _ = v9
            }
            _ = v7
        } else {
            if (32767 < 255) {
                var v10 uint16 = 65535
                v11 := GenS2{F0: 100}
                _ = v10
                _ = v11
            } else {
                var v12 uint64 = 18446744073709551615
                _ = v12
            }
            var v13 bool = ((v3 < v5) && false)
            v14 := GenS1{F0: 2, F1: 683, F2: (^(v1))}
            _ = v13
            _ = v14
        }
        v15 := GenS2{F0: 1}
        _ = v15
        _ = v6
    }
    var v16 uint32 = 2147483648
    switch 2147483647 {
    case 0:
        v5 = v3
    default:
        v1 = (v1 & v1)
    }
    var v17 uint = 309
    var v18 int64 = -529
    for fi19 := 0; fi19 < 20; fi19++ {
        if (3.14159 <= (-2.25 * float64(float32((0.1 / 1e10))))) {
            if (!true) {
                switch (fi19 ^ (^(fi19))) {
                case 969:
                    v18 += v18
                case 2147483647:
                    var v20 uint8 = v5
                    _ = v20
                case -121:
                    v21 := GenS3{F0: 0, F1: (!(((((0.1 * 1e10) - 0.0) + 1.5) != 0.0) && (18446744073709551615 > 0)))}
                    _ = v21
                default:
                    fmt.Println(v18, v1)
                }
            }
            v22 := v2[(fi19 / ((fi19) | 1))]
            _ = v22
        } else {
            if ("a" != ((((("hello" + ("\n\t" + "foo bar")) + "a") + ("foo bar" + ((("foo bar" + "\n\t") + "a") + ""))) + ("a" + "hello")) + (("\n\t" + "hello") + "foo bar"))) {
                v23 := GenS2{F0: -1}
                _ = v23
            }
            switch ("a" + "hello") {
            case "hello":
                var v24 uint = (v1 >> (((^(v17))) & 63))
                _ = v24
            case "\n\t":
                var v25 bool = (480 >= 0)
                _ = v25
            case "":
                var v26 uint = 65535
                _ = v26
            }
            var v27 int32 = 0
            _ = v27
        }
        var v28 uint8 = (^(v3))
        _ = fi19
        _ = v28
    }
    v16 = (v16 & (v16 / ((v16) | 1)))
    for fi29 := 0; fi29 < 12; fi29++ {
        v30 := GenS4{F0: 1}
        switch -62 {
        case 127:
            switch fi29 {
            case 788:
                v5 += (v3 + v3)
            case 128:
                fmt.Println(v5)
            case 32768:
                if len(v4) > 0 {
                    v4[0] = "\n\t"
                }
            }
        case -397:
            v31 := GenS4{F0: 32767}
            _ = v31
        case -9223372036854775808:
            var v32 uint = (v1 + (v17 << ((42) & 63)))
            _ = v32
        }
        switch (fi29 - -2147483648) {
        case 42:
            var v33 uint16 = 2
            _ = v33
        case 511:
            if (!false) {
                v1 -= v17
                v16 = (v16 | (v16 | v16))
            } else {
                for fi34 := 0; fi34 < 8; fi34++ {
                    switch -1 {
                    case 1:
                        v35 := GenS3{F0: 432, F1: false}
                        _ = v35
                    }
                    fmt.Println(fi29, fi29, v17)
                    _ = fi34
                }
            }
        }
        _ = fi29
        _ = v30
    }
    v17 = 0
    v18 += v18
    if true {
        if (-1 >= 42) {
            if (!true) {
                for fi36 := 0; fi36 < 15; fi36++ {
                    for fi37 := 0; fi37 < 4; fi37++ {
                        if (true && (!(1.5 != (0.1 / 3.14159)))) {
                            v18 += v18
                            if false {
                                v38 := GenS1{F0: 0, F1: v17, F2: (v17 >> ((v17) & 63))}
                                if false {
                                    var v39 uint = v17
                                    switch "a" {
                                    case "foo bar":
                                        if false {
                                            if false {
                                                var v40 int32 = 0
                                                _ = v40
                                            } else {
                                                v18 -= (v18 % (((v18 ^ v18)) | 1))
                                            }
                                            var v41 int32 = 0
                                            _ = v41
                                        }
                                    default:
                                        var v42 int64 = (v18 * (v18 * (v18 - 42)))
                                        _ = v42
                                    }
                                    _ = v39
                                }
                                v18 = int64(v5)
                                _ = v38
                            }
                            v43 := GenS5{F0: (fi36 | fi36), F1: 84}
                            _ = v43
                        } else {
                            if (3.14159 >= (0.0 + 1.5)) {
                                var v44 uint32 = v16
                                fmt.Println(v1, v44)
                                _ = v44
                            }
                            _ = len(v2)
                        }
                        switch "foo bar" {
                        case "":
                            v17 = v1
                        case "hello":
                            v4 = append(v4, ("\n\t" + "\n\t"))
                        case "foo bar":
                            var v45 int = fi37
                            _ = v45
                        }
                        _ = fi37
                    }
                    _ = fi36
                }
            } else {
                v46 := GenS6{F0: 1}
                _ = v46
            }
            switch "" {
            case "foo bar":
                var v47 uint = v1
                _ = v47
            }
            v18 = 9223372036854775807
        }
    } else {
        v48 := GenS3{F0: 2, F1: false}
        _ = v48
    }
    var v49 string = "foo bar"
    v50 := GenS5{F0: 4294967295, F1: v3}
    if false {
        fmt.Println(v17, v49, v5)
        switch ("a" + (v49 + "")) {
        case "\n\t":
            var v51 int16 = 562
            _ = v51
        case "hello":
            v18 = v18
        default:
            if false {
                if true {
                    fmt.Println(v18, v5)
                    v1 += v17
                    v5 += 59
                } else {
                    var v52 uint = v1
                    for fi53 := 0; fi53 < 3; fi53++ {
                        fmt.Println(v18)
                        for fi54 := 0; fi54 < 3; fi54++ {
                            switch v49 {
                            case "hello":
                                v2[fi54] = fi54
                            case "\n\t":
                                v55 := GenS5{F0: fi53, F1: v3}
                                _ = v55
                            default:
                                if len(v4) > 0 {
                                    v4[0] = v49
                                }
                            }
                            var v56 uint = (v1 / ((1) | 1))
                            v57 := GenS6{F0: 1}
                            _ = fi54
                            _ = v56
                            _ = v57
                        }
                        _ = fi53
                    }
                    fmt.Println(v17, v16, v3)
                    _ = v52
                }
                for fi58 := 0; fi58 < 19; fi58++ {
                    var v59 uint32 = v16
                    fmt.Println(v16)
                    switch (v49 + v49) {
                    case "hello":
                        if true {
                            for fi60 := 0; fi60 < 5; fi60++ {
                                v61 := GenS7{F0: 2, F1: 65535}
                                v62 := GenS2{F0: 0}
                                v63 := v2[(^(fi60))]
                                _ = fi60
                                _ = v61
                                _ = v62
                                _ = v63
                            }
                            fmt.Println(v3, fi58)
                        } else {
                            var v64 int8 = 83
                            for fi65 := 0; fi65 < 10; fi65++ {
                                var v66 uint = v1
                                v18 -= -1
                                _ = fi65
                                _ = v66
                            }
                            _ = v64
                        }
                    default:
                        for fi67 := 0; fi67 < 1; fi67++ {
                            v68 := GenS8{F0: 1, F1: -645}
                            v2[(fi58 / (((-(fi67))) | 1))] = 2147483647
                            _ = fi67
                            _ = v68
                        }
                    }
                    _ = fi58
                    _ = v59
                }
                v18 = 314
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                for fi69 := 0; fi69 < 11; fi69++ {
                    var v70 string = (("a" + (v49 + v49)) + (("" + (v49 + v49)) + (("\n\t" + "hello") + "a")))
                    fmt.Println(v18, v5, v17)
                    var v71 uint16 = uint16(v17)
                    _ = fi69
                    _ = v70
                    _ = v71
                }
            }
        }
    } else {
        var v72 string = v49
        _ = v72
    }
    var v73 int64 = (v18 / ((v18) | 1))
    v18 = v18
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v73
}
