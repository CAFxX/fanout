package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int32; F1 int16; F2 int8 }
func genFn0() uint {
    var v5 int64 = -373
    if false {
        v6 := []int{0}
        _ = v6
    }
    v5 = (v5 | v5)
    if true {
        for fi7 := 0; fi7 < 3; fi7++ {
            v5 += (v5 << ((2) & 63))
            v8 := []int{}
            _ = fi7
            _ = v8
        }
        v9 := []int{732, 726, 569}
        _ = v9
    }
    _ = v5
    return 695
}
func genFn1(p0 int32, p1 uint64, p2 int) uint32 {
    p1 -= (p1 - 321)
    p2 = int(p0)
    for fi10 := 0; fi10 < 15; fi10++ {
        for fi11 := 0; fi11 < 3; fi11++ {
            v12 := []int{(p2 + fi11), (fi10 - fi11)}
            var v13 uint = uint(p1)
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    if false {
        fmt.Println(p0, p1)
        var v14 uint = 168
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    var v3 float32 = 0.1
    v4 := make(map[string]string)
    var v15 int32 = -688
    var v16 uint16 = 322
    if false {
        for fi17 := 0; fi17 < 19; fi17++ {
            v18 := []int{32767, (fi17 | 32767), (fi17 % ((-365) | 1))}
            fmt.Println(v16, v3, v15)
            v1["hello"] = fi17
            _ = fi17
            _ = v18
        }
        v15 = (-(v15))
    }
    fmt.Println(v16)
    fmt.Println(v3)
    if true {
        v16 = (v16 << ((genFn0()) & 63))
        v16 = (v16 >> ((1) & 63))
        switch "\n\t" {
        case "\n\t":
            fmt.Println(v16, v16, v3)
        }
    } else {
        switch 1 {
        case 127:
            v16 -= (v16 / (((^(v16))) | 1))
        case 1:
            var v19 int = -905
            _ = v19
        case 0:
            v20 := GenS1{F0: 4294967295, F1: 128, F2: genFn1((v15 >> ((355) & 63)), 4294967295, 127)}
            _ = v20
        default:
            var v21 int16 = -733
            _ = v21
        }
    }
    v22 := []int{32767}
    for fi23 := 0; fi23 < 9; fi23++ {
        v16 -= (^(v16))
        if false {
            if false {
                var v24 int64 = 42
                v25 := GenS2{F0: 470, F1: 42, F2: -32}
                _ = v24
                _ = v25
            }
        } else {
            if false {
                if true {
                    if (!false) {
                        var v26 int32 = (^(v15))
                        _ = v26
                    }
                }
            }
            switch "a" {
            case "foo bar":
                for fi27 := 0; fi27 < 1; fi27++ {
                    v22 = append(v22, 42)
                    _ = fi27
                }
            }
            fmt.Println(fi23, v3)
        }
        v16 = v16
        _ = fi23
    }
    var v28 uint = 65535
    var v29 int64 = 42
    v22 = append(v22, -320)
    for fi30 := 0; fi30 < 4; fi30++ {
        v31 := GenS1{F0: 2147483647, F1: 128, F2: 65535}
        v29 -= (v29 & (^(v29)))
        _ = fi30
        _ = v31
    }
    var v32 int64 = (v29 ^ 0)
    v29 -= 434
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
}
