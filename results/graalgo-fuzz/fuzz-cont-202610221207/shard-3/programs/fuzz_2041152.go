package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 int32; F2 int16 }
func genFn0(p0 uint8, p1 uint32, p2 int8) uint32 {
    v6 := GenS1{F0: "foo bar", F1: 42, F2: 4294967295}
    v7 := []int{32768, 32767, 183}
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p1
}

func main() {
    var v1 int16 = 0
    var v2 int8 = 1
    v3 := []string{"foo bar", (("foo bar" + ("\n\t" + "")) + ("foo bar" + "\n\t")), "\n\t", "a"}
    var v4 float32 = 3.14159
    v5 := make(map[string]bool)
    if len(v3) > 0 {
        v3[0] = "a"
    }
    fmt.Println(v1, v1)
    delete(v5, "hello")
    var v8 float32 = 0.1
    if true {
        v4 += ((-2.25 + 0.1) / (float32(0.0) * (float32(0.1) * v8)))
    } else {
        for fi9 := 0; fi9 < 7; fi9++ {
            v10 := GenS2{F0: (389 != 2), F1: 459, F2: 4294967295}
            if false {
                if (!true) {
                    v3 = append(v3, "hello")
                    v1 += (v1 >> ((954) & 63))
                } else {
                    if ("" >= ("foo bar" + "a")) {
                        if (true || (false || true)) {
                            v4 = ((3.14159 * 0.1) - v8)
                            switch "\n\t" {
                            case "foo bar":
                                v1 += v1
                            default:
                                for fi11 := 0; fi11 < 9; fi11++ {
                                    fmt.Println(v8)
                                    _ = fi11
                                }
                            }
                        } else {
                            if true {
                                v1 += (v1 % (((v1 / (((v1 | (v1 * v1))) | 1))) | 1))
                            } else {
                                var v12 bool = (32767 <= 2)
                                var v13 uint16 = 885
                                _ = v12
                                _ = v13
                            }
                        }
                        var v14 float32 = (v8 + (v8 + v8))
                        v15 := GenS2{F0: ((((-2.25 * float32(1e10)) * v8) * v8) > v8), F1: 875, F2: 407}
                        _ = v14
                        _ = v15
                    }
                    v16 := GenS3{F0: -2147483648}
                    var v17 int8 = (v2 % ((v2) | 1))
                    _ = v16
                    _ = v17
                }
                fmt.Println(v8, v2)
                var v18 float32 = ((v8 / (v8 + v4)) + v8)
                _ = v18
            } else {
                var v19 int = fi9
                _ = v19
            }
            _ = fi9
            _ = v10
        }
    }
    if (4294967295 > 4294967295) {
        v4 = (0.0 + v4)
        var v20 int = -174
        var v21 bool = true
        _ = v20
        _ = v21
    } else {
        var v22 int64 = -9223372036854775808
        var v23 int64 = (v22 & (v22 % (((^(v22))) | 1)))
        _ = v22
        _ = v23
    }
    v4 = (float32(((3.14159 - 3.14159) / (((0.1 / 1.5) - -2.25) + -2.25))) / v4)
    var v24 float64 = 1.5
    var v25 bool = true
    var v26 int32 = 358
    for fi27 := 0; fi27 < 7; fi27++ {
        v5[("hello" + (((("hello" + "\n\t") + ("a" + ("hello" + ""))) + "a") + "hello"))] = false
        switch (((("hello" + "hello") + "") + "") + "\n\t") {
        case "foo bar":
            var v28 uint = 1
            _ = v28
        case "hello":
            if ("foo bar" > "hello") {
                fmt.Println(v4, v26)
                if (255 <= 65535) {
                    v24 -= v24
                    v29 := GenS2{F0: v25, F1: 108, F2: 2}
                    _ = v29
                }
            }
        }
        _ = fi27
    }
    v30 := GenS4{F0: 42, F1: -2147483648, F2: 1000}
    if true {
        fmt.Println(v1)
        var v31 uint32 = 2147483647
        v32 := v5[("\n\t" + "\n\t")]
        _ = v31
        _ = v32
    } else {
        var v33 int = 32767
        for fi34 := 0; fi34 < 15; fi34++ {
            var v35 int16 = 602
            _ = len(v5)
            if len(v3) > 0 {
                v3[0] = "a"
            }
            _ = fi34
            _ = v35
        }
        _ = v33
    }
    v36 := GenS4{F0: 4294967295, F1: v26, F2: 0}
    for fi37 := 0; fi37 < 12; fi37++ {
        var v38 float32 = ((-2.25 + v4) + 1e10)
        _ = fi37
        _ = v38
    }
    var v39 int64 = -655
    v2 = v2
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v36
    _ = v39
    _ = v4
    _ = v5
    _ = v8
}
