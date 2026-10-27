package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int64; F1 uint32; F2 bool }
type GenS8 struct { F0 int16; F1 uint64 }
func genFn0() bool {
    var v7 string = "\n\t"
    v8 := GenS1{F0: 16, F1: 42}
    v7 = v7
    _ = v7
    _ = v8
    return (-32768 <= -1)
}
func genFn1() int64 {
    for fi9 := 0; fi9 < 5; fi9++ {
        switch (("" + "a") + ("hello" + "hello")) {
        case "foo bar":
            switch fi9 {
            case -717:
                v10 := GenS1{F0: 42, F1: 0}
                _ = v10
            case 127:
                var v11 uint8 = 127
                _ = v11
            case 32767:
                for fi12 := 0; fi12 < 2; fi12++ {
                    v13 := GenS2{F0: (false && false), F1: 0}
                    var v14 float64 = 3.14159
                    v15 := GenS1{F0: 1, F1: 301}
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
            default:
                var v16 bool = false
                _ = v16
            }
        }
        _ = fi9
    }
    v17 := []int{}
    _ = v17
    return 1
}

func main() {
    v1 := []uint32{2147483647, 65535, 980}
    var v2 int64 = -9223372036854775808
    var v3 float32 = (0.0 + 1e10)
    v4 := make(map[int]string)
    var v5 float64 = 0.0
    var v6 uint32 = 894
    for fi18 := 0; fi18 < 13; fi18++ {
        switch fi18 {
        case 0:
            var v19 bool = (true && true)
            _ = v19
        default:
            if true {
                v2 = genFn1()
            } else {
                v20 := GenS1{F0: 42, F1: (v2 - (v2 - genFn1()))}
                _ = v20
            }
        }
        if true {
            fmt.Println(v6, v6)
            switch (^(fi18)) {
            case 32768:
                var v21 string = "a"
                _ = v21
            default:
                var v22 int8 = 127
                _ = v22
            }
        }
        if ((!true) && (2147483647 == 42)) {
            if (!true) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                switch "" {
                case "\n\t":
                    fmt.Println(v5, fi18)
                case "":
                    v23 := GenS3{F0: 17}
                    _ = v23
                case "hello":
                    fmt.Println(v2, v3)
                default:
                    v2 += (v2 + int64(fi18))
                }
            } else {
                for fi24 := 0; fi24 < 15; fi24++ {
                    var v25 uint64 = 655
                    v26 := GenS4{F0: v2, F1: fi24}
                    var v27 int32 = 129
                    if false {
                        fmt.Println(v2, v5)
                        v25 -= v25
                        fmt.Println(v6)
                        _ = len(v4)
                        v4[(fi18 << ((698) & 63))] = ("hello" + ("foo bar" + "hello"))
                        v28 := GenS2{F0: (false && true), F1: 5}
                        _ = v28
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
                switch (fi18 / (((fi18 >> ((0) & 63))) | 1)) {
                case 746:
                    v3 = v3
                case 2147483647:
                    var v29 int8 = 88
                    _ = v29
                case 4294967295:
                    switch (("\n\t" + "\n\t") + ("\n\t" + "\n\t")) {
                    case "hello":
                        v30 := v4[(fi18 - fi18)]
                        _ = v30
                    case "":
                        switch fi18 {
                        case 4294967295:
                            if true {
                                var v31 uint16 = 0
                                if len(v1) > 0 {
                                    v1[0] = 2147483648
                                }
                                _ = len(v4)
                                _ = v31
                            }
                        }
                    case "a":
                        var v32 bool = false
                        _ = v32
                    }
                default:
                    for fi33 := 0; fi33 < 5; fi33++ {
                        v5 = 1.5
                        fmt.Println(v5, v6)
                        v34 := GenS4{F0: (-(v2)), F1: (^(fi33))}
                        _ = fi33
                        _ = v34
                    }
                }
                fmt.Println(v5, v2, v2)
            }
            var v35 uint32 = v6
            switch "hello" {
            case "hello":
                var v36 int16 = 1
                _ = v36
            case "foo bar":
                fmt.Println(v35)
            case "\n\t":
                for fi37 := 0; fi37 < 3; fi37++ {
                    _ = len(v1)
                    _ = fi37
                }
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = v35
        } else {
            v38 := GenS1{F0: -102, F1: v2}
            for fi39 := 0; fi39 < 5; fi39++ {
                v40 := GenS5{F0: 127}
                switch fi39 {
                case 4294967295:
                    if false {
                        var v41 int32 = -2147483648
                        var v42 int32 = v41
                        var v43 uint8 = 2
                        _ = v41
                        _ = v42
                        _ = v43
                    } else {
                        switch ("a" + "foo bar") {
                        case "hello":
                            var v44 uint16 = 1
                            _ = v44
                        case "\n\t":
                            v4[int(v6)] = ("foo bar" + ("a" + ("hello" + "\n\t")))
                        case "a":
                            var v45 float64 = 0.0
                            _ = v45
                        }
                        var v46 uint = 2
                        var v47 float64 = v5
                        _ = v46
                        _ = v47
                    }
                }
                v2 -= (-(v2))
                _ = fi39
                _ = v40
            }
            if false {
                for fi48 := 0; fi48 < 16; fi48++ {
                    if (false || (-1 != 0)) {
                        var v49 uint16 = 27
                        v6 -= v6
                        _ = v49
                    }
                    if (-59 != 620) {
                        v6 -= (v6 + (v6 ^ (v6 << ((577) & 63))))
                    } else {
                        v50 := GenS3{F0: (v6 % (((^(v6))) | 1))}
                        v5 = 0.0
                        v3 = (-2.25 - v3)
                        _ = v50
                    }
                    v51 := v4[fi18]
                    _ = fi48
                    _ = v51
                }
                v52 := GenS4{F0: (v2 & (^(v2))), F1: (fi18 % ((fi18) | 1))}
                _ = v52
            } else {
                var v53 string = "hello"
                if (true || (true || true)) {
                    switch (("a" + v53) + (v53 + v53)) {
                    case "\n\t":
                        v54 := GenS4{F0: (v2 ^ 42), F1: (fi18 & fi18)}
                        _ = v54
                    default:
                        v53 = v53
                    }
                    delete(v4, fi18)
                }
                _ = v53
            }
            _ = v38
        }
        _ = fi18
    }
    if true {
        v55 := GenS6{F0: 0}
        switch 4294967295 {
        case 517:
            v56 := GenS7{F0: (v2 >> ((42) & 63)), F1: 2, F2: (1 < 18446744073709551615)}
            _ = v56
        }
        _ = v55
    } else {
        var v57 float64 = v5
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v3 = (v3 - (0.1 + 1e10))
        _ = v57
    }
    if (0 != 0) {
        if true {
            var v58 float64 = float64(1e10)
            switch 127 {
            case 42:
                var v59 uint8 = 42
                _ = v59
            default:
                var v60 uint32 = v6
                _ = v60
            }
            _ = v58
        }
    }
    var v61 uint = 0
    v6 -= v6
    for fi62 := 0; fi62 < 20; fi62++ {
        var v63 float64 = (float64(0.1) + v5)
        _ = fi62
        _ = v63
    }
    _ = len(v1)
    if ("" < "") {
        delete(v4, 407)
        if false {
            v2 = v2
        }
    }
    v3 = 0.1
    for fi64 := 0; fi64 < 3; fi64++ {
        for fi65 := 0; fi65 < 4; fi65++ {
            if true {
                if (-358 != 0) {
                    v61 = v61
                    var v66 int64 = genFn1()
                    fmt.Println(fi64, fi65)
                    _ = v66
                } else {
                    v2 += (v2 % ((genFn1()) | 1))
                }
                if true {
                    var v67 string = ("" + "hello")
                    switch v67 {
                    case "hello":
                        var v68 uint64 = 0
                        _ = v68
                    }
                    _ = v67
                } else {
                    var v69 float64 = 0.1
                    v70 := GenS8{F0: 841, F1: 18446744073709551615}
                    _ = v69
                    _ = v70
                }
            }
            var v71 string = ""
            var v72 int = -65
            _ = fi65
            _ = v71
            _ = v72
        }
        _ = fi64
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
}
