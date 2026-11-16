package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 uint64 }
func genFn0(p0 uint64) int {
    var v4 string = "hello"
    v5 := GenS1{F0: false}
    _ = p0
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 uint, p1 int64) int8 {
    var v6 string = "\n\t"
    v7 := GenS1{F0: (0 <= 976)}
    v6 = v6
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint = 536
    var v3 int8 = 127
    if true {
        v3 = v3
    } else {
        v8 := GenS1{F0: (!(998 >= genFn0(2)))}
        fmt.Println(v3, v3)
        for fi9 := 0; fi9 < 5; fi9++ {
            v10 := GenS2{F0: 523, F1: 18446744073709551615}
            if false {
                v2 = v2
                var v11 string = "\n\t"
                _ = v11
            } else {
                var v12 float32 = -2.25
                _ = v12
            }
            for fi13 := 0; fi13 < 12; fi13++ {
                var v14 uint = (v2 | v2)
                var v15 float64 = ((float64(1.5) * 3.14159) - 3.14159)
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    v1 += (v1 / ((4294967295) | 1))
    v1 = v1
    fmt.Println(v3, v2)
    v1 -= (v1 / ((101) | 1))
    if true {
        fmt.Println(v1, v1, v2)
        v16 := []int{}
        var v17 uint16 = uint16(v2)
        _ = v16
        _ = v17
    } else {
        v2 += (^(v2))
        if true {
            switch "" {
            case "hello":
                var v18 int8 = v3
                _ = v18
            default:
                for fi19 := 0; fi19 < 8; fi19++ {
                    v3 = genFn1((v2 * v2), -431)
                    v20 := []int{genFn0((^(v1))), fi19}
                    fmt.Println(v1, v1, v1)
                    _ = fi19
                    _ = v20
                }
            }
            v3 = v3
        } else {
            var v21 bool = ((!(!true)) && false)
            _ = v21
        }
    }
    var v22 int8 = -1
    for fi23 := 0; fi23 < 20; fi23++ {
        var v24 int32 = 0
        _ = fi23
        _ = v24
    }
    v25 := []int{}
    v2 = 2
    v26 := GenS2{F0: 0, F1: (v1 - v1)}
    var v27 uint32 = 2147483647
    var v28 int64 = 163
    v1 += 227
    switch int(v28) {
    case 1:
        if len(v25) > 0 {
            v25[0] = int(v28)
        }
    case -689:
        switch 2147483647 {
        case 32767:
            v28 = v28
        case 42:
            for fi29 := 0; fi29 < 16; fi29++ {
                v22 -= int8(v27)
                _ = fi29
            }
        case -9223372036854775808:
            v3 = (v3 >> ((uint(v1)) & 63))
        default:
            var v30 int64 = (v28 << ((v2) & 63))
            _ = v30
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
