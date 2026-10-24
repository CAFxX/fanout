package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }

func main() {
    var v1 uint16 = 2
    v2 := []string{((("hello" + ("\n\t" + "hello")) + "foo bar") + "\n\t"), ("foo bar" + "foo bar")}
    var v3 int32 = 1
    var v4 bool = false
    fmt.Println(v3, v3)
    var v5 int32 = 1
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    v6 := GenS1{F0: -983, F1: v5}
    v2 = append(v2, ("" + "hello"))
    if v4 {
        var v7 uint = 0
        switch (("foo bar" + ("" + ("\n\t" + (("" + ("a" + "hello")) + ("a" + "a"))))) + "") {
        case "hello":
            if v4 {
                fmt.Println(v1, v7, v7)
                _ = len(v2)
            }
        }
        _ = v7
    }
    v2 = append(v2, ("a" + "foo bar"))
    for fi8 := 0; fi8 < 18; fi8++ {
        _ = len(v2)
        if (!((1 == 0) && v4)) {
            if len(v2) > 0 {
                v2[0] = ("hello" + ("\n\t" + (((("a" + "a") + ("hello" + "")) + "\n\t") + "\n\t")))
            }
            var v9 bool = ((("hello" + ("hello" + "")) >= "a") || (true && (65535 <= 255)))
            _ = v9
        } else {
            switch fi8 {
            case -9223372036854775808:
                if true {
                    for fi10 := 0; fi10 < 16; fi10++ {
                        v4 = (!(351 >= (v5 * v5)))
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v3 -= (v5 ^ 709)
                        _ = fi10
                    }
                    if v4 {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v5 += 2147483647
                        fmt.Println(v3)
                    }
                    var v11 float64 = (1.5 - (0.1 - 1e10))
                    _ = v11
                } else {
                    var v12 uint8 = 127
                    var v13 int8 = -120
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v12
                    _ = v13
                }
            case 127:
                if len(v2) > 0 {
                    v2[0] = "a"
                }
            case -58:
                v14 := GenS1{F0: 1, F1: (v5 & (v5 & (v5 & 1)))}
                _ = v14
            }
        }
        _ = fi8
    }
    var v15 int32 = (-(v3))
    if (v4 || (-2.25 != float64(3.14159))) {
        v5 -= 2147483647
        v16 := GenS1{F0: -1, F1: (^(v5))}
        v15 = v3
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
