package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint }

func main() {
    var v1 int8 = 7
    v2 := make(map[string]string)
    var v3 int = 32767
    var v4 int32 = -578
    var v5 uint16 = 351
    var v6 uint32 = 2147483647
    switch (v3 % ((127) | 1)) {
    case -232:
        delete(v2, "")
    }
    fmt.Println(v6)
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 int16 = -32768
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        _ = len(v2)
        v10 := v2["hello"]
        _ = fi9
        _ = v10
    }
    fmt.Println(v4, v1, v3)
    if (32767 >= 1) {
        var v11 float64 = (1.5 / 1.5)
        if (810 <= v3) {
            _ = len(v2)
            v12 := GenS1{F0: -797, F1: -365, F2: 18446744073709551615}
            _ = v12
        } else {
            v5 += (v5 + v5)
            switch "\n\t" {
            case "hello":
                v2[""] = ""
            case "\n\t":
                if (v3 >= 127) {
                    var v13 int32 = v4
                    for fi14 := 0; fi14 < 17; fi14++ {
                        if false {
                            v3 = (-(v3))
                        }
                        v11 = v11
                        _ = fi14
                    }
                    v4 -= v13
                    _ = v13
                }
            case "":
                v3 -= v3
            }
        }
        var v15 string = "foo bar"
        _ = v11
        _ = v15
    }
    var v16 string = ""
    var v17 int = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
