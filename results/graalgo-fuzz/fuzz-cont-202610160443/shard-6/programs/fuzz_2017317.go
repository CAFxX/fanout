package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint; F1 int64 }
func genFn0(p0 string, p1 int) uint {
    v4 := []int{-1}
    var v5 uint = 18446744073709551615
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (v5 - (v5 | 18446744073709551615))
}

func main() {
    v1 := []int{1}
    v2 := []int{1, 159, 610}
    var v3 uint16 = 540
    v6 := GenS1{F0: 0}
    if len(v2) > 0 {
        v2[0] = 32768
    }
    switch 42 {
    case 112:
        if (false || (!(1.5 > 0.1))) {
            v3 = 42
        }
    case -1:
        v1 = append(v1, 547)
    default:
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS1{F0: 360}
            if false {
                var v9 uint = 301
                if false {
                    _ = len(v1)
                    v10 := GenS2{F0: 381, F1: fi7}
                    _ = v10
                }
                var v11 int = (fi7 >> (((^(v9))) & 63))
                _ = v11
                _ = v9
            }
            _ = fi7
            _ = v8
        }
    }
    if true {
        _ = len(v2)
        v3 -= 42
    } else {
        v12 := GenS3{F0: 4294967295, F1: -145}
        var v13 bool = (!(62 > -1))
        _ = v12
        _ = v13
    }
    var v14 int32 = 2147483647
    v3 = v3
    fmt.Println(v3)
    v15 := GenS3{F0: 0, F1: -1}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v6
}
