package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 uint; F2 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32; F1 uint8; F2 uint16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int16 }
func genFn0(p0 uint8) uint64 {
    fmt.Println(p0)
    v5 := GenS1{F0: 9223372036854775807}
    var v6 string = ("" + "foo bar")
    v7 := []int{-567, -2147483648}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 uint, p1 uint8) int64 {
    v8 := []int{-753, 4294967295, -9223372036854775808, 127}
    fmt.Println(p0, p1)
    _ = len(v8)
    _ = p0
    _ = p1
    _ = v8
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 101
    var v3 uint = v2
    var v4 uint64 = 2
    if false {
        v3 = v2
    } else {
        var v9 float32 = -2.25
        v10 := GenS2{F0: 306, F1: 255, F2: 178}
        _ = v10
        _ = v9
    }
    var v11 float64 = (0.0 / (float64(0.1) - 1e10))
    var v12 uint64 = (^(v4))
    var v13 float64 = 1e10
    var v14 uint = 74
    v15 := GenS3{F0: -1}
    var v16 int16 = 32767
    v17 := GenS3{F0: -1}
    v13 = v13
    v13 += float64(0.1)
    for fi18 := 0; fi18 < 13; fi18++ {
        v14 += 692
        _ = fi18
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        switch "\n\t" {
        case "hello":
            var v20 int16 = v16
            _ = v20
        case "":
            fmt.Println(v16, v3)
        case "foo bar":
            var v21 uint32 = uint32(v12)
            _ = v21
        default:
            v12 -= v12
        }
        var v22 uint64 = 42
        v22 = (v22 ^ (v4 >> ((65535) & 63)))
        _ = fi19
        _ = v22
    }
    v23 := GenS3{F0: -1}
    switch -2147483648 {
    case 127:
        for fi24 := 0; fi24 < 8; fi24++ {
            v25 := []int{127, (fi24 / ((fi24) | 1)), -9223372036854775808, 127}
            _ = fi24
            _ = v25
        }
    case -1:
        if true {
            if false {
                if true {
                    var v26 bool = false
                    switch "foo bar" {
                    case "a":
                        var v27 uint = v14
                        _ = v27
                    case "\n\t":
                        switch ("\n\t" + "hello") {
                        case "\n\t":
                            v26 = false
                        default:
                            v3 += v14
                        }
                    case "":
                        var v28 int64 = -1
                        _ = v28
                    }
                    fmt.Println(v3, v11, v12)
                    _ = v26
                } else {
                    v29 := GenS4{F0: 42}
                    if (true || false) {
                        if (0 < 891) {
                            v30 := GenS5{F0: 661}
                            var v31 int = 128
                            var v32 uint32 = 325
                            _ = v30
                            _ = v31
                            _ = v32
                        } else {
                            v33 := []int{42, 0}
                            var v34 bool = (false || (uint8(v4) < 128))
                            _ = v33
                            _ = v34
                        }
                        v13 += v13
                    } else {
                        if false {
                            v35 := GenS6{F0: 42, F1: 2, F2: 2}
                            _ = v35
                        }
                        var v36 int = int(v4)
                        _ = v36
                    }
                    _ = v29
                }
                fmt.Println(v12, v14)
            } else {
                var v37 int = -55
                for fi38 := 0; fi38 < 2; fi38++ {
                    for fi39 := 0; fi39 < 9; fi39++ {
                        switch ("foo bar" + ("a" + "a")) {
                        case "foo bar":
                            v40 := GenS7{F0: 128}
                            _ = v40
                        }
                        var v41 uint64 = (v12 - uint64(v14))
                        var v42 uint16 = 32767
                        _ = fi39
                        _ = v41
                        _ = v42
                    }
                    v43 := []int{}
                    var v44 uint = v3
                    _ = fi38
                    _ = v43
                    _ = v44
                }
                var v45 int64 = int64(v16)
                _ = v37
                _ = v45
            }
            fmt.Println(v3)
            for fi46 := 0; fi46 < 15; fi46++ {
                switch fi46 {
                case -906:
                    if (42 != 2) {
                        v2 = 29
                        var v47 uint = v3
                        _ = v47
                    }
                case 127:
                    var v48 int = fi46
                    _ = v48
                case 2147483647:
                    for fi49 := 0; fi49 < 14; fi49++ {
                        fmt.Println(v16, v11, fi46)
                        var v50 int16 = (-(v16))
                        _ = fi49
                        _ = v50
                    }
                }
                _ = len(v1)
                _ = fi46
            }
        } else {
            v51 := []int{}
            var v52 string = ("\n\t" + ((("a" + "hello") + ("\n\t" + "hello")) + ("a" + ("" + "\n\t"))))
            switch 2147483647 {
            case -1:
                v53 := GenS8{F0: (v16 / (((v16 % ((1000) | 1))) | 1))}
                _ = v53
            case 32767:
                v54 := GenS7{F0: 250}
                _ = v54
            case 603:
                switch -912 {
                case 74:
                    _ = len(v51)
                case 9223372036854775807:
                    var v55 int32 = 1
                    _ = v55
                default:
                    switch ("a" + "a") {
                    case "foo bar":
                        var v56 int32 = 0
                        _ = v56
                    }
                }
            default:
                fmt.Println(v16, v52, v52)
            }
            _ = v51
            _ = v52
        }
    }
    v57 := v1["hello"]
    v58 := []int{-2147483648, int(v16), int(v4)}
    for fi59 := 0; fi59 < 1; fi59++ {
        if (18446744073709551615 < (v4 / ((1) | 1))) {
            switch fi59 {
            case 1:
                var v60 int = 9223372036854775807
                _ = v60
            case 128:
                for fi61 := 0; fi61 < 18; fi61++ {
                    if (((1 == -51) && (v11 < v13)) || ((0 == 1) || v57)) {
                        fmt.Println(v2, v2)
                        fmt.Println(v13, v2)
                    } else {
                        v16 += -1
                        var v62 float32 = 0.0
                        if v57 {
                            var v63 int = -311
                            _ = v63
                        } else {
                            var v64 int8 = -1
                            if (true || (v4 <= v12)) {
                                var v65 string = "a"
                                _ = len(v58)
                                var v66 uint8 = 0
                                _ = v65
                                _ = v66
                            }
                            if true {
                                switch fi61 {
                                case -2147483648:
                                    if len(v58) >= 2 {
                                        v58 = v58[1:len(v58)]
                                    }
                                default:
                                    switch "\n\t" {
                                    case "foo bar":
                                        v1[""] = true
                                    default:
                                        var v67 string = "\n\t"
                                        var v68 int16 = -1
                                        v69 := GenS5{F0: 4294967295}
                                        _ = v67
                                        _ = v68
                                        _ = v69
                                    }
                                }
                                fmt.Println(v57)
                            }
                            _ = v64
                        }
                        _ = v62
                    }
                    var v70 string = ("\n\t" + (("a" + "hello") + "a"))
                    var v71 uint = (v14 << (((v2 >> ((v3) & 63))) & 63))
                    _ = fi61
                    _ = v70
                    _ = v71
                }
            default:
                var v72 int64 = 9223372036854775807
                _ = v72
            }
            var v73 int64 = -1
            _ = v73
        }
        var v74 int = -494
        v11 += float64(((1.5 - -2.25) + (1e10 + (0.0 * 3.14159))))
        _ = fi59
        _ = v74
    }
    var v75 int64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v57:", v57)
    fmt.Println("slicelen v58:", len(v58), v58)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v57
    _ = v58
    _ = v75
}
