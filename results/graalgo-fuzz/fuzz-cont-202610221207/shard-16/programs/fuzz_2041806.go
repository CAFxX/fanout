package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int; F1 uint8; F2 string }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 bool }

func main() {
    var v1 float64 = ((-2.25 - 1e10) - 1e10)
    v2 := []int{634}
    v3 := []string{("foo bar" + "foo bar"), ("\n\t" + (("" + (("hello" + "\n\t") + "hello")) + ("" + "hello"))), ("hello" + "\n\t")}
    v1 = v1
    if (!false) {
        v4 := GenS1{F0: 0, F1: 361}
        fmt.Println(v1)
        for fi5 := 0; fi5 < 13; fi5++ {
            v6 := GenS2{F0: 4294967295, F1: 940}
            for fi7 := 0; fi7 < 7; fi7++ {
                if ((838 >= 18446744073709551615) || (true && ("a" != "hello"))) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    if false {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        fmt.Println(fi7, fi5, fi7)
                    } else {
                        v1 = v1
                        v1 = 1e10
                    }
                }
                _ = len(v3)
                v8 := GenS2{F0: 1, F1: 4294967295}
                if (!true) {
                    if ((-764 == 2147483647) && false) {
                        v9 := GenS3{F0: 2147483647, F1: 0, F2: "hello"}
                        if true {
                            if (0.1 >= ((1.5 / float32(v1)) * 0.0)) {
                                v2 = append(v2, fi5)
                            }
                        }
                        var v10 string = "\n\t"
                        _ = v10
                        _ = v9
                    }
                    v1 -= v1
                    if (-1 <= 1) {
                        v11 := GenS4{F0: (fi7 & (fi7 + fi5)), F1: 1}
                        switch "a" {
                        case "\n\t":
                            var v12 int32 = 1
                            _ = v12
                        case "a":
                            var v13 float32 = (0.0 - 3.14159)
                            _ = v13
                        case "hello":
                            _ = len(v2)
                        default:
                            var v14 int32 = 2147483647
                            _ = v14
                        }
                        v1 += v1
                        _ = v11
                    }
                } else {
                    v15 := GenS3{F0: fi7, F1: 0, F2: "hello"}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v1 += v1
                    _ = v15
                }
                v1 += (v1 * v1)
                if (0 < 2) {
                    fmt.Println(fi7)
                    fmt.Println(fi7, fi5, v1)
                } else {
                    var v16 uint = 2
                    if true {
                        var v17 int32 = 42
                        var v18 uint = 65535
                        if false {
                            if (-1 <= 0) {
                                v19 := GenS2{F0: 4294967295, F1: 18446744073709551615}
                                fmt.Println(v18, v17)
                                _ = v19
                            }
                            fmt.Println(v17, fi7)
                            _ = len(v2)
                        } else {
                            v17 = v17
                        }
                        _ = v17
                        _ = v18
                    } else {
                        switch 0 {
                        case 37:
                            v16 = (v16 + (^(v16)))
                        default:
                            fmt.Println(v16)
                        }
                        v20 := GenS1{F0: 32767, F1: (v16 << ((v16) & 63))}
                        _ = v20
                    }
                    var v21 float64 = (v1 / 0.1)
                    _ = v16
                    _ = v21
                }
                _ = fi7
                _ = v8
            }
            var v22 float32 = (-2.25 * 0.1)
            _ = fi5
            _ = v22
            _ = v6
        }
        _ = v4
    } else {
        var v23 bool = true
        _ = v23
    }
    fmt.Println(v1)
    var v24 float64 = ((float64(0.0) - v1) - (v1 + 0.0))
    _ = len(v3)
    var v25 uint32 = 1
    var v26 uint = 255
    v24 = 3.14159
    v27 := GenS3{F0: 2147483647, F1: 249, F2: ("" + "\n\t")}
    if (153 <= 345) {
        v2 = append(v2, -1)
        v24 += -2.25
    } else {
        for fi28 := 0; fi28 < 3; fi28++ {
            var v29 uint16 = 751
            var v30 int = fi28
            _ = fi28
            _ = v29
            _ = v30
        }
        if (128 > 42) {
            var v31 float32 = -2.25
            _ = v31
        }
        for fi32 := 0; fi32 < 18; fi32++ {
            var v33 int16 = 32767
            switch "a" {
            case "foo bar":
                v34 := GenS5{F0: (((v25 / ((v25) | 1)) != (v25 % ((v25) | 1))) || (true || (133 != 42)))}
                _ = v34
            default:
                for fi35 := 0; fi35 < 15; fi35++ {
                    var v36 uint16 = 255
                    _ = fi35
                    _ = v36
                }
            }
            _ = fi32
            _ = v33
        }
    }
    var v37 int8 = -128
    v37 += -128
    v24 = v1
    for fi38 := 0; fi38 < 3; fi38++ {
        fmt.Println(v1, v24, v37)
        _ = len(v2)
        if true {
            fmt.Println(v26, v37, fi38)
        }
        _ = fi38
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v37
}
