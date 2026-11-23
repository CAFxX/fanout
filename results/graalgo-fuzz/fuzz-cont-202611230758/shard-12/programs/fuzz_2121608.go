package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint8 }
func genFn0(p0 uint32, p1 int) uint64 {
    v5 := GenS1{F0: 145}
    p0 = p0
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint64 = 0
        var v8 int16 = 0
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    return 2
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]int)
    var v4 bool = (-32768 < -369)
    var v9 float64 = 1e10
    if false {
        if false {
            if v4 {
                var v10 int = 2147483647
                v11 := GenS1{F0: 18446744073709551615}
                _ = v10
                _ = v11
            }
            var v12 uint8 = 2
            _ = v12
        } else {
            for fi13 := 0; fi13 < 17; fi13++ {
                var v14 uint32 = 65535
                _ = fi13
                _ = v14
            }
            if (v9 >= v9) {
                var v15 bool = false
                v16 := GenS1{F0: 2}
                v17 := GenS1{F0: 458}
                _ = v15
                _ = v16
                _ = v17
            }
        }
        var v18 int16 = 32767
        v19 := []int{0, 0}
        _ = v18
        _ = v19
    }
    v20 := []int{4294967295, 127, 2147483647, -602}
    v20 = append(v20, 0)
    var v21 uint8 = 152
    v4 = (!v4)
    var v22 int8 = -36
    var v23 int64 = -9223372036854775808
    var v24 bool = true
    if v4 {
        var v25 uint16 = 2
        v20 = append(v20, 128)
        switch "" {
        case "foo bar":
            fmt.Println(v4, v24)
        case "":
            if true {
                switch 2147483647 {
                case -275:
                    fmt.Println(v25, v25)
                default:
                    switch -755 {
                    case 2147483647:
                        v26 := GenS2{F0: v22, F1: 2, F2: (v21 * v21)}
                        _ = v26
                    }
                }
            }
        case "a":
            fmt.Println(v4, v9, v24)
        }
        _ = v25
    } else {
        if (1 <= 1) {
            for fi27 := 0; fi27 < 13; fi27++ {
                if (!(uint32(v21) >= 2)) {
                    var v28 bool = (2 == 2)
                    _ = v28
                }
                var v29 float64 = -2.25
                _ = fi27
                _ = v29
            }
        }
        v22 = v22
    }
    v21 += v21
    v21 = (v21 % ((132) | 1))
    v20 = append(v20, 32768)
    var v30 uint = 486
    if v4 {
        var v31 int16 = 42
        _ = v31
    } else {
        v23 += v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v9
}
