package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint32; F2 int64 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int; F1 int32; F2 int8 }
type GenS8 struct { F0 uint64; F1 uint32 }
type GenS9 struct { F0 uint64 }
func genFn0(p0 uint32, p1 uint16, p2 uint64) uint {
    p0 = (p0 + 4294967295)
    v7 := GenS1{F0: 128, F1: (!(!false)), F2: 1}
    p2 += 1
    switch 42 {
    case 42:
        p2 = (p2 ^ 4294967295)
    case -1:
        var v8 uint64 = p2
        _ = v8
    default:
        fmt.Println(p2)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    var v1 float64 = ((1.5 - 0.0) / (float64((-2.25 / 1.5)) / 1.0))
    var v2 string = ("" + ("\n\t" + ""))
    v3 := make(map[uint32]string)
    var v4 int16 = -741
    var v5 uint64 = 42
    var v6 uint32 = 2147483647
    switch "" {
    case "\n\t":
        if false {
            switch "foo bar" {
            case "foo bar":
                var v9 string = v2
                _ = v9
            }
            var v10 string = v2
            if (false && ((!false) && false)) {
                v6 = v6
            }
            _ = v10
        } else {
            fmt.Println(v6)
            v11 := GenS1{F0: 128, F1: (true || (!((!((v4 >> ((700) & 63)) > (v4 & (v4 / ((int16(v6)) | 1))))) || false))), F2: 22}
            _ = v11
        }
    case "a":
        v6 -= 42
    }
    fmt.Println(v2)
    v12 := GenS2{F0: 1, F1: v6, F2: 0}
    v5 = v5
    v13 := []int{-1, 32768, 1, 1}
    v3[(^(v6))] = ((("a" + ("" + v2)) + "a") + v2)
    v14 := GenS1{F0: 113, F1: false, F2: 93}
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        for fi16 := 0; fi16 < 9; fi16++ {
            fmt.Println(fi15, v1)
            _ = fi16
        }
        _ = len(v13)
        switch (fi15 + -9223372036854775808) {
        case -1:
            if (!(((v4 + v4) != 1) && false)) {
                if false {
                    if len(v13) > 0 {
                        v13[0] = fi15
                    }
                    fmt.Println(v2)
                    v1 = (float64(0.0) + float64(0.1))
                } else {
                    v1 += (-2.25 + 1.5)
                    if len(v13) > 0 {
                        v13[0] = (^(fi15))
                    }
                    fmt.Println(v2)
                }
            }
        case 42:
            v3[v6] = "hello"
        default:
            if false {
                if false {
                    v17 := GenS3{F0: -9223372036854775808, F1: 65535}
                    fmt.Println(v1, v2)
                    v2 = v2
                    _ = v17
                } else {
                    v18 := GenS4{F0: 42, F1: 0}
                    fmt.Println(fi15, v4, v1)
                    _ = v18
                }
                _ = len(v3)
            }
        }
        _ = fi15
    }
    var v19 int16 = -32768
    switch v2 {
    case "a":
        fmt.Println(v1, v4)
    case "hello":
        switch (v2 + v2) {
        case "a":
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
        case "":
            v6 = v6
        case "hello":
            fmt.Println(v4, v5)
        }
    case "\n\t":
        switch "" {
        case "":
            if (1.5 <= 1e10) {
                for fi20 := 0; fi20 < 8; fi20++ {
                    v21 := GenS3{F0: 42, F1: 1}
                    v13 = append(v13, fi20)
                    _ = fi20
                    _ = v21
                }
                switch v2 {
                case "\n\t":
                    if (false && (!((!((v5 >> ((255) & 63)) <= v5)) || true))) {
                        v22 := GenS5{F0: 0, F1: 127}
                        v1 -= float64(0.1)
                        for fi23 := 0; fi23 < 6; fi23++ {
                            var v24 uint64 = 4294967295
                            if false {
                                for fi25 := 0; fi25 < 6; fi25++ {
                                    delete(v3, 42)
                                    v4 += (v19 * v4)
                                    _ = fi25
                                }
                                for fi26 := 0; fi26 < 17; fi26++ {
                                    v4 -= (v4 & v4)
                                    _ = fi26
                                }
                            }
                            _ = len(v13)
                            _ = fi23
                            _ = v24
                        }
                        _ = v22
                    } else {
                        if true {
                            v2 = v2
                            v4 = v19
                        } else {
                            var v27 string = v2
                            _ = v27
                        }
                    }
                case "hello":
                    v2 = ((v2 + "a") + "a")
                }
                if false {
                    if (uint16(v6) >= 523) {
                        for fi28 := 0; fi28 < 10; fi28++ {
                            var v29 bool = (!true)
                            if (65535 == 179) {
                                if (v29 && v29) {
                                    _ = len(v13)
                                    if ((fi28 * fi28) <= fi28) {
                                        var v30 int16 = 32767
                                        v29 = (42 <= 0)
                                        _ = v30
                                    }
                                    if v29 {
                                        if (false && v29) {
                                            var v31 int = (-(fi28))
                                            if len(v13) >= 2 {
                                                v13 = v13[1:len(v13)]
                                            }
                                            switch "" {
                                            case "hello":
                                                v4 = (v4 & 0)
                                            default:
                                                var v32 float32 = float32(v1)
                                                _ = v32
                                            }
                                            _ = v31
                                        } else {
                                            v4 -= v4
                                            v33 := GenS3{F0: -943, F1: 162}
                                            var v34 uint64 = v5
                                            _ = v33
                                            _ = v34
                                        }
                                    }
                                } else {
                                    var v35 uint16 = 255
                                    switch v2 {
                                    case "foo bar":
                                        for fi36 := 0; fi36 < 1; fi36++ {
                                            fmt.Println(v19)
                                            switch (fi36 % ((404) | 1)) {
                                            case -76:
                                                fmt.Println(v6, v1, v1)
                                            case 734:
                                                v35 = (v35 & v35)
                                            case 32767:
                                                if v29 {
                                                    v4 -= (v19 | (v19 % ((v4) | 1)))
                                                    switch v2 {
                                                    case "foo bar":
                                                        if (v29 && v29) {
                                                            if v29 {
                                                                if v29 {
                                                                    var v37 int8 = 1
                                                                    _ = v37
                                                                }
                                                                v38 := GenS6{F0: 0}
                                                                v35 += v35
                                                                _ = v38
                                                            } else {
                                                                var v39 uint32 = 2147483647
                                                                var v40 bool = true
                                                                _ = v39
                                                                _ = v40
                                                            }
                                                        }
                                                    case "":
                                                        v5 = v5
                                                    }
                                                    switch (("hello" + "hello") + "\n\t") {
                                                    case "\n\t":
                                                        v41 := GenS7{F0: (fi36 % (((fi36 + (fi36 & fi36))) | 1)), F1: 0, F2: 1}
                                                        _ = v41
                                                    case "hello":
                                                        if ((1 < 42) || true) {
                                                            v1 = 0.1
                                                            var v42 uint = uint(v35)
                                                            v43 := v3[(^(v6))]
                                                            _ = v42
                                                            _ = v43
                                                        } else {
                                                            var v44 uint32 = (v6 & (v6 << ((381) & 63)))
                                                            fmt.Println(v44)
                                                            _ = v44
                                                        }
                                                    default:
                                                        fmt.Println(v4, v29)
                                                        var v45 float64 = float64((0.0 * (1e10 * 1.5)))
                                                        v19 -= (v19 - (v4 | -1))
                                                        _ = v45
                                                    }
                                                } else {
                                                    var v46 uint = 42
                                                    v13 = append(v13, (fi28 >> ((genFn0((v6 * (^(v6))), (v35 % (((v35 - v35)) | 1)), (v5 >> ((v46) & 63)))) & 63)))
                                                    v29 = (0 >= 2147483647)
                                                    _ = v46
                                                }
                                            default:
                                                var v47 int32 = 42
                                                _ = v47
                                            }
                                            _ = fi36
                                        }
                                    case "":
                                        v6 = 56
                                    case "hello":
                                        if v29 {
                                            var v48 uint = 18446744073709551615
                                            for fi49 := 0; fi49 < 19; fi49++ {
                                                v4 = v19
                                                var v50 bool = v29
                                                _ = fi49
                                                _ = v50
                                            }
                                            fmt.Println(fi28, v48, v48)
                                            _ = v48
                                        }
                                    default:
                                        fmt.Println(v6)
                                    }
                                    _ = v35
                                }
                                v51 := GenS8{F0: (v5 ^ 1), F1: (v6 + (v6 / (((v6 + (v6 / ((v6) | 1)))) | 1)))}
                                if false {
                                    v1 = -2.25
                                    var v52 float32 = (1e10 / 3.14159)
                                    var v53 int8 = 17
                                    _ = v52
                                    _ = v53
                                } else {
                                    v54 := GenS5{F0: 255, F1: 42}
                                    var v55 string = (v2 + "\n\t")
                                    _ = v54
                                    _ = v55
                                }
                                _ = v51
                            }
                            for fi56 := 0; fi56 < 18; fi56++ {
                                v19 -= v4
                                _ = fi56
                            }
                            _ = fi28
                            _ = v29
                        }
                        if len(v13) > 0 {
                            v13[0] = -585
                        }
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                    } else {
                        v4 = -1
                        if (255 == 208) {
                            for fi57 := 0; fi57 < 4; fi57++ {
                                var v58 float32 = (0.1 * -2.25)
                                v1 = 1e10
                                var v59 float64 = (1e10 + 3.14159)
                                _ = fi57
                                _ = v58
                                _ = v59
                            }
                        } else {
                            v5 -= (v5 | 1)
                            var v60 float32 = float32(0.1)
                            _ = v60
                        }
                    }
                } else {
                    v13 = append(v13, 245)
                }
            } else {
                v61 := GenS9{F0: 778}
                v5 = 42
                _ = v61
            }
        case "hello":
            if false {
                for fi62 := 0; fi62 < 15; fi62++ {
                    if false {
                        switch v2 {
                        case "a":
                            fmt.Println(v19, v19, v2)
                        case "\n\t":
                            var v63 uint64 = 529
                            _ = v63
                        }
                    }
                    _ = fi62
                }
                var v64 uint8 = 1
                _ = v64
            }
        }
    default:
        v3[4294967295] = (v2 + ("a" + (v2 + (("hello" + ("foo bar" + "hello")) + (v2 + ("a" + v2))))))
    }
    var v65 uint16 = 342
    switch -955 {
    case -2147483648:
        for fi66 := 0; fi66 < 9; fi66++ {
            fmt.Println(v6, v2)
            for fi67 := 0; fi67 < 10; fi67++ {
                delete(v3, 2)
                v65 -= v65
                if (255 != 127) {
                    if false {
                        var v68 float32 = float32((((v1 + v1) / (v1 * 1e10)) - 1.5))
                        var v69 string = (v2 + v2)
                        var v70 int8 = 1
                        _ = v68
                        _ = v69
                        _ = v70
                    }
                }
                _ = fi67
            }
            _ = fi66
        }
    }
    v71 := v3[v6]
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v71
}
