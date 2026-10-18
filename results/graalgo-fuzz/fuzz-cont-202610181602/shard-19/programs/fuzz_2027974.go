package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 int64; F2 uint16 }
func genFn0() uint16 {
    v5 := GenS1{F0: 42}
    v6 := []int{}
    var v7 uint8 = 127
    fmt.Println(v7)
    _ = v5
    _ = v6
    _ = v7
    return 735
}

func main() {
    var v1 bool = false
    v2 := []int{2147483647, 42, 874, -2147483648}
    v3 := []int{598, -9223372036854775808}
    var v4 int8 = 127
    switch -2147483648 {
    case 877:
        var v8 bool = true
        _ = v8
    case -352:
        v1 = v1
    default:
        v9 := GenS2{F0: ""}
        _ = v9
    }
    var v10 int32 = 2147483647
    if v1 {
        var v11 int8 = (v4 * (v4 - v4))
        _ = v11
    } else {
        v10 = 455
        v4 = 1
    }
    v12 := GenS1{F0: 1}
    if ("\n\t" >= "") {
        v4 = v4
    }
    v13 := GenS1{F0: 0}
    if false {
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 uint32 = 902
            if (!((v1 || ((3.14159 * 0.0) == 1e10)) || v1)) {
                for fi16 := 0; fi16 < 16; fi16++ {
                    _ = len(v3)
                    _ = fi16
                }
                v10 -= -475
                if len(v3) > 0 {
                    v3[0] = (fi14 / ((1) | 1))
                }
            }
            _ = fi14
            _ = v15
        }
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 float32 = 0.1
        v18 += v18
        fmt.Println(v10, fi17, v1)
        _ = fi17
        _ = v18
    }
    if v1 {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v4 = (v4 | (v4 % ((v4) | 1)))
        switch "a" {
        case "":
            var v19 float32 = 0.1
            _ = v19
        case "foo bar":
            v1 = (!v1)
        case "a":
            fmt.Println(v10, v10, v1)
        }
    } else {
        if (32767 == 42) {
            for fi20 := 0; fi20 < 10; fi20++ {
                if v1 {
                    v3 = append(v3, int(v4))
                } else {
                    for fi21 := 0; fi21 < 14; fi21++ {
                        fmt.Println(v4)
                        switch -9223372036854775808 {
                        case 0:
                            v2 = append(v2, (fi20 >> ((944) & 63)))
                            v2 = append(v2, 0)
                            fmt.Println(fi20)
                            var v22 bool = v1
                            v3 = append(v3, (fi20 % (((-(fi20))) | 1)))
                            _ = v22
                        default:
                            var v23 int8 = (v4 - v4)
                            _ = v23
                        }
                        fmt.Println(v1, fi21, v1)
                        _ = fi21
                    }
                    switch (-(fi20)) {
                    case 15:
                        fmt.Println(v10)
                    case -290:
                        var v24 int64 = -9223372036854775808
                        _ = v24
                    }
                    var v25 uint8 = 42
                    _ = v25
                }
                _ = fi20
            }
        } else {
            if false {
                if v1 {
                    v4 = v4
                    for fi26 := 0; fi26 < 1; fi26++ {
                        if ((4294967295 != 42) || v1) {
                            v2 = append(v2, fi26)
                            for fi27 := 0; fi27 < 5; fi27++ {
                                if (65535 >= 1) {
                                    v28 := GenS3{F0: v4, F1: 1, F2: 2}
                                    v3 = append(v3, fi27)
                                    _ = v28
                                }
                                _ = fi27
                            }
                        } else {
                            v4 -= -1
                        }
                        _ = fi26
                    }
                }
            } else {
                if (!(!v1)) {
                    var v29 int = 127
                    v4 = v4
                    _ = v29
                } else {
                    v1 = (1 >= -1)
                    fmt.Println(v10, v10, v1)
                }
            }
            var v30 int64 = 42
            for fi31 := 0; fi31 < 20; fi31++ {
                var v32 uint8 = 34
                _ = fi31
                _ = v32
            }
            _ = v30
        }
    }
    fmt.Println(v4, v1, v1)
    v10 -= v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
