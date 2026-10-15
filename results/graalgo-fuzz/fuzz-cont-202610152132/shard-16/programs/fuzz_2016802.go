package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS2 struct { F0 int16; F1 int16 }
func genFn0(p0 bool, p1 uint32) int8 {
    var v4 float64 = 1e10
    var v5 float64 = (0.0 / v4)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 68
}
func genFn1() uint32 {
    v6 := GenS1{F0: 159, F1: 972, F2: -1}
    var v7 int16 = 32767
    var v8 int64 = 779
    var v9 bool = (0.1 <= float32(1.5))
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 bool = (!false)
    v2 := make(map[uint32]string)
    var v3 uint32 = 42
    v10 := []int{0, -9223372036854775808}
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 int64 = 322
        for fi13 := 0; fi13 < 17; fi13++ {
            _ = len(v2)
            if (!(-2147483648 == -589)) {
                if (-1 > genFn0((127 >= 0), v3)) {
                    if false {
                        var v14 int16 = 47
                        if v1 {
                            switch ((("\n\t" + ("hello" + "hello")) + (("foo bar" + ("" + ("foo bar" + "foo bar"))) + "foo bar")) + "a") {
                            case "a":
                                switch 341 {
                                case 127:
                                    v10 = append(v10, 1)
                                    v15 := GenS2{F0: v14, F1: v14}
                                    v1 = v1
                                    _ = v15
                                case 2147483647:
                                    v12 = (v12 ^ (v12 & v12))
                                default:
                                    v1 = (!(v1 || (true && false)))
                                }
                            case "\n\t":
                                v3 -= 696
                            }
                            var v16 int32 = 1
                            var v17 uint32 = genFn1()
                            _ = v16
                            _ = v17
                        }
                        fmt.Println(v12)
                        _ = v14
                    } else {
                        if (0 > 148) {
                            v18 := GenS1{F0: 870, F1: 32767, F2: 100}
                            _ = v18
                        } else {
                            v3 += 2147483647
                        }
                        v12 -= (v12 ^ v12)
                    }
                    switch 0 {
                    case 9223372036854775807:
                        fmt.Println(fi13, fi13, fi13)
                    case 32768:
                        var v19 uint = 255
                        _ = v19
                    }
                } else {
                    fmt.Println(fi11, fi13)
                    var v20 bool = (((fi11 * (fi13 | fi13)) >= fi13) && v1)
                    _ = v20
                }
                var v21 uint32 = (v3 >> ((2) & 63))
                _ = v21
            } else {
                fmt.Println(fi11)
            }
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(v1)
    var v22 int32 = -750
    for fi23 := 0; fi23 < 8; fi23++ {
        if (475 == 4294967295) {
            _ = len(v2)
            v3 = (v3 >> ((1) & 63))
        }
        var v24 int8 = 42
        _ = fi23
        _ = v24
    }
    fmt.Println(v22, v22)
    v1 = (415 == 828)
    v25 := GenS1{F0: 266, F1: 0, F2: 1}
    if v1 {
        v3 += v3
    } else {
        v26 := GenS1{F0: 255, F1: 32767, F2: genFn0((!v1), v3)}
        v27 := GenS2{F0: -44, F1: -657}
        _ = v26
        _ = v27
    }
    var v28 bool = (!(!(0 <= -32768)))
    v28 = true
    var v29 string = ("foo bar" + ("" + ("foo bar" + "")))
    v30 := v2[v3]
    v22 = (v22 + (v22 | 1))
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    v29 = v30
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
