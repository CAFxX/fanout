package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 uint8; F2 uint16 }
func genFn0() uint8 {
    v4 := GenS1{F0: 100, F1: 127}
    var v5 string = ""
    v5 = v5
    if true {
        var v6 int8 = 42
        v7 := GenS1{F0: (v6 + v6), F1: 1}
        _ = v6
        _ = v7
    }
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 uint8) uint {
    fmt.Println(p0)
    v8 := GenS2{F0: -9223372036854775808, F1: "\n\t"}
    p0 = p0
    p0 -= p0
    _ = p0
    _ = v8
    return 0
}

func main() {
    v1 := []int{42}
    var v2 float32 = 1.5
    var v3 bool = (4294967295 < 255)
    var v9 uint64 = 824
    for fi10 := 0; fi10 < 13; fi10++ {
        for fi11 := 0; fi11 < 4; fi11++ {
            switch 2147483647 {
            case -2147483648:
                var v12 uint = 0
                var v13 int64 = 1
                var v14 int8 = 0
                _ = v12
                _ = v13
                _ = v14
            }
            _ = fi11
        }
        for fi15 := 0; fi15 < 15; fi15++ {
            fmt.Println(v9, fi10)
            v9 -= (^(v9))
            _ = fi15
        }
        _ = fi10
    }
    fmt.Println(v3, v2, v3)
    fmt.Println(v9, v3, v3)
    if (v3 || ((v2 > (v2 + v2)) || ((false || v3) || false))) {
        v2 = (((float32((3.14159 * 1e10)) * (v2 / 1.0)) * v2) / v2)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v16 float32 = 3.14159
    switch ("" + ("foo bar" + (("\n\t" + ("hello" + "")) + "foo bar"))) {
    case "hello":
        var v17 int8 = 0
        _ = v17
    case "a":
        var v18 float32 = (v2 * -2.25)
        _ = v18
    case "":
        v2 -= ((1e10 - ((0.0 / ((float32((0.0 - 1.5)) + (float32((((float64((1.5 + v16)) * (1e10 + (1e10 / 1.5))) / 1e10) + 1e10)) + 0.1)) / 1.5)) + v2)) + v16)
    default:
        for fi19 := 0; fi19 < 7; fi19++ {
            switch fi19 {
            case 42:
                for fi20 := 0; fi20 < 13; fi20++ {
                    switch "" {
                    case "a":
                        if true {
                            switch 0 {
                            case 32768:
                                var v21 bool = v3
                                _ = v21
                            case 32767:
                                v1 = append(v1, fi20)
                            case 2147483647:
                                if len(v1) > 0 {
                                    v1[0] = fi20
                                }
                            }
                            if len(v1) > 0 {
                                v1[0] = (fi19 | (fi19 << ((1) & 63)))
                            }
                            var v22 int64 = 42
                            _ = v22
                        }
                    case "foo bar":
                        v2 = (v2 * (v16 - v16))
                        if (v3 || false) {
                            var v23 string = (((("a" + "a") + "foo bar") + "a") + "\n\t")
                            if true {
                                fmt.Println(fi19, v3)
                                if v3 {
                                    _ = len(v1)
                                    v2 -= 0.0
                                    v23 = v23
                                } else {
                                    if false {
                                        var v24 uint64 = (v9 % ((2) | 1))
                                        if false {
                                            v23 = ("\n\t" + ("" + v23))
                                        } else {
                                            v24 = (v9 / ((v9) | 1))
                                            v1 = append(v1, fi19)
                                        }
                                        _ = v24
                                    } else {
                                        var v25 uint16 = 2
                                        var v26 int8 = 106
                                        v27 := GenS3{F0: v9}
                                        _ = v25
                                        _ = v26
                                        _ = v27
                                    }
                                }
                            } else {
                                v23 = (v23 + ("a" + (("a" + ((v23 + ("" + ("" + v23))) + v23)) + (v23 + v23))))
                            }
                            var v28 int32 = 42
                            _ = v23
                            _ = v28
                        } else {
                            if true {
                                v3 = (!v3)
                            }
                            switch "" {
                            case "\n\t":
                                if v3 {
                                    if (true || v3) {
                                        var v29 uint32 = 2
                                        var v30 uint = 670
                                        fmt.Println(v29, v2)
                                        v16 -= 3.14159
                                        _ = v29
                                        _ = v30
                                    }
                                } else {
                                    fmt.Println(v16, v3, fi19)
                                }
                            case "a":
                                v31 := GenS1{F0: 64, F1: -68}
                                _ = v31
                            }
                        }
                        fmt.Println(v9, v2, v16)
                    default:
                        fmt.Println(v9, v2, fi19)
                        v3 = (!v3)
                        v32 := GenS1{F0: 100, F1: 127}
                        _ = v32
                    }
                    if v3 {
                        if v3 {
                            var v33 uint = 2
                            _ = v33
                        }
                    } else {
                        v34 := GenS1{F0: 100, F1: 18}
                        v35 := GenS3{F0: v9}
                        _ = v34
                        _ = v35
                    }
                    _ = fi20
                }
            }
            fmt.Println(v16, v3, fi19)
            v36 := GenS4{F0: 42, F1: 128, F2: 0}
            _ = fi19
            _ = v36
        }
    }
    fmt.Println(v9)
    fmt.Println(v2)
    v1 = append(v1, -455)
    for fi37 := 0; fi37 < 5; fi37++ {
        if len(v1) > 0 {
            v1[0] = (fi37 | (fi37 & (-(fi37))))
        }
        _ = len(v1)
        _ = fi37
    }
    if false {
        var v38 uint8 = 42
        _ = v38
    }
    if len(v1) > 0 {
        v1[0] = 32767
    }
    v2 += (((-2.25 / 0.1) * 0.0) - v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v9
}
