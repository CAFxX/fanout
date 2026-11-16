package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint8, p1 bool, p2 uint8) int8 {
    v4 := []int{32768}
    var v5 int = -9223372036854775808
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 bool, p1 int16, p2 uint64) uint32 {
    var v6 uint32 = 65535
    p0 = ((9223372036854775807 >= 399) && (2147483647 <= 76))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return v6
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 1
    v3 := make(map[string]string)
    v2 = (^(v2))
    if false {
        v2 += v2
    } else {
        fmt.Println(v2)
    }
    if false {
        delete(v3, "hello")
        v2 = (v2 << ((479) & 63))
        switch -282 {
        case -1:
            var v7 string = "\n\t"
            _ = v7
        default:
            v2 -= (v2 ^ v2)
        }
    }
    var v8 int32 = 42
    var v9 string = "a"
    v8 -= (v8 << ((18446744073709551615) & 63))
    for fi10 := 0; fi10 < 17; fi10++ {
        v8 = 0
        _ = fi10
    }
    v2 += v2
    v11 := GenS1{F0: 1, F1: 42}
    var v12 bool = false
    if (0 >= 0) {
        v8 -= (v8 - v8)
    }
    var v13 int64 = -203
    if v12 {
        v12 = (1 >= 65535)
        v14 := v3[""]
        v9 = v14
        _ = v14
    } else {
        fmt.Println(v12, v8, v12)
        v15 := GenS1{F0: 154, F1: 2}
        var v16 int16 = -32768
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        v13 = int64(fi17)
        switch -2147483648 {
        case -2147483648:
            var v18 bool = (v2 >= (v2 + v2))
            _ = v18
        }
        _ = fi17
    }
    _ = len(v1)
    var v19 int = 32768
    v20 := v3["foo bar"]
    if (32767 <= 1) {
        switch v19 {
        case 2147483647:
            var v21 float32 = (0.0 - 0.1)
            _ = v21
        case 640:
            var v22 int32 = (v8 * v8)
            _ = v22
        case 646:
            for fi23 := 0; fi23 < 7; fi23++ {
                v24 := GenS1{F0: 2, F1: 18446744073709551615}
                _ = fi23
                _ = v24
            }
        default:
            var v25 uint = 4294967295
            _ = v25
        }
        v19 = v19
    } else {
        v8 += 42
        var v26 bool = (float64((0.0 * 0.1)) != float64((1e10 - float32(0.0))))
        var v27 int8 = genFn0(17, (!(1 >= 76)), 22)
        _ = v26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v8
    _ = v9
}
