package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 uint32 }
func genFn0() int64 {
    var v7 int = 235
    switch -2147483648 {
    case -292:
        v7 += v7
    case 4294967295:
        v8 := []int{v7, (v7 >> ((1) & 63)), (v7 * (v7 >> ((4294967295) & 63))), v7}
        _ = v8
    case -948:
        var v9 int16 = 337
        _ = v9
    default:
        var v10 uint8 = 0
        _ = v10
    }
    if (!true) {
        if false {
            for fi11 := 0; fi11 < 12; fi11++ {
                var v12 int8 = 95
                if false {
                    v13 := []int{fi11, v7}
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 15; fi14++ {
                        v12 = v12
                        _ = fi14
                    }
                    v15 := GenS1{F0: 128, F1: 255, F2: -2147483648}
                    _ = v15
                }
                _ = fi11
                _ = v12
            }
        }
        v7 = (v7 >> ((4294967295) & 63))
    } else {
        var v16 bool = (2 != 65535)
        fmt.Println(v7, v7)
        for fi17 := 0; fi17 < 11; fi17++ {
            v16 = true
            _ = fi17
        }
        _ = v16
    }
    v18 := GenS1{F0: 66, F1: 1, F2: 42}
    var v19 int16 = 1000
    _ = v18
    _ = v19
    _ = v7
    return 1
}

func main() {
    v1 := []int{-9223372036854775808, 508, 1, 128}
    var v2 uint8 = 150
    var v3 string = "hello"
    var v4 uint = 276
    v5 := make(map[uint32]uint32)
    var v6 float32 = (1e10 / -2.25)
    delete(v5, 42)
    v2 += v2
    for fi20 := 0; fi20 < 13; fi20++ {
        switch fi20 {
        case -1:
            fmt.Println(v6, v6)
        case -9223372036854775808:
            v6 = ((v6 + v6) / v6)
        default:
            v4 -= v4
        }
        _ = fi20
    }
    if (v2 != v2) {
        v4 -= v4
        switch -336 {
        case 32768:
            v21 := v5[495]
            _ = v21
        case 127:
            delete(v5, 0)
        case 128:
            v4 = v4
        }
        for fi22 := 0; fi22 < 16; fi22++ {
            var v23 int = (-(fi22))
            v23 = 9223372036854775807
            v4 = 65535
            _ = fi22
            _ = v23
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v24 := GenS1{F0: 139, F1: (v4 * 1), F2: 1}
        for fi25 := 0; fi25 < 3; fi25++ {
            v26 := GenS2{F0: -1}
            _ = fi25
            _ = v26
        }
        _ = v24
    }
    fmt.Println(v3)
    _ = len(v5)
    v3 = ""
    if len(v1) > 0 {
        v1[0] = 539
    }
    for fi27 := 0; fi27 < 11; fi27++ {
        fmt.Println(v2)
        if true {
            for fi28 := 0; fi28 < 10; fi28++ {
                var v29 int8 = -8
                var v30 uint8 = (v2 ^ 1)
                _ = fi28
                _ = v29
                _ = v30
            }
        }
        _ = fi27
    }
    for fi31 := 0; fi31 < 19; fi31++ {
        v2 = 1
        var v32 float64 = 1.5
        v33 := GenS3{F0: 2, F1: 2147483648}
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 int16 = 792
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
