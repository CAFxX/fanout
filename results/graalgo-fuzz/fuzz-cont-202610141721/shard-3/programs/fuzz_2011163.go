package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int8 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint }
func genFn0(p0 uint, p1 int8, p2 uint64) int8 {
    v6 := []int{}
    fmt.Println(p2)
    v7 := GenS1{F0: 1, F1: (0.0 > 0.1), F2: -113}
    switch -1 {
    case 963:
        v8 := GenS1{F0: -2147483648, F1: true, F2: p1}
        _ = v8
    case 4294967295:
        if (2147483647 >= 2147483647) {
            p2 = p2
        }
    default:
        v9 := GenS2{F0: 2, F1: 1, F2: p0}
        _ = v9
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        p2 = (p2 >> ((p0) & 63))
        v11 := GenS1{F0: 2147483647, F1: false, F2: p1}
        p0 += p0
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (p1 << ((80) & 63))
}
func genFn1() int32 {
    if (1000 != -430) {
        v12 := []int{-494, 128}
        _ = v12
    } else {
        for fi13 := 0; fi13 < 18; fi13++ {
            if (false || (2 != 2147483648)) {
                var v14 string = "\n\t"
                if (42 != -32768) {
                    v14 = "foo bar"
                    if (false || false) {
                        var v15 uint32 = 560
                        if (-708 == 9223372036854775807) {
                            v16 := []int{}
                            for fi17 := 0; fi17 < 20; fi17++ {
                                var v18 uint8 = 128
                                _ = fi17
                                _ = v18
                            }
                            _ = v16
                        } else {
                            v19 := []int{}
                            var v20 float32 = 0.0
                            var v21 int8 = 96
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                        _ = v15
                    } else {
                        var v22 int16 = 42
                        var v23 int16 = 615
                        var v24 float32 = 1e10
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                } else {
                    v25 := []int{32768, (fi13 + -240), fi13}
                    if len(v25) > 0 {
                        v25[0] = (fi13 >> ((255) & 63))
                    }
                    _ = v25
                }
                _ = v14
            } else {
                if false {
                    var v26 uint = 65535
                    _ = v26
                } else {
                    switch "a" {
                    case "foo bar":
                        var v27 uint64 = 0
                        _ = v27
                    case "\n\t":
                        var v28 uint8 = 127
                        _ = v28
                    default:
                        var v29 int16 = 1000
                        _ = v29
                    }
                }
                v30 := []int{(fi13 % ((fi13) | 1)), (fi13 ^ -131), fi13, (fi13 * (fi13 | fi13))}
                _ = v30
            }
            _ = fi13
        }
        var v31 uint64 = 42
        _ = v31
    }
    if (1 < 85) {
        var v32 uint8 = 42
        _ = v32
    }
    var v33 bool = false
    v34 := []int{}
    if len(v34) > 0 {
        v34[0] = 128
    }
    _ = v33
    _ = v34
    return 1
}

func main() {
    var v1 float64 = 3.14159
    var v2 int = 32768
    var v3 int16 = 42
    v4 := []string{"foo bar"}
    var v5 uint = 255
    fmt.Println(v2, v2, v2)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v4)
    for fi35 := 0; fi35 < 10; fi35++ {
        var v36 float64 = v1
        _ = fi35
        _ = v36
    }
    var v37 int16 = (v3 ^ 881)
    v3 = (^(v37))
    var v38 float32 = 1e10
    var v39 float64 = (v1 - v1)
    fmt.Println(v39, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
