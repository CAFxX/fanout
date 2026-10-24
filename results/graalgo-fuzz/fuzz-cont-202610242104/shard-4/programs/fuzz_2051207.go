package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int16 }

func main() {
    var v1 int32 = 2147483647
    var v2 int = 42
    v3 := []string{("hello" + "hello"), "foo bar"}
    var v4 uint64 = 42
    if false {
        _ = len(v3)
        v1 += v1
        var v5 float32 = (3.14159 + (0.1 - float32((3.14159 + -2.25))))
        _ = v5
    }
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 uint16 = 343
        var v8 uint = 654
        _ = fi6
        _ = v7
        _ = v8
    }
    switch 398 {
    case 392:
        for fi9 := 0; fi9 < 15; fi9++ {
            if false {
                switch (("a" + "\n\t") + ("" + "")) {
                case "\n\t":
                    if (4294967295 == 18446744073709551615) {
                        if (!(((v1 ^ 2147483647) != v1) || true)) {
                            if len(v3) > 0 {
                                v3[0] = ("\n\t" + "a")
                            }
                            _ = len(v3)
                            v3 = append(v3, "hello")
                        }
                        v2 = (^(v2))
                    } else {
                        v3 = append(v3, "\n\t")
                        _ = len(v3)
                    }
                case "":
                    switch "hello" {
                    case "a":
                        var v10 string = ""
                        _ = v10
                    }
                }
            }
            _ = fi9
        }
    }
    v3 = append(v3, (("\n\t" + "foo bar") + "\n\t"))
    if (!false) {
        for fi11 := 0; fi11 < 9; fi11++ {
            fmt.Println(v2, fi11, v4)
            _ = fi11
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v13 uint16 = 32767
            var v14 float32 = (1e10 + -2.25)
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    fmt.Println(v4, v1, v4)
    v2 = v2
    v1 += (v1 - 1)
    v2 -= 864
    if false {
        v15 := GenS1{F0: v2, F1: "a", F2: -32768}
        if len(v3) > 0 {
            v3[0] = ""
        }
        var v16 int64 = 0
        _ = v15
        _ = v16
    } else {
        v1 = (v1 % (((^(v1))) | 1))
    }
    fmt.Println(v4)
    v17 := GenS1{F0: 32768, F1: "a", F2: 32767}
    if false {
        for fi18 := 0; fi18 < 14; fi18++ {
            var v19 float64 = float64((1e10 + 0.0))
            var v20 float64 = float64(0.1)
            _ = fi18
            _ = v19
            _ = v20
        }
        fmt.Println(v1)
    } else {
        v21 := GenS1{F0: (v2 % (((v2 | (-(v2)))) | 1)), F1: "", F2: 1}
        v4 = 1
        _ = v21
    }
    if false {
        var v22 uint8 = 255
        _ = v22
    }
    fmt.Println(v4, v2, v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "foo bar" {
    case "foo bar":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v23 float64 = (-2.25 / 1.0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
