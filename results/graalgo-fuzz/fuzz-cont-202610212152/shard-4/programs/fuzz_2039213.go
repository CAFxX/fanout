package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int16 }
type GenS2 struct { F0 int64; F1 int64 }
func genFn0(p0 int8, p1 int16) uint8 {
    v7 := []int{1, 9223372036854775807, -9223372036854775808}
    if false {
        var v8 uint16 = 901
        var v9 uint64 = 0
        if true {
            var v10 int32 = -1
            _ = v10
        }
        _ = v8
        _ = v9
    }
    var v11 uint = 2
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    return 179
}
func genFn1(p0 uint32) uint64 {
    switch 4294967295 {
    case 32767:
        switch -9223372036854775808 {
        case -2147483648:
            for fi12 := 0; fi12 < 17; fi12++ {
                for fi13 := 0; fi13 < 12; fi13++ {
                    var v14 int32 = 1
                    var v15 string = ""
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                p0 += p0
                for fi16 := 0; fi16 < 6; fi16++ {
                    if ((p0 >= p0) && (!((1 > 42) && true))) {
                        v17 := []int{}
                        _ = v17
                    } else {
                        if false {
                            var v18 uint32 = p0
                            _ = v18
                        } else {
                            p0 = (p0 * p0)
                        }
                    }
                    _ = fi16
                }
                _ = fi12
            }
        default:
            fmt.Println(p0)
        }
    case 283:
        for fi19 := 0; fi19 < 20; fi19++ {
            var v20 float32 = 3.14159
            _ = fi19
            _ = v20
        }
    case 4294967295:
        p0 += p0
    default:
        var v21 uint16 = 2
        _ = v21
    }
    var v22 int32 = -1
    var v23 uint64 = 99
    v24 := GenS1{F0: -90, F1: (!(230 < 37)), F2: 428}
    _ = p0
    _ = v22
    _ = v23
    _ = v24
    return v23
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = -128
    v3 := []int{-2147483648, -9223372036854775808, 32767}
    v4 := []uint32{42, 42, 4294967295, 976}
    var v5 uint64 = 4294967295
    v6 := make(map[string]int)
    var v25 string = ((((("" + ("foo bar" + (("\n\t" + ("hello" + ("foo bar" + "\n\t"))) + ((("" + "\n\t") + "") + "")))) + "foo bar") + "a") + "foo bar") + "\n\t")
    for fi26 := 0; fi26 < 15; fi26++ {
        if ((v2 + v2) < v2) {
            if false {
                v25 = v25
                delete(v1, 42)
            }
            _ = len(v1)
        } else {
            v2 = v2
        }
        switch v25 {
        case "hello":
            var v27 float32 = float32(float64((float32(1.5) + 0.0)))
            _ = v27
        case "\n\t":
            for fi28 := 0; fi28 < 20; fi28++ {
                v29 := GenS1{F0: v2, F1: (true && true), F2: 42}
                v2 += v2
                _ = fi28
                _ = v29
            }
        case "a":
            v30 := GenS1{F0: (^(v2)), F1: (!(128 < 155)), F2: 45}
            _ = v30
        }
        v31 := GenS2{F0: 179, F1: -698}
        _ = fi26
        _ = v31
    }
    fmt.Println(v5, v5)
    v32 := GenS1{F0: 3, F1: (0.1 == 1e10), F2: 118}
    v33 := GenS2{F0: 0, F1: 9223372036854775807}
    v4 = append(v4, 2147483647)
    v5 = v5
    fmt.Println(v5, v25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
