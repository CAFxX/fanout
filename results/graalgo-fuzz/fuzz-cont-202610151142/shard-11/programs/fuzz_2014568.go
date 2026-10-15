package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int32 }
type GenS2 struct { F0 uint; F1 int16; F2 uint64 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 int16; F1 bool; F2 int32 }
type GenS5 struct { F0 uint8; F1 string }
type GenS6 struct { F0 int64 }

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{}
    v3 := make(map[int]int)
    var v4 uint64 = 412
    _ = len(v3)
    v4 = 213
    var v5 int8 = 0
    var v6 uint64 = (v4 - (^(v4)))
    for fi7 := 0; fi7 < 17; fi7++ {
        switch 0 {
        case 0:
            v5 = (v5 ^ 0)
        case 128:
            switch "" {
            case "a":
                for fi8 := 0; fi8 < 1; fi8++ {
                    v5 += (-(v5))
                    switch ("hello" + ("foo bar" + "")) {
                    case "hello":
                        switch 928 {
                        case 42:
                            var v9 float32 = 1.5
                            var v10 string = "foo bar"
                            if true {
                                var v11 int32 = 2147483647
                                if true {
                                    var v12 int64 = 0
                                    v13 := GenS1{F0: true, F1: (^(v5)), F2: -2147483648}
                                    _ = len(v2)
                                    _ = v12
                                    _ = v13
                                } else {
                                    v4 -= 0
                                    if true {
                                        v14 := GenS2{F0: 2, F1: -4, F2: 42}
                                        if (true && false) {
                                            v15 := GenS2{F0: uint(v6), F1: -420, F2: 4294967295}
                                            _ = v15
                                        } else {
                                            v9 -= float32(((3.14159 / (1.5 * (0.0 + ((float64(-2.25) - (((1e10 + 3.14159) * -2.25) - 0.1)) / (float64(float32(3.14159)) - ((0.1 / 1.0) - 3.14159)))))) / 1.0))
                                            v16 := GenS2{F0: 609, F1: 32767, F2: (^(v4))}
                                            fmt.Println(v11, fi7)
                                            if ((!true) || true) {
                                                v17 := GenS2{F0: 65535, F1: 1000, F2: (v6 << ((255) & 63))}
                                                v18 := GenS3{F0: v11, F1: 0}
                                                _ = v17
                                                _ = v18
                                            }
                                            switch (("hello" + v10) + v10) {
                                            case "hello":
                                                var v19 int64 = 42
                                                _ = v19
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = 2147483648
                                            }
                                            var v20 int8 = 0
                                            _ = v16
                                            _ = v20
                                        }
                                        v21 := GenS1{F0: false, F1: (v5 % (((v5 % ((v5) | 1))) | 1)), F2: (v11 - 507)}
                                        var v22 int8 = 109
                                        if (false || true) {
                                            var v23 float64 = 1.5
                                            var v24 float32 = v9
                                            v9 = 3.14159
                                            _ = v23
                                            _ = v24
                                        }
                                        _ = v14
                                        _ = v21
                                        _ = v22
                                    }
                                    var v25 bool = (true || false)
                                    _ = v25
                                }
                                v26 := GenS2{F0: 18446744073709551615, F1: 32767, F2: (v6 / ((2) | 1))}
                                _ = v11
                                _ = v26
                            }
                            _ = v10
                            _ = v9
                        case -54:
                            v6 += (^(v6))
                        case 422:
                            var v27 string = "hello"
                            fmt.Println(fi7)
                            var v28 bool = true
                            if len(v2) > 0 {
                                v2[0] = 2
                            }
                            _ = v27
                            _ = v28
                        }
                    case "":
                        var v29 float32 = 3.14159
                        _ = v29
                    default:
                        v30 := GenS2{F0: 65535, F1: 0, F2: (v6 / ((v6) | 1))}
                        _ = v30
                    }
                    v31 := GenS4{F0: 811, F1: true, F2: -2147483648}
                    _ = fi8
                    _ = v31
                }
            case "\n\t":
                fmt.Println(v5)
            default:
                if (v5 >= v5) {
                    v6 = v4
                } else {
                    v1["hello"] = false
                    fmt.Println(v6, v5)
                    var v32 uint64 = v6
                    _ = v32
                }
            }
        default:
            v4 -= v4
        }
        _ = fi7
    }
    v5 = v5
    var v33 int32 = 97
    if (230 >= 4294967295) {
        _ = len(v2)
        switch -2147483648 {
        case 353:
            _ = len(v1)
        case -2147483648:
            switch "" {
            case "a":
                var v34 bool = false
                _ = v34
            }
        default:
            if false {
                var v35 int64 = -731
                fmt.Println(v4, v6, v6)
                v36 := GenS5{F0: 190, F1: "hello"}
                _ = v35
                _ = v36
            }
        }
    }
    var v37 uint32 = 0
    v38 := GenS2{F0: 65535, F1: -377, F2: v4}
    switch (("\n\t" + ("foo bar" + "\n\t")) + ("a" + "foo bar")) {
    case "foo bar":
        v39 := GenS6{F0: -216}
        _ = v39
    }
    _ = len(v2)
    var v40 bool = (true || (-2.25 <= 0.0))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v37
    _ = v38
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
