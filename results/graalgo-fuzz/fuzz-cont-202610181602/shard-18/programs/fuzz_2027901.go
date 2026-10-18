package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 int32; F1 int32; F2 string }
type GenS7 struct { F0 int8; F1 bool; F2 string }
func genFn0(p0 int, p1 bool, p2 int64) int8 {
    v6 := GenS1{F0: p2, F1: -556, F2: 42}
    v7 := []int{p0, -1, (p0 % ((p0) | 1)), (p0 << ((988) & 63))}
    for fi8 := 0; fi8 < 13; fi8++ {
        p1 = (897 == 1)
        v9 := GenS1{F0: p2, F1: p0, F2: 127}
        var v10 int16 = 0
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 bool = (0.1 > 3.14159)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 uint = 590
    var v2 uint = v1
    var v3 string = ("" + (("hello" + "hello") + "a"))
    v4 := make(map[string]string)
    var v5 int8 = 42
    var v12 uint8 = 1
    v13 := v4[("foo bar" + (v3 + ""))]
    if true {
        v14 := GenS1{F0: -9223372036854775808, F1: -9223372036854775808, F2: (v12 ^ (v12 & v12))}
        for fi15 := 0; fi15 < 7; fi15++ {
            v5 += v5
            fmt.Println(v3, v12, v1)
            _ = fi15
        }
        var v16 float64 = (0.1 / float64(-2.25))
        _ = v14
        _ = v16
    } else {
        v17 := []int{676, 4294967295}
        _ = v17
    }
    fmt.Println(v13, v3, v3)
    for fi18 := 0; fi18 < 13; fi18++ {
        v12 = (^(v12))
        fmt.Println(v12)
        if true {
            v19 := GenS2{F0: 9223372036854775807}
            var v20 float32 = 3.14159
            if true {
                v21 := v4[("foo bar" + v3)]
                _ = v21
            } else {
                v22 := GenS1{F0: -1, F1: (fi18 >> ((v1) & 63)), F2: (v12 << ((v2) & 63))}
                v23 := []int{32767}
                switch "hello" {
                case "a":
                    for fi24 := 0; fi24 < 19; fi24++ {
                        var v25 uint16 = 349
                        var v26 uint32 = 42
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                default:
                    for fi27 := 0; fi27 < 7; fi27++ {
                        fmt.Println(v1, v12, v20)
                        v20 = v20
                        v13 = v13
                        _ = fi27
                    }
                }
                _ = v22
                _ = v23
            }
            _ = v19
            _ = v20
        } else {
            for fi28 := 0; fi28 < 13; fi28++ {
                if true {
                    fmt.Println(fi28)
                } else {
                    v29 := GenS1{F0: 9223372036854775807, F1: fi18, F2: 109}
                    _ = v29
                }
                var v30 uint64 = 2
                if (1 != -12) {
                    v31 := []int{fi18, -1, -9223372036854775808}
                    _ = v31
                } else {
                    var v32 uint32 = uint32(v12)
                    var v33 int = -543
                    v34 := GenS3{F0: v32, F1: (v2 ^ 255)}
                    v5 = -1
                    var v35 uint8 = 128
                    _ = v32
                    _ = v33
                    _ = v34
                    _ = v35
                }
                _ = fi28
                _ = v30
            }
        }
        _ = fi18
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        v37 := GenS4{F0: 989, F1: 1}
        v4["hello"] = (v3 + "hello")
        switch v3 {
        case "\n\t":
            switch "a" {
            case "hello":
                v38 := GenS3{F0: 4294967295, F1: v2}
                _ = v38
            case "a":
                for fi39 := 0; fi39 < 10; fi39++ {
                    if (v2 >= v1) {
                        fmt.Println(v1)
                        var v40 uint32 = 2
                        switch fi36 {
                        case 2147483647:
                            v41 := GenS4{F0: 388, F1: v2}
                            if (2 != 971) {
                                v2 = (v2 << ((v2) & 63))
                                if (((v5 % ((v5) | 1)) <= v5) && true) {
                                    v42 := GenS1{F0: 1, F1: (fi36 | fi36), F2: v12}
                                    _ = v42
                                } else {
                                    v43 := GenS5{F0: 634, F1: -252}
                                    _ = v43
                                }
                                v5 = genFn0((fi36 + fi39), false, -9223372036854775808)
                                fmt.Println(v2, fi36)
                                var v44 int32 = 0
                                _ = v44
                            }
                            v5 = genFn0(fi36, (v12 > (v12 >> ((v1) & 63))), 604)
                            _ = v41
                        case 128:
                            v45 := []int{(-(fi39)), fi36, (fi39 / (((fi39 / (((fi39 + fi36)) | 1))) | 1))}
                            _ = v45
                        }
                        _ = v40
                    }
                    switch 127 {
                    case -1:
                        var v46 bool = (2 == 42)
                        _ = v46
                    case 128:
                        v12 += v12
                        v13 = ((v3 + v13) + (("foo bar" + "") + v13))
                    case 9223372036854775807:
                        switch (fi39 / ((fi39) | 1)) {
                        case -444:
                            var v47 int64 = 9223372036854775807
                            _ = v47
                        default:
                            var v48 int64 = -9223372036854775808
                            _ = v48
                        }
                    }
                    var v49 string = v13
                    _ = fi39
                    _ = v49
                }
            }
        case "foo bar":
            switch -518 {
            case -2147483648:
                v12 -= (v12 | uint8(v1))
            case 32768:
                switch (fi36 % (((fi36 % ((2147483647) | 1))) | 1)) {
                case -278:
                    fmt.Println(v1, v12)
                case 417:
                    fmt.Println(v1, v13)
                }
            }
        default:
            var v50 uint = (v1 ^ (v1 % (((^(v1))) | 1)))
            _ = v50
        }
        _ = fi36
        _ = v37
    }
    for fi51 := 0; fi51 < 14; fi51++ {
        var v52 int64 = -226
        v52 = v52
        _ = fi51
        _ = v52
    }
    if (uint16(v1) < 0) {
        if (532 > -302) {
            v3 = v13
        }
    }
    v53 := GenS6{F0: 0, F1: 795, F2: "foo bar"}
    for fi54 := 0; fi54 < 17; fi54++ {
        v12 = v12
        _ = fi54
    }
    if (1e10 <= (-2.25 - (-2.25 + float64((-2.25 / 3.14159))))) {
        v55 := GenS7{F0: -128, F1: (true || (v2 == v2)), F2: v13}
        _ = v55
    }
    v12 = (^(v12))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
}
