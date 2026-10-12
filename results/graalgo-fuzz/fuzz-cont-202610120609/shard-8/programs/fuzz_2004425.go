package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() int8 {
    var v5 float32 = (1e10 + -2.25)
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 string = "foo bar"
        if false {
            var v8 float32 = ((0.0 - (float32(((0.1 * 1.5) + 3.14159)) - -2.25)) * float32(3.14159))
            var v9 int32 = 2147483647
            var v10 uint64 = 472
            _ = v10
            _ = v8
            _ = v9
        } else {
            var v11 float32 = float32((1e10 - 0.0))
            _ = v11
        }
        _ = fi6
        _ = v7
    }
    if (42 < 0) {
        var v12 uint8 = 190
        _ = v12
    }
    _ = v5
    return 42
}

func main() {
    v1 := []uint32{642, 917}
    var v2 int64 = -1
    var v3 float64 = ((((1e10 + float64((0.0 / ((float32(-2.25) / 1e10) + 1e10)))) - -2.25) * 3.14159) - (1e10 / 1e10))
    var v4 int32 = 2147483647
    v3 += v3
    v3 += (-2.25 / ((1.5 / v3) - 1e10))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 -= v4
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 int8 = 127
        fmt.Println(v4, v3)
        v3 = v3
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        v2 += (v2 >> ((810) & 63))
        var v16 bool = false
        var v17 uint16 = 0
        _ = fi15
        _ = v16
        _ = v17
    }
    if ((0.1 + -2.25) >= 3.14159) {
        if true {
            var v18 int16 = 1
            for fi19 := 0; fi19 < 9; fi19++ {
                v18 = (v18 % (((-(v18))) | 1))
                var v20 uint64 = 646
                if true {
                    fmt.Println(v18, v18, v4)
                }
                _ = fi19
                _ = v20
            }
            _ = v18
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if (32767 > 302) {
                v4 = (-(v4))
                v1 = append(v1, 2147483647)
            }
            v3 += v3
        }
        var v21 int32 = (v4 >> ((18446744073709551615) & 63))
        if (0 >= 65535) {
            var v22 uint8 = 127
            for fi23 := 0; fi23 < 20; fi23++ {
                fmt.Println(v4, v4)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v24 uint = 255
                _ = fi23
                _ = v24
            }
            _ = v22
        } else {
            v25 := GenS1{F0: (v2 - (-(v2)))}
            var v26 string = "\n\t"
            _ = v25
            _ = v26
        }
        _ = v21
    } else {
        if (((true || false) && true) && false) {
            _ = len(v1)
            for fi27 := 0; fi27 < 1; fi27++ {
                v28 := GenS1{F0: (v2 & (v2 % ((-902) | 1)))}
                var v29 uint16 = 1
                _ = fi27
                _ = v28
                _ = v29
            }
        }
        switch -322 {
        case 0:
            v4 = v4
        }
        fmt.Println(v3, v3, v2)
    }
    v30 := GenS1{F0: (v2 << ((18446744073709551615) & 63))}
    v2 = v2
    fmt.Println(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
