package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint64 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 uint8; F1 int }

func main() {
    var v1 uint64 = 0
    var v2 float64 = 1e10
    var v3 float32 = 1.5
    var v4 int16 = 699
    for fi5 := 0; fi5 < 7; fi5++ {
        v3 -= v3
        if false {
            v6 := GenS1{F0: 1, F1: 258, F2: (v1 << ((4294967295) & 63))}
            v1 = v1
            fmt.Println(fi5)
            _ = v6
        } else {
            v1 -= (^(v1))
        }
        v7 := GenS2{F0: v4, F1: 32767}
        _ = fi5
        _ = v7
    }
    if (false && (-128 != 1)) {
        fmt.Println(v4)
        v8 := GenS3{F0: 255, F1: 165}
        _ = v8
    } else {
        var v9 bool = ((!true) && false)
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        v11 := GenS1{F0: (fi10 | (-(fi10))), F1: 132, F2: (v1 | (^(v1)))}
        v3 -= float32(v2)
        _ = fi10
        _ = v11
    }
    switch 42 {
    case -903:
        fmt.Println(v3, v4)
    case 32768:
        v1 += 18446744073709551615
    case -736:
        v12 := []int{0, 0, 4294967295}
        _ = v12
    default:
        if (941 != 0) {
            v13 := []int{-621, 127}
            var v14 string = "\n\t"
            fmt.Println(v1, v4)
            _ = v13
            _ = v14
        } else {
            v15 := GenS2{F0: (-(v4)), F1: 422}
            var v16 uint16 = 65535
            _ = v15
            _ = v16
        }
    }
    v1 += v1
    v17 := []int{-9223372036854775808}
    v18 := GenS2{F0: v4, F1: 42}
    v19 := GenS1{F0: 4294967295, F1: 2, F2: 1}
    if true {
        for fi20 := 0; fi20 < 5; fi20++ {
            fmt.Println(v2, fi20)
            var v21 int16 = (v4 - (v4 / ((v4) | 1)))
            fmt.Println(v1, v3)
            _ = fi20
            _ = v21
        }
        fmt.Println(v1, v3, v3)
        var v22 float64 = 0.0
        _ = v22
    } else {
        v1 -= v1
        for fi23 := 0; fi23 < 3; fi23++ {
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
            _ = fi23
        }
        fmt.Println(v3, v3)
    }
    v24 := GenS3{F0: 42, F1: 438}
    v25 := GenS1{F0: -772, F1: 2, F2: 18446744073709551615}
    var v26 uint = 993
    fmt.Println(v1, v4, v4)
    v3 += v3
    v3 += v3
    for fi27 := 0; fi27 < 3; fi27++ {
        v28 := GenS1{F0: fi27, F1: (v26 >> ((255) & 63)), F2: (v1 << ((v26) & 63))}
        fmt.Println(v2)
        v29 := GenS4{F0: 215, F1: 127}
        _ = fi27
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
