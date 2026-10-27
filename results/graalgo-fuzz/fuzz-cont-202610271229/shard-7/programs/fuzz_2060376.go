package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 uint32; F1 int64; F2 int32 }
type GenS6 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS7 struct { F0 string; F1 int }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint32; F1 int64 }
type GenS10 struct { F0 uint8; F1 uint16 }
func genFn0(p0 int32) uint16 {
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 uint16 = 798
        var v8 bool = ((2 != 1) || false)
        _ = fi6
        _ = v7
        _ = v8
    }
    if (("foo bar" + ("" + "hello")) < "") {
        fmt.Println(p0)
        v9 := []int{1, 1}
        _ = v9
    } else {
        var v10 float32 = 3.14159
        var v11 string = "\n\t"
        _ = v10
        _ = v11
    }
    p0 = p0
    if false {
        v12 := GenS1{F0: 2, F1: 127, F2: 876}
        v13 := GenS2{F0: 2147483648}
        _ = v12
        _ = v13
    } else {
        v14 := GenS2{F0: 65535}
        var v15 int16 = 42
        v16 := GenS2{F0: 2147483647}
        _ = v14
        _ = v15
        _ = v16
    }
    _ = p0
    return 2
}

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{2147483648, 766, 784, 2}
    var v3 string = "hello"
    v4 := make(map[int]uint32)
    v5 := make(map[string]uint32)
    v3 = "foo bar"
    v1 = ((v3 + (v3 + ("a" + ("\n\t" + (v1 + (((v3 + v3) + v3) + "hello")))))) + ((v1 + (v1 + (("" + v1) + ("hello" + (((((((v1 + v1) + v3) + ((v3 + ("hello" + ((((((v3 + v1) + ("\n\t" + (("hello" + ("hello" + (v1 + v3))) + v1))) + ("" + "\n\t")) + "\n\t") + "\n\t") + v1))) + (("a" + ("foo bar" + "hello")) + v1))) + (("foo bar" + "") + "")) + v3) + v1) + ("\n\t" + "foo bar")))))) + "hello"))
    fmt.Println(v1)
    fmt.Println(v1, v1)
    var v17 int32 = -822
    v18 := GenS3{F0: -34, F1: 1, F2: 48}
    for fi19 := 0; fi19 < 10; fi19++ {
        switch (v1 + (v3 + ("a" + "hello"))) {
        case "\n\t":
            v4[(fi19 >> ((2) & 63))] = 42
        default:
            fmt.Println(v17)
        }
        _ = fi19
    }
    if false {
        for fi20 := 0; fi20 < 13; fi20++ {
            var v21 int64 = 9223372036854775807
            var v22 string = "a"
            v21 = (v21 * -9223372036854775808)
            _ = fi20
            _ = v21
            _ = v22
        }
        var v23 uint16 = 42
        v23 = (v23 / ((genFn0((v17 + -611))) | 1))
        _ = v23
    }
    v1 = v1
    v24 := GenS4{F0: 42, F1: 155}
    if (v1 > v3) {
        if len(v2) > 0 {
            v2[0] = 65535
        }
        if (-32768 == 1) {
            if (!true) {
                switch v3 {
                case "":
                    fmt.Println(v17)
                case "hello":
                    v25 := GenS4{F0: 521, F1: 0}
                    _ = v25
                case "foo bar":
                    if false {
                        var v26 int8 = -1
                        if false {
                            for fi27 := 0; fi27 < 8; fi27++ {
                                _ = len(v2)
                                v28 := GenS5{F0: 2147483648, F1: 451, F2: 0}
                                _ = fi27
                                _ = v28
                            }
                        }
                        _ = v26
                    } else {
                        for fi29 := 0; fi29 < 17; fi29++ {
                            if false {
                                v30 := GenS5{F0: 415, F1: 9223372036854775807, F2: v17}
                                v31 := GenS6{F0: 42, F1: genFn0((v17 * v17)), F2: genFn0((v17 * (v17 & (v17 >> ((4294967295) & 63)))))}
                                v32 := GenS7{F0: v1, F1: (fi29 / ((fi29) | 1))}
                                _ = v30
                                _ = v31
                                _ = v32
                            }
                            _ = len(v2)
                            for fi33 := 0; fi33 < 1; fi33++ {
                                var v34 int8 = 127
                                _ = fi33
                                _ = v34
                            }
                            _ = fi29
                        }
                    }
                default:
                    v35 := v5["\n\t"]
                    _ = v35
                }
                v17 += (-(v17))
            }
            v1 = v3
            v1 = "hello"
        }
        fmt.Println(v1)
    } else {
        for fi36 := 0; fi36 < 5; fi36++ {
            for fi37 := 0; fi37 < 4; fi37++ {
                var v38 int32 = (-(v17))
                v3 = (v3 + "foo bar")
                if len(v2) > 0 {
                    v2[0] = 860
                }
                _ = fi37
                _ = v38
            }
            _ = fi36
        }
        _ = len(v4)
    }
    var v39 string = v1
    for fi40 := 0; fi40 < 5; fi40++ {
        var v41 int64 = 767
        for fi42 := 0; fi42 < 20; fi42++ {
            switch fi42 {
            case 32767:
                if true {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    fmt.Println(fi40, v41, fi40)
                }
            case 127:
                v41 += v41
            default:
                if false {
                    _ = len(v5)
                } else {
                    var v43 uint = 416
                    v41 += 342
                    var v44 string = (((v1 + v3) + ((("" + v1) + v3) + (v1 + ("foo bar" + v39)))) + v39)
                    _ = v43
                    _ = v44
                }
            }
            _ = fi42
        }
        _ = fi40
        _ = v41
    }
    var v45 uint32 = 571
    var v46 int = -1
    for fi47 := 0; fi47 < 11; fi47++ {
        for fi48 := 0; fi48 < 20; fi48++ {
            var v49 int = (fi48 % ((fi48) | 1))
            switch -593 {
            case 128:
                v50 := v5[(("foo bar" + (("\n\t" + v3) + "\n\t")) + v39)]
                if (!false) {
                    v49 = v46
                    if true {
                        v51 := GenS8{F0: true}
                        var v52 uint = 0
                        _ = v51
                        _ = v52
                    }
                    switch v1 {
                    case "a":
                        var v53 string = ""
                        _ = v53
                    default:
                        fmt.Println(v46, v50)
                    }
                    v39 = (((v3 + v39) + "hello") + ("hello" + v3))
                    if len(v2) > 0 {
                        v2[0] = (^(v45))
                    }
                }
                v54 := GenS9{F0: (^(v45)), F1: 9223372036854775807}
                _ = v50
                _ = v54
            }
            var v55 uint16 = genFn0((v17 % (((v17 - v17)) | 1)))
            v56 := GenS10{F0: 127, F1: v55}
            _ = fi48
            _ = v49
            _ = v55
            _ = v56
        }
        _ = fi47
    }
    for fi57 := 0; fi57 < 1; fi57++ {
        var v58 float32 = 0.0
        v39 = (v1 + (("\n\t" + ("\n\t" + "")) + v39))
        _ = fi57
        _ = v58
    }
    v39 = ("\n\t" + v39)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v39
    _ = v4
    _ = v45
    _ = v46
    _ = v5
}
