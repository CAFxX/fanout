package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64 }
func genFn0(p0 int, p1 int64) uint64 {
    if false {
        if (2 >= 714) {
            v7 := GenS1{F0: (!(!(false || (18446744073709551615 != 2))))}
            var v8 int32 = 0
            p0 = (^(p0))
            _ = v7
            _ = v8
        }
    } else {
        p1 = int64(p0)
        v9 := GenS1{F0: true}
        var v10 uint64 = 1
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := []int{p0, (fi11 << ((65535) & 63)), 1, (-(fi11))}
        v13 := GenS2{F0: 255}
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := []int{717, (^(p0)), -9223372036854775808, p0}
    var v15 int = p0
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return 145
}
func genFn1() int {
    var v16 int = 0
    v16 = v16
    v17 := GenS2{F0: 835}
    v18 := []int{v16, 42, (v16 >> ((29) & 63)), (v16 * v16)}
    if false {
        for fi19 := 0; fi19 < 17; fi19++ {
            for fi20 := 0; fi20 < 17; fi20++ {
                v21 := GenS2{F0: 2}
                fmt.Println(fi20, v16, v16)
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
        if len(v18) > 0 {
            v18[0] = 4294967295
        }
        var v22 float32 = 1.5
        _ = v22
    }
    _ = v16
    _ = v17
    _ = v18
    return (v16 & v16)
}

func main() {
    v1 := []uint32{65535}
    var v2 int32 = 306
    var v3 uint8 = 245
    v4 := make(map[int]bool)
    v5 := make(map[uint32]bool)
    v6 := make(map[uint32]string)
    switch genFn1() {
    case -1:
        v3 -= (v3 - v3)
    case -2147483648:
        v23 := GenS3{F0: 127}
        _ = v23
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        for fi25 := 0; fi25 < 6; fi25++ {
            v26 := GenS4{F0: 656}
            v2 = (v2 >> ((2) & 63))
            var v27 bool = false
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = len(v1)
        _ = fi24
    }
    v28 := GenS4{F0: 1}
    v3 = v3
    v6[14] = ""
    fmt.Println(v2)
    v2 = (v2 & (v2 << ((4294967295) & 63)))
    v2 -= (-(v2))
    delete(v4, 42)
    v3 = v3
    v3 += (v3 + 1)
    v3 -= (v3 * (v3 * v3))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
