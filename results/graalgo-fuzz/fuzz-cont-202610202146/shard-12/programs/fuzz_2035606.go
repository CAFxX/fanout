package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint }
type GenS3 struct { F0 string }
func genFn0(p0 string, p1 bool, p2 uint) uint8 {
    for fi7 := 0; fi7 < 10; fi7++ {
        if p1 {
            v8 := GenS1{F0: 9223372036854775807, F1: uint16(p2), F2: (^(p2))}
            _ = v8
        }
        var v9 uint = p2
        _ = fi7
        _ = v9
    }
    v10 := GenS1{F0: 1, F1: 65535, F2: 363}
    var v11 int = -2147483648
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 42
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]string)
    var v3 int32 = 2147483647
    v4 := []string{("a" + "hello"), (("\n\t" + "foo bar") + "foo bar"), "foo bar", "foo bar"}
    v5 := []int{216, 9223372036854775807, 32767, 513}
    v6 := make(map[uint32]bool)
    fmt.Println(v3)
    v3 += v3
    v4 = append(v4, ("a" + "a"))
    var v12 uint16 = 49
    fmt.Println(v3, v3)
    v12 = v12
    var v13 uint8 = 255
    fmt.Println(v3, v3)
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 string = ""
        _ = fi14
        _ = v15
    }
    v13 += (^(v13))
    v12 += (^(v12))
    v12 -= (v12 ^ v12)
    for fi16 := 0; fi16 < 13; fi16++ {
        v3 = v3
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 bool = true
            if (v18 || v18) {
                var v19 int = 0
                if len(v5) > 0 {
                    v5[0] = v19
                }
                if false {
                    v20 := GenS2{F0: (v13 ^ 1), F1: 471, F2: 582}
                    if (3.14159 < (1e10 + (float64((-2.25 * (0.0 * ((-2.25 * (0.0 * (0.0 * 1e10))) - 0.1)))) - ((1e10 / 1.0) / 0.1)))) {
                        if v18 {
                            fmt.Println(v18, v12, fi16)
                        }
                    } else {
                        switch -453 {
                        case 0:
                            v12 = v12
                        case 42:
                            var v21 string = ("foo bar" + "foo bar")
                            _ = v21
                        default:
                            var v22 uint = 878
                            _ = v22
                        }
                        v18 = v18
                        var v23 int32 = (^(v3))
                        if v18 {
                            if false {
                                switch "foo bar" {
                                case "":
                                    v23 -= v23
                                case "\n\t":
                                    v5 = append(v5, (v19 - -1))
                                case "foo bar":
                                    v5 = append(v5, (fi17 * fi17))
                                }
                                var v24 int64 = 1
                                _ = v24
                            } else {
                                v25 := GenS3{F0: "\n\t"}
                                _ = v25
                            }
                        } else {
                            if true {
                                switch (("a" + "") + ("hello" + "foo bar")) {
                                case "a":
                                    switch (fi16 << ((261) & 63)) {
                                    case 948:
                                        var v26 uint8 = 144
                                        _ = v26
                                    case 65:
                                        var v27 bool = ((1e10 / 1.5) == float32((1.5 - -2.25)))
                                        _ = v27
                                    case 32767:
                                        var v28 uint16 = (^(v12))
                                        _ = v28
                                    default:
                                        fmt.Println(v18, v19)
                                    }
                                case "":
                                    _ = len(v4)
                                    v23 -= 1
                                    switch "a" {
                                    case "\n\t":
                                        if v18 {
                                            if false {
                                                var v29 int = fi17
                                                _ = v29
                                            }
                                        }
                                    case "foo bar":
                                        delete(v1, v19)
                                    }
                                case "foo bar":
                                    switch "a" {
                                    case "hello":
                                        fmt.Println(fi17, fi17, v12)
                                    case "":
                                        var v30 int16 = 612
                                        _ = v30
                                    default:
                                        var v31 int16 = 32767
                                        _ = v31
                                    }
                                default:
                                    fmt.Println(v3, v18)
                                }
                                if ("" == "foo bar") {
                                    var v32 uint8 = (v13 + genFn0("a", (!v18), 0))
                                    _ = v32
                                } else {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                            }
                        }
                        delete(v1, (fi16 % (((^(fi16))) | 1)))
                        _ = v23
                    }
                    _ = v20
                }
                _ = v19
            }
            if len(v5) > 0 {
                v5[0] = fi16
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    if (1 <= 42) {
        fmt.Println(v12, v13)
    }
    if false {
        for fi33 := 0; fi33 < 15; fi33++ {
            if (!(true || (1 != 457))) {
                v2[""] = (("a" + "\n\t") + ("a" + ("\n\t" + ("\n\t" + "foo bar"))))
            }
            _ = fi33
        }
        v34 := v1[1]
        _ = v34
    }
    var v35 int = -724
    v36 := GenS3{F0: (("hello" + "") + (("a" + ((("foo bar" + "a") + "a") + "a")) + ""))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
