package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int; F1 int32; F2 uint32 }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint32 }
func genFn0(p0 uint64, p1 uint16) uint64 {
    for fi7 := 0; fi7 < 18; fi7++ {
        p1 -= p1
        _ = fi7
    }
    v8 := GenS1{F0: (p1 - (p1 + 1)), F1: p0, F2: 2}
    var v9 uint16 = p1
    v10 := GenS2{F0: int(v9), F1: 0, F2: 2}
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return (p0 + (^(p0)))
}
func genFn1() int16 {
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 uint8 = 2
        _ = fi11
        _ = v12
    }
    if (-1 != 127) {
        switch ("hello" + ("foo bar" + "foo bar")) {
        case "a":
            switch 4294967295 {
            case 32768:
                switch ("foo bar" + "hello") {
                case "":
                    v13 := GenS2{F0: 4294967295, F1: 542, F2: 277}
                    _ = v13
                case "\n\t":
                    var v14 string = "a"
                    _ = v14
                case "a":
                    var v15 float32 = ((1e10 + float32(1e10)) * (1e10 * ((3.14159 / -2.25) - (1.5 + 1e10))))
                    _ = v15
                default:
                    var v16 uint64 = 0
                    _ = v16
                }
            case 860:
                switch 675 {
                case -2147483648:
                    var v17 uint8 = 128
                    _ = v17
                case -772:
                    if false {
                        var v18 uint16 = 32767
                        var v19 uint32 = 2147483648
                        _ = v18
                        _ = v19
                    } else {
                        if false {
                            v20 := []int{42}
                            var v21 string = ("foo bar" + ("a" + "\n\t"))
                            v21 = (v21 + ("" + ""))
                            _ = v20
                            _ = v21
                        } else {
                            switch 128 {
                            case -1:
                                v22 := GenS2{F0: 128, F1: -582, F2: 42}
                                _ = v22
                            case -90:
                                switch "foo bar" {
                                case "\n\t":
                                    v23 := GenS1{F0: 2, F1: 18446744073709551615, F2: 128}
                                    _ = v23
                                case "":
                                    var v24 uint64 = 379
                                    _ = v24
                                default:
                                    v25 := []int{32767}
                                    _ = v25
                                }
                            case -36:
                                var v26 int16 = -32768
                                _ = v26
                            }
                        }
                        if (!true) {
                            var v27 float64 = (((0.1 * (1.5 + (-2.25 + (0.0 / float64(-2.25))))) + 0.0) - (3.14159 - (1.5 - 0.1)))
                            _ = v27
                        } else {
                            for fi28 := 0; fi28 < 2; fi28++ {
                                fmt.Println(fi28)
                                var v29 int64 = -9223372036854775808
                                var v30 int64 = v29
                                _ = fi28
                                _ = v29
                                _ = v30
                            }
                        }
                    }
                default:
                    if (!(!false)) {
                        switch (("a" + "\n\t") + "") {
                        case "":
                            switch "" {
                            case "hello":
                                var v31 int32 = 42
                                _ = v31
                            default:
                                if (false || true) {
                                    if (42 == 127) {
                                        var v32 int8 = 1
                                        var v33 int32 = 1
                                        _ = v32
                                        _ = v33
                                    } else {
                                        switch 498 {
                                        case 4294967295:
                                            var v34 float32 = 3.14159
                                            _ = v34
                                        case 1:
                                            var v35 string = ""
                                            _ = v35
                                        case -9223372036854775808:
                                            var v36 uint64 = 119
                                            _ = v36
                                        }
                                    }
                                } else {
                                    for fi37 := 0; fi37 < 13; fi37++ {
                                        var v38 uint32 = 0
                                        _ = fi37
                                        _ = v38
                                    }
                                }
                            }
                        }
                    }
                }
            case 751:
                v39 := []int{2147483647, 128, -834}
                _ = v39
            default:
                v40 := []int{-556, 128}
                _ = v40
            }
        case "\n\t":
            var v41 uint32 = 4294967295
            _ = v41
        default:
            var v42 int64 = 42
            _ = v42
        }
    }
    var v43 uint8 = 127
    fmt.Println(v43)
    _ = v43
    return -732
}

func main() {
    v1 := []string{"\n\t", "hello"}
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    v4 := []string{"", "", "hello", ""}
    var v5 int32 = 632
    var v6 int32 = (v5 % (((^(v5))) | 1))
    v44 := GenS1{F0: 660, F1: 4294967295, F2: 2}
    v5 += (-(v6))
    fmt.Println(v6, v6)
    v6 = (v6 + -2147483648)
    switch ((("foo bar" + ("a" + "foo bar")) + ("foo bar" + "a")) + "a") {
    case "a":
        if (0.1 <= 1.5) {
            for fi45 := 0; fi45 < 3; fi45++ {
                var v46 uint64 = 42
                for fi47 := 0; fi47 < 2; fi47++ {
                    if (!true) {
                        v3[(fi47 % ((fi45) | 1))] = (fi47 << ((1) & 63))
                    }
                    v5 -= (v5 - v5)
                    _ = fi47
                }
                _ = fi45
                _ = v46
            }
            v5 = (^(v5))
        }
    }
    v5 -= (-(v6))
    var v48 int = 695
    if false {
        v4 = append(v4, "foo bar")
    }
    switch "hello" {
    case "":
        var v49 int = int(v5)
        _ = v49
    default:
        if (!true) {
            v5 = v6
            if false {
                v50 := GenS3{F0: 2147483648, F1: 2, F2: 42}
                v5 = (v5 | v6)
                var v51 uint32 = 346
                _ = v50
                _ = v51
            } else {
                switch (-(v48)) {
                case -1:
                    if (4294967295 <= 970) {
                        if true {
                            v52 := GenS2{F0: (v48 & v48), F1: v6, F2: 771}
                            _ = v52
                        }
                        v48 = v48
                        if (40 <= 255) {
                            var v53 uint32 = 937
                            _ = v53
                        }
                    }
                case 4294967295:
                    var v54 float64 = ((0.1 - (0.1 * 0.1)) / 1.0)
                    _ = v54
                case 9223372036854775807:
                    var v55 uint64 = 18446744073709551615
                    _ = v55
                default:
                    for fi56 := 0; fi56 < 12; fi56++ {
                        for fi57 := 0; fi57 < 6; fi57++ {
                            switch "foo bar" {
                            case "a":
                                if true {
                                    if true {
                                        v5 = v5
                                        v48 = fi56
                                        v58 := v3[(^(fi56))]
                                        _ = v58
                                    }
                                    var v59 bool = (true && (!(0 != 2)))
                                    _ = v59
                                } else {
                                    v6 = (^(v6))
                                }
                            case "hello":
                                v5 -= v6
                            case "\n\t":
                                var v60 uint8 = 255
                                v5 = v5
                                var v61 uint32 = 4294967295
                                _ = v60
                                _ = v61
                            }
                            v6 += (v5 >> ((980) & 63))
                            v62 := GenS1{F0: 110, F1: 523, F2: 2}
                            _ = fi57
                            _ = v62
                        }
                        v5 += (v6 * (v6 ^ (-(v6))))
                        v6 -= (v6 * (v5 * -356))
                        _ = fi56
                    }
                }
                v63 := GenS2{F0: 127, F1: v5, F2: 1}
                _ = v63
            }
            v48 += (-(v48))
        }
    }
    var v64 int64 = 9223372036854775807
    _ = len(v1)
    v6 -= (v6 | v6)
    if (false && false) {
        v3[v48] = (v48 - 32767)
        for fi65 := 0; fi65 < 2; fi65++ {
            if ((1e10 + 0.1) <= -2.25) {
                for fi66 := 0; fi66 < 17; fi66++ {
                    v6 = v5
                    _ = fi66
                }
                delete(v3, fi65)
                if (((!((^(v48)) <= (v48 + (-(v48))))) && (219 >= 158)) || true) {
                    _ = len(v4)
                    var v67 int64 = v64
                    _ = v67
                }
            }
            var v68 uint32 = 65535
            _ = fi65
            _ = v68
        }
    }
    fmt.Println(v48, v48)
    fmt.Println(v64, v6, v64)
    v6 += (v6 & v6)
    switch "" {
    case "foo bar":
        var v69 int8 = -1
        _ = v69
    case "hello":
        v6 -= -841
    case "":
        if (((3.14159 / 0.1) * 3.14159) < 3.14159) {
            if (0 >= 26) {
                switch "\n\t" {
                case "":
                    if (v6 >= v5) {
                        var v70 bool = (!true)
                        _ = v70
                    }
                case "foo bar":
                    var v71 int32 = -2147483648
                    _ = v71
                case "hello":
                    v72 := GenS2{F0: v48, F1: (v6 * v6), F2: 65535}
                    _ = v72
                default:
                    v6 += (v6 / ((923) | 1))
                }
                _ = len(v2)
            } else {
                for fi73 := 0; fi73 < 6; fi73++ {
                    v74 := GenS1{F0: 255, F1: genFn0(4294967295, 32767), F2: 127}
                    v6 += v5
                    _ = fi73
                    _ = v74
                }
                var v75 int8 = 80
                if (!(-128 == 1)) {
                    v5 -= 2147483647
                    switch v48 {
                    case 32767:
                        if false {
                            for fi76 := 0; fi76 < 14; fi76++ {
                                var v77 uint16 = 2
                                var v78 float32 = 0.0
                                v64 -= (v64 * (v64 >> ((65535) & 63)))
                                _ = fi76
                                _ = v77
                                _ = v78
                            }
                            fmt.Println(v64)
                        }
                    }
                }
                _ = v75
            }
            v6 = int32(v48)
            for fi79 := 0; fi79 < 20; fi79++ {
                switch (fi79 / ((fi79) | 1)) {
                case -736:
                    var v80 int8 = -1
                    _ = v80
                }
                fmt.Println(fi79, v64)
                var v81 int32 = v5
                _ = fi79
                _ = v81
            }
        }
    default:
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
    }
    var v82 float32 = ((1e10 * 0.1) - -2.25)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v48
    _ = v5
    _ = v6
    _ = v64
    _ = v82
}
