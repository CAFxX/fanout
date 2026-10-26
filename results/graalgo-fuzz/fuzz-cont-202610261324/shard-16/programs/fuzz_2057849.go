package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int64 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32; F1 uint32; F2 uint64 }
func genFn0() uint32 {
    var v7 uint16 = 902
    var v8 int16 = -1
    var v9 float64 = 1.5
    fmt.Println(v7, v7, v9)
    v10 := GenS1{F0: v7, F1: -124, F2: -1}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 int, p1 bool, p2 int) int {
    v11 := GenS2{F0: 32767, F1: 2147483647}
    p0 = -7
    v12 := GenS1{F0: 65535, F1: 0, F2: -1}
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 32768
}

func main() {
    var v1 uint16 = 42
    v2 := []int{9223372036854775807, 2147483647, 1}
    v3 := []int{128, 4294967295, -400, 128}
    var v4 int = 32768
    var v5 uint32 = 2147483648
    var v6 float64 = float64((1e10 / 3.14159))
    var v13 uint32 = genFn0()
    if true {
        var v14 uint = 0
        var v15 float32 = 0.0
        _ = v14
        _ = v15
    }
    var v16 float64 = ((-2.25 / v6) * (v6 + (v6 - v6)))
    switch "\n\t" {
    case "\n\t":
        v4 = (v4 % ((-9223372036854775808) | 1))
    }
    v17 := GenS2{F0: -1, F1: 253}
    var v18 uint = 255
    v19 := GenS1{F0: v1, F1: -108, F2: 1}
    for fi20 := 0; fi20 < 6; fi20++ {
        var v21 uint64 = 2
        _ = fi20
        _ = v21
    }
    if len(v2) > 0 {
        v2[0] = (-(v4))
    }
    v1 = 255
    var v22 uint = v18
    if false {
        v23 := GenS1{F0: v1, F1: 68, F2: 132}
        _ = v23
    } else {
        if len(v3) > 0 {
            v3[0] = (v4 >> ((uint(v5)) & 63))
        }
        var v24 uint16 = v1
        switch "hello" {
        case "\n\t":
            switch (v4 << ((v18) & 63)) {
            case 4294967295:
                v2 = append(v2, v4)
            case -1:
                for fi25 := 0; fi25 < 16; fi25++ {
                    switch 964 {
                    case 32768:
                        v2 = append(v2, v4)
                    case 42:
                        var v26 int16 = 369
                        _ = v26
                    case 100:
                        if false {
                            var v27 int = v4
                            v16 = v6
                            var v28 uint32 = (v5 | (v13 % ((2) | 1)))
                            _ = v27
                            _ = v28
                        }
                    }
                    v29 := GenS1{F0: v24, F1: -122, F2: 0}
                    _ = fi25
                    _ = v29
                }
            case 127:
                for fi30 := 0; fi30 < 12; fi30++ {
                    if len(v3) > 0 {
                        v3[0] = fi30
                    }
                    if (((0.0 * 1e10) + -2.25) != -2.25) {
                        v31 := GenS2{F0: -1, F1: -1}
                        if len(v2) > 0 {
                            v2[0] = (^(fi30))
                        }
                        for fi32 := 0; fi32 < 11; fi32++ {
                            if false {
                                var v33 float64 = 1.5
                                v34 := GenS1{F0: v24, F1: -37, F2: 687}
                                _ = v33
                                _ = v34
                            }
                            fmt.Println(v4)
                            var v35 int64 = 42
                            _ = fi32
                            _ = v35
                        }
                        _ = v31
                    } else {
                        switch ((((("foo bar" + ("\n\t" + "foo bar")) + ("hello" + ("\n\t" + "hello"))) + ("a" + (("hello" + (("hello" + ("\n\t" + "hello")) + "foo bar")) + "a"))) + ("a" + (("\n\t" + "hello") + ((("hello" + "hello") + "a") + "a")))) + "a") {
                        case "a":
                            v36 := GenS3{F0: (^(v22)), F1: 4294967295}
                            _ = v36
                        case "\n\t":
                            for fi37 := 0; fi37 < 14; fi37++ {
                                if len(v2) > 0 {
                                    v2[0] = (-(v4))
                                }
                                v38 := GenS4{F0: (1 == -941)}
                                v39 := GenS5{F0: -1, F1: 42, F2: 2}
                                var v40 int16 = 0
                                _ = fi37
                                _ = v38
                                _ = v39
                                _ = v40
                            }
                        case "":
                            var v41 uint64 = 1
                            _ = v41
                        }
                        for fi42 := 0; fi42 < 10; fi42++ {
                            v13 = uint32(v1)
                            _ = fi42
                        }
                        v24 -= (v24 / (((v24 % (((v24 / ((v24) | 1))) | 1))) | 1))
                    }
                    var v43 int8 = -111
                    _ = fi30
                    _ = v43
                }
            }
        }
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
