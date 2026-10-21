package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 string; F1 uint8; F2 uint64 }
func genFn0(p0 int64, p1 uint32) uint64 {
    v4 := GenS1{F0: 158}
    p0 = 607
    switch "a" {
    case "":
        p0 = -549
    case "foo bar":
        p0 = -9223372036854775808
    case "\n\t":
        var v5 uint8 = uint8(p1)
        _ = v5
    default:
        var v6 uint64 = 4294967295
        _ = v6
    }
    v7 := GenS2{F0: 1}
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    return 175
}

func main() {
    v1 := []string{"a"}
    var v2 uint32 = 2147483648
    v3 := make(map[int]bool)
    v1 = append(v1, ("hello" + "\n\t"))
    fmt.Println(v2)
    v8 := GenS3{F0: 0, F1: 224}
    if true {
        var v9 uint8 = 42
        _ = v9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 int32 = 0
        v2 -= (^(v2))
        _ = fi10
        _ = v11
    }
    switch "foo bar" {
    case "foo bar":
        if true {
            if true {
                var v12 int = 361
                _ = v12
            }
            fmt.Println(v2)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            if (false && true) {
                var v13 int32 = 1
                if len(v1) > 0 {
                    v1[0] = ("\n\t" + (("hello" + "foo bar") + ""))
                }
                v2 -= (v2 + v2)
                _ = v13
            } else {
                _ = len(v1)
            }
            v14 := GenS3{F0: 2147483647, F1: 4294967295}
            _ = v14
        }
    case "":
        v2 = v2
    default:
        switch "hello" {
        case "\n\t":
            switch 1 {
            case 9223372036854775807:
                var v15 int32 = 1
                _ = v15
            case 333:
                v16 := GenS1{F0: 12}
                _ = v16
            case -749:
                if false {
                    v2 = 0
                    var v17 uint16 = 0
                    _ = v17
                } else {
                    var v18 float32 = 1e10
                    _ = v18
                }
            }
        case "":
            switch -325 {
            case -9223372036854775808:
                if false {
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                } else {
                    v2 -= (v2 % (((v2 ^ v2)) | 1))
                }
            case 2147483647:
                for fi19 := 0; fi19 < 19; fi19++ {
                    var v20 uint = 989
                    if true {
                        var v21 bool = false
                        var v22 uint32 = (v2 ^ (^(v2)))
                        _ = v21
                        _ = v22
                    }
                    var v23 int32 = -1
                    _ = fi19
                    _ = v20
                    _ = v23
                }
            case 303:
                fmt.Println(v2)
            default:
                var v24 bool = true
                _ = v24
            }
        case "a":
            fmt.Println(v2)
        default:
            v25 := GenS1{F0: uint8(v2)}
            _ = v25
        }
    }
    var v26 int64 = 42
    v1 = append(v1, "a")
    fmt.Println(v26, v2)
    var v27 float32 = 0.0
    var v28 int = -572
    if true {
        _ = len(v3)
    } else {
        for fi29 := 0; fi29 < 18; fi29++ {
            v28 += v28
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v30 uint32 = 42
            _ = fi29
            _ = v30
        }
        if false {
            for fi31 := 0; fi31 < 8; fi31++ {
                v32 := v3[v28]
                switch v28 {
                case 32768:
                    for fi33 := 0; fi33 < 1; fi33++ {
                        switch 302 {
                        case 154:
                            v28 = 2147483647
                        case 9223372036854775807:
                            v27 += v27
                        case -2147483648:
                            switch "\n\t" {
                            case "foo bar":
                                if ((2 >= 18446744073709551615) || true) {
                                    v32 = (!("a" >= ""))
                                }
                            case "a":
                                v34 := GenS4{F0: "a", F1: 0, F2: 216}
                                _ = v34
                            case "hello":
                                var v35 uint8 = 0
                                _ = v35
                            }
                        }
                        if (float64((1e10 / v27)) < (1e10 - 3.14159)) {
                            v36 := GenS3{F0: 262, F1: 18446744073709551615}
                            _ = v36
                        } else {
                            var v37 int64 = v26
                            _ = v37
                        }
                        _ = fi33
                    }
                case 9223372036854775807:
                    switch (fi31 * 9223372036854775807) {
                    case -9223372036854775808:
                        v26 = v26
                    default:
                        v26 -= 42
                    }
                case 42:
                    fmt.Println(fi31)
                default:
                    if false {
                        var v38 uint64 = 18446744073709551615
                        _ = v38
                    }
                }
                fmt.Println(v26, fi31)
                _ = fi31
                _ = v32
            }
            if (42 > 128) {
                for fi39 := 0; fi39 < 2; fi39++ {
                    v2 = (^(v2))
                    _ = fi39
                }
                v26 += 42
                var v40 float32 = v27
                _ = v40
            } else {
                for fi41 := 0; fi41 < 17; fi41++ {
                    switch "" {
                    case "\n\t":
                        var v42 int64 = (v26 | (v26 & (v26 << ((39) & 63))))
                        _ = v42
                    case "foo bar":
                        v2 = (v2 ^ (^(v2)))
                    case "a":
                        switch (fi41 / (((-(fi41))) | 1)) {
                        case 0:
                            var v43 string = "foo bar"
                            _ = v43
                        case -1:
                            v26 = 1
                        case -9223372036854775808:
                            var v44 int64 = (v26 / ((v26) | 1))
                            _ = v44
                        }
                    }
                    if ((-640 >= -32768) && ((127 > 127) && (32768 < fi41))) {
                        for fi45 := 0; fi45 < 6; fi45++ {
                            var v46 uint16 = 168
                            v27 -= (v27 * v27)
                            if (-32768 < 42) {
                                var v47 uint8 = 1
                                if ((0.0 * ((0.1 + 0.0) / (float64((v27 + ((v27 * 0.1) - -2.25))) * -2.25))) < 1.5) {
                                    var v48 int32 = 0
                                    fmt.Println(v47, fi45)
                                    _ = v48
                                } else {
                                    var v49 float64 = 0.1
                                    var v50 uint8 = (v47 >> ((uint(v47)) & 63))
                                    fmt.Println(v50, v2)
                                    _ = v49
                                    _ = v50
                                }
                                fmt.Println(fi45, v28)
                                _ = v47
                            }
                            _ = fi45
                            _ = v46
                        }
                        var v51 uint32 = v2
                        _ = v51
                    } else {
                        fmt.Println(v26)
                        fmt.Println(v26, v28)
                        var v52 uint64 = 4294967295
                        _ = v52
                    }
                    _ = len(v1)
                    _ = fi41
                }
                v28 += 9223372036854775807
                if len(v1) > 0 {
                    v1[0] = "a"
                }
            }
        }
    }
    switch "" {
    case "foo bar":
        for fi53 := 0; fi53 < 18; fi53++ {
            for fi54 := 0; fi54 < 19; fi54++ {
                fmt.Println(v28, fi53)
                fmt.Println(v26)
                _ = fi54
            }
            var v55 uint16 = 1
            _ = fi53
            _ = v55
        }
    case "\n\t":
        v27 = v27
    case "hello":
        var v56 uint8 = 1
        _ = v56
    }
    var v57 float64 = -2.25
    v57 -= (((v57 + (0.1 - v57)) - 3.14159) - 0.1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v57
    _ = v8
}
