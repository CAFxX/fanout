package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint32 }
func genFn0() uint32 {
    var v5 string = "hello"
    for fi6 := 0; fi6 < 12; fi6++ {
        v5 = (v5 + v5)
        v7 := GenS1{F0: (-(fi6)), F1: -48, F2: 87}
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 uint8 = 1
            var v10 uint = 329
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    fmt.Println(v5)
    var v11 bool = false
    _ = v11
    _ = v5
    return 2147483648
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    var v3 int32 = 2147483647
    var v4 uint32 = 2147483648
    switch 4294967295 {
    case 1:
        if (!(true || false)) {
            v3 -= (v3 ^ v3)
            fmt.Println(v3)
            v3 -= -1
        }
    case 32767:
        _ = len(v2)
    case 32768:
        v3 = (v3 % ((v3) | 1))
    default:
        var v12 int = 775
        _ = v12
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 int8 = 127
        if true {
            if false {
                _ = len(v2)
                v15 := GenS1{F0: fi13, F1: 0, F2: (v4 << ((0) & 63))}
                _ = v15
            }
            v4 += 65535
        } else {
            if false {
                if false {
                    v14 += 100
                    v4 -= genFn0()
                    _ = len(v2)
                }
                fmt.Println(v3)
                v3 += (-(v3))
            }
        }
        _ = fi13
        _ = v14
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        v4 -= v4
        v17 := []int{fi16, fi16, (fi16 << ((255) & 63))}
        _ = len(v17)
        _ = fi16
        _ = v17
    }
    if false {
        switch ("a" + "a") {
        case "hello":
            var v18 uint = 399
            _ = v18
        }
        v3 = v3
        v4 = (v4 | (v4 >> ((42) & 63)))
    } else {
        v3 = v3
    }
    v19 := []int{1, 32767, 4294967295, 210}
    var v20 uint64 = 0
    if (0 != 1) {
        for fi21 := 0; fi21 < 19; fi21++ {
            fmt.Println(v3)
            if false {
                _ = len(v1)
                var v22 uint64 = v20
                _ = v22
            }
            _ = len(v2)
            _ = fi21
        }
        if ((true && false) && true) {
            fmt.Println(v20)
        }
    }
    v3 += (v3 - (^(v3)))
    var v23 float64 = 0.0
    v24 := GenS1{F0: 363, F1: v3, F2: (^(v4))}
    var v25 int = -1
    v23 -= v23
    var v26 float32 = -2.25
    var v27 bool = false
    for fi28 := 0; fi28 < 19; fi28++ {
        var v29 uint32 = (v4 | (v4 >> ((65535) & 63)))
        v25 = 128
        _ = fi28
        _ = v29
    }
    v3 -= v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
