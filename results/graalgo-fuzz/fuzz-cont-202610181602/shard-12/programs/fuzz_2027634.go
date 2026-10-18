package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 uint8 }
func genFn0() int32 {
    var v7 int64 = -9223372036854775808
    v8 := []int{int(v7), 32767}
    for fi9 := 0; fi9 < 9; fi9++ {
        switch -1 {
        case 32768:
            v7 += 991
        case 2147483647:
            v7 = v7
        case 42:
            var v10 int32 = 42
            _ = v10
        }
        _ = fi9
    }
    var v11 uint8 = 0
    _ = v11
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := []string{("\n\t" + ("hello" + "foo bar")), "\n\t"}
    var v2 int32 = 2147483647
    var v3 uint = 694
    var v4 uint16 = 868
    v5 := make(map[string]string)
    v6 := []string{((("hello" + (("a" + "hello") + ("hello" + (("foo bar" + "hello") + ("foo bar" + ("" + "a")))))) + "hello") + ("a" + "hello")), "foo bar"}
    v12 := GenS1{F0: (!false), F1: ("hello" + ""), F2: 2147483647}
    for fi13 := 0; fi13 < 12; fi13++ {
        switch fi13 {
        case -9223372036854775808:
            for fi14 := 0; fi14 < 12; fi14++ {
                switch fi13 {
                case 128:
                    var v15 int32 = 1
                    _ = v15
                case 4294967295:
                    v16 := v5["foo bar"]
                    switch fi13 {
                    case 367:
                        var v17 int = fi14
                        _ = v17
                    }
                    if (((0.1 + (1.5 / 3.14159)) <= 1.5) || false) {
                        v3 = (^(v3))
                        if true {
                            v2 -= (v2 * v2)
                            var v18 int64 = -607
                            _ = v18
                        }
                    } else {
                        var v19 bool = ((v2 ^ -1) == 123)
                        if (v16 < "a") {
                            fmt.Println(v16, v19)
                            v3 -= (v3 - (^(v3)))
                            if (((v4 + v4) == (v4 - 42)) && v19) {
                                v20 := GenS1{F0: v19, F1: v16, F2: 805}
                                _ = v20
                            } else {
                                var v21 uint16 = (v4 / ((v4) | 1))
                                fmt.Println(v4, fi13, fi14)
                                var v22 float32 = 1e10
                                var v23 uint8 = 1
                                v24 := GenS1{F0: v19, F1: v16, F2: 42}
                                _ = v21
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                        }
                        _ = v19
                    }
                    _ = v16
                default:
                    fmt.Println(fi14, v3, fi13)
                }
                _ = fi14
            }
        case -217:
            v25 := GenS2{F0: (^(fi13))}
            _ = v25
        }
        v2 -= (v2 >> (((^(v3))) & 63))
        _ = fi13
    }
    v3 = 491
    var v26 uint = 255
    v4 = v4
    var v27 uint64 = 42
    fmt.Println(v4)
    var v28 bool = false
    v3 -= v3
    var v29 string = "\n\t"
    v30 := GenS2{F0: 128}
    for fi31 := 0; fi31 < 1; fi31++ {
        var v32 uint32 = 211
        v33 := v5[v29]
        _ = fi31
        _ = v32
        _ = v33
    }
    if (!v28) {
        var v34 string = "foo bar"
        _ = v34
    } else {
        if v28 {
            if (42 > 155) {
                v35 := GenS3{F0: 32767, F1: 42}
                _ = v35
            } else {
                for fi36 := 0; fi36 < 20; fi36++ {
                    var v37 int32 = genFn0()
                    _ = fi36
                    _ = v37
                }
            }
        } else {
            var v38 float64 = 1e10
            _ = v38
        }
        var v39 string = "foo bar"
        for fi40 := 0; fi40 < 9; fi40++ {
            fmt.Println(fi40, v29)
            fmt.Println(v2)
            var v41 string = v39
            _ = fi40
            _ = v41
        }
        _ = v39
    }
    var v42 int16 = -1
    fmt.Println(v42, v29, v26)
    v26 = (v26 % (((v3 & (v3 | (v26 ^ (v26 % ((4294967295) | 1)))))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
