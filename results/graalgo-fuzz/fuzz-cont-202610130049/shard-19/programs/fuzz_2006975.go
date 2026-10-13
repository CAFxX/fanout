package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
func genFn0(p0 int64, p1 uint8) uint8 {
    switch -395 {
    case 128:
        p1 -= (p1 + p1)
    case 666:
        p1 += p1
    case 830:
        p1 = p1
    default:
        var v5 uint = 1
        _ = v5
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        switch "hello" {
        case "foo bar":
            if (!(!false)) {
                if false {
                    for fi7 := 0; fi7 < 8; fi7++ {
                        fmt.Println(p0, p0, p0)
                        fmt.Println(p1)
                        _ = fi7
                    }
                    v8 := []int{}
                    _ = v8
                } else {
                    p0 = p0
                    for fi9 := 0; fi9 < 7; fi9++ {
                        p0 += p0
                        v10 := []int{(^(fi6)), (^(fi9)), fi6, (fi6 & fi9)}
                        _ = fi9
                        _ = v10
                    }
                    if false {
                        v11 := []int{fi6}
                        _ = v11
                    } else {
                        if ((!false) && true) {
                            v12 := []int{fi6, fi6, fi6}
                            _ = v12
                        } else {
                            var v13 float64 = 1.5
                            v13 = v13
                            v14 := []int{(fi6 % ((42) | 1)), -689}
                            _ = v13
                            _ = v14
                        }
                    }
                }
                v15 := []int{fi6}
                _ = v15
            }
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    return 128
}
func genFn1(p0 uint16, p1 uint32, p2 int32) int {
    p1 -= (p1 >> ((uint(p0)) & 63))
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return 128
}

func main() {
    var v1 uint32 = 21
    v2 := make(map[uint32]string)
    v3 := make(map[string]bool)
    var v4 string = ("\n\t" + ("foo bar" + "a"))
    var v16 uint32 = (v1 ^ v1)
    var v17 uint64 = 42
    var v18 uint16 = 425
    v4 = ((v4 + (v4 + ((("a" + (("hello" + (("a" + "\n\t") + "\n\t")) + "")) + v4) + (v4 + ((v4 + "foo bar") + v4))))) + v4)
    if true {
        if ((true && (!(-1 == -45))) || true) {
            var v19 float64 = float64(float32(0.0))
            _ = v19
        }
    }
    var v20 int8 = 1
    var v21 int16 = 32767
    v4 = v4
    v4 = v4
    v18 = uint16(v1)
    if true {
        v1 -= (v16 % ((2147483647) | 1))
        var v22 int16 = v21
        _ = v22
    } else {
        var v23 int8 = (v20 + v20)
        if true {
            var v24 int = -1
            v24 += (v24 * (^(v24)))
            _ = v24
        } else {
            var v25 int8 = (v23 << ((165) & 63))
            v26 := GenS1{F0: 18446744073709551615, F1: v16}
            var v27 float64 = 0.0
            _ = v25
            _ = v26
            _ = v27
        }
        _ = v23
    }
    fmt.Println(v16)
    fmt.Println(v4, v16, v21)
    var v28 int32 = 511
    var v29 int8 = 94
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
