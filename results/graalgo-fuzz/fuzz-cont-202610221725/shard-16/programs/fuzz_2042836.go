package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int16 = -814
    var v2 bool = (995 == (-(v1)))
    var v3 float64 = 1e10
    var v4 int8 = -17
    switch "" {
    case "\n\t":
        fmt.Println(v3, v2, v3)
    default:
        v2 = ((!(!false)) && (1 >= -143))
    }
    var v5 bool = v2
    v2 = (475 != 2147483647)
    v2 = ((float32((3.14159 - (v3 / ((1.5 - float64((1e10 + 1.5))) + ((v3 - (v3 / v3)) + v3))))) / 0.1) > 3.14159)
    v6 := []int{-607, 1, 32767, 32768}
    if len(v6) > 0 {
        v6[0] = -1
    }
    if v2 {
        var v7 uint16 = 692
        _ = v7
    }
    var v8 int8 = (v4 | (v4 - 127))
    if (4294967295 != -2147483648) {
        if v5 {
            var v9 float64 = (0.0 / (v3 - (v3 / v3)))
            _ = v9
        } else {
            var v10 int8 = 0
            v1 += v1
            if (0 > 0) {
                fmt.Println(v3, v3, v1)
                for fi11 := 0; fi11 < 12; fi11++ {
                    if v5 {
                        if len(v6) > 0 {
                            v6[0] = fi11
                        }
                    } else {
                        if v2 {
                            v12 := GenS1{F0: 0}
                            var v13 uint64 = 880
                            var v14 uint32 = 2147483647
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                    }
                    _ = fi11
                }
            }
            _ = v10
        }
    } else {
        v15 := GenS1{F0: -1}
        _ = v15
    }
    var v16 int64 = 42
    if false {
        v8 += (v4 - v4)
        var v17 int8 = (v8 / ((v4) | 1))
        var v18 int8 = (v17 << ((42) & 63))
        _ = v17
        _ = v18
    } else {
        if len(v6) > 0 {
            v6[0] = int(v1)
        }
        var v19 float64 = ((v3 / -2.25) * ((((v3 - (1.5 * v3)) + -2.25) / 1.0) - float64(float32(-2.25))))
        _ = v19
    }
    v5 = v2
    fmt.Println(v16, v5)
    fmt.Println(v1)
    var v20 uint = 2
    for fi21 := 0; fi21 < 6; fi21++ {
        _ = len(v6)
        fmt.Println(v4, v4, v16)
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
