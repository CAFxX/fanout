package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32; F1 bool }
func genFn0(p0 int, p1 string, p2 bool) int {
    if (574 <= 671) {
        p2 = (572 >= 919)
    } else {
        var v7 uint64 = 1
        _ = v7
    }
    var v8 int32 = 0
    v9 := []int{p0, (p0 + -9223372036854775808)}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 590
}
func genFn1(p0 uint16, p1 int32) string {
    v10 := []int{-705, 42, 726}
    v11 := GenS1{F0: 0, F1: false, F2: -2147483648}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return "a"
}

func main() {
    var v1 uint64 = 111
    var v2 float32 = 1e10
    v3 := make(map[uint32]string)
    v4 := []string{}
    var v5 string = "a"
    v6 := []int{1, 42, 0}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v3[2147483647] = "\n\t"
    v12 := GenS2{F0: 1}
    if len(v4) > 0 {
        v4[0] = v5
    }
    delete(v3, 103)
    var v13 int64 = int64(v1)
    if (false && false) {
        for fi14 := 0; fi14 < 12; fi14++ {
            for fi15 := 0; fi15 < 12; fi15++ {
                if false {
                    delete(v3, 892)
                    switch v5 {
                    case "\n\t":
                        if true {
                            v1 -= (v1 - v1)
                            if false {
                                if len(v4) > 0 {
                                    v4[0] = ("hello" + "")
                                }
                            }
                        } else {
                            var v16 bool = (110 != 75)
                            if true {
                                fmt.Println(v1, v16, v1)
                                var v17 int = (fi14 - (fi15 / (((fi14 * 767)) | 1)))
                                delete(v3, 0)
                                v16 = true
                                v5 = v5
                                _ = v17
                            }
                            _ = v16
                        }
                    case "foo bar":
                        if false {
                            v13 = -1
                        } else {
                            v2 = -2.25
                            var v18 float32 = (1.5 / -2.25)
                            v13 = -9223372036854775808
                            _ = v18
                        }
                    case "hello":
                        var v19 int = (fi15 | 42)
                        v20 := GenS3{F0: 0}
                        _ = v19
                        _ = v20
                    default:
                        fmt.Println(v2, v13, fi15)
                        v3[2147483648] = v5
                        var v21 float32 = (v2 * v2)
                        _ = v21
                    }
                } else {
                    v5 = ""
                }
                v2 = v2
                _ = fi15
            }
            v6 = append(v6, fi14)
            v6 = append(v6, fi14)
            _ = fi14
        }
    } else {
        var v22 string = (v5 + "")
        if (2147483647 < 127) {
            for fi23 := 0; fi23 < 6; fi23++ {
                var v24 int = fi23
                var v25 bool = (false || (900 != 42))
                _ = fi23
                _ = v24
                _ = v25
            }
        } else {
            v26 := GenS2{F0: 2147483647}
            v13 = v13
            _ = v26
        }
        if (-32 == 42) {
            v2 -= 0.0
            if (!((1e10 / 0.1) == (1.5 * (((0.1 / (((float64(((0.1 / (-2.25 / (v2 + v2))) - (float32(3.14159) * v2))) / ((-2.25 + (1e10 - 0.0)) / ((3.14159 - 1e10) / ((0.0 + 0.0) * float64(v2))))) / 1.5) + 3.14159)) - -2.25) * ((1e10 * 1.5) + 0.0))))) {
                if (185 == -2147483648) {
                    _ = len(v3)
                    var v27 int8 = 100
                    if len(v4) > 0 {
                        v4[0] = (v5 + "a")
                    }
                    _ = v27
                } else {
                    for fi28 := 0; fi28 < 3; fi28++ {
                        var v29 float64 = 3.14159
                        v30 := GenS3{F0: 245}
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    v31 := GenS4{F0: 42}
                    v13 += v13
                    _ = v31
                }
            }
        }
        _ = v22
    }
    v4 = append(v4, ("a" + v5))
    var v32 int64 = v13
    if true {
        if true {
            for fi33 := 0; fi33 < 16; fi33++ {
                var v34 float32 = 0.1
                var v35 int8 = 41
                for fi36 := 0; fi36 < 2; fi36++ {
                    v32 -= v13
                    fmt.Println(v1, v1)
                    _ = fi36
                }
                _ = fi33
                _ = v34
                _ = v35
            }
            fmt.Println(v5)
        } else {
            v37 := GenS3{F0: 1}
            _ = v37
        }
        if true {
            v2 -= float32(3.14159)
        } else {
            var v38 uint = 65535
            _ = v38
        }
        v39 := GenS3{F0: 65535}
        _ = v39
    }
    switch ((v5 + (v5 + "\n\t")) + ("" + ("\n\t" + "foo bar"))) {
    case "hello":
        for fi40 := 0; fi40 < 9; fi40++ {
            delete(v3, 1)
            _ = fi40
        }
    case "foo bar":
        v41 := GenS5{F0: 4294967295, F1: true}
        _ = v41
    default:
        v13 = (v13 * (v13 + -1))
    }
    var v42 float32 = 0.1
    v5 = ("hello" + (v5 + (v5 + "")))
    v13 = (v32 % (((v32 | v32)) | 1))
    var v43 int8 = 127
    var v44 bool = true
    if (v44 || (v44 || false)) {
        v43 += -128
        var v45 int = genFn0(0, v5, (393 <= 155))
        _ = v45
    } else {
        if len(v6) > 0 {
            v6[0] = 32768
        }
        v5 = "a"
        v46 := GenS4{F0: 2147483648}
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
