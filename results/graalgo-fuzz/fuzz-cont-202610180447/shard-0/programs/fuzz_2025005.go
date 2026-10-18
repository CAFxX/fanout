package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint16; F1 bool; F2 uint64 }
func genFn0() int16 {
    v4 := []int{}
    var v5 int8 = -1
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 uint8, p1 int64) uint8 {
    v6 := GenS1{F0: 18446744073709551615, F1: "hello"}
    p1 -= (p1 << ((168) & 63))
    var v7 uint8 = (p0 >> ((4294967295) & 63))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return p0
}

func main() {
    var v1 bool = true
    var v2 uint64 = 980
    var v3 uint64 = (v2 | v2)
    for fi8 := 0; fi8 < 10; fi8++ {
        v1 = v1
        _ = fi8
    }
    var v9 int32 = int32(v2)
    var v10 float64 = 0.1
    var v11 float64 = v10
    switch ("hello" + (("hello" + "foo bar") + "")) {
    case "":
        if v1 {
            v12 := GenS2{F0: 288, F1: (1 != 1), F2: v3}
            _ = v12
        } else {
            for fi13 := 0; fi13 < 2; fi13++ {
                v14 := []int{127}
                v9 = v9
                _ = fi13
                _ = v14
            }
            if v1 {
                var v15 uint8 = 211
                _ = v15
            } else {
                v16 := GenS1{F0: 2, F1: ("" + "foo bar")}
                _ = v16
            }
        }
    default:
        v17 := []int{985}
        _ = v17
    }
    switch ("\n\t" + ("\n\t" + "a")) {
    case "\n\t":
        var v18 int = 339
        _ = v18
    case "a":
        v19 := GenS1{F0: v2, F1: "a"}
        _ = v19
    default:
        v20 := []int{-573}
        _ = v20
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        v2 -= (v3 / (((v3 + 127)) | 1))
        _ = fi21
    }
    v2 -= 2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v9
}
