package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int8 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 uint; F1 int64; F2 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint; F1 uint64; F2 int }
type GenS7 struct { F0 int64; F1 int8; F2 int16 }
type GenS8 struct { F0 uint32; F1 int8; F2 int16 }
func genFn0(p0 uint8, p1 int) uint8 {
    v5 := []int{(p1 | -770), (p1 / ((p1) | 1)), (^(p1))}
    var v6 float32 = (1.5 - (3.14159 * (1e10 + 3.14159)))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return p0
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := []int{786, 0, -2147483648, 673}
    v3 := make(map[int]int)
    v4 := []string{"", ("" + ("hello" + "")), "hello", (("" + "\n\t") + "hello")}
    fmt.Println(v1)
    fmt.Println(v1)
    v7 := v3[-124]
    if (224 >= 1) {
        fmt.Println(v1, v7)
        if true {
            fmt.Println(v7)
            var v8 uint32 = 0
            v9 := GenS1{F0: (v1 % ((v1) | 1)), F1: 2147483647, F2: 7}
            _ = v8
            _ = v9
        }
    }
    var v10 string = "hello"
    v11 := GenS2{F0: v10, F1: 42}
    delete(v3, v7)
    v12 := GenS3{F0: v1, F1: -891, F2: 20}
    if (-347 > 2147483647) {
        v1 = (^(v1))
        if false {
            switch ((v10 + "hello") + v10) {
            case "hello":
                var v13 int16 = 42
                _ = v13
            case "a":
                v10 = v10
            case "\n\t":
                fmt.Println(v7, v1)
            default:
                if true {
                    switch 127 {
                    case 4294967295:
                        var v14 uint = (^(v1))
                        _ = v14
                    case 1:
                        fmt.Println(v1, v1)
                    case 0:
                        var v15 int = (-(v7))
                        _ = v15
                    default:
                        if (!(32767 <= 0)) {
                            for fi16 := 0; fi16 < 12; fi16++ {
                                v17 := GenS4{F0: true}
                                var v18 uint32 = uint32(v1)
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                            var v19 int16 = 42
                            _ = v19
                        } else {
                            var v20 uint64 = 0
                            _ = v20
                        }
                    }
                    _ = len(v4)
                } else {
                    var v21 uint64 = 765
                    _ = v21
                }
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if false {
                if (float64(1e10) != 1.5) {
                    v22 := GenS4{F0: (!true)}
                    _ = v22
                } else {
                    for fi23 := 0; fi23 < 17; fi23++ {
                        v3[(^(v7))] = 127
                        for fi24 := 0; fi24 < 5; fi24++ {
                            v10 = v10
                            _ = fi24
                        }
                        delete(v3, (fi23 - -9223372036854775808))
                        _ = fi23
                    }
                    switch "" {
                    case "":
                        switch "\n\t" {
                        case "hello":
                            v1 = 4294967295
                        case "":
                            v25 := v3[v7]
                            _ = v25
                        case "foo bar":
                            for fi26 := 0; fi26 < 16; fi26++ {
                                v27 := GenS4{F0: ((false && true) || true)}
                                var v28 int32 = 0
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                        }
                    case "a":
                        _ = len(v3)
                    case "foo bar":
                        v1 = (v1 * v1)
                    default:
                        if (155 <= 127) {
                            v4 = append(v4, "")
                            var v29 int = (^(v7))
                            _ = v29
                        } else {
                            var v30 int = v7
                            v31 := GenS3{F0: v1, F1: 9223372036854775807, F2: 255}
                            if false {
                                switch (v10 + v10) {
                                case "":
                                    for fi32 := 0; fi32 < 7; fi32++ {
                                        for fi33 := 0; fi33 < 18; fi33++ {
                                            switch ("a" + ((v10 + v10) + "foo bar")) {
                                            case "a":
                                                switch v30 {
                                                case 9223372036854775807:
                                                    v30 += -856
                                                default:
                                                    v10 = "a"
                                                }
                                            default:
                                                if ((32767 >= -32768) && false) {
                                                    switch fi32 {
                                                    case 0:
                                                        v1 += 539
                                                    case 1:
                                                        if false {
                                                            switch v10 {
                                                            case "a":
                                                                _ = len(v2)
                                                            case "hello":
                                                                v34 := GenS3{F0: v1, F1: -1, F2: 65535}
                                                                _ = v34
                                                            }
                                                            if (!(false && true)) {
                                                                var v35 uint16 = 24
                                                                _ = v35
                                                            } else {
                                                                if (true && true) {
                                                                    v30 = (v7 | fi32)
                                                                    v36 := GenS5{F0: 169}
                                                                    _ = v36
                                                                } else {
                                                                    fmt.Println(v30, v30)
                                                                }
                                                                var v37 float64 = (-2.25 - (0.1 * (float64(1.5) / ((1.5 - 1.5) - -2.25))))
                                                                _ = v37
                                                            }
                                                            _ = len(v4)
                                                        }
                                                    case -863:
                                                        v38 := GenS6{F0: v1, F1: 12, F2: (fi32 + -9223372036854775808)}
                                                        _ = v38
                                                    default:
                                                        var v39 uint64 = 0
                                                        _ = v39
                                                    }
                                                } else {
                                                    v40 := GenS3{F0: (v1 | 680), F1: 1, F2: 106}
                                                    v1 += 18446744073709551615
                                                    _ = v40
                                                }
                                                var v41 int32 = 42
                                                _ = v41
                                            }
                                            v3[fi33] = fi33
                                            var v42 int16 = -32768
                                            _ = fi33
                                            _ = v42
                                        }
                                        v2 = append(v2, (v7 & (-(fi32))))
                                        v43 := v3[(fi32 % ((v7) | 1))]
                                        _ = fi32
                                        _ = v43
                                    }
                                }
                            } else {
                                v3[32767] = v7
                                fmt.Println(v30, v1, v1)
                                for fi44 := 0; fi44 < 15; fi44++ {
                                    v10 = "hello"
                                    v10 = (("foo bar" + ((("\n\t" + "hello") + v10) + ((("\n\t" + "a") + (v10 + "hello")) + ((v10 + v10) + v10)))) + v10)
                                    _ = fi44
                                }
                            }
                            _ = v30
                            _ = v31
                        }
                    }
                    var v45 int64 = -9223372036854775808
                    _ = v45
                }
            } else {
                var v46 int64 = -648
                for fi47 := 0; fi47 < 9; fi47++ {
                    v48 := GenS7{F0: (v46 ^ (^(v46))), F1: 1, F2: -1}
                    _ = fi47
                    _ = v48
                }
                _ = v46
            }
        }
    }
    switch v10 {
    case "hello":
        v1 -= (v1 % (((^(v1))) | 1))
    }
    switch (v7 - v7) {
    case 134:
        for fi49 := 0; fi49 < 10; fi49++ {
            delete(v3, v7)
            fmt.Println(v10, fi49)
            _ = fi49
        }
    case 42:
        var v50 int8 = -11
        _ = v50
    }
    for fi51 := 0; fi51 < 5; fi51++ {
        if (!((((v1 + 193) > v1) && true) && (127 < 127))) {
            v52 := GenS8{F0: 2147483647, F1: -113, F2: 1}
            if (128 > 1) {
                v7 -= (fi51 + 128)
                v53 := GenS5{F0: -125}
                _ = v53
            } else {
                for fi54 := 0; fi54 < 9; fi54++ {
                    var v55 float32 = (((((1.5 + (0.0 / 0.1)) - (-2.25 / 0.1)) * 3.14159) - (3.14159 + (1e10 + ((1.5 * (3.14159 + 0.1)) * 0.1)))) + 3.14159)
                    v10 = v10
                    v1 = 718
                    _ = fi54
                    _ = v55
                }
            }
            v7 -= (fi51 >> ((1) & 63))
            _ = v52
        }
        _ = fi51
    }
    for fi56 := 0; fi56 < 12; fi56++ {
        if ((!false) || (false && (!false))) {
            v1 -= v1
        } else {
            if (-56 >= 127) {
                if (false && (((1e10 <= (1e10 - 1.5)) && true) && true)) {
                    var v57 uint32 = 4294967295
                    _ = v57
                }
            } else {
                v58 := GenS3{F0: (^(v1)), F1: 0, F2: 255}
                _ = v58
            }
            fmt.Println(v10, v1)
        }
        _ = fi56
    }
    v10 = v10
    var v59 uint = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v59:", v59)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v59
    _ = v7
}
