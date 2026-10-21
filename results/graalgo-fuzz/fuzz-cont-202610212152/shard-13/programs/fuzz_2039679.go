package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int, p1 int64, p2 int) uint {
    v4 := GenS1{F0: p2}
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 2
}

func main() {
    v1 := []uint32{2, 640}
    var v2 uint8 = 255
    var v3 float64 = 1e10
    var v5 int8 = -1
    v5 += (v5 & (v5 & v5))
    var v6 float64 = (1e10 - (v3 + (float64(0.1) / v3)))
    switch ("" + "") {
    case "":
        switch -9223372036854775808 {
        case 660:
            var v7 uint16 = 0
            _ = v7
        case 127:
            var v8 int16 = 32767
            _ = v8
        case -2147483648:
            v5 = (v5 & 1)
        }
    default:
        for fi9 := 0; fi9 < 8; fi9++ {
            v6 = v6
            for fi10 := 0; fi10 < 20; fi10++ {
                var v11 uint8 = (^(v2))
                if len(v1) > 0 {
                    v1[0] = 65535
                }
                var v12 uint = 4294967295
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi9
        }
    }
    v5 = v5
    var v13 int8 = 100
    v1 = append(v1, 2147483648)
    var v14 uint64 = 1
    fmt.Println(v6, v14, v13)
    v13 -= 0
    if len(v1) > 0 {
        v1[0] = 0
    }
    var v15 float32 = 1.5
    var v16 int16 = -32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
