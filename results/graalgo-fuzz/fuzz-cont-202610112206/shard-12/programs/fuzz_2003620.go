package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS5 struct { F0 uint; F1 int64 }
func genFn0(p0 int32, p1 bool) int8 {
    switch "hello" {
    case "a":
        var v7 float32 = -2.25
        _ = v7
    default:
        switch "" {
        case "\n\t":
            var v8 uint = 255
            _ = v8
        case "foo bar":
            var v9 float32 = 1e10
            _ = v9
        }
    }
    p1 = (p1 || (689 >= 1))
    _ = p0
    _ = p1
    return -128
}
func genFn1() uint32 {
    var v10 uint32 = 2
    fmt.Println(v10)
    for fi11 := 0; fi11 < 11; fi11++ {
        v10 -= v10
        if (1.5 >= float32((0.1 / 0.1))) {
            v12 := []int{(fi11 * fi11), fi11}
            for fi13 := 0; fi13 < 10; fi13++ {
                v10 -= v10
                if true {
                    v10 += (v10 * (v10 - v10))
                    if (3.14159 != 1.5) {
                        v14 := GenS1{F0: (2 == 127)}
                        _ = v14
                    } else {
                        switch (-(fi11)) {
                        case -2147483648:
                            switch 42 {
                            case 30:
                                var v15 uint = 36
                                _ = v15
                            case -9223372036854775808:
                                v10 += (v10 | 0)
                            }
                        default:
                            var v16 float64 = (1.5 + float64((1.5 - 1.5)))
                            _ = v16
                        }
                        v17 := GenS2{F0: 65535}
                        switch fi13 {
                        case 42:
                            fmt.Println(v10)
                        case -1:
                            v18 := GenS1{F0: (52 < 2)}
                            _ = v18
                        }
                        _ = v17
                    }
                } else {
                    switch ("\n\t" + "foo bar") {
                    case "hello":
                        var v19 bool = true
                        fmt.Println(fi11, v19, v10)
                        _ = v19
                    default:
                        var v20 int64 = 0
                        _ = v20
                    }
                    switch "hello" {
                    case "\n\t":
                        if true {
                            v10 = 0
                        }
                    default:
                        v10 = v10
                    }
                    fmt.Println(v10)
                }
                _ = fi13
            }
            switch fi11 {
            case -131:
                var v21 uint32 = (v10 ^ 2)
                _ = v21
            case -1:
                switch -479 {
                case 127:
                    v22 := GenS1{F0: false}
                    _ = v22
                case 0:
                    var v23 int64 = -803
                    _ = v23
                default:
                    switch fi11 {
                    case 237:
                        if (167 > 213) {
                            v24 := GenS1{F0: true}
                            v10 += 121
                            _ = v24
                        } else {
                            if (2 >= 200) {
                                if true {
                                    v25 := GenS1{F0: false}
                                    v10 = v10
                                    v10 = (v10 << ((0) & 63))
                                    _ = v25
                                } else {
                                    var v26 float32 = (1.5 + -2.25)
                                    _ = v26
                                }
                                for fi27 := 0; fi27 < 10; fi27++ {
                                    if len(v12) > 0 {
                                        v12[0] = fi27
                                    }
                                    v28 := GenS2{F0: 1}
                                    var v29 uint = 255
                                    v12 = append(v12, fi11)
                                    _ = fi27
                                    _ = v28
                                    _ = v29
                                }
                                if true {
                                    var v30 uint32 = 627
                                    v10 += v10
                                    _ = v30
                                }
                            } else {
                                v10 -= (v10 + v10)
                            }
                        }
                    case 713:
                        if (false && (!false)) {
                            v10 -= 65535
                            var v31 string = "hello"
                            var v32 string = ("foo bar" + v31)
                            _ = v31
                            _ = v32
                        }
                    case -2147483648:
                        for fi33 := 0; fi33 < 5; fi33++ {
                            v34 := GenS3{F0: 255}
                            var v35 int = (fi11 + (fi33 - fi11))
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                    }
                }
            default:
                v10 = 0
            }
            _ = v12
        }
        v36 := GenS1{F0: ("\n\t" == "hello")}
        _ = fi11
        _ = v36
    }
    fmt.Println(v10)
    v37 := []int{-186, 801, 42, 4294967295}
    _ = v10
    _ = v37
    return v10
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 42
    v3 := []int{0, 127, 2147483647}
    var v4 bool = false
    v5 := []uint32{2147483647}
    var v6 int32 = 99
    v38 := GenS3{F0: 168}
    _ = len(v3)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v6, v6)
    v1[127] = 127
    for fi39 := 0; fi39 < 8; fi39++ {
        v40 := GenS2{F0: 1}
        _ = fi39
        _ = v40
    }
    if (v4 || false) {
        v41 := GenS4{F0: 597, F1: 4294967295, F2: genFn0(v2, (0 < 219))}
        _ = v41
    } else {
        v42 := GenS3{F0: 255}
        for fi43 := 0; fi43 < 15; fi43++ {
            for fi44 := 0; fi44 < 1; fi44++ {
                var v45 uint64 = 0
                fmt.Println(v45, v4)
                _ = fi44
                _ = v45
            }
            if ((!false) && ((false || v4) && (!((0 > 42) || v4)))) {
                delete(v1, -720)
            } else {
                v46 := GenS2{F0: 0}
                var v47 uint32 = 2147483647
                for fi48 := 0; fi48 < 9; fi48++ {
                    v49 := GenS4{F0: 1, F1: 2, F2: 86}
                    var v50 uint8 = 2
                    _ = fi48
                    _ = v49
                    _ = v50
                }
                _ = v46
                _ = v47
            }
            v51 := GenS5{F0: 697, F1: -673}
            _ = fi43
            _ = v51
        }
        _ = v42
    }
    v2 -= v6
    if (-676 >= 0) {
        v52 := GenS1{F0: (1.5 != 0.0)}
        _ = v52
    }
    v53 := v1[127]
    v53 -= (-(v53))
    var v54 int8 = 127
    var v55 uint16 = 32767
    var v56 bool = ((v55 - 942) != (^(v55)))
    if false {
        v4 = v56
        for fi57 := 0; fi57 < 6; fi57++ {
            v1[(fi57 + (fi57 % ((fi57) | 1)))] = v53
            _ = fi57
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v6
}
