package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int16 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint8; F1 uint64; F2 int32 }
type GenS4 struct { F0 int8; F1 int64 }
func genFn0(p0 uint, p1 int16) bool {
    v7 := GenS1{F0: "", F1: (p0 % (((^(p0))) | 1)), F2: 42}
    if true {
        var v8 int32 = -1
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v7
    return (!false)
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 42
    var v3 int32 = -1
    var v4 uint64 = 42
    var v5 uint8 = 113
    v6 := make(map[int]int)
    var v9 bool = (255 < 65535)
    var v10 int16 = 1000
    var v11 bool = v9
    if (!true) {
        var v12 int8 = -128
        v3 = (v2 >> ((255) & 63))
        _ = v12
    }
    if true {
        switch "foo bar" {
        case "":
            switch "a" {
            case "foo bar":
                for fi13 := 0; fi13 < 3; fi13++ {
                    if ((v9 || (true || v11)) && v9) {
                        switch "\n\t" {
                        case "foo bar":
                            if v11 {
                                v14 := GenS2{F0: 670, F1: -1}
                                _ = v14
                            }
                        case "":
                            v15 := []int{(fi13 << ((122) & 63))}
                            _ = v15
                        default:
                            fmt.Println(v5, fi13, v10)
                        }
                        var v16 uint = 697
                        var v17 uint16 = 886
                        _ = v16
                        _ = v17
                    } else {
                        var v18 uint64 = 720
                        _ = v18
                    }
                    _ = fi13
                }
            case "hello":
                fmt.Println(v3, v5)
            case "a":
                var v19 int8 = -1
                _ = v19
            default:
                if (v5 < v5) {
                    fmt.Println(v2)
                    var v20 uint32 = 221
                    var v21 float64 = (1e10 - -2.25)
                    _ = v20
                    _ = v21
                }
            }
        default:
            v9 = v11
        }
        for fi22 := 0; fi22 < 10; fi22++ {
            if true {
                var v23 int16 = v10
                _ = v23
            } else {
                var v24 uint = 255
                _ = v24
            }
            v25 := GenS2{F0: -9223372036854775808, F1: 47}
            _ = fi22
            _ = v25
        }
        for fi26 := 0; fi26 < 4; fi26++ {
            v3 = v2
            _ = fi26
        }
    }
    var v27 uint64 = v4
    fmt.Println(v5)
    v28 := []int{-672, 144}
    v29 := GenS2{F0: 940, F1: 1}
    fmt.Println(v3, v3, v27)
    if ((true || v11) && (v11 || (2147483647 != 0))) {
        for fi30 := 0; fi30 < 4; fi30++ {
            fmt.Println(v2, v4, v2)
            _ = fi30
        }
        v4 -= 0
    } else {
        v5 = v5
        fmt.Println(v27, v3)
        switch 42 {
        case 1:
            fmt.Println(v5)
        case 295:
            if ("hello" >= "\n\t") {
                switch ("hello" + "a") {
                case "":
                    if len(v28) > 0 {
                        v28[0] = -9223372036854775808
                    }
                case "\n\t":
                    var v31 uint8 = 128
                    _ = v31
                default:
                    switch ("" + "hello") {
                    case "\n\t":
                        v32 := GenS1{F0: (("\n\t" + "\n\t") + "\n\t"), F1: 207, F2: 440}
                        _ = v32
                    default:
                        if len(v28) >= 2 {
                            v28 = v28[1:len(v28)]
                        }
                    }
                }
                var v33 int8 = 1
                _ = v33
            } else {
                v10 = -1
            }
        default:
            var v34 string = "hello"
            _ = v34
        }
    }
    if (42 >= 9223372036854775807) {
        for fi35 := 0; fi35 < 3; fi35++ {
            for fi36 := 0; fi36 < 19; fi36++ {
                if (!v9) {
                    v37 := GenS1{F0: ("foo bar" + ""), F1: 65535, F2: (^(v10))}
                    var v38 bool = v11
                    var v39 float64 = 3.14159
                    _ = v37
                    _ = v38
                    _ = v39
                }
                var v40 uint8 = (v5 ^ v5)
                var v41 int64 = 42
                if v11 {
                    if (true && ((uint(v40) != 255) && v11)) {
                        if (v11 || v11) {
                            v42 := v1[fi36]
                            if (v5 < v5) {
                                var v43 int8 = 104
                                if (!v11) {
                                    var v44 bool = true
                                    _ = v44
                                } else {
                                    if v11 {
                                        v28 = append(v28, -2147483648)
                                        if v9 {
                                            var v45 float64 = ((1.5 + 0.1) - 3.14159)
                                            _ = v45
                                        } else {
                                            v46 := GenS1{F0: "", F1: 2, F2: 42}
                                            v2 = (v2 << ((953) & 63))
                                            if (2 < 396) {
                                                if (!((2 < (v40 | v5)) || true)) {
                                                    if (!(!(!(!((v9 || true) || v9))))) {
                                                        v42 -= fi35
                                                        fmt.Println(v43)
                                                    } else {
                                                        if ((v5 - 2) == (v40 >> ((1) & 63))) {
                                                            var v47 int16 = (v10 >> ((255) & 63))
                                                            _ = len(v28)
                                                            _ = v47
                                                        }
                                                        v48 := GenS2{F0: (v41 / (((v41 >> ((364) & 63))) | 1)), F1: v43}
                                                        var v49 uint = 342
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                } else {
                                                    var v50 int8 = (-(v43))
                                                    var v51 int = fi35
                                                    fmt.Println(v50, v4)
                                                    if len(v28) >= 2 {
                                                        v28 = v28[1:len(v28)]
                                                    }
                                                    var v52 int = (v42 << ((uint(v5)) & 63))
                                                    _ = v50
                                                    _ = v51
                                                    _ = v52
                                                }
                                            }
                                            _ = v46
                                        }
                                        if true {
                                            var v53 int32 = (v3 >> ((0) & 63))
                                            if true {
                                                v5 += (v40 >> ((255) & 63))
                                                var v54 float32 = 1e10
                                                switch (v42 & (fi35 % ((fi35) | 1))) {
                                                case -889:
                                                    if false {
                                                        v40 = (v40 << ((1) & 63))
                                                    }
                                                }
                                                _ = v54
                                            }
                                            var v55 uint64 = v27
                                            _ = v53
                                            _ = v55
                                        }
                                    }
                                    fmt.Println(v42, v3)
                                    var v56 int64 = v41
                                    _ = v56
                                }
                                _ = v43
                            }
                            v57 := GenS3{F0: 81, F1: v27, F2: (^(v3))}
                            _ = v42
                            _ = v57
                        }
                    }
                    v40 = (v5 | (^(v40)))
                } else {
                    v41 -= (v41 % ((9223372036854775807) | 1))
                    if v11 {
                        switch (fi35 * (fi35 / ((2147483647) | 1))) {
                        case -9223372036854775808:
                            var v58 float32 = 3.14159
                            _ = v58
                        case 680:
                            v41 = v41
                        case 4294967295:
                            v9 = true
                        default:
                            v28 = append(v28, fi36)
                        }
                    }
                }
                v10 = 136
                _ = fi36
                _ = v40
                _ = v41
            }
            _ = fi35
        }
    } else {
        for fi59 := 0; fi59 < 19; fi59++ {
            delete(v1, 9223372036854775807)
            _ = fi59
        }
        delete(v1, 9223372036854775807)
    }
    fmt.Println(v3, v4)
    v4 += 2
    if ((v4 - v27) != v27) {
        v10 = -1
        fmt.Println(v9, v5, v2)
        for fi60 := 0; fi60 < 8; fi60++ {
            var v61 int = (fi60 << ((255) & 63))
            _ = fi60
            _ = v61
        }
    } else {
        var v62 bool = (true || v9)
        if false {
            v62 = (false && true)
            fmt.Println(v27, v5)
        }
        _ = v62
    }
    v63 := GenS4{F0: 38, F1: -9223372036854775808}
    if (-128 > 34) {
        v64 := GenS3{F0: 128, F1: 321, F2: v2}
        fmt.Println(v3, v3)
        _ = v64
    } else {
        var v65 float64 = 0.1
        _ = v65
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v9
}
