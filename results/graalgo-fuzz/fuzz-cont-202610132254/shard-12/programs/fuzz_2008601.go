package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 bool }
type GenS2 struct { F0 uint; F1 int16; F2 uint }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 int32; F1 uint8 }
func genFn0(p0 uint64, p1 uint16) bool {
    for fi4 := 0; fi4 < 5; fi4++ {
        switch int(p1) {
        case 128:
            p1 = p1
        case 1:
            v5 := []int{}
            _ = v5
        case 9223372036854775807:
            v6 := []int{fi4}
            _ = v6
        }
        p1 += 309
        _ = fi4
    }
    v7 := []int{9223372036854775807}
    if false {
        var v8 uint64 = 0
        v9 := GenS1{F0: p1, F1: (^(p0)), F2: true}
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 65535, F1: 2, F2: true}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return (42 > 0)
}

func main() {
    var v1 bool = false
    var v2 float64 = -2.25
    var v3 bool = ((1 != 960) && v1)
    v11 := GenS2{F0: 256, F1: 32767, F2: 0}
    var v12 int8 = 26
    v13 := GenS2{F0: 682, F1: 198, F2: 42}
    if v3 {
        fmt.Println(v3, v12)
    }
    v3 = v3
    v14 := GenS3{F0: (v12 + (v12 ^ 0)), F1: -1}
    var v15 uint = 4294967295
    v16 := GenS4{F0: -236, F1: 0}
    for fi17 := 0; fi17 < 1; fi17++ {
        var v18 uint32 = 860
        fmt.Println(v1, v2, v12)
        v15 = v15
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 string = "foo bar"
        for fi21 := 0; fi21 < 4; fi21++ {
            fmt.Println(v3, fi21)
            _ = fi21
        }
        _ = fi19
        _ = v20
    }
    if v3 {
        fmt.Println(v3, v1)
    }
    v22 := GenS1{F0: 1, F1: 18446744073709551615, F2: v3}
    fmt.Println(v3, v12)
    v23 := []int{4294967295, 4294967295}
    if ((v15 / ((255) | 1)) > v15) {
        v3 = false
        var v24 bool = v1
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
