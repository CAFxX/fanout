package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint64; F1 int64; F2 uint16 }
func genFn0(p0 uint8, p1 uint64) uint8 {
    if true {
        for fi5 := 0; fi5 < 9; fi5++ {
            fmt.Println(fi5, p1)
            _ = fi5
        }
    } else {
        p0 -= 127
        switch 553 {
        case 4294967295:
            fmt.Println(p0)
        default:
            var v6 float64 = -2.25
            _ = v6
        }
        p1 += 365
    }
    p1 = 2
    var v7 bool = false
    _ = p0
    _ = p1
    _ = v7
    return (^(p0))
}
func genFn1(p0 uint64, p1 int8, p2 uint16) uint {
    fmt.Println(p2)
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := []int{(fi8 >> ((2) & 63)), fi8, (fi8 + 32768)}
        p1 += 0
        _ = fi8
        _ = v9
    }
    var v10 bool = true
    v11 := []int{32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 1
}

func main() {
    var v1 uint16 = 959
    v2 := make(map[uint32]bool)
    var v3 string = ("foo bar" + "")
    v4 := make(map[uint32]uint32)
    if (!(9223372036854775807 >= 42)) {
        v12 := []int{-2147483648, 42}
        for fi13 := 0; fi13 < 11; fi13++ {
            delete(v2, 2147483647)
            _ = fi13
        }
        var v14 uint32 = 4294967295
        _ = v12
        _ = v14
    }
    v15 := GenS1{F0: 255, F1: 65535}
    var v16 uint32 = 1
    v16 += 42
    if true {
        v3 = v3
    }
    fmt.Println(v1, v1)
    v17 := GenS1{F0: genFn1(uint64(v16), 127, v1), F1: 1}
    v16 = (v16 ^ v16)
    v18 := GenS1{F0: 2, F1: 4294967295}
    for fi19 := 0; fi19 < 7; fi19++ {
        for fi20 := 0; fi20 < 4; fi20++ {
            v16 += 2
            if false {
                if true {
                    v3 = ""
                    v1 -= v1
                    v21 := GenS1{F0: 0, F1: 1}
                    _ = v21
                } else {
                    v22 := GenS1{F0: 953, F1: 660}
                    _ = v22
                }
                var v23 bool = true
                _ = v23
            }
            _ = fi20
        }
        v24 := GenS2{F0: 665, F1: 0, F2: (v1 | (^(v1)))}
        _ = fi19
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
