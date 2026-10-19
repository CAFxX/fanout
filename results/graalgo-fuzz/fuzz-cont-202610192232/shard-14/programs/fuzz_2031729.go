package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 uint; F2 uint }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 int }
func genFn0(p0 int32) uint8 {
    v5 := GenS1{F0: -1, F1: 1, F2: int64(p0)}
    var v6 string = ("\n\t" + (("\n\t" + "a") + ((("hello" + (("" + ("foo bar" + "")) + "")) + "a") + "\n\t")))
    var v7 uint8 = 128
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return v7
}
func genFn1(p0 int64, p1 bool, p2 int64) int8 {
    p0 = -1
    v8 := GenS2{F0: 95}
    for fi9 := 0; fi9 < 4; fi9++ {
        v10 := GenS2{F0: fi9}
        for fi11 := 0; fi11 < 18; fi11++ {
            v12 := []int{}
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 11; fi13++ {
            v14 := GenS3{F0: -545, F1: 255, F2: 65535}
            _ = fi13
            _ = v14
        }
        _ = fi9
        _ = v10
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return -100
}

func main() {
    var v1 uint8 = 42
    var v2 uint32 = 705
    var v3 int = 4294967295
    var v4 int8 = -1
    v15 := []int{(v3 ^ (v3 % ((1) | 1))), (^(v3)), v3, v3}
    if len(v15) > 0 {
        v15[0] = v3
    }
    var v16 uint = 42
    var v17 uint16 = 1
    for fi18 := 0; fi18 < 17; fi18++ {
        fmt.Println(v17, fi18)
        _ = fi18
    }
    var v19 int32 = 42
    v20 := GenS4{F0: v2, F1: (v19 * (v19 ^ 42))}
    var v21 int32 = v19
    fmt.Println(v4)
    fmt.Println(v16)
    var v22 int32 = (v19 * v19)
    v3 = v3
    _ = len(v15)
    for fi23 := 0; fi23 < 4; fi23++ {
        if (!true) {
            v16 = 42
            if (true || (!(genFn1(9223372036854775807, true, -1) <= v4))) {
                v15 = append(v15, fi23)
                var v24 float32 = 1.5
                _ = v24
            } else {
                for fi25 := 0; fi25 < 4; fi25++ {
                    var v26 int32 = (v21 | v21)
                    v22 += v19
                    _ = fi25
                    _ = v26
                }
            }
        }
        _ = fi23
    }
    v22 = -135
    if ((!(!(!((uint16(v16) > v17) && true)))) && true) {
        v1 = (v1 / ((255) | 1))
        v27 := GenS5{F0: (v3 | v3)}
        _ = v27
    } else {
        v28 := GenS2{F0: v3}
        var v29 uint8 = (^(v1))
        _ = v28
        _ = v29
    }
    var v30 uint32 = 186
    v31 := GenS3{F0: (v21 << (((v16 / ((v16) | 1))) & 63)), F1: (v16 * v16), F2: (v16 / ((v16) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
