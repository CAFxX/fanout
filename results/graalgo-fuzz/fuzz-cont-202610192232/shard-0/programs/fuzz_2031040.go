package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }

func main() {
    var v1 uint = 559
    v2 := []uint32{2, 123, 65535}
    var v3 float64 = (0.0 / (float64(1e10) + 0.1))
    v4 := GenS1{F0: 126, F1: ("foo bar" + ("a" + "a"))}
    var v5 int8 = 75
    switch "hello" {
    case "foo bar":
        if (0 == 2147483647) {
            var v6 uint32 = 41
            fmt.Println(v3)
            if len(v2) > 0 {
                v2[0] = v6
            }
            _ = v6
        }
    case "":
        var v7 float64 = (1e10 / float64(0.1))
        _ = v7
    default:
        if len(v2) > 0 {
            v2[0] = 659
        }
    }
    v8 := GenS1{F0: 4294967295, F1: "a"}
    _ = len(v2)
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 string = ("a" + "")
        if true {
            var v11 uint8 = 128
            v3 += ((((v3 - ((-2.25 + v3) + (v3 + v3))) + float64(1.5)) + (v3 * v3)) / 1e10)
            _ = v11
        } else {
            for fi12 := 0; fi12 < 14; fi12++ {
                var v13 uint32 = 4294967295
                v14 := GenS1{F0: v13, F1: "foo bar"}
                var v15 int8 = (v5 | v5)
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
            v1 = 4294967295
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(v3)
    v1 = 861
    v2 = append(v2, 256)
    v16 := GenS1{F0: 956, F1: ""}
    v1 = 560
    v1 -= 871
    var v17 uint = (v1 - (v1 | (^(v1))))
    v18 := GenS1{F0: 602, F1: ("hello" + "\n\t")}
    v3 = v3
    v19 := GenS1{F0: 909, F1: ""}
    fmt.Println(v1, v5)
    v20 := GenS1{F0: 0, F1: ("a" + "\n\t")}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
