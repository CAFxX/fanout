package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint32, p1 uint32) int {
    v5 := GenS1{F0: 83}
    p0 -= p1
    switch ((("foo bar" + (("hello" + (("\n\t" + (("a" + "a") + "")) + ("\n\t" + ("" + "")))) + "a")) + "foo bar") + "") {
    case "hello":
        for fi6 := 0; fi6 < 4; fi6++ {
            var v7 int16 = 1000
            for fi8 := 0; fi8 < 9; fi8++ {
                if false {
                    v9 := GenS1{F0: 42}
                    v7 = (v7 / ((v7) | 1))
                    fmt.Println(p1, fi6)
                    v10 := []int{fi8, fi8}
                    _ = v10
                    _ = v9
                } else {
                    if false {
                        switch ("\n\t" + "a") {
                        case "\n\t":
                            fmt.Println(p1)
                        }
                        v7 -= v7
                        v11 := []int{(fi8 % (((^(fi6))) | 1))}
                        _ = v11
                    }
                    v12 := GenS2{F0: -1, F1: 0}
                    if true {
                        v13 := []int{32767, int(p0), 310}
                        if false {
                            var v14 uint64 = 1
                            switch "" {
                            case "hello":
                                switch (fi8 << ((65535) & 63)) {
                                case 42:
                                    var v15 uint64 = 88
                                    _ = v15
                                }
                            case "foo bar":
                                p1 = (^(p0))
                            case "a":
                                var v16 uint16 = 1
                                _ = v16
                            default:
                                v7 += v7
                            }
                            if ((((!true) && true) || true) && true) {
                                p0 += 1
                            }
                            switch "\n\t" {
                            case "\n\t":
                                p1 = p1
                            case "hello":
                                v17 := GenS1{F0: (v14 - v14)}
                                _ = v17
                            case "a":
                                v14 = 4294967295
                            }
                            _ = v14
                        }
                        p1 += p0
                        _ = v13
                    } else {
                        var v18 uint8 = 75
                        v19 := GenS3{F0: (^(v18)), F1: uint64(p0), F2: 646}
                        p0 = p1
                        _ = v18
                        _ = v19
                    }
                    _ = v12
                }
                _ = fi8
            }
            _ = fi6
            _ = v7
        }
    case "":
        var v20 bool = true
        _ = v20
    case "foo bar":
        fmt.Println(p0)
    default:
        for fi21 := 0; fi21 < 4; fi21++ {
            p0 -= (p1 | (^(p0)))
            v22 := []int{fi21, fi21, 497}
            _ = fi21
            _ = v22
        }
    }
    p0 = (^(p0))
    _ = p0
    _ = p1
    _ = v5
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 2147483647
    v3 := make(map[int]bool)
    var v4 bool = true
    v23 := []int{-255, 1, 9223372036854775807}
    switch genFn0((v2 / (((^(v2))) | 1)), v2) {
    case 127:
        if ((("" + "a") + ("a" + "hello")) >= "a") {
            if false {
                if (v4 || v4) {
                    for fi24 := 0; fi24 < 20; fi24++ {
                        var v25 uint = 1
                        _ = fi24
                        _ = v25
                    }
                    if ("a" != (((("foo bar" + "a") + (("hello" + (("hello" + "a") + "foo bar")) + "")) + "\n\t") + "a")) {
                        v2 += v2
                    }
                }
                var v26 int8 = 42
                _ = v26
            }
            v4 = (50 >= -67)
        } else {
            v4 = ((4294967295 == 194) || (!(1 == 831)))
            v27 := GenS1{F0: uint64(v2)}
            _ = v27
        }
    case 32768:
        v28 := v1[v2]
        _ = v28
    default:
        for fi29 := 0; fi29 < 17; fi29++ {
            for fi30 := 0; fi30 < 11; fi30++ {
                if (!v4) {
                    var v31 uint32 = (v2 >> ((86) & 63))
                    _ = v31
                } else {
                    if (v4 || v4) {
                        var v32 string = (("" + ("foo bar" + "a")) + "foo bar")
                        v33 := GenS4{F0: 375, F1: "hello"}
                        var v34 int64 = -562
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    var v35 int64 = 0
                    _ = v35
                }
                _ = fi30
            }
            var v36 float32 = ((0.0 + 1e10) + float32((0.1 - (-2.25 / (-2.25 * (float64(3.14159) * -2.25))))))
            v37 := GenS4{F0: 279, F1: "a"}
            _ = fi29
            _ = v36
            _ = v37
        }
    }
    v38 := GenS4{F0: 42, F1: ("foo bar" + "")}
    if true {
        v2 += (v2 / (((v2 << ((4294967295) & 63))) | 1))
    } else {
        v39 := v3[120]
        _ = v39
    }
    var v40 int32 = 263
    for fi41 := 0; fi41 < 11; fi41++ {
        v42 := GenS5{F0: 127}
        _ = fi41
        _ = v42
    }
    v40 -= v40
    var v43 float32 = ((float32(1.5) + float32(-2.25)) + 1.5)
    _ = len(v3)
    v44 := GenS3{F0: 255, F1: 42, F2: 2}
    v4 = (!true)
    v45 := v1[v2]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v38
    _ = v4
    _ = v40
    _ = v43
    _ = v44
    _ = v45
}
