package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint64, p1 int64) int64 {
    fmt.Println(p0)
    var v6 bool = true
    if v6 {
        v7 := GenS1{F0: 112}
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 int8 = 0
            _ = fi8
            _ = v9
        }
        v10 := []int{800}
        _ = v10
        _ = v7
    } else {
        p0 -= p0
        v11 := []int{127, 4294967295, 128}
        _ = v11
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = v6
    return p1
}
func genFn1() int {
    v12 := []int{}
    var v13 bool = (!true)
    for fi14 := 0; fi14 < 17; fi14++ {
        switch ("" + "\n\t") {
        case "a":
            var v15 uint16 = 0
            _ = v15
        case "hello":
            v13 = (42 <= 746)
        case "\n\t":
            v16 := GenS2{F0: 255}
            _ = v16
        }
        _ = fi14
    }
    var v17 float32 = 0.0
    _ = v12
    _ = v13
    _ = v17
    return 32767
}

func main() {
    var v1 int = 0
    var v2 int64 = 543
    var v3 uint64 = 612
    v4 := []int{(^(v1)), (^(v1)), v1}
    v5 := make(map[uint32]bool)
    var v18 int64 = (v2 >> ((900) & 63))
    v18 -= genFn0(4294967295, genFn0(808, genFn0(18446744073709551615, v18)))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        v2 += v18
        if false {
            if false {
                fmt.Println(v2)
            }
            v20 := v5[65535]
            _ = v20
        } else {
            var v21 int = fi19
            _ = v21
        }
        v22 := GenS2{F0: 0}
        _ = fi19
        _ = v22
    }
    var v23 uint32 = 65535
    var v24 int = (v1 << ((4294967295) & 63))
    switch "\n\t" {
    case "":
        v1 -= genFn1()
    case "a":
        v25 := GenS1{F0: (v24 - (v1 >> ((801) & 63)))}
        _ = v25
    case "\n\t":
        fmt.Println(v18, v2)
    default:
        v26 := v5[v23]
        _ = v26
    }
    fmt.Println(v1, v18, v2)
    v24 = v1
    var v27 uint32 = (v23 / ((v23) | 1))
    v5[v23] = false
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
