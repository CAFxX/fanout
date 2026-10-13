package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 int16; F1 string }
func genFn0(p0 uint, p1 bool) uint8 {
    fmt.Println(p0, p1)
    p0 -= p0
    _ = p0
    _ = p1
    return 255
}

func main() {
    v1 := make(map[string]int)
    v2 := []string{"", ("a" + ("hello" + "")), "hello"}
    var v3 int16 = 854
    var v4 int32 = 1
    var v5 int16 = 32767
    v6 := make(map[string]bool)
    var v7 int64 = 9223372036854775807
    v4 = (v4 & (v4 - 541))
    var v8 int16 = v3
    fmt.Println(v5)
    if true {
        var v9 int8 = -1
        _ = v9
    } else {
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 int = (^(fi10))
            var v12 uint8 = 128
            v13 := GenS1{F0: v7, F1: 2, F2: 42}
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        v14 := GenS1{F0: (v7 * (v7 / ((42) | 1))), F1: 0, F2: 42}
        v2 = append(v2, "foo bar")
        _ = v14
    }
    fmt.Println(v8, v3, v3)
    fmt.Println(v5, v8)
    var v15 uint32 = 65535
    _ = len(v1)
    v15 += 380
    if (v15 == v15) {
        var v16 float32 = -2.25
        _ = v16
    } else {
        v4 = 1
        v8 = (v8 / ((0) | 1))
    }
    v17 := GenS2{F0: uint64(v15), F1: 0}
    if false {
        switch 124 {
        case -135:
            v18 := GenS2{F0: 42, F1: 4294967295}
            _ = v18
        case -525:
            v19 := GenS1{F0: (v7 * -1), F1: v15, F2: 4294967295}
            _ = v19
        case 128:
            switch "\n\t" {
            case "hello":
                if ((!false) && true) {
                    for fi20 := 0; fi20 < 3; fi20++ {
                        switch ("a" + "foo bar") {
                        case "foo bar":
                            if len(v2) > 0 {
                                v2[0] = ""
                            }
                        case "":
                            v4 = (-(v4))
                        case "\n\t":
                            if ((159 <= 684) && (0 != 872)) {
                                v8 -= v8
                                v8 = v8
                            }
                        default:
                            v3 = 728
                        }
                        _ = fi20
                    }
                    v21 := GenS2{F0: 2, F1: 65535}
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                    _ = v21
                }
            default:
                var v22 int64 = (v7 - (v7 & v7))
                _ = v22
            }
        default:
            if (844 != 0) {
                v5 -= v8
                for fi23 := 0; fi23 < 16; fi23++ {
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                    v5 += 0
                    _ = fi23
                }
            } else {
                var v24 float32 = (((3.14159 / 1.5) * ((((((-2.25 / float32((-2.25 / 1.5))) + 0.1) / 1e10) / 3.14159) / 1.5) + 0.0)) / -2.25)
                _ = len(v2)
                _ = v24
            }
        }
        if false {
            for fi25 := 0; fi25 < 2; fi25++ {
                v26 := GenS1{F0: v7, F1: 292, F2: 2}
                var v27 float64 = ((0.1 + 1.5) * float64(3.14159))
                for fi28 := 0; fi28 < 15; fi28++ {
                    fmt.Println(v4, v7, v27)
                    _ = fi28
                }
                _ = fi25
                _ = v26
                _ = v27
            }
            for fi29 := 0; fi29 < 16; fi29++ {
                switch (-(fi29)) {
                case 2147483647:
                    v8 = 32767
                case -1:
                    var v30 uint16 = 255
                    _ = v30
                case 460:
                    v3 -= (^(v8))
                default:
                    var v31 int32 = 647
                    _ = v31
                }
                switch -1 {
                case 32768:
                    fmt.Println(v3, fi29, v7)
                case 0:
                    v32 := GenS1{F0: v7, F1: v15, F2: 1}
                    _ = v32
                case 2147483647:
                    var v33 int64 = 42
                    _ = v33
                }
                switch (((("hello" + ("a" + "a")) + "\n\t") + "a") + "hello") {
                case "hello":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    if (0 == -9223372036854775808) {
                        var v34 uint32 = (v15 >> ((655) & 63))
                        _ = v34
                    }
                }
                _ = fi29
            }
            for fi35 := 0; fi35 < 17; fi35++ {
                v15 = 1
                if (!true) {
                    if true {
                        _ = len(v1)
                    } else {
                        v3 -= (^(v5))
                    }
                    v36 := GenS1{F0: 1, F1: v15, F2: 255}
                    _ = v36
                }
                _ = fi35
            }
        } else {
            v8 += v5
            var v37 uint64 = 18446744073709551615
            var v38 float32 = float32(3.14159)
            _ = v37
            _ = v38
        }
    }
    var v39 uint64 = 18446744073709551615
    v40 := GenS3{F0: v7, F1: 117}
    v41 := GenS4{F0: (v8 - 42), F1: "foo bar"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
