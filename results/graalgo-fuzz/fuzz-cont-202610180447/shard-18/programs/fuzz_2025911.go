package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS4 struct { F0 string }
func genFn0(p0 int, p1 int) uint32 {
    switch 32767 {
    case 128:
        if false {
            for fi6 := 0; fi6 < 8; fi6++ {
                switch fi6 {
                case 0:
                    fmt.Println(fi6, fi6)
                default:
                    switch "" {
                    case "":
                        for fi7 := 0; fi7 < 20; fi7++ {
                            if true {
                                v8 := []int{(fi6 + (fi6 & fi6)), (fi6 << ((405) & 63))}
                                var v9 int8 = -128
                                p1 += -1
                                v9 += v9
                                switch p1 {
                                case 4294967295:
                                    fmt.Println(fi6, fi7)
                                case 816:
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                case 32768:
                                    p1 = -2147483648
                                default:
                                    if (false && false) {
                                        var v10 uint8 = 227
                                        var v11 uint64 = 256
                                        _ = v10
                                        _ = v11
                                    } else {
                                        switch ("foo bar" + ("" + "")) {
                                        case "":
                                            if ((1.5 + 1.5) != 1e10) {
                                                switch (("hello" + "hello") + "a") {
                                                case "hello":
                                                    var v12 string = "hello"
                                                    _ = v12
                                                default:
                                                    var v13 uint16 = 32767
                                                    _ = v13
                                                }
                                                v9 += v9
                                                p1 = 9223372036854775807
                                            } else {
                                                v14 := GenS1{F0: 0, F1: 1, F2: 1}
                                                v9 = (v9 ^ 34)
                                                _ = v14
                                            }
                                        default:
                                            var v15 uint64 = 0
                                            _ = v15
                                        }
                                        var v16 string = "\n\t"
                                        fmt.Println(fi7)
                                        _ = v16
                                    }
                                }
                                fmt.Println(v9, p1, fi7)
                                _ = v8
                                _ = v9
                            }
                            var v17 string = "hello"
                            _ = fi7
                            _ = v17
                        }
                    case "\n\t":
                        v18 := GenS1{F0: 127, F1: 689, F2: 662}
                        _ = v18
                    }
                }
                _ = fi6
            }
            p1 = p0
        } else {
            v19 := GenS2{F0: (-(p1)), F1: 234}
            _ = v19
        }
    }
    switch -560 {
    case -1:
        var v20 uint = 255
        _ = v20
    case 4294967295:
        switch (p0 ^ (p0 * p0)) {
        case -744:
            v21 := GenS2{F0: p1, F1: 0}
            _ = v21
        }
    default:
        var v22 uint8 = 1
        _ = v22
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]string)
    v3 := []uint32{2147483647, 0}
    var v4 int = 127
    var v5 int = v4
    switch 32767 {
    case -9223372036854775808:
        for fi23 := 0; fi23 < 16; fi23++ {
            if len(v3) > 0 {
                v3[0] = 696
            }
            switch (((("hello" + "a") + "\n\t") + (("" + "") + "a")) + ("" + (("foo bar" + ("\n\t" + "foo bar")) + ("hello" + "hello")))) {
            case "":
                v24 := GenS3{F0: true, F1: 0, F2: 2}
                _ = v24
            }
            _ = fi23
        }
    default:
        delete(v1, (^(v5)))
    }
    if (true || (!false)) {
        var v25 uint8 = 255
        v26 := GenS1{F0: -113, F1: 42, F2: 255}
        delete(v2, v4)
        _ = v25
        _ = v26
    } else {
        v4 -= 0
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v27 uint64 = 42
    for fi28 := 0; fi28 < 9; fi28++ {
        v5 = (v5 + (v5 >> ((0) & 63)))
        _ = len(v3)
        _ = fi28
    }
    fmt.Println(v27, v5, v4)
    v5 = 128
    var v29 int64 = 724
    v5 += (v4 >> ((1) & 63))
    v29 = (v29 | 42)
    _ = len(v3)
    for fi30 := 0; fi30 < 8; fi30++ {
        var v31 float64 = 3.14159
        fmt.Println(v29, v31)
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 14; fi32++ {
        switch int(v27) {
        case 900:
            var v33 uint8 = 128
            _ = v33
        default:
            v34 := GenS2{F0: (v4 ^ fi32), F1: 3}
            _ = v34
        }
        for fi35 := 0; fi35 < 17; fi35++ {
            fmt.Println(v4, fi32, v29)
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            v36 := GenS4{F0: "foo bar"}
            _ = fi35
            _ = v36
        }
        var v37 string = ""
        _ = fi32
        _ = v37
    }
    fmt.Println(v5, v4)
    v29 -= v29
    v38 := GenS2{F0: v5, F1: 128}
    var v39 int = (v4 - 128)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
