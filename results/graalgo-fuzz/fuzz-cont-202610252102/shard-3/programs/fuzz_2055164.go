package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 uint64; F2 bool }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 int8; F1 bool }
func genFn0(p0 string, p1 uint16, p2 int8) uint32 {
    v5 := GenS1{F0: 726}
    if (!(true && false)) {
        v6 := GenS1{F0: 18446744073709551615}
        p1 += p1
        _ = v6
    } else {
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := []int{fi7}
            for fi9 := 0; fi9 < 7; fi9++ {
                if (!(false || true)) {
                    switch "" {
                    case "":
                        p1 -= 255
                    }
                    if ((1e10 > (3.14159 + float32(0.0))) && false) {
                        _ = len(v8)
                    }
                }
                if true {
                    var v10 int = fi7
                    p2 -= -128
                    v11 := GenS2{F0: (fi9 - v10), F1: 4294967295, F2: (!(0 < 0))}
                    _ = v10
                    _ = v11
                }
                _ = fi9
            }
            _ = fi7
            _ = v8
        }
    }
    var v12 string = ((p0 + "\n\t") + p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v5
    return 372
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -2147483648
    v3 := []string{"a", "foo bar"}
    v4 := make(map[int]uint32)
    v2 = -2147483648
    v2 = (v2 & v2)
    if false {
        for fi13 := 0; fi13 < 14; fi13++ {
            switch (fi13 - (fi13 % (((fi13 / ((fi13) | 1))) | 1))) {
            case 0:
                v2 = 42
            case 1:
                if true {
                    if (!true) {
                        switch ("a" + "hello") {
                        case "a":
                            var v14 uint16 = 32767
                            _ = v14
                        default:
                            var v15 bool = true
                            _ = v15
                        }
                        v2 = v2
                        var v16 uint64 = 468
                        _ = v16
                    } else {
                        v2 -= (v2 * v2)
                        v17 := GenS3{F0: -1, F1: 18446744073709551615}
                        v2 += -2147483648
                        _ = v17
                    }
                } else {
                    var v18 bool = (4294967295 < genFn0(("hello" + "hello"), 255, 11))
                    for fi19 := 0; fi19 < 13; fi19++ {
                        switch "foo bar" {
                        case "a":
                            v18 = v18
                        case "":
                            var v20 int64 = 42
                            _ = v20
                        case "foo bar":
                            var v21 int8 = -1
                            _ = v21
                        }
                        _ = fi19
                    }
                    var v22 uint = 1
                    _ = v18
                    _ = v22
                }
            case 127:
                if (!false) {
                    for fi23 := 0; fi23 < 11; fi23++ {
                        fmt.Println(fi23)
                        _ = fi23
                    }
                    if len(v3) > 0 {
                        v3[0] = ("\n\t" + ("" + "hello"))
                    }
                    _ = len(v1)
                }
            }
            var v24 uint = 2
            _ = fi13
            _ = v24
        }
        var v25 int32 = v2
        _ = v25
    }
    v26 := v1[4294967295]
    v27 := GenS4{F0: 42, F1: (-1 <= 0)}
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    switch "" {
    case "a":
        delete(v4, -374)
    }
    v28 := v4[42]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
