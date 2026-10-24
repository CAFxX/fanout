package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint8, p1 int32) int {
    p1 = (p1 + (p1 >> ((436) & 63)))
    p1 -= 42
    _ = p0
    _ = p1
    return 140
}

func main() {
    var v1 uint = 250
    var v2 float32 = (1.5 - 1.5)
    var v3 int8 = 42
    var v4 float64 = (float64(v2) * -2.25)
    fmt.Println(v4, v1, v1)
    var v5 int64 = 1
    v6 := GenS1{F0: -121, F1: 1000}
    v5 += v5
    switch "\n\t" {
    case "foo bar":
        var v7 uint32 = 413
        _ = v7
    case "a":
        v8 := []int{}
        _ = v8
    case "\n\t":
        v1 += v1
    default:
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 uint8 = 1
            _ = fi9
            _ = v10
        }
    }
    var v11 string = "a"
    if false {
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 uint32 = 994
            var v14 uint32 = (^(v13))
            switch 2147483647 {
            case 1:
                v15 := GenS2{F0: v5}
                _ = v15
            case 32768:
                v16 := []int{genFn0(2, 0), genFn0(1, 42), genFn0(80, -908), (-(fi12))}
                _ = v16
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        var v17 int64 = 42
        _ = v17
    } else {
        if false {
            for fi18 := 0; fi18 < 5; fi18++ {
                fmt.Println(v5, v4, fi18)
                v19 := []int{fi18, (fi18 + fi18)}
                v20 := GenS1{F0: -1, F1: 1}
                _ = fi18
                _ = v19
                _ = v20
            }
            switch v11 {
            case "hello":
                v21 := GenS1{F0: 816, F1: 32767}
                _ = v21
            case "\n\t":
                v22 := []int{366}
                _ = v22
            case "a":
                v2 = 1.5
            }
        }
    }
    var v23 uint = v1
    switch ("\n\t" + "a") {
    case "a":
        v24 := []int{128, 4294967295, -497, genFn0(255, -2147483648)}
        _ = v24
    default:
        var v25 int8 = (v3 + 97)
        _ = v25
    }
    v5 = 9223372036854775807
    v26 := []int{127, 1, genFn0(60, 1), 0}
    var v27 uint64 = 42
    v2 += v2
    v5 -= (v5 / ((v5) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
