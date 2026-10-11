package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 string; F2 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 int; F2 int8 }
type GenS5 struct { F0 uint8; F1 string }
func genFn0(p0 uint32) string {
    v6 := GenS1{F0: 42}
    v7 := []int{}
    for fi8 := 0; fi8 < 14; fi8++ {
        fmt.Println(fi8)
        _ = fi8
    }
    v9 := GenS2{F0: 1, F1: ((("\n\t" + "hello") + "foo bar") + ("a" + ("" + "foo bar"))), F2: 1}
    _ = p0
    _ = v6
    _ = v7
    _ = v9
    return ""
}
func genFn1(p0 uint, p1 int16, p2 uint32) bool {
    var v10 int16 = (p1 - p1)
    p1 = (v10 >> (((p0 / (((p0 % ((uint(p2)) | 1))) | 1))) & 63))
    v11 := GenS1{F0: 716}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return false
}

func main() {
    var v1 float32 = (((3.14159 - (1e10 * (1e10 - 1e10))) * 0.1) + 0.1)
    v2 := make(map[uint32]string)
    v3 := make(map[string]int)
    var v4 string = "a"
    var v5 bool = (true || false)
    v4 = v4
    v4 = ((v4 + (v4 + v4)) + (v4 + v4))
    var v12 uint64 = 4294967295
    for fi13 := 0; fi13 < 20; fi13++ {
        switch 9223372036854775807 {
        case 32768:
            fmt.Println(fi13)
        case 9223372036854775807:
            v14 := GenS3{F0: v4}
            _ = v14
        }
        if (632 >= 2147483647) {
            if ((1 == 10) || (!(!((!false) || v5)))) {
                v15 := GenS4{F0: -9223372036854775808, F1: fi13, F2: 0}
                switch v4 {
                case "a":
                    if ("foo bar" <= v4) {
                        for fi16 := 0; fi16 < 10; fi16++ {
                            var v17 uint16 = 65535
                            _ = fi16
                            _ = v17
                        }
                        var v18 int16 = 1
                        var v19 string = ""
                        _ = v18
                        _ = v19
                    } else {
                        v4 = "hello"
                        var v20 int32 = -2147483648
                        _ = v20
                    }
                case "foo bar":
                    switch "hello" {
                    case "\n\t":
                        _ = len(v3)
                    case "hello":
                        var v21 int8 = 127
                        _ = v21
                    default:
                        for fi22 := 0; fi22 < 15; fi22++ {
                            v4 = ((v4 + "\n\t") + "hello")
                            _ = fi22
                        }
                    }
                }
                _ = v15
            } else {
                _ = len(v2)
                _ = len(v2)
                if v5 {
                    v1 = 3.14159
                    v12 = v12
                }
            }
        } else {
            v12 = (v12 - (v12 / ((v12) | 1)))
            var v23 uint8 = 128
            if false {
                var v24 uint8 = (^(v23))
                if v5 {
                    v25 := []int{fi13, (fi13 + (fi13 * fi13))}
                    _ = v25
                } else {
                    v12 = v12
                }
                v26 := []int{fi13, fi13, 4294967295, fi13}
                _ = v24
                _ = v26
            } else {
                v23 -= (v23 / ((42) | 1))
                switch 42 {
                case 32767:
                    for fi27 := 0; fi27 < 5; fi27++ {
                        if (1 > -940) {
                            v1 = v1
                            v3[((v4 + v4) + v4)] = (fi13 - (-(fi27)))
                            var v28 string = v4
                            v28 = "hello"
                            v29 := GenS4{F0: 1, F1: fi13, F2: 0}
                            _ = v28
                            _ = v29
                        }
                        if v5 {
                            v12 -= (^(v12))
                            switch 32768 {
                            case -1:
                                v1 = v1
                            case 32768:
                                var v30 int16 = -1
                                _ = v30
                            }
                            delete(v3, (v4 + ((("\n\t" + "a") + v4) + "foo bar")))
                        }
                        var v31 string = v4
                        _ = fi27
                        _ = v31
                    }
                case 9223372036854775807:
                    v32 := []int{(fi13 ^ (-(fi13))), (^(fi13)), 42, fi13}
                    _ = v32
                }
                if (-95 > 1) {
                    if false {
                        var v33 int32 = -2147483648
                        v33 = (v33 << ((4294967295) & 63))
                        switch v4 {
                        case "":
                            v34 := []int{}
                            _ = v34
                        case "\n\t":
                            if true {
                                if true {
                                    var v35 int32 = (v33 - v33)
                                    _ = v35
                                } else {
                                    var v36 uint = uint(v23)
                                    _ = v36
                                }
                            }
                        case "foo bar":
                            if v5 {
                                var v37 uint = 458
                                v33 = (^(v33))
                                _ = v37
                            }
                        }
                        _ = v33
                    }
                    v38 := GenS2{F0: 987, F1: v4, F2: -198}
                    if (42 > (v12 ^ 18446744073709551615)) {
                        v1 = v1
                    } else {
                        fmt.Println(v4, v12)
                        var v39 bool = (!(34 > 0))
                        v4 = (v4 + "hello")
                        _ = v39
                    }
                    _ = v38
                }
            }
            _ = v23
        }
        _ = fi13
    }
    for fi40 := 0; fi40 < 6; fi40++ {
        if ((false && true) || (!(1 <= -1))) {
            for fi41 := 0; fi41 < 18; fi41++ {
                var v42 uint8 = 0
                switch (v4 + v4) {
                case "\n\t":
                    v43 := GenS5{F0: (v42 + v42), F1: v4}
                    _ = v43
                default:
                    var v44 uint = 508
                    _ = v44
                }
                if v5 {
                    if (v42 <= 14) {
                        _ = len(v3)
                    } else {
                        var v45 float64 = 1e10
                        v12 = v12
                        _ = v45
                    }
                }
                v46 := GenS3{F0: v4}
                fmt.Println(v42)
                _ = fi41
                _ = v42
                _ = v46
            }
            delete(v2, 851)
        }
        _ = fi40
    }
    v1 += float32(1e10)
    var v47 uint = 42
    if (true && false) {
        v48 := []int{32768, 451, -9223372036854775808, 42}
        var v49 bool = (65535 < 2)
        _ = v48
        _ = v49
    }
    for fi50 := 0; fi50 < 16; fi50++ {
        switch fi50 {
        case 4294967295:
            fmt.Println(fi50, v5, v47)
        }
        _ = fi50
    }
    v51 := GenS4{F0: -9223372036854775808, F1: 0, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v51
}
