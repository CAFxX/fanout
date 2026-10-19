package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint32 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 uint }

func main() {
    var v1 uint8 = 2
    var v2 bool = false
    var v3 bool = true
    var v4 string = ("foo bar" + ("hello" + "a"))
    var v5 uint64 = 0
    for fi6 := 0; fi6 < 13; fi6++ {
        v3 = (v2 && (!((v5 ^ v5) >= (v5 << ((128) & 63)))))
        v5 = v5
        v4 = v4
        _ = fi6
    }
    v7 := []int{-302, 535, int(v5)}
    fmt.Println(v5, v1)
    v4 = (v4 + "")
    if (v2 || v2) {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    var v8 bool = v3
    switch 42 {
    case -24:
        v5 -= v5
    case -2147483648:
        if v2 {
            fmt.Println(v8, v3, v5)
            var v9 int = 127
            _ = v9
        } else {
            v10 := GenS1{F0: 65535, F1: 4294967295, F2: 1}
            v2 = v2
            for fi11 := 0; fi11 < 18; fi11++ {
                _ = len(v7)
                v12 := GenS1{F0: 255, F1: 1, F2: 424}
                if (true && v2) {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                } else {
                    switch 1 {
                    case 9223372036854775807:
                        for fi13 := 0; fi13 < 6; fi13++ {
                            v14 := GenS1{F0: 42, F1: 0, F2: 4294967295}
                            switch v4 {
                            case "hello":
                                if v3 {
                                    v5 -= (^(v5))
                                    v7 = append(v7, fi13)
                                    var v15 bool = (127 == 0)
                                    _ = v15
                                } else {
                                    v4 = ""
                                    v16 := GenS2{F0: (false && (((0 >= 42) && v3) && true)), F1: v2}
                                    var v17 bool = v2
                                    _ = v16
                                    _ = v17
                                }
                            }
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            _ = fi13
                            _ = v14
                        }
                    }
                    var v18 bool = ((1e10 / 3.14159) < 1.5)
                    v19 := GenS3{F0: int8(fi11), F1: fi11}
                    _ = v18
                    _ = v19
                }
                _ = fi11
                _ = v12
            }
            _ = v10
        }
    case -9223372036854775808:
        var v20 int = 0
        _ = v20
    }
    v2 = v8
    _ = len(v7)
    if (814 != 32767) {
        fmt.Println(v1, v4, v4)
        v3 = (false && ((3.14159 * ((0.1 - ((0.1 - (-2.25 * (0.1 + (1.5 * 0.1)))) + (((1.5 / 0.1) - (float32(1e10) + ((float32((0.1 / 1.0)) * 3.14159) - 3.14159))) / (((1.5 + (-2.25 / 1e10)) + 1e10) * -2.25)))) * (1.5 - 0.1))) <= (3.14159 - (1.5 / 1.0))))
    } else {
        if (v1 > (v1 % ((v1) | 1))) {
            v1 -= 145
        } else {
            switch "hello" {
            case "":
                if v8 {
                    fmt.Println(v3)
                    fmt.Println(v2)
                } else {
                    for fi21 := 0; fi21 < 12; fi21++ {
                        fmt.Println(v1)
                        _ = fi21
                    }
                    if (v8 || v3) {
                        if false {
                            fmt.Println(v2, v5)
                            for fi22 := 0; fi22 < 4; fi22++ {
                                v23 := GenS4{F0: 756, F1: (false && false), F2: 255}
                                var v24 int8 = 58
                                var v25 uint16 = 0
                                _ = fi22
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                        } else {
                            v26 := GenS1{F0: 652, F1: 2, F2: 2}
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 3; fi27++ {
                            fmt.Println(v4, v1, v5)
                            _ = fi27
                        }
                        fmt.Println(v3, v8)
                    } else {
                        var v28 float32 = 0.0
                        _ = v28
                    }
                }
            case "hello":
                v4 = v4
            case "\n\t":
                v3 = false
            }
            if false {
                for fi29 := 0; fi29 < 8; fi29++ {
                    v30 := GenS5{F0: 2147483647, F1: 0}
                    _ = fi29
                    _ = v30
                }
                if false {
                    v4 = "a"
                    v5 += 18446744073709551615
                }
                v2 = (v2 && (42 <= 0))
            } else {
                if v2 {
                    v1 = (v1 << ((65535) & 63))
                }
                _ = len(v7)
                if (100 > 285) {
                    v1 -= (v1 / (((v1 % ((v1) | 1))) | 1))
                    for fi31 := 0; fi31 < 2; fi31++ {
                        var v32 string = v4
                        if (v5 == 4294967295) {
                            var v33 int = fi31
                            _ = v33
                        }
                        _ = fi31
                        _ = v32
                    }
                }
            }
            var v34 float64 = 0.0
            _ = v34
        }
        v35 := GenS2{F0: (true && v3), F1: v2}
        if len(v7) > 0 {
            v7[0] = -610
        }
        _ = v35
    }
    if len(v7) > 0 {
        v7[0] = 731
    }
    v36 := GenS6{F0: 2}
    var v37 int32 = -689
    v38 := GenS2{F0: false, F1: (978 == 2147483647)}
    for fi39 := 0; fi39 < 14; fi39++ {
        fmt.Println(v4, v1)
        _ = fi39
    }
    var v40 bool = (v8 || v8)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v40
    _ = v5
    _ = v7
    _ = v8
}
