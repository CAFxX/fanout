package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint64; F1 int8; F2 uint8 }

func main() {
    var v1 int8 = -128
    var v2 uint16 = 1
    var v3 int32 = int32(v1)
    v4 := []int{-500, -9223372036854775808}
    var v5 bool = ((((("\n\t" + ("hello" + ("a" + (("a" + ("" + ("hello" + "a"))) + "")))) + "") + ("\n\t" + ("a" + (((("foo bar" + ("foo bar" + "a")) + "hello") + "hello") + "")))) + "") != "\n\t")
    var v6 float32 = 3.14159
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi7 := 0; fi7 < 19; fi7++ {
        if (-1 < (v3 % ((v3) | 1))) {
            if len(v4) > 0 {
                v4[0] = fi7
            }
            for fi8 := 0; fi8 < 9; fi8++ {
                if (1.5 > 1.5) {
                    var v9 string = "foo bar"
                    var v10 int16 = 122
                    _ = len(v4)
                    _ = v10
                    _ = v9
                }
                var v11 int64 = 9223372036854775807
                _ = fi8
                _ = v11
            }
            if (v5 || (0 > 713)) {
                if (1000 != 0) {
                    fmt.Println(v2, v6, v3)
                }
                var v12 uint8 = 42
                var v13 uint = 65535
                _ = v12
                _ = v13
            }
        }
        fmt.Println(v1, v3, v5)
        _ = fi7
    }
    var v14 int8 = -93
    v15 := GenS1{F0: "hello", F1: "\n\t"}
    var v16 int64 = 1
    v16 = (^(v16))
    for fi17 := 0; fi17 < 12; fi17++ {
        v18 := GenS1{F0: "hello", F1: "foo bar"}
        var v19 bool = (!true)
        var v20 uint64 = 254
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    v1 = (v14 ^ (v1 ^ (v1 * 35)))
    switch "foo bar" {
    case "":
        var v21 bool = v5
        _ = v21
    }
    v3 = -253
    if v5 {
        var v22 bool = (0 <= 2)
        v2 = 65535
        var v23 float64 = 0.0
        _ = v22
        _ = v23
    } else {
        if ("\n\t" >= "a") {
            v16 = v16
            for fi24 := 0; fi24 < 1; fi24++ {
                if v5 {
                    if (42 > 0) {
                        _ = len(v4)
                        v14 = (v14 % ((v1) | 1))
                    }
                    switch 32768 {
                    case 1:
                        var v25 int8 = (v1 | (v1 - v1))
                        _ = v25
                    }
                }
                for fi26 := 0; fi26 < 19; fi26++ {
                    v27 := GenS2{F0: 0, F1: (^(v1)), F2: 173}
                    v16 -= (v16 / ((v16) | 1))
                    var v28 uint32 = 386
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                _ = fi24
            }
            v6 += float32(3.14159)
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        switch "" {
        case "\n\t":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "":
            if v5 {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                fmt.Println(v2, v1)
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
