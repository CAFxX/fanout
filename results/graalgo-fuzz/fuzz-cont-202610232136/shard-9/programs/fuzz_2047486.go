package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int; F1 int32; F2 uint64 }
func genFn0(p0 uint8, p1 uint64, p2 int32) uint32 {
    if (1 >= -672) {
        v4 := []int{1}
        fmt.Println(p1, p1, p0)
        _ = v4
    } else {
        fmt.Println(p0, p1)
        fmt.Println(p0, p2, p1)
        v5 := []int{}
        _ = v5
    }
    fmt.Println(p0, p0)
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := GenS1{F0: 1, F1: 4294967295}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 967
}

func main() {
    v1 := []int{9223372036854775807}
    v2 := []uint32{947, 42, 1, 566}
    var v3 float32 = 0.0
    v3 -= 0.1
    _ = len(v1)
    for fi8 := 0; fi8 < 2; fi8++ {
        v3 = -2.25
        if true {
            var v9 uint64 = 4294967295
            if ((0 < 1) && false) {
                v3 += v3
            }
            _ = v9
        } else {
            for fi10 := 0; fi10 < 14; fi10++ {
                var v11 uint64 = 1
                if (int8(fi8) > -21) {
                    fmt.Println(fi8, v3)
                    v12 := GenS1{F0: 32767, F1: uint(v11)}
                    _ = v12
                }
                var v13 string = ((("foo bar" + "hello") + "\n\t") + "foo bar")
                v2 = append(v2, 968)
                v11 = v11
                _ = fi10
                _ = v11
                _ = v13
            }
            var v14 string = "foo bar"
            v15 := GenS1{F0: 2, F1: 42}
            _ = v14
            _ = v15
        }
        fmt.Println(v3, v3)
        _ = fi8
    }
    v16 := GenS1{F0: 491, F1: 120}
    var v17 uint8 = 132
    for fi18 := 0; fi18 < 12; fi18++ {
        v17 -= 2
        v17 += (v17 | v17)
        switch -871 {
        case 580:
            for fi19 := 0; fi19 < 20; fi19++ {
                v3 = 1.5
                v17 += (v17 << ((588) & 63))
                _ = len(v1)
                _ = fi19
            }
        case 80:
            var v20 uint32 = 2147483648
            _ = v20
        case -874:
            v1 = append(v1, 0)
        }
        _ = fi18
    }
    var v21 int = 127
    if (-1 >= 306) {
        v22 := GenS2{F0: v21, F1: 1, F2: 634}
        var v23 int = v21
        var v24 bool = (-1 == 42)
        _ = v22
        _ = v23
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
}
