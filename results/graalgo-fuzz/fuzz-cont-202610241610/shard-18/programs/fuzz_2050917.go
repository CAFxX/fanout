package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 uint64 }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 bool }

func main() {
    var v1 int32 = 42
    v2 := []string{"a", "\n\t"}
    v3 := make(map[int]bool)
    v4 := []uint32{42, 930, 399}
    var v5 int64 = -9223372036854775808
    var v6 uint8 = 127
    fmt.Println(v5, v6)
    if (!(42 > -1)) {
        switch int(v6) {
        case 0:
            v7 := GenS1{F0: "hello", F1: 2147483647}
            _ = v7
        }
        v8 := GenS2{F0: 18446744073709551615, F1: (v1 <= v1)}
        _ = v8
    } else {
        switch -9223372036854775808 {
        case 0:
            v5 = v5
        case 32768:
            if false {
                switch -555 {
                case -445:
                    v3[388] = true
                case 42:
                    v9 := GenS3{F0: 255, F1: "hello"}
                    _ = v9
                default:
                    fmt.Println(v5)
                }
            } else {
                v5 -= int64(v6)
                v1 = (v1 << ((1) & 63))
                for fi10 := 0; fi10 < 18; fi10++ {
                    if (((0.1 / (((0.1 - 1.5) / (((float32(float64((((3.14159 - (-2.25 / 1.0)) + (((0.0 * 0.0) - 1e10) + 0.0)) * 3.14159))) + 3.14159) * ((1e10 + 1.5) / 1e10)) / 3.14159)) - (3.14159 * 3.14159))) + 0.0) == 0.1) {
                        v6 = (v6 & v6)
                        v1 = v1
                        if (!(127 != (v6 * 255))) {
                            v11 := GenS3{F0: 18446744073709551615, F1: "foo bar"}
                            var v12 uint8 = (v6 % (((v6 | v6)) | 1))
                            for fi13 := 0; fi13 < 8; fi13++ {
                                v1 = -2147483648
                                switch 32768 {
                                case 2147483647:
                                    v14 := GenS3{F0: 2, F1: "hello"}
                                    _ = v14
                                }
                                if ((18446744073709551615 == 267) || (255 >= 255)) {
                                    v6 += v12
                                    var v15 int = fi13
                                    if true {
                                        switch fi10 {
                                        case -186:
                                            var v16 uint = 1
                                            _ = v16
                                        case 32767:
                                            v4 = append(v4, 2)
                                            v17 := GenS4{F0: v12}
                                            _ = len(v3)
                                            switch 9223372036854775807 {
                                            case 127:
                                                var v18 uint16 = 2
                                                _ = v18
                                            case -535:
                                                v19 := v3[(-(v15))]
                                                _ = v19
                                            default:
                                                v12 = v6
                                            }
                                            var v20 bool = (!("foo bar" != "\n\t"))
                                            _ = v17
                                            _ = v20
                                        default:
                                            v21 := GenS3{F0: 29, F1: ("foo bar" + "a")}
                                            _ = v21
                                        }
                                    }
                                    if (!false) {
                                        var v22 uint16 = 482
                                        v6 = v12
                                        var v23 int = fi10
                                        _ = v22
                                        _ = v23
                                    } else {
                                        if (false || false) {
                                            v6 += v12
                                        }
                                        _ = len(v3)
                                        if (false || true) {
                                            v5 -= v5
                                            var v24 int = (fi13 % (((^(fi10))) | 1))
                                            v25 := GenS5{F0: 127, F1: 18446744073709551615}
                                            _ = v24
                                            _ = v25
                                        } else {
                                            fmt.Println(v5)
                                        }
                                    }
                                    delete(v3, -1)
                                    _ = v15
                                }
                                _ = fi13
                            }
                            _ = v11
                            _ = v12
                        } else {
                            var v26 float32 = (3.14159 + 3.14159)
                            _ = v26
                        }
                    }
                    var v27 uint64 = 18446744073709551615
                    _ = fi10
                    _ = v27
                }
            }
        case 2147483647:
            switch 65 {
            case -1:
                var v28 int = -1
                _ = v28
            case -2147483648:
                if (!(v6 == 30)) {
                    var v29 float64 = 0.1
                    _ = v29
                }
            case 128:
                var v30 int64 = 0
                _ = v30
            default:
                v4 = append(v4, 1)
            }
        default:
            v31 := GenS1{F0: "\n\t", F1: 65535}
            _ = v31
        }
        v1 = (^(v1))
    }
    v6 = v6
    _ = len(v4)
    for fi32 := 0; fi32 < 17; fi32++ {
        var v33 uint = 658
        _ = fi32
        _ = v33
    }
    var v34 uint = 2
    v35 := GenS5{F0: 67, F1: 4294967295}
    if true {
        for fi36 := 0; fi36 < 12; fi36++ {
            for fi37 := 0; fi37 < 1; fi37++ {
                v38 := GenS1{F0: ("hello" + ""), F1: 0}
                _ = fi37
                _ = v38
            }
            _ = fi36
        }
    } else {
        var v39 uint8 = v6
        for fi40 := 0; fi40 < 5; fi40++ {
            var v41 uint8 = 127
            switch fi40 {
            case 216:
                fmt.Println(v39, v6, v6)
            case 4294967295:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case -264:
                for fi42 := 0; fi42 < 4; fi42++ {
                    v43 := GenS2{F0: 719, F1: true}
                    v6 -= 207
                    v44 := GenS1{F0: "foo bar", F1: uint32(v34)}
                    _ = fi42
                    _ = v43
                    _ = v44
                }
            default:
                v39 = v39
            }
            v34 = v34
            _ = fi40
            _ = v41
        }
        _ = v39
    }
    if true {
        var v45 int16 = 515
        fmt.Println(v5, v5)
        v34 += v34
        _ = v45
    } else {
        var v46 int = 9223372036854775807
        if (108 != 65535) {
            var v47 string = ("hello" + "a")
            var v48 bool = (!((true || (!true)) || false))
            _ = v47
            _ = v48
        } else {
            switch 128 {
            case 42:
                v49 := GenS2{F0: 371, F1: true}
                _ = v49
            }
            var v50 int64 = (v5 - (-(v5)))
            _ = v50
        }
        _ = v46
    }
    v4 = append(v4, 7)
    v5 = (v5 / (((^(v5))) | 1))
    for fi51 := 0; fi51 < 4; fi51++ {
        v52 := GenS6{F0: 749, F1: 427}
        v1 += v1
        v1 = v1
        _ = fi51
        _ = v52
    }
    if len(v4) > 0 {
        v4[0] = 916
    }
    for fi53 := 0; fi53 < 17; fi53++ {
        switch (-(fi53)) {
        case 127:
            if true {
                for fi54 := 0; fi54 < 20; fi54++ {
                    fmt.Println(v34, fi53)
                    v6 = (v6 - v6)
                    var v55 bool = true
                    _ = fi54
                    _ = v55
                }
            }
        case 2147483647:
            var v56 uint8 = 255
            _ = v56
        default:
            v57 := GenS6{F0: 42, F1: 2}
            _ = v57
        }
        if false {
            v5 = -9223372036854775808
            fmt.Println(v6, fi53, fi53)
        }
        var v58 int = (fi53 * fi53)
        _ = fi53
        _ = v58
    }
    v2 = append(v2, "a")
    v59 := GenS7{F0: false}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v59
    _ = v6
}
