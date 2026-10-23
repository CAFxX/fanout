package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int16 }
func genFn0() int32 {
    var v5 uint = 38
    var v6 float64 = (float64((float32(1.5) + float32(1.5))) - 0.1)
    v6 -= v6
    _ = v5
    _ = v6
    return 6
}

func main() {
    var v1 int64 = -1
    var v2 int = -438
    v3 := make(map[uint32]uint32)
    var v4 float32 = 3.14159
    v7 := GenS1{F0: 1, F1: 0, F2: -780}
    v8 := v3[2147483648]
    if (false || true) {
        v2 -= (v2 & v2)
    } else {
        var v9 uint8 = 42
        v8 -= 65535
        _ = v9
    }
    v10 := []int{v2, -2147483648}
    v8 -= v8
    fmt.Println(v8, v2)
    var v11 int = (v2 | v2)
    v11 -= v2
    v12 := v3[(v8 | (^(v8)))]
    v4 += (v4 * (3.14159 * 3.14159))
    for fi13 := 0; fi13 < 3; fi13++ {
        if false {
            var v14 int16 = -1
            for fi15 := 0; fi15 < 19; fi15++ {
                switch (^(fi15)) {
                case -1:
                    v16 := GenS1{F0: 0, F1: 1, F2: -835}
                    _ = v16
                case -700:
                    var v17 string = "foo bar"
                    _ = v17
                }
                _ = fi15
            }
            for fi18 := 0; fi18 < 11; fi18++ {
                v19 := GenS1{F0: 4294967295, F1: 1, F2: 42}
                v20 := GenS1{F0: 4294967295, F1: 79, F2: v14}
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(fi13, v11)
                case "\n\t":
                    v12 += v12
                default:
                    var v21 uint64 = 2
                    _ = v21
                }
                var v22 bool = true
                _ = fi18
                _ = v19
                _ = v20
                _ = v22
            }
            _ = v14
        }
        _ = fi13
    }
    var v23 bool = ("" <= "foo bar")
    if (18446744073709551615 >= 686) {
        var v24 uint16 = 1
        v4 -= v4
        _ = v24
    }
    var v25 bool = true
    for fi26 := 0; fi26 < 16; fi26++ {
        if v23 {
            var v27 int8 = 100
            v28 := GenS1{F0: 357, F1: 255, F2: 31}
            _ = v27
            _ = v28
        }
        v10 = append(v10, (fi26 * -2147483648))
        if len(v10) > 0 {
            v10[0] = -95
        }
        _ = fi26
    }
    var v29 int64 = -1
    var v30 int64 = v29
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v7
    _ = v8
}
