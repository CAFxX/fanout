package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 string }

func main() {
    var v1 int = 32767
    var v2 int64 = 9223372036854775807
    var v3 int8 = 104
    v4 := make(map[uint32]string)
    var v5 float64 = (((1.5 - (0.0 * (3.14159 / 1e10))) / 1.0) / (((1.5 - (0.0 - (float64(0.0) * 1e10))) * 1.5) * 1.5))
    var v6 float64 = (v5 * v5)
    v1 = (v1 | v1)
    v7 := GenS1{F0: 42, F1: 835, F2: "foo bar"}
    if (v5 != (-2.25 + float64((float32(v6) - (0.1 * 3.14159))))) {
        var v8 uint64 = 18446744073709551615
        var v9 int32 = -1
        _ = v8
        _ = v9
    } else {
        switch "hello" {
        case "":
            var v10 string = ""
            _ = v10
        default:
            fmt.Println(v6, v6, v5)
        }
    }
    v2 -= (v2 % ((v2) | 1))
    var v11 int = 9223372036854775807
    v3 = v3
    switch (v1 | (^(v1))) {
    case -2147483648:
        v4[65535] = "hello"
    }
    v5 -= (-2.25 - 3.14159)
    v2 = (v2 / (((v2 << ((277) & 63))) | 1))
    switch "hello" {
    case "a":
        switch -624 {
        case -958:
            switch ("\n\t" + "a") {
            case "":
                if ((v3 * (v3 % ((-116) | 1))) < 1) {
                    var v12 int16 = -32768
                    var v13 bool = true
                    v5 += (v5 - (v5 + (v6 - (((v5 - float64(0.1)) - (v6 - v5)) * ((float64(1.5) * v6) - float64((1e10 + (0.0 / 1e10))))))))
                    _ = v12
                    _ = v13
                }
            case "foo bar":
                var v14 int64 = v2
                _ = v14
            case "a":
                switch (("\n\t" + "hello") + "") {
                case "hello":
                    fmt.Println(v6, v1)
                case "foo bar":
                    v4[1] = "\n\t"
                default:
                    var v15 int8 = v3
                    _ = v15
                }
            default:
                for fi16 := 0; fi16 < 5; fi16++ {
                    v2 -= v2
                    _ = len(v4)
                    _ = fi16
                }
            }
        }
    case "foo bar":
        var v17 uint8 = 2
        _ = v17
    case "hello":
        fmt.Println(v5, v1, v6)
    }
    v18 := v4[1]
    v19 := []int{v11, -1, v1, (-(v11))}
    var v20 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
