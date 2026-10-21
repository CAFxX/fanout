package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64 }
func genFn0(p0 uint32, p1 string) uint {
    var v5 int8 = -1
    v5 = v5
    var v6 int64 = -955
    if false {
        var v7 int64 = 42
        switch 222 {
        case -782:
            var v8 uint64 = 242
            _ = v8
        case -941:
            if (-837 < 698) {
                v9 := GenS1{F0: 2, F1: 842, F2: (v7 / ((0) | 1))}
                v6 -= v7
                _ = v9
            } else {
                if ((true || ((-1 < 32767) || (65535 <= 42))) || ((1.5 != 0.0) && true)) {
                    if (true || (!true)) {
                        v10 := GenS1{F0: 255, F1: p0, F2: v7}
                        _ = v10
                    }
                    v11 := []int{}
                    _ = v11
                } else {
                    switch 428 {
                    case 9223372036854775807:
                        p1 = (p1 + ("a" + p1))
                    }
                    v12 := []int{-9223372036854775808, 42}
                    _ = v12
                }
                v13 := []int{2147483647, -9223372036854775808}
                _ = v13
            }
        case 128:
            switch 515 {
            case -758:
                var v14 int32 = 2147483647
                _ = v14
            case 32767:
                v15 := []int{}
                _ = v15
            }
        }
        _ = v7
    }
    v16 := GenS1{F0: 904, F1: p0, F2: (v6 / (((-(v6))) | 1))}
    _ = p0
    _ = p1
    _ = v16
    _ = v5
    _ = v6
    return 441
}
func genFn1(p0 string, p1 bool) bool {
    v17 := GenS1{F0: 669, F1: 4294967295, F2: 1}
    switch ((("a" + p0) + "") + "\n\t") {
    case "foo bar":
        var v18 int64 = 0
        _ = v18
    }
    if (p1 || false) {
        v19 := GenS2{F0: -714}
        var v20 float64 = -2.25
        for fi21 := 0; fi21 < 14; fi21++ {
            p1 = p1
            if (!p1) {
                if (42 > 78) {
                    p1 = (-459 > -32768)
                    v22 := GenS2{F0: (fi21 << ((42) & 63))}
                    _ = v22
                } else {
                    if (18446744073709551615 == 1) {
                        switch -382 {
                        case -509:
                            var v23 string = ("foo bar" + p0)
                            _ = v23
                        case 2147483647:
                            for fi24 := 0; fi24 < 11; fi24++ {
                                p0 = "a"
                                if (p1 && p1) {
                                    var v25 int16 = 989
                                    switch "a" {
                                    case "a":
                                        var v26 string = (("\n\t" + (((p0 + p0) + p0) + "hello")) + "foo bar")
                                        _ = v26
                                    default:
                                        v27 := []int{(^(fi21))}
                                        _ = v27
                                    }
                                    var v28 uint64 = 839
                                    _ = v25
                                    _ = v28
                                } else {
                                    v29 := GenS1{F0: 42, F1: 4294967295, F2: 42}
                                    _ = v29
                                }
                                var v30 uint64 = 18446744073709551615
                                _ = fi24
                                _ = v30
                            }
                        }
                        if (p1 && true) {
                            var v31 uint64 = 4294967295
                            _ = v31
                        }
                    } else {
                        for fi32 := 0; fi32 < 6; fi32++ {
                            var v33 int = fi21
                            if (!((!(4294967295 >= 0)) && (2147483647 > 537))) {
                                var v34 uint = 2
                                _ = v34
                            }
                            if p1 {
                                v33 = fi32
                            }
                            _ = fi32
                            _ = v33
                        }
                    }
                    v35 := []int{4294967295}
                    _ = v35
                }
            }
            _ = fi21
        }
        _ = v19
        _ = v20
    }
    if (105 != 127) {
        for fi36 := 0; fi36 < 11; fi36++ {
            var v37 bool = p1
            p0 = p0
            _ = fi36
            _ = v37
        }
        v38 := []int{-162, 1, 9223372036854775807, -9223372036854775808}
        _ = v38
    }
    switch p0 {
    case "hello":
        v39 := GenS3{F0: 1, F1: (!((p1 || (!true)) || p1))}
        _ = v39
    default:
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    _ = v17
    return (2 > 32767)
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 42
    v3 := make(map[string]uint32)
    var v4 uint8 = 255
    v40 := GenS2{F0: 838}
    var v41 int16 = 42
    if true {
        var v42 bool = false
        _ = v42
    } else {
        if true {
            v43 := GenS4{F0: 0, F1: 42}
            _ = v43
        }
    }
    switch "hello" {
    case "a":
        v4 = 27
    case "hello":
        v44 := GenS2{F0: 872}
        _ = v44
    case "foo bar":
        for fi45 := 0; fi45 < 6; fi45++ {
            if (!false) {
                v46 := GenS5{F0: uint32(v2)}
                _ = v46
            }
            _ = fi45
        }
    }
    if true {
        v47 := GenS1{F0: uint16(v2), F1: 2147483648, F2: -98}
        var v48 float32 = 1e10
        v49 := v3[""]
        _ = v47
        _ = v48
        _ = v49
    } else {
        for fi50 := 0; fi50 < 6; fi50++ {
            v51 := []int{-9223372036854775808, (^(fi50)), 32768}
            v1[42] = ("" == "hello")
            v4 += (v4 + 66)
            _ = fi50
            _ = v51
        }
    }
    v52 := []int{4294967295, 32767}
    v53 := GenS6{F0: 0}
    var v54 int32 = -1
    delete(v3, (("a" + "hello") + "hello"))
    for fi55 := 0; fi55 < 13; fi55++ {
        switch "\n\t" {
        case "hello":
            delete(v3, ((("\n\t" + "") + (("" + "\n\t") + "\n\t")) + "a"))
        case "":
            fmt.Println(v54, v54)
        case "foo bar":
            var v56 int32 = v54
            _ = v56
        default:
            var v57 int16 = (v41 | v41)
            _ = v57
        }
        fmt.Println(fi55, v2, v2)
        var v58 uint16 = 32767
        _ = fi55
        _ = v58
    }
    var v59 string = "\n\t"
    v60 := GenS1{F0: 42, F1: 2, F2: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v52:", len(v52), v52)
    fmt.Println("v54:", v54)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v52
    _ = v53
    _ = v54
    _ = v59
    _ = v60
}
