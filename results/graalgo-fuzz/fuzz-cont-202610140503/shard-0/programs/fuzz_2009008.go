package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 int; F1 int64; F2 int8 }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 uint32; F1 uint }
func genFn0(p0 string) int {
    fmt.Println(p0)
    v6 := GenS1{F0: 0, F1: -9223372036854775808, F2: 952}
    var v7 uint32 = 404
    fmt.Println(p0, p0)
    if true {
        v8 := []int{32768, 1, -994}
        v9 := GenS1{F0: -203, F1: 1, F2: 2}
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v6
    _ = v7
    return 332
}

func main() {
    v1 := []string{"", "a", "\n\t"}
    var v2 uint16 = 32767
    var v3 uint16 = 65535
    var v4 bool = (1 <= 57)
    v5 := []int{}
    for fi10 := 0; fi10 < 18; fi10++ {
        v11 := GenS2{F0: 333, F1: 569}
        var v12 float32 = ((-2.25 * 0.0) - 0.0)
        _ = fi10
        _ = v11
        _ = v12
    }
    v3 += (^(v2))
    if ((float32(0.1) - 0.0) > 0.0) {
        if false {
            var v13 int32 = -1
            switch "" {
            case "":
                fmt.Println(v13)
            case "hello":
                v14 := GenS2{F0: 588, F1: 32767}
                _ = v14
            case "foo bar":
                var v15 string = "a"
                _ = v15
            default:
                fmt.Println(v2, v3)
            }
            for fi16 := 0; fi16 < 12; fi16++ {
                var v17 int = fi16
                var v18 uint16 = 65535
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = v13
        } else {
            fmt.Println(v2, v4)
        }
    } else {
        for fi19 := 0; fi19 < 15; fi19++ {
            for fi20 := 0; fi20 < 19; fi20++ {
                v21 := GenS3{F0: 128, F1: 42, F2: -128}
                if v4 {
                    switch -970 {
                    case 32768:
                        switch (fi19 ^ fi20) {
                        case 32767:
                            var v22 uint64 = 18446744073709551615
                            v2 = v3
                            var v23 string = ""
                            _ = v22
                            _ = v23
                        case 9223372036854775807:
                            if v4 {
                                var v24 int8 = 127
                                v25 := GenS4{F0: 86, F1: 42}
                                _ = v24
                                _ = v25
                            } else {
                                v2 -= v2
                                v4 = true
                            }
                        case 2147483647:
                            v3 = 32767
                        }
                    case -1:
                        v3 = v2
                    }
                } else {
                    v2 += 32767
                }
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
    }
    v4 = ((-477 == 32768) || (2 >= v3))
    if v4 {
        v26 := GenS1{F0: 1, F1: -9223372036854775808, F2: 517}
        _ = v26
    }
    for fi27 := 0; fi27 < 10; fi27++ {
        switch (fi27 ^ 32768) {
        case -1:
            var v28 int8 = 127
            _ = v28
        default:
            if len(v5) > 0 {
                v5[0] = (fi27 + (-(fi27)))
            }
        }
        switch (fi27 % (((fi27 | fi27)) | 1)) {
        case -2147483648:
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
        case -9223372036854775808:
            switch (fi27 % (((-(fi27))) | 1)) {
            case 1:
                if true {
                    var v29 uint32 = 561
                    var v30 bool = ((0.0 / ((0.0 - 0.1) * -2.25)) == 1e10)
                    switch "a" {
                    case "hello":
                        var v31 uint64 = 311
                        _ = v31
                    case "":
                        switch -98 {
                        case 128:
                            for fi32 := 0; fi32 < 18; fi32++ {
                                _ = len(v1)
                                v33 := GenS3{F0: fi32, F1: -9223372036854775808, F2: -128}
                                _ = fi32
                                _ = v33
                            }
                        default:
                            for fi34 := 0; fi34 < 18; fi34++ {
                                v3 = v3
                                var v35 uint32 = v29
                                if (0 <= 42) {
                                    if v30 {
                                        v29 = (v35 + 2)
                                        v1 = append(v1, ("hello" + ("foo bar" + "hello")))
                                        fmt.Println(v29, v4)
                                        var v36 int32 = 462
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        v35 -= (v35 + 118)
                                        _ = v36
                                    } else {
                                        if len(v5) > 0 {
                                            v5[0] = fi27
                                        }
                                        v37 := GenS1{F0: fi34, F1: -1, F2: 18446744073709551615}
                                        v29 -= 42
                                        _ = v37
                                    }
                                    var v38 bool = ("hello" < "foo bar")
                                    switch "hello" {
                                    case "":
                                        if (270 < 492) {
                                            fmt.Println(fi34)
                                            v5 = append(v5, fi34)
                                        } else {
                                            v29 = (v35 * v29)
                                            var v39 float32 = 1.5
                                            _ = v39
                                        }
                                    case "hello":
                                        var v40 uint = 65535
                                        _ = v40
                                    case "a":
                                        var v41 uint16 = v2
                                        _ = v41
                                    }
                                    _ = v38
                                } else {
                                    fmt.Println(v3, v29)
                                    var v42 bool = true
                                    if false {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        var v43 float32 = 0.0
                                        var v44 uint16 = v3
                                        _ = v43
                                        _ = v44
                                    } else {
                                        if (255 < v3) {
                                            v2 -= v2
                                            var v45 int64 = -985
                                            fmt.Println(fi34, v45, v2)
                                            _ = v45
                                        } else {
                                            var v46 int = -2147483648
                                            _ = v46
                                        }
                                        v35 = 1
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v42
                                }
                                v29 += (v29 / (((v35 * v29)) | 1))
                                _ = fi34
                                _ = v35
                            }
                        }
                    }
                    _ = v29
                    _ = v30
                }
            case 32767:
                v4 = ((v4 || true) || (-1 >= 1))
            }
        default:
            if v4 {
                _ = len(v5)
            }
        }
        _ = fi27
    }
    var v47 uint16 = v2
    v4 = (0 != 417)
    fmt.Println(v47, v2, v4)
    switch "a" {
    case "\n\t":
        var v48 float32 = 1e10
        _ = v48
    }
    if ((1.5 < (0.1 * 0.0)) || ((2147483648 > 888) && (293 >= 811))) {
        if v4 {
            switch ("" + ("" + "a")) {
            case "foo bar":
                if (0.0 > 0.0) {
                    switch -9223372036854775808 {
                    case 513:
                        var v49 int16 = 1000
                        _ = v49
                    case 2147483647:
                        if v4 {
                            var v50 string = "a"
                            if (("\n\t" + "foo bar") <= v50) {
                                fmt.Println(v47, v4, v4)
                                v51 := GenS4{F0: 255, F1: 0}
                                v2 = 2
                                _ = v51
                            }
                            _ = v50
                        } else {
                            for fi52 := 0; fi52 < 10; fi52++ {
                                var v53 float32 = (1.5 - -2.25)
                                _ = fi52
                                _ = v53
                            }
                            var v54 float32 = (1e10 - 1.5)
                            _ = v54
                        }
                    case 0:
                        v55 := GenS5{F0: 982, F1: 0}
                        _ = v55
                    default:
                        if len(v5) > 0 {
                            v5[0] = 0
                        }
                    }
                } else {
                    v1 = append(v1, ("hello" + "foo bar"))
                }
            case "\n\t":
                var v56 uint8 = 1
                _ = v56
            case "a":
                switch 42 {
                case 320:
                    var v57 bool = true
                    _ = v57
                case 127:
                    if (false && v4) {
                        v47 = v2
                        var v58 float64 = 1.5
                        v2 -= (v2 * v2)
                        _ = v58
                    }
                case 32768:
                    fmt.Println(v3, v2, v3)
                }
            default:
                v5 = append(v5, 4294967295)
            }
        }
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v59 := GenS2{F0: 1000, F1: 32767}
    var v60 string = ""
    for fi61 := 0; fi61 < 15; fi61++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi61
    }
    for fi62 := 0; fi62 < 13; fi62++ {
        fmt.Println(v60)
        _ = fi62
    }
    var v63 int8 = 0
    v5 = append(v5, -757)
    fmt.Println(v47, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v60:", v60)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v59
    _ = v60
    _ = v63
}
