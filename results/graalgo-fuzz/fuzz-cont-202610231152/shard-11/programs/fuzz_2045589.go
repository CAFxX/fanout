package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int32, p1 string) uint64 {
    fmt.Println(p1)
    if (!(!true)) {
        if (885 != 983) {
            v5 := []int{32768, -2147483648, -501, -309}
            if true {
                if false {
                    p0 = p0
                    fmt.Println(p1)
                    var v6 int32 = p0
                    _ = v6
                } else {
                    var v7 string = p1
                    _ = v7
                }
                p1 = ((p1 + p1) + p1)
                for fi8 := 0; fi8 < 5; fi8++ {
                    fmt.Println(p1, p1)
                    var v9 int8 = 127
                    _ = fi8
                    _ = v9
                }
            } else {
                var v10 uint64 = 4294967295
                _ = v10
            }
            p1 = p1
            _ = v5
        } else {
            var v11 int = -752
            _ = v11
        }
        v12 := GenS1{F0: 990}
        _ = v12
    }
    v13 := []int{}
    _ = p0
    _ = p1
    _ = v13
    return 18446744073709551615
}
func genFn1(p0 int) uint64 {
    fmt.Println(p0)
    for fi14 := 0; fi14 < 2; fi14++ {
        p0 = (fi14 << ((1) & 63))
        _ = fi14
    }
    _ = p0
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 547
    var v3 int16 = 94
    v4 := make(map[int]uint32)
    if true {
        for fi15 := 0; fi15 < 8; fi15++ {
            v1["a"] = 4294967295
            fmt.Println(v2, fi15, v2)
            _ = fi15
        }
        var v16 uint8 = 255
        _ = v16
    }
    var v17 int64 = 42
    var v18 int16 = v3
    var v19 int64 = v17
    v20 := []int{}
    if (uint32(v2) < 1) {
        for fi21 := 0; fi21 < 18; fi21++ {
            for fi22 := 0; fi22 < 15; fi22++ {
                var v23 uint = 1
                v24 := GenS1{F0: v19}
                v3 = v3
                v19 = (v17 | (v19 % (((v17 % (((v17 | v19)) | 1))) | 1)))
                v25 := GenS1{F0: v17}
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            _ = fi21
        }
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        var v27 int64 = (v19 << ((4294967295) & 63))
        _ = fi26
        _ = v27
    }
    v17 = (^(v19))
    v2 = v2
    var v28 uint = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
}
