package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 bool, p1 int16) uint8 {
    v5 := GenS1{F0: 42}
    var v6 int32 = -112
    v7 := GenS1{F0: 524}
    var v8 uint = 42
    if (1 >= 2) {
        if p0 {
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 bool = false
                v10 = (0 < 128)
                _ = fi9
                _ = v10
            }
        }
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 bool, p1 uint16) int32 {
    if true {
        var v11 int32 = 944
        v12 := []int{}
        _ = v11
        _ = v12
    } else {
        p0 = (0 != -9223372036854775808)
        var v13 int8 = -1
        _ = v13
    }
    if false {
        v14 := []int{}
        _ = v14
    } else {
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := []uint32{895, 232, 2, 4294967295}
    var v2 float32 = (0.0 + 0.0)
    var v3 uint8 = 1
    var v4 uint32 = 2
    for fi15 := 0; fi15 < 3; fi15++ {
        v4 = 568
        v4 -= 0
        _ = fi15
    }
    v2 -= v2
    var v16 int = 32768
    if (false && false) {
        v1 = append(v1, 42)
    }
    var v17 bool = false
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v3)
    if true {
        fmt.Println(v3, v16)
        for fi18 := 0; fi18 < 9; fi18++ {
            for fi19 := 0; fi19 < 17; fi19++ {
                fmt.Println(v3, v17, fi18)
                if true {
                    v17 = true
                    v3 = v3
                    if len(v1) > 0 {
                        v1[0] = (v4 + (v4 | 2))
                    }
                }
                _ = fi19
            }
            _ = fi18
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        if false {
            fmt.Println(v3, v4)
        }
        if len(v1) > 0 {
            v1[0] = (v4 >> ((18446744073709551615) & 63))
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
