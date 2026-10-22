package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint; F1 uint8; F2 int64 }
type GenS3 struct { F0 string; F1 int8; F2 uint32 }

func main() {
    v1 := make(map[int]string)
    var v2 int32 = 27
    var v3 uint32 = 2147483647
    v4 := []uint32{v3, v3, (v3 % ((42) | 1)), v3}
    var v5 uint8 = 1
    v6 := []uint32{65535, 1, v3, v3}
    if true {
        v5 = v5
        if len(v6) > 0 {
            v6[0] = v3
        }
        if true {
            fmt.Println(v5, v2)
        } else {
            if true {
                v3 += 2147483648
            } else {
                if ((((true && true) && (false && (true && (2 >= uint64(v5))))) || (127 > 0)) || true) {
                    v2 = v2
                    for fi7 := 0; fi7 < 6; fi7++ {
                        for fi8 := 0; fi8 < 1; fi8++ {
                            if false {
                                var v9 uint64 = 18446744073709551615
                                _ = v9
                            } else {
                                v5 = (^(v5))
                                v10 := GenS1{F0: "\n\t", F1: (("foo bar" + "\n\t") + "hello")}
                                _ = v10
                            }
                            if ((1000 > -147) || false) {
                                var v11 int16 = -992
                                _ = v11
                            }
                            _ = fi8
                        }
                        v12 := v1[fi7]
                        _ = fi7
                        _ = v12
                    }
                    v13 := GenS1{F0: ("foo bar" + "\n\t"), F1: ""}
                    _ = v13
                } else {
                    v5 = (v5 + (v5 * (v5 % (((v5 << ((uint(v5)) & 63))) | 1))))
                }
                switch ((("" + "foo bar") + (("\n\t" + "foo bar") + "")) + "\n\t") {
                case "\n\t":
                    if (false || false) {
                        var v14 float32 = 0.1
                        _ = v14
                    } else {
                        v15 := GenS1{F0: "a", F1: "\n\t"}
                        var v16 int16 = -628
                        _ = len(v6)
                        _ = v15
                        _ = v16
                    }
                default:
                    if true {
                        var v17 uint64 = 279
                        var v18 int32 = (v2 | 1)
                        _ = v17
                        _ = v18
                    }
                }
                v19 := GenS2{F0: 18446744073709551615, F1: (v5 / (((^(v5))) | 1)), F2: 264}
                _ = v19
            }
        }
    }
    var v20 int64 = 1
    v20 = v20
    if false {
        v3 -= v3
    } else {
        switch 285 {
        case 32768:
            if (!false) {
                delete(v1, 0)
                if false {
                    v1[9223372036854775807] = ("hello" + "\n\t")
                    for fi21 := 0; fi21 < 12; fi21++ {
                        v22 := GenS2{F0: uint(v5), F1: v5, F2: v20}
                        _ = len(v6)
                        _ = fi21
                        _ = v22
                    }
                } else {
                    for fi23 := 0; fi23 < 12; fi23++ {
                        v3 -= 671
                        _ = fi23
                    }
                    switch "\n\t" {
                    case "\n\t":
                        delete(v1, 1)
                    }
                }
            } else {
                v2 = -1
            }
        case -582:
            for fi24 := 0; fi24 < 3; fi24++ {
                if false {
                    v20 += v20
                    delete(v1, -439)
                    v25 := GenS1{F0: ("" + "hello"), F1: "foo bar"}
                    _ = v25
                }
                _ = fi24
            }
        case 127:
            v6 = append(v6, v3)
        default:
            if true {
                if (false || (!(!false))) {
                    v26 := GenS1{F0: "\n\t", F1: ("" + "hello")}
                    v5 = v5
                    v27 := GenS1{F0: ("\n\t" + ("" + "foo bar")), F1: "a"}
                    _ = v26
                    _ = v27
                }
            }
        }
        if ((v5 <= (v5 + (v5 & 0))) && (1e10 > 0.1)) {
            for fi28 := 0; fi28 < 5; fi28++ {
                for fi29 := 0; fi29 < 4; fi29++ {
                    if len(v6) > 0 {
                        v6[0] = (v3 - (v3 / ((v3) | 1)))
                    }
                    var v30 uint = 65535
                    v31 := GenS2{F0: (v30 << ((18446744073709551615) & 63)), F1: uint8(v3), F2: (v20 * int64(fi28))}
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                if (1.5 != 3.14159) {
                    if true {
                        fmt.Println(v5)
                    }
                    var v32 uint32 = 1
                    v32 += 65535
                    _ = v32
                }
                _ = fi28
            }
            v33 := GenS3{F0: "foo bar", F1: 100, F2: (v3 | v3)}
            var v34 uint8 = (v5 | (v5 & 42))
            _ = v33
            _ = v34
        }
        var v35 float64 = 1.5
        _ = v35
    }
    v36 := GenS3{F0: "\n\t", F1: 100, F2: v3}
    switch 2147483647 {
    case 9223372036854775807:
        v37 := GenS3{F0: ("foo bar" + ""), F1: 38, F2: (v3 & 2147483647)}
        _ = v37
    default:
        var v38 bool = true
        _ = v38
    }
    for fi39 := 0; fi39 < 13; fi39++ {
        fmt.Println(v5)
        for fi40 := 0; fi40 < 20; fi40++ {
            v2 += (-(v2))
            v3 += 65535
            _ = len(v1)
            _ = fi40
        }
        v41 := v1[(fi39 / ((fi39) | 1))]
        _ = fi39
        _ = v41
    }
    v42 := GenS1{F0: "\n\t", F1: "hello"}
    var v43 float32 = (1e10 - 1e10)
    for fi44 := 0; fi44 < 3; fi44++ {
        for fi45 := 0; fi45 < 18; fi45++ {
            if true {
                v3 += 0
                var v46 int16 = -1
                var v47 uint = 42
                var v48 int = (^(fi44))
                _ = v46
                _ = v47
                _ = v48
            } else {
                v3 += 1
            }
            v49 := GenS1{F0: "\n\t", F1: "a"}
            v50 := GenS2{F0: 818, F1: (v5 / ((v5) | 1)), F2: 9223372036854775807}
            _ = fi45
            _ = v49
            _ = v50
        }
        v3 = v3
        _ = fi44
    }
    if ((1 == int(v20)) && true) {
        switch ("foo bar" + "hello") {
        case "\n\t":
            for fi51 := 0; fi51 < 7; fi51++ {
                for fi52 := 0; fi52 < 9; fi52++ {
                    fmt.Println(v2, v2, v5)
                    _ = fi52
                }
                var v53 int = 297
                _ = fi51
                _ = v53
            }
        case "":
            v43 -= (v43 / ((v43 + v43) - v43))
        case "foo bar":
            var v54 int = -470
            _ = v54
        }
        v5 -= 2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v36
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
