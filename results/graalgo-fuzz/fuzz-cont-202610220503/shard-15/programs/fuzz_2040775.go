package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint64 }
func genFn0(p0 uint64) int {
    p0 = (^(p0))
    for fi5 := 0; fi5 < 16; fi5++ {
        p0 += (^(p0))
        var v6 uint64 = p0
        fmt.Println(fi5)
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 7; fi7++ {
        if true {
            if (920 < -1) {
                if false {
                    var v8 uint16 = 1
                    var v9 int64 = 0
                    var v10 bool = (((p0 * p0) == p0) && false)
                    _ = v10
                    _ = v8
                    _ = v9
                }
                var v11 uint = 568
                var v12 int = fi7
                _ = v11
                _ = v12
            }
        }
        _ = fi7
    }
    v13 := GenS1{F0: 0}
    _ = p0
    _ = v13
    return -2147483648
}
func genFn1() uint32 {
    for fi14 := 0; fi14 < 6; fi14++ {
        v15 := GenS1{F0: 1}
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: -128}
    _ = v16
    return 2147483648
}

func main() {
    v1 := make(map[int]string)
    var v2 int16 = 689
    var v3 uint32 = 4294967295
    var v4 uint8 = 42
    var v17 string = "hello"
    v17 = "hello"
    v17 = v17
    for fi18 := 0; fi18 < 9; fi18++ {
        delete(v1, -1)
        _ = fi18
    }
    fmt.Println(v3, v4)
    _ = len(v1)
    v2 = (v2 * -303)
    for fi19 := 0; fi19 < 18; fi19++ {
        v20 := GenS2{F0: (v4 <= v4), F1: 903}
        var v21 int32 = -1
        fmt.Println(fi19, fi19)
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        fmt.Println(v17, v4, v3)
        if true {
            fmt.Println(fi22, fi22, v17)
            v2 -= v2
            v2 -= (v2 << ((2) & 63))
        }
        _ = fi22
    }
    fmt.Println(v4)
    v3 = v3
    delete(v1, genFn0(18446744073709551615))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
