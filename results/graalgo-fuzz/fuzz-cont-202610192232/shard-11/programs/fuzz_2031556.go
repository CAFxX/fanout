package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int8 }
type GenS2 struct { F0 int16; F1 int; F2 int32 }
func genFn0() string {
    v4 := []int{-223}
    var v5 int32 = 1
    v5 -= v5
    switch -623 {
    case 0:
        if true {
            fmt.Println(v5)
        }
    }
    _ = v4
    _ = v5
    return ""
}
func genFn1(p0 int8, p1 int64) string {
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := []int{fi6, (fi6 + (^(fi6))), -9223372036854775808, fi6}
        _ = fi6
        _ = v7
    }
    v8 := []int{}
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            p0 = (p0 ^ p0)
            _ = fi10
        }
        _ = fi9
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        switch 42 {
        case 128:
            switch "foo bar" {
            case "a":
                v12 := GenS1{F0: 4294967295, F1: 998, F2: p0}
                _ = v12
            }
        case 32767:
            var v13 uint32 = 2147483648
            _ = v13
        default:
            fmt.Println(p0, p0)
        }
        _ = fi11
    }
    var v14 int16 = -164
    _ = p0
    _ = p1
    _ = v14
    _ = v8
    return (("foo bar" + "") + "hello")
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]int)
    var v3 string = ("foo bar" + "")
    var v15 uint64 = 4294967295
    _ = len(v2)
    if ((0 != 4294967295) && (1 < 1)) {
        v3 = (v3 + "\n\t")
        if (!true) {
            v16 := GenS1{F0: 18446744073709551615, F1: 328, F2: 65}
            for fi17 := 0; fi17 < 18; fi17++ {
                v15 += v15
                v18 := GenS1{F0: v15, F1: 1, F2: 100}
                switch "a" {
                case "foo bar":
                    for fi19 := 0; fi19 < 5; fi19++ {
                        v1 = append(v1, 122)
                        var v20 int32 = 42
                        _ = fi19
                        _ = v20
                    }
                case "hello":
                    v21 := GenS1{F0: v15, F1: 0, F2: 0}
                    _ = v21
                case "a":
                    var v22 uint32 = 2
                    _ = v22
                }
                _ = fi17
                _ = v18
            }
            var v23 float64 = 1e10
            _ = v16
            _ = v23
        } else {
            v24 := GenS1{F0: (v15 - v15), F1: 72, F2: 100}
            _ = v24
        }
    }
    v25 := v2[1]
    v26 := GenS1{F0: 1, F1: 577, F2: 127}
    v25 -= (v25 % ((v25) | 1))
    fmt.Println(v15, v15)
    _ = len(v1)
    switch ("\n\t" + v3) {
    case "a":
        if (65535 < 0) {
            var v27 uint64 = (^(v15))
            for fi28 := 0; fi28 < 6; fi28++ {
                if (true || (0.1 < 1.5)) {
                    for fi29 := 0; fi29 < 14; fi29++ {
                        if ((127 >= 100) || (!(42 != 42))) {
                            delete(v2, uint32(v27))
                        } else {
                            v3 = v3
                            v3 = v3
                            v27 -= (v15 | 42)
                        }
                        _ = fi29
                    }
                    switch (v25 * -946) {
                    case 127:
                        if true {
                            for fi30 := 0; fi30 < 1; fi30++ {
                                var v31 string = "a"
                                switch "" {
                                case "":
                                    var v32 int64 = 23
                                    _ = v32
                                case "a":
                                    _ = len(v2)
                                case "hello":
                                    var v33 float32 = 3.14159
                                    _ = v33
                                }
                                _ = fi30
                                _ = v31
                            }
                            v34 := v2[253]
                            _ = v34
                        }
                    default:
                        v3 = v3
                    }
                }
                v15 += (v27 + 42)
                var v35 string = ""
                _ = fi28
                _ = v35
            }
            var v36 int64 = 9223372036854775807
            _ = v27
            _ = v36
        }
    case "foo bar":
        var v37 bool = (!false)
        _ = v37
    case "hello":
        for fi38 := 0; fi38 < 5; fi38++ {
            var v39 int8 = 50
            fmt.Println(v15)
            for fi40 := 0; fi40 < 16; fi40++ {
                switch "foo bar" {
                case "\n\t":
                    v41 := GenS2{F0: 339, F1: (fi40 - -1), F2: -505}
                    _ = v41
                case "foo bar":
                    var v42 int32 = 42
                    _ = v42
                }
                if (42 <= 935) {
                    v39 -= v39
                    var v43 float64 = 3.14159
                    _ = v43
                } else {
                    if true {
                        var v44 int = fi40
                        _ = v44
                    }
                    fmt.Println(fi38)
                    fmt.Println(fi38, fi40, v3)
                }
                switch v25 {
                case 32768:
                    var v45 float32 = float32(1.5)
                    _ = v45
                }
                _ = fi40
            }
            _ = fi38
            _ = v39
        }
    }
    v3 = v3
    switch v3 {
    case "\n\t":
        v46 := GenS1{F0: 0, F1: 18446744073709551615, F2: -64}
        _ = v46
    case "a":
        var v47 bool = (false || false)
        _ = v47
    default:
        if ((^(v15)) == v15) {
            var v48 int8 = 127
            _ = len(v1)
            _ = len(v2)
            _ = v48
        }
    }
    if len(v1) > 0 {
        v1[0] = 784
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v3
}
