package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint32, p1 int16, p2 bool) uint64 {
    var v6 int64 = 1
    if (674 < 9223372036854775807) {
        if (0 >= 40) {
            v6 += v6
            v6 -= 0
            var v7 uint8 = 2
            _ = v7
        } else {
            for fi8 := 0; fi8 < 12; fi8++ {
                for fi9 := 0; fi9 < 11; fi9++ {
                    v10 := GenS1{F0: (^(p0))}
                    _ = fi9
                    _ = v10
                }
                _ = fi8
            }
            v11 := []int{39, 436, 2147483647}
            if (!false) {
                p2 = false
            } else {
                for fi12 := 0; fi12 < 19; fi12++ {
                    var v13 int16 = (^(p1))
                    v14 := GenS2{F0: 0}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                p1 = (p1 ^ 0)
                var v15 int16 = (-(p1))
                _ = v15
            }
            _ = v11
        }
        p0 += p0
    } else {
        var v16 int8 = 104
        if (1e10 > 0.1) {
            v17 := []int{int(p0), 2147483647}
            var v18 int16 = (p1 ^ 1000)
            _ = v17
            _ = v18
        } else {
            var v19 uint16 = 0
            _ = v19
        }
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}

func main() {
    var v1 uint32 = 65535
    var v2 int16 = 0
    v3 := make(map[string]uint32)
    v4 := []int{1, -9223372036854775808}
    var v5 uint8 = 2
    v1 = v1
    v2 += (^(v2))
    var v20 int32 = 1
    v1 = v1
    if len(v4) > 0 {
        v4[0] = 9223372036854775807
    }
    if ((18446744073709551615 != 255) && ((0.1 != 1e10) && true)) {
        var v21 int32 = v20
        _ = v21
    } else {
        v5 = 2
    }
    switch 32768 {
    case 89:
        v5 += v5
    case 2147483647:
        v20 -= v20
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v5 = v5
    var v22 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
