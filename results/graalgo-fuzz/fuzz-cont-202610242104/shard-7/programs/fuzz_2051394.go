package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 int32; F2 bool }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 uint; F1 uint64; F2 int64 }
func genFn0(p0 string, p1 uint, p2 uint) uint16 {
    var v7 uint = 318
    p0 = (p0 + p0)
    v7 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}
func genFn1(p0 uint16, p1 string, p2 uint16) bool {
    for fi8 := 0; fi8 < 8; fi8++ {
        switch fi8 {
        case -1:
            for fi9 := 0; fi9 < 3; fi9++ {
                if false {
                    if (true && (!false)) {
                        var v10 uint = 498
                        _ = v10
                    } else {
                        switch fi8 {
                        case 1:
                            v11 := []int{fi8}
                            if (42 < -32768) {
                                p0 = (p2 * (p2 << ((18446744073709551615) & 63)))
                                p0 -= p2
                                var v12 int = fi8
                                _ = v12
                            }
                            _ = len(v11)
                            _ = v11
                        case 0:
                            var v13 uint8 = 2
                            _ = v13
                        case -790:
                            var v14 int16 = -32768
                            fmt.Println(v14, p1)
                            p0 -= (p2 >> ((489) & 63))
                            _ = v14
                        default:
                            var v15 uint = 65535
                            _ = v15
                        }
                        var v16 uint32 = 0
                        v17 := []int{}
                        p1 = p1
                        _ = v16
                        _ = v17
                    }
                } else {
                    switch ((p1 + (p1 + p1)) + p1) {
                    case "foo bar":
                        if (59 <= 42) {
                            var v18 string = p1
                            p1 = (v18 + (v18 + ("\n\t" + (("foo bar" + (("" + v18) + (p1 + p1))) + "a"))))
                            p1 = v18
                            _ = v18
                        }
                        switch fi8 {
                        case 128:
                            p2 = p0
                        case 4294967295:
                            fmt.Println(fi8)
                        default:
                            var v19 float32 = 3.14159
                            _ = v19
                        }
                        v20 := []int{1, 4294967295, fi8, fi9}
                        var v21 string = "a"
                        _ = v20
                        _ = v21
                    case "\n\t":
                        var v22 uint32 = uint32(p0)
                        _ = v22
                    case "a":
                        if (false || (-911 == 1)) {
                            var v23 int8 = -1
                            v24 := GenS1{F0: 0, F1: 931}
                            switch p1 {
                            case "foo bar":
                                p2 = 255
                            }
                            _ = v23
                            _ = v24
                        } else {
                            p1 = "foo bar"
                            p0 = (p0 ^ (p0 + genFn0("foo bar", 18446744073709551615, 770)))
                            v25 := GenS2{F0: (fi9 - 665)}
                            _ = v25
                        }
                    default:
                        var v26 uint = 617
                        _ = v26
                    }
                    if ((1.5 / -2.25) >= float32(1e10)) {
                        var v27 int64 = 9223372036854775807
                        _ = v27
                    }
                }
                p1 = "\n\t"
                _ = fi9
            }
        }
        _ = fi8
    }
    fmt.Println(p2)
    v28 := GenS3{F0: 2147483647, F1: 822, F2: (127 != 0)}
    var v29 float64 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    _ = v29
    return true
}

func main() {
    v1 := make(map[int]int)
    var v2 float32 = 0.0
    v3 := make(map[string]uint32)
    v4 := make(map[string]int)
    v5 := []int{-1}
    var v6 uint64 = 4294967295
    switch "" {
    case "a":
        v30 := GenS2{F0: 127}
        _ = v30
    default:
        for fi31 := 0; fi31 < 10; fi31++ {
            _ = len(v5)
            _ = fi31
        }
    }
    v32 := GenS3{F0: -595, F1: 1, F2: (1e10 != (1.5 + 1e10))}
    var v33 float32 = v2
    if true {
        var v34 int = 127
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v34
    }
    v6 += (v6 << ((292) & 63))
    for fi35 := 0; fi35 < 13; fi35++ {
        if (255 == 258) {
            v36 := GenS1{F0: -1, F1: 2147483647}
            if (!false) {
                if len(v5) > 0 {
                    v5[0] = -2147483648
                }
                switch (fi35 << ((0) & 63)) {
                case 2147483647:
                    for fi37 := 0; fi37 < 8; fi37++ {
                        var v38 int32 = -213
                        fmt.Println(v2)
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        var v39 string = ("hello" + "foo bar")
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                case -488:
                    fmt.Println(v6, fi35, v6)
                }
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            _ = v36
        } else {
            fmt.Println(v2)
        }
        _ = fi35
    }
    if false {
        switch ("\n\t" + "hello") {
        case "foo bar":
            if (42 < -9223372036854775808) {
                v2 = (3.14159 / ((v33 - v2) * 1.5))
            } else {
                var v40 float32 = v2
                delete(v4, "hello")
                _ = v40
            }
        case "hello":
            var v41 string = ""
            _ = v41
        default:
            for fi42 := 0; fi42 < 7; fi42++ {
                v43 := GenS1{F0: -1, F1: 0}
                _ = fi42
                _ = v43
            }
        }
    } else {
        _ = len(v3)
        switch "hello" {
        case "foo bar":
            if true {
                v6 += (v6 - (v6 >> ((uint(v6)) & 63)))
                var v44 bool = (!(769 <= 0))
                var v45 uint8 = 63
                _ = v44
                _ = v45
            } else {
                v46 := GenS4{F0: ("a" + ""), F1: 37}
                fmt.Println(v2)
                v47 := GenS5{F0: uint(v6), F1: 4294967295, F2: 311}
                _ = v46
                _ = v47
            }
        }
    }
    switch "\n\t" {
    case "hello":
        v2 = 1.5
    default:
        v3["hello"] = 42
    }
    var v48 string = "hello"
    if (-2147483648 > 1) {
        if (100 >= 0) {
            switch 42 {
            case 32768:
                if false {
                    var v49 uint8 = 208
                    v33 -= float32(1.5)
                    v2 = (v33 / 1.0)
                    _ = v49
                } else {
                    v50 := GenS5{F0: 255, F1: v6, F2: -333}
                    _ = v50
                }
            default:
                v33 = float32(float64((v33 + (v33 - v33))))
            }
            if (!true) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if false {
                    var v51 uint32 = 42
                    if (190 < 127) {
                        fmt.Println(v48, v2, v2)
                    }
                    switch (v48 + (("" + ("" + v48)) + "hello")) {
                    case "foo bar":
                        var v52 int = 128
                        _ = v52
                    default:
                        v33 = v2
                    }
                    _ = v51
                }
            }
            v33 += v2
        }
        var v53 uint = 100
        var v54 uint32 = 0
        _ = v53
        _ = v54
    }
    var v55 uint16 = genFn0(v48, 42, 225)
    if len(v5) > 0 {
        v5[0] = 32768
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v48
    _ = v5
    _ = v55
    _ = v6
}
