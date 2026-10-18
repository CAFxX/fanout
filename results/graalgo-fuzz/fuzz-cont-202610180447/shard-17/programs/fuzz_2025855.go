package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint; F1 int; F2 int16 }
func genFn0(p0 int8, p1 uint32, p2 uint16) uint64 {
    v4 := GenS1{F0: 1, F1: 123}
    p1 -= 432
    p0 = p0
    for fi5 := 0; fi5 < 17; fi5++ {
        switch (fi5 >> ((255) & 63)) {
        case -955:
            v6 := GenS1{F0: int32(p2), F1: 119}
            _ = v6
        default:
            if (false || ((616 != 4294967295) && false)) {
                if (-126 < 42) {
                    if (((128 <= 127) || (!false)) && (((p1 ^ 65535) == p1) || (p0 == (p0 * p0)))) {
                        fmt.Println(p2)
                    } else {
                        p0 = p0
                        switch fi5 {
                        case 1:
                            var v7 uint64 = 724
                            _ = v7
                        case -884:
                            v8 := []int{}
                            _ = v8
                        case -998:
                            if (4294967295 > 42) {
                                v9 := []int{fi5, fi5, (fi5 << ((0) & 63))}
                                _ = v9
                            }
                        }
                    }
                    var v10 bool = false
                    _ = v10
                }
                p0 += p0
                if false {
                    v11 := GenS1{F0: -1, F1: 2}
                    _ = v11
                } else {
                    p2 = p2
                    v12 := []int{-91, 1, fi5}
                    _ = v12
                }
            } else {
                p0 = p0
            }
        }
        var v13 bool = (true || true)
        fmt.Println(v13, v13)
        _ = fi5
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 18446744073709551615
}

func main() {
    var v1 uint64 = 1
    var v2 float64 = (-2.25 / 1e10)
    var v3 uint32 = 42
    switch 32768 {
    case 652:
        v14 := GenS2{F0: 1, F1: -9223372036854775808, F2: 0}
        _ = v14
    case -886:
        var v15 int16 = 119
        _ = v15
    }
    if (((1e10 + float32(v2)) >= 0.0) || true) {
        var v16 string = ""
        v1 += genFn0(-5, (v3 * 2147483647), 65535)
        var v17 int8 = int8(v3)
        _ = v16
        _ = v17
    }
    v3 = v3
    var v18 uint16 = 42
    v3 += v3
    for fi19 := 0; fi19 < 13; fi19++ {
        switch 625 {
        case 4294967295:
            if true {
                v2 = float64((3.14159 - float32((-2.25 + v2))))
            }
        }
        v20 := []int{}
        _ = fi19
        _ = v20
    }
    v21 := GenS1{F0: 2147483647, F1: 0}
    for fi22 := 0; fi22 < 13; fi22++ {
        if false {
            fmt.Println(fi22, v2)
            v1 -= v1
        }
        v23 := []int{128, fi22}
        _ = fi22
        _ = v23
    }
    var v24 float32 = float32(3.14159)
    for fi25 := 0; fi25 < 10; fi25++ {
        var v26 int = 42
        v27 := GenS1{F0: int32(v1), F1: 255}
        _ = fi25
        _ = v26
        _ = v27
    }
    v28 := []int{32768, -106}
    v28 = append(v28, 127)
    fmt.Println(v24, v18, v3)
    v3 -= (v3 << ((65535) & 63))
    var v29 bool = (-176 != 1)
    v18 -= v18
    v18 = (v18 << ((4294967295) & 63))
    var v30 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
