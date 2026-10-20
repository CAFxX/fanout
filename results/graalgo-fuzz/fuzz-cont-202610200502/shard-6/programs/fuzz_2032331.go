package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
func genFn0(p0 bool, p1 int32) uint16 {
    if p0 {
        var v5 int16 = 1
        _ = v5
    }
    v6 := []int{}
    _ = p0
    _ = p1
    _ = v6
    return 32767
}
func genFn1() uint32 {
    if false {
        var v7 uint8 = 127
        var v8 float32 = 1.5
        var v9 uint64 = 42
        _ = v7
        _ = v8
        _ = v9
    } else {
        var v10 uint8 = 168
        _ = v10
    }
    if (true && false) {
        switch 32767 {
        case 4294967295:
            v11 := GenS1{F0: 2, F1: 32767}
            _ = v11
        case 916:
            var v12 uint64 = 2
            _ = v12
        default:
            v13 := []int{499}
            _ = v13
        }
        for fi14 := 0; fi14 < 2; fi14++ {
            if false {
                v15 := GenS1{F0: 127, F1: -1}
                var v16 uint8 = 42
                _ = v15
                _ = v16
            } else {
                if ((true || true) && true) {
                    var v17 uint = 42
                    _ = v17
                }
                var v18 uint16 = 0
                v19 := []int{fi14, 2147483647}
                _ = v18
                _ = v19
            }
            _ = fi14
        }
    }
    if (0.0 != 1.5) {
        var v20 int32 = 0
        fmt.Println(v20)
        v21 := GenS1{F0: 255, F1: 0}
        _ = v20
        _ = v21
    } else {
        var v22 uint64 = 42
        switch 42 {
        case 42:
            for fi23 := 0; fi23 < 7; fi23++ {
                v24 := []int{fi23}
                var v25 string = ("hello" + ("" + (("a" + "a") + "")))
                fmt.Println(fi23)
                _ = fi23
                _ = v24
                _ = v25
            }
        case -746:
            v22 += v22
        case -9223372036854775808:
            for fi26 := 0; fi26 < 18; fi26++ {
                v27 := []int{-176, 32768, fi26}
                if len(v27) > 0 {
                    v27[0] = (^(fi26))
                }
                _ = fi26
                _ = v27
            }
        default:
            var v28 uint16 = 2
            _ = v28
        }
        _ = v22
    }
    return 204
}

func main() {
    var v1 float64 = 1.5
    var v2 bool = (2 != 42)
    v3 := make(map[uint32]bool)
    var v4 string = ""
    if (1 == 0) {
        var v29 int16 = 32767
        for fi30 := 0; fi30 < 5; fi30++ {
            v2 = (v2 || (887 > 605))
            fmt.Println(v4, v4)
            v4 = (v4 + "foo bar")
            _ = fi30
        }
        v3[2147483647] = false
        _ = v29
    }
    var v31 uint32 = 0
    v1 = 3.14159
    var v32 int = 127
    v33 := []int{-9223372036854775808, v32}
    fmt.Println(v31, v4)
    if v2 {
        v3[v31] = (42 != 982)
        v32 -= v32
        fmt.Println(v4, v1)
    } else {
        v34 := v3[(v31 * (v31 % (((^(v31))) | 1)))]
        v2 = (false || ((((float32(v1) * -2.25) - 0.0) * 1e10) == -2.25))
        v31 = v31
        _ = v34
    }
    var v35 bool = ((!(2147483647 < (^(v31)))) || v2)
    switch v32 {
    case -631:
        v36 := GenS1{F0: 128, F1: 0}
        _ = v36
    case -2147483648:
        var v37 int = (-(v32))
        _ = v37
    case -733:
        delete(v3, (v31 & genFn1()))
    default:
        var v38 int = (v32 ^ v32)
        _ = v38
    }
    for fi39 := 0; fi39 < 14; fi39++ {
        v2 = (!((!v35) && false))
        _ = fi39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v35
    _ = v4
}
