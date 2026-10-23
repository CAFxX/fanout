package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int8 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int; F1 bool }
func genFn0(p0 int16, p1 int8) int64 {
    switch 0 {
    case 32767:
        var v6 string = "foo bar"
        _ = v6
    case 32768:
        if true {
            for fi7 := 0; fi7 < 5; fi7++ {
                fmt.Println(p0)
                v8 := []int{fi7, (fi7 >> ((396) & 63)), (^(fi7))}
                _ = fi7
                _ = v8
            }
        } else {
            fmt.Println(p1)
            v9 := GenS1{F0: "a", F1: 1000, F2: p1}
            _ = v9
        }
    case -2147483648:
        p0 = p0
    }
    var v10 uint32 = 4294967295
    if true {
        var v11 int64 = 9223372036854775807
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v10
    return -1
}
func genFn1(p0 string, p1 uint64) uint {
    v12 := []int{32767}
    var v13 int32 = 370
    for fi14 := 0; fi14 < 13; fi14++ {
        for fi15 := 0; fi15 < 15; fi15++ {
            v16 := GenS1{F0: ((("foo bar" + (p0 + p0)) + p0) + p0), F1: -1, F2: -103}
            fmt.Println(fi14)
            var v17 uint16 = 2
            _ = fi15
            _ = v16
            _ = v17
        }
        switch ("foo bar" + p0) {
        case "foo bar":
            v18 := GenS2{F0: (v13 ^ v13), F1: -128}
            _ = v18
        }
        _ = fi14
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 18446744073709551615
}

func main() {
    v1 := []string{"foo bar"}
    var v2 int = 128
    v3 := make(map[uint32]string)
    v4 := make(map[int]bool)
    var v5 float32 = 0.1
    if false {
        v2 = (v2 << ((42) & 63))
    }
    switch v2 {
    case 9223372036854775807:
        for fi19 := 0; fi19 < 12; fi19++ {
            for fi20 := 0; fi20 < 2; fi20++ {
                fmt.Println(v5, fi20)
                if true {
                    var v21 uint64 = 18446744073709551615
                    if false {
                        var v22 uint32 = 385
                        if (-1 == 0) {
                            if true {
                                switch fi19 {
                                case -9223372036854775808:
                                    v1 = append(v1, "")
                                case -496:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case -2147483648:
                                    var v23 float64 = 1e10
                                    _ = v23
                                default:
                                    if ((2 != 2) || (-1 >= 0)) {
                                        var v24 uint16 = 42
                                        _ = v24
                                    }
                                }
                                var v25 int64 = -847
                                fmt.Println(v2)
                                _ = v25
                            }
                            var v26 uint16 = 2
                            v27 := GenS1{F0: "", F1: 208, F2: 51}
                            _ = v26
                            _ = v27
                        }
                        if false {
                            var v28 uint32 = (v22 * v22)
                            v22 += 2147483647
                            _ = v28
                        }
                        _ = v22
                    } else {
                        var v29 bool = (18446744073709551615 != 18446744073709551615)
                        fmt.Println(v2, fi20, v2)
                        _ = v29
                    }
                    fmt.Println(v21, v2, fi20)
                    _ = v21
                } else {
                    fmt.Println(fi20, fi20)
                    var v30 float64 = float64(1.5)
                    if false {
                        fmt.Println(v30, v5)
                        var v31 uint64 = 0
                        _ = v31
                    }
                    var v32 uint64 = 2
                    v33 := GenS1{F0: ("" + (("" + (("a" + ("\n\t" + (("a" + "foo bar") + ""))) + "")) + "")), F1: 1, F2: 100}
                    v2 += (^(v2))
                    _ = v30
                    _ = v32
                    _ = v33
                }
                fmt.Println(fi19, v5, fi19)
                fmt.Println(v2)
                var v34 int16 = -652
                _ = fi20
                _ = v34
            }
            _ = fi19
        }
    case 128:
        var v35 uint64 = 268
        _ = v35
    }
    fmt.Println(v2, v5)
    for fi36 := 0; fi36 < 9; fi36++ {
        v1 = append(v1, ("hello" + ""))
        _ = fi36
    }
    if false {
        var v37 uint64 = 487
        var v38 int16 = 32767
        _ = v37
        _ = v38
    }
    var v39 string = "\n\t"
    v40 := GenS3{F0: v2, F1: (!(!true))}
    fmt.Println(v39, v2)
    var v41 int8 = 0
    v42 := GenS2{F0: -1, F1: (v41 / ((int8(v2)) | 1))}
    var v43 uint64 = 0
    v41 = v41
    var v44 float32 = 1e10
    v41 = (^(v41))
    fmt.Println(v41, v2, v41)
    if (-1 >= 204) {
        for fi45 := 0; fi45 < 20; fi45++ {
            v2 -= (v2 * 2147483647)
            _ = fi45
        }
        if ((v39 != ((v39 + "hello") + (v39 + v39))) && false) {
            var v46 uint8 = 42
            v47 := GenS2{F0: -2147483648, F1: (v41 >> ((73) & 63))}
            _ = v46
            _ = v47
        }
        var v48 int8 = 100
        _ = v48
    }
    v49 := GenS2{F0: 42, F1: v41}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v49
    _ = v5
}
