package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := []int{-2147483648, -2147483648, -30, 620}
    v2 := make(map[int]string)
    var v3 string = "foo bar"
    var v4 uint32 = 690
    v5 := make(map[uint32]string)
    var v6 uint32 = 1
    delete(v2, 357)
    var v7 float64 = (1.5 * (-2.25 - 0.0))
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 5; fi9++ {
            if false {
                switch fi8 {
                case -439:
                    v6 = v6
                case -713:
                    v3 = v3
                case -2147483648:
                    var v10 uint = 51
                    _ = v10
                default:
                    var v11 bool = (!true)
                    _ = v11
                }
                v12 := GenS1{F0: 802}
                v6 -= v4
                _ = v12
            }
            var v13 int16 = 32767
            var v14 uint32 = (v4 & v6)
            _ = fi9
            _ = v13
            _ = v14
        }
        _ = fi8
    }
    if (!false) {
        var v15 string = v3
        for fi16 := 0; fi16 < 14; fi16++ {
            v17 := GenS1{F0: 2}
            var v18 uint32 = (v6 + (v6 - 0))
            if (1 == 23) {
                var v19 int8 = 85
                fmt.Println(v19)
                v20 := GenS1{F0: 2}
                _ = v19
                _ = v20
            } else {
                var v21 int = -498
                _ = v21
            }
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = v15
    }
    switch -338 {
    case -843:
        v6 = v6
    case 1:
        var v22 int16 = 42
        _ = v22
    case -2147483648:
        v23 := GenS1{F0: 42}
        _ = v23
    default:
        switch -2147483648 {
        case 0:
            var v24 int16 = 1
            _ = v24
        case -9223372036854775808:
            var v25 uint = 900
            _ = v25
        default:
            fmt.Println(v7, v3)
        }
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        if len(v1) > 0 {
            v1[0] = fi26
        }
        if false {
            v27 := GenS1{F0: 465}
            for fi28 := 0; fi28 < 2; fi28++ {
                delete(v2, (fi26 + 0))
                _ = fi28
            }
            _ = v27
        } else {
            fmt.Println(v7, fi26)
            delete(v2, (fi26 - fi26))
        }
        _ = fi26
    }
    var v29 uint32 = v6
    var v30 float32 = 0.0
    if ((-982 > 1) || (!false)) {
        v30 = v30
        if len(v1) > 0 {
            v1[0] = -2147483648
        }
        if len(v1) > 0 {
            v1[0] = 0
        }
    }
    if (0 < -90) {
        v31 := GenS1{F0: 42}
        var v32 float64 = float64(v30)
        _ = v31
        _ = v32
    }
    v30 = (v30 * (-2.25 * v30))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
