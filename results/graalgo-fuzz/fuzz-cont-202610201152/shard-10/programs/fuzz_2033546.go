package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint; F1 int16; F2 uint8 }
func genFn0(p0 int32) uint8 {
    if (-208 > 2147483647) {
        if true {
            p0 = (p0 << ((4294967295) & 63))
        } else {
            if (255 <= 65535) {
                v7 := GenS1{F0: ""}
                _ = v7
            } else {
                v8 := GenS1{F0: "hello"}
                _ = v8
            }
        }
    }
    v9 := []int{2147483647, -1}
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 int64 = -1
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = v9
    return 177
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint32 = 4294967295
    var v3 uint = 938
    var v4 int8 = -1
    var v5 int32 = -2147483648
    v6 := make(map[uint32]string)
    var v12 uint32 = (^(v2))
    fmt.Println(v5)
    for fi13 := 0; fi13 < 15; fi13++ {
        for fi14 := 0; fi14 < 13; fi14++ {
            switch "\n\t" {
            case "hello":
                fmt.Println(v2)
                v15 := GenS2{F0: (-(v5)), F1: v1}
                _ = v15
            }
            _ = fi14
        }
        if (v3 != (v3 << ((v3) & 63))) {
            v16 := []int{fi13, fi13, fi13}
            _ = v16
        }
        _ = fi13
    }
    var v17 float64 = 1e10
    switch "hello" {
    case "hello":
        var v18 int64 = 922
        _ = v18
    case "\n\t":
        fmt.Println(v5)
    default:
        switch -102 {
        case 2147483647:
            v19 := []int{-30}
            _ = v19
        }
    }
    if (-9223372036854775808 > -323) {
        if false {
            var v20 uint8 = 255
            var v21 uint16 = 2
            var v22 float64 = v17
            _ = v20
            _ = v21
            _ = v22
        }
        v1 -= uint64(v12)
        v1 -= 575
    } else {
        for fi23 := 0; fi23 < 16; fi23++ {
            fmt.Println(v17, v4, v5)
            _ = fi23
        }
    }
    if false {
        fmt.Println(v17, v4, v4)
        v1 += 2
        v6[395] = ("" + "\n\t")
    }
    if (true && (v1 < v1)) {
        v12 += (v2 % ((v12) | 1))
        v24 := GenS3{F0: (v3 * 2), F1: 1000, F2: uint8(v12)}
        _ = v24
    } else {
        v2 = v2
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        for fi26 := 0; fi26 < 1; fi26++ {
            var v27 int16 = 0
            _ = fi26
            _ = v27
        }
        var v28 uint = v3
        _ = fi25
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
