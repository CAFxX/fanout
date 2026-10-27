package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int64 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint16; F1 int; F2 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 int64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{("foo bar" + "\n\t")}
    var v3 int16 = 0
    v4 := make(map[int]int)
    if ((62 < 42) || (!(int32(v3) >= 2147483647))) {
        var v5 int8 = 25
        _ = v5
    }
    var v6 int = 32768
    if (v6 != v6) {
        var v7 uint = 42
        switch v6 {
        case 1:
            v7 -= 18446744073709551615
        case -2147483648:
            var v8 int32 = 42
            _ = v8
        }
        _ = v7
    }
    var v9 float64 = float64((((3.14159 + 3.14159) - 1e10) / (1e10 - 1.5)))
    _ = len(v4)
    var v10 uint64 = 2
    var v11 uint32 = 4294967295
    switch (v6 | v6) {
    case -414:
        for fi12 := 0; fi12 < 11; fi12++ {
            v6 = (v6 | (v6 + fi12))
            var v13 int64 = 42
            _ = fi12
            _ = v13
        }
    case -2147483648:
        v14 := GenS1{F0: 0, F1: (v11 - uint32(v10)), F2: -346}
        _ = v14
    default:
        if ((!(v9 == v9)) && true) {
            if false {
                for fi15 := 0; fi15 < 18; fi15++ {
                    if false {
                        switch "" {
                        case "\n\t":
                            var v16 float64 = v9
                            _ = v16
                        case "foo bar":
                            var v17 uint16 = 42
                            _ = v17
                        }
                    } else {
                        var v18 int32 = 1
                        for fi19 := 0; fi19 < 4; fi19++ {
                            var v20 uint64 = v10
                            var v21 bool = false
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                        var v22 string = "a"
                        _ = v18
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 14; fi23++ {
                        var v24 int64 = -634
                        fmt.Println(v6, v10, v9)
                        _ = len(v4)
                        _ = fi23
                        _ = v24
                    }
                    if (2 <= 847) {
                        for fi25 := 0; fi25 < 2; fi25++ {
                            v26 := GenS2{F0: fi25, F1: ""}
                            _ = fi25
                            _ = v26
                        }
                        v10 -= (v10 >> ((582) & 63))
                        v10 -= (v10 >> ((490) & 63))
                    }
                    _ = fi15
                }
                for fi27 := 0; fi27 < 16; fi27++ {
                    v28 := GenS3{F0: 684, F1: (fi27 / (((v6 | (fi27 << ((0) & 63)))) | 1)), F2: v3}
                    _ = len(v4)
                    v2 = append(v2, ("hello" + "hello"))
                    _ = fi27
                    _ = v28
                }
            }
            if ((1 < -1) || (((1.5 / v9) * v9) != ((v9 / (1e10 / 1e10)) * (v9 + v9)))) {
                for fi29 := 0; fi29 < 6; fi29++ {
                    var v30 int8 = 100
                    _ = fi29
                    _ = v30
                }
                v9 = 3.14159
            }
        }
    }
    v31 := GenS3{F0: 873, F1: v6, F2: (^(v3))}
    for fi32 := 0; fi32 < 16; fi32++ {
        for fi33 := 0; fi33 < 17; fi33++ {
            v9 -= v9
            var v34 int = 1
            switch 418 {
            case -700:
                if false {
                    switch (v34 << ((42) & 63)) {
                    case 9223372036854775807:
                        v2 = append(v2, ("hello" + ("foo bar" + "\n\t")))
                    case -9223372036854775808:
                        fmt.Println(v9, v10)
                    case 32767:
                        switch v6 {
                        case 127:
                            v10 += 4294967295
                        default:
                            v6 = v6
                        }
                    default:
                        _ = len(v2)
                    }
                    v10 += 1
                    v9 = v9
                    v9 -= (0.0 + 0.0)
                    v35 := v1[(fi33 * (v34 ^ fi32))]
                    _ = v35
                } else {
                    var v36 uint = 0
                    v11 += (v11 + v11)
                    switch (-(fi33)) {
                    case 0:
                        fmt.Println(v3)
                    }
                    var v37 int8 = 1
                    v10 -= 4294967295
                    _ = v36
                    _ = v37
                }
            case 147:
                v38 := v4[(v6 - 127)]
                _ = v38
            case 94:
                fmt.Println(v9)
            default:
                v39 := GenS4{F0: 31}
                _ = v39
            }
            _ = fi33
            _ = v34
        }
        v40 := GenS3{F0: 0, F1: 493, F2: (v3 ^ -666)}
        _ = fi32
        _ = v40
    }
    var v41 float32 = 0.1
    var v42 int16 = v3
    var v43 uint8 = 1
    var v44 uint16 = 0
    if (!false) {
        switch v6 {
        case 1:
            var v45 uint64 = uint64(v44)
            _ = v45
        case -1:
            v46 := GenS2{F0: v6, F1: ("" + "foo bar")}
            _ = v46
        case -442:
            if false {
                v47 := GenS1{F0: -9223372036854775808, F1: v11, F2: -9223372036854775808}
                var v48 uint = 4294967295
                _ = v47
                _ = v48
            }
        default:
            if (!false) {
                v49 := GenS4{F0: (v43 / (((v43 - 255)) | 1))}
                _ = len(v1)
                _ = v49
            } else {
                if false {
                    var v50 int8 = 0
                    _ = v50
                } else {
                    var v51 uint16 = 2
                    if (!((v9 / ((v9 * v9) / v9)) >= v9)) {
                        v52 := GenS2{F0: 9223372036854775807, F1: "\n\t"}
                        _ = v52
                    } else {
                        for fi53 := 0; fi53 < 6; fi53++ {
                            var v54 int = (-(v6))
                            _ = fi53
                            _ = v54
                        }
                        for fi55 := 0; fi55 < 12; fi55++ {
                            fmt.Println(v6)
                            for fi56 := 0; fi56 < 18; fi56++ {
                                switch -89 {
                                case 32768:
                                    v57 := GenS5{F0: (v42 / ((v42) | 1)), F1: 1}
                                    var v58 uint32 = (v11 + v11)
                                    if false {
                                        var v59 int = -2147483648
                                        var v60 float32 = 0.1
                                        _ = v59
                                        _ = v60
                                    } else {
                                        v44 += (v51 + (^(v51)))
                                        v58 += v58
                                        var v61 bool = (false && true)
                                        _ = v61
                                    }
                                    _ = v57
                                    _ = v58
                                }
                                var v62 int = -666
                                _ = fi56
                                _ = v62
                            }
                            _ = fi55
                        }
                    }
                    v63 := GenS4{F0: (v43 << ((4294967295) & 63))}
                    _ = v51
                    _ = v63
                }
                v64 := GenS3{F0: (v44 >> ((1) & 63)), F1: v6, F2: (v42 >> ((111) & 63))}
                _ = v64
            }
        }
        var v65 int16 = v3
        switch ("foo bar" + ("a" + ("" + "a"))) {
        case "foo bar":
            v65 -= v42
        default:
            v6 -= v6
        }
        _ = v65
    } else {
        v6 = (v6 << ((364) & 63))
        v66 := GenS1{F0: -885, F1: v11, F2: -1}
        var v67 uint64 = (^(v10))
        _ = v66
        _ = v67
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v6
    _ = v9
}
