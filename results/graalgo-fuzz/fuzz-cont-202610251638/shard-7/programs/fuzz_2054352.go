package main

import (
    "fmt"
)

func genFn0() uint {
    var v7 uint16 = 276
    v7 += 32767
    var v8 int64 = -234
    _ = v7
    _ = v8
    return 323
}

func main() {
    var v1 uint8 = 161
    var v2 int = 362
    v3 := []uint32{2147483648, 170}
    var v4 string = ""
    v5 := make(map[uint32]bool)
    var v6 uint = 948
    v3 = append(v3, 42)
    v6 += (v6 << ((v6) & 63))
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 uint32 = 2147483647
        _ = fi9
        _ = v10
    }
    var v11 uint8 = (^(v1))
    if false {
        fmt.Println(v11)
    }
    v1 += (^(v1))
    v11 -= v1
    if (4294967295 <= 0) {
        for fi12 := 0; fi12 < 15; fi12++ {
            v2 = (fi12 / (((v2 + -197)) | 1))
            for fi13 := 0; fi13 < 5; fi13++ {
                if len(v3) > 0 {
                    v3[0] = 194
                }
                var v14 int8 = -30
                fmt.Println(fi13)
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    } else {
        var v15 uint = 2
        _ = v15
    }
    var v16 uint64 = 504
    for fi17 := 0; fi17 < 8; fi17++ {
        if ((v2 / (((v2 * v2)) | 1)) != (-(fi17))) {
            fmt.Println(v11, v1)
        } else {
            if len(v3) > 0 {
                v3[0] = 202
            }
        }
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
