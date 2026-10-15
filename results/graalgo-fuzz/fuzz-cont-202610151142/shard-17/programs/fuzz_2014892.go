package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int8 = 100
    v2 := make(map[uint32]int)
    var v3 int16 = 1
    var v4 int8 = (-(v1))
    var v5 float64 = (0.1 + 3.14159)
    var v6 uint = 2
    v3 -= v3
    var v7 int16 = (v3 & (v3 ^ (v3 % ((42) | 1))))
    if (1.5 <= 1e10) {
        var v8 int8 = (v4 & (v4 * v4))
        for fi9 := 0; fi9 < 1; fi9++ {
            for fi10 := 0; fi10 < 15; fi10++ {
                var v11 float32 = 1.5
                _ = fi10
                _ = v11
            }
            if ("hello" > "") {
                for fi12 := 0; fi12 < 16; fi12++ {
                    fmt.Println(fi9, v1, v6)
                    v13 := []int{}
                    _ = fi12
                    _ = v13
                }
                if (-1 < 0) {
                    var v14 uint = v6
                    v15 := []int{(fi9 * (fi9 >> (((v14 << ((796) & 63))) & 63))), 352}
                    _ = v14
                    _ = v15
                }
            } else {
                v6 = 65535
                var v16 int16 = 1
                _ = v16
            }
            _ = fi9
        }
        for fi17 := 0; fi17 < 20; fi17++ {
            v6 = v6
            _ = len(v2)
            _ = fi17
        }
        _ = v8
    } else {
        v6 = v6
    }
    v18 := GenS1{F0: 4294967295}
    fmt.Println(v7, v1)
    v19 := GenS1{F0: 0}
    v20 := []int{-1, 4294967295, -9223372036854775808}
    v21 := GenS1{F0: 4294967295}
    for fi22 := 0; fi22 < 5; fi22++ {
        v23 := GenS1{F0: 261}
        v1 = (v4 | v1)
        v3 = -1
        _ = fi22
        _ = v23
    }
    var v24 int8 = (^(v4))
    fmt.Println(v1, v7)
    v3 = v3
    var v25 uint8 = 128
    switch -349 {
    case 42:
        switch 9223372036854775807 {
        case 0:
            v5 += ((((1e10 * v5) + 0.0) - v5) / 1.0)
        case 32768:
            v1 -= (v24 << (((v6 ^ (v6 / (((^(v6))) | 1)))) & 63))
        case -2147483648:
            v1 = v24
        default:
            v3 += (v3 & (v7 << ((v6) & 63)))
        }
    case 755:
        v20 = append(v20, 32767)
    }
    v1 = 100
    fmt.Println(v25, v7)
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
