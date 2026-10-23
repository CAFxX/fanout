package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 uint; F2 string }
func genFn0(p0 int) int16 {
    v5 := []int{-1, -651}
    var v6 float32 = ((1e10 - ((1e10 * -2.25) + float32(-2.25))) - -2.25)
    _ = p0
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 int, p1 string) uint32 {
    fmt.Println(p1)
    p1 = p1
    switch p0 {
    case 0:
        for fi7 := 0; fi7 < 9; fi7++ {
            var v8 int8 = 52
            for fi9 := 0; fi9 < 3; fi9++ {
                p0 = (fi7 >> ((18446744073709551615) & 63))
                v10 := []int{(p0 % ((579) | 1))}
                _ = fi9
                _ = v10
            }
            var v11 float32 = (0.1 - 3.14159)
            _ = fi7
            _ = v11
            _ = v8
        }
    case 1:
        var v12 float64 = (-2.25 + 0.0)
        _ = v12
    default:
        var v13 uint64 = 1
        _ = v13
    }
    var v14 float64 = (3.14159 - ((1e10 - -2.25) / 1e10))
    if false {
        v15 := []int{-770, (p0 | p0), -1, p0}
        p1 = p1
        p0 = p0
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v14
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 string = ("\n\t" + "hello")
    v3 := make(map[string]uint32)
    var v4 int16 = 0
    delete(v1, 2147483648)
    for fi16 := 0; fi16 < 19; fi16++ {
        v1[1] = fi16
        var v17 uint = 693
        var v18 float64 = (1e10 / 1.5)
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println(v2)
    v19 := GenS1{F0: 42, F1: 255}
    var v20 float32 = 0.0
    v21 := []int{}
    switch "\n\t" {
    case "a":
        switch -1 {
        case 2147483647:
            v22 := v3["\n\t"]
            _ = v22
        case -2147483648:
            if len(v21) >= 2 {
                v21 = v21[1:len(v21)]
            }
        }
    case "":
        fmt.Println(v4, v2)
    case "foo bar":
        if true {
            if (1 <= -323) {
                var v23 bool = false
                if len(v21) > 0 {
                    v21[0] = -1
                }
                _ = v23
            }
            var v24 float32 = float32(0.1)
            _ = v24
        } else {
            switch 127 {
            case 0:
                v25 := GenS2{F0: 32768}
                _ = v25
            case 42:
                for fi26 := 0; fi26 < 2; fi26++ {
                    v27 := v3[("hello" + (v2 + v2))]
                    _ = fi26
                    _ = v27
                }
            default:
                fmt.Println(v4, v4, v2)
            }
        }
    default:
        if len(v21) >= 2 {
            v21 = v21[1:len(v21)]
        }
    }
    v20 = 1.5
    var v28 uint = 255
    fmt.Println(v2, v20)
    v20 -= v20
    v20 = v20
    v29 := GenS2{F0: -171}
    v4 = (v4 & genFn0(1))
    v30 := GenS3{F0: -1, F1: 42, F2: v2}
    var v31 int32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
