package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint; F1 uint32; F2 uint }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 string) int8 {
    v6 := []int{128, -2147483648, 32768}
    p0 = (p0 + ("a" + (p0 + p0)))
    _ = p0
    _ = v6
    return 127
}
func genFn1() int32 {
    switch 32767 {
    case 127:
        var v7 uint64 = 2
        _ = v7
    case -9223372036854775808:
        if true {
            var v8 uint64 = 775
            _ = v8
        }
    case 32768:
        var v9 uint64 = 2
        _ = v9
    default:
        v10 := GenS1{F0: "hello", F1: 32767}
        _ = v10
    }
    v11 := GenS1{F0: "\n\t", F1: 0}
    v12 := GenS1{F0: "hello", F1: -382}
    var v13 string = ""
    v14 := []int{-403}
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return 0
}

func main() {
    v1 := []uint32{2147483648, 2147483647, 471, 2147483648}
    v2 := make(map[string]string)
    v3 := make(map[string]int)
    v4 := make(map[int]uint32)
    v5 := []int{498, 32767, 322}
    v15 := v2["hello"]
    for fi16 := 0; fi16 < 4; fi16++ {
        for fi17 := 0; fi17 < 19; fi17++ {
            v18 := GenS1{F0: ("a" + v15), F1: 0}
            v15 = v15
            v15 = v15
            var v19 int = fi17
            _ = fi17
            _ = v18
            _ = v19
        }
        switch v15 {
        case "":
            v15 = v15
        default:
            switch v15 {
            case "hello":
                var v20 string = ""
                _ = v20
            default:
                v21 := v3["a"]
                _ = v21
            }
        }
        _ = fi16
    }
    var v22 int16 = 42
    var v23 int16 = v22
    if ((true && false) || (0 == 2)) {
        var v24 int64 = -9223372036854775808
        _ = len(v1)
        _ = v24
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        switch v15 {
        case "hello":
            v25 := GenS2{F0: 1, F1: 0, F2: 2}
            _ = v25
        case "\n\t":
            for fi26 := 0; fi26 < 13; fi26++ {
                for fi27 := 0; fi27 < 16; fi27++ {
                    fmt.Println(fi26, v23)
                    var v28 uint16 = 32767
                    _ = fi27
                    _ = v28
                }
                v29 := GenS1{F0: ((((("foo bar" + "\n\t") + (v15 + v15)) + ("hello" + (v15 + (v15 + "a")))) + (("" + v15) + "hello")) + (v15 + "a")), F1: (v23 / ((-1) | 1))}
                if false {
                    for fi30 := 0; fi30 < 11; fi30++ {
                        fmt.Println(v22)
                        _ = fi30
                    }
                    v22 -= (v23 << ((346) & 63))
                } else {
                    v31 := GenS2{F0: 1, F1: 2147483648, F2: 42}
                    if ((448 >= 4294967295) || false) {
                        if ((false && true) || (fi26 > fi26)) {
                            if (255 < 0) {
                                if (42 != -128) {
                                    switch fi26 {
                                    case -579:
                                        v15 = ""
                                    case -329:
                                        fmt.Println(v23, v22, v15)
                                    default:
                                        var v32 uint64 = 42
                                        _ = v32
                                    }
                                } else {
                                    if true {
                                        v33 := GenS1{F0: v15, F1: v23}
                                        if len(v5) > 0 {
                                            v5[0] = fi26
                                        }
                                        _ = v33
                                    }
                                    v15 = ("" + ((v15 + ((v15 + v15) + (("hello" + v15) + v15))) + v15))
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                fmt.Println(fi26, v22, v23)
                            } else {
                                for fi34 := 0; fi34 < 3; fi34++ {
                                    if true {
                                        if true {
                                            var v35 int64 = 42
                                            switch (fi26 * fi34) {
                                            case 42:
                                                v36 := GenS3{F0: genFn0("a"), F1: 2}
                                                _ = v36
                                            }
                                            var v37 uint32 = 1
                                            _ = v35
                                            _ = v37
                                        } else {
                                            if (("hello" + ("foo bar" + v15)) > "") {
                                                delete(v3, v15)
                                                v22 += (v22 << ((1) & 63))
                                                v38 := GenS2{F0: 2, F1: 2147483647, F2: 0}
                                                _ = v38
                                            } else {
                                                if false {
                                                    v39 := GenS2{F0: 42, F1: 65535, F2: 266}
                                                    var v40 float64 = 0.0
                                                    switch -285 {
                                                    case 1:
                                                        var v41 uint64 = 508
                                                        _ = v41
                                                    }
                                                    _ = v39
                                                    _ = v40
                                                } else {
                                                    v22 += (^(v23))
                                                    if true {
                                                        v1 = append(v1, 1)
                                                    } else {
                                                        if len(v5) > 0 {
                                                            v5[0] = fi34
                                                        }
                                                    }
                                                }
                                                if (32768 == fi26) {
                                                    fmt.Println(fi34)
                                                }
                                                if false {
                                                    v15 = (v15 + "hello")
                                                    v23 -= 0
                                                    var v42 uint = 18446744073709551615
                                                    _ = v42
                                                } else {
                                                    if true {
                                                        v15 = ""
                                                        fmt.Println(fi26, fi34, v15)
                                                    } else {
                                                        _ = len(v5)
                                                        fmt.Println(v15, fi34)
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    }
                                                    switch (fi26 - fi34) {
                                                    case 1:
                                                        v43 := GenS1{F0: "\n\t", F1: (v22 / (((v23 | -32768)) | 1))}
                                                        _ = v43
                                                    case -2147483648:
                                                        var v44 uint16 = 0
                                                        _ = v44
                                                    case 269:
                                                        v2[v15] = (v15 + "foo bar")
                                                    }
                                                    var v45 float64 = 1.5
                                                    v46 := GenS2{F0: 255, F1: 756, F2: 361}
                                                    _ = v45
                                                    _ = v46
                                                }
                                                v15 = ""
                                                var v47 uint16 = 32767
                                                _ = v47
                                            }
                                        }
                                    } else {
                                        v1 = append(v1, 1)
                                        v48 := GenS4{F0: 255}
                                        fmt.Println(fi34, fi34)
                                        _ = v48
                                    }
                                    _ = fi34
                                }
                                v22 += -32768
                                v15 = "a"
                            }
                        } else {
                            var v49 uint = 1
                            _ = v49
                        }
                    }
                    _ = v31
                }
                _ = fi26
                _ = v29
            }
        case "a":
            delete(v4, -664)
        default:
            switch 0 {
            case 0:
                if (true || true) {
                    for fi50 := 0; fi50 < 10; fi50++ {
                        var v51 int8 = -1
                        _ = fi50
                        _ = v51
                    }
                    for fi52 := 0; fi52 < 17; fi52++ {
                        v53 := GenS3{F0: -1, F1: 65535}
                        _ = fi52
                        _ = v53
                    }
                    fmt.Println(v15, v22, v23)
                }
            case 32767:
                _ = len(v5)
            }
        }
    }
    for fi54 := 0; fi54 < 19; fi54++ {
        if (42 > -399) {
            var v55 bool = true
            var v56 int64 = 1
            _ = v55
            _ = v56
        }
        if (!(false || (false && (!(false || true))))) {
            var v57 uint64 = 758
            for fi58 := 0; fi58 < 9; fi58++ {
                var v59 float64 = 3.14159
                _ = fi58
                _ = v59
            }
            var v60 int32 = -1
            _ = v57
            _ = v60
        }
        v61 := GenS1{F0: (v15 + "hello"), F1: (v22 >> ((65535) & 63))}
        _ = fi54
        _ = v61
    }
    v4[-2147483648] = 320
    v23 += v22
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
