package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int, p1 int64) int32 {
    var v5 float64 = ((0.0 + 1.5) / (1e10 * 0.1))
    var v6 float32 = 0.0
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 825
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]int)
    var v3 uint16 = 0
    var v4 float64 = 1e10
    var v7 int8 = 42
    v3 = v3
    v3 = (v3 | 42)
    delete(v1, 1)
    switch -1 {
    case -190:
        v3 += v3
    case -763:
        var v8 bool = false
        _ = v8
    }
    var v9 int = 4294967295
    var v10 uint16 = v3
    if (false && (3.14159 > 1e10)) {
        var v11 bool = false
        var v12 int = (v9 ^ (v9 * v9))
        _ = v11
        _ = v12
    }
    var v13 float64 = (1.5 / 1.5)
    if false {
        var v14 uint16 = v10
        for fi15 := 0; fi15 < 19; fi15++ {
            v1[880] = (-(fi15))
            _ = fi15
        }
        _ = v14
    }
    v16 := []int{v9, v9, v9, (-(v9))}
    for fi17 := 0; fi17 < 3; fi17++ {
        if (!true) {
            v18 := GenS1{F0: 230}
            _ = v18
        } else {
            for fi19 := 0; fi19 < 6; fi19++ {
                var v20 float64 = 3.14159
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                if (false || true) {
                    delete(v2, -922)
                    v7 += v7
                } else {
                    var v21 float64 = (v20 / (((float64((float32((v20 * (1e10 * (v13 * v4)))) - 0.0)) - (v4 + ((v20 / 1.0) / v4))) + float64(float32(v13))) + (v20 + v20)))
                    _ = v21
                }
                if false {
                    switch v9 {
                    case 32768:
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                    case -463:
                        v10 = (v3 + v10)
                    case 0:
                        var v22 bool = true
                        _ = v22
                    }
                    var v23 uint64 = 0
                    _ = v23
                } else {
                    delete(v2, 128)
                    v20 = v13
                    v20 = v13
                    if false {
                        v1[1] = fi17
                        fmt.Println(v3, v10)
                    }
                    var v24 string = "\n\t"
                    var v25 uint16 = 923
                    v1[2147483647] = fi17
                    v10 += v3
                    _ = v24
                    _ = v25
                }
                _ = fi19
                _ = v20
            }
            var v26 float32 = 3.14159
            _ = v26
        }
        var v27 uint16 = (v10 % ((65535) | 1))
        _ = fi17
        _ = v27
    }
    var v28 int64 = 699
    if (1 >= 0) {
        v13 = v4
        v7 = (-(v7))
        v4 -= (3.14159 * (-2.25 * 3.14159))
    } else {
        var v29 bool = true
        _ = v29
    }
    fmt.Println(v7, v4)
    var v30 uint64 = uint64(v3)
    var v31 uint16 = (v10 % ((v3) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v7
    _ = v9
}
