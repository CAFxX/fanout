package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 int8 }
func genFn0(p0 int8) int32 {
    switch "foo bar" {
    case "hello":
        for fi6 := 0; fi6 < 20; fi6++ {
            var v7 uint64 = 2
            fmt.Println(v7, p0, p0)
            p0 = (p0 << ((1) & 63))
            _ = fi6
            _ = v7
        }
    case "":
        var v8 int64 = 130
        _ = v8
    default:
        p0 -= (p0 / ((1) | 1))
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        p0 = (p0 % (((p0 ^ (p0 | -1))) | 1))
        _ = fi9
    }
    var v10 int64 = 42
    for fi11 := 0; fi11 < 6; fi11++ {
        if true {
            p0 = p0
        }
        v12 := []int{32768, (fi11 << ((42) & 63)), fi11, -422}
        var v13 string = "hello"
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = v10
    return 1
}
func genFn1(p0 uint, p1 bool) int {
    var v14 int16 = 0
    if p1 {
        var v15 float32 = 1.5
        v16 := GenS1{F0: 0, F1: (v14 + v14), F2: ((65535 < p0) || p1)}
        _ = v15
        _ = v16
    } else {
        var v17 bool = false
        if (!(p1 && (784 >= 18446744073709551615))) {
            var v18 float64 = 0.1
            _ = v18
        } else {
            var v19 uint64 = 1
            if (42 > 563) {
                v20 := []int{2147483647}
                _ = v20
            } else {
                var v21 float64 = 3.14159
                v22 := GenS1{F0: (v19 | 18446744073709551615), F1: (v14 ^ (v14 / (((v14 % ((v14) | 1))) | 1))), F2: v17}
                _ = v21
                _ = v22
            }
            if (!(!(1e10 <= 0.0))) {
                var v23 uint8 = 91
                v24 := []int{}
                var v25 float64 = (3.14159 - 0.0)
                _ = v23
                _ = v24
                _ = v25
            } else {
                for fi26 := 0; fi26 < 18; fi26++ {
                    if (v17 || true) {
                        for fi27 := 0; fi27 < 11; fi27++ {
                            v28 := GenS1{F0: v19, F1: -1, F2: ((p0 % ((404) | 1)) != (p0 * p0))}
                            if ((-2147483648 == 628) || (!v17)) {
                                v14 = v14
                                p0 = 216
                                var v29 float32 = 3.14159
                                _ = v29
                            }
                            _ = fi27
                            _ = v28
                        }
                        p1 = true
                    }
                    _ = fi26
                }
                v17 = true
            }
            _ = v19
        }
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v14
    return -1
}

func main() {
    var v1 int32 = 42
    v2 := make(map[uint32]int)
    v3 := []string{("a" + "\n\t"), "a"}
    var v4 int64 = -9223372036854775808
    var v5 uint = 18446744073709551615
    v1 = v1
    var v30 bool = ((62 > 72) && (!true))
    v5 += 927
    v31 := GenS2{F0: 1}
    fmt.Println(v5)
    v4 += v4
    fmt.Println(v5, v1, v5)
    var v32 int16 = 42
    var v33 bool = (v30 || true)
    v1 += v1
    var v34 string = ""
    _ = len(v2)
    fmt.Println(v5, v34, v5)
    switch (v34 + v34) {
    case "hello":
        for fi35 := 0; fi35 < 17; fi35++ {
            v36 := GenS2{F0: (v5 | (v5 & (v5 | (^(v5)))))}
            _ = fi35
            _ = v36
        }
    case "a":
        var v37 uint = v5
        _ = v37
    default:
        if (((v5 % ((591) | 1)) == v5) && true) {
            v4 = 1
            for fi38 := 0; fi38 < 1; fi38++ {
                for fi39 := 0; fi39 < 7; fi39++ {
                    _ = len(v2)
                    v5 -= (^(v5))
                    if (!(!(((1.5 * 0.0) * 0.1) == 0.0))) {
                        if v30 {
                            v34 = "a"
                            v40 := GenS1{F0: 18446744073709551615, F1: v32, F2: v33}
                            if (!(int8(v1) > 80)) {
                                v32 = v32
                                var v41 uint8 = 0
                                if (!v30) {
                                    v42 := GenS1{F0: 380, F1: v32, F2: (fi39 > (fi39 * int(v32)))}
                                    _ = v42
                                } else {
                                    if (false && v33) {
                                        v33 = false
                                        v41 -= v41
                                        v5 = (v5 % (((v5 & v5)) | 1))
                                        v5 = 993
                                        fmt.Println(fi39, fi38)
                                    }
                                    if v33 {
                                        fmt.Println(v4)
                                        var v43 bool = (v32 <= v32)
                                        _ = v43
                                    }
                                }
                                v44 := GenS1{F0: 951, F1: v32, F2: v33}
                                v30 = true
                                if v30 {
                                    _ = len(v3)
                                    v34 = "a"
                                    fmt.Println(v30, v5, v34)
                                    var v45 int16 = -1
                                    _ = v45
                                }
                                v46 := GenS3{F0: 858, F1: -1}
                                _ = v41
                                _ = v44
                                _ = v46
                            } else {
                                v32 = (v32 * (v32 | int16(v1)))
                            }
                            _ = v40
                        } else {
                            var v47 int8 = -38
                            _ = v47
                        }
                    }
                    v4 -= v4
                    v33 = v33
                    _ = fi39
                }
                var v48 uint = (v5 & v5)
                _ = fi38
                _ = v48
            }
        } else {
            var v49 uint64 = 828
            _ = v49
        }
    }
    delete(v2, 594)
    var v50 bool = v30
    if (!((687 >= -9223372036854775808) || ((255 > 2) || v30))) {
        var v51 string = "a"
        _ = v51
    }
    fmt.Println(v50, v33, v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v50
}
