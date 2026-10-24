package main

import (
    "fmt"
)


func main() {
    var v1 int = -996
    var v2 int32 = 42
    var v3 int32 = (v2 / ((v2) | 1))
    v4 := make(map[uint32]int)
    var v5 int64 = 0
    var v6 uint16 = 180
    var v7 uint32 = 65535
    if false {
        _ = len(v4)
    }
    v8 := []int{}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = len(v8)
    var v9 float32 = float32(((-2.25 * 1e10) + float64((3.14159 * 3.14159))))
    var v10 uint32 = v7
    if (0 <= (v2 ^ v3)) {
        var v11 uint8 = 10
        var v12 uint64 = 18446744073709551615
        switch v1 {
        case -2147483648:
            for fi13 := 0; fi13 < 1; fi13++ {
                var v14 int = (v1 >> ((18446744073709551615) & 63))
                _ = fi13
                _ = v14
            }
        case 42:
            var v15 float64 = 3.14159
            _ = v15
        default:
            switch (v1 ^ 32767) {
            case 590:
                if (v1 > v1) {
                    v2 = v2
                    for fi16 := 0; fi16 < 1; fi16++ {
                        var v17 int32 = v3
                        _ = fi16
                        _ = v17
                    }
                }
            }
        }
        _ = v11
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
