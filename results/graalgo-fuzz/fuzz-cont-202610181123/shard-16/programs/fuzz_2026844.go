package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint64; F2 bool }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 int64; F1 uint }
func genFn0(p0 uint64, p1 int8, p2 bool) uint32 {
    if ("" < "a") {
        v5 := GenS1{F0: 1000, F1: 255, F2: 255}
        var v6 int64 = -119
        _ = v5
        _ = v6
    } else {
        var v7 int = 2147483647
        if p2 {
            for fi8 := 0; fi8 < 12; fi8++ {
                p2 = p2
                if p2 {
                    fmt.Println(p1, fi8, p1)
                    var v9 uint16 = 32767
                    _ = v9
                } else {
                    if p2 {
                        fmt.Println(v7)
                    } else {
                        var v10 bool = (!(-2.25 < 1.5))
                        if v10 {
                            var v11 int8 = (p1 << ((845) & 63))
                            var v12 uint = 0
                            for fi13 := 0; fi13 < 4; fi13++ {
                                var v14 uint64 = (^(p0))
                                v15 := []int{}
                                v11 += 100
                                _ = fi13
                                _ = v14
                                _ = v15
                            }
                            _ = v11
                            _ = v12
                        }
                        var v16 int8 = p1
                        _ = v10
                        _ = v16
                    }
                }
                var v17 uint32 = 59
                _ = fi8
                _ = v17
            }
            var v18 bool = (1 >= 42)
            _ = v18
        }
        v7 = v7
        _ = v7
    }
    var v19 uint = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return 0
}
func genFn1(p0 int32, p1 uint8) int32 {
    v20 := []int{-21}
    fmt.Println(p1)
    for fi21 := 0; fi21 < 19; fi21++ {
        fmt.Println(fi21)
        if (2 >= 42) {
            var v22 bool = (0.1 <= 1e10)
            v23 := GenS2{F0: p1, F1: 665, F2: v22}
            var v24 string = ("hello" + ((("foo bar" + "a") + "foo bar") + (("" + ("a" + "\n\t")) + "a")))
            _ = v22
            _ = v23
            _ = v24
        }
        _ = fi21
    }
    var v25 float32 = 3.14159
    _ = len(v20)
    _ = p0
    _ = p1
    _ = v20
    _ = v25
    return (p0 / ((p0) | 1))
}

func main() {
    var v1 int64 = 949
    v2 := make(map[uint32]string)
    var v3 int64 = 42
    v4 := make(map[string]bool)
    var v26 uint8 = 127
    var v27 uint64 = 332
    v26 += (v26 << ((42) & 63))
    v27 += v27
    var v28 string = ("a" + "\n\t")
    for fi29 := 0; fi29 < 9; fi29++ {
        for fi30 := 0; fi30 < 18; fi30++ {
            fmt.Println(fi29)
            var v31 float64 = ((3.14159 - ((float64(-2.25) + (3.14159 / 1.5)) * ((0.0 - 0.0) / 1e10))) + 0.0)
            fmt.Println(v31, fi30, v28)
            _ = fi30
            _ = v31
        }
        v28 = ""
        v1 = v1
        _ = fi29
    }
    switch v28 {
    case "foo bar":
        for fi32 := 0; fi32 < 16; fi32++ {
            v28 = v28
            if false {
                var v33 string = v28
                _ = v33
            } else {
                switch v28 {
                case "\n\t":
                    var v34 uint64 = 4294967295
                    _ = v34
                case "":
                    var v35 float32 = 0.1
                    _ = v35
                }
                _ = len(v4)
            }
            var v36 float32 = 0.1
            _ = fi32
            _ = v36
        }
    case "\n\t":
        delete(v2, 1)
    }
    v37 := GenS3{F0: -2147483648, F1: 128}
    v38 := []int{32767}
    v39 := GenS2{F0: (v26 % ((v26) | 1)), F1: (v27 | 360), F2: (((!true) && true) || (((float64((1.5 - ((((0.1 - (1.5 * 1.5)) * float32(1e10)) * 0.0) * (0.0 + (1.5 + float32((0.1 + 0.1))))))) * 1e10) * (0.0 - 3.14159)) >= -2.25))}
    if (true && ((v26 >> ((65535) & 63)) <= 1)) {
        v28 = (v28 + (v28 + v28))
    } else {
        v2[1] = v28
        v3 -= (^(v3))
    }
    switch (((("" + "\n\t") + v28) + "foo bar") + ((((v28 + v28) + v28) + ("\n\t" + v28)) + v28)) {
    case "foo bar":
        v40 := GenS4{F0: "hello", F1: v3}
        _ = v40
    case "":
        if true {
            var v41 uint = 65535
            fmt.Println(v3, v1)
            if true {
                v28 = "a"
            }
            _ = v41
        } else {
            fmt.Println(v28, v28)
            for fi42 := 0; fi42 < 8; fi42++ {
                if (true || true) {
                    v1 += -1
                    v1 = (v3 * (v3 & v1))
                }
                var v43 int32 = 0
                switch v28 {
                case "\n\t":
                    v44 := GenS5{F0: (v3 & 539), F1: 4294967295}
                    _ = v44
                case "a":
                    v4[""] = false
                default:
                    v1 += 0
                }
                _ = fi42
                _ = v43
            }
        }
    case "\n\t":
        v45 := v2[42]
        _ = v45
    default:
        fmt.Println(v28, v27, v3)
    }
    var v46 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v46
}
