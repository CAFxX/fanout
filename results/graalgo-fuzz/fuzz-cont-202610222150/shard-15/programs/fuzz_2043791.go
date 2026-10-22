package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 uint64; F2 int64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 string; F1 int16 }
type GenS8 struct { F0 int16; F1 int16 }
func genFn0(p0 int8) uint64 {
    var v5 bool = false
    for fi6 := 0; fi6 < 5; fi6++ {
        if v5 {
            v7 := GenS1{F0: 1}
            for fi8 := 0; fi8 < 11; fi8++ {
                switch 32768 {
                case -491:
                    switch (fi6 % ((-861) | 1)) {
                    case 4294967295:
                        v9 := []int{(fi6 / ((fi6) | 1)), fi8, -149, -213}
                        _ = v9
                    case 1:
                        if true {
                            v5 = v5
                            fmt.Println(fi6, fi8)
                            fmt.Println(v5)
                        } else {
                            if v5 {
                                p0 -= (p0 | (p0 * p0))
                                p0 -= p0
                                v10 := GenS1{F0: 730}
                                _ = v10
                            }
                            p0 = p0
                        }
                    case -610:
                        fmt.Println(fi8)
                    default:
                        v11 := []int{}
                        _ = v11
                    }
                case 0:
                    switch fi8 {
                    case -49:
                        v5 = v5
                        switch ("" + "a") {
                        case "hello":
                            v12 := []int{(fi8 >> ((255) & 63)), fi6, -716, 127}
                            _ = v12
                        case "\n\t":
                            if (!(v5 && (v5 || (v5 || false)))) {
                                switch 32767 {
                                case 565:
                                    v13 := []int{(fi6 << ((18446744073709551615) & 63)), fi6, 127}
                                    _ = v13
                                default:
                                    fmt.Println(fi6, fi8)
                                    fmt.Println(v5, p0, v5)
                                }
                                v14 := []int{fi6}
                                if v5 {
                                    if len(v14) > 0 {
                                        v14[0] = fi8
                                    }
                                }
                                _ = v14
                            }
                        }
                        if v5 {
                            fmt.Println(p0, v5)
                            v15 := GenS1{F0: -1}
                            _ = v15
                        }
                    case 127:
                        v16 := GenS1{F0: 42}
                        _ = v16
                    }
                case 28:
                    v17 := GenS2{F0: 405}
                    v18 := GenS2{F0: 255}
                    switch "\n\t" {
                    case "a":
                        v19 := []int{(fi6 << ((18446744073709551615) & 63)), (fi6 + (fi6 | fi8)), fi8, -1}
                        _ = v19
                    }
                    if (2 != 42) {
                        var v20 float64 = (0.0 / (3.14159 / 1.5))
                        var v21 uint = 517
                        _ = v20
                        _ = v21
                    } else {
                        p0 += (-(p0))
                    }
                    _ = v17
                    _ = v18
                }
                _ = fi8
            }
            _ = v7
        } else {
            switch 1 {
            case 1:
                v5 = (!((!v5) || ((float64(float32(3.14159)) + 0.0) == (-2.25 * 3.14159))))
            case -319:
                var v22 uint = 763
                _ = v22
            case 32768:
                p0 = (^(p0))
            }
            v23 := GenS3{F0: 363}
            _ = v23
        }
        switch fi6 {
        case 128:
            for fi24 := 0; fi24 < 20; fi24++ {
                v25 := GenS4{F0: 0}
                var v26 float32 = (3.14159 / 0.1)
                switch ("\n\t" + "foo bar") {
                case "a":
                    v26 += v26
                case "":
                    fmt.Println(v26, v26, p0)
                default:
                    v5 = v5
                }
                _ = fi24
                _ = v25
                _ = v26
            }
        case -9223372036854775808:
            var v27 bool = v5
            _ = v27
        case -729:
            v28 := GenS5{F0: 0, F1: 851, F2: int64(fi6)}
            _ = v28
        }
        p0 += -76
        _ = fi6
    }
    switch (("\n\t" + ("a" + ("" + "hello"))) + "hello") {
    case "hello":
        var v29 uint8 = 1
        _ = v29
    case "\n\t":
        v30 := GenS5{F0: -1, F1: 842, F2: 1}
        _ = v30
    }
    var v31 float64 = 0.1
    var v32 uint64 = 42
    _ = p0
    _ = v31
    _ = v32
    _ = v5
    return v32
}

func main() {
    var v1 uint = 4294967295
    v2 := make(map[string]string)
    v3 := make(map[int]bool)
    var v4 int64 = int64(v1)
    var v33 float64 = 0.0
    var v34 int64 = (^(v4))
    for fi35 := 0; fi35 < 20; fi35++ {
        v34 = (v34 + v4)
        var v36 int32 = -1
        _ = fi35
        _ = v36
    }
    v34 = v4
    if false {
        for fi37 := 0; fi37 < 15; fi37++ {
            fmt.Println(fi37, fi37)
            v1 = v1
            v38 := GenS6{F0: -2147483648}
            _ = fi37
            _ = v38
        }
        fmt.Println(v33)
    }
    var v39 int64 = v34
    for fi40 := 0; fi40 < 4; fi40++ {
        if true {
            v34 = (v4 * (v39 | (-(v39))))
            v41 := []int{32768, fi40, (-(fi40))}
            for fi42 := 0; fi42 < 3; fi42++ {
                v43 := GenS7{F0: ("hello" + "\n\t"), F1: -32768}
                var v44 bool = (!("foo bar" >= "\n\t"))
                fmt.Println(v33, v39, v34)
                if true {
                    fmt.Println(v1, v34, v1)
                } else {
                    v45 := GenS3{F0: 836}
                    _ = v45
                }
                v34 -= (v39 / ((v34) | 1))
                v33 = (v33 + (0.1 * v33))
                _ = fi42
                _ = v43
                _ = v44
            }
            _ = v41
        } else {
            v2["a"] = ""
            var v46 uint = (v1 - v1)
            v47 := GenS6{F0: 230}
            _ = v46
            _ = v47
        }
        var v48 uint32 = 65535
        _ = fi40
        _ = v48
    }
    v33 = ((v33 - (float64((0.1 / (1e10 * -2.25))) + v33)) + (v33 - float64((float32(0.1) - (-2.25 + 0.0)))))
    v33 = 1.5
    v49 := GenS8{F0: 1, F1: -868}
    var v50 uint32 = 2147483647
    if (!(!(!false))) {
        var v51 int16 = 1000
        var v52 float64 = v33
        _ = v51
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v39
    _ = v4
    _ = v49
    _ = v50
}
