package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint8; F1 int8; F2 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 int8; F1 uint32 }
func genFn0(p0 int8, p1 uint8, p2 int) uint16 {
    fmt.Println(p0, p1)
    v4 := GenS1{F0: 1, F1: 255, F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 1
}
func genFn1(p0 int8, p1 string, p2 uint8) string {
    fmt.Println(p0, p2, p2)
    for fi5 := 0; fi5 < 12; fi5++ {
        fmt.Println(p2, p1)
        var v6 uint16 = 65535
        _ = fi5
        _ = v6
    }
    v7 := []int{-2147483648, -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p1 + "\n\t")
}

func main() {
    v1 := []int{-2147483648, 9223372036854775807}
    var v2 uint64 = 985
    v3 := make(map[uint32]bool)
    var v8 int32 = -925
    var v9 uint32 = 65535
    _ = len(v1)
    v3[v9] = (true || true)
    if true {
        v10 := GenS2{F0: 127, F1: -22, F2: v8}
        switch 32767 {
        case -703:
            var v11 uint8 = 172
            _ = v11
        case 2147483647:
            v9 -= (v9 | v9)
        default:
            v1 = append(v1, 1)
        }
        _ = v10
    }
    var v12 uint8 = 0
    v12 = v12
    if true {
        if false {
            var v13 float64 = 1e10
            if true {
                for fi14 := 0; fi14 < 2; fi14++ {
                    v13 = ((-2.25 + 0.1) + (0.1 * (1.5 - 0.0)))
                    _ = fi14
                }
                var v15 int16 = 1
                _ = v15
            }
            v16 := v3[(v9 / ((4294967295) | 1))]
            _ = v13
            _ = v16
        }
        for fi17 := 0; fi17 < 4; fi17++ {
            if (int(v2) > 2147483647) {
                if (!false) {
                    v18 := GenS3{F0: -1}
                    var v19 int32 = (v8 * (v8 & (v8 | v8)))
                    switch (fi17 / (((fi17 & fi17)) | 1)) {
                    case 750:
                        switch (fi17 - (-(fi17))) {
                        case 4294967295:
                            var v20 int8 = -1
                            _ = v20
                        case -534:
                            for fi21 := 0; fi21 < 18; fi21++ {
                                switch "a" {
                                case "\n\t":
                                    if false {
                                        v3[v9] = (true || false)
                                    } else {
                                        _ = len(v3)
                                        v22 := GenS4{F0: 4294967295, F1: -1}
                                        _ = v22
                                    }
                                case "foo bar":
                                    v3[1] = (661 < 2)
                                }
                                _ = fi21
                            }
                        default:
                            v23 := GenS1{F0: 1, F1: 18446744073709551615, F2: v9}
                            _ = v23
                        }
                    }
                    _ = v18
                    _ = v19
                }
                v9 = (^(v9))
            }
            for fi24 := 0; fi24 < 11; fi24++ {
                v12 = uint8(v9)
                v9 += v9
                switch "foo bar" {
                case "":
                    v25 := GenS5{F0: 0, F1: v9}
                    _ = v25
                }
                _ = fi24
            }
            var v26 int16 = 1000
            _ = fi17
            _ = v26
        }
    } else {
        v27 := GenS2{F0: 84, F1: -106, F2: (-(v8))}
        v28 := GenS5{F0: -30, F1: 0}
        var v29 int32 = 639
        _ = v27
        _ = v28
        _ = v29
    }
    v9 = (v9 | v9)
    var v30 uint32 = v9
    switch 4294967295 {
    case 1:
        v1 = append(v1, 0)
    case 633:
        var v31 uint32 = (^(v9))
        _ = v31
    }
    v2 = v2
    for fi32 := 0; fi32 < 9; fi32++ {
        v9 = (v30 & 4294967295)
        var v33 bool = true
        _ = fi32
        _ = v33
    }
    if (true && true) {
        _ = len(v1)
        var v34 bool = (!(!(!false)))
        var v35 string = "\n\t"
        _ = v34
        _ = v35
    }
    if len(v1) > 0 {
        v1[0] = 860
    }
    v30 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v8
    _ = v9
}
