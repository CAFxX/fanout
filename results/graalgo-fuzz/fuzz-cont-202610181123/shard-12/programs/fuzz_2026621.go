package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint8; F1 uint; F2 int }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool; F1 int16 }

func main() {
    var v1 float64 = 0.1
    v2 := []string{(("" + "\n\t") + "\n\t"), "", "a", "hello"}
    v3 := make(map[string]int)
    var v4 uint64 = 219
    v5 := make(map[int]string)
    var v6 uint16 = 954
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: (!false), F1: false}
        if (95 <= v4) {
            for fi9 := 0; fi9 < 14; fi9++ {
                v1 -= v1
                switch ("a" + "a") {
                case "foo bar":
                    if (false && false) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        switch "\n\t" {
                        case "\n\t":
                            v10 := GenS2{F0: 27, F1: 4294967295, F2: fi9}
                            _ = v10
                        case "foo bar":
                            var v11 int64 = 42
                            _ = v11
                        default:
                            var v12 uint8 = 255
                            _ = v12
                        }
                    }
                default:
                    v4 = 4294967295
                }
                _ = fi9
            }
        } else {
            v3[("foo bar" + "\n\t")] = (fi7 ^ (fi7 & 424))
            var v13 bool = true
            if (!true) {
                var v14 float64 = 3.14159
                if v13 {
                    v15 := GenS2{F0: 91, F1: 255, F2: fi7}
                    var v16 uint32 = 138
                    v17 := GenS3{F0: fi7, F1: int8(v6)}
                    _ = v15
                    _ = v16
                    _ = v17
                }
                var v18 int8 = 127
                _ = v14
                _ = v18
            } else {
                switch (fi7 % (((fi7 + fi7)) | 1)) {
                case 855:
                    var v19 int32 = 449
                    _ = v19
                case 4294967295:
                    for fi20 := 0; fi20 < 9; fi20++ {
                        var v21 int = (fi20 >> ((0) & 63))
                        _ = fi20
                        _ = v21
                    }
                }
            }
            _ = v13
        }
        _ = fi7
        _ = v8
    }
    _ = len(v2)
    for fi22 := 0; fi22 < 7; fi22++ {
        if ((fi22 >> ((2) & 63)) > 306) {
            switch (fi22 + (-(fi22))) {
            case 575:
                var v23 uint64 = 1
                _ = v23
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        v24 := GenS3{F0: fi22, F1: -1}
        fmt.Println(v4)
        _ = fi22
        _ = v24
    }
    _ = len(v3)
    switch 32768 {
    case 127:
        var v25 int32 = 0
        _ = v25
    default:
        if (522 != -136) {
            _ = len(v2)
            _ = len(v3)
        } else {
            var v26 int64 = 42
            for fi27 := 0; fi27 < 13; fi27++ {
                v4 = 211
                var v28 uint = 0
                var v29 int64 = v26
                _ = fi27
                _ = v28
                _ = v29
            }
            v26 = 0
            _ = v26
        }
    }
    var v30 uint32 = 2
    v4 = (^(v4))
    var v31 uint = 0
    if true {
        if (!(42 <= 1)) {
            fmt.Println(v30, v4, v4)
            switch -2147483648 {
            case -9223372036854775808:
                if (9223372036854775807 != -1) {
                    v1 = v1
                    v32 := GenS2{F0: 42, F1: (v31 % ((v31) | 1)), F2: -796}
                    switch "" {
                    case "\n\t":
                        v4 += 4294967295
                    case "a":
                        fmt.Println(v31, v6, v31)
                    }
                    _ = v32
                } else {
                    v33 := GenS1{F0: true, F1: false}
                    _ = v33
                }
            default:
                if ((112 != 0) && (-62 <= -1)) {
                    v31 += v31
                    _ = len(v2)
                }
            }
        }
    }
    if ((v30 >> (((v31 / ((v31) | 1))) & 63)) > 2147483647) {
        v6 = (^(v6))
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v34 uint32 = v30
        _ = v34
    }
    v35 := GenS4{F0: 2}
    v6 = v6
    v36 := GenS4{F0: (v4 & v4)}
    v37 := GenS5{F0: ((70 <= 42) && (false || (!((!true) || false)))), F1: 1000}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
