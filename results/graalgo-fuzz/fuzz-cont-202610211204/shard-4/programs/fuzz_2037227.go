package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint16 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64; F1 string }
func genFn0(p0 bool, p1 int64) int32 {
    if p0 {
        for fi7 := 0; fi7 < 20; fi7++ {
            p1 -= p1
            for fi8 := 0; fi8 < 15; fi8++ {
                fmt.Println(fi8, fi8, p0)
                v9 := GenS1{F0: 65535, F1: -1, F2: 0}
                _ = fi8
                _ = v9
            }
            if true {
                v10 := GenS2{F0: -71, F1: p0}
                v11 := []int{(fi7 + (fi7 / ((32767) | 1))), (fi7 >> ((42) & 63))}
                _ = v10
                _ = v11
            } else {
                p0 = (1.5 <= 0.0)
                if p0 {
                    v12 := []int{-589, (^(fi7))}
                    var v13 uint16 = 242
                    _ = v12
                    _ = v13
                }
                var v14 bool = (862 != 847)
                _ = v14
            }
            _ = fi7
        }
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 int32 = 2147483647
            var v17 uint16 = 65535
            _ = fi15
            _ = v16
            _ = v17
        }
        p0 = (42 >= 0)
    } else {
        var v18 uint64 = 805
        fmt.Println(p0)
        _ = v18
    }
    v19 := GenS3{F0: 2, F1: 9223372036854775807}
    _ = p0
    _ = p1
    _ = v19
    return 572
}
func genFn1(p0 uint, p1 uint32) int32 {
    p1 += 1
    if false {
        if (!false) {
            v20 := []int{1, 970, 332, 32768}
            switch ("\n\t" + "a") {
            case "a":
                fmt.Println(p1, p1)
            case "hello":
                var v21 int8 = 100
                _ = v21
            case "foo bar":
                p0 = p0
            }
            _ = v20
        } else {
            for fi22 := 0; fi22 < 17; fi22++ {
                if false {
                    p0 += (^(p0))
                }
                _ = fi22
            }
            fmt.Println(p1, p1)
            for fi23 := 0; fi23 < 11; fi23++ {
                p0 -= 806
                switch int(p0) {
                case 127:
                    if true {
                        for fi24 := 0; fi24 < 1; fi24++ {
                            var v25 bool = true
                            fmt.Println(fi23, p1, p0)
                            p0 = (^(p0))
                            v26 := GenS1{F0: 2, F1: -964, F2: 0}
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        p1 -= (p1 + (^(p1)))
                    } else {
                        p1 = 2
                        v27 := GenS4{F0: uint16(p0)}
                        var v28 int16 = -32768
                        _ = v27
                        _ = v28
                    }
                case 2147483647:
                    v29 := []int{fi23, (fi23 << (((p0 & p0)) & 63)), fi23}
                    _ = v29
                case -1:
                    fmt.Println(p0, p1)
                default:
                    p1 += (p1 & 298)
                }
                _ = fi23
            }
        }
    }
    p1 += 42
    v30 := GenS5{F0: -908}
    _ = p0
    _ = p1
    _ = v30
    return 1
}

func main() {
    v1 := []string{}
    var v2 uint8 = 168
    var v3 uint = 358
    var v4 float64 = 0.1
    var v5 int = 2147483647
    var v6 int32 = -2147483648
    v2 = (v2 ^ v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v31 := GenS6{F0: genFn0(false, -1)}
    var v32 int32 = genFn0(false, 0)
    v6 = (v32 << ((380) & 63))
    var v33 uint32 = 65535
    v6 += (v32 >> (((v3 + v3)) & 63))
    switch v5 {
    case 4294967295:
        if true {
            v34 := GenS5{F0: -32768}
            fmt.Println(v6, v4, v33)
            _ = v34
        }
    case -352:
        for fi35 := 0; fi35 < 10; fi35++ {
            switch fi35 {
            case 9223372036854775807:
                v32 = -2147483648
            default:
                var v36 int8 = -17
                _ = v36
            }
            v37 := GenS7{F0: 9223372036854775807, F1: "foo bar"}
            for fi38 := 0; fi38 < 10; fi38++ {
                if false {
                    switch 127 {
                    case -333:
                        var v39 uint8 = v2
                        _ = v39
                    case -151:
                        v1 = append(v1, (("foo bar" + "foo bar") + "\n\t"))
                    case -9223372036854775808:
                        var v40 uint = (v3 % (((v3 % (((v3 | v3)) | 1))) | 1))
                        switch (("\n\t" + "hello") + ((("a" + (((("a" + ("" + "a")) + "") + ("a" + "foo bar")) + "")) + "") + "\n\t")) {
                        case "\n\t":
                            v41 := GenS3{F0: uint64(v2), F1: fi38}
                            _ = v41
                        case "a":
                            var v42 int = fi35
                            _ = v42
                        }
                        _ = v40
                    default:
                        switch fi35 {
                        case 9223372036854775807:
                            fmt.Println(fi38)
                        case -155:
                            if len(v1) > 0 {
                                v1[0] = "hello"
                            }
                        }
                    }
                }
                _ = fi38
            }
            _ = fi35
            _ = v37
        }
    }
    var v43 float64 = (v4 / (v4 / float64(1.5)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
