package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint }

func main() {
    v1 := []string{"", "foo bar", ("\n\t" + (("a" + (("\n\t" + ("a" + ((("\n\t" + "foo bar") + ("foo bar" + "")) + "\n\t"))) + ((("\n\t" + (("" + (("foo bar" + (("foo bar" + "foo bar") + "hello")) + ("\n\t" + "\n\t"))) + "")) + ("" + "hello")) + "a"))) + "")), ""}
    var v2 uint8 = 121
    var v3 int8 = 42
    var v4 int64 = -9223372036854775808
    var v5 uint16 = 42
    var v6 bool = (!true)
    var v7 bool = ((!(!false)) && v6)
    if (float64(-2.25) == 0.0) {
        var v8 uint32 = 802
        for fi9 := 0; fi9 < 8; fi9++ {
            switch (-(fi9)) {
            case 4294967295:
                v10 := GenS1{F0: ("\n\t" + "a")}
                _ = v10
            case 32768:
                fmt.Println(v2)
            }
            if len(v1) > 0 {
                v1[0] = ""
            }
            _ = fi9
        }
        var v11 uint64 = 18446744073709551615
        _ = v11
        _ = v8
    } else {
        if (1 == 0) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            v4 += (v4 * -576)
            fmt.Println(v6, v2, v3)
        }
    }
    switch 127 {
    case 816:
        v2 += (^(v2))
    case -9223372036854775808:
        var v12 uint64 = 0
        _ = v12
    default:
        if v7 {
            for fi13 := 0; fi13 < 16; fi13++ {
                var v14 string = "a"
                _ = fi13
                _ = v14
            }
            for fi15 := 0; fi15 < 16; fi15++ {
                v16 := GenS1{F0: "foo bar"}
                switch fi15 {
                case -9223372036854775808:
                    if (!(v4 < -9223372036854775808)) {
                        for fi17 := 0; fi17 < 17; fi17++ {
                            fmt.Println(v2, v2)
                            _ = fi17
                        }
                        v18 := GenS2{F0: -32768, F1: 65535}
                        v1 = append(v1, "\n\t")
                        _ = v18
                    }
                }
                _ = fi15
                _ = v16
            }
        } else {
            if false {
                v19 := GenS1{F0: "a"}
                v20 := GenS1{F0: "foo bar"}
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
                _ = v19
                _ = v20
            }
            for fi21 := 0; fi21 < 16; fi21++ {
                v2 -= v2
                v2 += v2
                if v6 {
                    var v22 uint = 0
                    var v23 uint32 = 65535
                    _ = v22
                    _ = v23
                }
                _ = fi21
            }
        }
    }
    v1 = append(v1, "")
    v5 = (v5 / (((v5 & (^(v5)))) | 1))
    var v24 int16 = -652
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
