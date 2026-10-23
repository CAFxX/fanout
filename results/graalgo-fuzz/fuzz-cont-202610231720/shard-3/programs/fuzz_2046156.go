package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 bool }
type GenS2 struct { F0 uint64; F1 int16 }
func genFn0(p0 int32) string {
    var v6 int32 = p0
    var v7 string = ""
    if (true && (!true)) {
        v8 := GenS1{F0: 4294967295, F1: v7, F2: true}
        var v9 int32 = (p0 / ((v6) | 1))
        _ = v8
        _ = v9
    }
    if false {
        v10 := []int{1}
        var v11 float64 = -2.25
        for fi12 := 0; fi12 < 10; fi12++ {
            if (v7 != ("\n\t" + "hello")) {
                var v13 uint = 77
                var v14 int64 = 9223372036854775807
                var v15 int32 = (p0 >> (((v13 >> (((^(v13))) & 63))) & 63))
                _ = v13
                _ = v14
                _ = v15
            } else {
                v16 := GenS1{F0: 42, F1: v7, F2: (42 < 328)}
                _ = v16
            }
            _ = fi12
        }
        _ = v10
        _ = v11
    } else {
        var v17 uint64 = 314
        v18 := []int{}
        _ = v17
        _ = v18
    }
    _ = p0
    _ = v6
    _ = v7
    return ("\n\t" + "")
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    var v3 float32 = 0.0
    v4 := []string{"", ("\n\t" + ("" + "")), ("" + "hello")}
    var v5 uint64 = 0
    _ = len(v2)
    var v19 float32 = v3
    switch 128 {
    case -2147483648:
        var v20 int64 = 893
        _ = v20
    case -421:
        v3 += v19
    }
    var v21 string = ("" + "a")
    fmt.Println(v5)
    v4 = append(v4, v21)
    v21 = "hello"
    var v22 int64 = -9223372036854775808
    for fi23 := 0; fi23 < 5; fi23++ {
        v24 := GenS1{F0: 4294967295, F1: v21, F2: true}
        _ = fi23
        _ = v24
    }
    var v25 int = 42
    v25 = (v25 ^ v25)
    if (47 < 42) {
        if (1 <= int8(v22)) {
            _ = len(v4)
        }
        var v26 string = v21
        _ = v26
    } else {
        for fi27 := 0; fi27 < 1; fi27++ {
            var v28 uint32 = 693
            if false {
                v5 = (v5 * (^(v5)))
                var v29 uint32 = 0
                v25 += (v25 - fi27)
                _ = v29
            } else {
                if ((v5 + (v5 | v5)) == 0) {
                    if (false && (!(v19 < v19))) {
                        _ = len(v4)
                    }
                    for fi30 := 0; fi30 < 16; fi30++ {
                        if (!(0 != 65535)) {
                            var v31 uint16 = 255
                            switch -2147483648 {
                            case 124:
                                v1[(v28 * (v28 % (((v28 + (v28 % (((v28 >> ((18446744073709551615) & 63))) | 1)))) | 1)))] = 490
                            case 32767:
                                fmt.Println(v3)
                            default:
                                if false {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                } else {
                                    v32 := GenS2{F0: v5, F1: 0}
                                    var v33 string = v21
                                    _ = v32
                                    _ = v33
                                }
                            }
                            _ = v31
                        }
                        var v34 int = fi30
                        switch v21 {
                        case "hello":
                            v5 = (v5 & (v5 + v5))
                        }
                        _ = fi30
                        _ = v34
                    }
                } else {
                    v25 = (v25 & v25)
                    for fi35 := 0; fi35 < 16; fi35++ {
                        v21 = (((v21 + "\n\t") + "hello") + "a")
                        _ = fi35
                    }
                }
                var v36 float64 = -2.25
                _ = v36
            }
            var v37 float32 = v19
            _ = fi27
            _ = v28
            _ = v37
        }
        for fi38 := 0; fi38 < 11; fi38++ {
            v39 := GenS2{F0: 547, F1: 1000}
            _ = fi38
            _ = v39
        }
        fmt.Println(v19, v5, v19)
    }
    v3 -= v3
    v25 -= (v25 >> ((4294967295) & 63))
    v3 -= (1.5 - v19)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
