package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32; F1 int32 }
func genFn0() uint64 {
    var v7 uint64 = 42
    v7 = (v7 + (v7 ^ v7))
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            fmt.Println(fi9, fi8)
            var v10 uint16 = 32767
            var v11 string = ("hello" + "hello")
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    v12 := GenS1{F0: true, F1: 128}
    var v13 uint = 18446744073709551615
    _ = v12
    _ = v13
    _ = v7
    return v7
}
func genFn1() uint32 {
    v14 := GenS2{F0: 255}
    var v15 float64 = 3.14159
    fmt.Println(v15)
    _ = v14
    _ = v15
    return 1
}

func main() {
    var v1 bool = false
    var v2 int32 = -1
    var v3 int = 42
    var v4 int64 = -1
    var v5 int8 = 42
    var v6 float32 = ((((1.5 / 1.0) / 0.1) * 1e10) + float32(((-2.25 / 1.5) / 3.14159)))
    var v16 int = int(v5)
    switch v3 {
    case 0:
        switch (v3 / (((-(v16))) | 1)) {
        case 983:
            v17 := GenS1{F0: true, F1: 39}
            _ = v17
        default:
            v5 = (v5 >> ((852) & 63))
        }
    default:
        for fi18 := 0; fi18 < 10; fi18++ {
            v19 := []int{fi18, (v16 - fi18)}
            if (v6 == v6) {
                if (0 <= v5) {
                    for fi20 := 0; fi20 < 4; fi20++ {
                        fmt.Println(v16)
                        v19 = append(v19, 21)
                        v1 = true
                        v21 := GenS3{F0: 2147483648}
                        _ = fi20
                        _ = v21
                    }
                    v1 = v1
                    for fi22 := 0; fi22 < 11; fi22++ {
                        var v23 int16 = 976
                        v24 := GenS2{F0: 0}
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                }
            } else {
                v16 = (v3 >> ((255) & 63))
            }
            v16 += (v16 / ((-1) | 1))
            _ = fi18
            _ = v19
        }
    }
    fmt.Println(v3, v6)
    for fi25 := 0; fi25 < 11; fi25++ {
        v5 = (v5 | v5)
        for fi26 := 0; fi26 < 18; fi26++ {
            v3 = (v3 >> ((0) & 63))
            var v27 uint16 = 32767
            _ = fi26
            _ = v27
        }
        for fi28 := 0; fi28 < 9; fi28++ {
            v29 := GenS4{F0: 876}
            v1 = false
            _ = fi28
            _ = v29
        }
        _ = fi25
    }
    for fi30 := 0; fi30 < 4; fi30++ {
        v31 := []int{382, fi30, (-(v3)), fi30}
        if len(v31) > 0 {
            v31[0] = (v3 / ((-1) | 1))
        }
        v32 := GenS5{F0: 2147483647, F1: (v2 & v2)}
        _ = fi30
        _ = v31
        _ = v32
    }
    var v33 float64 = (1e10 - -2.25)
    v16 = (^(v3))
    v1 = v1
    var v34 string = "foo bar"
    fmt.Println(v4)
    v35 := []int{(v3 + (v3 * v3)), (v16 * v16)}
    if v1 {
        var v36 uint8 = 56
        var v37 string = ((v34 + "") + (v34 + v34))
        var v38 string = v34
        _ = v36
        _ = v37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
