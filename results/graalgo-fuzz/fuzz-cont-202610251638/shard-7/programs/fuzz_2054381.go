package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
func genFn0(p0 uint) int32 {
    if true {
        if (255 != 0) {
            if true {
                if (!false) {
                    if false {
                        fmt.Println(p0)
                        v5 := GenS1{F0: 402, F1: 1}
                        p0 -= p0
                        _ = v5
                    }
                    v6 := []int{}
                    _ = v6
                }
                for fi7 := 0; fi7 < 7; fi7++ {
                    if true {
                        p0 = (p0 << (((p0 ^ (p0 & (p0 * (p0 & p0))))) & 63))
                    }
                    _ = fi7
                }
                p0 = p0
            } else {
                if true {
                    v8 := []int{-106, 691, 32767}
                    _ = v8
                }
                v9 := []int{2147483647, 854, 917, 4294967295}
                _ = v9
            }
        }
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := []int{fi10, -307, -611, 216}
            p0 = 0
            p0 = p0
            _ = fi10
            _ = v11
        }
        p0 = p0
    } else {
        v12 := []int{}
        var v13 uint64 = 0
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 uint16 = 65535
        _ = fi14
        _ = v15
    }
    _ = p0
    return -982
}
func genFn1(p0 bool, p1 int, p2 int64) uint32 {
    v16 := []int{(p1 & 9223372036854775807), p1, p1, 32767}
    switch p1 {
    case 2147483647:
        switch "\n\t" {
        case "":
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 407
}

func main() {
    var v1 uint64 = 0
    var v2 bool = true
    v3 := []uint32{406}
    var v4 bool = (1000 > -732)
    if v2 {
        var v17 float32 = 3.14159
        _ = v17
    }
    var v18 uint = 735
    for fi19 := 0; fi19 < 8; fi19++ {
        if v2 {
            fmt.Println(fi19, v4, v4)
            if (42 > uint32(v1)) {
                if (42 == 1) {
                    v18 = v18
                } else {
                    var v20 bool = v2
                    fmt.Println(v1, v4)
                    v21 := GenS1{F0: uint32(v18), F1: 1}
                    _ = v20
                    _ = v21
                }
            }
        }
        v4 = v2
        for fi22 := 0; fi22 < 6; fi22++ {
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
            _ = fi22
        }
        _ = fi19
    }
    var v23 uint16 = 2
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v1)
    if len(v3) > 0 {
        v3[0] = 2
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v18, v1, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
