package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int16 }

func main() {
    var v1 string = ""
    v2 := make(map[string]string)
    v3 := []string{"foo bar", (v1 + "foo bar")}
    v4 := make(map[string]string)
    var v5 uint32 = 2
    _ = len(v3)
    v5 += v5
    var v6 int64 = -9223372036854775808
    _ = len(v3)
    delete(v2, "hello")
    v7 := v2[((v1 + "") + ((v1 + v1) + v1))]
    v1 = v1
    _ = len(v3)
    fmt.Println(v5, v5)
    for fi8 := 0; fi8 < 12; fi8++ {
        if (2 < 42) {
            var v9 int32 = -1
            v9 = v9
            _ = v9
        } else {
            for fi10 := 0; fi10 < 20; fi10++ {
                if false {
                    v7 = "hello"
                    var v11 int8 = 1
                    _ = v11
                } else {
                    fmt.Println(v7)
                    fmt.Println(fi10)
                    v5 += (v5 / ((2147483648) | 1))
                }
                var v12 string = v1
                _ = fi10
                _ = v12
            }
        }
        _ = fi8
    }
    switch v1 {
    case "hello":
        if true {
            var v13 int = 1
            v13 = (-(v13))
            _ = v13
        } else {
            v6 -= (v6 | (v6 / ((1) | 1)))
            switch 245 {
            case 701:
                v14 := GenS1{F0: 440, F1: v6, F2: 0}
                _ = v14
            }
        }
    }
    var v15 int16 = 42
    _ = len(v2)
    for fi16 := 0; fi16 < 17; fi16++ {
        for fi17 := 0; fi17 < 10; fi17++ {
            switch ("" + ((v7 + v7) + (v1 + "\n\t"))) {
            case "foo bar":
                if false {
                    var v18 float64 = 0.0
                    _ = v18
                } else {
                    var v19 float64 = (0.1 / ((3.14159 / 1.0) - (0.0 + 1e10)))
                    if false {
                        v20 := GenS1{F0: 959, F1: v6, F2: 0}
                        _ = v20
                    }
                    _ = v19
                }
            case "\n\t":
                v6 -= v6
            case "hello":
                if (false || true) {
                    if (false || (true || ((!false) || false))) {
                        if len(v3) > 0 {
                            v3[0] = v7
                        }
                    }
                }
            default:
                var v21 int32 = int32(v6)
                _ = v21
            }
            var v22 string = v7
            fmt.Println(v6, v5, fi16)
            _ = fi17
            _ = v22
        }
        fmt.Println(fi16, v5, fi16)
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
