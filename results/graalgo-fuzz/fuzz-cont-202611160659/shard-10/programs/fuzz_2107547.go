package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int64; F1 int8 }
type GenS7 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS8 struct { F0 int; F1 int32; F2 string }
func genFn0(p0 string, p1 uint16, p2 uint64) int16 {
    v7 := []int{-2147483648, -1}
    p1 = (p1 ^ (^(p1)))
    p1 = p1
    p1 += p1
    var v8 string = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 32767
}
func genFn1(p0 uint64, p1 uint) int64 {
    switch "" {
    case "hello":
        var v9 float64 = 0.1
        _ = v9
    case "":
        var v10 int64 = -552
        _ = v10
    default:
        p0 -= p0
    }
    p1 -= uint(p0)
    if (p0 == (p0 - 817)) {
        fmt.Println(p1)
    }
    p0 = p0
    _ = p0
    _ = p1
    return 600
}

func main() {
    var v1 bool = false
    var v2 float64 = (((3.14159 / (3.14159 / -2.25)) - 3.14159) + -2.25)
    v3 := make(map[uint32]string)
    v4 := make(map[string]bool)
    var v5 uint = 65535
    var v6 uint16 = 542
    fmt.Println(v5, v2)
    v11 := GenS1{F0: 65535, F1: "foo bar"}
    delete(v3, 34)
    v12 := GenS1{F0: 1, F1: "hello"}
    v13 := []int{}
    var v14 float32 = (float32(v2) - 3.14159)
    var v15 int = 873
    for fi16 := 0; fi16 < 20; fi16++ {
        if len(v13) > 0 {
            v13[0] = v15
        }
        _ = fi16
    }
    fmt.Println(v15, v6, v14)
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    v17 := GenS2{F0: 0}
    switch 4294967295 {
    case 2147483647:
        for fi18 := 0; fi18 < 18; fi18++ {
            if v1 {
                v19 := GenS3{F0: "hello", F1: 0}
                fmt.Println(v6, v14, v2)
                if (true && v1) {
                    fmt.Println(v14, v2)
                    v20 := GenS3{F0: ("a" + "a"), F1: 129}
                    for fi21 := 0; fi21 < 15; fi21++ {
                        if ("\n\t" < (("a" + "\n\t") + (("\n\t" + "") + "a"))) {
                            v1 = false
                            v22 := v3[2147483647]
                            _ = v22
                        } else {
                            v23 := GenS4{F0: 128}
                            var v24 uint = 2
                            _ = v23
                            _ = v24
                        }
                        var v25 uint16 = (^(v6))
                        fmt.Println(fi18)
                        v26 := GenS5{F0: ""}
                        var v27 int16 = 0
                        var v28 int = (v15 % ((v15) | 1))
                        if true {
                            var v29 int = -301
                            delete(v3, 2147483648)
                            _ = v29
                        }
                        _ = fi21
                        _ = v25
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    _ = v20
                }
                _ = v19
            }
            v30 := GenS6{F0: 852, F1: 127}
            switch -263 {
            case 9223372036854775807:
                var v31 uint8 = 88
                _ = v31
            case 128:
                if true {
                    v32 := GenS2{F0: fi18}
                    if v1 {
                        v1 = true
                    } else {
                        if v1 {
                            var v33 uint = (v5 % ((v5) | 1))
                            v1 = false
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            _ = v33
                        }
                        if len(v13) > 0 {
                            v13[0] = 130
                        }
                        switch 193 {
                        case -1:
                            v14 = 0.0
                        case -9223372036854775808:
                            for fi34 := 0; fi34 < 8; fi34++ {
                                v35 := GenS6{F0: 1, F1: 1}
                                v15 = fi18
                                v36 := GenS5{F0: ""}
                                _ = fi34
                                _ = v35
                                _ = v36
                            }
                        }
                    }
                    _ = v32
                } else {
                    if (!(v5 > v5)) {
                        v14 = v14
                        v37 := GenS5{F0: ("a" + ("foo bar" + "a"))}
                        _ = v37
                    } else {
                        for fi38 := 0; fi38 < 8; fi38++ {
                            v14 = ((float32(1.5) / v14) - (((v14 + v14) * v14) / v14))
                            v2 = 0.1
                            v13 = append(v13, (fi38 / ((-2147483648) | 1)))
                            _ = fi38
                        }
                    }
                    v39 := GenS5{F0: "hello"}
                    v1 = v1
                    _ = v39
                }
            }
            _ = fi18
            _ = v30
        }
    case 127:
        fmt.Println(v1, v6)
    default:
        v14 = v14
    }
    fmt.Println(v5)
    var v40 int16 = 42
    v2 = v2
    if true {
        v1 = v1
        v41 := GenS7{F0: (v6 * 65535), F1: 2, F2: 42}
        for fi42 := 0; fi42 < 10; fi42++ {
            if (!v1) {
                v43 := GenS2{F0: fi42}
                var v44 bool = (v5 > v5)
                v14 -= v14
                _ = v43
                _ = v44
            } else {
                v45 := GenS8{F0: v15, F1: -190, F2: (("a" + "a") + "foo bar")}
                _ = v45
            }
            _ = fi42
        }
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
