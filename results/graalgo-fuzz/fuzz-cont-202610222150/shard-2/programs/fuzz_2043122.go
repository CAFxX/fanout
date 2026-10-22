package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() bool {
    var v6 uint16 = 940
    switch ("a" + "foo bar") {
    case "hello":
        var v7 uint8 = 128
        _ = v7
    case "":
        fmt.Println(v6)
    }
    _ = v6
    return false
}

func main() {
    v1 := []string{"hello", ("hello" + "\n\t"), "a", "\n\t"}
    v2 := make(map[string]bool)
    var v3 int32 = 42
    v4 := []string{"a", "a"}
    var v5 int32 = (v3 | (v3 % (((^(v3))) | 1)))
    var v8 float64 = 0.0
    for fi9 := 0; fi9 < 15; fi9++ {
        v5 += v3
        v8 = v8
        var v10 int16 = 42
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 2147483647}
    for fi12 := 0; fi12 < 3; fi12++ {
        for fi13 := 0; fi13 < 10; fi13++ {
            if false {
                delete(v2, ("hello" + ("hello" + "a")))
                if (0 <= 127) {
                    var v14 uint = 470
                    _ = v14
                } else {
                    fmt.Println(v5, fi13)
                }
            }
            switch (^(fi13)) {
            case 32768:
                if (760 < 2) {
                    var v15 uint32 = 122
                    v8 -= ((-2.25 - v8) / ((v8 / (1e10 * v8)) * (v8 / (float64(0.1) * (((1e10 * v8) / v8) * v8)))))
                    v3 = v3
                    _ = v15
                }
            case 9223372036854775807:
                var v16 int16 = 1000
                _ = v16
            case -1:
                var v17 float64 = (v8 - v8)
                _ = v17
            default:
                if len(v4) > 0 {
                    v4[0] = ""
                }
            }
            delete(v2, ("\n\t" + ((("hello" + "hello") + ("hello" + (("foo bar" + "a") + (("" + ("" + ((("a" + "a") + "foo bar") + "foo bar"))) + "\n\t")))) + (("\n\t" + "\n\t") + ""))))
            _ = fi13
        }
        var v18 float64 = v8
        _ = fi12
        _ = v18
    }
    switch "hello" {
    case "hello":
        v8 = float64(((float32((v8 / v8)) * 3.14159) - 3.14159))
    case "":
        var v19 int32 = 1
        _ = v19
    case "a":
        if true {
            fmt.Println(v5)
            switch "\n\t" {
            case "a":
                var v20 string = "\n\t"
                _ = v20
            case "":
                var v21 int64 = -94
                _ = v21
            case "hello":
                fmt.Println(v5, v3)
            default:
                v5 = (^(v5))
            }
            v22 := GenS1{F0: 128}
            _ = v22
        }
    default:
        fmt.Println(v8)
    }
    var v23 int32 = 709
    fmt.Println(v5, v5, v23)
    delete(v2, "a")
    v8 += v8
    fmt.Println(v23, v23)
    var v24 int64 = 0
    v25 := GenS1{F0: -1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
