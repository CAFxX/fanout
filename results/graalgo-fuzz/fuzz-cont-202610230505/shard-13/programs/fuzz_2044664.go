package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int; F1 bool; F2 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 int16; F2 int64 }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 int8; F1 int; F2 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 552
    v3 := []uint32{2147483647, 2147483648}
    var v4 int32 = 842
    v5 := []uint32{42}
    var v6 uint32 = 4294967295
    fmt.Println(v6, v6, v2)
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := GenS1{F0: true, F1: 0}
            v2 = v2
            _ = fi8
            _ = v9
        }
        fmt.Println(v2)
        var v10 uint32 = (v6 >> ((966) & 63))
        _ = fi7
        _ = v10
    }
    var v11 uint = 826
    for fi12 := 0; fi12 < 18; fi12++ {
        if (42 != 42) {
            var v13 int64 = -9223372036854775808
            for fi14 := 0; fi14 < 17; fi14++ {
                if (-1 < 42) {
                    switch 2147483647 {
                    case 689:
                        if false {
                            switch ("\n\t" + "hello") {
                            case "\n\t":
                                v11 -= 1
                            case "hello":
                                if (!(!false)) {
                                    v2 = (v2 % ((v2) | 1))
                                    var v15 int8 = 113
                                    v13 = (v13 >> ((255) & 63))
                                    v16 := GenS1{F0: (0 > 42), F1: v6}
                                    v11 = (v11 + uint(v6))
                                    _ = v15
                                    _ = v16
                                } else {
                                    var v17 int8 = 24
                                    var v18 string = "\n\t"
                                    var v19 int16 = 1000
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                }
                            case "":
                                v13 += (v13 << (((v11 >> ((v11) & 63))) & 63))
                            default:
                                var v20 uint64 = 2
                                switch -1 {
                                case -9223372036854775808:
                                    v21 := GenS1{F0: true, F1: (v6 ^ 1)}
                                    _ = v21
                                }
                                if ((3.14159 != 0.1) && false) {
                                    v13 = (^(v13))
                                    if true {
                                        if true {
                                            fmt.Println(v6, v13, v20)
                                            v20 -= (v20 + (v2 << ((v11) & 63)))
                                        }
                                    } else {
                                        v6 = uint32(v20)
                                        v13 = (v13 - v13)
                                    }
                                } else {
                                    v22 := GenS2{F0: fi12, F1: ((v6 * v6) < (v6 + (^(v6)))), F2: 891}
                                    switch "a" {
                                    case "a":
                                        if (false || ((true && (true && false)) && true)) {
                                            _ = len(v1)
                                            fmt.Println(v20, fi12, v11)
                                        } else {
                                            fmt.Println(fi12)
                                            v6 = v6
                                        }
                                    case "":
                                        v1[v6] = v6
                                    default:
                                        switch (("" + "") + "hello") {
                                        case "foo bar":
                                            var v23 uint32 = uint32(v11)
                                            _ = v23
                                        default:
                                            fmt.Println(fi14)
                                        }
                                    }
                                    fmt.Println(fi14)
                                    _ = v22
                                }
                                var v24 uint16 = 65535
                                _ = v20
                                _ = v24
                            }
                            if true {
                                v11 = 1
                                v25 := GenS1{F0: ((^(v11)) == (v11 % ((65535) | 1))), F1: v6}
                                if false {
                                    fmt.Println(v11, fi14)
                                } else {
                                    fmt.Println(v4)
                                }
                                fmt.Println(v2)
                                v4 += -2147483648
                                _ = v25
                            }
                            v26 := v1[(^(v6))]
                            _ = v26
                        } else {
                            var v27 uint8 = uint8(v2)
                            var v28 uint32 = v6
                            var v29 uint = 65535
                            var v30 uint64 = (^(v2))
                            var v31 uint16 = 417
                            _ = v27
                            _ = v28
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                    case 7:
                        var v32 uint = (v11 % ((65535) | 1))
                        _ = v32
                    case 2147483647:
                        v13 = (^(v13))
                    default:
                        delete(v1, v6)
                    }
                    switch "foo bar" {
                    case "a":
                        v33 := GenS3{F0: v13}
                        _ = v33
                    }
                } else {
                    fmt.Println(v11, v2)
                    switch ("\n\t" + ("foo bar" + "hello")) {
                    case "":
                        v34 := GenS4{F0: v13, F1: -1, F2: (v13 % (((v13 >> ((v11) & 63))) | 1))}
                        _ = v34
                    default:
                        var v35 int = (fi12 + fi14)
                        _ = v35
                    }
                }
                _ = len(v3)
                _ = fi14
            }
            _ = v13
        }
        _ = fi12
    }
    for fi36 := 0; fi36 < 10; fi36++ {
        if false {
            if (399 == 42) {
                if ((fi36 << ((261) & 63)) < fi36) {
                    switch "a" {
                    case "hello":
                        for fi37 := 0; fi37 < 19; fi37++ {
                            fmt.Println(v11, fi36, v6)
                            _ = fi37
                        }
                    case "":
                        var v38 uint = uint(v2)
                        _ = v38
                    case "foo bar":
                        fmt.Println(v4)
                    default:
                        v39 := GenS1{F0: false, F1: uint32(v11)}
                        _ = v39
                    }
                }
                for fi40 := 0; fi40 < 20; fi40++ {
                    if (-1 != int64(v2)) {
                        if (!false) {
                            switch -9223372036854775808 {
                            case 128:
                                fmt.Println(v6, v11)
                            case 9223372036854775807:
                                fmt.Println(fi40)
                            case 32767:
                                v6 -= 1
                            }
                            var v41 int64 = 42
                            var v42 uint64 = 807
                            _ = v41
                            _ = v42
                        }
                    } else {
                        v6 = uint32(v2)
                        if false {
                            if ((((3.14159 * 0.0) * 3.14159) + 1e10) != -2.25) {
                                if ("hello" >= "a") {
                                    fmt.Println(fi36, v4)
                                    var v43 int32 = v4
                                    _ = v43
                                } else {
                                    v11 = (v11 + v11)
                                }
                            } else {
                                v44 := GenS3{F0: -238}
                                _ = v44
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v45 int32 = v4
                            _ = v45
                        } else {
                            v5 = append(v5, (v6 + (v6 * 166)))
                        }
                        var v46 uint64 = (v2 + v2)
                        _ = v46
                    }
                    var v47 int32 = v4
                    _ = fi40
                    _ = v47
                }
            }
            v2 = (v2 - (v2 + (v2 / ((355) | 1))))
            if (true || true) {
                v48 := GenS1{F0: false, F1: v6}
                _ = v48
            }
        }
        _ = fi36
    }
    var v49 uint8 = 128
    switch "" {
    case "":
        var v50 int8 = 111
        _ = v50
    }
    switch 1 {
    case 9223372036854775807:
        var v51 uint64 = 4294967295
        _ = v51
    }
    v2 -= v2
    v5 = append(v5, (v6 >> (((v11 - v11)) & 63)))
    var v52 int32 = -373
    v53 := GenS5{F0: v6, F1: false}
    v4 += -1
    var v54 uint32 = v6
    if false {
        fmt.Println(v54, v11, v6)
    } else {
        v5 = append(v5, 2147483647)
        if false {
            _ = len(v3)
            v55 := GenS1{F0: (v2 == (v2 + (v2 - v2))), F1: v54}
            var v56 uint = (v11 & (v11 | v11))
            _ = v55
            _ = v56
        } else {
            for fi57 := 0; fi57 < 8; fi57++ {
                switch ("a" + "") {
                case "foo bar":
                    v4 = -1
                case "\n\t":
                    _ = len(v1)
                default:
                    v2 = 0
                }
                v58 := GenS2{F0: (fi57 - -131), F1: (((^(v11)) < (v11 >> (((^(v11))) & 63))) && ((!(65535 >= 1)) || false)), F2: v54}
                v59 := v1[v6]
                _ = fi57
                _ = v58
                _ = v59
            }
        }
        for fi60 := 0; fi60 < 16; fi60++ {
            v61 := GenS6{F0: 127, F1: (fi60 << ((v11) & 63)), F2: 602}
            v62 := v1[v6]
            v54 = 2147483648
            _ = fi60
            _ = v61
            _ = v62
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v6
}
