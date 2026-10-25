package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 int32; F1 int }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint8; F1 int; F2 int64 }
func genFn0(p0 uint8, p1 int) uint64 {
    fmt.Println(p0)
    if false {
        var v4 float64 = 3.14159
        _ = v4
    }
    p0 += 2
    if false {
        fmt.Println(p0, p0)
    }
    v5 := GenS1{F0: 2, F1: 2}
    _ = p0
    _ = p1
    _ = v5
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]int)
    var v2 int = -577
    var v3 uint32 = 2147483647
    v6 := GenS1{F0: uint16(v3), F1: 42}
    var v7 int32 = 2147483647
    v2 += v2
    v8 := GenS2{F0: 65535, F1: 0}
    if (!true) {
        v3 -= v3
        fmt.Println(v7)
    }
    v2 = v2
    if (1000 != 42) {
        switch (v2 << ((2) & 63)) {
        case 42:
            var v9 float32 = (((((((0.1 * (((-2.25 / ((1.5 / 1e10) + (0.1 / 1.0))) / 1.0) - 1e10)) - 3.14159) / 1.5) * (0.0 * -2.25)) + 0.1) / 1e10) - (float32((((-2.25 * 0.0) - -2.25) + 0.0)) + 1e10))
            _ = v9
        case 4294967295:
            for fi10 := 0; fi10 < 20; fi10++ {
                v11 := []int{v2, v2, (v2 << ((2) & 63))}
                for fi12 := 0; fi12 < 5; fi12++ {
                    switch fi10 {
                    case -2147483648:
                        var v13 int = (^(v2))
                        v2 += v13
                        _ = v13
                    case 2147483647:
                        fmt.Println(v7, v7)
                        delete(v1, "foo bar")
                        switch "" {
                        case "hello":
                            var v14 float64 = -2.25
                            _ = v14
                        case "foo bar":
                            v2 = v2
                        default:
                            switch "hello" {
                            case "\n\t":
                                fmt.Println(v3, fi10, fi10)
                            case "hello":
                                if (!true) {
                                    var v15 float32 = (float32((float64((1e10 * ((3.14159 * ((float32(0.1) * 0.0) / (1e10 + 3.14159))) - 0.0))) - 1.5)) - -2.25)
                                    _ = v15
                                } else {
                                    v11 = append(v11, (fi10 & fi12))
                                    switch (fi10 & (v2 % ((v2) | 1))) {
                                    case 42:
                                        v16 := GenS2{F0: 351, F1: -128}
                                        _ = v16
                                    case 2147483647:
                                        var v17 int32 = (v7 | (v7 - v7))
                                        var v18 int64 = -914
                                        var v19 uint8 = 0
                                        _ = v17
                                        _ = v18
                                        _ = v19
                                    }
                                    var v20 int8 = 24
                                    _ = v20
                                }
                                switch (v2 / ((0) | 1)) {
                                case 0:
                                    var v21 uint = 4294967295
                                    v22 := GenS2{F0: 2, F1: -128}
                                    _ = v21
                                    _ = v22
                                case 127:
                                    var v23 uint64 = 208
                                    _ = v23
                                case -970:
                                    v24 := GenS3{F0: (-(v7))}
                                    _ = v24
                                }
                            }
                        }
                    default:
                        if true {
                            v25 := GenS4{F0: genFn0(203, fi10)}
                            _ = v25
                        } else {
                            if true {
                                v3 -= v3
                                v11 = append(v11, fi12)
                            } else {
                                if ((true && false) && (1 > 18446744073709551615)) {
                                    delete(v1, ("" + ""))
                                    v26 := GenS2{F0: 1, F1: 0}
                                    switch "foo bar" {
                                    case "foo bar":
                                        v3 += v3
                                    case "\n\t":
                                        v27 := GenS4{F0: 2}
                                        _ = v27
                                    }
                                    _ = v26
                                }
                                var v28 bool = false
                                _ = v28
                            }
                            v2 += 0
                            switch "foo bar" {
                            case "hello":
                                if len(v11) > 0 {
                                    v11[0] = (fi10 | fi12)
                                }
                            }
                        }
                    }
                    switch (v2 << ((0) & 63)) {
                    case 684:
                        v29 := GenS5{F0: -395, F1: 0}
                        _ = v29
                    default:
                        var v30 uint64 = 561
                        _ = v30
                    }
                    v31 := v1["\n\t"]
                    _ = fi12
                    _ = v31
                }
                if (true && (!false)) {
                    v2 = -1
                    v2 = -1
                } else {
                    for fi32 := 0; fi32 < 4; fi32++ {
                        v33 := GenS4{F0: genFn0(128, (v2 & fi32))}
                        v34 := GenS6{F0: fi32, F1: "a"}
                        v35 := GenS6{F0: (^(v2)), F1: ("hello" + "hello")}
                        fmt.Println(fi32, v2)
                        _ = fi32
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                }
                _ = fi10
                _ = v11
            }
        default:
            if (((1e10 / 1.0) > 0.1) && (32767 >= 42)) {
                v1["hello"] = 4294967295
                var v36 uint = 1
                var v37 uint64 = 4294967295
                _ = v36
                _ = v37
            }
        }
        var v38 int = v2
        _ = v38
    }
    v39 := GenS1{F0: 2, F1: 4294967295}
    v40 := GenS1{F0: 846, F1: 1}
    _ = len(v1)
    switch 2147483647 {
    case -349:
        fmt.Println(v3, v7)
    case 127:
        for fi41 := 0; fi41 < 10; fi41++ {
            v42 := GenS7{F0: v7, F1: (fi41 ^ -2147483648)}
            for fi43 := 0; fi43 < 14; fi43++ {
                if true {
                    if (127 < -95) {
                        v2 = fi41
                    }
                } else {
                    v44 := GenS8{F0: v3}
                    var v45 uint64 = 251
                    _ = v44
                    _ = v45
                }
                v1["hello"] = -319
                v46 := v1[((("" + ("foo bar" + "a")) + (("foo bar" + "hello") + "hello")) + ("a" + "hello"))]
                fmt.Println(v7)
                _ = fi43
                _ = v46
            }
            _ = fi41
            _ = v42
        }
    case 32767:
        v7 += (v7 / ((v7) | 1))
    }
    v47 := GenS9{F0: 131, F1: (-(v2)), F2: 42}
    v2 = (v2 / ((v2) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v47
    _ = v6
    _ = v7
    _ = v8
}
