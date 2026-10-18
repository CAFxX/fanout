package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
func genFn0() uint16 {
    v4 := []int{-2147483648, 42}
    var v5 int64 = 473
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = v4
    _ = v5
    return 957
}
func genFn1(p0 string, p1 int16, p2 uint64) bool {
    switch "a" {
    case "hello":
        for fi6 := 0; fi6 < 5; fi6++ {
            for fi7 := 0; fi7 < 19; fi7++ {
                var v8 string = p0
                _ = fi7
                _ = v8
            }
            v9 := GenS1{F0: 847}
            _ = fi6
            _ = v9
        }
    case "":
        p1 = p1
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        if false {
            for fi11 := 0; fi11 < 20; fi11++ {
                fmt.Println(p2, p2, p2)
                fmt.Println(fi11)
                var v12 float64 = 1.5
                _ = fi11
                _ = v12
            }
            p0 = ""
            v13 := GenS2{F0: 2147483648}
            _ = v13
        }
        _ = fi10
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        v15 := GenS1{F0: 897}
        var v16 float32 = float32((-2.25 * (-2.25 * ((0.1 + 1.5) + 0.1))))
        v17 := GenS2{F0: 0}
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    return (p1 <= p1)
}

func main() {
    var v1 uint64 = 42
    var v2 int16 = 844
    var v3 uint8 = 255
    v3 += 0
    var v18 uint32 = 65535
    fmt.Println(v3, v3)
    v19 := []int{0, -572, 9223372036854775807}
    var v20 bool = (!false)
    var v21 float32 = 0.1
    var v22 bool = v20
    var v23 int64 = 905
    var v24 float64 = (3.14159 + 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
