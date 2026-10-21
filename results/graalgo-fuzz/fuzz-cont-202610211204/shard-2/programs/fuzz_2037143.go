package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS4 struct { F0 int32 }
func genFn0() uint64 {
    var v5 uint = 1
    var v6 int8 = 100
    v6 += (v6 + v6)
    var v7 uint = v5
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := []int{fi8, (fi8 % ((fi8) | 1))}
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        if ("hello" < ((((("" + "foo bar") + "a") + ("hello" + "foo bar")) + "") + "a")) {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        } else {
            switch fi8 {
            case 0:
                var v10 int64 = 1
                _ = v10
            }
        }
        _ = fi8
        _ = v9
    }
    _ = v5
    _ = v6
    _ = v7
    return 793
}

func main() {
    var v1 bool = true
    var v2 float64 = 1.5
    var v3 uint16 = 0
    v4 := make(map[int]bool)
    var v11 bool = (!(!false))
    if v11 {
        if (0.1 >= (v2 - 0.0)) {
            var v12 uint16 = (^(v3))
            _ = v12
        }
        v13 := []int{42}
        var v14 int16 = 1
        _ = v13
        _ = v14
    }
    if ((1 < 0) || v1) {
        for fi15 := 0; fi15 < 2; fi15++ {
            if v1 {
                var v16 uint32 = 1
                if (1 < 42) {
                    v17 := GenS1{F0: 255}
                    _ = v17
                } else {
                    var v18 bool = ((fi15 + (fi15 / ((32767) | 1))) != fi15)
                    if (!v18) {
                        fmt.Println(v1, v3, fi15)
                    } else {
                        var v19 bool = (!true)
                        v20 := []int{}
                        _ = v19
                        _ = v20
                    }
                    var v21 uint8 = 111
                    _ = v18
                    _ = v21
                }
                var v22 uint16 = (v3 * 42)
                _ = v16
                _ = v22
            } else {
                fmt.Println(v3, v1, v2)
            }
            v11 = false
            var v23 float32 = 3.14159
            _ = fi15
            _ = v23
        }
        v11 = (v1 || v1)
    }
    v24 := v4[0]
    switch -9223372036854775808 {
    case -67:
        if ("a" > "") {
            v25 := GenS1{F0: 32767}
            _ = v25
        } else {
            var v26 uint32 = 332
            v27 := []int{-2147483648, 32767}
            _ = v26
            _ = v27
        }
    case 128:
        if (247 <= 18446744073709551615) {
            if false {
                if v24 {
                    fmt.Println(v3)
                } else {
                    var v28 int16 = 32767
                    if (!(-1 > (v28 | v28))) {
                        delete(v4, 612)
                        fmt.Println(v24, v2, v2)
                    } else {
                        fmt.Println(v2, v3, v24)
                        for fi29 := 0; fi29 < 5; fi29++ {
                            var v30 int64 = 671
                            _ = fi29
                            _ = v30
                        }
                    }
                    v24 = v11
                    _ = v28
                }
                fmt.Println(v11, v11)
            } else {
                v24 = v11
            }
            var v31 uint16 = 531
            _ = v31
        } else {
            switch -771 {
            case -9223372036854775808:
                v24 = (1 > 468)
            case 127:
                if (42 > 100) {
                    var v32 uint8 = 249
                    if false {
                        v33 := GenS2{F0: -1}
                        delete(v4, 128)
                        v34 := GenS2{F0: 1}
                        _ = v33
                        _ = v34
                    }
                    _ = v32
                } else {
                    v35 := GenS2{F0: 481}
                    switch "\n\t" {
                    case "foo bar":
                        var v36 bool = (-2.25 >= float32(v2))
                        _ = v36
                    case "a":
                        switch "\n\t" {
                        case "":
                            for fi37 := 0; fi37 < 6; fi37++ {
                                var v38 int = 32768
                                _ = fi37
                                _ = v38
                            }
                        case "a":
                            switch 2147483647 {
                            case 32768:
                                v2 -= float64((1e10 + float32(v2)))
                            case -801:
                                fmt.Println(v24, v11)
                            }
                        case "foo bar":
                            v39 := v4[-9223372036854775808]
                            _ = v39
                        }
                    case "hello":
                        v40 := GenS1{F0: v3}
                        _ = v40
                    }
                    _ = v35
                }
            case 265:
                for fi41 := 0; fi41 < 9; fi41++ {
                    if (v11 && (false || v24)) {
                        v24 = ((fi41 % (((fi41 | -712)) | 1)) < fi41)
                    } else {
                        v4[0] = ((!((!false) || (1 <= 1))) || (!false))
                        if ((410 <= 0) || (!(!false))) {
                            var v42 int64 = 42
                            for fi43 := 0; fi43 < 19; fi43++ {
                                var v44 uint64 = 39
                                v45 := GenS1{F0: (v3 ^ v3)}
                                _ = fi43
                                _ = v44
                                _ = v45
                            }
                            _ = v42
                        } else {
                            switch "" {
                            case "\n\t":
                                var v46 int = (fi41 + (fi41 << ((4294967295) & 63)))
                                _ = v46
                            default:
                                v47 := GenS3{F0: -857, F1: 4294967295, F2: v3}
                                _ = v47
                            }
                            v48 := GenS1{F0: (v3 ^ v3)}
                            if false {
                                if v24 {
                                    v24 = v11
                                    if v1 {
                                        var v49 uint32 = 4294967295
                                        _ = v49
                                    }
                                } else {
                                    v50 := []int{732, -9223372036854775808, fi41, (fi41 + fi41)}
                                    _ = v50
                                }
                                var v51 uint8 = 127
                                switch 32768 {
                                case 4294967295:
                                    if v11 {
                                        v52 := GenS4{F0: -1}
                                        _ = v52
                                    }
                                case 0:
                                    switch 357 {
                                    case 128:
                                        for fi53 := 0; fi53 < 11; fi53++ {
                                            var v54 float64 = float64(1e10)
                                            v55 := GenS1{F0: v3}
                                            _ = fi53
                                            _ = v54
                                            _ = v55
                                        }
                                    case 1:
                                        v51 = v51
                                    default:
                                        var v56 uint64 = 18446744073709551615
                                        _ = v56
                                    }
                                }
                                _ = v51
                            } else {
                                var v57 bool = (true && true)
                                v58 := []int{(fi41 ^ fi41), (fi41 | (fi41 ^ 32768)), fi41, (-(fi41))}
                                _ = v57
                                _ = v58
                            }
                            _ = v48
                        }
                    }
                    for fi59 := 0; fi59 < 9; fi59++ {
                        if v24 {
                            v2 += (v2 + v2)
                            fmt.Println(v11)
                            var v60 int64 = 939
                            _ = v60
                        }
                        _ = fi59
                    }
                    fmt.Println(v1)
                    _ = fi41
                }
            }
        }
    }
    v61 := []int{}
    if v1 {
        v3 += 434
        v62 := GenS3{F0: -176, F1: 42, F2: (^(v3))}
        fmt.Println(v2, v3)
        _ = v62
    } else {
        fmt.Println(v24, v3)
        var v63 uint32 = 4294967295
        _ = v63
    }
    for fi64 := 0; fi64 < 20; fi64++ {
        if v11 {
            var v65 bool = false
            for fi66 := 0; fi66 < 15; fi66++ {
                if (!(!((fi66 << ((881) & 63)) <= -662))) {
                    v67 := GenS3{F0: 791, F1: genFn0(), F2: 268}
                    fmt.Println(fi64, v65)
                    _ = v67
                } else {
                    var v68 string = "hello"
                    if len(v61) >= 2 {
                        v61 = v61[1:len(v61)]
                    }
                    _ = v68
                }
                switch (fi66 | fi66) {
                case 128:
                    if v1 {
                        fmt.Println(v24, v3, fi64)
                    } else {
                        v2 = v2
                    }
                case -9223372036854775808:
                    var v69 uint8 = 153
                    _ = v69
                }
                var v70 string = (("\n\t" + ("hello" + ("" + "hello"))) + "foo bar")
                var v71 uint8 = 0
                _ = fi66
                _ = v70
                _ = v71
            }
            for fi72 := 0; fi72 < 19; fi72++ {
                if false {
                    fmt.Println(v24, v1, v11)
                    v1 = true
                    var v73 int8 = 90
                    _ = v73
                }
                _ = fi72
            }
            _ = v65
        }
        fmt.Println(v11, v2)
        _ = fi64
    }
    var v74 bool = true
    v2 = (((v2 * v2) + v2) / v2)
    v11 = true
    v11 = v1
    var v75 float64 = v2
    v76 := GenS2{F0: -1}
    if len(v61) >= 2 {
        v61 = v61[1:len(v61)]
    }
    var v77 uint = 1
    if (false && true) {
        v11 = (0 <= 903)
    } else {
        var v78 uint32 = 1
        v1 = true
        v24 = false
        _ = v78
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v61:", len(v61), v61)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v61
    _ = v74
    _ = v75
    _ = v76
    _ = v77
}
