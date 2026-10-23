package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 uint32) uint64 {
    var v7 int16 = -891
    p0 = p0
    if ("a" < "a") {
        p0 = (p0 / (((p0 * p0)) | 1))
    }
    v7 -= 0
    _ = p0
    _ = v7
    return 473
}
func genFn1(p0 int64, p1 int64, p2 uint32) int32 {
    switch "foo bar" {
    case "":
        var v8 uint16 = 2
        _ = v8
    case "\n\t":
        var v9 uint8 = 183
        _ = v9
    default:
        v10 := GenS1{F0: 0}
        _ = v10
    }
    var v11 int8 = 30
    var v12 int32 = 42
    var v13 int = 42
    v14 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return (v12 & 2147483647)
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = ("hello" + "")
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]bool)
    var v5 int8 = 1
    v6 := make(map[uint32]string)
    var v15 uint16 = 0
    fmt.Println(v15)
    fmt.Println(v5, v2, v2)
    var v16 uint64 = 879
    v17 := []int{-1, -629, 928, 659}
    if (1 > 0) {
        switch 2147483647 {
        case 27:
            for fi18 := 0; fi18 < 8; fi18++ {
                fmt.Println(fi18, fi18)
                _ = fi18
            }
        case 2147483647:
            if true {
                v15 += (v15 + uint16(v16))
            }
        default:
            for fi19 := 0; fi19 < 9; fi19++ {
                for fi20 := 0; fi20 < 18; fi20++ {
                    v21 := v1[fi20]
                    _ = fi20
                    _ = v21
                }
                _ = fi19
            }
        }
        v16 = v16
        v22 := GenS1{F0: 2}
        _ = v22
    }
    v2 = v2
    fmt.Println(v2, v15)
    if true {
        v2 = v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
