package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint64 }
type GenS2 struct { F0 int; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 uint }

func main() {
    var v1 uint16 = 65535
    v2 := make(map[uint32]int)
    var v3 float64 = float64(-2.25)
    var v4 uint = 1
    if false {
        if (!true) {
            var v5 uint8 = 241
            var v6 float32 = 0.1
            _ = v5
            _ = v6
        } else {
            var v7 int8 = -98
            var v8 float64 = v3
            _ = v7
            _ = v8
        }
    }
    var v9 uint32 = 648
    v3 += v3
    var v10 uint = 865
    v11 := GenS1{F0: 999, F1: 0, F2: 4294967295}
    v12 := []int{-81}
    var v13 string = "\n\t"
    switch v13 {
    case "":
        fmt.Println(v9, v10)
    default:
        var v14 uint32 = v9
        _ = v14
    }
    v15 := GenS1{F0: -788, F1: 0, F2: 2}
    if (36 < 100) {
        if false {
            var v16 uint8 = 187
            v1 = 255
            v13 = ("hello" + v13)
            _ = v16
        } else {
            var v17 uint64 = 42
            var v18 uint8 = 171
            _ = v17
            _ = v18
        }
    } else {
        v19 := GenS1{F0: int64(v10), F1: 1, F2: uint64(v9)}
        _ = v19
    }
    switch v13 {
    case "hello":
        switch 4294967295 {
        case -1:
            for fi20 := 0; fi20 < 17; fi20++ {
                fmt.Println(v13, v9)
                v21 := GenS2{F0: fi20, F1: (v1 | 32767), F2: int32(v10)}
                v2[(v9 ^ 42)] = (fi20 | fi20)
                _ = fi20
                _ = v21
            }
        case 1:
            fmt.Println(v3, v9, v10)
        }
    case "\n\t":
        for fi22 := 0; fi22 < 2; fi22++ {
            v3 += (0.0 + 0.1)
            v10 = v10
            v23 := GenS3{F0: 18446744073709551615}
            _ = fi22
            _ = v23
        }
    case "foo bar":
        if false {
            fmt.Println(v13, v1)
            if false {
                for fi24 := 0; fi24 < 16; fi24++ {
                    for fi25 := 0; fi25 < 7; fi25++ {
                        if false {
                            v9 += (v9 & 0)
                        } else {
                            _ = len(v2)
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                            v13 = v13
                            if true {
                                switch ("a" + v13) {
                                case "a":
                                    v12 = append(v12, fi25)
                                }
                                switch (-(fi25)) {
                                case 2147483647:
                                    if true {
                                        v26 := GenS4{F0: v9}
                                        _ = v26
                                    } else {
                                        var v27 bool = true
                                        v28 := GenS5{F0: 62, F1: 1}
                                        _ = v27
                                        _ = v28
                                    }
                                }
                            } else {
                                switch v13 {
                                case "foo bar":
                                    v10 = 42
                                case "a":
                                    var v29 string = (v13 + (v13 + v13))
                                    _ = v29
                                }
                                if true {
                                    switch v13 {
                                    case "hello":
                                        var v30 uint8 = 94
                                        _ = v30
                                    case "a":
                                        switch v13 {
                                        case "a":
                                            v3 -= v3
                                        default:
                                            v31 := GenS5{F0: 4294967295, F1: 618}
                                            _ = v31
                                        }
                                    case "foo bar":
                                        var v32 int32 = 1
                                        _ = v32
                                    }
                                    if (-284 < 42) {
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                        var v33 int32 = 1
                                        _ = v33
                                    } else {
                                        var v34 uint32 = v9
                                        if true {
                                            if ((1e10 * (float32(((v3 + 3.14159) * ((v3 + -2.25) * v3))) + 3.14159)) <= 0.1) {
                                                var v35 float64 = float64(-2.25)
                                                v10 = 944
                                                var v36 uint32 = 42
                                                v37 := GenS6{F0: 204}
                                                var v38 int64 = 9223372036854775807
                                                if true {
                                                    if len(v12) > 0 {
                                                        v12[0] = fi24
                                                    }
                                                }
                                                _ = v35
                                                _ = v36
                                                _ = v37
                                                _ = v38
                                            } else {
                                                var v39 uint16 = (v1 ^ (v1 << ((4294967295) & 63)))
                                                var v40 float64 = v3
                                                _ = v39
                                                _ = v40
                                            }
                                            v34 += v34
                                        } else {
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                            var v41 uint64 = 18446744073709551615
                                            v13 = v13
                                            _ = v41
                                        }
                                        fmt.Println(fi25)
                                        _ = v34
                                    }
                                    _ = len(v2)
                                }
                                fmt.Println(fi25, v3, v10)
                            }
                            v9 += (v9 % ((v9) | 1))
                            var v42 uint16 = 947
                            _ = v42
                        }
                        v10 += v10
                        if false {
                            v43 := GenS4{F0: (v9 << (((v4 << (((v10 >> ((v4) & 63))) & 63))) & 63))}
                            if ((^(v10)) > v4) {
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                            }
                            v10 += v4
                            _ = v43
                        }
                        _ = fi25
                    }
                    v44 := GenS5{F0: uint64(v1), F1: 0}
                    _ = fi24
                    _ = v44
                }
                fmt.Println(v10, v10)
                var v45 int16 = 1000
                _ = v45
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
