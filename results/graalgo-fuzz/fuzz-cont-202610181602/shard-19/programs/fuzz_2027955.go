package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int16; F1 int8 }
func genFn0(p0 uint32) bool {
    var v5 int64 = 0
    var v6 uint16 = 255
    _ = p0
    _ = v5
    _ = v6
    return false
}

func main() {
    var v1 uint32 = 2147483648
    var v2 string = ("" + "\n\t")
    v3 := []int{658, 2147483647, 128, 42}
    var v4 uint32 = (v1 * v1)
    v4 = v1
    for fi7 := 0; fi7 < 15; fi7++ {
        v2 = (v2 + "")
        var v8 uint64 = 2
        if true {
            fmt.Println(v4, v4)
        } else {
            switch 439 {
            case 128:
                fmt.Println(v1, v1, v1)
            case -2147483648:
                fmt.Println(v2, v2)
            case 32767:
                _ = len(v3)
            default:
                v9 := GenS1{F0: 778, F1: 42}
                _ = v9
            }
            v8 = v8
            v4 += (v4 >> ((0) & 63))
        }
        _ = fi7
        _ = v8
    }
    if false {
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 float64 = 1.5
            var v12 bool = false
            v3 = append(v3, (fi10 ^ (fi10 / ((4294967295) | 1))))
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 uint = 351
        for fi14 := 0; fi14 < 20; fi14++ {
            for fi15 := 0; fi15 < 2; fi15++ {
                var v16 uint64 = 431
                var v17 float32 = 3.14159
                var v18 uint = (v13 * v13)
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            v19 := GenS1{F0: 2, F1: 32767}
            var v20 int = (fi14 << ((42) & 63))
            _ = fi14
            _ = v19
            _ = v20
        }
        _ = v13
    } else {
        v2 = ("" + ("\n\t" + ""))
        v2 = v2
    }
    v4 += 4294967295
    fmt.Println(v2, v4)
    var v21 uint32 = v1
    if ((!(1 == 9223372036854775807)) || (42 < 2147483647)) {
        if len(v3) > 0 {
            v3[0] = -9223372036854775808
        }
    }
    var v22 uint16 = 183
    for fi23 := 0; fi23 < 6; fi23++ {
        switch 0 {
        case 511:
            var v24 uint = 604
            _ = v24
        case 1:
            if true {
                v25 := GenS2{F0: 42, F1: 981}
                v1 = (v21 + (v4 / ((673) | 1)))
                _ = v25
            } else {
                var v26 float32 = (0.0 * 1.5)
                _ = v26
            }
        }
        switch ("hello" + (((((v2 + v2) + v2) + "foo bar") + "foo bar") + v2)) {
        case "hello":
            if (0.0 < (1e10 - -2.25)) {
                fmt.Println(v22, v21)
                var v27 uint32 = v1
                _ = v27
            }
        case "":
            var v28 bool = (fi23 >= 1)
            _ = v28
        }
        v21 += 421
        _ = fi23
    }
    v29 := GenS3{F0: 1000, F1: -1}
    v2 = v2
    for fi30 := 0; fi30 < 1; fi30++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi30
    }
    for fi31 := 0; fi31 < 12; fi31++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if len(v3) > 0 {
            v3[0] = (fi31 / ((592) | 1))
        }
        v22 = v22
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v4
}
