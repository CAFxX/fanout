package main

import (
    "fmt"
)


func main() {
    v1 := []int{0, 0, -665, 42}
    var v2 int = 32768
    var v3 int32 = 1
    var v4 uint64 = 4294967295
    v5 := []uint32{2147483647}
    var v6 uint32 = 65535
    _ = len(v1)
    for fi7 := 0; fi7 < 10; fi7++ {
        fmt.Println(v2)
        for fi8 := 0; fi8 < 8; fi8++ {
            switch fi8 {
            case -940:
                if (false || true) {
                    fmt.Println(v3, v4)
                    v2 = fi7
                } else {
                    var v9 string = ""
                    v6 = (v6 - v6)
                    _ = v9
                }
            case 429:
                var v10 bool = (fi7 == (-(fi8)))
                _ = v10
            case 127:
                _ = len(v5)
            }
            switch v2 {
            case -2147483648:
                v1 = append(v1, (fi7 & (fi7 >> ((4294967295) & 63))))
            default:
                fmt.Println(v4)
            }
            _ = fi8
        }
        for fi11 := 0; fi11 < 20; fi11++ {
            fmt.Println(v6, fi11)
            _ = fi11
        }
        _ = fi7
    }
    fmt.Println(v3, v3)
    var v12 int8 = -78
    fmt.Println(v12, v12, v12)
    var v13 uint8 = 1
    v3 -= v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
