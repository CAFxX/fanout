package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int; F2 int64 }
type GenS3 struct { F0 uint; F1 uint64; F2 bool }
func genFn0(p0 uint16, p1 string) uint64 {
    v6 := GenS1{F0: 32767, F1: 255, F2: 129}
    v7 := GenS2{F0: 251, F1: 32768, F2: 1}
    switch 2147483647 {
    case 225:
        v8 := []int{1, 32768, -2147483648}
        _ = v8
    case 42:
        p0 += (p0 ^ 2)
    case 514:
        fmt.Println(p1, p0)
    default:
        p1 = ((p1 + p1) + p1)
    }
    p0 += (p0 >> ((3) & 63))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 bool, p1 int32) uint {
    switch -2147483648 {
    case 840:
        for fi9 := 0; fi9 < 18; fi9++ {
            v10 := GenS1{F0: 32767, F1: 132, F2: 1}
            var v11 int = -222
            var v12 int32 = (-(p1))
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    default:
        var v13 int64 = 0
        _ = v13
    }
    var v14 float32 = 1e10
    var v15 uint = 0
    var v16 int32 = p1
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v16
    return v15
}

func main() {
    var v1 int32 = -6
    var v2 uint16 = 32767
    var v3 string = ""
    v4 := make(map[uint32]bool)
    v5 := []uint32{929, 698, 0}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v3 = v3
    v1 = v1
    v1 -= v1
    v3 = v3
    var v17 float64 = ((((1e10 / 0.1) * 3.14159) / 3.14159) + 3.14159)
    if true {
        switch 0 {
        case 1:
            var v18 int8 = 42
            _ = v18
        case 32768:
            v19 := GenS1{F0: 2, F1: 239, F2: 2}
            _ = v19
        }
        var v20 uint = 42
        var v21 uint64 = 1
        _ = v20
        _ = v21
    } else {
        var v22 int32 = (v1 << ((42) & 63))
        for fi23 := 0; fi23 < 12; fi23++ {
            if true {
                if ((0 >= v1) && (genFn0((v2 >> ((genFn1(false, 708)) & 63)), v3) <= 2)) {
                    v1 += (v1 - (v22 >> ((42) & 63)))
                    fmt.Println(v22, v3)
                }
            } else {
                v2 = 2
                fmt.Println(fi23)
            }
            _ = fi23
        }
        v24 := GenS2{F0: 0, F1: 4294967295, F2: -9223372036854775808}
        _ = v22
        _ = v24
    }
    v3 = ("a" + v3)
    v25 := GenS3{F0: 2, F1: 42, F2: (false || false)}
    switch -935 {
    case -573:
        v2 = 0
    case 493:
        var v26 string = (("foo bar" + (v3 + v3)) + ((v3 + "a") + ""))
        _ = v26
    default:
        for fi27 := 0; fi27 < 14; fi27++ {
            v3 = v3
            for fi28 := 0; fi28 < 19; fi28++ {
                switch (fi28 << ((305) & 63)) {
                case 370:
                    fmt.Println(v17, v1, v1)
                case 32767:
                    v1 = -2147483648
                case 127:
                    fmt.Println(fi28, v2)
                }
                _ = len(v5)
                switch 2147483647 {
                case -9223372036854775808:
                    v17 = v17
                case 32767:
                    v2 = (v2 * v2)
                }
                _ = fi28
            }
            _ = fi27
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
