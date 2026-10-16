package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 bool }
type GenS2 struct { F0 int8; F1 string }
func genFn0(p0 int16) int64 {
    v6 := []int{233, -697, 0, 32768}
    fmt.Println(p0)
    var v7 bool = false
    if v7 {
        var v8 string = ("" + "hello")
        _ = v8
    }
    switch 257 {
    case -2147483648:
        var v9 int64 = 9223372036854775807
        _ = v9
    default:
        var v10 bool = v7
        _ = v10
    }
    _ = p0
    _ = v6
    _ = v7
    return 9223372036854775807
}

func main() {
    v1 := []string{("\n\t" + ("" + (((("a" + "") + "a") + ((("" + ("a" + ((("hello" + "hello") + "foo bar") + "foo bar"))) + ("a" + "a")) + "hello")) + "\n\t"))), ("a" + ((("foo bar" + "hello") + ("hello" + "\n\t")) + "\n\t")), (("\n\t" + "hello") + "hello")}
    var v2 int32 = -1
    var v3 int16 = 508
    v4 := []string{("hello" + "foo bar")}
    var v5 int16 = 581
    if false {
        var v11 float64 = 0.0
        var v12 uint8 = 1
        if len(v4) > 0 {
            v4[0] = ("\n\t" + "hello")
        }
        _ = v11
        _ = v12
    } else {
        switch "a" {
        case "\n\t":
            for fi13 := 0; fi13 < 6; fi13++ {
                v5 = v5
                _ = fi13
            }
        case "a":
            var v14 uint32 = 838
            _ = v14
        case "":
            v3 = v3
        default:
            var v15 int8 = 100
            _ = v15
        }
    }
    v2 = (^(v2))
    fmt.Println(v3, v3)
    for fi16 := 0; fi16 < 2; fi16++ {
        v17 := GenS1{F0: (v3 - v5), F1: (fi16 * fi16), F2: false}
        var v18 string = (("\n\t" + "foo bar") + "hello")
        var v19 uint64 = 1
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    fmt.Println(v3)
    if false {
        if true {
            v5 = (v5 + -1)
            if (1.5 > 0.1) {
                fmt.Println(v2, v2, v2)
                var v20 uint64 = 42
                _ = v20
            } else {
                if ((2 != 18446744073709551615) || ((((((1e10 + float32((1e10 + (((1.5 * -2.25) - -2.25) + 1e10)))) + (0.0 + float32(1e10))) - ((1.5 + 0.1) * 3.14159)) + (3.14159 + float32((0.0 + 3.14159)))) + 3.14159) < 1e10)) {
                    v21 := GenS1{F0: (v5 % (((v3 ^ (v3 % (((v3 * v5)) | 1)))) | 1)), F1: 32768, F2: false}
                    var v22 float64 = 1e10
                    _ = v21
                    _ = v22
                } else {
                    v23 := GenS1{F0: 0, F1: 4294967295, F2: (2 == 65535)}
                    var v24 int8 = 1
                    _ = v23
                    _ = v24
                }
                for fi25 := 0; fi25 < 13; fi25++ {
                    for fi26 := 0; fi26 < 2; fi26++ {
                        var v27 uint32 = 2147483647
                        _ = len(v4)
                        v3 = (^(v5))
                        _ = fi26
                        _ = v27
                    }
                    for fi28 := 0; fi28 < 11; fi28++ {
                        fmt.Println(fi25)
                        _ = fi28
                    }
                    for fi29 := 0; fi29 < 17; fi29++ {
                        fmt.Println(fi29, v3, v5)
                        v4 = append(v4, ((("" + "hello") + "foo bar") + ("" + "foo bar")))
                        fmt.Println(v5, fi25, v2)
                        fmt.Println(fi29)
                        if len(v4) > 0 {
                            v4[0] = (("foo bar" + "\n\t") + ("hello" + "\n\t"))
                        }
                        _ = fi29
                    }
                    _ = fi25
                }
                for fi30 := 0; fi30 < 3; fi30++ {
                    fmt.Println(fi30, fi30, fi30)
                    _ = fi30
                }
            }
            v5 -= v3
        } else {
            v2 = (-(v2))
        }
    } else {
        v1 = append(v1, "hello")
        switch "a" {
        case "a":
            if len(v4) > 0 {
                v4[0] = ("foo bar" + ("a" + ((("a" + "") + ((("" + "a") + ("foo bar" + "a")) + "")) + (("\n\t" + "foo bar") + "\n\t"))))
            }
        default:
            var v31 int64 = -9223372036854775808
            _ = v31
        }
        switch ((("a" + "a") + "foo bar") + "a") {
        case "\n\t":
            for fi32 := 0; fi32 < 1; fi32++ {
                v5 += (v3 ^ 374)
                _ = fi32
            }
        }
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        v34 := GenS2{F0: 100, F1: "a"}
        var v35 uint16 = 32767
        for fi36 := 0; fi36 < 18; fi36++ {
            var v37 int8 = 100
            _ = fi36
            _ = v37
        }
        _ = fi33
        _ = v34
        _ = v35
    }
    fmt.Println(v2)
    if (659 <= 1) {
        var v38 int32 = (v2 / (((v2 & (v2 * 1))) | 1))
        _ = v38
    } else {
        var v39 string = ("hello" + "a")
        _ = v39
    }
    var v40 int = 127
    var v41 float32 = (0.1 / (3.14159 / -2.25))
    var v42 float32 = -2.25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
