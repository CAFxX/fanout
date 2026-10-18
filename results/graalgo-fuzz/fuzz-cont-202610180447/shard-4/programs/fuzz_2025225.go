package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint64 }
func genFn0(p0 int64, p1 uint, p2 string) int {
    p0 -= 0
    var v7 uint8 = 127
    if false {
        var v8 uint16 = 937
        if false {
            var v9 uint16 = (^(v8))
            fmt.Println(p2)
            var v10 float64 = 1e10
            _ = v10
            _ = v9
        } else {
            fmt.Println(p1)
            var v11 uint16 = v8
            v8 = v8
            _ = v11
        }
        _ = v8
    } else {
        v12 := []int{32767}
        _ = v12
    }
    var v13 int32 = -1
    var v14 int32 = v13
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v7
    return 0
}
func genFn1(p0 int8) uint {
    var v15 int32 = 631
    v16 := GenS1{F0: 1, F1: 20, F2: 2}
    _ = p0
    _ = v15
    _ = v16
    return 297
}

func main() {
    v1 := []string{}
    var v2 uint64 = 2
    v3 := make(map[uint32]string)
    var v4 uint64 = v2
    v5 := []int{449}
    v6 := []uint32{0, 252}
    var v17 float64 = ((-2.25 - (-2.25 + -2.25)) - ((float64(1e10) + (0.1 + -2.25)) * 0.0))
    v17 = ((float64(1.5) + (((0.1 + (-2.25 - ((float64(1.5) * float64(((((3.14159 - (0.1 * 0.0)) - 3.14159) + ((0.1 / 0.1) - (1.5 / 1.0))) * float32(v17)))) - v17))) * v17) - 3.14159)) * ((0.1 / v17) - v17))
    v4 -= 2
    switch "" {
    case "":
        var v18 float32 = 0.0
        _ = v18
    case "hello":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v19 float64 = 0.1
    v20 := GenS1{F0: 65535, F1: 44, F2: v2}
    var v21 int16 = 42
    v21 += (v21 ^ v21)
    switch -1 {
    case -216:
        var v22 uint16 = 871
        _ = v22
    default:
        for fi23 := 0; fi23 < 16; fi23++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi23
        }
    }
    var v24 uint64 = (v4 >> ((2) & 63))
    var v25 int32 = 0
    for fi26 := 0; fi26 < 10; fi26++ {
        for fi27 := 0; fi27 < 1; fi27++ {
            if ((!false) || (("hello" + "\n\t") == "a")) {
                v3[2] = "foo bar"
                switch "\n\t" {
                case "a":
                    v25 = int32(v21)
                case "\n\t":
                    switch "a" {
                    case "\n\t":
                        var v28 int = (fi27 | fi27)
                        _ = v28
                    case "":
                        var v29 uint = 482
                        _ = v29
                    default:
                        _ = len(v3)
                    }
                }
                var v30 bool = (!("" < "\n\t"))
                fmt.Println(fi27, v21, v24)
                _ = len(v3)
                v2 -= 18446744073709551615
                v31 := GenS1{F0: 814, F1: 128, F2: 0}
                var v32 uint = genFn1(int8(v21))
                _ = v30
                _ = v31
                _ = v32
            } else {
                v33 := GenS1{F0: 1, F1: 114, F2: (v4 * v24)}
                if len(v6) > 0 {
                    v6[0] = 2147483647
                }
                v34 := GenS1{F0: 792, F1: 127, F2: 700}
                _ = v33
                _ = v34
            }
            if len(v5) > 0 {
                v5[0] = -1
            }
            v21 += (v21 - -362)
            _ = fi27
        }
        _ = fi26
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v35 uint32 = 1
    var v36 int32 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
