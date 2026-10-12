package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int32; F1 uint; F2 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 uint64; F2 int16 }
func genFn0(p0 uint8, p1 string, p2 uint64) uint64 {
    var v7 uint8 = p0
    if true {
        v8 := []int{}
        var v9 uint = 412
        var v10 uint16 = 710
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        switch (("foo bar" + "") + "") {
        case "foo bar":
            if (!false) {
                for fi12 := 0; fi12 < 18; fi12++ {
                    v7 -= p0
                    _ = fi12
                }
            } else {
                if (((127 != -120) && (!false)) || true) {
                    var v13 int64 = -1
                    if true {
                        var v14 uint = 959
                        _ = v14
                    } else {
                        for fi15 := 0; fi15 < 17; fi15++ {
                            var v16 uint16 = 555
                            if true {
                                if true {
                                    var v17 uint8 = (^(p0))
                                    var v18 float32 = 0.1
                                    p2 = 0
                                    v19 := []int{fi15, (fi11 | fi11), (fi15 / (((fi15 | fi11)) | 1)), (fi11 % (((fi15 * (fi15 ^ 914))) | 1))}
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                }
                                if false {
                                    fmt.Println(v16, v16, fi11)
                                    fmt.Println(fi15, v13)
                                }
                                v16 += 497
                            } else {
                                fmt.Println(v7)
                                p1 = p1
                                v20 := GenS1{F0: 399, F1: (false && true)}
                                fmt.Println(v13, p1, p2)
                                _ = v20
                            }
                            var v21 int8 = 42
                            _ = fi15
                            _ = v16
                            _ = v21
                        }
                    }
                    switch fi11 {
                    case 32768:
                        v22 := GenS2{F0: 2147483647, F1: 1, F2: true}
                        _ = v22
                    case 1:
                        v23 := []int{(fi11 * 32768)}
                        _ = v23
                    case -2147483648:
                        for fi24 := 0; fi24 < 2; fi24++ {
                            v25 := []int{(fi11 % ((127) | 1)), (fi24 / (((^(fi11))) | 1))}
                            _ = fi24
                            _ = v25
                        }
                    }
                    _ = v13
                } else {
                    fmt.Println(p0)
                    var v26 uint16 = uint16(p0)
                    var v27 int = fi11
                    _ = v26
                    _ = v27
                }
            }
        }
        _ = fi11
    }
    var v28 bool = false
    for fi29 := 0; fi29 < 15; fi29++ {
        var v30 string = ((p1 + "\n\t") + p1)
        _ = fi29
        _ = v30
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    _ = v7
    return (^(p2))
}
func genFn1(p0 uint32, p1 int8) uint32 {
    for fi31 := 0; fi31 < 15; fi31++ {
        var v32 int32 = 0
        var v33 uint64 = 4294967295
        _ = fi31
        _ = v32
        _ = v33
    }
    v34 := GenS1{F0: uint16(p0), F1: false}
    fmt.Println(p0, p0)
    switch -986 {
    case -1:
        for fi35 := 0; fi35 < 7; fi35++ {
            fmt.Println(fi35, p1, p0)
            _ = fi35
        }
    case 128:
        p0 = p0
    case -117:
        if (2 >= 1) {
            var v36 int = -9223372036854775808
            p1 += 123
            _ = v36
        }
    }
    _ = p0
    _ = p1
    _ = v34
    return p0
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{"\n\t", (("hello" + "hello") + "\n\t"), "\n\t"}
    v3 := make(map[int]bool)
    var v4 int16 = -939
    var v5 int64 = 0
    var v6 bool = false
    switch 2147483647 {
    case 32767:
        for fi37 := 0; fi37 < 18; fi37++ {
            fmt.Println(v5, v4, v5)
            if v6 {
                var v38 float64 = ((3.14159 - (-2.25 + (float64(0.0) - 0.0))) + 0.1)
                if v6 {
                    var v39 int64 = v5
                    _ = v39
                }
                _ = v38
            }
            var v40 bool = v6
            _ = fi37
            _ = v40
        }
    case 2147483647:
        var v41 uint64 = 4294967295
        _ = v41
    case 343:
        fmt.Println(v5)
    }
    v42 := GenS3{F0: false}
    var v43 float64 = -2.25
    for fi44 := 0; fi44 < 11; fi44++ {
        v1[("\n\t" + "a")] = "\n\t"
        v6 = ((fi44 ^ (fi44 << ((2) & 63))) <= 42)
        var v45 string = "hello"
        _ = fi44
        _ = v45
    }
    for fi46 := 0; fi46 < 4; fi46++ {
        fmt.Println(v4)
        v47 := GenS1{F0: 107, F1: v6}
        if true {
            if v6 {
                v2 = append(v2, "\n\t")
                v43 = float64((-2.25 - 0.1))
            }
            v48 := GenS2{F0: -717, F1: 889, F2: v6}
            _ = v48
        } else {
            v4 += (v4 % ((v4) | 1))
        }
        _ = fi46
        _ = v47
    }
    var v49 uint16 = 65535
    v6 = false
    for fi50 := 0; fi50 < 19; fi50++ {
        v4 = (v4 + (v4 & (-(v4))))
        _ = fi50
    }
    switch "foo bar" {
    case "":
        v49 -= 65535
    case "a":
        for fi51 := 0; fi51 < 6; fi51++ {
            for fi52 := 0; fi52 < 2; fi52++ {
                v5 -= v5
                _ = fi52
            }
            _ = fi51
        }
    default:
        var v53 int32 = 0
        _ = v53
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v54 := GenS4{F0: v5}
    var v55 uint8 = 128
    if v6 {
        var v56 uint32 = genFn1(1, 0)
        switch -689 {
        case -519:
            _ = len(v2)
        case -561:
            if (false && v6) {
                v1["a"] = ("foo bar" + "hello")
            } else {
                fmt.Println(v4, v6)
            }
        case -9223372036854775808:
            var v57 uint64 = 509
            _ = v57
        default:
            for fi58 := 0; fi58 < 20; fi58++ {
                v43 += (1e10 / 1.0)
                _ = fi58
            }
        }
        fmt.Println(v4)
        _ = v56
    }
    switch "foo bar" {
    case "hello":
        for fi59 := 0; fi59 < 5; fi59++ {
            var v60 float32 = 0.1
            if (v55 > v55) {
                for fi61 := 0; fi61 < 5; fi61++ {
                    fmt.Println(v4, fi61, v49)
                    v62 := GenS5{F0: 1, F1: 2, F2: (v4 << ((2) & 63))}
                    _ = fi61
                    _ = v62
                }
                v2 = append(v2, ("hello" + "foo bar"))
                var v63 int32 = 2147483647
                _ = v63
            }
            var v64 int32 = 1
            _ = fi59
            _ = v60
            _ = v64
        }
    case "a":
        for fi65 := 0; fi65 < 3; fi65++ {
            fmt.Println(v4, v4)
            _ = fi65
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v49
    _ = v5
    _ = v54
    _ = v55
    _ = v6
}
