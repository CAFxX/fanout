package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 bool }
type GenS6 struct { F0 string; F1 string }
type GenS7 struct { F0 uint32 }
func genFn0(p0 int16, p1 int32, p2 int) int32 {
    p2 = (p2 * (p2 & (p2 << ((255) & 63))))
    v7 := GenS1{F0: 0, F1: -159}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p1 + 1)
}

func main() {
    var v1 int = 128
    v2 := make(map[uint32]string)
    var v3 string = "hello"
    v4 := make(map[string]string)
    v5 := make(map[int]string)
    var v6 uint = 83
    fmt.Println(v6, v3, v1)
    if (2 != 0) {
        v8 := GenS1{F0: -9223372036854775808, F1: -1}
        var v9 float32 = float32(0.0)
        _ = v8
        _ = v9
    } else {
        _ = len(v5)
        if true {
            for fi10 := 0; fi10 < 6; fi10++ {
                v11 := []int{}
                _ = fi10
                _ = v11
            }
            v12 := GenS2{F0: 1}
            v1 = v1
            _ = v12
        } else {
            for fi13 := 0; fi13 < 14; fi13++ {
                v1 = v1
                fmt.Println(v3)
                _ = fi13
            }
            switch -9223372036854775808 {
            case 128:
                for fi14 := 0; fi14 < 6; fi14++ {
                    for fi15 := 0; fi15 < 11; fi15++ {
                        var v16 bool = (139 != 1)
                        v17 := GenS3{F0: -2147483648, F1: v6}
                        v18 := GenS2{F0: 32767}
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    var v19 float32 = float32(0.0)
                    _ = fi14
                    _ = v19
                }
            default:
                for fi20 := 0; fi20 < 13; fi20++ {
                    v21 := GenS1{F0: 0, F1: -1}
                    v22 := GenS4{F0: v3}
                    var v23 uint64 = 4294967295
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
            }
        }
    }
    var v24 uint32 = 2147483648
    if true {
        for fi25 := 0; fi25 < 15; fi25++ {
            if ((v24 | 42) < v24) {
                var v26 uint = (^(v6))
                _ = v26
            } else {
                if true {
                    v6 -= v6
                    v27 := GenS4{F0: v3}
                    _ = v27
                }
                var v28 bool = true
                for fi29 := 0; fi29 < 18; fi29++ {
                    v6 = v6
                    v30 := GenS2{F0: 850}
                    _ = fi29
                    _ = v30
                }
                _ = v28
            }
            for fi31 := 0; fi31 < 3; fi31++ {
                v3 = v3
                v32 := v5[fi25]
                v1 = -9223372036854775808
                var v33 uint8 = 175
                if true {
                    v34 := GenS1{F0: 106, F1: -9223372036854775808}
                    v24 += v24
                    _ = v34
                } else {
                    if false {
                        var v35 int32 = genFn0(-1, 1, -622)
                        _ = v35
                    } else {
                        v36 := v4[v32]
                        _ = v36
                    }
                    v37 := GenS5{F0: 42, F1: (2 >= 2)}
                    v32 = v3
                    _ = v37
                }
                _ = fi31
                _ = v32
                _ = v33
            }
            _ = fi25
        }
    } else {
        fmt.Println(v3)
        for fi38 := 0; fi38 < 10; fi38++ {
            var v39 int16 = 1
            _ = fi38
            _ = v39
        }
        switch v1 {
        case -2147483648:
            v40 := GenS1{F0: 9223372036854775807, F1: -371}
            _ = v40
        case 1:
            v24 = uint32(v6)
        case 484:
            v41 := GenS4{F0: v3}
            _ = v41
        default:
            if true {
                if true {
                    for fi42 := 0; fi42 < 5; fi42++ {
                        v43 := GenS4{F0: ("foo bar" + v3)}
                        _ = fi42
                        _ = v43
                    }
                    fmt.Println(v24, v6)
                } else {
                    v24 = (v24 << (((v6 >> ((v6) & 63))) & 63))
                }
                switch v3 {
                case "\n\t":
                    var v44 uint16 = 1
                    _ = v44
                case "hello":
                    fmt.Println(v24, v6, v1)
                case "foo bar":
                    v1 -= (v1 & v1)
                }
            } else {
                var v45 bool = (4294967295 < 4294967295)
                _ = v45
            }
        }
    }
    switch v3 {
    case "a":
        if true {
            fmt.Println(v6, v3, v6)
        } else {
            v6 += 408
        }
    case "hello":
        var v46 float32 = 1.5
        _ = v46
    case "":
        if (!(true && (v3 <= "a"))) {
            var v47 float32 = -2.25
            var v48 uint = v6
            _ = v47
            _ = v48
        }
    }
    v3 = (("foo bar" + (v3 + "\n\t")) + (((v3 + v3) + (v3 + v3)) + "a"))
    var v49 int16 = 1000
    v3 = "foo bar"
    fmt.Println(v6)
    switch v3 {
    case "a":
        var v50 uint32 = (^(v24))
        _ = v50
    case "\n\t":
        switch v3 {
        case "":
            for fi51 := 0; fi51 < 5; fi51++ {
                v52 := GenS2{F0: 42}
                _ = fi51
                _ = v52
            }
        case "\n\t":
            switch 1 {
            case 128:
                v6 = v6
            case 32767:
                var v53 int64 = 42
                _ = v53
            case 127:
                v49 = (v49 + v49)
            default:
                var v54 int64 = 9223372036854775807
                _ = v54
            }
        default:
            v55 := GenS2{F0: 42}
            _ = v55
        }
    }
    v3 = v3
    for fi56 := 0; fi56 < 6; fi56++ {
        _ = len(v5)
        _ = fi56
    }
    if (true || ((212 != 116) && true)) {
        if (true && true) {
            v1 += (v1 << ((0) & 63))
            _ = len(v2)
        }
    }
    var v57 uint8 = 255
    v58 := GenS6{F0: "", F1: v3}
    if false {
        var v59 int8 = 0
        _ = v59
    } else {
        fmt.Println(v6)
        v6 = (v6 - (v6 - v6))
    }
    switch -718 {
    case 9223372036854775807:
        var v60 uint = v6
        _ = v60
    case 32767:
        switch v3 {
        case "a":
            v1 = (v1 | 4294967295)
        case "":
            v1 = v1
        case "foo bar":
            if (1.5 >= 0.0) {
                v4[(v3 + "\n\t")] = ((((v3 + (("\n\t" + (("\n\t" + "hello") + v3)) + "hello")) + v3) + v3) + v3)
                var v61 float32 = 3.14159
                _ = v61
            }
        }
    case -9223372036854775808:
        if (7 >= 11) {
            if (genFn0((v49 ^ 1), 713, (-(v1))) != 0) {
                var v62 string = v3
                _ = v62
            } else {
                for fi63 := 0; fi63 < 5; fi63++ {
                    switch int(v57) {
                    case 2147483647:
                        for fi64 := 0; fi64 < 3; fi64++ {
                            if (42 <= -907) {
                                v65 := []int{(fi63 & 384)}
                                var v66 int32 = genFn0(0, 813, -2147483648)
                                fmt.Println(v24, fi64)
                                _ = v65
                                _ = v66
                            } else {
                                var v67 int64 = -1
                                _ = v67
                            }
                            var v68 int16 = v49
                            _ = fi64
                            _ = v68
                        }
                    case 97:
                        var v69 uint8 = 42
                        _ = v69
                    default:
                        var v70 uint = v6
                        _ = v70
                    }
                    _ = fi63
                }
                v1 -= (v1 * v1)
                v71 := GenS1{F0: 42, F1: -1}
                _ = v71
            }
        } else {
            for fi72 := 0; fi72 < 7; fi72++ {
                fmt.Println(v49, v49)
                _ = fi72
            }
            v1 = v1
        }
    }
    if ((!(true || true)) && false) {
        switch ("" + (v3 + "")) {
        case "a":
            if true {
                v73 := GenS7{F0: (v24 & 42)}
                _ = v73
            }
        case "hello":
            v74 := []int{v1, v1, 4294967295, 9223372036854775807}
            _ = v74
        }
        v75 := []int{}
        v49 -= (-(v49))
        _ = v75
    } else {
        v6 = v6
        for fi76 := 0; fi76 < 4; fi76++ {
            if (genFn0(v49, genFn0(v49, -1, v1), fi76) <= 2147483647) {
                for fi77 := 0; fi77 < 18; fi77++ {
                    v3 = ((v3 + "\n\t") + v3)
                    _ = fi77
                }
                switch "" {
                case "a":
                    v6 = v6
                case "hello":
                    v78 := GenS3{F0: -263, F1: v6}
                    _ = v78
                }
                var v79 uint = (v6 - 465)
                _ = v79
            }
            _ = fi76
        }
        v80 := GenS2{F0: 712}
        _ = v80
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v57
    _ = v58
    _ = v6
}
