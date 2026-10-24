package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint8; F1 string; F2 uint32 }
type GenS3 struct { F0 string; F1 int32; F2 uint16 }
func genFn0(p0 bool) bool {
    if p0 {
        v4 := []int{607, 4294967295}
        _ = v4
    }
    switch -1 {
    case -2147483648:
        v5 := GenS1{F0: 266, F1: 0}
        _ = v5
    case 50:
        fmt.Println(p0)
    default:
        v6 := []int{0, 32768, -490, 487}
        _ = v6
    }
    _ = p0
    return false
}
func genFn1(p0 uint16, p1 int32, p2 uint16) uint {
    p2 = (p2 | p0)
    switch 32767 {
    case 342:
        p2 += p0
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 bool = true
    var v2 bool = false
    var v3 bool = (!((v2 && v2) && v1))
    if (v3 || true) {
        var v7 uint64 = 42
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 uint16 = 32767
            if v2 {
                var v10 uint16 = (v9 - (v9 * (^(v9))))
                var v11 float64 = 0.1
                _ = v10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    v2 = v1
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := GenS2{F0: 2, F1: ((((("a" + "") + "") + "hello") + "a") + "hello"), F2: 710}
        var v14 int8 = 100
        _ = fi12
        _ = v13
        _ = v14
    }
    switch 1 {
    case 457:
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 int8 = 100
            _ = fi15
            _ = v16
        }
    case -1:
        v2 = v3
    case 128:
        var v17 int16 = -32768
        _ = v17
    }
    var v18 bool = v3
    v19 := GenS3{F0: "", F1: 42, F2: 1}
    v20 := []int{166, 9223372036854775807, 127}
    if (!(!v1)) {
        v18 = (!v1)
        v21 := GenS2{F0: 1, F1: "a", F2: 2147483648}
        _ = v21
    }
    if (true && v1) {
        fmt.Println(v2)
        switch "a" {
        case "\n\t":
            v20 = append(v20, 32767)
        }
    }
    v3 = (!true)
    for fi22 := 0; fi22 < 19; fi22++ {
        var v23 uint32 = 2
        _ = fi22
        _ = v23
    }
    v3 = (!(42 > 2147483647))
    switch 42 {
    case 753:
        v2 = true
    case 32768:
        fmt.Println(v3, v2)
    default:
        fmt.Println(v1, v18)
    }
    v1 = true
    for fi24 := 0; fi24 < 20; fi24++ {
        v2 = v2
        for fi25 := 0; fi25 < 9; fi25++ {
            var v26 uint16 = 899
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
