package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint32 }
func genFn0(p0 int8) uint16 {
    var v5 bool = (true && (!false))
    var v6 string = "a"
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 float64 = ((-2.25 - 0.0) + 1.5)
        v9 := []int{}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v6
    return 32767
}
func genFn1(p0 int32) bool {
    fmt.Println(p0)
    p0 = p0
    var v10 float64 = -2.25
    switch "a" {
    case "hello":
        if (!true) {
            v11 := GenS1{F0: -476, F1: 1, F2: 1}
            _ = v11
        } else {
            v12 := GenS1{F0: 856, F1: 40, F2: 648}
            _ = v12
        }
    }
    _ = p0
    _ = v10
    return (!false)
}

func main() {
    v1 := []int{532, 2147483647, -74}
    var v2 string = ""
    var v3 uint16 = 65535
    v4 := make(map[uint32]string)
    v2 = ("\n\t" + "foo bar")
    fmt.Println(v3)
    _ = len(v4)
    for fi13 := 0; fi13 < 4; fi13++ {
        for fi14 := 0; fi14 < 15; fi14++ {
            switch 189 {
            case -1:
                v15 := GenS1{F0: 9223372036854775807, F1: 127, F2: 2}
                _ = v15
            case 0:
                fmt.Println(fi14, fi14, fi14)
            default:
                delete(v4, 2)
            }
            _ = fi14
        }
        for fi16 := 0; fi16 < 4; fi16++ {
            v3 -= v3
            v17 := GenS1{F0: -9223372036854775808, F1: 219, F2: 4294967295}
            v3 += v3
            var v18 uint64 = 0
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi13
    }
    v19 := GenS1{F0: 9223372036854775807, F1: 87, F2: 65535}
    var v20 float32 = ((0.1 / 3.14159) * 0.1)
    fmt.Println(v2)
    _ = len(v1)
    v21 := GenS1{F0: 42, F1: 206, F2: 542}
    v22 := GenS1{F0: -756, F1: 2, F2: 635}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch "a" {
    case "foo bar":
        v23 := v4[0]
        _ = v23
    default:
        var v24 string = (v2 + "a")
        _ = v24
    }
    if true {
        if ((0 >= uint64(v3)) && false) {
            var v25 float64 = -2.25
            for fi26 := 0; fi26 < 1; fi26++ {
                v4[1] = ("\n\t" + (v2 + v2))
                fmt.Println(fi26, v25)
                _ = fi26
            }
            v20 -= v20
            _ = v25
        } else {
            v27 := GenS1{F0: 0, F1: 42, F2: 1}
            var v28 int16 = -765
            _ = v27
            _ = v28
        }
        fmt.Println(v3, v20)
        if false {
            for fi29 := 0; fi29 < 12; fi29++ {
                v20 -= -2.25
                _ = fi29
            }
            var v30 float32 = v20
            _ = v30
        }
    } else {
        var v31 bool = (true && (-608 < -120))
        v2 = v2
        var v32 uint64 = 4294967295
        _ = v31
        _ = v32
    }
    v33 := GenS1{F0: 174, F1: 2, F2: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v33
    _ = v4
}
