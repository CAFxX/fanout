package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int; F2 int16 }
type GenS3 struct { F0 uint64; F1 int }
func genFn0() int8 {
    v5 := GenS1{F0: "hello"}
    if false {
        v6 := []int{}
        _ = v6
    }
    v7 := GenS2{F0: -128, F1: 32767, F2: 1000}
    var v8 int64 = -1
    v8 = (^(v8))
    _ = v5
    _ = v7
    _ = v8
    return 123
}
func genFn1() int32 {
    var v9 bool = (false && false)
    v9 = v9
    v9 = (-111 == genFn0())
    _ = v9
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = (!false)
    var v3 uint = 42
    v4 := make(map[int]uint32)
    var v10 float64 = (-2.25 * 3.14159)
    var v11 bool = true
    v11 = ((v3 > v3) || (-2147483648 > -1))
    if v11 {
        delete(v4, 1)
        v3 += (v3 * (v3 + v3))
    }
    v2 = v2
    for fi12 := 0; fi12 < 3; fi12++ {
        if (!(!v11)) {
            v13 := GenS3{F0: 0, F1: (-(fi12))}
            var v14 float64 = (float64(-2.25) * (-2.25 * float64((1.5 * 3.14159))))
            v15 := GenS3{F0: 115, F1: (-(fi12))}
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    switch 76 {
    case -278:
        switch 371 {
        case 127:
            var v16 int16 = -32768
            _ = v16
        case 596:
            v11 = v11
        default:
            var v17 int32 = genFn1()
            _ = v17
        }
    default:
        var v18 uint16 = 32767
        _ = v18
    }
    if (42 >= 32767) {
        for fi19 := 0; fi19 < 19; fi19++ {
            fmt.Println(v11, v10)
            switch (fi19 ^ -9223372036854775808) {
            case 0:
                v3 = 42
            }
            _ = fi19
        }
        if (v2 || (!true)) {
            fmt.Println(v10, v11, v2)
            v20 := []int{1}
            if (!false) {
                delete(v4, 0)
            }
            _ = v20
        }
        v21 := GenS3{F0: 2, F1: -9223372036854775808}
        _ = v21
    } else {
        v11 = (42 > genFn1())
        switch ("a" + "a") {
        case "\n\t":
            v22 := v4[-246]
            _ = v22
        case "":
            v23 := []int{42, -9223372036854775808, 1, -362}
            _ = v23
        }
    }
    var v24 int32 = -902
    v25 := GenS2{F0: -1, F1: -1, F2: -32768}
    v24 = -594
    for fi26 := 0; fi26 < 7; fi26++ {
        v27 := []int{fi26, (fi26 | fi26)}
        var v28 uint64 = 0
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 int32 = 0
    for fi30 := 0; fi30 < 16; fi30++ {
        for fi31 := 0; fi31 < 4; fi31++ {
            fmt.Println(fi31)
            _ = fi31
        }
        v10 -= v10
        v32 := GenS1{F0: ""}
        _ = fi30
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
}
