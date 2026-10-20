package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 uint32; F2 uint8 }
func genFn0(p0 int64, p1 uint) int64 {
    v5 := GenS1{F0: 0}
    p0 = (p0 >> ((4294967295) & 63))
    fmt.Println(p0, p0)
    v6 := []int{42, 32768}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (^(p0))
}
func genFn1(p0 int) int16 {
    if true {
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 int64 = 217
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 uint8 = 1
                _ = fi9
                _ = v10
            }
            switch "\n\t" {
            case "":
                fmt.Println(p0)
            case "foo bar":
                v11 := []int{}
                _ = v11
            case "\n\t":
                if false {
                    v12 := []int{fi7, -450, (^(p0))}
                    _ = v12
                }
            default:
                if true {
                    for fi13 := 0; fi13 < 20; fi13++ {
                        p0 = p0
                        var v14 uint8 = 139
                        if (false || (!false)) {
                            var v15 int32 = 320
                            _ = v15
                        } else {
                            v16 := []int{(fi13 & 0), (p0 / (((fi13 / (((fi13 % ((fi7) | 1))) | 1))) | 1)), -16, (fi7 & fi7)}
                            v17 := GenS1{F0: (^(v14))}
                            v14 = (^(v14))
                            _ = v16
                            _ = v17
                        }
                        _ = fi13
                        _ = v14
                    }
                    v8 = (v8 << ((4294967295) & 63))
                    var v18 uint64 = 714
                    _ = v18
                } else {
                    v19 := []int{(p0 << ((18446744073709551615) & 63))}
                    p0 = fi7
                    p0 -= (p0 | 2147483647)
                    _ = v19
                }
            }
            _ = fi7
            _ = v8
        }
        var v20 int64 = -1
        _ = v20
    }
    if ((((!false) || false) || false) && false) {
        v21 := GenS1{F0: 0}
        _ = v21
    }
    p0 = p0
    _ = p0
    return 42
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[int]bool)
    v3 := []string{}
    var v4 int = 318
    v1 -= v1
    switch "" {
    case "foo bar":
        if (857 <= 995) {
            fmt.Println(v4, v1)
            v4 -= (v4 & (v4 & 127))
        }
    case "a":
        v1 = 1e10
    case "\n\t":
        v1 = (0.1 - 0.1)
    }
    v22 := GenS1{F0: 204}
    v1 += (v1 / -2.25)
    v1 = v1
    switch v4 {
    case -2147483648:
        switch "hello" {
        case "foo bar":
            if false {
                v23 := v2[(v4 << ((531) & 63))]
                v4 = 32767
                _ = v23
            } else {
                v4 = (v4 * v4)
            }
        }
    case 32768:
        v24 := GenS1{F0: 128}
        _ = v24
    }
    if (((1.5 - -2.25) + (1e10 - 0.0)) < ((1.5 / 1e10) * float32((v1 - 0.1)))) {
        switch "foo bar" {
        case "foo bar":
            v25 := GenS2{F0: 1}
            _ = v25
        case "a":
            var v26 float64 = (float64(((0.1 * 0.1) * 0.0)) + v1)
            _ = v26
        case "\n\t":
            if len(v3) > 0 {
                v3[0] = ""
            }
        }
        v27 := GenS1{F0: 42}
        fmt.Println(v1, v4)
        _ = v27
    }
    var v28 int16 = 0
    var v29 int8 = 1
    var v30 float64 = ((float64(1.5) / (((v1 + v1) / float64(3.14159)) - v1)) - ((((0.1 - v1) - v1) - 0.0) - v1))
    var v31 uint64 = 898
    if ((114 > v28) || true) {
        v32 := v2[(v4 | (v4 * -168))]
        var v33 uint = 750
        _ = v32
        _ = v33
    } else {
        var v34 string = "a"
        _ = len(v2)
        v35 := GenS3{F0: "", F1: 887, F2: 187}
        _ = v34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
