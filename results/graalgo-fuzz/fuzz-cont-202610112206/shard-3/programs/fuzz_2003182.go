package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint; F1 int16; F2 uint16 }
type GenS5 struct { F0 int8; F1 int }
type GenS6 struct { F0 int; F1 uint32; F2 int16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint8 }

func main() {
    var v1 string = (("a" + "") + "")
    var v2 bool = true
    var v3 uint32 = 0
    var v4 uint64 = 4294967295
    var v5 int = -9223372036854775808
    v6 := []int{}
    v1 = "\n\t"
    v7 := GenS1{F0: 42, F1: 409}
    var v8 uint32 = 65535
    var v9 uint32 = (v3 + v8)
    var v10 uint64 = v4
    v2 = v2
    v6 = append(v6, -9223372036854775808)
    v11 := GenS2{F0: (v5 << ((4294967295) & 63)), F1: 100}
    if ((^(v10)) > uint64(v9)) {
        if ((0.1 >= float32(0.0)) || v2) {
            var v12 uint8 = 127
            _ = v12
        } else {
            if (false && (v10 >= v4)) {
                v13 := GenS1{F0: 42, F1: 994}
                var v14 uint32 = (v9 & 2147483648)
                _ = v13
                _ = v14
            }
            v10 = 2
            var v15 uint32 = v9
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        v6 = append(v6, 201)
        _ = fi16
    }
    v17 := GenS2{F0: 1, F1: 127}
    var v18 uint = 505
    switch v5 {
    case 2147483647:
        v19 := GenS2{F0: 127, F1: 43}
        _ = v19
    case 76:
        fmt.Println(v2, v3, v18)
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        v4 -= v10
        for fi21 := 0; fi21 < 7; fi21++ {
            if false {
                if v2 {
                    if len(v6) > 0 {
                        v6[0] = (fi20 * (fi21 + fi20))
                    }
                    switch v1 {
                    case "":
                        if (v4 <= (v4 % ((uint64(v8)) | 1))) {
                            var v22 int16 = 42
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            _ = v22
                        } else {
                            var v23 uint64 = uint64(v9)
                            v24 := GenS3{F0: 100, F1: v10, F2: 896}
                            if (2 > 255) {
                                if v2 {
                                    fmt.Println(v9)
                                    v10 += v10
                                    v9 = v9
                                } else {
                                    switch (v5 % ((42) | 1)) {
                                    case -9223372036854775808:
                                        v25 := GenS1{F0: 42, F1: (^(v18))}
                                        _ = v25
                                    case -321:
                                        v6 = append(v6, v5)
                                    default:
                                        v26 := GenS4{F0: 4294967295, F1: 905, F2: 42}
                                        _ = v26
                                    }
                                    if len(v6) > 0 {
                                        v6[0] = fi21
                                    }
                                    fmt.Println(v2, v5, v8)
                                }
                                if v2 {
                                    _ = len(v6)
                                    _ = len(v6)
                                }
                            } else {
                                var v27 bool = (!v2)
                                fmt.Println(v2, v9, v8)
                                _ = v27
                            }
                            _ = v23
                            _ = v24
                        }
                    }
                } else {
                    switch v5 {
                    case -924:
                        if (((1.5 * 0.0) >= (0.1 / 1.5)) && v2) {
                            v28 := GenS1{F0: 1, F1: v18}
                            v29 := GenS2{F0: fi20, F1: 6}
                            v10 += v4
                            var v30 uint = (v18 ^ (v18 - v18))
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                    case 443:
                        if false {
                            var v31 uint8 = 255
                            var v32 int64 = -832
                            _ = v31
                            _ = v32
                        }
                    }
                    v9 -= (v8 + v8)
                    var v33 uint16 = 0
                    _ = v33
                }
            } else {
                v34 := GenS5{F0: 127, F1: -724}
                fmt.Println(fi20)
                if (!v2) {
                    if (1000 > 1) {
                        v3 += v3
                        v6 = append(v6, 127)
                    } else {
                        switch fi20 {
                        case 418:
                            v6 = append(v6, (fi20 & 128))
                        case 128:
                            v35 := GenS6{F0: (v5 + fi21), F1: v3, F2: 860}
                            _ = v35
                        default:
                            v6 = append(v6, fi21)
                        }
                        if true {
                            v36 := GenS7{F0: -971}
                            if false {
                                switch (fi20 ^ fi20) {
                                case 1:
                                    switch ((v1 + (v1 + v1)) + v1) {
                                    case "\n\t":
                                        fmt.Println(v18, v4)
                                    }
                                case 9223372036854775807:
                                    v37 := GenS8{F0: 255}
                                    _ = v37
                                case -755:
                                    v1 = ""
                                default:
                                    var v38 int16 = -442
                                    _ = v38
                                }
                                if len(v6) > 0 {
                                    v6[0] = 2147483647
                                }
                            }
                            var v39 uint32 = v8
                            v5 = (v5 + (fi20 + (v5 ^ 32767)))
                            _ = v36
                            _ = v39
                        } else {
                            v40 := GenS4{F0: 825, F1: 32767, F2: 0}
                            v2 = (1 > 0)
                            fmt.Println(v2, v10, v18)
                            _ = v40
                        }
                    }
                } else {
                    switch v1 {
                    case "a":
                        var v41 string = "hello"
                        _ = v41
                    }
                    if ((97 != 42) || (v2 || true)) {
                        var v42 uint8 = uint8(v4)
                        if len(v6) > 0 {
                            v6[0] = v5
                        }
                        var v43 uint8 = 127
                        fmt.Println(v3, fi20)
                        if true {
                            if len(v6) > 0 {
                                v6[0] = v5
                            }
                            v5 += (fi20 | -9223372036854775808)
                        } else {
                            _ = len(v6)
                            fmt.Println(v5)
                            var v44 int32 = 25
                            _ = v44
                        }
                        _ = v42
                        _ = v43
                    }
                }
                _ = v34
            }
            _ = fi21
        }
        _ = fi20
    }
    var v45 int32 = 1
    var v46 uint64 = (v10 << ((v18) & 63))
    for fi47 := 0; fi47 < 20; fi47++ {
        if len(v6) > 0 {
            v6[0] = 128
        }
        _ = fi47
    }
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
