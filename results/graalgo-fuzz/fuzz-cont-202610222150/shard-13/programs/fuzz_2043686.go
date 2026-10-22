package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 uint; F2 int64 }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 int8; F1 int16; F2 int8 }
func genFn0(p0 int8) int8 {
    p0 += (p0 * (p0 % ((p0) | 1)))
    p0 = (p0 % (((p0 << ((2) & 63))) | 1))
    fmt.Println(p0)
    switch -321 {
    case -2147483648:
        v6 := GenS1{F0: 803}
        _ = v6
    case 42:
        v7 := GenS1{F0: 0}
        _ = v7
    default:
        var v8 uint16 = 65535
        _ = v8
    }
    _ = p0
    return (p0 - 114)
}

func main() {
    var v1 uint32 = 65535
    v2 := []int{0, -240}
    var v3 uint16 = 0
    v4 := []int{1, -1, 724, 509}
    var v5 string = "hello"
    v9 := GenS2{F0: 1}
    v10 := GenS3{F0: 59, F1: 2, F2: 42}
    fmt.Println(v5)
    v11 := GenS4{F0: 255, F1: (v3 << ((1) & 63))}
    v12 := GenS5{F0: -128, F1: 838, F2: 127}
    v5 = ((v5 + "\n\t") + (v5 + ((v5 + (v5 + "")) + "\n\t")))
    v3 = v3
    v4 = append(v4, -9223372036854775808)
    var v13 int32 = 2147483647
    v4 = append(v4, 234)
    v3 = v3
    if false {
        if ((v13 - 42) < (v13 + (v13 % (((v13 / ((-1) | 1))) | 1)))) {
            v5 = v5
        } else {
            for fi14 := 0; fi14 < 13; fi14++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v15 uint = 4294967295
                _ = fi14
                _ = v15
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            fmt.Println(v5, v5)
        }
    }
    fmt.Println(v1, v3, v5)
    var v16 uint = 65535
    if len(v4) > 0 {
        v4[0] = -350
    }
    fmt.Println(v13)
    var v17 uint = (v16 + 711)
    for fi18 := 0; fi18 < 16; fi18++ {
        v4 = append(v4, int(v13))
        v5 = ("hello" + "foo bar")
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
