package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint32; F1 int16; F2 uint64 }
type GenS6 struct { F0 bool; F1 uint32; F2 int }
type GenS7 struct { F0 uint; F1 int64; F2 int64 }
type GenS8 struct { F0 uint64; F1 bool; F2 int32 }

func main() {
    var v1 int8 = -1
    v2 := make(map[string]int)
    var v3 uint32 = 936
    var v4 uint = 4294967295
    var v5 uint64 = uint64(v4)
    v5 = (^(v5))
    var v6 int16 = 1000
    v4 = v4
    var v7 int32 = 71
    v2[("hello" + "")] = 0
    var v8 float32 = float32(-2.25)
    v9 := GenS1{F0: 777, F1: 1}
    var v10 uint = v4
    for fi11 := 0; fi11 < 18; fi11++ {
        if false {
            switch (fi11 + (fi11 * 9223372036854775807)) {
            case 32767:
                fmt.Println(v6, v5, fi11)
            }
            fmt.Println(v1, fi11)
            v12 := GenS2{F0: v5}
            _ = v12
        } else {
            v13 := GenS3{F0: v1, F1: 226}
            _ = v13
        }
        if true {
            v1 = 100
            var v14 int8 = 73
            switch ("hello" + "foo bar") {
            case "":
                if (!true) {
                    var v15 int16 = (v6 ^ 1)
                    var v16 int64 = 0
                    _ = v15
                    _ = v16
                }
            case "\n\t":
                var v17 int16 = (v6 % (((v6 << ((18446744073709551615) & 63))) | 1))
                _ = v17
            }
            _ = v14
        } else {
            if (!true) {
                v18 := GenS3{F0: v1, F1: 432}
                _ = v18
            } else {
                if (v3 >= (v3 % ((v3) | 1))) {
                    var v19 string = (((("\n\t" + "hello") + ("" + ("hello" + "foo bar"))) + "a") + "a")
                    _ = v19
                }
                if true {
                    if true {
                        if (1 > v1) {
                            fmt.Println(fi11, v6)
                        } else {
                            var v20 uint = (v4 ^ (v4 & (v4 * v4)))
                            v21 := []int{(fi11 + (fi11 & (fi11 + (fi11 << ((4294967295) & 63))))), fi11, int(v3), fi11}
                            switch fi11 {
                            case -858:
                                _ = len(v2)
                            case -561:
                                if (832 > 255) {
                                    switch fi11 {
                                    case -313:
                                        _ = len(v21)
                                    }
                                    switch ("\n\t" + "foo bar") {
                                    case "foo bar":
                                        if true {
                                            var v22 float32 = (v8 * v8)
                                            _ = v22
                                        }
                                    case "hello":
                                        v1 -= v1
                                    case "":
                                        if (!((1e10 * (0.0 / 0.1)) <= 1.5)) {
                                            for fi23 := 0; fi23 < 10; fi23++ {
                                                var v24 int16 = v6
                                                _ = fi23
                                                _ = v24
                                            }
                                            v7 = v7
                                            fmt.Println(fi11, v3, fi11)
                                        } else {
                                            v1 += (v1 / (((v1 & v1)) | 1))
                                            var v25 uint16 = 42
                                            if false {
                                                v26 := GenS2{F0: v5}
                                                _ = v26
                                            } else {
                                                for fi27 := 0; fi27 < 18; fi27++ {
                                                    v3 -= 664
                                                    _ = fi27
                                                }
                                                v28 := GenS4{F0: (^(v25)), F1: -81, F2: uint8(v3)}
                                                _ = v28
                                            }
                                            _ = v25
                                        }
                                    }
                                    if ((7 > 20) && true) {
                                        v29 := v2[("hello" + ("a" + "a"))]
                                        var v30 uint = (^(v20))
                                        _ = v29
                                        _ = v30
                                    } else {
                                        v7 += v7
                                        _ = len(v21)
                                        delete(v2, "")
                                    }
                                } else {
                                    v31 := GenS5{F0: v3, F1: -1, F2: v5}
                                    _ = v31
                                }
                            case 1:
                                var v32 int8 = v1
                                _ = v32
                            }
                            _ = v20
                            _ = v21
                        }
                    }
                }
            }
            if (0 >= 65535) {
                v33 := GenS6{F0: (0 >= 255), F1: v3, F2: fi11}
                _ = v33
            } else {
                v34 := GenS1{F0: int(v3), F1: v4}
                var v35 uint64 = (v5 % (((v5 % ((110) | 1))) | 1))
                _ = v34
                _ = v35
            }
            if (2 != 416) {
                v1 += (-(v1))
            } else {
                var v36 uint32 = v3
                if false {
                    fmt.Println(v3, v7, v5)
                    fmt.Println(v10, fi11)
                    var v37 uint64 = (v5 | (v5 % ((2) | 1)))
                    _ = v37
                }
                _ = v36
            }
        }
        _ = fi11
    }
    if ((v3 << (((^(v4))) & 63)) <= (v3 % (((v3 | v3)) | 1))) {
        v38 := []int{}
        v1 = (v1 >> ((v4) & 63))
        _ = v38
    } else {
        if false {
            switch "hello" {
            case "":
                switch -1 {
                case 4294967295:
                    var v39 float32 = -2.25
                    _ = v39
                default:
                    v40 := GenS1{F0: -9223372036854775808, F1: (^(v10))}
                    _ = v40
                }
            default:
                for fi41 := 0; fi41 < 6; fi41++ {
                    switch "foo bar" {
                    case "a":
                        switch "hello" {
                        case "":
                            v10 -= uint(v5)
                        }
                    default:
                        for fi42 := 0; fi42 < 10; fi42++ {
                            var v43 uint = v10
                            _ = fi42
                            _ = v43
                        }
                    }
                    _ = fi41
                }
            }
        }
    }
    var v44 int = -1
    v45 := GenS7{F0: 0, F1: 0, F2: 0}
    v46 := GenS8{F0: (^(v5)), F1: true, F2: v7}
    v47 := []int{(-(v44)), (v44 / ((-9223372036854775808) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
