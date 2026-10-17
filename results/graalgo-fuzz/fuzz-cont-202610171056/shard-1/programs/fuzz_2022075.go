package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint32; F2 string }
type GenS3 struct { F0 uint16; F1 int8; F2 string }
type GenS4 struct { F0 int32; F1 uint; F2 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS8 struct { F0 int16; F1 string; F2 string }
func genFn0() int16 {
    v7 := GenS1{F0: 2147483648}
    var v8 uint16 = 708
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := []int{(fi9 << ((859) & 63)), fi9, fi9}
        _ = fi9
        _ = v10
    }
    if true {
        v8 = 563
        var v11 int32 = 0
        _ = v11
    }
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 string) int32 {
    v12 := GenS1{F0: 2}
    v13 := GenS1{F0: 76}
    _ = p0
    _ = v12
    _ = v13
    return 0
}

func main() {
    var v1 int8 = -128
    var v2 bool = (42 < v1)
    var v3 uint64 = 148
    v4 := make(map[int]string)
    v5 := []int{127, 32768}
    var v6 string = ("hello" + "a")
    var v14 int16 = -435
    v6 = v6
    _ = len(v5)
    if v2 {
        var v15 uint16 = 2
        _ = v15
    } else {
        fmt.Println(v1)
    }
    var v16 string = ("foo bar" + "a")
    fmt.Println(v3)
    if false {
        var v17 uint16 = 253
        fmt.Println(v3, v1)
        if (true && ((!v2) || v2)) {
            var v18 string = v16
            _ = v18
        } else {
            var v19 int8 = (v1 ^ v1)
            for fi20 := 0; fi20 < 20; fi20++ {
                if ((^(v17)) < 2) {
                    if (v1 <= (v1 - v19)) {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                } else {
                    for fi21 := 0; fi21 < 18; fi21++ {
                        var v22 uint32 = 65535
                        v3 -= 42
                        v23 := GenS2{F0: 1, F1: 92, F2: v16}
                        var v24 int32 = genFn1(v6)
                        _ = fi21
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    if v2 {
                        v6 = v6
                        v16 = ((v6 + v16) + v6)
                        var v25 int = fi20
                        _ = v25
                    }
                    if ((156 == 8) || ((v3 % ((42) | 1)) > v3)) {
                        delete(v4, (fi20 ^ fi20))
                    }
                }
                _ = fi20
            }
            switch v16 {
            case "a":
                fmt.Println(v16, v17, v16)
            }
            _ = v19
        }
        _ = v17
    }
    var v26 uint64 = 0
    if len(v5) > 0 {
        v5[0] = 42
    }
    if ((32768 != -9223372036854775808) || v2) {
        if (!(9223372036854775807 <= 128)) {
            if false {
                switch 830 {
                case -9223372036854775808:
                    var v27 bool = v2
                    _ = v27
                case -1:
                    if (int64(v1) >= 9223372036854775807) {
                        if (float32(0.0) >= 3.14159) {
                            if len(v5) > 0 {
                                v5[0] = 42
                            }
                            switch 128 {
                            case -2147483648:
                                var v28 bool = false
                                _ = v28
                            case 53:
                                v29 := GenS3{F0: 1, F1: (v1 >> ((1) & 63)), F2: v6}
                                _ = v29
                            case 0:
                                var v30 uint8 = 0
                                _ = v30
                            }
                            delete(v4, int(v1))
                        }
                        var v31 int32 = 425
                        v32 := GenS4{F0: int32(v26), F1: 524, F2: 62}
                        _ = v31
                        _ = v32
                    }
                case 491:
                    if false {
                        var v33 uint = 65535
                        v26 = (v3 | (v3 / ((954) | 1)))
                        for fi34 := 0; fi34 < 5; fi34++ {
                            switch 1 {
                            case 32768:
                                v33 += (v33 + 18446744073709551615)
                            case 0:
                                v1 = (v1 & (^(v1)))
                            case 127:
                                if v2 {
                                    var v35 uint16 = 255
                                    v36 := v4[4294967295]
                                    _ = v35
                                    _ = v36
                                }
                            }
                            if (!true) {
                                v37 := v4[2147483647]
                                v33 = v33
                                switch fi34 {
                                case 1:
                                    v4[fi34] = v16
                                case -1:
                                    v38 := GenS5{F0: ((((((("\n\t" + "") + v37) + "hello") + v6) + "a") + "") + v37)}
                                    _ = v38
                                case -501:
                                    var v39 uint = (v33 / (((v33 ^ (v33 >> ((v33) & 63)))) | 1))
                                    _ = v39
                                default:
                                    fmt.Println(v33, fi34)
                                }
                                _ = v37
                            }
                            for fi40 := 0; fi40 < 3; fi40++ {
                                v41 := GenS6{F0: 890}
                                _ = fi40
                                _ = v41
                            }
                            _ = fi34
                        }
                        _ = v33
                    }
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                for fi42 := 0; fi42 < 7; fi42++ {
                    for fi43 := 0; fi43 < 10; fi43++ {
                        v44 := GenS7{F0: 2, F1: 165, F2: -9223372036854775808}
                        v45 := GenS8{F0: (v14 ^ 0), F1: "\n\t", F2: ""}
                        v16 = (v6 + (v16 + "hello"))
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                    _ = fi42
                }
            }
            v26 += (v3 - (v26 / (((v26 % ((2) | 1))) | 1)))
            v2 = (!v2)
        } else {
            v4[128] = "a"
            for fi46 := 0; fi46 < 9; fi46++ {
                delete(v4, (fi46 ^ (fi46 ^ fi46)))
                var v47 uint32 = 0
                var v48 int64 = -1
                _ = fi46
                _ = v47
                _ = v48
            }
            var v49 int32 = -2147483648
            _ = v49
        }
    } else {
        delete(v4, -130)
        v6 = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
