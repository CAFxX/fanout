package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 int8; F1 uint64 }
func genFn0(p0 uint16, p1 int64, p2 int64) uint16 {
    p1 += 9223372036854775807
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1(p0 int16, p1 string, p2 int64) uint16 {
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := GenS1{F0: 42}
        var v8 int8 = 1
        v9 := GenS1{F0: 967}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := GenS2{F0: p1}
    if false {
        var v11 int8 = 53
        _ = v11
    } else {
        p2 = p2
        v12 := GenS2{F0: (p1 + "hello")}
        _ = v12
    }
    if ((-128 < -1) && true) {
        p0 -= 0
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 936
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{}
    var v3 float32 = 1.5
    var v4 bool = true
    var v5 int16 = -1
    if v4 {
        var v13 bool = v4
        if (!true) {
            v14 := v1[65535]
            for fi15 := 0; fi15 < 17; fi15++ {
                for fi16 := 0; fi16 < 6; fi16++ {
                    var v17 string = "foo bar"
                    var v18 string = (v17 + v17)
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                v19 := GenS3{F0: 2147483647, F1: 0}
                _ = fi15
                _ = v19
            }
            v3 = ((v3 - 1e10) + v3)
            _ = v14
        }
        _ = v13
    }
    var v20 int = 42
    for fi21 := 0; fi21 < 14; fi21++ {
        var v22 uint32 = 293
        switch v20 {
        case -885:
            switch v20 {
            case 42:
                if false {
                    v20 += 366
                    v23 := GenS4{F0: 42, F1: 1, F2: 255}
                    _ = v23
                }
            case 178:
                v22 += (v22 * v22)
            case -9223372036854775808:
                v3 -= 0.1
            }
        default:
            v22 -= 143
        }
        fmt.Println(v22, v20)
        _ = fi21
        _ = v22
    }
    if true {
        v20 += v20
    } else {
        v2 = append(v2, ("a" + "a"))
    }
    var v24 uint = 4294967295
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    var v25 int = (v20 | (v20 / ((-647) | 1)))
    v5 = v5
    fmt.Println(v5, v20)
    v20 = v25
    for fi26 := 0; fi26 < 11; fi26++ {
        delete(v1, 4294967295)
        _ = fi26
    }
    var v27 uint32 = 932
    switch "a" {
    case "foo bar":
        for fi28 := 0; fi28 < 4; fi28++ {
            switch "\n\t" {
            case "hello":
                for fi29 := 0; fi29 < 18; fi29++ {
                    var v30 int8 = 123
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi29
                    _ = v30
                }
            case "":
                if ((!true) || (!((v24 * 1) == 629))) {
                    switch 537 {
                    case 16:
                        v27 -= (v27 % (((^(v27))) | 1))
                    }
                    var v31 int32 = 234
                    v25 = v20
                    _ = v31
                } else {
                    switch "\n\t" {
                    case "a":
                        var v32 int32 = -1
                        _ = v32
                    case "":
                        if (-1 != 42) {
                            if v4 {
                                switch 2147483647 {
                                case 42:
                                    v4 = true
                                case 128:
                                    switch ("hello" + (("foo bar" + ("hello" + "hello")) + "a")) {
                                    case "a":
                                        var v33 string = "a"
                                        _ = v33
                                    case "\n\t":
                                        var v34 int16 = v5
                                        _ = v34
                                    case "":
                                        if (1 < 2) {
                                            for fi35 := 0; fi35 < 10; fi35++ {
                                                if false {
                                                    if (!false) {
                                                        if (0 > 2) {
                                                            var v36 bool = ((v27 % ((v27) | 1)) > v27)
                                                            _ = v36
                                                        }
                                                        _ = len(v1)
                                                        delete(v1, v27)
                                                    }
                                                    if v4 {
                                                        v25 = fi35
                                                    }
                                                    v1[(v27 % (((^(v27))) | 1))] = (fi28 << (((^(v24))) & 63))
                                                } else {
                                                    fmt.Println(v3)
                                                    if (!v4) {
                                                        if v4 {
                                                            v37 := GenS5{F0: 1, F1: (v5 | 0)}
                                                            _ = v37
                                                        }
                                                        v25 += -724
                                                        var v38 float64 = (1e10 * (float64((0.1 * v3)) + 3.14159))
                                                        _ = v38
                                                    }
                                                    switch v20 {
                                                    case 128:
                                                        var v39 float64 = 1e10
                                                        _ = v39
                                                    default:
                                                        v40 := GenS5{F0: 104, F1: v5}
                                                        _ = v40
                                                    }
                                                }
                                                _ = fi35
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                    default:
                                        switch v20 {
                                        case 2147483647:
                                            _ = len(v2)
                                        case -373:
                                            var v41 int64 = 1
                                            _ = v41
                                        case -1:
                                            if len(v2) > 0 {
                                                v2[0] = "hello"
                                            }
                                        }
                                    }
                                }
                                switch "\n\t" {
                                case "\n\t":
                                    if v4 {
                                        v42 := v1[v27]
                                        _ = v42
                                    }
                                case "":
                                    if v4 {
                                        v43 := GenS1{F0: 22}
                                        var v44 bool = v4
                                        _ = v43
                                        _ = v44
                                    }
                                case "hello":
                                    v5 = v5
                                default:
                                    for fi45 := 0; fi45 < 15; fi45++ {
                                        var v46 uint64 = 2
                                        _ = fi45
                                        _ = v46
                                    }
                                }
                            }
                            fmt.Println(v3, v4, v25)
                        }
                    }
                }
            case "a":
                v1[v27] = fi28
            default:
                if (!v4) {
                    var v47 int8 = 0
                    v48 := GenS6{F0: -1, F1: 408}
                    for fi49 := 0; fi49 < 1; fi49++ {
                        var v50 int = v25
                        var v51 uint64 = 42
                        var v52 uint = v24
                        _ = fi49
                        _ = v50
                        _ = v51
                        _ = v52
                    }
                    _ = v47
                    _ = v48
                }
            }
            var v53 int8 = -72
            if v4 {
                for fi54 := 0; fi54 < 19; fi54++ {
                    v27 = (^(v27))
                    switch fi28 {
                    case -653:
                        var v55 int16 = 1000
                        _ = v55
                    default:
                        fmt.Println(fi28, v5)
                    }
                    v4 = v4
                    _ = fi54
                }
            }
            _ = fi28
            _ = v53
        }
    case "\n\t":
        for fi56 := 0; fi56 < 13; fi56++ {
            if false {
                _ = len(v1)
            }
            _ = fi56
        }
    }
    var v57 uint16 = 42
    v2 = append(v2, "\n\t")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v57
}
