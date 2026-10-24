package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 int32 }
func genFn0(p0 int32, p1 uint8, p2 uint) bool {
    p1 = 77
    if false {
        v4 := []int{}
        _ = v4
    }
    switch -9223372036854775808 {
    case 32768:
        fmt.Println(p0)
    case 507:
        for fi5 := 0; fi5 < 5; fi5++ {
            switch -2147483648 {
            case -1:
                var v6 int16 = 32767
                _ = v6
            default:
                p1 -= (p1 >> ((p2) & 63))
            }
            if true {
                p2 -= (^(p2))
                v7 := []int{(fi5 % (((-(fi5))) | 1))}
                _ = v7
            }
            _ = fi5
        }
    }
    switch 32767 {
    case 2147483647:
        if (!false) {
            for fi8 := 0; fi8 < 17; fi8++ {
                for fi9 := 0; fi9 < 6; fi9++ {
                    var v10 float64 = (float64(1e10) - -2.25)
                    var v11 bool = (("foo bar" + "a") != "hello")
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                _ = fi8
            }
        } else {
            var v12 bool = true
            fmt.Println(p0, p2)
            _ = v12
        }
    }
    var v13 uint64 = uint64(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return false
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = -734
    var v3 uint64 = 18446744073709551615
    var v14 int16 = 42
    var v15 uint32 = 2147483647
    if false {
        fmt.Println(v2)
        var v16 float64 = float64(1e10)
        fmt.Println(v16, v14)
        _ = v16
    }
    v3 += 2
    var v17 uint64 = (v3 & (^(v3)))
    var v18 int32 = -2147483648
    v19 := GenS1{F0: 550, F1: 1}
    var v20 uint32 = v15
    for fi21 := 0; fi21 < 17; fi21++ {
        for fi22 := 0; fi22 < 14; fi22++ {
            v23 := GenS2{F0: 18446744073709551615}
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    for fi24 := 0; fi24 < 17; fi24++ {
        v25 := GenS3{F0: 230, F1: (v17 & 42)}
        v26 := GenS4{F0: v2}
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
