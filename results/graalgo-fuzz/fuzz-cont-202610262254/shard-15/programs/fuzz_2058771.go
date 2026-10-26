package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
func genFn0(p0 uint16, p1 int8, p2 uint64) int32 {
    var v5 uint64 = (p2 * (p2 / ((61) | 1)))
    if ((27 >= 255) || false) {
        if true {
            v6 := []int{}
            if ((!(-549 != 127)) || true) {
                p2 += (^(p2))
                switch 4294967295 {
                case 9223372036854775807:
                    if len(v6) > 0 {
                        v6[0] = 9223372036854775807
                    }
                case 128:
                    if false {
                        var v7 uint = 966
                        v7 += v7
                        v8 := GenS1{F0: 42, F1: p0}
                        _ = v7
                        _ = v8
                    }
                case 1:
                    fmt.Println(p2, p2)
                }
                for fi9 := 0; fi9 < 4; fi9++ {
                    for fi10 := 0; fi10 < 11; fi10++ {
                        v11 := GenS1{F0: 155, F1: p0}
                        fmt.Println(fi10)
                        v6 = append(v6, (fi9 ^ int(p0)))
                        _ = fi10
                        _ = v11
                    }
                    var v12 uint8 = 0
                    p0 = uint16(p2)
                    _ = fi9
                    _ = v12
                }
            }
            _ = v6
        } else {
            var v13 uint8 = 128
            fmt.Println(v5, v13, v5)
            var v14 int8 = p1
            _ = v13
            _ = v14
        }
    }
    var v15 float32 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v5
    return 0
}

func main() {
    var v1 string = "hello"
    var v2 int64 = -9223372036854775808
    var v3 string = (v1 + v1)
    v4 := make(map[int]uint32)
    v2 += v2
    var v16 float32 = 3.14159
    v1 = "a"
    v17 := []int{128, -850}
    for fi18 := 0; fi18 < 18; fi18++ {
        if (184 >= 255) {
            var v19 uint32 = 65535
            fmt.Println(v2, v2, v1)
            _ = v19
        }
        fmt.Println(v2)
        var v20 int32 = 42
        _ = fi18
        _ = v20
    }
    v17 = append(v17, 674)
    v21 := GenS1{F0: 255, F1: 0}
    if false {
        _ = len(v4)
        var v22 uint16 = 65535
        if true {
            var v23 uint64 = 230
            var v24 bool = false
            switch 32768 {
            case -437:
                var v25 float32 = v16
                _ = v25
            }
            _ = v23
            _ = v24
        }
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
