package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int8 }
type GenS2 struct { F0 int32; F1 int64; F2 int16 }

func main() {
    var v1 uint = 1
    v2 := make(map[int]int)
    v3 := []string{("" + "\n\t")}
    if (42 >= 85) {
        v4 := GenS1{F0: true, F1: 9223372036854775807, F2: -127}
        _ = v4
    } else {
        v2[-415] = 1
    }
    v5 := GenS1{F0: true, F1: 895, F2: 127}
    fmt.Println(v1)
    if true {
        var v6 uint8 = 2
        if ((!(145 >= -1)) || true) {
            var v7 uint32 = 743
            switch -635 {
            case 186:
                switch ("" + "hello") {
                case "a":
                    for fi8 := 0; fi8 < 14; fi8++ {
                        v9 := GenS1{F0: ((1.5 * ((1e10 + (0.0 - 1e10)) / 1.0)) <= float64(3.14159)), F1: -9223372036854775808, F2: 72}
                        v6 = (v6 & 128)
                        v10 := GenS1{F0: (0 <= 365), F1: 9223372036854775807, F2: 42}
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                case "\n\t":
                    _ = len(v3)
                }
            }
            _ = v7
        }
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        _ = v6
    } else {
        v1 = (v1 - (v1 / ((v1) | 1)))
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        fmt.Println(v1, fi11)
        var v12 float64 = ((0.0 + (3.14159 - (float64(3.14159) * 1.5))) + -2.25)
        _ = fi11
        _ = v12
    }
    v1 = v1
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 int = 4294967295
        fmt.Println(v14)
        _ = fi13
        _ = v14
    }
    v15 := GenS2{F0: -1, F1: -1, F2: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v5
}
