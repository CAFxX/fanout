package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint16; F1 bool; F2 uint }
type GenS4 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS5 struct { F0 uint }
func genFn0(p0 int32) int8 {
    p0 -= p0
    v7 := GenS1{F0: -1, F1: -1, F2: 42}
    _ = p0
    _ = v7
    return -128
}
func genFn1(p0 int, p1 uint) bool {
    v8 := GenS2{F0: 32767, F1: "\n\t"}
    p0 = (p0 - (p0 | p0))
    v9 := GenS2{F0: 1, F1: "foo bar"}
    switch (p0 / ((p0) | 1)) {
    case -2147483648:
        v10 := GenS3{F0: 2, F1: (false && false), F2: p1}
        _ = v10
    }
    if (("hello" + "a") <= "a") {
        v11 := []int{(p0 & (p0 ^ (p0 - (p0 / ((127) | 1)))))}
        v12 := GenS2{F0: 0, F1: "hello"}
        var v13 int64 = 9223372036854775807
        _ = v11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return (214 <= 0)
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 0
    var v3 bool = (true || true)
    v4 := []string{"a", "a"}
    v5 := make(map[uint32]int)
    v6 := make(map[int]uint32)
    fmt.Println(v2)
    var v14 uint8 = 128
    v2 = v2
    if false {
        var v15 bool = (4294967295 == 0)
        _ = v15
    } else {
        for fi16 := 0; fi16 < 15; fi16++ {
            if v3 {
                if v3 {
                    if false {
                        v3 = true
                    }
                } else {
                    v5[0] = fi16
                    var v17 float64 = 0.0
                    _ = len(v6)
                    _ = v17
                }
            } else {
                v14 += (v14 ^ v14)
                var v18 float32 = 1.5
                v19 := GenS1{F0: 960, F1: genFn0(-1), F2: 247}
                _ = v18
                _ = v19
            }
            _ = fi16
        }
        if (65535 == 611) {
            v20 := GenS2{F0: 553, F1: "\n\t"}
            if (1 <= 4294967295) {
                v2 = (v2 & (^(v2)))
                if (!(0 < 42)) {
                    v3 = (float64((-2.25 / 0.1)) >= (1e10 / -2.25))
                }
            } else {
                for fi21 := 0; fi21 < 14; fi21++ {
                    if v3 {
                        switch 9223372036854775807 {
                        case -2147483648:
                            switch fi21 {
                            case 9223372036854775807:
                                fmt.Println(v3)
                            case 880:
                                var v22 bool = v3
                                _ = v22
                            case -9223372036854775808:
                                for fi23 := 0; fi23 < 20; fi23++ {
                                    switch 4294967295 {
                                    case 2147483647:
                                        v24 := GenS1{F0: -1, F1: 127, F2: 18446744073709551615}
                                        _ = v24
                                    case 1:
                                        v3 = true
                                    case 42:
                                        v3 = (!(!(65535 > 0)))
                                    default:
                                        v25 := GenS4{F0: 2, F1: 605, F2: 0}
                                        _ = v25
                                    }
                                    var v26 int64 = -9223372036854775808
                                    _ = fi23
                                    _ = v26
                                }
                            }
                        case 9223372036854775807:
                            var v27 string = "\n\t"
                            _ = v27
                        }
                        v3 = v3
                    }
                    for fi28 := 0; fi28 < 9; fi28++ {
                        if true {
                            v14 = (v14 % ((97) | 1))
                        } else {
                            v3 = ((v14 << ((970) & 63)) < v14)
                        }
                        var v29 uint8 = 42
                        _ = fi28
                        _ = v29
                    }
                    v14 -= v14
                    _ = fi21
                }
                for fi30 := 0; fi30 < 8; fi30++ {
                    var v31 bool = false
                    v2 = v2
                    _ = fi30
                    _ = v31
                }
                var v32 int = -1
                _ = v32
            }
            v33 := GenS1{F0: 0, F1: 124, F2: 293}
            _ = v20
            _ = v33
        } else {
            var v34 int32 = -1
            var v35 bool = (!((!true) && v3))
            _ = v34
            _ = v35
        }
        if v3 {
            if (9223372036854775807 != 684) {
                var v36 int = -156
                v3 = true
                v37 := GenS3{F0: 255, F1: v3, F2: 42}
                _ = v36
                _ = v37
            } else {
                var v38 uint16 = 581
                var v39 int32 = 0
                _ = v38
                _ = v39
            }
            fmt.Println(v2, v14)
        } else {
            v4 = append(v4, "hello")
        }
    }
    if (-2.25 == ((1.5 / 3.14159) / 1.0)) {
        v40 := GenS5{F0: 269}
        _ = v40
    }
    if v3 {
        switch 32767 {
        case 42:
            var v41 int16 = 1
            _ = v41
        case 348:
            v14 = 0
        default:
            for fi42 := 0; fi42 < 2; fi42++ {
                var v43 int8 = 60
                _ = fi42
                _ = v43
            }
        }
    }
    var v44 uint32 = 2147483648
    var v45 int32 = -889
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
