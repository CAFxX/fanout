package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int32 }
type GenS2 struct { F0 int64; F1 int16; F2 int64 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 uint16; F1 int; F2 bool }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int16; F1 uint8 }
type GenS9 struct { F0 uint; F1 uint64; F2 string }
type GenS10 struct { F0 uint8; F1 bool }

func main() {
    v1 := make(map[int]int)
    var v2 int32 = -826
    v3 := make(map[uint32]string)
    v4 := []string{"hello", "foo bar", ("a" + "a")}
    v5 := []string{}
    switch 9223372036854775807 {
    case 128:
        var v6 string = "a"
        _ = v6
    }
    v7 := GenS1{F0: 32768, F1: 0, F2: 535}
    v2 -= (v2 ^ (v2 / (((v2 >> ((80) & 63))) | 1)))
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v8 int8 = 0
        var v9 uint8 = 2
        _ = v8
        _ = v9
    } else {
        if len(v4) > 0 {
            v4[0] = ("foo bar" + "hello")
        }
    }
    var v10 int8 = 61
    var v11 int = 42
    v2 += v2
    var v12 string = ("foo bar" + "hello")
    v13 := GenS1{F0: (v11 - (v11 ^ (v11 | (v11 << ((1) & 63))))), F1: 166, F2: v2}
    if true {
        var v14 int64 = 9223372036854775807
        _ = v14
    }
    var v15 uint64 = 42
    switch (v11 >> ((0) & 63)) {
    case 128:
        for fi16 := 0; fi16 < 20; fi16++ {
            if (!(v15 != (v15 | (v15 / ((v15) | 1))))) {
                if (false || true) {
                    var v17 int8 = 1
                    v18 := GenS1{F0: (fi16 - -364), F1: 0, F2: -1}
                    v19 := GenS2{F0: 0, F1: 592, F2: 42}
                    _ = v17
                    _ = v18
                    _ = v19
                }
            }
            if ((v10 % ((v10) | 1)) >= v10) {
                if false {
                    switch (v12 + (v12 + "a")) {
                    case "foo bar":
                        if true {
                            if ((^(v10)) < v10) {
                                v5 = append(v5, v12)
                                v10 = (v10 % (((v10 - (v10 * v10))) | 1))
                            } else {
                                _ = len(v4)
                                v15 += 531
                                v10 = (-(v10))
                            }
                            if ((fi16 == v11) || true) {
                                v20 := GenS2{F0: -1, F1: -32768, F2: 0}
                                for fi21 := 0; fi21 < 16; fi21++ {
                                    _ = len(v1)
                                    switch (v12 + v12) {
                                    case "hello":
                                        v10 = 100
                                        v10 -= (v10 & (v10 / ((-128) | 1)))
                                        var v22 uint8 = 171
                                        if true {
                                            switch 808 {
                                            case 79:
                                                if (0 >= 0) {
                                                    fmt.Println(v2, v10, fi21)
                                                    var v23 uint64 = v15
                                                    v23 = v23
                                                    _ = v23
                                                } else {
                                                    if (v10 < (v10 >> ((40) & 63))) {
                                                        v2 = 1
                                                        fmt.Println(v22)
                                                        v1[(fi21 >> ((42) & 63))] = v11
                                                    }
                                                    if (false && true) {
                                                        fmt.Println(v22, v2)
                                                    } else {
                                                        if false {
                                                            if (!false) {
                                                                _ = len(v1)
                                                                var v24 int8 = (v10 * v10)
                                                                _ = v24
                                                            } else {
                                                                fmt.Println(v2, v15, fi16)
                                                                var v25 uint8 = 128
                                                                v25 = (v25 ^ v25)
                                                                _ = v25
                                                            }
                                                        } else {
                                                            var v26 uint32 = 2147483647
                                                            var v27 uint16 = 65535
                                                            _ = v26
                                                            _ = v27
                                                        }
                                                        v15 = (v15 - v15)
                                                        if true {
                                                            v28 := GenS3{F0: 65535, F1: 42}
                                                            var v29 int = (fi21 << ((2) & 63))
                                                            if false {
                                                                var v30 uint64 = (v15 / ((2) | 1))
                                                                _ = v30
                                                            }
                                                            _ = v28
                                                            _ = v29
                                                        }
                                                    }
                                                    v15 += v15
                                                }
                                            case 127:
                                                v31 := GenS4{F0: 42}
                                                _ = v31
                                            default:
                                                v10 = (v10 / ((v10) | 1))
                                            }
                                            v22 = v22
                                            v32 := v1[v11]
                                            _ = v32
                                        }
                                        switch (("\n\t" + v12) + ("hello" + v12)) {
                                        case "":
                                            var v33 uint8 = v22
                                            _ = v33
                                        }
                                        _ = v22
                                    case "a":
                                        if false {
                                            v34 := GenS4{F0: -9223372036854775808}
                                            v2 = v2
                                            v10 = 60
                                            var v35 uint64 = 0
                                            _ = v34
                                            _ = v35
                                        }
                                    case "":
                                        v36 := GenS3{F0: 270, F1: 255}
                                        _ = v36
                                    default:
                                        if false {
                                            delete(v3, 0)
                                            if (false && true) {
                                                v15 += v15
                                                var v37 int64 = 696
                                                fmt.Println(v12)
                                                _ = v37
                                            } else {
                                                if (-2.25 <= 1e10) {
                                                    v38 := v3[65535]
                                                    var v39 uint = 4294967295
                                                    _ = v38
                                                    _ = v39
                                                } else {
                                                    v10 = (-(v10))
                                                    var v40 string = v12
                                                    _ = v40
                                                }
                                                var v41 int8 = 49
                                                var v42 uint = 0
                                                _ = v41
                                                _ = v42
                                            }
                                        } else {
                                            v4 = append(v4, ((v12 + "") + v12))
                                            delete(v3, 42)
                                        }
                                    }
                                    fmt.Println(v15, v11)
                                    _ = fi21
                                }
                                _ = v20
                            } else {
                                for fi43 := 0; fi43 < 9; fi43++ {
                                    delete(v1, fi16)
                                    switch 4294967295 {
                                    case 42:
                                        delete(v3, 2)
                                    case 4294967295:
                                        if (!((false || (v12 == "")) || true)) {
                                            var v44 float32 = (((0.1 * 0.0) - 1.5) * ((-2.25 / 1e10) / -2.25))
                                            var v45 int8 = (v10 >> ((255) & 63))
                                            v46 := GenS1{F0: (fi16 << ((255) & 63)), F1: 2, F2: -989}
                                            _ = v44
                                            _ = v45
                                            _ = v46
                                        }
                                    case -553:
                                        v10 += int8(v2)
                                    }
                                    v2 -= (v2 % ((v2) | 1))
                                    _ = fi43
                                }
                                fmt.Println(v11, v10, v10)
                            }
                        } else {
                            var v47 string = (((v12 + "a") + v12) + ((v12 + "") + "\n\t"))
                            if false {
                                v11 = 940
                                v48 := GenS1{F0: 127, F1: 243, F2: v2}
                                _ = v48
                            } else {
                                var v49 int8 = v10
                                v1[(v11 & -962)] = int(v49)
                                fmt.Println(fi16, v12, v2)
                                _ = v49
                            }
                            _ = v47
                        }
                    }
                    if (255 >= 255) {
                        v12 = v12
                    }
                } else {
                    if (!(!(!true))) {
                        v12 = "foo bar"
                        switch 2147483647 {
                        case 32768:
                            var v50 int64 = 42
                            _ = v50
                        case 2147483647:
                            if true {
                                for fi51 := 0; fi51 < 12; fi51++ {
                                    if true {
                                        v52 := GenS1{F0: 32768, F1: 140, F2: (v2 & 42)}
                                        v12 = "a"
                                        _ = v52
                                    }
                                    _ = fi51
                                }
                                v2 += (v2 << ((785) & 63))
                                var v53 uint = 42
                                _ = v53
                            }
                        default:
                            fmt.Println(v2)
                        }
                        fmt.Println(v12, v2, fi16)
                    } else {
                        fmt.Println(v10, v12, v12)
                    }
                }
                if (-9223372036854775808 <= 9223372036854775807) {
                    v12 = v12
                    fmt.Println(v10)
                    switch v12 {
                    case "":
                        fmt.Println(v11, v2, v11)
                    case "\n\t":
                        v54 := GenS5{F0: "foo bar", F1: 255}
                        _ = v54
                    default:
                        v55 := GenS1{F0: (v11 >> ((1) & 63)), F1: 0, F2: -2147483648}
                        _ = v55
                    }
                }
                if ("foo bar" < (v12 + "\n\t")) {
                    switch 832 {
                    case 32768:
                        v5 = append(v5, v12)
                    case -850:
                        var v56 int32 = v2
                        _ = v56
                    case -631:
                        if true {
                            v10 -= v10
                        } else {
                            v57 := GenS6{F0: 236, F1: fi16, F2: ((!false) && true)}
                            _ = v57
                        }
                    }
                } else {
                    fmt.Println(fi16, v10, v12)
                    if true {
                        switch (-(v11)) {
                        case 32768:
                            var v58 float32 = 0.1
                            _ = v58
                        default:
                            if (!(54 > (fi16 * fi16))) {
                                for fi59 := 0; fi59 < 17; fi59++ {
                                    switch 656 {
                                    case 278:
                                        fmt.Println(v12, v15, v12)
                                    }
                                    if (false && (true && true)) {
                                        v2 += v2
                                        switch (((v12 + (v12 + ((v12 + "foo bar") + "hello"))) + ("foo bar" + (((("a" + "foo bar") + "foo bar") + "\n\t") + v12))) + v12) {
                                        case "\n\t":
                                            var v60 float64 = (-2.25 - 3.14159)
                                            _ = v60
                                        case "":
                                            v10 = int8(v2)
                                        default:
                                            fmt.Println(v12, v2, v10)
                                        }
                                        v61 := GenS3{F0: 65535, F1: 2}
                                        _ = v61
                                    }
                                    v4 = append(v4, v12)
                                    _ = fi59
                                }
                                var v62 float32 = (-2.25 + 0.0)
                                v63 := GenS7{F0: 32767}
                                _ = v62
                                _ = v63
                            }
                        }
                        fmt.Println(v15, v15, v11)
                    } else {
                        var v64 uint = 4294967295
                        for fi65 := 0; fi65 < 15; fi65++ {
                            var v66 int32 = (v2 ^ v2)
                            if ((((0.1 * (1.5 - ((1e10 + (1.5 / 1e10)) - 0.0))) / (1e10 / 1e10)) + float64(1e10)) < 1e10) {
                                if true {
                                    v67 := GenS8{F0: 32767, F1: 231}
                                    switch 0 {
                                    case -2147483648:
                                        v11 -= (fi16 & 0)
                                    default:
                                        var v68 int16 = -1
                                        _ = v68
                                    }
                                    var v69 float64 = 1.5
                                    _ = v67
                                    _ = v69
                                }
                                if len(v5) > 0 {
                                    v5[0] = ((v12 + "a") + "\n\t")
                                }
                                v70 := GenS7{F0: 9223372036854775807}
                                var v71 bool = false
                                if v71 {
                                    v72 := GenS9{F0: (^(v64)), F1: 2, F2: v12}
                                    _ = v72
                                }
                                _ = v70
                                _ = v71
                            } else {
                                if (v64 > 65535) {
                                    if (uint(v15) > 18446744073709551615) {
                                        var v73 int64 = int64(v15)
                                        _ = v73
                                    }
                                    if true {
                                        v12 = v12
                                        fmt.Println(v15, v11, v12)
                                    } else {
                                        v4 = append(v4, ((v12 + v12) + v12))
                                        var v74 int64 = 0
                                        v11 -= (fi16 & -9223372036854775808)
                                        _ = v74
                                    }
                                }
                                fmt.Println(v2, fi65)
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            _ = fi65
                            _ = v66
                        }
                        _ = v64
                    }
                    var v75 int64 = -657
                    _ = v75
                }
            }
            switch v12 {
            case "a":
                if false {
                    v11 = v11
                } else {
                    var v76 uint8 = 0
                    v77 := GenS10{F0: (v76 >> ((64) & 63)), F1: false}
                    fmt.Println(v11)
                    _ = v76
                    _ = v77
                }
            case "hello":
                var v78 int16 = -32768
                _ = v78
            default:
                for fi79 := 0; fi79 < 9; fi79++ {
                    v1[fi16] = 32767
                    v1[int(v15)] = (-(fi16))
                    _ = fi79
                }
            }
            _ = fi16
        }
    }
    v11 = 535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
