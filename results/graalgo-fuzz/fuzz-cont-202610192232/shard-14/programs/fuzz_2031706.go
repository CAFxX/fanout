package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint64 }
func genFn0(p0 uint64, p1 int32) int16 {
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 uint8 = uint8(p0)
        _ = fi6
        _ = v7
    }
    v8 := []int{42, 0, 128}
    if true {
        var v9 uint16 = 2
        if len(v8) > 0 {
            v8[0] = 42
        }
        if len(v8) > 0 {
            v8[0] = 586
        }
        _ = v9
    }
    var v10 uint16 = 326
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    return 70
}
func genFn1(p0 uint32, p1 uint16) int64 {
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := []int{(fi11 * (fi11 << ((255) & 63)))}
        var v13 string = (("" + ("hello" + "\n\t")) + ("\n\t" + ""))
        v13 = v13
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := []int{127, -1, -158}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v14
    return 235
}

func main() {
    var v1 uint = 936
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    var v4 int16 = 32767
    v5 := make(map[string]string)
    v1 = (v1 ^ (v1 & v1))
    for fi15 := 0; fi15 < 18; fi15++ {
        v4 -= v4
        _ = fi15
    }
    v1 += (v1 + (^(v1)))
    v16 := []int{0}
    for fi17 := 0; fi17 < 6; fi17++ {
        v4 += (v4 | 32767)
        fmt.Println(v4)
        v18 := GenS1{F0: (v4 / ((v4) | 1)), F1: 1, F2: 994}
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        switch (-(fi19)) {
        case 2147483647:
            var v20 float32 = ((-2.25 - -2.25) + 1.5)
            _ = v20
        default:
            delete(v5, ("hello" + "a"))
        }
        v16 = append(v16, 32767)
        _ = fi19
    }
    switch 399 {
    case 42:
        v4 -= (v4 + genFn0(0, 1))
    case -1:
        v1 -= (v1 - v1)
    case -2147483648:
        var v21 int32 = 0
        _ = v21
    default:
        fmt.Println(v1)
    }
    var v22 string = "foo bar"
    var v23 uint64 = 2
    v4 = genFn0(v23, int32(v4))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
