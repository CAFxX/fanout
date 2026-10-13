package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 int32; F2 int64 }
type GenS4 struct { F0 bool; F1 int; F2 int16 }
type GenS5 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint; F1 bool; F2 int32 }
type GenS8 struct { F0 uint8 }
func genFn0() uint64 {
    var v4 int8 = 42
    if false {
        v4 = 42
        v4 += (v4 % ((1) | 1))
    } else {
        if (0 < 42) {
            v4 = v4
        } else {
            var v5 uint = 513
            _ = v5
        }
        var v6 int32 = -989
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 bool = true
            v9 := GenS1{F0: fi7}
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = v6
    }
    _ = v4
    return 18446744073709551615
}
func genFn1() int {
    var v10 int16 = 507
    v10 = v10
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := []int{fi11, fi11, fi11, fi11}
        var v13 uint64 = 2
        v10 = v10
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = v10
    return 32768
}

func main() {
    var v1 int16 = 1
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]int)
    if true {
        if (true || true) {
            switch "a" {
            case "a":
                v1 = (v1 | v1)
            default:
                switch "" {
                case "foo bar":
                    fmt.Println(v1)
                case "\n\t":
                    v14 := v2[2147483648]
                    _ = v14
                }
            }
            v15 := GenS2{F0: -2147483648}
            if false {
                v2[588] = 42
                fmt.Println(v1)
                if true {
                    for fi16 := 0; fi16 < 7; fi16++ {
                        switch genFn1() {
                        case 9223372036854775807:
                            var v17 int32 = -924
                            _ = v17
                        case -540:
                            var v18 float64 = float64(0.0)
                            _ = v18
                        case -9223372036854775808:
                            fmt.Println(v1, fi16)
                        default:
                            for fi19 := 0; fi19 < 11; fi19++ {
                                var v20 int = -1
                                v21 := v2[446]
                                if (0 == 1) {
                                    var v22 int16 = (v1 * (v1 >> ((18446744073709551615) & 63)))
                                    _ = v22
                                }
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                        }
                        fmt.Println(fi16)
                        _ = fi16
                    }
                    var v23 float32 = 1e10
                    _ = v23
                } else {
                    v1 = -1
                    for fi24 := 0; fi24 < 8; fi24++ {
                        fmt.Println(fi24, v1)
                        switch fi24 {
                        case 1:
                            for fi25 := 0; fi25 < 7; fi25++ {
                                var v26 string = "hello"
                                _ = fi25
                                _ = v26
                            }
                        case -711:
                            if (0 == 18446744073709551615) {
                                if (!false) {
                                    v1 += v1
                                    var v27 int = genFn1()
                                    _ = v27
                                } else {
                                    var v28 int16 = 849
                                    var v29 uint64 = 0
                                    _ = v28
                                    _ = v29
                                }
                                var v30 int32 = 807
                                _ = v30
                            } else {
                                switch genFn1() {
                                case 32768:
                                    if true {
                                        fmt.Println(v1, v1)
                                        v31 := GenS3{F0: v1, F1: 0, F2: 1}
                                        _ = v31
                                    }
                                case 0:
                                    var v32 uint64 = 2
                                    _ = v32
                                case 127:
                                    switch "hello" {
                                    case "a":
                                        fmt.Println(fi24)
                                    case "foo bar":
                                        for fi33 := 0; fi33 < 18; fi33++ {
                                            fmt.Println(v1, fi33, fi33)
                                            v1 += (-(v1))
                                            v34 := GenS4{F0: true, F1: (^(fi33)), F2: v1}
                                            _ = fi33
                                            _ = v34
                                        }
                                    default:
                                        if false {
                                            switch (fi24 - genFn1()) {
                                            case 255:
                                                v35 := []int{42, (fi24 + fi24), fi24}
                                                _ = v35
                                            }
                                            var v36 int = 2147483647
                                            _ = v36
                                        }
                                    }
                                default:
                                    var v37 uint16 = 0
                                    _ = v37
                                }
                            }
                        default:
                            fmt.Println(fi24)
                        }
                        fmt.Println(fi24, fi24)
                        _ = fi24
                    }
                }
            }
            _ = v15
        } else {
            v38 := GenS4{F0: (!(!(891 == 2147483647))), F1: -767, F2: v1}
            _ = v38
        }
    }
    var v39 uint = 0
    _ = len(v2)
    switch 0 {
    case 0:
        v40 := GenS2{F0: -445}
        _ = v40
    case 4294967295:
        v1 -= (v1 | (-(v1)))
    }
    if false {
        fmt.Println(v1)
        v1 = v1
        if (255 <= 65535) {
            var v41 int8 = 89
            _ = v41
        } else {
            fmt.Println(v39)
            if (v1 < v1) {
                v42 := GenS2{F0: 927}
                var v43 int16 = 42
                v1 -= (v43 + (v1 | v1))
                _ = v42
                _ = v43
            }
        }
    } else {
        var v44 float32 = 0.0
        v45 := GenS5{F0: v39, F1: 42, F2: 18446744073709551615}
        _ = v44
        _ = v45
    }
    var v46 bool = true
    var v47 int32 = -1
    for fi48 := 0; fi48 < 11; fi48++ {
        switch "foo bar" {
        case "\n\t":
            var v49 int8 = 127
            _ = v49
        default:
            v39 = (v39 >> ((255) & 63))
        }
        v50 := GenS6{F0: (^(v39))}
        v51 := GenS7{F0: 18446744073709551615, F1: (127 > -128), F2: v47}
        _ = fi48
        _ = v50
        _ = v51
    }
    v1 += v1
    v52 := GenS8{F0: 128}
    v47 -= (v47 % ((-2147483648) | 1))
    if ("" > "hello") {
        var v53 int64 = -1
        _ = v53
    }
    var v54 int8 = 100
    delete(v3, 61)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v46
    _ = v47
    _ = v52
    _ = v54
}
