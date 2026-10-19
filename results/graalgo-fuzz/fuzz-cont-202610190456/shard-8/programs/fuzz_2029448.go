package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 bool; F1 uint8 }
func genFn0(p0 bool, p1 string) bool {
    if p0 {
        var v7 float32 = -2.25
        var v8 bool = (false && (-128 > -128))
        if (348 == 205) {
            if (100 <= -1) {
                for fi9 := 0; fi9 < 18; fi9++ {
                    for fi10 := 0; fi10 < 15; fi10++ {
                        if (0.0 != (3.14159 + (0.1 - 0.1))) {
                            if v8 {
                                v8 = (true && (false || ((0 == 2147483647) || false)))
                            } else {
                                p0 = (0 != 128)
                                var v11 int = fi9
                                v12 := []int{}
                                _ = v11
                                _ = v12
                            }
                            p1 = p1
                            var v13 int32 = -582
                            v14 := GenS1{F0: 975, F1: 827}
                            switch (fi9 - (fi9 / ((fi10) | 1))) {
                            case 127:
                                var v15 int8 = 42
                                _ = v15
                            case 4294967295:
                                var v16 int64 = 0
                                _ = v16
                            default:
                                fmt.Println(v7, p1)
                            }
                            p0 = false
                            _ = v13
                            _ = v14
                        }
                        _ = fi10
                    }
                    v17 := []int{fi9}
                    _ = fi9
                    _ = v17
                }
            } else {
                p1 = (p1 + "a")
                v8 = (0 <= -1)
                v18 := []int{32768, 999, -9223372036854775808, 42}
                _ = v18
            }
            v19 := []int{-2147483648, 343, 2147483647, 661}
            var v20 int32 = -392
            _ = v19
            _ = v20
        }
        _ = v7
        _ = v8
    } else {
        for fi21 := 0; fi21 < 13; fi21++ {
            var v22 uint = 18446744073709551615
            p0 = p0
            v23 := GenS2{F0: 1, F1: 1}
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    if p0 {
        fmt.Println(p1, p0)
    } else {
        v24 := []int{32767, 4294967295, 9223372036854775807, 0}
        _ = v24
    }
    p1 = "hello"
    fmt.Println(p1)
    _ = p0
    _ = p1
    return (0 < 583)
}
func genFn1(p0 int64, p1 bool, p2 int) uint32 {
    v25 := GenS2{F0: 609, F1: p0}
    fmt.Println(p2, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    return 4294967295
}

func main() {
    var v1 int16 = 1000
    var v2 bool = ((v1 + v1) == (v1 + v1))
    var v3 uint64 = 1
    v4 := make(map[uint32]bool)
    var v5 uint8 = 255
    var v6 float32 = (float32(3.14159) - 1.5)
    v4[2147483647] = (-263 == 1)
    var v26 bool = v2
    v5 = 128
    var v27 uint32 = 736
    delete(v4, v27)
    var v28 int8 = 1
    v1 = (v1 << ((0) & 63))
    v29 := GenS2{F0: -2147483648, F1: 351}
    if v26 {
        v30 := GenS3{F0: false, F1: v5}
        _ = v30
    } else {
        switch -2147483648 {
        case 42:
            fmt.Println(v28, v26)
        case 774:
            if (2147483647 <= -2147483648) {
                var v31 int = 705
                if ((v1 / (((^(v1))) | 1)) != -32768) {
                    if v2 {
                        switch ("hello" + "a") {
                        case "hello":
                            v32 := []int{(v31 * 128), (v31 - (v31 - v31)), v31, v31}
                            _ = v32
                        }
                    }
                    fmt.Println(v31, v31)
                    switch "a" {
                    case "a":
                        var v33 string = ("" + "\n\t")
                        _ = v33
                    case "":
                        fmt.Println(v6)
                    case "hello":
                        for fi34 := 0; fi34 < 12; fi34++ {
                            v3 = v3
                            _ = fi34
                        }
                    }
                } else {
                    v5 = (^(v5))
                    v2 = false
                    if false {
                        v28 = (-(v28))
                        for fi35 := 0; fi35 < 11; fi35++ {
                            var v36 int8 = -1
                            for fi37 := 0; fi37 < 1; fi37++ {
                                var v38 int = (v31 % ((fi35) | 1))
                                v39 := GenS3{F0: v2, F1: 128}
                                v27 += genFn1(42, true, fi37)
                                v40 := []int{(fi37 << ((4294967295) & 63)), (fi37 | 387), int(v5)}
                                v2 = (241 > 42)
                                _ = fi37
                                _ = v38
                                _ = v39
                                _ = v40
                            }
                            _ = fi35
                            _ = v36
                        }
                    }
                }
                _ = v31
            }
        default:
            switch "" {
            case "\n\t":
                var v41 bool = v2
                _ = v41
            case "hello":
                v5 = 255
            }
        }
        v1 = (^(v1))
    }
    v4[genFn1(int64(v5), (("hello" + "foo bar") != "hello"), 4294967295)] = false
    if v2 {
        var v42 float32 = v6
        fmt.Println(v6, v42, v28)
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
