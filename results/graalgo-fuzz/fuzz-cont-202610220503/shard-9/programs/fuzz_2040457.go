package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint16; F1 int64; F2 bool }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint; F1 bool }
type GenS8 struct { F0 int8; F1 int32 }
func genFn0(p0 int8) uint {
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS1{F0: 42}
        for fi9 := 0; fi9 < 17; fi9++ {
            if false {
                v10 := []int{}
                p0 += 127
                fmt.Println(p0)
                _ = v10
            } else {
                if true {
                    v11 := []int{fi7, -790}
                    _ = v11
                }
                v12 := []int{fi9, (fi7 - -799), -333}
                _ = v12
            }
            var v13 uint16 = 32767
            p0 -= (p0 - p0)
            _ = fi9
            _ = v13
        }
        _ = fi7
        _ = v8
    }
    switch "" {
    case "hello":
        var v14 uint8 = 75
        _ = v14
    default:
        v15 := GenS1{F0: 1}
        _ = v15
    }
    if (false || false) {
        fmt.Println(p0)
        v16 := []int{128, -2147483648, 32768, 4294967295}
        p0 += 42
        _ = v16
    }
    var v17 int32 = -63
    _ = p0
    _ = v17
    return 2
}
func genFn1(p0 int64) int8 {
    for fi18 := 0; fi18 < 2; fi18++ {
        switch fi18 {
        case 127:
            switch ("" + "a") {
            case "a":
                for fi19 := 0; fi19 < 8; fi19++ {
                    p0 += p0
                    v20 := GenS1{F0: 61}
                    v21 := []int{fi19}
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            case "hello":
                v22 := GenS1{F0: 29}
                _ = v22
            case "":
                var v23 int = 432
                _ = v23
            }
        case -152:
            if (83 < 222) {
                p0 += (p0 % (((-(p0))) | 1))
                fmt.Println(p0, fi18)
                v24 := []int{1, (^(fi18))}
                _ = v24
            }
        }
        v25 := GenS2{F0: 42, F1: p0}
        _ = fi18
        _ = v25
    }
    if false {
        for fi26 := 0; fi26 < 5; fi26++ {
            v27 := GenS2{F0: 18446744073709551615, F1: -1}
            var v28 uint64 = 18446744073709551615
            fmt.Println(fi26, fi26)
            _ = fi26
            _ = v27
            _ = v28
        }
        var v29 int32 = 2147483647
        v30 := GenS2{F0: 4294967295, F1: (p0 * p0)}
        _ = v29
        _ = v30
    } else {
        var v31 uint = 65535
        var v32 uint64 = 18446744073709551615
        _ = v31
        _ = v32
    }
    _ = p0
    return 71
}

func main() {
    var v1 uint = 65535
    var v2 uint = 65535
    var v3 int = 1
    v4 := make(map[string]int)
    v5 := []uint32{577, 790}
    var v6 bool = false
    v5 = append(v5, 4294967295)
    var v33 float64 = 0.1
    for fi34 := 0; fi34 < 17; fi34++ {
        if v6 {
            if v6 {
                var v35 uint64 = 2
                _ = v35
            }
        } else {
            for fi36 := 0; fi36 < 17; fi36++ {
                switch (fi36 | fi36) {
                case 1:
                    if v6 {
                        var v37 float64 = (0.0 * (1e10 - v33))
                        _ = v37
                    } else {
                        var v38 int32 = -755
                        _ = v38
                    }
                case 0:
                    fmt.Println(v6, v2)
                    delete(v4, "\n\t")
                    switch v3 {
                    case 9223372036854775807:
                        v39 := GenS3{F0: "\n\t", F1: 4294967295}
                        _ = v39
                    case 4294967295:
                        v1 -= 2
                    case 0:
                        var v40 uint64 = 18446744073709551615
                        _ = v40
                    default:
                        v41 := GenS3{F0: "hello", F1: 18446744073709551615}
                        if (int8(v2) <= genFn1(0)) {
                            delete(v4, "\n\t")
                            fmt.Println(v6, v1, v33)
                            var v42 uint = v2
                            _ = v42
                        } else {
                            var v43 bool = (v6 || (2 < 1))
                            v44 := v4[("foo bar" + "")]
                            if (!v6) {
                                if ((434 >= 42) && (32767 >= 32767)) {
                                    var v45 uint = v1
                                    var v46 uint32 = 2147483648
                                    _ = v45
                                    _ = v46
                                }
                                v47 := GenS2{F0: 18446744073709551615, F1: -9223372036854775808}
                                v48 := GenS4{F0: 65535, F1: -9223372036854775808, F2: v43}
                                _ = v47
                                _ = v48
                            } else {
                                v3 = (fi36 - v44)
                                v44 += fi34
                            }
                            _ = v43
                            _ = v44
                        }
                        _ = v41
                    }
                case 737:
                    fmt.Println(fi36)
                }
                _ = fi36
            }
            if false {
                v2 = 65535
            }
        }
        switch ("\n\t" + "\n\t") {
        case "a":
            var v49 string = (("hello" + ("" + "hello")) + "hello")
            _ = v49
        case "\n\t":
            v2 = 65535
        case "foo bar":
            var v50 uint64 = 42
            _ = v50
        default:
            for fi51 := 0; fi51 < 19; fi51++ {
                var v52 int16 = 871
                var v53 uint64 = 18446744073709551615
                v52 = (v52 >> ((uint(v53)) & 63))
                _ = fi51
                _ = v52
                _ = v53
            }
        }
        _ = fi34
    }
    if false {
        for fi54 := 0; fi54 < 16; fi54++ {
            if v6 {
                v6 = ((fi54 > v3) && (v6 || false))
            }
            v55 := v4["foo bar"]
            _ = fi54
            _ = v55
        }
        v56 := GenS2{F0: 4294967295, F1: 514}
        _ = v56
    }
    v4["\n\t"] = v3
    for fi57 := 0; fi57 < 1; fi57++ {
        fmt.Println(v1, fi57, v2)
        _ = fi57
    }
    v3 -= (^(v3))
    var v58 uint64 = 1
    v4[""] = v3
    var v59 int = v3
    v4[(("\n\t" + "\n\t") + "a")] = (-(v3))
    for fi60 := 0; fi60 < 18; fi60++ {
        for fi61 := 0; fi61 < 19; fi61++ {
            var v62 uint32 = 991
            _ = fi61
            _ = v62
        }
        _ = fi60
    }
    v3 = (v59 ^ (v3 | (v3 << ((v2) & 63))))
    for fi63 := 0; fi63 < 7; fi63++ {
        switch ("foo bar" + "a") {
        case "\n\t":
            for fi64 := 0; fi64 < 12; fi64++ {
                switch (fi63 << ((42) & 63)) {
                case 4294967295:
                    var v65 uint16 = 42
                    _ = v65
                case 20:
                    var v66 int = v3
                    _ = v66
                case 32767:
                    delete(v4, "hello")
                    var v67 int32 = -177
                    v68 := GenS1{F0: 255}
                    _ = v67
                    _ = v68
                }
                v69 := GenS5{F0: (v2 - 59), F1: v2}
                if (false && (!true)) {
                    if false {
                        v2 = genFn0(-1)
                    }
                } else {
                    v70 := GenS6{F0: uint32(v1)}
                    _ = v70
                }
                _ = fi64
                _ = v69
            }
        }
        _ = fi63
    }
    v71 := GenS7{F0: v2, F1: (((0 <= 32767) || (4294967295 > (^(v58)))) || (-128 > 127))}
    for fi72 := 0; fi72 < 2; fi72++ {
        v33 += 1.5
        if len(v5) > 0 {
            v5[0] = 1
        }
        fmt.Println(fi72, v1)
        _ = fi72
    }
    fmt.Println(v1)
    switch v59 {
    case -503:
        var v73 int32 = 207
        _ = v73
    case 42:
        if v6 {
            v74 := GenS7{F0: 4294967295, F1: v6}
            if v6 {
                fmt.Println(v1)
                v1 += v2
                if (255 != 42) {
                    v75 := GenS5{F0: v1, F1: v2}
                    v1 -= genFn0(83)
                    for fi76 := 0; fi76 < 1; fi76++ {
                        _ = len(v4)
                        v77 := GenS8{F0: -114, F1: 503}
                        _ = fi76
                        _ = v77
                    }
                    _ = v75
                }
            }
            v3 = -1
            _ = v74
        }
    default:
        _ = len(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v71
}
