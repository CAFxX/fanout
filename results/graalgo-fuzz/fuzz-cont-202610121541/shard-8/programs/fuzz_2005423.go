package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS4 struct { F0 int }
func genFn0(p0 int, p1 int32) int64 {
    for fi5 := 0; fi5 < 11; fi5++ {
        p1 += (p1 % ((156) | 1))
        _ = fi5
    }
    v6 := GenS1{F0: 255, F1: (!(26 <= 219))}
    _ = p0
    _ = p1
    _ = v6
    return 42
}
func genFn1() uint32 {
    v7 := []int{42, 42}
    switch "\n\t" {
    case "\n\t":
        v8 := GenS1{F0: 463, F1: true}
        _ = v8
    case "a":
        if (!("foo bar" < "hello")) {
            for fi9 := 0; fi9 < 16; fi9++ {
                var v10 uint16 = 600
                if (0 == 2147483647) {
                    for fi11 := 0; fi11 < 17; fi11++ {
                        v10 -= 704
                        _ = fi11
                    }
                    var v12 bool = true
                    for fi13 := 0; fi13 < 14; fi13++ {
                        var v14 int8 = 87
                        _ = fi13
                        _ = v14
                    }
                    _ = v12
                } else {
                    var v15 int64 = genFn0(fi9, 475)
                    _ = v15
                }
                if (!false) {
                    var v16 int32 = 1
                    if ((!(false && false)) && ("hello" == "hello")) {
                        v17 := GenS1{F0: uint(v10), F1: false}
                        if (-189 >= 1000) {
                            v18 := GenS2{F0: 894, F1: "foo bar"}
                            if (v16 >= 1) {
                                var v19 uint16 = v10
                                _ = v19
                            } else {
                                var v20 uint16 = 65535
                                var v21 int = 128
                                v20 -= 0
                                _ = v20
                                _ = v21
                            }
                            _ = v18
                        } else {
                            var v22 bool = true
                            _ = v22
                        }
                        var v23 int8 = -3
                        _ = v17
                        _ = v23
                    } else {
                        if true {
                            var v24 int64 = 544
                            _ = v24
                        } else {
                            v25 := GenS2{F0: 255, F1: ("" + "hello")}
                            for fi26 := 0; fi26 < 4; fi26++ {
                                v16 -= (v16 + 1)
                                _ = fi26
                            }
                            _ = v25
                        }
                        for fi27 := 0; fi27 < 1; fi27++ {
                            v10 += v10
                            fmt.Println(fi9)
                            _ = fi27
                        }
                        var v28 float32 = 1e10
                        _ = v28
                    }
                    v7 = append(v7, 1)
                    _ = v16
                }
                _ = fi9
                _ = v10
            }
            v29 := GenS3{F0: 18446744073709551615, F1: 171, F2: 0}
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = v29
        }
    case "foo bar":
        var v30 uint64 = 908
        _ = v30
    default:
        if true {
            var v31 int64 = 509
            for fi32 := 0; fi32 < 10; fi32++ {
                var v33 uint8 = 1
                _ = fi32
                _ = v33
            }
            _ = v31
        } else {
            var v34 int8 = 50
            fmt.Println(v34)
            _ = len(v7)
            _ = v34
        }
    }
    v35 := GenS3{F0: 18, F1: 127, F2: 127}
    _ = v35
    _ = v7
    return 65535
}

func main() {
    var v1 int32 = -644
    v2 := make(map[uint32]int)
    v3 := []string{"hello", "foo bar"}
    v4 := []int{32768, 296, -1, 2147483647}
    if (true && (true && false)) {
        var v36 int64 = 42
        _ = v36
    }
    for fi37 := 0; fi37 < 9; fi37++ {
        v1 += v1
        v1 += v1
        _ = fi37
    }
    if false {
        switch "hello" {
        case "\n\t":
            if (!false) {
                for fi38 := 0; fi38 < 19; fi38++ {
                    if (!(127 != 0)) {
                        v1 = 174
                        if (!true) {
                            if len(v3) > 0 {
                                v3[0] = ""
                            }
                        } else {
                            _ = len(v4)
                        }
                    }
                    _ = len(v4)
                    _ = fi38
                }
            } else {
                v39 := GenS4{F0: 0}
                v2[2] = 397
                var v40 float64 = -2.25
                _ = v39
                _ = v40
            }
        default:
            var v41 int = -122
            _ = v41
        }
        for fi42 := 0; fi42 < 9; fi42++ {
            v1 += v1
            v1 = 1
            _ = fi42
        }
    } else {
        delete(v2, 999)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v1 -= 0
    v43 := GenS2{F0: 0, F1: "\n\t"}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "hello" {
    case "a":
        v3 = append(v3, (("foo bar" + "foo bar") + (((("a" + ("" + "a")) + "hello") + ("foo bar" + (("\n\t" + "foo bar") + "foo bar"))) + "\n\t")))
    }
    if ("hello" >= "foo bar") {
        _ = len(v2)
    }
    for fi44 := 0; fi44 < 15; fi44++ {
        v45 := GenS3{F0: 1, F1: 2, F2: 170}
        _ = fi44
        _ = v45
    }
    switch 942 {
    case 0:
        delete(v2, 2)
    case 4294967295:
        if true {
            var v46 int8 = 127
            v1 -= (v1 % ((-2147483648) | 1))
            v46 = v46
            _ = v46
        }
    }
    v1 += v1
    v2[2147483647] = 4294967295
    var v47 float64 = (0.1 / float64((-2.25 + 3.14159)))
    v1 += -2147483648
    var v48 string = (("foo bar" + "hello") + "")
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v47
    _ = v48
}
