package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint64 }

func main() {
    var v1 int32 = -1
    v2 := make(map[uint32]int)
    var v3 uint16 = 0
    v4 := []int{1}
    var v5 int32 = (v1 & v1)
    v1 += v1
    fmt.Println(v3, v3, v1)
    var v6 uint64 = 4294967295
    if true {
        v1 = v5
        if false {
            v3 += 147
            _ = len(v2)
        }
        switch 42 {
        case 2147483647:
            v7 := GenS1{F0: 0, F1: 42, F2: 42}
            _ = v7
        case -2147483648:
            v8 := GenS1{F0: 10, F1: -18, F2: 2}
            _ = v8
        default:
            for fi9 := 0; fi9 < 17; fi9++ {
                for fi10 := 0; fi10 < 3; fi10++ {
                    v4 = append(v4, 0)
                    if true {
                        v5 -= -2147483648
                    }
                    v11 := GenS1{F0: -103, F1: (fi10 % ((0) | 1)), F2: 2}
                    _ = fi10
                    _ = v11
                }
                fmt.Println(v3, v3)
                _ = len(v2)
                _ = fi9
            }
        }
    } else {
        v3 -= v3
        var v12 uint8 = 0
        _ = v12
    }
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v13 uint16 = 255
    var v14 uint = 2
    v6 = (v6 & v6)
    if (v5 > v5) {
        for fi15 := 0; fi15 < 8; fi15++ {
            _ = len(v4)
            _ = fi15
        }
        v16 := GenS1{F0: 100, F1: 32768, F2: v6}
        var v17 int = -9223372036854775808
        _ = v16
        _ = v17
    }
    if false {
        v14 -= 65535
    } else {
        fmt.Println(v5, v6)
        var v18 float32 = 3.14159
        _ = v18
    }
    if (false && (true || true)) {
        if true {
            for fi19 := 0; fi19 < 14; fi19++ {
                fmt.Println(v14, v13)
                var v20 int32 = (v1 - 0)
                if (!false) {
                    v3 = (v13 ^ (v3 * v3))
                    var v21 uint = (^(v14))
                    v3 += 42
                    _ = v21
                } else {
                    v2[0] = (fi19 | (fi19 / (((fi19 % (((^(fi19))) | 1))) | 1)))
                    v14 = 2
                }
                _ = fi19
                _ = v20
            }
        } else {
            v5 += (v1 - (v1 | v5))
            var v22 string = "a"
            _ = v22
        }
    }
    var v23 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
