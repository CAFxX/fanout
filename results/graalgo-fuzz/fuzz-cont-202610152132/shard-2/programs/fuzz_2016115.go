package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := []string{"", "foo bar", "\n\t", ("" + ("a" + ("\n\t" + ("foo bar" + "a"))))}
    var v2 int = -1
    v3 := make(map[string]string)
    v4 := make(map[uint32]string)
    var v5 int64 = 998
    if true {
        fmt.Println(v2, v5)
        for fi6 := 0; fi6 < 10; fi6++ {
            switch 2147483647 {
            case -264:
                for fi7 := 0; fi7 < 12; fi7++ {
                    var v8 int8 = 0
                    if (856 > 397) {
                        v2 += -9223372036854775808
                    }
                    _ = fi7
                    _ = v8
                }
            case -2147483648:
                for fi9 := 0; fi9 < 7; fi9++ {
                    v10 := GenS1{F0: (v5 % ((v5) | 1))}
                    if true {
                        var v11 uint16 = 93
                        _ = v11
                    } else {
                        v12 := v3[("\n\t" + ("foo bar" + ""))]
                        var v13 int8 = 0
                        _ = v12
                        _ = v13
                    }
                    _ = fi9
                    _ = v10
                }
            case 451:
                v2 -= -832
            }
            _ = fi6
        }
    }
    v5 += (v5 << ((1) & 63))
    switch -1 {
    case 32767:
        v5 = (v5 + (v5 & (-(v5))))
    case 42:
        for fi14 := 0; fi14 < 17; fi14++ {
            v5 -= v5
            var v15 float64 = 0.0
            v2 = v2
            _ = fi14
            _ = v15
        }
    }
    delete(v4, 994)
    v5 += (^(v5))
    var v16 string = "\n\t"
    if (true || false) {
        _ = len(v1)
        var v17 float32 = 3.14159
        _ = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
