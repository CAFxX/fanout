package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string }

func main() {
    var v1 int8 = -73
    v2 := []uint32{42}
    var v3 int32 = 2147483647
    var v4 string = ""
    if ("foo bar" >= v4) {
        v1 += v1
        if false {
            v3 -= v3
            var v5 float32 = -2.25
            var v6 uint32 = 42
            _ = v5
            _ = v6
        } else {
            var v7 int64 = -694
            if true {
                v7 += 9223372036854775807
                if (((32767 < 2) && false) && true) {
                    for fi8 := 0; fi8 < 6; fi8++ {
                        var v9 int16 = 715
                        v9 = -1
                        switch (-(fi8)) {
                        case 128:
                            v10 := GenS1{F0: v1, F1: v4, F2: 65535}
                            _ = v10
                        case 9223372036854775807:
                            v9 += 1000
                        default:
                            if false {
                                v11 := GenS2{F0: 2, F1: v1}
                                _ = v11
                            } else {
                                v12 := GenS2{F0: 745, F1: (v1 % ((v1) | 1))}
                                _ = v12
                            }
                        }
                        _ = fi8
                        _ = v9
                    }
                    var v13 int8 = 100
                    _ = v13
                } else {
                    var v14 int64 = (v7 ^ -9223372036854775808)
                    if true {
                        v2 = append(v2, 2)
                    }
                    if true {
                        var v15 int16 = int16(v7)
                        for fi16 := 0; fi16 < 10; fi16++ {
                            switch "foo bar" {
                            case "\n\t":
                                var v17 uint = 0
                                _ = v17
                            default:
                                if (fi16 == fi16) {
                                    var v18 uint64 = 2
                                    fmt.Println(fi16)
                                    v19 := GenS3{F0: (v18 & 42), F1: true, F2: 440}
                                    _ = v18
                                    _ = v19
                                } else {
                                    var v20 int64 = v7
                                    for fi21 := 0; fi21 < 3; fi21++ {
                                        v3 -= v3
                                        _ = fi21
                                    }
                                    for fi22 := 0; fi22 < 3; fi22++ {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        var v23 string = ("\n\t" + ("a" + v4))
                                        _ = fi22
                                        _ = v23
                                    }
                                    _ = v20
                                }
                            }
                            _ = fi16
                        }
                        _ = v15
                    } else {
                        var v24 string = "a"
                        var v25 bool = (1000 <= -1)
                        _ = v24
                        _ = v25
                    }
                    _ = v14
                }
            } else {
                var v26 int64 = -1
                _ = v26
            }
            _ = v7
        }
    }
    if len(v2) > 0 {
        v2[0] = 0
    }
    _ = len(v2)
    var v27 uint64 = 896
    for fi28 := 0; fi28 < 5; fi28++ {
        var v29 string = v4
        if ((("a" + "\n\t") > "\n\t") && false) {
            fmt.Println(v27, v27)
            if true {
                v30 := GenS4{F0: 42}
                v27 -= (v27 & v27)
                var v31 uint8 = 1
                _ = v30
                _ = v31
            } else {
                _ = len(v2)
            }
        }
        _ = fi28
        _ = v29
    }
    fmt.Println(v3, v1)
    switch 128 {
    case -521:
        v3 += 497
    case 128:
        _ = len(v2)
    default:
        v32 := GenS2{F0: v27, F1: (v1 << ((2) & 63))}
        _ = v32
    }
    v3 = v3
    v33 := GenS1{F0: 19, F1: "", F2: 65535}
    if ((1e10 < 1e10) || (v4 < (v4 + "foo bar"))) {
        var v34 float32 = ((-2.25 - (3.14159 * 1e10)) * 3.14159)
        v3 -= -2147483648
        _ = v34
    }
    var v35 float32 = 0.1
    var v36 int16 = 0
    v36 += (^(v36))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch v4 {
    case "\n\t":
        for fi37 := 0; fi37 < 5; fi37++ {
            _ = len(v2)
            for fi38 := 0; fi38 < 1; fi38++ {
                v39 := GenS5{F0: (v4 + "foo bar")}
                _ = fi38
                _ = v39
            }
            var v40 uint8 = 128
            _ = fi37
            _ = v40
        }
    default:
        v27 += 18446744073709551615
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v35
    _ = v36
    _ = v4
}
