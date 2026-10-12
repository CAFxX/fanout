package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{42, 4294967295, 54}
    var v2 uint16 = 938
    var v3 float32 = float32(1.5)
    var v4 uint32 = 639
    var v5 uint8 = 127
    v5 -= v5
    v4 = 2147483647
    fmt.Println(v4, v4, v3)
    v4 += (v4 & (v4 - v4))
    v2 = 32767
    var v6 bool = true
    _ = len(v1)
    fmt.Println(v4, v6, v6)
    switch 0 {
    case -635:
        switch 2147483647 {
        case 127:
            var v7 float64 = -2.25
            _ = v7
        case 627:
            for fi8 := 0; fi8 < 3; fi8++ {
                var v9 uint8 = 8
                var v10 uint = 18446744073709551615
                fmt.Println(v4, v6, v5)
                _ = fi8
                _ = v10
                _ = v9
            }
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    default:
        fmt.Println(v4, v4, v4)
    }
    var v11 uint = 65535
    v2 += (v2 << (((v11 << ((v11) & 63))) & 63))
    v5 = v5
    for fi12 := 0; fi12 < 5; fi12++ {
        v11 = uint(v2)
        _ = fi12
    }
    if v6 {
        var v13 uint64 = 711
        var v14 int = 32768
        if true {
            fmt.Println(v6)
        } else {
            if (("\n\t" + "") >= "hello") {
                v13 = uint64(v11)
                fmt.Println(v11, v13)
                v4 -= v4
            } else {
                v5 += v5
                var v15 int32 = 1
                _ = v15
            }
            for fi16 := 0; fi16 < 17; fi16++ {
                for fi17 := 0; fi17 < 10; fi17++ {
                    if len(v1) > 0 {
                        v1[0] = v4
                    }
                    _ = len(v1)
                    var v18 uint8 = (v5 / ((v5) | 1))
                    fmt.Println(v14, fi17)
                    v13 = v13
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
        }
        _ = v13
        _ = v14
    } else {
        var v19 int = 42
        _ = v19
    }
    v4 = (v4 * v4)
    var v20 uint64 = 0
    v6 = v6
    var v21 int8 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
