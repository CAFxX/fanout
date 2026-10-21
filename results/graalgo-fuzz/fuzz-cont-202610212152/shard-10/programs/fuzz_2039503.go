package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int64 }
type GenS2 struct { F0 int64; F1 int32; F2 int }
func genFn0(p0 uint, p1 int32) int16 {
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := []int{fi7, (fi7 - (fi8 % (((fi8 ^ fi7)) | 1)))}
            var v10 int64 = 1
            switch 9223372036854775807 {
            case 0:
                v10 += 0
            case 4294967295:
                _ = len(v9)
            default:
                v9 = append(v9, (fi7 ^ int(p0)))
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    if (42 > 42) {
        v11 := GenS1{F0: 2, F1: 812, F2: 876}
        if false {
            fmt.Println(p0, p1)
            switch "" {
            case "foo bar":
                for fi12 := 0; fi12 < 10; fi12++ {
                    var v13 string = "a"
                    v14 := []int{fi12, 4294967295, 42}
                    var v15 int = -517
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
            }
            switch ("a" + "foo bar") {
            case "":
                p0 = 0
            default:
                if true {
                    for fi16 := 0; fi16 < 9; fi16++ {
                        var v17 int = (fi16 & -1)
                        var v18 uint16 = 65535
                        var v19 uint64 = 306
                        _ = fi16
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    p0 = p0
                } else {
                    switch "" {
                    case "a":
                        v20 := []int{}
                        _ = v20
                    case "":
                        for fi21 := 0; fi21 < 19; fi21++ {
                            v22 := GenS1{F0: 128, F1: fi21, F2: 42}
                            for fi23 := 0; fi23 < 19; fi23++ {
                                v24 := GenS2{F0: 1, F1: (-(p1)), F2: (fi21 ^ fi23)}
                                _ = fi23
                                _ = v24
                            }
                            _ = fi21
                            _ = v22
                        }
                    case "\n\t":
                        switch 4294967295 {
                        case -9223372036854775808:
                            p1 += 1
                        }
                    }
                    v25 := []int{674, 42, -2147483648, 9223372036854775807}
                    var v26 int8 = 100
                    _ = v25
                    _ = v26
                }
            }
        }
        _ = v11
    } else {
        p0 -= (p0 / (((p0 * (p0 | p0))) | 1))
        v27 := GenS1{F0: 61, F1: 469, F2: 0}
        _ = v27
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := []uint32{2147483647, 932}
    v2 := []string{"\n\t", "", ""}
    var v3 bool = (1 <= 4294967295)
    var v4 uint32 = 899
    var v5 int8 = 100
    var v6 uint32 = 33
    v28 := GenS2{F0: 0, F1: 0, F2: 9223372036854775807}
    v29 := GenS1{F0: 229, F1: -1, F2: -1}
    v6 = v4
    var v30 uint32 = (^(v4))
    var v31 int8 = int8(v4)
    for fi32 := 0; fi32 < 5; fi32++ {
        v5 -= v5
        v6 = v4
        _ = fi32
    }
    v33 := GenS2{F0: -845, F1: 1, F2: 4294967295}
    for fi34 := 0; fi34 < 2; fi34++ {
        if (1 > -916) {
            if len(v2) > 0 {
                v2[0] = "hello"
            }
        }
        for fi35 := 0; fi35 < 12; fi35++ {
            fmt.Println(v30, v6, v31)
            var v36 uint64 = 1
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    var v37 uint32 = v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
