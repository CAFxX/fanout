package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 int64; F2 uint64 }

func main() {
    v1 := []int{128, 2147483647, -9223372036854775808}
    var v2 uint = 42
    v3 := make(map[int]int)
    v4 := []int{-81, -208}
    v5 := []uint32{}
    var v6 uint = (^(v2))
    switch "\n\t" {
    case "a":
        v2 -= (v2 & (v6 + 42))
    case "foo bar":
        var v7 uint16 = 2
        _ = v7
    case "":
        var v8 int8 = 53
        _ = v8
    }
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    switch 812 {
    case 4294967295:
        switch "foo bar" {
        case "":
            var v9 int8 = 127
            _ = v9
        case "\n\t":
            switch ("" + "a") {
            case "":
                v10 := v3[32767]
                _ = v10
            case "foo bar":
                switch -564 {
                case 32767:
                    for fi11 := 0; fi11 < 12; fi11++ {
                        var v12 int = 127
                        _ = fi11
                        _ = v12
                    }
                case 32768:
                    v2 = v6
                default:
                    v2 += 255
                }
            case "hello":
                switch 405 {
                case 2147483647:
                    switch "hello" {
                    case "hello":
                        var v13 uint64 = 4294967295
                        _ = v13
                    case "\n\t":
                        v2 += v2
                    }
                default:
                    var v14 string = "a"
                    _ = v14
                }
            default:
                if true {
                    var v15 float64 = 0.0
                    _ = v15
                } else {
                    _ = len(v5)
                }
            }
        case "foo bar":
            var v16 float32 = 0.0
            _ = v16
        }
    case 127:
        v17 := GenS1{F0: 42, F1: 4294967295, F2: (^(v2))}
        _ = v17
    default:
        for fi18 := 0; fi18 < 11; fi18++ {
            v19 := GenS1{F0: 339, F1: 42, F2: 42}
            _ = fi18
            _ = v19
        }
    }
    switch "" {
    case "":
        v2 += v6
    }
    var v20 uint = (v6 << ((v2) & 63))
    v2 -= (v2 + 42)
    switch "\n\t" {
    case "foo bar":
        switch ("foo bar" + "\n\t") {
        case "\n\t":
            switch "hello" {
            case "":
                var v21 int32 = 42
                _ = v21
            case "\n\t":
                if (843 == 296) {
                    v6 -= v6
                    v22 := GenS1{F0: -320, F1: 2147483647, F2: v2}
                    _ = v22
                } else {
                    var v23 string = ""
                    _ = v23
                }
            default:
                if (1 >= 2) {
                    for fi24 := 0; fi24 < 2; fi24++ {
                        fmt.Println(fi24, v20)
                        _ = fi24
                    }
                    var v25 string = "hello"
                    v26 := GenS1{F0: 209, F1: 42, F2: v20}
                    _ = v25
                    _ = v26
                }
            }
        case "a":
            fmt.Println(v20, v2)
        }
    case "a":
        v27 := v3[32767]
        _ = v27
    case "hello":
        v6 += (v20 % (((v2 >> ((v2) & 63))) | 1))
    }
    var v28 int = -842
    switch ("" + "\n\t") {
    case "hello":
        var v29 int64 = 9223372036854775807
        _ = v29
    case "foo bar":
        for fi30 := 0; fi30 < 18; fi30++ {
            _ = len(v3)
            for fi31 := 0; fi31 < 10; fi31++ {
                v28 += (fi30 + 165)
                var v32 uint16 = 704
                var v33 uint32 = 1
                _ = fi31
                _ = v32
                _ = v33
            }
            v34 := GenS1{F0: 0, F1: 65535, F2: (v6 << ((1) & 63))}
            _ = fi30
            _ = v34
        }
    case "":
        v35 := GenS1{F0: 724, F1: 42, F2: 65535}
        _ = v35
    default:
        var v36 uint64 = 4294967295
        _ = v36
    }
    if len(v4) > 0 {
        v4[0] = -1
    }
    v5 = append(v5, 1)
    for fi37 := 0; fi37 < 16; fi37++ {
        var v38 bool = false
        if ((537 >= 0) && v38) {
            v39 := v3[v28]
            for fi40 := 0; fi40 < 15; fi40++ {
                v41 := GenS1{F0: 32767, F1: 780, F2: (v20 | v2)}
                v6 += v20
                switch (v39 + fi40) {
                case 0:
                    var v42 uint = v2
                    _ = v42
                }
                _ = fi40
                _ = v41
            }
            var v43 int32 = -2147483648
            _ = v39
            _ = v43
        }
        for fi44 := 0; fi44 < 12; fi44++ {
            if v38 {
                if (1 < 893) {
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v38)
                    case "a":
                        v2 += v6
                    }
                    v45 := GenS2{F0: 809}
                    v46 := GenS1{F0: 32767, F1: 1, F2: 271}
                    _ = v45
                    _ = v46
                } else {
                    var v47 int32 = 42
                    _ = v47
                }
                var v48 int8 = -64
                _ = v48
            } else {
                var v49 uint16 = 552
                var v50 uint16 = v49
                v1 = append(v1, v28)
                if (!false) {
                    v51 := GenS3{F0: 0, F1: 1, F2: 570}
                    _ = v51
                } else {
                    if (!v38) {
                        v20 += 255
                        var v52 uint16 = 2
                        _ = v52
                    }
                    var v53 uint = v20
                    var v54 int64 = 1
                    var v55 int8 = 1
                    _ = v53
                    _ = v54
                    _ = v55
                }
                var v56 float32 = float32(1e10)
                _ = v49
                _ = v50
                _ = v56
            }
            var v57 int16 = 42
            switch (("a" + "foo bar") + ("hello" + ("" + ("hello" + "foo bar")))) {
            case "hello":
                v58 := GenS1{F0: (v57 - (v57 / (((v57 * v57)) | 1))), F1: 2147483648, F2: v2}
                _ = v58
            case "foo bar":
                var v59 uint64 = 18446744073709551615
                _ = v59
            }
            switch (("a" + ("\n\t" + "hello")) + "\n\t") {
            case "a":
                fmt.Println(v2, fi37, fi44)
            }
            var v60 float64 = 1.5
            _ = fi44
            _ = v57
            _ = v60
        }
        _ = fi37
        _ = v38
    }
    if (0 != 1) {
        v3[v28] = v28
    } else {
        for fi61 := 0; fi61 < 8; fi61++ {
            v62 := GenS1{F0: 1000, F1: 527, F2: 65535}
            v28 -= v28
            _ = fi61
            _ = v62
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
