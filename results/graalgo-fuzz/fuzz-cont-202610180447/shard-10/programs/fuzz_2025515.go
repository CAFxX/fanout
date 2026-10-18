package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int8; F2 uint8 }
type GenS3 struct { F0 bool }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = ((0.1 / -2.25) - 1e10)
    v3 := make(map[uint32]uint32)
    var v4 uint64 = 18446744073709551615
    var v5 uint64 = 115
    v6 := make(map[int]bool)
    v7 := v3[1]
    if true {
        if false {
            _ = len(v1)
        } else {
            switch "" {
            case "foo bar":
                if false {
                    v8 := []int{}
                    _ = v8
                }
            case "hello":
                if false {
                    if (0 <= 462) {
                        v9 := GenS1{F0: 4294967295}
                        for fi10 := 0; fi10 < 9; fi10++ {
                            v11 := []int{fi10, (fi10 % ((127) | 1)), 9223372036854775807, (fi10 - fi10)}
                            switch fi10 {
                            case 1:
                                switch "hello" {
                                case "\n\t":
                                    fmt.Println(fi10, v2)
                                case "":
                                    v4 += 166
                                default:
                                    switch "a" {
                                    case "\n\t":
                                        var v12 int32 = 42
                                        _ = v12
                                    }
                                }
                            case 42:
                                for fi13 := 0; fi13 < 6; fi13++ {
                                    var v14 int16 = 1000
                                    if len(v11) > 0 {
                                        v11[0] = (fi10 + -9223372036854775808)
                                    }
                                    var v15 string = ""
                                    v7 = v7
                                    _ = fi13
                                    _ = v14
                                    _ = v15
                                }
                            case 128:
                                v16 := GenS1{F0: 2}
                                _ = v16
                            }
                            _ = fi10
                            _ = v11
                        }
                        var v17 float64 = float64((v2 * v2))
                        _ = v17
                        _ = v9
                    }
                    switch "foo bar" {
                    case "\n\t":
                        v5 = (^(v4))
                    default:
                        v4 = v4
                    }
                    fmt.Println(v5, v4, v4)
                }
            }
            v4 += 0
        }
        for fi18 := 0; fi18 < 13; fi18++ {
            v19 := GenS1{F0: uint(v5)}
            var v20 int16 = 260
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        v7 = (v7 | (v7 * v7))
        var v22 int = fi21
        for fi23 := 0; fi23 < 12; fi23++ {
            if false {
                v24 := GenS1{F0: uint(v4)}
                var v25 string = (("\n\t" + ((("hello" + ("hello" + ("foo bar" + "hello"))) + (("foo bar" + "") + (("\n\t" + "foo bar") + ("hello" + "")))) + "foo bar")) + "foo bar")
                if true {
                    v26 := v1[(fi23 ^ (fi23 * (fi21 << ((993) & 63))))]
                    _ = v26
                } else {
                    v27 := []int{(v22 + fi23), (fi23 << ((uint(v4)) & 63))}
                    v7 -= v7
                    v4 = 243
                    _ = v27
                }
                _ = v24
                _ = v25
            }
            var v28 int8 = 127
            _ = fi23
            _ = v28
        }
        _ = fi21
        _ = v22
    }
    v29 := []int{0, 32767, -1}
    if true {
        v30 := GenS1{F0: 1}
        _ = v30
    }
    for fi31 := 0; fi31 < 16; fi31++ {
        v5 += 533
        v4 -= v4
        _ = fi31
    }
    fmt.Println(v2, v2, v5)
    if len(v29) >= 2 {
        v29 = v29[1:len(v29)]
    }
    v32 := GenS1{F0: 756}
    for fi33 := 0; fi33 < 8; fi33++ {
        var v34 float32 = 1e10
        switch ("a" + "foo bar") {
        case "foo bar":
            var v35 float64 = (-2.25 * 0.0)
            _ = v35
        default:
            if true {
                v36 := GenS2{F0: 9223372036854775807, F1: 37, F2: 2}
                _ = v36
            }
        }
        var v37 int = fi33
        _ = fi33
        _ = v34
        _ = v37
    }
    v4 -= v4
    for fi38 := 0; fi38 < 9; fi38++ {
        delete(v6, fi38)
        fmt.Println(fi38, v7, fi38)
        if false {
            var v39 int32 = 1
            _ = v39
        } else {
            v40 := GenS1{F0: uint(v7)}
            v41 := GenS3{F0: (-32768 >= -36)}
            fmt.Println(v2, fi38, fi38)
            _ = v40
            _ = v41
        }
        _ = fi38
    }
    for fi42 := 0; fi42 < 7; fi42++ {
        for fi43 := 0; fi43 < 19; fi43++ {
            switch ((("a" + ("a" + "\n\t")) + "") + "foo bar") {
            case "a":
                v2 += v2
            default:
                fmt.Println(fi43, v2, fi42)
            }
            v44 := GenS1{F0: 78}
            _ = fi43
            _ = v44
        }
        if ((v2 < v2) && true) {
            fmt.Println(v2)
        }
        var v45 uint32 = (v7 + (v7 - (v7 % ((2) | 1))))
        _ = fi42
        _ = v45
    }
    v2 -= v2
    v5 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
