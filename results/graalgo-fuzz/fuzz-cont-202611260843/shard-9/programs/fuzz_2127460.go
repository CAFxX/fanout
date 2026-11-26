package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint32; F2 string }
type GenS3 struct { F0 int; F1 int32; F2 int32 }
func genFn0(p0 int64, p1 int64) uint32 {
    p1 = (p1 - (-(p1)))
    var v5 uint16 = 32767
    for fi6 := 0; fi6 < 4; fi6++ {
        v5 = (v5 - v5)
        _ = fi6
    }
    v7 := []int{42, 1}
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    return 2147483648
}
func genFn1(p0 uint, p1 int32) uint {
    v8 := GenS1{F0: 4294967295}
    p1 -= p1
    v9 := GenS2{F0: 32768, F1: 209, F2: "foo bar"}
    if true {
        var v10 uint8 = 0
        p0 += p0
        _ = v10
    } else {
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return p0
}

func main() {
    var v1 int = 1
    v2 := []int{v1, (v1 * v1), v1}
    v3 := make(map[int]string)
    var v4 uint32 = 65535
    for fi11 := 0; fi11 < 10; fi11++ {
        delete(v3, -9223372036854775808)
        v1 = (^(v1))
        if false {
            for fi12 := 0; fi12 < 12; fi12++ {
                if len(v2) > 0 {
                    v2[0] = (fi12 << ((0) & 63))
                }
                _ = len(v2)
                _ = fi12
            }
            var v13 int64 = -1
            v13 = v13
            _ = v13
        }
        _ = fi11
    }
    var v14 int16 = 1000
    switch 127 {
    case -45:
        v15 := GenS3{F0: v1, F1: 1, F2: 42}
        _ = v15
    }
    var v16 int64 = int64(v14)
    v16 += 42
    switch -131 {
    case -2147483648:
        var v17 uint8 = 2
        _ = v17
    }
    var v18 uint32 = v4
    for fi19 := 0; fi19 < 1; fi19++ {
        v16 = (v16 ^ 9223372036854775807)
        _ = fi19
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        var v21 string = ""
        var v22 int32 = 449
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 int64 = v16
    v18 += genFn0((-(v16)), int64(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v16
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
