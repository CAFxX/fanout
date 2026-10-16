package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 uint32 }
func genFn0() uint64 {
    switch 42 {
    case -9223372036854775808:
        switch "a" {
        case "foo bar":
            var v5 int = -9223372036854775808
            _ = v5
        case "a":
            var v6 int = 4294967295
            _ = v6
        case "\n\t":
            if (true && (100 <= -2147483648)) {
                var v7 int64 = -950
                _ = v7
            }
        }
    case 1:
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 int32 = 1
            _ = fi8
            _ = v9
        }
    }
    var v10 int32 = 1
    switch "\n\t" {
    case "":
        if (70 < 635) {
            v11 := GenS1{F0: 65535, F1: 4294967295, F2: 128}
            v10 -= v10
            v12 := GenS2{F0: ((-32768 > -711) && (("foo bar" + (("hello" + "") + ("foo bar" + ("foo bar" + "a")))) != ("a" + ("a" + "hello"))))}
            _ = v11
            _ = v12
        } else {
            v13 := []int{-2147483648}
            _ = v13
        }
    default:
        for fi14 := 0; fi14 < 19; fi14++ {
            v10 -= (v10 | (v10 % (((v10 % ((0) | 1))) | 1)))
            v10 = (^(v10))
            _ = fi14
        }
    }
    v15 := GenS3{F0: 139, F1: true}
    if false {
        for fi16 := 0; fi16 < 20; fi16++ {
            v10 = 0
            if ((100 == -1) && (-1 == -78)) {
                for fi17 := 0; fi17 < 13; fi17++ {
                    fmt.Println(fi17, v10, fi17)
                    _ = fi17
                }
                var v18 int16 = 1000
                _ = v18
            } else {
                v10 = -875
            }
            _ = fi16
        }
        v19 := []int{974}
        _ = v19
    }
    _ = v10
    _ = v15
    return 2
}

func main() {
    var v1 int64 = 42
    var v2 uint = 4294967295
    var v3 uint16 = 314
    var v4 bool = (v1 != 21)
    for fi20 := 0; fi20 < 18; fi20++ {
        v21 := []int{-796, 94}
        _ = fi20
        _ = v21
    }
    var v22 uint64 = 0
    switch -1 {
    case -2147483648:
        v22 -= v22
    case 0:
        switch "foo bar" {
        case "hello":
            for fi23 := 0; fi23 < 9; fi23++ {
                if v4 {
                    for fi24 := 0; fi24 < 4; fi24++ {
                        var v25 string = (("" + ("hello" + ("a" + "foo bar"))) + ("" + ("\n\t" + (("hello" + "foo bar") + ("" + "foo bar")))))
                        fmt.Println(fi23)
                        _ = fi24
                        _ = v25
                    }
                    switch 42 {
                    case 432:
                        v26 := []int{}
                        _ = v26
                    }
                    if (v4 && false) {
                        for fi27 := 0; fi27 < 19; fi27++ {
                            v28 := GenS3{F0: 127, F1: false}
                            v29 := []int{fi23, (fi27 % (((^(fi23))) | 1))}
                            if ((false || v4) && v4) {
                                switch fi23 {
                                case -2147483648:
                                    var v30 int16 = 0
                                    _ = v30
                                case 1:
                                    v31 := GenS1{F0: 506, F1: v2, F2: fi23}
                                    _ = v31
                                case 32767:
                                    var v32 float64 = ((1.5 - (0.1 + 0.0)) * 1e10)
                                    _ = v32
                                default:
                                    var v33 int = fi27
                                    _ = v33
                                }
                            }
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                        var v34 int16 = 180
                        _ = v34
                    } else {
                        if (1 < -1) {
                            v4 = (163 > 15)
                            v35 := []int{(fi23 | -628)}
                            _ = v35
                        } else {
                            v1 = 656
                            if v4 {
                                var v36 uint32 = 0
                                fmt.Println(v3, v36, v3)
                                var v37 int16 = 1000
                                _ = v36
                                _ = v37
                            }
                        }
                    }
                }
                _ = fi23
            }
        }
    default:
        fmt.Println(v4, v22, v4)
    }
    var v38 int64 = (v1 | (v1 << (((v2 % ((v2) | 1))) & 63)))
    v39 := []int{307, 505, 42}
    for fi40 := 0; fi40 < 6; fi40++ {
        if (v4 && (1 > 802)) {
            v41 := GenS4{F0: 2147483648}
            v38 = v1
            _ = v41
        }
        _ = len(v39)
        if ((1e10 + 0.0) < float32(3.14159)) {
            var v42 uint16 = v3
            _ = v42
        }
        _ = fi40
    }
    var v43 uint16 = (v3 % (((v3 | uint16(v22))) | 1))
    v38 = v38
    if len(v39) > 0 {
        v39[0] = 2147483647
    }
    var v44 bool = v4
    v1 -= v38
    var v45 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v43
    _ = v44
    _ = v45
}
