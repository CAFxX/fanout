package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int32 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 int32; F1 uint32; F2 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS6 struct { F0 uint; F1 string; F2 int64 }
type GenS7 struct { F0 uint }
func genFn0(p0 string, p1 int32, p2 int16) int16 {
    p0 = "hello"
    p0 = p0
    var v5 float64 = ((-2.25 * 0.1) + 0.0)
    v6 := []int{}
    fmt.Println(p2, v5, v5)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint, p1 string) int64 {
    fmt.Println(p0)
    var v7 int32 = 1
    fmt.Println(p1)
    v8 := []int{2147483647, 2147483647, 9223372036854775807, -535}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 988
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"\n\t", "\n\t", "\n\t"}
    v3 := []string{(("a" + (("\n\t" + "hello") + ("foo bar" + (("foo bar" + ("a" + "a")) + "a")))) + ("" + "foo bar"))}
    var v4 float64 = ((3.14159 / -2.25) / 3.14159)
    if (!true) {
        if (!(42 >= 577)) {
            if (!true) {
                switch "foo bar" {
                case "\n\t":
                    v9 := GenS1{F0: 32767, F1: 2147483647, F2: -1}
                    _ = v9
                case "foo bar":
                    var v10 int64 = 0
                    _ = v10
                case "":
                    var v11 uint64 = 1
                    _ = v11
                }
            } else {
                var v12 int = -327
                _ = v12
            }
        } else {
            switch "a" {
            case "foo bar":
                v13 := GenS1{F0: 127, F1: 667, F2: 2147483647}
                _ = v13
            case "hello":
                v14 := GenS2{F0: 1, F1: true}
                _ = v14
            case "a":
                for fi15 := 0; fi15 < 2; fi15++ {
                    var v16 string = "\n\t"
                    for fi17 := 0; fi17 < 2; fi17++ {
                        var v18 float32 = 1.5
                        var v19 uint16 = 0
                        var v20 int16 = 1000
                        _ = fi17
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 3; fi21++ {
                        v16 = "a"
                        fmt.Println(v16, fi21, fi15)
                        _ = fi21
                    }
                    _ = fi15
                    _ = v16
                }
            }
            v4 += 1e10
            fmt.Println(v4)
        }
        for fi22 := 0; fi22 < 11; fi22++ {
            for fi23 := 0; fi23 < 10; fi23++ {
                var v24 int = fi23
                switch 0 {
                case 32768:
                    v25 := GenS3{F0: -1, F1: 2, F2: false}
                    _ = v25
                case -9223372036854775808:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case 0:
                    fmt.Println(v24, fi23)
                }
                _ = fi23
                _ = v24
            }
            var v26 int = (fi22 - 0)
            for fi27 := 0; fi27 < 7; fi27++ {
                _ = len(v3)
                var v28 uint = 4294967295
                v4 = v4
                _ = fi27
                _ = v28
            }
            _ = fi22
            _ = v26
        }
    } else {
        var v29 int = -562
        var v30 uint64 = 18446744073709551615
        switch "hello" {
        case "hello":
            var v31 float32 = 1e10
            _ = v31
        case "a":
            var v32 int32 = 42
            _ = v32
        default:
            if true {
                if false {
                    _ = len(v2)
                    if false {
                        if false {
                            v4 += v4
                            v4 = ((v4 + 1.5) / (v4 + ((v4 + (v4 / (((v4 / (v4 - ((v4 - (v4 / v4)) + v4))) - v4) / float64(float32(v4))))) - v4)))
                        }
                        for fi33 := 0; fi33 < 13; fi33++ {
                            v34 := GenS1{F0: 32768, F1: (^(fi33)), F2: 171}
                            _ = fi33
                            _ = v34
                        }
                    }
                    v1[(v29 % ((v29) | 1))] = (0 != 100)
                } else {
                    v35 := GenS2{F0: uint8(v30), F1: false}
                    v36 := GenS4{F0: 2}
                    fmt.Println(v30, v30, v29)
                    _ = v35
                    _ = v36
                }
            } else {
                fmt.Println(v30, v4, v4)
                v4 += (v4 * ((-2.25 / (-2.25 / v4)) + (v4 / v4)))
            }
        }
        _ = v29
        _ = v30
    }
    fmt.Println(v4)
    var v37 int8 = -53
    _ = len(v1)
    fmt.Println(v37, v37)
    fmt.Println(v37, v4)
    for fi38 := 0; fi38 < 4; fi38++ {
        fmt.Println(v4, v4)
        v39 := GenS5{F0: 0, F1: 2, F2: 255}
        _ = fi38
        _ = v39
    }
    v4 = v4
    for fi40 := 0; fi40 < 9; fi40++ {
        var v41 string = ("\n\t" + ("\n\t" + ""))
        if false {
            v42 := GenS5{F0: -1, F1: 32767, F2: 125}
            _ = v42
        } else {
            v41 = (v41 + ((v41 + "foo bar") + v41))
            if (false && true) {
                for fi43 := 0; fi43 < 19; fi43++ {
                    fmt.Println(v4, fi43)
                    var v44 uint = 636
                    _ = fi43
                    _ = v44
                }
                switch fi40 {
                case -478:
                    v4 = float64(1e10)
                default:
                    var v45 string = v41
                    _ = v45
                }
            } else {
                v46 := GenS6{F0: 448, F1: v41, F2: 459}
                _ = v46
            }
            fmt.Println(fi40)
        }
        _ = fi40
        _ = v41
    }
    v47 := GenS7{F0: 42}
    if true {
        var v48 int = -40
        for fi49 := 0; fi49 < 20; fi49++ {
            fmt.Println(v4, fi49)
            _ = fi49
        }
        _ = v48
    } else {
        var v50 int16 = 0
        for fi51 := 0; fi51 < 17; fi51++ {
            v37 -= (-(v37))
            _ = fi51
        }
        if ((504 >= 111) || ((v50 ^ v50) > 342)) {
            v52 := GenS3{F0: -1, F1: 377, F2: ((-32768 >= (v50 % ((v50) | 1))) && false)}
            _ = v52
        }
        _ = v50
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v47
}
