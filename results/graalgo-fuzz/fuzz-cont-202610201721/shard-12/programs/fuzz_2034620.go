package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 uint64; F1 int64 }
type GenS6 struct { F0 bool; F1 uint; F2 int32 }
func genFn0(p0 uint32) string {
    var v5 uint64 = 42
    fmt.Println(p0)
    _ = p0
    _ = v5
    return "\n\t"
}
func genFn1(p0 int64) int8 {
    v6 := GenS1{F0: 32767, F1: 42}
    if false {
        p0 = (p0 - (p0 * 0))
        p0 -= p0
        if (("hello" < (("\n\t" + ((("hello" + "\n\t") + ("hello" + "a")) + ("" + ""))) + ("a" + ("a" + "\n\t")))) && true) {
            var v7 float32 = -2.25
            v8 := []int{144}
            _ = len(v8)
            _ = v7
            _ = v8
        }
    } else {
        v9 := GenS2{F0: 65535, F1: (p0 ^ p0)}
        if false {
            v10 := GenS3{F0: 18446744073709551615, F1: 339, F2: 781}
            if false {
                for fi11 := 0; fi11 < 19; fi11++ {
                    var v12 int64 = p0
                    var v13 uint8 = 1
                    fmt.Println(fi11, fi11, v13)
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                fmt.Println(p0)
            }
            v14 := GenS1{F0: 65535, F1: p0}
            _ = v10
            _ = v14
        }
        if true {
            var v15 int64 = p0
            _ = v15
        } else {
            for fi16 := 0; fi16 < 7; fi16++ {
                var v17 uint32 = 4294967295
                v18 := GenS1{F0: 681, F1: -1}
                _ = fi16
                _ = v17
                _ = v18
            }
            var v19 float64 = 0.0
            _ = v19
        }
        _ = v9
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 int8 = -70
        _ = fi20
        _ = v21
    }
    _ = p0
    _ = v6
    return -128
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 127
    v3 := []string{"foo bar", ("a" + "\n\t"), "\n\t", ("hello" + (("\n\t" + "foo bar") + "\n\t"))}
    var v4 float64 = (-2.25 + float64(-2.25))
    v2 = (v2 % (((v2 & -128)) | 1))
    var v22 int32 = 42
    v23 := GenS4{F0: 636, F1: -452}
    if (42 >= 1) {
        v22 -= -2147483648
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        switch fi24 {
        case 32768:
            fmt.Println(v2)
        }
        for fi25 := 0; fi25 < 7; fi25++ {
            v26 := GenS2{F0: 65535, F1: -9223372036854775808}
            var v27 float64 = float64(0.0)
            v4 = v27
            var v28 int = fi24
            fmt.Println(fi24, v27)
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
        }
        var v29 float64 = float64(0.1)
        _ = fi24
        _ = v29
    }
    v2 += (v2 % ((genFn1(42)) | 1))
    v30 := GenS5{F0: 42, F1: 9223372036854775807}
    v22 -= v22
    var v31 float32 = 0.1
    for fi32 := 0; fi32 < 14; fi32++ {
        if true {
            fmt.Println(v4, fi32, v4)
            for fi33 := 0; fi33 < 16; fi33++ {
                fmt.Println(v22, fi33)
                switch 32767 {
                case 127:
                    var v34 uint8 = 116
                    _ = v34
                default:
                    if (-1 <= 0) {
                        fmt.Println(fi32)
                        switch fi33 {
                        case 0:
                            switch "a" {
                            case "\n\t":
                                var v35 float64 = v4
                                _ = v35
                            case "foo bar":
                                v36 := GenS4{F0: fi33, F1: v22}
                                _ = v36
                            default:
                                if true {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = len(v1)
                                } else {
                                    _ = len(v1)
                                    if (1 != 32767) {
                                        v22 = -542
                                    }
                                }
                            }
                        case 739:
                            fmt.Println(v31, v22)
                        case 32768:
                            v37 := GenS2{F0: 450, F1: 0}
                            _ = v37
                        default:
                            v31 = 0.0
                        }
                        if false {
                            var v38 uint32 = 1
                            _ = v38
                        }
                    }
                }
                var v39 uint8 = 34
                var v40 string = ""
                _ = fi33
                _ = v39
                _ = v40
            }
        } else {
            v41 := GenS6{F0: true, F1: 18446744073709551615, F2: v22}
            fmt.Println(v22, v2, v31)
            if (fi32 >= (fi32 % ((fi32) | 1))) {
                var v42 bool = false
                _ = v42
            }
            _ = v41
        }
        _ = fi32
    }
    if false {
        v2 = (^(v2))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
