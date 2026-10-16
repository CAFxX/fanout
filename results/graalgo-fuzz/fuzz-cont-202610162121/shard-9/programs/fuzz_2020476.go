package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 uint; F2 bool }
func genFn0(p0 string) uint8 {
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := []int{}
        for fi9 := 0; fi9 < 13; fi9++ {
            p0 = ((p0 + "foo bar") + "hello")
            v10 := GenS1{F0: "\n\t"}
            switch (fi7 & (-(fi7))) {
            case -2147483648:
                var v11 uint8 = 127
                _ = v11
            default:
                var v12 float32 = 3.14159
                _ = v12
            }
            _ = fi9
            _ = v10
        }
        var v13 uint32 = 599
        _ = fi7
        _ = v13
        _ = v8
    }
    if (26 <= 91) {
        var v14 uint64 = 0
        var v15 bool = ("" >= (p0 + p0))
        _ = v14
        _ = v15
    }
    var v16 uint32 = 644
    if false {
        if false {
            var v17 int16 = 0
            v18 := GenS1{F0: ((p0 + p0) + "foo bar")}
            var v19 uint16 = 723
            _ = v17
            _ = v18
            _ = v19
        }
        p0 = ("hello" + (((p0 + (((((((p0 + (p0 + ("" + p0))) + p0) + (p0 + p0)) + "a") + "a") + p0) + p0)) + p0) + p0))
        fmt.Println(v16)
    } else {
        var v20 uint = 4294967295
        var v21 string = "a"
        _ = v20
        _ = v21
    }
    _ = p0
    _ = v16
    return 116
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 77
    v3 := make(map[uint32]uint32)
    var v4 int32 = -2147483648
    var v5 uint16 = 65535
    v6 := []int{32768, -181, 9223372036854775807}
    if true {
        switch 42 {
        case -2147483648:
            var v22 string = "foo bar"
            _ = v22
        case -1:
            if true {
                var v23 uint8 = 110
                _ = v23
            } else {
                if true {
                    var v24 int = 0
                    _ = v24
                } else {
                    for fi25 := 0; fi25 < 7; fi25++ {
                        if false {
                            if false {
                                var v26 int64 = 431
                                if len(v6) > 0 {
                                    v6[0] = 42
                                }
                                _ = v26
                            }
                            v27 := GenS1{F0: "foo bar"}
                            _ = v27
                        } else {
                            if true {
                                v3[4294967295] = 915
                                var v28 uint64 = 383
                                _ = v28
                            } else {
                                fmt.Println(v4, v5)
                                v5 += 229
                                fmt.Println(v2, v4)
                            }
                            if false {
                                v29 := GenS1{F0: ("" + "a")}
                                if true {
                                    fmt.Println(v2, v2)
                                    for fi30 := 0; fi30 < 19; fi30++ {
                                        var v31 uint32 = 42
                                        fmt.Println(fi25, v4, fi30)
                                        _ = fi30
                                        _ = v31
                                    }
                                }
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                _ = v29
                            } else {
                                for fi32 := 0; fi32 < 17; fi32++ {
                                    fmt.Println(v4)
                                    v2 = -128
                                    v33 := v1[0]
                                    _ = fi32
                                    _ = v33
                                }
                                switch fi25 {
                                case -1:
                                    for fi34 := 0; fi34 < 2; fi34++ {
                                        v35 := GenS2{F0: (^(v2))}
                                        switch (-(fi34)) {
                                        case -391:
                                            var v36 uint = 2
                                            v37 := GenS1{F0: (((("hello" + "") + "hello") + (("\n\t" + "\n\t") + ((("a" + "a") + "\n\t") + ("a" + "foo bar")))) + "hello")}
                                            var v38 uint32 = uint32(v5)
                                            switch 32768 {
                                            case 4294967295:
                                                v39 := GenS2{F0: (v2 ^ 38)}
                                                _ = v39
                                            case 0:
                                                switch fi34 {
                                                case 614:
                                                    v5 -= v5
                                                default:
                                                    if (!true) {
                                                        v36 -= 0
                                                    } else {
                                                        v40 := GenS2{F0: (v2 % ((0) | 1))}
                                                        _ = v40
                                                    }
                                                }
                                            case 2147483647:
                                                v41 := GenS3{F0: 65535}
                                                _ = v41
                                            default:
                                                v42 := GenS2{F0: 38}
                                                _ = v42
                                            }
                                            _ = v36
                                            _ = v37
                                            _ = v38
                                        case -801:
                                            var v43 int64 = -633
                                            _ = v43
                                        default:
                                            fmt.Println(fi34, v4)
                                        }
                                        if true {
                                            switch "a" {
                                            case "foo bar":
                                                v44 := GenS3{F0: 0}
                                                _ = v44
                                            case "\n\t":
                                                v45 := GenS4{F0: 42, F1: 42, F2: false}
                                                _ = v45
                                            }
                                            var v46 float64 = 1.5
                                            _ = v46
                                        }
                                        _ = fi34
                                        _ = v35
                                    }
                                case 127:
                                    fmt.Println(v4, v2)
                                default:
                                    var v47 int8 = (v2 >> ((2) & 63))
                                    _ = v47
                                }
                                v1[42] = ("a" + ("a" + "\n\t"))
                            }
                            for fi48 := 0; fi48 < 16; fi48++ {
                                v2 += v2
                                v3[0] = 2147483648
                                _ = fi48
                            }
                        }
                        var v49 int16 = 742
                        _ = fi25
                        _ = v49
                    }
                    fmt.Println(v5)
                }
            }
        case -304:
            v2 = (v2 % (((-(v2))) | 1))
        }
        var v50 uint = 177
        _ = v50
    }
    var v51 bool = (v5 != v5)
    for fi52 := 0; fi52 < 16; fi52++ {
        var v53 uint8 = 127
        _ = fi52
        _ = v53
    }
    fmt.Println(v5, v4)
    v2 = (v2 + v2)
    v4 -= (v4 & (^(v4)))
    var v54 uint64 = 0
    var v55 uint64 = 574
    var v56 uint16 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v54
    _ = v55
    _ = v56
    _ = v6
}
