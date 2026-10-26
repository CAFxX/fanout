package main

import (
    "fmt"
)

func genFn0(p0 uint8, p1 uint64, p2 bool) uint32 {
    p1 += (p1 / ((p1) | 1))
    p0 = 1
    v4 := []int{9223372036854775807, 128, 2147483647}
    if len(v4) > 0 {
        v4[0] = 127
    }
    if p2 {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 65535
}

func main() {
    var v1 uint = 0
    v2 := []uint32{}
    var v3 uint32 = 65535
    if (!(((1.5 * 0.1) / -2.25) < (3.14159 / -2.25))) {
        var v5 int32 = -1
        var v6 float32 = (1e10 - (0.0 - 0.0))
        if len(v2) > 0 {
            v2[0] = v3
        }
        _ = v5
        _ = v6
    }
    _ = len(v2)
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 bool = ((fi8 / ((fi8) | 1)) <= fi8)
            v9 = true
            var v10 uint8 = 168
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    var v11 int16 = -1
    for fi12 := 0; fi12 < 11; fi12++ {
        switch "a" {
        case "a":
            for fi13 := 0; fi13 < 19; fi13++ {
                v3 += (v3 << (((v1 * v1)) & 63))
                var v14 bool = (false && (true || true))
                _ = fi13
                _ = v14
            }
        default:
            var v15 int = -9223372036854775808
            _ = v15
        }
        for fi16 := 0; fi16 < 12; fi16++ {
            v3 = v3
            _ = fi16
        }
        if (1e10 >= -2.25) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v17 uint8 = 2
            if len(v2) > 0 {
                v2[0] = 713
            }
            _ = v17
        }
        _ = fi12
    }
    if (!false) {
        v3 -= v3
    } else {
        var v18 uint16 = 711
        var v19 int64 = -588
        _ = v18
        _ = v19
    }
    v3 -= 1
    var v20 string = "hello"
    var v21 int16 = (v11 & v11)
    v20 = v20
    for fi22 := 0; fi22 < 4; fi22++ {
        var v23 uint = 18446744073709551615
        switch v20 {
        case "foo bar":
            var v24 int = (fi22 + (fi22 * fi22))
            _ = v24
        }
        _ = fi22
        _ = v23
    }
    var v25 uint = 4294967295
    v11 -= (v21 | v21)
    v2 = append(v2, (v3 * genFn0(77, uint64(v1), ((^(v1)) == (v25 - (^(v25)))))))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
}
