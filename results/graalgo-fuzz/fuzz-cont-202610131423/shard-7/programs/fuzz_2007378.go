package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int8) uint64 {
    if false {
        p0 += -128
    } else {
        for fi4 := 0; fi4 < 10; fi4++ {
            v5 := []int{-394, fi4}
            _ = fi4
            _ = v5
        }
        var v6 int16 = 42
        v7 := []int{748}
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: (989 == 574), F1: 455}
    switch ("hello" + "a") {
    case "hello":
        if ((true && (false || false)) && (float64(-2.25) < 3.14159)) {
            v9 := GenS2{F0: 438}
            var v10 float64 = (1e10 + 0.1)
            _ = v10
            _ = v9
        }
    case "":
        var v11 uint64 = 1
        _ = v11
    case "\n\t":
        for fi12 := 0; fi12 < 18; fi12++ {
            v13 := []int{(^(fi12)), (^(fi12))}
            _ = fi12
            _ = v13
        }
    default:
        v14 := []int{32768, 42, 128}
        _ = v14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 bool = (2147483647 <= -342)
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = v8
    return 637
}

func main() {
    var v1 uint8 = 0
    var v2 uint = 65535
    var v3 bool = true
    v17 := GenS2{F0: 1}
    var v18 string = (("foo bar" + "\n\t") + "foo bar")
    v3 = ((-2.25 >= 3.14159) || (!v3))
    v1 += v1
    for fi19 := 0; fi19 < 17; fi19++ {
        if (!(32767 < -32768)) {
            if ((false && v3) && (42 != 42)) {
                v20 := []int{-140, (fi19 | fi19), 32768}
                _ = v20
            }
            var v21 int = (-(fi19))
            _ = v21
        } else {
            fmt.Println(v1, v3)
            v3 = (0.1 >= ((0.1 + 1e10) + 0.1))
            fmt.Println(v18)
        }
        v22 := []int{fi19, 446, (^(fi19))}
        _ = fi19
        _ = v22
    }
    v23 := []int{0, 1}
    v24 := GenS2{F0: 925}
    var v25 float64 = (1e10 + 3.14159)
    v2 = (v2 / ((v2) | 1))
    for fi26 := 0; fi26 < 10; fi26++ {
        var v27 int32 = -660
        var v28 int16 = -928
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println(v1)
    if len(v23) >= 2 {
        v23 = v23[1:len(v23)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
