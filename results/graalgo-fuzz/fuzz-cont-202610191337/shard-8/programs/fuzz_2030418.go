package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = -292
    var v3 int64 = 42
    var v4 uint = 2
    v5 := []string{"a", "\n\t", ""}
    v6 := make(map[uint32]uint32)
    switch 9223372036854775807 {
    case 32768:
        for fi7 := 0; fi7 < 13; fi7++ {
            if true {
                if false {
                    for fi8 := 0; fi8 < 16; fi8++ {
                        v9 := GenS1{F0: 1, F1: "a"}
                        _ = fi8
                        _ = v9
                    }
                    var v10 uint = v4
                    _ = v10
                } else {
                    _ = len(v5)
                    var v11 bool = (671 == 1)
                    _ = v11
                }
            }
            _ = fi7
        }
    case 0:
        for fi12 := 0; fi12 < 2; fi12++ {
            fmt.Println(fi12, v4, fi12)
            _ = fi12
        }
    case 346:
        switch "\n\t" {
        case "\n\t":
            switch (("" + ("\n\t" + "")) + "foo bar") {
            case "foo bar":
                v4 = v4
            default:
                var v13 float64 = 0.0
                _ = v13
            }
        case "foo bar":
            var v14 uint32 = 0
            _ = v14
        case "a":
            v15 := GenS1{F0: 100, F1: ((("\n\t" + "hello") + "foo bar") + "hello")}
            _ = v15
        default:
            if (false && true) {
                v16 := GenS1{F0: int8(v2), F1: "foo bar"}
                _ = v16
            }
        }
    default:
        v17 := GenS1{F0: -75, F1: "a"}
        _ = v17
    }
    v18 := GenS1{F0: 100, F1: ("\n\t" + "\n\t")}
    var v19 float64 = 3.14159
    v19 += -2.25
    v2 -= v2
    var v20 int16 = 32767
    fmt.Println(v19, v4, v4)
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
