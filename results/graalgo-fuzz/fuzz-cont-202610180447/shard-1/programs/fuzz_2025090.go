package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint32; F2 uint }
type GenS3 struct { F0 string; F1 int; F2 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint8; F1 uint }

func main() {
    v1 := []int{1, 328, 9223372036854775807}
    v2 := make(map[string]uint32)
    var v3 uint64 = 4294967295
    v4 := make(map[int]string)
    var v5 int32 = 42
    v3 -= 4294967295
    if ("foo bar" < "hello") {
        switch ("a" + ("" + "foo bar")) {
        case "foo bar":
            if (!((!false) || true)) {
                v5 += (v5 ^ v5)
                v3 = (v3 >> ((0) & 63))
                v5 -= -482
            }
        case "hello":
            v6 := GenS1{F0: 500, F1: 1}
            _ = v6
        default:
            v5 = v5
        }
        if (-919 < 2147483647) {
            v7 := GenS2{F0: true, F1: 65535, F2: 621}
            _ = len(v1)
            for fi8 := 0; fi8 < 12; fi8++ {
                if true {
                    fmt.Println(v5, fi8, fi8)
                    v9 := GenS3{F0: "a", F1: fi8, F2: int64(v5)}
                    v5 -= v5
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 19; fi10++ {
                        if false {
                            var v11 bool = true
                            v11 = false
                            _ = v11
                        }
                        v12 := GenS4{F0: 4294967295}
                        _ = fi10
                        _ = v12
                    }
                }
                _ = fi8
            }
            _ = v7
        } else {
            for fi13 := 0; fi13 < 16; fi13++ {
                var v14 int32 = v5
                if true {
                    if false {
                        fmt.Println(v14, fi13)
                    } else {
                        var v15 float64 = (float64(1e10) * 3.14159)
                        fmt.Println(v14, v3, v5)
                        v16 := GenS4{F0: 541}
                        _ = v15
                        _ = v16
                    }
                    if false {
                        v17 := GenS4{F0: 18446744073709551615}
                        if true {
                            var v18 int8 = 127
                            switch (-(fi13)) {
                            case -983:
                                var v19 uint8 = 127
                                _ = v19
                            case 632:
                                if len(v1) > 0 {
                                    v1[0] = fi13
                                }
                            default:
                                switch fi13 {
                                case 42:
                                    v20 := GenS2{F0: ((fi13 <= -458) || false), F1: 903, F2: 791}
                                    _ = v20
                                }
                            }
                            for fi21 := 0; fi21 < 6; fi21++ {
                                v14 -= 0
                                var v22 float64 = 1.5
                                v1 = append(v1, -437)
                                _ = fi21
                                _ = v22
                            }
                            _ = v18
                        }
                        v23 := GenS3{F0: "", F1: (fi13 * 615), F2: 9223372036854775807}
                        _ = v17
                        _ = v23
                    } else {
                        if false {
                            if true {
                                var v24 uint16 = 32767
                                v5 = v14
                                var v25 int16 = 0
                                _ = v24
                                _ = v25
                            }
                            v3 = v3
                        } else {
                            v4[(fi13 & (fi13 >> ((628) & 63)))] = ""
                        }
                        for fi26 := 0; fi26 < 2; fi26++ {
                            v5 = (v5 / ((v5) | 1))
                            _ = fi26
                        }
                    }
                    v27 := v2[""]
                    _ = v27
                } else {
                    var v28 uint32 = 65535
                    var v29 uint8 = 67
                    _ = v28
                    _ = v29
                }
                _ = fi13
                _ = v14
            }
        }
        var v30 uint32 = 2
        _ = v30
    }
    var v31 float32 = 3.14159
    v2[""] = 65535
    v3 -= (^(v3))
    v31 = v31
    v31 = 1e10
    switch "a" {
    case "\n\t":
        if true {
            if false {
                v32 := GenS4{F0: 112}
                _ = v32
            }
        } else {
            v33 := GenS5{F0: -172}
            _ = v33
        }
    case "foo bar":
        v3 = (^(v3))
    case "":
        if true {
            var v34 int32 = (v5 >> ((65535) & 63))
            var v35 uint16 = 65535
            _ = v34
            _ = v35
        } else {
            switch 32768 {
            case 295:
                var v36 int32 = v5
                _ = v36
            case 4294967295:
                var v37 uint32 = 2147483648
                _ = v37
            }
            for fi38 := 0; fi38 < 2; fi38++ {
                v31 += v31
                _ = fi38
            }
        }
    default:
        for fi39 := 0; fi39 < 5; fi39++ {
            switch (fi39 & fi39) {
            case 32768:
                var v40 uint64 = v3
                _ = v40
            case -793:
                switch (fi39 & (fi39 << ((1) & 63))) {
                case 0:
                    v41 := GenS6{F0: 522}
                    _ = v41
                case 1:
                    var v42 string = "a"
                    _ = v42
                case 2147483647:
                    var v43 string = "foo bar"
                    _ = v43
                }
            }
            var v44 int32 = v5
            fmt.Println(fi39)
            _ = fi39
            _ = v44
        }
    }
    if len(v1) > 0 {
        v1[0] = -465
    }
    v45 := GenS7{F0: (!((^(v3)) <= v3))}
    if true {
        var v46 uint64 = 42
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v46
    } else {
        v1 = append(v1, 9223372036854775807)
        for fi47 := 0; fi47 < 7; fi47++ {
            var v48 int = (fi47 * fi47)
            if ((!(!(1 >= 255))) && true) {
                v5 += v5
                for fi49 := 0; fi49 < 20; fi49++ {
                    var v50 bool = ((v3 | (^(v3))) >= 4294967295)
                    if v50 {
                        v50 = (false || (2147483648 == 576))
                        fmt.Println(v31)
                    } else {
                        var v51 int8 = 0
                        fmt.Println(v50, v3)
                        _ = v51
                    }
                    fmt.Println(v50)
                    _ = fi49
                    _ = v50
                }
                fmt.Println(v3, v48)
            } else {
                for fi52 := 0; fi52 < 17; fi52++ {
                    var v53 uint32 = 2
                    v53 += v53
                    _ = fi52
                    _ = v53
                }
                for fi54 := 0; fi54 < 3; fi54++ {
                    var v55 bool = true
                    switch "foo bar" {
                    case "a":
                        v56 := GenS8{F0: 0, F1: 255}
                        _ = v56
                    case "":
                        var v57 int64 = 0
                        _ = v57
                    }
                    if (v5 < 1) {
                        var v58 int32 = v5
                        v58 += (v58 + (v58 ^ (v5 - v58)))
                        fmt.Println(v5, v58, v3)
                        var v59 uint = 18446744073709551615
                        v60 := v2["hello"]
                        _ = v58
                        _ = v59
                        _ = v60
                    } else {
                        v61 := GenS3{F0: "\n\t", F1: 127, F2: -719}
                        v31 = v31
                        _ = v61
                    }
                    v4[(v48 * fi54)] = "a"
                    switch "foo bar" {
                    case "\n\t":
                        v55 = (2 >= 719)
                    }
                    _ = fi54
                    _ = v55
                }
                var v62 int32 = 2147483647
                _ = v62
            }
            _ = fi47
            _ = v48
        }
    }
    v1 = append(v1, 32768)
    v3 -= (v3 + 353)
    v31 = v31
    v5 -= (v5 - -1)
    switch "" {
    case "\n\t":
        fmt.Println(v5)
    case "":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "foo bar":
        v5 = v5
    }
    for fi63 := 0; fi63 < 7; fi63++ {
        v3 -= 42
        var v64 uint8 = 127
        v64 = v64
        _ = fi63
        _ = v64
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v45
    _ = v5
}
