package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 int64 = 42
    v2 := make(map[int]uint32)
    var v3 uint64 = 1
    v4 := []uint32{0, 42, 157, 1}
    if (((v1 % ((v1) | 1)) > (^(v1))) && false) {
        v5 := GenS1{F0: 0, F1: 698}
        var v6 int32 = 42
        v1 = 1
        _ = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 15; fi7++ {
            v3 = 42
            v8 := GenS2{F0: -2147483648}
            _ = fi7
            _ = v8
        }
    }
    var v9 uint8 = 255
    v3 = v3
    if false {
        v9 = v9
        if (true || (v9 != v9)) {
            v1 = (v1 << ((18446744073709551615) & 63))
            if len(v4) > 0 {
                v4[0] = 2147483647
            }
        }
        v10 := GenS3{F0: 42}
        _ = v10
    } else {
        for fi11 := 0; fi11 < 17; fi11++ {
            if (((v1 + v1) < (^(v1))) && true) {
                v3 += v3
            } else {
                v3 -= v3
            }
            v9 += 1
            _ = fi11
        }
    }
    v3 = 4294967295
    var v12 float32 = 1.5
    var v13 int64 = int64(v3)
    if false {
        v14 := GenS3{F0: 42}
        _ = v14
    }
    v1 = v1
    var v15 uint16 = 2
    for fi16 := 0; fi16 < 13; fi16++ {
        var v17 int8 = 127
        var v18 uint16 = (v15 + (v15 + 42))
        for fi19 := 0; fi19 < 8; fi19++ {
            v18 = v15
            if (179 >= -1) {
                var v20 int64 = (-(v1))
                _ = v20
            } else {
                delete(v2, (fi19 >> ((260) & 63)))
                fmt.Println(v12, v12, v17)
            }
            v21 := GenS3{F0: 42}
            _ = fi19
            _ = v21
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    if (((-2147483648 >= 50) && (0 != 917)) && false) {
        var v22 string = "hello"
        var v23 int16 = 1000
        _ = v22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
