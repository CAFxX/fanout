package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int8 }

func main() {
    var v1 int8 = 4
    var v2 bool = (!((-219 <= 1) && false))
    v3 := make(map[uint32]int)
    v1 -= (v1 * v1)
    v4 := GenS1{F0: 1000, F1: 128, F2: v1}
    v2 = (!((((0.0 / 1.5) / 1.0) * 0.0) >= ((1.5 / (1.5 / 3.14159)) / 1.5)))
    var v5 string = "hello"
    v5 = v5
    v1 += 100
    v6 := GenS1{F0: 65535, F1: 128, F2: -1}
    switch 2147483647 {
    case 4294967295:
        v2 = (0 != 921)
    default:
        switch (("" + v5) + "foo bar") {
        case "foo bar":
            if (0 <= 1) {
                switch 4294967295 {
                case 319:
                    fmt.Println(v5, v2, v2)
                case 1:
                    fmt.Println(v5, v2, v2)
                default:
                    var v7 uint64 = 42
                    _ = v7
                }
                fmt.Println(v2)
            } else {
                for fi8 := 0; fi8 < 20; fi8++ {
                    for fi9 := 0; fi9 < 4; fi9++ {
                        var v10 float32 = 3.14159
                        _ = fi9
                        _ = v10
                    }
                    _ = fi8
                }
            }
        case "":
            var v11 uint16 = 1
            _ = v11
        default:
            _ = len(v3)
        }
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        switch v5 {
        case "hello":
            var v13 int = fi12
            _ = v13
        case "foo bar":
            switch (fi12 >> ((65535) & 63)) {
            case 32768:
                v1 += 42
            }
        case "\n\t":
            switch fi12 {
            case -431:
                if (("\n\t" + v5) == "a") {
                    var v14 uint32 = 42
                    for fi15 := 0; fi15 < 18; fi15++ {
                        var v16 uint32 = v14
                        var v17 uint8 = 128
                        var v18 int32 = 42
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    _ = v14
                }
            case 872:
                for fi19 := 0; fi19 < 1; fi19++ {
                    v5 = ("" + v5)
                    _ = fi19
                }
            }
        default:
            v20 := GenS1{F0: 0, F1: 1, F2: 1}
            _ = v20
        }
        _ = fi12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
