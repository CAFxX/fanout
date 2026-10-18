package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 string; F1 string }
func genFn0(p0 bool) uint64 {
    p0 = (0 > 1)
    if p0 {
        var v4 float64 = (3.14159 / float64(((0.0 * (((1.5 + 0.0) + (((0.1 + 1e10) / -2.25) / (0.1 - 0.0))) + 3.14159)) + 0.1)))
        switch 0 {
        case -2147483648:
            v5 := []int{346, 32768}
            _ = v5
        case 868:
            fmt.Println(v4, p0)
        case 4294967295:
            switch ("hello" + ("\n\t" + "\n\t")) {
            case "foo bar":
                v4 = 3.14159
            default:
                fmt.Println(v4, p0)
            }
        default:
            v6 := []int{}
            _ = v6
        }
        var v7 string = "\n\t"
        _ = v4
        _ = v7
    } else {
        v8 := GenS1{F0: 1, F1: 2147483647, F2: 0}
        v9 := []int{1, 42, 32767, -679}
        for fi10 := 0; fi10 < 5; fi10++ {
            v9 = append(v9, 32767)
            _ = fi10
        }
        _ = v8
        _ = v9
    }
    p0 = false
    _ = p0
    return 4294967295
}

func main() {
    var v1 int16 = 568
    var v2 uint16 = 1
    var v3 string = "\n\t"
    for fi11 := 0; fi11 < 17; fi11++ {
        v12 := []int{}
        _ = fi11
        _ = v12
    }
    v13 := GenS2{F0: 1, F1: 42}
    fmt.Println(v1)
    v14 := []int{32767, -538}
    fmt.Println(v1, v2, v2)
    var v15 int64 = 0
    v2 = v2
    for fi16 := 0; fi16 < 13; fi16++ {
        if true {
            for fi17 := 0; fi17 < 9; fi17++ {
                v18 := GenS1{F0: 810, F1: 0, F2: genFn0(false)}
                _ = fi17
                _ = v18
            }
        }
        var v19 uint32 = 0
        _ = fi16
        _ = v19
    }
    fmt.Println(v1, v1)
    fmt.Println(v3, v15, v2)
    v20 := GenS3{F0: "hello", F1: v3}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
}
