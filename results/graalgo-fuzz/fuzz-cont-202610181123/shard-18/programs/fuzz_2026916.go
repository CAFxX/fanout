package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint }
type GenS2 struct { F0 uint8; F1 bool; F2 uint }
type GenS3 struct { F0 uint8; F1 int8; F2 uint32 }
func genFn0(p0 uint8) uint8 {
    if (1 > 293) {
        var v4 int8 = -1
        fmt.Println(v4, p0)
        fmt.Println(p0)
        _ = v4
    } else {
        var v5 string = ((("hello" + "hello") + "a") + "")
        _ = v5
    }
    p0 = p0
    p0 = 2
    p0 -= (p0 * 127)
    _ = p0
    return p0
}
func genFn1(p0 int) bool {
    var v6 int8 = 1
    fmt.Println(p0, v6)
    var v7 int64 = 42
    v8 := []int{p0, (p0 | p0)}
    v8 = append(v8, (p0 * (p0 ^ (p0 & p0))))
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return ((float64((0.1 * (1e10 + -2.25))) / (0.1 * (-2.25 + float64(0.0)))) < 3.14159)
}

func main() {
    var v1 string = "foo bar"
    var v2 uint8 = 42
    var v3 float32 = ((1e10 + 3.14159) - 1.5)
    v2 -= genFn0(v2)
    v9 := GenS1{F0: 79, F1: 42, F2: 65535}
    v10 := []int{0, 32768}
    v11 := GenS1{F0: 42, F1: 2, F2: 65535}
    switch 483 {
    case 32768:
        v12 := GenS1{F0: 1, F1: 2147483648, F2: 42}
        _ = v12
    default:
        for fi13 := 0; fi13 < 4; fi13++ {
            var v14 uint32 = 1
            var v15 int16 = -99
            fmt.Println(v14, v14)
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    if false {
        v16 := GenS1{F0: 4294967295, F1: 42, F2: 73}
        _ = v16
    }
    switch v1 {
    case "":
        v17 := GenS2{F0: (v2 / ((genFn0(42)) | 1)), F1: true, F2: 406}
        _ = v17
    case "hello":
        var v18 uint16 = 32767
        _ = v18
    case "a":
        v19 := GenS1{F0: 0, F1: 1, F2: 255}
        _ = v19
    default:
        v20 := GenS2{F0: 114, F1: (!true), F2: 255}
        _ = v20
    }
    _ = len(v10)
    v21 := GenS1{F0: 42, F1: 65535, F2: 4294967295}
    for fi22 := 0; fi22 < 6; fi22++ {
        switch (fi22 * -585) {
        case 127:
            var v23 bool = true
            _ = v23
        case -355:
            var v24 string = v1
            _ = v24
        }
        _ = fi22
    }
    switch ("\n\t" + (v1 + (v1 + v1))) {
    case "":
        v25 := GenS2{F0: v2, F1: (false || false), F2: 1}
        _ = v25
    case "a":
        var v26 float64 = -2.25
        _ = v26
    }
    var v27 int64 = -1
    fmt.Println(v2, v3, v27)
    v28 := GenS1{F0: 42, F1: 2147483647, F2: 543}
    if (true && (v2 > genFn0(215))) {
        var v29 bool = ((!true) || false)
        var v30 uint = 255
        switch 0 {
        case 127:
            var v31 bool = v29
            _ = v31
        case 366:
            switch 32767 {
            case 2147483647:
                var v32 string = v1
                _ = v32
            case 0:
                var v33 uint8 = genFn0(v2)
                _ = v33
            case 306:
                if v29 {
                    v34 := GenS2{F0: (v2 | genFn0(uint8(v30))), F1: false, F2: v30}
                    _ = v34
                } else {
                    for fi35 := 0; fi35 < 2; fi35++ {
                        v36 := GenS1{F0: 4294967295, F1: 2147483647, F2: (v30 << ((v30) & 63))}
                        _ = fi35
                        _ = v36
                    }
                    fmt.Println(v2, v2)
                }
            }
        }
        _ = v29
        _ = v30
    } else {
        for fi37 := 0; fi37 < 20; fi37++ {
            if len(v10) > 0 {
                v10[0] = (fi37 & fi37)
            }
            for fi38 := 0; fi38 < 7; fi38++ {
                if true {
                    var v39 string = ("a" + ((((("\n\t" + v1) + v1) + "foo bar") + ((v1 + v1) + v1)) + (v1 + v1)))
                    var v40 uint8 = (v2 * 255)
                    if ((v2 << ((999) & 63)) >= genFn0(v2)) {
                        v39 = ((v1 + ((v39 + "hello") + v1)) + (v39 + (v39 + v39)))
                        if (969 != 4294967295) {
                            if false {
                                var v41 float32 = v3
                                _ = v41
                            } else {
                                v42 := GenS3{F0: (^(v40)), F1: 30, F2: 744}
                                v3 = ((v3 + (0.0 - float32((0.0 + float64(1e10))))) * v3)
                                _ = v42
                            }
                        }
                    }
                    _ = v39
                    _ = v40
                } else {
                    var v43 float32 = (v3 + v3)
                    if true {
                        v27 -= 1
                        var v44 bool = ((false && true) || ((-(fi38)) == -1))
                        v27 += (v27 * (v27 - v27))
                        if v44 {
                            var v45 uint64 = 2
                            var v46 uint32 = 42
                            _ = v45
                            _ = v46
                        }
                        var v47 float64 = float64((v3 - 3.14159))
                        v1 = (v1 + v1)
                        v43 += v3
                        _ = v44
                        _ = v47
                    } else {
                        var v48 int64 = 1
                        var v49 int16 = 42
                        _ = v48
                        _ = v49
                    }
                    _ = v43
                }
                _ = len(v10)
                v50 := GenS3{F0: v2, F1: -128, F2: 4294967295}
                _ = fi38
                _ = v50
            }
            var v51 uint8 = (v2 * (v2 >> ((2) & 63)))
            _ = fi37
            _ = v51
        }
        v3 += v3
        _ = len(v10)
    }
    _ = len(v10)
    if (2147483647 > 0) {
        v1 = "a"
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v3
    _ = v9
}
