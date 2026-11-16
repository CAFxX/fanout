package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint64; F1 int32; F2 int32 }
func genFn0() uint64 {
    for fi6 := 0; fi6 < 20; fi6++ {
        for fi7 := 0; fi7 < 6; fi7++ {
            if (42 <= int64(fi6)) {
                switch "a" {
                case "":
                    switch "hello" {
                    case "":
                        if (42 != 400) {
                            var v8 uint = 42
                            _ = v8
                        }
                    case "a":
                        var v9 uint32 = 1
                        if (!false) {
                            if (!true) {
                                v9 = v9
                                v9 -= v9
                                v10 := []int{-9223372036854775808, fi6, fi7}
                                _ = v10
                            }
                            v11 := GenS1{F0: ("" + "hello")}
                            v9 -= (v9 ^ (v9 % ((v9) | 1)))
                            _ = v11
                        }
                        _ = v9
                    default:
                        if false {
                            switch fi6 {
                            case -872:
                                v12 := GenS1{F0: "hello"}
                                _ = v12
                            case 60:
                                var v13 bool = (42 > 25)
                                _ = v13
                            }
                            if (1 > 259) {
                                v14 := []int{}
                                _ = v14
                            }
                            var v15 int = (fi6 & fi7)
                            _ = v15
                        } else {
                            if false {
                                switch "\n\t" {
                                case "":
                                    v16 := []int{(fi7 >> ((2) & 63)), (fi7 | (^(fi6)))}
                                    _ = v16
                                }
                            } else {
                                var v17 bool = (42 >= 0)
                                _ = v17
                            }
                            var v18 float32 = 0.0
                            _ = v18
                        }
                    }
                }
                v19 := GenS2{F0: -9223372036854775808, F1: 0}
                v20 := []int{(-(fi7)), fi6}
                v21 := GenS2{F0: -9223372036854775808, F1: -701}
                _ = v19
                _ = v20
                _ = v21
            }
            if true {
                if true {
                    v22 := GenS3{F0: 0, F1: -1, F2: 1}
                    _ = v22
                }
            }
            _ = fi7
        }
        _ = fi6
    }
    var v23 int64 = 9223372036854775807
    v24 := GenS2{F0: (v23 & 0), F1: v23}
    switch 9223372036854775807 {
    case 127:
        v25 := []int{128, -9223372036854775808}
        _ = v25
    case -9223372036854775808:
        if true {
            v23 += (-(v23))
        }
    default:
        switch "a" {
        case "":
            v26 := GenS3{F0: 2, F1: -2147483648, F2: -307}
            _ = v26
        case "foo bar":
            v23 += -1
        case "\n\t":
            if ((391 == 0) || (!false)) {
                for fi27 := 0; fi27 < 15; fi27++ {
                    v23 -= 9223372036854775807
                    v23 = int64(fi27)
                    v28 := GenS2{F0: 218, F1: (v23 / ((v23) | 1))}
                    _ = fi27
                    _ = v28
                }
            } else {
                switch 9223372036854775807 {
                case 32768:
                    v23 = v23
                default:
                    var v29 uint = 65535
                    _ = v29
                }
                if false {
                    switch 32767 {
                    case 127:
                        var v30 uint = 1
                        _ = v30
                    case 32767:
                        var v31 int64 = (-(v23))
                        _ = v31
                    case -1:
                        v32 := []int{-9223372036854775808}
                        _ = v32
                    default:
                        v33 := GenS2{F0: v23, F1: v23}
                        _ = v33
                    }
                    v34 := GenS3{F0: 42, F1: 181, F2: 0}
                    _ = v34
                }
                v23 = 42
            }
        }
    }
    _ = v23
    _ = v24
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{("foo bar" + ""), "foo bar", "a"}
    var v3 uint64 = 42
    var v4 uint64 = (^(v3))
    var v5 int32 = -1
    v3 -= 0
    v5 -= (-(v5))
    var v35 uint16 = 0
    fmt.Println(v3, v5, v5)
    v36 := GenS3{F0: v4, F1: (v5 % ((v5) | 1)), F2: v5}
    switch 2147483647 {
    case 2147483647:
        fmt.Println(v5)
    case -778:
        var v37 uint8 = 1
        _ = v37
    case -546:
        var v38 int16 = -1
        _ = v38
    }
    v39 := GenS3{F0: v4, F1: (v5 | v5), F2: v5}
    fmt.Println(v35)
    var v40 string = (("" + ("" + "foo bar")) + ((((("" + "") + (("" + "\n\t") + "hello")) + ("\n\t" + "hello")) + "") + "foo bar"))
    var v41 string = ("hello" + ("foo bar" + v40))
    v5 += v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
