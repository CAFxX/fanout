package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int32; F2 int64 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int16, p1 uint64) int {
    var v5 uint16 = 0
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v5
    return 537
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -1
    var v3 int16 = 42
    var v4 int64 = -961
    v3 = (v3 & (-(v3)))
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int8 = 20
        v7 = v7
        switch ("foo bar" + "") {
        case "a":
            if (2 == 197) {
                var v8 float32 = 3.14159
                if true {
                    v9 := GenS1{F0: "hello"}
                    _ = v9
                } else {
                    v8 = float32(1e10)
                    if (v3 > (v3 << ((970) & 63))) {
                        for fi10 := 0; fi10 < 19; fi10++ {
                            v7 += 100
                            _ = fi10
                        }
                        var v11 int32 = (^(v2))
                        fmt.Println(v4)
                        _ = v11
                    } else {
                        if true {
                            v12 := GenS1{F0: (("hello" + ("foo bar" + "hello")) + (("a" + ((("a" + "\n\t") + "") + "hello")) + ("a" + "foo bar")))}
                            v8 += v8
                            _ = v12
                        }
                    }
                }
                _ = v8
            } else {
                delete(v1, 0)
                v13 := GenS2{F0: (v2 ^ v2), F1: v2, F2: v4}
                v14 := GenS1{F0: ("hello" + "hello")}
                _ = v13
                _ = v14
            }
        case "foo bar":
            fmt.Println(v4, v3)
        default:
            v15 := GenS1{F0: ("foo bar" + "\n\t")}
            _ = v15
        }
        _ = fi6
        _ = v7
    }
    if true {
        v2 = -788
        var v16 int32 = 300
        v3 += v3
        _ = v16
    }
    v17 := []int{128}
    if (true || (true || false)) {
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        fmt.Println(v2, v3)
        v18 := GenS1{F0: "a"}
        _ = v18
    }
    _ = len(v1)
    v4 -= (v4 ^ v4)
    if false {
        var v19 float32 = -2.25
        fmt.Println(v2, v19, v2)
        _ = v19
    } else {
        if len(v17) > 0 {
            v17[0] = 128
        }
        switch "foo bar" {
        case "\n\t":
            if true {
                for fi20 := 0; fi20 < 7; fi20++ {
                    v2 = v2
                    var v21 uint32 = 65535
                    v3 = v3
                    _ = fi20
                    _ = v21
                }
                for fi22 := 0; fi22 < 16; fi22++ {
                    for fi23 := 0; fi23 < 14; fi23++ {
                        if (!(436 < 65535)) {
                            v2 = v2
                            fmt.Println(v2, fi23)
                        } else {
                            var v24 uint16 = 0
                            v25 := GenS3{F0: v3, F1: 255}
                            switch "hello" {
                            case "foo bar":
                                if false {
                                    v26 := GenS3{F0: (^(v3)), F1: v24}
                                    switch ("" + "foo bar") {
                                    case "a":
                                        var v27 uint64 = 42
                                        _ = v27
                                    case "\n\t":
                                        var v28 bool = false
                                        _ = v28
                                    case "hello":
                                        fmt.Println(v4, v2)
                                    default:
                                        var v29 uint32 = 1
                                        _ = v29
                                    }
                                    var v30 int64 = -1
                                    _ = v26
                                    _ = v30
                                } else {
                                    v31 := v1[4294967295]
                                    fmt.Println(v2, v24, fi22)
                                    switch "" {
                                    case "a":
                                        var v32 int = int(v4)
                                        _ = v32
                                    case "":
                                        v33 := GenS3{F0: int16(fi23), F1: 2}
                                        _ = v33
                                    }
                                    v34 := GenS4{F0: 18446744073709551615}
                                    if (v4 == (v4 << ((1) & 63))) {
                                        var v35 int8 = 63
                                        if false {
                                            var v36 float32 = 3.14159
                                            if (!(!((0 < 1) && (v2 >= (v2 >> ((477) & 63)))))) {
                                                v36 -= (v36 - (v36 + (((((v36 - (3.14159 * (v36 - (float32(float64((3.14159 - v36))) * (v36 - v36))))) - v36) - v36) + v36) / 0.1)))
                                                var v37 int = fi22
                                                _ = v37
                                            }
                                            _ = v36
                                        }
                                        if (fi22 != fi22) {
                                            _ = len(v1)
                                            fmt.Println(fi23)
                                            v38 := v1[v31]
                                            _ = v38
                                        } else {
                                            var v39 bool = (0 >= 1)
                                            fmt.Println(v3, v4)
                                            delete(v1, v31)
                                            _ = v39
                                        }
                                        var v40 string = "\n\t"
                                        _ = v35
                                        _ = v40
                                    } else {
                                        if true {
                                            if false {
                                                v41 := v1[(v31 / (((v31 * (v31 % ((863) | 1)))) | 1))]
                                                _ = v41
                                            } else {
                                                var v42 string = ("foo bar" + ("foo bar" + (("a" + "") + ("hello" + "foo bar"))))
                                                v43 := GenS3{F0: v3, F1: v24}
                                                switch genFn0((v3 / ((v3) | 1)), 150) {
                                                case 2147483647:
                                                    v42 = v42
                                                case 557:
                                                    var v44 bool = (v31 > v31)
                                                    v24 += (v24 ^ v24)
                                                    _ = v44
                                                default:
                                                    fmt.Println(v4, v3)
                                                }
                                                if (1e10 != -2.25) {
                                                    var v45 float32 = 0.1
                                                    var v46 int64 = v4
                                                    _ = v45
                                                    _ = v46
                                                }
                                                _ = v42
                                                _ = v43
                                            }
                                            if len(v17) > 0 {
                                                v17[0] = genFn0((v3 * (v3 | 1000)), 1)
                                            }
                                        } else {
                                            v3 = (v3 | v3)
                                            var v47 string = (("" + ("" + "")) + "foo bar")
                                            _ = v47
                                        }
                                    }
                                    _ = v31
                                    _ = v34
                                }
                            case "\n\t":
                                if true {
                                    fmt.Println(v4, v2)
                                    v24 = (v24 + v24)
                                    fmt.Println(v2, v24)
                                }
                            }
                            _ = v24
                            _ = v25
                        }
                        _ = fi23
                    }
                    v48 := GenS2{F0: 731, F1: 42, F2: v4}
                    switch "\n\t" {
                    case "hello":
                        var v49 bool = false
                        _ = v49
                    case "":
                        var v50 float32 = 1.5
                        _ = v50
                    case "a":
                        var v51 float32 = 1.5
                        _ = v51
                    default:
                        v52 := GenS2{F0: v2, F1: 398, F2: v4}
                        _ = v52
                    }
                    _ = fi22
                    _ = v48
                }
            }
        }
    }
    for fi53 := 0; fi53 < 13; fi53++ {
        v3 = (v3 - (v3 / ((32767) | 1)))
        _ = fi53
    }
    v4 -= (^(v4))
    v4 += v4
    var v54 string = ""
    fmt.Println(v2, v2, v4)
    var v55 int64 = v4
    switch 2147483647 {
    case -169:
        fmt.Println(v3, v2)
    case 265:
        for fi56 := 0; fi56 < 3; fi56++ {
            v57 := GenS3{F0: v3, F1: 0}
            _ = fi56
            _ = v57
        }
    case -1:
        var v58 string = ((v54 + "a") + v54)
        _ = v58
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v55
}
