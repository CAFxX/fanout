package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 string }
type GenS2 struct { F0 uint16; F1 int32; F2 int32 }
type GenS3 struct { F0 uint8; F1 int; F2 int16 }

func main() {
    var v1 int64 = -620
    v2 := make(map[uint32]bool)
    v3 := []int{-9223372036854775808, 128}
    var v4 float64 = 3.14159
    switch "" {
    case "":
        var v5 float32 = 1.5
        _ = v5
    case "hello":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    delete(v2, 65535)
    v4 += (0.1 * (v4 * v4))
    if false {
        if len(v3) > 0 {
            v3[0] = -2147483648
        }
        var v6 uint = 4294967295
        switch "foo bar" {
        case "foo bar":
            v1 = 9223372036854775807
        case "":
            v7 := GenS1{F0: 2147483648, F1: v1, F2: ""}
            _ = v7
        default:
            if true {
                var v8 int32 = 0
                _ = v8
            } else {
                if false {
                    var v9 uint16 = 2
                    var v10 uint16 = (v9 & v9)
                    _ = v10
                    _ = v9
                } else {
                    var v11 float32 = 0.0
                    _ = v11
                }
                if ((65535 != 587) && (true && false)) {
                    if true {
                        var v12 int64 = (-(v1))
                        _ = v12
                    } else {
                        if true {
                            if false {
                                v13 := GenS1{F0: 707, F1: -1, F2: "a"}
                                var v14 int = 9223372036854775807
                                _ = v13
                                _ = v14
                            } else {
                                if (2 > 42) {
                                    if len(v3) > 0 {
                                        v3[0] = 1
                                    }
                                } else {
                                    for fi15 := 0; fi15 < 5; fi15++ {
                                        if true {
                                            for fi16 := 0; fi16 < 8; fi16++ {
                                                if true {
                                                    fmt.Println(v6)
                                                    v4 = v4
                                                    if false {
                                                        v1 += v1
                                                        v17 := v2[2147483647]
                                                        _ = v17
                                                    }
                                                    v2[616] = (v1 == v1)
                                                    v4 = v4
                                                }
                                                var v18 uint = 942
                                                var v19 int16 = -1
                                                switch fi15 {
                                                case -1:
                                                    var v20 float64 = (float64(0.1) / v4)
                                                    v4 = v20
                                                    _ = v20
                                                case 127:
                                                    fmt.Println(v4, v18, fi16)
                                                default:
                                                    var v21 float32 = (1e10 - (0.0 + (0.0 - 0.0)))
                                                    _ = v21
                                                }
                                                _ = fi16
                                                _ = v18
                                                _ = v19
                                            }
                                            switch (((((("hello" + "hello") + ("foo bar" + "a")) + "foo bar") + ("" + ("a" + "a"))) + ("\n\t" + "\n\t")) + ("" + "hello")) {
                                            case "a":
                                                var v22 float32 = (1.5 * (-2.25 + 0.0))
                                                _ = v22
                                            case "\n\t":
                                                _ = len(v2)
                                            case "hello":
                                                if (!(!false)) {
                                                    var v23 int32 = 42
                                                    switch (((("\n\t" + ("hello" + ("foo bar" + ("a" + "a")))) + ("a" + "")) + "\n\t") + "") {
                                                    case "\n\t":
                                                        var v24 int16 = 1000
                                                        _ = v24
                                                    default:
                                                        var v25 uint32 = 849
                                                        _ = v25
                                                    }
                                                    _ = v23
                                                }
                                            }
                                            switch (fi15 - (fi15 << (((^(v6))) & 63))) {
                                            case 737:
                                                _ = len(v2)
                                            }
                                        } else {
                                            var v26 bool = true
                                            _ = v26
                                        }
                                        for fi27 := 0; fi27 < 4; fi27++ {
                                            _ = len(v3)
                                            _ = fi27
                                        }
                                        _ = fi15
                                    }
                                }
                                delete(v2, 0)
                            }
                        }
                        if false {
                            fmt.Println(v1)
                            var v28 int = -86
                            fmt.Println(v1, v28)
                            _ = v28
                        }
                    }
                }
                fmt.Println(v4, v1, v6)
            }
        }
        _ = v6
    } else {
        fmt.Println(v1)
        var v29 bool = (v1 < (^(v1)))
        _ = v29
    }
    _ = len(v2)
    if (false || (false && (0 > 42))) {
        if (false || true) {
            if true {
                var v30 int32 = 0
                if false {
                    for fi31 := 0; fi31 < 16; fi31++ {
                        fmt.Println(v4, fi31)
                        fmt.Println(v4, v1)
                        _ = fi31
                    }
                    switch -558 {
                    case 0:
                        v1 = (^(v1))
                    case 42:
                        if false {
                            var v32 int = 4294967295
                            _ = v32
                        }
                    }
                } else {
                    v3 = append(v3, -293)
                    for fi33 := 0; fi33 < 3; fi33++ {
                        v4 += (((v4 + ((0.0 - (v4 * ((v4 / (v4 - float64(1.5))) / ((v4 / (1e10 * v4)) * 0.0)))) * float64(0.1))) + 1e10) + 1.5)
                        v34 := v2[137]
                        _ = fi33
                        _ = v34
                    }
                    var v35 int = 2147483647
                    _ = v35
                }
                fmt.Println(v1)
                _ = v30
            } else {
                if len(v3) > 0 {
                    v3[0] = 32768
                }
                v2[42] = (731 < 181)
                var v36 uint16 = 2
                _ = v36
            }
            var v37 int32 = -1
            _ = len(v2)
            _ = v37
        } else {
            v1 = v1
            switch ("hello" + "a") {
            case "a":
                if ((true && (true || (false || (v4 != (-2.25 + v4))))) && true) {
                    v38 := GenS2{F0: 620, F1: 42, F2: -536}
                    v39 := GenS2{F0: 65535, F1: -1, F2: -1}
                    _ = v38
                    _ = v39
                }
            default:
                v40 := GenS3{F0: 1, F1: 1, F2: 0}
                _ = v40
            }
            v41 := GenS1{F0: 2147483647, F1: (v1 + (v1 - v1)), F2: (("a" + "a") + (("hello" + "hello") + ((("foo bar" + "a") + ("\n\t" + ("\n\t" + ("hello" + "\n\t")))) + ("a" + "hello"))))}
            _ = v41
        }
        for fi42 := 0; fi42 < 3; fi42++ {
            for fi43 := 0; fi43 < 18; fi43++ {
                v1 -= v1
                var v44 bool = true
                _ = fi43
                _ = v44
            }
            _ = fi42
        }
    } else {
        v1 += v1
    }
    var v45 string = "hello"
    fmt.Println(v4, v45, v45)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
}
