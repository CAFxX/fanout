package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }

func main() {
    var v1 uint16 = 32767
    var v2 uint16 = 2
    v3 := make(map[uint32]uint32)
    if (!true) {
        v4 := GenS1{F0: 1, F1: -454}
        _ = v4
    }
    switch 2147483647 {
    case -2147483648:
        for fi5 := 0; fi5 < 5; fi5++ {
            fmt.Println(v2, v2, v1)
            _ = fi5
        }
    default:
        v1 = 2
    }
    v2 += 2
    fmt.Println(v1, v1)
    var v6 bool = (true || (0.1 < ((((((1.5 + 0.0) + 0.0) * float64(1e10)) * (-2.25 - 3.14159)) - 1.5) * (3.14159 / ((-2.25 + -2.25) * 1.5)))))
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := GenS1{F0: 2, F1: -932}
            var v10 uint64 = 682
            _ = fi8
            _ = v10
            _ = v9
        }
        switch (("" + "foo bar") + "foo bar") {
        case "\n\t":
            var v11 int8 = -94
            _ = v11
        case "hello":
            v2 = 42
        case "":
            var v12 uint16 = v2
            _ = v12
        default:
            for fi13 := 0; fi13 < 7; fi13++ {
                var v14 uint16 = (v2 - v1)
                if v6 {
                    v6 = v6
                    v15 := v3[0]
                    _ = v15
                } else {
                    v1 -= 815
                }
                var v16 float32 = 0.1
                v17 := []int{}
                _ = fi13
                _ = v14
                _ = v16
                _ = v17
            }
        }
        var v18 uint8 = 128
        _ = fi7
        _ = v18
    }
    v1 -= v2
    for fi19 := 0; fi19 < 4; fi19++ {
        v6 = (v6 || (1 > 32767))
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
}
