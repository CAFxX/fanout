package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint; F2 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 uint16 }
func genFn0(p0 uint64, p1 uint, p2 uint16) int64 {
    switch 42 {
    case 1:
        p1 = (p1 * (p1 - p1))
    case -2147483648:
        var v7 int16 = 42
        _ = v7
    case 0:
        p1 += (^(p1))
    }
    var v8 int64 = -539
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return v8
}
func genFn1(p0 string) uint8 {
    p0 = p0
    p0 = (p0 + p0)
    var v9 bool = true
    for fi10 := 0; fi10 < 18; fi10++ {
        p0 = p0
        if (2147483647 != 2147483647) {
            if v9 {
                var v11 uint8 = 128
                if ((^(v11)) < (v11 << ((255) & 63))) {
                    var v12 int16 = 926
                    v13 := GenS1{F0: 65535}
                    var v14 float64 = 3.14159
                    _ = v12
                    _ = v13
                    _ = v14
                }
                _ = v11
            } else {
                if v9 {
                    v15 := []int{}
                    if (134 != 255) {
                        v16 := GenS1{F0: 2147483648}
                        _ = len(v15)
                        _ = v16
                    } else {
                        if len(v15) > 0 {
                            v15[0] = fi10
                        }
                        v17 := GenS1{F0: 42}
                        _ = v17
                    }
                    var v18 bool = (((3.14159 * 0.1) / 1e10) < ((-2.25 * ((float64((0.1 + (-2.25 * -2.25))) - -2.25) + (-2.25 + 1e10))) + -2.25))
                    _ = v15
                    _ = v18
                } else {
                    switch p0 {
                    case "":
                        var v19 int8 = 1
                        _ = v19
                    }
                    var v20 int16 = 0
                    _ = v20
                }
                p0 = ("hello" + "hello")
            }
            var v21 float32 = -2.25
            _ = v21
        }
        p0 = ((p0 + p0) + "foo bar")
        _ = fi10
    }
    fmt.Println(v9)
    _ = p0
    _ = v9
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 2
    var v3 int32 = 1
    v4 := make(map[uint32]string)
    var v5 uint8 = 85
    v6 := make(map[int]string)
    v3 = v3
    fmt.Println(v3, v5)
    if false {
        for fi22 := 0; fi22 < 1; fi22++ {
            if false {
                v1[2] = ((65535 == 472) || false)
            } else {
                v4[645] = "foo bar"
            }
            v2 = 255
            var v23 uint32 = 4294967295
            _ = fi22
            _ = v23
        }
        v2 += v2
    }
    fmt.Println(v5, v5)
    var v24 uint16 = 65535
    switch "\n\t" {
    case "\n\t":
        v25 := []int{1}
        _ = v25
    case "a":
        fmt.Println(v3)
    }
    v24 += v24
    fmt.Println(v2, v24)
    if false {
        v26 := GenS2{F0: -9223372036854775808, F1: 207, F2: -32768}
        for fi27 := 0; fi27 < 4; fi27++ {
            v5 += (^(v5))
            _ = fi27
        }
        _ = v26
    } else {
        fmt.Println(v24, v24, v5)
        v28 := GenS2{F0: 182, F1: 42, F2: 1}
        for fi29 := 0; fi29 < 9; fi29++ {
            v30 := []int{(^(fi29)), (fi29 >> ((4294967295) & 63))}
            v24 += (^(v24))
            _ = fi29
            _ = v30
        }
        _ = v28
    }
    v31 := []int{617, 4294967295, -1}
    for fi32 := 0; fi32 < 5; fi32++ {
        for fi33 := 0; fi33 < 4; fi33++ {
            fmt.Println(fi32)
            fmt.Println(v2)
            _ = fi33
        }
        _ = fi32
    }
    switch "hello" {
    case "":
        v24 = 0
    case "hello":
        delete(v1, 4294967295)
    default:
        if (10 != -128) {
            for fi34 := 0; fi34 < 14; fi34++ {
                v24 = (v24 % ((348) | 1))
                _ = fi34
            }
            if (917 != 4294967295) {
                if (!true) {
                    fmt.Println(v24, v5)
                    v2 -= (v2 / ((255) | 1))
                }
                var v35 int16 = 1000
                var v36 int = 128
                _ = v35
                _ = v36
            } else {
                fmt.Println(v3, v2)
            }
        }
    }
    for fi37 := 0; fi37 < 15; fi37++ {
        switch "foo bar" {
        case "foo bar":
            var v38 int = (fi37 << ((42) & 63))
            _ = v38
        case "hello":
            switch ((("" + "\n\t") + "") + "\n\t") {
            case "":
                switch "" {
                case "\n\t":
                    var v39 float64 = float64((0.1 / ((0.1 - 3.14159) + (float32((float64(1.5) / 3.14159)) - 0.1))))
                    _ = v39
                }
            case "foo bar":
                v2 += v5
            case "a":
                _ = len(v31)
            default:
                v40 := v4[42]
                _ = v40
            }
        case "":
            _ = len(v4)
        default:
            var v41 int = fi37
            _ = v41
        }
        _ = fi37
    }
    for fi42 := 0; fi42 < 4; fi42++ {
        if (!false) {
            v24 = v24
        }
        _ = fi42
    }
    switch (("" + "hello") + (("hello" + "") + ("foo bar" + "\n\t"))) {
    case "\n\t":
        for fi43 := 0; fi43 < 13; fi43++ {
            if (!(int(v24) <= -1)) {
                var v44 int64 = -1
                _ = v44
            } else {
                v2 += (v2 * v2)
                v2 = (v5 / ((v5) | 1))
            }
            for fi45 := 0; fi45 < 3; fi45++ {
                if ((1.5 / 1e10) != float64(0.1)) {
                    delete(v6, -107)
                    if false {
                        var v46 int16 = 42
                        v5 += (v2 & genFn1(("a" + ((("hello" + "a") + ("foo bar" + ("" + "a"))) + "foo bar"))))
                        v4[2147483648] = "foo bar"
                        _ = v46
                    } else {
                        if false {
                            switch fi43 {
                            case -667:
                                v24 = (v24 & (^(v24)))
                            case 223:
                                switch ("" + "hello") {
                                case "foo bar":
                                    var v47 string = ""
                                    _ = v47
                                case "":
                                    if (!true) {
                                        fmt.Println(fi43, fi43)
                                        var v48 uint32 = 1
                                        v3 += v3
                                        _ = v48
                                    } else {
                                        var v49 int16 = 1000
                                        v5 += v2
                                        _ = v49
                                    }
                                }
                            }
                            v50 := GenS1{F0: 4294967295}
                            v51 := GenS1{F0: 2147483648}
                            _ = v50
                            _ = v51
                        } else {
                            switch "hello" {
                            case "foo bar":
                                var v52 uint16 = 1
                                _ = v52
                            case "hello":
                                v31 = append(v31, fi43)
                            case "\n\t":
                                _ = len(v31)
                            default:
                                if (v24 <= 65535) {
                                    v53 := GenS3{F0: -128}
                                    var v54 uint16 = (v24 % ((uint16(v2)) | 1))
                                    _ = v53
                                    _ = v54
                                } else {
                                    _ = len(v4)
                                    switch ("\n\t" + "hello") {
                                    case "\n\t":
                                        if ((v24 & v24) > v24) {
                                            v55 := GenS4{F0: 1}
                                            fmt.Println(v2, v3, v2)
                                            _ = v55
                                        }
                                    case "a":
                                        var v56 uint32 = 0
                                        _ = v56
                                    case "":
                                        fmt.Println(v24)
                                    }
                                }
                            }
                            fmt.Println(fi45)
                            v6[(^(fi43))] = "\n\t"
                            var v57 int64 = 214
                            if true {
                                if (327 > 255) {
                                    v58 := GenS5{F0: fi45, F1: (fi45 & -889)}
                                    _ = v58
                                }
                                fmt.Println(v24)
                            } else {
                                var v59 int16 = -1
                                if len(v31) >= 2 {
                                    v31 = v31[1:len(v31)]
                                }
                                _ = v59
                            }
                            _ = v57
                        }
                        var v60 int32 = (v3 | v3)
                        _ = v60
                    }
                    v5 = 1
                } else {
                    if false {
                        v24 += (^(v24))
                        v5 = (v2 + v5)
                        var v61 uint = 337
                        _ = v61
                    }
                    v2 -= (v5 + v5)
                    if (1 == uint(v24)) {
                        var v62 int16 = -719
                        _ = v62
                    } else {
                        delete(v1, 65535)
                        v24 = (v24 % ((785) | 1))
                        v63 := GenS6{F0: (v24 & 926)}
                        _ = v63
                    }
                }
                v64 := GenS5{F0: (-(fi43)), F1: 1}
                fmt.Println(v5, v3)
                _ = fi45
                _ = v64
            }
            _ = fi43
        }
    case "hello":
        v2 += 148
    default:
        v65 := GenS5{F0: 42, F1: 127}
        _ = v65
    }
    v24 -= v24
    for fi66 := 0; fi66 < 4; fi66++ {
        var v67 int8 = int8(fi66)
        _ = fi66
        _ = v67
    }
    v5 = (v2 - v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
