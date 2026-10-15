package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 int64 }
func genFn0() int64 {
    v4 := GenS1{F0: 65535}
    v5 := GenS2{F0: 314, F1: 32767, F2: 32767}
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := []int{fi6}
        if false {
            fmt.Println(fi6)
            if len(v7) > 0 {
                v7[0] = 1
            }
            var v8 uint64 = 564
            _ = v8
        }
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 float32 = 0.1
            v7 = append(v7, 9223372036854775807)
            _ = fi9
            _ = v10
        }
        _ = fi6
        _ = v7
    }
    if true {
        for fi11 := 0; fi11 < 1; fi11++ {
            fmt.Println(fi11)
            fmt.Println(fi11)
            fmt.Println(fi11)
            _ = fi11
        }
    } else {
        var v12 int = -2147483648
        _ = v12
    }
    switch 127 {
    case 9223372036854775807:
        v13 := []int{-2147483648}
        _ = v13
    case -103:
        var v14 int16 = 42
        _ = v14
    case -1:
        var v15 uint32 = 886
        _ = v15
    default:
        v16 := GenS3{F0: "foo bar", F1: 0}
        _ = v16
    }
    _ = v4
    _ = v5
    return 0
}

func main() {
    var v1 uint64 = 912
    var v2 uint32 = 2
    var v3 float64 = 1.5
    fmt.Println(v1, v3, v2)
    for fi17 := 0; fi17 < 8; fi17++ {
        switch "hello" {
        case "foo bar":
            v1 = v1
        default:
            v18 := []int{(-(fi17))}
            _ = v18
        }
        var v19 bool = true
        var v20 int16 = 1000
        _ = fi17
        _ = v19
        _ = v20
    }
    fmt.Println(v3, v3)
    v21 := GenS4{F0: 1, F1: 65535}
    switch 127 {
    case 9223372036854775807:
        if (!(-9223372036854775808 == genFn0())) {
            for fi22 := 0; fi22 < 6; fi22++ {
                v23 := GenS1{F0: 4294967295}
                _ = fi22
                _ = v23
            }
        }
    case -663:
        for fi24 := 0; fi24 < 6; fi24++ {
            if true {
                v25 := []int{}
                _ = v25
            } else {
                v26 := []int{(fi24 ^ 127)}
                v27 := GenS5{F0: (fi24 << ((569) & 63)), F1: 1}
                fmt.Println(v2)
                _ = v26
                _ = v27
            }
            if true {
                v28 := GenS6{F0: 485}
                _ = v28
            }
            for fi29 := 0; fi29 < 11; fi29++ {
                if false {
                    v30 := GenS1{F0: (v2 << ((525) & 63))}
                    if (!(!true)) {
                        if (!(898 < -805)) {
                            var v31 float64 = (v3 + (v3 / (3.14159 - ((v3 + (1e10 - 3.14159)) - (1e10 + v3)))))
                            _ = v31
                        } else {
                            v3 = v3
                            v1 += uint64(v2)
                        }
                    } else {
                        var v32 bool = (!false)
                        _ = v32
                    }
                    _ = v30
                }
                _ = fi29
            }
            _ = fi24
        }
    default:
        for fi33 := 0; fi33 < 2; fi33++ {
            v1 += (^(v1))
            fmt.Println(fi33, v3)
            _ = fi33
        }
    }
    v34 := []int{}
    for fi35 := 0; fi35 < 8; fi35++ {
        fmt.Println(fi35, v1, v1)
        var v36 uint = 176
        v3 += v3
        _ = fi35
        _ = v36
    }
    if (42 >= 255) {
        v34 = append(v34, 654)
        var v37 uint = 196
        fmt.Println(v3, v3, v37)
        _ = v37
    } else {
        for fi38 := 0; fi38 < 13; fi38++ {
            if true {
                fmt.Println(v2)
                if (!(int64(v1) >= -9223372036854775808)) {
                    if (997 >= 711) {
                        v39 := GenS6{F0: 1}
                        _ = v39
                    }
                    v2 = (v2 / ((v2) | 1))
                } else {
                    v1 += (^(v1))
                    for fi40 := 0; fi40 < 6; fi40++ {
                        var v41 int64 = int64(v1)
                        _ = fi40
                        _ = v41
                    }
                }
            }
            var v42 uint64 = (v1 + 351)
            fmt.Println(v42, v3, v2)
            _ = fi38
            _ = v42
        }
    }
    v43 := GenS6{F0: 42}
    var v44 uint16 = 1
    v2 = v2
    switch -1 {
    case 128:
        v2 -= v2
    }
    for fi45 := 0; fi45 < 6; fi45++ {
        v3 += (v3 * 0.1)
        _ = fi45
    }
    v3 = v3
    v1 = (v1 - 535)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v34
    _ = v43
    _ = v44
}
