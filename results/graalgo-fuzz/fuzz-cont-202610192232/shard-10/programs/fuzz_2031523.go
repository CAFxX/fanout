package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 uint64) int16 {
    if false {
        v4 := []int{-238}
        if false {
            for fi5 := 0; fi5 < 15; fi5++ {
                p0 = p0
                _ = fi5
            }
        }
        _ = v4
    }
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := []int{4294967295, (fi6 % ((fi6) | 1)), (-(fi6)), -9223372036854775808}
        for fi8 := 0; fi8 < 9; fi8++ {
            v9 := GenS1{F0: 32767}
            _ = fi8
            _ = v9
        }
        fmt.Println(fi6, p0)
        _ = fi6
        _ = v7
    }
    v10 := GenS1{F0: 589}
    v11 := []int{0, 127, 926, 2147483647}
    _ = p0
    _ = v10
    _ = v11
    return 0
}

func main() {
    var v1 uint8 = 42
    var v2 int32 = 1
    var v3 string = "a"
    for fi12 := 0; fi12 < 11; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 int64 = -1
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    v15 := []int{32767, 1, -9223372036854775808, int(v1)}
    var v16 int16 = -32768
    if len(v15) > 0 {
        v15[0] = 9223372036854775807
    }
    if true {
        fmt.Println(v2)
        for fi17 := 0; fi17 < 13; fi17++ {
            var v18 int64 = 0
            var v19 uint32 = 1
            _ = fi17
            _ = v18
            _ = v19
        }
    } else {
        v1 -= (v1 ^ v1)
        for fi20 := 0; fi20 < 19; fi20++ {
            v21 := GenS1{F0: 2}
            v3 = (v3 + v3)
            v1 = v1
            _ = fi20
            _ = v21
        }
        var v22 int8 = 100
        _ = v22
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        if false {
            v24 := GenS1{F0: 1}
            v1 = (v1 / (((^(v1))) | 1))
            if false {
                v16 = genFn0(476)
            }
            _ = v24
        } else {
            var v25 string = ""
            _ = v25
        }
        _ = fi23
    }
    var v26 string = (v3 + "a")
    var v27 float32 = (1.5 * (1.5 + 3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v3
}
