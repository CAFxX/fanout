package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 uint; F2 uint }
type GenS5 struct { F0 uint32; F1 uint8 }

func main() {
    var v1 int32 = 2147483647
    v2 := []string{(("a" + "hello") + (("" + "") + "a")), "a"}
    v3 := []int{42, 128, 127}
    v4 := make(map[uint32]string)
    var v5 int8 = 0
    v5 = 0
    fmt.Println(v5)
    var v6 float32 = (1.5 - 1.5)
    for fi7 := 0; fi7 < 14; fi7++ {
        switch "" {
        case "hello":
            v5 -= (-(v5))
        case "\n\t":
            switch "\n\t" {
            case "a":
                var v8 float32 = v6
                _ = v8
            case "\n\t":
                switch (fi7 << ((888) & 63)) {
                case -38:
                    var v9 int8 = (v5 ^ -85)
                    _ = v9
                case 1:
                    var v10 uint64 = 1
                    _ = v10
                }
            default:
                v1 -= v1
            }
        case "a":
            fmt.Println(fi7, fi7)
        default:
            if false {
                var v11 int16 = 1000
                _ = v11
            } else {
                v12 := GenS1{F0: 42}
                var v13 uint16 = 729
                _ = v12
                _ = v13
            }
        }
        v1 -= (v1 ^ 0)
        if false {
            if false {
                v14 := GenS2{F0: 1}
                if false {
                    delete(v4, 171)
                    v15 := GenS2{F0: 1}
                    v16 := GenS2{F0: 4294967295}
                    _ = v15
                    _ = v16
                } else {
                    switch "\n\t" {
                    case "a":
                        for fi17 := 0; fi17 < 14; fi17++ {
                            v18 := GenS1{F0: (v5 * v5)}
                            v6 = v6
                            _ = fi17
                            _ = v18
                        }
                    default:
                        var v19 float32 = v6
                        _ = v19
                    }
                }
                _ = v14
            } else {
                v20 := GenS2{F0: 884}
                for fi21 := 0; fi21 < 1; fi21++ {
                    if (v5 != 100) {
                        if false {
                            v22 := v4[2147483647]
                            switch (("\n\t" + (v22 + (((v22 + v22) + v22) + v22))) + "\n\t") {
                            case "foo bar":
                                if true {
                                    _ = len(v3)
                                    var v23 uint32 = 2147483647
                                    v24 := v4[42]
                                    _ = v23
                                    _ = v24
                                } else {
                                    var v25 uint8 = 119
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    var v26 bool = false
                                    var v27 float32 = v6
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                }
                            default:
                                v6 = float32(3.14159)
                            }
                            _ = v22
                        }
                        fmt.Println(fi21)
                    }
                    v4[0] = ("a" + "")
                    if ((v1 * v1) < (-(v1))) {
                        delete(v4, 561)
                        if true {
                            switch "\n\t" {
                            case "":
                                v1 -= -137
                            case "\n\t":
                                var v28 uint64 = 258
                                _ = v28
                            default:
                                _ = len(v4)
                            }
                            v29 := GenS3{F0: false}
                            v5 = (v5 * (v5 * (v5 + 0)))
                            _ = v29
                        }
                    } else {
                        v1 -= v1
                        var v30 uint32 = 2147483648
                        _ = v30
                    }
                    _ = fi21
                }
                _ = v20
            }
            var v31 int32 = v1
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v31
        } else {
            v6 += ((((v6 - v6) - float32((-2.25 + (float64((1e10 / float32(-2.25))) / (1e10 - -2.25))))) / (-2.25 - (0.0 + v6))) + float32(-2.25))
        }
        _ = fi7
    }
    var v32 int = -1
    fmt.Println(v1)
    var v33 int64 = -9223372036854775808
    v34 := GenS4{F0: 18446744073709551615, F1: 838, F2: 136}
    v3 = append(v3, -955)
    delete(v4, 454)
    v35 := GenS5{F0: 477, F1: 2}
    fmt.Println(v32, v6)
    fmt.Println(v32)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
