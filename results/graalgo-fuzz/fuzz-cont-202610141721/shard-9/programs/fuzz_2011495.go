package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32; F1 uint32; F2 int64 }
type GenS6 struct { F0 uint64; F1 uint32; F2 int64 }
func genFn0(p0 int8) string {
    v4 := []int{1, -327, 9223372036854775807, 1}
    v5 := GenS1{F0: 2147483647}
    if len(v4) > 0 {
        v4[0] = -1
    }
    _ = p0
    _ = v4
    _ = v5
    return "foo bar"
}
func genFn1(p0 uint32, p1 int) int8 {
    if (("hello" + "\n\t") < ((("" + ("a" + "\n\t")) + (("" + "foo bar") + ("hello" + ""))) + "hello")) {
        p0 += 1
        var v6 float32 = (float32(((0.1 + 3.14159) / 1e10)) + float32(1.5))
        p0 += p0
        _ = v6
    } else {
        for fi7 := 0; fi7 < 10; fi7++ {
            v8 := []int{(fi7 - (p1 | (p1 ^ fi7))), 2147483647, 127}
            _ = fi7
            _ = v8
        }
        if ((p1 >= p1) && false) {
            p1 = (-(p1))
            var v9 uint16 = 1
            switch 9223372036854775807 {
            case -1:
                if (v9 != (v9 << ((42) & 63))) {
                    if (!true) {
                        var v10 int8 = -55
                        var v11 int64 = 9223372036854775807
                        _ = v10
                        _ = v11
                    }
                }
            case 2147483647:
                fmt.Println(p1, p0, p0)
            }
            _ = v9
        } else {
            var v12 int32 = 754
            if (false && true) {
                p1 = (p1 * (p1 & 0))
                var v13 int = p1
                _ = v13
            }
            v14 := GenS1{F0: (v12 * 42)}
            _ = v12
            _ = v14
        }
        v15 := GenS2{F0: 896}
        _ = v15
    }
    switch p1 {
    case -2147483648:
        var v16 int32 = int32(p1)
        _ = v16
    case 435:
        var v17 float32 = -2.25
        _ = v17
    case 475:
        var v18 string = ("a" + "a")
        _ = v18
    default:
        switch "\n\t" {
        case "hello":
            var v19 float64 = 0.0
            _ = v19
        case "":
            p0 -= 2147483648
        case "foo bar":
            v20 := GenS2{F0: uint64(p0)}
            _ = v20
        default:
            for fi21 := 0; fi21 < 7; fi21++ {
                var v22 int = p1
                _ = fi21
                _ = v22
            }
        }
    }
    v23 := []int{9223372036854775807, 1}
    for fi24 := 0; fi24 < 5; fi24++ {
        for fi25 := 0; fi25 < 17; fi25++ {
            v26 := GenS1{F0: -1}
            var v27 uint64 = 1
            p1 -= fi25
            v28 := GenS3{F0: 157, F1: 32767}
            p1 = fi24
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
        }
        fmt.Println(p0)
        for fi29 := 0; fi29 < 20; fi29++ {
            p0 = p0
            p1 += p1
            var v30 uint64 = 18446744073709551615
            _ = fi29
            _ = v30
        }
        _ = fi24
    }
    v31 := GenS2{F0: 1}
    _ = p0
    _ = p1
    _ = v23
    _ = v31
    return -1
}

func main() {
    v1 := []uint32{65535, 387}
    var v2 float32 = (3.14159 * float32(1.5))
    v3 := make(map[uint32]bool)
    var v32 bool = (0.0 <= 3.14159)
    for fi33 := 0; fi33 < 17; fi33++ {
        var v34 uint32 = 42
        _ = fi33
        _ = v34
    }
    fmt.Println(v2)
    switch -297 {
    case 128:
        switch 32767 {
        case 0:
            v32 = (v32 || (683 > 1))
        default:
            var v35 int = 139
            _ = v35
        }
    case -862:
        v36 := GenS4{F0: 1}
        _ = v36
    case 1:
        fmt.Println(v2, v32)
    default:
        if (!(2147483647 == -1)) {
            for fi37 := 0; fi37 < 19; fi37++ {
                for fi38 := 0; fi38 < 18; fi38++ {
                    v2 += float32(1e10)
                    fmt.Println(fi37, fi37, fi38)
                    if (!false) {
                        v39 := v3[149]
                        _ = v39
                    } else {
                        var v40 int = fi37
                        if v32 {
                            v32 = v32
                        } else {
                            v3[827] = v32
                            v41 := GenS2{F0: 2}
                            var v42 int16 = -1
                            _ = v41
                            _ = v42
                        }
                        _ = v40
                    }
                    _ = fi38
                }
                switch "a" {
                case "a":
                    delete(v3, 65535)
                case "\n\t":
                    var v43 string = ""
                    _ = v43
                case "foo bar":
                    v1 = append(v1, 2147483648)
                default:
                    v44 := GenS3{F0: 464, F1: 42}
                    _ = v44
                }
                _ = fi37
            }
            var v45 bool = true
            _ = v45
        }
    }
    var v46 bool = v32
    var v47 int64 = -9223372036854775808
    fmt.Println(v46)
    v48 := GenS5{F0: 42, F1: 4294967295, F2: (v47 >> ((2) & 63))}
    v49 := GenS6{F0: 18446744073709551615, F1: 2147483647, F2: 9223372036854775807}
    v50 := GenS3{F0: v47, F1: 835}
    switch "\n\t" {
    case "\n\t":
        if false {
            var v51 int = -1
            _ = v51
        }
    case "foo bar":
        var v52 uint = 827
        _ = v52
    }
    var v53 int = 9223372036854775807
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v50
    _ = v53
}
