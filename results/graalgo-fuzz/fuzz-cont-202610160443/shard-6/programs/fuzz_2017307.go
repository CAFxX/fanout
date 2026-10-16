package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 int8; F1 uint; F2 int }
type GenS6 struct { F0 int16; F1 int }
type GenS7 struct { F0 uint8; F1 int64; F2 string }

func main() {
    v1 := make(map[string]bool)
    var v2 int = -552
    var v3 uint32 = 957
    var v4 uint8 = 127
    v5 := make(map[uint32]int)
    v6 := make(map[int]uint32)
    v3 += v3
    if true {
        if false {
            for fi7 := 0; fi7 < 7; fi7++ {
                switch ("a" + ("" + "a")) {
                case "hello":
                    if true {
                        if (255 > 42) {
                            if ((v4 >= (v4 << ((uint(v3)) & 63))) && true) {
                                switch v2 {
                                case -1:
                                    var v8 float64 = 0.0
                                    _ = v8
                                case 32768:
                                    fmt.Println(v4, v3)
                                default:
                                    v9 := GenS1{F0: 65535, F1: 32767, F2: 2}
                                    _ = v9
                                }
                            }
                            var v10 uint = 65535
                            _ = v10
                        } else {
                            v2 = 42
                        }
                        for fi11 := 0; fi11 < 18; fi11++ {
                            if ((v3 + 2147483647) >= 2147483648) {
                                v12 := GenS2{F0: "\n\t", F1: 18446744073709551615}
                                v13 := GenS3{F0: -790, F1: 42}
                                _ = v12
                                _ = v13
                            } else {
                                v4 = (^(v4))
                                var v14 int16 = 769
                                _ = v14
                            }
                            _ = fi11
                        }
                    }
                case "foo bar":
                    v3 = (v3 | v3)
                }
                var v15 int16 = -182
                for fi16 := 0; fi16 < 15; fi16++ {
                    if false {
                        v2 += fi16
                    } else {
                        switch "a" {
                        case "a":
                            _ = len(v6)
                        case "\n\t":
                            v17 := v1[("hello" + ("" + "\n\t"))]
                            _ = v17
                        default:
                            v18 := []int{(fi7 & (v2 >> ((255) & 63))), fi7}
                            _ = v18
                        }
                    }
                    _ = fi16
                }
                _ = fi7
                _ = v15
            }
        }
    } else {
        for fi19 := 0; fi19 < 11; fi19++ {
            v4 -= v4
            _ = fi19
        }
        var v20 int8 = -1
        var v21 uint = 4294967295
        _ = v20
        _ = v21
    }
    v22 := []int{v2, v2, 127, (-(v2))}
    v3 -= v3
    switch v2 {
    case 32767:
        for fi23 := 0; fi23 < 8; fi23++ {
            v24 := GenS4{F0: ("\n\t" + ("" + ("hello" + ""))), F1: 262}
            var v25 int = (v2 - fi23)
            if (int(v3) <= 210) {
                for fi26 := 0; fi26 < 3; fi26++ {
                    v27 := v1[(("hello" + "") + "hello")]
                    fmt.Println(v25)
                    v28 := v5[(v3 + v3)]
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
            _ = fi23
            _ = v24
            _ = v25
        }
    case -1:
        fmt.Println(v3, v3)
    }
    v22 = append(v22, 101)
    v1["\n\t"] = false
    for fi29 := 0; fi29 < 18; fi29++ {
        for fi30 := 0; fi30 < 16; fi30++ {
            switch (-(fi30)) {
            case 4294967295:
                switch "foo bar" {
                case "":
                    v31 := GenS5{F0: 42, F1: 323, F2: fi30}
                    var v32 string = "\n\t"
                    if false {
                        if ((961 == 65535) || (0 < -107)) {
                            var v33 float32 = 0.0
                            _ = v33
                        } else {
                            if len(v22) >= 2 {
                                v22 = v22[1:len(v22)]
                            }
                            switch (v32 + v32) {
                            case "\n\t":
                                v34 := GenS5{F0: 75, F1: 255, F2: (-(fi30))}
                                _ = v34
                            case "foo bar":
                                v35 := GenS2{F0: v32, F1: 0}
                                _ = v35
                            case "":
                                if true {
                                    if (!(!true)) {
                                        v36 := GenS3{F0: -1, F1: 42}
                                        _ = v36
                                    } else {
                                        fmt.Println(v3, fi29)
                                    }
                                    v37 := GenS6{F0: 0, F1: fi30}
                                    delete(v5, v3)
                                    _ = v37
                                } else {
                                    v3 -= 2147483647
                                }
                            }
                        }
                        if (231 == 2) {
                            switch v2 {
                            case -2147483648:
                                v38 := v1[(v32 + ((v32 + "foo bar") + "a"))]
                                _ = v38
                            }
                            v32 = v32
                            var v39 uint16 = 255
                            _ = v39
                        } else {
                            v40 := GenS4{F0: v32, F1: 255}
                            _ = v40
                        }
                    }
                    _ = v31
                    _ = v32
                }
                switch (("\n\t" + ("hello" + (("" + "foo bar") + ("hello" + "foo bar")))) + ((("" + "\n\t") + "hello") + (((("hello" + ("foo bar" + "\n\t")) + ("\n\t" + "foo bar")) + ((("foo bar" + "a") + "\n\t") + ("hello" + "a"))) + (((((("foo bar" + "hello") + "a") + ("a" + ("\n\t" + "hello"))) + ("\n\t" + "\n\t")) + "a") + "hello")))) {
                case "a":
                    if (9223372036854775807 < -9223372036854775808) {
                        var v41 float32 = (0.0 + 0.0)
                        switch 128 {
                        case 128:
                            var v42 uint16 = 2
                            _ = v42
                        default:
                            v43 := GenS5{F0: 103, F1: 42, F2: (-(v2))}
                            _ = v43
                        }
                        switch 128 {
                        case 1:
                            v2 = 32767
                        }
                        _ = v41
                    }
                }
                v44 := GenS7{F0: v4, F1: 9223372036854775807, F2: ""}
                _ = v44
            case -2147483648:
                if len(v22) > 0 {
                    v22[0] = (v2 / ((-1) | 1))
                }
            default:
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
            }
            _ = fi30
        }
        v2 = -9223372036854775808
        var v45 uint32 = (v3 + v3)
        _ = fi29
        _ = v45
    }
    var v46 int16 = 42
    if len(v22) > 0 {
        v22[0] = v2
    }
    v3 = (v3 & (v3 ^ 917))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
