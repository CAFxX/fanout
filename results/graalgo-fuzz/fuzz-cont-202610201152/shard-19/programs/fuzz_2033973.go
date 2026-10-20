package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 uint64; F1 uint }
type GenS7 struct { F0 uint; F1 string }
func genFn0(p0 int64, p1 int16, p2 uint64) int {
    var v6 int8 = 1
    v7 := GenS1{F0: 579}
    var v8 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 128
}
func genFn1(p0 uint64) uint {
    var v9 int16 = 0
    var v10 int32 = 2147483647
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 int = fi11
        _ = fi11
        _ = v12
    }
    v13 := GenS2{F0: 42, F1: false}
    var v14 string = ("hello" + (("hello" + (("\n\t" + ("" + ("hello" + ("\n\t" + ("" + "a"))))) + "hello")) + (("foo bar" + "foo bar") + "")))
    _ = p0
    _ = v10
    _ = v13
    _ = v14
    _ = v9
    return 136
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 623
    var v3 uint8 = 2
    v4 := []uint32{}
    v5 := []string{"a", "", ("hello" + (("foo bar" + ("foo bar" + "foo bar")) + "hello")), "foo bar"}
    var v15 int32 = -2147483648
    var v16 string = (((("foo bar" + "a") + ("hello" + "\n\t")) + ("\n\t" + ("hello" + ""))) + ("" + ("" + "")))
    v17 := GenS3{F0: "foo bar", F1: 65535}
    delete(v1, -647)
    v18 := GenS3{F0: "a", F1: 297}
    v19 := GenS4{F0: (((v16 + "") + v16) + v16), F1: 42}
    for fi20 := 0; fi20 < 13; fi20++ {
        if false {
            if true {
                v16 = v16
            }
            var v21 uint8 = v3
            switch fi20 {
            case 42:
                fmt.Println(v2, v15)
            case 128:
                var v22 uint16 = 939
                _ = v22
            }
            _ = v21
        } else {
            v23 := GenS3{F0: (v16 + ("a" + (("foo bar" + v16) + ""))), F1: 1}
            _ = v23
        }
        if false {
            var v24 uint16 = 42
            if (false || true) {
                if (!false) {
                    v16 = (v16 + v16)
                    delete(v1, (fi20 & (fi20 - (fi20 / (((fi20 / ((fi20) | 1))) | 1)))))
                    v25 := GenS5{F0: (v24 + (v24 / ((v24) | 1))), F1: -1}
                    _ = v25
                } else {
                    var v26 uint8 = 65
                    var v27 uint = 65535
                    for fi28 := 0; fi28 < 15; fi28++ {
                        _ = len(v4)
                        if (!true) {
                            v16 = v16
                            v24 = v24
                            v3 += v3
                            if (fi20 == -9223372036854775808) {
                                v26 = (v26 & (v26 ^ (v3 / ((121) | 1))))
                            }
                        } else {
                            v3 -= v3
                        }
                        if (v24 >= 0) {
                            fmt.Println(v2, fi28)
                        } else {
                            fmt.Println(v3)
                            v27 -= (v27 & v27)
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        var v29 uint32 = 0
                        switch v16 {
                        case "hello":
                            var v30 bool = ((v24 / (((v24 & v24)) | 1)) != (^(v24)))
                            _ = v30
                        default:
                            if false {
                                var v31 uint8 = (v3 ^ (v26 | (v3 + v26)))
                                switch v16 {
                                case "foo bar":
                                    fmt.Println(v16, v26)
                                }
                                v32 := GenS6{F0: 42, F1: 65535}
                                var v33 uint8 = (v3 / ((255) | 1))
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                        }
                        var v34 int16 = 1
                        _ = fi28
                        _ = v29
                        _ = v34
                    }
                    _ = v26
                    _ = v27
                }
                if true {
                    for fi35 := 0; fi35 < 8; fi35++ {
                        v36 := GenS7{F0: 1, F1: ""}
                        v37 := GenS6{F0: 2, F1: 829}
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                }
                v16 = ((("a" + ((v16 + (v16 + "\n\t")) + v16)) + v16) + (v16 + v16))
            }
            v15 = -2147483648
            _ = v24
        } else {
            v38 := GenS6{F0: 731, F1: 4294967295}
            var v39 uint32 = v2
            v3 = (^(v3))
            _ = v38
            _ = v39
        }
        if len(v5) > 0 {
            v5[0] = (v16 + "a")
        }
        _ = fi20
    }
    v16 = (v16 + v16)
    v40 := v1[42]
    for fi41 := 0; fi41 < 6; fi41++ {
        var v42 bool = (v3 > v3)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v2 -= v40
        _ = fi41
        _ = v42
    }
    v5 = append(v5, v16)
    v15 = v15
    var v43 int8 = -31
    fmt.Println(v40, v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v43
    _ = v5
}
