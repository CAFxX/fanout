package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 int64 }
func genFn0(p0 uint64, p1 int16, p2 int) int32 {
    p1 -= (p1 / ((p1) | 1))
    v7 := GenS1{F0: (-540 >= 332)}
    switch (p2 & (p2 - (p2 >> ((255) & 63)))) {
    case 236:
        var v8 uint = 42
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -2147483648
}

func main() {
    v1 := []string{"a", ""}
    var v2 float32 = float32(float64((float32(0.1) + 1e10)))
    var v3 int16 = 559
    var v4 uint64 = 0
    var v5 int16 = v3
    var v6 bool = false
    if (428 <= v5) {
        var v9 int8 = 100
        _ = v9
    } else {
        if len(v1) > 0 {
            v1[0] = ("" + "hello")
        }
    }
    v10 := GenS2{F0: 65535, F1: -9223372036854775808}
    v5 -= (v3 * v5)
    v11 := GenS1{F0: (-485 >= 228)}
    v12 := GenS2{F0: 483, F1: 0}
    if (0 < 127) {
        for fi13 := 0; fi13 < 17; fi13++ {
            if len(v1) > 0 {
                v1[0] = "a"
            }
            var v14 float64 = (-2.25 * ((-2.25 * 1e10) / (1e10 - (0.0 * 3.14159))))
            _ = fi13
            _ = v14
        }
        v1 = append(v1, "foo bar")
        if (true || v6) {
            v2 += v2
            fmt.Println(v5, v3)
            fmt.Println(v3, v5, v2)
        } else {
            for fi15 := 0; fi15 < 9; fi15++ {
                v3 = (v5 / ((v3) | 1))
                _ = fi15
            }
            v16 := GenS2{F0: 223, F1: 9223372036854775807}
            switch "\n\t" {
            case "\n\t":
                fmt.Println(v4, v2)
            }
            _ = v16
        }
    }
    v17 := GenS1{F0: v6}
    switch (("a" + (("\n\t" + "hello") + "foo bar")) + "") {
    case "":
        for fi18 := 0; fi18 < 16; fi18++ {
            if v6 {
                v19 := GenS1{F0: (848 > 2)}
                if true {
                    var v20 int64 = 0
                    var v21 int = (fi18 + fi18)
                    _ = v20
                    _ = v21
                } else {
                    var v22 float32 = 1e10
                    _ = v22
                }
                var v23 int64 = 1
                _ = v19
                _ = v23
            } else {
                v24 := GenS1{F0: v6}
                _ = v24
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi18
        }
    case "\n\t":
        for fi25 := 0; fi25 < 6; fi25++ {
            var v26 int16 = v3
            _ = fi25
            _ = v26
        }
    default:
        fmt.Println(v4)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
