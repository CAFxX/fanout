package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint32 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]int)
    var v3 int16 = 32767
    var v4 uint8 = 1
    var v5 uint8 = (v4 >> ((18446744073709551615) & 63))
    var v6 int8 = 0
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint32 = 157
        v9 := []int{(fi7 / ((793) | 1)), 4294967295}
        for fi10 := 0; fi10 < 1; fi10++ {
            fmt.Println(v3)
            var v11 int32 = 2147483647
            v12 := GenS1{F0: 427, F1: (v6 ^ 127)}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    delete(v2, "\n\t")
    v4 = (^(v5))
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := v1["\n\t"]
        for fi15 := 0; fi15 < 1; fi15++ {
            v16 := GenS1{F0: 1, F1: (v6 << ((555) & 63))}
            v17 := GenS2{F0: (^(v6)), F1: 622, F2: 0}
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 int32 = -482
        _ = fi13
        _ = v14
        _ = v18
    }
    if false {
        v19 := GenS2{F0: (v6 | v6), F1: 42, F2: 2}
        if ((!true) && (-1 > -1)) {
            var v20 uint = 977
            var v21 uint = (v20 / ((775) | 1))
            v5 = 128
            _ = v20
            _ = v21
        }
        _ = v19
    }
    v2["foo bar"] = 32768
    if (1 < 0) {
        for fi22 := 0; fi22 < 10; fi22++ {
            v23 := GenS1{F0: 0, F1: (v6 + v6)}
            _ = fi22
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        v25 := []int{}
        var v26 uint16 = 325
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
