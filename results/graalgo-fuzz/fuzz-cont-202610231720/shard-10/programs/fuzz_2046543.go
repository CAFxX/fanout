package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS2 struct { F0 bool; F1 bool; F2 uint32 }

func main() {
    var v1 int64 = 131
    var v2 uint16 = 32767
    var v3 bool = false
    v4 := make(map[string]int)
    v5 := make(map[uint32]uint32)
    v6 := make(map[uint32]int)
    v7 := GenS1{F0: 1, F1: (^(v2)), F2: uint32(v2)}
    var v8 bool = (42 == 2)
    for fi9 := 0; fi9 < 16; fi9++ {
        v3 = false
        _ = fi9
    }
    v10 := []int{2147483647}
    v8 = (true || (2147483647 > 42))
    var v11 int64 = (v1 >> ((42) & 63))
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    if ((231 < 0) || (v3 && ("\n\t" > ""))) {
        switch "hello" {
        case "a":
            var v12 float64 = 1.5
            _ = v12
        }
        var v13 uint32 = 2147483647
        v13 = (v13 / ((v13) | 1))
        _ = v13
    } else {
        v14 := GenS2{F0: (!v3), F1: (18446744073709551615 != 228), F2: 65535}
        var v15 string = "\n\t"
        for fi16 := 0; fi16 < 2; fi16++ {
            var v17 float32 = -2.25
            for fi18 := 0; fi18 < 11; fi18++ {
                v5[926] = 2147483648
                _ = fi18
            }
            if v3 {
                var v19 uint8 = 222
                _ = v19
            }
            _ = fi16
            _ = v17
        }
        _ = v14
        _ = v15
    }
    v1 = (v1 & v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
