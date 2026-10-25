package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
func genFn0() uint8 {
    var v5 int32 = -374
    v6 := GenS1{F0: 106, F1: false}
    var v7 uint = 0
    v8 := []int{}
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 190
    v3 := make(map[string]string)
    v4 := []string{"", "\n\t"}
    var v9 float64 = 1.5
    if false {
        v9 -= -2.25
        v10 := GenS1{F0: 65535, F1: ((!true) && (true && true))}
        v9 -= (v9 / v9)
        _ = v10
    } else {
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
        v3[((("foo bar" + ("hello" + "a")) + "hello") + "")] = (("hello" + "a") + "foo bar")
        switch ("\n\t" + (("hello" + "") + "\n\t")) {
        case "foo bar":
            switch 88 {
            case -48:
                for fi11 := 0; fi11 < 16; fi11++ {
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(v2, fi11, v2)
                    case "\n\t":
                        v2 = (v2 + 663)
                    case "foo bar":
                        switch (fi11 * (fi11 | -1)) {
                        case 1:
                            switch fi11 {
                            case 128:
                                var v12 float64 = -2.25
                                _ = v12
                            default:
                                for fi13 := 0; fi13 < 12; fi13++ {
                                    if false {
                                        v9 = (0.0 + v9)
                                    }
                                    _ = fi13
                                }
                            }
                        case 32768:
                            switch "\n\t" {
                            case "hello":
                                delete(v1, ("\n\t" + "a"))
                            }
                        default:
                            if len(v4) > 0 {
                                v4[0] = "a"
                            }
                        }
                    }
                    v14 := v3[(("foo bar" + ("hello" + ("hello" + "hello"))) + "hello")]
                    _ = fi11
                    _ = v14
                }
            case 9223372036854775807:
                var v15 bool = (532 == 42)
                _ = v15
            case 128:
                v2 = (v2 & 18446744073709551615)
            }
        case "a":
            if len(v4) > 0 {
                v4[0] = ""
            }
        case "\n\t":
            for fi16 := 0; fi16 < 4; fi16++ {
                fmt.Println(fi16, fi16)
                var v17 float64 = v9
                v18 := GenS1{F0: 382, F1: (!false)}
                _ = fi16
                _ = v17
                _ = v18
            }
        }
    }
    v2 += v2
    v2 = v2
    fmt.Println(v2)
    for fi19 := 0; fi19 < 17; fi19++ {
        v9 -= float64(-2.25)
        switch ("hello" + "hello") {
        case "a":
            var v20 uint16 = 910
            _ = v20
        case "":
            if len(v4) > 0 {
                v4[0] = ("foo bar" + "a")
            }
        }
        fmt.Println(v9)
        _ = fi19
    }
    if len(v4) > 0 {
        v4[0] = ("\n\t" + ((("" + "foo bar") + "") + "a"))
    }
    v9 -= v9
    var v21 float64 = float64((0.0 - 3.14159))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v9
}
