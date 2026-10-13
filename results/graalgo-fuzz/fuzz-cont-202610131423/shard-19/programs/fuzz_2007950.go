package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int32 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS3 struct { F0 bool; F1 uint8; F2 string }
func genFn0(p0 int32, p1 uint64) int {
    if (!(9223372036854775807 == 9223372036854775807)) {
        var v5 bool = false
        var v6 int8 = 0
        _ = v5
        _ = v6
    } else {
        switch "\n\t" {
        case "\n\t":
            if true {
                v7 := GenS1{F0: 1, F1: 4294967295, F2: (p0 & 42)}
                var v8 int32 = -593
                _ = v7
                _ = v8
            }
        case "hello":
            var v9 int32 = (p0 % (((p0 | p0)) | 1))
            _ = v9
        case "":
            v10 := GenS2{F0: -1, F1: 752, F2: 128}
            _ = v10
        default:
            v11 := GenS1{F0: 0, F1: 4294967295, F2: p0}
            _ = v11
        }
        if (!(!true)) {
            switch "foo bar" {
            case "":
                v12 := []int{855, 128, 32767}
                _ = v12
            case "foo bar":
                var v13 string = ("a" + "a")
                _ = v13
            default:
                var v14 uint64 = p1
                _ = v14
            }
            var v15 float32 = 1e10
            var v16 uint64 = (^(p1))
            _ = v15
            _ = v16
        }
    }
    var v17 int64 = 644
    switch ("a" + "\n\t") {
    case "\n\t":
        if true {
            var v18 uint64 = 4294967295
            v18 = 1
            for fi19 := 0; fi19 < 5; fi19++ {
                fmt.Println(fi19, fi19, fi19)
                _ = fi19
            }
            _ = v18
        } else {
            fmt.Println(p0)
            for fi20 := 0; fi20 < 16; fi20++ {
                v17 += 42
                p1 -= (p1 / (((p1 / ((0) | 1))) | 1))
                _ = fi20
            }
            if (1 < p0) {
                v21 := GenS1{F0: 0, F1: -1, F2: p0}
                if (2 == 2147483647) {
                    v22 := []int{}
                    _ = v22
                } else {
                    v23 := GenS2{F0: p0, F1: 255, F2: 214}
                    for fi24 := 0; fi24 < 4; fi24++ {
                        if false {
                            var v25 bool = true
                            fmt.Println(p0)
                            for fi26 := 0; fi26 < 19; fi26++ {
                                var v27 float32 = 0.1
                                v28 := []int{32768, (-(fi26)), -1}
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                            _ = v25
                        } else {
                            var v29 int32 = p0
                            var v30 uint64 = p1
                            for fi31 := 0; fi31 < 3; fi31++ {
                                v32 := GenS1{F0: 191, F1: (fi24 * 42), F2: (-(p0))}
                                fmt.Println(v29)
                                v33 := GenS1{F0: 2, F1: 32768, F2: 42}
                                _ = fi31
                                _ = v32
                                _ = v33
                            }
                            _ = v29
                            _ = v30
                        }
                        switch "" {
                        case "foo bar":
                            var v34 uint32 = uint32(p1)
                            _ = v34
                        case "":
                            v35 := GenS2{F0: (p0 >> ((395) & 63)), F1: 971, F2: 255}
                            _ = v35
                        default:
                            p0 -= p0
                        }
                        var v36 float32 = 0.1
                        _ = fi24
                        _ = v36
                    }
                    _ = v23
                }
                switch "\n\t" {
                case "a":
                    switch ("" + "a") {
                    case "hello":
                        var v37 int32 = (p0 >> ((1) & 63))
                        _ = v37
                    case "":
                        if ("\n\t" >= "hello") {
                            for fi38 := 0; fi38 < 17; fi38++ {
                                v17 -= (v17 + 42)
                                var v39 uint8 = 255
                                _ = fi38
                                _ = v39
                            }
                        }
                    }
                case "hello":
                    fmt.Println(p1, p0, v17)
                default:
                    var v40 bool = (!true)
                    _ = v40
                }
                _ = v21
            } else {
                v41 := GenS2{F0: (-(p0)), F1: 2, F2: 0}
                v42 := []int{4294967295, 1, -53}
                if true {
                    p0 -= (p0 * p0)
                    p1 = 1
                    if false {
                        var v43 int32 = (p0 - (p0 / ((p0) | 1)))
                        v42 = append(v42, 0)
                        v43 = (v43 * p0)
                        _ = v43
                    }
                }
                _ = v41
                _ = v42
            }
        }
    case "a":
        v44 := GenS2{F0: p0, F1: 32767, F2: 0}
        _ = v44
    case "foo bar":
        v45 := []int{-910}
        _ = v45
    default:
        p0 = (p0 + p0)
    }
    _ = p0
    _ = p1
    _ = v17
    return int(p1)
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]int)
    v3 := []int{}
    var v4 uint8 = 73
    v4 = v4
    _ = len(v1)
    if (1 <= 1) {
        var v46 int8 = 61
        switch "a" {
        case "foo bar":
            if true {
                v4 += v4
            }
        }
        _ = v46
    }
    switch "hello" {
    case "\n\t":
        v4 -= (v4 ^ v4)
    case "a":
        v47 := GenS3{F0: (-2147483648 <= -1), F1: (^(v4)), F2: ""}
        _ = v47
    case "foo bar":
        var v48 uint8 = (^(v4))
        _ = v48
    }
    var v49 int8 = 114
    v4 += (v4 ^ v4)
    v4 += (v4 & 2)
    v49 = v49
    if (!false) {
        switch genFn0(-375, 257) {
        case -9223372036854775808:
            v4 -= (v4 | 0)
        case -1:
            if true {
                fmt.Println(v4)
                v4 += v4
            } else {
                v2["hello"] = -77
            }
        case 128:
            var v50 int16 = -409
            _ = v50
        }
        switch 42 {
        case 2147483647:
            var v51 int64 = -9223372036854775808
            _ = v51
        case 366:
            v3 = append(v3, -2147483648)
        default:
            v4 += v4
        }
        var v52 int32 = 42
        _ = v52
    }
    delete(v1, 4294967295)
    if (1000 == 0) {
        for fi53 := 0; fi53 < 10; fi53++ {
            v49 = 127
            v49 += 100
            v54 := GenS1{F0: v4, F1: (-(fi53)), F2: -407}
            _ = fi53
            _ = v54
        }
    } else {
        v49 -= (v49 / ((28) | 1))
        var v55 uint16 = 804
        _ = v55
    }
    var v56 uint64 = 885
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v56
}
