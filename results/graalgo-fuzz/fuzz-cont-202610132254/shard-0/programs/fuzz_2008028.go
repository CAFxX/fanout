package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
func genFn0(p0 int8) bool {
    switch 0 {
    case 35:
        v6 := []int{-9223372036854775808, -9223372036854775808, int(p0)}
        _ = v6
    case -259:
        switch 1 {
        case -909:
            v7 := []int{-1, -579, -705}
            _ = v7
        }
    default:
        v8 := []int{1, 1, -2147483648, -339}
        _ = v8
    }
    switch ((("a" + "hello") + "a") + "\n\t") {
    case "a":
        v9 := GenS1{F0: -840, F1: -946}
        _ = v9
    case "foo bar":
        for fi10 := 0; fi10 < 14; fi10++ {
            fmt.Println(p0)
            p0 = p0
            var v11 uint64 = 821
            _ = fi10
            _ = v11
        }
    }
    v12 := GenS1{F0: 4294967295, F1: 128}
    _ = p0
    _ = v12
    return false
}
func genFn1() int32 {
    var v13 uint16 = 1
    v14 := []int{0, 1, -683}
    var v15 uint8 = 191
    if (!(!false)) {
        v15 += (v15 / ((42) | 1))
        if (4294967295 > 32767) {
            v14 = append(v14, -857)
        } else {
            for fi16 := 0; fi16 < 1; fi16++ {
                for fi17 := 0; fi17 < 12; fi17++ {
                    fmt.Println(fi17, fi17)
                    _ = fi17
                }
                var v18 float64 = (float64(3.14159) + 0.0)
                _ = fi16
                _ = v18
            }
            switch -2147483648 {
            case 4294967295:
                v13 += v13
            }
        }
    } else {
        v13 += (v13 ^ v13)
        fmt.Println(v15, v15)
    }
    var v19 int64 = 42
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"\n\t", "a", "a"}
    var v3 int = 32767
    var v4 int16 = 32767
    v5 := []int{v3}
    var v20 uint64 = 4294967295
    var v21 int32 = 0
    v22 := GenS1{F0: v3, F1: (-(v3))}
    fmt.Println(v20)
    for fi23 := 0; fi23 < 12; fi23++ {
        for fi24 := 0; fi24 < 3; fi24++ {
            var v25 float64 = -2.25
            _ = fi24
            _ = v25
        }
        v20 -= (v20 ^ (v20 / ((v20) | 1)))
        v3 += (v3 + 9223372036854775807)
        _ = fi23
    }
    var v26 uint32 = 65535
    v27 := v1[4294967295]
    if v27 {
        v21 = -377
        if v27 {
            if len(v5) > 0 {
                v5[0] = v3
            }
            v27 = false
            v28 := GenS1{F0: (v3 ^ (v3 | v3)), F1: v3}
            _ = v28
        }
    } else {
        v3 -= 221
        switch v3 {
        case 127:
            var v29 float64 = 3.14159
            _ = v29
        case 1:
            v30 := GenS1{F0: 128, F1: -9223372036854775808}
            _ = v30
        default:
            v26 -= 65535
        }
        var v31 int8 = 100
        _ = v31
    }
    fmt.Println(v27, v26, v4)
    if (2 == 100) {
        for fi32 := 0; fi32 < 4; fi32++ {
            if true {
                var v33 string = "foo bar"
                switch -9223372036854775808 {
                case -345:
                    if len(v5) > 0 {
                        v5[0] = 417
                    }
                default:
                    if (!v27) {
                        v34 := GenS1{F0: (fi32 * 313), F1: (v3 - (v3 + 32768))}
                        _ = v34
                    }
                }
                for fi35 := 0; fi35 < 2; fi35++ {
                    _ = len(v5)
                    fmt.Println(fi32, v21)
                    _ = fi35
                }
                _ = v33
            }
            _ = fi32
        }
        fmt.Println(v21)
    }
    v36 := GenS1{F0: (v3 * (-(v3))), F1: v3}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
