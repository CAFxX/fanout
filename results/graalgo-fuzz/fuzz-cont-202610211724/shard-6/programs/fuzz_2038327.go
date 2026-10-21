package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 int16; F2 string }
func genFn0(p0 bool, p1 string) bool {
    v6 := GenS1{F0: 101}
    if true {
        v7 := GenS1{F0: 1}
        var v8 int32 = 1
        var v9 bool = false
        _ = v7
        _ = v8
        _ = v9
    } else {
        var v10 uint32 = 2
        v11 := []int{-862, 151, 32768, 128}
        _ = v10
        _ = v11
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v6
    return false
}

func main() {
    var v1 uint64 = 42
    var v2 int16 = 1000
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]string)
    v5 := []int{}
    v4[0] = "\n\t"
    fmt.Println(v1)
    var v12 uint32 = 0
    v12 = v12
    if (-124 < 658) {
        v13 := GenS1{F0: 238}
        fmt.Println(v12, v2)
        if true {
            var v14 uint64 = 1
            _ = v14
        }
        _ = v13
    } else {
        var v15 int32 = 1
        if false {
            switch 366 {
            case 9223372036854775807:
                for fi16 := 0; fi16 < 12; fi16++ {
                    fmt.Println(v2, v12, v15)
                    _ = fi16
                }
            case 32767:
                v1 = 745
            }
            if true {
                v5 = append(v5, 647)
                v17 := GenS1{F0: 0}
                for fi18 := 0; fi18 < 20; fi18++ {
                    v19 := GenS1{F0: 2}
                    _ = fi18
                    _ = v19
                }
                _ = v17
            }
            v1 -= v1
        }
        if ((128 < 202) && true) {
            for fi20 := 0; fi20 < 15; fi20++ {
                var v21 int = (^(fi20))
                v15 -= v15
                _ = fi20
                _ = v21
            }
            v22 := GenS1{F0: 2}
            _ = v22
        }
        _ = v15
    }
    v1 += 4294967295
    var v23 int64 = 0
    v2 -= 0
    v24 := GenS2{F0: 128}
    v1 -= (v1 - v1)
    v25 := GenS3{F0: 0, F1: (v2 - (v2 * (v2 | (-(v2))))), F2: ("\n\t" + ("hello" + "\n\t"))}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
