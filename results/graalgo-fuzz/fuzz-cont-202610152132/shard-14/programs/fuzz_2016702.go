package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 uint32; F1 int32 }

func main() {
    var v1 int16 = -599
    v2 := make(map[uint32]string)
    v3 := make(map[int]int)
    var v4 int8 = -1
    var v5 int32 = 2147483647
    if true {
        delete(v3, 2147483647)
        var v6 uint = 0
        if true {
            for fi7 := 0; fi7 < 17; fi7++ {
                var v8 float32 = (3.14159 / float32(0.1))
                var v9 int = 32768
                v10 := []int{(v9 + (v9 ^ 9223372036854775807)), v9, fi7}
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
        }
        _ = v6
    } else {
        if ((0.0 < -2.25) || false) {
            v4 += (v4 - v4)
        }
        v1 = (v1 + 32767)
    }
    if false {
        _ = len(v3)
    } else {
        fmt.Println(v5, v5)
        var v11 uint8 = 174
        for fi12 := 0; fi12 < 6; fi12++ {
            v3[(fi12 >> ((984) & 63))] = 853
            _ = fi12
        }
        _ = v11
    }
    v13 := []int{9223372036854775807}
    delete(v3, 2147483647)
    if false {
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 int32 = (v5 << ((4294967295) & 63))
            v15 += v5
            var v16 float32 = float32((3.14159 - -2.25))
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    _ = len(v13)
    switch "\n\t" {
    case "a":
        _ = len(v13)
    case "\n\t":
        switch "" {
        case "foo bar":
            var v17 int8 = 42
            _ = v17
        case "hello":
            v18 := GenS1{F0: 32767, F1: 122, F2: 9223372036854775807}
            _ = v18
        case "":
            for fi19 := 0; fi19 < 9; fi19++ {
                v20 := GenS1{F0: 1, F1: 128, F2: -1}
                _ = fi19
                _ = v20
            }
        default:
            v21 := v2[1]
            _ = v21
        }
    }
    v5 -= 42
    var v22 uint32 = 65535
    v23 := GenS1{F0: 0, F1: 127, F2: -3}
    v24 := GenS1{F0: 32767, F1: 26, F2: 979}
    v5 = v5
    v4 = v4
    if (("" < ("" + ("foo bar" + "\n\t"))) || ("a" <= "foo bar")) {
        if false {
            v25 := GenS2{F0: (float32((3.14159 - 3.14159)) < 0.0)}
            if true {
                if (!true) {
                    v4 += 42
                    v5 -= (v5 + int32(v22))
                }
                v5 -= v5
                v26 := v2[65535]
                _ = v26
            } else {
                v4 = (v4 + v4)
            }
            _ = v25
        } else {
            switch ("a" + "\n\t") {
            case "":
                v5 -= (^(v5))
            case "\n\t":
                var v27 uint32 = (v22 << ((4294967295) & 63))
                _ = v27
            case "foo bar":
                if false {
                    var v28 uint = 4294967295
                    var v29 uint8 = 1
                    switch 0 {
                    case 759:
                        v2[(^(v22))] = ("" + "a")
                    case 128:
                        fmt.Println(v5, v4)
                    case -175:
                        if ((true && (v1 == v1)) && (((v5 & int32(v1)) < v5) && ((v29 | v29) < 147))) {
                            v1 -= (-(v1))
                            v4 = (^(v4))
                        } else {
                            if (0 <= 42) {
                                v30 := GenS2{F0: ((14 >= (v5 - (v5 | (v5 & (v5 % (((v5 ^ (v5 + v5))) | 1)))))) && ((v5 << ((v28) & 63)) > v5))}
                                _ = v30
                            } else {
                                _ = len(v13)
                                var v31 float32 = 0.0
                                _ = v31
                            }
                            var v32 float64 = 0.1
                            for fi33 := 0; fi33 < 16; fi33++ {
                                var v34 float64 = v32
                                v35 := GenS1{F0: 32767, F1: 127, F2: 1}
                                _ = fi33
                                _ = v34
                                _ = v35
                            }
                            _ = v32
                        }
                    default:
                        for fi36 := 0; fi36 < 13; fi36++ {
                            if ((v29 % ((v29) | 1)) > (v29 >> ((v28) & 63))) {
                                for fi37 := 0; fi37 < 13; fi37++ {
                                    fmt.Println(v5, v1)
                                    var v38 float32 = -2.25
                                    _ = fi37
                                    _ = v38
                                }
                                v39 := GenS1{F0: 0, F1: (v29 ^ 128), F2: 42}
                                v40 := GenS2{F0: true}
                                _ = v39
                                _ = v40
                            }
                            _ = fi36
                        }
                    }
                    _ = v28
                    _ = v29
                } else {
                    v4 = v4
                }
            }
            v2[v22] = (("hello" + "a") + (("foo bar" + "a") + ""))
            v41 := GenS1{F0: 8, F1: 0, F2: -9223372036854775808}
            _ = v41
        }
        if true {
            v42 := GenS1{F0: 2, F1: 139, F2: 42}
            v43 := GenS3{F0: 1, F1: 37}
            for fi44 := 0; fi44 < 7; fi44++ {
                v13 = append(v13, fi44)
                fmt.Println(v1)
                for fi45 := 0; fi45 < 10; fi45++ {
                    v46 := GenS3{F0: 2, F1: v4}
                    switch ("\n\t" + (("foo bar" + "hello") + "\n\t")) {
                    case "foo bar":
                        delete(v3, -252)
                    case "a":
                        var v47 uint16 = 593
                        _ = v47
                    default:
                        v48 := GenS2{F0: (v1 <= v1)}
                        _ = v48
                    }
                    v49 := GenS1{F0: 1, F1: 127, F2: -1}
                    _ = fi45
                    _ = v46
                    _ = v49
                }
                _ = fi44
            }
            _ = v42
            _ = v43
        } else {
            v22 = 42
            for fi50 := 0; fi50 < 13; fi50++ {
                var v51 uint32 = v22
                var v52 int = fi50
                fmt.Println(v5, fi50)
                _ = fi50
                _ = v51
                _ = v52
            }
        }
    }
    switch ("hello" + (("hello" + "\n\t") + "a")) {
    case "hello":
        if true {
            switch (("\n\t" + "") + ("a" + ((("\n\t" + ("hello" + "foo bar")) + (("foo bar" + "") + "a")) + ""))) {
            case "a":
                v53 := GenS4{F0: v22, F1: (v5 >> ((564) & 63))}
                _ = v53
            case "foo bar":
                var v54 uint = 2
                _ = v54
            case "":
                switch -9223372036854775808 {
                case 9223372036854775807:
                    v22 += (v22 * v22)
                default:
                    fmt.Println(v1, v1)
                }
            default:
                var v55 int = -759
                _ = v55
            }
        } else {
            v5 += v5
            _ = len(v3)
            v22 += 2147483647
        }
    case "a":
        v5 = (v5 & (v5 & (v5 | 2147483647)))
    default:
        fmt.Println(v22, v4, v4)
    }
    if (v22 == 0) {
        switch 127 {
        case 128:
            v5 = v5
        case -2147483648:
            fmt.Println(v22, v22)
        case -9223372036854775808:
            switch 0 {
            case -1:
                for fi56 := 0; fi56 < 20; fi56++ {
                    var v57 uint16 = 1
                    _ = fi56
                    _ = v57
                }
            }
        default:
            var v58 int64 = 0
            _ = v58
        }
        var v59 float32 = (0.0 * 0.1)
        _ = v59
    } else {
        v5 -= v5
        var v60 uint8 = 218
        for fi61 := 0; fi61 < 2; fi61++ {
            switch fi61 {
            case -2147483648:
                if len(v13) > 0 {
                    v13[0] = (fi61 - fi61)
                }
            case 0:
                v5 -= (v5 | int32(v4))
            case 2147483647:
                for fi62 := 0; fi62 < 17; fi62++ {
                    v63 := GenS4{F0: (v22 + (v22 ^ (^(v22)))), F1: v5}
                    if false {
                        if true {
                            v60 -= 225
                            var v64 int16 = 0
                            v4 = (-(v4))
                            _ = v64
                        } else {
                            switch (fi61 >> ((2) & 63)) {
                            case 793:
                                v5 = v5
                            case 625:
                                switch "" {
                                case "hello":
                                    fmt.Println(fi62)
                                case "\n\t":
                                    fmt.Println(fi62, v5, v60)
                                case "":
                                    _ = len(v3)
                                default:
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                    v65 := GenS2{F0: ((v4 & v4) <= v4)}
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                    _ = v65
                                }
                            default:
                                v66 := GenS1{F0: 42, F1: v60, F2: 9223372036854775807}
                                _ = v66
                            }
                        }
                    }
                    fmt.Println(v22, fi61)
                    _ = fi62
                    _ = v63
                }
            }
            v4 -= (v4 % (((v4 >> ((42) & 63))) | 1))
            var v67 uint = 580
            _ = fi61
            _ = v67
        }
        _ = v60
    }
    var v68 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v68
}
