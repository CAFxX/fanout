package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 int32; F1 uint32; F2 uint }
type GenS4 struct { F0 uint8; F1 bool }
func genFn0() uint32 {
    var v7 uint8 = 26
    var v8 string = "foo bar"
    v7 -= v7
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 bool = (!(!false))
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: -39}
    _ = v11
    _ = v7
    _ = v8
    return 834
}

func main() {
    v1 := []string{"a", "foo bar", "\n\t", ("" + "hello")}
    v2 := []int{-679}
    var v3 int8 = -1
    var v4 int32 = 77
    var v5 int = 517
    var v6 int16 = 1000
    var v12 int8 = -128
    if false {
        switch 42 {
        case 32767:
            if (!false) {
                var v13 uint8 = 255
                v12 -= v12
                _ = v13
            }
        case -2147483648:
            if (42 > 399) {
                for fi14 := 0; fi14 < 15; fi14++ {
                    v15 := GenS1{F0: v3}
                    _ = fi14
                    _ = v15
                }
                for fi16 := 0; fi16 < 20; fi16++ {
                    fmt.Println(v4)
                    var v17 int16 = v6
                    var v18 float64 = 1e10
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
        case 475:
            v12 += v12
        default:
            fmt.Println(v3, v3)
        }
    }
    var v19 uint = 65535
    switch "hello" {
    case "":
        for fi20 := 0; fi20 < 4; fi20++ {
            v21 := GenS2{F0: 18446744073709551615, F1: 2147483648}
            switch -1 {
            case 796:
                var v22 bool = false
                _ = v22
            case 67:
                switch v5 {
                case -25:
                    fmt.Println(v19, v19, v4)
                }
            case -2147483648:
                v5 += v5
            }
            var v23 uint16 = 65535
            _ = fi20
            _ = v21
            _ = v23
        }
    case "a":
        fmt.Println(v5, v3)
    case "\n\t":
        var v24 uint32 = genFn0()
        _ = v24
    }
    if true {
        for fi25 := 0; fi25 < 7; fi25++ {
            v3 += (^(v12))
            _ = fi25
        }
        v12 = v12
        switch v5 {
        case 0:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            var v26 bool = (!true)
            _ = v26
        }
    } else {
        if (false || false) {
            v12 = -1
        } else {
            switch ("foo bar" + ("hello" + "foo bar")) {
            case "":
                v4 -= (^(v4))
            default:
                if (v5 != -1) {
                    for fi27 := 0; fi27 < 10; fi27++ {
                        switch "\n\t" {
                        case "":
                            var v28 int = fi27
                            _ = v28
                        case "a":
                            v29 := GenS3{F0: (v4 + -2147483648), F1: 65535, F2: (^(v19))}
                            _ = v29
                        case "foo bar":
                            if false {
                                var v30 uint8 = 127
                                _ = v30
                            }
                        default:
                            v31 := GenS3{F0: 47, F1: 777, F2: 255}
                            _ = v31
                        }
                        _ = fi27
                    }
                    v6 = v6
                }
            }
            switch v5 {
            case 32767:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 2147483647:
                switch (("a" + ("\n\t" + ((("hello" + "foo bar") + ((("a" + "hello") + ("foo bar" + "a")) + "")) + "a"))) + "foo bar") {
                case "\n\t":
                    var v32 uint = (v19 + 255)
                    _ = v32
                }
            default:
                v4 += 1
            }
            v19 += (^(v19))
        }
        if (1.5 >= 3.14159) {
            if false {
                if (!true) {
                    if (0 > 271) {
                        var v33 int64 = 9223372036854775807
                        _ = v33
                    }
                    fmt.Println(v4, v19, v4)
                }
                var v34 float32 = 3.14159
                v5 = (v5 | 32768)
                _ = v34
            } else {
                v35 := GenS3{F0: v4, F1: 1, F2: v19}
                _ = v35
            }
        }
        fmt.Println(v6)
    }
    v12 += v3
    if (genFn0() != 42) {
        if true {
            v12 = v3
        } else {
            v36 := GenS2{F0: 579, F1: 2}
            var v37 float64 = (1.5 * 0.1)
            _ = v36
            _ = v37
        }
        v3 = (v3 - (v12 << ((v19) & 63)))
    } else {
        v19 -= v19
        var v38 int = (v5 % ((v5) | 1))
        if (v6 <= (v6 * (v6 % (((v6 - (v6 * (-(v6))))) | 1)))) {
            v39 := GenS2{F0: 4294967295, F1: genFn0()}
            _ = v39
        }
        _ = v38
    }
    fmt.Println(v12, v6, v6)
    v19 = v19
    for fi40 := 0; fi40 < 17; fi40++ {
        var v41 uint32 = 2147483647
        fmt.Println(fi40, v6, v5)
        _ = fi40
        _ = v41
    }
    fmt.Println(v3, v12, v12)
    if len(v2) > 0 {
        v2[0] = 1
    }
    var v42 bool = true
    v43 := GenS4{F0: 255, F1: v42}
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
