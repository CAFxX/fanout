package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
func genFn0(p0 int8, p1 string, p2 uint32) uint32 {
    p1 = p1
    if ((^(p2)) != 2147483648) {
        p2 = 0
        for fi6 := 0; fi6 < 12; fi6++ {
            for fi7 := 0; fi7 < 3; fi7++ {
                fmt.Println(fi7)
                _ = fi7
            }
            _ = fi6
        }
        if false {
            fmt.Println(p2, p2, p1)
            for fi8 := 0; fi8 < 18; fi8++ {
                fmt.Println(p1, p1)
                _ = fi8
            }
            var v9 uint32 = p2
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := []int{fi10, fi10, (fi10 / ((4294967295) | 1))}
        var v12 uint64 = 42
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return 432
}

func main() {
    var v1 int16 = -1
    v2 := []int{-9223372036854775808, 244, 2147483647}
    v3 := make(map[int]int)
    var v4 bool = (698 != 42)
    var v5 float64 = 0.1
    v4 = v4
    if ("" != "") {
        delete(v3, -2147483648)
    } else {
        switch 0 {
        case 604:
            for fi13 := 0; fi13 < 1; fi13++ {
                for fi14 := 0; fi14 < 17; fi14++ {
                    if v4 {
                        _ = len(v3)
                        v5 = v5
                    } else {
                        var v15 int8 = 100
                        v2 = append(v2, 9223372036854775807)
                        fmt.Println(v15)
                        _ = v15
                    }
                    v16 := GenS1{F0: 0, F1: 42}
                    fmt.Println(fi14, fi14)
                    _ = fi14
                    _ = v16
                }
                v4 = ((1e10 / 0.1) <= -2.25)
                _ = fi13
            }
        case -949:
            fmt.Println(v4)
        default:
            var v17 int = 128
            _ = v17
        }
        var v18 string = (("a" + "a") + ("foo bar" + ("a" + (("foo bar" + (("foo bar" + "\n\t") + ("hello" + ("\n\t" + "foo bar")))) + ("hello" + ((("\n\t" + ("hello" + ("" + "\n\t"))) + "a") + ""))))))
        _ = v18
    }
    var v19 int32 = -442
    var v20 uint = 2
    v21 := GenS1{F0: 128, F1: 0}
    delete(v3, 128)
    var v22 uint = (v20 | (v20 ^ v20))
    v1 = (v1 + v1)
    for fi23 := 0; fi23 < 6; fi23++ {
        var v24 int64 = 706
        for fi25 := 0; fi25 < 5; fi25++ {
            if (1 < 255) {
                var v26 bool = v4
                var v27 int32 = v19
                v22 -= 18446744073709551615
                _ = v26
                _ = v27
            } else {
                v2 = append(v2, (fi23 * (fi23 << ((v22) & 63))))
                if len(v2) > 0 {
                    v2[0] = fi23
                }
                fmt.Println(fi25, v19, v22)
                var v28 string = "\n\t"
                _ = v28
            }
            var v29 bool = true
            _ = fi25
            _ = v29
        }
        v30 := GenS1{F0: 36, F1: (v24 + -1)}
        _ = fi23
        _ = v24
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
