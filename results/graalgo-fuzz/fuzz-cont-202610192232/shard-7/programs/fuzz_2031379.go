package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint16 = 255
    var v2 uint8 = 126
    v3 := make(map[int]uint32)
    v4 := []int{-762, -2147483648}
    var v5 int64 = 733
    var v6 int8 = 65
    if (!true) {
        if (true && (325 >= uint32(v1))) {
            v7 := GenS1{F0: (v5 >> ((1) & 63)), F1: (-(v6))}
            _ = v7
        } else {
            var v8 int8 = int8(v1)
            v4 = append(v4, int(v2))
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        if false {
            var v10 uint16 = v1
            for fi11 := 0; fi11 < 9; fi11++ {
                if (0 > 47) {
                    var v12 int64 = (-(v5))
                    v5 = (v5 >> ((2) & 63))
                    _ = v12
                } else {
                    if true {
                        var v13 int64 = v5
                        _ = v13
                    }
                    v1 += (v10 % ((32767) | 1))
                }
                if true {
                    switch (fi9 >> ((4294967295) & 63)) {
                    case 547:
                        v14 := GenS2{F0: 439}
                        var v15 int32 = 1
                        _ = v14
                        _ = v15
                    case 128:
                        var v16 int64 = v5
                        _ = v16
                    default:
                        delete(v3, (^(fi11)))
                    }
                    var v17 uint16 = (v10 * v1)
                    v10 = v17
                    _ = v17
                }
                var v18 int = fi11
                _ = fi11
                _ = v18
            }
            _ = v10
        }
        _ = fi9
    }
    switch -2147483648 {
    case 382:
        _ = len(v4)
    case -905:
        if false {
            switch "a" {
            case "foo bar":
                v3[116] = 208
            default:
                v2 -= v2
            }
            var v19 int16 = 822
            var v20 bool = true
            _ = v19
            _ = v20
        } else {
            var v21 int = 127
            for fi22 := 0; fi22 < 12; fi22++ {
                _ = len(v3)
                v21 += (v21 - fi22)
                _ = fi22
            }
            var v23 string = "a"
            _ = v21
            _ = v23
        }
    case 42:
        v2 -= 42
    default:
        var v24 float32 = -2.25
        _ = v24
    }
    var v25 uint32 = 42
    var v26 int8 = -128
    fmt.Println(v1, v25)
    for fi27 := 0; fi27 < 12; fi27++ {
        var v28 uint8 = 101
        _ = fi27
        _ = v28
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v6 = v6
    switch -618 {
    case 2147483647:
        v29 := GenS1{F0: v5, F1: v26}
        _ = v29
    default:
        v1 = uint16(v2)
    }
    fmt.Println(v2)
    var v30 string = ""
    v31 := GenS1{F0: int64(v2), F1: -128}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
