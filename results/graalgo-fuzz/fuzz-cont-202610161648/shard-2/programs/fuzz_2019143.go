package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 int64; F1 string; F2 uint64 }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 int8; F1 int8; F2 uint16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint32; F1 int8 }
func genFn0(p0 uint32) string {
    v6 := []int{32768, 2147483647, 32767}
    var v7 float32 = (0.1 / 1.0)
    var v8 uint8 = 0
    switch ("" + ("hello" + "")) {
    case "hello":
        v8 = (v8 + v8)
    }
    var v9 bool = (0 != -1)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return ("" + (("hello" + "") + ""))
}
func genFn1(p0 uint8, p1 int32, p2 int32) int64 {
    if (1 >= -1) {
        v10 := GenS1{F0: 255}
        _ = v10
    } else {
        for fi11 := 0; fi11 < 8; fi11++ {
            for fi12 := 0; fi12 < 9; fi12++ {
                p1 -= (p1 - 110)
                if false {
                    if false {
                        v13 := GenS2{F0: false, F1: 0}
                        if false {
                            p2 = (^(p2))
                        }
                        if true {
                            var v14 int16 = -1
                            _ = v14
                        } else {
                            var v15 int32 = p2
                            _ = v15
                        }
                        _ = v13
                    }
                } else {
                    if false {
                        p0 = (p0 | 255)
                        var v16 int16 = 934
                        _ = v16
                    } else {
                        if (0 != 18) {
                            v17 := GenS2{F0: (42 != 47), F1: 198}
                            var v18 int64 = -697
                            _ = v17
                            _ = v18
                        } else {
                            if true {
                                fmt.Println(fi11, p2)
                                if true {
                                    fmt.Println(p1)
                                    p2 -= (p1 - p2)
                                } else {
                                    var v19 uint = 613
                                    _ = v19
                                }
                            }
                        }
                        p2 += p2
                        v20 := []int{}
                        var v21 string = ""
                        v22 := GenS2{F0: false, F1: 42}
                        v23 := GenS3{F0: (fi12 % ((fi12) | 1)), F1: 939}
                        switch "" {
                        case "":
                            var v24 string = "a"
                            _ = v24
                        case "foo bar":
                            p1 = 1
                        case "hello":
                            if false {
                                var v25 bool = true
                                _ = v25
                            }
                        default:
                            _ = len(v20)
                        }
                        _ = v20
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    var v26 uint = 2
                    _ = v26
                }
                p2 = p2
                _ = fi12
            }
            switch "" {
            case "":
                var v27 int32 = p2
                _ = v27
            case "a":
                p2 = p2
            case "foo bar":
                var v28 uint16 = 42
                _ = v28
            default:
                if (true || false) {
                    p1 -= 42
                } else {
                    p2 += p2
                    p1 = 42
                }
            }
            for fi29 := 0; fi29 < 15; fi29++ {
                fmt.Println(fi11)
                _ = fi29
            }
            _ = fi11
        }
        var v30 uint = 510
        p0 = (p0 & p0)
        _ = v30
    }
    v31 := []int{128, 0, 1, 127}
    if len(v31) > 0 {
        v31[0] = -2147483648
    }
    fmt.Println(p2, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v31
    return -1
}

func main() {
    v1 := []int{-644}
    v2 := []string{("foo bar" + ("a" + "foo bar")), "hello"}
    v3 := []string{"", ("a" + "hello")}
    var v4 uint8 = 120
    v5 := make(map[uint32]string)
    switch 1 {
    case 42:
        fmt.Println(v4)
    case 2147483647:
        fmt.Println(v4)
    case -274:
        var v32 int8 = -73
        _ = v32
    default:
        v4 -= (^(v4))
    }
    v4 = v4
    v4 -= 128
    if (((3.14159 + (((3.14159 * 0.1) * 1.5) * 0.1)) + -2.25) < -2.25) {
        if (((3.14159 / (float64(float32(1.5)) + 1e10)) + 0.1) < -2.25) {
            if (0 == 394) {
                for fi33 := 0; fi33 < 5; fi33++ {
                    fmt.Println(fi33, fi33)
                    v34 := GenS4{F0: 1, F1: "foo bar", F2: 2}
                    if len(v1) > 0 {
                        v1[0] = (fi33 / ((fi33) | 1))
                    }
                    _ = fi33
                    _ = v34
                }
            }
            for fi35 := 0; fi35 < 4; fi35++ {
                var v36 int8 = 42
                v5[65535] = "foo bar"
                _ = fi35
                _ = v36
            }
            var v37 uint16 = 0
            _ = v37
        }
    }
    fmt.Println(v4)
    if (603 > 678) {
        var v38 int = -2147483648
        delete(v5, 760)
        v39 := GenS5{F0: 2, F1: (42 >= 1)}
        _ = v38
        _ = v39
    } else {
        v4 = (v4 + (v4 - v4))
        if true {
            var v40 float32 = 3.14159
            _ = v40
        } else {
            v4 = 42
            for fi41 := 0; fi41 < 1; fi41++ {
                for fi42 := 0; fi42 < 17; fi42++ {
                    v43 := GenS6{F0: 0, F1: -68, F2: 255}
                    switch (fi42 >> ((0) & 63)) {
                    case -9223372036854775808:
                        var v44 int16 = 1
                        _ = v44
                    case 0:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case 9223372036854775807:
                        var v45 int = fi42
                        _ = v45
                    default:
                        switch (fi42 ^ (fi42 & fi41)) {
                        case 127:
                            v4 = 2
                        }
                    }
                    _ = fi42
                    _ = v43
                }
                switch (("" + "foo bar") + ("hello" + ("a" + ("foo bar" + ("foo bar" + "\n\t"))))) {
                case "":
                    v4 += (v4 / ((v4) | 1))
                case "a":
                    var v46 int32 = 0
                    _ = v46
                default:
                    for fi47 := 0; fi47 < 11; fi47++ {
                        v4 -= v4
                        v2 = append(v2, "\n\t")
                        var v48 int64 = 227
                        _ = fi47
                        _ = v48
                    }
                }
                _ = fi41
            }
            switch 32767 {
            case -180:
                v4 -= v4
            case -13:
                fmt.Println(v4)
            case 9223372036854775807:
                var v49 uint = 65535
                _ = v49
            default:
                v50 := GenS3{F0: 1, F1: 1}
                _ = v50
            }
        }
    }
    v4 -= 0
    v51 := GenS7{F0: 1}
    if (255 > 42) {
        v4 += 255
    } else {
        v4 = 0
    }
    v52 := GenS8{F0: 42, F1: 70}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
}
