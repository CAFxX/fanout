package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 bool; F1 int64; F2 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 uint32; F2 string }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 string }
func genFn0() uint16 {
    if false {
        var v7 int16 = 0
        _ = v7
    }
    if false {
        if true {
            v8 := []int{-2147483648, 64, 263, 9223372036854775807}
            _ = v8
        } else {
            v9 := GenS1{F0: 184, F1: 463}
            if true {
                switch ("hello" + "hello") {
                case "hello":
                    v10 := GenS2{F0: (true || (!((!false) || (0 != 119)))), F1: -9223372036854775808, F2: 338}
                    _ = v10
                default:
                    var v11 bool = (1e10 != 1e10)
                    _ = v11
                }
            } else {
                v12 := []int{}
                if (915 >= 2147483647) {
                    if len(v12) > 0 {
                        v12[0] = -593
                    }
                    v13 := GenS2{F0: false, F1: 602, F2: 2}
                    var v14 int32 = 2147483647
                    _ = v13
                    _ = v14
                } else {
                    v15 := GenS1{F0: 63, F1: 65535}
                    if len(v12) > 0 {
                        v12[0] = 128
                    }
                    _ = v15
                }
                _ = v12
            }
            _ = v9
        }
    }
    return 32767
}
func genFn1(p0 uint, p1 bool) uint64 {
    if (!(6 >= 1)) {
        if (958 < p0) {
            v16 := GenS1{F0: (p0 + (p0 & 42)), F1: 255}
            v17 := GenS3{F0: -143}
            _ = v16
            _ = v17
        } else {
            var v18 string = "a"
            for fi19 := 0; fi19 < 14; fi19++ {
                if p1 {
                    var v20 float32 = 0.0
                    fmt.Println(v20, v18, p0)
                    _ = v20
                }
                v21 := GenS4{F0: (p0 << (((p0 >> ((0) & 63))) & 63)), F1: 1, F2: v18}
                _ = fi19
                _ = v21
            }
            if p1 {
                p1 = ((542 == -2147483648) && (2 < 2147483648))
                v22 := GenS1{F0: (p0 << ((p0) & 63)), F1: p0}
                if (!(82 > int8(p0))) {
                    p0 += 324
                    v23 := GenS4{F0: p0, F1: 2147483648, F2: v18}
                    _ = v23
                }
                _ = v22
            }
            _ = v18
        }
    } else {
        if (p1 && false) {
            fmt.Println(p0)
            p1 = (!(p1 || false))
            for fi24 := 0; fi24 < 13; fi24++ {
                v25 := []int{fi24, int(p0)}
                switch "" {
                case "foo bar":
                    v25 = append(v25, (fi24 / ((-9223372036854775808) | 1)))
                case "hello":
                    v26 := GenS5{F0: -9223372036854775808, F1: 295}
                    _ = v26
                }
                if len(v25) > 0 {
                    v25[0] = (fi24 >> ((p0) & 63))
                }
                _ = fi24
                _ = v25
            }
        } else {
            p0 = p0
            if (false || p1) {
                p0 = 65535
                if (-1 > -905) {
                    v27 := []int{-2147483648, 838, 127}
                    _ = v27
                } else {
                    if p1 {
                        v28 := []int{}
                        _ = v28
                    }
                }
            } else {
                p0 -= (^(p0))
                v29 := []int{}
                switch -205 {
                case 980:
                    for fi30 := 0; fi30 < 19; fi30++ {
                        p1 = p1
                        if (42 < 1000) {
                            v29 = append(v29, (fi30 + (fi30 / (((fi30 << ((4294967295) & 63))) | 1))))
                            p0 = 65535
                            p0 = (^(p0))
                        }
                        switch "a" {
                        case "hello":
                            var v31 uint32 = 2147483648
                            _ = v31
                        }
                        _ = fi30
                    }
                case -722:
                    if true {
                        var v32 int32 = 2147483647
                        _ = v32
                    }
                }
                _ = v29
            }
        }
        v33 := GenS3{F0: 1000}
        _ = v33
    }
    var v34 int = 127
    _ = p0
    _ = p1
    _ = v34
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = -32768
    var v3 int = 741
    var v4 float64 = (-2.25 * 0.0)
    v5 := make(map[string]uint32)
    var v6 uint16 = 28
    for fi35 := 0; fi35 < 13; fi35++ {
        var v36 int64 = 1
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v6, fi35)
        case "hello":
            var v37 int16 = (v2 << ((18446744073709551615) & 63))
            _ = v37
        case "a":
            v38 := GenS4{F0: 1, F1: 4294967295, F2: (("" + "foo bar") + (((("foo bar" + ("foo bar" + (("\n\t" + "foo bar") + ""))) + "a") + ("" + "a")) + ("hello" + "hello")))}
            _ = v38
        }
        for fi39 := 0; fi39 < 19; fi39++ {
            var v40 uint64 = 1
            var v41 int = (fi39 % ((0) | 1))
            if false {
                switch "foo bar" {
                case "hello":
                    v42 := []int{fi35, int(v36), fi39}
                    _ = v42
                default:
                    fmt.Println(v6, v41, fi35)
                    var v43 uint8 = 255
                    _ = v43
                }
                v36 += (v36 ^ int64(v40))
            } else {
                var v44 int8 = 100
                v45 := []int{fi35, (-(fi39)), v41, 4294967295}
                _ = len(v5)
                _ = v44
                _ = v45
            }
            v46 := v5["a"]
            v47 := []int{4294967295, 941, (fi39 + (v3 << ((1) & 63)))}
            v48 := GenS5{F0: v36, F1: fi35}
            var v49 uint8 = 94
            _ = fi39
            _ = v40
            _ = v41
            _ = v46
            _ = v47
            _ = v48
            _ = v49
        }
        _ = fi35
        _ = v36
    }
    if true {
        switch 436 {
        case -2147483648:
            if (4294967295 != 265) {
                var v50 int32 = 42
                v50 = (v50 / (((^(v50))) | 1))
                _ = v50
            } else {
                switch (("a" + (("a" + ("" + "a")) + "a")) + ((("" + "hello") + ("" + "")) + ("" + ("hello" + ("" + ("\n\t" + "foo bar")))))) {
                case "a":
                    for fi51 := 0; fi51 < 18; fi51++ {
                        var v52 int32 = -1
                        if true {
                            if true {
                                delete(v5, ("" + "\n\t"))
                                for fi53 := 0; fi53 < 7; fi53++ {
                                    v54 := v1[("hello" + ((((("\n\t" + "") + "a") + ("foo bar" + "hello")) + "\n\t") + ("foo bar" + "")))]
                                    var v55 bool = true
                                    v56 := []int{int(v52), (^(fi51)), (fi51 / ((fi51) | 1)), 1}
                                    _ = fi53
                                    _ = v54
                                    _ = v55
                                    _ = v56
                                }
                            }
                            for fi57 := 0; fi57 < 19; fi57++ {
                                v6 -= v6
                                fmt.Println(v4, fi57, fi51)
                                v2 -= v2
                                _ = fi57
                            }
                        }
                        _ = fi51
                        _ = v52
                    }
                case "":
                    v3 = -9223372036854775808
                case "foo bar":
                    var v58 float32 = 3.14159
                    _ = v58
                default:
                    if (((0.0 + (v4 * 0.0)) - 0.1) < float64(0.1)) {
                        v6 = 172
                    }
                }
                v6 -= (v6 >> ((2) & 63))
                fmt.Println(v6, v2, v2)
            }
        }
    } else {
        v2 += v2
        for fi59 := 0; fi59 < 19; fi59++ {
            for fi60 := 0; fi60 < 20; fi60++ {
                switch 389 {
                case -900:
                    v61 := GenS2{F0: false, F1: 1, F2: 543}
                    _ = v61
                default:
                    v4 += v4
                }
                delete(v1, ("" + ("\n\t" + ("foo bar" + "a"))))
                v3 -= -571
                _ = fi60
            }
            delete(v5, (("" + "") + (("a" + ((("a" + ("a" + "")) + "foo bar") + ("" + "a"))) + "hello")))
            switch -755 {
            case 145:
                v2 -= (v2 >> ((255) & 63))
            case 1:
                delete(v1, ("" + "hello"))
            default:
                var v62 int = v3
                _ = v62
            }
            _ = fi59
        }
    }
    switch ("" + "") {
    case "a":
        v3 += v3
    case "":
        fmt.Println(v4)
    }
    if true {
        var v63 uint8 = 50
        _ = v63
    }
    v1["hello"] = uint32(v6)
    var v64 uint16 = v6
    v65 := GenS4{F0: uint(v6), F1: 884, F2: "a"}
    v64 += v64
    v66 := GenS6{F0: "hello"}
    var v67 uint16 = v6
    for fi68 := 0; fi68 < 11; fi68++ {
        v69 := v1[("\n\t" + "foo bar")]
        if (1 < 2) {
            if false {
                var v70 bool = true
                v67 += (v6 % (((v64 >> ((18446744073709551615) & 63))) | 1))
                _ = v70
            } else {
                v4 = (v4 + v4)
                v71 := []int{(v3 | fi68), -548, (v3 * (^(fi68)))}
                _ = v71
            }
        }
        delete(v5, "foo bar")
        _ = fi68
        _ = v69
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v65
    _ = v66
    _ = v67
}
