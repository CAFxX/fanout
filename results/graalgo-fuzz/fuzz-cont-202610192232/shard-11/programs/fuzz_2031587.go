package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
func genFn0(p0 int16, p1 uint) int32 {
    p1 += p1
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(p1, p0, p1)
        var v8 uint32 = 564
        _ = fi7
        _ = v8
    }
    p1 = (^(p1))
    fmt.Println(p1, p1)
    var v9 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = v9
    return 1
}

func main() {
    var v1 uint = 4294967295
    v2 := []string{}
    v3 := make(map[int]string)
    v4 := make(map[int]int)
    var v5 float64 = 0.1
    v6 := make(map[string]bool)
    switch -331 {
    case 503:
        var v10 uint16 = 255
        _ = v10
    case -9223372036854775808:
        v11 := GenS1{F0: 127, F1: 130}
        _ = v11
    case 654:
        v12 := GenS1{F0: 127, F1: 42}
        _ = v12
    }
    switch ("hello" + ((((("\n\t" + ("" + "\n\t")) + (("a" + "foo bar") + (("a" + "a") + ("" + (("\n\t" + "") + "foo bar"))))) + "") + "foo bar") + ("a" + "a"))) {
    case "\n\t":
        if (65535 >= 2147483648) {
            v1 = (^(v1))
            var v13 int = -548
            _ = v13
        }
    case "a":
        switch (("a" + ("hello" + "a")) + ("foo bar" + "hello")) {
        case "\n\t":
            switch "foo bar" {
            case "a":
                var v14 float32 = (3.14159 * float32((((1.5 - (float64(((((1.5 - 0.1) - 0.1) / -2.25) + (1e10 + 0.1))) - 0.0)) / (float64(0.0) + v5)) + v5)))
                _ = v14
            case "foo bar":
                v4[32768] = 2147483647
            case "":
                _ = len(v6)
            }
        case "foo bar":
            fmt.Println(v5, v1)
        default:
            for fi15 := 0; fi15 < 4; fi15++ {
                fmt.Println(fi15, v1, fi15)
                if (3.14159 <= ((((1e10 + -2.25) + float32((v5 / float64((3.14159 + 0.1))))) * 3.14159) + 1e10)) {
                    var v16 bool = (42 >= 255)
                    v1 -= v1
                    _ = v16
                }
                v4[fi15] = (fi15 & fi15)
                _ = fi15
            }
        }
    case "":
        var v17 string = "\n\t"
        _ = v17
    default:
        fmt.Println(v1, v5)
    }
    v18 := GenS1{F0: 88, F1: 2147483648}
    for fi19 := 0; fi19 < 13; fi19++ {
        delete(v3, (fi19 % ((fi19) | 1)))
        v5 = (v5 - float64(1.5))
        fmt.Println(v1, fi19, fi19)
        _ = fi19
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        var v21 uint = (^(v1))
        _ = fi20
        _ = v21
    }
    v22 := GenS1{F0: 2, F1: 42}
    v23 := v3[2147483647]
    v24 := GenS1{F0: 128, F1: 957}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
