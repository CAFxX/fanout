package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 uint8; F2 uint }
type GenS3 struct { F0 int64; F1 uint16; F2 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint8, p1 int64) int16 {
    switch "a" {
    case "foo bar":
        v5 := GenS1{F0: -13}
        _ = v5
    }
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 bool = false
        v8 := GenS1{F0: 2147483647}
        _ = fi6
        _ = v7
        _ = v8
    }
    if ((!false) && (3.14159 > float64(0.1))) {
        if true {
            if (!("" < "")) {
                v9 := GenS2{F0: 4294967295, F1: 42, F2: 52}
                switch -970 {
                case 4294967295:
                    var v10 uint32 = 2147483647
                    _ = v10
                case 927:
                    for fi11 := 0; fi11 < 7; fi11++ {
                        if true {
                            var v12 int8 = -26
                            _ = v12
                        } else {
                            if true {
                                for fi13 := 0; fi13 < 17; fi13++ {
                                    v14 := GenS3{F0: -9223372036854775808, F1: 2, F2: 4294967295}
                                    if (!true) {
                                        switch "" {
                                        case "":
                                            var v15 int64 = p1
                                            _ = v15
                                        case "foo bar":
                                            v16 := []int{(fi11 >> ((4294967295) & 63)), (fi11 << ((41) & 63))}
                                            _ = v16
                                        case "\n\t":
                                            fmt.Println(p1, p0)
                                        default:
                                            var v17 uint = 2
                                            _ = v17
                                        }
                                    }
                                    p1 = p1
                                    _ = fi13
                                    _ = v14
                                }
                                var v18 float32 = 0.0
                                _ = v18
                            }
                            var v19 uint16 = 751
                            _ = v19
                        }
                        _ = fi11
                    }
                default:
                    v20 := GenS3{F0: 9223372036854775807, F1: 47, F2: 2}
                    _ = v20
                }
                _ = v9
            } else {
                fmt.Println(p0, p0)
            }
        }
        p0 -= (p0 | p0)
        for fi21 := 0; fi21 < 13; fi21++ {
            switch 248 {
            case 722:
                v22 := GenS3{F0: (p1 >> ((1) & 63)), F1: 32767, F2: 65535}
                _ = v22
            default:
                var v23 uint8 = 42
                _ = v23
            }
            for fi24 := 0; fi24 < 5; fi24++ {
                v25 := GenS3{F0: p1, F1: 398, F2: uint(p0)}
                _ = fi24
                _ = v25
            }
            switch (fi21 * -9223372036854775808) {
            case -385:
                for fi26 := 0; fi26 < 12; fi26++ {
                    if (278 >= 32767) {
                        var v27 bool = false
                        v28 := GenS4{F0: -171}
                        _ = v27
                        _ = v28
                    } else {
                        if false {
                            v29 := []int{-270}
                            var v30 int16 = 338
                            fmt.Println(v30, v30)
                            _ = v29
                            _ = v30
                        }
                        var v31 string = ("a" + ("\n\t" + "hello"))
                        p1 += (p1 ^ p1)
                        _ = v31
                    }
                    p0 = p0
                    var v32 int = (fi26 >> ((4294967295) & 63))
                    v33 := []int{(fi21 >> ((0) & 63))}
                    _ = fi26
                    _ = v32
                    _ = v33
                }
            case -1:
                fmt.Println(fi21, p1, fi21)
            default:
                var v34 int = fi21
                _ = v34
            }
            _ = fi21
        }
    }
    p1 += (p1 & (p1 & p1))
    _ = p0
    _ = p1
    return 1000
}
func genFn1(p0 string, p1 uint16, p2 uint16) int32 {
    switch -529 {
    case 0:
        fmt.Println(p2, p0, p1)
    case 42:
        v35 := GenS5{F0: (p0 + "\n\t")}
        _ = v35
    }
    var v36 uint = 18446744073709551615
    fmt.Println(p1, v36)
    _ = p0
    _ = p1
    _ = p2
    _ = v36
    return -1
}

func main() {
    v1 := []uint32{65535, 65535, 248}
    var v2 int32 = 633
    var v3 int16 = 42
    var v4 string = (("hello" + (("foo bar" + "") + "hello")) + "hello")
    switch (v4 + "\n\t") {
    case "foo bar":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "a":
        var v37 float64 = 0.1
        _ = v37
    }
    fmt.Println(v3)
    switch ("hello" + v4) {
    case "hello":
        for fi38 := 0; fi38 < 4; fi38++ {
            fmt.Println(v3, fi38, v2)
            v3 += (v3 ^ -1)
            _ = fi38
        }
    case "a":
        fmt.Println(v2, v3, v2)
    }
    v2 += -211
    for fi39 := 0; fi39 < 11; fi39++ {
        v40 := GenS6{F0: -1, F1: 0}
        v2 = (v2 % ((genFn1(v4, 13, 0)) | 1))
        var v41 int64 = -9223372036854775808
        _ = fi39
        _ = v40
        _ = v41
    }
    fmt.Println(v2, v2, v4)
    for fi42 := 0; fi42 < 10; fi42++ {
        var v43 uint16 = 977
        var v44 uint8 = 186
        _ = fi42
        _ = v43
        _ = v44
    }
    switch ("foo bar" + v4) {
    case "\n\t":
        var v45 int = 1
        _ = v45
    case "":
        v4 = v4
    default:
        v46 := GenS6{F0: (v3 | v3), F1: 255}
        _ = v46
    }
    for fi47 := 0; fi47 < 17; fi47++ {
        switch (fi47 ^ 1) {
        case 127:
            var v48 uint64 = 351
            _ = v48
        case 283:
            var v49 uint64 = 42
            _ = v49
        case -756:
            var v50 float32 = 0.0
            _ = v50
        }
        var v51 int8 = -123
        _ = fi47
        _ = v51
    }
    v4 = ((("foo bar" + "a") + "a") + v4)
    var v52 string = v4
    fmt.Println(v4)
    var v53 uint64 = 4294967295
    v53 += v53
    var v54 int32 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v53
    _ = v54
}
