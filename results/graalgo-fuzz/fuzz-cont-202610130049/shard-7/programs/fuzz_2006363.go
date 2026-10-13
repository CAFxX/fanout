package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int; F2 int32 }
func genFn0() uint64 {
    var v5 int = 0
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint8 = 2
        _ = fi6
        _ = v7
    }
    if true {
        var v8 uint8 = 127
        if ((!true) && false) {
            var v9 int32 = 942
            _ = v9
        }
        _ = v8
    } else {
        var v10 int64 = -9223372036854775808
        var v11 uint64 = 1
        for fi12 := 0; fi12 < 13; fi12++ {
            if (false || false) {
                if true {
                    for fi13 := 0; fi13 < 5; fi13++ {
                        var v14 int32 = 120
                        fmt.Println(fi12, v11)
                        _ = fi13
                        _ = v14
                    }
                    v10 = v10
                    v15 := GenS1{F0: 255, F1: -2147483648, F2: 1}
                    _ = v15
                }
            }
            fmt.Println(v11, fi12)
            _ = fi12
        }
        _ = v10
        _ = v11
    }
    _ = v5
    return 42
}

func main() {
    var v1 uint = 2
    var v2 int64 = -1
    var v3 uint8 = 2
    v4 := make(map[uint32]int)
    fmt.Println(v1)
    v16 := []int{0, 42}
    var v17 uint = (v1 + v1)
    v18 := GenS2{F0: -2147483648, F1: 255}
    for fi19 := 0; fi19 < 20; fi19++ {
        var v20 uint = v17
        if false {
            for fi21 := 0; fi21 < 12; fi21++ {
                var v22 bool = (4294967295 == 42)
                v23 := GenS3{F0: 308, F1: (fi21 % (((^(fi21))) | 1)), F2: 0}
                v24 := v4[280]
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            v3 += (v3 ^ 245)
        } else {
            var v25 uint16 = 264
            _ = v25
        }
        _ = fi19
        _ = v20
    }
    for fi26 := 0; fi26 < 14; fi26++ {
        if true {
            if (((false || (32767 >= 42)) && (v17 == (v1 >> ((847) & 63)))) || (!true)) {
                var v27 bool = true
                _ = v27
            }
            _ = len(v4)
        }
        _ = fi26
    }
    for fi28 := 0; fi28 < 1; fi28++ {
        if len(v16) > 0 {
            v16[0] = fi28
        }
        _ = fi28
    }
    var v29 bool = (837 >= 384)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
