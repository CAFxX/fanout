package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 int32; F2 string }
func genFn0(p0 int16, p1 uint32) bool {
    if false {
        v5 := []int{42, 580, 127, 42}
        fmt.Println(p0)
        var v6 uint8 = 24
        _ = v5
        _ = v6
    } else {
        if (44 == 255) {
            for fi7 := 0; fi7 < 8; fi7++ {
                var v8 uint8 = 128
                var v9 int16 = p0
                if ((-1 >= 42) && true) {
                    v9 -= (p0 + int16(p1))
                    var v10 uint32 = 2
                    _ = v10
                }
                _ = fi7
                _ = v8
                _ = v9
            }
            var v11 uint64 = 998
            var v12 int32 = 2147483647
            _ = v11
            _ = v12
        }
    }
    if false {
        var v13 int32 = 2147483647
        _ = v13
    }
    p1 += p1
    v14 := GenS1{F0: true, F1: 757}
    v15 := []int{2147483647, 9223372036854775807, 128}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return (!(100 == 42))
}
func genFn1(p0 uint16) uint {
    v16 := GenS1{F0: (true && (0 < 255)), F1: (p0 | 1)}
    v17 := GenS1{F0: true, F1: (^(p0))}
    if ("hello" == (((("hello" + "") + "a") + (("a" + "foo bar") + "a")) + "foo bar")) {
        p0 = (p0 - (^(p0)))
        v18 := GenS1{F0: false, F1: 456}
        switch ("\n\t" + "hello") {
        case "a":
            var v19 float32 = float32(0.1)
            _ = v19
        case "hello":
            p0 = p0
        case "foo bar":
            if (-358 <= -32768) {
                if ((p0 >= 531) && true) {
                    p0 -= (p0 << ((42) & 63))
                    v20 := GenS1{F0: (!false), F1: 1}
                    _ = v20
                }
            }
        }
        _ = v18
    }
    _ = p0
    _ = v16
    _ = v17
    return 18446744073709551615
}

func main() {
    var v1 int16 = -32768
    var v2 int16 = (v1 >> ((777) & 63))
    var v3 int8 = 0
    var v4 int32 = -482
    v21 := GenS1{F0: false, F1: 42}
    var v22 int8 = v3
    var v23 uint8 = 128
    if (false || (-1 > 0)) {
        fmt.Println(v3, v23)
        var v24 float64 = 0.0
        _ = v24
    } else {
        v25 := GenS1{F0: (v4 >= v4), F1: 1}
        fmt.Println(v23, v22, v22)
        _ = v25
    }
    for fi26 := 0; fi26 < 7; fi26++ {
        var v27 uint32 = 0
        if (42 != 0) {
            var v28 bool = true
            fmt.Println(v4, v27)
            _ = v28
        } else {
            for fi29 := 0; fi29 < 8; fi29++ {
                v23 += (v23 << ((genFn1(32767)) & 63))
                v1 -= v1
                _ = fi29
            }
            if true {
                var v30 bool = (!false)
                var v31 int16 = (^(v2))
                _ = v30
                _ = v31
            } else {
                for fi32 := 0; fi32 < 12; fi32++ {
                    v27 -= (^(v27))
                    var v33 uint64 = 18446744073709551615
                    v34 := []int{fi32}
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            }
        }
        _ = fi26
        _ = v27
    }
    fmt.Println(v22)
    if (309 != 42) {
        for fi35 := 0; fi35 < 8; fi35++ {
            switch "hello" {
            case "":
                v36 := GenS1{F0: true, F1: 0}
                _ = v36
            case "a":
                var v37 uint = 2
                _ = v37
            case "foo bar":
                for fi38 := 0; fi38 < 19; fi38++ {
                    if (42 == (v2 | (v1 % (((v1 % ((v2) | 1))) | 1)))) {
                        var v39 int32 = -2147483648
                        if true {
                            v1 -= v2
                            v40 := GenS1{F0: true, F1: 533}
                            _ = v40
                        }
                        v41 := GenS2{F0: (v4 - (-(v39)))}
                        _ = v39
                        _ = v41
                    } else {
                        v42 := GenS1{F0: true, F1: 0}
                        v43 := GenS3{F0: v1, F1: 42, F2: ("hello" + ("hello" + "foo bar"))}
                        var v44 uint32 = 2
                        _ = v42
                        _ = v43
                        _ = v44
                    }
                    if true {
                        v2 = v2
                        var v45 uint64 = uint64(v23)
                        v46 := GenS1{F0: false, F1: 2}
                        _ = v45
                        _ = v46
                    }
                    v47 := []int{-2147483648, 4294967295, 32768, (fi38 + 679)}
                    _ = fi38
                    _ = v47
                }
            default:
                v3 += -128
            }
            _ = fi35
        }
        fmt.Println(v4)
    } else {
        v48 := []int{32767, 32768, 9223372036854775807, 1}
        v2 = (v1 - (v2 >> ((42) & 63)))
        var v49 int8 = v3
        _ = v48
        _ = v49
    }
    v3 = v3
    v50 := []int{42, 2147483647, 128}
    var v51 float32 = (1.5 / (1e10 / (((0.0 - 1.5) / ((0.0 - (3.14159 / 0.1)) + (-2.25 * 1e10))) + 1.5)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v50
    _ = v51
}
