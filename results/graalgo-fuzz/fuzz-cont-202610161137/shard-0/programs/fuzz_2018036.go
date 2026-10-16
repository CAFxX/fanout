package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 int64; F2 int8 }
func genFn0(p0 int64, p1 uint32, p2 int16) uint64 {
    v5 := GenS1{F0: p0}
    var v6 string = "hello"
    v7 := GenS2{F0: 1}
    v8 := []int{42}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int32, p1 bool, p2 uint64) uint64 {
    v9 := GenS2{F0: -2147483648}
    p1 = (575 < -9223372036854775808)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return genFn0(42, 965, 0)
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = "foo bar"
    var v3 uint = 0
    var v4 uint16 = 32767
    v10 := []int{56, 32768}
    if (0 == 2147483647) {
        if false {
            if (575 >= 32767) {
                v11 := GenS2{F0: -824}
                _ = v11
            } else {
                var v12 uint16 = v4
                fmt.Println(v3, v3)
                v13 := v1[42]
                _ = v12
                _ = v13
            }
        }
    } else {
        for fi14 := 0; fi14 < 15; fi14++ {
            if (!(true || (!false))) {
                fmt.Println(v4, v4)
                v1[fi14] = true
            }
            fmt.Println(v3)
            fmt.Println(fi14, v2)
            _ = fi14
        }
        var v15 int64 = -302
        var v16 float64 = (3.14159 - float64(0.1))
        _ = v15
        _ = v16
    }
    v1[0] = false
    for fi17 := 0; fi17 < 16; fi17++ {
        v4 += v4
        v4 = 255
        _ = fi17
    }
    if (!(!false)) {
        if (42 >= 0) {
            v18 := GenS3{F0: v2, F1: -533, F2: 0}
            _ = v18
        } else {
            v19 := v1[int(v4)]
            for fi20 := 0; fi20 < 16; fi20++ {
                fmt.Println(v3, v2, v2)
                var v21 float64 = -2.25
                for fi22 := 0; fi22 < 8; fi22++ {
                    var v23 int32 = 1
                    delete(v1, fi20)
                    _ = fi22
                    _ = v23
                }
                _ = fi20
                _ = v21
            }
            _ = v19
        }
    } else {
        v3 += (v3 << (((v3 << ((v3) & 63))) & 63))
        fmt.Println(v3, v3)
    }
    v4 = v4
    for fi24 := 0; fi24 < 13; fi24++ {
        if len(v10) > 0 {
            v10[0] = (fi24 << (((v3 * v3)) & 63))
        }
        _ = fi24
    }
    v2 = "\n\t"
    v3 = (^(v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
}
