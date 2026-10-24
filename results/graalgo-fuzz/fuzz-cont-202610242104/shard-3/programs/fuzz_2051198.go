package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint16, p1 int16) uint32 {
    if false {
        if false {
            v5 := []int{-1, -2147483648, -651}
            if len(v5) > 0 {
                v5[0] = 32767
            }
            switch 32768 {
            case 0:
                _ = len(v5)
            default:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            _ = v5
        }
        for fi6 := 0; fi6 < 14; fi6++ {
            p1 = -1
            if true {
                v7 := GenS1{F0: 675}
                _ = v7
            }
            for fi8 := 0; fi8 < 16; fi8++ {
                if true {
                    v9 := GenS1{F0: 9223372036854775807}
                    if true {
                        p0 += p0
                        fmt.Println(p0)
                    } else {
                        var v10 float32 = 0.1
                        v10 = ((-2.25 - 1.5) + v10)
                        p0 = (p0 ^ (p0 << ((2) & 63)))
                        _ = v10
                    }
                    _ = v9
                }
                fmt.Println(p0, p0, p1)
                _ = fi8
            }
            _ = fi6
        }
        if false {
            v11 := []int{42, -483, 1, -589}
            _ = v11
        } else {
            v12 := []int{4294967295, 4294967295}
            v12 = append(v12, 408)
            _ = v12
        }
    }
    var v13 float64 = 3.14159
    v14 := []int{420, -2147483648, -9223372036854775808}
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 65535
}
func genFn1() uint8 {
    var v15 int = 4294967295
    v16 := []int{v15, (^(v15)), v15, (v15 / ((v15) | 1))}
    var v17 string = ""
    v15 = (v15 * v15)
    _ = v15
    _ = v16
    _ = v17
    return 2
}

func main() {
    var v1 int32 = -678
    var v2 uint16 = 2
    var v3 float32 = 1.5
    var v4 float64 = 3.14159
    v18 := []int{-1, 127, 903, -9223372036854775808}
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 uint = 18446744073709551615
        if len(v18) > 0 {
            v18[0] = fi19
        }
        _ = fi19
        _ = v20
    }
    switch 2147483647 {
    case 42:
        v3 += (-2.25 - 1e10)
    case 4294967295:
        var v21 bool = true
        _ = v21
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        v1 = v1
        _ = fi22
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        v4 += ((v4 + v4) * (v4 / (v4 + v4)))
        v1 += v1
        _ = fi23
    }
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    var v24 uint = 18446744073709551615
    if (false || (false || true)) {
        v2 -= v2
        fmt.Println(v1, v2, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
