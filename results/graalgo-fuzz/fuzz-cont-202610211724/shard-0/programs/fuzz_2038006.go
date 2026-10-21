package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 int; F1 bool; F2 uint16 }
func genFn0(p0 uint8, p1 int16) bool {
    var v5 int32 = 348
    p0 += 42
    v6 := GenS1{F0: ("hello" + "\n\t")}
    v5 += (v5 & v5)
    v7 := GenS2{F0: 2, F1: 1}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return false
}

func main() {
    var v1 uint8 = 42
    var v2 float64 = (((1.5 / float64(1e10)) / ((float64(0.1) - -2.25) - 1.5)) / 1e10)
    var v3 bool = (false || (!false))
    var v4 int16 = 42
    v8 := GenS2{F0: 620, F1: -161}
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS3{F0: -9223372036854775808, F1: false, F2: 916}
        _ = fi9
        _ = v10
    }
    var v11 int = 4294967295
    v3 = (v3 || (9223372036854775807 >= (v11 ^ v11)))
    var v12 int16 = (^(v4))
    v1 = v1
    if v3 {
        for fi13 := 0; fi13 < 7; fi13++ {
            for fi14 := 0; fi14 < 5; fi14++ {
                v15 := []int{}
                var v16 uint = 42
                if (false && (!(!v3))) {
                    var v17 string = ("a" + "a")
                    v3 = (-2147483648 > 1)
                    _ = v17
                } else {
                    var v18 uint8 = 0
                    if len(v15) > 0 {
                        v15[0] = 9223372036854775807
                    }
                    _ = len(v15)
                    _ = v18
                }
                _ = fi14
                _ = v15
                _ = v16
            }
            v19 := GenS2{F0: 2, F1: (-(v11))}
            for fi20 := 0; fi20 < 17; fi20++ {
                if v3 {
                    var v21 float32 = (1.5 / ((0.0 / float32(1e10)) - ((1.5 / 0.1) / (3.14159 + float32((v2 / v2))))))
                    switch -1 {
                    case 130:
                        if v3 {
                            if v3 {
                                v1 += v1
                                var v22 uint64 = 4294967295
                                fmt.Println(v22, v21)
                                _ = v22
                            }
                        } else {
                            v23 := []int{fi20, -584, (^(v11))}
                            v23 = append(v23, -187)
                            var v24 float32 = 0.1
                            _ = v23
                            _ = v24
                        }
                    }
                    _ = v21
                }
                var v25 int8 = 100
                _ = fi20
                _ = v25
            }
            _ = fi13
            _ = v19
        }
    }
    var v26 float32 = (0.1 * 0.0)
    if false {
        v11 = (-(v11))
    } else {
        v27 := []int{v11, (v11 + v11), v11, 0}
        if len(v27) > 0 {
            v27[0] = v11
        }
        switch (v11 & v11) {
        case 2147483647:
            fmt.Println(v11)
        }
        _ = v27
    }
    v28 := []int{v11, -1, 237}
    v29 := GenS2{F0: 18446744073709551615, F1: (v11 & 4294967295)}
    if len(v28) > 0 {
        v28[0] = v11
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        switch "" {
        case "a":
            for fi31 := 0; fi31 < 3; fi31++ {
                var v32 string = "foo bar"
                _ = fi31
                _ = v32
            }
        default:
            var v33 int64 = 1
            _ = v33
        }
        v4 -= (v12 + 348)
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v8
}
