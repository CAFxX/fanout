package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 int32; F2 int8 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 int16 }

func main() {
    var v1 uint32 = 465
    v2 := []int{-325}
    var v3 float64 = float64((1e10 + 3.14159))
    v4 := make(map[string]uint32)
    var v5 int16 = -1
    var v6 uint32 = 2147483647
    v6 -= v1
    if (((-2.25 + -2.25) * (-2.25 / 1e10)) < 0.1) {
        v1 += v6
    } else {
        for fi7 := 0; fi7 < 17; fi7++ {
            delete(v4, "")
            _ = fi7
        }
    }
    var v8 int32 = 0
    if (638 != 393) {
        _ = len(v4)
    }
    if true {
        v9 := GenS1{F0: false}
        v10 := GenS1{F0: false}
        _ = v10
        _ = v9
    } else {
        if ((-29 == 0) && false) {
            var v11 uint = 42
            var v12 uint8 = 118
            var v13 uint32 = 134
            _ = v11
            _ = v12
            _ = v13
        } else {
            if false {
                for fi14 := 0; fi14 < 18; fi14++ {
                    v15 := GenS2{F0: (^(fi14)), F1: 5, F2: 42}
                    switch "\n\t" {
                    case "foo bar":
                        delete(v4, "\n\t")
                    case "a":
                        _ = len(v2)
                    case "\n\t":
                        for fi16 := 0; fi16 < 8; fi16++ {
                            var v17 uint64 = 2
                            if true {
                                switch fi14 {
                                case 32768:
                                    v5 -= (v5 & v5)
                                }
                            }
                            var v18 int8 = 82
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    }
                    switch (fi14 - (fi14 * fi14)) {
                    case 655:
                        for fi19 := 0; fi19 < 6; fi19++ {
                            var v20 uint32 = v1
                            var v21 bool = false
                            v4["\n\t"] = v20
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                    }
                    _ = fi14
                    _ = v15
                }
                var v22 int64 = -80
                v23 := GenS2{F0: 42, F1: v8, F2: 42}
                _ = v22
                _ = v23
            } else {
                if (!(!(42 >= 1))) {
                    _ = len(v2)
                    if true {
                        for fi24 := 0; fi24 < 4; fi24++ {
                            fmt.Println(v1, fi24, v1)
                            switch "foo bar" {
                            case "":
                                v1 = v6
                            case "\n\t":
                                var v25 float64 = (v3 - v3)
                                _ = v25
                            case "foo bar":
                                delete(v4, ("\n\t" + "\n\t"))
                            }
                            _ = fi24
                        }
                        if true {
                            v26 := GenS3{F0: 401, F1: -106}
                            if false {
                                v1 += (v6 * (v6 - v1))
                            } else {
                                v3 = v3
                                switch -736 {
                                case 32768:
                                    var v27 uint32 = 4294967295
                                    _ = v27
                                default:
                                    var v28 int16 = (v5 - 32767)
                                    _ = v28
                                }
                                v29 := GenS4{F0: (v5 >> ((368) & 63))}
                                _ = v29
                            }
                            _ = v26
                        } else {
                            fmt.Println(v3)
                            for fi30 := 0; fi30 < 3; fi30++ {
                                for fi31 := 0; fi31 < 17; fi31++ {
                                    v32 := v4[""]
                                    _ = fi31
                                    _ = v32
                                }
                                switch fi30 {
                                case -1:
                                    v33 := v4[("" + "a")]
                                    _ = v33
                                }
                                _ = fi30
                            }
                        }
                    } else {
                        var v34 uint32 = v1
                        if (false || false) {
                            if false {
                                switch 771 {
                                case -9223372036854775808:
                                    v4["hello"] = (v1 | v34)
                                case -2147483648:
                                    v35 := GenS1{F0: false}
                                    _ = v35
                                }
                            }
                            v1 += (v1 + (v6 * (v6 | v6)))
                            if false {
                                v36 := GenS4{F0: v5}
                                v3 -= ((v3 * float64(0.1)) - (float64(float32((v3 * (v3 * (((v3 - v3) / (0.1 * (v3 - v3))) + (v3 * v3)))))) - float64(1e10)))
                                _ = v36
                            }
                        } else {
                            var v37 int8 = -128
                            var v38 uint32 = v6
                            _ = v37
                            _ = v38
                        }
                        _ = v34
                    }
                }
                if (v5 >= (v5 % ((v5) | 1))) {
                    for fi39 := 0; fi39 < 6; fi39++ {
                        fmt.Println(v6)
                        v4[(((("hello" + ((("foo bar" + "a") + (("a" + "") + "a")) + "a")) + ("foo bar" + "")) + "\n\t") + (("hello" + "hello") + (("hello" + ("\n\t" + "\n\t")) + "foo bar")))] = v1
                        v6 += v6
                        _ = fi39
                    }
                    v8 = (v8 - v8)
                    delete(v4, ("" + ("hello" + "hello")))
                } else {
                    fmt.Println(v6, v1)
                }
            }
            if len(v2) > 0 {
                v2[0] = 0
            }
        }
        v4["hello"] = (v1 + v6)
    }
    v1 += (v1 * v1)
    var v40 uint8 = 153
    fmt.Println(v40, v5)
    var v41 uint64 = 0
    var v42 bool = (42 >= 42)
    fmt.Println(v6, v40, v8)
    var v43 bool = ((v1 - v1) == v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v8
}
