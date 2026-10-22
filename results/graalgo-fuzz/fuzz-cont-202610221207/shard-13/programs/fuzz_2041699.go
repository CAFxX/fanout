package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int8 }
func genFn0(p0 uint32, p1 uint16, p2 int64) int8 {
    v5 := GenS1{F0: (p1 & 657), F1: 42, F2: 1}
    for fi6 := 0; fi6 < 5; fi6++ {
        p0 = p0
        v7 := GenS1{F0: (^(p1)), F1: (p2 >> ((18446744073709551615) & 63)), F2: -128}
        v8 := GenS1{F0: 0, F1: 42, F2: 42}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int32 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v9
    return 0
}

func main() {
    var v1 float64 = 0.1
    var v2 int = -863
    var v3 string = ""
    var v4 int16 = -1
    v10 := []int{}
    v11 := GenS1{F0: 32767, F1: 864, F2: 100}
    fmt.Println(v2, v3)
    var v12 int64 = -213
    if len(v10) > 0 {
        v10[0] = 2147483647
    }
    v13 := GenS1{F0: 255, F1: 739, F2: 56}
    v3 = ""
    v14 := GenS1{F0: 282, F1: v12, F2: 0}
    if true {
        var v15 uint8 = 2
        v10 = append(v10, (-(v2)))
        if false {
            v15 += v15
        } else {
            if len(v10) > 0 {
                v10[0] = (v2 ^ (v2 & (v2 & (v2 * v2))))
            }
            var v16 int8 = int8(v4)
            v4 = -347
            _ = v16
        }
        _ = v15
    } else {
        fmt.Println(v1)
        for fi17 := 0; fi17 < 11; fi17++ {
            _ = len(v10)
            fmt.Println(v4)
            _ = fi17
        }
        switch (v2 + (v2 ^ -9223372036854775808)) {
        case 2147483647:
            switch v3 {
            case "\n\t":
                for fi18 := 0; fi18 < 2; fi18++ {
                    fmt.Println(fi18, v12, v1)
                    _ = fi18
                }
            case "a":
                fmt.Println(v3)
            default:
                if (2 >= 128) {
                    v2 = (v2 << ((911) & 63))
                    var v19 bool = ((!true) || (false && (0 < 1)))
                    _ = v19
                } else {
                    var v20 uint16 = 2
                    var v21 bool = (v1 > (v1 / (v1 / float64((1.5 * 1.5)))))
                    var v22 int16 = (v4 / ((v4) | 1))
                    _ = v20
                    _ = v21
                    _ = v22
                }
            }
        }
    }
    var v23 int16 = (-(v4))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
