package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 string }

func main() {
    var v1 int = -822
    var v2 uint = 4294967295
    v3 := []string{"", ("a" + "\n\t"), (("" + "a") + (("hello" + ("hello" + ("hello" + ("" + "")))) + (("foo bar" + ("hello" + ("" + "foo bar"))) + ("a" + ((("a" + "hello") + "a") + "\n\t"))))), "\n\t"}
    var v4 int8 = -128
    var v5 uint16 = 1
    var v6 uint = (v2 + (v2 + (v2 << (((v2 >> (((v2 - v2)) & 63))) & 63))))
    v7 := GenS1{F0: 65535}
    v8 := GenS1{F0: 4294967295}
    v9 := GenS1{F0: 423}
    for fi10 := 0; fi10 < 6; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            v1 += (fi11 + fi10)
            _ = fi11
        }
        _ = fi10
    }
    fmt.Println(v5, v6, v5)
    var v12 int64 = 849
    var v13 int64 = (^(v12))
    switch 4294967295 {
    case -105:
        if len(v3) > 0 {
            v3[0] = ""
        }
    case 128:
        v1 = -81
    }
    v14 := GenS1{F0: 42}
    var v15 float32 = (1e10 - ((0.0 - 3.14159) + (0.0 / -2.25)))
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 int64 = -1
        if true {
            v15 += v15
            if ("" >= "\n\t") {
                v5 += (v5 + v5)
                for fi18 := 0; fi18 < 9; fi18++ {
                    fmt.Println(v12)
                    _ = fi18
                }
            } else {
                var v19 uint32 = 2147483647
                if len(v3) > 0 {
                    v3[0] = ("\n\t" + ("" + (((((("hello" + "") + "a") + ("" + ("a" + ("hello" + (("hello" + ("a" + "")) + "hello"))))) + ((("a" + "foo bar") + "") + "\n\t")) + (("hello" + "\n\t") + "hello")) + "\n\t")))
                }
                _ = v19
            }
            if true {
                if ((fi16 * fi16) <= -177) {
                    switch "" {
                    case "":
                        if false {
                            v13 += v13
                            _ = len(v3)
                        } else {
                            fmt.Println(v15)
                        }
                    case "\n\t":
                        v15 = float32(1.5)
                    }
                } else {
                    var v20 int32 = 1
                    _ = v20
                }
                v21 := GenS1{F0: 848}
                _ = v21
            } else {
                _ = len(v3)
                var v22 int64 = v13
                v23 := GenS2{F0: 18446744073709551615, F1: "a"}
                _ = v22
                _ = v23
            }
        }
        fmt.Println(v13, v4)
        _ = fi16
        _ = v17
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        fmt.Println(v2, v5)
        _ = fi24
    }
    v4 += (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
