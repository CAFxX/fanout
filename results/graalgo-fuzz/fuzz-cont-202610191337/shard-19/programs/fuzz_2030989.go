package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int }

func main() {
    var v1 string = ("foo bar" + ("" + "hello"))
    var v2 int8 = 42
    var v3 uint32 = 2147483648
    var v4 float32 = 1.5
    var v5 uint16 = 32767
    v6 := GenS1{F0: (true || true), F1: 1, F2: -9223372036854775808}
    v4 += float32(float64((((v4 / v4) + (((-2.25 + v4) / float32(((float64(float32(((1.5 / ((((((1.5 + (((0.0 + 3.14159) - ((float64(((v4 * (0.0 / float32(0.1))) - v4)) / 1.5) + float64(v4))) + 0.0)) - 0.1) / -2.25) / (-2.25 / (((1.5 * (1e10 * 0.0)) - 3.14159) / -2.25))) * 0.1) - 1.5)) - 0.1))) * 1e10) * 1.5))) - v4)) / (float32(3.14159) - ((0.0 / 3.14159) + 0.0)))))
    v1 = ((((v1 + "\n\t") + v1) + "a") + v1)
    fmt.Println(v3, v1)
    if false {
        var v7 int16 = -1
        v8 := []int{}
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 int64 = -34
            if true {
                v11 := []int{(fi9 >> ((18446744073709551615) & 63)), 32767, (fi9 - 2147483647), fi9}
                _ = v11
            } else {
                if ((v5 << ((4294967295) & 63)) > v5) {
                    v12 := []int{4294967295, (fi9 ^ (^(fi9))), (fi9 - 520), fi9}
                    switch (fi9 ^ fi9) {
                    case 878:
                        _ = len(v12)
                    }
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 13; fi13++ {
                        v14 := GenS1{F0: (1e10 == -2.25), F1: 22, F2: -1}
                        if (42 >= 1) {
                            v15 := []int{}
                            var v16 bool = (false || ((true || (-1 >= 42)) && false))
                            _ = v15
                            _ = v16
                        } else {
                            v2 += v2
                        }
                        _ = fi13
                        _ = v14
                    }
                    var v17 uint64 = 1
                    _ = v17
                }
            }
            _ = fi9
            _ = v10
        }
        var v18 uint64 = 1
        _ = v18
    }
    if true {
        v2 -= (v2 - 1)
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        for fi20 := 0; fi20 < 11; fi20++ {
            v21 := []int{fi20, (fi19 / (((fi19 ^ -1)) | 1))}
            v2 = (v2 << ((1) & 63))
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    if (true || (-1 >= int16(v5))) {
        v4 -= 1.5
    }
    if (!((1.5 * 0.0) != ((1.5 - 0.0) * 1e10))) {
        v3 += v3
        fmt.Println(v4)
        v2 = (v2 & (v2 - v2))
    } else {
        v2 -= (v2 >> ((18446744073709551615) & 63))
        for fi22 := 0; fi22 < 15; fi22++ {
            v1 = v1
            _ = fi22
        }
        var v23 uint8 = 1
        _ = v23
    }
    v24 := []int{32768}
    if (true || (1 <= 0)) {
        fmt.Println(v3, v4)
        var v25 int8 = v2
        v2 = (-(v2))
        _ = v25
    } else {
        fmt.Println(v4, v5)
        var v26 float64 = 0.0
        v24 = append(v24, -437)
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
