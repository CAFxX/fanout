package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int64 }
type GenS2 struct { F0 bool; F1 uint; F2 int64 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64; F1 int64 }
type GenS7 struct { F0 bool }

func main() {
    var v1 uint32 = 1
    v2 := []uint32{(v1 * v1), 0}
    v3 := make(map[int]bool)
    var v4 uint = 18446744073709551615
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 int64 = -9223372036854775808
        v7 := GenS1{F0: -547, F1: 18446744073709551615, F2: v6}
        if (float64(1.5) != ((float64(0.0) - 1e10) * 1e10)) {
            v8 := GenS2{F0: (808 > 709), F1: (v4 << (((v4 << ((42) & 63))) & 63)), F2: 776}
            v4 -= v4
            _ = v8
        } else {
            for fi9 := 0; fi9 < 15; fi9++ {
                v10 := GenS2{F0: true, F1: v4, F2: v6}
                v1 = (v1 | 2147483648)
                var v11 int16 = 1
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    v12 := GenS1{F0: 42, F1: 2, F2: 0}
    for fi13 := 0; fi13 < 17; fi13++ {
        v14 := GenS2{F0: false, F1: uint(v1), F2: 0}
        for fi15 := 0; fi15 < 8; fi15++ {
            v16 := v3[(fi15 - fi15)]
            v16 = v16
            v17 := v3[fi15]
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi13
        _ = v14
    }
    v2 = append(v2, v1)
    var v18 uint16 = 65535
    v18 = v18
    v1 += 2147483647
    for fi19 := 0; fi19 < 13; fi19++ {
        if true {
            _ = len(v3)
        }
        v20 := GenS1{F0: 1000, F1: 329, F2: 42}
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        v1 += 47
        _ = fi21
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        if (true && false) {
            v2 = append(v2, v1)
        } else {
            fmt.Println(v18, fi22)
            for fi23 := 0; fi23 < 4; fi23++ {
                fmt.Println(v1, v1)
                var v24 float64 = ((-2.25 * -2.25) - (1.5 + 1e10))
                _ = fi23
                _ = v24
            }
            v25 := GenS1{F0: 0, F1: 18446744073709551615, F2: 42}
            _ = v25
        }
        _ = fi22
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        for fi27 := 0; fi27 < 5; fi27++ {
            if false {
                if false {
                    switch (fi27 + (fi26 % ((fi26) | 1))) {
                    case 503:
                        var v28 uint = (^(v4))
                        _ = v28
                    case 127:
                        fmt.Println(v4, fi26)
                        var v29 string = "foo bar"
                        v18 += v18
                        _ = v29
                    case 1:
                        var v30 int32 = 535
                        _ = v30
                    }
                    var v31 float32 = 0.0
                    var v32 float32 = (v31 - v31)
                    _ = v31
                    _ = v32
                }
                if true {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    if (!false) {
                        if (true || false) {
                            if (v1 != 754) {
                                if true {
                                    v33 := v3[fi27]
                                    _ = v33
                                }
                                if false {
                                    switch (^(fi27)) {
                                    case 2147483647:
                                        fmt.Println(fi27)
                                    case -9223372036854775808:
                                        v4 += 42
                                    case 32767:
                                        var v34 float64 = (-2.25 / (float64(0.1) / 1e10))
                                        _ = v34
                                    }
                                    fmt.Println(v1)
                                    v3[0] = (!(1000 > 798))
                                }
                                if len(v2) > 0 {
                                    v2[0] = v1
                                }
                            } else {
                                fmt.Println(fi27)
                            }
                            if ((-64 < -32768) && true) {
                                fmt.Println(fi27, fi27)
                                v4 -= (v4 ^ v4)
                                v18 = (v18 & v18)
                            } else {
                                v1 = (v1 % ((v1) | 1))
                                v35 := GenS1{F0: 0, F1: 1, F2: -9223372036854775808}
                                _ = v35
                            }
                            switch "hello" {
                            case "a":
                                v1 = 42
                            case "":
                                var v36 uint64 = 18446744073709551615
                                if true {
                                    v37 := GenS3{F0: 42, F1: v18}
                                    switch (^(fi27)) {
                                    case -173:
                                        fmt.Println(v4, v18, v4)
                                    case -395:
                                        v1 = uint32(v4)
                                    case -2147483648:
                                        v38 := GenS1{F0: -65, F1: v36, F2: 42}
                                        _ = v38
                                    }
                                    _ = v37
                                }
                                v39 := GenS4{F0: 32767, F1: (v36 ^ (v36 * v36)), F2: v36}
                                _ = v36
                                _ = v39
                            }
                        }
                        fmt.Println(fi27, fi27)
                        fmt.Println(fi26, fi27, v18)
                    }
                    var v40 int8 = 100
                    _ = v40
                }
            } else {
                v18 = (v18 & 42)
            }
            _ = fi27
        }
        var v41 uint8 = uint8(v4)
        if true {
            for fi42 := 0; fi42 < 6; fi42++ {
                var v43 uint = v4
                var v44 int = fi26
                var v45 uint32 = (v1 ^ (v1 << (((v4 - (v4 ^ 4294967295))) & 63)))
                _ = fi42
                _ = v43
                _ = v44
                _ = v45
            }
            for fi46 := 0; fi46 < 19; fi46++ {
                v47 := GenS3{F0: 0, F1: (v18 + (v18 & (^(v18))))}
                switch fi26 {
                case 2147483647:
                    v1 -= v1
                case 4294967295:
                    v41 = (v41 * 42)
                default:
                    if false {
                        var v48 uint32 = v1
                        var v49 float64 = (1.5 - (1.5 / 3.14159))
                        v48 += 1
                        _ = v48
                        _ = v49
                    } else {
                        if (true && (false && ((v1 <= 1) && (((978 == 0) || true) || false)))) {
                            v1 -= v1
                            var v50 uint64 = 1
                            _ = v50
                        }
                        var v51 uint = v4
                        var v52 uint64 = 0
                        _ = v51
                        _ = v52
                    }
                }
                v1 = (v1 << ((v4) & 63))
                _ = fi46
                _ = v47
            }
            for fi53 := 0; fi53 < 13; fi53++ {
                v54 := GenS1{F0: 1000, F1: 18446744073709551615, F2: -127}
                if true {
                    v18 -= (v18 - v18)
                } else {
                    if ((v4 != v4) || false) {
                        var v55 string = "a"
                        _ = v55
                    }
                    var v56 int64 = -963
                    v2 = append(v2, v1)
                    _ = v56
                }
                v57 := v3[(-(fi26))]
                v58 := GenS5{F0: (v4 + (v4 * v4))}
                _ = fi53
                _ = v54
                _ = v57
                _ = v58
            }
        } else {
            if len(v2) > 0 {
                v2[0] = v1
            }
            v59 := GenS3{F0: -272, F1: v18}
            for fi60 := 0; fi60 < 15; fi60++ {
                v4 = (v4 ^ v4)
                var v61 uint16 = (^(v18))
                var v62 int32 = -1
                _ = fi60
                _ = v61
                _ = v62
            }
            _ = v59
        }
        _ = fi26
        _ = v41
    }
    fmt.Println(v18)
    fmt.Println(v4, v18, v4)
    v63 := GenS2{F0: true, F1: 624, F2: 9223372036854775807}
    for fi64 := 0; fi64 < 14; fi64++ {
        fmt.Println(v4)
        v18 = v18
        switch ("\n\t" + ("hello" + "\n\t")) {
        case "":
            for fi65 := 0; fi65 < 4; fi65++ {
                var v66 int16 = 583
                v67 := GenS6{F0: -1, F1: -41}
                _ = fi65
                _ = v66
                _ = v67
            }
        case "a":
            if true {
                if false {
                    v18 = 65535
                    v18 -= (^(v18))
                    v1 = v1
                } else {
                    v68 := GenS7{F0: ((0 != uint64(v1)) && (false || false))}
                    _ = v68
                }
            }
        case "foo bar":
            for fi69 := 0; fi69 < 18; fi69++ {
                var v70 int8 = int8(fi64)
                _ = fi69
                _ = v70
            }
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v63
}
