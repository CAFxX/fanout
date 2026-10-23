package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint8; F2 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32; F1 uint16; F2 int8 }
func genFn0(p0 uint, p1 string) int64 {
    var v7 uint32 = 943
    switch 841 {
    case -241:
        for fi8 := 0; fi8 < 18; fi8++ {
            if true {
                if ((!false) || false) {
                    v9 := GenS1{F0: 42}
                    var v10 int16 = -32768
                    v11 := GenS1{F0: 2147483647}
                    _ = v10
                    _ = v11
                    _ = v9
                }
                v7 += (v7 % ((0) | 1))
            } else {
                if true {
                    v12 := []int{fi8, fi8}
                    p1 = "foo bar"
                    _ = v12
                }
                p0 -= 255
                if true {
                    v13 := []int{fi8, 32767, 32768, fi8}
                    p1 = p1
                    fmt.Println(fi8, p0, fi8)
                    _ = v13
                } else {
                    v14 := GenS1{F0: 379}
                    for fi15 := 0; fi15 < 15; fi15++ {
                        var v16 float32 = (0.1 - 0.0)
                        var v17 uint = (p0 << ((255) & 63))
                        v7 = (v7 % ((2147483648) | 1))
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 10; fi18++ {
                        var v19 uint = 539
                        v19 = 520
                        switch fi18 {
                        case -218:
                            fmt.Println(v7)
                        case -561:
                            v20 := []int{fi18, (^(fi8)), fi8, fi18}
                            if (!true) {
                                if (!(!true)) {
                                    v21 := GenS2{F0: 1, F1: 1, F2: (p1 + p1)}
                                    v22 := GenS2{F0: 122, F1: 134, F2: ("" + "")}
                                    v23 := GenS2{F0: 593, F1: 1, F2: "foo bar"}
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                }
                                p1 = p1
                            } else {
                                _ = len(v20)
                            }
                            _ = v20
                        default:
                            v24 := []int{(fi8 | (fi8 | (fi18 + -1))), (fi8 ^ int(v7)), fi18, int(v7)}
                            v25 := GenS3{F0: 648}
                            p0 = v19
                            _ = v24
                            _ = v25
                        }
                        _ = fi18
                        _ = v19
                    }
                    _ = v14
                }
            }
            if (((((false && true) || false) || false) && true) && (false && false)) {
                if (p1 >= (p1 + (p1 + p1))) {
                    v26 := []int{fi8, (fi8 ^ (^(fi8))), fi8, int(p0)}
                    if false {
                        switch 32767 {
                        case 4294967295:
                            p1 = p1
                        default:
                            for fi27 := 0; fi27 < 17; fi27++ {
                                p0 += p0
                                _ = len(v26)
                                _ = fi27
                            }
                        }
                    } else {
                        if true {
                            var v28 int32 = 0
                            _ = v28
                        } else {
                            switch (fi8 & -137) {
                            case 4294967295:
                                for fi29 := 0; fi29 < 13; fi29++ {
                                    var v30 int64 = int64(fi29)
                                    v7 = (v7 % ((v7) | 1))
                                    v30 = (v30 - v30)
                                    _ = fi29
                                    _ = v30
                                }
                            case 128:
                                fmt.Println(p1)
                            }
                            p1 = (p1 + ((("hello" + "\n\t") + "a") + (p1 + p1)))
                        }
                        var v31 uint64 = 827
                        var v32 int32 = 42
                        _ = v31
                        _ = v32
                    }
                    var v33 float64 = float64(1.5)
                    _ = v26
                    _ = v33
                } else {
                    v34 := GenS1{F0: 42}
                    v7 = (^(v7))
                    var v35 int16 = 42
                    _ = v34
                    _ = v35
                }
            }
            v7 -= v7
            _ = fi8
        }
    default:
        p0 += p0
    }
    var v36 bool = true
    v7 -= v7
    _ = p0
    _ = p1
    _ = v36
    _ = v7
    return -1
}
func genFn1(p0 int16, p1 bool, p2 string) int32 {
    v37 := GenS4{F0: 987, F1: 42}
    var v38 uint = 858
    if (0 != 18446744073709551615) {
        p0 -= p0
        var v39 int8 = 127
        _ = v39
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v37
    _ = v38
    return -9
}

func main() {
    v1 := []string{"a"}
    var v2 int16 = 18
    v3 := make(map[int]bool)
    var v4 int = 32768
    var v5 float32 = 1.5
    v6 := []string{"a"}
    v40 := GenS4{F0: 653, F1: 32767}
    v3[(v4 << ((2) & 63))] = false
    fmt.Println(v4, v2, v5)
    delete(v3, (v4 | 0))
    switch (v4 % (((v4 / (((v4 * v4)) | 1))) | 1)) {
    case 32767:
        v4 = -910
    }
    v5 = (v5 / 1.0)
    v5 += -2.25
    v41 := GenS5{F0: 255}
    for fi42 := 0; fi42 < 9; fi42++ {
        for fi43 := 0; fi43 < 2; fi43++ {
            var v44 int16 = v2
            v45 := GenS6{F0: 42, F1: 0, F2: -1}
            switch "a" {
            case "\n\t":
                if true {
                    v46 := GenS3{F0: (-(v2))}
                    var v47 int = 4294967295
                    _ = v46
                    _ = v47
                }
            case "foo bar":
                if true {
                    v5 -= 1e10
                    fmt.Println(v44, fi43, v4)
                    var v48 uint = 2
                    _ = v48
                }
            default:
                fmt.Println(fi43)
            }
            _ = fi43
            _ = v44
            _ = v45
        }
        v2 = -550
        _ = fi42
    }
    for fi49 := 0; fi49 < 12; fi49++ {
        var v50 int64 = 1
        switch ("" + "foo bar") {
        case "\n\t":
            v51 := GenS4{F0: 42, F1: 32767}
            _ = v51
        case "a":
            var v52 int64 = genFn0(65535, ((("" + "foo bar") + ("a" + "\n\t")) + ""))
            _ = v52
        case "":
            var v53 bool = false
            _ = v53
        default:
            v54 := GenS1{F0: genFn1((^(v2)), false, "hello")}
            _ = v54
        }
        _ = fi49
        _ = v50
    }
    var v55 int32 = -1
    var v56 int = v4
    if ((false && (true || (1 <= 42))) || ((-(v55)) > (v55 - v55))) {
        if (!(!false)) {
            delete(v3, 42)
            delete(v3, 1)
            switch (v4 ^ (v4 << ((42) & 63))) {
            case 9223372036854775807:
                v5 += -2.25
            case 2147483647:
                v4 = 127
            default:
                v55 -= (v55 % ((genFn1((v2 << ((1) & 63)), (1 >= 1), "a")) | 1))
            }
        } else {
            switch (v4 - 0) {
            case -989:
                v4 = 128
            }
            _ = len(v3)
        }
        _ = len(v3)
        v57 := GenS6{F0: 65535, F1: 42, F2: -1}
        _ = v57
    }
    var v58 string = ""
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v55
    _ = v56
    _ = v58
    _ = v6
}
