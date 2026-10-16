package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint16; F2 bool }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64; F1 uint64 }

func main() {
    v1 := []int{448}
    v2 := []uint32{2147483648}
    var v3 uint32 = 42
    var v4 uint8 = 116
    v5 := []uint32{(v3 - v3), (^(v3)), 65535, v3}
    v3 = (v3 ^ 2147483647)
    var v6 bool = true
    if (-2147483648 == 572) {
        if ((v6 || true) || false) {
            for fi7 := 0; fi7 < 7; fi7++ {
                _ = len(v2)
                _ = fi7
            }
            if v6 {
                switch -9223372036854775808 {
                case 4294967295:
                    v6 = (9223372036854775807 != 769)
                case -231:
                    for fi8 := 0; fi8 < 13; fi8++ {
                        switch 0 {
                        case -80:
                            v6 = (v6 && (1 >= 42))
                        default:
                            for fi9 := 0; fi9 < 11; fi9++ {
                                v10 := GenS1{F0: fi8}
                                v11 := GenS1{F0: 149}
                                _ = len(v5)
                                _ = fi9
                                _ = v10
                                _ = v11
                            }
                        }
                        v4 = 128
                        _ = fi8
                    }
                }
            }
            var v12 uint32 = (v3 + (v3 + 2147483648))
            _ = v12
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v3 = (v3 / ((v3) | 1))
            var v13 string = ("" + "")
            _ = v13
        }
    }
    v6 = (255 < 654)
    v3 += v3
    v3 += 990
    v4 -= (^(v4))
    fmt.Println(v6)
    switch -368 {
    case -1:
        v6 = (!(!(v3 < v3)))
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 int = 1
        switch 0 {
        case 42:
            if true {
                var v16 string = "a"
                if (0 == -66) {
                    var v17 uint64 = 4294967295
                    _ = v17
                } else {
                    v3 = (v3 * v3)
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v15 = fi14
                }
                _ = v16
            }
        }
        for fi18 := 0; fi18 < 14; fi18++ {
            var v19 uint64 = 0
            _ = fi18
            _ = v19
        }
        _ = fi14
        _ = v15
    }
    var v20 uint8 = 130
    v6 = v6
    v3 -= v3
    switch 128 {
    case -1:
        v20 += (v4 & (v4 / ((v20) | 1)))
    case 2147483647:
        var v21 uint16 = 149
        _ = v21
    }
    var v22 string = "foo bar"
    fmt.Println(v22, v3, v20)
    for fi23 := 0; fi23 < 6; fi23++ {
        if (float32(3.14159) > 0.0) {
            switch -506 {
            case 128:
                v6 = ((-2.25 == (0.1 + 0.0)) || (32767 <= 32767))
            default:
                for fi24 := 0; fi24 < 2; fi24++ {
                    var v25 int64 = -39
                    switch -511 {
                    case -1:
                        switch fi24 {
                        case 4294967295:
                            if v6 {
                                var v26 uint32 = (v3 << ((969) & 63))
                                if false {
                                    if (!v6) {
                                        v22 = "hello"
                                    }
                                    if (false && v6) {
                                        if v6 {
                                            v27 := GenS2{F0: true, F1: 2, F2: (65535 <= 1)}
                                            switch v22 {
                                            case "\n\t":
                                                _ = len(v2)
                                            case "hello":
                                                v28 := GenS3{F0: true, F1: 243}
                                                _ = v28
                                            case "a":
                                                var v29 int64 = (v25 << ((50) & 63))
                                                _ = v29
                                            default:
                                                switch 42 {
                                                case 1:
                                                    v26 = v26
                                                default:
                                                    var v30 uint8 = (v4 >> ((65535) & 63))
                                                    var v31 uint = 2
                                                    _ = v30
                                                    _ = v31
                                                }
                                            }
                                            v3 = 806
                                            _ = v27
                                        }
                                        if ((("" + v22) > (v22 + v22)) || v6) {
                                            if true {
                                                v32 := GenS1{F0: 127}
                                                var v33 float32 = -2.25
                                                _ = v32
                                                _ = v33
                                            }
                                            v20 = 42
                                        }
                                        fmt.Println(v22)
                                    }
                                    v4 = v20
                                } else {
                                    var v34 string = "\n\t"
                                    v1 = append(v1, fi24)
                                    v3 += v26
                                    _ = v34
                                }
                                _ = v26
                            }
                            v20 -= 42
                            switch v22 {
                            case "a":
                                _ = len(v5)
                            }
                            v20 = v20
                        }
                    case 2147483647:
                        var v35 uint64 = 1
                        _ = v35
                    case 32767:
                        if v6 {
                            var v36 int64 = -243
                            _ = v36
                        } else {
                            v37 := GenS4{F0: 281}
                            var v38 uint32 = v3
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v37
                            _ = v38
                        }
                    default:
                        _ = len(v5)
                    }
                    switch ("foo bar" + v22) {
                    case "foo bar":
                        var v39 string = v22
                        _ = v39
                    default:
                        _ = len(v5)
                        var v40 uint64 = 2
                        v25 = (v25 % (((v25 / ((607) | 1))) | 1))
                        switch "\n\t" {
                        case "a":
                            fmt.Println(v3, fi24)
                        case "":
                            switch v22 {
                            case "\n\t":
                                var v41 uint32 = uint32(v20)
                                _ = v41
                            case "":
                                fmt.Println(v20, v3, fi23)
                                v2 = append(v2, v3)
                                if (!v6) {
                                    switch -9223372036854775808 {
                                    case 42:
                                        switch v22 {
                                        case "a":
                                            fmt.Println(v6, fi24, v4)
                                        case "hello":
                                            if v6 {
                                                v4 -= (v4 + (v20 + (v20 % ((uint8(v3)) | 1))))
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                        }
                                    case 67:
                                        switch fi23 {
                                        case 128:
                                            _ = len(v5)
                                        }
                                    }
                                    fmt.Println(v3, v25)
                                    var v42 int32 = 0
                                    var v43 uint = 0
                                    _ = v42
                                    _ = v43
                                }
                                var v44 uint32 = (v3 >> ((981) & 63))
                                if len(v2) > 0 {
                                    v2[0] = v44
                                }
                                v4 -= (v20 + v4)
                                _ = v44
                            default:
                                switch fi24 {
                                case -545:
                                    var v45 uint = 65535
                                    v25 += (v25 >> ((v45) & 63))
                                    _ = v45
                                }
                            }
                        }
                        _ = v40
                    }
                    v46 := GenS3{F0: v6, F1: (v4 << ((4294967295) & 63))}
                    _ = fi24
                    _ = v25
                    _ = v46
                }
            }
            var v47 float32 = 0.0
            for fi48 := 0; fi48 < 13; fi48++ {
                v49 := GenS5{F0: 4294967295, F1: 42}
                fmt.Println(v47)
                _ = fi48
                _ = v49
            }
            _ = v47
        }
        v20 = v4
        if (v6 && false) {
            _ = len(v2)
            v4 = v20
        }
        _ = fi23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
