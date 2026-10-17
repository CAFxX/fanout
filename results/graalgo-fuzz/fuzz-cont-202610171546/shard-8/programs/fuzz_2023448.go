package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool; F1 int16; F2 int }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int64; F1 uint32; F2 bool }
type GenS9 struct { F0 int64; F1 bool }
type GenS10 struct { F0 uint32; F1 int32 }
type GenS11 struct { F0 string; F1 int8; F2 int32 }
func genFn0(p0 uint64) int16 {
    p0 -= (p0 << ((371) & 63))
    for fi4 := 0; fi4 < 10; fi4++ {
        v5 := GenS1{F0: 2}
        _ = fi4
        _ = v5
    }
    if (0 != 32) {
        if false {
            p0 = p0
        }
        if (42 != 32767) {
            v6 := []int{-2147483648, 1, 42, 9223372036854775807}
            _ = v6
        }
    } else {
        var v7 int16 = 42
        fmt.Println(v7, p0)
        v7 += (v7 | (v7 / ((v7) | 1)))
        _ = v7
    }
    _ = p0
    return -32768
}

func main() {
    var v1 int32 = 1
    var v2 uint64 = 18446744073709551615
    v3 := make(map[string]string)
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := []int{}
        fmt.Println(fi8, fi8)
        var v10 float64 = (-2.25 / 1.5)
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        switch "\n\t" {
        case "a":
            v12 := GenS2{F0: true}
            _ = v12
        case "hello":
            fmt.Println(v2, v1, v2)
        }
        _ = fi11
    }
    var v13 int64 = 497
    fmt.Println(v13, v13)
    v14 := GenS3{F0: "hello"}
    fmt.Println(v13, v2, v13)
    v15 := []int{127, 1}
    fmt.Println(v2, v13)
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 uint16 = 835
        for fi18 := 0; fi18 < 4; fi18++ {
            if false {
                switch int(v17) {
                case 32768:
                    v1 -= (v1 / ((-675) | 1))
                case 128:
                    v19 := GenS2{F0: false}
                    _ = v19
                default:
                    v3["\n\t"] = ""
                }
                v20 := GenS4{F0: 42, F1: -137}
                if true {
                    switch (("foo bar" + "hello") + "foo bar") {
                    case "\n\t":
                        if ((!false) || false) {
                            fmt.Println(fi16, v2)
                        } else {
                            v21 := GenS4{F0: 127, F1: 42}
                            _ = v21
                        }
                    default:
                        if true {
                            if false {
                                switch (fi16 + 11) {
                                case 32768:
                                    fmt.Println(fi18)
                                case 127:
                                    v22 := GenS5{F0: 272}
                                    _ = v22
                                case 42:
                                    if (177 != 1) {
                                        v23 := GenS6{F0: false, F1: 1, F2: -703}
                                        _ = v23
                                    } else {
                                        var v24 bool = (v13 >= (v13 >> ((uint(v2)) & 63)))
                                        _ = v24
                                    }
                                default:
                                    var v25 bool = true
                                    _ = v25
                                }
                                fmt.Println(fi18, fi16)
                            } else {
                                if (1.5 < 0.1) {
                                    v26 := GenS5{F0: 42}
                                    _ = v26
                                } else {
                                    var v27 uint8 = 0
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                    _ = v27
                                }
                            }
                            fmt.Println(fi16, v2, fi16)
                            if true {
                                if (18446744073709551615 != 65535) {
                                    v28 := GenS7{F0: 886}
                                    v13 = 9223372036854775807
                                    _ = v28
                                }
                                v29 := GenS7{F0: uint(v17)}
                                _ = v29
                            } else {
                                v2 -= (v2 - (^(v2)))
                                v2 -= (v2 / ((42) | 1))
                                v30 := GenS8{F0: v13, F1: 374, F2: ((32767 != v17) && true)}
                                _ = v30
                            }
                            fmt.Println(v17, fi18, v13)
                            switch -9223372036854775808 {
                            case -9223372036854775808:
                                var v31 uint64 = 42
                                _ = v31
                            }
                            if (1 <= 271) {
                                if (!(v13 >= v13)) {
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                    var v32 bool = true
                                    _ = v32
                                }
                                v33 := v3["a"]
                                _ = len(v15)
                                _ = v33
                            }
                        } else {
                            fmt.Println(v13, v2, fi18)
                            var v34 uint8 = 173
                            v35 := GenS1{F0: (v17 + (v17 - (^(v17))))}
                            _ = v34
                            _ = v35
                        }
                    }
                    if ((-(v1)) > -2147483648) {
                        v13 += (v13 * v13)
                        v36 := GenS1{F0: 421}
                        _ = v36
                    } else {
                        var v37 int32 = 0
                        fmt.Println(v13, v2)
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        _ = v37
                    }
                }
                _ = v20
            } else {
                v38 := GenS9{F0: -838, F1: (!("\n\t" >= "\n\t"))}
                v13 += (v13 & (v13 ^ v13))
                _ = v38
            }
            var v39 int64 = (v13 & -9223372036854775808)
            v39 = (v13 << ((4294967295) & 63))
            fmt.Println(v13, v1, v39)
            var v40 int32 = (v1 << ((274) & 63))
            _ = fi18
            _ = v39
            _ = v40
        }
        _ = fi16
        _ = v17
    }
    switch "\n\t" {
    case "hello":
        _ = len(v15)
    case "foo bar":
        v15 = append(v15, 1)
    }
    v3["a"] = "\n\t"
    switch "a" {
    case "a":
        if (v2 >= 4294967295) {
            var v41 int64 = -547
            var v42 uint = uint(v2)
            _ = v41
            _ = v42
        }
    case "":
        for fi43 := 0; fi43 < 16; fi43++ {
            var v44 uint = 362
            var v45 int64 = 42
            if false {
                v46 := GenS10{F0: 2147483648, F1: v1}
                v47 := GenS11{F0: "hello", F1: 127, F2: (v1 * -766)}
                if true {
                    for fi48 := 0; fi48 < 16; fi48++ {
                        var v49 uint16 = 32767
                        _ = fi48
                        _ = v49
                    }
                    switch (fi43 * 877) {
                    case 678:
                        fmt.Println(v13)
                    }
                } else {
                    fmt.Println(v13, v45)
                }
                _ = v46
                _ = v47
            } else {
                fmt.Println(v1, v13)
                var v50 int32 = 42
                switch (fi43 & 9223372036854775807) {
                case 127:
                    v13 = (v13 / ((42) | 1))
                case 32767:
                    fmt.Println(v44, v1, v2)
                case -1:
                    if len(v15) > 0 {
                        v15[0] = (fi43 << (((v44 / ((v44) | 1))) & 63))
                    }
                default:
                    v51 := GenS3{F0: "a"}
                    _ = v51
                }
                _ = v50
            }
            _ = fi43
            _ = v44
            _ = v45
        }
    case "\n\t":
        fmt.Println(v13, v2, v1)
    }
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    v13 = v13
    var v52 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v52:", v52)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v52
}
