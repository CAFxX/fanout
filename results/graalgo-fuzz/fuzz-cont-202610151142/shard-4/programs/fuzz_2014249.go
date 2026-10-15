package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 string }
type GenS2 struct { F0 string; F1 int64 }
func genFn0(p0 string) uint64 {
    var v5 int32 = -638
    fmt.Println(v5, v5)
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 int8 = 1
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := []int{fi8, fi8, 2147483647, (fi8 ^ (fi8 - 128))}
        _ = fi8
        _ = v9
    }
    var v10 float32 = 0.0
    _ = p0
    _ = v10
    _ = v5
    return 0
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    var v3 int64 = 0
    var v4 string = ("a" + "\n\t")
    v3 -= (v3 << ((162) & 63))
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := GenS1{F0: 127, F1: 42, F2: "hello"}
        if (!(!(false || true))) {
            var v13 bool = true
            v14 := []int{fi11, fi11, 0}
            _ = v13
            _ = v14
        } else {
            var v15 uint = 65535
            if (false || true) {
                v3 = (^(v3))
                for fi16 := 0; fi16 < 3; fi16++ {
                    var v17 int16 = 1000
                    _ = fi16
                    _ = v17
                }
            }
            _ = v15
        }
        var v18 uint8 = 1
        _ = fi11
        _ = v12
        _ = v18
    }
    if (95 >= 255) {
        delete(v2, 65535)
        if true {
            v4 = v4
            switch -9223372036854775808 {
            case 303:
                delete(v1, 1)
            case 4294967295:
                v19 := v1[387]
                _ = v19
            default:
                var v20 float32 = (-2.25 / 1.0)
                _ = v20
            }
        }
        v21 := GenS1{F0: -1, F1: 4294967295, F2: ((((("" + (v4 + "hello")) + ("" + ("" + (((v4 + v4) + "") + v4)))) + (v4 + v4)) + v4) + v4)}
        _ = v21
    }
    fmt.Println(v3, v4)
    switch -2147483648 {
    case 745:
        v22 := GenS1{F0: 42, F1: 2, F2: v4}
        _ = v22
    case -92:
        v4 = (v4 + v4)
    case -1:
        if true {
            switch 1 {
            case 127:
                v23 := []int{}
                _ = v23
            }
        } else {
            v24 := v1[-9223372036854775808]
            _ = v24
        }
    default:
        v25 := GenS1{F0: 6, F1: 77, F2: v4}
        _ = v25
    }
    for fi26 := 0; fi26 < 2; fi26++ {
        v27 := []int{-838, (fi26 | 346), fi26, -620}
        _ = fi26
        _ = v27
    }
    var v28 uint32 = 547
    fmt.Println(v4)
    v3 = (-(v3))
    v29 := GenS2{F0: "hello", F1: v3}
    var v30 uint32 = v28
    fmt.Println(v3, v30)
    fmt.Println(v28, v28)
    for fi31 := 0; fi31 < 15; fi31++ {
        fmt.Println(v4)
        v30 -= v28
        _ = fi31
    }
    v32 := []int{-577}
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
}
