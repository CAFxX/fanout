package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 uint32 }
func genFn0(p0 int) uint16 {
    var v5 uint16 = 42
    var v6 uint = 1
    _ = p0
    _ = v5
    _ = v6
    return 939
}

func main() {
    var v1 uint32 = 0
    v2 := []uint32{v1, v1, 2147483648}
    var v3 uint32 = v1
    var v4 uint64 = 4294967295
    v3 -= 65535
    v1 = (v3 >> ((0) & 63))
    var v7 string = ("\n\t" + "a")
    switch -372 {
    case 1:
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := GenS1{F0: v3, F1: v7}
            if false {
                v2 = append(v2, 4294967295)
                for fi10 := 0; fi10 < 15; fi10++ {
                    var v11 uint64 = (v4 - 0)
                    var v12 bool = false
                    var v13 uint32 = (^(v3))
                    var v14 uint32 = (v1 >> ((463) & 63))
                    if true {
                        if v12 {
                            if (0.0 == 0.1) {
                                var v15 uint = 65535
                                _ = v15
                            }
                        }
                        v4 += v4
                    }
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                    _ = v14
                }
                var v16 uint64 = 2
                _ = v16
            }
            _ = fi8
            _ = v9
        }
    default:
        v4 = 18446744073709551615
    }
    var v17 int64 = -9223372036854775808
    switch 1 {
    case 0:
        var v18 string = "hello"
        _ = v18
    case -423:
        for fi19 := 0; fi19 < 20; fi19++ {
            if (-1 == -1) {
                var v20 uint16 = 207
                var v21 string = ("\n\t" + ("" + v7))
                var v22 float64 = 1.5
                _ = v20
                _ = v21
                _ = v22
            }
            _ = fi19
        }
    default:
        v2 = append(v2, v1)
    }
    v23 := GenS2{F0: 221}
    fmt.Println(v1)
    var v24 uint = 0
    var v25 uint = (v24 % ((v24) | 1))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 = (^(v1))
    v2 = append(v2, 42)
    if false {
        fmt.Println(v1, v25)
        var v26 int = 4294967295
        _ = v26
    } else {
        for fi27 := 0; fi27 < 20; fi27++ {
            if (((v17 + v17) != 9223372036854775807) || (((3.14159 - 1.5) >= (0.1 / 1.0)) || (v4 == (v4 % ((1) | 1))))) {
                var v28 int64 = (v17 / ((v17) | 1))
                var v29 int8 = 0
                _ = v28
                _ = v29
            } else {
                v2 = append(v2, (v3 - v1))
                _ = len(v2)
                if len(v2) > 0 {
                    v2[0] = (v1 ^ (v1 | v3))
                }
            }
            if len(v2) > 0 {
                v2[0] = (v1 | 65535)
            }
            _ = fi27
        }
        fmt.Println(v24, v7)
        for fi30 := 0; fi30 < 8; fi30++ {
            v31 := GenS3{F0: -364, F1: (v3 & (v3 >> ((v25) & 63)))}
            fmt.Println(v1, v1, v25)
            _ = fi30
            _ = v31
        }
    }
    fmt.Println(v7, v7)
    v7 = (v7 + v7)
    if len(v2) > 0 {
        v2[0] = 0
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v7
}
