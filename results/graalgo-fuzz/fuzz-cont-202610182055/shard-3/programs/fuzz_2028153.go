package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := []string{(("a" + "foo bar") + "\n\t"), ("hello" + "foo bar"), "foo bar"}
    v2 := []int{}
    var v3 int = -1
    var v4 uint = 42
    v5 := []string{("foo bar" + "\n\t"), ((("" + "a") + "\n\t") + (("\n\t" + "hello") + "hello"))}
    if false {
        v4 = 4294967295
        var v6 int64 = -427
        _ = v6
    } else {
        if (!true) {
            if (((0.1 + (((1.5 * 3.14159) * ((float64(0.1) / 1e10) + (3.14159 / 0.1))) - float64(1e10))) / float64(1e10)) != 3.14159) {
                v3 -= v3
            } else {
                switch "a" {
                case "a":
                    if (1e10 > -2.25) {
                        v4 -= 4294967295
                        var v7 uint8 = 255
                        for fi8 := 0; fi8 < 8; fi8++ {
                            if len(v2) > 0 {
                                v2[0] = fi8
                            }
                            _ = fi8
                        }
                        _ = v7
                    }
                default:
                    fmt.Println(v3, v3)
                }
                v5 = append(v5, (("foo bar" + "hello") + "hello"))
            }
            var v9 int8 = 1
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 int = v3
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: -32768}
    var v13 uint64 = 185
    for fi14 := 0; fi14 < 7; fi14++ {
        if ((-(fi14)) <= v3) {
            fmt.Println(v4)
            v13 -= 42
        }
        _ = fi14
    }
    if false {
        for fi15 := 0; fi15 < 7; fi15++ {
            if len(v1) > 0 {
                v1[0] = ((("\n\t" + "a") + ("" + "hello")) + "")
            }
            v16 := GenS2{F0: 42}
            _ = fi15
            _ = v16
        }
        _ = len(v2)
        v17 := GenS2{F0: -1}
        _ = v17
    }
    v4 = 1
    switch "foo bar" {
    case "":
        if (v4 != 237) {
            if (true && false) {
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
                var v18 int16 = -32768
                _ = v18
            }
        } else {
            var v19 uint = v4
            v20 := GenS2{F0: -464}
            _ = v19
            _ = v20
        }
    }
    if ((v3 % (((v3 % (((^(v3))) | 1))) | 1)) > 32768) {
        var v21 int64 = 0
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
