package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int32 = 42
    var v2 uint64 = 0
    var v3 uint = 370
    var v4 int16 = 1000
    v5 := []int{2147483647, 32767, 42, -357}
    v6 := GenS1{F0: 0}
    fmt.Println(v3)
    fmt.Println(v1)
    for fi7 := 0; fi7 < 17; fi7++ {
        v2 = v2
        for fi8 := 0; fi8 < 3; fi8++ {
            v2 = v2
            var v9 uint32 = 42
            var v10 uint32 = 2147483647
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 int = fi7
        _ = fi7
        _ = v11
    }
    fmt.Println(v4, v1, v2)
    if ((v1 * v1) < (v1 & v1)) {
        v2 -= 0
    } else {
        if true {
            v3 = v3
        } else {
            v5 = append(v5, -725)
            if true {
                if len(v5) > 0 {
                    v5[0] = 9223372036854775807
                }
                v5 = append(v5, 32768)
                var v12 int8 = 123
                _ = v12
            }
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 uint8 = uint8(v2)
                _ = fi13
                _ = v14
            }
        }
    }
    v15 := GenS1{F0: -9223372036854775808}
    v16 := GenS1{F0: 0}
    if (!false) {
        fmt.Println(v1, v1)
    } else {
        var v17 uint64 = v2
        _ = v17
    }
    if (!true) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
