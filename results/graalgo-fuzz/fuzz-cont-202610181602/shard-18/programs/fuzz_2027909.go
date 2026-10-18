package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int }
type GenS2 struct { F0 int64 }
func genFn0(p0 int, p1 uint32, p2 string) uint32 {
    v7 := []int{127, p0}
    p1 -= 2
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 float32 = (((1.5 - float32(-2.25)) - float32((1e10 + ((1e10 * 0.0) + 0.1)))) * ((float32(((1e10 * 1e10) / -2.25)) + -2.25) * float32(1.5)))
        var v10 uint64 = 196
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p1
}
func genFn1(p0 bool) uint8 {
    v11 := GenS1{F0: "", F1: 18446744073709551615, F2: -9223372036854775808}
    v12 := GenS2{F0: -1}
    for fi13 := 0; fi13 < 12; fi13++ {
        p0 = ("\n\t" != "")
        for fi14 := 0; fi14 < 19; fi14++ {
            if p0 {
                v15 := []int{fi13, fi14, fi14}
                _ = v15
            }
            _ = fi14
        }
        v16 := []int{fi13, (fi13 << ((255) & 63)), fi13}
        _ = fi13
        _ = v16
    }
    v17 := []int{}
    _ = p0
    _ = v11
    _ = v12
    _ = v17
    return 203
}

func main() {
    var v1 int16 = 866
    var v2 uint64 = 656
    v3 := []string{("foo bar" + ""), "foo bar"}
    var v4 uint = 2
    var v5 int16 = (v1 << (((v4 - v4)) & 63))
    var v6 uint = v4
    fmt.Println(v4, v6, v5)
    switch 484 {
    case -550:
        var v18 bool = (-2147483648 < -601)
        _ = v18
    case 196:
        switch ("a" + "\n\t") {
        case "\n\t":
            fmt.Println(v6, v2, v1)
        case "a":
            fmt.Println(v4)
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v2, v4)
    var v19 bool = true
    if len(v3) > 0 {
        v3[0] = "a"
    }
    _ = len(v3)
    _ = len(v3)
    for fi20 := 0; fi20 < 7; fi20++ {
        var v21 int8 = 84
        _ = fi20
        _ = v21
    }
    var v22 int32 = 0
    v6 = (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
