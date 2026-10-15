package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint64 }

func main() {
    v1 := []uint32{}
    var v2 int64 = -320
    v3 := make(map[string]bool)
    var v4 int64 = 965
    v5 := []uint32{}
    if ((0.1 / -2.25) < 0.1) {
        v4 += -393
        if true {
            fmt.Println(v2, v4)
        } else {
            if ((1 < 0) && (265 != 1)) {
                var v6 uint8 = 128
                _ = v6
            }
        }
    }
    _ = len(v1)
    _ = len(v5)
    switch 6 {
    case 0:
        if (!(!(-1 != 2147483647))) {
            fmt.Println(v4, v2)
        }
    default:
        if (42 == 2147483648) {
            if false {
                fmt.Println(v2, v2)
            }
        }
    }
    v7 := v3[""]
    v7 = (100 >= 31)
    v2 = v2
    if (!(!true)) {
        v3["\n\t"] = (((3.14159 * float32(0.0)) != -2.25) && true)
        v8 := GenS1{F0: "", F1: 2, F2: 0}
        _ = v8
    }
    var v9 int32 = 42
    for fi10 := 0; fi10 < 18; fi10++ {
        for fi11 := 0; fi11 < 1; fi11++ {
            v7 = v7
            v12 := GenS1{F0: ("\n\t" + "\n\t"), F1: 107, F2: 118}
            var v13 float32 = (((((1.5 - 0.1) + 0.1) * 3.14159) * -2.25) / 0.1)
            fmt.Println(v2, v13, v4)
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
