package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8; F1 int64; F2 bool }
type GenS5 struct { F0 string; F1 uint32; F2 uint16 }
type GenS6 struct { F0 int; F1 int }
type GenS7 struct { F0 uint64; F1 int32; F2 int32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 32767
    var v3 uint8 = uint8(v2)
    var v4 bool = true
    v5 := make(map[int]uint32)
    var v6 int32 = 42
    v4 = v4
    if (1.5 <= (float64((1.5 + ((1e10 + float32(1e10)) - 3.14159))) / 3.14159)) {
        for fi7 := 0; fi7 < 2; fi7++ {
            if (((!(v2 <= (^(v2)))) && true) && false) {
                var v8 float64 = (1e10 - 0.0)
                _ = len(v5)
                for fi9 := 0; fi9 < 12; fi9++ {
                    v10 := v1[("\n\t" + ("" + (("" + "\n\t") + "a")))]
                    v3 = (^(v3))
                    _ = fi9
                    _ = v10
                }
                _ = v8
            } else {
                for fi11 := 0; fi11 < 1; fi11++ {
                    fmt.Println(v2, v4, v3)
                    v3 = v3
                    fmt.Println(fi7, fi11)
                    if true {
                        v12 := v1[("foo bar" + ("foo bar" + "a"))]
                        _ = v12
                    }
                    _ = fi11
                }
                fmt.Println(v3, v2)
                _ = len(v5)
            }
            _ = fi7
        }
        v13 := GenS1{F0: 4294967295}
        _ = v13
    } else {
        if v4 {
            switch ("" + "") {
            case "foo bar":
                v3 -= v3
            case "hello":
                var v14 uint32 = 65535
                _ = v14
            }
            switch "\n\t" {
            case "foo bar":
                var v15 string = "hello"
                _ = v15
            default:
                v16 := GenS2{F0: -7}
                _ = v16
            }
        } else {
            var v17 int = 1
            switch -997 {
            case 740:
                switch "" {
                case "":
                    var v18 uint64 = 190
                    _ = v18
                default:
                    if v4 {
                        if v4 {
                            v19 := GenS3{F0: 18446744073709551615}
                            v20 := GenS1{F0: 604}
                            _ = v19
                            _ = v20
                        }
                        v21 := GenS1{F0: 923}
                        fmt.Println(v2)
                        _ = v21
                    } else {
                        v22 := GenS3{F0: uint64(v3)}
                        _ = v22
                    }
                }
            case 254:
                var v23 uint = uint(v3)
                _ = v23
            case 32767:
                v24 := GenS4{F0: v3, F1: -9223372036854775808, F2: true}
                _ = v24
            }
            _ = v17
        }
        for fi25 := 0; fi25 < 8; fi25++ {
            if true {
                if (fi25 > (fi25 / (((fi25 | (fi25 >> ((897) & 63)))) | 1))) {
                    v26 := GenS1{F0: 0}
                    _ = v26
                }
                v27 := []int{(fi25 | fi25), (fi25 >> ((2) & 63)), (fi25 * (fi25 * (fi25 >> ((369) & 63))))}
                v6 -= -433
                _ = v27
            } else {
                if v4 {
                    fmt.Println(v6, v3, v2)
                    var v28 int32 = v6
                    for fi29 := 0; fi29 < 4; fi29++ {
                        fmt.Println(v28, fi25, v4)
                        v1["a"] = 2147483648
                        _ = fi29
                    }
                    _ = v28
                } else {
                    var v30 float32 = 1e10
                    _ = v30
                }
                if (!v4) {
                    var v31 uint16 = v2
                    v31 = v2
                    v32 := GenS3{F0: 42}
                    _ = v31
                    _ = v32
                }
                if ("foo bar" >= ("" + "\n\t")) {
                    v6 = v6
                    for fi33 := 0; fi33 < 5; fi33++ {
                        switch "foo bar" {
                        case "hello":
                            var v34 uint16 = v2
                            _ = v34
                        case "foo bar":
                            var v35 uint16 = 255
                            _ = v35
                        default:
                            fmt.Println(fi33, v6, fi33)
                        }
                        var v36 uint64 = 42
                        if v4 {
                            if true {
                                v37 := []int{2147483647, -1, (-(fi25))}
                                var v38 int8 = 100
                                var v39 int = 127
                                v40 := GenS4{F0: 127, F1: 0, F2: (false && v4)}
                                var v41 int = 32768
                                _ = v37
                                _ = v38
                                _ = v39
                                _ = v40
                                _ = v41
                            }
                            v42 := []int{(^(fi33))}
                            switch (fi25 + fi25) {
                            case 128:
                                v2 = (^(v2))
                            case -2147483648:
                                v43 := GenS5{F0: (("foo bar" + (("" + ("foo bar" + "hello")) + "")) + "hello"), F1: 65535, F2: uint16(v36)}
                                _ = v43
                            default:
                                var v44 int = fi25
                                _ = v44
                            }
                            _ = v42
                        }
                        if true {
                            fmt.Println(v36, v3, v3)
                            if (true && true) {
                                if v4 {
                                    v36 = 0
                                    v2 = (v2 >> ((uint(v36)) & 63))
                                    v3 -= v3
                                }
                            }
                            if ((v3 - v3) == v3) {
                                var v45 int64 = 9223372036854775807
                                _ = v45
                            }
                        } else {
                            v46 := []int{32, fi33, fi33}
                            v6 = v6
                            v47 := GenS6{F0: (fi33 + fi33), F1: 9223372036854775807}
                            _ = v46
                            _ = v47
                        }
                        _ = fi33
                        _ = v36
                    }
                    v3 = v3
                }
            }
            _ = fi25
        }
        fmt.Println(v2, v4, v3)
    }
    for fi48 := 0; fi48 < 17; fi48++ {
        v3 -= (^(v3))
        _ = fi48
    }
    v49 := []int{1, 1, -657, 32768}
    var v50 int = 9223372036854775807
    switch (v50 + v50) {
    case 127:
        if v4 {
            v51 := GenS7{F0: uint64(v3), F1: (v6 / ((v6) | 1)), F2: (v6 | v6)}
            _ = v51
        } else {
            v4 = true
        }
    case -9223372036854775808:
        switch (v50 >> ((42) & 63)) {
        case 32768:
            v3 = (v3 ^ (v3 * (^(v3))))
        case 2147483647:
            v6 -= v6
        case -927:
            fmt.Println(v6, v50, v50)
        default:
            if (32768 == v50) {
                switch 2147483647 {
                case 42:
                    for fi52 := 0; fi52 < 5; fi52++ {
                        v53 := GenS4{F0: 2, F1: 9223372036854775807, F2: false}
                        var v54 int = -1
                        _ = fi52
                        _ = v53
                        _ = v54
                    }
                case 32767:
                    var v55 int = v50
                    _ = v55
                case 163:
                    fmt.Println(v4, v2, v50)
                default:
                    fmt.Println(v50, v2, v6)
                }
            } else {
                for fi56 := 0; fi56 < 13; fi56++ {
                    var v57 int64 = -99
                    v2 -= (v2 & (^(v2)))
                    _ = fi56
                    _ = v57
                }
                var v58 float32 = 3.14159
                v3 += v3
                _ = v58
            }
        }
    }
    v2 = (v2 % ((v2) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
