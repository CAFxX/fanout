package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 int16; F1 uint64; F2 int }
func genFn0() uint {
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: 128, F1: "foo bar", F2: 9223372036854775807}
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 int32 = 2147483647
            if true {
                fmt.Println(fi8, fi8)
                if false {
                    if true {
                        v9 += 16
                        v10 := []int{fi6}
                        fmt.Println(v9)
                        _ = v10
                    } else {
                        v11 := GenS1{F0: 42, F1: ("foo bar" + ("\n\t" + "foo bar")), F2: fi8}
                        fmt.Println(fi6, fi8)
                        _ = v11
                    }
                } else {
                    var v12 int = 2147483647
                    switch "a" {
                    case "hello":
                        fmt.Println(v12, fi8)
                    case "\n\t":
                        var v13 uint = 65535
                        _ = v13
                    }
                    var v14 int = fi8
                    _ = v12
                    _ = v14
                }
            }
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 int16 = 42
        _ = fi15
        _ = v16
    }
    return 414
}

func main() {
    var v1 uint32 = 765
    var v2 float64 = 3.14159
    v3 := make(map[uint32]string)
    v4 := make(map[string]uint32)
    var v5 uint16 = 250
    fmt.Println(v2)
    v1 = v1
    v17 := GenS1{F0: 0, F1: "a", F2: 677}
    v18 := []int{733}
    if false {
        var v19 string = "foo bar"
        v20 := GenS1{F0: 0, F1: "a", F2: 42}
        var v21 int64 = -9223372036854775808
        _ = v19
        _ = v20
        _ = v21
    }
    v22 := GenS2{F0: 0}
    var v23 int16 = -960
    fmt.Println(v2, v5)
    v5 -= v5
    if (!(724 < 1000)) {
        v2 += -2.25
        v24 := v4["hello"]
        _ = v24
    } else {
        if (false && false) {
            if false {
                for fi25 := 0; fi25 < 10; fi25++ {
                    if (!((!(false || (1 == 0))) && (((v2 + v2) * float64(((((1.5 * 1.5) + -2.25) / (3.14159 - 0.0)) / (3.14159 + 0.1)))) < float64((0.0 + 1e10))))) {
                        var v26 string = "a"
                        _ = v26
                    }
                    switch 872 {
                    case 2147483647:
                        if ((!(0 >= -9223372036854775808)) || (65535 >= 65535)) {
                            for fi27 := 0; fi27 < 9; fi27++ {
                                delete(v4, "a")
                                _ = fi27
                            }
                            v28 := GenS1{F0: 127, F1: "a", F2: fi25}
                            _ = v28
                        } else {
                            if (v5 != (v5 >> ((genFn0()) & 63))) {
                                if false {
                                    for fi29 := 0; fi29 < 16; fi29++ {
                                        v5 += (v5 - (v5 + (v5 * v5)))
                                        v30 := GenS3{F0: 100}
                                        if len(v18) >= 2 {
                                            v18 = v18[1:len(v18)]
                                        }
                                        var v31 int64 = 42
                                        var v32 float64 = float64((float32(float64((0.1 + 0.1))) + ((1e10 - 1e10) * 1e10)))
                                        if len(v18) > 0 {
                                            v18[0] = (^(fi25))
                                        }
                                        v33 := GenS1{F0: 1, F1: "\n\t", F2: (fi29 + int(v31))}
                                        _ = fi29
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    }
                                } else {
                                    v5 += v5
                                    for fi34 := 0; fi34 < 11; fi34++ {
                                        if true {
                                            if false {
                                                fmt.Println(v5, v2, v23)
                                            }
                                            var v35 uint = 4294967295
                                            if (-226 == 576) {
                                                if len(v18) > 0 {
                                                    v18[0] = fi25
                                                }
                                                v36 := GenS2{F0: 255}
                                                switch 1 {
                                                case 40:
                                                    fmt.Println(fi25)
                                                case -2147483648:
                                                    if false {
                                                        var v37 uint64 = 702
                                                        _ = v37
                                                    }
                                                }
                                                _ = v36
                                            }
                                            _ = v35
                                        } else {
                                            if true {
                                                switch (fi34 & (-(fi25))) {
                                                case 128:
                                                    if false {
                                                        var v38 int64 = -1
                                                        _ = v38
                                                    } else {
                                                        _ = len(v3)
                                                        var v39 bool = (false || true)
                                                        var v40 uint64 = 1
                                                        _ = v39
                                                        _ = v40
                                                    }
                                                case 2147483647:
                                                    var v41 uint8 = 2
                                                    _ = v41
                                                default:
                                                    if false {
                                                        fmt.Println(fi25, fi25, v2)
                                                        v42 := GenS4{F0: 235, F1: 508}
                                                        _ = v42
                                                    }
                                                }
                                                var v43 int64 = 1
                                                switch (fi25 >> ((0) & 63)) {
                                                case 2147483647:
                                                    if (v43 != (v43 >> ((genFn0()) & 63))) {
                                                        var v44 uint32 = 42
                                                        v1 = v44
                                                        v3[0] = ("" + (("" + ("a" + "hello")) + ("\n\t" + "foo bar")))
                                                        _ = v44
                                                    } else {
                                                        var v45 uint64 = 230
                                                        var v46 uint = uint(v1)
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                    fmt.Println(v23, fi25)
                                                    fmt.Println(v23)
                                                case 343:
                                                    v47 := GenS5{F0: (v23 ^ v23), F1: 2, F2: fi25}
                                                    _ = v47
                                                }
                                                var v48 string = ""
                                                _ = v43
                                                _ = v48
                                            }
                                            v49 := v4["hello"]
                                            _ = v49
                                        }
                                        _ = len(v18)
                                        fmt.Println(v2)
                                        _ = fi34
                                    }
                                }
                            }
                            if len(v18) > 0 {
                                v18[0] = -830
                            }
                        }
                    case 9223372036854775807:
                        var v50 uint = 65535
                        _ = v50
                    }
                    if false {
                        delete(v4, "a")
                    } else {
                        v23 = v23
                    }
                    _ = fi25
                }
                fmt.Println(v23, v5)
            } else {
                var v51 int = -441
                for fi52 := 0; fi52 < 4; fi52++ {
                    switch 816 {
                    case 128:
                        fmt.Println(v23)
                    default:
                        if true {
                            if len(v18) >= 2 {
                                v18 = v18[1:len(v18)]
                            }
                        } else {
                            var v53 bool = true
                            _ = v53
                        }
                    }
                    v18 = append(v18, 42)
                    if len(v18) > 0 {
                        v18[0] = v51
                    }
                    _ = fi52
                }
                _ = v51
            }
            v54 := v3[65535]
            var v55 int = -316
            _ = v54
            _ = v55
        } else {
            var v56 float64 = v2
            var v57 bool = true
            v57 = (2 >= 0)
            _ = v56
            _ = v57
        }
    }
    v1 += (v1 + 838)
    v1 += v1
    var v58 bool = (2 < uint8(v1))
    var v59 int8 = -1
    v60 := GenS2{F0: 127}
    var v61 uint = 762
    for fi62 := 0; fi62 < 6; fi62++ {
        for fi63 := 0; fi63 < 15; fi63++ {
            v64 := GenS3{F0: (v59 | int8(v23))}
            fmt.Println(v58)
            _ = fi63
            _ = v64
        }
        v23 = (v23 * (v23 % ((v23) | 1)))
        v61 = v61
        _ = fi62
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v60
    _ = v61
}
