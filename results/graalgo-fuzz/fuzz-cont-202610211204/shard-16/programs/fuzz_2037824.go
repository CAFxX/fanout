package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }

func main() {
    var v1 int32 = 42
    var v2 int32 = 0
    v3 := make(map[uint32]int)
    var v4 float32 = 0.0
    var v5 float32 = ((v4 * ((v4 - 3.14159) - 0.0)) * 0.1)
    v5 = v5
    for fi6 := 0; fi6 < 20; fi6++ {
        fmt.Println(fi6, v4, v4)
        v7 := GenS1{F0: (!true), F1: -32768}
        fmt.Println(v2)
        _ = fi6
        _ = v7
    }
    var v8 int = 127
    var v9 int64 = -1
    v2 += v2
    if (32767 < 255) {
        fmt.Println(v5)
        if true {
            for fi10 := 0; fi10 < 4; fi10++ {
                switch "\n\t" {
                case "\n\t":
                    v4 = 1.5
                default:
                    v1 += v1
                }
                _ = fi10
            }
            v8 = (v8 ^ (v8 % ((4294967295) | 1)))
        }
        for fi11 := 0; fi11 < 8; fi11++ {
            v1 = (v1 + v2)
            var v12 uint16 = 42
            _ = fi11
            _ = v12
        }
    } else {
        if (!(-2.25 == v4)) {
            v5 -= v5
            v13 := []int{v8}
            var v14 int8 = -1
            _ = v13
            _ = v14
        }
        v8 = -2147483648
    }
    v15 := GenS1{F0: false, F1: 42}
    for fi16 := 0; fi16 < 7; fi16++ {
        if ((!false) && false) {
            var v17 int8 = 42
            if (true && (true && (!(!(1 < 1000))))) {
                var v18 uint8 = 135
                _ = len(v3)
                if false {
                    for fi19 := 0; fi19 < 15; fi19++ {
                        var v20 int64 = (v9 + (v9 % (((^(v9))) | 1)))
                        var v21 uint16 = uint16(v18)
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    v22 := []int{405, -930}
                    _ = v22
                }
                _ = v18
            } else {
                v23 := []int{(-(v8)), (fi16 & (fi16 - 0)), (v8 ^ (v8 + fi16)), v8}
                v24 := GenS1{F0: ((!false) && ((fi16 * fi16) <= fi16)), F1: -7}
                _ = v23
                _ = v24
            }
            _ = v17
        }
        _ = fi16
    }
    v1 += 150
    var v25 float32 = 1.5
    var v26 uint64 = 18446744073709551615
    for fi27 := 0; fi27 < 7; fi27++ {
        fmt.Println(v2, v1)
        delete(v3, 42)
        _ = fi27
    }
    var v28 float32 = (v25 / (float32(3.14159) / v4))
    for fi29 := 0; fi29 < 10; fi29++ {
        var v30 int16 = 1
        v26 = v26
        for fi31 := 0; fi31 < 16; fi31++ {
            var v32 int32 = (v1 / ((0) | 1))
            _ = fi31
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
