package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 int }
func genFn0(p0 int16, p1 int32) uint32 {
    switch -753 {
    case 32767:
        var v4 int16 = (p0 - 1000)
        _ = v4
    case 128:
        p0 = (^(p0))
    default:
        p1 -= (p1 * p1)
    }
    fmt.Println(p0)
    switch 0 {
    case 42:
        var v5 uint64 = 248
        _ = v5
    default:
        if true {
            if (-2.25 <= (1e10 * 1.5)) {
                p1 -= (p1 >> ((65535) & 63))
            } else {
                switch -870 {
                case 4294967295:
                    switch -181 {
                    case 127:
                        p0 = p0
                    }
                case -183:
                    v6 := []int{32767, 455}
                    _ = v6
                case 2147483647:
                    v7 := []int{-2147483648, 853, -1}
                    _ = v7
                default:
                    v8 := []int{-1}
                    _ = v8
                }
            }
            for fi9 := 0; fi9 < 20; fi9++ {
                p1 -= (^(p1))
                _ = fi9
            }
        } else {
            if false {
                v10 := []int{}
                if false {
                    switch -9223372036854775808 {
                    case 42:
                        var v11 string = "\n\t"
                        _ = v11
                    case 32767:
                        v10 = append(v10, -1)
                    default:
                        switch "a" {
                        case "":
                            if false {
                                v10 = append(v10, -677)
                                var v12 uint32 = 42
                                fmt.Println(v12, v12, p0)
                                _ = v12
                            }
                        }
                    }
                    p1 = (p1 - p1)
                }
                _ = v10
            }
        }
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        fmt.Println(fi13, p1, p0)
        p1 = (-(p1))
        var v14 int = fi13
        _ = fi13
        _ = v14
    }
    var v15 int32 = (^(p1))
    _ = p0
    _ = p1
    _ = v15
    return 606
}

func main() {
    var v1 int = -1
    var v2 uint32 = 2147483648
    var v3 uint8 = 2
    v16 := GenS1{F0: 42, F1: 32767, F2: "hello"}
    v1 += 0
    v17 := []int{v1}
    if (((true || false) || (!(1 != -9223372036854775808))) && (1 > 42)) {
        _ = len(v17)
        fmt.Println(v3)
    }
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    fmt.Println(v2)
    if false {
        var v18 float64 = -2.25
        if len(v17) > 0 {
            v17[0] = v1
        }
        if (v3 == 205) {
            switch "hello" {
            case "foo bar":
                if true {
                    var v19 uint32 = genFn0(-1, 633)
                    _ = v19
                }
            case "\n\t":
                for fi20 := 0; fi20 < 4; fi20++ {
                    v3 = 255
                    var v21 uint = 2
                    v22 := GenS2{F0: -2147483648}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            case "a":
                var v23 uint32 = genFn0(0, -2147483648)
                _ = v23
            }
            for fi24 := 0; fi24 < 6; fi24++ {
                for fi25 := 0; fi25 < 15; fi25++ {
                    var v26 uint = 255
                    switch "" {
                    case "a":
                        v3 = (v3 / ((v3) | 1))
                    case "hello":
                        v3 = v3
                    case "foo bar":
                        if (!false) {
                            v18 = v18
                        } else {
                            if false {
                                fmt.Println(fi24)
                            } else {
                                var v27 float32 = 0.0
                                if len(v17) > 0 {
                                    v17[0] = -358
                                }
                                v3 = v3
                                var v28 float32 = v27
                                var v29 int32 = 2147483647
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                        }
                    }
                    _ = fi25
                    _ = v26
                }
                v3 = 1
                for fi30 := 0; fi30 < 14; fi30++ {
                    fmt.Println(v1, v3)
                    fmt.Println(v2, v18, v18)
                    v17 = append(v17, v1)
                    _ = fi30
                }
                _ = fi24
            }
        }
        _ = v18
    } else {
        if len(v17) > 0 {
            v17[0] = v1
        }
        if (152 < 2) {
            if (0 < 762) {
                v1 = 42
            } else {
                fmt.Println(v3)
            }
            for fi31 := 0; fi31 < 13; fi31++ {
                var v32 uint16 = 42
                _ = fi31
                _ = v32
            }
        } else {
            var v33 string = "foo bar"
            _ = v33
        }
        v34 := GenS1{F0: (^(v2)), F1: (v1 * v1), F2: ("" + ("\n\t" + ("foo bar" + ((("" + "hello") + (("\n\t" + "") + "a")) + ""))))}
        _ = v34
    }
    var v35 string = (("hello" + (("foo bar" + (("" + ("\n\t" + ((((("hello" + "foo bar") + "a") + "hello") + "hello") + "hello"))) + "a")) + ("a" + "a"))) + "hello")
    v36 := GenS3{F0: -115, F1: v1}
    _ = len(v17)
    _ = len(v17)
    var v37 string = v35
    var v38 int8 = -18
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
}
