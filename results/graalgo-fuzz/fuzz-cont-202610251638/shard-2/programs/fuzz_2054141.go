package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }

func main() {
    var v1 uint = 42
    var v2 int8 = 15
    var v3 uint16 = 65535
    v4 := make(map[string]bool)
    var v5 int8 = 127
    var v6 bool = false
    if (3.14159 <= (0.1 * 1e10)) {
        for fi7 := 0; fi7 < 3; fi7++ {
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 bool = v6
                _ = fi8
                _ = v9
            }
            switch fi7 {
            case 32768:
                v3 = v3
            case 127:
                v10 := GenS1{F0: 1, F1: v3}
                _ = v10
            }
            v11 := v4["foo bar"]
            _ = fi7
            _ = v11
        }
    }
    v12 := []int{32767, 127, 127, -9223372036854775808}
    var v13 int = 399
    var v14 uint = (v1 - (v1 ^ v1))
    var v15 int32 = 1
    var v16 bool = (!v6)
    v3 = v3
    for fi17 := 0; fi17 < 10; fi17++ {
        var v18 float64 = 0.1
        _ = fi17
        _ = v18
    }
    delete(v4, ("" + ""))
    v16 = v6
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
