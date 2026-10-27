package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint; F1 uint; F2 uint }
type GenS3 struct { F0 uint64; F1 uint32; F2 int }
type GenS4 struct { F0 int; F1 uint8; F2 int8 }
type GenS5 struct { F0 bool; F1 int32 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int16; F1 uint64 }
func genFn0(p0 uint32, p1 bool, p2 string) uint8 {
    if (153 >= 42) {
        v6 := []int{}
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := GenS1{F0: 965, F1: 42}
            v9 := GenS1{F0: -1, F1: 923}
            for fi10 := 0; fi10 < 13; fi10++ {
                p2 = ((p2 + p2) + "a")
                var v11 int8 = 109
                _ = fi10
                _ = v11
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = v6
    }
    if (p0 != 910) {
        switch "foo bar" {
        case "a":
            v12 := GenS1{F0: 1000, F1: -1}
            _ = v12
        case "hello":
            var v13 int = -586
            _ = v13
        case "foo bar":
            var v14 float64 = (-2.25 * 1.5)
            _ = v14
        }
        for fi15 := 0; fi15 < 15; fi15++ {
            var v16 int32 = 491
            _ = fi15
            _ = v16
        }
        var v17 bool = true
        _ = v17
    } else {
        v18 := GenS2{F0: 65535, F1: 4294967295, F2: uint(p0)}
        switch -1 {
        case 42:
            if (1 <= -1) {
                var v19 uint8 = 100
                _ = v19
            } else {
                if (true && (255 != 103)) {
                    for fi20 := 0; fi20 < 11; fi20++ {
                        if (!(p1 || ((p1 || (1 == -115)) && (127 == 1)))) {
                            v21 := []int{(fi20 - fi20), -1, fi20, (fi20 / ((fi20) | 1))}
                            v22 := GenS3{F0: 18446744073709551615, F1: p0, F2: fi20}
                            for fi23 := 0; fi23 < 7; fi23++ {
                                fmt.Println(p2, p0)
                                var v24 int64 = 165
                                if len(v21) >= 2 {
                                    v21 = v21[1:len(v21)]
                                }
                                v25 := GenS4{F0: fi20, F1: 1, F2: -1}
                                v24 -= (v24 % (((v24 | (v24 | (^(v24))))) | 1))
                                _ = len(v21)
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                            _ = v21
                            _ = v22
                        } else {
                            v26 := []int{(fi20 >> ((0) & 63)), 1}
                            if p1 {
                                v27 := GenS5{F0: (!false), F1: 42}
                                fmt.Println(p1)
                                _ = v27
                            }
                            _ = v26
                        }
                        if ((0.0 + 1e10) < (3.14159 / 1.0)) {
                            p2 = p2
                        }
                        _ = fi20
                    }
                }
            }
        case 4294967295:
            for fi28 := 0; fi28 < 13; fi28++ {
                v29 := []int{(fi28 - fi28)}
                if (1.5 <= (0.1 + (3.14159 * 0.0))) {
                    v30 := GenS6{F0: 0}
                    _ = v30
                }
                _ = fi28
                _ = v29
            }
        case 128:
            switch -151 {
            case -2147483648:
                var v31 int16 = 1000
                _ = v31
            case 0:
                fmt.Println(p1, p1, p1)
            }
        default:
            p1 = p1
        }
        _ = v18
    }
    v32 := []int{}
    if len(v32) >= 2 {
        v32 = v32[1:len(v32)]
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v32
    return 2
}

func main() {
    var v1 int8 = -124
    var v2 uint32 = 401
    var v3 uint8 = 127
    var v4 string = "hello"
    var v5 int32 = 42
    v33 := []int{127, 42, 202}
    if true {
        v4 = "foo bar"
        var v34 int8 = 127
        _ = v34
    } else {
        var v35 float64 = float64(-2.25)
        v33 = append(v33, 4294967295)
        _ = v35
    }
    var v36 int = 127
    v4 = v4
    var v37 uint8 = genFn0((v2 & (v2 | v2)), true, ("\n\t" + "hello"))
    if true {
        var v38 int = (v36 / (((v36 + 32768)) | 1))
        fmt.Println(v38)
        var v39 uint = 569
        _ = v38
        _ = v39
    }
    fmt.Println(v4)
    for fi40 := 0; fi40 < 20; fi40++ {
        v33 = append(v33, (-(fi40)))
        _ = fi40
    }
    var v41 float32 = float32(1.5)
    var v42 float64 = 1e10
    var v43 uint64 = 1
    if ((^(v5)) >= (^(v5))) {
        switch v4 {
        case "":
            v44 := GenS7{F0: 1000, F1: uint64(v37)}
            _ = v44
        case "a":
            if true {
                var v45 int8 = (v1 % ((int8(v43)) | 1))
                var v46 float32 = (0.0 + v41)
                _ = v45
                _ = v46
            } else {
                switch v4 {
                case "hello":
                    _ = len(v33)
                case "":
                    if false {
                        v42 += v42
                        v33 = append(v33, v36)
                        fmt.Println(v4, v2, v5)
                    }
                }
            }
        }
        for fi47 := 0; fi47 < 6; fi47++ {
            v42 = v42
            fmt.Println(v3)
            _ = fi47
        }
    }
    var v48 float32 = float32((0.1 * (1.5 + 1.5)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v48
    _ = v5
}
