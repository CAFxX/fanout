package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 bool, p1 int64, p2 uint64) int {
    var v6 uint16 = 0
    fmt.Println(p2, p2, v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 128
}
func genFn1(p0 uint32, p1 uint) uint {
    if (42 > 894) {
        if true {
            v7 := GenS1{F0: (^(p0))}
            var v8 uint8 = 255
            var v9 uint32 = (p0 - p0)
            _ = v7
            _ = v8
            _ = v9
        }
        v10 := []int{4294967295, 2147483647}
        _ = v10
    }
    if ((1.5 + ((3.14159 - ((1.5 + 3.14159) / (0.1 + 1.5))) / 1.0)) > 1.5) {
        v11 := GenS1{F0: (p0 * 96)}
        p1 += (p1 | 129)
        _ = v11
    }
    if true {
        switch genFn0(true, -1, 42) {
        case 42:
            p0 = p0
        case -9223372036854775808:
            v12 := []int{32767, 32768, genFn0(false, int64(p1), 42), genFn0((4294967295 < 677), -454, 2)}
            _ = v12
        case 32768:
            if (p1 != (p1 * (^(p1)))) {
                switch -693 {
                case -104:
                    p1 = (p1 << ((p1) & 63))
                case -793:
                    var v13 uint16 = 32767
                    _ = v13
                }
            }
        }
    }
    switch -1 {
    case -2147483648:
        var v14 uint8 = 42
        _ = v14
    case 0:
        var v15 int64 = 9223372036854775807
        _ = v15
    case 1:
        fmt.Println(p0, p1)
    default:
        v16 := []int{4294967295, -234}
        _ = v16
    }
    _ = p0
    _ = p1
    return uint(p0)
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 1
    v3 := []int{891}
    var v4 int16 = v2
    var v5 bool = true
    var v17 float32 = ((-2.25 / 3.14159) * (0.1 * ((-2.25 / 0.1) - -2.25)))
    _ = len(v3)
    switch 42 {
    case 32768:
        fmt.Println(v17, v5, v17)
    default:
        var v18 uint16 = 65535
        _ = v18
    }
    v17 += float32((1.5 / 1.0))
    v19 := GenS1{F0: 65535}
    v2 += 0
    for fi20 := 0; fi20 < 16; fi20++ {
        var v21 bool = false
        v22 := GenS1{F0: 655}
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 uint16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
