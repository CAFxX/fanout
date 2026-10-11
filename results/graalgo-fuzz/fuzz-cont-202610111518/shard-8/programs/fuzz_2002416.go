package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int8 }
type GenS2 struct { F0 int16; F1 string; F2 bool }
type GenS3 struct { F0 uint; F1 uint; F2 int }
type GenS4 struct { F0 int64; F1 uint8; F2 uint8 }
func genFn0(p0 string) uint64 {
    if true {
        p0 = ("" + p0)
        v7 := []int{}
        _ = v7
    }
    v8 := []int{}
    if (!false) {
        p0 = p0
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        switch p0 {
        case "":
            _ = len(v8)
        case "hello":
            v10 := GenS1{F0: 228, F1: 255, F2: 41}
            _ = v10
        default:
            fmt.Println(p0)
        }
        v11 := GenS2{F0: -1, F1: (p0 + p0), F2: true}
        _ = fi9
        _ = v11
    }
    var v12 uint = 18446744073709551615
    _ = p0
    _ = v12
    _ = v8
    return 2
}

func main() {
    var v1 uint16 = 873
    v2 := make(map[uint32]bool)
    var v3 float64 = 3.14159
    var v4 int32 = 42
    var v5 bool = ((v4 % (((v4 >> ((942) & 63))) | 1)) > v4)
    v6 := make(map[int]bool)
    v13 := []int{9223372036854775807}
    if (false && ((v1 + v1) != (v1 & (v1 - v1)))) {
        if len(v13) > 0 {
            v13[0] = 127
        }
        if v5 {
            switch -2147483648 {
            case 32768:
                if (42 > 4294967295) {
                    switch "a" {
                    case "a":
                        var v14 int64 = 0
                        _ = v14
                    }
                }
            case 2147483647:
                switch 590 {
                case 1:
                    v5 = false
                default:
                    for fi15 := 0; fi15 < 16; fi15++ {
                        var v16 uint64 = 755
                        _ = fi15
                        _ = v16
                    }
                }
            }
            var v17 int8 = 79
            _ = v17
        } else {
            var v18 float32 = -2.25
            _ = v18
        }
        var v19 uint64 = 2
        _ = v19
    }
    v4 = (v4 - -1)
    if (!true) {
        if true {
            v20 := GenS1{F0: (v1 | (v1 | v1)), F1: 1, F2: 81}
            fmt.Println(v5)
            fmt.Println(v5, v4, v4)
            _ = v20
        }
    }
    v21 := GenS2{F0: 42, F1: "", F2: v5}
    for fi22 := 0; fi22 < 19; fi22++ {
        if v5 {
            if len(v13) > 0 {
                v13[0] = (-(fi22))
            }
        }
        for fi23 := 0; fi23 < 5; fi23++ {
            v24 := GenS2{F0: -528, F1: ("\n\t" + "foo bar"), F2: v5}
            var v25 int8 = -39
            _ = fi23
            _ = v24
            _ = v25
        }
        if ((false && (-770 <= -9223372036854775808)) || (v5 && (!true))) {
            v26 := GenS2{F0: -821, F1: "foo bar", F2: (65 == v1)}
            if (505 > 326) {
                if true {
                    if true {
                        var v27 uint = 65535
                        v5 = v5
                        _ = v27
                    } else {
                        switch fi22 {
                        case -599:
                            v3 = v3
                        default:
                            v3 = v3
                        }
                    }
                    v4 -= 600
                }
            } else {
                v1 -= v1
                v28 := GenS3{F0: 18446744073709551615, F1: 317, F2: (fi22 << ((4294967295) & 63))}
                _ = v28
            }
            _ = v26
        }
        _ = fi22
    }
    v3 = float64(1.5)
    if ((1 < 2) && v5) {
        v5 = v5
        switch ("" + "a") {
        case "hello":
            v4 += 115
        case "":
            for fi29 := 0; fi29 < 2; fi29++ {
                _ = len(v13)
                _ = fi29
            }
        case "a":
            var v30 bool = ((!v5) || false)
            _ = v30
        }
    } else {
        switch "\n\t" {
        case "":
            if false {
                var v31 uint16 = (^(v1))
                _ = len(v2)
                var v32 bool = (v5 && (v31 < v1))
                _ = v31
                _ = v32
            }
        case "hello":
            v1 -= 795
        case "a":
            if (-76 == 9223372036854775807) {
                delete(v2, 456)
                fmt.Println(v4, v4)
            } else {
                if ("\n\t" < "hello") {
                    v33 := GenS3{F0: 617, F1: 65535, F2: 9223372036854775807}
                    v2[2147483648] = (!false)
                    _ = v33
                } else {
                    v4 = 0
                    if false {
                        var v34 float64 = 1e10
                        if len(v13) > 0 {
                            v13[0] = 1
                        }
                        _ = v34
                    }
                }
            }
        default:
            v1 = v1
        }
        var v35 float64 = v3
        _ = v35
    }
    v36 := GenS4{F0: 0, F1: 182, F2: 0}
    var v37 uint = 255
    switch 127 {
    case 4294967295:
        var v38 int = -1
        _ = v38
    default:
        var v39 int = 32768
        _ = v39
    }
    var v40 int16 = 803
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
