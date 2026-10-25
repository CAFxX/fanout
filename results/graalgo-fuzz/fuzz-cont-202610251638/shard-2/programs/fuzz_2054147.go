package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint64; F2 int }
type GenS3 struct { F0 uint }
func genFn0() uint16 {
    v5 := []int{}
    switch "\n\t" {
    case "foo bar":
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 uint = 135
            v7 -= (v7 - 2)
            v7 = (v7 - (v7 << ((v7) & 63)))
            _ = fi6
            _ = v7
        }
    case "":
        var v8 int32 = 1
        _ = v8
    case "a":
        switch -2147483648 {
        case -515:
            if (("a" < (("hello" + ("foo bar" + "foo bar")) + ("a" + ("\n\t" + (("hello" + "a") + (("a" + "") + "foo bar")))))) || ((((42 > 65535) || false) && true) && ("a" <= "hello"))) {
                if (-567 != 1000) {
                    var v9 string = "a"
                    _ = v9
                }
                switch "foo bar" {
                case "\n\t":
                    var v10 float64 = 1e10
                    _ = v10
                case "hello":
                    var v11 float64 = 0.0
                    _ = v11
                case "":
                    var v12 uint8 = 127
                    _ = v12
                default:
                    if len(v5) > 0 {
                        v5[0] = -9223372036854775808
                    }
                }
                if false {
                    var v13 uint8 = 1
                    for fi14 := 0; fi14 < 16; fi14++ {
                        if true {
                            var v15 float64 = 1e10
                            fmt.Println(v15)
                            _ = v15
                        }
                        var v16 uint = 1
                        _ = fi14
                        _ = v16
                    }
                    v17 := GenS1{F0: -125}
                    _ = v13
                    _ = v17
                } else {
                    var v18 bool = (-1 != 1)
                    if (29 != -968) {
                        var v19 uint16 = 1
                        var v20 int8 = -128
                        _ = v19
                        _ = v20
                    }
                    v18 = (100 <= 80)
                    _ = v18
                }
            } else {
                var v21 int32 = 920
                _ = v21
            }
        case -9223372036854775808:
            for fi22 := 0; fi22 < 18; fi22++ {
                var v23 int32 = 1
                _ = fi22
                _ = v23
            }
        default:
            if (53 <= 128) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        }
    default:
        for fi24 := 0; fi24 < 13; fi24++ {
            switch ("foo bar" + "\n\t") {
            case "foo bar":
                if len(v5) > 0 {
                    v5[0] = 949
                }
            default:
                var v25 string = ""
                _ = v25
            }
            var v26 float32 = (float32(float64(float32(0.0))) / 1e10)
            _ = fi24
            _ = v26
        }
    }
    v27 := GenS2{F0: 1, F1: 727, F2: 127}
    for fi28 := 0; fi28 < 9; fi28++ {
        var v29 string = ("" + (((("hello" + "foo bar") + "foo bar") + ("a" + ("foo bar" + "foo bar"))) + "hello"))
        v29 = v29
        v29 = v29
        _ = fi28
        _ = v29
    }
    _ = v27
    _ = v5
    return 255
}
func genFn1(p0 int, p1 uint, p2 int8) uint {
    fmt.Println(p1, p2, p1)
    fmt.Println(p2, p1, p0)
    for fi30 := 0; fi30 < 16; fi30++ {
        if (0 >= 430) {
            var v31 uint8 = 42
            _ = v31
        } else {
            p2 = (p2 % ((p2) | 1))
            v32 := []int{555}
            _ = v32
        }
        if ((true && (float64((float32(float64(0.0)) * (0.0 + 0.0))) >= 0.1)) && true) {
            v33 := GenS1{F0: -128}
            var v34 float32 = 1e10
            _ = v33
            _ = v34
        } else {
            var v35 float32 = 3.14159
            v36 := GenS2{F0: fi30, F1: 216, F2: 0}
            switch p0 {
            case -2147483648:
                fmt.Println(p2, v35, p2)
            }
            _ = v35
            _ = v36
        }
        p1 = p1
        _ = fi30
    }
    _ = p0
    _ = p1
    _ = p2
    return (p1 & 18446744073709551615)
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]uint32)
    v3 := []string{"a", ("" + "foo bar")}
    var v4 string = "\n\t"
    var v37 int32 = 42
    var v38 string = v4
    var v39 uint16 = 65535
    if ((-2147483648 == -716) || true) {
        v2[2147483648] = 65535
        if false {
            for fi40 := 0; fi40 < 5; fi40++ {
                var v41 int64 = -9223372036854775808
                v42 := GenS2{F0: fi40, F1: 442, F2: (fi40 - (^(fi40)))}
                _ = fi40
                _ = v41
                _ = v42
            }
            if false {
                var v43 int32 = 2147483647
                fmt.Println(v43, v43, v43)
                _ = v43
            }
            v44 := GenS3{F0: 795}
            _ = v44
        } else {
            v39 -= v39
            var v45 int32 = (^(v37))
            _ = v45
        }
        if false {
            var v46 int32 = -490
            _ = v46
        } else {
            var v47 bool = (genFn0() == v39)
            _ = v47
        }
    }
    for fi48 := 0; fi48 < 11; fi48++ {
        _ = len(v1)
        _ = fi48
    }
    for fi49 := 0; fi49 < 19; fi49++ {
        if false {
            _ = len(v3)
        } else {
            var v50 int = 127
            var v51 bool = (v37 != 42)
            _ = v50
            _ = v51
        }
        _ = len(v3)
        var v52 float64 = 0.1
        _ = fi49
        _ = v52
    }
    if (-251 != -1) {
        fmt.Println(v37, v38)
        fmt.Println(v38, v37)
        v53 := GenS1{F0: -1}
        _ = v53
    }
    fmt.Println(v4, v39)
    if true {
        _ = len(v2)
    }
    for fi54 := 0; fi54 < 8; fi54++ {
        var v55 int32 = (v37 ^ (v37 >> ((genFn1((fi54 & (fi54 ^ 535)), 356, 127)) & 63)))
        _ = fi54
        _ = v55
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
