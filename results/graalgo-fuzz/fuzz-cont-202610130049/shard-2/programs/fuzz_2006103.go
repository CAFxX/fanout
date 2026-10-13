package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 uint64; F1 int; F2 string }
type GenS5 struct { F0 uint64 }
func genFn0(p0 int32, p1 int64, p2 uint8) int64 {
    for fi5 := 0; fi5 < 20; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            var v7 float64 = 3.14159
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    if ((p1 % ((p1) | 1)) > (p1 | 42)) {
        fmt.Println(p0, p0, p2)
        p2 = (p2 * (p2 >> ((2) & 63)))
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        p2 -= p2
        v9 := []int{fi8}
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: "foo bar", F1: 4294967295}
    v11 := []int{127, int(p2), 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return p1
}
func genFn1(p0 int64, p1 uint8) int8 {
    v12 := GenS2{F0: 42, F1: false}
    for fi13 := 0; fi13 < 10; fi13++ {
        p0 = p0
        _ = fi13
    }
    p0 = (p0 | p0)
    _ = p0
    _ = p1
    _ = v12
    return 1
}

func main() {
    var v1 int8 = 100
    var v2 int8 = 89
    var v3 uint = 1
    v4 := []int{32768, 0, 32768}
    fmt.Println(v1, v3)
    if true {
        var v14 uint64 = 412
        _ = v14
    } else {
        if true {
            if len(v4) > 0 {
                v4[0] = 776
            }
            switch "" {
            case "hello":
                switch -735 {
                case 9223372036854775807:
                    v1 += (v1 / ((100) | 1))
                }
            default:
                var v15 float64 = 3.14159
                _ = v15
            }
            if len(v4) > 0 {
                v4[0] = 42
            }
        } else {
            v2 -= v2
        }
    }
    v16 := GenS3{F0: 2147483647, F1: "a"}
    switch "hello" {
    case "hello":
        fmt.Println(v1, v3)
    case "\n\t":
        for fi17 := 0; fi17 < 20; fi17++ {
            v18 := GenS4{F0: 42, F1: (fi17 | fi17), F2: "foo bar"}
            v2 -= -64
            if ((0.0 + 3.14159) > 0.1) {
                var v19 uint8 = 162
                if len(v4) > 0 {
                    v4[0] = (fi17 >> ((v3) & 63))
                }
                _ = v19
            } else {
                v2 = 127
                v3 = v3
                switch fi17 {
                case -99:
                    var v20 int8 = (-(v2))
                    _ = v20
                case -541:
                    var v21 uint64 = 1
                    _ = v21
                default:
                    var v22 uint8 = 1
                    _ = v22
                }
            }
            _ = fi17
            _ = v18
        }
    case "foo bar":
        _ = len(v4)
    default:
        v23 := GenS5{F0: 18446744073709551615}
        _ = v23
    }
    var v24 int16 = 1
    _ = len(v4)
    for fi25 := 0; fi25 < 16; fi25++ {
        fmt.Println(v2, v1, v3)
        if ("foo bar" >= "") {
            if true {
                fmt.Println(v3, v2)
                if (42 == (v24 - (v24 + -758))) {
                    for fi26 := 0; fi26 < 18; fi26++ {
                        if (-15 <= -1) {
                            fmt.Println(fi25)
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v27 uint = (v3 | v3)
                            switch "foo bar" {
                            case "foo bar":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case "\n\t":
                                v28 := GenS1{F0: ((("" + ("foo bar" + "\n\t")) + "") + "a"), F1: 10}
                                _ = v28
                            }
                            _ = v27
                        }
                        var v29 bool = (false && true)
                        _ = fi26
                        _ = v29
                    }
                    var v30 uint8 = 255
                    fmt.Println(v3, v1, v2)
                    _ = v30
                } else {
                    var v31 uint64 = 0
                    _ = v31
                }
            }
            v2 += 127
            for fi32 := 0; fi32 < 2; fi32++ {
                var v33 float32 = 0.1
                var v34 int64 = 102
                _ = fi32
                _ = v33
                _ = v34
            }
        }
        v3 -= 906
        _ = fi25
    }
    for fi35 := 0; fi35 < 3; fi35++ {
        v2 -= (-(v2))
        var v36 int = fi35
        _ = fi35
        _ = v36
    }
    var v37 string = ("hello" + "\n\t")
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v37)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v37
    _ = v4
}
