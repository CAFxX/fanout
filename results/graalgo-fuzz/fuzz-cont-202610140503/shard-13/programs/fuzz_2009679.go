package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int8; F1 int8 }
func genFn0(p0 int64, p1 int) string {
    switch (^(p1)) {
    case 795:
        var v7 int32 = 1
        _ = v7
    }
    fmt.Println(p1)
    p0 = (p0 << ((255) & 63))
    p0 = -9223372036854775808
    _ = p0
    _ = p1
    return "\n\t"
}
func genFn1(p0 int64) bool {
    var v8 uint32 = 2147483647
    v9 := GenS1{F0: int8(v8), F1: 1}
    _ = p0
    _ = v8
    _ = v9
    return false
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 0
    var v3 uint = 143
    var v4 int16 = 939
    var v5 bool = (false && (741 != 448))
    var v6 int32 = 220
    fmt.Println(v5)
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 bool = v5
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        for fi13 := 0; fi13 < 4; fi13++ {
            v14 := []int{(fi12 ^ fi13), (fi13 * -912), (fi12 % ((fi12) | 1))}
            v15 := v1[65535]
            v16 := GenS2{F0: 100, F1: 1}
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        for fi17 := 0; fi17 < 11; fi17++ {
            switch "a" {
            case "\n\t":
                if (42 >= uint8(v2)) {
                    v18 := []int{(fi12 & (-(fi17))), fi17}
                    var v19 int16 = (^(v4))
                    v19 = (v19 / (((v19 % ((v19) | 1))) | 1))
                    v3 -= v3
                    _ = v18
                    _ = v19
                }
            case "a":
                if (-128 != 0) {
                    switch ("hello" + "") {
                    case "\n\t":
                        var v20 float64 = ((1.5 / (3.14159 * 1e10)) * (-2.25 * 3.14159))
                        _ = v20
                    case "hello":
                        v21 := GenS1{F0: -128, F1: 215}
                        _ = v21
                    case "":
                        v5 = false
                    }
                    fmt.Println(fi17, v6, v3)
                    v22 := []int{-9223372036854775808, fi12, fi12}
                    _ = v22
                } else {
                    var v23 bool = false
                    _ = v23
                }
            case "":
                v6 = 42
            default:
                fmt.Println(v4)
            }
            v1[2147483648] = ("" + "a")
            v5 = v5
            _ = fi17
        }
        _ = fi12
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        var v25 int64 = 42
        var v26 uint = 873
        _ = fi24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        if false {
            v1[0] = "a"
        } else {
            for fi28 := 0; fi28 < 15; fi28++ {
                v4 = v4
                v4 = (v4 / ((v4) | 1))
                v2 += v2
                _ = fi28
            }
            var v29 string = "hello"
            _ = v29
        }
        _ = fi27
    }
    v6 = (v6 - (v6 << ((1) & 63)))
    var v30 int64 = 33
    v3 -= 65535
    var v31 int64 = v30
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
