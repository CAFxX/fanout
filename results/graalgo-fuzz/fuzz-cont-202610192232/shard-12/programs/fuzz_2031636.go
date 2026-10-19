package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 int64; F2 int16 }
func genFn0(p0 int, p1 int8, p2 int) string {
    var v5 int = (p2 & -166)
    for fi6 := 0; fi6 < 3; fi6++ {
        if false {
            v7 := []int{p0}
            var v8 uint8 = 2
            _ = v7
            _ = v8
        } else {
            var v9 int16 = 0
            for fi10 := 0; fi10 < 20; fi10++ {
                p1 += p1
                _ = fi10
            }
            switch "" {
            case "\n\t":
                for fi11 := 0; fi11 < 11; fi11++ {
                    v9 -= v9
                    _ = fi11
                }
            default:
                v12 := []int{fi6}
                _ = v12
            }
            _ = v9
        }
        fmt.Println(p1, v5, p1)
        _ = fi6
    }
    switch -624 {
    case 32768:
        var v13 float64 = (1.5 + -2.25)
        _ = v13
    default:
        p1 -= 42
    }
    v14 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v5
    return "\n\t"
}
func genFn1(p0 bool) int64 {
    fmt.Println(p0)
    v15 := GenS1{F0: 846}
    switch 2147483647 {
    case -806:
        var v16 uint16 = 920
        _ = v16
    case 0:
        var v17 int16 = -1
        _ = v17
    }
    _ = p0
    _ = v15
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = -871
    var v3 float32 = 1.5
    v4 := make(map[uint32]string)
    var v18 uint8 = 127
    var v19 uint32 = 909
    if (!false) {
        if ("\n\t" == ("" + "a")) {
            var v20 float32 = ((-2.25 + float32(1.5)) + v3)
            if false {
                _ = len(v4)
            } else {
                v21 := []int{2147483647}
                _ = v21
            }
            for fi22 := 0; fi22 < 8; fi22++ {
                var v23 uint = 255
                v24 := GenS2{F0: -640}
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = v20
        }
        v2 += v2
    } else {
        v25 := GenS1{F0: 285}
        v26 := GenS3{F0: (v19 << ((4294967295) & 63)), F1: 1, F2: -640}
        _ = v25
        _ = v26
    }
    v3 -= (v3 - (-2.25 - 0.1))
    var v27 uint16 = 65535
    var v28 float64 = 0.0
    var v29 uint64 = 18446744073709551615
    v30 := []int{2147483647, 4294967295}
    for fi31 := 0; fi31 < 14; fi31++ {
        v32 := GenS2{F0: -9223372036854775808}
        v33 := GenS1{F0: 2}
        _ = fi31
        _ = v32
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
