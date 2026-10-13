package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint64, p1 uint64, p2 int) int {
    p1 -= p0
    var v5 float32 = (0.1 - (float32((0.0 + ((1.5 / 1.5) + 0.0))) + 1e10))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p2
}
func genFn1(p0 int8, p1 int8, p2 int64) uint8 {
    fmt.Println(p1, p0)
    fmt.Println(p0, p1)
    if false {
        var v6 int16 = 42
        if (3.14159 >= float64((1e10 + 1.5))) {
            p2 += (p2 % ((p2) | 1))
        } else {
            p2 = 0
            p2 = (p2 >> ((331) & 63))
            for fi7 := 0; fi7 < 7; fi7++ {
                switch "hello" {
                case "\n\t":
                    switch "hello" {
                    case "a":
                        v8 := GenS1{F0: 630, F1: -2147483648}
                        _ = v8
                    case "hello":
                        switch "" {
                        case "hello":
                            fmt.Println(p0)
                        case "":
                            v9 := []int{}
                            _ = v9
                        default:
                            p1 += int8(fi7)
                        }
                    case "\n\t":
                        var v10 uint16 = 65535
                        _ = v10
                    default:
                        v11 := GenS1{F0: 2, F1: 42}
                        _ = v11
                    }
                case "hello":
                    v12 := []int{(fi7 - fi7), int(v6), -66}
                    _ = v12
                }
                for fi13 := 0; fi13 < 3; fi13++ {
                    if (-128 >= 127) {
                        p1 += p0
                        fmt.Println(p1, v6, v6)
                        switch "" {
                        case "\n\t":
                            v14 := GenS1{F0: 374, F1: -276}
                            _ = v14
                        }
                    } else {
                        p0 -= (p1 ^ p0)
                        var v15 uint16 = 42
                        var v16 uint32 = 4294967295
                        if true {
                            p1 -= p1
                        }
                        v17 := GenS1{F0: 0, F1: -1}
                        switch 1 {
                        case 128:
                            v18 := []int{32767, (fi7 - 1), (-(fi13)), (fi13 >> ((42) & 63))}
                            _ = v18
                        case 4294967295:
                            var v19 int64 = (p2 + p2)
                            _ = v19
                        default:
                            var v20 uint8 = 255
                            var v21 float64 = 0.0
                            _ = v20
                            _ = v21
                        }
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    p0 = -61
                    _ = fi13
                }
                fmt.Println(p0, v6)
                _ = fi7
            }
        }
        _ = v6
    } else {
        p0 = (p0 / (((p0 ^ -1)) | 1))
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 int64 = -1
    v2 := make(map[uint32]int)
    var v3 bool = (1 != 100)
    var v4 int = 609
    var v22 uint32 = 1
    v2[v22] = v4
    var v23 bool = (68 >= 455)
    v1 = v1
    v1 = 0
    switch v4 {
    case 2147483647:
        fmt.Println(v3, v23, v4)
    case 813:
        if v3 {
            v24 := GenS2{F0: (-(v1))}
            fmt.Println(v3, v3, v23)
            v25 := []int{(v4 - -2147483648)}
            _ = v24
            _ = v25
        }
    case -2147483648:
        for fi26 := 0; fi26 < 17; fi26++ {
            for fi27 := 0; fi27 < 15; fi27++ {
                fmt.Println(v4, v22, v22)
                v1 -= -869
                _ = fi27
            }
            _ = fi26
        }
    }
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        v4 -= int(v1)
    }
    for fi28 := 0; fi28 < 15; fi28++ {
        var v29 string = "hello"
        var v30 int = genFn0(813, 121, (v4 >> ((255) & 63)))
        var v31 string = (v29 + v29)
        _ = fi28
        _ = v29
        _ = v30
        _ = v31
    }
    var v32 uint32 = (v22 * (v22 % ((v22) | 1)))
    v23 = v23
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v32
    _ = v4
}
