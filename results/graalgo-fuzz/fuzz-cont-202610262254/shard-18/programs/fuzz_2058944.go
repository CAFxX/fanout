package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }

func main() {
    v1 := []string{"", "foo bar"}
    var v2 float64 = (1e10 - ((1e10 / (-2.25 * ((3.14159 - (float64(1e10) * -2.25)) - float64(0.0)))) + float64(((((1e10 / 3.14159) - 0.1) * -2.25) + 1.5))))
    v3 := []int{-2147483648}
    var v4 float64 = v2
    v5 := make(map[string]string)
    v2 += 1e10
    v6 := v5["\n\t"]
    v4 = float64((3.14159 * ((1.5 / 1e10) - 0.0)))
    delete(v5, "foo bar")
    v5[(("" + ("foo bar" + ((((v6 + v6) + "\n\t") + "hello") + (("foo bar" + (((v6 + v6) + ("\n\t" + "")) + "\n\t")) + "\n\t")))) + "\n\t")] = "hello"
    fmt.Println(v4, v2)
    v4 = ((v2 + 0.1) + 1e10)
    switch "a" {
    case "a":
        if false {
            if (!(!((!(!(65535 < 4294967295))) && false))) {
                fmt.Println(v2, v6, v2)
            }
        } else {
            _ = len(v5)
            for fi7 := 0; fi7 < 17; fi7++ {
                delete(v5, "\n\t")
                _ = fi7
            }
        }
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 uint64 = 1
        _ = fi8
        _ = v9
    }
    if false {
        if false {
            for fi10 := 0; fi10 < 2; fi10++ {
                if false {
                    v11 := GenS1{F0: 127, F1: fi10}
                    _ = v11
                }
                v12 := v5[v6]
                switch (fi10 << ((18446744073709551615) & 63)) {
                case 32767:
                    v13 := v5[v12]
                    _ = v13
                case -352:
                    if (fi10 < 32767) {
                        switch v6 {
                        case "":
                            v14 := GenS1{F0: 255, F1: 4294967295}
                            _ = v14
                        default:
                            var v15 int64 = 1
                            _ = v15
                        }
                    }
                case 0:
                    fmt.Println(v4, v6, v4)
                }
                _ = fi10
                _ = v12
            }
            var v16 uint8 = 255
            _ = v16
        } else {
            fmt.Println(v6)
        }
        fmt.Println(v4, v4)
        var v17 int32 = 2147483647
        _ = v17
    }
    var v18 float32 = (3.14159 * 3.14159)
    var v19 uint64 = 1
    var v20 uint16 = 823
    fmt.Println(v20, v4)
    var v21 float32 = float32(0.1)
    v6 = (v6 + "foo bar")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
