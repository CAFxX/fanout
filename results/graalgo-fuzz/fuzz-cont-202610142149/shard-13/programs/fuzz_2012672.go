package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS7 struct { F0 int32 }
func genFn0(p0 int16, p1 uint, p2 int8) int64 {
    var v5 uint = 555
    switch "" {
    case "\n\t":
        v6 := []int{-9223372036854775808, 2147483647, -629}
        _ = v6
    case "a":
        v7 := []int{32768}
        _ = v7
    case "foo bar":
        if true {
            var v8 int32 = 2147483647
            _ = v8
        } else {
            v9 := []int{-2147483648}
            switch 9223372036854775807 {
            case -844:
                var v10 float32 = 0.0
                _ = v10
            case 4294967295:
                v11 := GenS1{F0: 18446744073709551615, F1: 246}
                _ = v11
            case 0:
                v9 = append(v9, 32768)
            default:
                p0 = p0
            }
            _ = v9
        }
    default:
        var v12 int32 = 1
        _ = v12
    }
    if ("a" == ("" + ("\n\t" + "\n\t"))) {
        v13 := GenS2{F0: 801, F1: 18446744073709551615}
        if (32767 > 98) {
            var v14 uint16 = 891
            v14 += v14
            _ = v14
        }
        v15 := GenS3{F0: v5, F1: v5}
        _ = v13
        _ = v15
    } else {
        if (127 != 42) {
            v16 := GenS4{F0: int64(p1), F1: -2147483648}
            v17 := []int{}
            _ = v16
            _ = v17
        }
        v18 := GenS3{F0: p1, F1: (^(p1))}
        _ = v18
    }
    if (!(-696 <= 112)) {
        p2 = p2
        v19 := []int{670, -563}
        _ = v19
    }
    if true {
        var v20 int = -427
        switch v20 {
        case 127:
            v21 := GenS4{F0: -154, F1: 1}
            _ = v21
        case 42:
            fmt.Println(p0, p1, v20)
        case 32768:
            fmt.Println(p1, p0)
        }
        _ = v20
    } else {
        v22 := []int{}
        fmt.Println(v5, v5, v5)
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -1
}

func main() {
    var v1 uint32 = 42
    v2 := []string{"foo bar"}
    v3 := make(map[string]int)
    var v4 float64 = 1.5
    _ = len(v2)
    var v23 int16 = -32768
    for fi24 := 0; fi24 < 20; fi24++ {
        var v25 uint32 = 441
        v26 := GenS5{F0: -254}
        v27 := GenS5{F0: (v23 / (((^(v23))) | 1))}
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    var v28 uint = 1
    var v29 int = 42
    if (!false) {
        v1 += (v1 | 2)
        v23 += v23
    }
    v30 := GenS6{F0: 32767, F1: 0, F2: 127}
    v31 := GenS2{F0: (v23 + v23), F1: 18446744073709551615}
    fmt.Println(v23, v1, v29)
    for fi32 := 0; fi32 < 4; fi32++ {
        var v33 float64 = v4
        var v34 int = fi32
        for fi35 := 0; fi35 < 19; fi35++ {
            switch v29 {
            case -2147483648:
                v3["hello"] = fi32
                v36 := GenS6{F0: 32767, F1: 551, F2: 212}
                if false {
                    if false {
                        if true {
                            v37 := GenS4{F0: int64(v28), F1: -1}
                            _ = v37
                        }
                    } else {
                        v28 += (v28 % ((42) | 1))
                        v38 := GenS5{F0: 566}
                        var v39 int8 = 1
                        _ = v38
                        _ = v39
                    }
                } else {
                    v29 += (fi32 | 2147483647)
                }
                var v40 int16 = v23
                _ = v36
                _ = v40
            case 42:
                var v41 int8 = -1
                v1 -= (v1 % (((v1 & 812)) | 1))
                _ = len(v3)
                v42 := GenS7{F0: 42}
                _ = v41
                _ = v42
            case -9223372036854775808:
                v2 = append(v2, "a")
            default:
                var v43 uint64 = 871
                _ = v43
            }
            v29 = -1
            if false {
                v44 := GenS5{F0: (v23 | v23)}
                switch ("foo bar" + "foo bar") {
                case "foo bar":
                    v4 = float64(1.5)
                }
                _ = v44
            }
            var v45 uint = (v28 % ((v28) | 1))
            _ = len(v3)
            v46 := GenS2{F0: (v23 & v23), F1: 0}
            _ = fi35
            _ = v45
            _ = v46
        }
        _ = fi32
        _ = v33
        _ = v34
    }
    for fi47 := 0; fi47 < 3; fi47++ {
        v48 := GenS7{F0: 42}
        for fi49 := 0; fi49 < 18; fi49++ {
            fmt.Println(v4)
            _ = fi49
        }
        _ = fi47
        _ = v48
    }
    v28 = (^(v28))
    v50 := GenS1{F0: (v28 / (((v28 | 643)) | 1)), F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v50
}
