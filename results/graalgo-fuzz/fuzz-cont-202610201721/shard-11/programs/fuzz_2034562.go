package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint16 }

func main() {
    var v1 float64 = 1.5
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    v4 := make(map[int]bool)
    var v5 bool = (1 <= 0)
    v6 := []string{}
    for fi7 := 0; fi7 < 18; fi7++ {
        v5 = v5
        v8 := GenS1{F0: v5}
        _ = fi7
        _ = v8
    }
    if (!(-633 == -1)) {
        var v9 int8 = -1
        switch -2147483648 {
        case 127:
            var v10 float64 = v1
            _ = v10
        case -34:
            var v11 uint = 18446744073709551615
            _ = v11
        }
        _ = v9
    } else {
        var v12 uint8 = 245
        _ = v12
    }
    if (666 != 910) {
        v13 := GenS2{F0: "", F1: 42}
        for fi14 := 0; fi14 < 1; fi14++ {
            v2[436] = "hello"
            _ = len(v6)
            v5 = (363 <= 255)
            _ = fi14
        }
        _ = v13
    }
    var v15 float32 = 1e10
    var v16 uint32 = 526
    var v17 float64 = (1e10 * -2.25)
    var v18 float64 = ((1.5 - 0.0) * ((0.1 + 0.0) - ((0.0 + float64(0.1)) - 0.0)))
    for fi19 := 0; fi19 < 7; fi19++ {
        switch "" {
        case "a":
            v17 = v17
        }
        fmt.Println(v16, v15, v17)
        _ = fi19
    }
    var v20 uint16 = 42
    fmt.Println(v20)
    v15 = -2.25
    if len(v6) > 0 {
        v6[0] = "hello"
    }
    if v5 {
        switch "" {
        case "":
            switch ("" + (("foo bar" + "hello") + ("\n\t" + "\n\t"))) {
            case "\n\t":
                if (float32(1.5) == (v15 - ((0.1 - (1.5 * v15)) * v15))) {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
            case "a":
                var v21 bool = v5
                _ = v21
            case "":
                v22 := v3[-1]
                _ = v22
            }
        case "hello":
            var v23 bool = v5
            _ = v23
        case "a":
            v24 := GenS1{F0: true}
            _ = v24
        }
    } else {
        v15 += v15
    }
    var v25 bool = v5
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
