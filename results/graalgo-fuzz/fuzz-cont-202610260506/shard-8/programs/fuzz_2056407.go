package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int; F1 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 0
    v3 := []uint32{}
    var v4 uint64 = 18446744073709551615
    var v5 bool = false
    for fi6 := 0; fi6 < 17; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 uint8 = 53
            if v5 {
                v1[4294967295] = ((((((float64(3.14159) * 3.14159) * (0.1 / 1.5)) / -2.25) + (1.5 * 0.0)) - 0.0) > -2.25)
            }
            v5 = ((^(v8)) <= v8)
            _ = fi7
            _ = v8
        }
        v5 = true
        _ = fi6
    }
    v9 := GenS1{F0: 537, F1: (^(v2)), F2: (0.1 >= ((3.14159 / 1.0) + 1e10))}
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 float32 = 0.0
        _ = fi10
        _ = v11
    }
    v4 -= 455
    v5 = false
    var v12 uint32 = 42
    for fi13 := 0; fi13 < 10; fi13++ {
        if (!v5) {
            var v14 uint16 = 591
            if (4294967295 > 18446744073709551615) {
                v5 = false
                v3 = append(v3, v12)
            } else {
                fmt.Println(v2, fi13)
                v4 = (v4 >> ((42) & 63))
            }
            if (!v5) {
                switch "\n\t" {
                case "\n\t":
                    v15 := GenS2{F0: -1}
                    _ = v15
                case "foo bar":
                    for fi16 := 0; fi16 < 18; fi16++ {
                        switch "" {
                        case "\n\t":
                            var v17 string = "a"
                            _ = v17
                        default:
                            var v18 int16 = 42
                            _ = v18
                        }
                        switch int(v12) {
                        case -418:
                            fmt.Println(v4, v5)
                        case 32768:
                            if v5 {
                                v19 := GenS1{F0: 0, F1: 0, F2: v5}
                                switch (^(fi13)) {
                                case -2147483648:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case 1:
                                    var v20 float32 = 1.5
                                    _ = v20
                                default:
                                    var v21 float64 = (1.5 + 0.0)
                                    _ = v21
                                }
                                v4 = (v4 % ((v4) | 1))
                                var v22 float64 = 1.5
                                _ = v19
                                _ = v22
                            } else {
                                var v23 bool = v5
                                v24 := GenS2{F0: -32768}
                                var v25 uint = 255
                                v26 := GenS3{F0: v14, F1: 255}
                                var v27 int32 = 42
                                if v5 {
                                    v28 := v1[v12]
                                    v25 -= 2
                                    _ = v28
                                }
                                _ = v23
                                _ = v24
                                _ = v25
                                _ = v26
                                _ = v27
                            }
                        case 42:
                            var v29 uint32 = v12
                            _ = v29
                        default:
                            v30 := GenS4{F0: uint64(v12), F1: 192}
                            _ = v30
                        }
                        _ = fi16
                    }
                }
                switch (fi13 | (fi13 + fi13)) {
                case -826:
                    v31 := GenS3{F0: (v14 / ((v2) | 1)), F1: 50}
                    _ = v31
                }
            } else {
                fmt.Println(v14)
            }
            _ = v14
        }
        if (v4 > (v4 >> ((299) & 63))) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if ((3.14159 + (1.5 - -2.25)) == 3.14159) {
                var v32 uint = 963
                v32 += 4294967295
                _ = v32
            } else {
                v12 += uint32(v4)
                for fi33 := 0; fi33 < 19; fi33++ {
                    var v34 float32 = 0.1
                    var v35 float32 = float32(0.1)
                    v5 = false
                    v4 = (v4 >> ((18446744073709551615) & 63))
                    if (((3.14159 - 0.1) < 1e10) && false) {
                        var v36 uint = 4294967295
                        _ = v36
                    } else {
                        v35 = v34
                        switch (fi13 << ((630) & 63)) {
                        case 951:
                            if v5 {
                                v5 = (!false)
                                if true {
                                    v5 = (!(3.14159 != 0.1))
                                    switch "foo bar" {
                                    case "\n\t":
                                        fmt.Println(v35)
                                    case "":
                                        if (v5 || (fi33 == (fi33 * 0))) {
                                            v5 = true
                                            v37 := GenS5{F0: (^(v4))}
                                            _ = v37
                                        } else {
                                            v12 = v12
                                        }
                                    case "a":
                                        fmt.Println(fi13, v4)
                                    default:
                                        var v38 uint64 = v4
                                        _ = v38
                                    }
                                } else {
                                    fmt.Println(v12)
                                    var v39 int8 = -51
                                    fmt.Println(v2, v2, v12)
                                    var v40 int64 = -9223372036854775808
                                    _ = v39
                                    _ = v40
                                }
                                if len(v3) > 0 {
                                    v3[0] = v12
                                }
                            }
                        default:
                            v35 = 0.0
                        }
                    }
                    _ = fi33
                    _ = v34
                    _ = v35
                }
                v1[v12] = v5
            }
            for fi41 := 0; fi41 < 5; fi41++ {
                if false {
                    v42 := GenS4{F0: 516, F1: (v12 >> ((617) & 63))}
                    v4 += 1
                    _ = v42
                } else {
                    v2 -= v2
                    fmt.Println(fi41)
                }
                var v43 uint64 = v4
                _ = len(v3)
                _ = fi41
                _ = v43
            }
        }
        _ = fi13
    }
    var v44 int = 576
    for fi45 := 0; fi45 < 8; fi45++ {
        switch "\n\t" {
        case "a":
            if (!true) {
                var v46 string = "\n\t"
                var v47 float32 = 0.1
                _ = v46
                _ = v47
            } else {
                if (0 != 2147483647) {
                    v44 += (fi45 & fi45)
                }
            }
        case "hello":
            var v48 float32 = (((float32(1e10) / float32(1.5)) + 0.1) + -2.25)
            _ = v48
        default:
            v49 := v1[v12]
            _ = v49
        }
        v50 := GenS2{F0: 789}
        switch "a" {
        case "":
            v12 += v12
        case "\n\t":
            v4 += v4
        case "foo bar":
            fmt.Println(v2, v5)
        default:
            for fi51 := 0; fi51 < 16; fi51++ {
                v44 -= fi45
                v52 := GenS6{F0: 538}
                switch "" {
                case "":
                    fmt.Println(v44)
                case "foo bar":
                    if v5 {
                        v53 := v1[(v12 >> ((65535) & 63))]
                        if (2 <= 42) {
                            v4 += (v4 * 246)
                            v54 := GenS1{F0: 276, F1: (^(v2)), F2: false}
                            _ = v54
                        }
                        _ = v53
                    } else {
                        v55 := GenS7{F0: fi51, F1: (v2 | 1)}
                        _ = v55
                    }
                default:
                    v56 := GenS6{F0: 1000}
                    _ = v56
                }
                _ = fi51
                _ = v52
            }
        }
        _ = fi45
        _ = v50
    }
    v4 = v4
    fmt.Println(v5)
    v4 = 937
    for fi57 := 0; fi57 < 13; fi57++ {
        v5 = v5
        var v58 int32 = -2147483648
        _ = fi57
        _ = v58
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v9
}
