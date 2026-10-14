package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint32; F2 bool }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint8 }

func main() {
    v1 := []string{"hello", ("foo bar" + "hello"), "\n\t", ("foo bar" + "hello")}
    var v2 uint32 = 2147483648
    v3 := make(map[uint32]uint32)
    var v4 float64 = ((1.5 + 0.0) / 1e10)
    v5 := []string{"hello", "a", "\n\t", ((("foo bar" + "a") + ("\n\t" + ("" + "\n\t"))) + "a")}
    var v6 int8 = -117
    if (32767 != 65535) {
        v7 := GenS1{F0: 42}
        _ = v7
    } else {
        v4 = 0.1
    }
    var v8 uint64 = 2
    var v9 uint32 = (v2 >> ((507) & 63))
    if false {
        var v10 int64 = 0
        _ = v10
    }
    v11 := GenS2{F0: (^(v6)), F1: 4294967295, F2: true}
    v4 += 0.0
    if true {
        if false {
            v6 -= (v6 << ((0) & 63))
            if (-2147483648 < 676) {
                if false {
                    if (579 != -335) {
                        fmt.Println(v8)
                        switch "" {
                        case "\n\t":
                            var v12 uint = 4294967295
                            _ = v12
                        case "a":
                            switch "\n\t" {
                            case "":
                                for fi13 := 0; fi13 < 18; fi13++ {
                                    var v14 float32 = 0.0
                                    _ = fi13
                                    _ = v14
                                }
                            case "a":
                                if (((false || (555 != 51)) && ((!false) && false)) && (v8 == 942)) {
                                    fmt.Println(v6, v8, v2)
                                    v4 = v4
                                    switch -2147483648 {
                                    case -806:
                                        var v15 uint8 = 128
                                        _ = v15
                                    case 9223372036854775807:
                                        var v16 string = ("hello" + "foo bar")
                                        _ = v16
                                    case 0:
                                        v8 = v8
                                    default:
                                        if (((0 == 2) && (!(true && true))) || true) {
                                            var v17 string = (("hello" + "a") + "")
                                            v4 -= 0.0
                                            var v18 int32 = 2147483647
                                            _ = v17
                                            _ = v18
                                        }
                                    }
                                }
                            case "hello":
                                if false {
                                    var v19 int16 = 32767
                                    _ = v19
                                }
                            default:
                                var v20 uint8 = uint8(v9)
                                _ = v20
                            }
                        case "foo bar":
                            v21 := GenS2{F0: -9, F1: v9, F2: true}
                            _ = v21
                        }
                        v22 := GenS3{F0: "", F1: -9223372036854775808}
                        _ = v22
                    } else {
                        fmt.Println(v2, v2)
                        fmt.Println(v6, v9)
                        if false {
                            if false {
                                v2 = v2
                                v6 += (v6 ^ v6)
                                for fi23 := 0; fi23 < 16; fi23++ {
                                    for fi24 := 0; fi24 < 10; fi24++ {
                                        v4 = 0.0
                                        if false {
                                            v25 := GenS3{F0: "\n\t", F1: (fi24 * -9223372036854775808)}
                                            var v26 int64 = 9223372036854775807
                                            _ = v25
                                            _ = v26
                                        }
                                        v8 += v8
                                        v27 := GenS2{F0: v6, F1: 42, F2: false}
                                        v2 += (v2 | 200)
                                        v4 -= v4
                                        var v28 uint = 1
                                        _ = fi24
                                        _ = v27
                                        _ = v28
                                    }
                                    _ = fi23
                                }
                            } else {
                                v6 = v6
                                if (!true) {
                                    if (!false) {
                                        var v29 string = "foo bar"
                                        var v30 bool = (v2 == 480)
                                        var v31 uint8 = 128
                                        _ = v29
                                        _ = v30
                                        _ = v31
                                    } else {
                                        var v32 int16 = 522
                                        _ = v32
                                    }
                                    v2 = (v9 | (^(v2)))
                                    v4 = (v4 * (0.1 * v4))
                                }
                            }
                            var v33 int32 = -26
                            if true {
                                v33 = (v33 * (v33 ^ 0))
                                fmt.Println(v9, v6, v2)
                            }
                            _ = v33
                        } else {
                            if (277 > 600) {
                                for fi34 := 0; fi34 < 16; fi34++ {
                                    var v35 int = fi34
                                    v4 = float64((-2.25 + (0.0 - ((0.0 / ((1e10 / (0.1 + (float32((v4 * 1e10)) / 0.1))) / (1e10 - 1.5))) / 3.14159))))
                                    if false {
                                        if false {
                                            var v36 float32 = float32(v4)
                                            v2 = 645
                                            _ = v36
                                        } else {
                                            _ = len(v5)
                                        }
                                        fmt.Println(v8)
                                    }
                                    _ = fi34
                                    _ = v35
                                }
                                v37 := GenS1{F0: 534}
                                v3[v2] = v9
                                _ = v37
                            }
                            v2 = (v2 << ((0) & 63))
                        }
                    }
                }
            } else {
                if false {
                    var v38 int = 32768
                    _ = v38
                } else {
                    v2 += (v9 | v9)
                }
            }
        } else {
            v39 := GenS4{F0: (v8 ^ v8), F1: 2}
            _ = v39
        }
    } else {
        v40 := GenS5{F0: -9223372036854775808}
        switch "foo bar" {
        case "":
            switch "" {
            case "hello":
                v2 = v2
            case "a":
                for fi41 := 0; fi41 < 6; fi41++ {
                    for fi42 := 0; fi42 < 1; fi42++ {
                        if true {
                            switch fi42 {
                            case 907:
                                v43 := GenS6{F0: int64(v6)}
                                _ = v43
                            case -386:
                                var v44 string = "a"
                                _ = v44
                            case 32768:
                                if (18446744073709551615 <= (v8 * (v8 + 289))) {
                                    if true {
                                        delete(v3, v9)
                                        fmt.Println(v2, v9, fi41)
                                        var v45 int8 = v6
                                        v4 = 1.5
                                        var v46 bool = (!(v6 != (v45 + v6)))
                                        _ = v45
                                        _ = v46
                                    }
                                    if (((((v2 / (((v9 >> ((255) & 63))) | 1)) != (^(v9))) && true) && true) || false) {
                                        var v47 uint32 = v2
                                        var v48 bool = false
                                        if true {
                                            v49 := GenS2{F0: -37, F1: 0, F2: v48}
                                            v50 := v3[0]
                                            var v51 int32 = -234
                                            _ = v49
                                            _ = v50
                                            _ = v51
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = "a"
                                        }
                                        _ = len(v5)
                                        v52 := GenS7{F0: 16}
                                        _ = v47
                                        _ = v48
                                        _ = v52
                                    }
                                } else {
                                    fmt.Println(v8, v4, v4)
                                }
                            default:
                                var v53 int = 4294967295
                                _ = v53
                            }
                        } else {
                            v54 := GenS8{F0: 42}
                            v9 = (^(v9))
                            _ = v54
                        }
                        _ = fi42
                    }
                    v9 = (v9 ^ v9)
                    v2 -= (v2 >> ((816) & 63))
                    _ = fi41
                }
            case "foo bar":
                switch 707 {
                case 9223372036854775807:
                    switch 128 {
                    case 832:
                        v8 = v8
                    default:
                        v9 = 95
                    }
                case 42:
                    if false {
                        v6 = 42
                    }
                default:
                    var v55 uint32 = v9
                    _ = v55
                }
            }
        default:
            if (1 != 42) {
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
            }
        }
        _ = v40
    }
    var v56 float32 = ((1e10 - float32(1.5)) / -2.25)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v6
    _ = v8
    _ = v9
}
