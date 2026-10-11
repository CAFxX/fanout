package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int8, p1 string, p2 uint32) uint64 {
    v7 := []int{}
    if ((1 >= -455) || (-97 < -542)) {
        var v8 float32 = 0.1
        v7 = append(v7, -2147483648)
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}
func genFn1(p0 int16, p1 uint, p2 uint16) string {
    var v9 uint16 = (p2 >> ((65535) & 63))
    v9 = v9
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return "\n\t"
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 2
    v3 := []int{-780}
    v4 := []int{-276, 1}
    v5 := make(map[uint32]string)
    var v6 uint32 = 2147483647
    var v10 int = 128
    for fi11 := 0; fi11 < 12; fi11++ {
        v12 := GenS1{F0: 18446744073709551615}
        fmt.Println(v6, v10, v10)
        _ = fi11
        _ = v12
    }
    if true {
        fmt.Println(v6)
        v10 -= v10
        v10 = v10
    }
    fmt.Println(v2, v2, v10)
    var v13 uint64 = (v2 / ((uint64(v6)) | 1))
    switch (v10 / (((v10 * 4294967295)) | 1)) {
    case 710:
        if false {
            v1[""] = (v6 + v6)
        }
    case -564:
        if false {
            switch (v10 / (((v10 >> ((4294967295) & 63))) | 1)) {
            case 2147483647:
                for fi14 := 0; fi14 < 19; fi14++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi14
                }
            case 961:
                v10 += (-(v10))
            case -9223372036854775808:
                fmt.Println(v13)
            }
        } else {
            _ = len(v3)
            var v15 bool = (false && false)
            _ = v15
        }
    }
    fmt.Println(v6)
    v2 += 4294967295
    v16 := GenS1{F0: 943}
    for fi17 := 0; fi17 < 8; fi17++ {
        v6 -= 536
        fmt.Println(v6, v6)
        _ = fi17
    }
    v13 += 453
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
