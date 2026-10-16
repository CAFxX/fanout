package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 int32; F2 int16 }
type GenS4 struct { F0 int64; F1 uint64; F2 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint32; F1 int }

func main() {
    var v1 int32 = -1
    v2 := make(map[uint32]uint32)
    v3 := []uint32{165, 453, 1, 4294967295}
    var v4 int64 = 9223372036854775807
    v5 := make(map[uint32]string)
    v6 := make(map[string]string)
    var v7 uint32 = 65535
    v8 := GenS1{F0: 2147483647, F1: 65535, F2: 21}
    if (false && (!(1 <= 128))) {
        var v9 int8 = -128
        var v10 uint64 = 604
        for fi11 := 0; fi11 < 14; fi11++ {
            v9 = 127
            _ = fi11
        }
        _ = v10
        _ = v9
    } else {
        delete(v2, v7)
    }
    if false {
        if false {
            v4 += 0
        } else {
            switch 597 {
            case 9223372036854775807:
                for fi12 := 0; fi12 < 18; fi12++ {
                    if false {
                        v13 := GenS2{F0: (v4 - int64(v7))}
                        _ = v13
                    }
                    v14 := GenS3{F0: 164, F1: (v1 << ((270) & 63)), F2: 1000}
                    _ = fi12
                    _ = v14
                }
            }
            var v15 int32 = (v1 << ((239) & 63))
            _ = v15
        }
    } else {
        if len(v3) > 0 {
            v3[0] = (v7 % ((2147483648) | 1))
        }
        _ = len(v3)
    }
    fmt.Println(v1, v7)
    if ((-(v4)) == (v4 ^ (v4 / (((v4 * (v4 - (-(v4))))) | 1)))) {
        for fi16 := 0; fi16 < 8; fi16++ {
            if false {
                fmt.Println(v7, fi16, fi16)
            } else {
                var v17 uint32 = v7
                _ = v17
            }
            fmt.Println(v1, v4)
            if false {
                fmt.Println(v4)
            } else {
                fmt.Println(v1, v4, fi16)
                for fi18 := 0; fi18 < 10; fi18++ {
                    v19 := GenS3{F0: 18446744073709551615, F1: (v1 % ((v1) | 1)), F2: 1000}
                    var v20 uint = 1
                    if false {
                        if (255 > 2) {
                            v3 = append(v3, 2)
                            v21 := GenS2{F0: v4}
                            _ = len(v3)
                            fmt.Println(fi18)
                            fmt.Println(v4, fi16)
                            v1 += (^(v1))
                            _ = v21
                        } else {
                            var v22 uint32 = v7
                            if true {
                                var v23 float32 = 1e10
                                var v24 uint16 = 1
                                _ = v23
                                _ = v24
                            }
                            var v25 uint = v20
                            var v26 float32 = (((0.0 + 0.1) / 0.1) * 0.1)
                            v27 := GenS1{F0: v7, F1: 482, F2: (v20 - v20)}
                            v28 := GenS3{F0: uint64(v25), F1: (v1 / ((v1) | 1)), F2: -32768}
                            var v29 bool = (!true)
                            _ = v22
                            _ = v25
                            _ = v26
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                        v30 := GenS4{F0: v4, F1: 0, F2: (^(v1))}
                        _ = v30
                    } else {
                        v31 := GenS5{F0: 173}
                        if len(v3) > 0 {
                            v3[0] = v7
                        }
                        v32 := GenS5{F0: 111}
                        _ = v31
                        _ = v32
                    }
                    if (!(284 <= 42)) {
                        switch "" {
                        case "a":
                            v33 := GenS4{F0: v4, F1: 4294967295, F2: (v1 >> (((^(v20))) & 63))}
                            if false {
                                if true {
                                    v34 := GenS6{F0: 522, F1: 2, F2: 1}
                                    switch (((("" + ("a" + "")) + "foo bar") + "hello") + ("a" + "hello")) {
                                    case "":
                                        var v35 float32 = 1e10
                                        _ = v35
                                    case "\n\t":
                                        fmt.Println(v1, v20)
                                    }
                                    v7 = (v7 | (v7 + v7))
                                    _ = v34
                                }
                            }
                            _ = v33
                        case "\n\t":
                            fmt.Println(fi18, fi16, fi18)
                        default:
                            v20 += v20
                        }
                    }
                    var v36 uint = 255
                    v1 = v1
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v36
                }
                var v37 string = "\n\t"
                _ = v37
            }
            _ = fi16
        }
    } else {
        if false {
            for fi38 := 0; fi38 < 19; fi38++ {
                var v39 uint8 = 95
                v7 = 2147483648
                _ = fi38
                _ = v39
            }
            var v40 int16 = -32768
            _ = v40
        }
    }
    for fi41 := 0; fi41 < 8; fi41++ {
        if false {
            var v42 int32 = (-(v1))
            v4 = v4
            _ = v42
        }
        var v43 int8 = -1
        _ = fi41
        _ = v43
    }
    for fi44 := 0; fi44 < 15; fi44++ {
        var v45 uint = 806
        if false {
            var v46 float64 = -2.25
            var v47 float32 = 1.5
            _ = v46
            _ = v47
        }
        v1 += (v1 - v1)
        _ = fi44
        _ = v45
    }
    for fi48 := 0; fi48 < 18; fi48++ {
        v49 := GenS3{F0: 180, F1: (v1 + v1), F2: 1}
        var v50 int = (fi48 / (((fi48 % ((1) | 1))) | 1))
        _ = fi48
        _ = v49
        _ = v50
    }
    fmt.Println(v4, v7, v7)
    v7 = (v7 << ((65535) & 63))
    v7 = v7
    if false {
        if (255 > 128) {
            v51 := GenS2{F0: (v4 - 1)}
            for fi52 := 0; fi52 < 9; fi52++ {
                v1 += v1
                v4 = -1
                _ = fi52
            }
            switch 42 {
            case -834:
                v53 := GenS7{F0: 53}
                _ = v53
            }
            _ = v51
        }
        v2[0] = v7
        v7 = (v7 % ((0) | 1))
    }
    switch ("foo bar" + (("hello" + "") + "")) {
    case "\n\t":
        v54 := GenS8{F0: (v7 & 2147483647), F1: 336}
        _ = v54
    case "hello":
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v4, v7)
        case "\n\t":
            fmt.Println(v7, v7)
        case "":
            for fi55 := 0; fi55 < 19; fi55++ {
                fmt.Println(fi55, v7, v1)
                _ = fi55
            }
        }
    case "":
        for fi56 := 0; fi56 < 6; fi56++ {
            var v57 bool = true
            var v58 float64 = float64((1e10 - 1e10))
            _ = fi56
            _ = v57
            _ = v58
        }
    default:
        if (100 == 127) {
            _ = len(v3)
            if false {
                if true {
                    if (true && false) {
                        fmt.Println(v4)
                        var v59 int64 = (-(v4))
                        if true {
                            v59 = (^(v59))
                            switch "hello" {
                            case "":
                                v4 = v4
                            case "foo bar":
                                var v60 float64 = ((1.5 - 1e10) + 0.0)
                                _ = v60
                            case "\n\t":
                                var v61 uint8 = 0
                                _ = v61
                            }
                        }
                        _ = v59
                    }
                } else {
                    var v62 uint16 = 65535
                    _ = v62
                }
                v4 = v4
            }
            fmt.Println(v1, v4)
        }
    }
    if ("foo bar" == "") {
        var v63 string = "foo bar"
        if (!true) {
            v3 = append(v3, (v7 + v7))
        }
        _ = v63
    } else {
        fmt.Println(v1, v1, v1)
    }
    _ = len(v3)
    fmt.Println(v1)
    switch 220 {
    case 876:
        v64 := v5[v7]
        _ = v64
    case -9223372036854775808:
        v5[v7] = (("" + ("" + ("\n\t" + "a"))) + "a")
    case 0:
        var v65 uint = 65535
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
