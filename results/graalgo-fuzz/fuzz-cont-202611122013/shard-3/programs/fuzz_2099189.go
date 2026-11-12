package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint8 }
type GenS2 struct { F0 string }
func genFn0(p0 uint64) uint64 {
    if false {
        var v7 int16 = -32768
        var v8 int16 = (v7 * (v7 & -1))
        for fi9 := 0; fi9 < 20; fi9++ {
            switch ("hello" + "foo bar") {
            case "hello":
                v8 -= (v8 & (v8 + v7))
            case "foo bar":
                for fi10 := 0; fi10 < 3; fi10++ {
                    if true {
                        switch "" {
                        case "":
                            var v11 int32 = -2147483648
                            _ = v11
                        case "hello":
                            v12 := []int{}
                            _ = v12
                        case "a":
                            fmt.Println(p0, v7, v8)
                        }
                    } else {
                        v13 := GenS1{F0: v8, F1: -1, F2: 0}
                        _ = v13
                    }
                    _ = fi10
                }
            case "":
                v8 = v7
            }
            _ = fi9
        }
        _ = v7
        _ = v8
    }
    p0 += (p0 | p0)
    v14 := []int{42, 127, -2147483648, 127}
    p0 = p0
    _ = p0
    _ = v14
    return p0
}
func genFn1(p0 uint16, p1 int8, p2 uint) int {
    p0 = (p0 >> ((255) & 63))
    var v15 int16 = 683
    v16 := GenS2{F0: "\n\t"}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return -9223372036854775808
}

func main() {
    var v1 uint16 = 0
    v2 := make(map[int]uint32)
    var v3 int32 = -1
    var v4 int64 = -9223372036854775808
    v5 := make(map[int]string)
    v6 := make(map[string]uint32)
    v4 += v4
    v1 -= 17
    var v17 int64 = v4
    var v18 string = "\n\t"
    delete(v5, -444)
    for fi19 := 0; fi19 < 14; fi19++ {
        fmt.Println(v17, v1, v18)
        var v20 bool = true
        _ = fi19
        _ = v20
    }
    delete(v6, "a")
    if (!(!(false && (32768 < genFn1((v1 << ((746) & 63)), -1, 2))))) {
        for fi21 := 0; fi21 < 2; fi21++ {
            switch (-(fi21)) {
            case 32768:
                v18 = ((v18 + "") + (v18 + ""))
            case 850:
                var v22 float32 = 0.0
                _ = v22
            default:
                v1 = (^(v1))
            }
            v23 := []int{fi21, (-(fi21))}
            _ = fi21
            _ = v23
        }
        v24 := GenS1{F0: 638, F1: v3, F2: 176}
        _ = v24
    } else {
        v25 := []int{}
        v4 = (v4 % (((-(v17))) | 1))
        _ = v25
    }
    v17 = -456
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
