package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int8; F1 uint8 }
func genFn0(p0 uint16, p1 int64, p2 uint16) int64 {
    var v5 uint8 = 37
    var v6 int16 = 926
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p1
}

func main() {
    var v1 int = 32768
    var v2 int16 = -1
    v3 := make(map[string]bool)
    var v4 uint8 = 42
    if (2 < 2147483648) {
        switch v1 {
        case 127:
            v7 := []int{v1, 1000, 4294967295}
            _ = v7
        case -9223372036854775808:
            fmt.Println(v4, v4, v1)
        default:
            if (42 < 0) {
                for fi8 := 0; fi8 < 14; fi8++ {
                    for fi9 := 0; fi9 < 19; fi9++ {
                        var v10 float32 = 1e10
                        _ = fi9
                        _ = v10
                    }
                    _ = fi8
                }
                delete(v3, "")
                var v11 uint32 = 65535
                _ = v11
            }
        }
        var v12 uint = 2
        var v13 int16 = v2
        _ = v12
        _ = v13
    } else {
        delete(v3, (("a" + "hello") + ""))
        v1 = (v1 >> ((1) & 63))
        v14 := GenS1{F0: (^(v4)), F1: 42}
        _ = v14
    }
    var v15 uint = 4294967295
    var v16 uint64 = 2
    for fi17 := 0; fi17 < 15; fi17++ {
        fmt.Println(v15, v4, v16)
        _ = fi17
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        v3["a"] = ((32767 < 255) || false)
        for fi19 := 0; fi19 < 11; fi19++ {
            var v20 float32 = 1e10
            v21 := GenS1{F0: (^(v4)), F1: v15}
            var v22 uint32 = 2147483648
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        _ = fi18
    }
    v23 := []int{}
    v3[("" + "a")] = (false && false)
    v24 := GenS1{F0: (^(v4)), F1: (^(v15))}
    v25 := GenS1{F0: v4, F1: 364}
    v2 -= -1
    for fi26 := 0; fi26 < 17; fi26++ {
        v2 = v2
        _ = fi26
    }
    v2 += (v2 + v2)
    var v27 int64 = -1
    v28 := v3["hello"]
    v1 = v1
    for fi29 := 0; fi29 < 8; fi29++ {
        v30 := GenS2{F0: 1, F1: 9}
        _ = fi29
        _ = v30
    }
    var v31 uint32 = 180
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
}
