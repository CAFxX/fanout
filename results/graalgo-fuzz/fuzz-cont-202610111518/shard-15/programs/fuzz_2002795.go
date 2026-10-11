package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 int8; F2 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint; F1 uint64; F2 int32 }
type GenS8 struct { F0 int16; F1 int16 }
type GenS9 struct { F0 bool }
func genFn0(p0 uint, p1 int64, p2 uint8) uint16 {
    v6 := []int{127}
    if true {
        v7 := GenS1{F0: p0, F1: 27}
        var v8 uint32 = 42
        _ = v7
        _ = v8
    } else {
        if ((false && (255 != uint8(p0))) && false) {
            v9 := GenS1{F0: p0, F1: (p2 - p2)}
            _ = v9
        } else {
            fmt.Println(p2, p0, p0)
            p2 = 127
            if (true && true) {
                var v10 int16 = 1000
                if len(v6) > 0 {
                    v6[0] = 32768
                }
                switch 42 {
                case -9223372036854775808:
                    for fi11 := 0; fi11 < 7; fi11++ {
                        for fi12 := 0; fi12 < 20; fi12++ {
                            switch "foo bar" {
                            case "a":
                                switch "\n\t" {
                                case "hello":
                                    if ((p2 == p2) && (((p2 & 127) > 127) || true)) {
                                        if false {
                                            fmt.Println(p1, fi12, p0)
                                            v13 := GenS1{F0: p0, F1: (^(p2))}
                                            _ = v13
                                        }
                                    } else {
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        var v14 int16 = (^(v10))
                                        _ = v14
                                    }
                                case "\n\t":
                                    var v15 uint = (p0 | (p0 & 0))
                                    _ = v15
                                default:
                                    if (176 <= 2147483648) {
                                        var v16 int64 = p1
                                        v17 := GenS1{F0: (p0 >> (((p0 >> ((18446744073709551615) & 63))) & 63)), F1: p2}
                                        _ = v16
                                        _ = v17
                                    } else {
                                        p0 = (p0 & p0)
                                        v6 = append(v6, (fi11 >> ((933) & 63)))
                                        if (!(3.14159 != 1.5)) {
                                            var v18 int16 = -32768
                                            if (false && ("" < "\n\t")) {
                                                v18 = (-(v18))
                                                if len(v6) > 0 {
                                                    v6[0] = fi11
                                                }
                                                if (false && ("\n\t" != "foo bar")) {
                                                    var v19 uint = (^(p0))
                                                    if (!((1e10 * (0.1 * 1e10)) > 0.0)) {
                                                        var v20 uint8 = (p2 << ((v19) & 63))
                                                        var v21 int32 = 42
                                                        _ = v20
                                                        _ = v21
                                                    }
                                                    if false {
                                                        v19 = 65535
                                                        var v22 int8 = 19
                                                        _ = v22
                                                    }
                                                    _ = v19
                                                }
                                            }
                                            fmt.Println(p0)
                                            v6 = append(v6, 32767)
                                            v23 := GenS1{F0: p0, F1: (p2 | (p2 + p2))}
                                            _ = v18
                                            _ = v23
                                        } else {
                                            var v24 int32 = 366
                                            var v25 bool = (0 < 0)
                                            if false {
                                                v26 := GenS1{F0: (p0 * p0), F1: 2}
                                                var v27 int32 = (v24 << ((1) & 63))
                                                v28 := GenS2{F0: 2}
                                                _ = v26
                                                _ = v27
                                                _ = v28
                                            } else {
                                                p0 = (p0 & p0)
                                            }
                                            p1 = (^(p1))
                                            _ = v24
                                            _ = v25
                                        }
                                    }
                                }
                            case "\n\t":
                                switch fi11 {
                                case 336:
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                            case "foo bar":
                                var v29 int16 = -699
                                _ = v29
                            }
                            var v30 int32 = 0
                            v31 := GenS3{F0: v30, F1: 127, F2: -41}
                            _ = fi12
                            _ = v30
                            _ = v31
                        }
                        _ = fi11
                    }
                case 2147483647:
                    p2 = p2
                default:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                _ = v10
            }
        }
    }
    if true {
        p1 = int64(p0)
        p1 -= (p1 % ((0) | 1))
    }
    v6 = append(v6, -93)
    if false {
        v6 = append(v6, 9223372036854775807)
        if false {
            var v32 uint64 = 1
            var v33 int = 4294967295
            p1 += -699
            _ = v32
            _ = v33
        } else {
            p0 -= 42
            p0 = (p0 ^ (p0 % (((p0 & (p0 + p0))) | 1)))
        }
        var v34 int16 = 1000
        _ = v34
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 364
}

func main() {
    v1 := []uint32{42}
    var v2 string = "a"
    var v3 bool = true
    var v4 uint = 0
    v5 := make(map[string]string)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v35 := GenS4{F0: 128}
    if v3 {
        v3 = (v3 || ((((!v3) || false) && (1 <= 2147483647)) || (((v2 + v2) + "") > (v2 + v2))))
    } else {
        delete(v5, ((v2 + v2) + "a"))
        v36 := GenS5{F0: 223, F1: 65535, F2: 1}
        v4 += 637
        _ = v36
    }
    v37 := GenS6{F0: 9223372036854775807}
    for fi38 := 0; fi38 < 9; fi38++ {
        if (!v3) {
            v2 = "hello"
        }
        var v39 bool = ((0.0 - 1.5) != float64(((-2.25 * 1e10) * (0.1 - ((1.5 - (3.14159 / 3.14159)) * 1e10)))))
        v3 = v3
        _ = fi38
        _ = v39
    }
    if false {
        v4 += v4
    }
    v3 = ((v4 << ((v4) & 63)) == (v4 & (v4 >> ((18446744073709551615) & 63))))
    fmt.Println(v3, v3, v2)
    var v40 string = v2
    if (((1e10 - (3.14159 + (3.14159 / float32(-2.25)))) == float32((0.0 + (0.0 - ((0.0 + 3.14159) + (float64(0.0) + 1e10)))))) || v3) {
        if (v3 && true) {
            var v41 int16 = 32767
            for fi42 := 0; fi42 < 17; fi42++ {
                if false {
                    switch ("" + v40) {
                    case "":
                        _ = len(v1)
                    case "a":
                        v43 := v5[(("" + ("hello" + "a")) + "hello")]
                        _ = v43
                    default:
                        v44 := v5[(v40 + v40)]
                        _ = v44
                    }
                    switch ("hello" + "") {
                    case "hello":
                        v3 = false
                    default:
                        var v45 uint64 = 18446744073709551615
                        _ = v45
                    }
                } else {
                    for fi46 := 0; fi46 < 18; fi46++ {
                        v3 = (!true)
                        if ((608 > 255) && v3) {
                            if len(v1) > 0 {
                                v1[0] = 4294967295
                            }
                            v47 := GenS5{F0: 53, F1: 64, F2: 0}
                            v48 := GenS7{F0: (v4 & 4294967295), F1: 1, F2: 0}
                            _ = v47
                            _ = v48
                        }
                        var v49 uint32 = 4294967295
                        v5[(v40 + (v40 + (v40 + v2)))] = "\n\t"
                        var v50 uint32 = 1
                        _ = fi46
                        _ = v49
                        _ = v50
                    }
                }
                _ = fi42
            }
            _ = v41
        }
        _ = len(v5)
    } else {
        if v3 {
            if v3 {
                if ((0 < -1) && (((!(9223372036854775807 != 677)) && (1000 > -941)) || v3)) {
                    fmt.Println(v3)
                } else {
                    v4 = v4
                }
            }
            switch 42 {
            case 953:
                v2 = v40
            }
        } else {
            if (true && v3) {
                for fi51 := 0; fi51 < 18; fi51++ {
                    var v52 uint = (^(v4))
                    delete(v5, (v40 + (v2 + ((((v2 + v2) + (v40 + "\n\t")) + ("a" + ("hello" + "\n\t"))) + v2))))
                    switch (^(fi51)) {
                    case -126:
                        if (((0.1 == 0.0) && (("hello" + v2) > v2)) && (-1 <= 32767)) {
                            var v53 int = (fi51 ^ fi51)
                            fmt.Println(v52, v2)
                            var v54 int32 = 327
                            _ = v53
                            _ = v54
                        }
                    case 352:
                        var v55 int16 = 1000
                        _ = v55
                    }
                    _ = fi51
                    _ = v52
                }
                var v56 uint8 = 44
                fmt.Println(v4)
                _ = v56
            } else {
                for fi57 := 0; fi57 < 5; fi57++ {
                    if true {
                        v58 := GenS8{F0: 1000, F1: -930}
                        var v59 string = v40
                        _ = v58
                        _ = v59
                    }
                    var v60 uint64 = 4294967295
                    _ = fi57
                    _ = v60
                }
            }
            v4 = v4
            if ((^(v4)) <= (^(v4))) {
                var v61 uint8 = 224
                _ = v61
            } else {
                var v62 int64 = 0
                v63 := GenS9{F0: v3}
                var v64 string = ("hello" + "hello")
                _ = v62
                _ = v63
                _ = v64
            }
        }
        v65 := GenS7{F0: v4, F1: 824, F2: -822}
        for fi66 := 0; fi66 < 20; fi66++ {
            switch v40 {
            case "foo bar":
                var v67 uint16 = 379
                _ = v67
            }
            for fi68 := 0; fi68 < 15; fi68++ {
                v2 = v2
                _ = fi68
            }
            v2 = v2
            _ = fi66
        }
        _ = v65
    }
    v2 = "foo bar"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v4
    _ = v40
    _ = v5
}
