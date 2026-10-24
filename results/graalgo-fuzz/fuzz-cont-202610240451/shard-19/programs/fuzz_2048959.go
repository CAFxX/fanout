package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint64; F1 bool; F2 int8 }
func genFn0(p0 int8, p1 uint, p2 int64) bool {
    var v6 uint16 = 689
    switch -1 {
    case -1:
        var v7 int32 = int32(p2)
        _ = v7
    case -2147483648:
        var v8 int = 4294967295
        _ = v8
    default:
        for fi9 := 0; fi9 < 11; fi9++ {
            v10 := GenS1{F0: (p1 % ((65535) | 1)), F1: 171}
            for fi11 := 0; fi11 < 20; fi11++ {
                v12 := []int{(fi9 * fi11)}
                p1 -= (p1 >> (((p1 * p1)) & 63))
                _ = fi11
                _ = v12
            }
            var v13 float64 = 0.0
            _ = fi9
            _ = v10
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return true
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 2147483648
    var v3 float64 = 1.5
    var v4 int = 1
    var v5 int16 = 32767
    var v14 float64 = v3
    if (((0.0 * v14) != v3) || true) {
        var v15 bool = true
        _ = v15
    } else {
        v16 := GenS2{F0: 1, F1: -38}
        _ = v16
    }
    v14 = v14
    if (-1 <= 860) {
        v3 += ((v14 + v3) - (0.0 - v14))
        switch ("hello" + "") {
        case "a":
            var v17 uint16 = 255
            _ = v17
        case "hello":
            for fi18 := 0; fi18 < 7; fi18++ {
                v14 = (0.0 * 1.5)
                _ = fi18
            }
        case "":
            v5 -= v5
        }
    }
    v19 := GenS3{F0: 2, F1: (((v2 - (v2 - 817)) >= v2) && true), F2: -1}
    v5 -= v5
    var v20 int16 = 679
    fmt.Println(v3, v3, v2)
    fmt.Println(v14)
    fmt.Println(v4, v4, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
