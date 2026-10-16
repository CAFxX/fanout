package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint32; F1 int; F2 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 int8 }

func main() {
    var v1 uint32 = 631
    v2 := []uint32{(^(v1)), (v1 | (v1 >> ((676) & 63)))}
    v3 := make(map[int]int)
    var v4 int64 = -9223372036854775808
    v5 := []string{("" + "hello"), "\n\t"}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi6 := 0; fi6 < 15; fi6++ {
        if ((v1 & v1) >= (v1 << ((522) & 63))) {
            fmt.Println(v1, v4, v1)
        }
        for fi7 := 0; fi7 < 20; fi7++ {
            if false {
                v4 = v4
            }
            if false {
                v8 := GenS1{F0: 153}
                if (((v4 - (v4 * -600)) == (v4 | v4)) || false) {
                    var v9 uint32 = v1
                    if (!false) {
                        var v10 uint8 = 240
                        if ((!(!true)) || false) {
                            fmt.Println(fi7, fi6)
                        }
                        _ = v10
                    }
                    _ = v9
                }
                fmt.Println(fi7, v1, v4)
                v1 += 2147483648
                _ = v8
            }
            v4 = v4
            _ = fi7
        }
        _ = fi6
    }
    v11 := GenS2{F0: -1, F1: -102}
    for fi12 := 0; fi12 < 3; fi12++ {
        v4 -= -1
        var v13 uint64 = 1
        if (!(v1 > v1)) {
            var v14 float64 = 1e10
            var v15 int16 = -791
            v16 := GenS3{F0: (v1 << ((142) & 63)), F1: 127, F2: (v15 >> ((0) & 63))}
            _ = v14
            _ = v15
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    v3[-2147483648] = -2147483648
    if (220 < 0) {
        v17 := GenS1{F0: 154}
        var v18 float64 = ((3.14159 * 1.5) / 3.14159)
        v19 := GenS2{F0: -128, F1: 127}
        _ = v17
        _ = v18
        _ = v19
    }
    var v20 bool = true
    var v21 int8 = 7
    v1 = (v1 / ((2) | 1))
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 uint32 = 0
        delete(v3, fi22)
        v20 = (!(!true))
        _ = fi22
        _ = v23
    }
    switch "foo bar" {
    case "\n\t":
        for fi24 := 0; fi24 < 14; fi24++ {
            fmt.Println(v21)
            v25 := GenS3{F0: (v1 | v1), F1: (fi24 % ((fi24) | 1)), F2: -32768}
            var v26 int = fi24
            _ = fi24
            _ = v25
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 3; fi27++ {
        var v28 int16 = 686
        var v29 int16 = (v28 * v28)
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 int64 = v4
    if v20 {
        if false {
            fmt.Println(v21, v1)
        } else {
            if (v21 != 102) {
                var v31 int16 = 32767
                switch "hello" {
                case "hello":
                    var v32 int32 = -80
                    _ = v32
                case "":
                    if false {
                        v2 = append(v2, v1)
                        fmt.Println(v4, v20)
                        var v33 uint64 = 0
                        _ = v33
                    } else {
                        v34 := GenS2{F0: v21, F1: v21}
                        _ = v34
                    }
                default:
                    var v35 float32 = ((0.0 - (-2.25 - (3.14159 / 0.1))) - ((((3.14159 - (-2.25 / 1.0)) * 0.0) / (((1.5 * (1e10 / 1.0)) - 1.5) + -2.25)) + float32(1e10)))
                    _ = v35
                }
                _ = v31
            } else {
                v21 -= (-(v21))
                var v36 uint8 = 243
                switch "foo bar" {
                case "a":
                    for fi37 := 0; fi37 < 20; fi37++ {
                        if ("\n\t" < "hello") {
                            v38 := v3[fi37]
                            if v20 {
                                v39 := GenS3{F0: v1, F1: -62, F2: 522}
                                v40 := GenS3{F0: (v1 | (^(v1))), F1: (fi37 << ((42) & 63)), F2: -32768}
                                _ = v39
                                _ = v40
                            } else {
                                var v41 float64 = 3.14159
                                v1 += (v1 % (((v1 | v1)) | 1))
                                v38 = fi37
                                _ = v41
                            }
                            _ = v38
                        }
                        _ = fi37
                    }
                case "\n\t":
                    var v42 string = "a"
                    _ = v42
                default:
                    v43 := GenS4{F0: -813}
                    _ = v43
                }
                _ = v36
            }
            fmt.Println(v20, v1)
            v44 := GenS3{F0: (v1 ^ 65535), F1: 0, F2: 0}
            _ = v44
        }
        v20 = v20
    }
    switch -746 {
    case -1:
        v21 = 11
    case 598:
        v20 = true
    default:
        for fi45 := 0; fi45 < 18; fi45++ {
            var v46 float64 = ((float64(3.14159) * -2.25) - 3.14159)
            v47 := GenS5{F0: "foo bar", F1: (v21 >> ((0) & 63))}
            _ = fi45
            _ = v46
            _ = v47
        }
    }
    var v48 uint = 0
    if ((3.14159 + 3.14159) < float64(1.5)) {
        var v49 float64 = 1e10
        switch 128 {
        case 9223372036854775807:
            switch -9223372036854775808 {
            case 2147483647:
                fmt.Println(v20, v49)
            default:
                var v50 uint32 = 15
                _ = v50
            }
        case -1:
            v21 -= v21
        default:
            v51 := GenS4{F0: int(v1)}
            _ = v51
        }
        var v52 int32 = -2147483648
        _ = v49
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v4
    _ = v48
    _ = v5
}
