package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{42}
    v3 := make(map[string]string)
    var v4 int64 = 9223372036854775807
    var v5 float64 = (1.5 * ((3.14159 - -2.25) / 0.1))
    var v6 float64 = v5
    v6 = v5
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        switch 0 {
        case 32768:
            v5 -= ((v5 + ((0.1 - float64(3.14159)) * 0.1)) * v6)
        case -249:
            var v7 int16 = 0
            _ = v7
        case 40:
            var v8 uint32 = 2147483647
            _ = v8
        }
        delete(v1, 1)
    }
    if (!false) {
        var v9 uint64 = 120
        var v10 uint64 = v9
        _ = v10
        _ = v9
    }
    var v11 uint64 = 1
    switch "\n\t" {
    case "a":
        if true {
            v2 = append(v2, -966)
            var v12 int32 = 1
            _ = v12
        }
    case "\n\t":
        if false {
            fmt.Println(v4)
            if true {
                v5 -= (v6 + v5)
                var v13 float32 = 1e10
                _ = v13
            }
        } else {
            v6 -= v6
            for fi14 := 0; fi14 < 1; fi14++ {
                v6 += ((v5 * v6) * -2.25)
                _ = fi14
            }
            v5 -= ((v5 + (0.1 / v6)) - float64(1.5))
        }
    }
    if true {
        v15 := GenS1{F0: v4, F1: 863}
        _ = v15
    } else {
        fmt.Println(v5)
        if true {
            v5 = v6
        } else {
            switch -2147483648 {
            case 32767:
                var v16 int8 = -128
                _ = v16
            case 883:
                switch "" {
                case "a":
                    if false {
                        delete(v3, "foo bar")
                    } else {
                        delete(v3, "a")
                    }
                case "\n\t":
                    fmt.Println(v11, v5, v6)
                case "foo bar":
                    v2 = append(v2, 401)
                default:
                    for fi17 := 0; fi17 < 13; fi17++ {
                        if (420 == -426) {
                            var v18 uint = 659
                            v3["a"] = ((("\n\t" + (("a" + "hello") + ("" + "a"))) + ("" + "hello")) + "hello")
                            switch (fi17 << ((65535) & 63)) {
                            case 269:
                                switch fi17 {
                                case -48:
                                    var v19 uint32 = 42
                                    _ = v19
                                case 128:
                                    v6 = float64(-2.25)
                                case -549:
                                    switch "" {
                                    case "hello":
                                        v11 = uint64(v18)
                                    }
                                default:
                                    v4 = v4
                                }
                            default:
                                _ = len(v3)
                            }
                            _ = v18
                        }
                        if false {
                            switch fi17 {
                            case -9223372036854775808:
                                for fi20 := 0; fi20 < 10; fi20++ {
                                    fmt.Println(fi17, fi17)
                                    switch 1 {
                                    case 4294967295:
                                        if true {
                                            if (1 == 198) {
                                                _ = len(v1)
                                                v5 = float64(0.0)
                                            } else {
                                                v4 = 42
                                                v5 += float64(0.1)
                                            }
                                            v21 := GenS2{F0: v4}
                                            _ = v21
                                        }
                                    default:
                                        v2 = append(v2, fi20)
                                    }
                                    v22 := GenS2{F0: 42}
                                    delete(v1, 821)
                                    _ = len(v2)
                                    _ = fi20
                                    _ = v22
                                }
                            case 42:
                                for fi23 := 0; fi23 < 15; fi23++ {
                                    var v24 int8 = -1
                                    _ = fi23
                                    _ = v24
                                }
                            case 1:
                                v1[4294967295] = "foo bar"
                            default:
                                var v25 float32 = 0.0
                                _ = v25
                            }
                            fmt.Println(v4, v5)
                        }
                        _ = fi17
                    }
                }
            case -847:
                v6 = (1e10 * -2.25)
            default:
                fmt.Println(v6)
            }
            if true {
                if ((((v6 / (((-2.25 + 0.1) * v5) + -2.25)) - 3.14159) - v5) <= (v5 * (v6 - (3.14159 - v5)))) {
                    var v26 bool = (!true)
                    _ = v26
                } else {
                    switch 2147483647 {
                    case 42:
                        fmt.Println(v5)
                    }
                    for fi27 := 0; fi27 < 14; fi27++ {
                        fmt.Println(v5, v4)
                        v6 = (3.14159 * v6)
                        _ = fi27
                    }
                    var v28 int64 = -673
                    _ = v28
                }
                v5 = v6
                fmt.Println(v6, v6, v4)
            } else {
                var v29 float64 = v6
                _ = v29
            }
        }
    }
    v4 = v4
    var v30 string = "foo bar"
    if false {
        var v31 string = v30
        _ = v31
    }
    if (float64(float32(float64((1.5 * float32(v6))))) > 3.14159) {
        var v32 int8 = 94
        v32 += (-(v32))
        _ = v32
    } else {
        fmt.Println(v6, v4, v6)
        v2 = append(v2, 32768)
        var v33 int64 = v4
        _ = v33
    }
    v34 := v1[598]
    v30 = ((v34 + v30) + (v30 + "foo bar"))
    v30 = (v34 + "foo bar")
    v11 = (v11 - v11)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
