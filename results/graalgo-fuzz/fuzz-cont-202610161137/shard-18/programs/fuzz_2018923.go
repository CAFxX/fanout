package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int32, p1 int32) uint8 {
    p0 += p1
    var v6 int64 = -494
    switch -445 {
    case 761:
        p1 -= p1
    case -37:
        for fi7 := 0; fi7 < 18; fi7++ {
            fmt.Println(fi7)
            p0 -= 689
            _ = fi7
        }
    default:
        fmt.Println(v6)
    }
    fmt.Println(p0, p0)
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: 2147483647}
        var v10 uint64 = 1
        switch fi8 {
        case -1:
            var v11 int64 = 1
            _ = v11
        default:
            v12 := []int{(fi8 | (-(fi8))), fi8, (fi8 - fi8)}
            _ = v12
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    return 94
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = 9223372036854775807
    v3 := make(map[string]uint32)
    v4 := []int{2147483647}
    v5 := make(map[uint32]string)
    if false {
        for fi13 := 0; fi13 < 20; fi13++ {
            switch fi13 {
            case 0:
                if false {
                    v14 := GenS1{F0: 65535}
                    _ = v14
                } else {
                    if (!false) {
                        v15 := GenS1{F0: 4294967295}
                        for fi16 := 0; fi16 < 11; fi16++ {
                            v2 += 9223372036854775807
                            v2 -= (fi13 >> ((573) & 63))
                            _ = fi16
                        }
                        _ = v15
                    } else {
                        var v17 string = ""
                        v2 = fi13
                        _ = v17
                    }
                }
            default:
                delete(v5, 2147483648)
            }
            _ = fi13
        }
        v2 += v2
    }
    v18 := GenS2{F0: ""}
    for fi19 := 0; fi19 < 14; fi19++ {
        if false {
            v20 := GenS3{F0: false, F1: (fi19 << ((4294967295) & 63))}
            fmt.Println(fi19, fi19)
            for fi21 := 0; fi21 < 11; fi21++ {
                var v22 int = (fi19 | (fi19 & -1))
                _ = fi21
                _ = v22
            }
            _ = v20
        } else {
            v2 = (fi19 << ((774) & 63))
        }
        v2 = 128
        _ = len(v3)
        _ = fi19
    }
    if false {
        var v23 uint = 971
        _ = v23
    }
    v2 = (v2 * v2)
    _ = len(v5)
    switch 355 {
    case 32768:
        v24 := GenS1{F0: 404}
        _ = v24
    }
    v25 := GenS4{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
