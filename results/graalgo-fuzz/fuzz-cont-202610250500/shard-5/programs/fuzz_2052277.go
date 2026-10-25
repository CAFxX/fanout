package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8 }
func genFn0() uint32 {
    var v5 int64 = -160
    v6 := []int{748, -433, 0}
    switch ("hello" + (("hello" + "\n\t") + "hello")) {
    case "a":
        switch "\n\t" {
        case "":
            v6 = append(v6, 4294967295)
        case "\n\t":
            if false {
                v6 = append(v6, -171)
            } else {
                var v7 uint32 = 42
                v7 += v7
                _ = v7
            }
        default:
            v8 := GenS1{F0: 60, F1: -215}
            _ = v8
        }
    case "":
        var v9 uint = 18446744073709551615
        _ = v9
    default:
        v5 = v5
    }
    _ = v5
    _ = v6
    return 2147483647
}
func genFn1(p0 uint, p1 int64) uint16 {
    switch -9223372036854775808 {
    case -763:
        v10 := []int{}
        _ = v10
    case 1:
        if (499 <= 973) {
            v11 := []int{}
            var v12 int8 = -1
            _ = v11
            _ = v12
        }
    }
    v13 := []int{32767, 4294967295, 9223372036854775807}
    v14 := GenS1{F0: 0, F1: int32(p1)}
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 424
}

func main() {
    v1 := []uint32{42, 2}
    v2 := make(map[string]bool)
    v3 := []int{128, -9223372036854775808}
    var v4 int64 = -636
    for fi15 := 0; fi15 < 10; fi15++ {
        if false {
            v16 := GenS1{F0: 127, F1: 2147483647}
            v17 := GenS1{F0: 255, F1: -336}
            if true {
                if true {
                    v4 -= (v4 | v4)
                } else {
                    v4 = -789
                    switch fi15 {
                    case 878:
                        fmt.Println(v4, fi15)
                    default:
                        var v18 uint8 = 171
                        _ = v18
                    }
                    var v19 bool = true
                    _ = v19
                }
                v4 -= 1
            }
            _ = v16
            _ = v17
        } else {
            if true {
                v20 := GenS1{F0: 128, F1: 42}
                var v21 uint32 = 0
                _ = v20
                _ = v21
            }
            v2[(("" + ("" + "\n\t")) + "\n\t")] = true
            var v22 int64 = v4
            _ = v22
        }
        _ = fi15
    }
    v4 -= (v4 % (((-(v4))) | 1))
    var v23 int32 = 0
    var v24 float32 = ((float32(3.14159) + 3.14159) / 1.5)
    var v25 uint8 = 98
    for fi26 := 0; fi26 < 14; fi26++ {
        v27 := GenS1{F0: v25, F1: v23}
        _ = fi26
        _ = v27
    }
    v28 := GenS1{F0: v25, F1: (v23 ^ 2147483647)}
    var v29 bool = true
    v3 = append(v3, 32768)
    v4 = -9223372036854775808
    var v30 float32 = float32((1e10 + 1e10))
    delete(v2, ("foo bar" + "\n\t"))
    if (32767 <= 1) {
        var v31 string = "hello"
        for fi32 := 0; fi32 < 6; fi32++ {
            v33 := GenS2{F0: genFn0(), F1: 65535}
            v34 := GenS2{F0: 652, F1: 42}
            _ = fi32
            _ = v33
            _ = v34
        }
        fmt.Println(v24, v4, v31)
        _ = v31
    } else {
        fmt.Println(v23, v23, v30)
        if false {
            var v35 bool = v29
            _ = v35
        } else {
            var v36 uint64 = 1
            if (true || (!(2 < 4294967295))) {
                var v37 float64 = (3.14159 / 1.0)
                _ = v37
            }
            _ = v36
        }
    }
    v38 := GenS3{F0: ("foo bar" < ("hello" + "\n\t"))}
    if (!true) {
        v39 := GenS4{F0: v25}
        _ = v39
    } else {
        v25 -= 153
        v23 = (^(v23))
        for fi40 := 0; fi40 < 8; fi40++ {
            for fi41 := 0; fi41 < 5; fi41++ {
                switch "\n\t" {
                case "a":
                    var v42 int = (fi41 ^ (fi41 + -922))
                    _ = v42
                case "hello":
                    if (!(!(42 <= 0))) {
                        fmt.Println(v30, v24, v29)
                        _ = len(v3)
                    } else {
                        var v43 uint64 = 205
                        switch fi41 {
                        case 675:
                            if true {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v44 int8 = 100
                                v29 = v29
                                _ = v44
                            } else {
                                var v45 float32 = (v30 / ((0.0 / ((1.5 - 1e10) / float32(3.14159))) + 0.1))
                                var v46 float32 = 0.0
                                _ = len(v2)
                                _ = v45
                                _ = v46
                            }
                            switch ("\n\t" + ("foo bar" + "a")) {
                            case "hello":
                                var v47 int = (fi41 % ((int(v43)) | 1))
                                _ = v47
                            default:
                                if (32767 <= 514) {
                                    v48 := v2["foo bar"]
                                    var v49 bool = v29
                                    _ = v48
                                    _ = v49
                                }
                            }
                        default:
                            v24 -= 0.1
                        }
                        v50 := GenS2{F0: uint32(v25), F1: uint32(v43)}
                        _ = v43
                        _ = v50
                    }
                }
                v24 -= v24
                switch fi40 {
                case 32768:
                    if v29 {
                        v51 := GenS2{F0: 0, F1: 0}
                        if (!v29) {
                            v23 = 0
                            var v52 uint8 = 42
                            _ = v52
                        } else {
                            v53 := GenS2{F0: 655, F1: 2147483647}
                            v25 -= v25
                            v25 = 255
                            _ = v53
                        }
                        _ = v51
                    } else {
                        var v54 int16 = 42
                        _ = v54
                    }
                default:
                    v55 := v2["\n\t"]
                    _ = v55
                }
                _ = fi41
            }
            var v56 string = "hello"
            _ = len(v1)
            _ = fi40
            _ = v56
        }
    }
    var v57 uint64 = 18446744073709551615
    var v58 int = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v38
    _ = v4
    _ = v57
    _ = v58
}
