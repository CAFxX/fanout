package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 bool; F1 int32; F2 int32 }
func genFn0(p0 string) int {
    p0 = p0
    if false {
        v6 := GenS1{F0: 1}
        _ = v6
    }
    _ = p0
    return 127
}
func genFn1(p0 int64, p1 uint32, p2 bool) int16 {
    v7 := []int{715}
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 int16 = -32768
        fmt.Println(v9)
        v10 := GenS2{F0: genFn0("hello"), F1: 100}
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi11
    }
    var v12 int = 4294967295
    var v13 string = ""
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v7
    return 32767
}

func main() {
    var v1 int16 = -1
    var v2 uint64 = 4294967295
    v3 := []int{-9223372036854775808, -1, -1}
    var v4 int = -9223372036854775808
    var v5 bool = false
    v2 = (^(v2))
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 int16 = (^(v1))
        _ = fi14
        _ = v15
    }
    var v16 int16 = v1
    if v5 {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v17 int8 = -1
        _ = v17
    }
    if false {
        var v18 int32 = -2147483648
        _ = v18
    }
    v1 -= (v1 << ((648) & 63))
    fmt.Println(v1)
    if len(v3) > 0 {
        v3[0] = (-(v4))
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        if (!(v5 || (!v5))) {
            if v5 {
                if (!(18446744073709551615 <= 0)) {
                    var v20 uint8 = 127
                    v21 := GenS3{F0: v5, F1: 2147483647, F2: -2147483648}
                    _ = v20
                    _ = v21
                } else {
                    v4 = fi19
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v2 = 1
            } else {
                v5 = (!v5)
            }
        } else {
            var v22 uint32 = uint32(v2)
            for fi23 := 0; fi23 < 10; fi23++ {
                fmt.Println(v2, v4, v2)
                _ = fi23
            }
            _ = v22
        }
        var v24 float32 = 1.5
        _ = fi19
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
