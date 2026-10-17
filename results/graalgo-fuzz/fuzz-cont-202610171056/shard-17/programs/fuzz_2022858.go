package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 string }
type GenS2 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS3 struct { F0 int8; F1 string; F2 int8 }
type GenS4 struct { F0 int8 }
func genFn0(p0 int16, p1 uint16, p2 int32) int32 {
    var v7 uint8 = 30
    v8 := GenS1{F0: 0, F1: 207, F2: (("foo bar" + "") + (("hello" + ("\n\t" + "a")) + ("foo bar" + "hello")))}
    fmt.Println(p2, p0)
    v9 := GenS1{F0: v7, F1: 2, F2: ("a" + "\n\t")}
    v10 := []int{int(v7)}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return p2
}
func genFn1(p0 uint64, p1 uint) uint64 {
    v11 := []int{}
    var v12 uint8 = 255
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return p0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 42
    v3 := []int{32768, -764}
    v4 := []uint32{65535, 441, 2}
    v5 := []int{-205, -2147483648, 32767, -292}
    var v6 uint = 97
    switch "hello" {
    case "":
        v13 := GenS1{F0: 127, F1: 696, F2: "hello"}
        _ = v13
    case "hello":
        v2 = genFn1((v2 * (v2 | v2)), (v6 % ((v6) | 1)))
    case "a":
        v2 = v2
    default:
        v14 := v1[4294967295]
        _ = v14
    }
    v15 := GenS2{F0: 61, F1: -82, F2: 42}
    var v16 float64 = (1e10 - (1e10 * 0.0))
    switch "" {
    case "hello":
        switch 32767 {
        case 9223372036854775807:
            if true {
                v17 := GenS1{F0: 14, F1: 498, F2: ""}
                if true {
                    for fi18 := 0; fi18 < 17; fi18++ {
                        v6 += v6
                        for fi19 := 0; fi19 < 13; fi19++ {
                            if false {
                                fmt.Println(v6, v2, fi18)
                                v16 += v16
                                v20 := v1[401]
                                _ = v20
                            } else {
                                fmt.Println(v16, fi18, fi18)
                            }
                            v21 := GenS1{F0: 42, F1: 2147483648, F2: "hello"}
                            if (0 < 0) {
                                var v22 uint64 = v2
                                _ = v22
                            } else {
                                _ = len(v1)
                                v2 = (v2 & (v2 + (^(v2))))
                                v16 += 1.5
                            }
                            _ = fi19
                            _ = v21
                        }
                        if false {
                            for fi23 := 0; fi23 < 5; fi23++ {
                                fmt.Println(fi23, v2)
                                var v24 string = "\n\t"
                                var v25 uint32 = 65535
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                            fmt.Println(v16)
                            if false {
                                if true {
                                    var v26 uint = (v6 / ((v6) | 1))
                                    _ = v26
                                } else {
                                    for fi27 := 0; fi27 < 16; fi27++ {
                                        v6 += 1
                                        _ = fi27
                                    }
                                }
                                v28 := v1[65535]
                                v29 := GenS2{F0: 42, F1: -1, F2: 127}
                                _ = v28
                                _ = v29
                            }
                        }
                        _ = fi18
                    }
                }
                _ = v17
            } else {
                if true {
                    v3 = append(v3, 0)
                    v16 = v16
                } else {
                    v6 -= 65535
                }
                var v30 int32 = 42
                for fi31 := 0; fi31 < 15; fi31++ {
                    v32 := GenS1{F0: 255, F1: 660, F2: (("foo bar" + "\n\t") + "hello")}
                    _ = fi31
                    _ = v32
                }
                _ = v30
            }
        case 922:
            if (1 <= 2147483647) {
                if false {
                    for fi33 := 0; fi33 < 5; fi33++ {
                        var v34 int16 = 0
                        var v35 int64 = -70
                        var v36 int = 9223372036854775807
                        _ = fi33
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                    v37 := GenS2{F0: 2, F1: 493, F2: 42}
                    fmt.Println(v2)
                    _ = v37
                }
                for fi38 := 0; fi38 < 11; fi38++ {
                    v39 := GenS3{F0: -128, F1: "foo bar", F2: 42}
                    _ = fi38
                    _ = v39
                }
                if true {
                    for fi40 := 0; fi40 < 12; fi40++ {
                        switch (fi40 >> (((v6 % ((v6) | 1))) & 63)) {
                        case 32768:
                            var v41 int16 = 0
                            _ = v41
                        default:
                            for fi42 := 0; fi42 < 18; fi42++ {
                                var v43 int16 = -315
                                switch (("hello" + "") + "foo bar") {
                                case "":
                                    delete(v1, 2147483648)
                                case "a":
                                    fmt.Println(v43)
                                case "\n\t":
                                    v44 := GenS4{F0: -1}
                                    _ = v44
                                }
                                _ = fi42
                                _ = v43
                            }
                        }
                        var v45 uint32 = 2147483647
                        _ = fi40
                        _ = v45
                    }
                    fmt.Println(v2, v2)
                }
            } else {
                v16 += -2.25
                if false {
                    v2 -= (v2 ^ (v2 % ((2) | 1)))
                    var v46 int16 = -32768
                    v47 := GenS2{F0: 127, F1: -1, F2: 100}
                    _ = v46
                    _ = v47
                }
                fmt.Println(v16)
            }
        }
    }
    if len(v3) > 0 {
        v3[0] = 128
    }
    fmt.Println(v2, v16)
    v2 = (^(v2))
    var v48 uint64 = v2
    if (42 < 127) {
        if (true || false) {
            var v49 bool = (((false || true) || false) || true)
            _ = v49
        } else {
            var v50 int64 = 814
            _ = v50
        }
        var v51 float32 = (1e10 * 1.5)
        v51 -= v51
        _ = v51
    } else {
        var v52 int32 = genFn0(0, 81, -547)
        for fi53 := 0; fi53 < 7; fi53++ {
            var v54 uint32 = 2147483648
            v55 := GenS4{F0: -118}
            var v56 int16 = -1
            _ = fi53
            _ = v54
            _ = v55
            _ = v56
        }
        _ = v52
    }
    for fi57 := 0; fi57 < 14; fi57++ {
        fmt.Println(v16, v48, v16)
        var v58 int32 = -2147483648
        v48 = genFn1(v2, (v6 << ((uint(v2)) & 63)))
        _ = fi57
        _ = v58
    }
    var v59 uint16 = 42
    var v60 bool = (uint16(v2) <= (v59 - (v59 << ((v6) & 63))))
    var v61 int8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v59
    _ = v6
    _ = v60
    _ = v61
}
