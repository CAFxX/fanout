package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int32 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int; F1 int; F2 uint8 }
type GenS7 struct { F0 string; F1 uint }
func genFn0(p0 uint, p1 uint32) int64 {
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := []int{fi6, (fi6 - (fi6 | (fi6 << ((p0) & 63)))), -1}
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi6
        _ = v7
    }
    switch 32767 {
    case -143:
        fmt.Println(p0)
    case 32:
        switch 0 {
        case 32768:
            var v8 string = "hello"
            _ = v8
        case 859:
            var v9 int64 = -9223372036854775808
            _ = v9
        case -184:
            switch (("a" + "a") + ((("hello" + "") + "") + ("a" + ""))) {
            case "a":
                var v10 float32 = 3.14159
                _ = v10
            case "foo bar":
                v11 := GenS1{F0: 2, F1: 1, F2: 2147483647}
                _ = v11
            case "hello":
                v12 := GenS1{F0: 1, F1: -57, F2: -490}
                _ = v12
            default:
                var v13 int32 = -317
                _ = v13
            }
        default:
            fmt.Println(p1, p1)
        }
    }
    p0 = 18446744073709551615
    v14 := []int{-9223372036854775808, 42, -9223372036854775808, 1}
    _ = p0
    _ = p1
    _ = v14
    return -321
}

func main() {
    var v1 int16 = -610
    v2 := make(map[string]int)
    var v3 string = "hello"
    var v4 uint64 = 1
    v5 := []string{"hello", "foo bar"}
    var v15 int8 = 100
    delete(v2, ((("a" + "a") + v3) + ("" + v3)))
    v15 -= (v15 * (-(v15)))
    fmt.Println(v4, v15, v15)
    v16 := GenS2{F0: 1, F1: "foo bar"}
    for fi17 := 0; fi17 < 15; fi17++ {
        if false {
            switch fi17 {
            case 128:
                var v18 string = ("a" + v3)
                _ = v18
            case 756:
                v19 := GenS2{F0: (v4 >> ((0) & 63)), F1: "foo bar"}
                _ = v19
            case -1:
                v1 = (^(v1))
            default:
                if (809 < 42) {
                    v4 -= (v4 + v4)
                    if false {
                        v20 := GenS2{F0: (v4 & v4), F1: v3}
                        var v21 int16 = (^(v1))
                        if false {
                            for fi22 := 0; fi22 < 9; fi22++ {
                                v21 += v21
                                _ = fi22
                            }
                        }
                        _ = v20
                        _ = v21
                    } else {
                        fmt.Println(v4)
                    }
                    for fi23 := 0; fi23 < 12; fi23++ {
                        _ = len(v2)
                        if true {
                            switch 9223372036854775807 {
                            case -9223372036854775808:
                                fmt.Println(fi17, v15, v1)
                            case 42:
                                var v24 int32 = 2147483647
                                _ = v24
                            }
                        } else {
                            v5 = append(v5, ("hello" + (v3 + v3)))
                            var v25 int = fi17
                            var v26 int32 = 1
                            _ = v25
                            _ = v26
                        }
                        if (v3 <= "foo bar") {
                            v27 := v2[("\n\t" + v3)]
                            _ = v27
                        }
                        _ = fi23
                    }
                } else {
                    for fi28 := 0; fi28 < 19; fi28++ {
                        if (468 == 32767) {
                            switch v3 {
                            case "\n\t":
                                if (!true) {
                                    if (42 != 0) {
                                        if (!true) {
                                            if true {
                                                v5 = append(v5, "")
                                                var v29 int = fi28
                                                v3 = v3
                                                _ = v29
                                            }
                                            _ = len(v5)
                                            v30 := GenS3{F0: 4294967295, F1: 0}
                                            _ = v30
                                        }
                                        var v31 int8 = (v15 % ((v15) | 1))
                                        fmt.Println(fi28, v3)
                                        _ = v31
                                    }
                                } else {
                                    v4 = (^(v4))
                                }
                            case "a":
                                if false {
                                    v32 := GenS1{F0: (v4 - v4), F1: 829, F2: 2147483647}
                                    v15 -= int8(v1)
                                    _ = v32
                                } else {
                                    var v33 int = fi17
                                    var v34 uint16 = 991
                                    v35 := GenS4{F0: 42, F1: 310}
                                    switch v3 {
                                    case "\n\t":
                                        switch (fi28 & (fi28 / (((-(v33))) | 1))) {
                                        case 0:
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        case 656:
                                            v36 := GenS1{F0: v4, F1: 1, F2: 377}
                                            if (!true) {
                                                var v37 int16 = (v1 - (v1 << ((4294967295) & 63)))
                                                var v38 int64 = genFn0(18446744073709551615, 4294967295)
                                                fmt.Println(v33, fi28, fi17)
                                                _ = v37
                                                _ = v38
                                            }
                                            v39 := GenS5{F0: (-(v1))}
                                            _ = v36
                                            _ = v39
                                        }
                                    case "foo bar":
                                        v3 = v3
                                    case "hello":
                                        fmt.Println(v1)
                                        v40 := GenS4{F0: 127, F1: 9223372036854775807}
                                        _ = v40
                                    }
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                }
                            }
                            var v41 uint64 = v4
                            _ = v41
                        } else {
                            fmt.Println(v3)
                            var v42 uint8 = 1
                            if true {
                                fmt.Println(fi17, v1)
                                var v43 int = -2147483648
                                var v44 int64 = -1
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v42 = 128
                                _ = v43
                                _ = v44
                            }
                            v45 := GenS6{F0: (-(fi17)), F1: (fi17 ^ fi28), F2: v42}
                            _ = v42
                            _ = v45
                        }
                        v3 = "a"
                        _ = fi28
                    }
                }
            }
            var v46 int32 = 1
            _ = v46
        } else {
            v47 := GenS7{F0: "", F1: 1}
            _ = v47
        }
        v5 = append(v5, v3)
        _ = fi17
    }
    var v48 bool = false
    v15 += v15
    var v49 int8 = 37
    if len(v5) > 0 {
        v5[0] = "\n\t"
    }
    v15 = v49
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
}
