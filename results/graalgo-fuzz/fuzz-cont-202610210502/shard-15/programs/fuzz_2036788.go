package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int; F1 uint16; F2 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 string; F2 uint8 }
type GenS6 struct { F0 uint16; F1 int64 }
func genFn0() int8 {
    v5 := []int{42, 127}
    var v6 uint64 = 340
    if len(v5) > 0 {
        v5[0] = 42
    }
    _ = v5
    _ = v6
    return 127
}
func genFn1() uint32 {
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := GenS1{F0: -27, F1: 61, F2: 18446744073709551615}
        var v9 uint32 = 2
        if true {
            v9 = (^(v9))
        } else {
            if (!false) {
                v10 := GenS1{F0: -1, F1: 65535, F2: 144}
                v9 -= (v9 - 2147483647)
                if (-2.25 == (1e10 * -2.25)) {
                    var v11 bool = (-575 == -1)
                    v12 := []int{fi7, (-(fi7))}
                    v12 = append(v12, fi7)
                    _ = v11
                    _ = v12
                }
                _ = v10
            }
            if true {
                v13 := []int{(fi7 | (fi7 ^ fi7))}
                v14 := GenS2{F0: ((0 == 488) || true), F1: true}
                if false {
                    var v15 int16 = -1
                    v16 := GenS2{F0: false, F1: ((v15 / (((v15 >> ((4294967295) & 63))) | 1)) != v15)}
                    var v17 float64 = (0.0 - 3.14159)
                    _ = v15
                    _ = v16
                    _ = v17
                }
                _ = v13
                _ = v14
            }
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    v18 := GenS1{F0: 974, F1: 0, F2: 2}
    for fi19 := 0; fi19 < 16; fi19++ {
        if true {
            var v20 uint32 = 0
            for fi21 := 0; fi21 < 13; fi21++ {
                var v22 int = (fi21 + (fi19 * (fi19 ^ (fi19 >> ((1) & 63)))))
                switch "a" {
                case "foo bar":
                    if (!(!false)) {
                        v23 := GenS2{F0: false, F1: (!(170 < 354))}
                        var v24 float64 = 1e10
                        _ = v23
                        _ = v24
                    } else {
                        fmt.Println(fi21)
                        if true {
                            switch fi21 {
                            case -2147483648:
                                v25 := GenS3{F0: (fi21 % ((fi21) | 1)), F1: 123, F2: 0}
                                _ = v25
                            }
                            var v26 uint32 = v20
                            _ = v26
                        } else {
                            v27 := GenS3{F0: fi21, F1: 546, F2: 49}
                            v28 := GenS3{F0: 9223372036854775807, F1: 0, F2: 100}
                            v29 := GenS3{F0: fi19, F1: 32767, F2: 27}
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                    }
                default:
                    v20 -= 4294967295
                }
                _ = fi21
                _ = v22
            }
            _ = v20
        } else {
            fmt.Println(fi19)
        }
        v30 := []int{-238}
        _ = fi19
        _ = v30
    }
    _ = v18
    return 2147483648
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int64 = 1
    v3 := make(map[string]string)
    var v4 uint = 42
    if ("foo bar" >= "") {
        var v31 uint8 = 196
        fmt.Println(v4, v1, v31)
        fmt.Println(v1)
        _ = v31
    } else {
        var v32 uint16 = uint16(v1)
        switch 991 {
        case 0:
            switch ("foo bar" + "") {
            case "\n\t":
                v3[("a" + "\n\t")] = ("foo bar" + "a")
            case "":
                for fi33 := 0; fi33 < 1; fi33++ {
                    switch ("foo bar" + "foo bar") {
                    case "":
                        v32 = (v32 * 18)
                    }
                    v1 = v1
                    v34 := []int{fi33, (fi33 << ((v4) & 63)), fi33}
                    _ = fi33
                    _ = v34
                }
            default:
                switch -795 {
                case 42:
                    var v35 uint = (^(v4))
                    _ = v35
                case -2147483648:
                    fmt.Println(v32)
                case 573:
                    v32 += v32
                default:
                    v36 := GenS1{F0: 793, F1: (v32 & (^(v32))), F2: v1}
                    _ = v36
                }
            }
        case 796:
            v1 += (^(v1))
        default:
            var v37 string = ""
            _ = v37
        }
        _ = v32
    }
    for fi38 := 0; fi38 < 18; fi38++ {
        for fi39 := 0; fi39 < 5; fi39++ {
            var v40 uint16 = 65535
            if false {
                var v41 uint32 = 42
                v40 = v40
                _ = v41
            } else {
                v40 = v40
            }
            if (2 <= 1) {
                v42 := []int{}
                v43 := GenS3{F0: (fi39 % ((-9223372036854775808) | 1)), F1: (v40 ^ 898), F2: 100}
                _ = v42
                _ = v43
            }
            var v44 int8 = -1
            _ = fi39
            _ = v40
            _ = v44
        }
        if true {
            if (fi38 >= 2147483647) {
                var v45 uint16 = 42
                _ = v45
            }
            fmt.Println(v2, v1)
        }
        if (true || ((v2 | v2) == -431)) {
            switch (fi38 | -108) {
            case 1:
                for fi46 := 0; fi46 < 19; fi46++ {
                    v47 := GenS2{F0: false, F1: (348 <= int32(v1))}
                    switch "hello" {
                    case "foo bar":
                        var v48 float64 = 0.1
                        _ = v48
                    default:
                        switch "hello" {
                        case "foo bar":
                            var v49 uint = 184
                            _ = v49
                        case "a":
                            fmt.Println(fi38, fi46, fi46)
                        case "":
                            v4 = v4
                        }
                    }
                    _ = fi46
                    _ = v47
                }
            case 0:
                var v50 float32 = 0.0
                _ = v50
            case -9223372036854775808:
                v51 := GenS3{F0: fi38, F1: 2, F2: 0}
                _ = v51
            default:
                var v52 string = "\n\t"
                _ = v52
            }
            v53 := []int{}
            _ = v53
        }
        _ = fi38
    }
    if (v1 < (v1 << ((v4) & 63))) {
        v1 -= (v1 >> ((v4) & 63))
        v4 = v4
    }
    for fi54 := 0; fi54 < 19; fi54++ {
        v1 -= v1
        v4 = v4
        v1 = (^(v1))
        _ = fi54
    }
    var v55 uint32 = 1
    v56 := GenS4{F0: (-40 == -128)}
    for fi57 := 0; fi57 < 15; fi57++ {
        for fi58 := 0; fi58 < 13; fi58++ {
            var v59 float32 = 0.1
            if (!false) {
                v60 := v3[("\n\t" + "hello")]
                v61 := GenS2{F0: ((v4 - v4) != (v4 ^ (^(v4)))), F1: true}
                _ = v60
                _ = v61
            } else {
                switch 32767 {
                case 2147483647:
                    v2 += 11
                case -1:
                    var v62 uint64 = v1
                    _ = v62
                }
                var v63 uint16 = 736
                switch "foo bar" {
                case "hello":
                    v64 := GenS5{F0: 673, F1: ("" + (((("\n\t" + (("" + "\n\t") + "a")) + "a") + ("\n\t" + ("\n\t" + "\n\t"))) + ("hello" + "a"))), F2: 127}
                    _ = v64
                case "foo bar":
                    var v65 int64 = (v2 >> ((v4) & 63))
                    _ = v65
                default:
                    if ((true && (-2147483648 == fi57)) || true) {
                        v55 += (v55 % (((v55 * 2147483648)) | 1))
                        if false {
                            v66 := []int{fi58, (-(fi57)), fi58, int(v4)}
                            switch (fi57 * (fi58 << (((v4 + (v4 % ((uint(v63)) | 1)))) & 63))) {
                            case 1:
                                delete(v3, "hello")
                            default:
                                var v67 float32 = 1e10
                                _ = v67
                            }
                            v59 = (3.14159 / v59)
                            _ = v66
                        }
                    }
                }
                _ = v63
            }
            _ = fi58
            _ = v59
        }
        var v68 float64 = -2.25
        for fi69 := 0; fi69 < 16; fi69++ {
            v1 -= v1
            var v70 int8 = -44
            v71 := GenS2{F0: true, F1: (!(false && true))}
            _ = fi69
            _ = v70
            _ = v71
        }
        _ = fi57
        _ = v68
    }
    for fi72 := 0; fi72 < 16; fi72++ {
        if true {
            var v73 float32 = 1e10
            v1 = 18446744073709551615
            var v74 int16 = 0
            _ = v73
            _ = v74
        } else {
            v75 := []int{fi72, fi72, fi72, (fi72 & fi72)}
            v2 = (-(v2))
            _ = v75
        }
        if (!((0.1 < (3.14159 / 3.14159)) || false)) {
            v2 += v2
            var v76 uint8 = 128
            fmt.Println(v2)
            _ = v76
        }
        _ = fi72
    }
    v77 := GenS5{F0: 729, F1: "", F2: 128}
    var v78 uint32 = 2147483648
    v79 := GenS6{F0: 2, F1: v2}
    var v80 uint = v4
    v81 := v3["hello"]
    v82 := []int{2147483647, -9223372036854775808, 0, -365}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v55:", v55)
    fmt.Println("v78:", v78)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    fmt.Println("slicelen v82:", len(v82), v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v56
    _ = v77
    _ = v78
    _ = v79
    _ = v80
    _ = v81
    _ = v82
}
