package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 string; F1 int }
func genFn0(p0 uint64) uint32 {
    v7 := []int{127}
    for fi8 := 0; fi8 < 10; fi8++ {
        if (0 >= 127) {
            fmt.Println(fi8)
            for fi9 := 0; fi9 < 2; fi9++ {
                switch (fi9 | (fi9 | (fi8 << ((748) & 63)))) {
                case 657:
                    p0 += p0
                case 32768:
                    var v10 uint32 = 4294967295
                    _ = v10
                default:
                    v11 := GenS1{F0: 1}
                    _ = v11
                }
                var v12 float32 = (0.0 * 1.5)
                p0 -= 18446744073709551615
                _ = fi9
                _ = v12
            }
            v7 = append(v7, -1)
        } else {
            for fi13 := 0; fi13 < 16; fi13++ {
                fmt.Println(p0)
                _ = fi13
            }
            if (4294967295 < 4294967295) {
                fmt.Println(p0, fi8)
            }
        }
        if false {
            var v14 uint32 = 1
            for fi15 := 0; fi15 < 13; fi15++ {
                switch (fi15 & (fi8 % ((9223372036854775807) | 1))) {
                case 128:
                    var v16 int = (-(fi8))
                    _ = v16
                case -2147483648:
                    v7 = append(v7, (fi8 | 2147483647))
                case 0:
                    var v17 int16 = 623
                    _ = v17
                default:
                    p0 = (p0 / (((p0 / ((p0) | 1))) | 1))
                }
                v18 := GenS2{F0: -9223372036854775808, F1: 42}
                fmt.Println(fi15, fi15)
                _ = fi15
                _ = v18
            }
            var v19 float64 = 0.1
            _ = v14
            _ = v19
        } else {
            if true {
                switch (fi8 >> ((155) & 63)) {
                case -2147483648:
                    var v20 float32 = 1e10
                    _ = v20
                case -9223372036854775808:
                    var v21 bool = (true && false)
                    _ = v21
                case -195:
                    p0 = p0
                default:
                    p0 = p0
                }
                var v22 uint64 = p0
                _ = v22
            }
            for fi23 := 0; fi23 < 15; fi23++ {
                v24 := GenS1{F0: 153}
                _ = fi23
                _ = v24
            }
        }
        p0 += (p0 >> ((42) & 63))
        _ = fi8
    }
    _ = p0
    _ = v7
    return 2147483647
}
func genFn1(p0 uint8, p1 uint64, p2 uint8) uint64 {
    for fi25 := 0; fi25 < 13; fi25++ {
        var v26 uint32 = 2
        if (p1 >= p1) {
            if false {
                v27 := GenS3{F0: genFn0((p1 >> ((1) & 63))), F1: 42}
                _ = v27
            } else {
                v26 = v26
                if ((p0 | (p2 | p2)) >= p2) {
                    var v28 string = "hello"
                    _ = v28
                } else {
                    if (p2 >= 255) {
                        p1 = 283
                        var v29 float64 = 1.5
                        p2 += p0
                        _ = v29
                    } else {
                        v30 := []int{(^(fi25)), fi25, 127, (fi25 - fi25)}
                        switch (fi25 >> ((18446744073709551615) & 63)) {
                        case 0:
                            v26 -= v26
                        case -1:
                            v30 = append(v30, fi25)
                        case 32767:
                            if false {
                                v26 = (^(v26))
                                v30 = append(v30, (fi25 % (((^(fi25))) | 1)))
                            }
                        default:
                            var v31 uint8 = 0
                            _ = v31
                        }
                        _ = v30
                    }
                    p2 -= p2
                    v32 := []int{fi25}
                    _ = v32
                }
            }
        }
        for fi33 := 0; fi33 < 7; fi33++ {
            fmt.Println(p1, p0)
            v26 = v26
            if true {
                var v34 uint64 = p1
                p2 += p2
                var v35 int = 128
                _ = v34
                _ = v35
            } else {
                v36 := GenS4{F0: "a", F1: fi33}
                v37 := []int{(fi33 % (((fi25 & -9223372036854775808)) | 1)), fi25, 127, fi25}
                if false {
                    switch ("a" + ("\n\t" + "")) {
                    case "a":
                        var v38 uint32 = 561
                        _ = v38
                    default:
                        if (0 == int16(fi33)) {
                            if len(v37) > 0 {
                                v37[0] = 32768
                            }
                            p0 -= (p2 ^ (^(p0)))
                            fmt.Println(p1, v26, fi33)
                        }
                    }
                }
                _ = v36
                _ = v37
            }
            _ = fi33
        }
        _ = fi25
        _ = v26
    }
    p1 += p1
    v39 := []int{128, 127, -823, -342}
    _ = p0
    _ = p1
    _ = p2
    _ = v39
    return p1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 2
    var v3 uint = 18446744073709551615
    v4 := make(map[uint32]bool)
    v5 := make(map[string]bool)
    v6 := []int{-38, 4294967295, 42, 0}
    v40 := GenS3{F0: 638, F1: 937}
    var v41 uint32 = 65535
    v41 += 764
    v42 := GenS2{F0: 42, F1: 448}
    delete(v4, v41)
    v2 -= 2
    var v43 int32 = -788
    for fi44 := 0; fi44 < 11; fi44++ {
        _ = len(v6)
        var v45 int16 = 828
        var v46 int64 = -918
        _ = fi44
        _ = v45
        _ = v46
    }
    for fi47 := 0; fi47 < 11; fi47++ {
        var v48 bool = true
        _ = fi47
        _ = v48
    }
    v3 += v3
    v3 -= 42
    switch -1 {
    case -2147483648:
        v49 := GenS1{F0: (v2 ^ (v2 / (((v2 << ((v3) & 63))) | 1)))}
        _ = v49
    case -9223372036854775808:
        if (true || false) {
            v50 := GenS2{F0: -9223372036854775808, F1: -516}
            switch 0 {
            case 151:
                if (42 < 26) {
                    v2 = 255
                    _ = len(v6)
                }
            }
            _ = v50
        }
    case 1:
        fmt.Println(v41)
    default:
        fmt.Println(v2, v3, v41)
    }
    var v51 int32 = v43
    fmt.Println(v2, v2, v3)
    var v52 int16 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v51
    _ = v52
    _ = v6
}
