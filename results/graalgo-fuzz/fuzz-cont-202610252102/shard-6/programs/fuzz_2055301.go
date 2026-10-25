package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int64; F2 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 int64; F2 uint8 }
type GenS5 struct { F0 int64 }

func main() {
    v1 := []uint32{}
    var v2 int64 = 9223372036854775807
    var v3 string = "foo bar"
    if (1 <= int32(v2)) {
        for fi4 := 0; fi4 < 5; fi4++ {
            switch 9223372036854775807 {
            case -2147483648:
                v2 = -1
            case 9223372036854775807:
                v5 := GenS1{F0: 4294967295}
                _ = v5
            }
            v2 = (v2 + 42)
            _ = fi4
        }
    } else {
        for fi6 := 0; fi6 < 16; fi6++ {
            if (false || true) {
                if false {
                    v1 = append(v1, 555)
                    for fi7 := 0; fi7 < 20; fi7++ {
                        _ = len(v1)
                        v2 = v2
                        switch (-(fi6)) {
                        case -9223372036854775808:
                            if false {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v8 := GenS1{F0: 2}
                                v3 = ("foo bar" + v3)
                                _ = v8
                            }
                        case 377:
                            var v9 string = ("" + "foo bar")
                            _ = v9
                        default:
                            var v10 int8 = 42
                            v10 = (v10 ^ v10)
                            v3 = v3
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v11 float64 = 0.1
                            var v12 uint32 = 65535
                            _ = v10
                            _ = v11
                            _ = v12
                        }
                        _ = fi7
                    }
                }
            } else {
                var v13 int64 = v2
                _ = v13
            }
            v1 = append(v1, 4294967295)
            _ = fi6
        }
    }
    switch v3 {
    case "hello":
        fmt.Println(v2)
    }
    var v14 uint64 = 737
    v2 = (v2 << ((42) & 63))
    v3 = v3
    var v15 uint8 = 255
    for fi16 := 0; fi16 < 4; fi16++ {
        v2 = -982
        _ = fi16
    }
    v1 = append(v1, 0)
    v17 := GenS1{F0: v14}
    switch ((("\n\t" + v3) + v3) + (v3 + v3)) {
    case "\n\t":
        if (false || (4294967295 < 4294967295)) {
            if true {
                fmt.Println(v15)
            } else {
                if (0.1 <= (0.0 / 1.0)) {
                    v3 = ""
                    var v18 int16 = 42
                    if (false && false) {
                        var v19 float64 = 0.0
                        v20 := GenS2{F0: (v14 | v14), F1: v2, F2: 1}
                        _ = v19
                        _ = v20
                    } else {
                        if true {
                            for fi21 := 0; fi21 < 15; fi21++ {
                                if (42 == 11) {
                                    for fi22 := 0; fi22 < 3; fi22++ {
                                        _ = len(v1)
                                        switch ((v3 + v3) + v3) {
                                        case "a":
                                            v2 = 1
                                        case "":
                                            if true {
                                                switch v3 {
                                                case "foo bar":
                                                    var v23 string = "a"
                                                    _ = v23
                                                case "\n\t":
                                                    var v24 float64 = (3.14159 - 0.1)
                                                    _ = v24
                                                case "":
                                                    var v25 int = fi21
                                                    _ = v25
                                                default:
                                                    var v26 uint = 65535
                                                    _ = v26
                                                }
                                                v14 = v14
                                                var v27 int32 = 112
                                                _ = v27
                                            }
                                        case "\n\t":
                                            v14 += v14
                                        default:
                                            var v28 int = (fi22 - int(v18))
                                            v14 = v14
                                            switch "" {
                                            case "":
                                                switch (fi22 - 128) {
                                                case 406:
                                                    v14 += v14
                                                case -1:
                                                    if (!false) {
                                                        v15 += v15
                                                        var v29 int16 = 42
                                                        _ = v29
                                                    }
                                                default:
                                                    v14 += v14
                                                }
                                            case "a":
                                                v30 := GenS3{F0: v14}
                                                _ = v30
                                            case "\n\t":
                                                v2 -= v2
                                            }
                                            if (-32 == 0) {
                                                v14 = 18446744073709551615
                                                v31 := GenS3{F0: (^(v14))}
                                                var v32 uint16 = 1
                                                if (v32 == (v32 * 65535)) {
                                                    switch v3 {
                                                    case "":
                                                        v1 = append(v1, 206)
                                                    default:
                                                        v33 := GenS2{F0: (v14 | (v14 - 42)), F1: (v2 & (v2 % ((v2) | 1))), F2: fi21}
                                                        _ = v33
                                                    }
                                                } else {
                                                    v34 := GenS4{F0: 18446744073709551615, F1: v2, F2: v15}
                                                    v1 = append(v1, 2147483647)
                                                    var v35 int8 = 127
                                                    _ = v34
                                                    _ = v35
                                                }
                                                v36 := GenS5{F0: 1}
                                                _ = v31
                                                _ = v32
                                                _ = v36
                                            } else {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                fmt.Println(v18, v3, fi22)
                                            }
                                            var v37 uint16 = 1
                                            var v38 int8 = -99
                                            if (v15 < 255) {
                                                v2 -= -9223372036854775808
                                                v3 = v3
                                                _ = len(v1)
                                            }
                                            v1 = append(v1, 65535)
                                            fmt.Println(v38, v2, v3)
                                            _ = v28
                                            _ = v37
                                            _ = v38
                                        }
                                        _ = fi22
                                    }
                                    v3 = ("a" + "")
                                }
                                for fi39 := 0; fi39 < 12; fi39++ {
                                    v14 = v14
                                    v40 := GenS2{F0: 18446744073709551615, F1: (v2 / (((v2 % ((v2) | 1))) | 1)), F2: -9223372036854775808}
                                    var v41 int16 = -379
                                    _ = fi39
                                    _ = v40
                                    _ = v41
                                }
                                _ = fi21
                            }
                        } else {
                            switch (v3 + v3) {
                            case "\n\t":
                                v42 := GenS1{F0: 584}
                                _ = v42
                            case "hello":
                                var v43 int16 = (v18 | (v18 % ((v18) | 1)))
                                _ = v43
                            }
                            fmt.Println(v14, v18)
                        }
                    }
                    _ = v18
                }
            }
            var v44 int32 = -1
            _ = v44
        } else {
            if len(v1) > 0 {
                v1[0] = 973
            }
        }
    case "":
        var v45 int8 = 60
        _ = v45
    case "hello":
        fmt.Println(v15)
    }
    if true {
        v46 := GenS5{F0: 9223372036854775807}
        fmt.Println(v2, v14)
        v47 := GenS5{F0: 0}
        _ = v46
        _ = v47
    } else {
        if (4294967295 != 2147483648) {
            fmt.Println(v15, v15)
        } else {
            var v48 float32 = (-2.25 * (3.14159 - -2.25))
            v15 = v15
            _ = v48
        }
    }
    v49 := GenS1{F0: (v14 >> ((42) & 63))}
    v50 := GenS3{F0: (v14 << ((42) & 63))}
    v3 = ((v3 + ("a" + ((v3 + (v3 + (("" + (v3 + v3)) + ""))) + "\n\t"))) + v3)
    var v51 bool = true
    fmt.Println(v3, v51, v2)
    for fi52 := 0; fi52 < 7; fi52++ {
        v3 = ("" + ((v3 + "") + v3))
        _ = fi52
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v49
    _ = v50
    _ = v51
}
