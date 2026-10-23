package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{}
    var v3 int32 = 2147483647
    var v4 uint = 297
    var v5 float32 = float32(((((float64((1e10 * 3.14159)) + (3.14159 * 3.14159)) / (((3.14159 * (0.1 * (0.1 / (((1.5 / 0.1) - -2.25) - ((1e10 / (0.1 * 1.5)) / 1e10))))) * 0.1) / 1e10)) * 0.1) - 3.14159))
    if true {
        v5 += v5
        fmt.Println(v4, v4, v4)
        fmt.Println(v4, v5, v4)
    }
    switch ("foo bar" + "\n\t") {
    case "a":
        fmt.Println(v5, v5, v3)
    case "hello":
        var v6 int64 = 42
        _ = v6
    case "foo bar":
        v7 := GenS1{F0: 2, F1: (v4 | (v4 + 318)), F2: v3}
        _ = v7
    }
    var v8 int64 = 42
    var v9 uint16 = 255
    var v10 float32 = (v5 - ((1.5 / (float32((3.14159 / (-2.25 + 1e10))) - v5)) + 0.0))
    for fi11 := 0; fi11 < 4; fi11++ {
        if false {
            fmt.Println(v10, v4, v9)
            switch fi11 {
            case -85:
                var v12 int32 = 300
                _ = v12
            case 128:
                if true {
                    fmt.Println(v8, fi11)
                    v5 = 1.5
                }
            case 127:
                var v13 uint = v4
                _ = v13
            default:
                if (("hello" + (("a" + "foo bar") + "a")) < "a") {
                    var v14 uint64 = 1
                    for fi15 := 0; fi15 < 2; fi15++ {
                        var v16 string = "foo bar"
                        v17 := GenS1{F0: (v14 + (v14 % ((856) | 1))), F1: v4, F2: v3}
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    v14 = v14
                    _ = v14
                }
            }
        } else {
            v3 += (v3 / (((^(v3))) | 1))
            for fi18 := 0; fi18 < 5; fi18++ {
                v19 := v1[("foo bar" + ("foo bar" + ("" + "\n\t")))]
                _ = fi18
                _ = v19
            }
        }
        _ = fi11
    }
    if (("a" + "a") > "foo bar") {
        var v20 int = -2147483648
        var v21 int8 = 59
        fmt.Println(v9, v4)
        _ = v20
        _ = v21
    }
    var v22 bool = true
    _ = len(v1)
    var v23 uint = 255
    v10 += v5
    v4 -= (^(v23))
    fmt.Println(v10)
    if true {
        var v24 bool = v22
        var v25 float64 = 3.14159
        var v26 int64 = v8
        _ = v24
        _ = v25
        _ = v26
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v8 = v8
    var v27 int64 = v8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
