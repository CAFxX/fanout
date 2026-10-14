package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 string }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64 }

func main() {
    var v1 uint64 = 4294967295
    v2 := []uint32{}
    var v3 string = "a"
    v4 := make(map[int]int)
    v5 := make(map[int]uint32)
    v1 += v1
    if true {
        switch 32767 {
        case 0:
            for fi6 := 0; fi6 < 4; fi6++ {
                if ((true && false) && (true && false)) {
                    fmt.Println(v1, fi6)
                }
                if len(v2) > 0 {
                    v2[0] = 2147483647
                }
                _ = fi6
            }
        default:
            if (((!false) || (int64(v1) > -1)) || false) {
                v7 := GenS1{F0: 255, F1: 127}
                _ = v7
            }
        }
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        if false {
            v9 := GenS2{F0: 42, F1: 4294967295, F2: v3}
            var v10 int16 = -1
            _ = v10
            _ = v9
        } else {
            for fi11 := 0; fi11 < 8; fi11++ {
                switch fi8 {
                case 128:
                    v1 += (^(v1))
                case 116:
                    if (100 <= 28) {
                        if (false && false) {
                            if (int64(v1) != 126) {
                                v12 := GenS3{F0: 2, F1: -429}
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v12
                            }
                            var v13 int64 = 9223372036854775807
                            if false {
                                var v14 uint16 = 77
                                v13 = (v13 << ((4294967295) & 63))
                                switch v3 {
                                case "a":
                                    var v15 uint = 2
                                    _ = v15
                                case "":
                                    if (97 < 1) {
                                        v14 = v14
                                    }
                                case "\n\t":
                                    var v16 float64 = (3.14159 - 0.0)
                                    _ = v16
                                default:
                                    if false {
                                        fmt.Println(v3)
                                    } else {
                                        v14 -= 255
                                        var v17 int32 = -109
                                        var v18 int32 = v17
                                        _ = v17
                                        _ = v18
                                    }
                                }
                                _ = v14
                            } else {
                                var v19 uint8 = 128
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v19
                            }
                            _ = v13
                        }
                        v20 := GenS4{F0: 9223372036854775807}
                        _ = v20
                    } else {
                        if true {
                            if false {
                                v21 := GenS3{F0: 2147483647, F1: -2147483648}
                                v3 = (v3 + "foo bar")
                                if len(v2) > 0 {
                                    v2[0] = 0
                                }
                                _ = v21
                            }
                            var v22 uint = 255
                            var v23 int8 = -1
                            _ = v22
                            _ = v23
                        } else {
                            var v24 uint32 = 4294967295
                            if (true || false) {
                                var v25 uint8 = 42
                                switch (v3 + (v3 + "\n\t")) {
                                case "hello":
                                    if (!(!true)) {
                                        v1 = (v1 / ((978) | 1))
                                    } else {
                                        _ = len(v5)
                                        if true {
                                            v3 = v3
                                        } else {
                                            fmt.Println(v25, fi11, v3)
                                            switch (fi11 + int(v24)) {
                                            case 9223372036854775807:
                                                if false {
                                                    fmt.Println(v24, fi11)
                                                    v26 := GenS4{F0: 9223372036854775807}
                                                    v27 := GenS2{F0: uint32(v25), F1: (v24 >> ((18446744073709551615) & 63)), F2: v3}
                                                    _ = v26
                                                    _ = v27
                                                }
                                            case 639:
                                                if true {
                                                    if false {
                                                        var v28 int = (-(fi11))
                                                        _ = len(v4)
                                                        fmt.Println(v24)
                                                        _ = v28
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    var v29 int = (fi11 & (fi8 / ((fi8) | 1)))
                                                    _ = len(v2)
                                                    fmt.Println(fi8)
                                                    _ = v29
                                                } else {
                                                    v3 = "hello"
                                                    fmt.Println(fi8)
                                                }
                                            }
                                        }
                                    }
                                case "foo bar":
                                    var v30 float64 = 0.0
                                    _ = v30
                                }
                                v3 = v3
                                _ = v25
                            }
                            v31 := GenS5{F0: 42}
                            _ = v24
                            _ = v31
                        }
                    }
                case 1:
                    v3 = ((v3 + (v3 + v3)) + v3)
                }
                v32 := GenS3{F0: 1, F1: -2147483648}
                fmt.Println(v3)
                _ = fi11
                _ = v32
            }
            delete(v4, -2147483648)
            v1 = 61
        }
        for fi33 := 0; fi33 < 11; fi33++ {
            var v34 int64 = -1
            var v35 float32 = (1e10 + 0.0)
            v36 := GenS1{F0: 1, F1: 128}
            _ = fi33
            _ = v34
            _ = v35
            _ = v36
        }
        var v37 uint16 = 747
        _ = fi8
        _ = v37
    }
    fmt.Println(v1, v3)
    v38 := v5[4294967295]
    var v39 int32 = -2147483648
    fmt.Println(v1, v1)
    switch 159 {
    case 9223372036854775807:
        for fi40 := 0; fi40 < 9; fi40++ {
            fmt.Println(v1, v39)
            for fi41 := 0; fi41 < 13; fi41++ {
                v38 += v38
                v1 = v1
                if false {
                    v38 = (v38 << ((4294967295) & 63))
                    v3 = "foo bar"
                }
                _ = fi41
            }
            if len(v2) > 0 {
                v2[0] = v38
            }
            _ = fi40
        }
    case 32767:
        var v42 int64 = -9223372036854775808
        _ = v42
    case 127:
        v38 = v38
    }
    var v43 int16 = 1000
    if (((42 < 1) || true) || true) {
        if true {
            v5[int(v38)] = v38
        } else {
            if (65535 >= 42) {
                var v44 uint32 = 4294967295
                v45 := v4[369]
                for fi46 := 0; fi46 < 11; fi46++ {
                    if true {
                        var v47 int16 = v43
                        _ = v47
                    }
                    v48 := GenS6{F0: (v1 | v1)}
                    if (true && true) {
                        var v49 int8 = -1
                        v43 += v43
                        _ = v49
                    }
                    _ = fi46
                    _ = v48
                }
                _ = v44
                _ = v45
            }
            var v50 int64 = 27
            if (((false && true) && true) || ((v39 % ((-403) | 1)) != 833)) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v2 = append(v2, (v38 * v38))
            } else {
                for fi51 := 0; fi51 < 1; fi51++ {
                    switch ((v3 + "") + v3) {
                    case "foo bar":
                        var v52 uint = 1
                        _ = v52
                    default:
                        for fi53 := 0; fi53 < 17; fi53++ {
                            switch (-(fi51)) {
                            case 42:
                                v54 := GenS5{F0: int64(v1)}
                                _ = v54
                            case -2147483648:
                                var v55 uint = 18446744073709551615
                                _ = v55
                            case 32767:
                                var v56 int16 = v43
                                _ = v56
                            }
                            fmt.Println(v39, v1, v3)
                            v38 += v38
                            _ = fi53
                        }
                    }
                    _ = fi51
                }
                var v57 uint32 = (v38 << ((42) & 63))
                _ = v57
            }
            _ = v50
        }
    }
    for fi58 := 0; fi58 < 8; fi58++ {
        var v59 bool = true
        if v59 {
            switch (fi58 / ((fi58) | 1)) {
            case 9223372036854775807:
                var v60 float32 = 3.14159
                _ = v60
            case 1:
                var v61 uint = 42
                _ = v61
            default:
                fmt.Println(v43, v39, v39)
            }
            if v59 {
                var v62 uint32 = (^(v38))
                _ = v62
            } else {
                v1 = 2
                fmt.Println(v3, v38, v43)
            }
            var v63 float32 = -2.25
            _ = v63
        } else {
            var v64 string = (v3 + v3)
            if false {
                var v65 float64 = (0.1 * (-2.25 * (float64((0.0 / 1e10)) / 1.0)))
                _ = v65
            }
            if true {
                v64 = ((("" + (("hello" + v64) + v64)) + "hello") + v64)
            } else {
                v2 = append(v2, (v38 - (v38 & v38)))
                if (((("" + "a") != v3) && (!true)) && (false && v59)) {
                    for fi66 := 0; fi66 < 1; fi66++ {
                        v59 = v59
                        _ = len(v2)
                        if v59 {
                            var v67 float32 = float32(0.1)
                            if (v59 || (-2147483648 >= fi66)) {
                                var v68 string = v3
                                if v59 {
                                    var v69 float32 = v67
                                    var v70 int8 = 100
                                    _ = v69
                                    _ = v70
                                }
                                var v71 float32 = (v67 - 3.14159)
                                _ = v68
                                _ = v71
                            }
                            _ = v67
                        } else {
                            if (0 > -1) {
                                var v72 uint64 = 42
                                if (v72 == (v72 + v72)) {
                                    v43 = v43
                                    fmt.Println(fi66)
                                    v72 = 4294967295
                                } else {
                                    var v73 int16 = (v43 & v43)
                                    _ = v73
                                }
                                _ = v72
                            } else {
                                var v74 uint = 581
                                _ = v74
                            }
                        }
                        _ = fi66
                    }
                }
            }
            _ = v64
        }
        v1 = v1
        _ = fi58
        _ = v59
    }
    var v75 uint32 = 42
    fmt.Println(v43, v1)
    v3 = v3
    var v76 int8 = -1
    var v77 uint32 = 65535
    if len(v2) > 0 {
        v2[0] = v77
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v43
    _ = v5
    _ = v75
    _ = v76
    _ = v77
}
