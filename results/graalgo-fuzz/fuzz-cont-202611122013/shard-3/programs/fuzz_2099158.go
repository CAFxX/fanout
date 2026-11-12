package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16; F1 uint16; F2 uint64 }

func main() {
    var v1 float64 = (((1.5 + 0.1) + 0.1) / 1e10)
    v2 := []uint32{289, 2, 4294967295}
    var v3 string = "\n\t"
    var v4 int16 = 1
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 bool = ((v4 - v4) > -32768)
        v7 := GenS1{F0: 124}
        v8 := GenS1{F0: 18446744073709551615}
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 int64 = 1
    for fi10 := 0; fi10 < 1; fi10++ {
        if (!(88 < 156)) {
            if false {
                var v11 float64 = 0.0
                _ = v11
            }
        }
        _ = fi10
    }
    var v12 float32 = 1e10
    if (v12 == 3.14159) {
        v4 = (v4 % (((-(v4))) | 1))
        fmt.Println(v12)
    }
    if (true && (false || (!false))) {
        for fi13 := 0; fi13 < 7; fi13++ {
            fmt.Println(v1)
            if true {
                if (!(1 != 154)) {
                    for fi14 := 0; fi14 < 16; fi14++ {
                        v4 += (-(v4))
                        if false {
                            if len(v2) > 0 {
                                v2[0] = 4294967295
                            }
                        }
                        switch (v3 + ("a" + "hello")) {
                        case "hello":
                            v15 := GenS1{F0: 0}
                            _ = v15
                        case "a":
                            var v16 float64 = v1
                            if false {
                                var v17 uint32 = 65535
                                _ = v17
                            }
                            var v18 int16 = (v4 ^ -667)
                            v3 = ""
                            _ = v16
                            _ = v18
                        case "\n\t":
                            if (1 != 2) {
                                var v19 int = (fi14 >> ((1) & 63))
                                var v20 uint8 = 0
                                var v21 int64 = 882
                                if len(v2) > 0 {
                                    v2[0] = 4294967295
                                }
                                switch ((v3 + ((((v3 + "foo bar") + v3) + "hello") + "a")) + "a") {
                                case "a":
                                    v22 := GenS1{F0: 19}
                                    _ = v22
                                case "\n\t":
                                    if (false && (-2.25 != (v12 + v12))) {
                                        switch (fi14 & fi14) {
                                        case -877:
                                            v23 := GenS1{F0: 2}
                                            v21 += (v21 + v21)
                                            var v24 string = v3
                                            _ = v23
                                            _ = v24
                                        case -1:
                                            v25 := GenS1{F0: 428}
                                            _ = v25
                                        case 127:
                                            fmt.Println(v1, v9)
                                            v26 := GenS2{F0: (v4 * (v4 | 0))}
                                            _ = v26
                                        default:
                                            v9 = -1
                                        }
                                        var v27 string = v3
                                        _ = v27
                                    }
                                case "foo bar":
                                    v20 = v20
                                }
                                _ = v19
                                _ = v20
                                _ = v21
                            } else {
                                if true {
                                    v12 += v12
                                }
                                _ = len(v2)
                                v28 := GenS1{F0: 65535}
                                fmt.Println(fi13, v12, fi14)
                                if (-2.25 >= ((((float32(v1) + -2.25) / ((v12 - float32(0.0)) + -2.25)) - 3.14159) + ((v12 - v12) / v12))) {
                                    v29 := GenS2{F0: (v4 % ((v4) | 1))}
                                    _ = v29
                                }
                                _ = v28
                            }
                        default:
                            switch fi13 {
                            case 738:
                                v30 := GenS2{F0: (v4 | v4)}
                                _ = v30
                            default:
                                v31 := GenS3{F0: 250, F1: 0}
                                _ = v31
                            }
                            _ = len(v2)
                            _ = len(v2)
                            var v32 int = fi13
                            _ = v32
                        }
                        if (128 != 2) {
                            if true {
                                if len(v2) > 0 {
                                    v2[0] = 714
                                }
                                v12 = float32(v1)
                                v33 := GenS3{F0: 2, F1: 1}
                                _ = v33
                            } else {
                                v9 += v9
                                v12 += (float32(v1) + ((1.5 * (float32(v1) - v12)) * float32(1e10)))
                                v3 = "foo bar"
                                v9 -= v9
                            }
                            fmt.Println(v9, fi14, v1)
                        }
                        _ = fi14
                    }
                    var v34 uint = 2
                    v35 := GenS4{F0: 1}
                    _ = v34
                    _ = v35
                } else {
                    if len(v2) > 0 {
                        v2[0] = 65535
                    }
                }
            }
            if true {
                if len(v2) > 0 {
                    v2[0] = 34
                }
                var v36 uint = 18446744073709551615
                for fi37 := 0; fi37 < 16; fi37++ {
                    v38 := GenS1{F0: (v36 + (v36 ^ 42))}
                    if (3.14159 != ((float32(v1) - (0.0 * v12)) + v12)) {
                        _ = len(v2)
                        if true {
                            v39 := GenS3{F0: 1, F1: 0}
                            _ = v39
                        }
                        if (v9 != v9) {
                            var v40 int8 = 8
                            v41 := GenS4{F0: 2}
                            fmt.Println(v40, v40)
                            v36 -= (v36 << ((2) & 63))
                            _ = v40
                            _ = v41
                        }
                        var v42 int8 = -1
                        var v43 int = -9223372036854775808
                        _ = v42
                        _ = v43
                    }
                    v2 = append(v2, 2)
                    _ = fi37
                    _ = v38
                }
                _ = v36
            }
            _ = fi13
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v4 = v4
    }
    if true {
        v44 := GenS1{F0: 18446744073709551615}
        _ = v44
    }
    v45 := GenS2{F0: v4}
    switch -1 {
    case 32768:
        _ = len(v2)
    case 127:
        v46 := GenS5{F0: (v4 | (v4 + (v4 ^ (v4 + (v4 + 42))))), F1: 32767, F2: 4294967295}
        _ = v46
    case -1:
        var v47 int32 = 1
        _ = v47
    }
    for fi48 := 0; fi48 < 9; fi48++ {
        switch fi48 {
        case 127:
            fmt.Println(v4, v1)
        case -946:
            for fi49 := 0; fi49 < 8; fi49++ {
                fmt.Println(v3, fi48, fi49)
                var v50 uint = 422
                _ = fi49
                _ = v50
            }
        }
        _ = fi48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v9
}
