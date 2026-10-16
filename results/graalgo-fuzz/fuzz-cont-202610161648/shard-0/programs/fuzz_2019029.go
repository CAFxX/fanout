package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int8; F2 uint }
type GenS3 struct { F0 string; F1 int32; F2 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64; F1 int8; F2 bool }
func genFn0(p0 int8, p1 uint8, p2 int16) int64 {
    p2 = -32768
    p0 -= p0
    var v5 int32 = -2147483648
    p0 = 63
    for fi6 := 0; fi6 < 2; fi6++ {
        if true {
            var v7 uint = 1
            p1 -= p1
            var v8 float32 = 1e10
            _ = v7
            _ = v8
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}
func genFn1(p0 uint64, p1 int64, p2 int64) int8 {
    if false {
        fmt.Println(p2)
    } else {
        v9 := GenS1{F0: 1}
        _ = v9
    }
    p1 = (p2 * genFn0(100, 255, 42))
    _ = p0
    _ = p1
    _ = p2
    return 100
}

func main() {
    v1 := []int{-1, 9223372036854775807, 415}
    var v2 uint16 = 65535
    var v3 float64 = (float64((float32((3.14159 / float64(1.5))) + ((3.14159 / 1.0) / 3.14159))) * float64((-2.25 * ((0.1 / (1.5 * 3.14159)) * float32(0.1)))))
    v4 := []uint32{}
    v3 = v3
    _ = len(v1)
    v1 = append(v1, 9223372036854775807)
    var v10 float32 = float32((v3 - v3))
    fmt.Println(v10, v10)
    v10 -= (v10 - -2.25)
    if true {
        switch -1 {
        case -516:
            v11 := GenS1{F0: 931}
            _ = v11
        case 0:
            v12 := GenS2{F0: 4294967295, F1: genFn1(380, -9223372036854775808, genFn0(-128, 0, -1)), F2: 1}
            _ = v12
        }
        switch "foo bar" {
        case "":
            if (-128 <= genFn1(18446744073709551615, -9223372036854775808, 1)) {
                var v13 int32 = 0
                if true {
                    var v14 int16 = 1
                    v2 -= (v2 | v2)
                    v15 := GenS2{F0: 383, F1: -1, F2: 65535}
                    _ = v14
                    _ = v15
                }
                _ = v13
            }
        case "\n\t":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            var v16 string = "\n\t"
            _ = v16
        }
        v17 := GenS1{F0: 977}
        _ = v17
    } else {
        var v18 int32 = 1
        if (1.5 <= 1.5) {
            v19 := GenS1{F0: 2}
            v20 := GenS3{F0: "foo bar", F1: (-(v18)), F2: v18}
            _ = v19
            _ = v20
        } else {
            if false {
                var v21 int16 = 569
                if (291 >= 0) {
                    for fi22 := 0; fi22 < 5; fi22++ {
                        var v23 bool = (true && true)
                        for fi24 := 0; fi24 < 12; fi24++ {
                            if v23 {
                                if v23 {
                                    if (!(("hello" + "hello") != ("" + ("hello" + "")))) {
                                        fmt.Println(v21, v3, v23)
                                        switch "\n\t" {
                                        case "foo bar":
                                            fmt.Println(v21, fi24)
                                        }
                                    }
                                    fmt.Println(v23, fi22, v2)
                                }
                                var v25 int64 = 0
                                _ = v25
                            }
                            _ = fi24
                        }
                        _ = fi22
                        _ = v23
                    }
                }
                _ = v21
            } else {
                var v26 bool = false
                var v27 bool = (!((^(v18)) < (^(v18))))
                _ = v26
                _ = v27
            }
            _ = len(v4)
        }
        switch "foo bar" {
        case "":
            if (v2 > 1) {
                v28 := GenS4{F0: -1}
                _ = v28
            } else {
                if false {
                    v29 := GenS5{F0: 40}
                    _ = v29
                } else {
                    if (4294967295 < 42) {
                        v3 += float64(v10)
                        v18 = int32(v2)
                    } else {
                        if true {
                            switch "hello" {
                            case "":
                                if len(v4) > 0 {
                                    v4[0] = 235
                                }
                            case "\n\t":
                                fmt.Println(v10, v3, v2)
                            }
                            for fi30 := 0; fi30 < 7; fi30++ {
                                if (!false) {
                                    for fi31 := 0; fi31 < 14; fi31++ {
                                        if len(v1) > 0 {
                                            v1[0] = int(v18)
                                        }
                                        var v32 float64 = (3.14159 / v3)
                                        _ = fi31
                                        _ = v32
                                    }
                                    for fi33 := 0; fi33 < 6; fi33++ {
                                        v34 := GenS4{F0: -1}
                                        if len(v4) > 0 {
                                            v4[0] = 1
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = fi33
                                        _ = v34
                                    }
                                } else {
                                    var v35 int64 = genFn0(1, 2, 0)
                                    if len(v4) > 0 {
                                        v4[0] = 1
                                    }
                                    _ = v35
                                }
                                _ = fi30
                            }
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            switch "" {
                            case "\n\t":
                                fmt.Println(v18, v2)
                            case "a":
                                var v36 float32 = v10
                                _ = v36
                            case "foo bar":
                                v3 += (v3 * v3)
                            }
                        }
                    }
                }
            }
        case "\n\t":
            fmt.Println(v10, v18)
        default:
            var v37 uint8 = 158
            _ = v37
        }
        _ = v18
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v38 := GenS6{F0: v2}
    v39 := GenS7{F0: 1, F1: 127, F2: false}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
}
